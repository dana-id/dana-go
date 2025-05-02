package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dana-id/dana-go-api-client/config"
	uuid "github.com/google/uuid"
)

// SetSnapHeaders applies all Snap API authentication headers to the provided headers map
// This centralized function helps avoid duplicate authentication code across API methods
func SetSnapHeaders(headerParams map[string]string, apiKey *config.APIKey, body string, method string, endpoint string) {

	if apiKey == nil {
		ReportError("API key is required for authentication")
		return
	}

	// Set ORIGIN header
	if apiKey.ORIGIN != "" {
		headerParams["ORIGIN"] = apiKey.ORIGIN
	}

	// Set X-PARTNER-ID header
	if apiKey.X_PARTNER_ID != "" {
		headerParams["X-PARTNER-ID"] = apiKey.X_PARTNER_ID
	}

	// Set CHANNEL-ID header
	if apiKey.CHANNEL_ID != "" {
		headerParams["CHANNEL-ID"] = apiKey.CHANNEL_ID
	}

	runtimeHeaders := getRuntimeHeaders(body, apiKey, method, endpoint)

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

func getRuntimeHeaders(body string, apiKey *config.APIKey, method string, endpoint string) map[string]string {
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

	// Format in ISO 8601 with fixed +07:00 timezone indicator for Jakarta
	return map[string]string{
		"X-TIMESTAMP":   timestamp,
		"X-EXTERNAL-ID": uuid.New().String(),
		"X-SIGNATURE":   generateSignature(body, apiKey, method, endpoint, timestamp),
	}
}

func generateSignature(body string, apiKey *config.APIKey, method string, endpoint string, timestamp string) string {
	// Validate we have the required data
	if apiKey == nil {
		return ""
	}

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
	urlPath = editPath(urlPath)

	// Format the string to sign:
	// "<HTTP METHOD>:<RELATIVE PATH URL>:<LOWERCASE_HEX_ENCODED_SHA_256(MINIFIED_HTTP_BODY)>:<X-TIMESTAMP>"
	data := fmt.Sprintf("%s:%s:%s:%s", method, urlPath, hashedPayload, timestamp)

	// Get usable private key
	privateKeyData, err := getUsablePrivateKey(apiKey)
	if err != nil || privateKeyData == nil {
		fmt.Println("Failed to get usable private key:", err)
		return ""
	}

	// Generate signature
	signature, err := signWithPrivateKey([]byte(data), privateKeyData)
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

// getUsablePrivateKey retrieves and prepares the private key from either direct input or file path
// Private key path is the preferred method as it avoids issues with newlines and formatting
func getUsablePrivateKey(apiKey *config.APIKey) ([]byte, error) {
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

// signWithPrivateKey generates an RSA signature using the provided private key
func signWithPrivateKey(data []byte, privateKeyPEM []byte) (string, error) {
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

func editPath(path string) string {
	RESOURCE_PATH_TO_EDIT := []string{"/payment-gateway/v1.0/debit/status.htm"}

	if StrContains(RESOURCE_PATH_TO_EDIT, path) {
		return strings.Replace(path, "/payment-gateway/v1.0", "/v1.0", 1)
	}

	return path
}
