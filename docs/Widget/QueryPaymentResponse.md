# QueryPaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **string** | Refer to response code list:<br /> * 2005500 - Successful<br /> * 4005500 - Bad Request - Retry request with proper parameter<br /> * 4005501 - Invalid Field Format - Retry request with proper parameter<br /> * 4005502 - Invalid Mandatory Field - Retry request with proper parameter<br /> * 4015500 - Unauthorized. [reason] - Retry request with proper parameter<br /> * 4015501 - Invalid Token (B2B) - Retry request with proper parameter<br /> * 4045501 - Transaction Not Found - Try to create a new order<br /> * 4295500 - Too Many Requests - Retry request periodically<br /> * 5005500 - General Error - Retry request periodically<br /> * 5005501 - Internal Server Error - Retry request periodically<br />  | 
**ResponseMessage** | **string** | Refer to response code list  | 
**OriginalPartnerReferenceNo** | Pointer to **string** | Original transaction identifier on partner system | [optional] 
**OriginalReferenceNo** | Pointer to **string** | Original transaction identifier on DANA system | [optional] 
**OriginalExternalId** | Pointer to **string** | Original external identifier on header message | [optional] 
**ServiceCode** | **string** | Transaction type indicator:<br /> - IPG Cashier Pay - SNAP: 54<br /> - QRIS CPM (Acquirer) - SNAP: 60<br /> - QRIS MPM (Acquirer) - SNAP: 47<br /> - Payment Gateway: 54<br />  | [default to "54"]
**LatestTransactionStatus** | **string** | Status code:<br /> - 00 &#x3D; Success. Order has been successfully in final state and paid<br /> - 01 &#x3D; Initiated. Waiting for payment. Mark Payment as Pending<br /> - 02 &#x3D; Paying. The order is in process, not in final state, payment is success. Mark Payment as Success<br /> - 05 &#x3D; Cancelled. Order has been cancelled. Mark Payment as Failed<br /> - 07 &#x3D; Not found. Order is not found. Mark Payment as Failed<br />  | 
**TransactionStatusDesc** | Pointer to **string** | Description of transaction status | [optional] 
**OriginalResponseCode** | Pointer to **string** | Original response code | [optional] 
**OriginalResponseMessage** | Pointer to **string** | Original response message | [optional] 
**SessionId** | Pointer to **string** | Session identifier | [optional] 
**RequestID** | Pointer to **string** | Transaction request identifier | [optional] 
**TransAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**Amount** | Pointer to [**Money**](Money.md) |  | [optional] 
**FeeAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**PaidTime** | Pointer to **string** | Payment timestamp in format YYYY-MM-DDTHH:mm:ss+07:00 (Jakarta time) | [optional] 
**Title** | Pointer to **string** | Brief description of transaction | [optional] 
**AdditionalInfo** | Pointer to [**QueryPaymentResponseAdditionalInfo**](QueryPaymentResponseAdditionalInfo.md) |  | [optional] 

## Methods

### NewQueryPaymentResponse

`func NewQueryPaymentResponse(responseCode string, responseMessage string, serviceCode string, latestTransactionStatus string, ) *QueryPaymentResponse`

NewQueryPaymentResponse instantiates a new QueryPaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPaymentResponseWithDefaults

`func NewQueryPaymentResponseWithDefaults() *QueryPaymentResponse`

NewQueryPaymentResponseWithDefaults instantiates a new QueryPaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *QueryPaymentResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *QueryPaymentResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *QueryPaymentResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseMessage

