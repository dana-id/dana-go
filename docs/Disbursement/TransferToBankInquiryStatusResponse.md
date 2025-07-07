# TransferToBankInquiryStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **string** | Refer to response code list | 
**ResponseMessage** | **string** | Refer to response code list | 
**OriginalPartnerReferenceNo** | Pointer to **string** | Original transaction identifier on partner system | [optional] 
**OriginalReferenceNo** | Pointer to **string** | Original transaction identifier on DANA system | [optional] 
**OriginalExternalId** | Pointer to **string** | Original external identifier on header message | [optional] 
**ServiceCode** | **string** | Transaction type indicator is based on the service code of the original transaction request, value always 00 | [default to "00"]
**Amount** | Pointer to [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO  | [optional] 
**LatestTransactionStatus** | **string** | Status of latest transaction:<br /> 00 - Success<br /> 01 - Initiated<br /> 02 - Paying<br /> 03 - Pending<br /> 04 - Refunded<br /> 05 - Canceled<br /> 06 - Failed<br /> 07 - Not found  | 
**TransactionStatusDesc** | Pointer to **string** | Description of transaction status | [optional] 
**AdditionalInfo** | Pointer to **map[string]interface{}** | Additional information | [optional] 

## Methods

### NewTransferToBankInquiryStatusResponse

`func NewTransferToBankInquiryStatusResponse(responseCode string, responseMessage string, serviceCode string, latestTransactionStatus string, ) *TransferToBankInquiryStatusResponse`

NewTransferToBankInquiryStatusResponse instantiates a new TransferToBankInquiryStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferToBankInquiryStatusResponseWithDefaults

`func NewTransferToBankInquiryStatusResponseWithDefaults() *TransferToBankInquiryStatusResponse`

NewTransferToBankInquiryStatusResponseWithDefaults instantiates a new TransferToBankInquiryStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *TransferToBankInquiryStatusResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *TransferToBankInquiryStatusResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *TransferToBankInquiryStatusResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseMessage

`func (o *TransferToBankInquiryStatusResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *TransferToBankInquiryStatusResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *TransferToBankInquiryStatusResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetOriginalPartnerReferenceNo

`func (o *TransferToBankInquiryStatusResponse) GetOriginalPartnerReferenceNo() string`

GetOriginalPartnerReferenceNo returns the OriginalPartnerReferenceNo field if non-nil, zero value otherwise.

### GetOriginalPartnerReferenceNoOk

`func (o *TransferToBankInquiryStatusResponse) GetOriginalPartnerReferenceNoOk() (*string, bool)`

GetOriginalPartnerReferenceNoOk returns a tuple with the OriginalPartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPartnerReferenceNo

`func (o *TransferToBankInquiryStatusResponse) SetOriginalPartnerReferenceNo(v string)`

SetOriginalPartnerReferenceNo sets OriginalPartnerReferenceNo field to given value.

### HasOriginalPartnerReferenceNo

`func (o *TransferToBankInquiryStatusResponse) HasOriginalPartnerReferenceNo() bool`

HasOriginalPartnerReferenceNo returns a boolean if a field has been set.

### GetOriginalReferenceNo

`func (o *TransferToBankInquiryStatusResponse) GetOriginalReferenceNo() string`

GetOriginalReferenceNo returns the OriginalReferenceNo field if non-nil, zero value otherwise.

### GetOriginalReferenceNoOk

`func (o *TransferToBankInquiryStatusResponse) GetOriginalReferenceNoOk() (*string, bool)`

GetOriginalReferenceNoOk returns a tuple with the OriginalReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReferenceNo

`func (o *TransferToBankInquiryStatusResponse) SetOriginalReferenceNo(v string)`

SetOriginalReferenceNo sets OriginalReferenceNo field to given value.

### HasOriginalReferenceNo

`func (o *TransferToBankInquiryStatusResponse) HasOriginalReferenceNo() bool`

HasOriginalReferenceNo returns a boolean if a field has been set.

### GetOriginalExternalId

`func (o *TransferToBankInquiryStatusResponse) GetOriginalExternalId() string`

GetOriginalExternalId returns the OriginalExternalId field if non-nil, zero value otherwise.

### GetOriginalExternalIdOk

`func (o *TransferToBankInquiryStatusResponse) GetOriginalExternalIdOk() (*string, bool)`

GetOriginalExternalIdOk returns a tuple with the OriginalExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalExternalId

`func (o *TransferToBankInquiryStatusResponse) SetOriginalExternalId(v string)`

SetOriginalExternalId sets OriginalExternalId field to given value.

### HasOriginalExternalId

`func (o *TransferToBankInquiryStatusResponse) HasOriginalExternalId() bool`

HasOriginalExternalId returns a boolean if a field has been set.

### GetServiceCode

`func (o *TransferToBankInquiryStatusResponse) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *TransferToBankInquiryStatusResponse) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *TransferToBankInquiryStatusResponse) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.


### GetAmount

`func (o *TransferToBankInquiryStatusResponse) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferToBankInquiryStatusResponse) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferToBankInquiryStatusResponse) SetAmount(v Money)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransferToBankInquiryStatusResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetLatestTransactionStatus

`func (o *TransferToBankInquiryStatusResponse) GetLatestTransactionStatus() string`

GetLatestTransactionStatus returns the LatestTransactionStatus field if non-nil, zero value otherwise.

### GetLatestTransactionStatusOk

`func (o *TransferToBankInquiryStatusResponse) GetLatestTransactionStatusOk() (*string, bool)`

GetLatestTransactionStatusOk returns a tuple with the LatestTransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTransactionStatus

`func (o *TransferToBankInquiryStatusResponse) SetLatestTransactionStatus(v string)`

SetLatestTransactionStatus sets LatestTransactionStatus field to given value.


### GetTransactionStatusDesc

`func (o *TransferToBankInquiryStatusResponse) GetTransactionStatusDesc() string`

GetTransactionStatusDesc returns the TransactionStatusDesc field if non-nil, zero value otherwise.

### GetTransactionStatusDescOk

`func (o *TransferToBankInquiryStatusResponse) GetTransactionStatusDescOk() (*string, bool)`

GetTransactionStatusDescOk returns a tuple with the TransactionStatusDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatusDesc

`func (o *TransferToBankInquiryStatusResponse) SetTransactionStatusDesc(v string)`

SetTransactionStatusDesc sets TransactionStatusDesc field to given value.

### HasTransactionStatusDesc

`func (o *TransferToBankInquiryStatusResponse) HasTransactionStatusDesc() bool`

HasTransactionStatusDesc returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *TransferToBankInquiryStatusResponse) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *TransferToBankInquiryStatusResponse) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *TransferToBankInquiryStatusResponse) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *TransferToBankInquiryStatusResponse) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


