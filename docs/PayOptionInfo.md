# PayOptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayMethod** | **string** | Payment method name | 
**PayOption** | Pointer to **string** | Payment provider name | [optional] 
**PayAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TransAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**ChargeAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**PayOptionBillExtendInfo** | Pointer to **string** | Extended bill information for pay option | [optional] 
**ExtendInfo** | Pointer to **string** | Additional extend information | [optional] 
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


