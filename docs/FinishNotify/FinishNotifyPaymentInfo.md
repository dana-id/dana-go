# FinishNotifyPaymentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CashierRequestId** | **string** | Cashier request identifier | [optional]
**PaidTime** | **string** | Information of paid time, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | [optional]
**PayOptionInfos** | [**[]PayOptionInfo**](PayOptionInfo.md) | Information of pay option. Refer to payOptionInfos for the detailed | 
**PayRequestExtendInfo** | Pointer to **string** | Extend information of pay request | [optional] 
**ExtendInfo** | Pointer to **string** | Extend information | [optional] 

## Methods

### NewFinishNotifyPaymentInfo

`func NewFinishNotifyPaymentInfo(cashierRequestId string, paidTime string, payOptionInfos []PayOptionInfo, ) *FinishNotifyPaymentInfo`

NewFinishNotifyPaymentInfo instantiates a new FinishNotifyPaymentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinishNotifyPaymentInfoWithDefaults

`func NewFinishNotifyPaymentInfoWithDefaults() *FinishNotifyPaymentInfo`

NewFinishNotifyPaymentInfoWithDefaults instantiates a new FinishNotifyPaymentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCashierRequestId

`func (o *FinishNotifyPaymentInfo) GetCashierRequestId() string`

GetCashierRequestId returns the CashierRequestId field if non-nil, zero value otherwise.

### GetCashierRequestIdOk

`func (o *FinishNotifyPaymentInfo) GetCashierRequestIdOk() (*string, bool)`

GetCashierRequestIdOk returns a tuple with the CashierRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashierRequestId

`func (o *FinishNotifyPaymentInfo) SetCashierRequestId(v *string)`

SetCashierRequestId sets CashierRequestId field to given value.


### GetPaidTime

`func (o *FinishNotifyPaymentInfo) GetPaidTime() string`

GetPaidTime returns the PaidTime field if non-nil, zero value otherwise.

### GetPaidTimeOk

`func (o *FinishNotifyPaymentInfo) GetPaidTimeOk() (*string, bool)`

GetPaidTimeOk returns a tuple with the PaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTime

`func (o *FinishNotifyPaymentInfo) SetPaidTime(v *string)`

SetPaidTime sets PaidTime field to given value.


### GetPayOptionInfos

`func (o *FinishNotifyPaymentInfo) GetPayOptionInfos() []PayOptionInfo`

GetPayOptionInfos returns the PayOptionInfos field if non-nil, zero value otherwise.

### GetPayOptionInfosOk

`func (o *FinishNotifyPaymentInfo) GetPayOptionInfosOk() (*[]PayOptionInfo, bool)`

GetPayOptionInfosOk returns a tuple with the PayOptionInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayOptionInfos

`func (o *FinishNotifyPaymentInfo) SetPayOptionInfos(v []PayOptionInfo)`

SetPayOptionInfos sets PayOptionInfos field to given value.


### GetPayRequestExtendInfo

`func (o *FinishNotifyPaymentInfo) GetPayRequestExtendInfo() string`

GetPayRequestExtendInfo returns the PayRequestExtendInfo field if non-nil, zero value otherwise.

### GetPayRequestExtendInfoOk

`func (o *FinishNotifyPaymentInfo) GetPayRequestExtendInfoOk() (*string, bool)`

GetPayRequestExtendInfoOk returns a tuple with the PayRequestExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayRequestExtendInfo

`func (o *FinishNotifyPaymentInfo) SetPayRequestExtendInfo(v string)`

SetPayRequestExtendInfo sets PayRequestExtendInfo field to given value.

### HasPayRequestExtendInfo

`func (o *FinishNotifyPaymentInfo) HasPayRequestExtendInfo() bool`

HasPayRequestExtendInfo returns a boolean if a field has been set.

### GetExtendInfo

`func (o *FinishNotifyPaymentInfo) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *FinishNotifyPaymentInfo) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *FinishNotifyPaymentInfo) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *FinishNotifyPaymentInfo) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


