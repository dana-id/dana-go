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

package common

import (
	"context"
	"net/http"
	"net/url"

	config "github.com/dana-id/dana-go/config"
)

// This package provides common interfaces and types used across the Dana API client

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
// This is the legacy version maintained for backward compatibility
type APIKey struct {
	Key    string
	Prefix string
}

// FormFile represents a file to be submitted in a multipart/form-data request
type FormFile struct {
	FileBytes    []byte
	FileName     string
	FormFileName string
}

// ClientInterface defines the methods that service objects need from the API client
type ClientInterface interface {
	// Core request methods
	PrepareRequest(ctx context.Context, path string, method string,
		postBody interface{}, headerParams map[string]string,
		queryParams url.Values, formParams url.Values,
		formFiles []FormFile) (*http.Request, error)
	CallAPI(req *http.Request) (*http.Response, error)
	Decode(v interface{}, b []byte, contentType string) error
	GetConfig() *config.Configuration

	// High-level convenience method that combines the above methods
	ExecuteRequest(ctx context.Context, path string, method string,
		postBody interface{}, headerParams map[string]string,
		queryParams url.Values, formParams url.Values,
		formFiles []FormFile, result interface{}) (*http.Response, error)
}
