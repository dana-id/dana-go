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

package payment_gateway

/*
Payment Gateway API

API for doing operations in DANA Payment Gateway (Gapura)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.


import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
	utils "github.com/dana-id/dana-go/utils"
)

// CreateOrderRequest - This schema is a oneOf type with the following possible variants:<br> - CreateOrderByApiRequest<br> - CreateOrderByRedirectRequest<br> 
type CreateOrderRequest struct {
	CreateOrderByApiRequest *CreateOrderByApiRequest
	CreateOrderByRedirectRequest *CreateOrderByRedirectRequest
}

// CreateOrderByApiRequestAsCreateOrderRequest is a convenience function that returns CreateOrderByApiRequest wrapped in CreateOrderRequest
func CreateOrderByApiRequestAsCreateOrderRequest(v *CreateOrderByApiRequest) CreateOrderRequest {
	return CreateOrderRequest{
		CreateOrderByApiRequest: v,
	}
}

// CreateOrderByRedirectRequestAsCreateOrderRequest is a convenience function that returns CreateOrderByRedirectRequest wrapped in CreateOrderRequest
func CreateOrderByRedirectRequestAsCreateOrderRequest(v *CreateOrderByRedirectRequest) CreateOrderRequest {
	return CreateOrderRequest{
		CreateOrderByRedirectRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateOrderRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateOrderByApiRequest
	err = utils.NewStrictDecoder(data).Decode(&dst.CreateOrderByApiRequest)
	if err == nil {
		jsonCreateOrderByApiRequest, _ := json.Marshal(dst.CreateOrderByApiRequest)
		if string(jsonCreateOrderByApiRequest) == "{}" { // empty struct
			dst.CreateOrderByApiRequest = nil
		} else {
			if err = validator.Validate(dst.CreateOrderByApiRequest); err != nil {
				dst.CreateOrderByApiRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.CreateOrderByApiRequest = nil
	}

	// try to unmarshal data into CreateOrderByRedirectRequest
	err = utils.NewStrictDecoder(data).Decode(&dst.CreateOrderByRedirectRequest)
	if err == nil {
		jsonCreateOrderByRedirectRequest, _ := json.Marshal(dst.CreateOrderByRedirectRequest)
		if string(jsonCreateOrderByRedirectRequest) == "{}" { // empty struct
			dst.CreateOrderByRedirectRequest = nil
		} else {
			if err = validator.Validate(dst.CreateOrderByRedirectRequest); err != nil {
				dst.CreateOrderByRedirectRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.CreateOrderByRedirectRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateOrderByApiRequest = nil
		dst.CreateOrderByRedirectRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateOrderRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateOrderRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateOrderRequest) MarshalJSON() ([]byte, error) {
	if src.CreateOrderByApiRequest != nil {
		return json.Marshal(&src.CreateOrderByApiRequest)
	}

	if src.CreateOrderByRedirectRequest != nil {
		return json.Marshal(&src.CreateOrderByRedirectRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateOrderRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CreateOrderByApiRequest != nil {
		return obj.CreateOrderByApiRequest
	}

	if obj.CreateOrderByRedirectRequest != nil {
		return obj.CreateOrderByRedirectRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CreateOrderRequest) GetActualInstanceValue() (interface{}) {
	if obj.CreateOrderByApiRequest != nil {
		return *obj.CreateOrderByApiRequest
	}

	if obj.CreateOrderByRedirectRequest != nil {
		return *obj.CreateOrderByRedirectRequest
	}

	// all schemas are nil
	return nil
}

type NullableCreateOrderRequest struct {
	value *CreateOrderRequest
	isSet bool
}

func (v NullableCreateOrderRequest) Get() *CreateOrderRequest {
	return v.value
}

func (v *NullableCreateOrderRequest) Set(val *CreateOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrderRequest(val *CreateOrderRequest) *NullableCreateOrderRequest {
	return &NullableCreateOrderRequest{value: val, isSet: true}
}

func (v NullableCreateOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


