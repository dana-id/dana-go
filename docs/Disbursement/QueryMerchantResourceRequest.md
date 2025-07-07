# QueryMerchantResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | [**QueryMerchantResourceRequestRequest**](QueryMerchantResourceRequestRequest.md) |  | 
**Signature** | **string** | Signature string for request validation | 

## Methods

### NewQueryMerchantResourceRequest

`func NewQueryMerchantResourceRequest(request QueryMerchantResourceRequestRequest, signature string, ) *QueryMerchantResourceRequest`

NewQueryMerchantResourceRequest instantiates a new QueryMerchantResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMerchantResourceRequestWithDefaults

`func NewQueryMerchantResourceRequestWithDefaults() *QueryMerchantResourceRequest`

NewQueryMerchantResourceRequestWithDefaults instantiates a new QueryMerchantResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *QueryMerchantResourceRequest) GetRequest() QueryMerchantResourceRequestRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *QueryMerchantResourceRequest) GetRequestOk() (*QueryMerchantResourceRequestRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *QueryMerchantResourceRequest) SetRequest(v QueryMerchantResourceRequestRequest)`

SetRequest sets Request field to given value.


### GetSignature

`func (o *QueryMerchantResourceRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *QueryMerchantResourceRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *QueryMerchantResourceRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


