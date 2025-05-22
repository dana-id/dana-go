# CreateOrderBaseAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mcc** | **string** | Additional information of merchant category code. This parameter is used to identify the type of business in which a merchant is engaged. Refer to Details of https://dashboard.dana.id/api-docs/read/197#OpenAPI-MerchantCategoryCode | 
**ExtendInfo** | Pointer to **string** | Additional information of extend such as partner passthrough and risk information | [optional] 
**EnvInfo** | [**EnvInfo**](EnvInfo.md) | Additional information of environment info | 

## Methods

### NewCreateOrderBaseAdditionalInfo

`func NewCreateOrderBaseAdditionalInfo(mcc string, envInfo EnvInfo, ) *CreateOrderBaseAdditionalInfo`

NewCreateOrderBaseAdditionalInfo instantiates a new CreateOrderBaseAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderBaseAdditionalInfoWithDefaults

`func NewCreateOrderBaseAdditionalInfoWithDefaults() *CreateOrderBaseAdditionalInfo`

NewCreateOrderBaseAdditionalInfoWithDefaults instantiates a new CreateOrderBaseAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMcc

`func (o *CreateOrderBaseAdditionalInfo) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *CreateOrderBaseAdditionalInfo) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *CreateOrderBaseAdditionalInfo) SetMcc(v string)`

SetMcc sets Mcc field to given value.


### GetExtendInfo

`func (o *CreateOrderBaseAdditionalInfo) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *CreateOrderBaseAdditionalInfo) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *CreateOrderBaseAdditionalInfo) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *CreateOrderBaseAdditionalInfo) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.

### GetEnvInfo

`func (o *CreateOrderBaseAdditionalInfo) GetEnvInfo() EnvInfo`

GetEnvInfo returns the EnvInfo field if non-nil, zero value otherwise.

### GetEnvInfoOk

`func (o *CreateOrderBaseAdditionalInfo) GetEnvInfoOk() (*EnvInfo, bool)`

GetEnvInfoOk returns a tuple with the EnvInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvInfo

`func (o *CreateOrderBaseAdditionalInfo) SetEnvInfo(v EnvInfo)`

SetEnvInfo sets EnvInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


