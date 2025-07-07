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

package webhook

import (
	"encoding/json"
	"bytes"
	"fmt"
	utils "github.com/dana-id/dana-go/utils"
)

// checks if the FinishNotifyRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FinishNotifyRequest{}

// FinishNotifyRequest struct for FinishNotifyRequest
type FinishNotifyRequest struct {
	// Original transaction identifier on DANA system
	OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
	// Original transaction identifier on partner system
	OriginalReferenceNo string `json:"originalReferenceNo"`
	// Original external identifier on header message
	OriginalExternalId *string `json:"originalExternalId,omitempty"`
	// Unique identifier per each merchant
	MerchantId string `json:"merchantId"`
	// Information of sub merchant identifier
	SubMerchantId *string `json:"subMerchantId,omitempty"`
	// Amount. Contains two sub-fields:<br> 1. Value: Transaction amount, including the cents<br> 2. Currency: Currency code based on ISO<br> 
	Amount Money `json:"amount"`
	// Transaction status code:<br> - 00 = Success, the order has been paid<br> - 05 = Cancelled, the order has been closed because it is expired<br> 
	LatestTransactionStatus string `json:"latestTransactionStatus"`
	// Description of transaction status
	TransactionStatusDesc *string `json:"transactionStatusDesc,omitempty"`
	// Transaction created time, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time)
	CreatedTime string `json:"createdTime" validate:"regexp=^\\\\d{4}-\\\\d{2}-\\\\d{2}T\\\\d{2}:\\\\d{2}:\\\\d{2}\\\\+07:00$"`
	// Transaction finished time, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time)
	FinishedTime string `json:"finishedTime" validate:"regexp=^\\\\d{4}-\\\\d{2}-\\\\d{2}T\\\\d{2}:\\\\d{2}:\\\\d{2}\\\\+07:00$"`
	// Store identifier to indicate to which store this payment belongs to
	ExternalStoreId *string `json:"externalStoreId,omitempty"`
	// Additional information
	AdditionalInfo *FinishNotifyRequestAdditionalInfo `json:"additionalInfo,omitempty"`
}

type _FinishNotifyRequest FinishNotifyRequest

// NewFinishNotifyRequest instantiates a new FinishNotifyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinishNotifyRequest(originalPartnerReferenceNo string, originalReferenceNo string, merchantId string, amount Money, latestTransactionStatus string, createdTime string, finishedTime string) *FinishNotifyRequest {
	this := FinishNotifyRequest{}
	this.OriginalPartnerReferenceNo = originalPartnerReferenceNo
	this.OriginalReferenceNo = originalReferenceNo
	this.MerchantId = merchantId
	this.Amount = amount
	this.LatestTransactionStatus = latestTransactionStatus
	this.CreatedTime = createdTime
	this.FinishedTime = finishedTime
	return &this
}

// NewFinishNotifyRequestWithDefaults instantiates a new FinishNotifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinishNotifyRequestWithDefaults() *FinishNotifyRequest {
	this := FinishNotifyRequest{}
	return &this
}

// GetOriginalPartnerReferenceNo returns the OriginalPartnerReferenceNo field value
func (o *FinishNotifyRequest) GetOriginalPartnerReferenceNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalPartnerReferenceNo
}

// GetOriginalPartnerReferenceNoOk returns a tuple with the OriginalPartnerReferenceNo field value
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequest) GetOriginalPartnerReferenceNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalPartnerReferenceNo, true
}

// SetOriginalPartnerReferenceNo sets field value
func (o *FinishNotifyRequest) SetOriginalPartnerReferenceNo(v string) {
	o.OriginalPartnerReferenceNo = v
}

// GetOriginalReferenceNo returns the OriginalReferenceNo field value
func (o *FinishNotifyRequest) GetOriginalReferenceNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalReferenceNo
}

// GetOriginalReferenceNoOk returns a tuple with the OriginalReferenceNo field value
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequest) GetOriginalReferenceNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalReferenceNo, true
}

// SetOriginalReferenceNo sets field value
func (o *FinishNotifyRequest) SetOriginalReferenceNo(v string) {
	o.OriginalReferenceNo = v
}

