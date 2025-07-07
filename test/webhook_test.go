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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/dana-id/dana-go/test/fixtures"
	"github.com/dana-id/dana-go/webhook"
	"github.com/stretchr/testify/assert"
)

func TestWebhookParser_ParseWebhook(t *testing.T) {
	webhookHttpMethod := "POST"
	webhookRelativeURL := "/v1.0/debit/notify"
	webhookBodyMap := map[string]interface{}{
		"originalPartnerReferenceNo": "TESTPN20240101001",
		"originalReferenceNo":        "TESTREF20240101001",
		"merchantId":                 "TESTMERCH001",
		"subMerchantId":              "TESTSUBMERCH001",
		"amount": map[string]string{
			"value":    "15000.00",
			"currency": "IDR",
		},
		"latestTransactionStatus": "00",
		"transactionStatusDesc":   "Success",
		"createdTime":             "2024-01-01T10:00:00+07:00",
		"finishedTime":            "2024-01-01T10:00:05+07:00",
	}

	webhookBodyBytes, err := json.Marshal(webhookBodyMap)
	assert.NoError(t, err, "Error marshaling webhook body")
	webhookBodyStr := string(webhookBodyBytes)

	minifiedWebhookBodyStr, err := webhook.MinifyJSON(webhookBodyStr)
	assert.NoError(t, err, "Error minifying webhook body")

	generatedHeaders, err := fixtures.GenerateSnapAuthHeaders(webhookHttpMethod, webhookRelativeURL, minifiedWebhookBodyStr)
	assert.NoError(t, err, "Error generating SNAP headers")
	assert.NotEmpty(t, generatedHeaders["X-SIGNATURE"], "X-SIGNATURE should not be empty")
	assert.NotEmpty(t, generatedHeaders["X-TIMESTAMP"], "X-TIMESTAMP should not be empty")

	publicKey := os.Getenv("WEBHOOK_PUBLIC_KEY")
	parser, err := webhook.NewWebhookParser(&publicKey, nil)
	assert.NoError(t, err, "Error creating WebhookParser")
	assert.NotNil(t, parser, "WebhookParser should not be nil")

	reqURL, err := url.Parse(webhookRelativeURL)
	assert.NoError(t, err, "Error parsing webhook URL")

	req, err := http.NewRequest(
		webhookHttpMethod,
		reqURL.String(),
		bytes.NewBufferString(webhookBodyStr),
	)
	assert.NoError(t, err, "Error creating http.Request")

	req.Header.Set("X-SIGNATURE", generatedHeaders["X-SIGNATURE"])
	req.Header.Set("X-TIMESTAMP", generatedHeaders["X-TIMESTAMP"])

	extractedMethod := req.Method
	extractedPath := req.URL.Path

	bodyBytes, err := ioutil.ReadAll(req.Body)
	assert.NoError(t, err, "Error reading request body for test")
	defer req.Body.Close()
	extractedBodyStr := string(bodyBytes)

	parsedData, err := parser.ParseWebhook(
		extractedMethod,
		extractedPath,
		req.Header,
		extractedBodyStr,
	)
	assert.NoError(t, err, "Webhook parsing/verification failed")
	if err == nil {
		assert.NotNil(t, parsedData, "Parsed data should not be nil")

		assert.Equal(t, webhookBodyMap["originalPartnerReferenceNo"], parsedData.OriginalPartnerReferenceNo)
		assert.Equal(t, webhookBodyMap["originalReferenceNo"], parsedData.OriginalReferenceNo)
		amountMap := webhookBodyMap["amount"].(map[string]string)
		assert.Equal(t, amountMap["value"], parsedData.Amount.Value)
		assert.Equal(t, amountMap["currency"], parsedData.Amount.Currency)
		assert.Equal(t, webhookBodyMap["latestTransactionStatus"], parsedData.LatestTransactionStatus)
	}
}
