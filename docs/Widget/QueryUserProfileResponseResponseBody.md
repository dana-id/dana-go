# QueryUserProfileResponseResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**UserResourceInfos** | Pointer to [**[]UserResourceInfo**](UserResourceInfo.md) | The querying resource list value | [optional] 

## Methods

### NewQueryUserProfileResponseResponseBody

`func NewQueryUserProfileResponseResponseBody(resultInfo ResultInfo, ) *QueryUserProfileResponseResponseBody`

NewQueryUserProfileResponseResponseBody instantiates a new QueryUserProfileResponseResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryUserProfileResponseResponseBodyWithDefaults

`func NewQueryUserProfileResponseResponseBodyWithDefaults() *QueryUserProfileResponseResponseBody`

NewQueryUserProfileResponseResponseBodyWithDefaults instantiates a new QueryUserProfileResponseResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *QueryUserProfileResponseResponseBody) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *QueryUserProfileResponseResponseBody) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *QueryUserProfileResponseResponseBody) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetUserResourceInfos

`func (o *QueryUserProfileResponseResponseBody) GetUserResourceInfos() []UserResourceInfo`

GetUserResourceInfos returns the UserResourceInfos field if non-nil, zero value otherwise.

### GetUserResourceInfosOk

`func (o *QueryUserProfileResponseResponseBody) GetUserResourceInfosOk() (*[]UserResourceInfo, bool)`

GetUserResourceInfosOk returns a tuple with the UserResourceInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResourceInfos

`func (o *QueryUserProfileResponseResponseBody) SetUserResourceInfos(v []UserResourceInfo)`

SetUserResourceInfos sets UserResourceInfos field to given value.

### HasUserResourceInfos

`func (o *QueryUserProfileResponseResponseBody) HasUserResourceInfos() bool`

HasUserResourceInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


