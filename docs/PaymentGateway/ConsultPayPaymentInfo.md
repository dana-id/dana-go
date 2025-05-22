# ConsultPayPaymentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayMethod** | **string** | Payment method that used to payment. The enums:<br />   * BALANCE - Payment method with balance<br />   * COUPON - Payment method with coupon<br />   * NET_BANKING - Payment method with internet banking<br />   * CREDIT_CARD - Payment method with credit card<br />   * DEBIT_CARD - Payment method with debit card<br />   * VIRTUAL_ACCOUNT - Payment method with virtual account<br />   * OTC - Payment method with OTC<br />   * DIRECT_DEBIT_CREDIT_CARD - Payment method with direct debit of credit card<br />   * DIRECT_DEBIT_DEBIT_CARD - Payment method with direct debit of debit card<br />   * ONLINE_CREDIT - Payment method with online Credit<br />   * LOAN_CREDIT - Payment method with DANA Cicil<br />   * NETWORK_PAY - Payment method with e-wallet<br />  | 
**PayOption** | Pointer to **string** | Payment option that available to used to payment, depends on the payment method. The enums:<br />   * NETWORK_PAY_PG_SPAY - Payment method with ShopeePay e-wallet<br />   * NETWORK_PAY_PG_OVO - Payment method with OVO e-wallet<br />   * NETWORK_PAY_PG_GOPAY - Payment method with GoPay e-wallet<br />   * NETWORK_PAY_PG_LINKAJA - Payment method with LinkAja e-wallet<br />   * NETWORK_PAY_PG_CARD - Payment method with Card<br />   * VIRTUAL_ACCOUNT_BCA - Payment method with BCA virtual account<br />   * VIRTUAL_ACCOUNT_BNI - Payment method with BNI virtual account<br />   * VIRTUAL_ACCOUNT_MANDIRI - Payment method with Mandiri virtual account<br />   * VIRTUAL_ACCOUNT_BRI - Payment method with BRI virtual account<br />   * VIRTUAL_ACCOUNT_BTPN - Payment method with BTPN virtual account<br />   * VIRTUAL_ACCOUNT_CIMB - Payment method with CIMB virtual account<br />   * VIRTUAL_ACCOUNT_PERMATA - Payment method with Permata virtual account<br />  | [optional] 
**PromoInfos** | Pointer to [**[]PromoInfo**](PromoInfo.md) | Additional Information of promotion | [optional] 

## Methods

### NewConsultPayPaymentInfo

`func NewConsultPayPaymentInfo(payMethod string, ) *ConsultPayPaymentInfo`

NewConsultPayPaymentInfo instantiates a new ConsultPayPaymentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsultPayPaymentInfoWithDefaults

`func NewConsultPayPaymentInfoWithDefaults() *ConsultPayPaymentInfo`

NewConsultPayPaymentInfoWithDefaults instantiates a new ConsultPayPaymentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayMethod

`func (o *ConsultPayPaymentInfo) GetPayMethod() string`

GetPayMethod returns the PayMethod field if non-nil, zero value otherwise.

### GetPayMethodOk

`func (o *ConsultPayPaymentInfo) GetPayMethodOk() (*string, bool)`

GetPayMethodOk returns a tuple with the PayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayMethod

`func (o *ConsultPayPaymentInfo) SetPayMethod(v string)`

SetPayMethod sets PayMethod field to given value.


### GetPayOption

`func (o *ConsultPayPaymentInfo) GetPayOption() string`

GetPayOption returns the PayOption field if non-nil, zero value otherwise.

### GetPayOptionOk

`func (o *ConsultPayPaymentInfo) GetPayOptionOk() (*string, bool)`

GetPayOptionOk returns a tuple with the PayOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayOption

`func (o *ConsultPayPaymentInfo) SetPayOption(v string)`

SetPayOption sets PayOption field to given value.

### HasPayOption

`func (o *ConsultPayPaymentInfo) HasPayOption() bool`

HasPayOption returns a boolean if a field has been set.

### GetPromoInfos

`func (o *ConsultPayPaymentInfo) GetPromoInfos() []PromoInfo`

GetPromoInfos returns the PromoInfos field if non-nil, zero value otherwise.

### GetPromoInfosOk

`func (o *ConsultPayPaymentInfo) GetPromoInfosOk() (*[]PromoInfo, bool)`

GetPromoInfosOk returns a tuple with the PromoInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoInfos

`func (o *ConsultPayPaymentInfo) SetPromoInfos(v []PromoInfo)`

SetPromoInfos sets PromoInfos field to given value.

### HasPromoInfos

`func (o *ConsultPayPaymentInfo) HasPromoInfos() bool`

HasPromoInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


