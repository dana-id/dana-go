# QueryAssetCardListResponseResponseHead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | API version | [optional] [default to "2.0"]
**Function** | Pointer to **string** | API interface | [optional] 
**ClientId** | Pointer to **string** | Client ID provided by DANA, used to identify partner and application system | [optional] 
**RespTime** | Pointer to **string** | DateTime with timezone (ISO-8601) | [optional] 
**ReqMsgId** | Pointer to **string** | Request message ID | [optional] 
**ClientSecret** | Pointer to **string** | As a secret key of client. Assigned client secret during registration | [optional] 
**Reserve** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewQueryAssetCardListResponseResponseHead

`func NewQueryAssetCardListResponseResponseHead() *QueryAssetCardListResponseResponseHead`

NewQueryAssetCardListResponseResponseHead instantiates a new QueryAssetCardListResponseResponseHead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryAssetCardListResponseResponseHeadWithDefaults

`func NewQueryAssetCardListResponseResponseHeadWithDefaults() *QueryAssetCardListResponseResponseHead`

NewQueryAssetCardListResponseResponseHeadWithDefaults instantiates a new QueryAssetCardListResponseResponseHead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *QueryAssetCardListResponseResponseHead) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *QueryAssetCardListResponseResponseHead) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *QueryAssetCardListResponseResponseHead) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *QueryAssetCardListResponseResponseHead) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFunction

`func (o *QueryAssetCardListResponseResponseHead) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *QueryAssetCardListResponseResponseHead) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *QueryAssetCardListResponseResponseHead) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *QueryAssetCardListResponseResponseHead) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetClientId

`func (o *QueryAssetCardListResponseResponseHead) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *QueryAssetCardListResponseResponseHead) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *QueryAssetCardListResponseResponseHead) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *QueryAssetCardListResponseResponseHead) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetRespTime

`func (o *QueryAssetCardListResponseResponseHead) GetRespTime() string`

GetRespTime returns the RespTime field if non-nil, zero value otherwise.

### GetRespTimeOk

`func (o *QueryAssetCardListResponseResponseHead) GetRespTimeOk() (*string, bool)`

GetRespTimeOk returns a tuple with the RespTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespTime

`func (o *QueryAssetCardListResponseResponseHead) SetRespTime(v string)`

SetRespTime sets RespTime field to given value.

### HasRespTime

`func (o *QueryAssetCardListResponseResponseHead) HasRespTime() bool`

HasRespTime returns a boolean if a field has been set.

### GetReqMsgId

`func (o *QueryAssetCardListResponseResponseHead) GetReqMsgId() string`

GetReqMsgId returns the ReqMsgId field if non-nil, zero value otherwise.

### GetReqMsgIdOk

`func (o *QueryAssetCardListResponseResponseHead) GetReqMsgIdOk() (*string, bool)`

GetReqMsgIdOk returns a tuple with the ReqMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqMsgId

`func (o *QueryAssetCardListResponseResponseHead) SetReqMsgId(v string)`

SetReqMsgId sets ReqMsgId field to given value.

### HasReqMsgId

`func (o *QueryAssetCardListResponseResponseHead) HasReqMsgId() bool`

HasReqMsgId returns a boolean if a field has been set.

### GetClientSecret

`func (o *QueryAssetCardListResponseResponseHead) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *QueryAssetCardListResponseResponseHead) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *QueryAssetCardListResponseResponseHead) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *QueryAssetCardListResponseResponseHead) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetReserve

`func (o *QueryAssetCardListResponseResponseHead) GetReserve() interface{}`

GetReserve returns the Reserve field if non-nil, zero value otherwise.

### GetReserveOk

`func (o *QueryAssetCardListResponseResponseHead) GetReserveOk() (*interface{}, bool)`

GetReserveOk returns a tuple with the Reserve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserve

`func (o *QueryAssetCardListResponseResponseHead) SetReserve(v interface{})`

SetReserve sets Reserve field to given value.

### HasReserve

`func (o *QueryAssetCardListResponseResponseHead) HasReserve() bool`

HasReserve returns a boolean if a field has been set.

### SetReserveNil

`func (o *QueryAssetCardListResponseResponseHead) SetReserveNil(b bool)`

 SetReserveNil sets the value for Reserve to be an explicit nil

### UnsetReserve
`func (o *QueryAssetCardListResponseResponseHead) UnsetReserve()`

UnsetReserve ensures that no value is present for Reserve, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


