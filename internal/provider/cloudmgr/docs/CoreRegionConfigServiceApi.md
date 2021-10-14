# \CoreRegionConfigServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegionConfig**](CoreRegionConfigServiceApi.md#CreateRegionConfig) | **Post** /default/api/v2/regions/configs | Create a data center region configuration.
[**DeleteRegionConfig**](CoreRegionConfigServiceApi.md#DeleteRegionConfig) | **Delete** /default/api/v2/regions/configs/{id} | Delete a data center region configuration.
[**DescribeRegionConfig**](CoreRegionConfigServiceApi.md#DescribeRegionConfig) | **Get** /default/api/v2/regions/configs/{id} | Describe the detail of a given data center region configuration.
[**ListRegionConfig**](CoreRegionConfigServiceApi.md#ListRegionConfig) | **Get** /default/api/v2/regions/configs | List all data center region configuration.
[**UpdateRegionConfig**](CoreRegionConfigServiceApi.md#UpdateRegionConfig) | **Put** /default/api/v2/regions/configs/{id} | Update a data center region configuration.



## CreateRegionConfig

> CoreDescribeRegionConfigResponse CreateRegionConfig(ctx).Body(body).Execute()

Create a data center region configuration.

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
    body := *openapiclient.NewCoreCreateRegionConfigRequest() // CoreCreateRegionConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreRegionConfigServiceApi.CreateRegionConfig(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreRegionConfigServiceApi.CreateRegionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRegionConfig`: CoreDescribeRegionConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreRegionConfigServiceApi.CreateRegionConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateRegionConfigRequest**](CoreCreateRegionConfigRequest.md) |  | 

### Return type

[**CoreDescribeRegionConfigResponse**](CoreDescribeRegionConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegionConfig

> CommonActionResponse DeleteRegionConfig(ctx, id).Execute()

Delete a data center region configuration.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreRegionConfigServiceApi.DeleteRegionConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreRegionConfigServiceApi.DeleteRegionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRegionConfig`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreRegionConfigServiceApi.DeleteRegionConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegionConfigRequest struct via the builder pattern


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


## DescribeRegionConfig

> CoreDescribeRegionConfigResponse DescribeRegionConfig(ctx, id).Execute()

Describe the detail of a given data center region configuration.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreRegionConfigServiceApi.DescribeRegionConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreRegionConfigServiceApi.DescribeRegionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeRegionConfig`: CoreDescribeRegionConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreRegionConfigServiceApi.DescribeRegionConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeRegionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeRegionConfigResponse**](CoreDescribeRegionConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegionConfig

> CoreListRegionConfigResponse ListRegionConfig(ctx).Id(id).Tenant(tenant).Region(region).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all data center region configuration.

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
    tenant := []string{"Inner_example"} // []string |  (optional)
    region := []string{"Inner_example"} // []string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreRegionConfigServiceApi.ListRegionConfig(context.Background()).Id(id).Tenant(tenant).Region(region).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreRegionConfigServiceApi.ListRegionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRegionConfig`: CoreListRegionConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreRegionConfigServiceApi.ListRegionConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRegionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **tenant** | **[]string** |  | 
 **region** | **[]string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListRegionConfigResponse**](CoreListRegionConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRegionConfig

> CoreDescribeRegionConfigResponse UpdateRegionConfig(ctx, id).Body(body).Execute()

Update a data center region configuration.

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
    id := "id_example" // string | 
    body := *openapiclient.NewCoreUpdateRegionConfigRequest() // CoreUpdateRegionConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreRegionConfigServiceApi.UpdateRegionConfig(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreRegionConfigServiceApi.UpdateRegionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRegionConfig`: CoreDescribeRegionConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreRegionConfigServiceApi.UpdateRegionConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreUpdateRegionConfigRequest**](CoreUpdateRegionConfigRequest.md) |  | 

### Return type

[**CoreDescribeRegionConfigResponse**](CoreDescribeRegionConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

