# Oauth2UrlDataSeamlessData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BizScenario** | Pointer to **string** | User’s bizScenario | [optional] 
**MobileNumber** | Pointer to **string** | User&#39;s phone number. If this field is filled in, the user must log in with the number that has been included | [optional] 
**VerifiedTime** | Pointer to **string** | Value which states that the mobile number that has been included in seamlessData has verified ownership and does not require OTP verification by the provider, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | [optional] 
**ExternalUid** | Pointer to **string** | User identifier on partner application | [optional] 
**DeviceId** | Pointer to **string** | User&#39;s device identifier | [optional] 
**SkipRegisterConsult** | Pointer to **bool** | Identifier to differentiate seamless registration flow. The possible values are true or false | [optional] 

## Methods

### NewOauth2UrlDataSeamlessData

`func NewOauth2UrlDataSeamlessData() *Oauth2UrlDataSeamlessData`

NewOauth2UrlDataSeamlessData instantiates a new Oauth2UrlDataSeamlessData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2UrlDataSeamlessDataWithDefaults

`func NewOauth2UrlDataSeamlessDataWithDefaults() *Oauth2UrlDataSeamlessData`

NewOauth2UrlDataSeamlessDataWithDefaults instantiates a new Oauth2UrlDataSeamlessData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBizScenario

`func (o *Oauth2UrlDataSeamlessData) GetBizScenario() string`

GetBizScenario returns the BizScenario field if non-nil, zero value otherwise.

### GetBizScenarioOk

`func (o *Oauth2UrlDataSeamlessData) GetBizScenarioOk() (*string, bool)`

GetBizScenarioOk returns a tuple with the BizScenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizScenario

`func (o *Oauth2UrlDataSeamlessData) SetBizScenario(v string)`

SetBizScenario sets BizScenario field to given value.

### HasBizScenario

`func (o *Oauth2UrlDataSeamlessData) HasBizScenario() bool`

HasBizScenario returns a boolean if a field has been set.

### GetMobileNumber

`func (o *Oauth2UrlDataSeamlessData) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *Oauth2UrlDataSeamlessData) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *Oauth2UrlDataSeamlessData) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *Oauth2UrlDataSeamlessData) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### GetVerifiedTime

`func (o *Oauth2UrlDataSeamlessData) GetVerifiedTime() string`

GetVerifiedTime returns the VerifiedTime field if non-nil, zero value otherwise.

### GetVerifiedTimeOk

`func (o *Oauth2UrlDataSeamlessData) GetVerifiedTimeOk() (*string, bool)`

GetVerifiedTimeOk returns a tuple with the VerifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedTime

`func (o *Oauth2UrlDataSeamlessData) SetVerifiedTime(v string)`

SetVerifiedTime sets VerifiedTime field to given value.

### HasVerifiedTime

`func (o *Oauth2UrlDataSeamlessData) HasVerifiedTime() bool`

HasVerifiedTime returns a boolean if a field has been set.

### GetExternalUid

`func (o *Oauth2UrlDataSeamlessData) GetExternalUid() string`

GetExternalUid returns the ExternalUid field if non-nil, zero value otherwise.

### GetExternalUidOk

`func (o *Oauth2UrlDataSeamlessData) GetExternalUidOk() (*string, bool)`

GetExternalUidOk returns a tuple with the ExternalUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUid

`func (o *Oauth2UrlDataSeamlessData) SetExternalUid(v string)`

SetExternalUid sets ExternalUid field to given value.

### HasExternalUid

`func (o *Oauth2UrlDataSeamlessData) HasExternalUid() bool`

HasExternalUid returns a boolean if a field has been set.

### GetDeviceId

`func (o *Oauth2UrlDataSeamlessData) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *Oauth2UrlDataSeamlessData) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *Oauth2UrlDataSeamlessData) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *Oauth2UrlDataSeamlessData) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetSkipRegisterConsult

`func (o *Oauth2UrlDataSeamlessData) GetSkipRegisterConsult() bool`

GetSkipRegisterConsult returns the SkipRegisterConsult field if non-nil, zero value otherwise.

### GetSkipRegisterConsultOk

`func (o *Oauth2UrlDataSeamlessData) GetSkipRegisterConsultOk() (*bool, bool)`

GetSkipRegisterConsultOk returns a tuple with the SkipRegisterConsult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRegisterConsult

`func (o *Oauth2UrlDataSeamlessData) SetSkipRegisterConsult(v bool)`

SetSkipRegisterConsult sets SkipRegisterConsult field to given value.

### HasSkipRegisterConsult

`func (o *Oauth2UrlDataSeamlessData) HasSkipRegisterConsult() bool`

HasSkipRegisterConsult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


