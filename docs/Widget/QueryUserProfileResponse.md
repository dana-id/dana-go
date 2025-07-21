# QueryUserProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | [**QueryUserProfileResponseResponse**](QueryUserProfileResponseResponse.md) |  | 
**Signature** | Pointer to **string** | Signature string | [optional] 

## Methods

### NewQueryUserProfileResponse

`func NewQueryUserProfileResponse(response QueryUserProfileResponseResponse, ) *QueryUserProfileResponse`

NewQueryUserProfileResponse instantiates a new QueryUserProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryUserProfileResponseWithDefaults

`func NewQueryUserProfileResponseWithDefaults() *QueryUserProfileResponse`

NewQueryUserProfileResponseWithDefaults instantiates a new QueryUserProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *QueryUserProfileResponse) GetResponse() QueryUserProfileResponseResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *QueryUserProfileResponse) GetResponseOk() (*QueryUserProfileResponseResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *QueryUserProfileResponse) SetResponse(v QueryUserProfileResponseResponse)`

SetResponse sets Response field to given value.


### GetSignature

`func (o *QueryUserProfileResponse) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *QueryUserProfileResponse) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *QueryUserProfileResponse) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *QueryUserProfileResponse) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


