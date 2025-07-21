# QueryUserProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserResources** | **[]string** | The resource type list that the merchant server wants to get from DANA | 

## Methods

### NewQueryUserProfileRequest

`func NewQueryUserProfileRequest(userResources []string, ) *QueryUserProfileRequest`

NewQueryUserProfileRequest instantiates a new QueryUserProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryUserProfileRequestWithDefaults

`func NewQueryUserProfileRequestWithDefaults() *QueryUserProfileRequest`

NewQueryUserProfileRequestWithDefaults instantiates a new QueryUserProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserResources

`func (o *QueryUserProfileRequest) GetUserResources() []string`

GetUserResources returns the UserResources field if non-nil, zero value otherwise.

### GetUserResourcesOk

`func (o *QueryUserProfileRequest) GetUserResourcesOk() (*[]string, bool)`

GetUserResourcesOk returns a tuple with the UserResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResources

`func (o *QueryUserProfileRequest) SetUserResources(v []string)`

SetUserResources sets UserResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


