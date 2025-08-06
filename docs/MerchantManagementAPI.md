# MerchantManagementAPI

All URIs are relative to http://api.sandbox.dana.id for sandbox environment and https://api.saas.dana.id for production environment

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDivision**](MerchantManagementAPI.md#CreateDivision) | **Post** /dana/merchant/division/createDivision.htm | Create Division
[**CreateShop**](MerchantManagementAPI.md#CreateShop) | **Post** /dana/merchant/shop/createShop.htm | Member – Create Shop
[**QueryDivision**](MerchantManagementAPI.md#QueryDivision) | **Post** /dana/merchant/division/queryDivision.htm | Query Division
[**QueryMerchantResource**](MerchantManagementAPI.md#QueryMerchantResource) | **Post** /dana/merchant/queryMerchantResource.htm | Member – Merchant Open API Check Disbursement Account
[**QueryShop**](MerchantManagementAPI.md#QueryShop) | **Post** /dana/merchant/shop/queryShop.htm | Member – Query Shop
[**UpdateDivision**](MerchantManagementAPI.md#UpdateDivision) | **Post** /dana/merchant/division/updateDivision.htm | Update Division
[**UpdateShop**](MerchantManagementAPI.md#UpdateShop) | **Post** /dana/merchant/shop/updateShop.htm | Update Shop



## CreateDivision

> CreateDivisionResponse CreateDivision(ctx).CreateDivisionRequest(createDivisionRequest).Execute()

Create Division



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	merchant_management "github.com/dana-id/dana-go/merchant_management/v1"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.CreateDivisionRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.CreateDivision(context.Background()).CreateDivisionRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.CreateDivision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDivision`: CreateDivisionResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.CreateDivision`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDivisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDivisionRequest** | [**CreateDivisionRequest**](MerchantManagement/CreateDivisionRequest.md) |  | 

### Return type

[**CreateDivisionResponse**](MerchantManagement/CreateDivisionResponse.md)

### Authorization

[CLIENT_SECRET](../README.md#CLIENT_SECRET)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShop

> CreateShopResponse CreateShop(ctx).CreateShopRequest(createShopRequest).Execute()

Member – Create Shop



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	merchant_management "github.com/dana-id/dana-go/merchant_management/v1"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.CreateShopRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.CreateShop(context.Background()).CreateShopRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.CreateShop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShop`: CreateShopResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.CreateShop`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createShopRequest** | [**CreateShopRequest**](MerchantManagement/CreateShopRequest.md) |  | 

### Return type

[**CreateShopResponse**](MerchantManagement/CreateShopResponse.md)

### Authorization

[CLIENT_SECRET](../README.md#CLIENT_SECRET)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryDivision

> QueryDivisionResponse QueryDivision(ctx).QueryDivisionRequest(queryDivisionRequest).Execute()

Query Division



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	merchant_management "github.com/dana-id/dana-go/merchant_management/v1"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.QueryDivisionRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.QueryDivision(context.Background()).QueryDivisionRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.QueryDivision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryDivision`: QueryDivisionResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.QueryDivision`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryDivisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryDivisionRequest** | [**QueryDivisionRequest**](MerchantManagement/QueryDivisionRequest.md) |  | 

### Return type

[**QueryDivisionResponse**](MerchantManagement/QueryDivisionResponse.md)

### Authorization

[CLIENT_SECRET](../README.md#CLIENT_SECRET)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryMerchantResource

> QueryMerchantResourceResponse QueryMerchantResource(ctx).QueryMerchantResourceRequest(queryMerchantResourceRequest).Execute()

Member – Merchant Open API Check Disbursement Account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	merchant_management "github.com/dana-id/dana-go/merchant_management/v1"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.QueryMerchantResourceRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.QueryMerchantResource(context.Background()).QueryMerchantResourceRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.QueryMerchantResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryMerchantResource`: QueryMerchantResourceResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.QueryMerchantResource`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryMerchantResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryMerchantResourceRequest** | [**QueryMerchantResourceRequest**](MerchantManagement/QueryMerchantResourceRequest.md) |  | 

### Return type

[**QueryMerchantResourceResponse**](MerchantManagement/QueryMerchantResourceResponse.md)

### Authorization

[CLIENT_SECRET](../README.md#CLIENT_SECRET)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryShop

> QueryShopResponse QueryShop(ctx).QueryShopRequest(queryShopRequest).Execute()

Member – Query Shop



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	merchant_management "github.com/dana-id/dana-go/merchant_management/v1"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.QueryShopRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.QueryShop(context.Background()).QueryShopRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.QueryShop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryShop`: QueryShopResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.QueryShop`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryShopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryShopRequest** | [**QueryShopRequest**](MerchantManagement/QueryShopRequest.md) |  | 

### Return type

[**QueryShopResponse**](MerchantManagement/QueryShopResponse.md)

### Authorization

[CLIENT_SECRET](../README.md#CLIENT_SECRET)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDivision

> UpdateDivisionResponse UpdateDivision(ctx).UpdateDivisionRequest(updateDivisionRequest).Execute()

Update Division



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	merchant_management "github.com/dana-id/dana-go/merchant_management/v1"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.UpdateDivisionRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.UpdateDivision(context.Background()).UpdateDivisionRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.UpdateDivision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDivision`: UpdateDivisionResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.UpdateDivision`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDivisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDivisionRequest** | [**UpdateDivisionRequest**](MerchantManagement/UpdateDivisionRequest.md) |  | 

### Return type

[**UpdateDivisionResponse**](MerchantManagement/UpdateDivisionResponse.md)

### Authorization

[CLIENT_SECRET](../README.md#CLIENT_SECRET)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShop

> UpdateShopResponse UpdateShop(ctx).UpdateShopRequest(updateShopRequest).Execute()

Update Shop



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	dana "github.com/dana-id/dana-go"
	"github.com/dana-id/dana-go/config"
	merchant_management "github.com/dana-id/dana-go/merchant_management/v1"
)

func main() {
	// Define request struct directly (example)
	request := payment_gateway.UpdateShopRequest{
		// Fill in required fields here, e.g.:
		// Field1: "value1",
		// Field2: "value2",
	}

	configuration := config.NewConfiguration()
	// Set API keys
	configuration.APIKey = &config.APIKey{
		// ENV:          config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production. Can use DANA_ENV instead
		DANA_ENV:     config.ENV_SANDBOX, // use config.ENV_PRODUCTION for production
		X_PARTNER_ID: os.Getenv("X_PARTNER_ID"),
		PRIVATE_KEY:  os.Getenv("PRIVATE_KEY"),
		ORIGIN:       os.Getenv("ORIGIN"),
		// PRIVATE_KEY_PATH: os.Getenv("PRIVATE_KEY_PATH"),
	}
	apiClient := dana.NewAPIClient(configuration)
	_, r, err := apiClient.PaymentGatewayAPI.UpdateShop(context.Background()).UpdateShopRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentGatewayAPI.UpdateShop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateShop`: UpdateShopResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentGatewayAPI.UpdateShop`: %v\n", r.Body)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateShopRequest** | [**UpdateShopRequest**](MerchantManagement/UpdateShopRequest.md) |  | 

### Return type

[**UpdateShopResponse**](MerchantManagement/UpdateShopResponse.md)

### Authorization

[CLIENT_SECRET](../README.md#CLIENT_SECRET)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# Enum Types

```go
import merchant_management "github.com/dana-id/dana-go/merchant_management/v1"

ipg := string(merchant_management.SOURCEPLATFORM_IPG_)
```

## actorType
| Value | Description |
|-------|-------------|
| `USER` | User |
| `MERCHANT` | Merchant<br |
| `MERCHANT_OPERATOR` | Merchant operator |
| `BACK_OFFICE` | Back office |
| `SYSTEM` | System |


## businessEntity
| Value | Description |
|-------|-------------|
| `pt` |  |
| `cv` |  |
| `individu` |  |
| `usaha_dagang` |  |
| `yayasan` |  |
| `koperasi` |  |


## divisionIdType
| Value | Description |
|-------|-------------|
| `INNER_ID` |  |
| `EXTERNAL_ID` |  |


## divisionType
| Value | Description |
|-------|-------------|
| `REGION` |  |
| `AREA` |  |
| `BRANCH` |  |
| `OUTLET` |  |
| `STORE` |  |
| `KIOSK` |  |
| `STALL` |  |
| `COUNTER` |  |
| `BOOTH` |  |
| `POINT_OF_SALE` |  |
| `VENDING_MACHINE` |  |


## docType
| Value | Description |
|-------|-------------|
| `KTP` |  |
| `SIM` |  |
| `SIUP` |  |
| `NIB` |  |


## loyalty
| Value | Description |
|-------|-------------|
| `true` |  |
| `false` |  |


## orderTerminalType
| Value | Description |
|-------|-------------|
| `APP` | Mobile Application |
| `WEB` | Browser Web |
| `WAP` | Mobile Wap |
| `SYSTEM` | System Call |


## ownerIdType
| Value | Description |
|-------|-------------|
| `KTP` |  |
| `SIM` |  |
| `PASSPORT` |  |
| `SIUP` |  |
| `NIB` |  |


## parentRoleType
| Value | Description |
|-------|-------------|
| `MERCHANT` |  |
| `DIVISION` |  |
| `EXTERNAL_DIVISION` |  |


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


## pgDivisionFlag
| Value | Description |
|-------|-------------|
| `true` |  |
| `false` |  |


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


## shopBizType
| Value | Description |
|-------|-------------|
| `ONLINE` |  |
| `OFFLINE` |  |


## shopIdType
| Value | Description |
|-------|-------------|
| `INNER_ID` |  |
| `EXTERNAL_ID` |  |


## shopOwning
| Value | Description |
|-------|-------------|
| `DIRECT_OWNED` |  |
| `FRANCHISED` |  |


## shopParentType
| Value | Description |
|-------|-------------|
| `MERCHANT` |  |
| `DIVISION` |  |
| `EXTERNAL_DIVISION` |  |


## sizeType
| Value | Description |
|-------|-------------|
| `UMI` |  |
| `UKE` |  |
| `UME` |  |
| `UBE` |  |


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


## verified
| Value | Description |
|-------|-------------|
| `true` |  |
| `false` |  |




