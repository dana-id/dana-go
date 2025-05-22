# CancelOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalPartnerReferenceNo** | **string** | Original transaction identifier on partner system | 
**OriginalReferenceNo** | Pointer to **string** | Original transaction identifier on DANA system | [optional] 
**OriginalExternalId** | Pointer to **string** | Original external identifier on header message | [optional] 
**MerchantId** | **string** | Merchant identifier that is unique per each merchant | 
**SubMerchantId** | Pointer to **string** | Information of sub merchant identifier | [optional] 
**Reason** | Pointer to **string** | Cancellation reason | [optional] 
**ExternalStoreId** | Pointer to **string** | Store identifier to indicate to which store this payment belongs to | [optional] 
**Amount** | Pointer to [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | [optional] 
**AdditionalInfo** | Pointer to **map[string]interface{}** | Additional information | [optional] 

## Methods

### NewCancelOrderRequest

`func NewCancelOrderRequest(originalPartnerReferenceNo string, merchantId string, ) *CancelOrderRequest`

NewCancelOrderRequest instantiates a new CancelOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelOrderRequestWithDefaults

`func NewCancelOrderRequestWithDefaults() *CancelOrderRequest`

NewCancelOrderRequestWithDefaults instantiates a new CancelOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalPartnerReferenceNo

`func (o *CancelOrderRequest) GetOriginalPartnerReferenceNo() string`

GetOriginalPartnerReferenceNo returns the OriginalPartnerReferenceNo field if non-nil, zero value otherwise.

### GetOriginalPartnerReferenceNoOk

`func (o *CancelOrderRequest) GetOriginalPartnerReferenceNoOk() (*string, bool)`

GetOriginalPartnerReferenceNoOk returns a tuple with the OriginalPartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPartnerReferenceNo

`func (o *CancelOrderRequest) SetOriginalPartnerReferenceNo(v string)`

SetOriginalPartnerReferenceNo sets OriginalPartnerReferenceNo field to given value.


### GetOriginalReferenceNo

`func (o *CancelOrderRequest) GetOriginalReferenceNo() string`

GetOriginalReferenceNo returns the OriginalReferenceNo field if non-nil, zero value otherwise.

### GetOriginalReferenceNoOk

`func (o *CancelOrderRequest) GetOriginalReferenceNoOk() (*string, bool)`

GetOriginalReferenceNoOk returns a tuple with the OriginalReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReferenceNo

`func (o *CancelOrderRequest) SetOriginalReferenceNo(v string)`

SetOriginalReferenceNo sets OriginalReferenceNo field to given value.

### HasOriginalReferenceNo

`func (o *CancelOrderRequest) HasOriginalReferenceNo() bool`

HasOriginalReferenceNo returns a boolean if a field has been set.

### GetOriginalExternalId

`func (o *CancelOrderRequest) GetOriginalExternalId() string`

GetOriginalExternalId returns the OriginalExternalId field if non-nil, zero value otherwise.

### GetOriginalExternalIdOk

`func (o *CancelOrderRequest) GetOriginalExternalIdOk() (*string, bool)`

GetOriginalExternalIdOk returns a tuple with the OriginalExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalExternalId

`func (o *CancelOrderRequest) SetOriginalExternalId(v string)`

SetOriginalExternalId sets OriginalExternalId field to given value.

### HasOriginalExternalId

`func (o *CancelOrderRequest) HasOriginalExternalId() bool`

HasOriginalExternalId returns a boolean if a field has been set.

### GetMerchantId

`func (o *CancelOrderRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CancelOrderRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CancelOrderRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetSubMerchantId

`func (o *CancelOrderRequest) GetSubMerchantId() string`

GetSubMerchantId returns the SubMerchantId field if non-nil, zero value otherwise.

### GetSubMerchantIdOk

`func (o *CancelOrderRequest) GetSubMerchantIdOk() (*string, bool)`

GetSubMerchantIdOk returns a tuple with the SubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantId

`func (o *CancelOrderRequest) SetSubMerchantId(v string)`

SetSubMerchantId sets SubMerchantId field to given value.

### HasSubMerchantId

`func (o *CancelOrderRequest) HasSubMerchantId() bool`

HasSubMerchantId returns a boolean if a field has been set.

### GetReason

`func (o *CancelOrderRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CancelOrderRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CancelOrderRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CancelOrderRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetExternalStoreId

`func (o *CancelOrderRequest) GetExternalStoreId() string`

GetExternalStoreId returns the ExternalStoreId field if non-nil, zero value otherwise.

### GetExternalStoreIdOk

`func (o *CancelOrderRequest) GetExternalStoreIdOk() (*string, bool)`

GetExternalStoreIdOk returns a tuple with the ExternalStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoreId

`func (o *CancelOrderRequest) SetExternalStoreId(v string)`

SetExternalStoreId sets ExternalStoreId field to given value.

### HasExternalStoreId

`func (o *CancelOrderRequest) HasExternalStoreId() bool`

HasExternalStoreId returns a boolean if a field has been set.

### GetAmount

`func (o *CancelOrderRequest) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CancelOrderRequest) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CancelOrderRequest) SetAmount(v Money)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CancelOrderRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *CancelOrderRequest) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *CancelOrderRequest) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *CancelOrderRequest) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *CancelOrderRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


