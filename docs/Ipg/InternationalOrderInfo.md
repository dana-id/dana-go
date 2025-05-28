# InternationalOrderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginOrderAmount** | Pointer to [**Money**](Money.md) | Origin order amount in the original currency. Contains value (amount including cents) and currency (code based on ISO) | [optional] 
**ExchangeRate** | Pointer to [**InternationalOrderInfoExchangeRate**](InternationalOrderInfoExchangeRate.md) |  | [optional] 
**TotalAmount** | Pointer to [**Money**](Money.md) | Total amount after conversion. Contains value (amount including cents) and currency (code based on ISO) | [optional] 
**PaymentPromoInfo** | Pointer to [**PaymentPromoInfo**](PaymentPromoInfo.md) | Define the detail of payment promo information, contains promotion that handled and set by merchant | [optional] 
**RefundPromoInfo** | Pointer to [**RefundPromoInfo**](RefundPromoInfo.md) | Define the detail of refund promo information | [optional] 

## Methods

### NewInternationalOrderInfo

`func NewInternationalOrderInfo() *InternationalOrderInfo`

NewInternationalOrderInfo instantiates a new InternationalOrderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternationalOrderInfoWithDefaults

`func NewInternationalOrderInfoWithDefaults() *InternationalOrderInfo`

NewInternationalOrderInfoWithDefaults instantiates a new InternationalOrderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginOrderAmount

`func (o *InternationalOrderInfo) GetOriginOrderAmount() Money`

GetOriginOrderAmount returns the OriginOrderAmount field if non-nil, zero value otherwise.

### GetOriginOrderAmountOk

`func (o *InternationalOrderInfo) GetOriginOrderAmountOk() (*Money, bool)`

GetOriginOrderAmountOk returns a tuple with the OriginOrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOrderAmount

`func (o *InternationalOrderInfo) SetOriginOrderAmount(v Money)`

SetOriginOrderAmount sets OriginOrderAmount field to given value.

### HasOriginOrderAmount

`func (o *InternationalOrderInfo) HasOriginOrderAmount() bool`

HasOriginOrderAmount returns a boolean if a field has been set.

### GetExchangeRate

`func (o *InternationalOrderInfo) GetExchangeRate() InternationalOrderInfoExchangeRate`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *InternationalOrderInfo) GetExchangeRateOk() (*InternationalOrderInfoExchangeRate, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *InternationalOrderInfo) SetExchangeRate(v InternationalOrderInfoExchangeRate)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *InternationalOrderInfo) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetTotalAmount

`func (o *InternationalOrderInfo) GetTotalAmount() Money`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *InternationalOrderInfo) GetTotalAmountOk() (*Money, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *InternationalOrderInfo) SetTotalAmount(v Money)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *InternationalOrderInfo) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetPaymentPromoInfo

`func (o *InternationalOrderInfo) GetPaymentPromoInfo() PaymentPromoInfo`

GetPaymentPromoInfo returns the PaymentPromoInfo field if non-nil, zero value otherwise.

### GetPaymentPromoInfoOk

`func (o *InternationalOrderInfo) GetPaymentPromoInfoOk() (*PaymentPromoInfo, bool)`

GetPaymentPromoInfoOk returns a tuple with the PaymentPromoInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPromoInfo

`func (o *InternationalOrderInfo) SetPaymentPromoInfo(v PaymentPromoInfo)`

SetPaymentPromoInfo sets PaymentPromoInfo field to given value.

### HasPaymentPromoInfo

`func (o *InternationalOrderInfo) HasPaymentPromoInfo() bool`

HasPaymentPromoInfo returns a boolean if a field has been set.

### GetRefundPromoInfo

`func (o *InternationalOrderInfo) GetRefundPromoInfo() RefundPromoInfo`

GetRefundPromoInfo returns the RefundPromoInfo field if non-nil, zero value otherwise.

### GetRefundPromoInfoOk

`func (o *InternationalOrderInfo) GetRefundPromoInfoOk() (*RefundPromoInfo, bool)`

GetRefundPromoInfoOk returns a tuple with the RefundPromoInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundPromoInfo

`func (o *InternationalOrderInfo) SetRefundPromoInfo(v RefundPromoInfo)`

SetRefundPromoInfo sets RefundPromoInfo field to given value.

### HasRefundPromoInfo

`func (o *InternationalOrderInfo) HasRefundPromoInfo() bool`

HasRefundPromoInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


