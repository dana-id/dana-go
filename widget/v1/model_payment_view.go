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

package widget

import (
	"encoding/json"
	"fmt"
	utils "github.com/dana-id/dana-go/v2/utils"
)

// checks if the PaymentView type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentView{}

// PaymentView struct for PaymentView
type PaymentView struct {
	// Cashier request identifier
	CashierRequestId *string `json:"cashierRequestId,omitempty"`
	// Paid time in format YYYY-MM-DDTHH:mm:ss+07:00 (Jakarta time)
	PaidTime *string `json:"paidTime,omitempty" validate:"regexp=^\\\\d{4}-\\\\d{2}-\\\\d{2}T\\\\d{2}:\\\\d{2}:\\\\d{2}\\\\+07:00$"`
	// Information of pay options
	PayOptionInfos []PayOptionInfo `json:"payOptionInfos,omitempty"`
	// Extend information of pay request
	PayRequestExtendInfo *string `json:"payRequestExtendInfo,omitempty"`
	// Additional extend information
	ExtendInfo *string `json:"extendInfo,omitempty"`
}

// NewPaymentView instantiates a new PaymentView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentView() *PaymentView {
	this := PaymentView{}
	return &this
}

// NewPaymentViewWithDefaults instantiates a new PaymentView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentViewWithDefaults() *PaymentView {
	this := PaymentView{}
	return &this
}

// GetCashierRequestId returns the CashierRequestId field value if set, zero value otherwise.
func (o *PaymentView) GetCashierRequestId() string {
	if o == nil || utils.IsNil(o.CashierRequestId) {
		var ret string
		return ret
	}
	return *o.CashierRequestId
}

// GetCashierRequestIdOk returns a tuple with the CashierRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentView) GetCashierRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CashierRequestId) {
		return nil, false
	}
	return o.CashierRequestId, true
}

// HasCashierRequestId returns a boolean if a field has been set.
func (o *PaymentView) HasCashierRequestId() bool {
	if o != nil && !utils.IsNil(o.CashierRequestId) {
		return true
	}

	return false
}

// SetCashierRequestId gets a reference to the given string and assigns it to the CashierRequestId field.

func (o *PaymentView) SetCashierRequestId(v string) error {
	// Validate string length constraints
	if !utils.IsNil(v) {
		if utils.Strlen(v) > 64 {
			return fmt.Errorf("invalid length for CashierRequestId, must be smaller than or equal to 64")
		}
	}
	o.CashierRequestId = &v
	return nil
}




// GetPaidTime returns the PaidTime field value if set, zero value otherwise.
func (o *PaymentView) GetPaidTime() string {
	if o == nil || utils.IsNil(o.PaidTime) {
		var ret string
		return ret
	}
	return *o.PaidTime
}

// GetPaidTimeOk returns a tuple with the PaidTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentView) GetPaidTimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaidTime) {
		return nil, false
	}
	return o.PaidTime, true
}

// HasPaidTime returns a boolean if a field has been set.
func (o *PaymentView) HasPaidTime() bool {
	if o != nil && !utils.IsNil(o.PaidTime) {
		return true
	}

	return false
}

// SetPaidTime gets a reference to the given string and assigns it to the PaidTime field.

func (o *PaymentView) SetPaidTime(v string) error {
	// Validate string length constraints
	if !utils.IsNil(v) {
		if utils.Strlen(v) > 25 {
			return fmt.Errorf("invalid length for PaidTime, must be smaller than or equal to 25")
		}
	}
	o.PaidTime = &v
	return nil
}




// GetPayOptionInfos returns the PayOptionInfos field value if set, zero value otherwise.
func (o *PaymentView) GetPayOptionInfos() []PayOptionInfo {
	if o == nil || utils.IsNil(o.PayOptionInfos) {
		var ret []PayOptionInfo
		return ret
	}
	return o.PayOptionInfos
}

// GetPayOptionInfosOk returns a tuple with the PayOptionInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentView) GetPayOptionInfosOk() ([]PayOptionInfo, bool) {
	if o == nil || utils.IsNil(o.PayOptionInfos) {
		return nil, false
	}
	return o.PayOptionInfos, true
}