`func (o *QueryPaymentResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *QueryPaymentResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *QueryPaymentResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetOriginalPartnerReferenceNo

`func (o *QueryPaymentResponse) GetOriginalPartnerReferenceNo() string`

GetOriginalPartnerReferenceNo returns the OriginalPartnerReferenceNo field if non-nil, zero value otherwise.

### GetOriginalPartnerReferenceNoOk

`func (o *QueryPaymentResponse) GetOriginalPartnerReferenceNoOk() (*string, bool)`

GetOriginalPartnerReferenceNoOk returns a tuple with the OriginalPartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPartnerReferenceNo

`func (o *QueryPaymentResponse) SetOriginalPartnerReferenceNo(v string)`

SetOriginalPartnerReferenceNo sets OriginalPartnerReferenceNo field to given value.

### HasOriginalPartnerReferenceNo

`func (o *QueryPaymentResponse) HasOriginalPartnerReferenceNo() bool`

HasOriginalPartnerReferenceNo returns a boolean if a field has been set.

### GetOriginalReferenceNo

`func (o *QueryPaymentResponse) GetOriginalReferenceNo() string`

GetOriginalReferenceNo returns the OriginalReferenceNo field if non-nil, zero value otherwise.

### GetOriginalReferenceNoOk

`func (o *QueryPaymentResponse) GetOriginalReferenceNoOk() (*string, bool)`

GetOriginalReferenceNoOk returns a tuple with the OriginalReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReferenceNo

`func (o *QueryPaymentResponse) SetOriginalReferenceNo(v string)`

SetOriginalReferenceNo sets OriginalReferenceNo field to given value.

### HasOriginalReferenceNo

`func (o *QueryPaymentResponse) HasOriginalReferenceNo() bool`

HasOriginalReferenceNo returns a boolean if a field has been set.

### GetOriginalExternalId

`func (o *QueryPaymentResponse) GetOriginalExternalId() string`

GetOriginalExternalId returns the OriginalExternalId field if non-nil, zero value otherwise.

### GetOriginalExternalIdOk

`func (o *QueryPaymentResponse) GetOriginalExternalIdOk() (*string, bool)`

GetOriginalExternalIdOk returns a tuple with the OriginalExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalExternalId

`func (o *QueryPaymentResponse) SetOriginalExternalId(v string)`

SetOriginalExternalId sets OriginalExternalId field to given value.

### HasOriginalExternalId

`func (o *QueryPaymentResponse) HasOriginalExternalId() bool`

HasOriginalExternalId returns a boolean if a field has been set.

### GetServiceCode

`func (o *QueryPaymentResponse) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *QueryPaymentResponse) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *QueryPaymentResponse) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.


### GetLatestTransactionStatus

`func (o *QueryPaymentResponse) GetLatestTransactionStatus() string`

GetLatestTransactionStatus returns the LatestTransactionStatus field if non-nil, zero value otherwise.

### GetLatestTransactionStatusOk

`func (o *QueryPaymentResponse) GetLatestTransactionStatusOk() (*string, bool)`

GetLatestTransactionStatusOk returns a tuple with the LatestTransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTransactionStatus

`func (o *QueryPaymentResponse) SetLatestTransactionStatus(v string)`

SetLatestTransactionStatus sets LatestTransactionStatus field to given value.


### GetTransactionStatusDesc

`func (o *QueryPaymentResponse) GetTransactionStatusDesc() string`

GetTransactionStatusDesc returns the TransactionStatusDesc field if non-nil, zero value otherwise.

### GetTransactionStatusDescOk

`func (o *QueryPaymentResponse) GetTransactionStatusDescOk() (*string, bool)`

GetTransactionStatusDescOk returns a tuple with the TransactionStatusDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatusDesc

`func (o *QueryPaymentResponse) SetTransactionStatusDesc(v string)`

SetTransactionStatusDesc sets TransactionStatusDesc field to given value.

### HasTransactionStatusDesc

`func (o *QueryPaymentResponse) HasTransactionStatusDesc() bool`

HasTransactionStatusDesc returns a boolean if a field has been set.

### GetOriginalResponseCode

`func (o *QueryPaymentResponse) GetOriginalResponseCode() string`

GetOriginalResponseCode returns the OriginalResponseCode field if non-nil, zero value otherwise.

### GetOriginalResponseCodeOk

`func (o *QueryPaymentResponse) GetOriginalResponseCodeOk() (*string, bool)`

GetOriginalResponseCodeOk returns a tuple with the OriginalResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalResponseCode

