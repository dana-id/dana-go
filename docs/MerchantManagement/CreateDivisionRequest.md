# CreateDivisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version. As per the respective API reference. Must be &gt; 2 | 
**MerchantId** | **string** | Merchant identifier | 
**ParentDivisionId** | Pointer to **string** | Parent division identifier. Required when parentRoleType is DIVISION or EXTERNAL_DIVISION. Length depends on parentRoleType - DIVISION (21 max), EXTERNAL_DIVISION (64 max) | [optional] 
**ParentRoleType** | **string** | Type of parent role | 
**DivisionName** | **string** | Division name | 
**DivisionAddress** | [**AddressInfo**](AddressInfo.md) |  | 
**DivisionDescription** | Pointer to **string** | Division description | [optional] 
**DivisionType** | **string** | Division type | 
**ExternalDivisionId** | **string** | External division identifier | 
**LogoUrlMap** | Pointer to **map[string]string** | Logo URL map with base64 encoded images. Keys can be LOGO, PC_LOGO, MOBILE_LOGO | [optional] 
**ExtInfo** | [**CreateDivisionRequestExtInfo**](CreateDivisionRequestExtInfo.md) |  | 
**MccCodes** | **[]string** | Merchant category codes | 
**BusinessDocs** | [**[]BusinessDocs**](BusinessDocs.md) | Business documents. \&quot;individu\&quot; entity can only use KTP and SIM. Other entities can use SIUP and NIB | 
**BusinessEntity** | **string** | Business entity type | 
**OwnerName** | [**UserName**](UserName.md) |  | 
**OwnerPhoneNumber** | [**MobileNoInfo**](MobileNoInfo.md) |  | 
**OwnerIdType** | **string** | Owner identifier type | 
**OwnerIdNo** | **string** | Owner identifier number. Length depends on ownerIdType - KTP (16), SIM (12-14), Passport (8), NIB (&gt;&#x3D;13), SIUP (free text) | 
**OwnerAddress** | [**AddressInfo**](AddressInfo.md) |  | 
**DirectorPics** | [**[]PicInfo**](PicInfo.md) | Director as a PIC of sub merchant | 
**NonDirectorPics** | [**[]PicInfo**](PicInfo.md) | Non director which become a PIC of sub merchant | 
**SizeType** | **string** | Size type | 
**PgDivisionFlag** | Pointer to **string** | Flag if division is type PG | [optional] 

## Methods

### NewCreateDivisionRequest

`func NewCreateDivisionRequest(apiVersion string, merchantId string, parentRoleType string, divisionName string, divisionAddress AddressInfo, divisionType string, externalDivisionId string, extInfo CreateDivisionRequestExtInfo, mccCodes []string, businessDocs []BusinessDocs, businessEntity string, ownerName UserName, ownerPhoneNumber MobileNoInfo, ownerIdType string, ownerIdNo string, ownerAddress AddressInfo, directorPics []PicInfo, nonDirectorPics []PicInfo, sizeType string, ) *CreateDivisionRequest`

NewCreateDivisionRequest instantiates a new CreateDivisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDivisionRequestWithDefaults

`func NewCreateDivisionRequestWithDefaults() *CreateDivisionRequest`

NewCreateDivisionRequestWithDefaults instantiates a new CreateDivisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateDivisionRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateDivisionRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateDivisionRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMerchantId

`func (o *CreateDivisionRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CreateDivisionRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CreateDivisionRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetParentDivisionId

`func (o *CreateDivisionRequest) GetParentDivisionId() string`

GetParentDivisionId returns the ParentDivisionId field if non-nil, zero value otherwise.

### GetParentDivisionIdOk

`func (o *CreateDivisionRequest) GetParentDivisionIdOk() (*string, bool)`

GetParentDivisionIdOk returns a tuple with the ParentDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDivisionId

`func (o *CreateDivisionRequest) SetParentDivisionId(v string)`

SetParentDivisionId sets ParentDivisionId field to given value.

### HasParentDivisionId

`func (o *CreateDivisionRequest) HasParentDivisionId() bool`

HasParentDivisionId returns a boolean if a field has been set.

### GetParentRoleType

`func (o *CreateDivisionRequest) GetParentRoleType() string`

GetParentRoleType returns the ParentRoleType field if non-nil, zero value otherwise.

### GetParentRoleTypeOk

`func (o *CreateDivisionRequest) GetParentRoleTypeOk() (*string, bool)`

GetParentRoleTypeOk returns a tuple with the ParentRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoleType

`func (o *CreateDivisionRequest) SetParentRoleType(v string)`

SetParentRoleType sets ParentRoleType field to given value.


### GetDivisionName

`func (o *CreateDivisionRequest) GetDivisionName() string`

GetDivisionName returns the DivisionName field if non-nil, zero value otherwise.

### GetDivisionNameOk

`func (o *CreateDivisionRequest) GetDivisionNameOk() (*string, bool)`

GetDivisionNameOk returns a tuple with the DivisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionName

`func (o *CreateDivisionRequest) SetDivisionName(v string)`

SetDivisionName sets DivisionName field to given value.


### GetDivisionAddress

`func (o *CreateDivisionRequest) GetDivisionAddress() AddressInfo`

GetDivisionAddress returns the DivisionAddress field if non-nil, zero value otherwise.

### GetDivisionAddressOk

`func (o *CreateDivisionRequest) GetDivisionAddressOk() (*AddressInfo, bool)`

GetDivisionAddressOk returns a tuple with the DivisionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionAddress

`func (o *CreateDivisionRequest) SetDivisionAddress(v AddressInfo)`

SetDivisionAddress sets DivisionAddress field to given value.


### GetDivisionDescription

`func (o *CreateDivisionRequest) GetDivisionDescription() string`

GetDivisionDescription returns the DivisionDescription field if non-nil, zero value otherwise.

### GetDivisionDescriptionOk

`func (o *CreateDivisionRequest) GetDivisionDescriptionOk() (*string, bool)`

GetDivisionDescriptionOk returns a tuple with the DivisionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionDescription

`func (o *CreateDivisionRequest) SetDivisionDescription(v string)`

SetDivisionDescription sets DivisionDescription field to given value.

### HasDivisionDescription

`func (o *CreateDivisionRequest) HasDivisionDescription() bool`

HasDivisionDescription returns a boolean if a field has been set.

### GetDivisionType

`func (o *CreateDivisionRequest) GetDivisionType() string`

GetDivisionType returns the DivisionType field if non-nil, zero value otherwise.

### GetDivisionTypeOk

`func (o *CreateDivisionRequest) GetDivisionTypeOk() (*string, bool)`

GetDivisionTypeOk returns a tuple with the DivisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionType

`func (o *CreateDivisionRequest) SetDivisionType(v string)`

SetDivisionType sets DivisionType field to given value.


### GetExternalDivisionId

`func (o *CreateDivisionRequest) GetExternalDivisionId() string`

GetExternalDivisionId returns the ExternalDivisionId field if non-nil, zero value otherwise.

### GetExternalDivisionIdOk

`func (o *CreateDivisionRequest) GetExternalDivisionIdOk() (*string, bool)`

GetExternalDivisionIdOk returns a tuple with the ExternalDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDivisionId

`func (o *CreateDivisionRequest) SetExternalDivisionId(v string)`

SetExternalDivisionId sets ExternalDivisionId field to given value.


### GetLogoUrlMap

`func (o *CreateDivisionRequest) GetLogoUrlMap() map[string]string`

GetLogoUrlMap returns the LogoUrlMap field if non-nil, zero value otherwise.

### GetLogoUrlMapOk

`func (o *CreateDivisionRequest) GetLogoUrlMapOk() (*map[string]string, bool)`

GetLogoUrlMapOk returns a tuple with the LogoUrlMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrlMap

`func (o *CreateDivisionRequest) SetLogoUrlMap(v map[string]string)`

SetLogoUrlMap sets LogoUrlMap field to given value.

### HasLogoUrlMap

`func (o *CreateDivisionRequest) HasLogoUrlMap() bool`

HasLogoUrlMap returns a boolean if a field has been set.

### GetExtInfo

`func (o *CreateDivisionRequest) GetExtInfo() CreateDivisionRequestExtInfo`

GetExtInfo returns the ExtInfo field if non-nil, zero value otherwise.

### GetExtInfoOk

`func (o *CreateDivisionRequest) GetExtInfoOk() (*CreateDivisionRequestExtInfo, bool)`

GetExtInfoOk returns a tuple with the ExtInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtInfo

`func (o *CreateDivisionRequest) SetExtInfo(v CreateDivisionRequestExtInfo)`

SetExtInfo sets ExtInfo field to given value.


### GetMccCodes

`func (o *CreateDivisionRequest) GetMccCodes() []string`

GetMccCodes returns the MccCodes field if non-nil, zero value otherwise.

### GetMccCodesOk

`func (o *CreateDivisionRequest) GetMccCodesOk() (*[]string, bool)`

GetMccCodesOk returns a tuple with the MccCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccCodes

`func (o *CreateDivisionRequest) SetMccCodes(v []string)`

SetMccCodes sets MccCodes field to given value.


### GetBusinessDocs

`func (o *CreateDivisionRequest) GetBusinessDocs() []BusinessDocs`

GetBusinessDocs returns the BusinessDocs field if non-nil, zero value otherwise.

### GetBusinessDocsOk

`func (o *CreateDivisionRequest) GetBusinessDocsOk() (*[]BusinessDocs, bool)`

GetBusinessDocsOk returns a tuple with the BusinessDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDocs

`func (o *CreateDivisionRequest) SetBusinessDocs(v []BusinessDocs)`

SetBusinessDocs sets BusinessDocs field to given value.


### GetBusinessEntity

`func (o *CreateDivisionRequest) GetBusinessEntity() string`

GetBusinessEntity returns the BusinessEntity field if non-nil, zero value otherwise.

### GetBusinessEntityOk

`func (o *CreateDivisionRequest) GetBusinessEntityOk() (*string, bool)`

GetBusinessEntityOk returns a tuple with the BusinessEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessEntity

`func (o *CreateDivisionRequest) SetBusinessEntity(v string)`

SetBusinessEntity sets BusinessEntity field to given value.


### GetOwnerName

`func (o *CreateDivisionRequest) GetOwnerName() UserName`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *CreateDivisionRequest) GetOwnerNameOk() (*UserName, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *CreateDivisionRequest) SetOwnerName(v UserName)`

SetOwnerName sets OwnerName field to given value.


### GetOwnerPhoneNumber

`func (o *CreateDivisionRequest) GetOwnerPhoneNumber() MobileNoInfo`

GetOwnerPhoneNumber returns the OwnerPhoneNumber field if non-nil, zero value otherwise.

### GetOwnerPhoneNumberOk

`func (o *CreateDivisionRequest) GetOwnerPhoneNumberOk() (*MobileNoInfo, bool)`

GetOwnerPhoneNumberOk returns a tuple with the OwnerPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPhoneNumber

`func (o *CreateDivisionRequest) SetOwnerPhoneNumber(v MobileNoInfo)`

SetOwnerPhoneNumber sets OwnerPhoneNumber field to given value.


### GetOwnerIdType

`func (o *CreateDivisionRequest) GetOwnerIdType() string`

GetOwnerIdType returns the OwnerIdType field if non-nil, zero value otherwise.

### GetOwnerIdTypeOk

`func (o *CreateDivisionRequest) GetOwnerIdTypeOk() (*string, bool)`

GetOwnerIdTypeOk returns a tuple with the OwnerIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdType

`func (o *CreateDivisionRequest) SetOwnerIdType(v string)`

SetOwnerIdType sets OwnerIdType field to given value.


### GetOwnerIdNo

`func (o *CreateDivisionRequest) GetOwnerIdNo() string`

GetOwnerIdNo returns the OwnerIdNo field if non-nil, zero value otherwise.

### GetOwnerIdNoOk

`func (o *CreateDivisionRequest) GetOwnerIdNoOk() (*string, bool)`

GetOwnerIdNoOk returns a tuple with the OwnerIdNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdNo

`func (o *CreateDivisionRequest) SetOwnerIdNo(v string)`

SetOwnerIdNo sets OwnerIdNo field to given value.


### GetOwnerAddress

`func (o *CreateDivisionRequest) GetOwnerAddress() AddressInfo`

GetOwnerAddress returns the OwnerAddress field if non-nil, zero value otherwise.

### GetOwnerAddressOk

`func (o *CreateDivisionRequest) GetOwnerAddressOk() (*AddressInfo, bool)`

GetOwnerAddressOk returns a tuple with the OwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAddress

`func (o *CreateDivisionRequest) SetOwnerAddress(v AddressInfo)`

SetOwnerAddress sets OwnerAddress field to given value.


### GetDirectorPics

`func (o *CreateDivisionRequest) GetDirectorPics() []PicInfo`

GetDirectorPics returns the DirectorPics field if non-nil, zero value otherwise.

### GetDirectorPicsOk

`func (o *CreateDivisionRequest) GetDirectorPicsOk() (*[]PicInfo, bool)`

GetDirectorPicsOk returns a tuple with the DirectorPics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectorPics

`func (o *CreateDivisionRequest) SetDirectorPics(v []PicInfo)`

SetDirectorPics sets DirectorPics field to given value.


### GetNonDirectorPics

`func (o *CreateDivisionRequest) GetNonDirectorPics() []PicInfo`

GetNonDirectorPics returns the NonDirectorPics field if non-nil, zero value otherwise.

### GetNonDirectorPicsOk

`func (o *CreateDivisionRequest) GetNonDirectorPicsOk() (*[]PicInfo, bool)`

GetNonDirectorPicsOk returns a tuple with the NonDirectorPics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDirectorPics

`func (o *CreateDivisionRequest) SetNonDirectorPics(v []PicInfo)`

SetNonDirectorPics sets NonDirectorPics field to given value.


### GetSizeType

`func (o *CreateDivisionRequest) GetSizeType() string`

GetSizeType returns the SizeType field if non-nil, zero value otherwise.

### GetSizeTypeOk

`func (o *CreateDivisionRequest) GetSizeTypeOk() (*string, bool)`

GetSizeTypeOk returns a tuple with the SizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeType

`func (o *CreateDivisionRequest) SetSizeType(v string)`

SetSizeType sets SizeType field to given value.


### GetPgDivisionFlag

`func (o *CreateDivisionRequest) GetPgDivisionFlag() string`

GetPgDivisionFlag returns the PgDivisionFlag field if non-nil, zero value otherwise.

### GetPgDivisionFlagOk

`func (o *CreateDivisionRequest) GetPgDivisionFlagOk() (*string, bool)`

GetPgDivisionFlagOk returns a tuple with the PgDivisionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgDivisionFlag

`func (o *CreateDivisionRequest) SetPgDivisionFlag(v string)`

SetPgDivisionFlag sets PgDivisionFlag field to given value.

### HasPgDivisionFlag

`func (o *CreateDivisionRequest) HasPgDivisionFlag() bool`

HasPgDivisionFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


