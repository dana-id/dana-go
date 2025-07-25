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

// checks if the DanaAccountInquiryResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DanaAccountInquiryResponse{}

// DanaAccountInquiryResponse struct for DanaAccountInquiryResponse
type DanaAccountInquiryResponse struct {
	// Refer to response code list
	ResponseCode string `json:"responseCode"`
	// Refer to response code list
	ResponseMessage string `json:"responseMessage"`
	// Transaction identifier on DANA system
	ReferenceNo *string `json:"referenceNo,omitempty"`
	// Unique transaction identifier on partner system which assigned to each transaction<br> Notes:<br> If the partner receives a timeout or an unexpected response from DANA and partner expects to perform retry request to DANA, please use the partnerReferenceNo that is the same as the one used in the transaction request process before 
	PartnerReferenceNo *string `json:"partnerReferenceNo,omitempty"`
	// Session identifier
	SessionId *string `json:"sessionId,omitempty"`
	// Customer account number, in format 628xxx
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// Customer account name
	CustomerName string `json:"customerName"`
	// Limitation of transfer to DANA balance for customer per month
	CustomerMonthlyLimit *float32 `json:"customerMonthlyLimit,omitempty"`
	// Minimal amount. Contains two sub-fields:<br> 1. Value: Amount, including the cents<br> 2. Currency: Currency code based on ISO 
	MinAmount Money `json:"minAmount"`
	// Maximal amount. Contains two sub-fields:<br> 1. Value: Amount, including the cents<br> 2. Currency: Currency code based on ISO 
	MaxAmount Money `json:"maxAmount"`
	// Amount. Contains two sub-fields:<br> 1. Value: Transaction amount, including the cents<br> 2. Currency: Currency code based on ISO 
	Amount Money `json:"amount"`
	// Fee amount. Contains two sub-fields:<br> 1. Value: Amount, including the cents<br> 2. Currency: Currency code based on ISO 
	FeeAmount Money `json:"feeAmount"`
	// Type of fee for each transfer to DANA transaction. Such as admin fee
	FeeType *string `json:"feeType,omitempty"`
	// Additional information
	AdditionalInfo map[string]interface{} `json:"additionalInfo,omitempty"`
}

type _DanaAccountInquiryResponse DanaAccountInquiryResponse

// NewDanaAccountInquiryResponse instantiates a new DanaAccountInquiryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDanaAccountInquiryResponse(responseCode string, responseMessage string, customerName string, minAmount Money, maxAmount Money, amount Money, feeAmount Money) *DanaAccountInquiryResponse {
	this := DanaAccountInquiryResponse{}
	this.ResponseCode = responseCode
	this.ResponseMessage = responseMessage
	this.CustomerName = customerName
	this.MinAmount = minAmount
	this.MaxAmount = maxAmount
	this.Amount = amount
	this.FeeAmount = feeAmount
	return &this
}

// NewDanaAccountInquiryResponseWithDefaults instantiates a new DanaAccountInquiryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDanaAccountInquiryResponseWithDefaults() *DanaAccountInquiryResponse {
	this := DanaAccountInquiryResponse{}
	return &this
}

// GetResponseCode returns the ResponseCode field value
func (o *DanaAccountInquiryResponse) GetResponseCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseCode, true
}

// SetResponseCode sets field value
func (o *DanaAccountInquiryResponse) SetResponseCode(v string) {
	o.ResponseCode = v
}

// GetResponseMessage returns the ResponseMessage field value
func (o *DanaAccountInquiryResponse) GetResponseMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetResponseMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseMessage, true
}

// SetResponseMessage sets field value
func (o *DanaAccountInquiryResponse) SetResponseMessage(v string) {
	o.ResponseMessage = v
}

// GetReferenceNo returns the ReferenceNo field value if set, zero value otherwise.
func (o *DanaAccountInquiryResponse) GetReferenceNo() string {
	if o == nil || utils.IsNil(o.ReferenceNo) {
		var ret string
		return ret
	}
	return *o.ReferenceNo
}

// GetReferenceNoOk returns a tuple with the ReferenceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetReferenceNoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReferenceNo) {
		return nil, false
	}
	return o.ReferenceNo, true
}

// HasReferenceNo returns a boolean if a field has been set.
func (o *DanaAccountInquiryResponse) HasReferenceNo() bool {
	if o != nil && !utils.IsNil(o.ReferenceNo) {
		return true
	}

	return false
}

