# QueryMerchantResourceResponseResponseHead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | API version | [optional] [default to "2.0"]
**Function** | Pointer to **string** | API interface | [optional] 
**ClientId** | Pointer to **string** | Client ID provided by DANA, used to identify partner and application system | [optional] 
**RespTime** | Pointer to **string** | DateTime with timezone, which follows the ISO-8601 standard. Refer to RFC 3339 Section 5.6 | [optional] 
**ReqMsgId** | Pointer to **string** | Request message ID | [optional] 

## Methods

### NewQueryMerchantResourceResponseResponseHead

`func NewQueryMerchantResourceResponseResponseHead() *QueryMerchantResourceResponseResponseHead`

NewQueryMerchantResourceResponseResponseHead instantiates a new QueryMerchantResourceResponseResponseHead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMerchantResourceResponseResponseHeadWithDefaults

`func NewQueryMerchantResourceResponseResponseHeadWithDefaults() *QueryMerchantResourceResponseResponseHead`

NewQueryMerchantResourceResponseResponseHeadWithDefaults instantiates a new QueryMerchantResourceResponseResponseHead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *QueryMerchantResourceResponseResponseHead) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *QueryMerchantResourceResponseResponseHead) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *QueryMerchantResourceResponseResponseHead) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *QueryMerchantResourceResponseResponseHead) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFunction

`func (o *QueryMerchantResourceResponseResponseHead) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *QueryMerchantResourceResponseResponseHead) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *QueryMerchantResourceResponseResponseHead) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *QueryMerchantResourceResponseResponseHead) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetClientId

`func (o *QueryMerchantResourceResponseResponseHead) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *QueryMerchantResourceResponseResponseHead) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *QueryMerchantResourceResponseResponseHead) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *QueryMerchantResourceResponseResponseHead) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetRespTime

`func (o *QueryMerchantResourceResponseResponseHead) GetRespTime() string`

GetRespTime returns the RespTime field if non-nil, zero value otherwise.

### GetRespTimeOk

`func (o *QueryMerchantResourceResponseResponseHead) GetRespTimeOk() (*string, bool)`

GetRespTimeOk returns a tuple with the RespTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespTime

`func (o *QueryMerchantResourceResponseResponseHead) SetRespTime(v string)`

SetRespTime sets RespTime field to given value.

### HasRespTime

`func (o *QueryMerchantResourceResponseResponseHead) HasRespTime() bool`

HasRespTime returns a boolean if a field has been set.

### GetReqMsgId

`func (o *QueryMerchantResourceResponseResponseHead) GetReqMsgId() string`

GetReqMsgId returns the ReqMsgId field if non-nil, zero value otherwise.

### GetReqMsgIdOk

`func (o *QueryMerchantResourceResponseResponseHead) GetReqMsgIdOk() (*string, bool)`

GetReqMsgIdOk returns a tuple with the ReqMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqMsgId

`func (o *QueryMerchantResourceResponseResponseHead) SetReqMsgId(v string)`

SetReqMsgId sets ReqMsgId field to given value.

### HasReqMsgId

`func (o *QueryMerchantResourceResponseResponseHead) HasReqMsgId() bool`

HasReqMsgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


