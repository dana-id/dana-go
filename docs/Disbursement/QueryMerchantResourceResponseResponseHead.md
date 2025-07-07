# QueryMerchantResourceResponseResponseHead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | API version | [default to "2.0"]
**Function** | **string** | API interface | 
**ClientId** | **string** | Client ID provided by DANA, used to identify partner and application system | 
**RespTime** | **time.Time** | DateTime with timezone, which follows the ISO-8601 standard. Refer to RFC 3339 Section 5.6 | 
**ReqMsgId** | **string** | Request message ID | 

## Methods

### NewQueryMerchantResourceResponseResponseHead

`func NewQueryMerchantResourceResponseResponseHead(version string, function string, clientId string, respTime time.Time, reqMsgId string, ) *QueryMerchantResourceResponseResponseHead`

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


### GetRespTime

`func (o *QueryMerchantResourceResponseResponseHead) GetRespTime() time.Time`

GetRespTime returns the RespTime field if non-nil, zero value otherwise.

### GetRespTimeOk

`func (o *QueryMerchantResourceResponseResponseHead) GetRespTimeOk() (*time.Time, bool)`

GetRespTimeOk returns a tuple with the RespTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespTime

`func (o *QueryMerchantResourceResponseResponseHead) SetRespTime(v time.Time)`

SetRespTime sets RespTime field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


