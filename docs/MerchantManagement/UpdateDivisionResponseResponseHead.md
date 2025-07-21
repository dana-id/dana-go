# UpdateDivisionResponseResponseHead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | API version | [optional] [default to "2.0"]
**Function** | Pointer to **string** | API interface | [optional] 
**ClientId** | Pointer to **string** | Client ID provided by DANA, used to identify partner and application system | [optional] 
**ClientSecret** | Pointer to **string** | As a secret key of client. Assigned client secret during registration | [optional] 
**RespTime** | Pointer to **string** | Response time, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | [optional] 
**ReqMsgId** | Pointer to **string** | Request message ID | [optional] 
**Reserve** | Pointer to **string** | Reserved for future implementation (Key/Value) | [optional] 

## Methods

### NewUpdateDivisionResponseResponseHead

`func NewUpdateDivisionResponseResponseHead() *UpdateDivisionResponseResponseHead`

NewUpdateDivisionResponseResponseHead instantiates a new UpdateDivisionResponseResponseHead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDivisionResponseResponseHeadWithDefaults

`func NewUpdateDivisionResponseResponseHeadWithDefaults() *UpdateDivisionResponseResponseHead`

NewUpdateDivisionResponseResponseHeadWithDefaults instantiates a new UpdateDivisionResponseResponseHead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *UpdateDivisionResponseResponseHead) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateDivisionResponseResponseHead) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateDivisionResponseResponseHead) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateDivisionResponseResponseHead) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFunction

`func (o *UpdateDivisionResponseResponseHead) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *UpdateDivisionResponseResponseHead) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *UpdateDivisionResponseResponseHead) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *UpdateDivisionResponseResponseHead) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetClientId

`func (o *UpdateDivisionResponseResponseHead) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UpdateDivisionResponseResponseHead) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UpdateDivisionResponseResponseHead) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UpdateDivisionResponseResponseHead) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *UpdateDivisionResponseResponseHead) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *UpdateDivisionResponseResponseHead) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *UpdateDivisionResponseResponseHead) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *UpdateDivisionResponseResponseHead) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetRespTime

`func (o *UpdateDivisionResponseResponseHead) GetRespTime() string`

GetRespTime returns the RespTime field if non-nil, zero value otherwise.

### GetRespTimeOk

`func (o *UpdateDivisionResponseResponseHead) GetRespTimeOk() (*string, bool)`

GetRespTimeOk returns a tuple with the RespTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespTime

`func (o *UpdateDivisionResponseResponseHead) SetRespTime(v string)`

SetRespTime sets RespTime field to given value.

### HasRespTime

`func (o *UpdateDivisionResponseResponseHead) HasRespTime() bool`

HasRespTime returns a boolean if a field has been set.

### GetReqMsgId

`func (o *UpdateDivisionResponseResponseHead) GetReqMsgId() string`

GetReqMsgId returns the ReqMsgId field if non-nil, zero value otherwise.

### GetReqMsgIdOk

`func (o *UpdateDivisionResponseResponseHead) GetReqMsgIdOk() (*string, bool)`

GetReqMsgIdOk returns a tuple with the ReqMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqMsgId

`func (o *UpdateDivisionResponseResponseHead) SetReqMsgId(v string)`

SetReqMsgId sets ReqMsgId field to given value.

### HasReqMsgId

`func (o *UpdateDivisionResponseResponseHead) HasReqMsgId() bool`

HasReqMsgId returns a boolean if a field has been set.

### GetReserve

`func (o *UpdateDivisionResponseResponseHead) GetReserve() string`

GetReserve returns the Reserve field if non-nil, zero value otherwise.

### GetReserveOk

`func (o *UpdateDivisionResponseResponseHead) GetReserveOk() (*string, bool)`

GetReserveOk returns a tuple with the Reserve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserve

`func (o *UpdateDivisionResponseResponseHead) SetReserve(v string)`

SetReserve sets Reserve field to given value.

### HasReserve

`func (o *UpdateDivisionResponseResponseHead) HasReserve() bool`

HasReserve returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


