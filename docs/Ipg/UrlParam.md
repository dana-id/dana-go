# UrlParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL link | 
**Type** | **string** | Url param. The enums:<br /> * PAY_RETURN - When finish payment, DANA will notify to the URL that has been defined by<br /> * NOTIFICATION - After the payment, the user will be redirected to merchant page, this is mandatory<br />  | 
**IsDeeplink** | **string** | Deeplink URL or not | 

## Methods

### NewUrlParam

`func NewUrlParam(url string, type_ string, isDeeplink string, ) *UrlParam`

NewUrlParam instantiates a new UrlParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlParamWithDefaults

`func NewUrlParamWithDefaults() *UrlParam`

NewUrlParamWithDefaults instantiates a new UrlParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UrlParam) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UrlParam) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UrlParam) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetType

`func (o *UrlParam) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UrlParam) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UrlParam) SetType(v string)`

SetType sets Type field to given value.


### GetIsDeeplink

`func (o *UrlParam) GetIsDeeplink() string`

GetIsDeeplink returns the IsDeeplink field if non-nil, zero value otherwise.

### GetIsDeeplinkOk

`func (o *UrlParam) GetIsDeeplinkOk() (*string, bool)`

GetIsDeeplinkOk returns a tuple with the IsDeeplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeeplink

`func (o *UrlParam) SetIsDeeplink(v string)`

SetIsDeeplink sets IsDeeplink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


