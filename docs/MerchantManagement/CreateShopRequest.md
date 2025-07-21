# CreateShopRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | Merchant id | 
**ParentDivisionId** | Pointer to **string** | Parent division ID. Required when shopParentType is DIVISION or EXTERNAL_DIVISION | [optional] 
**ShopParentType** | **string** | Parent type for the shop | 
**MainName** | **string** | Shop name | 
**ShopAddress** | Pointer to [**AddressInfo**](AddressInfo.md) |  | [optional] 
**ShopDesc** | Pointer to **string** | Shop description | [optional] 
**ExternalShopId** | **string** | External shop identifier | 
**LogoUrlMap** | Pointer to **map[string]string** | Logo images encoded in base64. Keys can be LOGO, PC_LOGO, MOBILE_LOGO. Images must be PNG format. | [optional] 
**MccCodes** | Pointer to **[]string** | MCC codes | [optional] 
**SizeType** | **string** | Size type of the shop | 
**Ln** | Pointer to **string** | Longitude | [optional] 
**Lat** | Pointer to **string** | Latitude | [optional] 
**TaxNo** | Pointer to **string** | Tax number (NPWP). Must be 15 digits | [optional] 
**BrandName** | Pointer to **string** | Legal name/tax name | [optional] 
**TaxAddress** | Pointer to [**AddressInfo**](AddressInfo.md) |  | [optional] 
**Loyalty** | Pointer to **string** | Flag for loyalty category | [optional] 
**ApiVersion** | Pointer to **string** | API version flag. Use &gt; 2 for new attributes | [optional] 
**BusinessEntity** | Pointer to **string** | Business entity type. Required if apiVersion &gt; 2 | [optional] 
**BusinessDocs** | Pointer to [**[]BusinessDocs**](BusinessDocs.md) | Business documents. Required if apiVersion &gt; 2 | [optional] 
**OwnerName** | Pointer to [**UserName**](UserName.md) |  | [optional] 
**OwnerIdType** | Pointer to **string** | Owner ID type. Required if apiVersion &gt; 2 | [optional] 
**OwnerIdNo** | Pointer to **string** | Owner ID number. Required if apiVersion &gt; 2 | [optional] 
**OwnerAddress** | Pointer to [**AddressInfo**](AddressInfo.md) |  | [optional] 
**OwnerPhoneNumber** | Pointer to [**MobileNoInfo**](MobileNoInfo.md) |  | [optional] 
**ShopOwning** | Pointer to **string** | Shop ownership type. Required if apiVersion &gt; 2 | [optional] 
**DeviceNumber** | Pointer to **string** | Device number. Required if apiVersion &gt; 2 | [optional] 
**PosNumber** | Pointer to **string** | POS number. Required if apiVersion &gt; 2 | [optional] 
**DirectorPics** | Pointer to [**[]PicInfo**](PicInfo.md) | Director PICs. Required if apiVersion &gt; 2 | [optional] 
**NonDirectorPics** | Pointer to [**[]PicInfo**](PicInfo.md) | Non-director PICs. Required if apiVersion &gt; 2 | [optional] 
**ShopBizType** | Pointer to **string** | Shop business type | [optional] 
**ExtInfo** | Pointer to **map[string]interface{}** | Extended information | [optional] 

## Methods

### NewCreateShopRequest

`func NewCreateShopRequest(merchantId string, shopParentType string, mainName string, externalShopId string, sizeType string, ) *CreateShopRequest`

NewCreateShopRequest instantiates a new CreateShopRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShopRequestWithDefaults

`func NewCreateShopRequestWithDefaults() *CreateShopRequest`

NewCreateShopRequestWithDefaults instantiates a new CreateShopRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *CreateShopRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CreateShopRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CreateShopRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetParentDivisionId

`func (o *CreateShopRequest) GetParentDivisionId() string`

GetParentDivisionId returns the ParentDivisionId field if non-nil, zero value otherwise.

### GetParentDivisionIdOk

`func (o *CreateShopRequest) GetParentDivisionIdOk() (*string, bool)`

GetParentDivisionIdOk returns a tuple with the ParentDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDivisionId

`func (o *CreateShopRequest) SetParentDivisionId(v string)`

