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

// checks if the FinishNotifyRequestAdditionalInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FinishNotifyRequestAdditionalInfo{}

// FinishNotifyRequestAdditionalInfo struct for FinishNotifyRequestAdditionalInfo
type FinishNotifyRequestAdditionalInfo struct {
	// Additional information of payment
	PaymentInfo *FinishNotifyPaymentInfo `json:"paymentInfo,omitempty"`
	// Additional information of shop
	ShopInfo *ShopInfo `json:"shopInfo,omitempty"`
	// Additional information of extend (As a JSON string)
	ExtendInfo *string `json:"extendInfo,omitempty"`
	// Additional information of closed reason. Required if order is closed
	ExtendInfoClosedReason *string `json:"extendInfo.closedReason,omitempty"`
	// Additional information of paid time
	PaidTime *string `json:"paidTime,omitempty" validate:"omitempty,regexp=^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\+07:00$"`
}

// NewFinishNotifyRequestAdditionalInfo instantiates a new FinishNotifyRequestAdditionalInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinishNotifyRequestAdditionalInfo() *FinishNotifyRequestAdditionalInfo {
	this := FinishNotifyRequestAdditionalInfo{}
	return &this
}

// NewFinishNotifyRequestAdditionalInfoWithDefaults instantiates a new FinishNotifyRequestAdditionalInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinishNotifyRequestAdditionalInfoWithDefaults() *FinishNotifyRequestAdditionalInfo {
	this := FinishNotifyRequestAdditionalInfo{}
	return &this
}

// GetPaymentInfo returns the PaymentInfo field value if set, zero value otherwise.
func (o *FinishNotifyRequestAdditionalInfo) GetPaymentInfo() FinishNotifyPaymentInfo {
	if o == nil || utils.IsNil(o.PaymentInfo) {
		var ret FinishNotifyPaymentInfo
		return ret
	}
	return *o.PaymentInfo
}

// GetPaymentInfoOk returns a tuple with the PaymentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequestAdditionalInfo) GetPaymentInfoOk() (*FinishNotifyPaymentInfo, bool) {
	if o == nil || utils.IsNil(o.PaymentInfo) {
		return nil, false
	}
	return o.PaymentInfo, true
}

// HasPaymentInfo returns a boolean if a field has been set.
func (o *FinishNotifyRequestAdditionalInfo) HasPaymentInfo() bool {
	if o != nil && !utils.IsNil(o.PaymentInfo) {
		return true
	}

	return false
}

// SetPaymentInfo gets a reference to the given FinishNotifyPaymentInfo and assigns it to the PaymentInfo field.
func (o *FinishNotifyRequestAdditionalInfo) SetPaymentInfo(v FinishNotifyPaymentInfo) {
	o.PaymentInfo = &v
}

// GetShopInfo returns the ShopInfo field value if set, zero value otherwise.
func (o *FinishNotifyRequestAdditionalInfo) GetShopInfo() ShopInfo {
	if o == nil || utils.IsNil(o.ShopInfo) {
		var ret ShopInfo
		return ret
	}
	return *o.ShopInfo
}

// GetShopInfoOk returns a tuple with the ShopInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequestAdditionalInfo) GetShopInfoOk() (*ShopInfo, bool) {
	if o == nil || utils.IsNil(o.ShopInfo) {
		return nil, false
	}
	return o.ShopInfo, true
}

// HasShopInfo returns a boolean if a field has been set.
func (o *FinishNotifyRequestAdditionalInfo) HasShopInfo() bool {
	if o != nil && !utils.IsNil(o.ShopInfo) {
		return true
	}

	return false
}

// SetShopInfo gets a reference to the given ShopInfo and assigns it to the ShopInfo field.
func (o *FinishNotifyRequestAdditionalInfo) SetShopInfo(v ShopInfo) {
	o.ShopInfo = &v
}

// GetExtendInfo returns the ExtendInfo field value if set, zero value otherwise.
func (o *FinishNotifyRequestAdditionalInfo) GetExtendInfo() string {
	if o == nil || utils.IsNil(o.ExtendInfo) {
		var ret string
		return ret
	}
	return *o.ExtendInfo
}

// GetExtendInfoOk returns a tuple with the ExtendInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequestAdditionalInfo) GetExtendInfoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExtendInfo) {
		return nil, false
	}
	return o.ExtendInfo, true
}

