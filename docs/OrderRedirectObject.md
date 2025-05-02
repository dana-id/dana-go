# OrderRedirectObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderTitle** | **string** |  | 
**MerchantTransType** | Pointer to **string** |  | [optional] 
**Buyer** | Pointer to [**Buyer**](Buyer.md) |  | [optional] 
**Goods** | Pointer to [**[]Goods**](Goods.md) |  | [optional] 
**ShippingInfo** | Pointer to [**[]ShippingInfo**](ShippingInfo.md) |  | [optional] 
**ExtendInfo** | Pointer to **string** |  | [optional] 
**Scenario** | Pointer to **string** |  | [optional] 

## Methods

### NewOrderRedirectObject

`func NewOrderRedirectObject(orderTitle string, ) *OrderRedirectObject`

NewOrderRedirectObject instantiates a new OrderRedirectObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderRedirectObjectWithDefaults

`func NewOrderRedirectObjectWithDefaults() *OrderRedirectObject`

NewOrderRedirectObjectWithDefaults instantiates a new OrderRedirectObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderTitle

`func (o *OrderRedirectObject) GetOrderTitle() string`

GetOrderTitle returns the OrderTitle field if non-nil, zero value otherwise.

### GetOrderTitleOk

`func (o *OrderRedirectObject) GetOrderTitleOk() (*string, bool)`

GetOrderTitleOk returns a tuple with the OrderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTitle

`func (o *OrderRedirectObject) SetOrderTitle(v string)`

SetOrderTitle sets OrderTitle field to given value.


### GetMerchantTransType

`func (o *OrderRedirectObject) GetMerchantTransType() string`

GetMerchantTransType returns the MerchantTransType field if non-nil, zero value otherwise.

### GetMerchantTransTypeOk

`func (o *OrderRedirectObject) GetMerchantTransTypeOk() (*string, bool)`

GetMerchantTransTypeOk returns a tuple with the MerchantTransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTransType

`func (o *OrderRedirectObject) SetMerchantTransType(v string)`

SetMerchantTransType sets MerchantTransType field to given value.

### HasMerchantTransType

`func (o *OrderRedirectObject) HasMerchantTransType() bool`

HasMerchantTransType returns a boolean if a field has been set.

### GetBuyer

`func (o *OrderRedirectObject) GetBuyer() Buyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *OrderRedirectObject) GetBuyerOk() (*Buyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *OrderRedirectObject) SetBuyer(v Buyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *OrderRedirectObject) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetGoods

`func (o *OrderRedirectObject) GetGoods() []Goods`

GetGoods returns the Goods field if non-nil, zero value otherwise.

### GetGoodsOk

`func (o *OrderRedirectObject) GetGoodsOk() (*[]Goods, bool)`

GetGoodsOk returns a tuple with the Goods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoods

`func (o *OrderRedirectObject) SetGoods(v []Goods)`

SetGoods sets Goods field to given value.

### HasGoods

`func (o *OrderRedirectObject) HasGoods() bool`

HasGoods returns a boolean if a field has been set.

### GetShippingInfo

`func (o *OrderRedirectObject) GetShippingInfo() []ShippingInfo`

GetShippingInfo returns the ShippingInfo field if non-nil, zero value otherwise.

### GetShippingInfoOk

`func (o *OrderRedirectObject) GetShippingInfoOk() (*[]ShippingInfo, bool)`

GetShippingInfoOk returns a tuple with the ShippingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInfo

`func (o *OrderRedirectObject) SetShippingInfo(v []ShippingInfo)`

SetShippingInfo sets ShippingInfo field to given value.

### HasShippingInfo

`func (o *OrderRedirectObject) HasShippingInfo() bool`

HasShippingInfo returns a boolean if a field has been set.

### GetExtendInfo

`func (o *OrderRedirectObject) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *OrderRedirectObject) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *OrderRedirectObject) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *OrderRedirectObject) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.

### GetScenario

`func (o *OrderRedirectObject) GetScenario() string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *OrderRedirectObject) GetScenarioOk() (*string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *OrderRedirectObject) SetScenario(v string)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *OrderRedirectObject) HasScenario() bool`

HasScenario returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


