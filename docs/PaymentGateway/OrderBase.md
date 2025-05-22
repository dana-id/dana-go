# OrderBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderTitle** | **string** | Additional information of order title | 
**MerchantTransType** | Pointer to **string** | Additional information of merchant transaction type | [optional] 
**Buyer** | Pointer to [**Buyer**](Buyer.md) | Additional information of buyer | [optional] 
**Goods** | Pointer to [**[]Goods**](Goods.md) | Additional information of goods | [optional] 
**ShippingInfo** | Pointer to [**[]ShippingInfo**](ShippingInfo.md) | Additional information of shipping info | [optional] 
**ExtendInfo** | Pointer to **string** | Additional information of extend | [optional] 

## Methods

### NewOrderBase

`func NewOrderBase(orderTitle string, ) *OrderBase`

NewOrderBase instantiates a new OrderBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderBaseWithDefaults

`func NewOrderBaseWithDefaults() *OrderBase`

NewOrderBaseWithDefaults instantiates a new OrderBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderTitle

`func (o *OrderBase) GetOrderTitle() string`

GetOrderTitle returns the OrderTitle field if non-nil, zero value otherwise.

### GetOrderTitleOk

`func (o *OrderBase) GetOrderTitleOk() (*string, bool)`

GetOrderTitleOk returns a tuple with the OrderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTitle

`func (o *OrderBase) SetOrderTitle(v string)`

SetOrderTitle sets OrderTitle field to given value.


### GetMerchantTransType

`func (o *OrderBase) GetMerchantTransType() string`

GetMerchantTransType returns the MerchantTransType field if non-nil, zero value otherwise.

### GetMerchantTransTypeOk

`func (o *OrderBase) GetMerchantTransTypeOk() (*string, bool)`

GetMerchantTransTypeOk returns a tuple with the MerchantTransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTransType

`func (o *OrderBase) SetMerchantTransType(v string)`

SetMerchantTransType sets MerchantTransType field to given value.

### HasMerchantTransType

`func (o *OrderBase) HasMerchantTransType() bool`

HasMerchantTransType returns a boolean if a field has been set.

### GetBuyer

`func (o *OrderBase) GetBuyer() Buyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *OrderBase) GetBuyerOk() (*Buyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *OrderBase) SetBuyer(v Buyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *OrderBase) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetGoods

`func (o *OrderBase) GetGoods() []Goods`

GetGoods returns the Goods field if non-nil, zero value otherwise.

### GetGoodsOk

`func (o *OrderBase) GetGoodsOk() (*[]Goods, bool)`

GetGoodsOk returns a tuple with the Goods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoods

`func (o *OrderBase) SetGoods(v []Goods)`

SetGoods sets Goods field to given value.

### HasGoods

`func (o *OrderBase) HasGoods() bool`

HasGoods returns a boolean if a field has been set.

### GetShippingInfo

`func (o *OrderBase) GetShippingInfo() []ShippingInfo`

GetShippingInfo returns the ShippingInfo field if non-nil, zero value otherwise.

### GetShippingInfoOk

`func (o *OrderBase) GetShippingInfoOk() (*[]ShippingInfo, bool)`

GetShippingInfoOk returns a tuple with the ShippingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInfo

`func (o *OrderBase) SetShippingInfo(v []ShippingInfo)`

SetShippingInfo sets ShippingInfo field to given value.

### HasShippingInfo

`func (o *OrderBase) HasShippingInfo() bool`

HasShippingInfo returns a boolean if a field has been set.

### GetExtendInfo

`func (o *OrderBase) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *OrderBase) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *OrderBase) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *OrderBase) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


