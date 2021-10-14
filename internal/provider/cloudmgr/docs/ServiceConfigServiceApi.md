# \ServiceConfigServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceConfig**](ServiceConfigServiceApi.md#CreateServiceConfig) | **Post** /ksyun/ccb/api/v2/configs/creation | Save service configuration for creating, expanding or shrinking action.
[**DescribeServiceConfig**](ServiceConfigServiceApi.md#DescribeServiceConfig) | **Post** /ksyun/ccb/api/v2/configs/inquiration | Inquire result of service configuration executing
[**ExecuteServiceConfig**](ServiceConfigServiceApi.md#ExecuteServiceConfig) | **Post** /ksyun/ccb/api/v2/configs/execution | Create, expand or shrink service by configuration



## CreateServiceConfig

> KsyunCreateServiceConfigResponse CreateServiceConfig(ctx).Body(body).Execute()

Save service configuration for creating, expanding or shrinking action.

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
    body := *openapiclient.NewKsyunCreateServiceConfigRequest() // KsyunCreateServiceConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceConfigServiceApi.CreateServiceConfig(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceConfigServiceApi.CreateServiceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceConfig`: KsyunCreateServiceConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceConfigServiceApi.CreateServiceConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**KsyunCreateServiceConfigRequest**](KsyunCreateServiceConfigRequest.md) |  | 

### Return type

[**KsyunCreateServiceConfigResponse**](KsyunCreateServiceConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeServiceConfig

> KsyunDescribeServiceConfigResponse DescribeServiceConfig(ctx).Body(body).Execute()

Inquire result of service configuration executing

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
    body := *openapiclient.NewKsyunDescribeServiceConfigRequest() // KsyunDescribeServiceConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceConfigServiceApi.DescribeServiceConfig(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceConfigServiceApi.DescribeServiceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeServiceConfig`: KsyunDescribeServiceConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceConfigServiceApi.DescribeServiceConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeServiceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**KsyunDescribeServiceConfigRequest**](KsyunDescribeServiceConfigRequest.md) |  | 

### Return type

[**KsyunDescribeServiceConfigResponse**](KsyunDescribeServiceConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteServiceConfig

> KsyunExecuteServiceConfigResponse ExecuteServiceConfig(ctx).Body(body).Execute()

Create, expand or shrink service by configuration

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
    body := *openapiclient.NewKsyunExecuteServiceConfigRequest() // KsyunExecuteServiceConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceConfigServiceApi.ExecuteServiceConfig(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceConfigServiceApi.ExecuteServiceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteServiceConfig`: KsyunExecuteServiceConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceConfigServiceApi.ExecuteServiceConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteServiceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**KsyunExecuteServiceConfigRequest**](KsyunExecuteServiceConfigRequest.md) |  | 

### Return type

[**KsyunExecuteServiceConfigResponse**](KsyunExecuteServiceConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

