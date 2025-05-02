# CreateOrderByRedirectAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mcc** | **string** |  | 
**ExtendInfo** | Pointer to **string** |  | [optional] 
**EnvInfo** | [**EnvInfo**](EnvInfo.md) |  | 
**Order** | Pointer to [**OrderRedirectObject**](OrderRedirectObject.md) |  | [optional] 

## Methods

### NewCreateOrderByRedirectAdditionalInfo

`func NewCreateOrderByRedirectAdditionalInfo(mcc string, envInfo EnvInfo, ) *CreateOrderByRedirectAdditionalInfo`

NewCreateOrderByRedirectAdditionalInfo instantiates a new CreateOrderByRedirectAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderByRedirectAdditionalInfoWithDefaults

`func NewCreateOrderByRedirectAdditionalInfoWithDefaults() *CreateOrderByRedirectAdditionalInfo`

NewCreateOrderByRedirectAdditionalInfoWithDefaults instantiates a new CreateOrderByRedirectAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMcc

`func (o *CreateOrderByRedirectAdditionalInfo) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *CreateOrderByRedirectAdditionalInfo) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *CreateOrderByRedirectAdditionalInfo) SetMcc(v string)`

SetMcc sets Mcc field to given value.


### GetExtendInfo

`func (o *CreateOrderByRedirectAdditionalInfo) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *CreateOrderByRedirectAdditionalInfo) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *CreateOrderByRedirectAdditionalInfo) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *CreateOrderByRedirectAdditionalInfo) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.

### GetEnvInfo

`func (o *CreateOrderByRedirectAdditionalInfo) GetEnvInfo() EnvInfo`

GetEnvInfo returns the EnvInfo field if non-nil, zero value otherwise.

### GetEnvInfoOk

`func (o *CreateOrderByRedirectAdditionalInfo) GetEnvInfoOk() (*EnvInfo, bool)`

GetEnvInfoOk returns a tuple with the EnvInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvInfo

`func (o *CreateOrderByRedirectAdditionalInfo) SetEnvInfo(v EnvInfo)`

SetEnvInfo sets EnvInfo field to given value.


### GetOrder

`func (o *CreateOrderByRedirectAdditionalInfo) GetOrder() OrderRedirectObject`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CreateOrderByRedirectAdditionalInfo) GetOrderOk() (*OrderRedirectObject, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CreateOrderByRedirectAdditionalInfo) SetOrder(v OrderRedirectObject)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CreateOrderByRedirectAdditionalInfo) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


