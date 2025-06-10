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
