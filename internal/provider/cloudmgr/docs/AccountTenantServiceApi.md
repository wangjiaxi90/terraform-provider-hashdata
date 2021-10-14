# \AccountTenantServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalTenant**](AccountTenantServiceApi.md#CreateExternalTenant) | **Post** /account/api/v2/tenants/external | Create a new tenant from outside.
[**CreateTenant**](AccountTenantServiceApi.md#CreateTenant) | **Post** /account/api/v2/tenants | Create a new tenant.
[**DeleteTenant**](AccountTenantServiceApi.md#DeleteTenant) | **Delete** /account/api/v2/tenants/{tenantid} | Delete a tenant.
[**DescribeExternalTenant**](AccountTenantServiceApi.md#DescribeExternalTenant) | **Get** /account/api/v2/tenants/external | Describe the detail of tenant.
[**DescribeTenant**](AccountTenantServiceApi.md#DescribeTenant) | **Get** /account/api/v2/tenants/{tenantid} | Describe the detail of tenant.
[**ListTenant**](AccountTenantServiceApi.md#ListTenant) | **Get** /account/api/v2/tenants | List all tenants.



## CreateExternalTenant

> AccountDescribeTenantResponse CreateExternalTenant(ctx).Body(body).Execute()

Create a new tenant from outside.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAccountCreateExternalTenantRequest() // AccountCreateExternalTenantRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTenantServiceApi.CreateExternalTenant(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTenantServiceApi.CreateExternalTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExternalTenant`: AccountDescribeTenantResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTenantServiceApi.CreateExternalTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccountCreateExternalTenantRequest**](AccountCreateExternalTenantRequest.md) |  | 

### Return type

[**AccountDescribeTenantResponse**](AccountDescribeTenantResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTenant

> AccountDescribeTenantResponse CreateTenant(ctx).Body(body).Execute()

Create a new tenant.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAccountCreateTenantRequest() // AccountCreateTenantRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTenantServiceApi.CreateTenant(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTenantServiceApi.CreateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTenant`: AccountDescribeTenantResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTenantServiceApi.CreateTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccountCreateTenantRequest**](AccountCreateTenantRequest.md) |  | 

### Return type

[**AccountDescribeTenantResponse**](AccountDescribeTenantResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenant

> CommonActionResponse DeleteTenant(ctx, tenantid).Execute()

Delete a tenant.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantid := "tenantid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTenantServiceApi.DeleteTenant(context.Background(), tenantid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTenantServiceApi.DeleteTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTenant`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTenantServiceApi.DeleteTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommonActionResponse**](CommonActionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeExternalTenant

> AccountDescribeTenantResponse DescribeExternalTenant(ctx).ExternalId(externalId).Vendor(vendor).Execute()

Describe the detail of tenant.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    externalId := "externalId_example" // string |  (optional)
    vendor := "vendor_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTenantServiceApi.DescribeExternalTenant(context.Background()).ExternalId(externalId).Vendor(vendor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTenantServiceApi.DescribeExternalTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeExternalTenant`: AccountDescribeTenantResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTenantServiceApi.DescribeExternalTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeExternalTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalId** | **string** |  | 
 **vendor** | **string** |  | 

### Return type

[**AccountDescribeTenantResponse**](AccountDescribeTenantResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTenant

> AccountDescribeTenantResponse DescribeTenant(ctx, tenantid).Execute()

Describe the detail of tenant.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantid := "tenantid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTenantServiceApi.DescribeTenant(context.Background(), tenantid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTenantServiceApi.DescribeTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeTenant`: AccountDescribeTenantResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTenantServiceApi.DescribeTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountDescribeTenantResponse**](AccountDescribeTenantResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTenant

> AccountListTenantResponse ListTenant(ctx).Id(id).Name(name).Verified(verified).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all tenants.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := []string{"Inner_example"} // []string |  (optional)
    name := []string{"Inner_example"} // []string |  (optional)
    verified := true // bool |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTenantServiceApi.ListTenant(context.Background()).Id(id).Name(name).Verified(verified).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTenantServiceApi.ListTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTenant`: AccountListTenantResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTenantServiceApi.ListTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **name** | **[]string** |  | 
 **verified** | **bool** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**AccountListTenantResponse**](AccountListTenantResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

