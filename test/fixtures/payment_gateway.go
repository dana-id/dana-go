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

	payment_gateway "github.com/dana-id/go_client/payment_gateway/v1"
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
			PayMethod: string(payment_gateway.PAYMETHOD_VIRTUAL_ACCOUNT_),
			PayOption: string(payment_gateway.PAYOPTION_VIRTUAL_ACCOUNT_BNI_),
			TransAmount: payment_gateway.Money{
				Value:    "222000.00",
				Currency: "IDR",
			},
		},
	}

	// Add additional info
	envInfo := payment_gateway.EnvInfo{
		SourcePlatform: "IPG",
		TerminalType:   "SYSTEM",
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
// func GetRefundOrderRequest(orderRequest *payment_gateway.CreateOrderByApiRequest) payment_gateway.RefundOrderRequest {
// 	// Generate a unique partner refund number based on timestamp
// 	partnerRefundNo := "REFUND-" + time.Now().Format("20060102150405")

// 	// Helper for string pointer
// 	strPtr := func(s string) *string {
// 		return &s
// 	}

// 	return payment_gateway.RefundOrderRequest{
// 		MerchantId:                 GetMerchantId(),
// 		OriginalPartnerReferenceNo: orderRequest.GetPartnerReferenceNo(),
// 		PartnerRefundNo:            partnerRefundNo,
// 		RefundAmount: payment_gateway.Money{
// 			Value:    "10000.00",
// 			Currency: "IDR",
// 		},
// 		Reason: strPtr("Test refund"),
// 		// Add additional info as a map with the required fields for the special case
// 		AdditionalInfo: map[string]interface{}{
// 			"payoutAccountNo": "20050000000001503276",
// 			"actorType":       "BACK_OFFICE",
// 			"deviceId":       "TEST-DEVICE-ID",
// 			"terminalType":   "WEB",
// 		},
// 	}
// }
