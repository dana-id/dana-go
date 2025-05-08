# QueryPaymentResponseAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | Pointer to **string** | Merchant identifier | [optional] 
**Buyer** | Pointer to [**Buyer**](Buyer.md) |  | [optional] 
**Seller** | Pointer to [**Seller**](Seller.md) |  | [optional] 
**AmountDetail** | Pointer to [**AmountDetail**](AmountDetail.md) |  | [optional] 
**TimeDetail** | Pointer to [**TimeDetail**](TimeDetail.md) |  | [optional] 
**Goods** | Pointer to [**[]Goods**](Goods.md) | Additional information of goods | [optional] 
**ShippingInfo** | Pointer to [**[]ShippingInfo**](ShippingInfo.md) | Additional information of shipping | [optional] 
**OrderMemo** | Pointer to **string** | Additional information of memo | [optional] 
**PaymentViews** | Pointer to [**[]PaymentView**](PaymentView.md) | Additional information of payment views | [optional] 
**ExtendInfo** | Pointer to **string** | Additional information of extend | [optional] 
**StatusDetail** | Pointer to [**StatusDetail**](StatusDetail.md) |  | [optional] 
**CloseReason** | Pointer to **string** | Additional information of close reason | [optional] 
**VirtualAccountInfo** | Pointer to [**VirtualAccountInfo**](VirtualAccountInfo.md) |  | [optional] 

## Methods

### NewQueryPaymentResponseAdditionalInfo

`func NewQueryPaymentResponseAdditionalInfo() *QueryPaymentResponseAdditionalInfo`

NewQueryPaymentResponseAdditionalInfo instantiates a new QueryPaymentResponseAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPaymentResponseAdditionalInfoWithDefaults

`func NewQueryPaymentResponseAdditionalInfoWithDefaults() *QueryPaymentResponseAdditionalInfo`

NewQueryPaymentResponseAdditionalInfoWithDefaults instantiates a new QueryPaymentResponseAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *QueryPaymentResponseAdditionalInfo) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *QueryPaymentResponseAdditionalInfo) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *QueryPaymentResponseAdditionalInfo) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *QueryPaymentResponseAdditionalInfo) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetBuyer

`func (o *QueryPaymentResponseAdditionalInfo) GetBuyer() Buyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *QueryPaymentResponseAdditionalInfo) GetBuyerOk() (*Buyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *QueryPaymentResponseAdditionalInfo) SetBuyer(v Buyer)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *QueryPaymentResponseAdditionalInfo) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetSeller

`func (o *QueryPaymentResponseAdditionalInfo) GetSeller() Seller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *QueryPaymentResponseAdditionalInfo) GetSellerOk() (*Seller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *QueryPaymentResponseAdditionalInfo) SetSeller(v Seller)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *QueryPaymentResponseAdditionalInfo) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetAmountDetail

`func (o *QueryPaymentResponseAdditionalInfo) GetAmountDetail() AmountDetail`

GetAmountDetail returns the AmountDetail field if non-nil, zero value otherwise.

### GetAmountDetailOk

`func (o *QueryPaymentResponseAdditionalInfo) GetAmountDetailOk() (*AmountDetail, bool)`

GetAmountDetailOk returns a tuple with the AmountDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDetail

`func (o *QueryPaymentResponseAdditionalInfo) SetAmountDetail(v AmountDetail)`

SetAmountDetail sets AmountDetail field to given value.

### HasAmountDetail

`func (o *QueryPaymentResponseAdditionalInfo) HasAmountDetail() bool`

HasAmountDetail returns a boolean if a field has been set.

### GetTimeDetail

`func (o *QueryPaymentResponseAdditionalInfo) GetTimeDetail() TimeDetail`

GetTimeDetail returns the TimeDetail field if non-nil, zero value otherwise.

### GetTimeDetailOk

`func (o *QueryPaymentResponseAdditionalInfo) GetTimeDetailOk() (*TimeDetail, bool)`

GetTimeDetailOk returns a tuple with the TimeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeDetail

`func (o *QueryPaymentResponseAdditionalInfo) SetTimeDetail(v TimeDetail)`

SetTimeDetail sets TimeDetail field to given value.

### HasTimeDetail

`func (o *QueryPaymentResponseAdditionalInfo) HasTimeDetail() bool`

HasTimeDetail returns a boolean if a field has been set.

### GetGoods

`func (o *QueryPaymentResponseAdditionalInfo) GetGoods() []Goods`

GetGoods returns the Goods field if non-nil, zero value otherwise.

### GetGoodsOk

`func (o *QueryPaymentResponseAdditionalInfo) GetGoodsOk() (*[]Goods, bool)`

GetGoodsOk returns a tuple with the Goods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoods

`func (o *QueryPaymentResponseAdditionalInfo) SetGoods(v []Goods)`

