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

// checks if the TransferToBankRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransferToBankRequest{}

// TransferToBankRequest struct for TransferToBankRequest
type TransferToBankRequest struct {
	// Unique transaction identifier on partner system which assigned to each transaction<br> Notes:<br> If the partner receives a timeout or an unexpected response from DANA and partner expects to perform retry request to DANA, please use the partnerReferenceNo that is the same as the one used in the transaction request process before 
	PartnerReferenceNo *string `json:"partnerReferenceNo,omitempty"`
	// Customer account number, in format 628xxx
	CustomerNumber string `json:"customerNumber"`
	// Customer account type
	AccountType *string `json:"accountType,omitempty"`
	// Beneficiary account number
	BeneficiaryAccountNumber string `json:"beneficiaryAccountNumber"`
	// Beneficiary Bank code
	BeneficiaryBankCode string `json:"beneficiaryBankCode"`
	// Amount. Contains two sub-fields:<br> 1. Value: Transaction amount, including the cents<br> 2. Currency: Currency code based on ISO 
	Amount Money `json:"amount"`
	AdditionalInfo TransferToBankRequestAdditionalInfo `json:"additionalInfo"`
}

type _TransferToBankRequest TransferToBankRequest

// NewTransferToBankRequest instantiates a new TransferToBankRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferToBankRequest(customerNumber string, beneficiaryAccountNumber string, beneficiaryBankCode string, amount Money, additionalInfo TransferToBankRequestAdditionalInfo) *TransferToBankRequest {
	this := TransferToBankRequest{}
	this.CustomerNumber = customerNumber
	this.BeneficiaryAccountNumber = beneficiaryAccountNumber
	this.BeneficiaryBankCode = beneficiaryBankCode
	this.Amount = amount
	this.AdditionalInfo = additionalInfo
	return &this
}

// NewTransferToBankRequestWithDefaults instantiates a new TransferToBankRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferToBankRequestWithDefaults() *TransferToBankRequest {
	this := TransferToBankRequest{}
	return &this
}

// GetPartnerReferenceNo returns the PartnerReferenceNo field value if set, zero value otherwise.
func (o *TransferToBankRequest) GetPartnerReferenceNo() string {
	if o == nil || utils.IsNil(o.PartnerReferenceNo) {
		var ret string
		return ret
	}
	return *o.PartnerReferenceNo
}

// GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferToBankRequest) GetPartnerReferenceNoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PartnerReferenceNo) {
		return nil, false
	}
	return o.PartnerReferenceNo, true
}

// HasPartnerReferenceNo returns a boolean if a field has been set.
func (o *TransferToBankRequest) HasPartnerReferenceNo() bool {
	if o != nil && !utils.IsNil(o.PartnerReferenceNo) {
		return true
	}

	return false
}

// SetPartnerReferenceNo gets a reference to the given string and assigns it to the PartnerReferenceNo field.
func (o *TransferToBankRequest) SetPartnerReferenceNo(v string) {
	o.PartnerReferenceNo = &v
}

// GetCustomerNumber returns the CustomerNumber field value
func (o *TransferToBankRequest) GetCustomerNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerNumber
}

// GetCustomerNumberOk returns a tuple with the CustomerNumber field value
// and a boolean to check if the value has been set.
func (o *TransferToBankRequest) GetCustomerNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerNumber, true
}

// SetCustomerNumber sets field value
func (o *TransferToBankRequest) SetCustomerNumber(v string) {
	o.CustomerNumber = v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *TransferToBankRequest) GetAccountType() string {
	if o == nil || utils.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferToBankRequest) GetAccountTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *TransferToBankRequest) HasAccountType() bool {
	if o != nil && !utils.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *TransferToBankRequest) SetAccountType(v string) {
	o.AccountType = &v
}

// GetBeneficiaryAccountNumber returns the BeneficiaryAccountNumber field value
func (o *TransferToBankRequest) GetBeneficiaryAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeneficiaryAccountNumber
}

// GetBeneficiaryAccountNumberOk returns a tuple with the BeneficiaryAccountNumber field value
// and a boolean to check if the value has been set.
func (o *TransferToBankRequest) GetBeneficiaryAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BeneficiaryAccountNumber, true
}

// SetBeneficiaryAccountNumber sets field value
func (o *TransferToBankRequest) SetBeneficiaryAccountNumber(v string) {
	o.BeneficiaryAccountNumber = v
}

// GetBeneficiaryBankCode returns the BeneficiaryBankCode field value
func (o *TransferToBankRequest) GetBeneficiaryBankCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeneficiaryBankCode
}

// GetBeneficiaryBankCodeOk returns a tuple with the BeneficiaryBankCode field value
// and a boolean to check if the value has been set.
func (o *TransferToBankRequest) GetBeneficiaryBankCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BeneficiaryBankCode, true
}

// SetBeneficiaryBankCode sets field value
func (o *TransferToBankRequest) SetBeneficiaryBankCode(v string) {
	o.BeneficiaryBankCode = v
}

// GetAmount returns the Amount field value
func (o *TransferToBankRequest) GetAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferToBankRequest) GetAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferToBankRequest) SetAmount(v Money) {
	o.Amount = v
}

// GetAdditionalInfo returns the AdditionalInfo field value
func (o *TransferToBankRequest) GetAdditionalInfo() TransferToBankRequestAdditionalInfo {
	if o == nil {
		var ret TransferToBankRequestAdditionalInfo
		return ret
	}

	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value
// and a boolean to check if the value has been set.
func (o *TransferToBankRequest) GetAdditionalInfoOk() (*TransferToBankRequestAdditionalInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionalInfo, true
}

// SetAdditionalInfo sets field value
func (o *TransferToBankRequest) SetAdditionalInfo(v TransferToBankRequestAdditionalInfo) {
	o.AdditionalInfo = v
}

func (o TransferToBankRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferToBankRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PartnerReferenceNo) {
		toSerialize["partnerReferenceNo"] = o.PartnerReferenceNo
	}
	toSerialize["customerNumber"] = o.CustomerNumber
	if !utils.IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	toSerialize["beneficiaryAccountNumber"] = o.BeneficiaryAccountNumber
	toSerialize["beneficiaryBankCode"] = o.BeneficiaryBankCode
	toSerialize["amount"] = o.Amount
	toSerialize["additionalInfo"] = o.AdditionalInfo
	return toSerialize, nil
}

func (o *TransferToBankRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customerNumber",
		"beneficiaryAccountNumber",
		"beneficiaryBankCode",
		"amount",
		"additionalInfo",
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

	varTransferToBankRequest := _TransferToBankRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransferToBankRequest)

	if err != nil {
		return err
	}

	*o = TransferToBankRequest(varTransferToBankRequest)

	return err
}

type NullableTransferToBankRequest struct {
	value *TransferToBankRequest
	isSet bool
}

func (v NullableTransferToBankRequest) Get() *TransferToBankRequest {
	return v.value
}

func (v *NullableTransferToBankRequest) Set(val *TransferToBankRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferToBankRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferToBankRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferToBankRequest(val *TransferToBankRequest) *NullableTransferToBankRequest {
	return &NullableTransferToBankRequest{value: val, isSet: true}
}

func (v NullableTransferToBankRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferToBankRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


