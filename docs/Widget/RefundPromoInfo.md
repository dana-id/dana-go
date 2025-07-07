# RefundPromoInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromoId** | **string** | Promotion identifier | 
**PromoName** | **string** | Promotion name | 
**PromoType** | **string** | Type of promotion | 
**RefundAmount** | [**Money**](Money.md) | Refund amount from this promotion. Contains value (amount including cents) and currency (code based on ISO) | 

## Methods

### NewRefundPromoInfo

`func NewRefundPromoInfo(promoId string, promoName string, promoType string, refundAmount Money, ) *RefundPromoInfo`

NewRefundPromoInfo instantiates a new RefundPromoInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundPromoInfoWithDefaults

`func NewRefundPromoInfoWithDefaults() *RefundPromoInfo`

NewRefundPromoInfoWithDefaults instantiates a new RefundPromoInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromoId

`func (o *RefundPromoInfo) GetPromoId() string`

GetPromoId returns the PromoId field if non-nil, zero value otherwise.

### GetPromoIdOk

`func (o *RefundPromoInfo) GetPromoIdOk() (*string, bool)`

GetPromoIdOk returns a tuple with the PromoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoId

`func (o *RefundPromoInfo) SetPromoId(v string)`

SetPromoId sets PromoId field to given value.


### GetPromoName

`func (o *RefundPromoInfo) GetPromoName() string`

GetPromoName returns the PromoName field if non-nil, zero value otherwise.

### GetPromoNameOk

`func (o *RefundPromoInfo) GetPromoNameOk() (*string, bool)`

GetPromoNameOk returns a tuple with the PromoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoName

`func (o *RefundPromoInfo) SetPromoName(v string)`

SetPromoName sets PromoName field to given value.


### GetPromoType

`func (o *RefundPromoInfo) GetPromoType() string`

GetPromoType returns the PromoType field if non-nil, zero value otherwise.

### GetPromoTypeOk

`func (o *RefundPromoInfo) GetPromoTypeOk() (*string, bool)`

GetPromoTypeOk returns a tuple with the PromoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoType

`func (o *RefundPromoInfo) SetPromoType(v string)`

SetPromoType sets PromoType field to given value.


### GetRefundAmount

`func (o *RefundPromoInfo) GetRefundAmount() Money`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *RefundPromoInfo) GetRefundAmountOk() (*Money, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *RefundPromoInfo) SetRefundAmount(v Money)`

SetRefundAmount sets RefundAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


