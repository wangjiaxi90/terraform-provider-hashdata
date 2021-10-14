# \ResourceUsageServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateResourceUsage**](ResourceUsageServiceApi.md#CalculateResourceUsage) | **Get** /ksyun/ccb/api/v2/tenants/configs | Calculate resource usage for configuring cluster
[**DescribeResourceUsage**](ResourceUsageServiceApi.md#DescribeResourceUsage) | **Get** /ksyun/ccb/api/v2/tenants/regions/{region} | Describe the resource usage of a given region.
[**ListResourceUsage**](ResourceUsageServiceApi.md#ListResourceUsage) | **Get** /ksyun/ccb/api/v2/tenants | List resource usage for all tenants.



## CalculateResourceUsage

> KsyunCalculateResourceUsageResponse CalculateResourceUsage(ctx).ConfId(confId).Execute()

Calculate resource usage for configuring cluster

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
    confId := "confId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceUsageServiceApi.CalculateResourceUsage(context.Background()).ConfId(confId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceUsageServiceApi.CalculateResourceUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CalculateResourceUsage`: KsyunCalculateResourceUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceUsageServiceApi.CalculateResourceUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalculateResourceUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confId** | **string** |  | 

### Return type

[**KsyunCalculateResourceUsageResponse**](KsyunCalculateResourceUsageResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeResourceUsage

> KsyunDescribeResourceUsageResponse DescribeResourceUsage(ctx, region).Execute()

Describe the resource usage of a given region.

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
    region := "region_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceUsageServiceApi.DescribeResourceUsage(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceUsageServiceApi.DescribeResourceUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeResourceUsage`: KsyunDescribeResourceUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceUsageServiceApi.DescribeResourceUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeResourceUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KsyunDescribeResourceUsageResponse**](KsyunDescribeResourceUsageResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceUsage

> KsyunListResourceUsageResponse ListResourceUsage(ctx).Execute()

List resource usage for all tenants.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceUsageServiceApi.ListResourceUsage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceUsageServiceApi.ListResourceUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListResourceUsage`: KsyunListResourceUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceUsageServiceApi.ListResourceUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListResourceUsageRequest struct via the builder pattern


### Return type

[**KsyunListResourceUsageResponse**](KsyunListResourceUsageResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

