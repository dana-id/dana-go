# ApplyOTTResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **string** | Response code. https://dashboard.dana.id/api-docs/read/109#HTML-API-ApplyOTT-ResponseCodeandMessage | 
**ResponseMessage** | **string** | Response message. https://dashboard.dana.id/api-docs/read/109#HTML-API-ApplyOTT-ResponseCodeandMessage | 
**UserResources** | [**[]ApplyOTTResponseUserResourcesInner**](ApplyOTTResponseUserResourcesInner.md) | User resources | 
**AdditionalInfo** | Pointer to **map[string]interface{}** | Additional information | [optional] 

## Methods

### NewApplyOTTResponse

`func NewApplyOTTResponse(responseCode string, responseMessage string, userResources []ApplyOTTResponseUserResourcesInner, ) *ApplyOTTResponse`

NewApplyOTTResponse instantiates a new ApplyOTTResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyOTTResponseWithDefaults

`func NewApplyOTTResponseWithDefaults() *ApplyOTTResponse`

NewApplyOTTResponseWithDefaults instantiates a new ApplyOTTResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *ApplyOTTResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *ApplyOTTResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *ApplyOTTResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseMessage

`func (o *ApplyOTTResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *ApplyOTTResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *ApplyOTTResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetUserResources

`func (o *ApplyOTTResponse) GetUserResources() []ApplyOTTResponseUserResourcesInner`

GetUserResources returns the UserResources field if non-nil, zero value otherwise.

### GetUserResourcesOk

`func (o *ApplyOTTResponse) GetUserResourcesOk() (*[]ApplyOTTResponseUserResourcesInner, bool)`

GetUserResourcesOk returns a tuple with the UserResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResources

`func (o *ApplyOTTResponse) SetUserResources(v []ApplyOTTResponseUserResourcesInner)`

SetUserResources sets UserResources field to given value.


### GetAdditionalInfo

`func (o *ApplyOTTResponse) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ApplyOTTResponse) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ApplyOTTResponse) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ApplyOTTResponse) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


