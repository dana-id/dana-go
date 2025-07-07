# BankAccountInquiryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **string** | Refer to response code list | 
**ResponseMessage** | **string** | Refer to response code list | 
**ReferenceNo** | Pointer to **string** | Transaction identifier on DANA system | [optional] 
**PartnerReferenceNo** | Pointer to **string** | Unique transaction identifier on partner system which assigned to each transaction<br /> Notes:<br /> If the partner receives a timeout or an unexpected response from DANA and partner expects to perform retry request to DANA, please use the partnerReferenceNo that is the same as the one used in the transaction request process before  | [optional] 
**AccountType** | Pointer to **string** | Customer account type | [optional] 
**BeneficiaryAccountNumber** | **string** | Beneficiary account number | 
**BeneficiaryAccountName** | **string** | Beneficiary account name | 
**BeneficiaryBankCode** | Pointer to **string** | Beneficiary Bank code | [optional] 
**BeneficiaryBankShortName** | Pointer to **string** | Beneficiary Bank short name | [optional] 
**BeneficiaryBankName** | Pointer to **string** | Beneficiary Bank name | [optional] 
**Amount** | [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO  | 
**AdditionalInfo** | Pointer to [**BankAccountInquiryResponseAdditionalInfo**](BankAccountInquiryResponseAdditionalInfo.md) |  | [optional] 

## Methods

### NewBankAccountInquiryResponse

`func NewBankAccountInquiryResponse(responseCode string, responseMessage string, beneficiaryAccountNumber string, beneficiaryAccountName string, amount Money, ) *BankAccountInquiryResponse`

NewBankAccountInquiryResponse instantiates a new BankAccountInquiryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountInquiryResponseWithDefaults

`func NewBankAccountInquiryResponseWithDefaults() *BankAccountInquiryResponse`

NewBankAccountInquiryResponseWithDefaults instantiates a new BankAccountInquiryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *BankAccountInquiryResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *BankAccountInquiryResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *BankAccountInquiryResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseMessage

`func (o *BankAccountInquiryResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *BankAccountInquiryResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *BankAccountInquiryResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetReferenceNo

`func (o *BankAccountInquiryResponse) GetReferenceNo() string`

GetReferenceNo returns the ReferenceNo field if non-nil, zero value otherwise.

### GetReferenceNoOk

`func (o *BankAccountInquiryResponse) GetReferenceNoOk() (*string, bool)`

GetReferenceNoOk returns a tuple with the ReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNo

`func (o *BankAccountInquiryResponse) SetReferenceNo(v string)`

SetReferenceNo sets ReferenceNo field to given value.

### HasReferenceNo

`func (o *BankAccountInquiryResponse) HasReferenceNo() bool`

HasReferenceNo returns a boolean if a field has been set.

### GetPartnerReferenceNo

`func (o *BankAccountInquiryResponse) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *BankAccountInquiryResponse) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *BankAccountInquiryResponse) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.

### HasPartnerReferenceNo

`func (o *BankAccountInquiryResponse) HasPartnerReferenceNo() bool`

HasPartnerReferenceNo returns a boolean if a field has been set.

### GetAccountType

