# PromoInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromoAmount** | [**Money**](Money.md) |  | 
**PromoId** | **string** |  | 
**PromoType** | **string** |  | [default to "DIRECT_DISCOUNT"]

## Methods

### NewPromoInfo

`func NewPromoInfo(promoAmount Money, promoId string, promoType string, ) *PromoInfo`

NewPromoInfo instantiates a new PromoInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromoInfoWithDefaults

`func NewPromoInfoWithDefaults() *PromoInfo`

NewPromoInfoWithDefaults instantiates a new PromoInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromoAmount

`func (o *PromoInfo) GetPromoAmount() Money`

GetPromoAmount returns the PromoAmount field if non-nil, zero value otherwise.

### GetPromoAmountOk

`func (o *PromoInfo) GetPromoAmountOk() (*Money, bool)`

GetPromoAmountOk returns a tuple with the PromoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoAmount

`func (o *PromoInfo) SetPromoAmount(v Money)`

SetPromoAmount sets PromoAmount field to given value.


### GetPromoId

`func (o *PromoInfo) GetPromoId() string`

GetPromoId returns the PromoId field if non-nil, zero value otherwise.

### GetPromoIdOk

`func (o *PromoInfo) GetPromoIdOk() (*string, bool)`

GetPromoIdOk returns a tuple with the PromoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoId

`func (o *PromoInfo) SetPromoId(v string)`

SetPromoId sets PromoId field to given value.


### GetPromoType

`func (o *PromoInfo) GetPromoType() string`

GetPromoType returns the PromoType field if non-nil, zero value otherwise.

### GetPromoTypeOk

`func (o *PromoInfo) GetPromoTypeOk() (*string, bool)`

GetPromoTypeOk returns a tuple with the PromoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoType

`func (o *PromoInfo) SetPromoType(v string)`

SetPromoType sets PromoType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


