# UpdateDivisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | Merchant identifier | 
**DivisionId** | Pointer to **string** | Division identifier. Required when divisionIdType is INNER_ID | [optional] 
**DivisionName** | **string** | Division name | 
**DivisionAddress** | [**AddressInfo**](AddressInfo.md) |  | 
**DivisionDescription** | Pointer to **string** | Division description | [optional] 
**DivisionType** | **string** | Division type | 
**DivisionIdType** | **string** | Division identifier type | 
**ExternalDivisionId** | Pointer to **string** | External division identifier. Required when divisionIdType is EXTERNAL_ID | [optional] 
**NewExternalDivisionId** | **string** | New external division identifier | 
**LogoUrlMap** | Pointer to **map[string]string** | Logo URL map with base64 encoded images. Keys can be LOGO, PC_LOGO, MOBILE_LOGO | [optional] 
**MccCodes** | **[]string** | Merchant category codes | 
**ExtInfo** | **map[string]interface{}** | Extended information | 
**ApiVersion** | Pointer to **string** | API version flag. Use &gt; 2 for new attributes | [optional] 
**BusinessDocs** | Pointer to [**[]BusinessDocs**](BusinessDocs.md) | Business documents. Required when apiVersion &gt; 2. \&quot;individu\&quot; entity can only use KTP and SIM. Other entities can use SIUP and NIB | [optional] 
**BusinessEntity** | Pointer to **string** | Business entity type. Required when apiVersion &gt; 2 | [optional] 
**BusinessEndDate** | Pointer to **string** | Business end date, in format YYYY-MM-DD. Required when apiVersion &gt; 2 | [optional] 
**OwnerName** | Pointer to [**UserName**](UserName.md) |  | [optional] 
**OwnerPhoneNumber** | Pointer to [**MobileNoInfo**](MobileNoInfo.md) |  | [optional] 
**OwnerIdType** | Pointer to **string** | Owner identifier type. Required when apiVersion &gt; 2 | [optional] 
**OwnerIdNo** | Pointer to **string** | Owner identifier number. Required when apiVersion &gt; 2. Length depends on ownerIdType - KTP (16), SIM (12-14), Passport (8), NIB (&gt;&#x3D;13), SIUP (free text) | [optional] 
**OwnerAddress** | Pointer to [**AddressInfo**](AddressInfo.md) |  | [optional] 
**DirectorPics** | Pointer to [**[]PicInfo**](PicInfo.md) | Director as a PIC of sub merchant. Required when apiVersion &gt; 2 | [optional] 
**NonDirectorPics** | Pointer to [**[]PicInfo**](PicInfo.md) | Non director which become a PIC of sub merchant. Required when apiVersion &gt; 2 | [optional] 
**SizeType** | Pointer to **string** | Size type. Required when apiVersion &gt; 2 | [optional] 
**PgDivisionFlag** | Pointer to **string** | Flag if division is type PG | [optional] 

## Methods

### NewUpdateDivisionRequest

`func NewUpdateDivisionRequest(merchantId string, divisionName string, divisionAddress AddressInfo, divisionType string, divisionIdType string, newExternalDivisionId string, mccCodes []string, extInfo map[string]interface{}, ) *UpdateDivisionRequest`

NewUpdateDivisionRequest instantiates a new UpdateDivisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDivisionRequestWithDefaults

`func NewUpdateDivisionRequestWithDefaults() *UpdateDivisionRequest`

NewUpdateDivisionRequestWithDefaults instantiates a new UpdateDivisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *UpdateDivisionRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UpdateDivisionRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UpdateDivisionRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetDivisionId

`func (o *UpdateDivisionRequest) GetDivisionId() string`

GetDivisionId returns the DivisionId field if non-nil, zero value otherwise.

### GetDivisionIdOk

`func (o *UpdateDivisionRequest) GetDivisionIdOk() (*string, bool)`

GetDivisionIdOk returns a tuple with the DivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionId

`func (o *UpdateDivisionRequest) SetDivisionId(v string)`

SetDivisionId sets DivisionId field to given value.

### HasDivisionId

`func (o *UpdateDivisionRequest) HasDivisionId() bool`

HasDivisionId returns a boolean if a field has been set.

### GetDivisionName

`func (o *UpdateDivisionRequest) GetDivisionName() string`

GetDivisionName returns the DivisionName field if non-nil, zero value otherwise.

### GetDivisionNameOk

`func (o *UpdateDivisionRequest) GetDivisionNameOk() (*string, bool)`

GetDivisionNameOk returns a tuple with the DivisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionName

`func (o *UpdateDivisionRequest) SetDivisionName(v string)`

SetDivisionName sets DivisionName field to given value.


### GetDivisionAddress

`func (o *UpdateDivisionRequest) GetDivisionAddress() AddressInfo`

GetDivisionAddress returns the DivisionAddress field if non-nil, zero value otherwise.

### GetDivisionAddressOk

`func (o *UpdateDivisionRequest) GetDivisionAddressOk() (*AddressInfo, bool)`

GetDivisionAddressOk returns a tuple with the DivisionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionAddress

`func (o *UpdateDivisionRequest) SetDivisionAddress(v AddressInfo)`

SetDivisionAddress sets DivisionAddress field to given value.


### GetDivisionDescription

`func (o *UpdateDivisionRequest) GetDivisionDescription() string`

GetDivisionDescription returns the DivisionDescription field if non-nil, zero value otherwise.

### GetDivisionDescriptionOk

`func (o *UpdateDivisionRequest) GetDivisionDescriptionOk() (*string, bool)`

GetDivisionDescriptionOk returns a tuple with the DivisionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionDescription

`func (o *UpdateDivisionRequest) SetDivisionDescription(v string)`

SetDivisionDescription sets DivisionDescription field to given value.

### HasDivisionDescription

`func (o *UpdateDivisionRequest) HasDivisionDescription() bool`

HasDivisionDescription returns a boolean if a field has been set.

### GetDivisionType

`func (o *UpdateDivisionRequest) GetDivisionType() string`

GetDivisionType returns the DivisionType field if non-nil, zero value otherwise.

### GetDivisionTypeOk

`func (o *UpdateDivisionRequest) GetDivisionTypeOk() (*string, bool)`

GetDivisionTypeOk returns a tuple with the DivisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionType

`func (o *UpdateDivisionRequest) SetDivisionType(v string)`

SetDivisionType sets DivisionType field to given value.


### GetDivisionIdType

`func (o *UpdateDivisionRequest) GetDivisionIdType() string`

GetDivisionIdType returns the DivisionIdType field if non-nil, zero value otherwise.

### GetDivisionIdTypeOk

`func (o *UpdateDivisionRequest) GetDivisionIdTypeOk() (*string, bool)`

GetDivisionIdTypeOk returns a tuple with the DivisionIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionIdType

`func (o *UpdateDivisionRequest) SetDivisionIdType(v string)`

SetDivisionIdType sets DivisionIdType field to given value.


### GetExternalDivisionId

`func (o *UpdateDivisionRequest) GetExternalDivisionId() string`

GetExternalDivisionId returns the ExternalDivisionId field if non-nil, zero value otherwise.

### GetExternalDivisionIdOk

`func (o *UpdateDivisionRequest) GetExternalDivisionIdOk() (*string, bool)`

GetExternalDivisionIdOk returns a tuple with the ExternalDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDivisionId

`func (o *UpdateDivisionRequest) SetExternalDivisionId(v string)`

SetExternalDivisionId sets ExternalDivisionId field to given value.

### HasExternalDivisionId

`func (o *UpdateDivisionRequest) HasExternalDivisionId() bool`

HasExternalDivisionId returns a boolean if a field has been set.

### GetNewExternalDivisionId

`func (o *UpdateDivisionRequest) GetNewExternalDivisionId() string`

GetNewExternalDivisionId returns the NewExternalDivisionId field if non-nil, zero value otherwise.

### GetNewExternalDivisionIdOk

`func (o *UpdateDivisionRequest) GetNewExternalDivisionIdOk() (*string, bool)`

GetNewExternalDivisionIdOk returns a tuple with the NewExternalDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewExternalDivisionId

`func (o *UpdateDivisionRequest) SetNewExternalDivisionId(v string)`

SetNewExternalDivisionId sets NewExternalDivisionId field to given value.


### GetLogoUrlMap

`func (o *UpdateDivisionRequest) GetLogoUrlMap() map[string]string`

GetLogoUrlMap returns the LogoUrlMap field if non-nil, zero value otherwise.

### GetLogoUrlMapOk

`func (o *UpdateDivisionRequest) GetLogoUrlMapOk() (*map[string]string, bool)`

GetLogoUrlMapOk returns a tuple with the LogoUrlMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrlMap

`func (o *UpdateDivisionRequest) SetLogoUrlMap(v map[string]string)`

SetLogoUrlMap sets LogoUrlMap field to given value.

### HasLogoUrlMap

`func (o *UpdateDivisionRequest) HasLogoUrlMap() bool`

HasLogoUrlMap returns a boolean if a field has been set.

### GetMccCodes

`func (o *UpdateDivisionRequest) GetMccCodes() []string`

GetMccCodes returns the MccCodes field if non-nil, zero value otherwise.

### GetMccCodesOk

`func (o *UpdateDivisionRequest) GetMccCodesOk() (*[]string, bool)`

GetMccCodesOk returns a tuple with the MccCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccCodes

`func (o *UpdateDivisionRequest) SetMccCodes(v []string)`

SetMccCodes sets MccCodes field to given value.


### GetExtInfo

`func (o *UpdateDivisionRequest) GetExtInfo() map[string]interface{}`

GetExtInfo returns the ExtInfo field if non-nil, zero value otherwise.

### GetExtInfoOk

`func (o *UpdateDivisionRequest) GetExtInfoOk() (*map[string]interface{}, bool)`

GetExtInfoOk returns a tuple with the ExtInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtInfo

`func (o *UpdateDivisionRequest) SetExtInfo(v map[string]interface{})`

SetExtInfo sets ExtInfo field to given value.


### GetApiVersion

`func (o *UpdateDivisionRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UpdateDivisionRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UpdateDivisionRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *UpdateDivisionRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetBusinessDocs

`func (o *UpdateDivisionRequest) GetBusinessDocs() []BusinessDocs`

GetBusinessDocs returns the BusinessDocs field if non-nil, zero value otherwise.

### GetBusinessDocsOk

`func (o *UpdateDivisionRequest) GetBusinessDocsOk() (*[]BusinessDocs, bool)`

GetBusinessDocsOk returns a tuple with the BusinessDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDocs

`func (o *UpdateDivisionRequest) SetBusinessDocs(v []BusinessDocs)`

SetBusinessDocs sets BusinessDocs field to given value.

### HasBusinessDocs

`func (o *UpdateDivisionRequest) HasBusinessDocs() bool`

HasBusinessDocs returns a boolean if a field has been set.

### GetBusinessEntity

`func (o *UpdateDivisionRequest) GetBusinessEntity() string`

GetBusinessEntity returns the BusinessEntity field if non-nil, zero value otherwise.

### GetBusinessEntityOk

`func (o *UpdateDivisionRequest) GetBusinessEntityOk() (*string, bool)`

GetBusinessEntityOk returns a tuple with the BusinessEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessEntity

`func (o *UpdateDivisionRequest) SetBusinessEntity(v string)`

SetBusinessEntity sets BusinessEntity field to given value.

### HasBusinessEntity

`func (o *UpdateDivisionRequest) HasBusinessEntity() bool`

HasBusinessEntity returns a boolean if a field has been set.

### GetBusinessEndDate

`func (o *UpdateDivisionRequest) GetBusinessEndDate() string`

GetBusinessEndDate returns the BusinessEndDate field if non-nil, zero value otherwise.

### GetBusinessEndDateOk

`func (o *UpdateDivisionRequest) GetBusinessEndDateOk() (*string, bool)`

GetBusinessEndDateOk returns a tuple with the BusinessEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessEndDate

`func (o *UpdateDivisionRequest) SetBusinessEndDate(v string)`

SetBusinessEndDate sets BusinessEndDate field to given value.

### HasBusinessEndDate

`func (o *UpdateDivisionRequest) HasBusinessEndDate() bool`

HasBusinessEndDate returns a boolean if a field has been set.

### GetOwnerName

`func (o *UpdateDivisionRequest) GetOwnerName() UserName`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *UpdateDivisionRequest) GetOwnerNameOk() (*UserName, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *UpdateDivisionRequest) SetOwnerName(v UserName)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *UpdateDivisionRequest) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetOwnerPhoneNumber

`func (o *UpdateDivisionRequest) GetOwnerPhoneNumber() MobileNoInfo`

GetOwnerPhoneNumber returns the OwnerPhoneNumber field if non-nil, zero value otherwise.

### GetOwnerPhoneNumberOk

`func (o *UpdateDivisionRequest) GetOwnerPhoneNumberOk() (*MobileNoInfo, bool)`

GetOwnerPhoneNumberOk returns a tuple with the OwnerPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPhoneNumber

`func (o *UpdateDivisionRequest) SetOwnerPhoneNumber(v MobileNoInfo)`

SetOwnerPhoneNumber sets OwnerPhoneNumber field to given value.

### HasOwnerPhoneNumber

`func (o *UpdateDivisionRequest) HasOwnerPhoneNumber() bool`

HasOwnerPhoneNumber returns a boolean if a field has been set.

### GetOwnerIdType

`func (o *UpdateDivisionRequest) GetOwnerIdType() string`

GetOwnerIdType returns the OwnerIdType field if non-nil, zero value otherwise.

### GetOwnerIdTypeOk

`func (o *UpdateDivisionRequest) GetOwnerIdTypeOk() (*string, bool)`

GetOwnerIdTypeOk returns a tuple with the OwnerIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdType

`func (o *UpdateDivisionRequest) SetOwnerIdType(v string)`

SetOwnerIdType sets OwnerIdType field to given value.

### HasOwnerIdType

`func (o *UpdateDivisionRequest) HasOwnerIdType() bool`

HasOwnerIdType returns a boolean if a field has been set.

### GetOwnerIdNo

`func (o *UpdateDivisionRequest) GetOwnerIdNo() string`

GetOwnerIdNo returns the OwnerIdNo field if non-nil, zero value otherwise.

### GetOwnerIdNoOk

`func (o *UpdateDivisionRequest) GetOwnerIdNoOk() (*string, bool)`

GetOwnerIdNoOk returns a tuple with the OwnerIdNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdNo

`func (o *UpdateDivisionRequest) SetOwnerIdNo(v string)`

SetOwnerIdNo sets OwnerIdNo field to given value.

### HasOwnerIdNo

`func (o *UpdateDivisionRequest) HasOwnerIdNo() bool`

HasOwnerIdNo returns a boolean if a field has been set.

### GetOwnerAddress

`func (o *UpdateDivisionRequest) GetOwnerAddress() AddressInfo`

GetOwnerAddress returns the OwnerAddress field if non-nil, zero value otherwise.

### GetOwnerAddressOk

`func (o *UpdateDivisionRequest) GetOwnerAddressOk() (*AddressInfo, bool)`

GetOwnerAddressOk returns a tuple with the OwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAddress

`func (o *UpdateDivisionRequest) SetOwnerAddress(v AddressInfo)`

SetOwnerAddress sets OwnerAddress field to given value.

### HasOwnerAddress

`func (o *UpdateDivisionRequest) HasOwnerAddress() bool`

HasOwnerAddress returns a boolean if a field has been set.

### GetDirectorPics

`func (o *UpdateDivisionRequest) GetDirectorPics() []PicInfo`

GetDirectorPics returns the DirectorPics field if non-nil, zero value otherwise.

### GetDirectorPicsOk

`func (o *UpdateDivisionRequest) GetDirectorPicsOk() (*[]PicInfo, bool)`

GetDirectorPicsOk returns a tuple with the DirectorPics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectorPics

`func (o *UpdateDivisionRequest) SetDirectorPics(v []PicInfo)`

SetDirectorPics sets DirectorPics field to given value.

### HasDirectorPics

`func (o *UpdateDivisionRequest) HasDirectorPics() bool`

HasDirectorPics returns a boolean if a field has been set.

### GetNonDirectorPics

`func (o *UpdateDivisionRequest) GetNonDirectorPics() []PicInfo`

GetNonDirectorPics returns the NonDirectorPics field if non-nil, zero value otherwise.

### GetNonDirectorPicsOk

`func (o *UpdateDivisionRequest) GetNonDirectorPicsOk() (*[]PicInfo, bool)`

GetNonDirectorPicsOk returns a tuple with the NonDirectorPics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDirectorPics

`func (o *UpdateDivisionRequest) SetNonDirectorPics(v []PicInfo)`

SetNonDirectorPics sets NonDirectorPics field to given value.

### HasNonDirectorPics

`func (o *UpdateDivisionRequest) HasNonDirectorPics() bool`

HasNonDirectorPics returns a boolean if a field has been set.

### GetSizeType

`func (o *UpdateDivisionRequest) GetSizeType() string`

GetSizeType returns the SizeType field if non-nil, zero value otherwise.

### GetSizeTypeOk

`func (o *UpdateDivisionRequest) GetSizeTypeOk() (*string, bool)`

GetSizeTypeOk returns a tuple with the SizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeType

`func (o *UpdateDivisionRequest) SetSizeType(v string)`

SetSizeType sets SizeType field to given value.

### HasSizeType

`func (o *UpdateDivisionRequest) HasSizeType() bool`

HasSizeType returns a boolean if a field has been set.

### GetPgDivisionFlag

`func (o *UpdateDivisionRequest) GetPgDivisionFlag() string`

GetPgDivisionFlag returns the PgDivisionFlag field if non-nil, zero value otherwise.

### GetPgDivisionFlagOk

`func (o *UpdateDivisionRequest) GetPgDivisionFlagOk() (*string, bool)`

GetPgDivisionFlagOk returns a tuple with the PgDivisionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgDivisionFlag

`func (o *UpdateDivisionRequest) SetPgDivisionFlag(v string)`

SetPgDivisionFlag sets PgDivisionFlag field to given value.

### HasPgDivisionFlag

`func (o *UpdateDivisionRequest) HasPgDivisionFlag() bool`

HasPgDivisionFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


