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

// checks if the QueryAssetCardListResponseResponseBody type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QueryAssetCardListResponseResponseBody{}

// QueryAssetCardListResponseResponseBody Response body (`resultInfo` mandatory; `assetCardList` optional)
type QueryAssetCardListResponseResponseBody struct {
	ResultInfo MemberAssetResultInfo `json:"resultInfo"`
	// Asset card list detail (present when applicable)
	AssetCardList []AssetCardListItem `json:"assetCardList,omitempty"`
}

type _QueryAssetCardListResponseResponseBody QueryAssetCardListResponseResponseBody

// NewQueryAssetCardListResponseResponseBody instantiates a new QueryAssetCardListResponseResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAssetCardListResponseResponseBody(resultInfo MemberAssetResultInfo) *QueryAssetCardListResponseResponseBody {
	this := QueryAssetCardListResponseResponseBody{}
	this.ResultInfo = resultInfo
	return &this
}

// NewQueryAssetCardListResponseResponseBodyWithDefaults instantiates a new QueryAssetCardListResponseResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAssetCardListResponseResponseBodyWithDefaults() *QueryAssetCardListResponseResponseBody {
	this := QueryAssetCardListResponseResponseBody{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *QueryAssetCardListResponseResponseBody) GetResultInfo() MemberAssetResultInfo {
	if o == nil {
		var ret MemberAssetResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *QueryAssetCardListResponseResponseBody) GetResultInfoOk() (*MemberAssetResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *QueryAssetCardListResponseResponseBody) SetResultInfo(v MemberAssetResultInfo) {
	o.ResultInfo = v
}

// GetAssetCardList returns the AssetCardList field value if set, zero value otherwise.
func (o *QueryAssetCardListResponseResponseBody) GetAssetCardList() []AssetCardListItem {
	if o == nil || utils.IsNil(o.AssetCardList) {
		var ret []AssetCardListItem
		return ret
	}
	return o.AssetCardList
}

// GetAssetCardListOk returns a tuple with the AssetCardList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAssetCardListResponseResponseBody) GetAssetCardListOk() ([]AssetCardListItem, bool) {
	if o == nil || utils.IsNil(o.AssetCardList) {
		return nil, false
	}
	return o.AssetCardList, true
}

// HasAssetCardList returns a boolean if a field has been set.
func (o *QueryAssetCardListResponseResponseBody) HasAssetCardList() bool {
	if o != nil && !utils.IsNil(o.AssetCardList) {
		return true
	}

	return false
}

// SetAssetCardList gets a reference to the given []AssetCardListItem and assigns it to the AssetCardList field.
func (o *QueryAssetCardListResponseResponseBody) SetAssetCardList(v []AssetCardListItem) {
	o.AssetCardList = v
}

func (o QueryAssetCardListResponseResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAssetCardListResponseResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resultInfo"] = o.ResultInfo
	if !utils.IsNil(o.AssetCardList) {
		toSerialize["assetCardList"] = o.AssetCardList
	}
	return toSerialize, nil
}

func (o *QueryAssetCardListResponseResponseBody) UnmarshalJSON(data []byte) (err error) {
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

	varQueryAssetCardListResponseResponseBody := _QueryAssetCardListResponseResponseBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryAssetCardListResponseResponseBody)

	if err != nil {
		return err
	}

	*o = QueryAssetCardListResponseResponseBody(varQueryAssetCardListResponseResponseBody)

	return err
}

type NullableQueryAssetCardListResponseResponseBody struct {
	value *QueryAssetCardListResponseResponseBody
	isSet bool
}

func (v NullableQueryAssetCardListResponseResponseBody) Get() *QueryAssetCardListResponseResponseBody {
	return v.value
}

func (v *NullableQueryAssetCardListResponseResponseBody) Set(val *QueryAssetCardListResponseResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAssetCardListResponseResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAssetCardListResponseResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryAssetCardListResponseResponseBody(val *QueryAssetCardListResponseResponseBody) *NullableQueryAssetCardListResponseResponseBody {
	return &NullableQueryAssetCardListResponseResponseBody{value: val, isSet: true}
}

func (v NullableQueryAssetCardListResponseResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAssetCardListResponseResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


