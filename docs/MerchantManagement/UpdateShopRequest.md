# UpdateShopRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopId** | **string** | Shop identifier. Length depends on shopIdType - INNER_ID (21 max), EXTERNAL_ID (64 max) | 
**MerchantId** | **string** | Merchant identifier | 
**ShopIdType** | **string** | Shop identifier type | 
**MainName** | Pointer to **string** | Shop name | [optional] 
**ShopAddress** | [**AddressInfo**](AddressInfo.md) |  | 
**ShopDesc** | Pointer to **string** | Shop description | [optional] 
**NewExternalShopId** | Pointer to **string** | New external shop identifier | [optional] 
**MccCodes** | Pointer to **[]string** | Merchant category code | [optional] 
**LogoUrlMap** | Pointer to **map[string]string** | Logo URL map with base64 encoded images. Keys can be LOGO, PC_LOGO, MOBILE_LOGO | [optional] 
**ExtInfo** | Pointer to **map[string]interface{}** | Extend information | [optional] 
**SizeType** | Pointer to **string** | Size type | [optional] 
**Ln** | Pointer to **string** | Longitude of shop&#39;s location | [optional] 
**Lat** | Pointer to **string** | Latitude of shop&#39;s location | [optional] 
**Loyalty** | Pointer to **string** | Flag for loyalty category | [optional] 
**OwnerAddress** | Pointer to [**AddressInfo**](AddressInfo.md) |  | [optional] 
**OwnerName** | Pointer to [**UserName**](UserName.md) |  | [optional] 
**OwnerPhoneNumber** | Pointer to [**MobileNoInfo**](MobileNoInfo.md) |  | [optional] 
**OwnerIdType** | Pointer to **string** | Owner identifier type | [optional] 
**OwnerIdNo** | Pointer to **string** | Owner identifier number. Length depends on ownerIdType - KTP (16), SIM (12-14), Passport (8), NIB (&gt;&#x3D;13), SIUP (free text) | [optional] 
**DeviceNumber** | Pointer to **string** | Device number | [optional] 
**PosNumber** | Pointer to **string** | POS number | [optional] 
**ApiVersion** | Pointer to **string** | API version flag. Use &gt; 2 for new attributes | [optional] 
**BusinessEntity** | Pointer to **string** | Business entity type | [optional] 
**ShopOwning** | Pointer to **string** | Shop owning information | [optional] 
**ShopBizType** | Pointer to **string** | Shop business type | [optional] 
**BusinessDocs** | Pointer to [**[]BusinessDocs**](BusinessDocs.md) | Business documents. \&quot;individu\&quot; entity can only use KTP and SIM. Other entities can use SIUP and NIB | [optional] 
**BusinessEndDate** | Pointer to **string** | Business end date, in format YYYY-MM-dd | [optional] 
**TaxNo** | Pointer to **string** | Tax number (NPWP). Must be 15 digits | [optional] 
**TaxAddress** | Pointer to [**AddressInfo**](AddressInfo.md) |  | [optional] 
**BrandName** | Pointer to **string** | Brand name on legal name or tax name | [optional] 
**DirectorPics** | Pointer to [**[]PicInfo**](PicInfo.md) | Director as a PIC of shop | [optional] 
**NonDirectorPics** | Pointer to [**[]PicInfo**](PicInfo.md) | Non director which become an PIC of shop | [optional] 

## Methods

### NewUpdateShopRequest

`func NewUpdateShopRequest(shopId string, merchantId string, shopIdType string, shopAddress AddressInfo, ) *UpdateShopRequest`

NewUpdateShopRequest instantiates a new UpdateShopRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShopRequestWithDefaults

`func NewUpdateShopRequestWithDefaults() *UpdateShopRequest`

NewUpdateShopRequestWithDefaults instantiates a new UpdateShopRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopId

`func (o *UpdateShopRequest) GetShopId() string`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *UpdateShopRequest) GetShopIdOk() (*string, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *UpdateShopRequest) SetShopId(v string)`

SetShopId sets ShopId field to given value.


### GetMerchantId

`func (o *UpdateShopRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UpdateShopRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UpdateShopRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetShopIdType

`func (o *UpdateShopRequest) GetShopIdType() string`

GetShopIdType returns the ShopIdType field if non-nil, zero value otherwise.

### GetShopIdTypeOk

`func (o *UpdateShopRequest) GetShopIdTypeOk() (*string, bool)`

GetShopIdTypeOk returns a tuple with the ShopIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopIdType

`func (o *UpdateShopRequest) SetShopIdType(v string)`

SetShopIdType sets ShopIdType field to given value.


### GetMainName

`func (o *UpdateShopRequest) GetMainName() string`

GetMainName returns the MainName field if non-nil, zero value otherwise.

### GetMainNameOk

`func (o *UpdateShopRequest) GetMainNameOk() (*string, bool)`

GetMainNameOk returns a tuple with the MainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainName

`func (o *UpdateShopRequest) SetMainName(v string)`

SetMainName sets MainName field to given value.

### HasMainName

`func (o *UpdateShopRequest) HasMainName() bool`

HasMainName returns a boolean if a field has been set.

### GetShopAddress

`func (o *UpdateShopRequest) GetShopAddress() AddressInfo`

GetShopAddress returns the ShopAddress field if non-nil, zero value otherwise.

### GetShopAddressOk

`func (o *UpdateShopRequest) GetShopAddressOk() (*AddressInfo, bool)`

GetShopAddressOk returns a tuple with the ShopAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopAddress

`func (o *UpdateShopRequest) SetShopAddress(v AddressInfo)`

SetShopAddress sets ShopAddress field to given value.


### GetShopDesc

`func (o *UpdateShopRequest) GetShopDesc() string`

GetShopDesc returns the ShopDesc field if non-nil, zero value otherwise.

### GetShopDescOk

`func (o *UpdateShopRequest) GetShopDescOk() (*string, bool)`

GetShopDescOk returns a tuple with the ShopDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopDesc

`func (o *UpdateShopRequest) SetShopDesc(v string)`

SetShopDesc sets ShopDesc field to given value.

### HasShopDesc

`func (o *UpdateShopRequest) HasShopDesc() bool`

HasShopDesc returns a boolean if a field has been set.

### GetNewExternalShopId

`func (o *UpdateShopRequest) GetNewExternalShopId() string`

GetNewExternalShopId returns the NewExternalShopId field if non-nil, zero value otherwise.

### GetNewExternalShopIdOk

`func (o *UpdateShopRequest) GetNewExternalShopIdOk() (*string, bool)`

GetNewExternalShopIdOk returns a tuple with the NewExternalShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewExternalShopId

`func (o *UpdateShopRequest) SetNewExternalShopId(v string)`

SetNewExternalShopId sets NewExternalShopId field to given value.

### HasNewExternalShopId

`func (o *UpdateShopRequest) HasNewExternalShopId() bool`

HasNewExternalShopId returns a boolean if a field has been set.

### GetMccCodes

`func (o *UpdateShopRequest) GetMccCodes() []string`

GetMccCodes returns the MccCodes field if non-nil, zero value otherwise.

### GetMccCodesOk

`func (o *UpdateShopRequest) GetMccCodesOk() (*[]string, bool)`

GetMccCodesOk returns a tuple with the MccCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccCodes

`func (o *UpdateShopRequest) SetMccCodes(v []string)`

SetMccCodes sets MccCodes field to given value.

### HasMccCodes

`func (o *UpdateShopRequest) HasMccCodes() bool`

HasMccCodes returns a boolean if a field has been set.

### GetLogoUrlMap

`func (o *UpdateShopRequest) GetLogoUrlMap() map[string]string`

GetLogoUrlMap returns the LogoUrlMap field if non-nil, zero value otherwise.

### GetLogoUrlMapOk

`func (o *UpdateShopRequest) GetLogoUrlMapOk() (*map[string]string, bool)`

GetLogoUrlMapOk returns a tuple with the LogoUrlMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrlMap

`func (o *UpdateShopRequest) SetLogoUrlMap(v map[string]string)`

SetLogoUrlMap sets LogoUrlMap field to given value.

### HasLogoUrlMap

`func (o *UpdateShopRequest) HasLogoUrlMap() bool`

HasLogoUrlMap returns a boolean if a field has been set.

### GetExtInfo

`func (o *UpdateShopRequest) GetExtInfo() map[string]interface{}`

GetExtInfo returns the ExtInfo field if non-nil, zero value otherwise.

### GetExtInfoOk

`func (o *UpdateShopRequest) GetExtInfoOk() (*map[string]interface{}, bool)`

GetExtInfoOk returns a tuple with the ExtInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtInfo

`func (o *UpdateShopRequest) SetExtInfo(v map[string]interface{})`

SetExtInfo sets ExtInfo field to given value.

### HasExtInfo

`func (o *UpdateShopRequest) HasExtInfo() bool`

HasExtInfo returns a boolean if a field has been set.

### GetSizeType

`func (o *UpdateShopRequest) GetSizeType() string`

GetSizeType returns the SizeType field if non-nil, zero value otherwise.

### GetSizeTypeOk

`func (o *UpdateShopRequest) GetSizeTypeOk() (*string, bool)`

GetSizeTypeOk returns a tuple with the SizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeType

`func (o *UpdateShopRequest) SetSizeType(v string)`

SetSizeType sets SizeType field to given value.

### HasSizeType

`func (o *UpdateShopRequest) HasSizeType() bool`

HasSizeType returns a boolean if a field has been set.

### GetLn

`func (o *UpdateShopRequest) GetLn() string`

GetLn returns the Ln field if non-nil, zero value otherwise.

### GetLnOk

`func (o *UpdateShopRequest) GetLnOk() (*string, bool)`

GetLnOk returns a tuple with the Ln field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLn

`func (o *UpdateShopRequest) SetLn(v string)`

SetLn sets Ln field to given value.

### HasLn

`func (o *UpdateShopRequest) HasLn() bool`

HasLn returns a boolean if a field has been set.

### GetLat

`func (o *UpdateShopRequest) GetLat() string`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *UpdateShopRequest) GetLatOk() (*string, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *UpdateShopRequest) SetLat(v string)`

SetLat sets Lat field to given value.

### HasLat

`func (o *UpdateShopRequest) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLoyalty

