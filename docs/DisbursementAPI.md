# DisbursementAPI

All URIs are relative to http://api.sandbox.dana.id for sandbox environment and https://api.saas.dana.id for production environment

Method | HTTP request | Description
------------- | ------------- | -------------
[**BankAccountInquiry**](DisbursementAPI.md#BankAccountInquiry) | **Post** /v1.0/emoney/bank-account-inquiry.htm | Transfer to Bank Account Inquiry
[**DanaAccountInquiry**](DisbursementAPI.md#DanaAccountInquiry) | **Post** /v1.0/emoney/account-inquiry.htm | DANA Account Inquiry
[**QueryMerchantResource**](DisbursementAPI.md#QueryMerchantResource) | **Post** /merchant/queryMerchantResource.htm | Member – Merchant Open API Check Disbursement Account
[**TransferToBank**](DisbursementAPI.md#TransferToBank) | **Post** /v1.0/emoney/transfer-bank.htm | Transfer to Bank
[**TransferToBankInquiryStatus**](DisbursementAPI.md#TransferToBankInquiryStatus) | **Post** /v1.0/emoney/transfer-bank-status.htm | Transfer to Bank Inquiry Status
[**TransferToDana**](DisbursementAPI.md#TransferToDana) | **Post** /v1.0/emoney/topup.htm | Transfer to DANA
[**TransferToDanaInquiryStatus**](DisbursementAPI.md#TransferToDanaInquiryStatus) | **Post** /v1.0/emoney/topup-status.htm | Transfer to DANA Inquiry Status


## Additional Documentation

* [Enum Types](#enum-types) - List of available enum constants
* [Webhook Parser](#webhookparser) - Webhook handling


## BankAccountInquiry

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	disbursement "github.com/dana-id/dana-go/disbursement/v1"
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
	request := disbursement.BankAccountInquiryRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	_, r, err := apiClient.DisbursementAPI.BankAccountInquiry(context.Background()).BankAccountInquiryRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisbursementAPI.BankAccountInquiry`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BankAccountInquiry`: BankAccountInquiryResponse
	fmt.Fprintf(os.Stdout, "Response from `DisbursementAPI.BankAccountInquiry`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to BankAccountInquiryRequest (struct with type ApiBankAccountInquiryRequest)

[**BankAccountInquiryRequest**](Disbursement/BankAccountInquiryRequest.md)

### Return type

[**BankAccountInquiryResponse**](Disbursement/BankAccountInquiryResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## DanaAccountInquiry

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	disbursement "github.com/dana-id/dana-go/disbursement/v1"
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
	request := disbursement.DanaAccountInquiryRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	_, r, err := apiClient.DisbursementAPI.DanaAccountInquiry(context.Background()).DanaAccountInquiryRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisbursementAPI.DanaAccountInquiry`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DanaAccountInquiry`: DanaAccountInquiryResponse
	fmt.Fprintf(os.Stdout, "Response from `DisbursementAPI.DanaAccountInquiry`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to DanaAccountInquiryRequest (struct with type ApiDanaAccountInquiryRequest)

[**DanaAccountInquiryRequest**](Disbursement/DanaAccountInquiryRequest.md)

### Return type

[**DanaAccountInquiryResponse**](Disbursement/DanaAccountInquiryResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## QueryMerchantResource

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	disbursement "github.com/dana-id/dana-go/disbursement/v1"
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
	request := disbursement.QueryMerchantResourceRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	_, r, err := apiClient.DisbursementAPI.QueryMerchantResource(context.Background()).QueryMerchantResourceRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisbursementAPI.QueryMerchantResource`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryMerchantResource`: QueryMerchantResourceResponse
	fmt.Fprintf(os.Stdout, "Response from `DisbursementAPI.QueryMerchantResource`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to QueryMerchantResourceRequest (struct with type ApiQueryMerchantResourceRequest)

[**QueryMerchantResourceRequest**](Disbursement/QueryMerchantResourceRequest.md)

### Return type

[**QueryMerchantResourceResponse**](Disbursement/QueryMerchantResourceResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## TransferToBank

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	disbursement "github.com/dana-id/dana-go/disbursement/v1"
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
	request := disbursement.TransferToBankRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	_, r, err := apiClient.DisbursementAPI.TransferToBank(context.Background()).TransferToBankRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisbursementAPI.TransferToBank`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferToBank`: TransferToBankResponse
	fmt.Fprintf(os.Stdout, "Response from `DisbursementAPI.TransferToBank`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to TransferToBankRequest (struct with type ApiTransferToBankRequest)

[**TransferToBankRequest**](Disbursement/TransferToBankRequest.md)

### Return type

[**TransferToBankResponse**](Disbursement/TransferToBankResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## TransferToBankInquiryStatus

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	disbursement "github.com/dana-id/dana-go/disbursement/v1"
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
	request := disbursement.TransferToBankInquiryStatusRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	_, r, err := apiClient.DisbursementAPI.TransferToBankInquiryStatus(context.Background()).TransferToBankInquiryStatusRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisbursementAPI.TransferToBankInquiryStatus`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferToBankInquiryStatus`: TransferToBankInquiryStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DisbursementAPI.TransferToBankInquiryStatus`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to TransferToBankInquiryStatusRequest (struct with type ApiTransferToBankInquiryStatusRequest)

[**TransferToBankInquiryStatusRequest**](Disbursement/TransferToBankInquiryStatusRequest.md)

### Return type

[**TransferToBankInquiryStatusResponse**](Disbursement/TransferToBankInquiryStatusResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## TransferToDana

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	disbursement "github.com/dana-id/dana-go/disbursement/v1"
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
	request := disbursement.TransferToDanaRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	_, r, err := apiClient.DisbursementAPI.TransferToDana(context.Background()).TransferToDanaRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisbursementAPI.TransferToDana`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferToDana`: TransferToDanaResponse
	fmt.Fprintf(os.Stdout, "Response from `DisbursementAPI.TransferToDana`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to TransferToDanaRequest (struct with type ApiTransferToDanaRequest)

[**TransferToDanaRequest**](Disbursement/TransferToDanaRequest.md)

### Return type

[**TransferToDanaResponse**](Disbursement/TransferToDanaResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


## TransferToDanaInquiryStatus

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	disbursement "github.com/dana-id/dana-go/disbursement/v1"
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
	request := disbursement.TransferToDanaInquiryStatusRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	_, r, err := apiClient.DisbursementAPI.TransferToDanaInquiryStatus(context.Background()).TransferToDanaInquiryStatusRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisbursementAPI.TransferToDanaInquiryStatus`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferToDanaInquiryStatus`: TransferToDanaInquiryStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DisbursementAPI.TransferToDanaInquiryStatus`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to TransferToDanaInquiryStatusRequest (struct with type ApiTransferToDanaInquiryStatusRequest)

[**TransferToDanaInquiryStatusRequest**](Disbursement/TransferToDanaInquiryStatusRequest.md)

### Return type

[**TransferToDanaInquiryStatusResponse**](Disbursement/TransferToDanaInquiryStatusResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)


# Enum Types

```go
import disbursement "github.com/dana-id/dana-go/disbursement/v1"

ipg := string(disbursement.SOURCEPLATFORM_IPG_)
```

## actorType
| Value | Description |
|-------|-------------|
| `USER` | User |
| `MERCHANT` | Merchant<br |
| `MERCHANT_OPERATOR` | Merchant operator |
| `BACK_OFFICE` | Back office |
| `SYSTEM` | System |


## latestTransactionStatus
| Value | Description |
|-------|-------------|
| `00` |  |
| `01` |  |
| `02` |  |
| `03` |  |
| `04` |  |
| `05` |  |
| `06` |  |
| `07` |  |


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


## resourceType
| Value | Description |
|-------|-------------|
| `MERCHANT_DEPOSIT_BALANCE` |  |
| `MERCHANT_AVAILABLE_BALANCE` |  |
| `MERCHANT_TOTAL_BALANCE` |  |


## resultStatus
| Value | Description |
|-------|-------------|
| `S` |  |
| `F` |  |
| `U` |  |


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
	disbursement "github.com/dana-id/dana-go/disbursement/v1"
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

