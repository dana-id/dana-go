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

package test

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/dana-id/dana-go/test/fixtures"
	"github.com/dana-id/dana-go/widget/v1"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		return
	}

	// Navigate to the root dana-api-client-generator directory
	// From /sdk/dana-go-api-client/test we need to go up three levels
	rootDir := filepath.Dir(filepath.Dir(filepath.Dir(cwd)))

	// Load the .env file from dana-api-client-generator directory
	envPath := filepath.Join(rootDir, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Printf("Warning: Could not load .env file from %s: %v\n", envPath, err)
		log.Println("Using environment variables from the system.")
	} else {
		log.Printf("Successfully loaded environment variables from %s\n", envPath)
	}
}

func TestWidgetAPIService(t *testing.T) {
	// Get API client fixture once for all tests
	apiClient := fixtures.GetApiClient()

	t.Run("Test WidgetAPI Payment", func(t *testing.T) {
		// Skip test in CI environments with no credentials
		if apiClient.Client == nil {
			t.Skip("Skipping test: No API client credentials")
		}

		widgetPaymentRequest := fixtures.GetWidgetPaymentRequest()
		widgetPaymentResponse, httpRes, err := apiClient.Client.WidgetAPI.WidgetPayment(apiClient.Ctx).WidgetPaymentRequest(widgetPaymentRequest).Execute()

		require.NoError(t, err, "Error during WidgetPayment execution")
		require.NotNil(t, widgetPaymentResponse)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotEmpty(t, widgetPaymentResponse.WebRedirectUrl)
		assert.Equal(t, widgetPaymentRequest.PartnerReferenceNo, widgetPaymentResponse.PartnerReferenceNo)
	})

	t.Run("Test WidgetAPI QueryPayment", func(t *testing.T) {
		// Skip test in CI environments with no credentials
		if apiClient.Client == nil {
			t.Skip("Skipping test: No API client credentials")
		}

		// Create a payment first
		widgetPaymentRequest := fixtures.GetWidgetPaymentRequest()
		widgetPaymentResponse, httpRes, err := apiClient.Client.WidgetAPI.WidgetPayment(apiClient.Ctx).WidgetPaymentRequest(widgetPaymentRequest).Execute()

		require.NoError(t, err, "Failed to create payment before querying")
		require.Equal(t, 200, httpRes.StatusCode)

		// Test QueryPayment
		queryPaymentRequest := fixtures.GetWidgetQueryPaymentRequest(&widgetPaymentRequest, widgetPaymentResponse)
		queryPaymentResponse, httpRes, err := apiClient.Client.WidgetAPI.QueryPayment(apiClient.Ctx).QueryPaymentRequest(queryPaymentRequest).Execute()

		require.NoError(t, err)
		require.NotNil(t, queryPaymentResponse)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, widgetPaymentRequest.PartnerReferenceNo, *queryPaymentResponse.OriginalPartnerReferenceNo)
	})

	t.Run("Test WidgetAPI CancelOrder", func(t *testing.T) {
		// Skip test in CI environments with no credentials
		if apiClient.Client == nil {
			t.Skip("Skipping test: No API client credentials")
		}

		// Create a payment first
		widgetPaymentRequest := fixtures.GetWidgetPaymentRequest()
		_, httpRes, err := apiClient.Client.WidgetAPI.WidgetPayment(apiClient.Ctx).WidgetPaymentRequest(widgetPaymentRequest).Execute()

		require.NoError(t, err, "Failed to create payment before cancelling")
		require.Equal(t, 200, httpRes.StatusCode)

		// Test CancelOrder
		cancelOrderRequest := fixtures.GetWidgetCancelOrderRequest(&widgetPaymentRequest)
		cancelOrderResponse, httpRes, err := apiClient.Client.WidgetAPI.CancelOrder(apiClient.Ctx).CancelOrderRequest(cancelOrderRequest).Execute()

		require.NoError(t, err)
		require.NotNil(t, cancelOrderResponse)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, widgetPaymentRequest.PartnerReferenceNo, cancelOrderResponse.GetOriginalPartnerReferenceNo())
	})

	t.Run("Apply Token with Authorization Code", func(t *testing.T) {
		// Check Go version to skip OAuth tests on versions below 1.22
		goVersion := runtime.Version()
		isSupportedVersion := strings.HasPrefix(goVersion, "go1.22") || strings.HasPrefix(goVersion, "devel")
		if !isSupportedVersion {
			t.Skip("Skipping OAuth tests on Go version < 1.22, current version: " + goVersion)
			return
		}

		var authCode string

		log.Println("Starting OAuth automation flow...")

		// Run oauth_automation.go to get authorization code and set it to env variable DANA_AUTH_CODE
		authCode = os.Getenv("DANA_AUTH_CODE")

		require.NotEmpty(t, authCode, "Authorization code must not be empty")

		// Create Apply Token request with Authorization Code
		authCodeReq := widget.NewApplyTokenAuthorizationCodeRequest("AUTHORIZATION_CODE", authCode)
		applyTokenRequest := widget.ApplyTokenAuthorizationCodeRequestAsApplyTokenRequest(authCodeReq)

		partnerId := os.Getenv("X_PARTNER_ID")
		if partnerId == "" {
			log.Println("PARTNER_ID environment variable is not set")
			t.Skip("Skipping test: PARTNER_ID environment variable is not set")
			return
		}

		applyTokenResponse, httpRes, err := apiClient.Client.WidgetAPI.ApplyToken(apiClient.Ctx).ApplyTokenRequest(applyTokenRequest).Execute()
		require.NoError(t, err, "Apply Token request failed")
		require.NotNil(t, applyTokenResponse, "Apply Token response must not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "HTTP status code should be 200")

		assert.NotEmpty(t, applyTokenResponse.AccessToken, "AccessToken should not be empty")
		assert.NotEmpty(t, applyTokenResponse.RefreshToken, "RefreshToken should not be empty")

		accessToken := applyTokenResponse.AccessToken
		require.NotEmpty(t, accessToken, "Access token must not be empty")

		t.Run("Apply OTT with Access Token", func(t *testing.T) {
			applyOttRequest := fixtures.GetApplyOttRequest(accessToken)

			applyOttResponse, httpRes, err := apiClient.Client.WidgetAPI.ApplyOTT(apiClient.Ctx).ApplyOTTRequest(applyOttRequest).Execute()

			require.NoError(t, err, "Apply OTT request failed")
			require.NotNil(t, applyOttResponse, "Apply OTT response must not be nil")
			assert.Equal(t, 200, httpRes.StatusCode, "HTTP status code should be 200")

			require.NotEmpty(t, applyOttResponse.UserResources, "UserResources should not be empty")

			var ott string
			for _, resource := range applyOttResponse.UserResources {
				if resource.HasResourceType() && resource.GetResourceType() == "OTT" && resource.HasValue() {
					ott = resource.GetValue()
					break
				}
			}

			assert.NotEmpty(t, ott, "OTT value should be found in the response")
		})

		t.Run("Unbinding Account", func(t *testing.T) {
			unbindingAccountRequest := fixtures.GetUnbindingAccountRequest(accessToken)

			unbindingAccountResponse, httpRes, err := apiClient.Client.WidgetAPI.AccountUnbinding(apiClient.Ctx).AccountUnbindingRequest(unbindingAccountRequest).Execute()

			require.NoError(t, err, "Unbinding Account request failed")
			require.NotNil(t, unbindingAccountResponse, "Unbinding Account response must not be nil")
			assert.Equal(t, 200, httpRes.StatusCode, "HTTP status code should be 200")
			assert.Equal(t, "Successful", unbindingAccountResponse.ResponseMessage, "ResponseMessage should be Successful")
		})
	})
}
