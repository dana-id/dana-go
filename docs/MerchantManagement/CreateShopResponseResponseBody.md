# CreateShopResponseResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**ShopId** | Pointer to **string** | The shop ID that was created. Present when resultCodeId is 00000000 | [optional] 

## Methods

### NewCreateShopResponseResponseBody

`func NewCreateShopResponseResponseBody(resultInfo ResultInfo, ) *CreateShopResponseResponseBody`

NewCreateShopResponseResponseBody instantiates a new CreateShopResponseResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShopResponseResponseBodyWithDefaults

`func NewCreateShopResponseResponseBodyWithDefaults() *CreateShopResponseResponseBody`

NewCreateShopResponseResponseBodyWithDefaults instantiates a new CreateShopResponseResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *CreateShopResponseResponseBody) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *CreateShopResponseResponseBody) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *CreateShopResponseResponseBody) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetShopId

`func (o *CreateShopResponseResponseBody) GetShopId() string`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *CreateShopResponseResponseBody) GetShopIdOk() (*string, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *CreateShopResponseResponseBody) SetShopId(v string)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *CreateShopResponseResponseBody) HasShopId() bool`

HasShopId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


