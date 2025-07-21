# ShopResourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | Pointer to **string** | Merchant identifier | [optional] 
**ParentDivisionId** | Pointer to **string** | Parent division identifier | [optional] 
**ParentRoleType** | Pointer to **string** | Parent role type | [optional] 
**MainName** | Pointer to **string** | Shop name | [optional] 
**SizeType** | Pointer to **string** | Size type | [optional] 
**ShopAddress** | Pointer to [**AddressInfo**](AddressInfo.md) |  | [optional] 
**ExternalShopId** | Pointer to **string** | External shop identifier | [optional] 
**LogoUrlMap** | Pointer to **map[string]string** | Logo URL map with base64 encoded images | [optional] 
**ExtInfo** | Pointer to **map[string]interface{}** | Extended information | [optional] 
**Ln** | Pointer to **string** | Longitude | [optional] 
**Lat** | Pointer to **string** | Latitude | [optional] 
**Nmid** | Pointer to **string** | Network merchant identifier | [optional] 

## Methods

### NewShopResourceInfo

`func NewShopResourceInfo() *ShopResourceInfo`

NewShopResourceInfo instantiates a new ShopResourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopResourceInfoWithDefaults

`func NewShopResourceInfoWithDefaults() *ShopResourceInfo`

NewShopResourceInfoWithDefaults instantiates a new ShopResourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *ShopResourceInfo) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *ShopResourceInfo) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *ShopResourceInfo) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *ShopResourceInfo) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetParentDivisionId

`func (o *ShopResourceInfo) GetParentDivisionId() string`

GetParentDivisionId returns the ParentDivisionId field if non-nil, zero value otherwise.

### GetParentDivisionIdOk

`func (o *ShopResourceInfo) GetParentDivisionIdOk() (*string, bool)`

GetParentDivisionIdOk returns a tuple with the ParentDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDivisionId

`func (o *ShopResourceInfo) SetParentDivisionId(v string)`

SetParentDivisionId sets ParentDivisionId field to given value.

### HasParentDivisionId

`func (o *ShopResourceInfo) HasParentDivisionId() bool`

HasParentDivisionId returns a boolean if a field has been set.

### GetParentRoleType

`func (o *ShopResourceInfo) GetParentRoleType() string`

GetParentRoleType returns the ParentRoleType field if non-nil, zero value otherwise.

### GetParentRoleTypeOk

`func (o *ShopResourceInfo) GetParentRoleTypeOk() (*string, bool)`

GetParentRoleTypeOk returns a tuple with the ParentRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoleType

`func (o *ShopResourceInfo) SetParentRoleType(v string)`

SetParentRoleType sets ParentRoleType field to given value.

### HasParentRoleType

`func (o *ShopResourceInfo) HasParentRoleType() bool`

HasParentRoleType returns a boolean if a field has been set.

### GetMainName

`func (o *ShopResourceInfo) GetMainName() string`

GetMainName returns the MainName field if non-nil, zero value otherwise.

### GetMainNameOk

`func (o *ShopResourceInfo) GetMainNameOk() (*string, bool)`

GetMainNameOk returns a tuple with the MainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainName

`func (o *ShopResourceInfo) SetMainName(v string)`

SetMainName sets MainName field to given value.

### HasMainName

`func (o *ShopResourceInfo) HasMainName() bool`

HasMainName returns a boolean if a field has been set.

### GetSizeType

`func (o *ShopResourceInfo) GetSizeType() string`

GetSizeType returns the SizeType field if non-nil, zero value otherwise.

### GetSizeTypeOk

`func (o *ShopResourceInfo) GetSizeTypeOk() (*string, bool)`

GetSizeTypeOk returns a tuple with the SizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeType

`func (o *ShopResourceInfo) SetSizeType(v string)`

SetSizeType sets SizeType field to given value.

### HasSizeType

`func (o *ShopResourceInfo) HasSizeType() bool`

HasSizeType returns a boolean if a field has been set.

### GetShopAddress