SetParentDivisionId sets ParentDivisionId field to given value.

### HasParentDivisionId

`func (o *CreateShopRequest) HasParentDivisionId() bool`

HasParentDivisionId returns a boolean if a field has been set.

### GetShopParentType

`func (o *CreateShopRequest) GetShopParentType() string`

GetShopParentType returns the ShopParentType field if non-nil, zero value otherwise.

### GetShopParentTypeOk

`func (o *CreateShopRequest) GetShopParentTypeOk() (*string, bool)`

GetShopParentTypeOk returns a tuple with the ShopParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopParentType

`func (o *CreateShopRequest) SetShopParentType(v string)`

SetShopParentType sets ShopParentType field to given value.


### GetMainName

`func (o *CreateShopRequest) GetMainName() string`

GetMainName returns the MainName field if non-nil, zero value otherwise.

### GetMainNameOk

`func (o *CreateShopRequest) GetMainNameOk() (*string, bool)`

GetMainNameOk returns a tuple with the MainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainName

`func (o *CreateShopRequest) SetMainName(v string)`

SetMainName sets MainName field to given value.


### GetShopAddress

`func (o *CreateShopRequest) GetShopAddress() AddressInfo`

GetShopAddress returns the ShopAddress field if non-nil, zero value otherwise.

### GetShopAddressOk

`func (o *CreateShopRequest) GetShopAddressOk() (*AddressInfo, bool)`

GetShopAddressOk returns a tuple with the ShopAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopAddress

`func (o *CreateShopRequest) SetShopAddress(v AddressInfo)`

SetShopAddress sets ShopAddress field to given value.

### HasShopAddress

`func (o *CreateShopRequest) HasShopAddress() bool`

HasShopAddress returns a boolean if a field has been set.

### GetShopDesc

`func (o *CreateShopRequest) GetShopDesc() string`

GetShopDesc returns the ShopDesc field if non-nil, zero value otherwise.

### GetShopDescOk

`func (o *CreateShopRequest) GetShopDescOk() (*string, bool)`

GetShopDescOk returns a tuple with the ShopDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopDesc

`func (o *CreateShopRequest) SetShopDesc(v string)`

SetShopDesc sets ShopDesc field to given value.

### HasShopDesc

`func (o *CreateShopRequest) HasShopDesc() bool`

HasShopDesc returns a boolean if a field has been set.

### GetExternalShopId

`func (o *CreateShopRequest) GetExternalShopId() string`

GetExternalShopId returns the ExternalShopId field if non-nil, zero value otherwise.

### GetExternalShopIdOk

`func (o *CreateShopRequest) GetExternalShopIdOk() (*string, bool)`

GetExternalShopIdOk returns a tuple with the ExternalShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShopId

`func (o *CreateShopRequest) SetExternalShopId(v string)`

SetExternalShopId sets ExternalShopId field to given value.


### GetLogoUrlMap

`func (o *CreateShopRequest) GetLogoUrlMap() map[string]string`

GetLogoUrlMap returns the LogoUrlMap field if non-nil, zero value otherwise.

### GetLogoUrlMapOk

`func (o *CreateShopRequest) GetLogoUrlMapOk() (*map[string]string, bool)`

GetLogoUrlMapOk returns a tuple with the LogoUrlMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrlMap

`func (o *CreateShopRequest) SetLogoUrlMap(v map[string]string)`

SetLogoUrlMap sets LogoUrlMap field to given value.

### HasLogoUrlMap

`func (o *CreateShopRequest) HasLogoUrlMap() bool`

HasLogoUrlMap returns a boolean if a field has been set.

### GetMccCodes

`func (o *CreateShopRequest) GetMccCodes() []string`

GetMccCodes returns the MccCodes field if non-nil, zero value otherwise.

### GetMccCodesOk

`func (o *CreateShopRequest) GetMccCodesOk() (*[]string, bool)`

GetMccCodesOk returns a tuple with the MccCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccCodes

`func (o *CreateShopRequest) SetMccCodes(v []string)`

SetMccCodes sets MccCodes field to given value.

### HasMccCodes

`func (o *CreateShopRequest) HasMccCodes() bool`

HasMccCodes returns a boolean if a field has been set.

