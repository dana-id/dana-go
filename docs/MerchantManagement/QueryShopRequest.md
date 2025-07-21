# QueryShopRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | Pointer to **string** | Merchant identifier. Required when shopIdType is EXTERNAL_ID | [optional] 
**ShopId** | **string** | Shop identifier. Length depends on shopIdType - INNER_ID (21 max), EXTERNAL_ID (64 max) | 
**ShopIdType** | **string** | Shop identifier type | 

## Methods

### NewQueryShopRequest

`func NewQueryShopRequest(shopId string, shopIdType string, ) *QueryShopRequest`

NewQueryShopRequest instantiates a new QueryShopRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryShopRequestWithDefaults

`func NewQueryShopRequestWithDefaults() *QueryShopRequest`

NewQueryShopRequestWithDefaults instantiates a new QueryShopRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *QueryShopRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *QueryShopRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *QueryShopRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *QueryShopRequest) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetShopId

`func (o *QueryShopRequest) GetShopId() string`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *QueryShopRequest) GetShopIdOk() (*string, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *QueryShopRequest) SetShopId(v string)`

SetShopId sets ShopId field to given value.


### GetShopIdType

`func (o *QueryShopRequest) GetShopIdType() string`

GetShopIdType returns the ShopIdType field if non-nil, zero value otherwise.

### GetShopIdTypeOk

`func (o *QueryShopRequest) GetShopIdTypeOk() (*string, bool)`

GetShopIdTypeOk returns a tuple with the ShopIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopIdType

`func (o *QueryShopRequest) SetShopIdType(v string)`

SetShopIdType sets ShopIdType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


