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

package webhook

import (
	"encoding/json"
	"bytes"
	"fmt"
	utils "github.com/dana-id/dana-go/utils"
)

// checks if the FinishNotifyResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FinishNotifyResponse{}

// FinishNotifyResponse struct for FinishNotifyResponse
type FinishNotifyResponse struct {
	// Response code. Refer to https://dashboard.dana.id/api-docs/read/123#HTML-API-FinishNotify-ResponseCodeandMessage
	ResponseCode string `json:"responseCode"`
	// Response message. Refer to https://dashboard.dana.id/api-docs/read/123#HTML-API-FinishNotify-ResponseCodeandMessage
	ResponseMessage string `json:"responseMessage"`
}

type _FinishNotifyResponse FinishNotifyResponse

// NewFinishNotifyResponse instantiates a new FinishNotifyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinishNotifyResponse(responseCode string, responseMessage string) *FinishNotifyResponse {
	this := FinishNotifyResponse{}
	this.ResponseCode = responseCode
	this.ResponseMessage = responseMessage
	return &this
}

// NewFinishNotifyResponseWithDefaults instantiates a new FinishNotifyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinishNotifyResponseWithDefaults() *FinishNotifyResponse {
	this := FinishNotifyResponse{}
	return &this
}

// GetResponseCode returns the ResponseCode field value
func (o *FinishNotifyResponse) GetResponseCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value
// and a boolean to check if the value has been set.
func (o *FinishNotifyResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseCode, true
}

// SetResponseCode sets field value
func (o *FinishNotifyResponse) SetResponseCode(v string) {
	o.ResponseCode = v
}

// GetResponseMessage returns the ResponseMessage field value
func (o *FinishNotifyResponse) GetResponseMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value
// and a boolean to check if the value has been set.
func (o *FinishNotifyResponse) GetResponseMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseMessage, true
}

// SetResponseMessage sets field value
func (o *FinishNotifyResponse) SetResponseMessage(v string) {
	o.ResponseMessage = v
}

func (o FinishNotifyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinishNotifyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["responseCode"] = o.ResponseCode
	toSerialize["responseMessage"] = o.ResponseMessage
	return toSerialize, nil
}

func (o *FinishNotifyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"responseCode",
		"responseMessage",
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

	varFinishNotifyResponse := _FinishNotifyResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFinishNotifyResponse)

	if err != nil {
		return err
	}

	*o = FinishNotifyResponse(varFinishNotifyResponse)

	return err
}

type NullableFinishNotifyResponse struct {
	value *FinishNotifyResponse
	isSet bool
}

func (v NullableFinishNotifyResponse) Get() *FinishNotifyResponse {
	return v.value
}

func (v *NullableFinishNotifyResponse) Set(val *FinishNotifyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFinishNotifyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFinishNotifyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinishNotifyResponse(val *FinishNotifyResponse) *NullableFinishNotifyResponse {
	return &NullableFinishNotifyResponse{value: val, isSet: true}
}

func (v NullableFinishNotifyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinishNotifyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


