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

package disbursement

/*
Disbursement API

API for doing disbursement operations in DANA, including DANA account inquiry, transfer to DANA, and transfer to bank disbursement

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.


import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
	utils "github.com/dana-id/dana-go/utils"
)

// checks if the QueryMerchantResourceResponseResponseHead type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QueryMerchantResourceResponseResponseHead{}

// QueryMerchantResourceResponseResponseHead struct for QueryMerchantResourceResponseResponseHead
type QueryMerchantResourceResponseResponseHead struct {
	// API version
	Version string `json:"version"`
	// API interface
	Function string `json:"function"`
	// Client ID provided by DANA, used to identify partner and application system
	ClientId string `json:"clientId"`
	// DateTime with timezone, which follows the ISO-8601 standard. Refer to RFC 3339 Section 5.6
	RespTime time.Time `json:"respTime"`
	// Request message ID
	ReqMsgId string `json:"reqMsgId"`
}

type _QueryMerchantResourceResponseResponseHead QueryMerchantResourceResponseResponseHead

// NewQueryMerchantResourceResponseResponseHead instantiates a new QueryMerchantResourceResponseResponseHead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMerchantResourceResponseResponseHead(version string, function string, clientId string, respTime time.Time, reqMsgId string) *QueryMerchantResourceResponseResponseHead {
	this := QueryMerchantResourceResponseResponseHead{}
	this.Version = version
	this.Function = function
	this.ClientId = clientId
	this.RespTime = respTime
	this.ReqMsgId = reqMsgId
	return &this
}

// NewQueryMerchantResourceResponseResponseHeadWithDefaults instantiates a new QueryMerchantResourceResponseResponseHead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMerchantResourceResponseResponseHeadWithDefaults() *QueryMerchantResourceResponseResponseHead {
	this := QueryMerchantResourceResponseResponseHead{}
	var version string = "2.0"
	this.Version = version
	return &this
}

// GetVersion returns the Version field value
func (o *QueryMerchantResourceResponseResponseHead) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceResponseResponseHead) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *QueryMerchantResourceResponseResponseHead) SetVersion(v string) {
	o.Version = v
}

// GetFunction returns the Function field value
func (o *QueryMerchantResourceResponseResponseHead) GetFunction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceResponseResponseHead) GetFunctionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value
func (o *QueryMerchantResourceResponseResponseHead) SetFunction(v string) {
	o.Function = v
}

// GetClientId returns the ClientId field value
func (o *QueryMerchantResourceResponseResponseHead) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceResponseResponseHead) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *QueryMerchantResourceResponseResponseHead) SetClientId(v string) {
	o.ClientId = v
}

// GetRespTime returns the RespTime field value
func (o *QueryMerchantResourceResponseResponseHead) GetRespTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.RespTime
}

// GetRespTimeOk returns a tuple with the RespTime field value
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceResponseResponseHead) GetRespTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RespTime, true
}

// SetRespTime sets field value
func (o *QueryMerchantResourceResponseResponseHead) SetRespTime(v time.Time) {
	o.RespTime = v
}

// GetReqMsgId returns the ReqMsgId field value
func (o *QueryMerchantResourceResponseResponseHead) GetReqMsgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReqMsgId
}

// GetReqMsgIdOk returns a tuple with the ReqMsgId field value
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceResponseResponseHead) GetReqMsgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReqMsgId, true
}

// SetReqMsgId sets field value
func (o *QueryMerchantResourceResponseResponseHead) SetReqMsgId(v string) {
	o.ReqMsgId = v
}

func (o QueryMerchantResourceResponseResponseHead) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMerchantResourceResponseResponseHead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["function"] = o.Function
	toSerialize["clientId"] = o.ClientId
	toSerialize["respTime"] = o.RespTime
	toSerialize["reqMsgId"] = o.ReqMsgId
	return toSerialize, nil
}

func (o *QueryMerchantResourceResponseResponseHead) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"function",
		"clientId",
		"respTime",
		"reqMsgId",
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

	varQueryMerchantResourceResponseResponseHead := _QueryMerchantResourceResponseResponseHead{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryMerchantResourceResponseResponseHead)

	if err != nil {
		return err
	}

	*o = QueryMerchantResourceResponseResponseHead(varQueryMerchantResourceResponseResponseHead)

	return err
}

type NullableQueryMerchantResourceResponseResponseHead struct {
	value *QueryMerchantResourceResponseResponseHead
	isSet bool
}

func (v NullableQueryMerchantResourceResponseResponseHead) Get() *QueryMerchantResourceResponseResponseHead {
	return v.value
}

func (v *NullableQueryMerchantResourceResponseResponseHead) Set(val *QueryMerchantResourceResponseResponseHead) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMerchantResourceResponseResponseHead) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMerchantResourceResponseResponseHead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMerchantResourceResponseResponseHead(val *QueryMerchantResourceResponseResponseHead) *NullableQueryMerchantResourceResponseResponseHead {
	return &NullableQueryMerchantResourceResponseResponseHead{value: val, isSet: true}
}

func (v NullableQueryMerchantResourceResponseResponseHead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMerchantResourceResponseResponseHead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