`func (o *BankAccountInquiryResponse) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankAccountInquiryResponse) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankAccountInquiryResponse) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *BankAccountInquiryResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBeneficiaryAccountNumber

`func (o *BankAccountInquiryResponse) GetBeneficiaryAccountNumber() string`

GetBeneficiaryAccountNumber returns the BeneficiaryAccountNumber field if non-nil, zero value otherwise.

### GetBeneficiaryAccountNumberOk

`func (o *BankAccountInquiryResponse) GetBeneficiaryAccountNumberOk() (*string, bool)`

GetBeneficiaryAccountNumberOk returns a tuple with the BeneficiaryAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAccountNumber

`func (o *BankAccountInquiryResponse) SetBeneficiaryAccountNumber(v string)`

SetBeneficiaryAccountNumber sets BeneficiaryAccountNumber field to given value.


### GetBeneficiaryAccountName

`func (o *BankAccountInquiryResponse) GetBeneficiaryAccountName() string`

GetBeneficiaryAccountName returns the BeneficiaryAccountName field if non-nil, zero value otherwise.

### GetBeneficiaryAccountNameOk

`func (o *BankAccountInquiryResponse) GetBeneficiaryAccountNameOk() (*string, bool)`

GetBeneficiaryAccountNameOk returns a tuple with the BeneficiaryAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAccountName

`func (o *BankAccountInquiryResponse) SetBeneficiaryAccountName(v string)`

SetBeneficiaryAccountName sets BeneficiaryAccountName field to given value.


### GetBeneficiaryBankCode

`func (o *BankAccountInquiryResponse) GetBeneficiaryBankCode() string`

GetBeneficiaryBankCode returns the BeneficiaryBankCode field if non-nil, zero value otherwise.

### GetBeneficiaryBankCodeOk

`func (o *BankAccountInquiryResponse) GetBeneficiaryBankCodeOk() (*string, bool)`

GetBeneficiaryBankCodeOk returns a tuple with the BeneficiaryBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryBankCode

`func (o *BankAccountInquiryResponse) SetBeneficiaryBankCode(v string)`

SetBeneficiaryBankCode sets BeneficiaryBankCode field to given value.

### HasBeneficiaryBankCode

`func (o *BankAccountInquiryResponse) HasBeneficiaryBankCode() bool`

HasBeneficiaryBankCode returns a boolean if a field has been set.

### GetBeneficiaryBankShortName

`func (o *BankAccountInquiryResponse) GetBeneficiaryBankShortName() string`

GetBeneficiaryBankShortName returns the BeneficiaryBankShortName field if non-nil, zero value otherwise.

### GetBeneficiaryBankShortNameOk

`func (o *BankAccountInquiryResponse) GetBeneficiaryBankShortNameOk() (*string, bool)`

GetBeneficiaryBankShortNameOk returns a tuple with the BeneficiaryBankShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryBankShortName

`func (o *BankAccountInquiryResponse) SetBeneficiaryBankShortName(v string)`

SetBeneficiaryBankShortName sets BeneficiaryBankShortName field to given value.

### HasBeneficiaryBankShortName

`func (o *BankAccountInquiryResponse) HasBeneficiaryBankShortName() bool`

HasBeneficiaryBankShortName returns a boolean if a field has been set.

### GetBeneficiaryBankName

`func (o *BankAccountInquiryResponse) GetBeneficiaryBankName() string`

GetBeneficiaryBankName returns the BeneficiaryBankName field if non-nil, zero value otherwise.

### GetBeneficiaryBankNameOk

`func (o *BankAccountInquiryResponse) GetBeneficiaryBankNameOk() (*string, bool)`

GetBeneficiaryBankNameOk returns a tuple with the BeneficiaryBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryBankName

`func (o *BankAccountInquiryResponse) SetBeneficiaryBankName(v string)`

SetBeneficiaryBankName sets BeneficiaryBankName field to given value.

### HasBeneficiaryBankName

`func (o *BankAccountInquiryResponse) HasBeneficiaryBankName() bool`

HasBeneficiaryBankName returns a boolean if a field has been set.

### GetAmount

`func (o *BankAccountInquiryResponse) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BankAccountInquiryResponse) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BankAccountInquiryResponse) SetAmount(v Money)`

SetAmount sets Amount field to given value.


### GetAdditionalInfo

`func (o *BankAccountInquiryResponse) GetAdditionalInfo() BankAccountInquiryResponseAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *BankAccountInquiryResponse) GetAdditionalInfoOk() (*BankAccountInquiryResponseAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *BankAccountInquiryResponse) SetAdditionalInfo(v BankAccountInquiryResponseAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *BankAccountInquiryResponse) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