SetGoods sets Goods field to given value.

### HasGoods

`func (o *QueryPaymentResponseAdditionalInfo) HasGoods() bool`

HasGoods returns a boolean if a field has been set.

### GetShippingInfo

`func (o *QueryPaymentResponseAdditionalInfo) GetShippingInfo() []ShippingInfo`

GetShippingInfo returns the ShippingInfo field if non-nil, zero value otherwise.

### GetShippingInfoOk

`func (o *QueryPaymentResponseAdditionalInfo) GetShippingInfoOk() (*[]ShippingInfo, bool)`

GetShippingInfoOk returns a tuple with the ShippingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInfo

`func (o *QueryPaymentResponseAdditionalInfo) SetShippingInfo(v []ShippingInfo)`

SetShippingInfo sets ShippingInfo field to given value.

### HasShippingInfo

`func (o *QueryPaymentResponseAdditionalInfo) HasShippingInfo() bool`

HasShippingInfo returns a boolean if a field has been set.

### GetOrderMemo

`func (o *QueryPaymentResponseAdditionalInfo) GetOrderMemo() string`

GetOrderMemo returns the OrderMemo field if non-nil, zero value otherwise.

### GetOrderMemoOk

`func (o *QueryPaymentResponseAdditionalInfo) GetOrderMemoOk() (*string, bool)`

GetOrderMemoOk returns a tuple with the OrderMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderMemo

`func (o *QueryPaymentResponseAdditionalInfo) SetOrderMemo(v string)`

SetOrderMemo sets OrderMemo field to given value.

### HasOrderMemo

`func (o *QueryPaymentResponseAdditionalInfo) HasOrderMemo() bool`

HasOrderMemo returns a boolean if a field has been set.

### GetPaymentViews

`func (o *QueryPaymentResponseAdditionalInfo) GetPaymentViews() []PaymentView`

GetPaymentViews returns the PaymentViews field if non-nil, zero value otherwise.

### GetPaymentViewsOk

`func (o *QueryPaymentResponseAdditionalInfo) GetPaymentViewsOk() (*[]PaymentView, bool)`

GetPaymentViewsOk returns a tuple with the PaymentViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentViews

`func (o *QueryPaymentResponseAdditionalInfo) SetPaymentViews(v []PaymentView)`

SetPaymentViews sets PaymentViews field to given value.

### HasPaymentViews

`func (o *QueryPaymentResponseAdditionalInfo) HasPaymentViews() bool`

HasPaymentViews returns a boolean if a field has been set.

### GetExtendInfo

`func (o *QueryPaymentResponseAdditionalInfo) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *QueryPaymentResponseAdditionalInfo) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *QueryPaymentResponseAdditionalInfo) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *QueryPaymentResponseAdditionalInfo) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.

### GetStatusDetail

`func (o *QueryPaymentResponseAdditionalInfo) GetStatusDetail() StatusDetail`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *QueryPaymentResponseAdditionalInfo) GetStatusDetailOk() (*StatusDetail, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *QueryPaymentResponseAdditionalInfo) SetStatusDetail(v StatusDetail)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *QueryPaymentResponseAdditionalInfo) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.

### GetCloseReason

`func (o *QueryPaymentResponseAdditionalInfo) GetCloseReason() string`

GetCloseReason returns the CloseReason field if non-nil, zero value otherwise.

### GetCloseReasonOk

`func (o *QueryPaymentResponseAdditionalInfo) GetCloseReasonOk() (*string, bool)`

GetCloseReasonOk returns a tuple with the CloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseReason

`func (o *QueryPaymentResponseAdditionalInfo) SetCloseReason(v string)`

SetCloseReason sets CloseReason field to given value.

### HasCloseReason

`func (o *QueryPaymentResponseAdditionalInfo) HasCloseReason() bool`

HasCloseReason returns a boolean if a field has been set.

### GetVirtualAccountInfo

`func (o *QueryPaymentResponseAdditionalInfo) GetVirtualAccountInfo() VirtualAccountInfo`

GetVirtualAccountInfo returns the VirtualAccountInfo field if non-nil, zero value otherwise.

### GetVirtualAccountInfoOk

`func (o *QueryPaymentResponseAdditionalInfo) GetVirtualAccountInfoOk() (*VirtualAccountInfo, bool)`

GetVirtualAccountInfoOk returns a tuple with the VirtualAccountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccountInfo

`func (o *QueryPaymentResponseAdditionalInfo) SetVirtualAccountInfo(v VirtualAccountInfo)`

SetVirtualAccountInfo sets VirtualAccountInfo field to given value.

### HasVirtualAccountInfo

`func (o *QueryPaymentResponseAdditionalInfo) HasVirtualAccountInfo() bool`

HasVirtualAccountInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


