# ApplyTokenAuthorizationCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInfo** | Pointer to **map[string]interface{}** | Additional information | [optional] 
**GrantType** | **string** | Apply token request type. The value is AUTHORIZATION_CODE | 
**AuthCode** | **string** | Authorization code. Please refer to https://dashboard.dana.id/api-docs/read/125. Required if grantType is AUTHORIZATION_CODE | 
**RefreshToken** | Pointer to **string** | This token is used for refresh session if existing token has been expired | [optional] 

## Methods

### NewApplyTokenAuthorizationCodeRequest

`func NewApplyTokenAuthorizationCodeRequest(grantType string, authCode string, ) *ApplyTokenAuthorizationCodeRequest`

NewApplyTokenAuthorizationCodeRequest instantiates a new ApplyTokenAuthorizationCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyTokenAuthorizationCodeRequestWithDefaults

`func NewApplyTokenAuthorizationCodeRequestWithDefaults() *ApplyTokenAuthorizationCodeRequest`

NewApplyTokenAuthorizationCodeRequestWithDefaults instantiates a new ApplyTokenAuthorizationCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInfo

`func (o *ApplyTokenAuthorizationCodeRequest) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ApplyTokenAuthorizationCodeRequest) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ApplyTokenAuthorizationCodeRequest) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ApplyTokenAuthorizationCodeRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetGrantType

`func (o *ApplyTokenAuthorizationCodeRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *ApplyTokenAuthorizationCodeRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *ApplyTokenAuthorizationCodeRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetAuthCode

`func (o *ApplyTokenAuthorizationCodeRequest) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *ApplyTokenAuthorizationCodeRequest) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *ApplyTokenAuthorizationCodeRequest) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### GetRefreshToken

`func (o *ApplyTokenAuthorizationCodeRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *ApplyTokenAuthorizationCodeRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *ApplyTokenAuthorizationCodeRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *ApplyTokenAuthorizationCodeRequest) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


