# RefundOptionBill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayMethod** | Pointer to **string** | Payment method name. The enums:<br />   * BALANCE - Payment method with balance<br />   * COUPON - Payment method with coupon<br />   * NET_BANKING - Payment method with internet banking<br />   * CREDIT_CARD - Payment method with credit card<br />   * DEBIT_CARD - Payment method with debit card<br />   * VIRTUAL_ACCOUNT - Payment method with virtual account<br />   * OTC - Payment method with OTC<br />   * DIRECT_DEBIT_CREDIT_CARD - Payment method with direct debit of credit card<br />   * DIRECT_DEBIT_DEBIT_CARD - Payment method with direct debit of debit card<br />   * ONLINE_CREDIT - Payment method with online Credit<br />   * LOAN_CREDIT - Payment method with DANA Cicil<br />  | [optional] 
**TransAmount** | Pointer to [**Money**](Money.md) | Trans amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | [optional] 

## Methods

### NewRefundOptionBill

`func NewRefundOptionBill() *RefundOptionBill`

NewRefundOptionBill instantiates a new RefundOptionBill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundOptionBillWithDefaults

`func NewRefundOptionBillWithDefaults() *RefundOptionBill`

NewRefundOptionBillWithDefaults instantiates a new RefundOptionBill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayMethod

`func (o *RefundOptionBill) GetPayMethod() string`

GetPayMethod returns the PayMethod field if non-nil, zero value otherwise.

### GetPayMethodOk

`func (o *RefundOptionBill) GetPayMethodOk() (*string, bool)`

GetPayMethodOk returns a tuple with the PayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayMethod

`func (o *RefundOptionBill) SetPayMethod(v string)`

SetPayMethod sets PayMethod field to given value.

### HasPayMethod

`func (o *RefundOptionBill) HasPayMethod() bool`

HasPayMethod returns a boolean if a field has been set.

### GetTransAmount

`func (o *RefundOptionBill) GetTransAmount() Money`

GetTransAmount returns the TransAmount field if non-nil, zero value otherwise.

### GetTransAmountOk

`func (o *RefundOptionBill) GetTransAmountOk() (*Money, bool)`

GetTransAmountOk returns a tuple with the TransAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransAmount

`func (o *RefundOptionBill) SetTransAmount(v Money)`

SetTransAmount sets TransAmount field to given value.

### HasTransAmount

`func (o *RefundOptionBill) HasTransAmount() bool`

HasTransAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


