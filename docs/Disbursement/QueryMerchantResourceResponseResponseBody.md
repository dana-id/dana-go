# QueryMerchantResourceResponseResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**MerchantResourceInformations** | Pointer to [**[]MerchantResourceInformation**](MerchantResourceInformation.md) | Merchant resource information list - will be filled if success | [optional] 

## Methods

### NewQueryMerchantResourceResponseResponseBody

`func NewQueryMerchantResourceResponseResponseBody(resultInfo ResultInfo, ) *QueryMerchantResourceResponseResponseBody`

NewQueryMerchantResourceResponseResponseBody instantiates a new QueryMerchantResourceResponseResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMerchantResourceResponseResponseBodyWithDefaults

`func NewQueryMerchantResourceResponseResponseBodyWithDefaults() *QueryMerchantResourceResponseResponseBody`

NewQueryMerchantResourceResponseResponseBodyWithDefaults instantiates a new QueryMerchantResourceResponseResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *QueryMerchantResourceResponseResponseBody) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *QueryMerchantResourceResponseResponseBody) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *QueryMerchantResourceResponseResponseBody) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetMerchantResourceInformations

`func (o *QueryMerchantResourceResponseResponseBody) GetMerchantResourceInformations() []MerchantResourceInformation`

GetMerchantResourceInformations returns the MerchantResourceInformations field if non-nil, zero value otherwise.

### GetMerchantResourceInformationsOk

`func (o *QueryMerchantResourceResponseResponseBody) GetMerchantResourceInformationsOk() (*[]MerchantResourceInformation, bool)`

GetMerchantResourceInformationsOk returns a tuple with the MerchantResourceInformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantResourceInformations

`func (o *QueryMerchantResourceResponseResponseBody) SetMerchantResourceInformations(v []MerchantResourceInformation)`

SetMerchantResourceInformations sets MerchantResourceInformations field to given value.

### HasMerchantResourceInformations

`func (o *QueryMerchantResourceResponseResponseBody) HasMerchantResourceInformations() bool`

HasMerchantResourceInformations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


