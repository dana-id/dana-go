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

package fixtures

import (
	"os"
	"time"

	"github.com/google/uuid"
	utils "github.com/dana-id/dana-go/utils"
	widget "github.com/dana-id/dana-go/widget/v1"
)

// GetWidgetPaymentRequest returns a test WidgetPaymentRequest
func GetWidgetPaymentRequest() widget.WidgetPaymentRequest {
	order := widget.Order{
		OrderTitle:  "DANA Widget Test Order",
		CreatedTime: utils.PtrString(time.Now().Format("2006-01-02T15:04:05+07:00")),
	}
	envInfo := widget.NewEnvInfo(string(widget.SOURCEPLATFORM_IPG_), string(widget.TERMINALTYPE_SYSTEM_))
	additionalInfo := widget.NewWidgetPaymentRequestAdditionalInfo(
		"51051000100000000001", // ProductCode
		order,                  // Order
		"5732",                 // Mcc
		*envInfo,               // EnvInfo
	)

	amount := widget.NewMoney("10000.00", "IDR")
	partnerRefNo := uuid.New().String()

	request := widget.NewWidgetPaymentRequest(
		partnerRefNo,    // PartnerReferenceNo
		GetMerchantId(), // MerchantId
		*amount,
		*additionalInfo,
	)
	return *request
}

// GetWidgetQueryPaymentRequest returns a test QueryPaymentRequest
func GetWidgetQueryPaymentRequest(req *widget.WidgetPaymentRequest, resp *widget.WidgetPaymentResponse) widget.QueryPaymentRequest {
	strPtr := func(s string) *string {
		return &s
	}

	return widget.QueryPaymentRequest{
		// Just use simple string values
		ServiceCode:                "54",
		MerchantId:                 GetMerchantId(),
		OriginalPartnerReferenceNo: strPtr(req.PartnerReferenceNo),
		OriginalReferenceNo:        strPtr(*resp.ReferenceNo),
	}
}

// GetWidgetCancelOrderRequest returns a test CancelOrderRequest
func GetWidgetCancelOrderRequest(req *widget.WidgetPaymentRequest) widget.CancelOrderRequest {
	return widget.CancelOrderRequest{
		OriginalPartnerReferenceNo: req.PartnerReferenceNo,
		MerchantId:                 req.MerchantId,
		Reason:                     utils.PtrString("User cancelled order"),
	}
}

// GetApplyTokenRequest returns a test ApplyToken request using authorization code
func GetApplyTokenRequest() widget.ApplyTokenRequest {
	// Use sample authorization code or from environment
	authCode := os.Getenv("DANA_AUTH_CODE")
	if authCode == "" {
		// Use a default value for testing if not available
		// In a real scenario, this would come from user authorization flow
		authCode = "sample-auth-code-123"
	}

	// Create the authorization code request
	authCodeReq := widget.NewApplyTokenAuthorizationCodeRequest("AUTHORIZATION_CODE", authCode)

	// Convert to ApplyTokenRequest
	return widget.ApplyTokenAuthorizationCodeRequestAsApplyTokenRequest(authCodeReq)
}

// GetApplyOttRequest returns a test ApplyOTT request using the provided access token
func GetApplyOttRequest(accessToken string) widget.ApplyOTTRequest {
	// Device ID for the request
	deviceId := "test-device-id-" + time.Now().Format("20060102150405")

	// Create the additional info object with required fields
	additionalInfo := widget.NewApplyOTTRequestAdditionalInfo(accessToken, deviceId)

	// Set optional fields if needed
	endUserIp := "127.0.0.1"
	additionalInfo.SetEndUserIpAddress(endUserIp)

	// Create the main request with user resources and additional info
	userResources := []string{"OTT"}
	return *widget.NewApplyOTTRequest(userResources, *additionalInfo)
}

func GetUnbindingAccountRequest(accessToken string) widget.AccountUnbindingRequest {
	return widget.AccountUnbindingRequest{
		MerchantId: GetMerchantId(),
		AdditionalInfo: widget.AccountUnbindingRequestAdditionalInfo{
			AccessToken:      accessToken,
			DeviceId:         "test-device-id-" + time.Now().Format("20060102150405"),
			EndUserIpAddress: utils.PtrString("127.0.0.1"),
			Latitude:         utils.PtrString("-6.175"),
			Longitude:        utils.PtrString("106.865"),
		},
	}
}
