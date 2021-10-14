# \CoreHdfsConfigServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateHdfsConfig**](CoreHdfsConfigServiceApi.md#ActivateHdfsConfig) | **Put** /default/api/v2/services/{service}/hdfs/{timestamp}/activation | Activate a specified hdfs configuration of a service.
[**CreateHdfsConfig**](CoreHdfsConfigServiceApi.md#CreateHdfsConfig) | **Post** /default/api/v2/services/{service}/hdfs | Create a service hdfs configuration.
[**DeleteHdfsConfig**](CoreHdfsConfigServiceApi.md#DeleteHdfsConfig) | **Delete** /default/api/v2/service/{service}/hdfs/{timestamp} | Delete a hdfs configuration of a service.
[**DescribeHdfsConfig**](CoreHdfsConfigServiceApi.md#DescribeHdfsConfig) | **Get** /default/api/v2/services/{service}/hdfs/{timestamp} | Describe the detail of a given hdfs configuration.
[**ListHdfsConfig**](CoreHdfsConfigServiceApi.md#ListHdfsConfig) | **Get** /default/api/v2/services/{service}/hdfs | List the hdfs configuration of a given service.



## ActivateHdfsConfig

> CommonDescribeJobResponse ActivateHdfsConfig(ctx, service, timestamp).Execute()

Activate a specified hdfs configuration of a service.

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
    service := "service_example" // string | 
    timestamp := "timestamp_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreHdfsConfigServiceApi.ActivateHdfsConfig(context.Background(), service, timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreHdfsConfigServiceApi.ActivateHdfsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateHdfsConfig`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreHdfsConfigServiceApi.ActivateHdfsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**timestamp** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateHdfsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CommonDescribeJobResponse**](CommonDescribeJobResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHdfsConfig

> CommonDescribeJobResponse CreateHdfsConfig(ctx, service).Body(body).Execute()

Create a service hdfs configuration.

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
    service := "service_example" // string | 
    body := *openapiclient.NewCoreCreateHdfsConfigRequest() // CoreCreateHdfsConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreHdfsConfigServiceApi.CreateHdfsConfig(context.Background(), service).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreHdfsConfigServiceApi.CreateHdfsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHdfsConfig`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreHdfsConfigServiceApi.CreateHdfsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHdfsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreCreateHdfsConfigRequest**](CoreCreateHdfsConfigRequest.md) |  | 

### Return type

[**CommonDescribeJobResponse**](CommonDescribeJobResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHdfsConfig

> CommonActionResponse DeleteHdfsConfig(ctx, service, timestamp).Execute()

Delete a hdfs configuration of a service.

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
    service := "service_example" // string | 
    timestamp := "timestamp_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreHdfsConfigServiceApi.DeleteHdfsConfig(context.Background(), service, timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreHdfsConfigServiceApi.DeleteHdfsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHdfsConfig`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreHdfsConfigServiceApi.DeleteHdfsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**timestamp** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHdfsConfigRequest struct via the builder pattern


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


## DescribeHdfsConfig

> CoreDescribeHdfsConfigResponse DescribeHdfsConfig(ctx, service, timestamp).Execute()

Describe the detail of a given hdfs configuration.

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
    service := "service_example" // string | 
    timestamp := "timestamp_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreHdfsConfigServiceApi.DescribeHdfsConfig(context.Background(), service, timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreHdfsConfigServiceApi.DescribeHdfsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeHdfsConfig`: CoreDescribeHdfsConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreHdfsConfigServiceApi.DescribeHdfsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**timestamp** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeHdfsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CoreDescribeHdfsConfigResponse**](CoreDescribeHdfsConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHdfsConfig

> CoreListHdfsConfigResponse ListHdfsConfig(ctx, service).Offset(offset).Limit(limit).Execute()

List the hdfs configuration of a given service.

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
    service := "service_example" // string | 
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreHdfsConfigServiceApi.ListHdfsConfig(context.Background(), service).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreHdfsConfigServiceApi.ListHdfsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHdfsConfig`: CoreListHdfsConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreHdfsConfigServiceApi.ListHdfsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHdfsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListHdfsConfigResponse**](CoreListHdfsConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

