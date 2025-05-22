# AmountDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderAmount** | [**Money**](Money.md) | Order amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | 
**PayAmount** | Pointer to [**Money**](Money.md) | Pay amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | [optional] 
**VoidAmount** | Pointer to [**Money**](Money.md) | Void amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | [optional] 
**ConfirmAmount** | Pointer to [**Money**](Money.md) | Confirm amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | [optional] 
**RefundAmount** | Pointer to [**Money**](Money.md) | Refund amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | [optional] 
**ChargebackAmount** | Pointer to [**Money**](Money.md) | Chargeback amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | [optional] 
**ChargeAmount** | Pointer to [**Money**](Money.md) | Charge amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | [optional] 

## Methods

### NewAmountDetail

`func NewAmountDetail(orderAmount Money, ) *AmountDetail`

NewAmountDetail instantiates a new AmountDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmountDetailWithDefaults

`func NewAmountDetailWithDefaults() *AmountDetail`

NewAmountDetailWithDefaults instantiates a new AmountDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderAmount

`func (o *AmountDetail) GetOrderAmount() Money`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *AmountDetail) GetOrderAmountOk() (*Money, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *AmountDetail) SetOrderAmount(v Money)`

SetOrderAmount sets OrderAmount field to given value.


### GetPayAmount

`func (o *AmountDetail) GetPayAmount() Money`

GetPayAmount returns the PayAmount field if non-nil, zero value otherwise.

### GetPayAmountOk

`func (o *AmountDetail) GetPayAmountOk() (*Money, bool)`

GetPayAmountOk returns a tuple with the PayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAmount

`func (o *AmountDetail) SetPayAmount(v Money)`

SetPayAmount sets PayAmount field to given value.

### HasPayAmount

`func (o *AmountDetail) HasPayAmount() bool`

HasPayAmount returns a boolean if a field has been set.

### GetVoidAmount

`func (o *AmountDetail) GetVoidAmount() Money`

GetVoidAmount returns the VoidAmount field if non-nil, zero value otherwise.

### GetVoidAmountOk

`func (o *AmountDetail) GetVoidAmountOk() (*Money, bool)`

GetVoidAmountOk returns a tuple with the VoidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidAmount

`func (o *AmountDetail) SetVoidAmount(v Money)`

SetVoidAmount sets VoidAmount field to given value.

### HasVoidAmount

`func (o *AmountDetail) HasVoidAmount() bool`

HasVoidAmount returns a boolean if a field has been set.

### GetConfirmAmount

`func (o *AmountDetail) GetConfirmAmount() Money`

GetConfirmAmount returns the ConfirmAmount field if non-nil, zero value otherwise.

### GetConfirmAmountOk

`func (o *AmountDetail) GetConfirmAmountOk() (*Money, bool)`

GetConfirmAmountOk returns a tuple with the ConfirmAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmAmount

`func (o *AmountDetail) SetConfirmAmount(v Money)`

SetConfirmAmount sets ConfirmAmount field to given value.

### HasConfirmAmount

`func (o *AmountDetail) HasConfirmAmount() bool`

HasConfirmAmount returns a boolean if a field has been set.

### GetRefundAmount

`func (o *AmountDetail) GetRefundAmount() Money`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *AmountDetail) GetRefundAmountOk() (*Money, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *AmountDetail) SetRefundAmount(v Money)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *AmountDetail) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetChargebackAmount

`func (o *AmountDetail) GetChargebackAmount() Money`

GetChargebackAmount returns the ChargebackAmount field if non-nil, zero value otherwise.

### GetChargebackAmountOk

`func (o *AmountDetail) GetChargebackAmountOk() (*Money, bool)`

GetChargebackAmountOk returns a tuple with the ChargebackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargebackAmount

`func (o *AmountDetail) SetChargebackAmount(v Money)`

SetChargebackAmount sets ChargebackAmount field to given value.

### HasChargebackAmount

`func (o *AmountDetail) HasChargebackAmount() bool`

HasChargebackAmount returns a boolean if a field has been set.

### GetChargeAmount

`func (o *AmountDetail) GetChargeAmount() Money`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *AmountDetail) GetChargeAmountOk() (*Money, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *AmountDetail) SetChargeAmount(v Money)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *AmountDetail) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


