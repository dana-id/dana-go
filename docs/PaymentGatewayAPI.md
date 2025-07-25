# PaymentGatewayAPI

All URIs are relative to http://api.sandbox.dana.id for sandbox environment and https://api.saas.dana.id for production environment

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](PaymentGatewayAPI.md#CancelOrder) | **Post** /payment-gateway/v1.0/debit/cancel.htm | Cancel Order - Payment Gateway
[**ConsultPay**](PaymentGatewayAPI.md#ConsultPay) | **Post** /v1.0/payment-gateway/consult-pay.htm | Consult Pay - Payment Gateway
[**CreateOrder**](PaymentGatewayAPI.md#CreateOrder) | **Post** /payment-gateway/v1.0/debit/payment-host-to-host.htm | Create Order - Payment Gateway
[**QueryPayment**](PaymentGatewayAPI.md#QueryPayment) | **Post** /payment-gateway/v1.0/debit/status.htm | Query Payment - Payment Gateway
[**RefundOrder**](PaymentGatewayAPI.md#RefundOrder) | **Post** /payment-gateway/v1.0/debit/refund.htm | Refund Order - Payment Gateway


## Additional Documentation

* [Enum Types](#enum-types) - List of available enum constants
* [Webhook Parser](#webhookparser) - Webhook handling


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
	payment_gateway "github.com/dana-id/dana-go/payment_gateway/v1"
)

func main() {
	
	// Configuring api client
	// Api client should be singleton, can reuse the apiClient for multiple requests in various operations
	configuration := config.NewConfiguration()
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"), // Can provide the private key directly as a string or via a file path (PRIVATE_KEY_PATH). If both added, we will prioritize the path
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := payment_gateway.CancelOrderRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	_, r, err := apiClient.PaymentGatewayAPI.CancelOrder(context.Background()).CancelOrderRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.CancelOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelOrder`: CancelOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.CancelOrder`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to CancelOrderRequest (struct with type ApiCancelOrderRequest)

[**CancelOrderRequest**](PaymentGateway/CancelOrderRequest.md)

### Return type

[**CancelOrderResponse**](PaymentGateway/CancelOrderResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## ConsultPay

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	payment_gateway "github.com/dana-id/dana-go/payment_gateway/v1"
)

func main() {
	
	// Configuring api client
	// Api client should be singleton, can reuse the apiClient for multiple requests in various operations
	configuration := config.NewConfiguration()
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"), // Can provide the private key directly as a string or via a file path (PRIVATE_KEY_PATH). If both added, we will prioritize the path
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := payment_gateway.ConsultPayRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	_, r, err := apiClient.PaymentGatewayAPI.ConsultPay(context.Background()).ConsultPayRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.ConsultPay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsultPay`: ConsultPayResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.ConsultPay`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to ConsultPayRequest (struct with type ApiConsultPayRequest)

[**ConsultPayRequest**](PaymentGateway/ConsultPayRequest.md)

### Return type

[**ConsultPayResponse**](PaymentGateway/ConsultPayResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## CreateOrder

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	payment_gateway "github.com/dana-id/dana-go/payment_gateway/v1"
)

func main() {
	
	// Configuring api client
	// Api client should be singleton, can reuse the apiClient for multiple requests in various operations
	configuration := config.NewConfiguration()
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"), // Can provide the private key directly as a string or via a file path (PRIVATE_KEY_PATH). If both added, we will prioritize the path
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := payment_gateway.CreateOrderRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	_, r, err := apiClient.PaymentGatewayAPI.CreateOrder(context.Background()).CreateOrderRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.CreateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrder`: CreateOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.CreateOrder`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to CreateOrderRequest (struct with type ApiCreateOrderRequest)

[**CreateOrderRequest**](PaymentGateway/CreateOrderRequest.md)

### Return type

[**CreateOrderResponse**](PaymentGateway/CreateOrderResponse.md)

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
	payment_gateway "github.com/dana-id/dana-go/payment_gateway/v1"
)

func main() {
	
	// Configuring api client
	// Api client should be singleton, can reuse the apiClient for multiple requests in various operations
	configuration := config.NewConfiguration()
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"), // Can provide the private key directly as a string or via a file path (PRIVATE_KEY_PATH). If both added, we will prioritize the path
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := payment_gateway.QueryPaymentRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	_, r, err := apiClient.PaymentGatewayAPI.QueryPayment(context.Background()).QueryPaymentRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.QueryPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryPayment`: QueryPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.QueryPayment`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to QueryPaymentRequest (struct with type ApiQueryPaymentRequest)

[**QueryPaymentRequest**](PaymentGateway/QueryPaymentRequest.md)

### Return type

[**QueryPaymentResponse**](PaymentGateway/QueryPaymentResponse.md)

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
	payment_gateway "github.com/dana-id/dana-go/payment_gateway/v1"
)

func main() {
	
	// Configuring api client
	// Api client should be singleton, can reuse the apiClient for multiple requests in various operations
	configuration := config.NewConfiguration()
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"), // Can provide the private key directly as a string or via a file path (PRIVATE_KEY_PATH). If both added, we will prioritize the path
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := payment_gateway.RefundOrderRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	_, r, err := apiClient.PaymentGatewayAPI.RefundOrder(context.Background()).RefundOrderRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.RefundOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundOrder`: RefundOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.RefundOrder`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to RefundOrderRequest (struct with type ApiRefundOrderRequest)

[**RefundOrderRequest**](PaymentGateway/RefundOrderRequest.md)

### Return type

[**RefundOrderResponse**](PaymentGateway/RefundOrderResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


# Enum Types

```go
import payment_gateway "github.com/dana-id/dana-go/payment_gateway/v1"

ipg := string(payment_gateway.SOURCEPLATFORM_IPG_)
```

## acquirementStatus
| Value | Description |
|-------|-------------|
| `INIT` | Order is created but not paid yet |
| `SUCCESS` | Order is succeeded |
| `CLOSED` | Order is closed |
| `PAYING` | Order is paid but not finish |
| `MERCHANT_ACCEPT` | Order is accepted by merchant after order is paid for PAY-CONFIRM |
| `CANCELLED` | Order is cancelled |


## actorType
| Value | Description |
|-------|-------------|
| `USER` | User |
| `MERCHANT` | Merchant |
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
[**FinishNotifyRequest**](FinishNotify/FinishNotifyRequest.md) | Represents the standard notification payload for payment events.

