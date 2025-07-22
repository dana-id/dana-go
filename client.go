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

package dana

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"

	common "github.com/dana-id/dana-go/common"
	config "github.com/dana-id/dana-go/config"
	disbursement "github.com/dana-id/dana-go/disbursement/v1"
	exceptions "github.com/dana-id/dana-go/exceptions"
	merchant_management "github.com/dana-id/dana-go/merchant_management/v1"
	payment_gateway "github.com/dana-id/dana-go/payment_gateway/v1"
	widget "github.com/dana-id/dana-go/widget/v1"
)

var (
	JsonCheck       = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
	XmlCheck        = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer("%5B", "[", "%5D", "]")
)

// APIClient manages communication with the Payment Gateway API API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *config.Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	PaymentGatewayAPI     *payment_gateway.PaymentGatewayAPIService
	WidgetAPI             *widget.WidgetAPIService
	DisbursementAPI       *disbursement.DisbursementAPIService
	MerchantManagementAPI *merchant_management.MerchantManagementAPIService
}

// service is the base service implementation that can be embedded by specific services
type service struct {
	client *APIClient // self-reference back to the main client
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *config.Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	// Initialize services with client reference
	c.PaymentGatewayAPI = payment_gateway.NewPaymentGatewayAPIService(c)
	c.WidgetAPI = widget.NewWidgetAPIService(c)
	c.DisbursementAPI = disbursement.NewDisbursementAPIService(c)
	c.MerchantManagementAPI = merchant_management.NewMerchantManagementAPIService(c)
	return c
}

// Note: Utility functions have been moved to utils/parameter_utils.go

// executeRequestImpl is the private implementation of the ExecuteRequest method
func (c *APIClient) executeRequestImpl(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []common.FormFile,
	result interface{}) (*http.Response, error) {

	// Prepare the request
	req, err := c.prepareRequestImpl(ctx, path, method, postBody, headerParams, queryParams, formParams, formFiles)
	if err != nil {
		return nil, err
	}

	// Call the API
	resp, err := c.callAPIImpl(req)
	if err != nil || resp == nil {
		return resp, err
	}

	// Read the body
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		return resp, err
	}

	// Handle error status codes
	if resp.StatusCode >= 300 {
		return resp, &exceptions.GenericOpenAPIError{
			RawBody:  body,
			ErrorMsg: resp.Status,
		}
	}

	// Decode response into result
	if result != nil {
		err = c.decodeImpl(result, body, resp.Header.Get("Content-Type"))
		if err != nil {
			return resp, &exceptions.GenericOpenAPIError{
				RawBody:  body,
				ErrorMsg: err.Error(),
			}
		}
	}

	return resp, nil
}

// ExecuteRequest delegates to the private implementation
// This method provides a convenient high-level API for services
func (c *APIClient) ExecuteRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []common.FormFile,
	result interface{}) (*http.Response, error) {
	return c.executeRequestImpl(ctx, path, method, postBody, headerParams, queryParams, formParams, formFiles, result)
}

// callAPIImpl is the private implementation of the API call
func (c *APIClient) callAPIImpl(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// CallAPI delegates to the private implementation
// This method is maintained for interface compatibility
func (c *APIClient) CallAPI(request *http.Request) (*http.Response, error) {
	return c.callAPIImpl(request)
}

// GetConfig returns the configuration that was used to initialize the APIClient
func (c *APIClient) GetConfig() *config.Configuration {
	cfg := &config.Configuration{
		UserAgent:        c.cfg.UserAgent,
		Debug:            c.cfg.Debug,
		DefaultHeader:    c.cfg.DefaultHeader,
		Servers:          c.cfg.Servers,
		APIKey:           c.cfg.APIKey, // Make sure to copy the APIKey
		HTTPClient:       c.cfg.HTTPClient,
		OperationServers: c.cfg.OperationServers,
	}
	return cfg
}

// prepareRequestImpl is the private implementation of request preparation
func (c *APIClient) prepareRequestImpl(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []common.FormFile) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("cannot specify postBody and multipart form at the same time")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.FileBytes) > 0 && formFile.FileName != "" {
				w.Boundary()
				part, err := w.CreateFormFile(formFile.FormFileName, filepath.Base(formFile.FileName))
				if err != nil {
					return nil, err
				}
				_, err = part.Write(formFile.FileBytes)
				if err != nil {
					return nil, err
				}
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("cannot specify postBody and x-www-form-urlencoded form at the same time")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
		pieces := strings.Split(s, "=")
		pieces[0] = queryDescape.Replace(pieces[0])
		return strings.Join(pieces, "=")
	})

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

// PrepareRequest delegates to the private implementation
func (c *APIClient) PrepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []common.FormFile) (localVarRequest *http.Request, err error) {
	return c.prepareRequestImpl(ctx, path, method, postBody, headerParams, queryParams, formParams, formFiles)
}

// Decode delegates to the private implementation
// This method is maintained for interface compatibility
func (c *APIClient) Decode(v interface{}, b []byte, contentType string) (err error) {
	return c.decodeImpl(v, b, contentType)
}

// decodeImpl is the private implementation of the response decoding
func (c *APIClient) decodeImpl(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(*os.File); ok {
		f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = f.Write(b)
		if err != nil {
			return
		}
		_, err = f.Seek(0, io.SeekStart)
		return
	}
	if f, ok := v.(**os.File); ok {
		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if XmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				// Try direct unmarshaling first
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					// If it fails due to unknown fields, try using mapstructure
					if strings.Contains(err.Error(), "unknown field") {
						actualInstance := actualObj.GetActualInstance()
						if err = UnmarshalJson(b, actualInstance); err == nil {
							return nil
						}
					}
					return err
				}
			} else {
				return errors.New("unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else { // simple model
			// Try standard unmarshaling first
			if err = json.Unmarshal(b, v); err != nil {
				// If it fails due to unknown fields, try mapstructure
				if strings.Contains(err.Error(), "unknown field") {
					err = UnmarshalJson(b, v)
					if err == nil {
						return nil
					}
				}
				return err
			}
		}
		return nil
	}
	return errors.New("undefined response type")
}

func UnmarshalJson(input []byte, result interface{}) error {
	tempMap := make(map[string]interface{})
	if err := json.Unmarshal(input, &tempMap); err != nil {
		return err
	}

	var md mapstructure.Metadata
	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			Metadata: &md,
			Result:   result,
		})
	if err != nil {
		return err
	}

	if err := decoder.Decode(tempMap); err != nil {
		return err
	}

	return nil
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(*os.File); ok {
		_, err = bodyBuf.ReadFrom(fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if JsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if XmlCheck.MatchString(contentType) {
		var bs []byte
		bs, err = xml.Marshal(body)
		if err == nil {
			bodyBuf.Write(bs)
		}
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}