`func (o *QueryPaymentResponse) SetOriginalResponseCode(v string)`

SetOriginalResponseCode sets OriginalResponseCode field to given value.

### HasOriginalResponseCode

`func (o *QueryPaymentResponse) HasOriginalResponseCode() bool`

HasOriginalResponseCode returns a boolean if a field has been set.

### GetOriginalResponseMessage

`func (o *QueryPaymentResponse) GetOriginalResponseMessage() string`

GetOriginalResponseMessage returns the OriginalResponseMessage field if non-nil, zero value otherwise.

### GetOriginalResponseMessageOk

`func (o *QueryPaymentResponse) GetOriginalResponseMessageOk() (*string, bool)`

GetOriginalResponseMessageOk returns a tuple with the OriginalResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalResponseMessage

`func (o *QueryPaymentResponse) SetOriginalResponseMessage(v string)`

SetOriginalResponseMessage sets OriginalResponseMessage field to given value.

### HasOriginalResponseMessage

`func (o *QueryPaymentResponse) HasOriginalResponseMessage() bool`

HasOriginalResponseMessage returns a boolean if a field has been set.

### GetSessionId

`func (o *QueryPaymentResponse) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *QueryPaymentResponse) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *QueryPaymentResponse) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *QueryPaymentResponse) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetRequestID

`func (o *QueryPaymentResponse) GetRequestID() string`

GetRequestID returns the RequestID field if non-nil, zero value otherwise.

### GetRequestIDOk

`func (o *QueryPaymentResponse) GetRequestIDOk() (*string, bool)`

GetRequestIDOk returns a tuple with the RequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestID

`func (o *QueryPaymentResponse) SetRequestID(v string)`

SetRequestID sets RequestID field to given value.

### HasRequestID

`func (o *QueryPaymentResponse) HasRequestID() bool`

HasRequestID returns a boolean if a field has been set.

### GetTransAmount

`func (o *QueryPaymentResponse) GetTransAmount() Money`

GetTransAmount returns the TransAmount field if non-nil, zero value otherwise.

### GetTransAmountOk

`func (o *QueryPaymentResponse) GetTransAmountOk() (*Money, bool)`

GetTransAmountOk returns a tuple with the TransAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransAmount

`func (o *QueryPaymentResponse) SetTransAmount(v Money)`

SetTransAmount sets TransAmount field to given value.

### HasTransAmount

`func (o *QueryPaymentResponse) HasTransAmount() bool`

HasTransAmount returns a boolean if a field has been set.

### GetAmount

`func (o *QueryPaymentResponse) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QueryPaymentResponse) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QueryPaymentResponse) SetAmount(v Money)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QueryPaymentResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFeeAmount

`func (o *QueryPaymentResponse) GetFeeAmount() Money`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *QueryPaymentResponse) GetFeeAmountOk() (*Money, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *QueryPaymentResponse) SetFeeAmount(v Money)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *QueryPaymentResponse) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetPaidTime

`func (o *QueryPaymentResponse) GetPaidTime() string`

GetPaidTime returns the PaidTime field if non-nil, zero value otherwise.

### GetPaidTimeOk

`func (o *QueryPaymentResponse) GetPaidTimeOk() (*string, bool)`

GetPaidTimeOk returns a tuple with the PaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTime

`func (o *QueryPaymentResponse) SetPaidTime(v string)`

SetPaidTime sets PaidTime field to given value.

### HasPaidTime

`func (o *QueryPaymentResponse) HasPaidTime() bool`

HasPaidTime returns a boolean if a field has been set.

### GetTitle

`func (o *QueryPaymentResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QueryPaymentResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QueryPaymentResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *QueryPaymentResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *QueryPaymentResponse) GetAdditionalInfo() QueryPaymentResponseAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *QueryPaymentResponse) GetAdditionalInfoOk() (*QueryPaymentResponseAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *QueryPaymentResponse) SetAdditionalInfo(v QueryPaymentResponseAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *QueryPaymentResponse) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


