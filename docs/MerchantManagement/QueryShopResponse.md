# QueryShopResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | [**QueryShopResponseResponse**](QueryShopResponseResponse.md) |  | 
**Signature** | Pointer to **string** | Signature | [optional] 

## Methods

### NewQueryShopResponse

`func NewQueryShopResponse(response QueryShopResponseResponse, ) *QueryShopResponse`

NewQueryShopResponse instantiates a new QueryShopResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryShopResponseWithDefaults

`func NewQueryShopResponseWithDefaults() *QueryShopResponse`

NewQueryShopResponseWithDefaults instantiates a new QueryShopResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *QueryShopResponse) GetResponse() QueryShopResponseResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *QueryShopResponse) GetResponseOk() (*QueryShopResponseResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *QueryShopResponse) SetResponse(v QueryShopResponseResponse)`

SetResponse sets Response field to given value.


### GetSignature

`func (o *QueryShopResponse) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *QueryShopResponse) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *QueryShopResponse) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *QueryShopResponse) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


