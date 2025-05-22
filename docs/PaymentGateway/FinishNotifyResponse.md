# FinishNotifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **string** | Response code. Refer to https://dashboard.dana.id/api-docs/read/123#HTML-API-FinishNotify-ResponseCodeandMessage | 
**ResponseMessage** | **string** | Response message. Refer to https://dashboard.dana.id/api-docs/read/123#HTML-API-FinishNotify-ResponseCodeandMessage | 

## Methods

### NewFinishNotifyResponse

`func NewFinishNotifyResponse(responseCode string, responseMessage string, ) *FinishNotifyResponse`

NewFinishNotifyResponse instantiates a new FinishNotifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinishNotifyResponseWithDefaults

`func NewFinishNotifyResponseWithDefaults() *FinishNotifyResponse`

NewFinishNotifyResponseWithDefaults instantiates a new FinishNotifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *FinishNotifyResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *FinishNotifyResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *FinishNotifyResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseMessage

`func (o *FinishNotifyResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *FinishNotifyResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *FinishNotifyResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


