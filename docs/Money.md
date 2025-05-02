# Money

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Value of amount. Following ISO-4217, for IDR the value includes 2 decimal digits separated with point e.g. IDR 10.000,- will be placed with 10000.00 | 
**Currency** | **string** | Currency of money following ISO-4217 | 

## Methods

### NewMoney

`func NewMoney(value string, currency string, ) *Money`

NewMoney instantiates a new Money object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyWithDefaults

`func NewMoneyWithDefaults() *Money`

NewMoneyWithDefaults instantiates a new Money object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Money) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Money) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Money) SetValue(v string)`

SetValue sets Value field to given value.


### GetCurrency

`func (o *Money) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Money) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Money) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


