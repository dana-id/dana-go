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
	utils "github.com/dana-id/dana-go/utils"
)

// checks if the ShopInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ShopInfo{}

// ShopInfo struct for ShopInfo
type ShopInfo struct {
	// Information of shop identifier. Required if externalShopId is blank
	ShopId *string `json:"shopId,omitempty"`
	// Information of external shop identifier. Required if shopId is blank
	ExternalShopId *string `json:"externalShopId,omitempty"`
	// Information of operator identifier
	OperatorId *string `json:"operatorId,omitempty"`
	// Information of shop address
	ShopAddress *string `json:"shopAddress,omitempty"`
	// Information of division identifier
	DivisionId *string `json:"divisionId,omitempty"`
	// Information of external division identifier
	ExternalDivisionId *string `json:"externalDivisionId,omitempty"`
	// Information of division type
	DivisionType *string `json:"divisionType,omitempty"`
	// Information of shop name
	ShopName *string `json:"shopName,omitempty"`
}

// NewShopInfo instantiates a new ShopInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShopInfo() *ShopInfo {
	this := ShopInfo{}
	return &this
}

// NewShopInfoWithDefaults instantiates a new ShopInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShopInfoWithDefaults() *ShopInfo {
	this := ShopInfo{}
	return &this
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *ShopInfo) GetShopId() string {
	if o == nil || utils.IsNil(o.ShopId) {
		var ret string
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopInfo) GetShopIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *ShopInfo) HasShopId() bool {
	if o != nil && !utils.IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given string and assigns it to the ShopId field.
func (o *ShopInfo) SetShopId(v string) {
	o.ShopId = &v
}

// GetExternalShopId returns the ExternalShopId field value if set, zero value otherwise.
func (o *ShopInfo) GetExternalShopId() string {
	if o == nil || utils.IsNil(o.ExternalShopId) {
		var ret string
		return ret
	}
	return *o.ExternalShopId
}

// GetExternalShopIdOk returns a tuple with the ExternalShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopInfo) GetExternalShopIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExternalShopId) {
		return nil, false
	}
	return o.ExternalShopId, true
}

// HasExternalShopId returns a boolean if a field has been set.
func (o *ShopInfo) HasExternalShopId() bool {
	if o != nil && !utils.IsNil(o.ExternalShopId) {
		return true
	}

	return false
}

// SetExternalShopId gets a reference to the given string and assigns it to the ExternalShopId field.
func (o *ShopInfo) SetExternalShopId(v string) {
	o.ExternalShopId = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *ShopInfo) GetOperatorId() string {
	if o == nil || utils.IsNil(o.OperatorId) {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopInfo) GetOperatorIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OperatorId) {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *ShopInfo) HasOperatorId() bool {
	if o != nil && !utils.IsNil(o.OperatorId) {
		return true
	}

	return false
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *ShopInfo) SetOperatorId(v string) {
	o.OperatorId = &v
}

// GetShopAddress returns the ShopAddress field value if set, zero value otherwise.
func (o *ShopInfo) GetShopAddress() string {
	if o == nil || utils.IsNil(o.ShopAddress) {
		var ret string
		return ret
	}
	return *o.ShopAddress
}

// GetShopAddressOk returns a tuple with the ShopAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopInfo) GetShopAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ShopAddress) {
		return nil, false
	}
	return o.ShopAddress, true
}

// HasShopAddress returns a boolean if a field has been set.
func (o *ShopInfo) HasShopAddress() bool {
	if o != nil && !utils.IsNil(o.ShopAddress) {
		return true
	}

	return false
}

// SetShopAddress gets a reference to the given string and assigns it to the ShopAddress field.
func (o *ShopInfo) SetShopAddress(v string) {
	o.ShopAddress = &v
}

// GetDivisionId returns the DivisionId field value if set, zero value otherwise.
func (o *ShopInfo) GetDivisionId() string {
	if o == nil || utils.IsNil(o.DivisionId) {
		var ret string
		return ret
	}
	return *o.DivisionId
}

// GetDivisionIdOk returns a tuple with the DivisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopInfo) GetDivisionIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DivisionId) {
		return nil, false
	}
	return o.DivisionId, true
}

