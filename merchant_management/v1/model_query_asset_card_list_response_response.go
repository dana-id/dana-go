// Copyright 2026 PT Espay Debit Indonesia Koe
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

import (
	"encoding/json"
	"bytes"
	"fmt"
	utils "github.com/dana-id/dana-go/v2/utils"
)

// checks if the QueryAssetCardListResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QueryAssetCardListResponseResponse{}

// QueryAssetCardListResponseResponse struct for QueryAssetCardListResponseResponse
type QueryAssetCardListResponseResponse struct {
	Head QueryAssetCardListResponseResponseHead `json:"head"`
	Body QueryAssetCardListResponseResponseBody `json:"body"`
}

type _QueryAssetCardListResponseResponse QueryAssetCardListResponseResponse

// NewQueryAssetCardListResponseResponse instantiates a new QueryAssetCardListResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAssetCardListResponseResponse(head QueryAssetCardListResponseResponseHead, body QueryAssetCardListResponseResponseBody) *QueryAssetCardListResponseResponse {
	this := QueryAssetCardListResponseResponse{}
	this.Head = head
	this.Body = body
	return &this
}

// NewQueryAssetCardListResponseResponseWithDefaults instantiates a new QueryAssetCardListResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAssetCardListResponseResponseWithDefaults() *QueryAssetCardListResponseResponse {
	this := QueryAssetCardListResponseResponse{}
	return &this
}

// GetHead returns the Head field value
func (o *QueryAssetCardListResponseResponse) GetHead() QueryAssetCardListResponseResponseHead {
	if o == nil {
		var ret QueryAssetCardListResponseResponseHead
		return ret
	}

	return o.Head
}

// GetHeadOk returns a tuple with the Head field value
// and a boolean to check if the value has been set.
func (o *QueryAssetCardListResponseResponse) GetHeadOk() (*QueryAssetCardListResponseResponseHead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Head, true
}

// SetHead sets field value
func (o *QueryAssetCardListResponseResponse) SetHead(v QueryAssetCardListResponseResponseHead) {
	o.Head = v
}

// GetBody returns the Body field value
func (o *QueryAssetCardListResponseResponse) GetBody() QueryAssetCardListResponseResponseBody {
	if o == nil {
		var ret QueryAssetCardListResponseResponseBody
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *QueryAssetCardListResponseResponse) GetBodyOk() (*QueryAssetCardListResponseResponseBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *QueryAssetCardListResponseResponse) SetBody(v QueryAssetCardListResponseResponseBody) {
	o.Body = v
}

func (o QueryAssetCardListResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAssetCardListResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["head"] = o.Head
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

func (o *QueryAssetCardListResponseResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"head",
		"body",
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

	varQueryAssetCardListResponseResponse := _QueryAssetCardListResponseResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryAssetCardListResponseResponse)

	if err != nil {
		return err
	}

	*o = QueryAssetCardListResponseResponse(varQueryAssetCardListResponseResponse)

	return err
}

type NullableQueryAssetCardListResponseResponse struct {
	value *QueryAssetCardListResponseResponse
	isSet bool
}

func (v NullableQueryAssetCardListResponseResponse) Get() *QueryAssetCardListResponseResponse {
	return v.value
}

func (v *NullableQueryAssetCardListResponseResponse) Set(val *QueryAssetCardListResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAssetCardListResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAssetCardListResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryAssetCardListResponseResponse(val *QueryAssetCardListResponseResponse) *NullableQueryAssetCardListResponseResponse {
	return &NullableQueryAssetCardListResponseResponse{value: val, isSet: true}
}

func (v NullableQueryAssetCardListResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAssetCardListResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


