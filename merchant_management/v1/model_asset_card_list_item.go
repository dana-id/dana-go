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

// checks if the AssetCardListItem type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AssetCardListItem{}

// AssetCardListItem One row in `assetCardList` (response body)
type AssetCardListItem struct {
	// Contact business type (`ContactBizTypeEnum`)
	ContactBizType string `json:"contactBizType"`
	// Card index number
	CardIndexNo string `json:"cardIndexNo"`
	// Card number length based on card index number
	CardNoLength string `json:"cardNoLength"`
	// Card number (masked)
	MaskedCardNo string `json:"maskedCardNo"`
	// Asset / card type (`AssetCardTypeEnum`)
	AssetType string `json:"assetType"`
	// Holder name (JSON object per DANA spec)
	HolderName map[string]interface{} `json:"holderName"`
	// Institution logo URL
	InstLogoUrl *string `json:"instLogoUrl,omitempty"`
	// Institution identifier
	InstId string `json:"instId"`
	// Institution official name based on `instId`
	InstOfficialName string `json:"instOfficialName"`
	// Expiry year (e.g. debit, credit, virtual account)
	ExpiryYear string `json:"expiryYear"`
	// Expiry month
	ExpiryMonth string `json:"expiryMonth"`
	// Whether the user's card is verified
	Verified string `json:"verified"`
	// Asset card bind identifier
	BindingId *string `json:"bindingId,omitempty"`
	// Whether this asset is the user's default payment
	DefaultAsset *string `json:"defaultAsset,omitempty"`
	// Whether the card status is enabled; `\"true\"` when `enableOnly` in request was true
	EnableStatus *string `json:"enableStatus,omitempty"`
	// Whether payment uses direct debit
	DirectDebit *string `json:"directDebit,omitempty"`
}

type _AssetCardListItem AssetCardListItem

// NewAssetCardListItem instantiates a new AssetCardListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetCardListItem(contactBizType string, cardIndexNo string, cardNoLength string, maskedCardNo string, assetType string, holderName map[string]interface{}, instId string, instOfficialName string, expiryYear string, expiryMonth string, verified string) *AssetCardListItem {
	this := AssetCardListItem{}
	this.ContactBizType = contactBizType
	this.CardIndexNo = cardIndexNo
	this.CardNoLength = cardNoLength
	this.MaskedCardNo = maskedCardNo
	this.AssetType = assetType
	this.HolderName = holderName
	this.InstId = instId
	this.InstOfficialName = instOfficialName
	this.ExpiryYear = expiryYear
	this.ExpiryMonth = expiryMonth
	this.Verified = verified
	return &this
}

// NewAssetCardListItemWithDefaults instantiates a new AssetCardListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetCardListItemWithDefaults() *AssetCardListItem {
	this := AssetCardListItem{}
	return &this
}

// GetContactBizType returns the ContactBizType field value
func (o *AssetCardListItem) GetContactBizType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactBizType
}

// GetContactBizTypeOk returns a tuple with the ContactBizType field value
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetContactBizTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactBizType, true
}

// SetContactBizType sets field value
func (o *AssetCardListItem) SetContactBizType(v string) {
	o.ContactBizType = v
}

// GetCardIndexNo returns the CardIndexNo field value
func (o *AssetCardListItem) GetCardIndexNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardIndexNo
}

// GetCardIndexNoOk returns a tuple with the CardIndexNo field value
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetCardIndexNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardIndexNo, true
}

// SetCardIndexNo sets field value
func (o *AssetCardListItem) SetCardIndexNo(v string) {
	o.CardIndexNo = v
}

// GetCardNoLength returns the CardNoLength field value
func (o *AssetCardListItem) GetCardNoLength() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardNoLength
}

// GetCardNoLengthOk returns a tuple with the CardNoLength field value
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetCardNoLengthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardNoLength, true
}

// SetCardNoLength sets field value
func (o *AssetCardListItem) SetCardNoLength(v string) {
	o.CardNoLength = v
}

// GetMaskedCardNo returns the MaskedCardNo field value
func (o *AssetCardListItem) GetMaskedCardNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaskedCardNo
}

// GetMaskedCardNoOk returns a tuple with the MaskedCardNo field value
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetMaskedCardNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaskedCardNo, true
}

// SetMaskedCardNo sets field value
func (o *AssetCardListItem) SetMaskedCardNo(v string) {
	o.MaskedCardNo = v
}

