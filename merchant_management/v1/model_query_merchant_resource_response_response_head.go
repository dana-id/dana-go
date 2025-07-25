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
	utils "github.com/dana-id/dana-go/utils"
)

// checks if the QueryMerchantResourceResponseResponseHead type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QueryMerchantResourceResponseResponseHead{}

// QueryMerchantResourceResponseResponseHead struct for QueryMerchantResourceResponseResponseHead
type QueryMerchantResourceResponseResponseHead struct {
	// API version
	Version *string `json:"version,omitempty"`
	// API interface
	Function *string `json:"function,omitempty"`
	// Client ID provided by DANA, used to identify partner and application system
	ClientId *string `json:"clientId,omitempty"`
	// DateTime with timezone, which follows the ISO-8601 standard. Refer to RFC 3339 Section 5.6
	RespTime *string `json:"respTime,omitempty"`
	// Request message ID
	ReqMsgId *string `json:"reqMsgId,omitempty"`
}

// NewQueryMerchantResourceResponseResponseHead instantiates a new QueryMerchantResourceResponseResponseHead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMerchantResourceResponseResponseHead() *QueryMerchantResourceResponseResponseHead {
	this := QueryMerchantResourceResponseResponseHead{}
	var version string = "2.0"
	this.Version = &version
	return &this
}

// NewQueryMerchantResourceResponseResponseHeadWithDefaults instantiates a new QueryMerchantResourceResponseResponseHead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMerchantResourceResponseResponseHeadWithDefaults() *QueryMerchantResourceResponseResponseHead {
	this := QueryMerchantResourceResponseResponseHead{}
	var version string = "2.0"
	this.Version = &version
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *QueryMerchantResourceResponseResponseHead) GetVersion() string {
	if o == nil || utils.IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceResponseResponseHead) GetVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *QueryMerchantResourceResponseResponseHead) HasVersion() bool {
	if o != nil && !utils.IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *QueryMerchantResourceResponseResponseHead) SetVersion(v string) {
	o.Version = &v
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *QueryMerchantResourceResponseResponseHead) GetFunction() string {
	if o == nil || utils.IsNil(o.Function) {
		var ret string
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceResponseResponseHead) GetFunctionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Function) {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *QueryMerchantResourceResponseResponseHead) HasFunction() bool {
	if o != nil && !utils.IsNil(o.Function) {
		return true
	}

	return false
}

// SetFunction gets a reference to the given string and assigns it to the Function field.
func (o *QueryMerchantResourceResponseResponseHead) SetFunction(v string) {
	o.Function = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *QueryMerchantResourceResponseResponseHead) GetClientId() string {
	if o == nil || utils.IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceResponseResponseHead) GetClientIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *QueryMerchantResourceResponseResponseHead) HasClientId() bool {
	if o != nil && !utils.IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *QueryMerchantResourceResponseResponseHead) SetClientId(v string) {
	o.ClientId = &v
}

// GetRespTime returns the RespTime field value if set, zero value otherwise.
func (o *QueryMerchantResourceResponseResponseHead) GetRespTime() string {
	if o == nil || utils.IsNil(o.RespTime) {
		var ret string
		return ret
	}
	return *o.RespTime
}

// GetRespTimeOk returns a tuple with the RespTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceResponseResponseHead) GetRespTimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RespTime) {
		return nil, false
	}
	return o.RespTime, true
}

// HasRespTime returns a boolean if a field has been set.
func (o *QueryMerchantResourceResponseResponseHead) HasRespTime() bool {
	if o != nil && !utils.IsNil(o.RespTime) {
		return true
	}

	return false
}

// SetRespTime gets a reference to the given string and assigns it to the RespTime field.
func (o *QueryMerchantResourceResponseResponseHead) SetRespTime(v string) {
	o.RespTime = &v
}

// GetReqMsgId returns the ReqMsgId field value if set, zero value otherwise.
func (o *QueryMerchantResourceResponseResponseHead) GetReqMsgId() string {
	if o == nil || utils.IsNil(o.ReqMsgId) {
		var ret string
		return ret
	}
	return *o.ReqMsgId
}

// GetReqMsgIdOk returns a tuple with the ReqMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceResponseResponseHead) GetReqMsgIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReqMsgId) {
		return nil, false
	}
	return o.ReqMsgId, true
}

// HasReqMsgId returns a boolean if a field has been set.
func (o *QueryMerchantResourceResponseResponseHead) HasReqMsgId() bool {
	if o != nil && !utils.IsNil(o.ReqMsgId) {
		return true
	}

	return false
}

// SetReqMsgId gets a reference to the given string and assigns it to the ReqMsgId field.
func (o *QueryMerchantResourceResponseResponseHead) SetReqMsgId(v string) {
	o.ReqMsgId = &v
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
	if !utils.IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !utils.IsNil(o.Function) {
		toSerialize["function"] = o.Function
	}
	if !utils.IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !utils.IsNil(o.RespTime) {
		toSerialize["respTime"] = o.RespTime
	}
	if !utils.IsNil(o.ReqMsgId) {
		toSerialize["reqMsgId"] = o.ReqMsgId
	}
	return toSerialize, nil
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


