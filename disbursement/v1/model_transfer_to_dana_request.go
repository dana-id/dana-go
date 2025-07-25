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

// checks if the TransferToDanaRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransferToDanaRequest{}

// TransferToDanaRequest struct for TransferToDanaRequest
type TransferToDanaRequest struct {
	// Unique transaction identifier on partner system which assigned to each transaction<br> Notes:<br> If the partner receives a timeout or an unexpected response from DANA and partner expects to perform retry request to DANA, please use the partnerReferenceNo that is the same as the one used in the transaction request process before 
	PartnerReferenceNo string `json:"partnerReferenceNo"`
	// Customer account number, in format 628xxx
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// Amount. Contains two sub-fields:<br> 1. Value: Transaction amount, including the cents<br> 2. Currency: Currency code based on ISO 
	Amount Money `json:"amount"`
	// Fee amount. Contains two sub-fields:<br> 1. Value: Amount, including the cents<br> 2. Currency: Currency code based on ISO 
	FeeAmount Money `json:"feeAmount"`
	// Transaction date, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time)
	TransactionDate *string `json:"transactionDate,omitempty" validate:"regexp=^\\\\d{4}-\\\\d{2}-\\\\d{2}T\\\\d{2}:\\\\d{2}:\\\\d{2}\\\\+07:00$"`
	// Session identifier
	SessionId *string `json:"sessionId,omitempty"`
	// Category identifier
	CategoryId *float32 `json:"categoryId,omitempty"`
	// Transaction notes
	Notes *string `json:"notes,omitempty"`
	AdditionalInfo TransferToDanaRequestAdditionalInfo `json:"additionalInfo"`
}

type _TransferToDanaRequest TransferToDanaRequest

// NewTransferToDanaRequest instantiates a new TransferToDanaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferToDanaRequest(partnerReferenceNo string, amount Money, feeAmount Money, additionalInfo TransferToDanaRequestAdditionalInfo) *TransferToDanaRequest {
	this := TransferToDanaRequest{}
	this.PartnerReferenceNo = partnerReferenceNo
	this.Amount = amount
	this.FeeAmount = feeAmount
	this.AdditionalInfo = additionalInfo
	return &this
}

// NewTransferToDanaRequestWithDefaults instantiates a new TransferToDanaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferToDanaRequestWithDefaults() *TransferToDanaRequest {
	this := TransferToDanaRequest{}
	return &this
}

// GetPartnerReferenceNo returns the PartnerReferenceNo field value
func (o *TransferToDanaRequest) GetPartnerReferenceNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartnerReferenceNo
}

// GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field value
// and a boolean to check if the value has been set.
func (o *TransferToDanaRequest) GetPartnerReferenceNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerReferenceNo, true
}

// SetPartnerReferenceNo sets field value
func (o *TransferToDanaRequest) SetPartnerReferenceNo(v string) {
	o.PartnerReferenceNo = v
}

// GetCustomerNumber returns the CustomerNumber field value if set, zero value otherwise.
func (o *TransferToDanaRequest) GetCustomerNumber() string {
	if o == nil || utils.IsNil(o.CustomerNumber) {
		var ret string
		return ret
	}
	return *o.CustomerNumber
}

// GetCustomerNumberOk returns a tuple with the CustomerNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferToDanaRequest) GetCustomerNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CustomerNumber) {
		return nil, false
	}
	return o.CustomerNumber, true
}

// HasCustomerNumber returns a boolean if a field has been set.
func (o *TransferToDanaRequest) HasCustomerNumber() bool {
	if o != nil && !utils.IsNil(o.CustomerNumber) {
		return true
	}

	return false
}

// SetCustomerNumber gets a reference to the given string and assigns it to the CustomerNumber field.
func (o *TransferToDanaRequest) SetCustomerNumber(v string) {
	o.CustomerNumber = &v
}

// GetAmount returns the Amount field value
func (o *TransferToDanaRequest) GetAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferToDanaRequest) GetAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferToDanaRequest) SetAmount(v Money) {
	o.Amount = v
}

