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
	"crypto/sha256"
	"fmt"
	"os"
	"time"

	payment_gateway "github.com/dana-id/dana-go/payment_gateway/v1"
	utils "github.com/dana-id/dana-go/utils"
	"github.com/google/uuid"
)

// GetMerchantId returns the merchant ID from environment or a default value
func GetMerchantId() string {
	merchantId := os.Getenv("MERCHANT_ID")
	if merchantId == "" {
		merchantId = "MERCHANT_ID_DEFAULT"
	}
	return merchantId
}

// GetConsultPayRequest returns a fixture for ConsultPayRequest
func GetConsultPayRequest() payment_gateway.ConsultPayRequest {
	return payment_gateway.ConsultPayRequest{
		MerchantId: GetMerchantId(),
		Amount: payment_gateway.Money{
			Value:    "12345678.00",
			Currency: "IDR",
		},
		// Just basic fields based on what we've seen in the test
	}
}

// GetCreateOrderByApiRequest returns a fixture for CreateOrderByApiRequest
func GetCreateOrderByApiRequest() payment_gateway.CreateOrderByApiRequest {
	// Generate a unique partner reference number based on timestamp
	partnerRefNo := time.Now().Format(time.RFC3339)

	// Calculate valid date 1 day from now
	validUpTo := time.Now().Add(24 * time.Hour).Format("2006-01-02T15:04:05-07:00")

	// Helper for string pointer
	strPtr := func(s string) *string {
		return &s
	}

	// Create request with essential fields
	request := payment_gateway.CreateOrderByApiRequest{
		PartnerReferenceNo: partnerRefNo,
		MerchantId:         GetMerchantId(),
		Amount: payment_gateway.Money{
			Value:    "222000.00",
			Currency: "IDR",
		},
		ValidUpTo: strPtr(validUpTo),
	}

	// Add URL parameters
	request.UrlParams = []payment_gateway.UrlParam{
		{
			Url:        "https://example.com/return",
			Type:       "PAY_RETURN",
			IsDeeplink: "N",
		},
		{
			Url:        "https://example.com/notify",
			Type:       "NOTIFICATION",
			IsDeeplink: "N",
		},
	}

	// Add payment options
	request.PayOptionDetails = []payment_gateway.PayOptionDetail{
		{
			PayMethod: string(payment_gateway.PAYMETHOD_BALANCE_),
			PayOption: "",
			TransAmount: payment_gateway.Money{
				Value:    "222000.00",
				Currency: "IDR",
			},
		},
	}

	// Add additional info
	envInfo := payment_gateway.EnvInfo{
		SourcePlatform: string(payment_gateway.SOURCEPLATFORM_IPG_),
		TerminalType:   string(payment_gateway.TERMINALTYPE_SYSTEM_),
	}

	orderInfo := payment_gateway.OrderApiObject{
		OrderTitle:        "Test Product",
		Scenario:          strPtr("API"),
		MerchantTransType: strPtr("SPECIAL_MOVIE"),
	}

	request.AdditionalInfo = &payment_gateway.CreateOrderByApiAdditionalInfo{
		Order:   &orderInfo,
		Mcc:     "5732",
		EnvInfo: envInfo,
	}

	return request
}

// GetCreateOrderByApiPaidRequest returns a fixture for CreateOrderByApiRequest
func GetCreateOrderByApiPaidRequest() payment_gateway.CreateOrderByApiRequest {
	// Generate a unique partner reference number based on timestamp
	partnerRefNo := uuid.New().String()

	// Calculate valid date 1 day from now
	validUpTo := time.Now().Add(24 * time.Hour).Format("2026-01-02T15:04:05-07:00")

	// Helper for string pointer
	strPtr := func(s string) *string {
		return &s
	}

	// Create request with essential fields
	request := payment_gateway.CreateOrderByApiRequest{
		PartnerReferenceNo: partnerRefNo,
		MerchantId:         GetMerchantId(),
		Amount: payment_gateway.Money{
			Value:    "50001.00",
			Currency: "IDR",
		},
		ValidUpTo: strPtr(validUpTo),
	}

	// Add URL parameters
	request.UrlParams = []payment_gateway.UrlParam{
		{
			Url:        "https://example.com/return",
			Type:       "PAY_RETURN",
			IsDeeplink: "N",
		},
		{
			Url:        "https://example.com/notify",
			Type:       "NOTIFICATION",
			IsDeeplink: "N",
		},
	}

	// Add payment options
	request.PayOptionDetails = []payment_gateway.PayOptionDetail{
		{
			PayMethod: string(payment_gateway.PAYMETHOD_NETWORK_PAY_),
			PayOption: string(payment_gateway.PAYOPTION_NETWORK_PAY_PG_GOPAY_),
			TransAmount: payment_gateway.Money{
				Value:    "50001.00",
				Currency: "IDR",
			},
			FeeAmount: &payment_gateway.Money{
				Value:    "1000.00",
				Currency: "IDR",
			},
			AdditionalInfo: &payment_gateway.PayOptionAdditionalInfo{
				PhoneNumber: strPtr("08123456789"),
			},
		},
	}

	// Add additional info
	envInfo := payment_gateway.EnvInfo{
		SourcePlatform: string(payment_gateway.SOURCEPLATFORM_IPG_),
		TerminalType:   string(payment_gateway.TERMINALTYPE_SYSTEM_),
	}

	orderInfo := payment_gateway.OrderApiObject{
		OrderTitle:        "Test Product",
		Scenario:          strPtr("API"),
		MerchantTransType: strPtr("SPECIAL_MOVIE"),
	}

	request.AdditionalInfo = &payment_gateway.CreateOrderByApiAdditionalInfo{
		Order:   &orderInfo,
		Mcc:     "5732",
		EnvInfo: envInfo,
	}

	return request
}

