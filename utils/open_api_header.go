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

package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/dana-id/dana-go/config"
	uuid "github.com/google/uuid"
)

// OPEN_API Header constants
const (
	OpenApiVersion      = "version"
	OpenApiFunction     = "function"
	OpenApiClientID     = "clientId"
	OpenApiReqTime      = "reqTime"
	OpenApiReqMsgID     = "reqMsgId"
	OpenApiClientSecret = "clientSecret"
	OpenApiAccessToken  = "accessToken"
	OpenApiReserve      = "reserve"
)

// OpenApiRuntimeHeaders defines the basic headers for OPEN_API requests
var OpenApiRuntimeHeaders = []string{
	OpenApiVersion, OpenApiFunction, OpenApiClientID, OpenApiReqTime,
	OpenApiReqMsgID, OpenApiClientSecret, OpenApiReserve,
}

// OpenApiWithAccessTokenRuntimeHeaders defines headers including access token
var OpenApiWithAccessTokenRuntimeHeaders = []string{
	OpenApiVersion, OpenApiFunction, OpenApiClientID, OpenApiReqTime,
	OpenApiReqMsgID, OpenApiClientSecret, OpenApiAccessToken, OpenApiReserve,
}

// SetOpenApiHeaders applies all OPEN_API authentication headers to the provided headers map
// This centralized function helps avoid duplicate authentication code across API methods
func SetOpenApiHeaders(headerParams map[string]string, apiKey *config.APIKey, body string, method string, endpoint string, functionName string) {
	if apiKey == nil {
		fmt.Errorf("API key is required for OPEN_API authentication")
		return
	}

	if apiKey.CLIENT_SECRET == "" {
		fmt.Errorf("Client secret is required for OPEN_API authentication")
		return
	}

	// Get runtime headers with signature
	runtimeHeaders, signature := GetOpenApiGeneratedAuthWithSignature(
		body, apiKey, functionName,
	)

	// Apply all headers
	for k, v := range runtimeHeaders {
		headerParams[k] = v
	}

	// Apply signature at the body level (will be handled by request marshaling)
	if signature != "" {
		// Store signature for request body modification
		headerParams["X-OPEN-API-SIGNATURE"] = signature
	}

	// Remove sensitive data from headerParams
	for k := range headerParams {
		// Check for keys that start with sensitive prefixes (case insensitive)
		if hasOpenApiPrefix(k, "PRIVATE") || hasOpenApiPrefix(k, "private") ||
			hasOpenApiPrefix(k, "DANA_ENV") || hasOpenApiPrefix(k, "dana_env") ||
			hasOpenApiPrefix(k, "ENV") || hasOpenApiPrefix(k, "env") {
			delete(headerParams, k)
		}
	}
}

// hasOpenApiPrefix checks if a string starts with a given prefix (case-sensitive)
func hasOpenApiPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

// GenerateOpenApiSignature generates OPEN_API signature according to the official documentation:
// 1. Compose the string to sign: (<HTTP BODY>)
// 2. Apply SHA-256 with RSA-2048 encryption using pkcs8 private key
// 3. Encode the result to base64
func GenerateOpenApiSignature(body string, apiKey *config.APIKey) (string, error) {

	// Get usable private key
	privateKeyData, err := GetUsablePrivateKey(apiKey)
	if err != nil || privateKeyData == nil {
		return "", fmt.Errorf("failed to get usable private key: %v", err)
	}

	// Parse the private key
	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return "", fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		// Fallback to PKCS1 format
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return "", fmt.Errorf("failed to parse private key: %v", err)
		}
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("private key is not RSA")
	}

	// For OPEN_API, the string to sign is just the HTTP body
	stringToSign := body
	// Create hash of the data to sign
	hash := sha256.New()
	hash.Write([]byte(stringToSign))
	hashedData := hash.Sum(nil)

	// Generate signature using SHA256 with RSA
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hashedData)
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}

	// Encode to base64
	encodedSignature := base64.StdEncoding.EncodeToString(signature)
	return encodedSignature, nil
}

