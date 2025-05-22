# CreateOrderByApiAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mcc** | **string** | Additional information of merchant category code. This parameter is used to identify the type of business in which a merchant is engaged. Refer to Details of https://dashboard.dana.id/api-docs/read/197#OpenAPI-MerchantCategoryCode | 
**ExtendInfo** | Pointer to **string** | Additional information of extend such as partner passthrough and risk information | [optional] 
**EnvInfo** | [**EnvInfo**](EnvInfo.md) | Additional information of environment info | 
**Order** | Pointer to [**OrderApiObject**](OrderApiObject.md) | Additional information of order | [optional] 

## Methods

### NewCreateOrderByApiAdditionalInfo

`func NewCreateOrderByApiAdditionalInfo(mcc string, envInfo EnvInfo, ) *CreateOrderByApiAdditionalInfo`

NewCreateOrderByApiAdditionalInfo instantiates a new CreateOrderByApiAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderByApiAdditionalInfoWithDefaults

`func NewCreateOrderByApiAdditionalInfoWithDefaults() *CreateOrderByApiAdditionalInfo`

NewCreateOrderByApiAdditionalInfoWithDefaults instantiates a new CreateOrderByApiAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMcc

`func (o *CreateOrderByApiAdditionalInfo) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *CreateOrderByApiAdditionalInfo) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *CreateOrderByApiAdditionalInfo) SetMcc(v string)`

SetMcc sets Mcc field to given value.


### GetExtendInfo

`func (o *CreateOrderByApiAdditionalInfo) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *CreateOrderByApiAdditionalInfo) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *CreateOrderByApiAdditionalInfo) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *CreateOrderByApiAdditionalInfo) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.

### GetEnvInfo

`func (o *CreateOrderByApiAdditionalInfo) GetEnvInfo() EnvInfo`

GetEnvInfo returns the EnvInfo field if non-nil, zero value otherwise.

### GetEnvInfoOk

`func (o *CreateOrderByApiAdditionalInfo) GetEnvInfoOk() (*EnvInfo, bool)`

GetEnvInfoOk returns a tuple with the EnvInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvInfo

`func (o *CreateOrderByApiAdditionalInfo) SetEnvInfo(v EnvInfo)`

SetEnvInfo sets EnvInfo field to given value.


### GetOrder

`func (o *CreateOrderByApiAdditionalInfo) GetOrder() OrderApiObject`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CreateOrderByApiAdditionalInfo) GetOrderOk() (*OrderApiObject, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CreateOrderByApiAdditionalInfo) SetOrder(v OrderApiObject)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CreateOrderByApiAdditionalInfo) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