// GetQueryPaymentRequest returns a fixture for QueryPaymentRequest
func GetQueryPaymentRequest(orderRequest *payment_gateway.CreateOrderByApiRequest, order *payment_gateway.CreateOrderResponse) payment_gateway.QueryPaymentRequest {
	// Helper for string pointer
	strPtr := func(s string) *string {
		return &s
	}

	return payment_gateway.QueryPaymentRequest{
		// Just use simple string values
		ServiceCode:                "54", // Payment Gateway service code
		MerchantId:                 GetMerchantId(),
		OriginalPartnerReferenceNo: strPtr(orderRequest.GetPartnerReferenceNo()),
		OriginalReferenceNo:        strPtr(order.GetReferenceNo()),
	}
}

// GetCancelOrderRequest returns a fixture for CancelOrderRequest
func GetCancelOrderRequest(orderRequest *payment_gateway.CreateOrderByApiRequest) payment_gateway.CancelOrderRequest {
	// Helper for string pointer
	strPtr := func(s string) *string {
		return &s
	}

	return payment_gateway.CancelOrderRequest{
		MerchantId:                 GetMerchantId(),
		OriginalPartnerReferenceNo: orderRequest.GetPartnerReferenceNo(),
		Reason:                     strPtr("Test cancellation"),
	}
}

// GetRefundOrderRequest returns a fixture for RefundOrderRequest
func GetRefundOrderRequest(orderRequest *payment_gateway.CreateOrderByApiRequest) payment_gateway.RefundOrderRequest {
	// Generate a unique partner refund number based on timestamp
	// partnerRefundNo := "REFUND-" + time.Now().Format("20060102150405")

	// Helper for string pointer
	strPtr := func(s string) *string {
		return &s
	}

	return payment_gateway.RefundOrderRequest{
		MerchantId:                 GetMerchantId(),
		OriginalPartnerReferenceNo: orderRequest.GetPartnerReferenceNo(),
		PartnerRefundNo:            orderRequest.GetPartnerReferenceNo(),
		RefundAmount: payment_gateway.Money{
			Value:    "50001.00",
			Currency: "IDR",
		},
		Reason: strPtr("Customer complain"),
		// Add additional info as a map with the required fields for the special case
		AdditionalInfo: &payment_gateway.RefundOrderRequestAdditionalInfo{
			PayoutAccountNo: strPtr("20050000000001503276"),
			ActorType:       strPtr("BACK_OFFICE"),
		},
	}
}

// GenerateSnapAuthHeaders creates the X-SIGNATURE and X-TIMESTAMP headers for a request.
func GenerateSnapAuthHeaders(method, resourcePath, body string) (map[string]string, error) {
	jkt, errLoadLocation := time.LoadLocation("Asia/Jakarta")
	var jktTime time.Time
	if errLoadLocation != nil {
		now := time.Now().UTC()
		jktTime = now.Add(7 * time.Hour)
	} else {
		jktTime = time.Now().In(jkt)
	}
	timestamp := jktTime.Format("2006-01-02T15:04:05+07:00")

	bodyHasher := sha256.New()
	bodyHasher.Write([]byte(body))
	hashedPayload := fmt.Sprintf("%x", bodyHasher.Sum(nil))

	stringToSign := fmt.Sprintf("%s:%s:%s:%s", method, resourcePath, hashedPayload, timestamp)

	signature, err := utils.SignWithPrivateKey([]byte(stringToSign), []byte(os.Getenv("WEBHOOK_PRIVATE_KEY")))
	if err != nil {
		return nil, fmt.Errorf("GenerateSnapAuthHeaders: failed to sign data: %w", err)
	}

	headers := map[string]string{
		"X-TIMESTAMP": timestamp,
		"X-SIGNATURE": signature,
	}
	return headers, nil
}
