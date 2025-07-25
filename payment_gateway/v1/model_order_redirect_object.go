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

package payment_gateway

/*
Payment Gateway API

API for doing operations in DANA Payment Gateway (Gapura)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.


import (
	"encoding/json"
	"bytes"
	"fmt"
	utils "github.com/dana-id/dana-go/utils"
)

// checks if the OrderRedirectObject type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OrderRedirectObject{}

// OrderRedirectObject struct for OrderRedirectObject
type OrderRedirectObject struct {
	// Additional information of order title
	OrderTitle string `json:"orderTitle"`
	// Additional information of merchant transaction type
	MerchantTransType *string `json:"merchantTransType,omitempty"`
	// Additional information of buyer
	Buyer *Buyer `json:"buyer,omitempty"`
	// Additional information of goods
	Goods []Goods `json:"goods,omitempty"`
	// Additional information of shipping info
	ShippingInfo []ShippingInfo `json:"shippingInfo,omitempty"`
	// Additional information of extend
	ExtendInfo *string `json:"extendInfo,omitempty"`
	// Additional information of created time
	CreatedTime *string `json:"createdTime,omitempty"`
	// Additional information of order
	OrderMemo *string `json:"orderMemo,omitempty"`
	// For Payment Gateway Drop-in scenario, need to fill it as REDIRECT
	Scenario *string `json:"scenario,omitempty"`
}

type _OrderRedirectObject OrderRedirectObject

// NewOrderRedirectObject instantiates a new OrderRedirectObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderRedirectObject(orderTitle string) *OrderRedirectObject {
	this := OrderRedirectObject{}
	this.OrderTitle = orderTitle
	return &this
}

// NewOrderRedirectObjectWithDefaults instantiates a new OrderRedirectObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderRedirectObjectWithDefaults() *OrderRedirectObject {
	this := OrderRedirectObject{}
	return &this
}

// GetOrderTitle returns the OrderTitle field value
func (o *OrderRedirectObject) GetOrderTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderTitle
}

// GetOrderTitleOk returns a tuple with the OrderTitle field value
// and a boolean to check if the value has been set.
func (o *OrderRedirectObject) GetOrderTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderTitle, true
}

// SetOrderTitle sets field value
func (o *OrderRedirectObject) SetOrderTitle(v string) {
	o.OrderTitle = v
}

// GetMerchantTransType returns the MerchantTransType field value if set, zero value otherwise.
func (o *OrderRedirectObject) GetMerchantTransType() string {
	if o == nil || utils.IsNil(o.MerchantTransType) {
		var ret string
		return ret
	}
	return *o.MerchantTransType
}

// GetMerchantTransTypeOk returns a tuple with the MerchantTransType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRedirectObject) GetMerchantTransTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MerchantTransType) {
		return nil, false
	}
	return o.MerchantTransType, true
}

// HasMerchantTransType returns a boolean if a field has been set.
func (o *OrderRedirectObject) HasMerchantTransType() bool {
	if o != nil && !utils.IsNil(o.MerchantTransType) {
		return true
	}

	return false
}

// SetMerchantTransType gets a reference to the given string and assigns it to the MerchantTransType field.
func (o *OrderRedirectObject) SetMerchantTransType(v string) {
	o.MerchantTransType = &v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *OrderRedirectObject) GetBuyer() Buyer {
	if o == nil || utils.IsNil(o.Buyer) {
		var ret Buyer
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRedirectObject) GetBuyerOk() (*Buyer, bool) {
	if o == nil || utils.IsNil(o.Buyer) {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *OrderRedirectObject) HasBuyer() bool {
	if o != nil && !utils.IsNil(o.Buyer) {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given Buyer and assigns it to the Buyer field.
func (o *OrderRedirectObject) SetBuyer(v Buyer) {
	o.Buyer = &v
}

// GetGoods returns the Goods field value if set, zero value otherwise.
func (o *OrderRedirectObject) GetGoods() []Goods {
	if o == nil || utils.IsNil(o.Goods) {
		var ret []Goods
		return ret
	}
	return o.Goods
}

// GetGoodsOk returns a tuple with the Goods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRedirectObject) GetGoodsOk() ([]Goods, bool) {
	if o == nil || utils.IsNil(o.Goods) {
		return nil, false
	}
	return o.Goods, true
}

// HasGoods returns a boolean if a field has been set.
func (o *OrderRedirectObject) HasGoods() bool {
	if o != nil && !utils.IsNil(o.Goods) {
		return true
	}

	return false
}

// SetGoods gets a reference to the given []Goods and assigns it to the Goods field.
func (o *OrderRedirectObject) SetGoods(v []Goods) {
	o.Goods = v
}

// GetShippingInfo returns the ShippingInfo field value if set, zero value otherwise.
func (o *OrderRedirectObject) GetShippingInfo() []ShippingInfo {
	if o == nil || utils.IsNil(o.ShippingInfo) {
		var ret []ShippingInfo
		return ret
	}
	return o.ShippingInfo
}

// GetShippingInfoOk returns a tuple with the ShippingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRedirectObject) GetShippingInfoOk() ([]ShippingInfo, bool) {
	if o == nil || utils.IsNil(o.ShippingInfo) {
		return nil, false
	}
	return o.ShippingInfo, true
}

// HasShippingInfo returns a boolean if a field has been set.
func (o *OrderRedirectObject) HasShippingInfo() bool {
	if o != nil && !utils.IsNil(o.ShippingInfo) {
		return true
	}

	return false
}

// SetShippingInfo gets a reference to the given []ShippingInfo and assigns it to the ShippingInfo field.
func (o *OrderRedirectObject) SetShippingInfo(v []ShippingInfo) {
	o.ShippingInfo = v
}

// GetExtendInfo returns the ExtendInfo field value if set, zero value otherwise.
func (o *OrderRedirectObject) GetExtendInfo() string {
	if o == nil || utils.IsNil(o.ExtendInfo) {
		var ret string
		return ret
	}
	return *o.ExtendInfo
}

// GetExtendInfoOk returns a tuple with the ExtendInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRedirectObject) GetExtendInfoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExtendInfo) {
		return nil, false
	}
	return o.ExtendInfo, true
}

// HasExtendInfo returns a boolean if a field has been set.
func (o *OrderRedirectObject) HasExtendInfo() bool {
	if o != nil && !utils.IsNil(o.ExtendInfo) {
		return true
	}

	return false
}

// SetExtendInfo gets a reference to the given string and assigns it to the ExtendInfo field.
func (o *OrderRedirectObject) SetExtendInfo(v string) {
	o.ExtendInfo = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *OrderRedirectObject) GetCreatedTime() string {
	if o == nil || utils.IsNil(o.CreatedTime) {
		var ret string
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRedirectObject) GetCreatedTimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *OrderRedirectObject) HasCreatedTime() bool {
	if o != nil && !utils.IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given string and assigns it to the CreatedTime field.
func (o *OrderRedirectObject) SetCreatedTime(v string) {
	o.CreatedTime = &v
}

// GetOrderMemo returns the OrderMemo field value if set, zero value otherwise.
func (o *OrderRedirectObject) GetOrderMemo() string {
	if o == nil || utils.IsNil(o.OrderMemo) {
		var ret string
		return ret
	}
	return *o.OrderMemo
}

// GetOrderMemoOk returns a tuple with the OrderMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRedirectObject) GetOrderMemoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OrderMemo) {
		return nil, false
	}
	return o.OrderMemo, true
}

// HasOrderMemo returns a boolean if a field has been set.
func (o *OrderRedirectObject) HasOrderMemo() bool {
	if o != nil && !utils.IsNil(o.OrderMemo) {
		return true
	}

	return false
}

// SetOrderMemo gets a reference to the given string and assigns it to the OrderMemo field.
func (o *OrderRedirectObject) SetOrderMemo(v string) {
	o.OrderMemo = &v
}

// GetScenario returns the Scenario field value if set, zero value otherwise.
func (o *OrderRedirectObject) GetScenario() string {
	if o == nil || utils.IsNil(o.Scenario) {
		var ret string
		return ret
	}
	return *o.Scenario
}

// GetScenarioOk returns a tuple with the Scenario field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRedirectObject) GetScenarioOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Scenario) {
		return nil, false
	}
	return o.Scenario, true
}

// HasScenario returns a boolean if a field has been set.
func (o *OrderRedirectObject) HasScenario() bool {
	if o != nil && !utils.IsNil(o.Scenario) {
		return true
	}

	return false
}

// SetScenario gets a reference to the given string and assigns it to the Scenario field.
func (o *OrderRedirectObject) SetScenario(v string) {
	o.Scenario = &v
}

func (o OrderRedirectObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderRedirectObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderTitle"] = o.OrderTitle
	if !utils.IsNil(o.MerchantTransType) {
		toSerialize["merchantTransType"] = o.MerchantTransType
	}
	if !utils.IsNil(o.Buyer) {
		toSerialize["buyer"] = o.Buyer
	}
	if !utils.IsNil(o.Goods) {
		toSerialize["goods"] = o.Goods
	}
	if !utils.IsNil(o.ShippingInfo) {
		toSerialize["shippingInfo"] = o.ShippingInfo
	}
	if !utils.IsNil(o.ExtendInfo) {
		toSerialize["extendInfo"] = o.ExtendInfo
	}
	if !utils.IsNil(o.CreatedTime) {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if !utils.IsNil(o.OrderMemo) {
		toSerialize["orderMemo"] = o.OrderMemo
	}
	if !utils.IsNil(o.Scenario) {
		toSerialize["scenario"] = o.Scenario
	}
	return toSerialize, nil
}

func (o *OrderRedirectObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orderTitle",
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

	varOrderRedirectObject := _OrderRedirectObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderRedirectObject)

	if err != nil {
		return err
	}

	*o = OrderRedirectObject(varOrderRedirectObject)

	return err
}

type NullableOrderRedirectObject struct {
	value *OrderRedirectObject
	isSet bool
}

func (v NullableOrderRedirectObject) Get() *OrderRedirectObject {
	return v.value
}

func (v *NullableOrderRedirectObject) Set(val *OrderRedirectObject) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderRedirectObject) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderRedirectObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderRedirectObject(val *OrderRedirectObject) *NullableOrderRedirectObject {
	return &NullableOrderRedirectObject{value: val, isSet: true}
}

func (v NullableOrderRedirectObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderRedirectObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


