# FinishNotifyRequestAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentInfo** | Pointer to [**FinishNotifyPaymentInfo**](FinishNotifyPaymentInfo.md) |  | [optional] 
**ShopInfo** | Pointer to [**ShopInfo**](ShopInfo.md) |  | [optional] 
**ExtendInfo** | Pointer to **string** | Extended information (as a JSON string) | [optional] 
**ExtendInfoClosedReason** | Pointer to **string** | Reason for order closure (if order is closed) | [optional] 

## Methods

### NewFinishNotifyAdditionalInfo

`func NewFinishNotifyAdditionalInfo() *FinishNotifyAdditionalInfo`

NewFinishNotifyAdditionalInfo instantiates a new FinishNotifyAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinishNotifyAdditionalInfoWithDefaults

`func NewFinishNotifyAdditionalInfoWithDefaults() *FinishNotifyAdditionalInfo`

NewFinishNotifyAdditionalInfoWithDefaults instantiates a new FinishNotifyAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentInfo

`func (o *FinishNotifyAdditionalInfo) GetPaymentInfo() FinishNotifyPaymentInfo`

GetPaymentInfo returns the PaymentInfo field if non-nil, zero value otherwise.

### GetPaymentInfoOk

`func (o *FinishNotifyAdditionalInfo) GetPaymentInfoOk() (*FinishNotifyPaymentInfo, bool)`

GetPaymentInfoOk returns a tuple with the PaymentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInfo

`func (o *FinishNotifyAdditionalInfo) SetPaymentInfo(v FinishNotifyPaymentInfo)`

SetPaymentInfo sets PaymentInfo field to given value.

### HasPaymentInfo

`func (o *FinishNotifyAdditionalInfo) HasPaymentInfo() bool`

HasPaymentInfo returns a boolean if a field has been set.

### GetShopInfo

`func (o *FinishNotifyAdditionalInfo) GetShopInfo() ShopInfo`

GetShopInfo returns the ShopInfo field if non-nil, zero value otherwise.

### GetShopInfoOk

`func (o *FinishNotifyAdditionalInfo) GetShopInfoOk() (*ShopInfo, bool)`

GetShopInfoOk returns a tuple with the ShopInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopInfo

`func (o *FinishNotifyAdditionalInfo) SetShopInfo(v ShopInfo)`

SetShopInfo sets ShopInfo field to given value.

### HasShopInfo

`func (o *FinishNotifyAdditionalInfo) HasShopInfo() bool`

HasShopInfo returns a boolean if a field has been set.

### GetExtendInfo

`func (o *FinishNotifyAdditionalInfo) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *FinishNotifyAdditionalInfo) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *FinishNotifyAdditionalInfo) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *FinishNotifyAdditionalInfo) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.

### GetExtendInfoClosedReason

`func (o *FinishNotifyAdditionalInfo) GetExtendInfoClosedReason() string`

GetExtendInfoClosedReason returns the ExtendInfoClosedReason field if non-nil, zero value otherwise.

### GetExtendInfoClosedReasonOk

`func (o *FinishNotifyAdditionalInfo) GetExtendInfoClosedReasonOk() (*string, bool)`

GetExtendInfoClosedReasonOk returns a tuple with the ExtendInfoClosedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfoClosedReason

`func (o *FinishNotifyAdditionalInfo) SetExtendInfoClosedReason(v string)`

SetExtendInfoClosedReason sets ExtendInfoClosedReason field to given value.

### HasExtendInfoClosedReason

`func (o *FinishNotifyAdditionalInfo) HasExtendInfoClosedReason() bool`

HasExtendInfoClosedReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