// HasExtendInfo returns a boolean if a field has been set.
func (o *FinishNotifyRequestAdditionalInfo) HasExtendInfo() bool {
	if o != nil && !utils.IsNil(o.ExtendInfo) {
		return true
	}

	return false
}

// SetExtendInfo gets a reference to the given string and assigns it to the ExtendInfo field.
func (o *FinishNotifyRequestAdditionalInfo) SetExtendInfo(v string) {
	o.ExtendInfo = &v
}

// GetExtendInfoClosedReason returns the ExtendInfoClosedReason field value if set, zero value otherwise.
func (o *FinishNotifyRequestAdditionalInfo) GetExtendInfoClosedReason() string {
	if o == nil || utils.IsNil(o.ExtendInfoClosedReason) {
		var ret string
		return ret
	}
	return *o.ExtendInfoClosedReason
}

// GetExtendInfoClosedReasonOk returns a tuple with the ExtendInfoClosedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequestAdditionalInfo) GetExtendInfoClosedReasonOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExtendInfoClosedReason) {
		return nil, false
	}
	return o.ExtendInfoClosedReason, true
}

// HasExtendInfoClosedReason returns a boolean if a field has been set.
func (o *FinishNotifyRequestAdditionalInfo) HasExtendInfoClosedReason() bool {
	if o != nil && !utils.IsNil(o.ExtendInfoClosedReason) {
		return true
	}

	return false
}

// SetExtendInfoClosedReason gets a reference to the given string and assigns it to the ExtendInfoClosedReason field.
func (o *FinishNotifyRequestAdditionalInfo) SetExtendInfoClosedReason(v string) {
	o.ExtendInfoClosedReason = &v
}

// GetPaidTime returns the PaidTime field value if set, zero value otherwise.
func (o *FinishNotifyRequestAdditionalInfo) GetPaidTime() string {
	if o == nil || utils.IsNil(o.PaidTime) {
		var ret string
		return ret
	}
	return *o.PaidTime
}

// GetPaidTimeOk returns a tuple with the PaidTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequestAdditionalInfo) GetPaidTimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaidTime) {
		return nil, false
	}
	return o.PaidTime, true
}

// HasPaidTime returns a boolean if a field has been set.
func (o *FinishNotifyRequestAdditionalInfo) HasPaidTime() bool {
	if o != nil && !utils.IsNil(o.PaidTime) {
		return true
	}

	return false
}

// SetPaidTime gets a reference to the given string and assigns it to the PaidTime field.
func (o *FinishNotifyRequestAdditionalInfo) SetPaidTime(v string) {
	o.PaidTime = &v
}

func (o FinishNotifyRequestAdditionalInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinishNotifyRequestAdditionalInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PaymentInfo) {
		toSerialize["paymentInfo"] = o.PaymentInfo
	}
	if !utils.IsNil(o.ShopInfo) {
		toSerialize["shopInfo"] = o.ShopInfo
	}
	if !utils.IsNil(o.ExtendInfo) {
		toSerialize["extendInfo"] = o.ExtendInfo
	}
	if !utils.IsNil(o.ExtendInfoClosedReason) {
		toSerialize["extendInfo.closedReason"] = o.ExtendInfoClosedReason
	}
	if !utils.IsNil(o.PaidTime) {
		toSerialize["paidTime"] = o.PaidTime
	}
	return toSerialize, nil
}

type NullableFinishNotifyRequestAdditionalInfo struct {
	value *FinishNotifyRequestAdditionalInfo
	isSet bool
}

func (v NullableFinishNotifyRequestAdditionalInfo) Get() *FinishNotifyRequestAdditionalInfo {
	return v.value
}

func (v *NullableFinishNotifyRequestAdditionalInfo) Set(val *FinishNotifyRequestAdditionalInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFinishNotifyRequestAdditionalInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFinishNotifyRequestAdditionalInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinishNotifyRequestAdditionalInfo(val *FinishNotifyRequestAdditionalInfo) *NullableFinishNotifyRequestAdditionalInfo {
	return &NullableFinishNotifyRequestAdditionalInfo{value: val, isSet: true}
}

func (v NullableFinishNotifyRequestAdditionalInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinishNotifyRequestAdditionalInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