// GetFeeAmount returns the FeeAmount field value
func (o *TransferToDanaRequest) GetFeeAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value
// and a boolean to check if the value has been set.
func (o *TransferToDanaRequest) GetFeeAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeAmount, true
}

// SetFeeAmount sets field value
func (o *TransferToDanaRequest) SetFeeAmount(v Money) {
	o.FeeAmount = v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *TransferToDanaRequest) GetTransactionDate() string {
	if o == nil || utils.IsNil(o.TransactionDate) {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferToDanaRequest) GetTransactionDateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TransactionDate) {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *TransferToDanaRequest) HasTransactionDate() bool {
	if o != nil && !utils.IsNil(o.TransactionDate) {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *TransferToDanaRequest) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *TransferToDanaRequest) GetSessionId() string {
	if o == nil || utils.IsNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferToDanaRequest) GetSessionIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *TransferToDanaRequest) HasSessionId() bool {
	if o != nil && !utils.IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *TransferToDanaRequest) SetSessionId(v string) {
	o.SessionId = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *TransferToDanaRequest) GetCategoryId() float32 {
	if o == nil || utils.IsNil(o.CategoryId) {
		var ret float32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferToDanaRequest) GetCategoryIdOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *TransferToDanaRequest) HasCategoryId() bool {
	if o != nil && !utils.IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given float32 and assigns it to the CategoryId field.
func (o *TransferToDanaRequest) SetCategoryId(v float32) {
	o.CategoryId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *TransferToDanaRequest) GetNotes() string {
	if o == nil || utils.IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferToDanaRequest) GetNotesOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *TransferToDanaRequest) HasNotes() bool {
	if o != nil && !utils.IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *TransferToDanaRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value
func (o *TransferToDanaRequest) GetAdditionalInfo() TransferToDanaRequestAdditionalInfo {
	if o == nil {
		var ret TransferToDanaRequestAdditionalInfo
		return ret
	}

	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value
// and a boolean to check if the value has been set.
func (o *TransferToDanaRequest) GetAdditionalInfoOk() (*TransferToDanaRequestAdditionalInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionalInfo, true
}

// SetAdditionalInfo sets field value
func (o *TransferToDanaRequest) SetAdditionalInfo(v TransferToDanaRequestAdditionalInfo) {
	o.AdditionalInfo = v
}

func (o TransferToDanaRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferToDanaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["partnerReferenceNo"] = o.PartnerReferenceNo
	if !utils.IsNil(o.CustomerNumber) {
		toSerialize["customerNumber"] = o.CustomerNumber
	}
	toSerialize["amount"] = o.Amount
	toSerialize["feeAmount"] = o.FeeAmount
	if !utils.IsNil(o.TransactionDate) {
		toSerialize["transactionDate"] = o.TransactionDate
	}
	if !utils.IsNil(o.SessionId) {
		toSerialize["sessionId"] = o.SessionId
	}
	if !utils.IsNil(o.CategoryId) {
		toSerialize["categoryId"] = o.CategoryId
	}
	if !utils.IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["additionalInfo"] = o.AdditionalInfo
	return toSerialize, nil
}

func (o *TransferToDanaRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"partnerReferenceNo",
		"amount",
		"feeAmount",
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

	varTransferToDanaRequest := _TransferToDanaRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransferToDanaRequest)

	if err != nil {
		return err
	}

	*o = TransferToDanaRequest(varTransferToDanaRequest)

	return err
}

type NullableTransferToDanaRequest struct {
	value *TransferToDanaRequest
	isSet bool
}

func (v NullableTransferToDanaRequest) Get() *TransferToDanaRequest {
	return v.value
}

func (v *NullableTransferToDanaRequest) Set(val *TransferToDanaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferToDanaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferToDanaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferToDanaRequest(val *TransferToDanaRequest) *NullableTransferToDanaRequest {
	return &NullableTransferToDanaRequest{value: val, isSet: true}
}

func (v NullableTransferToDanaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferToDanaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


