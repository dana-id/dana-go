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

// checks if the QueryMerchantResourceResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QueryMerchantResourceResponseResponse{}

// QueryMerchantResourceResponseResponse struct for QueryMerchantResourceResponseResponse
type QueryMerchantResourceResponseResponse struct {
	Head QueryMerchantResourceResponseResponseHead `json:"head"`
	Body QueryMerchantResourceResponseResponseBody `json:"body"`
}

type _QueryMerchantResourceResponseResponse QueryMerchantResourceResponseResponse

// NewQueryMerchantResourceResponseResponse instantiates a new QueryMerchantResourceResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMerchantResourceResponseResponse(head QueryMerchantResourceResponseResponseHead, body QueryMerchantResourceResponseResponseBody) *QueryMerchantResourceResponseResponse {
	this := QueryMerchantResourceResponseResponse{}
	this.Head = head
	this.Body = body
	return &this
}

// NewQueryMerchantResourceResponseResponseWithDefaults instantiates a new QueryMerchantResourceResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMerchantResourceResponseResponseWithDefaults() *QueryMerchantResourceResponseResponse {
	this := QueryMerchantResourceResponseResponse{}
	return &this
}

// GetHead returns the Head field value
func (o *QueryMerchantResourceResponseResponse) GetHead() QueryMerchantResourceResponseResponseHead {
	if o == nil {
		var ret QueryMerchantResourceResponseResponseHead
		return ret
	}

	return o.Head
}

// GetHeadOk returns a tuple with the Head field value
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceResponseResponse) GetHeadOk() (*QueryMerchantResourceResponseResponseHead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Head, true
}

// SetHead sets field value
func (o *QueryMerchantResourceResponseResponse) SetHead(v QueryMerchantResourceResponseResponseHead) {
	o.Head = v
}

// GetBody returns the Body field value
func (o *QueryMerchantResourceResponseResponse) GetBody() QueryMerchantResourceResponseResponseBody {
	if o == nil {
		var ret QueryMerchantResourceResponseResponseBody
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceResponseResponse) GetBodyOk() (*QueryMerchantResourceResponseResponseBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *QueryMerchantResourceResponseResponse) SetBody(v QueryMerchantResourceResponseResponseBody) {
	o.Body = v
}

func (o QueryMerchantResourceResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMerchantResourceResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["head"] = o.Head
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

func (o *QueryMerchantResourceResponseResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"head",
		"body",
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

	varQueryMerchantResourceResponseResponse := _QueryMerchantResourceResponseResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryMerchantResourceResponseResponse)

	if err != nil {
		return err
	}

	*o = QueryMerchantResourceResponseResponse(varQueryMerchantResourceResponseResponse)

	return err
}

type NullableQueryMerchantResourceResponseResponse struct {
	value *QueryMerchantResourceResponseResponse
	isSet bool
}

func (v NullableQueryMerchantResourceResponseResponse) Get() *QueryMerchantResourceResponseResponse {
	return v.value
}

func (v *NullableQueryMerchantResourceResponseResponse) Set(val *QueryMerchantResourceResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMerchantResourceResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMerchantResourceResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMerchantResourceResponseResponse(val *QueryMerchantResourceResponseResponse) *NullableQueryMerchantResourceResponseResponse {
	return &NullableQueryMerchantResourceResponseResponse{value: val, isSet: true}
}

func (v NullableQueryMerchantResourceResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMerchantResourceResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


