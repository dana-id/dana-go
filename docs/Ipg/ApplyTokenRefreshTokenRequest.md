# ApplyTokenRefreshTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInfo** | Pointer to **map[string]interface{}** | Additional information | [optional] 
**GrantType** | **string** | Apply token request type. The value is REFRESH_TOKEN | 
**AuthCode** | Pointer to **string** | Authorization code. Please refer to https://dashboard.dana.id/api-docs/read/125. Required if grantType is AUTHORIZATION_CODE | [optional] 
**RefreshToken** | **string** | This token is used for refresh session if existing token has been expired | 

## Methods

### NewApplyTokenRefreshTokenRequest

`func NewApplyTokenRefreshTokenRequest(grantType string, refreshToken string, ) *ApplyTokenRefreshTokenRequest`

NewApplyTokenRefreshTokenRequest instantiates a new ApplyTokenRefreshTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyTokenRefreshTokenRequestWithDefaults

`func NewApplyTokenRefreshTokenRequestWithDefaults() *ApplyTokenRefreshTokenRequest`

NewApplyTokenRefreshTokenRequestWithDefaults instantiates a new ApplyTokenRefreshTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInfo

`func (o *ApplyTokenRefreshTokenRequest) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ApplyTokenRefreshTokenRequest) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ApplyTokenRefreshTokenRequest) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ApplyTokenRefreshTokenRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetGrantType

`func (o *ApplyTokenRefreshTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *ApplyTokenRefreshTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *ApplyTokenRefreshTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetAuthCode

`func (o *ApplyTokenRefreshTokenRequest) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *ApplyTokenRefreshTokenRequest) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *ApplyTokenRefreshTokenRequest) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *ApplyTokenRefreshTokenRequest) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetRefreshToken

`func (o *ApplyTokenRefreshTokenRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *ApplyTokenRefreshTokenRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *ApplyTokenRefreshTokenRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