// GetAssetType returns the AssetType field value
func (o *AssetCardListItem) GetAssetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetAssetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetType, true
}

// SetAssetType sets field value
func (o *AssetCardListItem) SetAssetType(v string) {
	o.AssetType = v
}

// GetHolderName returns the HolderName field value
func (o *AssetCardListItem) GetHolderName() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.HolderName
}

// GetHolderNameOk returns a tuple with the HolderName field value
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetHolderNameOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.HolderName, true
}

// SetHolderName sets field value
func (o *AssetCardListItem) SetHolderName(v map[string]interface{}) {
	o.HolderName = v
}

// GetInstLogoUrl returns the InstLogoUrl field value if set, zero value otherwise.
func (o *AssetCardListItem) GetInstLogoUrl() string {
	if o == nil || utils.IsNil(o.InstLogoUrl) {
		var ret string
		return ret
	}
	return *o.InstLogoUrl
}

// GetInstLogoUrlOk returns a tuple with the InstLogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetInstLogoUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.InstLogoUrl) {
		return nil, false
	}
	return o.InstLogoUrl, true
}

// HasInstLogoUrl returns a boolean if a field has been set.
func (o *AssetCardListItem) HasInstLogoUrl() bool {
	if o != nil && !utils.IsNil(o.InstLogoUrl) {
		return true
	}

	return false
}

// SetInstLogoUrl gets a reference to the given string and assigns it to the InstLogoUrl field.
func (o *AssetCardListItem) SetInstLogoUrl(v string) {
	o.InstLogoUrl = &v
}

// GetInstId returns the InstId field value
func (o *AssetCardListItem) GetInstId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetInstIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstId, true
}

// SetInstId sets field value
func (o *AssetCardListItem) SetInstId(v string) {
	o.InstId = v
}

// GetInstOfficialName returns the InstOfficialName field value
func (o *AssetCardListItem) GetInstOfficialName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstOfficialName
}

// GetInstOfficialNameOk returns a tuple with the InstOfficialName field value
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetInstOfficialNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstOfficialName, true
}

// SetInstOfficialName sets field value
func (o *AssetCardListItem) SetInstOfficialName(v string) {
	o.InstOfficialName = v
}

// GetExpiryYear returns the ExpiryYear field value
func (o *AssetCardListItem) GetExpiryYear() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryYear
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetExpiryYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryYear, true
}

// SetExpiryYear sets field value
func (o *AssetCardListItem) SetExpiryYear(v string) {
	o.ExpiryYear = v
}

// GetExpiryMonth returns the ExpiryMonth field value
func (o *AssetCardListItem) GetExpiryMonth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryMonth
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetExpiryMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryMonth, true
}

// SetExpiryMonth sets field value
func (o *AssetCardListItem) SetExpiryMonth(v string) {
	o.ExpiryMonth = v
}

// GetVerified returns the Verified field value
func (o *AssetCardListItem) GetVerified() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetVerifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *AssetCardListItem) SetVerified(v string) {
	o.Verified = v
}

// GetBindingId returns the BindingId field value if set, zero value otherwise.
func (o *AssetCardListItem) GetBindingId() string {
	if o == nil || utils.IsNil(o.BindingId) {
		var ret string
		return ret
	}
	return *o.BindingId
}

// GetBindingIdOk returns a tuple with the BindingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetBindingIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BindingId) {
		return nil, false
	}
	return o.BindingId, true
}

// HasBindingId returns a boolean if a field has been set.
func (o *AssetCardListItem) HasBindingId() bool {
	if o != nil && !utils.IsNil(o.BindingId) {
		return true
	}

	return false
}

// SetBindingId gets a reference to the given string and assigns it to the BindingId field.
func (o *AssetCardListItem) SetBindingId(v string) {
	o.BindingId = &v
}

// GetDefaultAsset returns the DefaultAsset field value if set, zero value otherwise.
func (o *AssetCardListItem) GetDefaultAsset() string {
	if o == nil || utils.IsNil(o.DefaultAsset) {
		var ret string
		return ret
	}
	return *o.DefaultAsset
}

// GetDefaultAssetOk returns a tuple with the DefaultAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetDefaultAssetOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DefaultAsset) {
		return nil, false
	}
	return o.DefaultAsset, true
}

