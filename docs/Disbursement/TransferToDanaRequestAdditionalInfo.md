# TransferToDanaRequestAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtendInfo** | Pointer to **string** | Additional information of extend | [optional] 
**AccountType** | Pointer to **string** | Additional information of account type | [optional] 
**FundType** | **string** | Additional information of transfer to DANA fund type, i.e.<br /> AGENT_TOPUP_FOR_USER_SETTLE  | 
**ExternalDivisionId** | Pointer to **string** | Additional information of external division identifier. This parameter only used for Transfer to DANA subMerchant (fundType : AGENT_TOPUP_FOR_USER_SETTLE)<br /> Notes:<br /> The required of this parameter is Optional, but if \&quot;additionalInfo.chargeTarget\&quot; has value DIVISION then the required of this parameter will be changed to Mandatory  | [optional] 
**ChargeTarget** | Pointer to **string** | Additional information of charge target. This parameter only used for Transfer to DANA subMerchant. The value are:<br /> • null<br /> • DIVISION<br /> • MERCHANT<br /> if the value is DIVISION, externalDivisionId will be Mandatory  | [optional] 
**AccessToken** | Pointer to **string** | Contains customer token, which has been obtained from binding process, refer to Account Binding &amp; Unbinding documentation<br /> If request is coming from user interaction, this field is mandatory. If not, just filled customerNumber  | [optional] 
**CustomerId** | Pointer to **string** | Public user identifier of DANA user.<br /> Notes: If used, requires customerNumber to be filled with default phone number literal \&quot;620000000000\&quot;  | [optional] 

## Methods

### NewTransferToDanaRequestAdditionalInfo

`func NewTransferToDanaRequestAdditionalInfo(fundType string, ) *TransferToDanaRequestAdditionalInfo`

NewTransferToDanaRequestAdditionalInfo instantiates a new TransferToDanaRequestAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferToDanaRequestAdditionalInfoWithDefaults

`func NewTransferToDanaRequestAdditionalInfoWithDefaults() *TransferToDanaRequestAdditionalInfo`

NewTransferToDanaRequestAdditionalInfoWithDefaults instantiates a new TransferToDanaRequestAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtendInfo

`func (o *TransferToDanaRequestAdditionalInfo) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *TransferToDanaRequestAdditionalInfo) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *TransferToDanaRequestAdditionalInfo) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *TransferToDanaRequestAdditionalInfo) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.

### GetAccountType

`func (o *TransferToDanaRequestAdditionalInfo) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *TransferToDanaRequestAdditionalInfo) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *TransferToDanaRequestAdditionalInfo) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *TransferToDanaRequestAdditionalInfo) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetFundType

`func (o *TransferToDanaRequestAdditionalInfo) GetFundType() string`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *TransferToDanaRequestAdditionalInfo) GetFundTypeOk() (*string, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *TransferToDanaRequestAdditionalInfo) SetFundType(v string)`

SetFundType sets FundType field to given value.


### GetExternalDivisionId

`func (o *TransferToDanaRequestAdditionalInfo) GetExternalDivisionId() string`

GetExternalDivisionId returns the ExternalDivisionId field if non-nil, zero value otherwise.

### GetExternalDivisionIdOk

`func (o *TransferToDanaRequestAdditionalInfo) GetExternalDivisionIdOk() (*string, bool)`

GetExternalDivisionIdOk returns a tuple with the ExternalDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDivisionId

`func (o *TransferToDanaRequestAdditionalInfo) SetExternalDivisionId(v string)`

SetExternalDivisionId sets ExternalDivisionId field to given value.

### HasExternalDivisionId

`func (o *TransferToDanaRequestAdditionalInfo) HasExternalDivisionId() bool`

HasExternalDivisionId returns a boolean if a field has been set.

### GetChargeTarget

`func (o *TransferToDanaRequestAdditionalInfo) GetChargeTarget() string`

GetChargeTarget returns the ChargeTarget field if non-nil, zero value otherwise.

### GetChargeTargetOk

`func (o *TransferToDanaRequestAdditionalInfo) GetChargeTargetOk() (*string, bool)`

GetChargeTargetOk returns a tuple with the ChargeTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeTarget

`func (o *TransferToDanaRequestAdditionalInfo) SetChargeTarget(v string)`

SetChargeTarget sets ChargeTarget field to given value.

### HasChargeTarget

`func (o *TransferToDanaRequestAdditionalInfo) HasChargeTarget() bool`

HasChargeTarget returns a boolean if a field has been set.

### GetAccessToken

`func (o *TransferToDanaRequestAdditionalInfo) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TransferToDanaRequestAdditionalInfo) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TransferToDanaRequestAdditionalInfo) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TransferToDanaRequestAdditionalInfo) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetCustomerId

`func (o *TransferToDanaRequestAdditionalInfo) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TransferToDanaRequestAdditionalInfo) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TransferToDanaRequestAdditionalInfo) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *TransferToDanaRequestAdditionalInfo) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


