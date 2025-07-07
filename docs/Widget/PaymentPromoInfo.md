# PaymentPromoInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromoId** | **string** | Promotion identifier | 
**PromoName** | **string** | Promotion name | 
**PromoType** | **string** | Type of promotion | 
**SavingsAmount** | [**Money**](Money.md) | Savings amount from this promotion. Contains value (amount including cents) and currency (code based on ISO) | 

## Methods

### NewPaymentPromoInfo

`func NewPaymentPromoInfo(promoId string, promoName string, promoType string, savingsAmount Money, ) *PaymentPromoInfo`

NewPaymentPromoInfo instantiates a new PaymentPromoInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPromoInfoWithDefaults

`func NewPaymentPromoInfoWithDefaults() *PaymentPromoInfo`

NewPaymentPromoInfoWithDefaults instantiates a new PaymentPromoInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromoId

`func (o *PaymentPromoInfo) GetPromoId() string`

GetPromoId returns the PromoId field if non-nil, zero value otherwise.

### GetPromoIdOk

`func (o *PaymentPromoInfo) GetPromoIdOk() (*string, bool)`

GetPromoIdOk returns a tuple with the PromoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoId

`func (o *PaymentPromoInfo) SetPromoId(v string)`

SetPromoId sets PromoId field to given value.


### GetPromoName

`func (o *PaymentPromoInfo) GetPromoName() string`

GetPromoName returns the PromoName field if non-nil, zero value otherwise.

### GetPromoNameOk

`func (o *PaymentPromoInfo) GetPromoNameOk() (*string, bool)`

GetPromoNameOk returns a tuple with the PromoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoName

`func (o *PaymentPromoInfo) SetPromoName(v string)`

SetPromoName sets PromoName field to given value.


### GetPromoType

`func (o *PaymentPromoInfo) GetPromoType() string`

GetPromoType returns the PromoType field if non-nil, zero value otherwise.

### GetPromoTypeOk

`func (o *PaymentPromoInfo) GetPromoTypeOk() (*string, bool)`

GetPromoTypeOk returns a tuple with the PromoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoType

`func (o *PaymentPromoInfo) SetPromoType(v string)`

SetPromoType sets PromoType field to given value.


### GetSavingsAmount

`func (o *PaymentPromoInfo) GetSavingsAmount() Money`

GetSavingsAmount returns the SavingsAmount field if non-nil, zero value otherwise.

### GetSavingsAmountOk

`func (o *PaymentPromoInfo) GetSavingsAmountOk() (*Money, bool)`

GetSavingsAmountOk returns a tuple with the SavingsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingsAmount

`func (o *PaymentPromoInfo) SetSavingsAmount(v Money)`

SetSavingsAmount sets SavingsAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


