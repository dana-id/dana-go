// Copyright 2025 PT Espay Debit Indonesia Koe
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// FlexibleUnmarshal attempts to unmarshal JSON into a struct, handling missing fields gracefully
func FlexibleUnmarshal(data []byte, v interface{}) error {
	// First, unmarshal into a generic map
	var rawData map[string]interface{}
	if err := json.Unmarshal(data, &rawData); err != nil {
		return err
	}

	// Use reflection to populate the struct fields
	return PopulateStructFromMap(v, rawData)
}

// PopulateStructFromMap uses reflection to populate struct fields from a map
func PopulateStructFromMap(v interface{}, data map[string]interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to struct")
	}

	structValue := rv.Elem()
	structType := structValue.Type()

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldType := structType.Field(i)

		// Skip unexported fields
		if !field.CanSet() {
			continue
		}

		// Get JSON tag name
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// Parse JSON tag (handle omitempty, etc.)
		jsonName := strings.Split(jsonTag, ",")[0]
		if jsonName == "" {
			jsonName = fieldType.Name
		}

		// Check if the field exists in the data
		if value, exists := data[jsonName]; exists && value != nil {
			if err := setFieldValue(field, value, fieldType.Type); err != nil {
				fmt.Printf("Error setting field %s: %v\n", jsonName, err)
				continue
			}
		}
	}

	return nil
}

// setFieldValue sets a struct field value from an interface{} value
func setFieldValue(field reflect.Value, value interface{}, fieldType reflect.Type) error {
	switch fieldType.Kind() {
	case reflect.String:
		if str, ok := value.(string); ok {
			field.SetString(str)
		}
	case reflect.Ptr:
		if fieldType.Elem().Kind() == reflect.String {
			if str, ok := value.(string); ok {
				field.Set(reflect.ValueOf(&str))
			}
		} else {
			// Handle pointer to struct
			if value != nil {
				newValue := reflect.New(fieldType.Elem())
				if mapValue, ok := value.(map[string]interface{}); ok {
					if err := PopulateStructFromMap(newValue.Interface(), mapValue); err == nil {
						field.Set(newValue)
					}
				}
			}
		}
	case reflect.Struct:
		if mapValue, ok := value.(map[string]interface{}); ok {
			newValue := reflect.New(fieldType).Elem()
			if err := PopulateStructFromMap(newValue.Addr().Interface(), mapValue); err == nil {
				field.Set(newValue)
			}
		}
	case reflect.Slice:
		if sliceValue, ok := value.([]interface{}); ok {
			elemType := fieldType.Elem()
			newSlice := reflect.MakeSlice(fieldType, len(sliceValue), len(sliceValue))

			for i, item := range sliceValue {
				elem := newSlice.Index(i)
				if elemType.Kind() == reflect.Struct {
					if mapItem, ok := item.(map[string]interface{}); ok {
						newElem := reflect.New(elemType)
						if err := PopulateStructFromMap(newElem.Interface(), mapItem); err == nil {
							elem.Set(newElem.Elem())
						}
					}
				} else {
					setFieldValue(elem, item, elemType)
				}
			}
			field.Set(newSlice)
		}
	}

	return nil
}

func processNestedJSONFields(jsonStr string) string {
	normalizedStr := strings.ReplaceAll(jsonStr, `\\"`, `\"`)
	re := regexp.MustCompile(`"(\w+)":"(\{.*?\})"`)

	return re.ReplaceAllStringFunc(normalizedStr, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) == 3 {
			fieldName := parts[1]
			jsonValue := parts[2]

			escapedValue := strings.ReplaceAll(jsonValue, `"`, `\"`)

			return fmt.Sprintf(`"%s":"%s"`, fieldName, escapedValue)
		}
		return match
	})
}

func EnsureMinifiedJSON(jsonStr string) (string, error) {
	normalizedStr := strings.ReplaceAll(jsonStr, `\\"`, `\"`)

	if isJSONMinified(normalizedStr) {
		return normalizedStr, nil
	}
	processedStr := processNestedJSONFields(normalizedStr)

	return minifyJSON(processedStr)
}

// isJSONMinified performs a quick heuristic check to see if JSON is already minified
func isJSONMinified(jsonStr string) bool {
	indicators := []string{
		": ",
		", ",
		"{ ",
		"[ ",
		"\n",
		"\t",
		"\r",
	}

	for _, indicator := range indicators {
		if strings.Contains(jsonStr, indicator) {
			return false
		}
	}

	return true
}

// MinifyJSON compacts a JSON string by removing unnecessary whitespace.
func minifyJSON(jsonStr string) (string, error) {
	var obj interface{}
	if err := json.Unmarshal([]byte(jsonStr), &obj); err != nil {
		return "", fmt.Errorf("MinifyJSON: failed to unmarshal JSON: %w", err)
	}

	minifiedBytes, err := json.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("MinifyJSON: failed to marshal JSON for minification: %w", err)
	}
	return string(minifiedBytes), nil
}