// HasDivisionId returns a boolean if a field has been set.
func (o *ShopInfo) HasDivisionId() bool {
	if o != nil && !utils.IsNil(o.DivisionId) {
		return true
	}

	return false
}

// SetDivisionId gets a reference to the given string and assigns it to the DivisionId field.
func (o *ShopInfo) SetDivisionId(v string) {
	o.DivisionId = &v
}

// GetExternalDivisionId returns the ExternalDivisionId field value if set, zero value otherwise.
func (o *ShopInfo) GetExternalDivisionId() string {
	if o == nil || utils.IsNil(o.ExternalDivisionId) {
		var ret string
		return ret
	}
	return *o.ExternalDivisionId
}

// GetExternalDivisionIdOk returns a tuple with the ExternalDivisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopInfo) GetExternalDivisionIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExternalDivisionId) {
		return nil, false
	}
	return o.ExternalDivisionId, true
}

// HasExternalDivisionId returns a boolean if a field has been set.
func (o *ShopInfo) HasExternalDivisionId() bool {
	if o != nil && !utils.IsNil(o.ExternalDivisionId) {
		return true
	}

	return false
}

// SetExternalDivisionId gets a reference to the given string and assigns it to the ExternalDivisionId field.
func (o *ShopInfo) SetExternalDivisionId(v string) {
	o.ExternalDivisionId = &v
}

// GetDivisionType returns the DivisionType field value if set, zero value otherwise.
func (o *ShopInfo) GetDivisionType() string {
	if o == nil || utils.IsNil(o.DivisionType) {
		var ret string
		return ret
	}
	return *o.DivisionType
}

// GetDivisionTypeOk returns a tuple with the DivisionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopInfo) GetDivisionTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DivisionType) {
		return nil, false
	}
	return o.DivisionType, true
}

// HasDivisionType returns a boolean if a field has been set.
func (o *ShopInfo) HasDivisionType() bool {
	if o != nil && !utils.IsNil(o.DivisionType) {
		return true
	}

	return false
}

// SetDivisionType gets a reference to the given string and assigns it to the DivisionType field.
func (o *ShopInfo) SetDivisionType(v string) {
	o.DivisionType = &v
}

// GetShopName returns the ShopName field value if set, zero value otherwise.
func (o *ShopInfo) GetShopName() string {
	if o == nil || utils.IsNil(o.ShopName) {
		var ret string
		return ret
	}
	return *o.ShopName
}

// GetShopNameOk returns a tuple with the ShopName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopInfo) GetShopNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ShopName) {
		return nil, false
	}
	return o.ShopName, true
}

// HasShopName returns a boolean if a field has been set.
func (o *ShopInfo) HasShopName() bool {
	if o != nil && !utils.IsNil(o.ShopName) {
		return true
	}

	return false
}

// SetShopName gets a reference to the given string and assigns it to the ShopName field.
func (o *ShopInfo) SetShopName(v string) {
	o.ShopName = &v
}

func (o ShopInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShopInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ShopId) {
		toSerialize["shopId"] = o.ShopId
	}
	if !utils.IsNil(o.ExternalShopId) {
		toSerialize["externalShopId"] = o.ExternalShopId
	}
	if !utils.IsNil(o.OperatorId) {
		toSerialize["operatorId"] = o.OperatorId
	}
	if !utils.IsNil(o.ShopAddress) {
		toSerialize["shopAddress"] = o.ShopAddress
	}
	if !utils.IsNil(o.DivisionId) {
		toSerialize["divisionId"] = o.DivisionId
	}
	if !utils.IsNil(o.ExternalDivisionId) {
		toSerialize["externalDivisionId"] = o.ExternalDivisionId
	}
	if !utils.IsNil(o.DivisionType) {
		toSerialize["divisionType"] = o.DivisionType
	}
	if !utils.IsNil(o.ShopName) {
		toSerialize["shopName"] = o.ShopName
	}
	return toSerialize, nil
}

type NullableShopInfo struct {
	value *ShopInfo
	isSet bool
}

func (v NullableShopInfo) Get() *ShopInfo {
	return v.value
}

func (v *NullableShopInfo) Set(val *ShopInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableShopInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableShopInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShopInfo(val *ShopInfo) *NullableShopInfo {
	return &NullableShopInfo{value: val, isSet: true}
}

func (v NullableShopInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShopInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


