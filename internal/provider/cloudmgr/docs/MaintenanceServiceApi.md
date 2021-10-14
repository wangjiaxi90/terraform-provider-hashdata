# \MaintenanceServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigThreadPool**](MaintenanceServiceApi.md#ConfigThreadPool) | **Post** /default/api/v2/system/maintenance/configure/thread-pool | Configure thread pool.



## ConfigThreadPool

> CommonActionResponse ConfigThreadPool(ctx).Body(body).Execute()

Configure thread pool.

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
    body := *openapiclient.NewCoreConfigThreadPoolRequest() // CoreConfigThreadPoolRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintenanceServiceApi.ConfigThreadPool(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceServiceApi.ConfigThreadPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigThreadPool`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceServiceApi.ConfigThreadPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigThreadPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreConfigThreadPoolRequest**](CoreConfigThreadPoolRequest.md) |  | 

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

