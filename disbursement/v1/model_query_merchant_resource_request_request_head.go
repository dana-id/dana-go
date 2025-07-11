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

// checks if the QueryMerchantResourceRequestRequestHead type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QueryMerchantResourceRequestRequestHead{}

// QueryMerchantResourceRequestRequestHead struct for QueryMerchantResourceRequestRequestHead
type QueryMerchantResourceRequestRequestHead struct {
	// API version
	Version string `json:"version"`
	// API interface
	Function string `json:"function"`
	// Client ID provided by DANA, used to identify partner and application system
	ClientId string `json:"clientId"`
	// DateTime with timezone, which follows the ISO-8601 standard. Refer to RFC 3339 Section 5.6
	ReqTime time.Time `json:"reqTime"`
	// The sequence id of request - Each request will be assigned with a unique id (uuid)
	ReqMsgId string `json:"reqMsgId"`
	// Secret key of client - Assigned client secret during registration
	ClientSecret string `json:"clientSecret"`
	// Reserved for future implementation
	Reserve *string `json:"reserve,omitempty"`
}

type _QueryMerchantResourceRequestRequestHead QueryMerchantResourceRequestRequestHead

// NewQueryMerchantResourceRequestRequestHead instantiates a new QueryMerchantResourceRequestRequestHead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMerchantResourceRequestRequestHead(version string, function string, clientId string, reqTime time.Time, reqMsgId string, clientSecret string) *QueryMerchantResourceRequestRequestHead {
	this := QueryMerchantResourceRequestRequestHead{}
	this.Version = version
	this.Function = function
	this.ClientId = clientId
	this.ReqTime = reqTime
	this.ReqMsgId = reqMsgId
	this.ClientSecret = clientSecret
	return &this
}

// NewQueryMerchantResourceRequestRequestHeadWithDefaults instantiates a new QueryMerchantResourceRequestRequestHead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMerchantResourceRequestRequestHeadWithDefaults() *QueryMerchantResourceRequestRequestHead {
	this := QueryMerchantResourceRequestRequestHead{}
	var version string = "2.0"
	this.Version = version
	var function string = "dana.merchant.queryMerchantResource"
	this.Function = function
	return &this
}

// GetVersion returns the Version field value
func (o *QueryMerchantResourceRequestRequestHead) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceRequestRequestHead) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *QueryMerchantResourceRequestRequestHead) SetVersion(v string) {
	o.Version = v
}

// GetFunction returns the Function field value
func (o *QueryMerchantResourceRequestRequestHead) GetFunction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceRequestRequestHead) GetFunctionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value
func (o *QueryMerchantResourceRequestRequestHead) SetFunction(v string) {
	o.Function = v
}

// GetClientId returns the ClientId field value
func (o *QueryMerchantResourceRequestRequestHead) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceRequestRequestHead) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *QueryMerchantResourceRequestRequestHead) SetClientId(v string) {
	o.ClientId = v
}

// GetReqTime returns the ReqTime field value
func (o *QueryMerchantResourceRequestRequestHead) GetReqTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ReqTime
}

// GetReqTimeOk returns a tuple with the ReqTime field value
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceRequestRequestHead) GetReqTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReqTime, true
}

// SetReqTime sets field value
func (o *QueryMerchantResourceRequestRequestHead) SetReqTime(v time.Time) {
	o.ReqTime = v
}

// GetReqMsgId returns the ReqMsgId field value
func (o *QueryMerchantResourceRequestRequestHead) GetReqMsgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReqMsgId
}

// GetReqMsgIdOk returns a tuple with the ReqMsgId field value
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceRequestRequestHead) GetReqMsgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReqMsgId, true
}

// SetReqMsgId sets field value
func (o *QueryMerchantResourceRequestRequestHead) SetReqMsgId(v string) {
	o.ReqMsgId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *QueryMerchantResourceRequestRequestHead) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceRequestRequestHead) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *QueryMerchantResourceRequestRequestHead) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetReserve returns the Reserve field value if set, zero value otherwise.
func (o *QueryMerchantResourceRequestRequestHead) GetReserve() string {
	if o == nil || utils.IsNil(o.Reserve) {
		var ret string
		return ret
	}
	return *o.Reserve
}

// GetReserveOk returns a tuple with the Reserve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMerchantResourceRequestRequestHead) GetReserveOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Reserve) {
		return nil, false
	}
	return o.Reserve, true
}

// HasReserve returns a boolean if a field has been set.
func (o *QueryMerchantResourceRequestRequestHead) HasReserve() bool {
	if o != nil && !utils.IsNil(o.Reserve) {
		return true
	}

	return false
}

// SetReserve gets a reference to the given string and assigns it to the Reserve field.
func (o *QueryMerchantResourceRequestRequestHead) SetReserve(v string) {
	o.Reserve = &v
}

func (o QueryMerchantResourceRequestRequestHead) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMerchantResourceRequestRequestHead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["function"] = o.Function
	toSerialize["clientId"] = o.ClientId
	toSerialize["reqTime"] = o.ReqTime
	toSerialize["reqMsgId"] = o.ReqMsgId
	toSerialize["clientSecret"] = o.ClientSecret
	if !utils.IsNil(o.Reserve) {
		toSerialize["reserve"] = o.Reserve
	}
	return toSerialize, nil
}

func (o *QueryMerchantResourceRequestRequestHead) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"function",
		"clientId",
		"reqTime",
		"reqMsgId",
		"clientSecret",
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

	varQueryMerchantResourceRequestRequestHead := _QueryMerchantResourceRequestRequestHead{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryMerchantResourceRequestRequestHead)

	if err != nil {
		return err
	}

	*o = QueryMerchantResourceRequestRequestHead(varQueryMerchantResourceRequestRequestHead)

	return err
}

type NullableQueryMerchantResourceRequestRequestHead struct {
	value *QueryMerchantResourceRequestRequestHead
	isSet bool
}

func (v NullableQueryMerchantResourceRequestRequestHead) Get() *QueryMerchantResourceRequestRequestHead {
	return v.value
}

func (v *NullableQueryMerchantResourceRequestRequestHead) Set(val *QueryMerchantResourceRequestRequestHead) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMerchantResourceRequestRequestHead) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMerchantResourceRequestRequestHead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMerchantResourceRequestRequestHead(val *QueryMerchantResourceRequestRequestHead) *NullableQueryMerchantResourceRequestRequestHead {
	return &NullableQueryMerchantResourceRequestRequestHead{value: val, isSet: true}
}

func (v NullableQueryMerchantResourceRequestRequestHead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMerchantResourceRequestRequestHead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