func GetOpenApiGeneratedHeaders(body string, apiKey *config.APIKey, functionName string) map[string]string {
	if apiKey.DANA_ENV == "" && apiKey.ENV == "" {
		fmt.Errorf("DANA_ENV or ENV is required for OPEN_API authentication")
		return nil
	}

	// Extract client_id from configuration
	clientID := apiKey.CLIENT_ID
	if clientID == "" {
		// Fallback to X_PARTNER_ID if CLIENT_ID is not set
		clientID = apiKey.X_PARTNER_ID
	}
	if clientID == "" {
		// Final fallback to default
		clientID = "2014000014442"
	}

	// Generate timestamp in Jakarta time (GMT+7)
	jkt, err := time.LoadLocation("Asia/Jakarta")
	var jktTime time.Time
	if err != nil {
		// Create time with fixed +7 hours offset from UTC
		now := time.Now().UTC()
		jktTime = now.Add(7 * time.Hour)
	} else {
		// Get current time in Jakarta timezone
		jktTime = time.Now().In(jkt)
	}

	timestamp := jktTime.Format("2006-01-02T15:04:05+07:00")

	// Generate unique request message ID
	reqMsgID := "sdk" + uuid.New().String()[3:]

	// Basic OPEN_API headers
	headers := map[string]string{
		OpenApiVersion:      "2.0",
		OpenApiFunction:     functionName,
		OpenApiClientID:     clientID,
		OpenApiReqTime:      timestamp,
		OpenApiReqMsgID:     reqMsgID,
		OpenApiClientSecret: getClientSecret(apiKey),
		OpenApiReserve:      "{}",
	}

	return headers
}

// GetOpenApiGeneratedAuthWithSignature generates OPEN_API authentication headers and signature
func GetOpenApiGeneratedAuthWithSignature(body string, apiKey *config.APIKey, functionName string) (map[string]string, string) {
	// Generate headers
	headers := GetOpenApiGeneratedHeaders(body, apiKey, functionName)

	// Generate signature if private key is available
	signature := ""
	if apiKey.PRIVATE_KEY != "" || apiKey.PRIVATE_KEY_PATH != "" {
		sig, err := GenerateOpenApiSignature(body, apiKey)
		if err != nil {
			fmt.Errorf("Failed to generate OPEN_API signature: %v\n", err)
		} else {
			signature = sig
		}
	}

	return headers, signature
}

// getClientSecret extracts client secret from API key configuration
func getClientSecret(apiKey *config.APIKey) string {
	// Return the CLIENT_SECRET from the API key configuration
	if apiKey != nil && apiKey.CLIENT_SECRET != "" {
		return apiKey.CLIENT_SECRET
	}

	// Fallback to empty string if not configured
	// In production, this should be properly configured
	return ""
}

// MergeWithOpenApiRuntimeHeaders removes any items containing 'private' or 'env' and merges with OPEN_API runtime headers
func MergeWithOpenApiRuntimeHeaders(authFromUsers []string, requiresAccessToken bool) []string {
	// Filter out auth items containing 'private' or 'env'
	var filteredAuth []string
	for _, auth := range authFromUsers {
		if !strings.Contains(strings.ToLower(auth), "private") &&
			!strings.Contains(strings.ToLower(auth), "env") {
			filteredAuth = append(filteredAuth, auth)
		}
	}

	// Create a map for deduplication
	headerMap := make(map[string]bool)

	// Add filtered auth headers
	for _, auth := range filteredAuth {
		headerMap[auth] = true
	}

	// Add runtime headers
	var runtimeHeaders []string
	if requiresAccessToken {
		runtimeHeaders = OpenApiWithAccessTokenRuntimeHeaders
	} else {
		runtimeHeaders = OpenApiRuntimeHeaders
	}

	for _, header := range runtimeHeaders {
		headerMap[header] = true
	}

	// Convert back to slice
	var result []string
	for header := range headerMap {
		result = append(result, header)
	}

	return result
}
