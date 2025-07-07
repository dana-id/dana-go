# BankAccountInquiryRequestAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FundType** | **string** | Additional information of withdraw fund type, i.e.<br /> MERCHANT_WITHDRAW_FOR_CORPORATE  | 
**ExternalDivisionId** | Pointer to **string** | Additional information of external division identifier. (fundType: MERCHANT_WITHDRAW_FOR_CORPORATE)<br /> Notes: The required of this parameter is Optional, but if \&quot;additionalInfo.chargeTarget\&quot; has value DIVISION then the required of this parameter will be changed to Mandatory  | [optional] 
**ChargeTarget** | Pointer to **string** | Additional information of charge target. The values are:<br /> • null<br /> • DIVISION<br /> • MERCHANT<br /> Notes: If the value is DIVISION, externalDivisionId will be Mandatory  | [optional] 
**BeneficiaryBankCode** | **string** | Additional information of beneficiary Bank code | 
**BeneficiaryAccountName** | Pointer to **string** | Additional information of beneficiary account name for validation purpose | [optional] 
**AccountType** | Pointer to **string** | Additional information of account type | [optional] 
**AccessToken** | Pointer to **string** | Contains customer token, which has been obtained from binding process, refer to Account Binding &amp; Unbinding documentation<br /> If request is coming from user interaction, this field is mandatory. If not, just filled customerNumber  | [optional] 

## Methods

### NewBankAccountInquiryRequestAdditionalInfo

`func NewBankAccountInquiryRequestAdditionalInfo(fundType string, beneficiaryBankCode string, ) *BankAccountInquiryRequestAdditionalInfo`

NewBankAccountInquiryRequestAdditionalInfo instantiates a new BankAccountInquiryRequestAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountInquiryRequestAdditionalInfoWithDefaults

`func NewBankAccountInquiryRequestAdditionalInfoWithDefaults() *BankAccountInquiryRequestAdditionalInfo`

NewBankAccountInquiryRequestAdditionalInfoWithDefaults instantiates a new BankAccountInquiryRequestAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFundType

`func (o *BankAccountInquiryRequestAdditionalInfo) GetFundType() string`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *BankAccountInquiryRequestAdditionalInfo) GetFundTypeOk() (*string, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *BankAccountInquiryRequestAdditionalInfo) SetFundType(v string)`

SetFundType sets FundType field to given value.


### GetExternalDivisionId

`func (o *BankAccountInquiryRequestAdditionalInfo) GetExternalDivisionId() string`

GetExternalDivisionId returns the ExternalDivisionId field if non-nil, zero value otherwise.

### GetExternalDivisionIdOk

`func (o *BankAccountInquiryRequestAdditionalInfo) GetExternalDivisionIdOk() (*string, bool)`

GetExternalDivisionIdOk returns a tuple with the ExternalDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDivisionId

`func (o *BankAccountInquiryRequestAdditionalInfo) SetExternalDivisionId(v string)`

SetExternalDivisionId sets ExternalDivisionId field to given value.

### HasExternalDivisionId

`func (o *BankAccountInquiryRequestAdditionalInfo) HasExternalDivisionId() bool`

HasExternalDivisionId returns a boolean if a field has been set.

### GetChargeTarget

`func (o *BankAccountInquiryRequestAdditionalInfo) GetChargeTarget() string`

GetChargeTarget returns the ChargeTarget field if non-nil, zero value otherwise.

### GetChargeTargetOk

`func (o *BankAccountInquiryRequestAdditionalInfo) GetChargeTargetOk() (*string, bool)`

GetChargeTargetOk returns a tuple with the ChargeTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeTarget

`func (o *BankAccountInquiryRequestAdditionalInfo) SetChargeTarget(v string)`

SetChargeTarget sets ChargeTarget field to given value.

### HasChargeTarget

`func (o *BankAccountInquiryRequestAdditionalInfo) HasChargeTarget() bool`

HasChargeTarget returns a boolean if a field has been set.

### GetBeneficiaryBankCode

`func (o *BankAccountInquiryRequestAdditionalInfo) GetBeneficiaryBankCode() string`

GetBeneficiaryBankCode returns the BeneficiaryBankCode field if non-nil, zero value otherwise.

### GetBeneficiaryBankCodeOk

`func (o *BankAccountInquiryRequestAdditionalInfo) GetBeneficiaryBankCodeOk() (*string, bool)`

GetBeneficiaryBankCodeOk returns a tuple with the BeneficiaryBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryBankCode

`func (o *BankAccountInquiryRequestAdditionalInfo) SetBeneficiaryBankCode(v string)`

SetBeneficiaryBankCode sets BeneficiaryBankCode field to given value.


### GetBeneficiaryAccountName

`func (o *BankAccountInquiryRequestAdditionalInfo) GetBeneficiaryAccountName() string`

GetBeneficiaryAccountName returns the BeneficiaryAccountName field if non-nil, zero value otherwise.

### GetBeneficiaryAccountNameOk

`func (o *BankAccountInquiryRequestAdditionalInfo) GetBeneficiaryAccountNameOk() (*string, bool)`

GetBeneficiaryAccountNameOk returns a tuple with the BeneficiaryAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAccountName

`func (o *BankAccountInquiryRequestAdditionalInfo) SetBeneficiaryAccountName(v string)`

SetBeneficiaryAccountName sets BeneficiaryAccountName field to given value.

### HasBeneficiaryAccountName

`func (o *BankAccountInquiryRequestAdditionalInfo) HasBeneficiaryAccountName() bool`

HasBeneficiaryAccountName returns a boolean if a field has been set.

### GetAccountType

`func (o *BankAccountInquiryRequestAdditionalInfo) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankAccountInquiryRequestAdditionalInfo) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankAccountInquiryRequestAdditionalInfo) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *BankAccountInquiryRequestAdditionalInfo) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAccessToken

`func (o *BankAccountInquiryRequestAdditionalInfo) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *BankAccountInquiryRequestAdditionalInfo) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *BankAccountInquiryRequestAdditionalInfo) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *BankAccountInquiryRequestAdditionalInfo) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


