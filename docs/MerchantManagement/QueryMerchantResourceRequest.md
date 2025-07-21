# QueryMerchantResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestMerchantId** | **string** | This is a merchantId of DANA | 
**MerchantResourceInfoList** | **[]string** | This is a constant merchant resource info request, currently available for value of constant these: MERCHANT_DEPOSIT_BALANCE MERCHANT_AVAILABLE_BALANCE MERCHANT_TOTAL_BALANCE value for this request can&#39;t be empty or wrong constant info  | 

## Methods

### NewQueryMerchantResourceRequest

`func NewQueryMerchantResourceRequest(requestMerchantId string, merchantResourceInfoList []string, ) *QueryMerchantResourceRequest`

NewQueryMerchantResourceRequest instantiates a new QueryMerchantResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMerchantResourceRequestWithDefaults

`func NewQueryMerchantResourceRequestWithDefaults() *QueryMerchantResourceRequest`

NewQueryMerchantResourceRequestWithDefaults instantiates a new QueryMerchantResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestMerchantId

`func (o *QueryMerchantResourceRequest) GetRequestMerchantId() string`

GetRequestMerchantId returns the RequestMerchantId field if non-nil, zero value otherwise.

### GetRequestMerchantIdOk

`func (o *QueryMerchantResourceRequest) GetRequestMerchantIdOk() (*string, bool)`

GetRequestMerchantIdOk returns a tuple with the RequestMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMerchantId

`func (o *QueryMerchantResourceRequest) SetRequestMerchantId(v string)`

SetRequestMerchantId sets RequestMerchantId field to given value.


### GetMerchantResourceInfoList

`func (o *QueryMerchantResourceRequest) GetMerchantResourceInfoList() []string`

GetMerchantResourceInfoList returns the MerchantResourceInfoList field if non-nil, zero value otherwise.

### GetMerchantResourceInfoListOk

`func (o *QueryMerchantResourceRequest) GetMerchantResourceInfoListOk() (*[]string, bool)`

GetMerchantResourceInfoListOk returns a tuple with the MerchantResourceInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantResourceInfoList

`func (o *QueryMerchantResourceRequest) SetMerchantResourceInfoList(v []string)`

SetMerchantResourceInfoList sets MerchantResourceInfoList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