// HasPayOptionInfos returns a boolean if a field has been set.
func (o *PaymentView) HasPayOptionInfos() bool {
	if o != nil && !utils.IsNil(o.PayOptionInfos) {
		return true
	}

	return false
}

// SetPayOptionInfos gets a reference to the given []PayOptionInfo and assigns it to the PayOptionInfos field.
func (o *PaymentView) SetPayOptionInfos(v []PayOptionInfo) {
	o.PayOptionInfos = v
}

// GetPayRequestExtendInfo returns the PayRequestExtendInfo field value if set, zero value otherwise.
func (o *PaymentView) GetPayRequestExtendInfo() string {
	if o == nil || utils.IsNil(o.PayRequestExtendInfo) {
		var ret string
		return ret
	}
	return *o.PayRequestExtendInfo
}

// GetPayRequestExtendInfoOk returns a tuple with the PayRequestExtendInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentView) GetPayRequestExtendInfoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PayRequestExtendInfo) {
		return nil, false
	}
	return o.PayRequestExtendInfo, true
}

// HasPayRequestExtendInfo returns a boolean if a field has been set.
func (o *PaymentView) HasPayRequestExtendInfo() bool {
	if o != nil && !utils.IsNil(o.PayRequestExtendInfo) {
		return true
	}

	return false
}

// SetPayRequestExtendInfo gets a reference to the given string and assigns it to the PayRequestExtendInfo field.

func (o *PaymentView) SetPayRequestExtendInfo(v string) error {
	// Validate string length constraints
	if !utils.IsNil(v) {
		if utils.Strlen(v) > 4096 {
			return fmt.Errorf("invalid length for PayRequestExtendInfo, must be smaller than or equal to 4096")
		}
	}
	o.PayRequestExtendInfo = &v
	return nil
}




// GetExtendInfo returns the ExtendInfo field value if set, zero value otherwise.
func (o *PaymentView) GetExtendInfo() string {
	if o == nil || utils.IsNil(o.ExtendInfo) {
		var ret string
		return ret
	}
	return *o.ExtendInfo
}

// GetExtendInfoOk returns a tuple with the ExtendInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentView) GetExtendInfoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExtendInfo) {
		return nil, false
	}
	return o.ExtendInfo, true
}

// HasExtendInfo returns a boolean if a field has been set.
func (o *PaymentView) HasExtendInfo() bool {
	if o != nil && !utils.IsNil(o.ExtendInfo) {
		return true
	}

	return false
}

// SetExtendInfo gets a reference to the given string and assigns it to the ExtendInfo field.

func (o *PaymentView) SetExtendInfo(v string) error {
	// Validate string length constraints
	if !utils.IsNil(v) {
		if utils.Strlen(v) > 4096 {
			return fmt.Errorf("invalid length for ExtendInfo, must be smaller than or equal to 4096")
		}
	}
	o.ExtendInfo = &v
	return nil
}




func (o PaymentView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CashierRequestId) {
		toSerialize["cashierRequestId"] = o.CashierRequestId
	}
	if !utils.IsNil(o.PaidTime) {
		toSerialize["paidTime"] = o.PaidTime
	}
	if !utils.IsNil(o.PayOptionInfos) {
		toSerialize["payOptionInfos"] = o.PayOptionInfos
	}
	if !utils.IsNil(o.PayRequestExtendInfo) {
		toSerialize["payRequestExtendInfo"] = o.PayRequestExtendInfo
	}
	if !utils.IsNil(o.ExtendInfo) {
		toSerialize["extendInfo"] = o.ExtendInfo
	}
	return toSerialize, nil
}

type NullablePaymentView struct {
	value *PaymentView
	isSet bool
}

func (v NullablePaymentView) Get() *PaymentView {
	return v.value
}

func (v *NullablePaymentView) Set(val *PaymentView) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentView) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentView(val *PaymentView) *NullablePaymentView {
	return &NullablePaymentView{value: val, isSet: true}
}

func (v NullablePaymentView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


