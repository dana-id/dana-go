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

package merchant_management

/*
Merchant Management API

API for merchant management operations in DANA

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.


import (
	"encoding/json"
	"bytes"
	"fmt"
	utils "github.com/dana-id/dana-go/utils"
)

// checks if the CreateShopResponseResponseBody type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateShopResponseResponseBody{}

// CreateShopResponseResponseBody struct for CreateShopResponseResponseBody
type CreateShopResponseResponseBody struct {
	ResultInfo ResultInfo `json:"resultInfo"`
	// The shop ID that was created. Present when resultCodeId is 00000000
	ShopId *string `json:"shopId,omitempty"`
}

type _CreateShopResponseResponseBody CreateShopResponseResponseBody

// NewCreateShopResponseResponseBody instantiates a new CreateShopResponseResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShopResponseResponseBody(resultInfo ResultInfo) *CreateShopResponseResponseBody {
	this := CreateShopResponseResponseBody{}
	this.ResultInfo = resultInfo
	return &this
}

// NewCreateShopResponseResponseBodyWithDefaults instantiates a new CreateShopResponseResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShopResponseResponseBodyWithDefaults() *CreateShopResponseResponseBody {
	this := CreateShopResponseResponseBody{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *CreateShopResponseResponseBody) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *CreateShopResponseResponseBody) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *CreateShopResponseResponseBody) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *CreateShopResponseResponseBody) GetShopId() string {
	if o == nil || utils.IsNil(o.ShopId) {
		var ret string
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopResponseResponseBody) GetShopIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *CreateShopResponseResponseBody) HasShopId() bool {
	if o != nil && !utils.IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given string and assigns it to the ShopId field.
func (o *CreateShopResponseResponseBody) SetShopId(v string) {
	o.ShopId = &v
}

func (o CreateShopResponseResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateShopResponseResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resultInfo"] = o.ResultInfo
	if !utils.IsNil(o.ShopId) {
		toSerialize["shopId"] = o.ShopId
	}
	return toSerialize, nil
}

func (o *CreateShopResponseResponseBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resultInfo",
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

	varCreateShopResponseResponseBody := _CreateShopResponseResponseBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateShopResponseResponseBody)

	if err != nil {
		return err
	}

	*o = CreateShopResponseResponseBody(varCreateShopResponseResponseBody)

	return err
}

type NullableCreateShopResponseResponseBody struct {
	value *CreateShopResponseResponseBody
	isSet bool
}

func (v NullableCreateShopResponseResponseBody) Get() *CreateShopResponseResponseBody {
	return v.value
}

func (v *NullableCreateShopResponseResponseBody) Set(val *CreateShopResponseResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShopResponseResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShopResponseResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShopResponseResponseBody(val *CreateShopResponseResponseBody) *NullableCreateShopResponseResponseBody {
	return &NullableCreateShopResponseResponseBody{value: val, isSet: true}
}

func (v NullableCreateShopResponseResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShopResponseResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


