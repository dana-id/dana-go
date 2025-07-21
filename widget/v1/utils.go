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

package widget

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	
	"github.com/dana-id/dana-go/utils"
)

func GenerateOauthUrl(data *Oauth2UrlData) (string, error) {
	// Use environment variable or default to sandbox
	env := os.Getenv("ENV")
	if env == "" {
		env = "sandbox"
	}

	// Determine base URL based on environment
	var baseUrl string
	if env == "sandbox" {
		baseUrl = "https://m.sandbox.dana.id/n/ipg/oauth"
	} else {
		baseUrl = "https://m.dana.id/n/ipg/oauth"
	}

	// Get partner ID from environment
	partnerId := os.Getenv("X_PARTNER_ID")
	if partnerId == "" {
		return "", fmt.Errorf("X_PARTNER_ID is not defined")
	}

	// Use provided state or generate a new one
	var state string
	if data.State != nil && *data.State != "" {
		state = *data.State
	} else {
		uuidStr := uuid.New().String()
		state = uuidStr
	}

	// Generate channel ID
	channelId := GenerateChannelId()

	// Generate scopes
	scopes := GenerateScopes()

	// Generate external ID
	externalId := GenerateExternalId(data.ExternalId)

	merchantId := data.MerchantId

	// Generate timestamp in Jakarta time format
	timestamp := GenerateTimestamp()

	// Build URL with required parameters
	urlParams := url.Values{}
	urlParams.Add("partnerId", partnerId)
	urlParams.Add("scopes", scopes)
	urlParams.Add("externalId", externalId)
	urlParams.Add("channelId", channelId)
	urlParams.Add("redirectUrl", data.RedirectUrl)
	urlParams.Add("timestamp", timestamp)
	urlParams.Add("state", state)
	urlParams.Add("isSnapBI", "true")

	if merchantId != "" {
		urlParams.Add("merchantId", merchantId)
	}

	if data.SubMerchantId != nil && *data.SubMerchantId != "" {
		urlParams.Add("subMerchantId", *data.SubMerchantId)
	}

	if data.Lang != nil && *data.Lang != "" {
		urlParams.Add("lang", *data.Lang)
	}

	if data.AllowRegistration != nil && *data.AllowRegistration != "" {
		urlParams.Add("allowRegistration", *data.AllowRegistration)
	}

	if data.SeamlessData != nil {
		// Convert seamlessData object to JSON string
		seamlessDataBytes, err := json.Marshal(data.SeamlessData)
		if err != nil {
			return "", fmt.Errorf("error marshaling seamlessData: %v", err)
		}
		urlParams.Add("seamlessData", string(seamlessDataBytes))

		// Get private key from environment
		privateKey := os.Getenv("PRIVATE_KEY")
		if privateKey != "" {
			// Calculate seamlessSign if private key is available
			seamlessSign, err := GenerateSeamlessSign(data.SeamlessData, privateKey)
			if err != nil {
				return "", fmt.Errorf("error generating seamlessSign: %v", err)
			}
			if seamlessSign != "" {
				urlParams.Add("seamlessSign", seamlessSign)
			}
		}
	}

	return fmt.Sprintf("%s?%s", baseUrl, urlParams.Encode()), nil
}

// GenerateChannelId generates a channel ID
func GenerateChannelId() string {
	return fmt.Sprintf("widget-%d", time.Now().UnixNano())
}

// GenerateScopes generates the OAuth scopes
func GenerateScopes() string {
	return "APP_AUTH,QUERY_BALANCE"
}

// GenerateExternalId generates an external ID based on input or creates a new one
func GenerateExternalId(input string) string {
	if input != "" {
		return input
	}
	// Generate a new UUID as string
	uuidStr := uuid.New().String()
	return uuidStr
}

// GenerateTimestamp generates a timestamp in Jakarta time format
func GenerateTimestamp() string {
	// Jakarta is UTC+7
	jakarta, _ := time.LoadLocation("Asia/Jakarta")
	if jakarta == nil {
		// Fallback if location is not available
		now := time.Now().UTC().Add(7 * time.Hour)
		return strings.Replace(now.Format(time.RFC3339), "Z", "+07:00", 1)
	}

	now := time.Now().In(jakarta)
	return strings.Replace(now.Format(time.RFC3339), "Z", "+07:00", 1)
}

// GenerateSeamlessSign generates a signature for the seamless data
func GenerateSeamlessSign(seamlessData interface{}, privateKey string) (string, error) {
	// If seamlessData is nil, return empty string
	if seamlessData == nil {
		return "", nil
	}

	// Check if privateKey is provided
	if privateKey == "" {
		return "", fmt.Errorf("private key is required for generating seamless sign")
	}

	// Convert seamlessData to JSON string
	seamlessDataBytes, err := json.Marshal(seamlessData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal seamless data: %v", err)
	}

	// Format the private key if needed
	formattedPrivateKey := utils.ConvertToPEM(privateKey, "PRIVATE")

	// Sign the data using RSA-SHA256
	signature, err := signData(seamlessDataBytes, []byte(formattedPrivateKey))
	if err != nil {
		return "", fmt.Errorf("failed to calculate seamlessSign: %v", err)
	}

	// URL encode the signature
	return url.QueryEscape(signature), nil
}

// signData signs data with the given private key using RSA-SHA256
func signData(data []byte, privateKeyPEM []byte) (string, error) {
	// Parse the PEM encoded private key
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing private key")
	}

	// Try parsing as PKCS1 first
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// Try parsing as PKCS8
		pkcs8Key, pkcs8Err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if pkcs8Err != nil {
			return "", fmt.Errorf("failed to parse private key: %v", pkcs8Err)
		}
		
		// Convert to RSA private key
		rsaKey, ok := pkcs8Key.(*rsa.PrivateKey)
		if !ok {
			return "", fmt.Errorf("parsed key is not an RSA private key")
		}
		privateKey = rsaKey
	}

	// Create hash
	hash := sha256.New()
	hash.Write(data)
	hashed := hash.Sum(nil)

	// Sign the hash
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}

	// Convert to Base64
	return base64.StdEncoding.EncodeToString(signature), nil
}
