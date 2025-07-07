# TransferToBankRequestAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FundType** | **string** | Additional information of withdraw fund type, i.e.<br /> MERCHANT_WITHDRAW_FOR_CORPORATE  | 
**ExternalDivisionId** | Pointer to **string** | Additional information of external division identifier. (fundType: MERCHANT_WITHDRAW_FOR_CORPORATE)<br /> Notes: The required of this parameter is Optional, but if \&quot;additionalInfo.chargeTarget\&quot; has value DIVISION then the required of this parameter will be changed to Mandatory  | [optional] 
**ChargeTarget** | Pointer to **string** | Additional information of charge target. The values are:<br /> • null<br /> • DIVISION<br /> • MERCHANT<br /> Notes: If the value is DIVISION, externalDivisionId will be Mandatory  | [optional] 
**NeedNotify** | Pointer to **bool** | Additional information of flag result notification on transaction completed (result sync/async) | [optional] 
**BeneficiaryAccountName** | Pointer to **string** | Additional information of beneficiary account name for validation purpose | [optional] 
**AccessToken** | Pointer to **string** | Contains customer token, which has been obtained from binding process, refer to Account Binding &amp; Unbinding documentation<br /> If request is coming from user interaction, this field is mandatory. If not, just filled customerNumber  | [optional] 

## Methods

### NewTransferToBankRequestAdditionalInfo

`func NewTransferToBankRequestAdditionalInfo(fundType string, ) *TransferToBankRequestAdditionalInfo`

NewTransferToBankRequestAdditionalInfo instantiates a new TransferToBankRequestAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferToBankRequestAdditionalInfoWithDefaults

`func NewTransferToBankRequestAdditionalInfoWithDefaults() *TransferToBankRequestAdditionalInfo`

NewTransferToBankRequestAdditionalInfoWithDefaults instantiates a new TransferToBankRequestAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFundType

`func (o *TransferToBankRequestAdditionalInfo) GetFundType() string`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *TransferToBankRequestAdditionalInfo) GetFundTypeOk() (*string, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *TransferToBankRequestAdditionalInfo) SetFundType(v string)`

SetFundType sets FundType field to given value.


### GetExternalDivisionId

`func (o *TransferToBankRequestAdditionalInfo) GetExternalDivisionId() string`

GetExternalDivisionId returns the ExternalDivisionId field if non-nil, zero value otherwise.

### GetExternalDivisionIdOk

`func (o *TransferToBankRequestAdditionalInfo) GetExternalDivisionIdOk() (*string, bool)`

GetExternalDivisionIdOk returns a tuple with the ExternalDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDivisionId

`func (o *TransferToBankRequestAdditionalInfo) SetExternalDivisionId(v string)`

SetExternalDivisionId sets ExternalDivisionId field to given value.

### HasExternalDivisionId

`func (o *TransferToBankRequestAdditionalInfo) HasExternalDivisionId() bool`

HasExternalDivisionId returns a boolean if a field has been set.

### GetChargeTarget

`func (o *TransferToBankRequestAdditionalInfo) GetChargeTarget() string`

GetChargeTarget returns the ChargeTarget field if non-nil, zero value otherwise.

### GetChargeTargetOk

`func (o *TransferToBankRequestAdditionalInfo) GetChargeTargetOk() (*string, bool)`

GetChargeTargetOk returns a tuple with the ChargeTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeTarget

`func (o *TransferToBankRequestAdditionalInfo) SetChargeTarget(v string)`

SetChargeTarget sets ChargeTarget field to given value.

### HasChargeTarget

`func (o *TransferToBankRequestAdditionalInfo) HasChargeTarget() bool`

HasChargeTarget returns a boolean if a field has been set.

### GetNeedNotify

`func (o *TransferToBankRequestAdditionalInfo) GetNeedNotify() bool`

GetNeedNotify returns the NeedNotify field if non-nil, zero value otherwise.

### GetNeedNotifyOk

`func (o *TransferToBankRequestAdditionalInfo) GetNeedNotifyOk() (*bool, bool)`

GetNeedNotifyOk returns a tuple with the NeedNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedNotify

`func (o *TransferToBankRequestAdditionalInfo) SetNeedNotify(v bool)`

SetNeedNotify sets NeedNotify field to given value.

### HasNeedNotify

`func (o *TransferToBankRequestAdditionalInfo) HasNeedNotify() bool`

HasNeedNotify returns a boolean if a field has been set.

### GetBeneficiaryAccountName

`func (o *TransferToBankRequestAdditionalInfo) GetBeneficiaryAccountName() string`

GetBeneficiaryAccountName returns the BeneficiaryAccountName field if non-nil, zero value otherwise.

### GetBeneficiaryAccountNameOk

`func (o *TransferToBankRequestAdditionalInfo) GetBeneficiaryAccountNameOk() (*string, bool)`

GetBeneficiaryAccountNameOk returns a tuple with the BeneficiaryAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAccountName

`func (o *TransferToBankRequestAdditionalInfo) SetBeneficiaryAccountName(v string)`

SetBeneficiaryAccountName sets BeneficiaryAccountName field to given value.

### HasBeneficiaryAccountName

`func (o *TransferToBankRequestAdditionalInfo) HasBeneficiaryAccountName() bool`

HasBeneficiaryAccountName returns a boolean if a field has been set.

### GetAccessToken

`func (o *TransferToBankRequestAdditionalInfo) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TransferToBankRequestAdditionalInfo) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TransferToBankRequestAdditionalInfo) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TransferToBankRequestAdditionalInfo) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


