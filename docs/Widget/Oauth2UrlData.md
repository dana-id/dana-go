# Oauth2UrlData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | Identifier from merchant | 
**MerchantId** | **string** | Merchant identifier that is unique per each merchant | 
**SubMerchantId** | Pointer to **string** | Information of sub merchant identifier | [optional] 
**SeamlessData** | Pointer to [**Oauth2UrlDataSeamlessData**](Oauth2UrlDataSeamlessData.md) |  | [optional] 
**Scopes** | Pointer to **[]string** | The scopes of the authorization | [optional] 
**RedirectUrl** | **string** | When user authorization is success, the user will be redirected to this URL | 
**State** | Pointer to **string** | Random string for CSRF protection purposes | [optional] 
**Lang** | Pointer to **string** | Service language code. ISO 639-1 | [optional] [default to "id"]
**AllowRegistration** | Pointer to **string** | If value equals true, provider may enable registration process during binding. Default true | [optional] [default to "true"]
**Mode** | Pointer to **string** | Mode of the authorization. The possible values are API or DEEPLINK | [optional] 

## Methods

### NewOauth2UrlData

`func NewOauth2UrlData(externalId string, merchantId string, redirectUrl string, ) *Oauth2UrlData`

NewOauth2UrlData instantiates a new Oauth2UrlData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2UrlDataWithDefaults

`func NewOauth2UrlDataWithDefaults() *Oauth2UrlData`

NewOauth2UrlDataWithDefaults instantiates a new Oauth2UrlData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *Oauth2UrlData) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Oauth2UrlData) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Oauth2UrlData) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetMerchantId

`func (o *Oauth2UrlData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *Oauth2UrlData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *Oauth2UrlData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetSubMerchantId

`func (o *Oauth2UrlData) GetSubMerchantId() string`

GetSubMerchantId returns the SubMerchantId field if non-nil, zero value otherwise.

### GetSubMerchantIdOk

`func (o *Oauth2UrlData) GetSubMerchantIdOk() (*string, bool)`

GetSubMerchantIdOk returns a tuple with the SubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantId

`func (o *Oauth2UrlData) SetSubMerchantId(v string)`

SetSubMerchantId sets SubMerchantId field to given value.

### HasSubMerchantId

`func (o *Oauth2UrlData) HasSubMerchantId() bool`

HasSubMerchantId returns a boolean if a field has been set.

### GetSeamlessData

`func (o *Oauth2UrlData) GetSeamlessData() Oauth2UrlDataSeamlessData`

GetSeamlessData returns the SeamlessData field if non-nil, zero value otherwise.

### GetSeamlessDataOk

`func (o *Oauth2UrlData) GetSeamlessDataOk() (*Oauth2UrlDataSeamlessData, bool)`

GetSeamlessDataOk returns a tuple with the SeamlessData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeamlessData

`func (o *Oauth2UrlData) SetSeamlessData(v Oauth2UrlDataSeamlessData)`

SetSeamlessData sets SeamlessData field to given value.

### HasSeamlessData

`func (o *Oauth2UrlData) HasSeamlessData() bool`

HasSeamlessData returns a boolean if a field has been set.

### GetScopes

`func (o *Oauth2UrlData) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Oauth2UrlData) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Oauth2UrlData) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *Oauth2UrlData) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *Oauth2UrlData) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *Oauth2UrlData) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *Oauth2UrlData) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### GetState

`func (o *Oauth2UrlData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Oauth2UrlData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Oauth2UrlData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Oauth2UrlData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLang

`func (o *Oauth2UrlData) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *Oauth2UrlData) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *Oauth2UrlData) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *Oauth2UrlData) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetAllowRegistration

`func (o *Oauth2UrlData) GetAllowRegistration() string`

GetAllowRegistration returns the AllowRegistration field if non-nil, zero value otherwise.

### GetAllowRegistrationOk

`func (o *Oauth2UrlData) GetAllowRegistrationOk() (*string, bool)`

GetAllowRegistrationOk returns a tuple with the AllowRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRegistration

`func (o *Oauth2UrlData) SetAllowRegistration(v string)`

SetAllowRegistration sets AllowRegistration field to given value.

### HasAllowRegistration

`func (o *Oauth2UrlData) HasAllowRegistration() bool`

HasAllowRegistration returns a boolean if a field has been set.

### GetMode

`func (o *Oauth2UrlData) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Oauth2UrlData) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Oauth2UrlData) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Oauth2UrlData) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


