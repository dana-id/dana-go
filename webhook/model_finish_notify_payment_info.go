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

// checks if the FinishNotifyPaymentInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FinishNotifyPaymentInfo{}

// FinishNotifyPaymentInfo struct for FinishNotifyPaymentInfo
type FinishNotifyPaymentInfo struct {
	// Cashier request identifier
	CashierRequestId string `json:"cashierRequestId"`
	// Information of paid time, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time)
	PaidTime string `json:"paidTime" validate:"regexp=^\\\\d{4}-\\\\d{2}-\\\\d{2}T\\\\d{2}:\\\\d{2}:\\\\d{2}\\\\+07:00$"`
	// Information of pay option. Refer to payOptionInfos for the detailed
	PayOptionInfos []PayOptionInfo `json:"payOptionInfos"`
	// Extend information of pay request
	PayRequestExtendInfo *string `json:"payRequestExtendInfo,omitempty"`
	// Extend information
	ExtendInfo *string `json:"extendInfo,omitempty"`
}

type _FinishNotifyPaymentInfo FinishNotifyPaymentInfo

// NewFinishNotifyPaymentInfo instantiates a new FinishNotifyPaymentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinishNotifyPaymentInfo(cashierRequestId string, paidTime string, payOptionInfos []PayOptionInfo) *FinishNotifyPaymentInfo {
	this := FinishNotifyPaymentInfo{}
	this.CashierRequestId = cashierRequestId
	this.PaidTime = paidTime
	this.PayOptionInfos = payOptionInfos
	return &this
}

// NewFinishNotifyPaymentInfoWithDefaults instantiates a new FinishNotifyPaymentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinishNotifyPaymentInfoWithDefaults() *FinishNotifyPaymentInfo {
	this := FinishNotifyPaymentInfo{}
	return &this
}

// GetCashierRequestId returns the CashierRequestId field value
func (o *FinishNotifyPaymentInfo) GetCashierRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CashierRequestId
}

// GetCashierRequestIdOk returns a tuple with the CashierRequestId field value
// and a boolean to check if the value has been set.
func (o *FinishNotifyPaymentInfo) GetCashierRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CashierRequestId, true
}

// SetCashierRequestId sets field value
func (o *FinishNotifyPaymentInfo) SetCashierRequestId(v string) {
	o.CashierRequestId = v
}

// GetPaidTime returns the PaidTime field value
func (o *FinishNotifyPaymentInfo) GetPaidTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaidTime
}

// GetPaidTimeOk returns a tuple with the PaidTime field value
// and a boolean to check if the value has been set.
func (o *FinishNotifyPaymentInfo) GetPaidTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaidTime, true
}

// SetPaidTime sets field value
func (o *FinishNotifyPaymentInfo) SetPaidTime(v string) {
	o.PaidTime = v
}

// GetPayOptionInfos returns the PayOptionInfos field value
func (o *FinishNotifyPaymentInfo) GetPayOptionInfos() []PayOptionInfo {
	if o == nil {
		var ret []PayOptionInfo
		return ret
	}

	return o.PayOptionInfos
}

// GetPayOptionInfosOk returns a tuple with the PayOptionInfos field value
// and a boolean to check if the value has been set.
func (o *FinishNotifyPaymentInfo) GetPayOptionInfosOk() ([]PayOptionInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayOptionInfos, true
}

// SetPayOptionInfos sets field value
func (o *FinishNotifyPaymentInfo) SetPayOptionInfos(v []PayOptionInfo) {
	o.PayOptionInfos = v
}

// GetPayRequestExtendInfo returns the PayRequestExtendInfo field value if set, zero value otherwise.
func (o *FinishNotifyPaymentInfo) GetPayRequestExtendInfo() string {
	if o == nil || utils.IsNil(o.PayRequestExtendInfo) {
		var ret string
		return ret
	}
	return *o.PayRequestExtendInfo
}

// GetPayRequestExtendInfoOk returns a tuple with the PayRequestExtendInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinishNotifyPaymentInfo) GetPayRequestExtendInfoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PayRequestExtendInfo) {
		return nil, false
	}
	return o.PayRequestExtendInfo, true
}

// HasPayRequestExtendInfo returns a boolean if a field has been set.
func (o *FinishNotifyPaymentInfo) HasPayRequestExtendInfo() bool {
	if o != nil && !utils.IsNil(o.PayRequestExtendInfo) {
		return true
	}

	return false
}

// SetPayRequestExtendInfo gets a reference to the given string and assigns it to the PayRequestExtendInfo field.
func (o *FinishNotifyPaymentInfo) SetPayRequestExtendInfo(v string) {
	o.PayRequestExtendInfo = &v
}

// GetExtendInfo returns the ExtendInfo field value if set, zero value otherwise.
func (o *FinishNotifyPaymentInfo) GetExtendInfo() string {
	if o == nil || utils.IsNil(o.ExtendInfo) {
		var ret string
		return ret
	}
	return *o.ExtendInfo
}

// GetExtendInfoOk returns a tuple with the ExtendInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinishNotifyPaymentInfo) GetExtendInfoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExtendInfo) {
		return nil, false
	}
	return o.ExtendInfo, true
}

// HasExtendInfo returns a boolean if a field has been set.
func (o *FinishNotifyPaymentInfo) HasExtendInfo() bool {
	if o != nil && !utils.IsNil(o.ExtendInfo) {
		return true
	}

	return false
}

// SetExtendInfo gets a reference to the given string and assigns it to the ExtendInfo field.
func (o *FinishNotifyPaymentInfo) SetExtendInfo(v string) {
	o.ExtendInfo = &v
}

func (o FinishNotifyPaymentInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinishNotifyPaymentInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cashierRequestId"] = o.CashierRequestId
	toSerialize["paidTime"] = o.PaidTime
	toSerialize["payOptionInfos"] = o.PayOptionInfos
	if !utils.IsNil(o.PayRequestExtendInfo) {
		toSerialize["payRequestExtendInfo"] = o.PayRequestExtendInfo
	}
	if !utils.IsNil(o.ExtendInfo) {
		toSerialize["extendInfo"] = o.ExtendInfo
	}
	return toSerialize, nil
}

func (o *FinishNotifyPaymentInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cashierRequestId",
		"paidTime",
		"payOptionInfos",
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

	varFinishNotifyPaymentInfo := _FinishNotifyPaymentInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFinishNotifyPaymentInfo)

	if err != nil {
		return err
	}

	*o = FinishNotifyPaymentInfo(varFinishNotifyPaymentInfo)

	return err
}

type NullableFinishNotifyPaymentInfo struct {
	value *FinishNotifyPaymentInfo
	isSet bool
}

func (v NullableFinishNotifyPaymentInfo) Get() *FinishNotifyPaymentInfo {
	return v.value
}

func (v *NullableFinishNotifyPaymentInfo) Set(val *FinishNotifyPaymentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFinishNotifyPaymentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFinishNotifyPaymentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinishNotifyPaymentInfo(val *FinishNotifyPaymentInfo) *NullableFinishNotifyPaymentInfo {
	return &NullableFinishNotifyPaymentInfo{value: val, isSet: true}
}

func (v NullableFinishNotifyPaymentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinishNotifyPaymentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


