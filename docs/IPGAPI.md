# IPGAPI

All URIs are relative to http://api.sandbox.dana.id for sandbox environment and https://api.saas.dana.id for production environment

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountUnbinding**](IPGAPI.md#AccountUnbinding) | **Post** /v1.0/registration-account-unbinding.htm | Account unbinding - Binding
[**ApplyOTT**](IPGAPI.md#ApplyOTT) | **Post** /rest/v1.1/qr/apply-ott | Apply OTT - IPG
[**ApplyToken**](IPGAPI.md#ApplyToken) | **Post** /v1.0/access-token/b2b2c.htm | Apply Token, required by Apply OTT - Binding
[**CancelOrder**](IPGAPI.md#CancelOrder) | **Post** /v1.0/debit/cancel.htm | Cancel Order - IPG
[**IpgPayment**](IPGAPI.md#IpgPayment) | **Post** /rest/redirection/v1.0/debit/payment-host-to-host | IPG payment - IPG
[**QueryPayment**](IPGAPI.md#QueryPayment) | **Post** /rest/v1.1/debit/status | Query Payment - IPG
[**RefundOrder**](IPGAPI.md#RefundOrder) | **Post** /v1.0/debit/refund.htm | Refund Order - IPG



## AccountUnbinding

> AccountUnbindingResponse AccountUnbinding(ctx).AccountUnbindingRequest(accountUnbindingRequest).Execute()

Account unbinding - Binding



### Example

```go
package main

import (
	dana "github.com/dana-id/go_client"
	"context"
	"fmt"
	"os"
	ipg "github.com/dana-id/go_client/ipg/v1"
	"github.com/dana-id/go_client/config"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.AccountUnbindingRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX,
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.AccountUnbinding(context.Background()).AccountUnbindingRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.AccountUnbinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountUnbinding`: AccountUnbindingResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.AccountUnbinding`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountUnbindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountUnbindingRequest** | [**AccountUnbindingRequest**](Ipg/AccountUnbindingRequest.md) |  | 

### Return type

[**AccountUnbindingResponse**](Ipg/AccountUnbindingResponse.md)

### Authorization

[X_PARTNER_ID](../README.md#X_PARTNER_ID), [PRIVATE_KEY](../README.md#PRIVATE_KEY), [PRIVATE_KEY_PATH](../README.md#PRIVATE_KEY_PATH), [ENV](../README.md#ENV)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyOTT

> ApplyOTTResponse ApplyOTT(ctx).ApplyOTTRequest(applyOTTRequest).Execute()

Apply OTT - IPG



### Example

```go
package main

import (
	dana "github.com/dana-id/go_client"
	"context"
	"fmt"
	"os"
	ipg "github.com/dana-id/go_client/ipg/v1"
	"github.com/dana-id/go_client/config"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.ApplyOTTRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX,
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.ApplyOTT(context.Background()).ApplyOTTRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.ApplyOTT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyOTT`: ApplyOTTResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.ApplyOTT`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyOTTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applyOTTRequest** | [**ApplyOTTRequest**](Ipg/ApplyOTTRequest.md) |  | 

### Return type

[**ApplyOTTResponse**](Ipg/ApplyOTTResponse.md)

### Authorization

[ORIGIN](../README.md#ORIGIN), [X_PARTNER_ID](../README.md#X_PARTNER_ID), [CHANNEL_ID](../README.md#CHANNEL_ID), [PRIVATE_KEY](../README.md#PRIVATE_KEY), [PRIVATE_KEY_PATH](../README.md#PRIVATE_KEY_PATH), [ENV](../README.md#ENV)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyToken

> ApplyTokenResponse ApplyToken(ctx).ApplyTokenRequest(applyTokenRequest).Execute()

Apply Token, required by Apply OTT - Binding



### Example

```go
package main

import (
	dana "github.com/dana-id/go_client"
	"context"
	"fmt"
	"os"
	ipg "github.com/dana-id/go_client/ipg/v1"
	"github.com/dana-id/go_client/config"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.ApplyTokenRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX,
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.ApplyToken(context.Background()).ApplyTokenRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.ApplyToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyToken`: ApplyTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.ApplyToken`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applyTokenRequest** | [**ApplyTokenRequest**](Ipg/ApplyTokenRequest.md) |  | 

### Return type

[**ApplyTokenResponse**](Ipg/ApplyTokenResponse.md)

### Authorization

[X_PARTNER_ID](../README.md#X_PARTNER_ID), [PRIVATE_KEY](../README.md#PRIVATE_KEY), [PRIVATE_KEY_PATH](../README.md#PRIVATE_KEY_PATH), [ENV](../README.md#ENV)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelOrder

> CancelOrderResponse CancelOrder(ctx).CancelOrderRequest(cancelOrderRequest).Execute()

Cancel Order - IPG



### Example

```go
package main

import (
	dana "github.com/dana-id/go_client"
	"context"
	"fmt"
	"os"
	ipg "github.com/dana-id/go_client/ipg/v1"
	"github.com/dana-id/go_client/config"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.CancelOrderRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX,
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.CancelOrder(context.Background()).CancelOrderRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.CancelOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelOrder`: CancelOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.CancelOrder`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancelOrderRequest** | [**CancelOrderRequest**](Ipg/CancelOrderRequest.md) |  | 

### Return type

[**CancelOrderResponse**](Ipg/CancelOrderResponse.md)

### Authorization

[ORIGIN](../README.md#ORIGIN), [X_PARTNER_ID](../README.md#X_PARTNER_ID), [CHANNEL_ID](../README.md#CHANNEL_ID), [PRIVATE_KEY](../README.md#PRIVATE_KEY), [PRIVATE_KEY_PATH](../README.md#PRIVATE_KEY_PATH)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpgPayment

> IPGPaymentResponse IpgPayment(ctx).IPGPaymentRequest(iPGPaymentRequest).Execute()

IPG payment - IPG



### Example

```go
package main

import (
	dana "github.com/dana-id/go_client"
	"context"
	"fmt"
	"os"
	ipg "github.com/dana-id/go_client/ipg/v1"
	"github.com/dana-id/go_client/config"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.IpgPaymentRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX,
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.IpgPayment(context.Background()).IpgPaymentRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.IpgPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IpgPayment`: IPGPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.IpgPayment`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIpgPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iPGPaymentRequest** | [**IPGPaymentRequest**](Ipg/IPGPaymentRequest.md) |  | 

### Return type

[**IPGPaymentResponse**](Ipg/IPGPaymentResponse.md)

### Authorization

[ORIGIN](../README.md#ORIGIN), [X_PARTNER_ID](../README.md#X_PARTNER_ID), [CHANNEL_ID](../README.md#CHANNEL_ID), [PRIVATE_KEY](../README.md#PRIVATE_KEY), [PRIVATE_KEY_PATH](../README.md#PRIVATE_KEY_PATH), [ENV](../README.md#ENV)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryPayment

> QueryPaymentResponse QueryPayment(ctx).QueryPaymentRequest(queryPaymentRequest).Execute()

Query Payment - IPG



### Example

```go
package main

import (
	dana "github.com/dana-id/go_client"
	"context"
	"fmt"
	"os"
	ipg "github.com/dana-id/go_client/ipg/v1"
	"github.com/dana-id/go_client/config"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.QueryPaymentRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX,
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.QueryPayment(context.Background()).QueryPaymentRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.QueryPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryPayment`: QueryPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.QueryPayment`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryPaymentRequest** | [**QueryPaymentRequest**](Ipg/QueryPaymentRequest.md) |  | 

### Return type

[**QueryPaymentResponse**](Ipg/QueryPaymentResponse.md)

### Authorization

[ORIGIN](../README.md#ORIGIN), [X_PARTNER_ID](../README.md#X_PARTNER_ID), [CHANNEL_ID](../README.md#CHANNEL_ID), [PRIVATE_KEY](../README.md#PRIVATE_KEY), [PRIVATE_KEY_PATH](../README.md#PRIVATE_KEY_PATH)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundOrder

> RefundOrderResponse RefundOrder(ctx).RefundOrderRequest(refundOrderRequest).Execute()

Refund Order - IPG



### Example

```go
package main

import (
	dana "github.com/dana-id/go_client"
	"context"
	"fmt"
	"os"
	ipg "github.com/dana-id/go_client/ipg/v1"
	"github.com/dana-id/go_client/config"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.RefundOrderRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		ENV:          config.ENV_SANDBOX,
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.RefundOrder(context.Background()).RefundOrderRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.RefundOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundOrder`: RefundOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.RefundOrder`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundOrderRequest** | [**RefundOrderRequest**](Ipg/RefundOrderRequest.md) |  | 

### Return type

[**RefundOrderResponse**](Ipg/RefundOrderResponse.md)

### Authorization

[ORIGIN](../README.md#ORIGIN), [X_PARTNER_ID](../README.md#X_PARTNER_ID), [CHANNEL_ID](../README.md#CHANNEL_ID), [PRIVATE_KEY](../README.md#PRIVATE_KEY), [PRIVATE_KEY_PATH](../README.md#PRIVATE_KEY_PATH)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# Enum Types
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
	dana "github.com/dana-id/go_client"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	ipg "github.com/dana-id/go_client/ipg/v1"
	ipg "github.com/dana-id/go_client/ipg/v1"
)

// This function would be your actual webhook handler in a real application.
func webhookNotificationHandler(req *http.Request) {
	// 1. Initialize the WebhookParser
	danaPublicKeyPEM := os.Getenv("DANA_PUBLIC_KEY")

	parser, err := webhook.NewWebhookParser(danaPublicKeyPEM)
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
func NewWebhookParser(gatewayPublicKeyPEM string) (*WebhookParser, error)
```
- `gatewayPublicKeyPEM` (string): The DANA gateway's public key in PEM format. Used to verify the webhook signature.
- **Returns:** A pointer to a `WebhookParser` instance and an error if the public key is invalid.

**Method:**

```go
func (p *WebhookParser) ParseWebhook(httpMethod string, relativePathURL string, headers map[string]string, body string) (*model.FinishNotify, error)
```
- `httpMethod` (string): The HTTP method of the incoming webhook request (e.g., `http.MethodPost`).
- `relativePathURL` (string): The relative URL path of the webhook endpoint that received the notification (e.g., `/v1.0/debit/notify`).
- `headers` (map[string]string): A map containing the HTTP request headers. This map must include `X-SIGNATURE` and `X-TIMESTAMP` headers provided by DANA for signature verification.
- `body` (string): The raw JSON string payload from the webhook request body.
- **Returns:** A pointer to a `model.FinishNotify` struct containing the parsed and verified webhook data, or an error if parsing or signature verification fails.


## Security Notes
- Always use the official public key provided by DANA for webhook verification. Store and load it securely.
- The `ParseWebhook` method handles both JSON parsing and cryptographic signature verification. If it returns an error, the payload should not be trusted.

## Webhook Notification Models

The following webhook notification models may be received:

Model | Description
------------- | -------------
[**FinishNotify**](PaymentGateway/FinishNotify.md) | Represents the standard notification payload for payment events.

