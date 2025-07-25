// Copyright 2025 PT Espay Debit Indonesia Koe
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package merchant_management

/*
Merchant Management API

API for merchant management operations in DANA

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.


import (
	"encoding/json"
	"bytes"
	"fmt"
	utils "github.com/dana-id/dana-go/utils"
)

// checks if the CreateShopRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateShopRequest{}

// CreateShopRequest struct for CreateShopRequest
type CreateShopRequest struct {
	// Merchant id
	MerchantId string `json:"merchantId"`
	// Parent division ID. Required when shopParentType is DIVISION or EXTERNAL_DIVISION
	ParentDivisionId *string `json:"parentDivisionId,omitempty"`
	// Parent type for the shop
	ShopParentType string `json:"shopParentType"`
	// Shop name
	MainName string `json:"mainName"`
	ShopAddress *AddressInfo `json:"shopAddress,omitempty"`
	// Shop description
	ShopDesc *string `json:"shopDesc,omitempty"`
	// External shop identifier
	ExternalShopId string `json:"externalShopId"`
	// Logo images encoded in base64. Keys can be LOGO, PC_LOGO, MOBILE_LOGO. Images must be PNG format.
	LogoUrlMap map[string]string `json:"logoUrlMap,omitempty"`
	// MCC codes
	MccCodes []string `json:"mccCodes,omitempty"`
	// Size type of the shop
	SizeType string `json:"sizeType"`
	// Longitude
	Ln *string `json:"ln,omitempty"`
	// Latitude
	Lat *string `json:"lat,omitempty"`
	// Tax number (NPWP). Must be 15 digits
	TaxNo *string `json:"taxNo,omitempty" validate:"regexp=^[0-9]{15}$"`
	// Legal name/tax name
	BrandName *string `json:"brandName,omitempty"`
	TaxAddress *AddressInfo `json:"taxAddress,omitempty"`
	// Flag for loyalty category
	Loyalty *string `json:"loyalty,omitempty"`
	// API version flag. Use > 2 for new attributes
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Business entity type. Required if apiVersion > 2
	BusinessEntity *string `json:"businessEntity,omitempty"`
	// Business documents. Required if apiVersion > 2
	BusinessDocs []BusinessDocs `json:"businessDocs,omitempty"`
	OwnerName *UserName `json:"ownerName,omitempty"`
	// Owner ID type. Required if apiVersion > 2
	OwnerIdType *string `json:"ownerIdType,omitempty"`
	// Owner ID number. Required if apiVersion > 2
	OwnerIdNo *string `json:"ownerIdNo,omitempty"`
	OwnerAddress *AddressInfo `json:"ownerAddress,omitempty"`
	OwnerPhoneNumber *MobileNoInfo `json:"ownerPhoneNumber,omitempty"`
	// Shop ownership type. Required if apiVersion > 2
	ShopOwning *string `json:"shopOwning,omitempty"`
	// Device number. Required if apiVersion > 2
	DeviceNumber *string `json:"deviceNumber,omitempty"`
	// POS number. Required if apiVersion > 2
	PosNumber *string `json:"posNumber,omitempty"`
	// Director PICs. Required if apiVersion > 2
	DirectorPics []PicInfo `json:"directorPics,omitempty"`
	// Non-director PICs. Required if apiVersion > 2
	NonDirectorPics []PicInfo `json:"nonDirectorPics,omitempty"`
	// Shop business type
	ShopBizType *string `json:"shopBizType,omitempty"`
	// Extended information
	ExtInfo map[string]interface{} `json:"extInfo,omitempty"`
}

type _CreateShopRequest CreateShopRequest

// NewCreateShopRequest instantiates a new CreateShopRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShopRequest(merchantId string, shopParentType string, mainName string, externalShopId string, sizeType string) *CreateShopRequest {
	this := CreateShopRequest{}
	this.MerchantId = merchantId
	this.ShopParentType = shopParentType
	this.MainName = mainName
	this.ExternalShopId = externalShopId
	this.SizeType = sizeType
	return &this
}

// NewCreateShopRequestWithDefaults instantiates a new CreateShopRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShopRequestWithDefaults() *CreateShopRequest {
	this := CreateShopRequest{}
	return &this
}

// GetMerchantId returns the MerchantId field value
func (o *CreateShopRequest) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *CreateShopRequest) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetParentDivisionId returns the ParentDivisionId field value if set, zero value otherwise.
func (o *CreateShopRequest) GetParentDivisionId() string {
	if o == nil || utils.IsNil(o.ParentDivisionId) {
		var ret string
		return ret
	}
	return *o.ParentDivisionId
}

// GetParentDivisionIdOk returns a tuple with the ParentDivisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetParentDivisionIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ParentDivisionId) {
		return nil, false
	}
	return o.ParentDivisionId, true
}

// HasParentDivisionId returns a boolean if a field has been set.
func (o *CreateShopRequest) HasParentDivisionId() bool {
	if o != nil && !utils.IsNil(o.ParentDivisionId) {
		return true
	}

	return false
}

// SetParentDivisionId gets a reference to the given string and assigns it to the ParentDivisionId field.
func (o *CreateShopRequest) SetParentDivisionId(v string) {
	o.ParentDivisionId = &v
}

// GetShopParentType returns the ShopParentType field value
func (o *CreateShopRequest) GetShopParentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopParentType
}

// GetShopParentTypeOk returns a tuple with the ShopParentType field value
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetShopParentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopParentType, true
}

// SetShopParentType sets field value
func (o *CreateShopRequest) SetShopParentType(v string) {
	o.ShopParentType = v
}

// GetMainName returns the MainName field value
func (o *CreateShopRequest) GetMainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MainName
}

// GetMainNameOk returns a tuple with the MainName field value
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetMainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MainName, true
}

// SetMainName sets field value
func (o *CreateShopRequest) SetMainName(v string) {
	o.MainName = v
}

// GetShopAddress returns the ShopAddress field value if set, zero value otherwise.
func (o *CreateShopRequest) GetShopAddress() AddressInfo {
	if o == nil || utils.IsNil(o.ShopAddress) {
		var ret AddressInfo
		return ret
	}
	return *o.ShopAddress
}

// GetShopAddressOk returns a tuple with the ShopAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetShopAddressOk() (*AddressInfo, bool) {
	if o == nil || utils.IsNil(o.ShopAddress) {
		return nil, false
	}
	return o.ShopAddress, true
}

// HasShopAddress returns a boolean if a field has been set.
func (o *CreateShopRequest) HasShopAddress() bool {
	if o != nil && !utils.IsNil(o.ShopAddress) {
		return true
	}

	return false
}

// SetShopAddress gets a reference to the given AddressInfo and assigns it to the ShopAddress field.
func (o *CreateShopRequest) SetShopAddress(v AddressInfo) {
	o.ShopAddress = &v
}

// GetShopDesc returns the ShopDesc field value if set, zero value otherwise.
func (o *CreateShopRequest) GetShopDesc() string {
	if o == nil || utils.IsNil(o.ShopDesc) {
		var ret string
		return ret
	}
	return *o.ShopDesc
}

// GetShopDescOk returns a tuple with the ShopDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetShopDescOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ShopDesc) {
		return nil, false
	}
	return o.ShopDesc, true
}

// HasShopDesc returns a boolean if a field has been set.
func (o *CreateShopRequest) HasShopDesc() bool {
	if o != nil && !utils.IsNil(o.ShopDesc) {
		return true
	}

	return false
}

// SetShopDesc gets a reference to the given string and assigns it to the ShopDesc field.
func (o *CreateShopRequest) SetShopDesc(v string) {
	o.ShopDesc = &v
}

// GetExternalShopId returns the ExternalShopId field value
func (o *CreateShopRequest) GetExternalShopId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalShopId
}

// GetExternalShopIdOk returns a tuple with the ExternalShopId field value
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetExternalShopIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalShopId, true
}

// SetExternalShopId sets field value
func (o *CreateShopRequest) SetExternalShopId(v string) {
	o.ExternalShopId = v
}

// GetLogoUrlMap returns the LogoUrlMap field value if set, zero value otherwise.
func (o *CreateShopRequest) GetLogoUrlMap() map[string]string {
	if o == nil || utils.IsNil(o.LogoUrlMap) {
		var ret map[string]string
		return ret
	}
	return o.LogoUrlMap
}

// GetLogoUrlMapOk returns a tuple with the LogoUrlMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetLogoUrlMapOk() (map[string]string, bool) {
	if o == nil || utils.IsNil(o.LogoUrlMap) {
		return map[string]string{}, false
	}
	return o.LogoUrlMap, true
}

// HasLogoUrlMap returns a boolean if a field has been set.
func (o *CreateShopRequest) HasLogoUrlMap() bool {
	if o != nil && !utils.IsNil(o.LogoUrlMap) {
		return true
	}

	return false
}

// SetLogoUrlMap gets a reference to the given map[string]string and assigns it to the LogoUrlMap field.
func (o *CreateShopRequest) SetLogoUrlMap(v map[string]string) {
	o.LogoUrlMap = v
}

// GetMccCodes returns the MccCodes field value if set, zero value otherwise.
func (o *CreateShopRequest) GetMccCodes() []string {
	if o == nil || utils.IsNil(o.MccCodes) {
		var ret []string
		return ret
	}
	return o.MccCodes
}

// GetMccCodesOk returns a tuple with the MccCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetMccCodesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.MccCodes) {
		return nil, false
	}
	return o.MccCodes, true
}

// HasMccCodes returns a boolean if a field has been set.
func (o *CreateShopRequest) HasMccCodes() bool {
	if o != nil && !utils.IsNil(o.MccCodes) {
		return true
	}

	return false
}

// SetMccCodes gets a reference to the given []string and assigns it to the MccCodes field.
func (o *CreateShopRequest) SetMccCodes(v []string) {
	o.MccCodes = v
}

// GetSizeType returns the SizeType field value
func (o *CreateShopRequest) GetSizeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SizeType
}

// GetSizeTypeOk returns a tuple with the SizeType field value
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetSizeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeType, true
}

// SetSizeType sets field value
func (o *CreateShopRequest) SetSizeType(v string) {
	o.SizeType = v
}

// GetLn returns the Ln field value if set, zero value otherwise.
func (o *CreateShopRequest) GetLn() string {
	if o == nil || utils.IsNil(o.Ln) {
		var ret string
		return ret
	}
	return *o.Ln
}

// GetLnOk returns a tuple with the Ln field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetLnOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Ln) {
		return nil, false
	}
	return o.Ln, true
}

// HasLn returns a boolean if a field has been set.
func (o *CreateShopRequest) HasLn() bool {
	if o != nil && !utils.IsNil(o.Ln) {
		return true
	}

	return false
}

// SetLn gets a reference to the given string and assigns it to the Ln field.
func (o *CreateShopRequest) SetLn(v string) {
	o.Ln = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *CreateShopRequest) GetLat() string {
	if o == nil || utils.IsNil(o.Lat) {
		var ret string
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetLatOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Lat) {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *CreateShopRequest) HasLat() bool {
	if o != nil && !utils.IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given string and assigns it to the Lat field.
func (o *CreateShopRequest) SetLat(v string) {
	o.Lat = &v
}

// GetTaxNo returns the TaxNo field value if set, zero value otherwise.
func (o *CreateShopRequest) GetTaxNo() string {
	if o == nil || utils.IsNil(o.TaxNo) {
		var ret string
		return ret
	}
	return *o.TaxNo
}

// GetTaxNoOk returns a tuple with the TaxNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetTaxNoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TaxNo) {
		return nil, false
	}
	return o.TaxNo, true
}

// HasTaxNo returns a boolean if a field has been set.
func (o *CreateShopRequest) HasTaxNo() bool {
	if o != nil && !utils.IsNil(o.TaxNo) {
		return true
	}

	return false
}

// SetTaxNo gets a reference to the given string and assigns it to the TaxNo field.
func (o *CreateShopRequest) SetTaxNo(v string) {
	o.TaxNo = &v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *CreateShopRequest) GetBrandName() string {
	if o == nil || utils.IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetBrandNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *CreateShopRequest) HasBrandName() bool {
	if o != nil && !utils.IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *CreateShopRequest) SetBrandName(v string) {
	o.BrandName = &v
}

// GetTaxAddress returns the TaxAddress field value if set, zero value otherwise.
func (o *CreateShopRequest) GetTaxAddress() AddressInfo {
	if o == nil || utils.IsNil(o.TaxAddress) {
		var ret AddressInfo
		return ret
	}
	return *o.TaxAddress
}

// GetTaxAddressOk returns a tuple with the TaxAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetTaxAddressOk() (*AddressInfo, bool) {
	if o == nil || utils.IsNil(o.TaxAddress) {
		return nil, false
	}
	return o.TaxAddress, true
}

// HasTaxAddress returns a boolean if a field has been set.
func (o *CreateShopRequest) HasTaxAddress() bool {
	if o != nil && !utils.IsNil(o.TaxAddress) {
		return true
	}

	return false
}

// SetTaxAddress gets a reference to the given AddressInfo and assigns it to the TaxAddress field.
func (o *CreateShopRequest) SetTaxAddress(v AddressInfo) {
	o.TaxAddress = &v
}

// GetLoyalty returns the Loyalty field value if set, zero value otherwise.
func (o *CreateShopRequest) GetLoyalty() string {
	if o == nil || utils.IsNil(o.Loyalty) {
		var ret string
		return ret
	}
	return *o.Loyalty
}

// GetLoyaltyOk returns a tuple with the Loyalty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetLoyaltyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Loyalty) {
		return nil, false
	}
	return o.Loyalty, true
}

// HasLoyalty returns a boolean if a field has been set.
func (o *CreateShopRequest) HasLoyalty() bool {
	if o != nil && !utils.IsNil(o.Loyalty) {
		return true
	}

	return false
}

// SetLoyalty gets a reference to the given string and assigns it to the Loyalty field.
func (o *CreateShopRequest) SetLoyalty(v string) {
	o.Loyalty = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *CreateShopRequest) GetApiVersion() string {
	if o == nil || utils.IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetApiVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *CreateShopRequest) HasApiVersion() bool {
	if o != nil && !utils.IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *CreateShopRequest) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetBusinessEntity returns the BusinessEntity field value if set, zero value otherwise.
func (o *CreateShopRequest) GetBusinessEntity() string {
	if o == nil || utils.IsNil(o.BusinessEntity) {
		var ret string
		return ret
	}
	return *o.BusinessEntity
}

// GetBusinessEntityOk returns a tuple with the BusinessEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetBusinessEntityOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BusinessEntity) {
		return nil, false
	}
	return o.BusinessEntity, true
}

// HasBusinessEntity returns a boolean if a field has been set.
func (o *CreateShopRequest) HasBusinessEntity() bool {
	if o != nil && !utils.IsNil(o.BusinessEntity) {
		return true
	}

	return false
}

// SetBusinessEntity gets a reference to the given string and assigns it to the BusinessEntity field.
func (o *CreateShopRequest) SetBusinessEntity(v string) {
	o.BusinessEntity = &v
}

// GetBusinessDocs returns the BusinessDocs field value if set, zero value otherwise.
func (o *CreateShopRequest) GetBusinessDocs() []BusinessDocs {
	if o == nil || utils.IsNil(o.BusinessDocs) {
		var ret []BusinessDocs
		return ret
	}
	return o.BusinessDocs
}

// GetBusinessDocsOk returns a tuple with the BusinessDocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetBusinessDocsOk() ([]BusinessDocs, bool) {
	if o == nil || utils.IsNil(o.BusinessDocs) {
		return nil, false
	}
	return o.BusinessDocs, true
}

// HasBusinessDocs returns a boolean if a field has been set.
func (o *CreateShopRequest) HasBusinessDocs() bool {
	if o != nil && !utils.IsNil(o.BusinessDocs) {
		return true
	}

	return false
}

// SetBusinessDocs gets a reference to the given []BusinessDocs and assigns it to the BusinessDocs field.
func (o *CreateShopRequest) SetBusinessDocs(v []BusinessDocs) {
	o.BusinessDocs = v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *CreateShopRequest) GetOwnerName() UserName {
	if o == nil || utils.IsNil(o.OwnerName) {
		var ret UserName
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetOwnerNameOk() (*UserName, bool) {
	if o == nil || utils.IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *CreateShopRequest) HasOwnerName() bool {
	if o != nil && !utils.IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given UserName and assigns it to the OwnerName field.
func (o *CreateShopRequest) SetOwnerName(v UserName) {
	o.OwnerName = &v
}

// GetOwnerIdType returns the OwnerIdType field value if set, zero value otherwise.
func (o *CreateShopRequest) GetOwnerIdType() string {
	if o == nil || utils.IsNil(o.OwnerIdType) {
		var ret string
		return ret
	}
	return *o.OwnerIdType
}

// GetOwnerIdTypeOk returns a tuple with the OwnerIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetOwnerIdTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OwnerIdType) {
		return nil, false
	}
	return o.OwnerIdType, true
}

// HasOwnerIdType returns a boolean if a field has been set.
func (o *CreateShopRequest) HasOwnerIdType() bool {
	if o != nil && !utils.IsNil(o.OwnerIdType) {
		return true
	}

	return false
}

// SetOwnerIdType gets a reference to the given string and assigns it to the OwnerIdType field.
func (o *CreateShopRequest) SetOwnerIdType(v string) {
	o.OwnerIdType = &v
}

// GetOwnerIdNo returns the OwnerIdNo field value if set, zero value otherwise.
func (o *CreateShopRequest) GetOwnerIdNo() string {
	if o == nil || utils.IsNil(o.OwnerIdNo) {
		var ret string
		return ret
	}
	return *o.OwnerIdNo
}

// GetOwnerIdNoOk returns a tuple with the OwnerIdNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetOwnerIdNoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OwnerIdNo) {
		return nil, false
	}
	return o.OwnerIdNo, true
}

// HasOwnerIdNo returns a boolean if a field has been set.
func (o *CreateShopRequest) HasOwnerIdNo() bool {
	if o != nil && !utils.IsNil(o.OwnerIdNo) {
		return true
	}

	return false
}

// SetOwnerIdNo gets a reference to the given string and assigns it to the OwnerIdNo field.
func (o *CreateShopRequest) SetOwnerIdNo(v string) {
	o.OwnerIdNo = &v
}

// GetOwnerAddress returns the OwnerAddress field value if set, zero value otherwise.
func (o *CreateShopRequest) GetOwnerAddress() AddressInfo {
	if o == nil || utils.IsNil(o.OwnerAddress) {
		var ret AddressInfo
		return ret
	}
	return *o.OwnerAddress
}

// GetOwnerAddressOk returns a tuple with the OwnerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetOwnerAddressOk() (*AddressInfo, bool) {
	if o == nil || utils.IsNil(o.OwnerAddress) {
		return nil, false
	}
	return o.OwnerAddress, true
}

// HasOwnerAddress returns a boolean if a field has been set.
func (o *CreateShopRequest) HasOwnerAddress() bool {
	if o != nil && !utils.IsNil(o.OwnerAddress) {
		return true
	}

	return false
}

// SetOwnerAddress gets a reference to the given AddressInfo and assigns it to the OwnerAddress field.
func (o *CreateShopRequest) SetOwnerAddress(v AddressInfo) {
	o.OwnerAddress = &v
}

// GetOwnerPhoneNumber returns the OwnerPhoneNumber field value if set, zero value otherwise.
func (o *CreateShopRequest) GetOwnerPhoneNumber() MobileNoInfo {
	if o == nil || utils.IsNil(o.OwnerPhoneNumber) {
		var ret MobileNoInfo
		return ret
	}
	return *o.OwnerPhoneNumber
}

// GetOwnerPhoneNumberOk returns a tuple with the OwnerPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetOwnerPhoneNumberOk() (*MobileNoInfo, bool) {
	if o == nil || utils.IsNil(o.OwnerPhoneNumber) {
		return nil, false
	}
	return o.OwnerPhoneNumber, true
}

// HasOwnerPhoneNumber returns a boolean if a field has been set.
func (o *CreateShopRequest) HasOwnerPhoneNumber() bool {
	if o != nil && !utils.IsNil(o.OwnerPhoneNumber) {
		return true
	}

	return false
}

// SetOwnerPhoneNumber gets a reference to the given MobileNoInfo and assigns it to the OwnerPhoneNumber field.
func (o *CreateShopRequest) SetOwnerPhoneNumber(v MobileNoInfo) {
	o.OwnerPhoneNumber = &v
}

// GetShopOwning returns the ShopOwning field value if set, zero value otherwise.
func (o *CreateShopRequest) GetShopOwning() string {
	if o == nil || utils.IsNil(o.ShopOwning) {
		var ret string
		return ret
	}
	return *o.ShopOwning
}

// GetShopOwningOk returns a tuple with the ShopOwning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetShopOwningOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ShopOwning) {
		return nil, false
	}
	return o.ShopOwning, true
}

// HasShopOwning returns a boolean if a field has been set.
func (o *CreateShopRequest) HasShopOwning() bool {
	if o != nil && !utils.IsNil(o.ShopOwning) {
		return true
	}

	return false
}

// SetShopOwning gets a reference to the given string and assigns it to the ShopOwning field.
func (o *CreateShopRequest) SetShopOwning(v string) {
	o.ShopOwning = &v
}

// GetDeviceNumber returns the DeviceNumber field value if set, zero value otherwise.
func (o *CreateShopRequest) GetDeviceNumber() string {
	if o == nil || utils.IsNil(o.DeviceNumber) {
		var ret string
		return ret
	}
	return *o.DeviceNumber
}

// GetDeviceNumberOk returns a tuple with the DeviceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetDeviceNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DeviceNumber) {
		return nil, false
	}
	return o.DeviceNumber, true
}

// HasDeviceNumber returns a boolean if a field has been set.
func (o *CreateShopRequest) HasDeviceNumber() bool {
	if o != nil && !utils.IsNil(o.DeviceNumber) {
		return true
	}

	return false
}

// SetDeviceNumber gets a reference to the given string and assigns it to the DeviceNumber field.
func (o *CreateShopRequest) SetDeviceNumber(v string) {
	o.DeviceNumber = &v
}

// GetPosNumber returns the PosNumber field value if set, zero value otherwise.
func (o *CreateShopRequest) GetPosNumber() string {
	if o == nil || utils.IsNil(o.PosNumber) {
		var ret string
		return ret
	}
	return *o.PosNumber
}

// GetPosNumberOk returns a tuple with the PosNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetPosNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PosNumber) {
		return nil, false
	}
	return o.PosNumber, true
}

// HasPosNumber returns a boolean if a field has been set.
func (o *CreateShopRequest) HasPosNumber() bool {
	if o != nil && !utils.IsNil(o.PosNumber) {
		return true
	}

	return false
}

// SetPosNumber gets a reference to the given string and assigns it to the PosNumber field.
func (o *CreateShopRequest) SetPosNumber(v string) {
	o.PosNumber = &v
}

// GetDirectorPics returns the DirectorPics field value if set, zero value otherwise.
func (o *CreateShopRequest) GetDirectorPics() []PicInfo {
	if o == nil || utils.IsNil(o.DirectorPics) {
		var ret []PicInfo
		return ret
	}
	return o.DirectorPics
}

// GetDirectorPicsOk returns a tuple with the DirectorPics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetDirectorPicsOk() ([]PicInfo, bool) {
	if o == nil || utils.IsNil(o.DirectorPics) {
		return nil, false
	}
	return o.DirectorPics, true
}

// HasDirectorPics returns a boolean if a field has been set.
func (o *CreateShopRequest) HasDirectorPics() bool {
	if o != nil && !utils.IsNil(o.DirectorPics) {
		return true
	}

	return false
}

// SetDirectorPics gets a reference to the given []PicInfo and assigns it to the DirectorPics field.
func (o *CreateShopRequest) SetDirectorPics(v []PicInfo) {
	o.DirectorPics = v
}

// GetNonDirectorPics returns the NonDirectorPics field value if set, zero value otherwise.
func (o *CreateShopRequest) GetNonDirectorPics() []PicInfo {
	if o == nil || utils.IsNil(o.NonDirectorPics) {
		var ret []PicInfo
		return ret
	}
	return o.NonDirectorPics
}

// GetNonDirectorPicsOk returns a tuple with the NonDirectorPics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetNonDirectorPicsOk() ([]PicInfo, bool) {
	if o == nil || utils.IsNil(o.NonDirectorPics) {
		return nil, false
	}
	return o.NonDirectorPics, true
}

// HasNonDirectorPics returns a boolean if a field has been set.
func (o *CreateShopRequest) HasNonDirectorPics() bool {
	if o != nil && !utils.IsNil(o.NonDirectorPics) {
		return true
	}

	return false
}

// SetNonDirectorPics gets a reference to the given []PicInfo and assigns it to the NonDirectorPics field.
func (o *CreateShopRequest) SetNonDirectorPics(v []PicInfo) {
	o.NonDirectorPics = v
}

// GetShopBizType returns the ShopBizType field value if set, zero value otherwise.
func (o *CreateShopRequest) GetShopBizType() string {
	if o == nil || utils.IsNil(o.ShopBizType) {
		var ret string
		return ret
	}
	return *o.ShopBizType
}

// GetShopBizTypeOk returns a tuple with the ShopBizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetShopBizTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ShopBizType) {
		return nil, false
	}
	return o.ShopBizType, true
}

// HasShopBizType returns a boolean if a field has been set.
func (o *CreateShopRequest) HasShopBizType() bool {
	if o != nil && !utils.IsNil(o.ShopBizType) {
		return true
	}

	return false
}

// SetShopBizType gets a reference to the given string and assigns it to the ShopBizType field.
func (o *CreateShopRequest) SetShopBizType(v string) {
	o.ShopBizType = &v
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *CreateShopRequest) GetExtInfo() map[string]interface{} {
	if o == nil || utils.IsNil(o.ExtInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetExtInfoOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.ExtInfo) {
		return map[string]interface{}{}, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *CreateShopRequest) HasExtInfo() bool {
	if o != nil && !utils.IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given map[string]interface{} and assigns it to the ExtInfo field.
func (o *CreateShopRequest) SetExtInfo(v map[string]interface{}) {
	o.ExtInfo = v
}

func (o CreateShopRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateShopRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantId"] = o.MerchantId
	if !utils.IsNil(o.ParentDivisionId) {
		toSerialize["parentDivisionId"] = o.ParentDivisionId
	}
	toSerialize["shopParentType"] = o.ShopParentType
	toSerialize["mainName"] = o.MainName
	if !utils.IsNil(o.ShopAddress) {
		toSerialize["shopAddress"] = o.ShopAddress
	}
	if !utils.IsNil(o.ShopDesc) {
		toSerialize["shopDesc"] = o.ShopDesc
	}
	toSerialize["externalShopId"] = o.ExternalShopId
	if !utils.IsNil(o.LogoUrlMap) {
		toSerialize["logoUrlMap"] = o.LogoUrlMap
	}
	if !utils.IsNil(o.MccCodes) {
		toSerialize["mccCodes"] = o.MccCodes
	}
	toSerialize["sizeType"] = o.SizeType
	if !utils.IsNil(o.Ln) {
		toSerialize["ln"] = o.Ln
	}
	if !utils.IsNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !utils.IsNil(o.TaxNo) {
		toSerialize["taxNo"] = o.TaxNo
	}
	if !utils.IsNil(o.BrandName) {
		toSerialize["brandName"] = o.BrandName
	}
	if !utils.IsNil(o.TaxAddress) {
		toSerialize["taxAddress"] = o.TaxAddress
	}
	if !utils.IsNil(o.Loyalty) {
		toSerialize["loyalty"] = o.Loyalty
	}
	if !utils.IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !utils.IsNil(o.BusinessEntity) {
		toSerialize["businessEntity"] = o.BusinessEntity
	}
	if !utils.IsNil(o.BusinessDocs) {
		toSerialize["businessDocs"] = o.BusinessDocs
	}
	if !utils.IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !utils.IsNil(o.OwnerIdType) {
		toSerialize["ownerIdType"] = o.OwnerIdType
	}
	if !utils.IsNil(o.OwnerIdNo) {
		toSerialize["ownerIdNo"] = o.OwnerIdNo
	}
	if !utils.IsNil(o.OwnerAddress) {
		toSerialize["ownerAddress"] = o.OwnerAddress
	}
	if !utils.IsNil(o.OwnerPhoneNumber) {
		toSerialize["ownerPhoneNumber"] = o.OwnerPhoneNumber
	}
	if !utils.IsNil(o.ShopOwning) {
		toSerialize["shopOwning"] = o.ShopOwning
	}
	if !utils.IsNil(o.DeviceNumber) {
		toSerialize["deviceNumber"] = o.DeviceNumber
	}
	if !utils.IsNil(o.PosNumber) {
		toSerialize["posNumber"] = o.PosNumber
	}
	if !utils.IsNil(o.DirectorPics) {
		toSerialize["directorPics"] = o.DirectorPics
	}
	if !utils.IsNil(o.NonDirectorPics) {
		toSerialize["nonDirectorPics"] = o.NonDirectorPics
	}
	if !utils.IsNil(o.ShopBizType) {
		toSerialize["shopBizType"] = o.ShopBizType
	}
	if !utils.IsNil(o.ExtInfo) {
		toSerialize["extInfo"] = o.ExtInfo
	}
	return toSerialize, nil
}

func (o *CreateShopRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"merchantId",
		"shopParentType",
		"mainName",
		"externalShopId",
		"sizeType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateShopRequest := _CreateShopRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateShopRequest)

	if err != nil {
		return err
	}

	*o = CreateShopRequest(varCreateShopRequest)

	return err
}

type NullableCreateShopRequest struct {
	value *CreateShopRequest
	isSet bool
}

func (v NullableCreateShopRequest) Get() *CreateShopRequest {
	return v.value
}

func (v *NullableCreateShopRequest) Set(val *CreateShopRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShopRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShopRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShopRequest(val *CreateShopRequest) *NullableCreateShopRequest {
	return &NullableCreateShopRequest{value: val, isSet: true}
}

func (v NullableCreateShopRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShopRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