`func (o *UpdateShopRequest) GetLoyalty() string`

GetLoyalty returns the Loyalty field if non-nil, zero value otherwise.

### GetLoyaltyOk

`func (o *UpdateShopRequest) GetLoyaltyOk() (*string, bool)`

GetLoyaltyOk returns a tuple with the Loyalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyalty

`func (o *UpdateShopRequest) SetLoyalty(v string)`

SetLoyalty sets Loyalty field to given value.

### HasLoyalty

`func (o *UpdateShopRequest) HasLoyalty() bool`

HasLoyalty returns a boolean if a field has been set.

### GetOwnerAddress

`func (o *UpdateShopRequest) GetOwnerAddress() AddressInfo`

GetOwnerAddress returns the OwnerAddress field if non-nil, zero value otherwise.

### GetOwnerAddressOk

`func (o *UpdateShopRequest) GetOwnerAddressOk() (*AddressInfo, bool)`

GetOwnerAddressOk returns a tuple with the OwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAddress

`func (o *UpdateShopRequest) SetOwnerAddress(v AddressInfo)`

SetOwnerAddress sets OwnerAddress field to given value.

### HasOwnerAddress

`func (o *UpdateShopRequest) HasOwnerAddress() bool`

HasOwnerAddress returns a boolean if a field has been set.

