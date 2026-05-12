// Copyright 2026 PT Espay Debit Indonesia Koe
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

import (
	"encoding/json"
	"bytes"
	"fmt"
	utils "github.com/dana-id/dana-go/v2/utils"
)

// checks if the MemberAssetResultInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MemberAssetResultInfo{}

// MemberAssetResultInfo Result information for member asset Open APIs (`resultInfo` per DANA spec)
type MemberAssetResultInfo struct {
	// Request status (`ResultStatus`). S=Success, F=Failure, U=Unknown
	ResultStatus string `json:"resultStatus"`
	// Result code identifier (`ResultCodeId`)
	ResultCodeId string `json:"resultCodeId"`
	// Result code (`ResultCode`)
	ResultCode string `json:"resultCode"`
	// Result message (`ResultMsg`). Optional per API contract
	ResultMsg *string `json:"resultMsg,omitempty"`
}

type _MemberAssetResultInfo MemberAssetResultInfo

// NewMemberAssetResultInfo instantiates a new MemberAssetResultInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberAssetResultInfo(resultStatus string, resultCodeId string, resultCode string) *MemberAssetResultInfo {
	this := MemberAssetResultInfo{}
	this.ResultStatus = resultStatus
	this.ResultCodeId = resultCodeId
	this.ResultCode = resultCode
	return &this
}

// NewMemberAssetResultInfoWithDefaults instantiates a new MemberAssetResultInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberAssetResultInfoWithDefaults() *MemberAssetResultInfo {
	this := MemberAssetResultInfo{}
	return &this
}

// GetResultStatus returns the ResultStatus field value
func (o *MemberAssetResultInfo) GetResultStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultStatus
}

// GetResultStatusOk returns a tuple with the ResultStatus field value
// and a boolean to check if the value has been set.
func (o *MemberAssetResultInfo) GetResultStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultStatus, true
}

// SetResultStatus sets field value
func (o *MemberAssetResultInfo) SetResultStatus(v string) {
	o.ResultStatus = v
}

// GetResultCodeId returns the ResultCodeId field value
func (o *MemberAssetResultInfo) GetResultCodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultCodeId
}

// GetResultCodeIdOk returns a tuple with the ResultCodeId field value
// and a boolean to check if the value has been set.
func (o *MemberAssetResultInfo) GetResultCodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultCodeId, true
}

// SetResultCodeId sets field value
func (o *MemberAssetResultInfo) SetResultCodeId(v string) {
	o.ResultCodeId = v
}

// GetResultCode returns the ResultCode field value
func (o *MemberAssetResultInfo) GetResultCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value
// and a boolean to check if the value has been set.
func (o *MemberAssetResultInfo) GetResultCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultCode, true
}

// SetResultCode sets field value
func (o *MemberAssetResultInfo) SetResultCode(v string) {
	o.ResultCode = v
}

// GetResultMsg returns the ResultMsg field value if set, zero value otherwise.
func (o *MemberAssetResultInfo) GetResultMsg() string {
	if o == nil || utils.IsNil(o.ResultMsg) {
		var ret string
		return ret
	}
	return *o.ResultMsg
}

// GetResultMsgOk returns a tuple with the ResultMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAssetResultInfo) GetResultMsgOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ResultMsg) {
		return nil, false
	}
	return o.ResultMsg, true
}

// HasResultMsg returns a boolean if a field has been set.
func (o *MemberAssetResultInfo) HasResultMsg() bool {
	if o != nil && !utils.IsNil(o.ResultMsg) {
		return true
	}

	return false
}

// SetResultMsg gets a reference to the given string and assigns it to the ResultMsg field.
func (o *MemberAssetResultInfo) SetResultMsg(v string) {
	o.ResultMsg = &v
}

func (o MemberAssetResultInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberAssetResultInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resultStatus"] = o.ResultStatus
	toSerialize["resultCodeId"] = o.ResultCodeId
	toSerialize["resultCode"] = o.ResultCode
	if !utils.IsNil(o.ResultMsg) {
		toSerialize["resultMsg"] = o.ResultMsg
	}
	return toSerialize, nil
}

func (o *MemberAssetResultInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resultStatus",
		"resultCodeId",
		"resultCode",
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

	varMemberAssetResultInfo := _MemberAssetResultInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMemberAssetResultInfo)

	if err != nil {
		return err
	}

	*o = MemberAssetResultInfo(varMemberAssetResultInfo)

	return err
}

type NullableMemberAssetResultInfo struct {
	value *MemberAssetResultInfo
	isSet bool
}

func (v NullableMemberAssetResultInfo) Get() *MemberAssetResultInfo {
	return v.value
}

func (v *NullableMemberAssetResultInfo) Set(val *MemberAssetResultInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberAssetResultInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberAssetResultInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberAssetResultInfo(val *MemberAssetResultInfo) *NullableMemberAssetResultInfo {
	return &NullableMemberAssetResultInfo{value: val, isSet: true}
}

func (v NullableMemberAssetResultInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberAssetResultInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


