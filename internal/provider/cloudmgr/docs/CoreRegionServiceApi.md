# \CoreRegionServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegion**](CoreRegionServiceApi.md#CreateRegion) | **Post** /default/api/v2/regions | Create a new data center region.
[**DeleteRegion**](CoreRegionServiceApi.md#DeleteRegion) | **Delete** /default/api/v2/regions/{name} | Delete a data center region.
[**DescribeRegion**](CoreRegionServiceApi.md#DescribeRegion) | **Get** /default/api/v2/regions/{name} | Describe the detail of a given data center region.
[**ListRegion**](CoreRegionServiceApi.md#ListRegion) | **Get** /default/api/v2/regions | List all data center region.



## CreateRegion

> CoreDescribeRegionResponse CreateRegion(ctx).Body(body).Execute()

Create a new data center region.

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
    body := *openapiclient.NewCoreCreateRegionRequest() // CoreCreateRegionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreRegionServiceApi.CreateRegion(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreRegionServiceApi.CreateRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRegion`: CoreDescribeRegionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreRegionServiceApi.CreateRegion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateRegionRequest**](CoreCreateRegionRequest.md) |  | 

### Return type

[**CoreDescribeRegionResponse**](CoreDescribeRegionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegion

> CommonActionResponse DeleteRegion(ctx, name).Execute()

Delete a data center region.

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreRegionServiceApi.DeleteRegion(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreRegionServiceApi.DeleteRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRegion`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreRegionServiceApi.DeleteRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegionRequest struct via the builder pattern


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


## DescribeRegion

> CoreDescribeRegionResponse DescribeRegion(ctx, name).Execute()

Describe the detail of a given data center region.

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreRegionServiceApi.DescribeRegion(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreRegionServiceApi.DescribeRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeRegion`: CoreDescribeRegionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreRegionServiceApi.DescribeRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeRegionResponse**](CoreDescribeRegionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegion

> CoreListRegionResponse ListRegion(ctx).Name(name).Vendor(vendor).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all data center region.

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
    name := []string{"Inner_example"} // []string |  (optional)
    vendor := []string{"Inner_example"} // []string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreRegionServiceApi.ListRegion(context.Background()).Name(name).Vendor(vendor).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreRegionServiceApi.ListRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRegion`: CoreListRegionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreRegionServiceApi.ListRegion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **[]string** |  | 
 **vendor** | **[]string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListRegionResponse**](CoreListRegionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