### GetOwnerName

`func (o *UpdateShopRequest) GetOwnerName() UserName`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *UpdateShopRequest) GetOwnerNameOk() (*UserName, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *UpdateShopRequest) SetOwnerName(v UserName)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *UpdateShopRequest) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetOwnerPhoneNumber

`func (o *UpdateShopRequest) GetOwnerPhoneNumber() MobileNoInfo`

GetOwnerPhoneNumber returns the OwnerPhoneNumber field if non-nil, zero value otherwise.

### GetOwnerPhoneNumberOk

`func (o *UpdateShopRequest) GetOwnerPhoneNumberOk() (*MobileNoInfo, bool)`

GetOwnerPhoneNumberOk returns a tuple with the OwnerPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPhoneNumber

`func (o *UpdateShopRequest) SetOwnerPhoneNumber(v MobileNoInfo)`

SetOwnerPhoneNumber sets OwnerPhoneNumber field to given value.

### HasOwnerPhoneNumber

`func (o *UpdateShopRequest) HasOwnerPhoneNumber() bool`

HasOwnerPhoneNumber returns a boolean if a field has been set.

### GetOwnerIdType

`func (o *UpdateShopRequest) GetOwnerIdType() string`

GetOwnerIdType returns the OwnerIdType field if non-nil, zero value otherwise.

### GetOwnerIdTypeOk

`func (o *UpdateShopRequest) GetOwnerIdTypeOk() (*string, bool)`

