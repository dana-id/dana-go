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

// checks if the QueryAssetCardListRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QueryAssetCardListRequest{}

// QueryAssetCardListRequest Business fields mapped to `request.body` in the JSON envelope (head/signature are handled by the client/runtime).
type QueryAssetCardListRequest struct {
	// User identifier or merchant identifier
	MemberId string `json:"memberId"`
	// Asset card bind identifier
	BindingId *string `json:"bindingId,omitempty"`
	// Return only active cards when `\"true\"`
	EnableOnly *string `json:"enableOnly,omitempty"`
	// Contact biz type list (`ContactBizTypeEnum`)
	ContactBizTypeList []string `json:"contactBizTypeList,omitempty"`
	// Asset type list (`AssetCardTypeEnum`)
	AssetTypeList []string `json:"assetTypeList,omitempty"`
}

type _QueryAssetCardListRequest QueryAssetCardListRequest

// NewQueryAssetCardListRequest instantiates a new QueryAssetCardListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAssetCardListRequest(memberId string) *QueryAssetCardListRequest {
	this := QueryAssetCardListRequest{}
	this.MemberId = memberId
	return &this
}

// NewQueryAssetCardListRequestWithDefaults instantiates a new QueryAssetCardListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAssetCardListRequestWithDefaults() *QueryAssetCardListRequest {
	this := QueryAssetCardListRequest{}
	return &this
}

// GetMemberId returns the MemberId field value
func (o *QueryAssetCardListRequest) GetMemberId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value
// and a boolean to check if the value has been set.
func (o *QueryAssetCardListRequest) GetMemberIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberId, true
}

// SetMemberId sets field value
func (o *QueryAssetCardListRequest) SetMemberId(v string) {
	o.MemberId = v
}

// GetBindingId returns the BindingId field value if set, zero value otherwise.
func (o *QueryAssetCardListRequest) GetBindingId() string {
	if o == nil || utils.IsNil(o.BindingId) {
		var ret string
		return ret
	}
	return *o.BindingId
}

// GetBindingIdOk returns a tuple with the BindingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAssetCardListRequest) GetBindingIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BindingId) {
		return nil, false
	}
	return o.BindingId, true
}

// HasBindingId returns a boolean if a field has been set.
func (o *QueryAssetCardListRequest) HasBindingId() bool {
	if o != nil && !utils.IsNil(o.BindingId) {
		return true
	}

	return false
}

// SetBindingId gets a reference to the given string and assigns it to the BindingId field.
func (o *QueryAssetCardListRequest) SetBindingId(v string) {
	o.BindingId = &v
}

// GetEnableOnly returns the EnableOnly field value if set, zero value otherwise.
func (o *QueryAssetCardListRequest) GetEnableOnly() string {
	if o == nil || utils.IsNil(o.EnableOnly) {
		var ret string
		return ret
	}
	return *o.EnableOnly
}

// GetEnableOnlyOk returns a tuple with the EnableOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAssetCardListRequest) GetEnableOnlyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EnableOnly) {
		return nil, false
	}
	return o.EnableOnly, true
}

// HasEnableOnly returns a boolean if a field has been set.
func (o *QueryAssetCardListRequest) HasEnableOnly() bool {
	if o != nil && !utils.IsNil(o.EnableOnly) {
		return true
	}

	return false
}

// SetEnableOnly gets a reference to the given string and assigns it to the EnableOnly field.
func (o *QueryAssetCardListRequest) SetEnableOnly(v string) {
	o.EnableOnly = &v
}

// GetContactBizTypeList returns the ContactBizTypeList field value if set, zero value otherwise.
func (o *QueryAssetCardListRequest) GetContactBizTypeList() []string {
	if o == nil || utils.IsNil(o.ContactBizTypeList) {
		var ret []string
		return ret
	}
	return o.ContactBizTypeList
}

// GetContactBizTypeListOk returns a tuple with the ContactBizTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAssetCardListRequest) GetContactBizTypeListOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ContactBizTypeList) {
		return nil, false
	}
	return o.ContactBizTypeList, true
}

// HasContactBizTypeList returns a boolean if a field has been set.
func (o *QueryAssetCardListRequest) HasContactBizTypeList() bool {
	if o != nil && !utils.IsNil(o.ContactBizTypeList) {
		return true
	}

	return false
}

// SetContactBizTypeList gets a reference to the given []string and assigns it to the ContactBizTypeList field.
func (o *QueryAssetCardListRequest) SetContactBizTypeList(v []string) {
	o.ContactBizTypeList = v
}

// GetAssetTypeList returns the AssetTypeList field value if set, zero value otherwise.
func (o *QueryAssetCardListRequest) GetAssetTypeList() []string {
	if o == nil || utils.IsNil(o.AssetTypeList) {
		var ret []string
		return ret
	}
	return o.AssetTypeList
}

// GetAssetTypeListOk returns a tuple with the AssetTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAssetCardListRequest) GetAssetTypeListOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AssetTypeList) {
		return nil, false
	}
	return o.AssetTypeList, true
}

// HasAssetTypeList returns a boolean if a field has been set.
func (o *QueryAssetCardListRequest) HasAssetTypeList() bool {
	if o != nil && !utils.IsNil(o.AssetTypeList) {
		return true
	}

	return false
}

// SetAssetTypeList gets a reference to the given []string and assigns it to the AssetTypeList field.
func (o *QueryAssetCardListRequest) SetAssetTypeList(v []string) {
	o.AssetTypeList = v
}

func (o QueryAssetCardListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAssetCardListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["memberId"] = o.MemberId
	if !utils.IsNil(o.BindingId) {
		toSerialize["bindingId"] = o.BindingId
	}
	if !utils.IsNil(o.EnableOnly) {
		toSerialize["enableOnly"] = o.EnableOnly
	}
	if !utils.IsNil(o.ContactBizTypeList) {
		toSerialize["contactBizTypeList"] = o.ContactBizTypeList
	}
	if !utils.IsNil(o.AssetTypeList) {
		toSerialize["assetTypeList"] = o.AssetTypeList
	}
	return toSerialize, nil
}

func (o *QueryAssetCardListRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"memberId",
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

	varQueryAssetCardListRequest := _QueryAssetCardListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryAssetCardListRequest)

	if err != nil {
		return err
	}

	*o = QueryAssetCardListRequest(varQueryAssetCardListRequest)

	return err
}

type NullableQueryAssetCardListRequest struct {
	value *QueryAssetCardListRequest
	isSet bool
}

func (v NullableQueryAssetCardListRequest) Get() *QueryAssetCardListRequest {
	return v.value
}

func (v *NullableQueryAssetCardListRequest) Set(val *QueryAssetCardListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAssetCardListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAssetCardListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryAssetCardListRequest(val *QueryAssetCardListRequest) *NullableQueryAssetCardListRequest {
	return &NullableQueryAssetCardListRequest{value: val, isSet: true}
}

func (v NullableQueryAssetCardListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAssetCardListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


