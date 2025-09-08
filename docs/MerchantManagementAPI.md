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
	request := merchant_management.CreateDivisionRequest{
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
	_, r, err := apiClient.MerchantManagementAPI.CreateDivision(context.Background()).CreateDivisionRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantManagementAPI.CreateDivision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDivision`: CreateDivisionResponse
	fmt.Fprintf(os.Stdout, "Response from `MerchantManagementAPI.CreateDivision`: %v\n", r.Body)
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
	request := merchant_management.CreateShopRequest{
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
	_, r, err := apiClient.MerchantManagementAPI.CreateShop(context.Background()).CreateShopRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantManagementAPI.CreateShop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShop`: CreateShopResponse
	fmt.Fprintf(os.Stdout, "Response from `MerchantManagementAPI.CreateShop`: %v\n", r.Body)
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
	request := merchant_management.QueryDivisionRequest{
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
	_, r, err := apiClient.MerchantManagementAPI.QueryDivision(context.Background()).QueryDivisionRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantManagementAPI.QueryDivision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryDivision`: QueryDivisionResponse
	fmt.Fprintf(os.Stdout, "Response from `MerchantManagementAPI.QueryDivision`: %v\n", r.Body)
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
	request := merchant_management.QueryMerchantResourceRequest{
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
	_, r, err := apiClient.MerchantManagementAPI.QueryMerchantResource(context.Background()).QueryMerchantResourceRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantManagementAPI.QueryMerchantResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryMerchantResource`: QueryMerchantResourceResponse
	fmt.Fprintf(os.Stdout, "Response from `MerchantManagementAPI.QueryMerchantResource`: %v\n", r.Body)
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
	request := merchant_management.QueryShopRequest{
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
	_, r, err := apiClient.MerchantManagementAPI.QueryShop(context.Background()).QueryShopRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantManagementAPI.QueryShop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryShop`: QueryShopResponse
	fmt.Fprintf(os.Stdout, "Response from `MerchantManagementAPI.QueryShop`: %v\n", r.Body)
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
	request := merchant_management.UpdateDivisionRequest{
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
	_, r, err := apiClient.MerchantManagementAPI.UpdateDivision(context.Background()).UpdateDivisionRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantManagementAPI.UpdateDivision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDivision`: UpdateDivisionResponse
	fmt.Fprintf(os.Stdout, "Response from `MerchantManagementAPI.UpdateDivision`: %v\n", r.Body)
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
	request := merchant_management.UpdateShopRequest{
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
	_, r, err := apiClient.MerchantManagementAPI.UpdateShop(context.Background()).UpdateShopRequest(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantManagementAPI.UpdateShop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateShop`: UpdateShopResponse
	fmt.Fprintf(os.Stdout, "Response from `MerchantManagementAPI.UpdateShop`: %v\n", r.Body)
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

value := string(merchant_management.SHOPPARENTTYPE_MERCHANT_)
```

## GOODS_SOLD_TYPE
| Value | Description |
|-------|-------------|
| `GOODS_SOLD_TYPE_DIGITAL_` |  |
| `GOODS_SOLD_TYPE_PHYSICAL_` |  |


## USER_PROFILING
| Value | Description |
|-------|-------------|
| `USER_PROFILING_B2B_` |  |
| `USER_PROFILING_B2C_` |  |


## BusinessEntity
| Value | Description |
|-------|-------------|
| `BUSINESSENTITY_PT_` |  |
| `BUSINESSENTITY_CV_` |  |
| `BUSINESSENTITY_INDIVIDU_` |  |
| `BUSINESSENTITY_USAHA_DAGANG_` |  |
| `BUSINESSENTITY_YAYASAN_` |  |
| `BUSINESSENTITY_KOPERASI_` |  |


## DivisionIdType
| Value | Description |
|-------|-------------|
| `DIVISIONIDTYPE_INNER_ID_` |  |
| `DIVISIONIDTYPE_EXTERNAL_ID_` |  |


## DivisionType
| Value | Description |
|-------|-------------|
| `DIVISIONTYPE_REGION_` |  |
| `DIVISIONTYPE_AREA_` |  |
| `DIVISIONTYPE_BRANCH_` |  |
| `DIVISIONTYPE_OUTLET_` |  |
| `DIVISIONTYPE_STORE_` |  |
| `DIVISIONTYPE_KIOSK_` |  |
| `DIVISIONTYPE_STALL_` |  |
| `DIVISIONTYPE_COUNTER_` |  |
| `DIVISIONTYPE_BOOTH_` |  |
| `DIVISIONTYPE_POINT_OF_SALE_` |  |
| `DIVISIONTYPE_VENDING_MACHINE_` |  |


## DocType
| Value | Description |
|-------|-------------|
| `DOCTYPE_KTP_` |  |
| `DOCTYPE_SIM_` |  |
| `DOCTYPE_SIUP_` |  |
| `DOCTYPE_NIB_` |  |


## Loyalty
| Value | Description |
|-------|-------------|
| `LOYALTY_TRUE_` |  |
| `LOYALTY_FALSE_` |  |


## OwnerIdType
| Value | Description |
|-------|-------------|
| `OWNERIDTYPE_KTP_` |  |
| `OWNERIDTYPE_SIM_` |  |
| `OWNERIDTYPE_PASSPORT_` |  |
| `OWNERIDTYPE_SIUP_` |  |
| `OWNERIDTYPE_NIB_` |  |


## ParentRoleType
| Value | Description |
|-------|-------------|
| `PARENTROLETYPE_MERCHANT_` |  |
| `PARENTROLETYPE_DIVISION_` |  |
| `PARENTROLETYPE_EXTERNAL_DIVISION_` |  |


## PgDivisionFlag
| Value | Description |
|-------|-------------|
| `PGDIVISIONFLAG_TRUE_` |  |
| `PGDIVISIONFLAG_FALSE_` |  |


## ResourceType
| Value | Description |
|-------|-------------|
| `RESOURCETYPE_MERCHANT_DEPOSIT_BALANCE_` |  |
| `RESOURCETYPE_MERCHANT_AVAILABLE_BALANCE_` |  |
| `RESOURCETYPE_MERCHANT_TOTAL_BALANCE_` |  |


## ResultStatus
| Value | Description |
|-------|-------------|
| `RESULTSTATUS_S_` |  |
| `RESULTSTATUS_F_` |  |
| `RESULTSTATUS_U_` |  |


## ShopBizType
| Value | Description |
|-------|-------------|
| `SHOPBIZTYPE_ONLINE_` |  |
| `SHOPBIZTYPE_OFFLINE_` |  |


## ShopIdType
| Value | Description |
|-------|-------------|
| `SHOPIDTYPE_INNER_ID_` |  |
| `SHOPIDTYPE_EXTERNAL_ID_` |  |


## ShopOwning
| Value | Description |
|-------|-------------|
| `SHOPOWNING_DIRECT_OWNED_` |  |
| `SHOPOWNING_FRANCHISED_` |  |


## ShopParentType
| Value | Description |
|-------|-------------|
| `SHOPPARENTTYPE_MERCHANT_` |  |
| `SHOPPARENTTYPE_DIVISION_` |  |
| `SHOPPARENTTYPE_EXTERNAL_DIVISION_` |  |


## SizeType
| Value | Description |
|-------|-------------|
| `SIZETYPE_UMI_` |  |
| `SIZETYPE_UKE_` |  |
| `SIZETYPE_UME_` |  |
| `SIZETYPE_UBE_` |  |


## Verified
| Value | Description |
|-------|-------------|
| `VERIFIED_TRUE_` |  |
| `VERIFIED_FALSE_` |  |




