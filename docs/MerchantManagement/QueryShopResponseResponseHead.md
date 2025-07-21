# QueryShopResponseResponseHead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | API version | [optional] [default to "2.0"]
**Function** | Pointer to **string** | API interface | [optional] 
**ClientId** | Pointer to **string** | Client ID provided by DANA, used to identify partner and application system | [optional] 
**RespTime** | Pointer to **string** | Response time, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | [optional] 
**ReqMsgId** | Pointer to **string** | Request message ID | [optional] 
**ClientSecret** | Pointer to **string** | As a secret key of client. Assigned client secret during registration | [optional] 
**Reserve** | Pointer to **string** | Reserved for future implementation (Key/Value) | [optional] 

## Methods

### NewQueryShopResponseResponseHead

`func NewQueryShopResponseResponseHead() *QueryShopResponseResponseHead`

NewQueryShopResponseResponseHead instantiates a new QueryShopResponseResponseHead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryShopResponseResponseHeadWithDefaults

`func NewQueryShopResponseResponseHeadWithDefaults() *QueryShopResponseResponseHead`

NewQueryShopResponseResponseHeadWithDefaults instantiates a new QueryShopResponseResponseHead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *QueryShopResponseResponseHead) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *QueryShopResponseResponseHead) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *QueryShopResponseResponseHead) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *QueryShopResponseResponseHead) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFunction

`func (o *QueryShopResponseResponseHead) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *QueryShopResponseResponseHead) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *QueryShopResponseResponseHead) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *QueryShopResponseResponseHead) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetClientId

`func (o *QueryShopResponseResponseHead) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *QueryShopResponseResponseHead) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *QueryShopResponseResponseHead) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *QueryShopResponseResponseHead) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetRespTime

`func (o *QueryShopResponseResponseHead) GetRespTime() string`

GetRespTime returns the RespTime field if non-nil, zero value otherwise.

### GetRespTimeOk

`func (o *QueryShopResponseResponseHead) GetRespTimeOk() (*string, bool)`

GetRespTimeOk returns a tuple with the RespTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespTime

`func (o *QueryShopResponseResponseHead) SetRespTime(v string)`

SetRespTime sets RespTime field to given value.

### HasRespTime

`func (o *QueryShopResponseResponseHead) HasRespTime() bool`

HasRespTime returns a boolean if a field has been set.

### GetReqMsgId

`func (o *QueryShopResponseResponseHead) GetReqMsgId() string`

GetReqMsgId returns the ReqMsgId field if non-nil, zero value otherwise.

### GetReqMsgIdOk

`func (o *QueryShopResponseResponseHead) GetReqMsgIdOk() (*string, bool)`

GetReqMsgIdOk returns a tuple with the ReqMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqMsgId

`func (o *QueryShopResponseResponseHead) SetReqMsgId(v string)`

SetReqMsgId sets ReqMsgId field to given value.

### HasReqMsgId

`func (o *QueryShopResponseResponseHead) HasReqMsgId() bool`

HasReqMsgId returns a boolean if a field has been set.

### GetClientSecret

`func (o *QueryShopResponseResponseHead) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *QueryShopResponseResponseHead) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *QueryShopResponseResponseHead) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *QueryShopResponseResponseHead) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetReserve

`func (o *QueryShopResponseResponseHead) GetReserve() string`

GetReserve returns the Reserve field if non-nil, zero value otherwise.

### GetReserveOk

`func (o *QueryShopResponseResponseHead) GetReserveOk() (*string, bool)`

GetReserveOk returns a tuple with the Reserve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserve

`func (o *QueryShopResponseResponseHead) SetReserve(v string)`

SetReserve sets Reserve field to given value.

### HasReserve

`func (o *QueryShopResponseResponseHead) HasReserve() bool`

HasReserve returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


