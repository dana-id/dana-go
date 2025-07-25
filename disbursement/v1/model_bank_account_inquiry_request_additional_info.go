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

package disbursement

/*
Disbursement API

API for doing disbursement operations in DANA, including DANA account inquiry, transfer to DANA, and transfer to bank disbursement

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.


import (
	"encoding/json"
	"bytes"
	"fmt"
	utils "github.com/dana-id/dana-go/utils"
)

// checks if the BankAccountInquiryRequestAdditionalInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BankAccountInquiryRequestAdditionalInfo{}

// BankAccountInquiryRequestAdditionalInfo Additional information
type BankAccountInquiryRequestAdditionalInfo struct {
	// Additional information of withdraw fund type, i.e.<br> MERCHANT_WITHDRAW_FOR_CORPORATE 
	FundType string `json:"fundType"`
	// Additional information of external division identifier. (fundType: MERCHANT_WITHDRAW_FOR_CORPORATE)<br> Notes: The required of this parameter is Optional, but if \"additionalInfo.chargeTarget\" has value DIVISION then the required of this parameter will be changed to Mandatory 
	ExternalDivisionId *string `json:"externalDivisionId,omitempty"`
	// Additional information of charge target. The values are:<br> • null<br> • DIVISION<br> • MERCHANT<br> Notes: If the value is DIVISION, externalDivisionId will be Mandatory 
	ChargeTarget *string `json:"chargeTarget,omitempty"`
	// Additional information of beneficiary Bank code
	BeneficiaryBankCode string `json:"beneficiaryBankCode"`
	// Additional information of beneficiary account name for validation purpose
	BeneficiaryAccountName *string `json:"beneficiaryAccountName,omitempty"`
	// Additional information of account type
	AccountType *string `json:"accountType,omitempty"`
	// Contains customer token, which has been obtained from binding process, refer to Account Binding & Unbinding documentation<br> If request is coming from user interaction, this field is mandatory. If not, just filled customerNumber 
	AccessToken *string `json:"accessToken,omitempty"`
}

type _BankAccountInquiryRequestAdditionalInfo BankAccountInquiryRequestAdditionalInfo

// NewBankAccountInquiryRequestAdditionalInfo instantiates a new BankAccountInquiryRequestAdditionalInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccountInquiryRequestAdditionalInfo(fundType string, beneficiaryBankCode string) *BankAccountInquiryRequestAdditionalInfo {
	this := BankAccountInquiryRequestAdditionalInfo{}
	this.FundType = fundType
	this.BeneficiaryBankCode = beneficiaryBankCode
	return &this
}

// NewBankAccountInquiryRequestAdditionalInfoWithDefaults instantiates a new BankAccountInquiryRequestAdditionalInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountInquiryRequestAdditionalInfoWithDefaults() *BankAccountInquiryRequestAdditionalInfo {
	this := BankAccountInquiryRequestAdditionalInfo{}
	return &this
}

// GetFundType returns the FundType field value
func (o *BankAccountInquiryRequestAdditionalInfo) GetFundType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundType
}

// GetFundTypeOk returns a tuple with the FundType field value
// and a boolean to check if the value has been set.
func (o *BankAccountInquiryRequestAdditionalInfo) GetFundTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundType, true
}

// SetFundType sets field value
func (o *BankAccountInquiryRequestAdditionalInfo) SetFundType(v string) {
	o.FundType = v
}

// GetExternalDivisionId returns the ExternalDivisionId field value if set, zero value otherwise.
func (o *BankAccountInquiryRequestAdditionalInfo) GetExternalDivisionId() string {
	if o == nil || utils.IsNil(o.ExternalDivisionId) {
		var ret string
		return ret
	}
	return *o.ExternalDivisionId
}

// GetExternalDivisionIdOk returns a tuple with the ExternalDivisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountInquiryRequestAdditionalInfo) GetExternalDivisionIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExternalDivisionId) {
		return nil, false
	}
	return o.ExternalDivisionId, true
}

// HasExternalDivisionId returns a boolean if a field has been set.
func (o *BankAccountInquiryRequestAdditionalInfo) HasExternalDivisionId() bool {
	if o != nil && !utils.IsNil(o.ExternalDivisionId) {
		return true
	}

	return false
}

// SetExternalDivisionId gets a reference to the given string and assigns it to the ExternalDivisionId field.
func (o *BankAccountInquiryRequestAdditionalInfo) SetExternalDivisionId(v string) {
	o.ExternalDivisionId = &v
}

// GetChargeTarget returns the ChargeTarget field value if set, zero value otherwise.
func (o *BankAccountInquiryRequestAdditionalInfo) GetChargeTarget() string {
	if o == nil || utils.IsNil(o.ChargeTarget) {
		var ret string
		return ret
	}
	return *o.ChargeTarget
}

// GetChargeTargetOk returns a tuple with the ChargeTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountInquiryRequestAdditionalInfo) GetChargeTargetOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ChargeTarget) {
		return nil, false
	}
	return o.ChargeTarget, true
}

// HasChargeTarget returns a boolean if a field has been set.
func (o *BankAccountInquiryRequestAdditionalInfo) HasChargeTarget() bool {
	if o != nil && !utils.IsNil(o.ChargeTarget) {
		return true
	}

	return false
}

// SetChargeTarget gets a reference to the given string and assigns it to the ChargeTarget field.
func (o *BankAccountInquiryRequestAdditionalInfo) SetChargeTarget(v string) {
	o.ChargeTarget = &v
}

// GetBeneficiaryBankCode returns the BeneficiaryBankCode field value
func (o *BankAccountInquiryRequestAdditionalInfo) GetBeneficiaryBankCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeneficiaryBankCode
}

// GetBeneficiaryBankCodeOk returns a tuple with the BeneficiaryBankCode field value
// and a boolean to check if the value has been set.
func (o *BankAccountInquiryRequestAdditionalInfo) GetBeneficiaryBankCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BeneficiaryBankCode, true
}

// SetBeneficiaryBankCode sets field value
func (o *BankAccountInquiryRequestAdditionalInfo) SetBeneficiaryBankCode(v string) {
	o.BeneficiaryBankCode = v
}

// GetBeneficiaryAccountName returns the BeneficiaryAccountName field value if set, zero value otherwise.
func (o *BankAccountInquiryRequestAdditionalInfo) GetBeneficiaryAccountName() string {
	if o == nil || utils.IsNil(o.BeneficiaryAccountName) {
		var ret string
		return ret
	}
	return *o.BeneficiaryAccountName
}

// GetBeneficiaryAccountNameOk returns a tuple with the BeneficiaryAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountInquiryRequestAdditionalInfo) GetBeneficiaryAccountNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BeneficiaryAccountName) {
		return nil, false
	}
	return o.BeneficiaryAccountName, true
}

// HasBeneficiaryAccountName returns a boolean if a field has been set.
func (o *BankAccountInquiryRequestAdditionalInfo) HasBeneficiaryAccountName() bool {
	if o != nil && !utils.IsNil(o.BeneficiaryAccountName) {
		return true
	}

	return false
}

// SetBeneficiaryAccountName gets a reference to the given string and assigns it to the BeneficiaryAccountName field.
func (o *BankAccountInquiryRequestAdditionalInfo) SetBeneficiaryAccountName(v string) {
	o.BeneficiaryAccountName = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *BankAccountInquiryRequestAdditionalInfo) GetAccountType() string {
	if o == nil || utils.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountInquiryRequestAdditionalInfo) GetAccountTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *BankAccountInquiryRequestAdditionalInfo) HasAccountType() bool {
	if o != nil && !utils.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *BankAccountInquiryRequestAdditionalInfo) SetAccountType(v string) {
	o.AccountType = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *BankAccountInquiryRequestAdditionalInfo) GetAccessToken() string {
	if o == nil || utils.IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountInquiryRequestAdditionalInfo) GetAccessTokenOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *BankAccountInquiryRequestAdditionalInfo) HasAccessToken() bool {
	if o != nil && !utils.IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *BankAccountInquiryRequestAdditionalInfo) SetAccessToken(v string) {
	o.AccessToken = &v
}

func (o BankAccountInquiryRequestAdditionalInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankAccountInquiryRequestAdditionalInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fundType"] = o.FundType
	if !utils.IsNil(o.ExternalDivisionId) {
		toSerialize["externalDivisionId"] = o.ExternalDivisionId
	}
	if !utils.IsNil(o.ChargeTarget) {
		toSerialize["chargeTarget"] = o.ChargeTarget
	}
	toSerialize["beneficiaryBankCode"] = o.BeneficiaryBankCode
	if !utils.IsNil(o.BeneficiaryAccountName) {
		toSerialize["beneficiaryAccountName"] = o.BeneficiaryAccountName
	}
	if !utils.IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !utils.IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	return toSerialize, nil
}

func (o *BankAccountInquiryRequestAdditionalInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fundType",
		"beneficiaryBankCode",
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

	varBankAccountInquiryRequestAdditionalInfo := _BankAccountInquiryRequestAdditionalInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBankAccountInquiryRequestAdditionalInfo)

	if err != nil {
		return err
	}

	*o = BankAccountInquiryRequestAdditionalInfo(varBankAccountInquiryRequestAdditionalInfo)

	return err
}

type NullableBankAccountInquiryRequestAdditionalInfo struct {
	value *BankAccountInquiryRequestAdditionalInfo
	isSet bool
}

func (v NullableBankAccountInquiryRequestAdditionalInfo) Get() *BankAccountInquiryRequestAdditionalInfo {
	return v.value
}

func (v *NullableBankAccountInquiryRequestAdditionalInfo) Set(val *BankAccountInquiryRequestAdditionalInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccountInquiryRequestAdditionalInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccountInquiryRequestAdditionalInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccountInquiryRequestAdditionalInfo(val *BankAccountInquiryRequestAdditionalInfo) *NullableBankAccountInquiryRequestAdditionalInfo {
	return &NullableBankAccountInquiryRequestAdditionalInfo{value: val, isSet: true}
}

func (v NullableBankAccountInquiryRequestAdditionalInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccountInquiryRequestAdditionalInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


