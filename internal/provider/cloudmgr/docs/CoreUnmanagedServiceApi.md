# \CoreUnmanagedServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUnmanagedService**](CoreUnmanagedServiceApi.md#CreateUnmanagedService) | **Post** /default/api/v2/services/unmanaged | Create a unmanaged service.



## CreateUnmanagedService

> CommonDescribeJobResponse CreateUnmanagedService(ctx).Body(body).Execute()

Create a unmanaged service.

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
    body := *openapiclient.NewCoreCreateUnmanagedServiceRequest() // CoreCreateUnmanagedServiceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreUnmanagedServiceApi.CreateUnmanagedService(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreUnmanagedServiceApi.CreateUnmanagedService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUnmanagedService`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreUnmanagedServiceApi.CreateUnmanagedService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUnmanagedServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateUnmanagedServiceRequest**](CoreCreateUnmanagedServiceRequest.md) |  | 

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

