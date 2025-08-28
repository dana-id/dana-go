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
	mathrand "math/rand"
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

	// Determine base URL based on environment and mode
	var baseUrl string
	mode := "API" // Default mode
	if data.Mode != nil {
		mode = *data.Mode
	}
	
	if mode == "DEEPLINK" {
		if env == "sandbox" {
			baseUrl = "https://m.sandbox.dana.id/n/link/binding"
		} else {
			baseUrl = "https://link.dana.id/bindSnap"
		}
	} else { // API mode
		if env == "sandbox" {
			baseUrl = "https://m.sandbox.dana.id/n/ipg/oauth"
		} else {
			baseUrl = "https://m.dana.id/n/ipg/oauth"
		}
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

	// Generate scopes
	scopes := GenerateScopes()

	// Generate external ID
	externalId := GenerateExternalId(data.ExternalId)

	merchantId := data.MerchantId
	if merchantId == "" {
		merchantId = os.Getenv("MERCHANT_ID")
	}

	// Generate timestamp in Jakarta time format
	timestamp := GenerateTimestamp()

	// Build URL with required parameters
	var params []string
	
	// Variable to store requestId for DEEPLINK mode
	var requestId string

	// Add common parameters
	params = append(params, "partnerId="+partnerId)
	params = append(params, "scopes="+scopes)
	params = append(params, "externalId="+externalId)
	
	if mode == "DEEPLINK" {
		// Generate UUID request ID for DEEPLINK mode
		requestId = uuid.New().String()
		
		// Add DEEPLINK specific parameters
		params = append(params, "terminalType=WEB")
		params = append(params, "requestId="+requestId)
		params = append(params, "redirectUrl="+url.QueryEscape(data.RedirectUrl))
		params = append(params, "state="+state)
	} else { // API mode
		// Generate channel ID for API mode
		channelId := GenerateChannelId()
		
		// Add API specific parameters
		params = append(params, "channelId="+channelId)
		params = append(params, "redirectUrl="+url.QueryEscape(data.RedirectUrl))
		params = append(params, "timestamp="+url.QueryEscape(timestamp))
		params = append(params, "state="+state)
		params = append(params, "isSnapBI=true")
	}

	if merchantId != "" {
		params = append(params, "merchantId="+merchantId)
	}

	if data.SubMerchantId != nil && *data.SubMerchantId != "" {
		params = append(params, "subMerchantId="+*data.SubMerchantId)
	}

	if data.Lang != nil && *data.Lang != "" {
		params = append(params, "lang="+*data.Lang)
	}

	if data.AllowRegistration != nil && *data.AllowRegistration != "" {
		params = append(params, "allowRegistration="+*data.AllowRegistration)
	}

	if data.SeamlessData != nil {
		// Use a map for seamless data modifications
		seamlessDataMap := make(map[string]interface{})
		rawData := data.SeamlessData
		
		// Convert to JSON and back to ensure we have a map
		seamlessDataBytes, err := json.Marshal(rawData)
		if err != nil {
			return "", fmt.Errorf("error marshaling seamlessData: %v", err)
		}
		
		err = json.Unmarshal(seamlessDataBytes, &seamlessDataMap)
		if err != nil {
			return "", fmt.Errorf("error unmarshaling seamlessData: %v", err)
		}
		
		// Handle mobile/mobileNumber conversion
		if mobileNumber, ok := seamlessDataMap["mobileNumber"]; ok {
			seamlessDataMap["mobile"] = mobileNumber
			seamlessDataMap["mobileNumber"] = nil
			delete(seamlessDataMap, "mobileNumber")
		}
		
		// Add DEEPLINK specific data
		if mode == "DEEPLINK" && requestId != "" {
			seamlessDataMap["externalUid"] = externalId
			seamlessDataMap["reqTime"] = timestamp
			seamlessDataMap["verifiedTime"] = "0"
			seamlessDataMap["reqMsgId"] = requestId
		}
		
		// Convert seamlessData object to JSON string
		seamlessDataBytes, err = json.Marshal(seamlessDataMap)
		if err != nil {
			return "", fmt.Errorf("error marshaling seamlessData: %v", err)
		}
		params = append(params, "seamlessData="+url.QueryEscape(string(seamlessDataBytes)))

		// Get private key from environment
		privateKey := os.Getenv("PRIVATE_KEY")
		if privateKey != "" {
			// Calculate seamlessSign if private key is available
			seamlessSign, err := GenerateSeamlessSign(seamlessDataMap, privateKey)
			if err != nil {
				return "", fmt.Errorf("error generating seamlessSign: %v", err)
			}
			if seamlessSign != "" {
				params = append(params, "seamlessSign="+seamlessSign)
			}
		}
	}

	// Join all parameters with & and append to base URL
	queryString := strings.Join(params, "&")
	return fmt.Sprintf("%s?%s", baseUrl, queryString), nil
}

// GenerateChannelId generates a channel ID in Jakarta time format (GMT+7): YYYYMMDDHHmmssSSSnnnnnnn
func GenerateChannelId() string {
	// Get current time
	now := time.Now()

	// Add 7 hours to get Jakarta time (GMT+7)
	jakartaTime := now.Add(7 * time.Hour)

	// Format: YYYYMMDDHHmmssSSSnnnnnnn
	year := jakartaTime.Format("2006")
	month := jakartaTime.Format("01")
	day := jakartaTime.Format("02")
	hours := jakartaTime.Format("15")
	minutes := jakartaTime.Format("04")
	seconds := jakartaTime.Format("05")

	// Go doesn't directly format milliseconds with padded zeros, so we extract and format them
	milliseconds := fmt.Sprintf("%03d", jakartaTime.Nanosecond()/1000000)

	// Generate a random 7-digit number for the nanopart
	mathrand.Seed(time.Now().UnixNano())
	nanopart := fmt.Sprintf("%07d", mathrand.Intn(10000000))

	return fmt.Sprintf("%s%s%s%s%s%s%s%s", year, month, day, hours, minutes, seconds, milliseconds, nanopart)
}

// GenerateScopes generates the OAuth scopes based on environment
func GenerateScopes() string {
	// Use environment variable or default to sandbox
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("DANA_ENV")
	}
	if env == "" {
		env = "sandbox"
	}

	// Return different scopes based on environment
	if strings.ToLower(env) != "production" {
		return "CASHIER,AGREEMENT_PAY,QUERY_BALANCE,DEFAULT_BASIC_PROFILE,MINI_DANA"
	} else {
		return "CASHIER"
	}
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
	jakarta, _ := time.LoadLocation("Asia/Jakarta")
	if jakarta == nil {
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
	return signature, nil
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

// GenerateCompletePaymentUrl combines the webRedirectUrl from WidgetPaymentResponse with the OTT token from ApplyOTTResponse
func GenerateCompletePaymentUrl(widgetPaymentResponse *WidgetPaymentResponse, applyOTTResponse *ApplyOTTResponse) string {
	// Check for nil pointers
	if widgetPaymentResponse == nil || applyOTTResponse == nil {
		return ""
	}

	// Get the WebRedirectUrl value safely
	webRedirectUrl := widgetPaymentResponse.GetWebRedirectUrl()

	// Check if there are user resources and safely access the first one's value
	if len(applyOTTResponse.UserResources) == 0 || applyOTTResponse.UserResources[0].Value == nil {
		return webRedirectUrl
	}

	ott := *applyOTTResponse.UserResources[0].Value

	// Combine the URL with the OTT token
	return fmt.Sprintf("%s&ott=%s", webRedirectUrl, ott)
}
