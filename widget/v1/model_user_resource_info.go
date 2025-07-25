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

package widget

/*
Widget API

API for enabling the user to make payment from merchant's platform with redirecting to DANA's platform.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.


import (
	"encoding/json"
	"bytes"
	"fmt"
	utils "github.com/dana-id/dana-go/utils"
)

// checks if the UserResourceInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UserResourceInfo{}

// UserResourceInfo struct for UserResourceInfo
type UserResourceInfo struct {
	// Type of user resource
	ResourceType string `json:"resourceType"`
	// Resource value
	Value string `json:"value"`
}

type _UserResourceInfo UserResourceInfo

// NewUserResourceInfo instantiates a new UserResourceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResourceInfo(resourceType string, value string) *UserResourceInfo {
	this := UserResourceInfo{}
	this.ResourceType = resourceType
	this.Value = value
	return &this
}

// NewUserResourceInfoWithDefaults instantiates a new UserResourceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResourceInfoWithDefaults() *UserResourceInfo {
	this := UserResourceInfo{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *UserResourceInfo) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *UserResourceInfo) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *UserResourceInfo) SetResourceType(v string) {
	o.ResourceType = v
}

// GetValue returns the Value field value
func (o *UserResourceInfo) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UserResourceInfo) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UserResourceInfo) SetValue(v string) {
	o.Value = v
}

func (o UserResourceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResourceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *UserResourceInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserResourceInfo := _UserResourceInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserResourceInfo)

	if err != nil {
		return err
	}

	*o = UserResourceInfo(varUserResourceInfo)

	return err
}

type NullableUserResourceInfo struct {
	value *UserResourceInfo
	isSet bool
}

func (v NullableUserResourceInfo) Get() *UserResourceInfo {
	return v.value
}

func (v *NullableUserResourceInfo) Set(val *UserResourceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResourceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResourceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResourceInfo(val *UserResourceInfo) *NullableUserResourceInfo {
	return &NullableUserResourceInfo{value: val, isSet: true}
}

func (v NullableUserResourceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResourceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


