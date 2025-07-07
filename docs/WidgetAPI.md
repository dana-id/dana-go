# WidgetAPI

All URIs are relative to http://api.sandbox.dana.id for sandbox environment and https://api.saas.dana.id for production environment

List of available operations:

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountUnbinding**](WidgetAPI.md#AccountUnbinding) | **Post** /v1.0/registration-account-unbinding.htm | Account unbinding - Binding
[**ApplyOTT**](WidgetAPI.md#ApplyOTT) | **Post** /rest/v1.1/qr/apply-ott | Apply OTT - Widget
[**ApplyToken**](WidgetAPI.md#ApplyToken) | **Post** /v1.0/access-token/b2b2c.htm | Apply Token, required by Apply OTT - Binding
[**CancelOrder**](WidgetAPI.md#CancelOrder) | **Post** /v1.0/debit/cancel.htm | Cancel Order - Widget
[**QueryPayment**](WidgetAPI.md#QueryPayment) | **Post** /rest/v1.1/debit/status | Query Payment - Widget
[**RefundOrder**](WidgetAPI.md#RefundOrder) | **Post** /v1.0/debit/refund.htm | Refund Order - Widget
[**WidgetPayment**](WidgetAPI.md#WidgetPayment) | **Post** /rest/redirection/v1.0/debit/payment-host-to-host | Widget Payment - Widget


## Additional Documentation

* [Enum Types](#enum-types) - List of available enum constants
* [Webhook Parser](#webhookparser) - Webhook handling documentation


## AccountUnbinding

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	widget "github.com/dana-id/dana-go/widget/v1"
)

func main() {
	
	// Configuring api client
	// Api client should be singleton, can reuse the apiClient for multiple requests in various operations
	configuration := config.NewConfiguration()
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := widget.AccountUnbindingRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}
	_, r, err := apiClient.WidgetAPI.AccountUnbinding(context.Background()).AccountUnbindingRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetAPI.AccountUnbinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountUnbinding`: AccountUnbindingResponse
	fmt.Fprintf(os.Stdout, "Response from `WidgetAPI.AccountUnbinding`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to AccountUnbindingRequest (struct with type ApiAccountUnbindingRequest)

[**AccountUnbindingRequest**](Widget/AccountUnbindingRequest.md)

### Return type

[**AccountUnbindingResponse**](Widget/AccountUnbindingResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## ApplyOTT

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	widget "github.com/dana-id/dana-go/widget/v1"
)

func main() {
	
	// Configuring api client
	// Api client should be singleton, can reuse the apiClient for multiple requests in various operations
	configuration := config.NewConfiguration()
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := widget.ApplyOTTRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}
	_, r, err := apiClient.WidgetAPI.ApplyOTT(context.Background()).ApplyOTTRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetAPI.ApplyOTT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyOTT`: ApplyOTTResponse
	fmt.Fprintf(os.Stdout, "Response from `WidgetAPI.ApplyOTT`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to ApplyOTTRequest (struct with type ApiApplyOTTRequest)

[**ApplyOTTRequest**](Widget/ApplyOTTRequest.md)

### Return type

[**ApplyOTTResponse**](Widget/ApplyOTTResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## ApplyToken

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	widget "github.com/dana-id/dana-go/widget/v1"
)

func main() {
	
	// Configuring api client
	// Api client should be singleton, can reuse the apiClient for multiple requests in various operations
	configuration := config.NewConfiguration()
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := widget.ApplyTokenRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}
	_, r, err := apiClient.WidgetAPI.ApplyToken(context.Background()).ApplyTokenRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetAPI.ApplyToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyToken`: ApplyTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `WidgetAPI.ApplyToken`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to ApplyTokenRequest (struct with type ApiApplyTokenRequest)

[**ApplyTokenRequest**](Widget/ApplyTokenRequest.md)

### Return type

[**ApplyTokenResponse**](Widget/ApplyTokenResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## CancelOrder

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	widget "github.com/dana-id/dana-go/widget/v1"
)

func main() {
	
	// Configuring api client
	// Api client should be singleton, can reuse the apiClient for multiple requests in various operations
	configuration := config.NewConfiguration()
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := widget.CancelOrderRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}
	_, r, err := apiClient.WidgetAPI.CancelOrder(context.Background()).CancelOrderRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetAPI.CancelOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelOrder`: CancelOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `WidgetAPI.CancelOrder`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to CancelOrderRequest (struct with type ApiCancelOrderRequest)

[**CancelOrderRequest**](Widget/CancelOrderRequest.md)

### Return type

[**CancelOrderResponse**](Widget/CancelOrderResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## QueryPayment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	widget "github.com/dana-id/dana-go/widget/v1"
)

func main() {
	
	// Configuring api client
	// Api client should be singleton, can reuse the apiClient for multiple requests in various operations
	configuration := config.NewConfiguration()
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := widget.QueryPaymentRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}
	_, r, err := apiClient.WidgetAPI.QueryPayment(context.Background()).QueryPaymentRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetAPI.QueryPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryPayment`: QueryPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `WidgetAPI.QueryPayment`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to QueryPaymentRequest (struct with type ApiQueryPaymentRequest)

[**QueryPaymentRequest**](Widget/QueryPaymentRequest.md)

### Return type

[**QueryPaymentResponse**](Widget/QueryPaymentResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## RefundOrder

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	widget "github.com/dana-id/dana-go/widget/v1"
)

func main() {
	
	// Configuring api client
	// Api client should be singleton, can reuse the apiClient for multiple requests in various operations
	configuration := config.NewConfiguration()
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := widget.RefundOrderRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}
	_, r, err := apiClient.WidgetAPI.RefundOrder(context.Background()).RefundOrderRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetAPI.RefundOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundOrder`: RefundOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `WidgetAPI.RefundOrder`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to RefundOrderRequest (struct with type ApiRefundOrderRequest)

[**RefundOrderRequest**](Widget/RefundOrderRequest.md)

### Return type

[**RefundOrderResponse**](Widget/RefundOrderResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## WidgetPayment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	widget "github.com/dana-id/dana-go/widget/v1"
)

func main() {
	
	// Configuring api client
	// Api client should be singleton, can reuse the apiClient for multiple requests in various operations
	configuration := config.NewConfiguration()
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := widget.WidgetPaymentRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}
	_, r, err := apiClient.WidgetAPI.WidgetPayment(context.Background()).WidgetPaymentRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetAPI.WidgetPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WidgetPayment`: WidgetPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `WidgetAPI.WidgetPayment`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to WidgetPaymentRequest (struct with type ApiWidgetPaymentRequest)

[**WidgetPaymentRequest**](Widget/WidgetPaymentRequest.md)

### Return type

[**WidgetPaymentResponse**](Widget/WidgetPaymentResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


# Enum Types

```go
import widget "github.com/dana-id/dana-go/widget/v1"

ipg := string(widget.SOURCEPLATFORM_IPG_)
```

## acquirementStatus
| Value | Description |
|-------|-------------|
| `INIT` |  |
| `SUCCESS` |  |
| `CLOSED` |  |
| `PAYING` |  |
| `MERCHANT_ACCEPT` |  |
| `CANCELLED` |  |


## actorType
| Value | Description |
|-------|-------------|
| `USER` | User |
| `MERCHANT` | Merchant<br |
| `MERCHANT_OPERATOR` | Merchant operator |
| `BACK_OFFICE` | Back office |
| `SYSTEM` | System |


## orderTerminalType
| Value | Description |
|-------|-------------|
| `APP` | Mobile Application |
| `WEB` | Browser Web |
| `WAP` | Mobile Wap |
| `SYSTEM` | System Call |


## payMethod
| Value | Description |
|-------|-------------|
| `BALANCE` | Payment method with balance |
| `COUPON` | Payment method with coupon |
| `NET_BANKING` | Payment method with internet banking |
| `CREDIT_CARD` | Payment method with credit card |
| `DEBIT_CARD` | Payment method with debit card |
| `VIRTUAL_ACCOUNT` | Payment method with virtual account |
| `OTC` | Payment method with OTC |
| `DIRECT_DEBIT_CREDIT_CARD` | Payment method with direct debit of credit card |
| `DIRECT_DEBIT_DEBIT_CARD` | Payment method with direct debit of debit card |
| `ONLINE_CREDIT` | Payment method with online Credit |
| `LOAN_CREDIT` | Payment method with DANA Cicil |
| `NETWORK_PAY` | Payment method with e-wallet |


## payOption
| Value | Description |
|-------|-------------|
| `NETWORK_PAY_PG_SPAY` | Payment method with ShopeePay e-wallet |
| `NETWORK_PAY_PG_OVO` | Payment method with OVO e-wallet |
| `NETWORK_PAY_PG_GOPAY` | Payment method with GoPay e-wallet |
| `NETWORK_PAY_PG_LINKAJA` | Payment method with LinkAja e-wallet |
| `NETWORK_PAY_PG_CARD` | Payment method with Card |
| `VIRTUAL_ACCOUNT_BCA` | Payment method with BCA virtual account |
| `VIRTUAL_ACCOUNT_BNI` | Payment method with BNI virtual account |
| `VIRTUAL_ACCOUNT_MANDIRI` | Payment method with Mandiri virtual account |
| `VIRTUAL_ACCOUNT_BRI` | Payment method with BRI virtual account |
| `VIRTUAL_ACCOUNT_BTPN` | Payment method with BTPN virtual account |
| `VIRTUAL_ACCOUNT_CIMB` | Payment method with CIMB virtual account |
| `VIRTUAL_ACCOUNT_PERMATA` | Payment method with Permata virtual account |


## promoType
| Value | Description |
|-------|-------------|
| `CASH_BACK` |  |
| `DISCOUNT` |  |
| `VOUCHER` |  |
| `POINT` |  |


## serviceScenario
| Value | Description |
|-------|-------------|
| `SCAN_AND_PAY` |  |
| `EXIT_AND_PAY` |  |
| `EMAS_PURCHASE` |  |


## serviceType
| Value | Description |
|-------|-------------|
| `PARKING` |  |
| `INVESTMENT` |  |


## sourcePlatform
| Value | Description |
|-------|-------------|
| `IPG` |  |


## terminalType
| Value | Description |
|-------|-------------|
| `APP` | Mobile Application |
| `WEB` | Browser Web |
| `WAP` | Mobile Wap |
| `SYSTEM` | System Call |


## type
| Value | Description |
|-------|-------------|
| `PAY_RETURN` | When finish payment, DANA will notify to the URL that has been defined by |
| `NOTIFICATION` | After the payment, the user will be redirected to merchant page, this is mandatory |




# WebhookParser

This section demonstrates how to securely verify and parse DANA webhook notifications using the `WebhookParser` utility from the Go SDK.

## Example

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	webhook "github.com/dana-id/dana-go/webhook"
	widget "github.com/dana-id/dana-go/widget/v1"
)

// This function would be your actual webhook handler in a real application.
func webhookNotificationHandler(req *http.Request) {
	// 1. Initialize the WebhookParser
	// You can provide the public key directly as a string or via a file path.
	// The parser will prioritize publicKeyPath if both are provided.

	// Option 1: Provide public key as a string
	// danaPublicKeyPEM := os.Getenv("DANA_PUBLIC_KEY")
	// parser, err := webhook.NewWebhookParser(&danaPublicKeyPEM, nil)

	// Option 2: Provide path to public key file 
	danaPublicKeyPath := os.Getenv("DANA_PUBLIC_KEY_PATH") // e.g., "/path/to/your/dana_public_key.pem"
	parser, err := webhook.NewWebhookParser(nil, &danaPublicKeyPath)
	if err != nil {
		fmt.Printf("Error creating WebhookParser: %v\n", err)
		return
	}

	// 2. Extract data from the incoming HTTP Request
	httpMethod := req.Method
	relativePathUrl := "/v1.0/debit/notify"
	// relativePathUrl := req.URL.Path // This should match the path DANA sends the webhook to for example: /v1.0/debit/notify

	// Read the request body
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Error reading request body: %v\n", err)
		return
	}
	defer req.Body.Close() // Important to close the body
	webhookBodyStr := string(bodyBytes)

	// Log received data for debugging (optional)
	fmt.Printf("Received webhook: Method=%s, Path=%s, Headers=%v, Body=%s\n",
		httpMethod, relativePathUrl, req.Header, webhookBodyStr)

	// 3. Parse and verify the webhook
	parsedData, err := parser.ParseWebhook(
		httpMethod,
		relativePathUrl,
		req.Header,
		webhookBodyStr,
	)

	if err != nil {
		fmt.Printf("Webhook parsing/verification failed: %v\n", err)
		// IMPORTANT: If verification fails, do not trust the payload.
		return
	}

	// 4. Use the parsed data
	fmt.Printf("Webhook parsed successfully!\n")
	fmt.Printf("Original Partner Reference No: %s\n", parsedData.OriginalPartnerReferenceNo)
	fmt.Printf("Amount: %s %s\n", parsedData.Amount.Value, parsedData.Amount.Currency)
	fmt.Printf("Status: %s\n", parsedData.LatestTransactionStatus)
	// Access other fields from parsedData as needed
}
```

## API Reference

### `WebhookParser`

The `WebhookParser` is part of the `webhook` package.

**Constructor:**

```go
func NewWebhookParser(publicKey *string, publicKeyPath *string) (*WebhookParser, error)
```
- `publicKey` (*string): Optional. The DANA gateway's public key as a PEM formatted string. This is used if `publicKeyPath` is not provided or is empty.
- `publicKeyPath` (*string): Optional. The file path to the DANA gateway's public key PEM file. If provided, this will be prioritized over the `publicKey` string.
- **Returns:** A pointer to a `WebhookParser` instance and an error if the public key cannot be loaded or is invalid.

**Method:**

```go
func (p *WebhookParser) ParseWebhook(httpMethod string, relativePathURL string, headers map[string]string, body string) (*webhook.FinishNotify, error)
```
- `httpMethod` (string): The HTTP method of the incoming webhook request (e.g., `http.MethodPost`).
- `relativePathURL` (string): The relative URL path of the webhook endpoint that received the notification (e.g., `/v1.0/debit/notify`).
- `headers` (map[string]string): A map containing the HTTP request headers. This map must include `X-SIGNATURE` and `X-TIMESTAMP` headers provided by DANA for signature verification.
- `body` (string): The raw JSON string payload from the webhook request body.
- **Returns:** A pointer to a `model.FinishNotifyRequest` struct containing the parsed and verified webhook data, or an error if parsing or signature verification fails.


## Security Notes
- Always use the official public key provided by DANA for webhook verification. Store and load it securely.
- The `ParseWebhook` method handles both JSON parsing and cryptographic signature verification. If it returns an error, the payload should not be trusted.

## Webhook Notification Models

The following webhook notification models may be received:

Model | Description
------------- | -------------
[**FinishNotifyRequest**](FinishNotify/FinishNotifyRequest.md) | Represents the standard notification payload for webhook payment events.

