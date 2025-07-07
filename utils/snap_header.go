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
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dana-id/dana-go/config"
	uuid "github.com/google/uuid"
)

// Header constants
const (
	HeaderOrigin                = "ORIGIN"
	HeaderChannelID             = "CHANNEL-ID"
	HeaderXPartnerID            = "X-PARTNER-ID"
	HeaderXClientKey            = "X-CLIENT-KEY"
	HeaderXTimestamp            = "X-TIMESTAMP"
	HeaderXSignature            = "X-SIGNATURE"
	HeaderXExternalID           = "X-EXTERNAL-ID"
	HeaderAuthorizationCustomer = "Authorization-Customer"
	HeaderXIPAddress            = "X-IP-ADDRESS"
	HeaderXDeviceID             = "X-DEVICE-ID"
	HeaderXLatitude             = "X-LATITUDE"
	HeaderXLongitude            = "X-LONGITUDE"
)

// Scenario constants
const (
	ScenarioApplyToken       = "apply_token"
	ScenarioApplyOTT         = "apply_ott"
	ScenarioUnbindingAccount = "unbinding_account"
)

// SetSnapHeaders applies all Snap API authentication headers to the provided headers map
// This centralized function helps avoid duplicate authentication code across API methods
func SetSnapHeaders(headerParams map[string]string, apiKey *config.APIKey, body string, method string, endpoint string, opts ...string) {
	if apiKey == nil {
		ReportError("API key is required for authentication")
		return
	}

	// Set ORIGIN header
	if apiKey.ORIGIN != "" {
		headerParams[HeaderOrigin] = apiKey.ORIGIN
	}

	// Set CHANNEL-ID header
	headerParams[HeaderChannelID] = "95221"

	// Get scenario from opts if provided
	var scenario string
	if len(opts) > 0 && opts[0] != "" {
		scenario = opts[0]
	}

	switch scenario {
	case ScenarioApplyToken:
		// Apply-Token uses X-CLIENT-KEY instead of X-PARTNER-ID
		if apiKey.X_PARTNER_ID != "" {
			headerParams[HeaderXClientKey] = apiKey.X_PARTNER_ID
		}
	default:
		// Other scenarios retain X-PARTNER-ID
		if apiKey.X_PARTNER_ID != "" {
			headerParams[HeaderXPartnerID] = apiKey.X_PARTNER_ID
		}
	}

	runtimeHeaders := getRuntimeHeaders(body, apiKey, method, endpoint, scenario)
	for k, v := range runtimeHeaders {
		headerParams[k] = v
	}

	// Remove sensitive data from headerParams
	for k := range headerParams {
		// Check for keys that start with sensitive prefixes (case insensitive)
		if hasPrefix(k, "PRIVATE") || hasPrefix(k, "private") ||
			hasPrefix(k, "ENV") || hasPrefix(k, "env") {
			delete(headerParams, k)
		}
	}
}

func getRuntimeHeaders(body string, apiKey *config.APIKey, method string, endpoint string, scenario string) map[string]string {
	// Define Jakarta location (GMT+7)
	jkt, err := time.LoadLocation("Asia/Jakarta")

	// Default to UTC+7 if location loading fails
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

	baseHeaders := map[string]string{
		HeaderXTimestamp: timestamp,
	}

	switch scenario {
	case ScenarioApplyToken:
		clientKey := apiKey.X_PARTNER_ID
		baseHeaders[HeaderXSignature] = generateSignature(body, apiKey, method, endpoint, timestamp, scenario)
		baseHeaders[HeaderXClientKey] = clientKey
		return baseHeaders

	case ScenarioApplyOTT, ScenarioUnbindingAccount:
		baseHeaders[HeaderXExternalID] = uuid.New().String()
		baseHeaders[HeaderXSignature] = generateSignature(body, apiKey, method, endpoint, timestamp, scenario)

		extractAdditionalHeaders(baseHeaders, body)
		return baseHeaders

	default:
		baseHeaders[HeaderXExternalID] = uuid.New().String()
		baseHeaders[HeaderXSignature] = generateSignature(body, apiKey, method, endpoint, timestamp, scenario)
		return baseHeaders
	}
}

func generateSignature(body string, apiKey *config.APIKey, method string, endpoint string, timestamp string, scenario string) string {
	// Validate we have the required data
	if apiKey == nil {
		return ""
	}

	var data string

	// Handle apply_token scenario differently
	if scenario == ScenarioApplyToken {
		// Apply token signature: "<clientKey>|<timestamp>"
		clientKey := apiKey.X_PARTNER_ID
		data = fmt.Sprintf("%s|%s", clientKey, timestamp)
	} else {
		// Standard SNAP signature
		// Calculate SHA-256 hash of the request body
		hash := sha256.New()
		hash.Write([]byte(body))
		hashedPayload := fmt.Sprintf("%x", hash.Sum(nil))

		// Extract just the path component from the full URL
		parsedURL, err := url.Parse(endpoint)
		if err != nil {
			fmt.Println("Error parsing URL:", err)
			return ""
		}

		// Use only the path from the URL for signing
		urlPath := parsedURL.Path

		// Format the string to sign:
		// "<HTTP METHOD>:<RELATIVE PATH URL>:<LOWERCASE_HEX_ENCODED_SHA_256(MINIFIED_HTTP_BODY)>:<X-TIMESTAMP>"
		data = fmt.Sprintf("%s:%s:%s:%s", method, urlPath, hashedPayload, timestamp)
	}

	// Get usable private key
	privateKeyData, err := GetUsablePrivateKey(apiKey)
	if err != nil || privateKeyData == nil {
		fmt.Println("Failed to get usable private key:", err)
		return ""
	}

	// Generate signature
	signature, err := SignWithPrivateKey([]byte(data), privateKeyData)
	if err != nil {
		fmt.Println("Failed to generate signature:", err)
		return ""
	}

	return signature
}