GetOwnerIdTypeOk returns a tuple with the OwnerIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdType

`func (o *UpdateShopRequest) SetOwnerIdType(v string)`

SetOwnerIdType sets OwnerIdType field to given value.

### HasOwnerIdType

`func (o *UpdateShopRequest) HasOwnerIdType() bool`

HasOwnerIdType returns a boolean if a field has been set.

### GetOwnerIdNo

`func (o *UpdateShopRequest) GetOwnerIdNo() string`

GetOwnerIdNo returns the OwnerIdNo field if non-nil, zero value otherwise.

### GetOwnerIdNoOk

`func (o *UpdateShopRequest) GetOwnerIdNoOk() (*string, bool)`

GetOwnerIdNoOk returns a tuple with the OwnerIdNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdNo

`func (o *UpdateShopRequest) SetOwnerIdNo(v string)`

SetOwnerIdNo sets OwnerIdNo field to given value.

### HasOwnerIdNo

`func (o *UpdateShopRequest) HasOwnerIdNo() bool`

HasOwnerIdNo returns a boolean if a field has been set.

### GetDeviceNumber

`func (o *UpdateShopRequest) GetDeviceNumber() string`

GetDeviceNumber returns the DeviceNumber field if non-nil, zero value otherwise.

### GetDeviceNumberOk

`func (o *UpdateShopRequest) GetDeviceNumberOk() (*string, bool)`

GetDeviceNumberOk returns a tuple with the DeviceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceNumber

`func (o *UpdateShopRequest) SetDeviceNumber(v string)`

SetDeviceNumber sets DeviceNumber field to given value.

### HasDeviceNumber

`func (o *UpdateShopRequest) HasDeviceNumber() bool`

HasDeviceNumber returns a boolean if a field has been set.

### GetPosNumber

`func (o *UpdateShopRequest) GetPosNumber() string`

GetPosNumber returns the PosNumber field if non-nil, zero value otherwise.

### GetPosNumberOk

`func (o *UpdateShopRequest) GetPosNumberOk() (*string, bool)`

GetPosNumberOk returns a tuple with the PosNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosNumber

`func (o *UpdateShopRequest) SetPosNumber(v string)`

SetPosNumber sets PosNumber field to given value.

### HasPosNumber

`func (o *UpdateShopRequest) HasPosNumber() bool`

HasPosNumber returns a boolean if a field has been set.

### GetApiVersion