// SetReferenceNo gets a reference to the given string and assigns it to the ReferenceNo field.
func (o *DanaAccountInquiryResponse) SetReferenceNo(v string) {
	o.ReferenceNo = &v
}

// GetPartnerReferenceNo returns the PartnerReferenceNo field value if set, zero value otherwise.
func (o *DanaAccountInquiryResponse) GetPartnerReferenceNo() string {
	if o == nil || utils.IsNil(o.PartnerReferenceNo) {
		var ret string
		return ret
	}
	return *o.PartnerReferenceNo
}

// GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetPartnerReferenceNoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PartnerReferenceNo) {
		return nil, false
	}
	return o.PartnerReferenceNo, true
}

// HasPartnerReferenceNo returns a boolean if a field has been set.
func (o *DanaAccountInquiryResponse) HasPartnerReferenceNo() bool {
	if o != nil && !utils.IsNil(o.PartnerReferenceNo) {
		return true
	}

	return false
}

// SetPartnerReferenceNo gets a reference to the given string and assigns it to the PartnerReferenceNo field.
func (o *DanaAccountInquiryResponse) SetPartnerReferenceNo(v string) {
	o.PartnerReferenceNo = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *DanaAccountInquiryResponse) GetSessionId() string {
	if o == nil || utils.IsNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetSessionIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *DanaAccountInquiryResponse) HasSessionId() bool {
	if o != nil && !utils.IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *DanaAccountInquiryResponse) SetSessionId(v string) {
	o.SessionId = &v
}

// GetCustomerNumber returns the CustomerNumber field value if set, zero value otherwise.
func (o *DanaAccountInquiryResponse) GetCustomerNumber() string {
	if o == nil || utils.IsNil(o.CustomerNumber) {
		var ret string
		return ret
	}
	return *o.CustomerNumber
}

// GetCustomerNumberOk returns a tuple with the CustomerNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetCustomerNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CustomerNumber) {
		return nil, false
	}
	return o.CustomerNumber, true
}

// HasCustomerNumber returns a boolean if a field has been set.
func (o *DanaAccountInquiryResponse) HasCustomerNumber() bool {
	if o != nil && !utils.IsNil(o.CustomerNumber) {
		return true
	}

	return false
}

// SetCustomerNumber gets a reference to the given string and assigns it to the CustomerNumber field.
func (o *DanaAccountInquiryResponse) SetCustomerNumber(v string) {
	o.CustomerNumber = &v
}

// GetCustomerName returns the CustomerName field value
func (o *DanaAccountInquiryResponse) GetCustomerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerName
}

// GetCustomerNameOk returns a tuple with the CustomerName field value
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetCustomerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerName, true
}

// SetCustomerName sets field value
func (o *DanaAccountInquiryResponse) SetCustomerName(v string) {
	o.CustomerName = v
}

// GetCustomerMonthlyLimit returns the CustomerMonthlyLimit field value if set, zero value otherwise.
func (o *DanaAccountInquiryResponse) GetCustomerMonthlyLimit() float32 {
	if o == nil || utils.IsNil(o.CustomerMonthlyLimit) {
		var ret float32
		return ret
	}
	return *o.CustomerMonthlyLimit
}

// GetCustomerMonthlyLimitOk returns a tuple with the CustomerMonthlyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetCustomerMonthlyLimitOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.CustomerMonthlyLimit) {
		return nil, false
	}
	return o.CustomerMonthlyLimit, true
}

// HasCustomerMonthlyLimit returns a boolean if a field has been set.
func (o *DanaAccountInquiryResponse) HasCustomerMonthlyLimit() bool {
	if o != nil && !utils.IsNil(o.CustomerMonthlyLimit) {
		return true
	}

	return false
}

// SetCustomerMonthlyLimit gets a reference to the given float32 and assigns it to the CustomerMonthlyLimit field.
func (o *DanaAccountInquiryResponse) SetCustomerMonthlyLimit(v float32) {
	o.CustomerMonthlyLimit = &v
}

// GetMinAmount returns the MinAmount field value
func (o *DanaAccountInquiryResponse) GetMinAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.MinAmount
}

// GetMinAmountOk returns a tuple with the MinAmount field value
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetMinAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinAmount, true
}

// SetMinAmount sets field value
func (o *DanaAccountInquiryResponse) SetMinAmount(v Money) {
	o.MinAmount = v
}

// GetMaxAmount returns the MaxAmount field value
func (o *DanaAccountInquiryResponse) GetMaxAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.MaxAmount
}

