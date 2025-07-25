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

// checks if the QueryShopResponseResponseHead type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QueryShopResponseResponseHead{}

// QueryShopResponseResponseHead struct for QueryShopResponseResponseHead
type QueryShopResponseResponseHead struct {
	// API version
	Version *string `json:"version,omitempty"`
	// API interface
	Function *string `json:"function,omitempty"`
	// Client ID provided by DANA, used to identify partner and application system
	ClientId *string `json:"clientId,omitempty"`
	// Response time, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time)
	RespTime *string `json:"respTime,omitempty"`
	// Request message ID
	ReqMsgId *string `json:"reqMsgId,omitempty"`
	// As a secret key of client. Assigned client secret during registration
	ClientSecret *string `json:"clientSecret,omitempty"`
	// Reserved for future implementation (Key/Value)
	Reserve *string `json:"reserve,omitempty"`
}

// NewQueryShopResponseResponseHead instantiates a new QueryShopResponseResponseHead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryShopResponseResponseHead() *QueryShopResponseResponseHead {
	this := QueryShopResponseResponseHead{}
	var version string = "2.0"
	this.Version = &version
	return &this
}

// NewQueryShopResponseResponseHeadWithDefaults instantiates a new QueryShopResponseResponseHead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryShopResponseResponseHeadWithDefaults() *QueryShopResponseResponseHead {
	this := QueryShopResponseResponseHead{}
	var version string = "2.0"
	this.Version = &version
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *QueryShopResponseResponseHead) GetVersion() string {
	if o == nil || utils.IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryShopResponseResponseHead) GetVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *QueryShopResponseResponseHead) HasVersion() bool {
	if o != nil && !utils.IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *QueryShopResponseResponseHead) SetVersion(v string) {
	o.Version = &v
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *QueryShopResponseResponseHead) GetFunction() string {
	if o == nil || utils.IsNil(o.Function) {
		var ret string
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryShopResponseResponseHead) GetFunctionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Function) {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *QueryShopResponseResponseHead) HasFunction() bool {
	if o != nil && !utils.IsNil(o.Function) {
		return true
	}

	return false
}

// SetFunction gets a reference to the given string and assigns it to the Function field.
func (o *QueryShopResponseResponseHead) SetFunction(v string) {
	o.Function = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *QueryShopResponseResponseHead) GetClientId() string {
	if o == nil || utils.IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryShopResponseResponseHead) GetClientIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *QueryShopResponseResponseHead) HasClientId() bool {
	if o != nil && !utils.IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *QueryShopResponseResponseHead) SetClientId(v string) {
	o.ClientId = &v
}

// GetRespTime returns the RespTime field value if set, zero value otherwise.
func (o *QueryShopResponseResponseHead) GetRespTime() string {
	if o == nil || utils.IsNil(o.RespTime) {
		var ret string
		return ret
	}
	return *o.RespTime
}

// GetRespTimeOk returns a tuple with the RespTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryShopResponseResponseHead) GetRespTimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RespTime) {
		return nil, false
	}
	return o.RespTime, true
}

// HasRespTime returns a boolean if a field has been set.
func (o *QueryShopResponseResponseHead) HasRespTime() bool {
	if o != nil && !utils.IsNil(o.RespTime) {
		return true
	}

	return false
}

// SetRespTime gets a reference to the given string and assigns it to the RespTime field.
func (o *QueryShopResponseResponseHead) SetRespTime(v string) {
	o.RespTime = &v
}

// GetReqMsgId returns the ReqMsgId field value if set, zero value otherwise.
func (o *QueryShopResponseResponseHead) GetReqMsgId() string {
	if o == nil || utils.IsNil(o.ReqMsgId) {
		var ret string
		return ret
	}
	return *o.ReqMsgId
}

// GetReqMsgIdOk returns a tuple with the ReqMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryShopResponseResponseHead) GetReqMsgIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReqMsgId) {
		return nil, false
	}
	return o.ReqMsgId, true
}

// HasReqMsgId returns a boolean if a field has been set.
func (o *QueryShopResponseResponseHead) HasReqMsgId() bool {
	if o != nil && !utils.IsNil(o.ReqMsgId) {
		return true
	}

	return false
}

// SetReqMsgId gets a reference to the given string and assigns it to the ReqMsgId field.
func (o *QueryShopResponseResponseHead) SetReqMsgId(v string) {
	o.ReqMsgId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *QueryShopResponseResponseHead) GetClientSecret() string {
	if o == nil || utils.IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryShopResponseResponseHead) GetClientSecretOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *QueryShopResponseResponseHead) HasClientSecret() bool {
	if o != nil && !utils.IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *QueryShopResponseResponseHead) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetReserve returns the Reserve field value if set, zero value otherwise.
func (o *QueryShopResponseResponseHead) GetReserve() string {
	if o == nil || utils.IsNil(o.Reserve) {
		var ret string
		return ret
	}
	return *o.Reserve
}

// GetReserveOk returns a tuple with the Reserve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryShopResponseResponseHead) GetReserveOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Reserve) {
		return nil, false
	}
	return o.Reserve, true
}

// HasReserve returns a boolean if a field has been set.
func (o *QueryShopResponseResponseHead) HasReserve() bool {
	if o != nil && !utils.IsNil(o.Reserve) {
		return true
	}

	return false
}

// SetReserve gets a reference to the given string and assigns it to the Reserve field.
func (o *QueryShopResponseResponseHead) SetReserve(v string) {
	o.Reserve = &v
}

func (o QueryShopResponseResponseHead) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryShopResponseResponseHead) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !utils.IsNil(o.Reserve) {
		toSerialize["reserve"] = o.Reserve
	}
	return toSerialize, nil
}

type NullableQueryShopResponseResponseHead struct {
	value *QueryShopResponseResponseHead
	isSet bool
}

func (v NullableQueryShopResponseResponseHead) Get() *QueryShopResponseResponseHead {
	return v.value
}

func (v *NullableQueryShopResponseResponseHead) Set(val *QueryShopResponseResponseHead) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryShopResponseResponseHead) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryShopResponseResponseHead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryShopResponseResponseHead(val *QueryShopResponseResponseHead) *NullableQueryShopResponseResponseHead {
	return &NullableQueryShopResponseResponseHead{value: val, isSet: true}
}

func (v NullableQueryShopResponseResponseHead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryShopResponseResponseHead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


