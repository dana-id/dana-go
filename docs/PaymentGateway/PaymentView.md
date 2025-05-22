# PaymentView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CashierRequestId** | Pointer to **string** | Cashier request identifier | [optional] 
**PaidTime** | Pointer to **string** | Information of paid time, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | [optional] 
**PayOptionInfos** | [**[]PayOptionInfo**](PayOptionInfo.md) | Information of pay option | 
**PayRequestExtendInfo** | Pointer to **string** | Extend information of pay request | [optional] 
**ExtendInfo** | Pointer to **string** | Extend information | [optional] 

## Methods

### NewPaymentView

`func NewPaymentView(payOptionInfos []PayOptionInfo, ) *PaymentView`

NewPaymentView instantiates a new PaymentView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentViewWithDefaults

`func NewPaymentViewWithDefaults() *PaymentView`

NewPaymentViewWithDefaults instantiates a new PaymentView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCashierRequestId

`func (o *PaymentView) GetCashierRequestId() string`

GetCashierRequestId returns the CashierRequestId field if non-nil, zero value otherwise.

### GetCashierRequestIdOk

`func (o *PaymentView) GetCashierRequestIdOk() (*string, bool)`

GetCashierRequestIdOk returns a tuple with the CashierRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashierRequestId

`func (o *PaymentView) SetCashierRequestId(v string)`

SetCashierRequestId sets CashierRequestId field to given value.

### HasCashierRequestId

`func (o *PaymentView) HasCashierRequestId() bool`

HasCashierRequestId returns a boolean if a field has been set.

### GetPaidTime

`func (o *PaymentView) GetPaidTime() string`

GetPaidTime returns the PaidTime field if non-nil, zero value otherwise.

### GetPaidTimeOk

`func (o *PaymentView) GetPaidTimeOk() (*string, bool)`

GetPaidTimeOk returns a tuple with the PaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTime

`func (o *PaymentView) SetPaidTime(v string)`

SetPaidTime sets PaidTime field to given value.

### HasPaidTime

`func (o *PaymentView) HasPaidTime() bool`

HasPaidTime returns a boolean if a field has been set.

### GetPayOptionInfos

`func (o *PaymentView) GetPayOptionInfos() []PayOptionInfo`

GetPayOptionInfos returns the PayOptionInfos field if non-nil, zero value otherwise.

### GetPayOptionInfosOk

`func (o *PaymentView) GetPayOptionInfosOk() (*[]PayOptionInfo, bool)`

GetPayOptionInfosOk returns a tuple with the PayOptionInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayOptionInfos

`func (o *PaymentView) SetPayOptionInfos(v []PayOptionInfo)`

SetPayOptionInfos sets PayOptionInfos field to given value.


### GetPayRequestExtendInfo

`func (o *PaymentView) GetPayRequestExtendInfo() string`

GetPayRequestExtendInfo returns the PayRequestExtendInfo field if non-nil, zero value otherwise.

### GetPayRequestExtendInfoOk

`func (o *PaymentView) GetPayRequestExtendInfoOk() (*string, bool)`

GetPayRequestExtendInfoOk returns a tuple with the PayRequestExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayRequestExtendInfo

`func (o *PaymentView) SetPayRequestExtendInfo(v string)`

SetPayRequestExtendInfo sets PayRequestExtendInfo field to given value.

### HasPayRequestExtendInfo

`func (o *PaymentView) HasPayRequestExtendInfo() bool`

HasPayRequestExtendInfo returns a boolean if a field has been set.

### GetExtendInfo

`func (o *PaymentView) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *PaymentView) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *PaymentView) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *PaymentView) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