`func (o *ShopResourceInfo) GetShopAddress() AddressInfo`

GetShopAddress returns the ShopAddress field if non-nil, zero value otherwise.

### GetShopAddressOk

`func (o *ShopResourceInfo) GetShopAddressOk() (*AddressInfo, bool)`

GetShopAddressOk returns a tuple with the ShopAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopAddress

`func (o *ShopResourceInfo) SetShopAddress(v AddressInfo)`

SetShopAddress sets ShopAddress field to given value.

### HasShopAddress

`func (o *ShopResourceInfo) HasShopAddress() bool`

HasShopAddress returns a boolean if a field has been set.

### GetExternalShopId

`func (o *ShopResourceInfo) GetExternalShopId() string`

GetExternalShopId returns the ExternalShopId field if non-nil, zero value otherwise.

### GetExternalShopIdOk

`func (o *ShopResourceInfo) GetExternalShopIdOk() (*string, bool)`

GetExternalShopIdOk returns a tuple with the ExternalShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShopId

`func (o *ShopResourceInfo) SetExternalShopId(v string)`

SetExternalShopId sets ExternalShopId field to given value.

### HasExternalShopId

`func (o *ShopResourceInfo) HasExternalShopId() bool`

HasExternalShopId returns a boolean if a field has been set.

### GetLogoUrlMap

`func (o *ShopResourceInfo) GetLogoUrlMap() map[string]string`

GetLogoUrlMap returns the LogoUrlMap field if non-nil, zero value otherwise.

### GetLogoUrlMapOk

`func (o *ShopResourceInfo) GetLogoUrlMapOk() (*map[string]string, bool)`

GetLogoUrlMapOk returns a tuple with the LogoUrlMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrlMap

`func (o *ShopResourceInfo) SetLogoUrlMap(v map[string]string)`

SetLogoUrlMap sets LogoUrlMap field to given value.

### HasLogoUrlMap

`func (o *ShopResourceInfo) HasLogoUrlMap() bool`

HasLogoUrlMap returns a boolean if a field has been set.

### GetExtInfo

`func (o *ShopResourceInfo) GetExtInfo() map[string]interface{}`

GetExtInfo returns the ExtInfo field if non-nil, zero value otherwise.

### GetExtInfoOk

`func (o *ShopResourceInfo) GetExtInfoOk() (*map[string]interface{}, bool)`

GetExtInfoOk returns a tuple with the ExtInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtInfo

`func (o *ShopResourceInfo) SetExtInfo(v map[string]interface{})`

SetExtInfo sets ExtInfo field to given value.

### HasExtInfo

`func (o *ShopResourceInfo) HasExtInfo() bool`

HasExtInfo returns a boolean if a field has been set.

### GetLn

`func (o *ShopResourceInfo) GetLn() string`

GetLn returns the Ln field if non-nil, zero value otherwise.

### GetLnOk

`func (o *ShopResourceInfo) GetLnOk() (*string, bool)`

GetLnOk returns a tuple with the Ln field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLn

`func (o *ShopResourceInfo) SetLn(v string)`

SetLn sets Ln field to given value.

### HasLn

`func (o *ShopResourceInfo) HasLn() bool`

HasLn returns a boolean if a field has been set.

### GetLat

`func (o *ShopResourceInfo) GetLat() string`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *ShopResourceInfo) GetLatOk() (*string, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *ShopResourceInfo) SetLat(v string)`

SetLat sets Lat field to given value.

### HasLat

`func (o *ShopResourceInfo) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetNmid

`func (o *ShopResourceInfo) GetNmid() string`

GetNmid returns the Nmid field if non-nil, zero value otherwise.

### GetNmidOk

`func (o *ShopResourceInfo) GetNmidOk() (*string, bool)`

GetNmidOk returns a tuple with the Nmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmid

`func (o *ShopResourceInfo) SetNmid(v string)`

SetNmid sets Nmid field to given value.

### HasNmid

`func (o *ShopResourceInfo) HasNmid() bool`

HasNmid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


