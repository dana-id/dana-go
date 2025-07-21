# DivisionResourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DivisionId** | Pointer to **string** | Division identifier | [optional] 
**MerchantId** | Pointer to **string** | Merchant identifier | [optional] 
**ParentRoleType** | Pointer to **string** | Parent role type | [optional] 
**ContactAddress** | Pointer to [**AddressInfo**](AddressInfo.md) |  | [optional] 
**DivisionDescription** | Pointer to **string** | Division description | [optional] 
**DivisionType** | Pointer to **string** | Division type | [optional] 
**DivisionName** | Pointer to **string** | Division name | [optional] 
**ExternalDivisionId** | Pointer to **string** | External division identifier | [optional] 
**PgDivisionFlag** | Pointer to **string** | Flag if division is type PG | [optional] 

## Methods

### NewDivisionResourceInfo

`func NewDivisionResourceInfo() *DivisionResourceInfo`

NewDivisionResourceInfo instantiates a new DivisionResourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDivisionResourceInfoWithDefaults

`func NewDivisionResourceInfoWithDefaults() *DivisionResourceInfo`

NewDivisionResourceInfoWithDefaults instantiates a new DivisionResourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivisionId

`func (o *DivisionResourceInfo) GetDivisionId() string`

GetDivisionId returns the DivisionId field if non-nil, zero value otherwise.

### GetDivisionIdOk

`func (o *DivisionResourceInfo) GetDivisionIdOk() (*string, bool)`

GetDivisionIdOk returns a tuple with the DivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionId

`func (o *DivisionResourceInfo) SetDivisionId(v string)`

SetDivisionId sets DivisionId field to given value.

### HasDivisionId

`func (o *DivisionResourceInfo) HasDivisionId() bool`

HasDivisionId returns a boolean if a field has been set.

### GetMerchantId

`func (o *DivisionResourceInfo) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *DivisionResourceInfo) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *DivisionResourceInfo) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *DivisionResourceInfo) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetParentRoleType

`func (o *DivisionResourceInfo) GetParentRoleType() string`

GetParentRoleType returns the ParentRoleType field if non-nil, zero value otherwise.

### GetParentRoleTypeOk

`func (o *DivisionResourceInfo) GetParentRoleTypeOk() (*string, bool)`

GetParentRoleTypeOk returns a tuple with the ParentRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoleType

`func (o *DivisionResourceInfo) SetParentRoleType(v string)`

SetParentRoleType sets ParentRoleType field to given value.

### HasParentRoleType

`func (o *DivisionResourceInfo) HasParentRoleType() bool`

HasParentRoleType returns a boolean if a field has been set.

### GetContactAddress

`func (o *DivisionResourceInfo) GetContactAddress() AddressInfo`

GetContactAddress returns the ContactAddress field if non-nil, zero value otherwise.

### GetContactAddressOk

`func (o *DivisionResourceInfo) GetContactAddressOk() (*AddressInfo, bool)`

GetContactAddressOk returns a tuple with the ContactAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAddress

`func (o *DivisionResourceInfo) SetContactAddress(v AddressInfo)`

SetContactAddress sets ContactAddress field to given value.

### HasContactAddress

`func (o *DivisionResourceInfo) HasContactAddress() bool`

HasContactAddress returns a boolean if a field has been set.

### GetDivisionDescription

`func (o *DivisionResourceInfo) GetDivisionDescription() string`

GetDivisionDescription returns the DivisionDescription field if non-nil, zero value otherwise.

### GetDivisionDescriptionOk

`func (o *DivisionResourceInfo) GetDivisionDescriptionOk() (*string, bool)`

GetDivisionDescriptionOk returns a tuple with the DivisionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionDescription

`func (o *DivisionResourceInfo) SetDivisionDescription(v string)`

SetDivisionDescription sets DivisionDescription field to given value.

### HasDivisionDescription

`func (o *DivisionResourceInfo) HasDivisionDescription() bool`

HasDivisionDescription returns a boolean if a field has been set.

### GetDivisionType

`func (o *DivisionResourceInfo) GetDivisionType() string`

GetDivisionType returns the DivisionType field if non-nil, zero value otherwise.

### GetDivisionTypeOk

`func (o *DivisionResourceInfo) GetDivisionTypeOk() (*string, bool)`

GetDivisionTypeOk returns a tuple with the DivisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionType

`func (o *DivisionResourceInfo) SetDivisionType(v string)`

SetDivisionType sets DivisionType field to given value.

### HasDivisionType

`func (o *DivisionResourceInfo) HasDivisionType() bool`

HasDivisionType returns a boolean if a field has been set.

### GetDivisionName

`func (o *DivisionResourceInfo) GetDivisionName() string`

GetDivisionName returns the DivisionName field if non-nil, zero value otherwise.

### GetDivisionNameOk

`func (o *DivisionResourceInfo) GetDivisionNameOk() (*string, bool)`

GetDivisionNameOk returns a tuple with the DivisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionName

`func (o *DivisionResourceInfo) SetDivisionName(v string)`

SetDivisionName sets DivisionName field to given value.

### HasDivisionName

`func (o *DivisionResourceInfo) HasDivisionName() bool`

HasDivisionName returns a boolean if a field has been set.

### GetExternalDivisionId

`func (o *DivisionResourceInfo) GetExternalDivisionId() string`

GetExternalDivisionId returns the ExternalDivisionId field if non-nil, zero value otherwise.

### GetExternalDivisionIdOk

`func (o *DivisionResourceInfo) GetExternalDivisionIdOk() (*string, bool)`

GetExternalDivisionIdOk returns a tuple with the ExternalDivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDivisionId

`func (o *DivisionResourceInfo) SetExternalDivisionId(v string)`

SetExternalDivisionId sets ExternalDivisionId field to given value.

### HasExternalDivisionId

`func (o *DivisionResourceInfo) HasExternalDivisionId() bool`

HasExternalDivisionId returns a boolean if a field has been set.

### GetPgDivisionFlag

`func (o *DivisionResourceInfo) GetPgDivisionFlag() string`

GetPgDivisionFlag returns the PgDivisionFlag field if non-nil, zero value otherwise.

### GetPgDivisionFlagOk

`func (o *DivisionResourceInfo) GetPgDivisionFlagOk() (*string, bool)`

GetPgDivisionFlagOk returns a tuple with the PgDivisionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgDivisionFlag

`func (o *DivisionResourceInfo) SetPgDivisionFlag(v string)`

SetPgDivisionFlag sets PgDivisionFlag field to given value.

### HasPgDivisionFlag

`func (o *DivisionResourceInfo) HasPgDivisionFlag() bool`

HasPgDivisionFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