### GetSizeType

`func (o *CreateShopRequest) GetSizeType() string`

GetSizeType returns the SizeType field if non-nil, zero value otherwise.

### GetSizeTypeOk

`func (o *CreateShopRequest) GetSizeTypeOk() (*string, bool)`

GetSizeTypeOk returns a tuple with the SizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeType

`func (o *CreateShopRequest) SetSizeType(v string)`

SetSizeType sets SizeType field to given value.


### GetLn

`func (o *CreateShopRequest) GetLn() string`

GetLn returns the Ln field if non-nil, zero value otherwise.

### GetLnOk

`func (o *CreateShopRequest) GetLnOk() (*string, bool)`

GetLnOk returns a tuple with the Ln field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLn

`func (o *CreateShopRequest) SetLn(v string)`

SetLn sets Ln field to given value.

### HasLn

`func (o *CreateShopRequest) HasLn() bool`

HasLn returns a boolean if a field has been set.

### GetLat

`func (o *CreateShopRequest) GetLat() string`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *CreateShopRequest) GetLatOk() (*string, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *CreateShopRequest) SetLat(v string)`

SetLat sets Lat field to given value.

### HasLat

`func (o *CreateShopRequest) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetTaxNo

`func (o *CreateShopRequest) GetTaxNo() string`

GetTaxNo returns the TaxNo field if non-nil, zero value otherwise.

### GetTaxNoOk

`func (o *CreateShopRequest) GetTaxNoOk() (*string, bool)`

GetTaxNoOk returns a tuple with the TaxNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNo

`func (o *CreateShopRequest) SetTaxNo(v string)`

SetTaxNo sets TaxNo field to given value.

### HasTaxNo

`func (o *CreateShopRequest) HasTaxNo() bool`

HasTaxNo returns a boolean if a field has been set.

### GetBrandName

`func (o *CreateShopRequest) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *CreateShopRequest) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *CreateShopRequest) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.

### HasBrandName

`func (o *CreateShopRequest) HasBrandName() bool`

HasBrandName returns a boolean if a field has been set.

### GetTaxAddress

`func (o *CreateShopRequest) GetTaxAddress() AddressInfo`

GetTaxAddress returns the TaxAddress field if non-nil, zero value otherwise.

### GetTaxAddressOk

`func (o *CreateShopRequest) GetTaxAddressOk() (*AddressInfo, bool)`

GetTaxAddressOk returns a tuple with the TaxAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAddress

`func (o *CreateShopRequest) SetTaxAddress(v AddressInfo)`

SetTaxAddress sets TaxAddress field to given value.

### HasTaxAddress

`func (o *CreateShopRequest) HasTaxAddress() bool`

HasTaxAddress returns a boolean if a field has been set.

### GetLoyalty

`func (o *CreateShopRequest) GetLoyalty() string`

GetLoyalty returns the Loyalty field if non-nil, zero value otherwise.

### GetLoyaltyOk

`func (o *CreateShopRequest) GetLoyaltyOk() (*string, bool)`

GetLoyaltyOk returns a tuple with the Loyalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyalty

`func (o *CreateShopRequest) SetLoyalty(v string)`

SetLoyalty sets Loyalty field to given value.

### HasLoyalty

`func (o *CreateShopRequest) HasLoyalty() bool`

HasLoyalty returns a boolean if a field has been set.

### GetApiVersion

`func (o *CreateShopRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateShopRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateShopRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateShopRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetBusinessEntity

`func (o *CreateShopRequest) GetBusinessEntity() string`

GetBusinessEntity returns the BusinessEntity field if non-nil, zero value otherwise.

### GetBusinessEntityOk

`func (o *CreateShopRequest) GetBusinessEntityOk() (*string, bool)`

GetBusinessEntityOk returns a tuple with the BusinessEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessEntity

`func (o *CreateShopRequest) SetBusinessEntity(v string)`

SetBusinessEntity sets BusinessEntity field to given value.

### HasBusinessEntity

`func (o *CreateShopRequest) HasBusinessEntity() bool`

HasBusinessEntity returns a boolean if a field has been set.

### GetBusinessDocs

`func (o *CreateShopRequest) GetBusinessDocs() []BusinessDocs`

GetBusinessDocs returns the BusinessDocs field if non-nil, zero value otherwise.

### GetBusinessDocsOk

`func (o *CreateShopRequest) GetBusinessDocsOk() (*[]BusinessDocs, bool)`

GetBusinessDocsOk returns a tuple with the BusinessDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDocs

`func (o *CreateShopRequest) SetBusinessDocs(v []BusinessDocs)`

SetBusinessDocs sets BusinessDocs field to given value.

### HasBusinessDocs

`func (o *CreateShopRequest) HasBusinessDocs() bool`

HasBusinessDocs returns a boolean if a field has been set.

### GetOwnerName

`func (o *CreateShopRequest) GetOwnerName() UserName`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *CreateShopRequest) GetOwnerNameOk() (*UserName, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *CreateShopRequest) SetOwnerName(v UserName)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *CreateShopRequest) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetOwnerIdType

`func (o *CreateShopRequest) GetOwnerIdType() string`

GetOwnerIdType returns the OwnerIdType field if non-nil, zero value otherwise.

### GetOwnerIdTypeOk

`func (o *CreateShopRequest) GetOwnerIdTypeOk() (*string, bool)`

GetOwnerIdTypeOk returns a tuple with the OwnerIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdType

`func (o *CreateShopRequest) SetOwnerIdType(v string)`

SetOwnerIdType sets OwnerIdType field to given value.

### HasOwnerIdType

`func (o *CreateShopRequest) HasOwnerIdType() bool`

HasOwnerIdType returns a boolean if a field has been set.

### GetOwnerIdNo

`func (o *CreateShopRequest) GetOwnerIdNo() string`

GetOwnerIdNo returns the OwnerIdNo field if non-nil, zero value otherwise.

### GetOwnerIdNoOk

`func (o *CreateShopRequest) GetOwnerIdNoOk() (*string, bool)`

GetOwnerIdNoOk returns a tuple with the OwnerIdNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdNo

`func (o *CreateShopRequest) SetOwnerIdNo(v string)`

SetOwnerIdNo sets OwnerIdNo field to given value.

### HasOwnerIdNo

`func (o *CreateShopRequest) HasOwnerIdNo() bool`

HasOwnerIdNo returns a boolean if a field has been set.

### GetOwnerAddress

`func (o *CreateShopRequest) GetOwnerAddress() AddressInfo`

GetOwnerAddress returns the OwnerAddress field if non-nil, zero value otherwise.

### GetOwnerAddressOk

`func (o *CreateShopRequest) GetOwnerAddressOk() (*AddressInfo, bool)`

GetOwnerAddressOk returns a tuple with the OwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAddress

`func (o *CreateShopRequest) SetOwnerAddress(v AddressInfo)`

SetOwnerAddress sets OwnerAddress field to given value.

### HasOwnerAddress

`func (o *CreateShopRequest) HasOwnerAddress() bool`

HasOwnerAddress returns a boolean if a field has been set.

### GetOwnerPhoneNumber

`func (o *CreateShopRequest) GetOwnerPhoneNumber() MobileNoInfo`

GetOwnerPhoneNumber returns the OwnerPhoneNumber field if non-nil, zero value otherwise.

### GetOwnerPhoneNumberOk

`func (o *CreateShopRequest) GetOwnerPhoneNumberOk() (*MobileNoInfo, bool)`

GetOwnerPhoneNumberOk returns a tuple with the OwnerPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPhoneNumber

`func (o *CreateShopRequest) SetOwnerPhoneNumber(v MobileNoInfo)`

SetOwnerPhoneNumber sets OwnerPhoneNumber field to given value.

### HasOwnerPhoneNumber

`func (o *CreateShopRequest) HasOwnerPhoneNumber() bool`

HasOwnerPhoneNumber returns a boolean if a field has been set.

### GetShopOwning

`func (o *CreateShopRequest) GetShopOwning() string`

GetShopOwning returns the ShopOwning field if non-nil, zero value otherwise.

### GetShopOwningOk

`func (o *CreateShopRequest) GetShopOwningOk() (*string, bool)`

GetShopOwningOk returns a tuple with the ShopOwning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopOwning

`func (o *CreateShopRequest) SetShopOwning(v string)`