// GetMaxAmountOk returns a tuple with the MaxAmount field value
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetMaxAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxAmount, true
}

// SetMaxAmount sets field value
func (o *DanaAccountInquiryResponse) SetMaxAmount(v Money) {
	o.MaxAmount = v
}

// GetAmount returns the Amount field value
func (o *DanaAccountInquiryResponse) GetAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *DanaAccountInquiryResponse) SetAmount(v Money) {
	o.Amount = v
}

// GetFeeAmount returns the FeeAmount field value
func (o *DanaAccountInquiryResponse) GetFeeAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetFeeAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeAmount, true
}

// SetFeeAmount sets field value
func (o *DanaAccountInquiryResponse) SetFeeAmount(v Money) {
	o.FeeAmount = v
}

// GetFeeType returns the FeeType field value if set, zero value otherwise.
func (o *DanaAccountInquiryResponse) GetFeeType() string {
	if o == nil || utils.IsNil(o.FeeType) {
		var ret string
		return ret
	}
	return *o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetFeeTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FeeType) {
		return nil, false
	}
	return o.FeeType, true
}

// HasFeeType returns a boolean if a field has been set.
func (o *DanaAccountInquiryResponse) HasFeeType() bool {
	if o != nil && !utils.IsNil(o.FeeType) {
		return true
	}

	return false
}

// SetFeeType gets a reference to the given string and assigns it to the FeeType field.
func (o *DanaAccountInquiryResponse) SetFeeType(v string) {
	o.FeeType = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *DanaAccountInquiryResponse) GetAdditionalInfo() map[string]interface{} {
	if o == nil || utils.IsNil(o.AdditionalInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DanaAccountInquiryResponse) GetAdditionalInfoOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.AdditionalInfo) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *DanaAccountInquiryResponse) HasAdditionalInfo() bool {
	if o != nil && !utils.IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given map[string]interface{} and assigns it to the AdditionalInfo field.
func (o *DanaAccountInquiryResponse) SetAdditionalInfo(v map[string]interface{}) {
	o.AdditionalInfo = v
}

func (o DanaAccountInquiryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DanaAccountInquiryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["responseCode"] = o.ResponseCode
	toSerialize["responseMessage"] = o.ResponseMessage
	if !utils.IsNil(o.ReferenceNo) {
		toSerialize["referenceNo"] = o.ReferenceNo
	}
	if !utils.IsNil(o.PartnerReferenceNo) {
		toSerialize["partnerReferenceNo"] = o.PartnerReferenceNo
	}
	if !utils.IsNil(o.SessionId) {
		toSerialize["sessionId"] = o.SessionId
	}
	if !utils.IsNil(o.CustomerNumber) {
		toSerialize["customerNumber"] = o.CustomerNumber
	}
	toSerialize["customerName"] = o.CustomerName
	if !utils.IsNil(o.CustomerMonthlyLimit) {
		toSerialize["customerMonthlyLimit"] = o.CustomerMonthlyLimit
	}
	toSerialize["minAmount"] = o.MinAmount
	toSerialize["maxAmount"] = o.MaxAmount
	toSerialize["amount"] = o.Amount
	toSerialize["feeAmount"] = o.FeeAmount
	if !utils.IsNil(o.FeeType) {
		toSerialize["feeType"] = o.FeeType
	}
	if !utils.IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	return toSerialize, nil
}

func (o *DanaAccountInquiryResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"responseCode",
		"responseMessage",
		"customerName",
		"minAmount",
		"maxAmount",
		"amount",
		"feeAmount",
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

	varDanaAccountInquiryResponse := _DanaAccountInquiryResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDanaAccountInquiryResponse)

	if err != nil {
		return err
	}

	*o = DanaAccountInquiryResponse(varDanaAccountInquiryResponse)

	return err
}

type NullableDanaAccountInquiryResponse struct {
	value *DanaAccountInquiryResponse
	isSet bool
}

func (v NullableDanaAccountInquiryResponse) Get() *DanaAccountInquiryResponse {
	return v.value
}

func (v *NullableDanaAccountInquiryResponse) Set(val *DanaAccountInquiryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDanaAccountInquiryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDanaAccountInquiryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDanaAccountInquiryResponse(val *DanaAccountInquiryResponse) *NullableDanaAccountInquiryResponse {
	return &NullableDanaAccountInquiryResponse{value: val, isSet: true}
}

func (v NullableDanaAccountInquiryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDanaAccountInquiryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


