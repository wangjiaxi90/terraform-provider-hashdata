# \CoreVolumeTypeServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolumeType**](CoreVolumeTypeServiceApi.md#CreateVolumeType) | **Post** /default/api/v2/volumes/types | Create a new volume type.
[**DeleteVolumeType**](CoreVolumeTypeServiceApi.md#DeleteVolumeType) | **Delete** /default/api/v2/volumes/types/{name} | Delete a volume type.
[**DescribeVolumeType**](CoreVolumeTypeServiceApi.md#DescribeVolumeType) | **Get** /default/api/v2/volumes/types/{name} | Describe the detail of a given volume type.
[**ListVolumeType**](CoreVolumeTypeServiceApi.md#ListVolumeType) | **Get** /default/api/v2/volumes/types | List all volume type.



## CreateVolumeType

> CoreDescribeVolumeTypeResponse CreateVolumeType(ctx).Body(body).Execute()

Create a new volume type.

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
    body := *openapiclient.NewCoreCreateVolumeTypeRequest() // CoreCreateVolumeTypeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreVolumeTypeServiceApi.CreateVolumeType(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreVolumeTypeServiceApi.CreateVolumeType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVolumeType`: CoreDescribeVolumeTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreVolumeTypeServiceApi.CreateVolumeType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateVolumeTypeRequest**](CoreCreateVolumeTypeRequest.md) |  | 

### Return type

[**CoreDescribeVolumeTypeResponse**](CoreDescribeVolumeTypeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolumeType

> CommonActionResponse DeleteVolumeType(ctx, name).Execute()

Delete a volume type.

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
    resp, r, err := api_client.CoreVolumeTypeServiceApi.DeleteVolumeType(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreVolumeTypeServiceApi.DeleteVolumeType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVolumeType`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreVolumeTypeServiceApi.DeleteVolumeType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeTypeRequest struct via the builder pattern


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


## DescribeVolumeType

> CoreDescribeVolumeTypeResponse DescribeVolumeType(ctx, name).Execute()

Describe the detail of a given volume type.

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
    resp, r, err := api_client.CoreVolumeTypeServiceApi.DescribeVolumeType(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreVolumeTypeServiceApi.DescribeVolumeType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeVolumeType`: CoreDescribeVolumeTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreVolumeTypeServiceApi.DescribeVolumeType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeVolumeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeVolumeTypeResponse**](CoreDescribeVolumeTypeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumeType

> CoreListVolumeTypeResponse ListVolumeType(ctx).Zone(zone).Vendor(vendor).InstanceType(instanceType).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all volume type.

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
    zone := []string{"Inner_example"} // []string |  (optional)
    vendor := []string{"Inner_example"} // []string |  (optional)
    instanceType := []string{"Inner_example"} // []string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreVolumeTypeServiceApi.ListVolumeType(context.Background()).Zone(zone).Vendor(vendor).InstanceType(instanceType).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreVolumeTypeServiceApi.ListVolumeType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVolumeType`: CoreListVolumeTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreVolumeTypeServiceApi.ListVolumeType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **[]string** |  | 
 **vendor** | **[]string** |  | 
 **instanceType** | **[]string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListVolumeTypeResponse**](CoreListVolumeTypeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

