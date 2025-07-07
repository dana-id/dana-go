# DanaAccountInquiryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerReferenceNo** | Pointer to **string** | Unique transaction identifier on partner system which assigned to each transaction<br /> Notes:<br /> If the partner receives a timeout or an unexpected response from DANA and partner expects to perform retry request to DANA, please use the partnerReferenceNo that is the same as the one used in the transaction request process before  | [optional] 
**CustomerNumber** | Pointer to **string** | Customer account number, in format 628xxx | [optional] 
**Amount** | [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO  | 
**TransactionDate** | Pointer to **string** | Transaction date, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | [optional] 
**AdditionalInfo** | [**DanaAccountInquiryRequestAdditionalInfo**](DanaAccountInquiryRequestAdditionalInfo.md) |  | 

## Methods

### NewDanaAccountInquiryRequest

`func NewDanaAccountInquiryRequest(amount Money, additionalInfo DanaAccountInquiryRequestAdditionalInfo, ) *DanaAccountInquiryRequest`

NewDanaAccountInquiryRequest instantiates a new DanaAccountInquiryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDanaAccountInquiryRequestWithDefaults

`func NewDanaAccountInquiryRequestWithDefaults() *DanaAccountInquiryRequest`

NewDanaAccountInquiryRequestWithDefaults instantiates a new DanaAccountInquiryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerReferenceNo

`func (o *DanaAccountInquiryRequest) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *DanaAccountInquiryRequest) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *DanaAccountInquiryRequest) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.

### HasPartnerReferenceNo

`func (o *DanaAccountInquiryRequest) HasPartnerReferenceNo() bool`

HasPartnerReferenceNo returns a boolean if a field has been set.

### GetCustomerNumber

`func (o *DanaAccountInquiryRequest) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *DanaAccountInquiryRequest) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *DanaAccountInquiryRequest) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.

### HasCustomerNumber

`func (o *DanaAccountInquiryRequest) HasCustomerNumber() bool`

HasCustomerNumber returns a boolean if a field has been set.

### GetAmount

`func (o *DanaAccountInquiryRequest) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DanaAccountInquiryRequest) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DanaAccountInquiryRequest) SetAmount(v Money)`

SetAmount sets Amount field to given value.


### GetTransactionDate

`func (o *DanaAccountInquiryRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *DanaAccountInquiryRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *DanaAccountInquiryRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *DanaAccountInquiryRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *DanaAccountInquiryRequest) GetAdditionalInfo() DanaAccountInquiryRequestAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *DanaAccountInquiryRequest) GetAdditionalInfoOk() (*DanaAccountInquiryRequestAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *DanaAccountInquiryRequest) SetAdditionalInfo(v DanaAccountInquiryRequestAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


