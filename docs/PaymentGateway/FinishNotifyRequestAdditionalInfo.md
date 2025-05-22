# FinishNotifyRequestAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentInfo** | Pointer to [**FinishNotifyPaymentInfo**](FinishNotifyPaymentInfo.md) | Additional information of payment | [optional] 
**ShopInfo** | Pointer to [**ShopInfo**](ShopInfo.md) | Additional information of shop | [optional] 
**ExtendInfo** | Pointer to **string** | Additional information of extend (As a JSON string) | [optional] 
**ExtendInfoClosedReason** | Pointer to **string** | Additional information of closed reason. Required if order is closed | [optional] 

## Methods

### NewFinishNotifyRequestAdditionalInfo

`func NewFinishNotifyRequestAdditionalInfo() *FinishNotifyRequestAdditionalInfo`

NewFinishNotifyRequestAdditionalInfo instantiates a new FinishNotifyRequestAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinishNotifyRequestAdditionalInfoWithDefaults

`func NewFinishNotifyRequestAdditionalInfoWithDefaults() *FinishNotifyRequestAdditionalInfo`

NewFinishNotifyRequestAdditionalInfoWithDefaults instantiates a new FinishNotifyRequestAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentInfo

`func (o *FinishNotifyRequestAdditionalInfo) GetPaymentInfo() FinishNotifyPaymentInfo`

GetPaymentInfo returns the PaymentInfo field if non-nil, zero value otherwise.

### GetPaymentInfoOk

`func (o *FinishNotifyRequestAdditionalInfo) GetPaymentInfoOk() (*FinishNotifyPaymentInfo, bool)`

GetPaymentInfoOk returns a tuple with the PaymentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInfo

`func (o *FinishNotifyRequestAdditionalInfo) SetPaymentInfo(v FinishNotifyPaymentInfo)`

SetPaymentInfo sets PaymentInfo field to given value.

### HasPaymentInfo

`func (o *FinishNotifyRequestAdditionalInfo) HasPaymentInfo() bool`

HasPaymentInfo returns a boolean if a field has been set.

### GetShopInfo

`func (o *FinishNotifyRequestAdditionalInfo) GetShopInfo() ShopInfo`

GetShopInfo returns the ShopInfo field if non-nil, zero value otherwise.

### GetShopInfoOk

`func (o *FinishNotifyRequestAdditionalInfo) GetShopInfoOk() (*ShopInfo, bool)`

GetShopInfoOk returns a tuple with the ShopInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopInfo

`func (o *FinishNotifyRequestAdditionalInfo) SetShopInfo(v ShopInfo)`

SetShopInfo sets ShopInfo field to given value.

### HasShopInfo

`func (o *FinishNotifyRequestAdditionalInfo) HasShopInfo() bool`

HasShopInfo returns a boolean if a field has been set.

### GetExtendInfo

`func (o *FinishNotifyRequestAdditionalInfo) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *FinishNotifyRequestAdditionalInfo) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *FinishNotifyRequestAdditionalInfo) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *FinishNotifyRequestAdditionalInfo) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.

### GetExtendInfoClosedReason

`func (o *FinishNotifyRequestAdditionalInfo) GetExtendInfoClosedReason() string`

GetExtendInfoClosedReason returns the ExtendInfoClosedReason field if non-nil, zero value otherwise.

### GetExtendInfoClosedReasonOk

`func (o *FinishNotifyRequestAdditionalInfo) GetExtendInfoClosedReasonOk() (*string, bool)`

GetExtendInfoClosedReasonOk returns a tuple with the ExtendInfoClosedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfoClosedReason

`func (o *FinishNotifyRequestAdditionalInfo) SetExtendInfoClosedReason(v string)`

SetExtendInfoClosedReason sets ExtendInfoClosedReason field to given value.

### HasExtendInfoClosedReason

`func (o *FinishNotifyRequestAdditionalInfo) HasExtendInfoClosedReason() bool`

HasExtendInfoClosedReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


