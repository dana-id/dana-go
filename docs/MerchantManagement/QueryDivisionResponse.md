# QueryDivisionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | [**QueryDivisionResponseResponse**](QueryDivisionResponseResponse.md) |  | 
**Signature** | Pointer to **string** | Signature | [optional] 

## Methods

### NewQueryDivisionResponse

`func NewQueryDivisionResponse(response QueryDivisionResponseResponse, ) *QueryDivisionResponse`

NewQueryDivisionResponse instantiates a new QueryDivisionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDivisionResponseWithDefaults

`func NewQueryDivisionResponseWithDefaults() *QueryDivisionResponse`

NewQueryDivisionResponseWithDefaults instantiates a new QueryDivisionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *QueryDivisionResponse) GetResponse() QueryDivisionResponseResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *QueryDivisionResponse) GetResponseOk() (*QueryDivisionResponseResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *QueryDivisionResponse) SetResponse(v QueryDivisionResponseResponse)`

SetResponse sets Response field to given value.


### GetSignature

`func (o *QueryDivisionResponse) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *QueryDivisionResponse) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *QueryDivisionResponse) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *QueryDivisionResponse) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