// hasPrefix checks if a string starts with a given prefix (case-sensitive)
func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

// GetUsablePrivateKey retrieves and prepares the private key from either direct input or file path
// Private key path is the preferred method as it avoids issues with newlines and formatting
func GetUsablePrivateKey(apiKey *config.APIKey) ([]byte, error) {
	// Check if we have conflicting config
	if apiKey.PRIVATE_KEY != "" && apiKey.PRIVATE_KEY_PATH != "" {
		return nil, fmt.Errorf("conflicting private key config: both PRIVATE_KEY and PRIVATE_KEY_PATH provided")
	}

	// Prioritize reading from file if path is available
	if apiKey.PRIVATE_KEY_PATH != "" {
		// Read private key from file
		fmt.Println("Reading private key from file:", apiKey.PRIVATE_KEY_PATH)
		privateKeyBytes, err := os.ReadFile(apiKey.PRIVATE_KEY_PATH)
		if err != nil {
			return nil, fmt.Errorf("failed to read private key file: %v", err)
		}

		// Print the first line of the key just for verification
		keyStr := string(privateKeyBytes)
		if len(keyStr) > 0 {
			lines := strings.Split(keyStr, "\n")
			if len(lines) > 0 {
				fmt.Println("Private key file loaded successfully, begins with:", lines[0])
			}
		}

		return privateKeyBytes, nil
	} else if apiKey.PRIVATE_KEY != "" {
		// Replace escaped newlines with actual newlines
		privateKey := strings.ReplaceAll(apiKey.PRIVATE_KEY, "\\n", "\n")

		return []byte(privateKey), nil
	}

	return nil, errors.New("no private key or private key path provided")
}

// SignWithPrivateKey generates an RSA signature using the provided private key
func SignWithPrivateKey(data []byte, privateKeyPEM []byte) (string, error) {
	// Parse the private key
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		return "", errors.New("failed to parse PEM block containing private key")
	}

	var privateKey *rsa.PrivateKey
	var err error

	// First try to parse as PKCS1 private key
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		pkcs8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return "", fmt.Errorf("failed to parse private key (both PKCS1 and PKCS8 attempts failed): %v", err)
		}

		// Convert to RSA private key
		rsaKey, ok := pkcs8Key.(*rsa.PrivateKey)
		if !ok {
			return "", fmt.Errorf("parsed key is not an RSA private key")
		}
		privateKey = rsaKey
	}

	// Create SHA-256 hash of the data
	hash := sha256.New()
	hash.Write(data)
	hashed := hash.Sum(nil)

	// Sign the hashed data with PKCS1v15
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}

	// Encode the signature as base64
	return base64.StdEncoding.EncodeToString(signature), nil
}

// checkPublicKeyMatch compares the public key extracted from the private key with the expected public key
func checkPublicKeyMatch(privateKey *rsa.PrivateKey, expectedPubKeyStr string) (bool, error) {
	pubKeyBytes, _ := base64.StdEncoding.DecodeString(expectedPubKeyStr)
	expectedPub, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return false, fmt.Errorf("error parsing expected public key: %v", err)
	}

	expectedRsaPub, ok := expectedPub.(*rsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("expected public key is not an RSA public key")
	}

	// Compare the two public keys
	keysEqual := expectedRsaPub.N.Cmp(privateKey.PublicKey.N) == 0 && expectedRsaPub.E == privateKey.PublicKey.E
	return keysEqual, nil
}

func extractAdditionalHeaders(headers map[string]string, body string) {
	if body == "" {
		return
	}
	var payload struct {
		AdditionalInfo map[string]interface{} `json:"additionalInfo"`
	}
	if err := json.Unmarshal([]byte(body), &payload); err != nil {
		return // silently ignore parse errors
	}

	if info := payload.AdditionalInfo; info != nil {
		if v, ok := info["accessToken"].(string); ok && v != "" {
			headers[HeaderAuthorizationCustomer] = fmt.Sprintf("Bearer %s", v)
		}
		if v, ok := info["endUserIpAddress"].(string); ok && v != "" {
			headers[HeaderXIPAddress] = v
		}
		if v, ok := info["deviceId"].(string); ok && v != "" {
			headers[HeaderXDeviceID] = v
		}
		if v, ok := info["latitude"].(string); ok && v != "" {
			headers[HeaderXLatitude] = v
		}
		if v, ok := info["longitude"].(string); ok && v != "" {
			headers[HeaderXLongitude] = v
		}
	}
}
