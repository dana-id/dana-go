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

package exceptions

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	RawBody  []byte `json:"body"`
	ErrorMsg string `json:"error"`
	ModelData any    `json:"model"`
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.ErrorMsg
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.RawBody
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() any {
	return e.ModelData
}
