# TransferToBankRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerReferenceNo** | Pointer to **string** | Unique transaction identifier on partner system which assigned to each transaction<br /> Notes:<br /> If the partner receives a timeout or an unexpected response from DANA and partner expects to perform retry request to DANA, please use the partnerReferenceNo that is the same as the one used in the transaction request process before  | [optional] 
**CustomerNumber** | **string** | Customer account number, in format 628xxx | 
**AccountType** | Pointer to **string** | Customer account type | [optional] 
**BeneficiaryAccountNumber** | **string** | Beneficiary account number | 
**BeneficiaryBankCode** | **string** | Beneficiary Bank code | 
**Amount** | [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO  | 
**AdditionalInfo** | [**TransferToBankRequestAdditionalInfo**](TransferToBankRequestAdditionalInfo.md) |  | 

## Methods

### NewTransferToBankRequest

`func NewTransferToBankRequest(customerNumber string, beneficiaryAccountNumber string, beneficiaryBankCode string, amount Money, additionalInfo TransferToBankRequestAdditionalInfo, ) *TransferToBankRequest`

NewTransferToBankRequest instantiates a new TransferToBankRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferToBankRequestWithDefaults

`func NewTransferToBankRequestWithDefaults() *TransferToBankRequest`

NewTransferToBankRequestWithDefaults instantiates a new TransferToBankRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerReferenceNo

`func (o *TransferToBankRequest) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *TransferToBankRequest) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *TransferToBankRequest) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.

### HasPartnerReferenceNo

`func (o *TransferToBankRequest) HasPartnerReferenceNo() bool`

HasPartnerReferenceNo returns a boolean if a field has been set.

### GetCustomerNumber

`func (o *TransferToBankRequest) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *TransferToBankRequest) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *TransferToBankRequest) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.


### GetAccountType

`func (o *TransferToBankRequest) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *TransferToBankRequest) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *TransferToBankRequest) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *TransferToBankRequest) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBeneficiaryAccountNumber

`func (o *TransferToBankRequest) GetBeneficiaryAccountNumber() string`

GetBeneficiaryAccountNumber returns the BeneficiaryAccountNumber field if non-nil, zero value otherwise.

### GetBeneficiaryAccountNumberOk

`func (o *TransferToBankRequest) GetBeneficiaryAccountNumberOk() (*string, bool)`

GetBeneficiaryAccountNumberOk returns a tuple with the BeneficiaryAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAccountNumber

`func (o *TransferToBankRequest) SetBeneficiaryAccountNumber(v string)`

SetBeneficiaryAccountNumber sets BeneficiaryAccountNumber field to given value.


### GetBeneficiaryBankCode

`func (o *TransferToBankRequest) GetBeneficiaryBankCode() string`

GetBeneficiaryBankCode returns the BeneficiaryBankCode field if non-nil, zero value otherwise.

### GetBeneficiaryBankCodeOk

`func (o *TransferToBankRequest) GetBeneficiaryBankCodeOk() (*string, bool)`

GetBeneficiaryBankCodeOk returns a tuple with the BeneficiaryBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryBankCode

`func (o *TransferToBankRequest) SetBeneficiaryBankCode(v string)`

SetBeneficiaryBankCode sets BeneficiaryBankCode field to given value.


### GetAmount

`func (o *TransferToBankRequest) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferToBankRequest) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferToBankRequest) SetAmount(v Money)`

SetAmount sets Amount field to given value.


### GetAdditionalInfo

`func (o *TransferToBankRequest) GetAdditionalInfo() TransferToBankRequestAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *TransferToBankRequest) GetAdditionalInfoOk() (*TransferToBankRequestAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *TransferToBankRequest) SetAdditionalInfo(v TransferToBankRequestAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


