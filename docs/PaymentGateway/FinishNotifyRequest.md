# FinishNotifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalPartnerReferenceNo** | **string** | Original transaction identifier on DANA system | 
**OriginalReferenceNo** | **string** | Original transaction identifier on partner system | 
**OriginalExternalId** | Pointer to **string** | Original external identifier on header message | [optional] 
**MerchantId** | **string** | Unique identifier per each merchant | 
**SubMerchantId** | Pointer to **string** | Information of sub merchant identifier | [optional] 
**Amount** | [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | 
**LatestTransactionStatus** | **string** | Transaction status code:<br /> - 00 &#x3D; Success, the order has been paid<br /> - 05 &#x3D; Cancelled, the order has been closed because it is expired<br />  | 
**TransactionStatusDesc** | Pointer to **string** | Description of transaction status | [optional] 
**CreatedTime** | **string** | Transaction created time, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | 
**FinishedTime** | **string** | Transaction finished time, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | 
**ExternalStoreId** | Pointer to **string** | Store identifier to indicate to which store this payment belongs to | [optional] 
**AdditionalInfo** | Pointer to [**FinishNotifyRequestAdditionalInfo**](FinishNotifyRequestAdditionalInfo.md) | Additional information | [optional] 

## Methods

### NewFinishNotifyRequest

`func NewFinishNotifyRequest(originalPartnerReferenceNo string, originalReferenceNo string, merchantId string, amount Money, latestTransactionStatus string, createdTime string, finishedTime string, ) *FinishNotifyRequest`

NewFinishNotifyRequest instantiates a new FinishNotifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinishNotifyRequestWithDefaults

`func NewFinishNotifyRequestWithDefaults() *FinishNotifyRequest`

NewFinishNotifyRequestWithDefaults instantiates a new FinishNotifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalPartnerReferenceNo

`func (o *FinishNotifyRequest) GetOriginalPartnerReferenceNo() string`

GetOriginalPartnerReferenceNo returns the OriginalPartnerReferenceNo field if non-nil, zero value otherwise.

### GetOriginalPartnerReferenceNoOk

`func (o *FinishNotifyRequest) GetOriginalPartnerReferenceNoOk() (*string, bool)`

GetOriginalPartnerReferenceNoOk returns a tuple with the OriginalPartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPartnerReferenceNo

`func (o *FinishNotifyRequest) SetOriginalPartnerReferenceNo(v string)`

SetOriginalPartnerReferenceNo sets OriginalPartnerReferenceNo field to given value.


### GetOriginalReferenceNo

`func (o *FinishNotifyRequest) GetOriginalReferenceNo() string`

GetOriginalReferenceNo returns the OriginalReferenceNo field if non-nil, zero value otherwise.

### GetOriginalReferenceNoOk

`func (o *FinishNotifyRequest) GetOriginalReferenceNoOk() (*string, bool)`

GetOriginalReferenceNoOk returns a tuple with the OriginalReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReferenceNo

`func (o *FinishNotifyRequest) SetOriginalReferenceNo(v string)`

SetOriginalReferenceNo sets OriginalReferenceNo field to given value.


### GetOriginalExternalId

`func (o *FinishNotifyRequest) GetOriginalExternalId() string`

GetOriginalExternalId returns the OriginalExternalId field if non-nil, zero value otherwise.

### GetOriginalExternalIdOk

`func (o *FinishNotifyRequest) GetOriginalExternalIdOk() (*string, bool)`

GetOriginalExternalIdOk returns a tuple with the OriginalExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalExternalId

`func (o *FinishNotifyRequest) SetOriginalExternalId(v string)`

SetOriginalExternalId sets OriginalExternalId field to given value.

### HasOriginalExternalId

`func (o *FinishNotifyRequest) HasOriginalExternalId() bool`

HasOriginalExternalId returns a boolean if a field has been set.

### GetMerchantId

`func (o *FinishNotifyRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *FinishNotifyRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *FinishNotifyRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetSubMerchantId

`func (o *FinishNotifyRequest) GetSubMerchantId() string`

GetSubMerchantId returns the SubMerchantId field if non-nil, zero value otherwise.

### GetSubMerchantIdOk

`func (o *FinishNotifyRequest) GetSubMerchantIdOk() (*string, bool)`

GetSubMerchantIdOk returns a tuple with the SubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantId

`func (o *FinishNotifyRequest) SetSubMerchantId(v string)`

SetSubMerchantId sets SubMerchantId field to given value.

### HasSubMerchantId

`func (o *FinishNotifyRequest) HasSubMerchantId() bool`

HasSubMerchantId returns a boolean if a field has been set.

### GetAmount

`func (o *FinishNotifyRequest) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FinishNotifyRequest) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FinishNotifyRequest) SetAmount(v Money)`

SetAmount sets Amount field to given value.


### GetLatestTransactionStatus

`func (o *FinishNotifyRequest) GetLatestTransactionStatus() string`

GetLatestTransactionStatus returns the LatestTransactionStatus field if non-nil, zero value otherwise.

### GetLatestTransactionStatusOk

`func (o *FinishNotifyRequest) GetLatestTransactionStatusOk() (*string, bool)`

GetLatestTransactionStatusOk returns a tuple with the LatestTransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTransactionStatus

`func (o *FinishNotifyRequest) SetLatestTransactionStatus(v string)`

SetLatestTransactionStatus sets LatestTransactionStatus field to given value.


### GetTransactionStatusDesc

`func (o *FinishNotifyRequest) GetTransactionStatusDesc() string`

GetTransactionStatusDesc returns the TransactionStatusDesc field if non-nil, zero value otherwise.

### GetTransactionStatusDescOk

`func (o *FinishNotifyRequest) GetTransactionStatusDescOk() (*string, bool)`

GetTransactionStatusDescOk returns a tuple with the TransactionStatusDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatusDesc

`func (o *FinishNotifyRequest) SetTransactionStatusDesc(v string)`

SetTransactionStatusDesc sets TransactionStatusDesc field to given value.

### HasTransactionStatusDesc

`func (o *FinishNotifyRequest) HasTransactionStatusDesc() bool`

HasTransactionStatusDesc returns a boolean if a field has been set.

### GetCreatedTime

`func (o *FinishNotifyRequest) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *FinishNotifyRequest) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *FinishNotifyRequest) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.


