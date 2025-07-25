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

package payment_gateway

/*
Payment Gateway API

API for doing operations in DANA Payment Gateway (Gapura)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.


import (
	"encoding/json"
	"bytes"
	"fmt"
	utils "github.com/dana-id/dana-go/utils"
)

// checks if the RefundOrderRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RefundOrderRequest{}

// RefundOrderRequest struct for RefundOrderRequest
type RefundOrderRequest struct {
	// Merchant identifier that is unique per each merchant
	MerchantId string `json:"merchantId"`
	// Information of sub merchant identifier
	SubMerchantId *string `json:"subMerchantId,omitempty"`
	// Original transaction identifier on DANA system
	OriginalReferenceNo *string `json:"originalReferenceNo,omitempty"`
	// Original transaction identifier on partner system
	OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
	// Original external identifier on header message
	OriginalExternalId *string `json:"originalExternalId,omitempty"`
	// DANA's capture identifier. Use to refund the corresponding capture order. Required if auth payment scenario
	OriginalCaptureNo *string `json:"originalCaptureNo,omitempty"`
	// Reference number from merchant for the refund
	PartnerRefundNo string `json:"partnerRefundNo"`
	// Refund amount. Contains two sub-fields:<br> 1. Value: Transaction amount, including the cents<br> 2. Currency: Currency code based on ISO<br> 
	RefundAmount Money `json:"refundAmount"`
	// Store identifier to indicate to which store this payment belongs to
	ExternalStoreId *string `json:"externalStoreId,omitempty"`
	// Refund reason
	Reason *string `json:"reason,omitempty"`
	// Additional information
	AdditionalInfo *RefundOrderRequestAdditionalInfo `json:"additionalInfo,omitempty"`
}

type _RefundOrderRequest RefundOrderRequest

// NewRefundOrderRequest instantiates a new RefundOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundOrderRequest(merchantId string, originalPartnerReferenceNo string, partnerRefundNo string, refundAmount Money) *RefundOrderRequest {
	this := RefundOrderRequest{}
	this.MerchantId = merchantId
	this.OriginalPartnerReferenceNo = originalPartnerReferenceNo
	this.PartnerRefundNo = partnerRefundNo
	this.RefundAmount = refundAmount
	return &this
}

// NewRefundOrderRequestWithDefaults instantiates a new RefundOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundOrderRequestWithDefaults() *RefundOrderRequest {
	this := RefundOrderRequest{}
	return &this
}

// GetMerchantId returns the MerchantId field value
func (o *RefundOrderRequest) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *RefundOrderRequest) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *RefundOrderRequest) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetSubMerchantId returns the SubMerchantId field value if set, zero value otherwise.
func (o *RefundOrderRequest) GetSubMerchantId() string {
	if o == nil || utils.IsNil(o.SubMerchantId) {
		var ret string
		return ret
	}
	return *o.SubMerchantId
}

// GetSubMerchantIdOk returns a tuple with the SubMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundOrderRequest) GetSubMerchantIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SubMerchantId) {
		return nil, false
	}
	return o.SubMerchantId, true
}

// HasSubMerchantId returns a boolean if a field has been set.
func (o *RefundOrderRequest) HasSubMerchantId() bool {
	if o != nil && !utils.IsNil(o.SubMerchantId) {
		return true
	}

	return false
}

// SetSubMerchantId gets a reference to the given string and assigns it to the SubMerchantId field.
func (o *RefundOrderRequest) SetSubMerchantId(v string) {
	o.SubMerchantId = &v
}

// GetOriginalReferenceNo returns the OriginalReferenceNo field value if set, zero value otherwise.
func (o *RefundOrderRequest) GetOriginalReferenceNo() string {
	if o == nil || utils.IsNil(o.OriginalReferenceNo) {
		var ret string
		return ret
	}
	return *o.OriginalReferenceNo
}

// GetOriginalReferenceNoOk returns a tuple with the OriginalReferenceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundOrderRequest) GetOriginalReferenceNoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OriginalReferenceNo) {
		return nil, false
	}
	return o.OriginalReferenceNo, true
}

// HasOriginalReferenceNo returns a boolean if a field has been set.
func (o *RefundOrderRequest) HasOriginalReferenceNo() bool {
	if o != nil && !utils.IsNil(o.OriginalReferenceNo) {
		return true
	}

	return false
}

// SetOriginalReferenceNo gets a reference to the given string and assigns it to the OriginalReferenceNo field.
func (o *RefundOrderRequest) SetOriginalReferenceNo(v string) {
	o.OriginalReferenceNo = &v
}

// GetOriginalPartnerReferenceNo returns the OriginalPartnerReferenceNo field value
func (o *RefundOrderRequest) GetOriginalPartnerReferenceNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalPartnerReferenceNo
}

// GetOriginalPartnerReferenceNoOk returns a tuple with the OriginalPartnerReferenceNo field value
// and a boolean to check if the value has been set.
func (o *RefundOrderRequest) GetOriginalPartnerReferenceNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalPartnerReferenceNo, true
}

// SetOriginalPartnerReferenceNo sets field value
func (o *RefundOrderRequest) SetOriginalPartnerReferenceNo(v string) {
	o.OriginalPartnerReferenceNo = v
}

// GetOriginalExternalId returns the OriginalExternalId field value if set, zero value otherwise.
func (o *RefundOrderRequest) GetOriginalExternalId() string {
	if o == nil || utils.IsNil(o.OriginalExternalId) {
		var ret string
		return ret
	}
	return *o.OriginalExternalId
}

// GetOriginalExternalIdOk returns a tuple with the OriginalExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundOrderRequest) GetOriginalExternalIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OriginalExternalId) {
		return nil, false
	}
	return o.OriginalExternalId, true
}

// HasOriginalExternalId returns a boolean if a field has been set.
func (o *RefundOrderRequest) HasOriginalExternalId() bool {
	if o != nil && !utils.IsNil(o.OriginalExternalId) {
		return true
	}

	return false
}

// SetOriginalExternalId gets a reference to the given string and assigns it to the OriginalExternalId field.
func (o *RefundOrderRequest) SetOriginalExternalId(v string) {
	o.OriginalExternalId = &v
}

// GetOriginalCaptureNo returns the OriginalCaptureNo field value if set, zero value otherwise.
func (o *RefundOrderRequest) GetOriginalCaptureNo() string {
	if o == nil || utils.IsNil(o.OriginalCaptureNo) {
		var ret string
		return ret
	}
	return *o.OriginalCaptureNo
}

// GetOriginalCaptureNoOk returns a tuple with the OriginalCaptureNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundOrderRequest) GetOriginalCaptureNoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OriginalCaptureNo) {
		return nil, false
	}
	return o.OriginalCaptureNo, true
}

// HasOriginalCaptureNo returns a boolean if a field has been set.
func (o *RefundOrderRequest) HasOriginalCaptureNo() bool {
	if o != nil && !utils.IsNil(o.OriginalCaptureNo) {
		return true
	}

	return false
}

// SetOriginalCaptureNo gets a reference to the given string and assigns it to the OriginalCaptureNo field.
func (o *RefundOrderRequest) SetOriginalCaptureNo(v string) {
	o.OriginalCaptureNo = &v
}

// GetPartnerRefundNo returns the PartnerRefundNo field value
func (o *RefundOrderRequest) GetPartnerRefundNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartnerRefundNo
}

// GetPartnerRefundNoOk returns a tuple with the PartnerRefundNo field value
// and a boolean to check if the value has been set.
func (o *RefundOrderRequest) GetPartnerRefundNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerRefundNo, true
}

// SetPartnerRefundNo sets field value
func (o *RefundOrderRequest) SetPartnerRefundNo(v string) {
	o.PartnerRefundNo = v
}

// GetRefundAmount returns the RefundAmount field value
func (o *RefundOrderRequest) GetRefundAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value
// and a boolean to check if the value has been set.
func (o *RefundOrderRequest) GetRefundAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefundAmount, true
}

// SetRefundAmount sets field value
func (o *RefundOrderRequest) SetRefundAmount(v Money) {
	o.RefundAmount = v
}

// GetExternalStoreId returns the ExternalStoreId field value if set, zero value otherwise.
func (o *RefundOrderRequest) GetExternalStoreId() string {
	if o == nil || utils.IsNil(o.ExternalStoreId) {
		var ret string
		return ret
	}
	return *o.ExternalStoreId
}

// GetExternalStoreIdOk returns a tuple with the ExternalStoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundOrderRequest) GetExternalStoreIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExternalStoreId) {
		return nil, false
	}
	return o.ExternalStoreId, true
}

// HasExternalStoreId returns a boolean if a field has been set.
func (o *RefundOrderRequest) HasExternalStoreId() bool {
	if o != nil && !utils.IsNil(o.ExternalStoreId) {
		return true
	}

	return false
}

// SetExternalStoreId gets a reference to the given string and assigns it to the ExternalStoreId field.
func (o *RefundOrderRequest) SetExternalStoreId(v string) {
	o.ExternalStoreId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RefundOrderRequest) GetReason() string {
	if o == nil || utils.IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundOrderRequest) GetReasonOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RefundOrderRequest) HasReason() bool {
	if o != nil && !utils.IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RefundOrderRequest) SetReason(v string) {
	o.Reason = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *RefundOrderRequest) GetAdditionalInfo() RefundOrderRequestAdditionalInfo {
	if o == nil || utils.IsNil(o.AdditionalInfo) {
		var ret RefundOrderRequestAdditionalInfo
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundOrderRequest) GetAdditionalInfoOk() (*RefundOrderRequestAdditionalInfo, bool) {
	if o == nil || utils.IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *RefundOrderRequest) HasAdditionalInfo() bool {
	if o != nil && !utils.IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given RefundOrderRequestAdditionalInfo and assigns it to the AdditionalInfo field.
func (o *RefundOrderRequest) SetAdditionalInfo(v RefundOrderRequestAdditionalInfo) {
	o.AdditionalInfo = &v
}

func (o RefundOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantId"] = o.MerchantId
	if !utils.IsNil(o.SubMerchantId) {
		toSerialize["subMerchantId"] = o.SubMerchantId
	}
	if !utils.IsNil(o.OriginalReferenceNo) {
		toSerialize["originalReferenceNo"] = o.OriginalReferenceNo
	}
	toSerialize["originalPartnerReferenceNo"] = o.OriginalPartnerReferenceNo
	if !utils.IsNil(o.OriginalExternalId) {
		toSerialize["originalExternalId"] = o.OriginalExternalId
	}
	if !utils.IsNil(o.OriginalCaptureNo) {
		toSerialize["originalCaptureNo"] = o.OriginalCaptureNo
	}
	toSerialize["partnerRefundNo"] = o.PartnerRefundNo
	toSerialize["refundAmount"] = o.RefundAmount
	if !utils.IsNil(o.ExternalStoreId) {
		toSerialize["externalStoreId"] = o.ExternalStoreId
	}
	if !utils.IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !utils.IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	return toSerialize, nil
}

func (o *RefundOrderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"merchantId",
		"originalPartnerReferenceNo",
		"partnerRefundNo",
		"refundAmount",
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

	varRefundOrderRequest := _RefundOrderRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRefundOrderRequest)

	if err != nil {
		return err
	}

	*o = RefundOrderRequest(varRefundOrderRequest)

	return err
}

type NullableRefundOrderRequest struct {
	value *RefundOrderRequest
	isSet bool
}

func (v NullableRefundOrderRequest) Get() *RefundOrderRequest {
	return v.value
}

func (v *NullableRefundOrderRequest) Set(val *RefundOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundOrderRequest(val *RefundOrderRequest) *NullableRefundOrderRequest {
	return &NullableRefundOrderRequest{value: val, isSet: true}
}

func (v NullableRefundOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


