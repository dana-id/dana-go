# Goods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Goods name | 
**MerchantGoodsId** | **string** | Goods identifier provided by merchant | 
**Description** | **string** | Goods description | 
**Category** | **string** | Goods category | 
**Price** | [**Money**](Money.md) | Goods price. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | 
**Unit** | Pointer to **string** | Goods unit | [optional] 
**Quantity** | **string** | Count of items | 
**MerchantShippingId** | Pointer to **string** | Shipment identifier provided by merchant | [optional] 
**SnapshotUrl** | Pointer to **string** | The URL of good&#39;s snapshot web page | [optional] 
**ExtendInfo** | Pointer to **string** | Extend information | [optional] 

## Methods

### NewGoods

`func NewGoods(name string, merchantGoodsId string, description string, category string, price Money, quantity string, ) *Goods`

NewGoods instantiates a new Goods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoodsWithDefaults

`func NewGoodsWithDefaults() *Goods`

NewGoodsWithDefaults instantiates a new Goods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Goods) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Goods) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Goods) SetName(v string)`

SetName sets Name field to given value.


### GetMerchantGoodsId

`func (o *Goods) GetMerchantGoodsId() string`

GetMerchantGoodsId returns the MerchantGoodsId field if non-nil, zero value otherwise.

### GetMerchantGoodsIdOk

`func (o *Goods) GetMerchantGoodsIdOk() (*string, bool)`

GetMerchantGoodsIdOk returns a tuple with the MerchantGoodsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGoodsId

`func (o *Goods) SetMerchantGoodsId(v string)`

SetMerchantGoodsId sets MerchantGoodsId field to given value.


### GetDescription

`func (o *Goods) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Goods) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Goods) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCategory

`func (o *Goods) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Goods) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Goods) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetPrice

`func (o *Goods) GetPrice() Money`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Goods) GetPriceOk() (*Money, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Goods) SetPrice(v Money)`

SetPrice sets Price field to given value.


### GetUnit

`func (o *Goods) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Goods) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Goods) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Goods) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetQuantity

`func (o *Goods) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Goods) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Goods) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetMerchantShippingId

`func (o *Goods) GetMerchantShippingId() string`

GetMerchantShippingId returns the MerchantShippingId field if non-nil, zero value otherwise.

### GetMerchantShippingIdOk

`func (o *Goods) GetMerchantShippingIdOk() (*string, bool)`

GetMerchantShippingIdOk returns a tuple with the MerchantShippingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantShippingId

`func (o *Goods) SetMerchantShippingId(v string)`

SetMerchantShippingId sets MerchantShippingId field to given value.

### HasMerchantShippingId

`func (o *Goods) HasMerchantShippingId() bool`

HasMerchantShippingId returns a boolean if a field has been set.

### GetSnapshotUrl

`func (o *Goods) GetSnapshotUrl() string`

GetSnapshotUrl returns the SnapshotUrl field if non-nil, zero value otherwise.

### GetSnapshotUrlOk

`func (o *Goods) GetSnapshotUrlOk() (*string, bool)`

GetSnapshotUrlOk returns a tuple with the SnapshotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotUrl

`func (o *Goods) SetSnapshotUrl(v string)`

SetSnapshotUrl sets SnapshotUrl field to given value.

### HasSnapshotUrl

`func (o *Goods) HasSnapshotUrl() bool`

HasSnapshotUrl returns a boolean if a field has been set.

### GetExtendInfo

`func (o *Goods) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *Goods) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *Goods) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *Goods) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