### GetFinishedTime

`func (o *FinishNotifyRequest) GetFinishedTime() string`

GetFinishedTime returns the FinishedTime field if non-nil, zero value otherwise.

### GetFinishedTimeOk

`func (o *FinishNotifyRequest) GetFinishedTimeOk() (*string, bool)`

GetFinishedTimeOk returns a tuple with the FinishedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedTime

`func (o *FinishNotifyRequest) SetFinishedTime(v string)`

SetFinishedTime sets FinishedTime field to given value.


### GetExternalStoreId

`func (o *FinishNotifyRequest) GetExternalStoreId() string`

GetExternalStoreId returns the ExternalStoreId field if non-nil, zero value otherwise.

### GetExternalStoreIdOk

`func (o *FinishNotifyRequest) GetExternalStoreIdOk() (*string, bool)`

GetExternalStoreIdOk returns a tuple with the ExternalStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoreId

`func (o *FinishNotifyRequest) SetExternalStoreId(v string)`

SetExternalStoreId sets ExternalStoreId field to given value.

### HasExternalStoreId

`func (o *FinishNotifyRequest) HasExternalStoreId() bool`

HasExternalStoreId returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *FinishNotifyRequest) GetAdditionalInfo() FinishNotifyRequestAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *FinishNotifyRequest) GetAdditionalInfoOk() (*FinishNotifyRequestAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *FinishNotifyRequest) SetAdditionalInfo(v FinishNotifyRequestAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *FinishNotifyRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


