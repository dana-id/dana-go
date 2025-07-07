# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buyer** | Pointer to [**Buyer**](Buyer.md) |  | [optional] 
**Seller** | Pointer to [**Seller**](Seller.md) |  | [optional] 
**OrderTitle** | **string** | Additional information of order title | 
**MerchantTransType** | Pointer to **string** | Additional information of merchant transaction type | [optional] 
**OrderMemo** | Pointer to **string** | Additional information of order memo | [optional] 
**CreatedTime** | Pointer to **string** | Additional information of created time, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | [optional] 
**Goods** | Pointer to [**[]Goods**](Goods.md) |  | [optional] 
**ShippingInfo** | Pointer to [**[]ShippingInfo**](ShippingInfo.md) | Additional information of shipping | [optional] 
**InternationalOrderInfo** | Pointer to [**InternationalOrderInfo**](InternationalOrderInfo.md) | Additional information of international order. Only use for Mini Program service | [optional] 
**ExtendInfo** | Pointer to **string** | Additional information of extend | [optional] 

## Methods

### NewOrder

`func NewOrder(orderTitle string, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyer

`func (o *Order) GetBuyer() Buyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *Order) GetBuyerOk() (*Buyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *Order) SetBuyer(v Buyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *Order) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetSeller

`func (o *Order) GetSeller() Seller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *Order) GetSellerOk() (*Seller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *Order) SetSeller(v Seller)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *Order) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetOrderTitle

`func (o *Order) GetOrderTitle() string`

GetOrderTitle returns the OrderTitle field if non-nil, zero value otherwise.

### GetOrderTitleOk

`func (o *Order) GetOrderTitleOk() (*string, bool)`

GetOrderTitleOk returns a tuple with the OrderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTitle

`func (o *Order) SetOrderTitle(v string)`

SetOrderTitle sets OrderTitle field to given value.


### GetMerchantTransType

`func (o *Order) GetMerchantTransType() string`

GetMerchantTransType returns the MerchantTransType field if non-nil, zero value otherwise.

### GetMerchantTransTypeOk

`func (o *Order) GetMerchantTransTypeOk() (*string, bool)`

GetMerchantTransTypeOk returns a tuple with the MerchantTransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTransType

`func (o *Order) SetMerchantTransType(v string)`

SetMerchantTransType sets MerchantTransType field to given value.

### HasMerchantTransType

`func (o *Order) HasMerchantTransType() bool`

HasMerchantTransType returns a boolean if a field has been set.

### GetOrderMemo

`func (o *Order) GetOrderMemo() string`

GetOrderMemo returns the OrderMemo field if non-nil, zero value otherwise.

### GetOrderMemoOk

`func (o *Order) GetOrderMemoOk() (*string, bool)`

GetOrderMemoOk returns a tuple with the OrderMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderMemo

`func (o *Order) SetOrderMemo(v string)`

SetOrderMemo sets OrderMemo field to given value.

### HasOrderMemo

`func (o *Order) HasOrderMemo() bool`

HasOrderMemo returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Order) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Order) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Order) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Order) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetGoods

`func (o *Order) GetGoods() []Goods`

GetGoods returns the Goods field if non-nil, zero value otherwise.

### GetGoodsOk

`func (o *Order) GetGoodsOk() (*[]Goods, bool)`

GetGoodsOk returns a tuple with the Goods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoods

`func (o *Order) SetGoods(v []Goods)`

SetGoods sets Goods field to given value.

### HasGoods

`func (o *Order) HasGoods() bool`

HasGoods returns a boolean if a field has been set.

### GetShippingInfo

`func (o *Order) GetShippingInfo() []ShippingInfo`

GetShippingInfo returns the ShippingInfo field if non-nil, zero value otherwise.

### GetShippingInfoOk

`func (o *Order) GetShippingInfoOk() (*[]ShippingInfo, bool)`

GetShippingInfoOk returns a tuple with the ShippingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInfo

`func (o *Order) SetShippingInfo(v []ShippingInfo)`

SetShippingInfo sets ShippingInfo field to given value.

### HasShippingInfo

`func (o *Order) HasShippingInfo() bool`

HasShippingInfo returns a boolean if a field has been set.

### GetInternationalOrderInfo

`func (o *Order) GetInternationalOrderInfo() InternationalOrderInfo`

GetInternationalOrderInfo returns the InternationalOrderInfo field if non-nil, zero value otherwise.

### GetInternationalOrderInfoOk

`func (o *Order) GetInternationalOrderInfoOk() (*InternationalOrderInfo, bool)`

GetInternationalOrderInfoOk returns a tuple with the InternationalOrderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalOrderInfo

`func (o *Order) SetInternationalOrderInfo(v InternationalOrderInfo)`

SetInternationalOrderInfo sets InternationalOrderInfo field to given value.

### HasInternationalOrderInfo

`func (o *Order) HasInternationalOrderInfo() bool`

HasInternationalOrderInfo returns a boolean if a field has been set.

### GetExtendInfo

`func (o *Order) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *Order) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *Order) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *Order) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


