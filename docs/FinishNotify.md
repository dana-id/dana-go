# FinishNotify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalPartnerReferenceNo** | **string** | Original transaction identifier on DANA system | 
**OriginalReferenceNo** | **string** | Original transaction identifier on partner system | 
**OriginalExternalId** | Pointer to **string** | Original external identifier on header message | [optional] 
**MerchantId** | **string** | Unique identifier for each merchant | 
**SubMerchantId** | Pointer to **string** | Sub merchant identifier | [optional] 
**Amount** | [**Money**](Money.md) |  | 
**LatestTransactionStatus** | **string** | Transaction status code:<br /> - 00 &#x3D; Success<br /> - 05 &#x3D; Cancelled (expired)<br />  | 
**TransactionStatusDesc** | Pointer to **string** | Description of transaction status | [optional] 
**CreatedTime** | **string** | Transaction creation time (GMT+7, Jakarta) | 
**FinishedTime** | **string** | Transaction completion time (GMT+7, Jakarta) | 
**ExternalStoreId** | Pointer to **string** | Store identifier | [optional] 
**AdditionalInfo** | Pointer to [**FinishNotifyAdditionalInfo**](FinishNotifyAdditionalInfo.md) |  | [optional] 

## Methods

### NewFinishNotify

`func NewFinishNotify(originalPartnerReferenceNo string, originalReferenceNo string, merchantId string, amount Money, latestTransactionStatus string, createdTime string, finishedTime string, ) *FinishNotify`

NewFinishNotify instantiates a new FinishNotify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinishNotifyWithDefaults

`func NewFinishNotifyWithDefaults() *FinishNotify`

NewFinishNotifyWithDefaults instantiates a new FinishNotify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalPartnerReferenceNo

`func (o *FinishNotify) GetOriginalPartnerReferenceNo() string`

GetOriginalPartnerReferenceNo returns the OriginalPartnerReferenceNo field if non-nil, zero value otherwise.

### GetOriginalPartnerReferenceNoOk

`func (o *FinishNotify) GetOriginalPartnerReferenceNoOk() (*string, bool)`

GetOriginalPartnerReferenceNoOk returns a tuple with the OriginalPartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPartnerReferenceNo

`func (o *FinishNotify) SetOriginalPartnerReferenceNo(v string)`

SetOriginalPartnerReferenceNo sets OriginalPartnerReferenceNo field to given value.


### GetOriginalReferenceNo

`func (o *FinishNotify) GetOriginalReferenceNo() string`

GetOriginalReferenceNo returns the OriginalReferenceNo field if non-nil, zero value otherwise.

### GetOriginalReferenceNoOk

`func (o *FinishNotify) GetOriginalReferenceNoOk() (*string, bool)`

GetOriginalReferenceNoOk returns a tuple with the OriginalReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReferenceNo

`func (o *FinishNotify) SetOriginalReferenceNo(v string)`

SetOriginalReferenceNo sets OriginalReferenceNo field to given value.


### GetOriginalExternalId

`func (o *FinishNotify) GetOriginalExternalId() string`

GetOriginalExternalId returns the OriginalExternalId field if non-nil, zero value otherwise.

### GetOriginalExternalIdOk

`func (o *FinishNotify) GetOriginalExternalIdOk() (*string, bool)`

GetOriginalExternalIdOk returns a tuple with the OriginalExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalExternalId

`func (o *FinishNotify) SetOriginalExternalId(v string)`

SetOriginalExternalId sets OriginalExternalId field to given value.

### HasOriginalExternalId

`func (o *FinishNotify) HasOriginalExternalId() bool`

HasOriginalExternalId returns a boolean if a field has been set.

### GetMerchantId

`func (o *FinishNotify) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *FinishNotify) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *FinishNotify) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetSubMerchantId

`func (o *FinishNotify) GetSubMerchantId() string`

GetSubMerchantId returns the SubMerchantId field if non-nil, zero value otherwise.

### GetSubMerchantIdOk

`func (o *FinishNotify) GetSubMerchantIdOk() (*string, bool)`

GetSubMerchantIdOk returns a tuple with the SubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantId

`func (o *FinishNotify) SetSubMerchantId(v string)`

SetSubMerchantId sets SubMerchantId field to given value.

### HasSubMerchantId

`func (o *FinishNotify) HasSubMerchantId() bool`

HasSubMerchantId returns a boolean if a field has been set.

### GetAmount

`func (o *FinishNotify) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FinishNotify) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FinishNotify) SetAmount(v Money)`

SetAmount sets Amount field to given value.


### GetLatestTransactionStatus

`func (o *FinishNotify) GetLatestTransactionStatus() string`

GetLatestTransactionStatus returns the LatestTransactionStatus field if non-nil, zero value otherwise.

### GetLatestTransactionStatusOk

`func (o *FinishNotify) GetLatestTransactionStatusOk() (*string, bool)`

GetLatestTransactionStatusOk returns a tuple with the LatestTransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTransactionStatus

`func (o *FinishNotify) SetLatestTransactionStatus(v string)`

SetLatestTransactionStatus sets LatestTransactionStatus field to given value.


### GetTransactionStatusDesc

`func (o *FinishNotify) GetTransactionStatusDesc() string`

GetTransactionStatusDesc returns the TransactionStatusDesc field if non-nil, zero value otherwise.

### GetTransactionStatusDescOk

`func (o *FinishNotify) GetTransactionStatusDescOk() (*string, bool)`

GetTransactionStatusDescOk returns a tuple with the TransactionStatusDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatusDesc

`func (o *FinishNotify) SetTransactionStatusDesc(v string)`

SetTransactionStatusDesc sets TransactionStatusDesc field to given value.

### HasTransactionStatusDesc

`func (o *FinishNotify) HasTransactionStatusDesc() bool`

HasTransactionStatusDesc returns a boolean if a field has been set.

### GetCreatedTime

`func (o *FinishNotify) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *FinishNotify) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *FinishNotify) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.


### GetFinishedTime

`func (o *FinishNotify) GetFinishedTime() string`

GetFinishedTime returns the FinishedTime field if non-nil, zero value otherwise.

### GetFinishedTimeOk

`func (o *FinishNotify) GetFinishedTimeOk() (*string, bool)`

GetFinishedTimeOk returns a tuple with the FinishedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedTime

`func (o *FinishNotify) SetFinishedTime(v string)`

SetFinishedTime sets FinishedTime field to given value.


### GetExternalStoreId

`func (o *FinishNotify) GetExternalStoreId() string`

GetExternalStoreId returns the ExternalStoreId field if non-nil, zero value otherwise.

### GetExternalStoreIdOk

`func (o *FinishNotify) GetExternalStoreIdOk() (*string, bool)`

GetExternalStoreIdOk returns a tuple with the ExternalStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoreId

`func (o *FinishNotify) SetExternalStoreId(v string)`

SetExternalStoreId sets ExternalStoreId field to given value.

### HasExternalStoreId

`func (o *FinishNotify) HasExternalStoreId() bool`

HasExternalStoreId returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *FinishNotify) GetAdditionalInfo() FinishNotifyAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *FinishNotify) GetAdditionalInfoOk() (*FinishNotifyAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *FinishNotify) SetAdditionalInfo(v FinishNotifyAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *FinishNotify) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


