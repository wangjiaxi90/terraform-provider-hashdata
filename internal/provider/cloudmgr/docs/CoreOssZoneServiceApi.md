# \CoreOssZoneServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOssZone**](CoreOssZoneServiceApi.md#CreateOssZone) | **Post** /default/api/v2/oss/zones | Create an object storage zone.
[**DeleteOssZone**](CoreOssZoneServiceApi.md#DeleteOssZone) | **Delete** /default/api/v2/oss/zones/{name} | Delete an object storage zone.
[**DescribeOssZone**](CoreOssZoneServiceApi.md#DescribeOssZone) | **Get** /default/api/v2/oss/zones/{name} | Describe the detail of a given object storage zone.
[**ListOssZone**](CoreOssZoneServiceApi.md#ListOssZone) | **Get** /default/api/v2/oss/zones | List all object storage zone.



## CreateOssZone

> CoreDescribeOssZoneResponse CreateOssZone(ctx).Body(body).Execute()

Create an object storage zone.

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
    body := *openapiclient.NewCoreCreateOssZoneRequest() // CoreCreateOssZoneRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreOssZoneServiceApi.CreateOssZone(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreOssZoneServiceApi.CreateOssZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOssZone`: CoreDescribeOssZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreOssZoneServiceApi.CreateOssZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOssZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateOssZoneRequest**](CoreCreateOssZoneRequest.md) |  | 

### Return type

[**CoreDescribeOssZoneResponse**](CoreDescribeOssZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOssZone

> CommonActionResponse DeleteOssZone(ctx, name).Execute()

Delete an object storage zone.

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
    resp, r, err := api_client.CoreOssZoneServiceApi.DeleteOssZone(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreOssZoneServiceApi.DeleteOssZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOssZone`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreOssZoneServiceApi.DeleteOssZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOssZoneRequest struct via the builder pattern


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


## DescribeOssZone

> CoreDescribeOssZoneResponse DescribeOssZone(ctx, name).Execute()

Describe the detail of a given object storage zone.

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
    resp, r, err := api_client.CoreOssZoneServiceApi.DescribeOssZone(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreOssZoneServiceApi.DescribeOssZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeOssZone`: CoreDescribeOssZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreOssZoneServiceApi.DescribeOssZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeOssZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeOssZoneResponse**](CoreDescribeOssZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOssZone

> CoreListOssZoneResponse ListOssZone(ctx).Name(name).Vendor(vendor).IaasZone(iaasZone).ResourcePool(resourcePool).AvailableOnly(availableOnly).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all object storage zone.

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
    iaasZone := []string{"Inner_example"} // []string |  (optional)
    resourcePool := []string{"Inner_example"} // []string |  (optional)
    availableOnly := true // bool |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreOssZoneServiceApi.ListOssZone(context.Background()).Name(name).Vendor(vendor).IaasZone(iaasZone).ResourcePool(resourcePool).AvailableOnly(availableOnly).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreOssZoneServiceApi.ListOssZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOssZone`: CoreListOssZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreOssZoneServiceApi.ListOssZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOssZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **[]string** |  | 
 **vendor** | **[]string** |  | 
 **iaasZone** | **[]string** |  | 
 **resourcePool** | **[]string** |  | 
 **availableOnly** | **bool** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListOssZoneResponse**](CoreListOssZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

