# DanaAccountInquiryRequestAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FundType** | **string** | Additional information of top up fund type, i.e.<br /> AGENT_TOPUP_FOR_  | 
**ExternalDivisionId** | Pointer to **string** | Additional information of external division identifier. This parameter only used for Top Up Disbursement subMerchant (fundType : AGENT_TOPUP_FOR_USER_SETTLE)<br /> Notes:<br /> The required of this parameter is Optional, but if \&quot;additionalInfo.chargeTarget\&quot; has value DIVISION then the required of this parameter will be changed to Mandatory  | [optional] 
**ChargeTarget** | Pointer to **string** | Additional information of charge target. This parameter only used for Top Up Disbursement subMerchant. The value are:<br /> • null<br /> • DIVISION<br /> • MERCHANT<br /> if the value is DIVISION, externalDivisionId will be Mandatory  | [optional] 
**AccessToken** | Pointer to **string** | 1. Contains customer token, which has been obtained from binding process, refer to Account Binding &amp; Unbinding documentation<br /> 2. If request is coming from user interaction, this field is mandatory. If not, just filled customerNumber  | [optional] 
**CustomerId** | Pointer to **string** | Public user identifier of DANA user<br /> Notes:<br /> If used, requires customerNumber to be filled with default phone number format \&quot;620000000000\&quot;  | [optional] 

## Methods

### NewDanaAccountInquiryRequestAdditionalInfo

`func NewDanaAccountInquiryRequestAdditionalInfo(fundType string, ) *DanaAccountInquiryRequestAdditionalInfo`

NewDanaAccountInquiryRequestAdditionalInfo instantiates a new DanaAccountInquiryRequestAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDanaAccountInquiryRequestAdditionalInfoWithDefaults

`func NewDanaAccountInquiryRequestAdditionalInfoWithDefaults() *DanaAccountInquiryRequestAdditionalInfo`

NewDanaAccountInquiryRequestAdditionalInfoWithDefaults instantiates a new DanaAccountInquiryRequestAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFundType

`func (o *DanaAccountInquiryRequestAdditionalInfo) GetFundType() string`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *DanaAccountInquiryRequestAdditionalInfo) GetFundTypeOk() (*string, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *DanaAccountInquiryRequestAdditionalInfo) SetFundType(v string)`

SetFundType sets FundType field to given value.


### GetExternalDivisionId

`func (o *DanaAccountInquiryRequestAdditionalInfo) GetExternalDivisionId() string`

GetExternalDivisionId returns the ExternalDivisionId field if non-nil, zero value otherwise.

### GetExternalDivisionIdOk

`func (o *DanaAccountInquiryRequestAdditionalInfo) GetExternalDivisionIdOk() (*string, bool)`

GetExternalDivisionIdOk returns a tuple with the ExternalDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDivisionId

`func (o *DanaAccountInquiryRequestAdditionalInfo) SetExternalDivisionId(v string)`

SetExternalDivisionId sets ExternalDivisionId field to given value.

### HasExternalDivisionId

`func (o *DanaAccountInquiryRequestAdditionalInfo) HasExternalDivisionId() bool`

HasExternalDivisionId returns a boolean if a field has been set.

### GetChargeTarget

`func (o *DanaAccountInquiryRequestAdditionalInfo) GetChargeTarget() string`

GetChargeTarget returns the ChargeTarget field if non-nil, zero value otherwise.

### GetChargeTargetOk

`func (o *DanaAccountInquiryRequestAdditionalInfo) GetChargeTargetOk() (*string, bool)`

GetChargeTargetOk returns a tuple with the ChargeTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeTarget

`func (o *DanaAccountInquiryRequestAdditionalInfo) SetChargeTarget(v string)`

SetChargeTarget sets ChargeTarget field to given value.

### HasChargeTarget

`func (o *DanaAccountInquiryRequestAdditionalInfo) HasChargeTarget() bool`

HasChargeTarget returns a boolean if a field has been set.

### GetAccessToken

`func (o *DanaAccountInquiryRequestAdditionalInfo) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *DanaAccountInquiryRequestAdditionalInfo) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *DanaAccountInquiryRequestAdditionalInfo) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *DanaAccountInquiryRequestAdditionalInfo) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetCustomerId

`func (o *DanaAccountInquiryRequestAdditionalInfo) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DanaAccountInquiryRequestAdditionalInfo) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DanaAccountInquiryRequestAdditionalInfo) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DanaAccountInquiryRequestAdditionalInfo) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