// HasDefaultAsset returns a boolean if a field has been set.
func (o *AssetCardListItem) HasDefaultAsset() bool {
	if o != nil && !utils.IsNil(o.DefaultAsset) {
		return true
	}

	return false
}

// SetDefaultAsset gets a reference to the given string and assigns it to the DefaultAsset field.
func (o *AssetCardListItem) SetDefaultAsset(v string) {
	o.DefaultAsset = &v
}

// GetEnableStatus returns the EnableStatus field value if set, zero value otherwise.
func (o *AssetCardListItem) GetEnableStatus() string {
	if o == nil || utils.IsNil(o.EnableStatus) {
		var ret string
		return ret
	}
	return *o.EnableStatus
}

// GetEnableStatusOk returns a tuple with the EnableStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetEnableStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EnableStatus) {
		return nil, false
	}
	return o.EnableStatus, true
}

// HasEnableStatus returns a boolean if a field has been set.
func (o *AssetCardListItem) HasEnableStatus() bool {
	if o != nil && !utils.IsNil(o.EnableStatus) {
		return true
	}

	return false
}

// SetEnableStatus gets a reference to the given string and assigns it to the EnableStatus field.
func (o *AssetCardListItem) SetEnableStatus(v string) {
	o.EnableStatus = &v
}

// GetDirectDebit returns the DirectDebit field value if set, zero value otherwise.
func (o *AssetCardListItem) GetDirectDebit() string {
	if o == nil || utils.IsNil(o.DirectDebit) {
		var ret string
		return ret
	}
	return *o.DirectDebit
}

// GetDirectDebitOk returns a tuple with the DirectDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCardListItem) GetDirectDebitOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DirectDebit) {
		return nil, false
	}
	return o.DirectDebit, true
}

// HasDirectDebit returns a boolean if a field has been set.
func (o *AssetCardListItem) HasDirectDebit() bool {
	if o != nil && !utils.IsNil(o.DirectDebit) {
		return true
	}

	return false
}

// SetDirectDebit gets a reference to the given string and assigns it to the DirectDebit field.
func (o *AssetCardListItem) SetDirectDebit(v string) {
	o.DirectDebit = &v
}

func (o AssetCardListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetCardListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contactBizType"] = o.ContactBizType
	toSerialize["cardIndexNo"] = o.CardIndexNo
	toSerialize["cardNoLength"] = o.CardNoLength
	toSerialize["maskedCardNo"] = o.MaskedCardNo
	toSerialize["assetType"] = o.AssetType
	toSerialize["holderName"] = o.HolderName
	if !utils.IsNil(o.InstLogoUrl) {
		toSerialize["instLogoUrl"] = o.InstLogoUrl
	}
	toSerialize["instId"] = o.InstId
	toSerialize["instOfficialName"] = o.InstOfficialName
	toSerialize["expiryYear"] = o.ExpiryYear
	toSerialize["expiryMonth"] = o.ExpiryMonth
	toSerialize["verified"] = o.Verified
	if !utils.IsNil(o.BindingId) {
		toSerialize["bindingId"] = o.BindingId
	}
	if !utils.IsNil(o.DefaultAsset) {
		toSerialize["defaultAsset"] = o.DefaultAsset
	}
	if !utils.IsNil(o.EnableStatus) {
		toSerialize["enableStatus"] = o.EnableStatus
	}
	if !utils.IsNil(o.DirectDebit) {
		toSerialize["directDebit"] = o.DirectDebit
	}
	return toSerialize, nil
}

func (o *AssetCardListItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contactBizType",
		"cardIndexNo",
		"cardNoLength",
		"maskedCardNo",
		"assetType",
		"holderName",
		"instId",
		"instOfficialName",
		"expiryYear",
		"expiryMonth",
		"verified",
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

	varAssetCardListItem := _AssetCardListItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssetCardListItem)

	if err != nil {
		return err
	}

	*o = AssetCardListItem(varAssetCardListItem)

	return err
}

type NullableAssetCardListItem struct {
	value *AssetCardListItem
	isSet bool
}

func (v NullableAssetCardListItem) Get() *AssetCardListItem {
	return v.value
}

func (v *NullableAssetCardListItem) Set(val *AssetCardListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetCardListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetCardListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetCardListItem(val *AssetCardListItem) *NullableAssetCardListItem {
	return &NullableAssetCardListItem{value: val, isSet: true}
}

func (v NullableAssetCardListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetCardListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


