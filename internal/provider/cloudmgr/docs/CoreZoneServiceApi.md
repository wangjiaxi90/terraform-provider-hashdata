# \CoreZoneServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateZone**](CoreZoneServiceApi.md#CreateZone) | **Post** /default/api/v2/zones | Create a new data center zone.
[**DeleteZone**](CoreZoneServiceApi.md#DeleteZone) | **Delete** /default/api/v2/zones/{name} | Delete a data center zone.
[**DescribeZone**](CoreZoneServiceApi.md#DescribeZone) | **Get** /default/api/v2/zones/{name} | Describe the detail of a given data center zone.
[**ListZone**](CoreZoneServiceApi.md#ListZone) | **Get** /default/api/v2/zones | List all data center zone.



## CreateZone

> CoreDescribeZoneResponse CreateZone(ctx).Body(body).Execute()

Create a new data center zone.

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
    body := *openapiclient.NewCoreCreateZoneRequest() // CoreCreateZoneRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreZoneServiceApi.CreateZone(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreZoneServiceApi.CreateZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateZone`: CoreDescribeZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreZoneServiceApi.CreateZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateZoneRequest**](CoreCreateZoneRequest.md) |  | 

### Return type

[**CoreDescribeZoneResponse**](CoreDescribeZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteZone

> CommonActionResponse DeleteZone(ctx, name).Execute()

Delete a data center zone.

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
    resp, r, err := api_client.CoreZoneServiceApi.DeleteZone(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreZoneServiceApi.DeleteZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteZone`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreZoneServiceApi.DeleteZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteZoneRequest struct via the builder pattern


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


## DescribeZone

> CoreDescribeZoneResponse DescribeZone(ctx, name).Execute()

Describe the detail of a given data center zone.

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
    resp, r, err := api_client.CoreZoneServiceApi.DescribeZone(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreZoneServiceApi.DescribeZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeZone`: CoreDescribeZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreZoneServiceApi.DescribeZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeZoneResponse**](CoreDescribeZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListZone

> CoreListZoneResponse ListZone(ctx).Name(name).Vendor(vendor).Region(region).AvailableOnly(availableOnly).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all data center zone.

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
    region := []string{"Inner_example"} // []string |  (optional)
    availableOnly := true // bool |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreZoneServiceApi.ListZone(context.Background()).Name(name).Vendor(vendor).Region(region).AvailableOnly(availableOnly).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreZoneServiceApi.ListZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListZone`: CoreListZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreZoneServiceApi.ListZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **[]string** |  | 
 **vendor** | **[]string** |  | 
 **region** | **[]string** |  | 
 **availableOnly** | **bool** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListZoneResponse**](CoreListZoneResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

