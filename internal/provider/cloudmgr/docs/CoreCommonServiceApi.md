# \CoreCommonServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DescribeAlertConfig**](CoreCommonServiceApi.md#DescribeAlertConfig) | **Get** /default/api/v2/commons/alert | Describe alert availability of system and a service.
[**InitializeTenant**](CoreCommonServiceApi.md#InitializeTenant) | **Post** /default/api/v2/commons/tenant/onboarding | Customer onboarding



## DescribeAlertConfig

> CoreDescribeAlertConfigResponse DescribeAlertConfig(ctx).Id(id).Execute()

Describe alert availability of system and a service.

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
    id := "id_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreCommonServiceApi.DescribeAlertConfig(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreCommonServiceApi.DescribeAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeAlertConfig`: CoreDescribeAlertConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreCommonServiceApi.DescribeAlertConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

[**CoreDescribeAlertConfigResponse**](CoreDescribeAlertConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeTenant

> map[string]interface{} InitializeTenant(ctx).Body(body).Execute()

Customer onboarding

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
    body := *openapiclient.NewCoreInitializeTenantRequest() // CoreInitializeTenantRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreCommonServiceApi.InitializeTenant(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreCommonServiceApi.InitializeTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitializeTenant`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CoreCommonServiceApi.InitializeTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitializeTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreInitializeTenantRequest**](CoreInitializeTenantRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

