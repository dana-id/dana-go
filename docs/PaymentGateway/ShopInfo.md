# ShopInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopId** | Pointer to **string** | Information of shop identifier. Required if externalShopId is blank | [optional] 
**ExternalShopId** | Pointer to **string** | Information of external shop identifier. Required if shopId is blank | [optional] 
**OperatorId** | Pointer to **string** | Information of operator identifier | [optional] 
**ShopAddress** | Pointer to **string** | Information of shop address | [optional] 
**DivisionId** | Pointer to **string** | Information of division identifier | [optional] 
**ExternalDivisionId** | Pointer to **string** | Information of external division identifier | [optional] 
**DivisionType** | Pointer to **string** | Information of division type | [optional] 
**ShopName** | Pointer to **string** | Information of shop name | [optional] 

## Methods

### NewShopInfo

`func NewShopInfo() *ShopInfo`

NewShopInfo instantiates a new ShopInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopInfoWithDefaults

`func NewShopInfoWithDefaults() *ShopInfo`

NewShopInfoWithDefaults instantiates a new ShopInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopId

`func (o *ShopInfo) GetShopId() string`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ShopInfo) GetShopIdOk() (*string, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ShopInfo) SetShopId(v string)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ShopInfo) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetExternalShopId

`func (o *ShopInfo) GetExternalShopId() string`

GetExternalShopId returns the ExternalShopId field if non-nil, zero value otherwise.

### GetExternalShopIdOk

`func (o *ShopInfo) GetExternalShopIdOk() (*string, bool)`

GetExternalShopIdOk returns a tuple with the ExternalShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShopId

`func (o *ShopInfo) SetExternalShopId(v string)`

SetExternalShopId sets ExternalShopId field to given value.

### HasExternalShopId

`func (o *ShopInfo) HasExternalShopId() bool`

HasExternalShopId returns a boolean if a field has been set.

### GetOperatorId

`func (o *ShopInfo) GetOperatorId() string`

GetOperatorId returns the OperatorId field if non-nil, zero value otherwise.

### GetOperatorIdOk

`func (o *ShopInfo) GetOperatorIdOk() (*string, bool)`

GetOperatorIdOk returns a tuple with the OperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorId

`func (o *ShopInfo) SetOperatorId(v string)`

SetOperatorId sets OperatorId field to given value.

### HasOperatorId

`func (o *ShopInfo) HasOperatorId() bool`

HasOperatorId returns a boolean if a field has been set.

### GetShopAddress

`func (o *ShopInfo) GetShopAddress() string`

GetShopAddress returns the ShopAddress field if non-nil, zero value otherwise.

### GetShopAddressOk

`func (o *ShopInfo) GetShopAddressOk() (*string, bool)`

GetShopAddressOk returns a tuple with the ShopAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopAddress

`func (o *ShopInfo) SetShopAddress(v string)`

SetShopAddress sets ShopAddress field to given value.

### HasShopAddress

`func (o *ShopInfo) HasShopAddress() bool`

HasShopAddress returns a boolean if a field has been set.

### GetDivisionId

`func (o *ShopInfo) GetDivisionId() string`

GetDivisionId returns the DivisionId field if non-nil, zero value otherwise.

### GetDivisionIdOk

`func (o *ShopInfo) GetDivisionIdOk() (*string, bool)`

GetDivisionIdOk returns a tuple with the DivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionId

`func (o *ShopInfo) SetDivisionId(v string)`

SetDivisionId sets DivisionId field to given value.

### HasDivisionId

`func (o *ShopInfo) HasDivisionId() bool`

HasDivisionId returns a boolean if a field has been set.

### GetExternalDivisionId

`func (o *ShopInfo) GetExternalDivisionId() string`

GetExternalDivisionId returns the ExternalDivisionId field if non-nil, zero value otherwise.

### GetExternalDivisionIdOk

`func (o *ShopInfo) GetExternalDivisionIdOk() (*string, bool)`

GetExternalDivisionIdOk returns a tuple with the ExternalDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDivisionId

`func (o *ShopInfo) SetExternalDivisionId(v string)`

SetExternalDivisionId sets ExternalDivisionId field to given value.

### HasExternalDivisionId

`func (o *ShopInfo) HasExternalDivisionId() bool`

HasExternalDivisionId returns a boolean if a field has been set.

### GetDivisionType

`func (o *ShopInfo) GetDivisionType() string`

GetDivisionType returns the DivisionType field if non-nil, zero value otherwise.

### GetDivisionTypeOk

`func (o *ShopInfo) GetDivisionTypeOk() (*string, bool)`

GetDivisionTypeOk returns a tuple with the DivisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionType

`func (o *ShopInfo) SetDivisionType(v string)`

SetDivisionType sets DivisionType field to given value.

### HasDivisionType

`func (o *ShopInfo) HasDivisionType() bool`

HasDivisionType returns a boolean if a field has been set.

### GetShopName

`func (o *ShopInfo) GetShopName() string`

GetShopName returns the ShopName field if non-nil, zero value otherwise.

### GetShopNameOk

`func (o *ShopInfo) GetShopNameOk() (*string, bool)`

GetShopNameOk returns a tuple with the ShopName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopName

`func (o *ShopInfo) SetShopName(v string)`

SetShopName sets ShopName field to given value.

### HasShopName

`func (o *ShopInfo) HasShopName() bool`

HasShopName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


