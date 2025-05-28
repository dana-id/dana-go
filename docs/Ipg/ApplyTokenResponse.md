# ApplyTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **string** | Response code. Refer to https://dashboard.dana.id/api-docs/read/110#HTML-ApplyToken-ResponseCodeandMessage | 
**ResponseMessage** | **string** | Response message. Refer to https://dashboard.dana.id/api-docs/read/110#HTML-ApplyToken-ResponseCodeandMessage | 
**TokenType** | Pointer to **string** | Token type. Present if successfully processed | [optional] 
**AccessToken** | **string** | This token is called Customer Token that will be used as a parameter on header in other API “Authorization-Customer”. Present if successfully processed | 
**AccessTokenExpiryTime** | Pointer to **string** | Expiry time for access token was given to user, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time). Present if successfully processed | [optional] 
**RefreshToken** | Pointer to **string** | This token is used for refresh session if existing token has been expired. Present if successfully processed | [optional] 
**RefreshTokenExpiryTime** | Pointer to **string** | Expiry time for refresh token was given to user, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time). Present if successfully processed | [optional] 
**AdditionalInfo** | Pointer to [**ApplyTokenResponseAdditionalInfo**](ApplyTokenResponseAdditionalInfo.md) | Additional information | [optional] 

## Methods

### NewApplyTokenResponse

`func NewApplyTokenResponse(responseCode string, responseMessage string, accessToken string, ) *ApplyTokenResponse`

NewApplyTokenResponse instantiates a new ApplyTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyTokenResponseWithDefaults

`func NewApplyTokenResponseWithDefaults() *ApplyTokenResponse`

NewApplyTokenResponseWithDefaults instantiates a new ApplyTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *ApplyTokenResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *ApplyTokenResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *ApplyTokenResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseMessage

`func (o *ApplyTokenResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *ApplyTokenResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *ApplyTokenResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetTokenType

`func (o *ApplyTokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ApplyTokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ApplyTokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *ApplyTokenResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetAccessToken

`func (o *ApplyTokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ApplyTokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ApplyTokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAccessTokenExpiryTime

`func (o *ApplyTokenResponse) GetAccessTokenExpiryTime() string`

GetAccessTokenExpiryTime returns the AccessTokenExpiryTime field if non-nil, zero value otherwise.

### GetAccessTokenExpiryTimeOk

`func (o *ApplyTokenResponse) GetAccessTokenExpiryTimeOk() (*string, bool)`

GetAccessTokenExpiryTimeOk returns a tuple with the AccessTokenExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiryTime

`func (o *ApplyTokenResponse) SetAccessTokenExpiryTime(v string)`

SetAccessTokenExpiryTime sets AccessTokenExpiryTime field to given value.

### HasAccessTokenExpiryTime

`func (o *ApplyTokenResponse) HasAccessTokenExpiryTime() bool`

HasAccessTokenExpiryTime returns a boolean if a field has been set.

### GetRefreshToken

`func (o *ApplyTokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *ApplyTokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *ApplyTokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *ApplyTokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetRefreshTokenExpiryTime

`func (o *ApplyTokenResponse) GetRefreshTokenExpiryTime() string`

GetRefreshTokenExpiryTime returns the RefreshTokenExpiryTime field if non-nil, zero value otherwise.

### GetRefreshTokenExpiryTimeOk

`func (o *ApplyTokenResponse) GetRefreshTokenExpiryTimeOk() (*string, bool)`

GetRefreshTokenExpiryTimeOk returns a tuple with the RefreshTokenExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpiryTime

`func (o *ApplyTokenResponse) SetRefreshTokenExpiryTime(v string)`

SetRefreshTokenExpiryTime sets RefreshTokenExpiryTime field to given value.

### HasRefreshTokenExpiryTime

`func (o *ApplyTokenResponse) HasRefreshTokenExpiryTime() bool`

HasRefreshTokenExpiryTime returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *ApplyTokenResponse) GetAdditionalInfo() ApplyTokenResponseAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ApplyTokenResponse) GetAdditionalInfoOk() (*ApplyTokenResponseAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ApplyTokenResponse) SetAdditionalInfo(v ApplyTokenResponseAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ApplyTokenResponse) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