// GetOriginalExternalId returns the OriginalExternalId field value if set, zero value otherwise.
func (o *FinishNotifyRequest) GetOriginalExternalId() string {
	if o == nil || utils.IsNil(o.OriginalExternalId) {
		var ret string
		return ret
	}
	return *o.OriginalExternalId
}

// GetOriginalExternalIdOk returns a tuple with the OriginalExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequest) GetOriginalExternalIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OriginalExternalId) {
		return nil, false
	}
	return o.OriginalExternalId, true
}

// HasOriginalExternalId returns a boolean if a field has been set.
func (o *FinishNotifyRequest) HasOriginalExternalId() bool {
	if o != nil && !utils.IsNil(o.OriginalExternalId) {
		return true
	}

	return false
}

// SetOriginalExternalId gets a reference to the given string and assigns it to the OriginalExternalId field.
func (o *FinishNotifyRequest) SetOriginalExternalId(v string) {
	o.OriginalExternalId = &v
}

// GetMerchantId returns the MerchantId field value
func (o *FinishNotifyRequest) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequest) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *FinishNotifyRequest) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetSubMerchantId returns the SubMerchantId field value if set, zero value otherwise.
func (o *FinishNotifyRequest) GetSubMerchantId() string {
	if o == nil || utils.IsNil(o.SubMerchantId) {
		var ret string
		return ret
	}
	return *o.SubMerchantId
}

// GetSubMerchantIdOk returns a tuple with the SubMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequest) GetSubMerchantIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SubMerchantId) {
		return nil, false
	}
	return o.SubMerchantId, true
}

// HasSubMerchantId returns a boolean if a field has been set.
func (o *FinishNotifyRequest) HasSubMerchantId() bool {
	if o != nil && !utils.IsNil(o.SubMerchantId) {
		return true
	}

	return false
}

// SetSubMerchantId gets a reference to the given string and assigns it to the SubMerchantId field.
func (o *FinishNotifyRequest) SetSubMerchantId(v string) {
	o.SubMerchantId = &v
}

// GetAmount returns the Amount field value
func (o *FinishNotifyRequest) GetAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequest) GetAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *FinishNotifyRequest) SetAmount(v Money) {
	o.Amount = v
}

// GetLatestTransactionStatus returns the LatestTransactionStatus field value
func (o *FinishNotifyRequest) GetLatestTransactionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LatestTransactionStatus
}

// GetLatestTransactionStatusOk returns a tuple with the LatestTransactionStatus field value
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequest) GetLatestTransactionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LatestTransactionStatus, true
}

// SetLatestTransactionStatus sets field value
func (o *FinishNotifyRequest) SetLatestTransactionStatus(v string) {
	o.LatestTransactionStatus = v
}

// GetTransactionStatusDesc returns the TransactionStatusDesc field value if set, zero value otherwise.
func (o *FinishNotifyRequest) GetTransactionStatusDesc() string {
	if o == nil || utils.IsNil(o.TransactionStatusDesc) {
		var ret string
		return ret
	}
	return *o.TransactionStatusDesc
}

// GetTransactionStatusDescOk returns a tuple with the TransactionStatusDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequest) GetTransactionStatusDescOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TransactionStatusDesc) {
		return nil, false
	}
	return o.TransactionStatusDesc, true
}

// HasTransactionStatusDesc returns a boolean if a field has been set.
func (o *FinishNotifyRequest) HasTransactionStatusDesc() bool {
	if o != nil && !utils.IsNil(o.TransactionStatusDesc) {
		return true
	}

	return false
}

// SetTransactionStatusDesc gets a reference to the given string and assigns it to the TransactionStatusDesc field.
func (o *FinishNotifyRequest) SetTransactionStatusDesc(v string) {
	o.TransactionStatusDesc = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *FinishNotifyRequest) GetCreatedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequest) GetCreatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *FinishNotifyRequest) SetCreatedTime(v string) {
	o.CreatedTime = v
}

// GetFinishedTime returns the FinishedTime field value
func (o *FinishNotifyRequest) GetFinishedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FinishedTime
}

