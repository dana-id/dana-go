# PaymentGatewayAPI

All URIs are relative to http://api.sandbox.dana.id for sandbox environment and https://api.saas.dana.id for production environment

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](PaymentGatewayAPI.md#CancelOrder) | **Post** /payment-gateway/v1.0/debit/cancel.htm | Cancel Order - Payment Gateway
[**ConsultPay**](PaymentGatewayAPI.md#ConsultPay) | **Post** /v1.0/payment-gateway/consult-pay.htm | Consult Pay - Payment Gateway
[**CreateOrder**](PaymentGatewayAPI.md#CreateOrder) | **Post** /payment-gateway/v1.0/debit/payment-host-to-host.htm | Create Order - Payment Gateway
[**QueryPayment**](PaymentGatewayAPI.md#QueryPayment) | **Post** /payment-gateway/v1.0/debit/status.htm | Query Payment - Payment Gateway
[**RefundOrder**](PaymentGatewayAPI.md#RefundOrder) | **Post** /payment-gateway/v1.0/debit/refund.htm | Refund Order - Payment Gateway



## CancelOrder

> CancelOrderResponse CancelOrder(ctx).CancelOrderRequest(cancelOrderRequest).Execute()

Cancel Order - Payment Gateway



### Example

```go
package main

import (
	ipg "github.com/dana-id/go_client/ipg/v1"
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


## ConsultPay

> ConsultPayResponse ConsultPay(ctx).ConsultPayRequest(consultPayRequest).Execute()

Consult Pay - Payment Gateway



### Example

```go
package main

import (
	ipg "github.com/dana-id/go_client/ipg/v1"
	dana "github.com/dana-id/go_client"
	"context"
	"fmt"
	"os"
	ipg "github.com/dana-id/go_client/ipg/v1"
	"github.com/dana-id/go_client/config"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.ConsultPayRequest{
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
	_, r, err := apiClient.PaymentGatewayAPI.ConsultPay(context.Background()).ConsultPayRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.ConsultPay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsultPay`: ConsultPayResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.ConsultPay`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsultPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consultPayRequest** | [**ConsultPayRequest**](Ipg/ConsultPayRequest.md) |  | 

### Return type

[**ConsultPayResponse**](Ipg/ConsultPayResponse.md)

### Authorization

[ORIGIN](../README.md#ORIGIN), [X_PARTNER_ID](../README.md#X_PARTNER_ID), [CHANNEL_ID](../README.md#CHANNEL_ID), [PRIVATE_KEY](../README.md#PRIVATE_KEY), [PRIVATE_KEY_PATH](../README.md#PRIVATE_KEY_PATH), [ENV](../README.md#ENV)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrder

> CreateOrderResponse CreateOrder(ctx).CreateOrderRequest(createOrderRequest).Execute()

Create Order - Payment Gateway



### Example

```go
package main

import (
	ipg "github.com/dana-id/go_client/ipg/v1"
	dana "github.com/dana-id/go_client"
	"context"
	"fmt"
	"os"
	ipg "github.com/dana-id/go_client/ipg/v1"
	"github.com/dana-id/go_client/config"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.CreateOrderRequest{
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
	_, r, err := apiClient.PaymentGatewayAPI.CreateOrder(context.Background()).CreateOrderRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.CreateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrder`: CreateOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.CreateOrder`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrderRequest** | [**CreateOrderRequest**](Ipg/CreateOrderRequest.md) |  | 

### Return type

[**CreateOrderResponse**](Ipg/CreateOrderResponse.md)

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

Query Payment - Payment Gateway



### Example

```go
package main

import (
	ipg "github.com/dana-id/go_client/ipg/v1"
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

Refund Order - Payment Gateway



### Example

```go
package main

import (
	ipg "github.com/dana-id/go_client/ipg/v1"
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