`func (o *UpdateShopRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UpdateShopRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UpdateShopRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *UpdateShopRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetBusinessEntity

`func (o *UpdateShopRequest) GetBusinessEntity() string`

GetBusinessEntity returns the BusinessEntity field if non-nil, zero value otherwise.

### GetBusinessEntityOk

`func (o *UpdateShopRequest) GetBusinessEntityOk() (*string, bool)`

GetBusinessEntityOk returns a tuple with the BusinessEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessEntity

`func (o *UpdateShopRequest) SetBusinessEntity(v string)`

SetBusinessEntity sets BusinessEntity field to given value.

### HasBusinessEntity

`func (o *UpdateShopRequest) HasBusinessEntity() bool`

HasBusinessEntity returns a boolean if a field has been set.

### GetShopOwning

`func (o *UpdateShopRequest) GetShopOwning() string`

GetShopOwning returns the ShopOwning field if non-nil, zero value otherwise.

### GetShopOwningOk

`func (o *UpdateShopRequest) GetShopOwningOk() (*string, bool)`

GetShopOwningOk returns a tuple with the ShopOwning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopOwning

`func (o *UpdateShopRequest) SetShopOwning(v string)`

SetShopOwning sets ShopOwning field to given value.

### HasShopOwning

`func (o *UpdateShopRequest) HasShopOwning() bool`

HasShopOwning returns a boolean if a field has been set.

### GetShopBizType

`func (o *UpdateShopRequest) GetShopBizType() string`

GetShopBizType returns the ShopBizType field if non-nil, zero value otherwise.

### GetShopBizTypeOk

`func (o *UpdateShopRequest) GetShopBizTypeOk() (*string, bool)`

GetShopBizTypeOk returns a tuple with the ShopBizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopBizType

`func (o *UpdateShopRequest) SetShopBizType(v string)`

SetShopBizType sets ShopBizType field to given value.

### HasShopBizType

`func (o *UpdateShopRequest) HasShopBizType() bool`

HasShopBizType returns a boolean if a field has been set.

### GetBusinessDocs

`func (o *UpdateShopRequest) GetBusinessDocs() []BusinessDocs`

GetBusinessDocs returns the BusinessDocs field if non-nil, zero value otherwise.

### GetBusinessDocsOk

`func (o *UpdateShopRequest) GetBusinessDocsOk() (*[]BusinessDocs, bool)`

GetBusinessDocsOk returns a tuple with the BusinessDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDocs

`func (o *UpdateShopRequest) SetBusinessDocs(v []BusinessDocs)`

SetBusinessDocs sets BusinessDocs field to given value.

### HasBusinessDocs

`func (o *UpdateShopRequest) HasBusinessDocs() bool`

HasBusinessDocs returns a boolean if a field has been set.

### GetBusinessEndDate

`func (o *UpdateShopRequest) GetBusinessEndDate() string`

GetBusinessEndDate returns the BusinessEndDate field if non-nil, zero value otherwise.

### GetBusinessEndDateOk

`func (o *UpdateShopRequest) GetBusinessEndDateOk() (*string, bool)`

GetBusinessEndDateOk returns a tuple with the BusinessEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessEndDate

`func (o *UpdateShopRequest) SetBusinessEndDate(v string)`

SetBusinessEndDate sets BusinessEndDate field to given value.

### HasBusinessEndDate

`func (o *UpdateShopRequest) HasBusinessEndDate() bool`

HasBusinessEndDate returns a boolean if a field has been set.

### GetTaxNo

`func (o *UpdateShopRequest) GetTaxNo() string`

GetTaxNo returns the TaxNo field if non-nil, zero value otherwise.

### GetTaxNoOk

`func (o *UpdateShopRequest) GetTaxNoOk() (*string, bool)`

GetTaxNoOk returns a tuple with the TaxNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNo

`func (o *UpdateShopRequest) SetTaxNo(v string)`

SetTaxNo sets TaxNo field to given value.

### HasTaxNo

`func (o *UpdateShopRequest) HasTaxNo() bool`

HasTaxNo returns a boolean if a field has been set.

### GetTaxAddress

`func (o *UpdateShopRequest) GetTaxAddress() AddressInfo`

GetTaxAddress returns the TaxAddress field if non-nil, zero value otherwise.

### GetTaxAddressOk

`func (o *UpdateShopRequest) GetTaxAddressOk() (*AddressInfo, bool)`

GetTaxAddressOk returns a tuple with the TaxAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAddress

`func (o *UpdateShopRequest) SetTaxAddress(v AddressInfo)`

SetTaxAddress sets TaxAddress field to given value.

### HasTaxAddress

`func (o *UpdateShopRequest) HasTaxAddress() bool`

HasTaxAddress returns a boolean if a field has been set.

### GetBrandName

`func (o *UpdateShopRequest) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *UpdateShopRequest) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *UpdateShopRequest) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.

### HasBrandName

`func (o *UpdateShopRequest) HasBrandName() bool`

HasBrandName returns a boolean if a field has been set.

### GetDirectorPics

`func (o *UpdateShopRequest) GetDirectorPics() []PicInfo`

GetDirectorPics returns the DirectorPics field if non-nil, zero value otherwise.

### GetDirectorPicsOk

`func (o *UpdateShopRequest) GetDirectorPicsOk() (*[]PicInfo, bool)`

GetDirectorPicsOk returns a tuple with the DirectorPics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectorPics

`func (o *UpdateShopRequest) SetDirectorPics(v []PicInfo)`

SetDirectorPics sets DirectorPics field to given value.

### HasDirectorPics

`func (o *UpdateShopRequest) HasDirectorPics() bool`

HasDirectorPics returns a boolean if a field has been set.

### GetNonDirectorPics

`func (o *UpdateShopRequest) GetNonDirectorPics() []PicInfo`

GetNonDirectorPics returns the NonDirectorPics field if non-nil, zero value otherwise.

### GetNonDirectorPicsOk

`func (o *UpdateShopRequest) GetNonDirectorPicsOk() (*[]PicInfo, bool)`

GetNonDirectorPicsOk returns a tuple with the NonDirectorPics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDirectorPics

`func (o *UpdateShopRequest) SetNonDirectorPics(v []PicInfo)`

SetNonDirectorPics sets NonDirectorPics field to given value.

### HasNonDirectorPics

`func (o *UpdateShopRequest) HasNonDirectorPics() bool`

HasNonDirectorPics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


