# TransferToDanaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerReferenceNo** | **string** | Unique transaction identifier on partner system which assigned to each transaction<br /> Notes:<br /> If the partner receives a timeout or an unexpected response from DANA and partner expects to perform retry request to DANA, please use the partnerReferenceNo that is the same as the one used in the transaction request process before  | 
**CustomerNumber** | Pointer to **string** | Customer account number, in format 628xxx | [optional] 
**Amount** | [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO  | 
**FeeAmount** | [**Money**](Money.md) | Fee amount. Contains two sub-fields:<br /> 1. Value: Amount, including the cents<br /> 2. Currency: Currency code based on ISO  | 
**TransactionDate** | Pointer to **string** | Transaction date, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | [optional] 
**SessionId** | Pointer to **string** | Session identifier | [optional] 
**CategoryId** | Pointer to **float32** | Category identifier | [optional] 
**Notes** | Pointer to **string** | Transaction notes | [optional] 
**AdditionalInfo** | [**TransferToDanaRequestAdditionalInfo**](TransferToDanaRequestAdditionalInfo.md) |  | 

## Methods

### NewTransferToDanaRequest

`func NewTransferToDanaRequest(partnerReferenceNo string, amount Money, feeAmount Money, additionalInfo TransferToDanaRequestAdditionalInfo, ) *TransferToDanaRequest`

NewTransferToDanaRequest instantiates a new TransferToDanaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferToDanaRequestWithDefaults

`func NewTransferToDanaRequestWithDefaults() *TransferToDanaRequest`

NewTransferToDanaRequestWithDefaults instantiates a new TransferToDanaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerReferenceNo

`func (o *TransferToDanaRequest) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *TransferToDanaRequest) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *TransferToDanaRequest) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.


### GetCustomerNumber

`func (o *TransferToDanaRequest) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *TransferToDanaRequest) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *TransferToDanaRequest) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.

### HasCustomerNumber

`func (o *TransferToDanaRequest) HasCustomerNumber() bool`

HasCustomerNumber returns a boolean if a field has been set.

### GetAmount

`func (o *TransferToDanaRequest) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferToDanaRequest) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferToDanaRequest) SetAmount(v Money)`

SetAmount sets Amount field to given value.


### GetFeeAmount

`func (o *TransferToDanaRequest) GetFeeAmount() Money`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *TransferToDanaRequest) GetFeeAmountOk() (*Money, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *TransferToDanaRequest) SetFeeAmount(v Money)`

SetFeeAmount sets FeeAmount field to given value.


### GetTransactionDate

`func (o *TransferToDanaRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *TransferToDanaRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *TransferToDanaRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *TransferToDanaRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetSessionId

`func (o *TransferToDanaRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *TransferToDanaRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *TransferToDanaRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *TransferToDanaRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetCategoryId

`func (o *TransferToDanaRequest) GetCategoryId() float32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *TransferToDanaRequest) GetCategoryIdOk() (*float32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *TransferToDanaRequest) SetCategoryId(v float32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *TransferToDanaRequest) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetNotes

`func (o *TransferToDanaRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TransferToDanaRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TransferToDanaRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TransferToDanaRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *TransferToDanaRequest) GetAdditionalInfo() TransferToDanaRequestAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *TransferToDanaRequest) GetAdditionalInfoOk() (*TransferToDanaRequestAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *TransferToDanaRequest) SetAdditionalInfo(v TransferToDanaRequestAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