SetShopOwning sets ShopOwning field to given value.

### HasShopOwning

`func (o *CreateShopRequest) HasShopOwning() bool`

HasShopOwning returns a boolean if a field has been set.

### GetDeviceNumber

`func (o *CreateShopRequest) GetDeviceNumber() string`

GetDeviceNumber returns the DeviceNumber field if non-nil, zero value otherwise.

### GetDeviceNumberOk

`func (o *CreateShopRequest) GetDeviceNumberOk() (*string, bool)`

GetDeviceNumberOk returns a tuple with the DeviceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceNumber

`func (o *CreateShopRequest) SetDeviceNumber(v string)`

SetDeviceNumber sets DeviceNumber field to given value.

### HasDeviceNumber

`func (o *CreateShopRequest) HasDeviceNumber() bool`

HasDeviceNumber returns a boolean if a field has been set.

### GetPosNumber

`func (o *CreateShopRequest) GetPosNumber() string`

GetPosNumber returns the PosNumber field if non-nil, zero value otherwise.

### GetPosNumberOk

`func (o *CreateShopRequest) GetPosNumberOk() (*string, bool)`

GetPosNumberOk returns a tuple with the PosNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosNumber

`func (o *CreateShopRequest) SetPosNumber(v string)`

SetPosNumber sets PosNumber field to given value.

### HasPosNumber

`func (o *CreateShopRequest) HasPosNumber() bool`

HasPosNumber returns a boolean if a field has been set.

### GetDirectorPics

`func (o *CreateShopRequest) GetDirectorPics() []PicInfo`

GetDirectorPics returns the DirectorPics field if non-nil, zero value otherwise.

### GetDirectorPicsOk

`func (o *CreateShopRequest) GetDirectorPicsOk() (*[]PicInfo, bool)`

GetDirectorPicsOk returns a tuple with the DirectorPics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectorPics

`func (o *CreateShopRequest) SetDirectorPics(v []PicInfo)`

SetDirectorPics sets DirectorPics field to given value.

### HasDirectorPics

`func (o *CreateShopRequest) HasDirectorPics() bool`

HasDirectorPics returns a boolean if a field has been set.

### GetNonDirectorPics

`func (o *CreateShopRequest) GetNonDirectorPics() []PicInfo`

GetNonDirectorPics returns the NonDirectorPics field if non-nil, zero value otherwise.

### GetNonDirectorPicsOk

`func (o *CreateShopRequest) GetNonDirectorPicsOk() (*[]PicInfo, bool)`

GetNonDirectorPicsOk returns a tuple with the NonDirectorPics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDirectorPics

`func (o *CreateShopRequest) SetNonDirectorPics(v []PicInfo)`

SetNonDirectorPics sets NonDirectorPics field to given value.

### HasNonDirectorPics

`func (o *CreateShopRequest) HasNonDirectorPics() bool`

HasNonDirectorPics returns a boolean if a field has been set.

### GetShopBizType

`func (o *CreateShopRequest) GetShopBizType() string`

GetShopBizType returns the ShopBizType field if non-nil, zero value otherwise.

### GetShopBizTypeOk

`func (o *CreateShopRequest) GetShopBizTypeOk() (*string, bool)`

GetShopBizTypeOk returns a tuple with the ShopBizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopBizType

`func (o *CreateShopRequest) SetShopBizType(v string)`

SetShopBizType sets ShopBizType field to given value.

### HasShopBizType

`func (o *CreateShopRequest) HasShopBizType() bool`

HasShopBizType returns a boolean if a field has been set.

### GetExtInfo

`func (o *CreateShopRequest) GetExtInfo() map[string]interface{}`

GetExtInfo returns the ExtInfo field if non-nil, zero value otherwise.

### GetExtInfoOk

`func (o *CreateShopRequest) GetExtInfoOk() (*map[string]interface{}, bool)`

GetExtInfoOk returns a tuple with the ExtInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtInfo

`func (o *CreateShopRequest) SetExtInfo(v map[string]interface{})`

SetExtInfo sets ExtInfo field to given value.

### HasExtInfo

`func (o *CreateShopRequest) HasExtInfo() bool`

HasExtInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


