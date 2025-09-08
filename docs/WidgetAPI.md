# WidgetAPI

All URIs are relative to http://api.sandbox.dana.id for sandbox environment and https://api.saas.dana.id for production environment

List of available operations:

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountUnbinding**](WidgetAPI.md#AccountUnbinding) | **Post** /v1.0/registration-account-unbinding.htm | Account unbinding - Binding
[**ApplyOTT**](WidgetAPI.md#ApplyOTT) | **Post** /rest/v1.1/qr/apply-ott | Apply OTT - Widget
[**ApplyToken**](WidgetAPI.md#ApplyToken) | **Post** /v1.0/access-token/b2b2c.htm | Apply Token, required by Apply OTT - Binding
[**BalanceInquiry**](WidgetAPI.md#BalanceInquiry) | **Post** /v1.0/balance-inquiry.htm | Balance Inquiry
[**CancelOrder**](WidgetAPI.md#CancelOrder) | **Post** /v1.0/debit/cancel.htm | Cancel Order - Widget
[**QueryPayment**](WidgetAPI.md#QueryPayment) | **Post** /rest/v1.1/debit/status | Query Payment - Widget
[**QueryUserProfile**](WidgetAPI.md#QueryUserProfile) | **Post** /dana/member/query/queryUserProfile.htm | Query User Profile
[**RefundOrder**](WidgetAPI.md#RefundOrder) | **Post** /v1.0/debit/refund.htm | Refund Order - Widget
[**WidgetPayment**](WidgetAPI.md#WidgetPayment) | **Post** /rest/redirection/v1.0/debit/payment-host-to-host | Widget Payment - Widget


## Additional Documentation

* [Enum Types](#enum-types) - List of available enum constants
* [Webhook Parser](#webhookparser) - Webhook handling documentation
* [OAuth URL Generation](#oauth-url-generation) - Generate OAuth URLs for authorization
* [Complete Payment URL Generation](#complete-payment-url-generation) - Generate URL to complete the payment by combining webRedirectUrl with OTT token


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
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
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
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
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
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
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


## BalanceInquiry

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
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := widget.BalanceInquiryRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}
	_, r, err := apiClient.WidgetAPI.BalanceInquiry(context.Background()).BalanceInquiryRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetAPI.BalanceInquiry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BalanceInquiry`: BalanceInquiryResponse
	fmt.Fprintf(os.Stdout, "Response from `WidgetAPI.BalanceInquiry`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to BalanceInquiryRequest (struct with type ApiBalanceInquiryRequest)

[**BalanceInquiryRequest**](Widget/BalanceInquiryRequest.md)

### Return type

[**BalanceInquiryResponse**](Widget/BalanceInquiryResponse.md)

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
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
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
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
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


## QueryUserProfile

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
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	
	// Define request struct directly (example)
	request := widget.QueryUserProfileRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}
	_, r, err := apiClient.WidgetAPI.QueryUserProfile(context.Background()).QueryUserProfileRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetAPI.QueryUserProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryUserProfile`: QueryUserProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `WidgetAPI.QueryUserProfile`: %v\n", r.Body)
}
```

### Payload

Payload is passed through a pointer to QueryUserProfileRequest (struct with type ApiQueryUserProfileRequest)

[**QueryUserProfileRequest**](Widget/QueryUserProfileRequest.md)

### Return type

[**QueryUserProfileResponse**](Widget/QueryUserProfileResponse.md)

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
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
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
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
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

value := string(widget.SERVICETYPE_PARKING_)
```

## AcquirementStatus
| Value | Description |
|-------|-------------|
| `ACQUIREMENTSTATUS_INIT_` |  |
| `ACQUIREMENTSTATUS_SUCCESS_` |  |
| `ACQUIREMENTSTATUS_CLOSED_` |  |
| `ACQUIREMENTSTATUS_PAYING_` |  |
| `ACQUIREMENTSTATUS_MERCHANT_ACCEPT_` |  |
| `ACQUIREMENTSTATUS_CANCELLED_` |  |


## Mode
| Value | Description |
|-------|-------------|
| `MODE_API_` |  |
| `MODE_DEEPLINK_` |  |


## OrderTerminalType
| Value | Description |
|-------|-------------|
| `ORDERTERMINALTYPE_APP_` | Mobile Application |
| `ORDERTERMINALTYPE_WEB_` | Browser Web |
| `ORDERTERMINALTYPE_WAP_` | Mobile Wap |
| `ORDERTERMINALTYPE_SYSTEM_` | System Call |


## PayMethod
| Value | Description |
|-------|-------------|
| `PAYMETHOD_BALANCE_` | Payment method with balance |
| `PAYMETHOD_COUPON_` | Payment method with coupon |
| `PAYMETHOD_NET_BANKING_` | Payment method with internet banking |
| `PAYMETHOD_CREDIT_CARD_` | Payment method with credit card |
| `PAYMETHOD_DEBIT_CARD_` | Payment method with debit card |
| `PAYMETHOD_VIRTUAL_ACCOUNT_` | Payment method with virtual account |
| `PAYMETHOD_OTC_` | Payment method with OTC |
| `PAYMETHOD_DIRECT_DEBIT_CREDIT_CARD_` | Payment method with direct debit of credit card |
| `PAYMETHOD_DIRECT_DEBIT_DEBIT_CARD_` | Payment method with direct debit of debit card |
| `PAYMETHOD_ONLINE_CREDIT_` | Payment method with online Credit |
| `PAYMETHOD_LOAN_CREDIT_` | Payment method with DANA Cicil |
| `PAYMETHOD_NETWORK_PAY_` | Payment method with e-wallet |
| `PAYMETHOD_CARD_` | Payment method with Card |


## PayOption
| Value | Description |
|-------|-------------|
| `PAYOPTION_NETWORK_PAY_PG_SPAY_` | Payment method with ShopeePay e-wallet |
| `PAYOPTION_NETWORK_PAY_PG_OVO_` | Payment method with OVO e-wallet |
| `PAYOPTION_NETWORK_PAY_PG_GOPAY_` | Payment method with GoPay e-wallet |
| `PAYOPTION_NETWORK_PAY_PG_LINKAJA_` | Payment method with LinkAja e-wallet |
| `PAYOPTION_NETWORK_PAY_PG_CARD_` | Payment method with Card |
| `PAYOPTION_VIRTUAL_ACCOUNT_BCA_` | Payment method with BCA virtual account |
| `PAYOPTION_VIRTUAL_ACCOUNT_BNI_` | Payment method with BNI virtual account |
| `PAYOPTION_VIRTUAL_ACCOUNT_MANDIRI_` | Payment method with Mandiri virtual account |
| `PAYOPTION_VIRTUAL_ACCOUNT_BRI_` | Payment method with BRI virtual account |
| `PAYOPTION_VIRTUAL_ACCOUNT_BTPN_` | Payment method with BTPN virtual account |
| `PAYOPTION_VIRTUAL_ACCOUNT_CIMB_` | Payment method with CIMB virtual account |
| `PAYOPTION_VIRTUAL_ACCOUNT_PERMATA_` | Payment method with Permata virtual account |


## PromoType
| Value | Description |
|-------|-------------|
| `PROMOTYPE_CASH_BACK_` |  |
| `PROMOTYPE_DISCOUNT_` |  |
| `PROMOTYPE_VOUCHER_` |  |
| `PROMOTYPE_POINT_` |  |


## ResourceType
| Value | Description |
|-------|-------------|
| `RESOURCETYPE_BALANCE_` |  |
| `RESOURCETYPE_TRANSACTION_URL_` |  |
| `RESOURCETYPE_MASK_DANA_ID_` |  |
| `RESOURCETYPE_TOPUP_URL_` |  |
| `RESOURCETYPE_OTT_` |  |
| `RESOURCETYPE_USER_KYC_` |  |


## ResultStatus
| Value | Description |
|-------|-------------|
| `RESULTSTATUS_S_` |  |
| `RESULTSTATUS_F_` |  |
| `RESULTSTATUS_U_` |  |


## ServiceScenario
| Value | Description |
|-------|-------------|
| `SERVICESCENARIO_SCAN_AND_PAY_` |  |
| `SERVICESCENARIO_EXIT_AND_PAY_` |  |
| `SERVICESCENARIO_EMAS_PURCHASE_` |  |


## ServiceType
| Value | Description |
|-------|-------------|
| `SERVICETYPE_PARKING_` |  |
| `SERVICETYPE_INVESTMENT_` |  |


## SourcePlatform
| Value | Description |
|-------|-------------|
| `SOURCEPLATFORM_IPG_` |  |


## TerminalType
| Value | Description |
|-------|-------------|
| `TERMINALTYPE_APP_` | Mobile Application |
| `TERMINALTYPE_WEB_` | Browser Web |
| `TERMINALTYPE_WAP_` | Mobile Wap |
| `TERMINALTYPE_SYSTEM_` | System Call |


## Type
| Value | Description |
|-------|-------------|
| `TYPE_PAY_RETURN_` | When finish payment, DANA will notify to the URL that has been defined by |
| `TYPE_NOTIFICATION_` | After the payment, the user will be redirected to merchant page, this is mandatory |




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


# OAuth URL Generation

This section demonstrates how to generate OAuth URLs for widget authorization using the Go SDK.

## Example

```go
package main

import (
	"fmt"

	"github.com/dana-team/dana-go-api-client/widget/v1"
	"github.com/dana-team/dana-go-api-client/widget/v1/model"
)

func main() {
	// Set up OAuth2 URL data
	oauth2UrlData := &model.Oauth2UrlData{
		RedirectUrl: "https://google.com",
		MerchantId:  merchantId,
		SeamlessData: map[string]interface{}{
			"mobileNumber": "087875849373",
		},
        Mode: "DEEPLINK", // or "API" (default value is "API")
	}

	// Generate the OAuth URL
	oauthUrl := widget.GenerateOauthUrl(oauth2UrlData, privateKey)
	fmt.Println("Generated OAuth URL:", oauthUrl)
}
```

The generated URL can be used to redirect users to DANA's authorization page.

# Complete Payment URL Generation

You can generate URL to complete the payment by combining the webRedirectUrl from a WidgetPaymentResponse with an OTT token from an ApplyOTTResponse.

## Example

```go
package main

import (
	"fmt"

	"github.com/dana-team/dana-go-api-client/widget/v1"
	"github.com/dana-team/dana-go-api-client/widget/v1/model"
)

func main() {
	// Example response from createWidgetPayment
	widgetPaymentResponse := &model.WidgetPaymentResponse{
		WebRedirectUrl: "https://example.com/payment?token=abc123",
	}
	// This should be from createPayment Widget API

	// Example response from applyOTT
	applyOTTResponse := &model.ApplyOTTResponse{
		UserResources: []model.UserResource{
			{
				Value: "ott_token_value",
			},
		},
	}
	// This should be from applyOTT Widget API

	// Generate the payment URL
	paymentUrl := widget.GenerateCompletePaymentUrl(widgetPaymentResponse, applyOTTResponse)
	fmt.Println("Generated Complete Payment URL:", paymentUrl)
}
```

The generated URL can be used to redirect users to DANA's payment page.

