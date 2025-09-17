# PayOptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayMethod** | **string** | Payment method name. The enums:<br />   * BALANCE - Payment method with balance<br />   * COUPON - Payment method with coupon<br />   * NET_BANKING - Payment method with internet banking<br />   * CREDIT_CARD - Payment method with credit card<br />   * DEBIT_CARD - Payment method with debit card<br />   * VIRTUAL_ACCOUNT - Payment method with virtual account<br />   * OTC - Payment method with OTC<br />   * DIRECT_DEBIT_CREDIT_CARD - Payment method with direct debit of credit card<br />   * DIRECT_DEBIT_DEBIT_CARD - Payment method with direct debit of debit card<br />   * ONLINE_CREDIT - Payment method with online Credit<br />   * LOAN_CREDIT - Payment method with DANA Cicil<br />   * NETWORK_PAY - Payment method with e-wallet   * CARD - Payment method with Card  | 
**PayOption** | Pointer to **string** | Payment option which shows the provider of this payment. The enums:<br />   * NETWORK_PAY_PG_SPAY - Payment method with ShopeePay e-wallet<br />   * NETWORK_PAY_PG_OVO - Payment method with OVO e-wallet<br />   * NETWORK_PAY_PG_GOPAY - Payment method with GoPay e-wallet<br />   * NETWORK_PAY_PG_LINKAJA - Payment method with LinkAja e-wallet<br />   * NETWORK_PAY_PG_CARD - Payment method with Card<br />   * NETWORK_PAY_PC_INDOMARET - Payment method with Indomaret<br />   * VIRTUAL_ACCOUNT_BCA - Payment method with BCA virtual account<br />   * VIRTUAL_ACCOUNT_BNI - Payment method with BNI virtual account<br />   * VIRTUAL_ACCOUNT_MANDIRI - Payment method with Mandiri virtual account<br />   * VIRTUAL_ACCOUNT_BRI - Payment method with BRI virtual account<br />   * VIRTUAL_ACCOUNT_BTPN - Payment method with BTPN virtual account<br />   * VIRTUAL_ACCOUNT_CIMB - Payment method with CIMB virtual account<br />   * VIRTUAL_ACCOUNT_PERMATA - Payment method with Permata virtual account<br />  | [optional] 
**PayAmount** | Pointer to [**Money**](Money.md) | Pay amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | [optional] 
**TransAmount** | Pointer to [**Money**](Money.md) | Trans amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | [optional] 
**ChargeAmount** | Pointer to [**Money**](Money.md) | Charge amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | [optional] 
**PayOptionBillExtendInfo** | Pointer to **string** | Extend information of pay option bill | [optional] 
**ExtendInfo** | Pointer to **string** | Extend information | [optional] 
**PaymentCode** | Pointer to **string** | Payment code | [optional] 

## Methods

### NewPayOptionInfo

`func NewPayOptionInfo(payMethod string, ) *PayOptionInfo`

NewPayOptionInfo instantiates a new PayOptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayOptionInfoWithDefaults

`func NewPayOptionInfoWithDefaults() *PayOptionInfo`

NewPayOptionInfoWithDefaults instantiates a new PayOptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayMethod

`func (o *PayOptionInfo) GetPayMethod() string`

GetPayMethod returns the PayMethod field if non-nil, zero value otherwise.

### GetPayMethodOk

`func (o *PayOptionInfo) GetPayMethodOk() (*string, bool)`

GetPayMethodOk returns a tuple with the PayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayMethod

`func (o *PayOptionInfo) SetPayMethod(v string)`

SetPayMethod sets PayMethod field to given value.


### GetPayOption

`func (o *PayOptionInfo) GetPayOption() string`

GetPayOption returns the PayOption field if non-nil, zero value otherwise.

### GetPayOptionOk

`func (o *PayOptionInfo) GetPayOptionOk() (*string, bool)`

GetPayOptionOk returns a tuple with the PayOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayOption

`func (o *PayOptionInfo) SetPayOption(v string)`

SetPayOption sets PayOption field to given value.

### HasPayOption

`func (o *PayOptionInfo) HasPayOption() bool`

HasPayOption returns a boolean if a field has been set.

### GetPayAmount

`func (o *PayOptionInfo) GetPayAmount() Money`

GetPayAmount returns the PayAmount field if non-nil, zero value otherwise.

### GetPayAmountOk

`func (o *PayOptionInfo) GetPayAmountOk() (*Money, bool)`

GetPayAmountOk returns a tuple with the PayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAmount

`func (o *PayOptionInfo) SetPayAmount(v Money)`

SetPayAmount sets PayAmount field to given value.

### HasPayAmount

`func (o *PayOptionInfo) HasPayAmount() bool`

HasPayAmount returns a boolean if a field has been set.

### GetTransAmount

`func (o *PayOptionInfo) GetTransAmount() Money`

GetTransAmount returns the TransAmount field if non-nil, zero value otherwise.

### GetTransAmountOk

`func (o *PayOptionInfo) GetTransAmountOk() (*Money, bool)`

GetTransAmountOk returns a tuple with the TransAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransAmount

`func (o *PayOptionInfo) SetTransAmount(v Money)`

SetTransAmount sets TransAmount field to given value.

### HasTransAmount

`func (o *PayOptionInfo) HasTransAmount() bool`

HasTransAmount returns a boolean if a field has been set.

### GetChargeAmount

`func (o *PayOptionInfo) GetChargeAmount() Money`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *PayOptionInfo) GetChargeAmountOk() (*Money, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *PayOptionInfo) SetChargeAmount(v Money)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *PayOptionInfo) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.

### GetPayOptionBillExtendInfo

`func (o *PayOptionInfo) GetPayOptionBillExtendInfo() string`

GetPayOptionBillExtendInfo returns the PayOptionBillExtendInfo field if non-nil, zero value otherwise.

### GetPayOptionBillExtendInfoOk

`func (o *PayOptionInfo) GetPayOptionBillExtendInfoOk() (*string, bool)`

GetPayOptionBillExtendInfoOk returns a tuple with the PayOptionBillExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayOptionBillExtendInfo

`func (o *PayOptionInfo) SetPayOptionBillExtendInfo(v string)`

SetPayOptionBillExtendInfo sets PayOptionBillExtendInfo field to given value.

### HasPayOptionBillExtendInfo

`func (o *PayOptionInfo) HasPayOptionBillExtendInfo() bool`

HasPayOptionBillExtendInfo returns a boolean if a field has been set.

### GetExtendInfo

`func (o *PayOptionInfo) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *PayOptionInfo) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *PayOptionInfo) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *PayOptionInfo) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.

### GetPaymentCode

`func (o *PayOptionInfo) GetPaymentCode() string`

GetPaymentCode returns the PaymentCode field if non-nil, zero value otherwise.

### GetPaymentCodeOk

`func (o *PayOptionInfo) GetPaymentCodeOk() (*string, bool)`

GetPaymentCodeOk returns a tuple with the PaymentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCode

`func (o *PayOptionInfo) SetPaymentCode(v string)`

SetPaymentCode sets PaymentCode field to given value.

### HasPaymentCode

`func (o *PayOptionInfo) HasPaymentCode() bool`

HasPaymentCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