// GetFinishedTimeOk returns a tuple with the FinishedTime field value
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequest) GetFinishedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishedTime, true
}

// SetFinishedTime sets field value
func (o *FinishNotifyRequest) SetFinishedTime(v string) {
	o.FinishedTime = v
}

// GetExternalStoreId returns the ExternalStoreId field value if set, zero value otherwise.
func (o *FinishNotifyRequest) GetExternalStoreId() string {
	if o == nil || utils.IsNil(o.ExternalStoreId) {
		var ret string
		return ret
	}
	return *o.ExternalStoreId
}

// GetExternalStoreIdOk returns a tuple with the ExternalStoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequest) GetExternalStoreIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExternalStoreId) {
		return nil, false
	}
	return o.ExternalStoreId, true
}

// HasExternalStoreId returns a boolean if a field has been set.
func (o *FinishNotifyRequest) HasExternalStoreId() bool {
	if o != nil && !utils.IsNil(o.ExternalStoreId) {
		return true
	}

	return false
}

// SetExternalStoreId gets a reference to the given string and assigns it to the ExternalStoreId field.
func (o *FinishNotifyRequest) SetExternalStoreId(v string) {
	o.ExternalStoreId = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *FinishNotifyRequest) GetAdditionalInfo() FinishNotifyRequestAdditionalInfo {
	if o == nil || utils.IsNil(o.AdditionalInfo) {
		var ret FinishNotifyRequestAdditionalInfo
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinishNotifyRequest) GetAdditionalInfoOk() (*FinishNotifyRequestAdditionalInfo, bool) {
	if o == nil || utils.IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *FinishNotifyRequest) HasAdditionalInfo() bool {
	if o != nil && !utils.IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given FinishNotifyRequestAdditionalInfo and assigns it to the AdditionalInfo field.
func (o *FinishNotifyRequest) SetAdditionalInfo(v FinishNotifyRequestAdditionalInfo) {
	o.AdditionalInfo = &v
}

func (o FinishNotifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinishNotifyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["originalPartnerReferenceNo"] = o.OriginalPartnerReferenceNo
	toSerialize["originalReferenceNo"] = o.OriginalReferenceNo
	if !utils.IsNil(o.OriginalExternalId) {
		toSerialize["originalExternalId"] = o.OriginalExternalId
	}
	toSerialize["merchantId"] = o.MerchantId
	if !utils.IsNil(o.SubMerchantId) {
		toSerialize["subMerchantId"] = o.SubMerchantId
	}
	toSerialize["amount"] = o.Amount
	toSerialize["latestTransactionStatus"] = o.LatestTransactionStatus
	if !utils.IsNil(o.TransactionStatusDesc) {
		toSerialize["transactionStatusDesc"] = o.TransactionStatusDesc
	}
	toSerialize["createdTime"] = o.CreatedTime
	toSerialize["finishedTime"] = o.FinishedTime
	if !utils.IsNil(o.ExternalStoreId) {
		toSerialize["externalStoreId"] = o.ExternalStoreId
	}
	if !utils.IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	return toSerialize, nil
}

func (o *FinishNotifyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"originalPartnerReferenceNo",
		"originalReferenceNo",
		"merchantId",
		"amount",
		"latestTransactionStatus",
		"createdTime",
		"finishedTime",
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

	varFinishNotifyRequest := _FinishNotifyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFinishNotifyRequest)

	if err != nil {
		return err
	}

	*o = FinishNotifyRequest(varFinishNotifyRequest)

	return err
}

type NullableFinishNotifyRequest struct {
	value *FinishNotifyRequest
	isSet bool
}

func (v NullableFinishNotifyRequest) Get() *FinishNotifyRequest {
	return v.value
}

func (v *NullableFinishNotifyRequest) Set(val *FinishNotifyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFinishNotifyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFinishNotifyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinishNotifyRequest(val *FinishNotifyRequest) *NullableFinishNotifyRequest {
	return &NullableFinishNotifyRequest{value: val, isSet: true}
}

func (v NullableFinishNotifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinishNotifyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


