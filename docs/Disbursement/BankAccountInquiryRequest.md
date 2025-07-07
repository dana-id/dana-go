# BankAccountInquiryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerReferenceNo** | Pointer to **string** | Unique transaction identifier on partner system which assigned to each transaction<br /> Notes:<br /> If the partner receives a timeout or an unexpected response from DANA and partner expects to perform retry request to DANA, please use the partnerReferenceNo that is the same as the one used in the transaction request process before  | [optional] 
**CustomerNumber** | **string** | Customer account number, in format 628xxx | 
**BeneficiaryAccountNumber** | **string** | Beneficiary account number | 
**Amount** | [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO  | 
**AdditionalInfo** | [**BankAccountInquiryRequestAdditionalInfo**](BankAccountInquiryRequestAdditionalInfo.md) |  | 

## Methods

### NewBankAccountInquiryRequest

`func NewBankAccountInquiryRequest(customerNumber string, beneficiaryAccountNumber string, amount Money, additionalInfo BankAccountInquiryRequestAdditionalInfo, ) *BankAccountInquiryRequest`

NewBankAccountInquiryRequest instantiates a new BankAccountInquiryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountInquiryRequestWithDefaults

`func NewBankAccountInquiryRequestWithDefaults() *BankAccountInquiryRequest`

NewBankAccountInquiryRequestWithDefaults instantiates a new BankAccountInquiryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerReferenceNo

`func (o *BankAccountInquiryRequest) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *BankAccountInquiryRequest) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *BankAccountInquiryRequest) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.

### HasPartnerReferenceNo

`func (o *BankAccountInquiryRequest) HasPartnerReferenceNo() bool`

HasPartnerReferenceNo returns a boolean if a field has been set.

### GetCustomerNumber

`func (o *BankAccountInquiryRequest) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *BankAccountInquiryRequest) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *BankAccountInquiryRequest) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.


### GetBeneficiaryAccountNumber

`func (o *BankAccountInquiryRequest) GetBeneficiaryAccountNumber() string`

GetBeneficiaryAccountNumber returns the BeneficiaryAccountNumber field if non-nil, zero value otherwise.

### GetBeneficiaryAccountNumberOk

`func (o *BankAccountInquiryRequest) GetBeneficiaryAccountNumberOk() (*string, bool)`

GetBeneficiaryAccountNumberOk returns a tuple with the BeneficiaryAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAccountNumber

`func (o *BankAccountInquiryRequest) SetBeneficiaryAccountNumber(v string)`

SetBeneficiaryAccountNumber sets BeneficiaryAccountNumber field to given value.


### GetAmount

`func (o *BankAccountInquiryRequest) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BankAccountInquiryRequest) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BankAccountInquiryRequest) SetAmount(v Money)`

SetAmount sets Amount field to given value.


### GetAdditionalInfo

`func (o *BankAccountInquiryRequest) GetAdditionalInfo() BankAccountInquiryRequestAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *BankAccountInquiryRequest) GetAdditionalInfoOk() (*BankAccountInquiryRequestAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *BankAccountInquiryRequest) SetAdditionalInfo(v BankAccountInquiryRequestAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


