# EnvInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** | Session identifier | [optional] 
**TokenId** | Pointer to **string** | Token identifier | [optional] 
**WebsiteLanguage** | Pointer to **string** | Website language | [optional] 
**ClientIp** | Pointer to **string** | Client IP address | [optional] 
**OsType** | Pointer to **string** | Operating system type | [optional] 
**AppVersion** | Pointer to **string** | App version | [optional] 
**SdkVersion** | Pointer to **string** | SDK version | [optional] 
**SourcePlatform** | **string** | The source platform is always independent payment gateway (IPG) | 
**OrderOsType** | Pointer to **string** | Order operating system type | [optional] 
**MerchantAppVersion** | Pointer to **string** | Merchant App version | [optional] 
**TerminalType** | **string** | Terminal type. The enums:<br /> * APP - Mobile Application<br /> * WEB - Browser Web<br /> * WAP - Mobile Wap<br /> * SYSTEM - System Call<br />  | 
**OrderTerminalType** | Pointer to **string** | Order terminal type. The enums:<br /> * APP - Mobile Application<br /> * WEB - Browser Web<br /> * WAP - Mobile Wap<br /> * SYSTEM - System Call<br />  | [optional] 
**ExtendInfo** | Pointer to **string** | Extend information | [optional] 

## Methods

### NewEnvInfo

`func NewEnvInfo(sourcePlatform string, terminalType string, ) *EnvInfo`

NewEnvInfo instantiates a new EnvInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvInfoWithDefaults

`func NewEnvInfoWithDefaults() *EnvInfo`

NewEnvInfoWithDefaults instantiates a new EnvInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *EnvInfo) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *EnvInfo) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *EnvInfo) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *EnvInfo) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTokenId

`func (o *EnvInfo) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EnvInfo) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EnvInfo) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *EnvInfo) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetWebsiteLanguage

`func (o *EnvInfo) GetWebsiteLanguage() string`

GetWebsiteLanguage returns the WebsiteLanguage field if non-nil, zero value otherwise.

### GetWebsiteLanguageOk

`func (o *EnvInfo) GetWebsiteLanguageOk() (*string, bool)`

GetWebsiteLanguageOk returns a tuple with the WebsiteLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteLanguage

`func (o *EnvInfo) SetWebsiteLanguage(v string)`

SetWebsiteLanguage sets WebsiteLanguage field to given value.

### HasWebsiteLanguage

`func (o *EnvInfo) HasWebsiteLanguage() bool`

HasWebsiteLanguage returns a boolean if a field has been set.

### GetClientIp

`func (o *EnvInfo) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *EnvInfo) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *EnvInfo) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *EnvInfo) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetOsType

`func (o *EnvInfo) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *EnvInfo) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *EnvInfo) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *EnvInfo) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetAppVersion

`func (o *EnvInfo) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *EnvInfo) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *EnvInfo) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *EnvInfo) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetSdkVersion

`func (o *EnvInfo) GetSdkVersion() string`

GetSdkVersion returns the SdkVersion field if non-nil, zero value otherwise.

### GetSdkVersionOk

`func (o *EnvInfo) GetSdkVersionOk() (*string, bool)`

GetSdkVersionOk returns a tuple with the SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersion

`func (o *EnvInfo) SetSdkVersion(v string)`

SetSdkVersion sets SdkVersion field to given value.

### HasSdkVersion

`func (o *EnvInfo) HasSdkVersion() bool`

HasSdkVersion returns a boolean if a field has been set.

### GetSourcePlatform

`func (o *EnvInfo) GetSourcePlatform() string`

GetSourcePlatform returns the SourcePlatform field if non-nil, zero value otherwise.

### GetSourcePlatformOk

`func (o *EnvInfo) GetSourcePlatformOk() (*string, bool)`

GetSourcePlatformOk returns a tuple with the SourcePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePlatform

`func (o *EnvInfo) SetSourcePlatform(v string)`

SetSourcePlatform sets SourcePlatform field to given value.


### GetOrderOsType

`func (o *EnvInfo) GetOrderOsType() string`

GetOrderOsType returns the OrderOsType field if non-nil, zero value otherwise.

### GetOrderOsTypeOk

`func (o *EnvInfo) GetOrderOsTypeOk() (*string, bool)`

GetOrderOsTypeOk returns a tuple with the OrderOsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderOsType

`func (o *EnvInfo) SetOrderOsType(v string)`

SetOrderOsType sets OrderOsType field to given value.

### HasOrderOsType

`func (o *EnvInfo) HasOrderOsType() bool`

HasOrderOsType returns a boolean if a field has been set.

### GetMerchantAppVersion

`func (o *EnvInfo) GetMerchantAppVersion() string`

GetMerchantAppVersion returns the MerchantAppVersion field if non-nil, zero value otherwise.

### GetMerchantAppVersionOk

`func (o *EnvInfo) GetMerchantAppVersionOk() (*string, bool)`

GetMerchantAppVersionOk returns a tuple with the MerchantAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAppVersion

`func (o *EnvInfo) SetMerchantAppVersion(v string)`

SetMerchantAppVersion sets MerchantAppVersion field to given value.

### HasMerchantAppVersion

`func (o *EnvInfo) HasMerchantAppVersion() bool`

HasMerchantAppVersion returns a boolean if a field has been set.

### GetTerminalType

`func (o *EnvInfo) GetTerminalType() string`

GetTerminalType returns the TerminalType field if non-nil, zero value otherwise.

### GetTerminalTypeOk

`func (o *EnvInfo) GetTerminalTypeOk() (*string, bool)`

GetTerminalTypeOk returns a tuple with the TerminalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalType

`func (o *EnvInfo) SetTerminalType(v string)`

SetTerminalType sets TerminalType field to given value.


### GetOrderTerminalType

`func (o *EnvInfo) GetOrderTerminalType() string`

GetOrderTerminalType returns the OrderTerminalType field if non-nil, zero value otherwise.

### GetOrderTerminalTypeOk

`func (o *EnvInfo) GetOrderTerminalTypeOk() (*string, bool)`

GetOrderTerminalTypeOk returns a tuple with the OrderTerminalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTerminalType

`func (o *EnvInfo) SetOrderTerminalType(v string)`

SetOrderTerminalType sets OrderTerminalType field to given value.

### HasOrderTerminalType

`func (o *EnvInfo) HasOrderTerminalType() bool`

HasOrderTerminalType returns a boolean if a field has been set.

### GetExtendInfo

`func (o *EnvInfo) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *EnvInfo) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *EnvInfo) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *EnvInfo) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


