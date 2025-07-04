/*
Payment Gateway API

API for doing operations in DANA Payment Gateway (Gapura)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payment_gateway

import (
	"encoding/json"
	"bytes"
	"fmt"
	utils "github.com/dana-id/dana-go/utils"
)

// checks if the OrderBase type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OrderBase{}

// OrderBase struct for OrderBase
type OrderBase struct {
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
}

type _OrderBase OrderBase

// NewOrderBase instantiates a new OrderBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderBase(orderTitle string) *OrderBase {
	this := OrderBase{}
	this.OrderTitle = orderTitle
	return &this
}

// NewOrderBaseWithDefaults instantiates a new OrderBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderBaseWithDefaults() *OrderBase {
	this := OrderBase{}
	return &this
}

// GetOrderTitle returns the OrderTitle field value
func (o *OrderBase) GetOrderTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderTitle
}

// GetOrderTitleOk returns a tuple with the OrderTitle field value
// and a boolean to check if the value has been set.
func (o *OrderBase) GetOrderTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderTitle, true
}

// SetOrderTitle sets field value
func (o *OrderBase) SetOrderTitle(v string) {
	o.OrderTitle = v
}

// GetMerchantTransType returns the MerchantTransType field value if set, zero value otherwise.
func (o *OrderBase) GetMerchantTransType() string {
	if o == nil || utils.IsNil(o.MerchantTransType) {
		var ret string
		return ret
	}
	return *o.MerchantTransType
}

// GetMerchantTransTypeOk returns a tuple with the MerchantTransType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBase) GetMerchantTransTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MerchantTransType) {
		return nil, false
	}
	return o.MerchantTransType, true
}

// HasMerchantTransType returns a boolean if a field has been set.
func (o *OrderBase) HasMerchantTransType() bool {
	if o != nil && !utils.IsNil(o.MerchantTransType) {
		return true
	}

	return false
}

// SetMerchantTransType gets a reference to the given string and assigns it to the MerchantTransType field.
func (o *OrderBase) SetMerchantTransType(v string) {
	o.MerchantTransType = &v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *OrderBase) GetBuyer() Buyer {
	if o == nil || utils.IsNil(o.Buyer) {
		var ret Buyer
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBase) GetBuyerOk() (*Buyer, bool) {
	if o == nil || utils.IsNil(o.Buyer) {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *OrderBase) HasBuyer() bool {
	if o != nil && !utils.IsNil(o.Buyer) {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given Buyer and assigns it to the Buyer field.
func (o *OrderBase) SetBuyer(v Buyer) {
	o.Buyer = &v
}

// GetGoods returns the Goods field value if set, zero value otherwise.
func (o *OrderBase) GetGoods() []Goods {
	if o == nil || utils.IsNil(o.Goods) {
		var ret []Goods
		return ret
	}
	return o.Goods
}

// GetGoodsOk returns a tuple with the Goods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBase) GetGoodsOk() ([]Goods, bool) {
	if o == nil || utils.IsNil(o.Goods) {
		return nil, false
	}
	return o.Goods, true
}

// HasGoods returns a boolean if a field has been set.
func (o *OrderBase) HasGoods() bool {
	if o != nil && !utils.IsNil(o.Goods) {
		return true
	}

	return false
}

// SetGoods gets a reference to the given []Goods and assigns it to the Goods field.
func (o *OrderBase) SetGoods(v []Goods) {
	o.Goods = v
}

// GetShippingInfo returns the ShippingInfo field value if set, zero value otherwise.
func (o *OrderBase) GetShippingInfo() []ShippingInfo {
	if o == nil || utils.IsNil(o.ShippingInfo) {
		var ret []ShippingInfo
		return ret
	}
	return o.ShippingInfo
}

// GetShippingInfoOk returns a tuple with the ShippingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBase) GetShippingInfoOk() ([]ShippingInfo, bool) {
	if o == nil || utils.IsNil(o.ShippingInfo) {
		return nil, false
	}
	return o.ShippingInfo, true
}

// HasShippingInfo returns a boolean if a field has been set.
func (o *OrderBase) HasShippingInfo() bool {
	if o != nil && !utils.IsNil(o.ShippingInfo) {
		return true
	}

	return false
}

// SetShippingInfo gets a reference to the given []ShippingInfo and assigns it to the ShippingInfo field.
func (o *OrderBase) SetShippingInfo(v []ShippingInfo) {
	o.ShippingInfo = v
}

// GetExtendInfo returns the ExtendInfo field value if set, zero value otherwise.
func (o *OrderBase) GetExtendInfo() string {
	if o == nil || utils.IsNil(o.ExtendInfo) {
		var ret string
		return ret
	}
	return *o.ExtendInfo
}

// GetExtendInfoOk returns a tuple with the ExtendInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBase) GetExtendInfoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExtendInfo) {
		return nil, false
	}
	return o.ExtendInfo, true
}

// HasExtendInfo returns a boolean if a field has been set.
func (o *OrderBase) HasExtendInfo() bool {
	if o != nil && !utils.IsNil(o.ExtendInfo) {
		return true
	}

	return false
}

// SetExtendInfo gets a reference to the given string and assigns it to the ExtendInfo field.
func (o *OrderBase) SetExtendInfo(v string) {
	o.ExtendInfo = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *OrderBase) GetCreatedTime() string {
	if o == nil || utils.IsNil(o.CreatedTime) {
		var ret string
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBase) GetCreatedTimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *OrderBase) HasCreatedTime() bool {
	if o != nil && !utils.IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given string and assigns it to the CreatedTime field.
func (o *OrderBase) SetCreatedTime(v string) {
	o.CreatedTime = &v
}

// GetOrderMemo returns the OrderMemo field value if set, zero value otherwise.
func (o *OrderBase) GetOrderMemo() string {
	if o == nil || utils.IsNil(o.OrderMemo) {
		var ret string
		return ret
	}
	return *o.OrderMemo
}

// GetOrderMemoOk returns a tuple with the OrderMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBase) GetOrderMemoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OrderMemo) {
		return nil, false
	}
	return o.OrderMemo, true
}

// HasOrderMemo returns a boolean if a field has been set.
func (o *OrderBase) HasOrderMemo() bool {
	if o != nil && !utils.IsNil(o.OrderMemo) {
		return true
	}

	return false
}

// SetOrderMemo gets a reference to the given string and assigns it to the OrderMemo field.
func (o *OrderBase) SetOrderMemo(v string) {
	o.OrderMemo = &v
}

func (o OrderBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderBase) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

func (o *OrderBase) UnmarshalJSON(data []byte) (err error) {
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

	varOrderBase := _OrderBase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderBase)

	if err != nil {
		return err
	}

	*o = OrderBase(varOrderBase)

	return err
}

type NullableOrderBase struct {
	value *OrderBase
	isSet bool
}

func (v NullableOrderBase) Get() *OrderBase {
	return v.value
}

func (v *NullableOrderBase) Set(val *OrderBase) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderBase) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderBase(val *OrderBase) *NullableOrderBase {
	return &NullableOrderBase{value: val, isSet: true}
}

func (v NullableOrderBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


