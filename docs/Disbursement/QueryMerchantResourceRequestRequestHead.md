# QueryMerchantResourceRequestRequestHead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | API version | [default to "2.0"]
**Function** | **string** | API interface | [default to "dana.merchant.queryMerchantResource"]
**ClientId** | **string** | Client ID provided by DANA, used to identify partner and application system | 
**ReqTime** | **time.Time** | DateTime with timezone, which follows the ISO-8601 standard. Refer to RFC 3339 Section 5.6 | 
**ReqMsgId** | **string** | The sequence id of request - Each request will be assigned with a unique id (uuid) | 
**ClientSecret** | **string** | Secret key of client - Assigned client secret during registration | 
**Reserve** | Pointer to **string** | Reserved for future implementation | [optional] 

## Methods

### NewQueryMerchantResourceRequestRequestHead

`func NewQueryMerchantResourceRequestRequestHead(version string, function string, clientId string, reqTime time.Time, reqMsgId string, clientSecret string, ) *QueryMerchantResourceRequestRequestHead`

NewQueryMerchantResourceRequestRequestHead instantiates a new QueryMerchantResourceRequestRequestHead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMerchantResourceRequestRequestHeadWithDefaults

`func NewQueryMerchantResourceRequestRequestHeadWithDefaults() *QueryMerchantResourceRequestRequestHead`

NewQueryMerchantResourceRequestRequestHeadWithDefaults instantiates a new QueryMerchantResourceRequestRequestHead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *QueryMerchantResourceRequestRequestHead) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *QueryMerchantResourceRequestRequestHead) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *QueryMerchantResourceRequestRequestHead) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetFunction

`func (o *QueryMerchantResourceRequestRequestHead) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *QueryMerchantResourceRequestRequestHead) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *QueryMerchantResourceRequestRequestHead) SetFunction(v string)`

SetFunction sets Function field to given value.


### GetClientId

`func (o *QueryMerchantResourceRequestRequestHead) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *QueryMerchantResourceRequestRequestHead) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *QueryMerchantResourceRequestRequestHead) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetReqTime

`func (o *QueryMerchantResourceRequestRequestHead) GetReqTime() time.Time`

GetReqTime returns the ReqTime field if non-nil, zero value otherwise.

### GetReqTimeOk

`func (o *QueryMerchantResourceRequestRequestHead) GetReqTimeOk() (*time.Time, bool)`

GetReqTimeOk returns a tuple with the ReqTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTime

`func (o *QueryMerchantResourceRequestRequestHead) SetReqTime(v time.Time)`

SetReqTime sets ReqTime field to given value.


### GetReqMsgId

`func (o *QueryMerchantResourceRequestRequestHead) GetReqMsgId() string`

GetReqMsgId returns the ReqMsgId field if non-nil, zero value otherwise.

### GetReqMsgIdOk

`func (o *QueryMerchantResourceRequestRequestHead) GetReqMsgIdOk() (*string, bool)`

GetReqMsgIdOk returns a tuple with the ReqMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqMsgId

`func (o *QueryMerchantResourceRequestRequestHead) SetReqMsgId(v string)`

SetReqMsgId sets ReqMsgId field to given value.


### GetClientSecret

`func (o *QueryMerchantResourceRequestRequestHead) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *QueryMerchantResourceRequestRequestHead) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *QueryMerchantResourceRequestRequestHead) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetReserve

`func (o *QueryMerchantResourceRequestRequestHead) GetReserve() string`

GetReserve returns the Reserve field if non-nil, zero value otherwise.

### GetReserveOk

`func (o *QueryMerchantResourceRequestRequestHead) GetReserveOk() (*string, bool)`

GetReserveOk returns a tuple with the Reserve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserve

`func (o *QueryMerchantResourceRequestRequestHead) SetReserve(v string)`

SetReserve sets Reserve field to given value.

### HasReserve

`func (o *QueryMerchantResourceRequestRequestHead) HasReserve() bool`

HasReserve returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


