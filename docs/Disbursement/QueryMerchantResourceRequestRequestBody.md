# QueryMerchantResourceRequestRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestMerchantId** | **string** | This is a merchantId of DANA | 
**MerchantResourceInfoList** | **[]string** | This is a constant merchant resource info request, currently available for value of constant these:<br /> MERCHANT_DEPOSIT_BALANCE<br /> MERCHANT_AVAILABLE_BALANCE<br /> MERCHANT_TOTAL_BALANCE<br /> value for this request can&#39;t be empty or wrong constant info  | 

## Methods

### NewQueryMerchantResourceRequestRequestBody

`func NewQueryMerchantResourceRequestRequestBody(requestMerchantId string, merchantResourceInfoList []string, ) *QueryMerchantResourceRequestRequestBody`

NewQueryMerchantResourceRequestRequestBody instantiates a new QueryMerchantResourceRequestRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMerchantResourceRequestRequestBodyWithDefaults

`func NewQueryMerchantResourceRequestRequestBodyWithDefaults() *QueryMerchantResourceRequestRequestBody`

NewQueryMerchantResourceRequestRequestBodyWithDefaults instantiates a new QueryMerchantResourceRequestRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestMerchantId

`func (o *QueryMerchantResourceRequestRequestBody) GetRequestMerchantId() string`

GetRequestMerchantId returns the RequestMerchantId field if non-nil, zero value otherwise.

### GetRequestMerchantIdOk

`func (o *QueryMerchantResourceRequestRequestBody) GetRequestMerchantIdOk() (*string, bool)`

GetRequestMerchantIdOk returns a tuple with the RequestMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMerchantId

`func (o *QueryMerchantResourceRequestRequestBody) SetRequestMerchantId(v string)`

SetRequestMerchantId sets RequestMerchantId field to given value.


### GetMerchantResourceInfoList

`func (o *QueryMerchantResourceRequestRequestBody) GetMerchantResourceInfoList() []string`

GetMerchantResourceInfoList returns the MerchantResourceInfoList field if non-nil, zero value otherwise.

### GetMerchantResourceInfoListOk

`func (o *QueryMerchantResourceRequestRequestBody) GetMerchantResourceInfoListOk() (*[]string, bool)`

GetMerchantResourceInfoListOk returns a tuple with the MerchantResourceInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantResourceInfoList

`func (o *QueryMerchantResourceRequestRequestBody) SetMerchantResourceInfoList(v []string)`

SetMerchantResourceInfoList sets MerchantResourceInfoList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


