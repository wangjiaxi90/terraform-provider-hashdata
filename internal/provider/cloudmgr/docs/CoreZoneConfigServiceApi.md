# \CoreZoneConfigServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateZoneConfig**](CoreZoneConfigServiceApi.md#CreateZoneConfig) | **Post** /default/api/v2/zones/configs | Create a data center zone configuration.
[**DeleteZoneConfig**](CoreZoneConfigServiceApi.md#DeleteZoneConfig) | **Delete** /default/api/v2/zones/configs/{id} | Delete a data center zone configuration.
[**DescribeZoneConfig**](CoreZoneConfigServiceApi.md#DescribeZoneConfig) | **Get** /default/api/v2/zones/configs/{id} | Describe the detail of a given data center zone configuration.
[**ListZoneConfig**](CoreZoneConfigServiceApi.md#ListZoneConfig) | **Get** /default/api/v2/zones/configs | List all data center zone configuration.
[**UpdateZoneConfig**](CoreZoneConfigServiceApi.md#UpdateZoneConfig) | **Put** /default/api/v2/zones/configs/{id} | Update a data center zone configuration.



## CreateZoneConfig

> CoreDescribeZoneConfigResponse CreateZoneConfig(ctx).Body(body).Execute()

Create a data center zone configuration.

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
    body := *openapiclient.NewCoreCreateZoneConfigRequest() // CoreCreateZoneConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreZoneConfigServiceApi.CreateZoneConfig(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreZoneConfigServiceApi.CreateZoneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateZoneConfig`: CoreDescribeZoneConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreZoneConfigServiceApi.CreateZoneConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateZoneConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateZoneConfigRequest**](CoreCreateZoneConfigRequest.md) |  | 

### Return type

[**CoreDescribeZoneConfigResponse**](CoreDescribeZoneConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteZoneConfig

> CommonActionResponse DeleteZoneConfig(ctx, id).Execute()

Delete a data center zone configuration.

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
    resp, r, err := api_client.CoreZoneConfigServiceApi.DeleteZoneConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreZoneConfigServiceApi.DeleteZoneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteZoneConfig`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreZoneConfigServiceApi.DeleteZoneConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteZoneConfigRequest struct via the builder pattern


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


## DescribeZoneConfig

> CoreDescribeZoneConfigResponse DescribeZoneConfig(ctx, id).Execute()

Describe the detail of a given data center zone configuration.

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
    resp, r, err := api_client.CoreZoneConfigServiceApi.DescribeZoneConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreZoneConfigServiceApi.DescribeZoneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeZoneConfig`: CoreDescribeZoneConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreZoneConfigServiceApi.DescribeZoneConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeZoneConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeZoneConfigResponse**](CoreDescribeZoneConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListZoneConfig

> CoreListZoneConfigResponse ListZoneConfig(ctx).Id(id).Tenant(tenant).Zone(zone).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Region(region).Execute()

List all data center zone configuration.

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
    zone := []string{"Inner_example"} // []string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    region := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreZoneConfigServiceApi.ListZoneConfig(context.Background()).Id(id).Tenant(tenant).Zone(zone).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreZoneConfigServiceApi.ListZoneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListZoneConfig`: CoreListZoneConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreZoneConfigServiceApi.ListZoneConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListZoneConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **tenant** | **[]string** |  | 
 **zone** | **[]string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **region** | **[]string** |  | 

### Return type

[**CoreListZoneConfigResponse**](CoreListZoneConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateZoneConfig

> CoreDescribeZoneConfigResponse UpdateZoneConfig(ctx, id).Body(body).Execute()

Update a data center zone configuration.

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
    body := *openapiclient.NewCoreUpdateZoneConfigRequest() // CoreUpdateZoneConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreZoneConfigServiceApi.UpdateZoneConfig(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreZoneConfigServiceApi.UpdateZoneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateZoneConfig`: CoreDescribeZoneConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreZoneConfigServiceApi.UpdateZoneConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateZoneConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreUpdateZoneConfigRequest**](CoreUpdateZoneConfigRequest.md) |  | 

### Return type

[**CoreDescribeZoneConfigResponse**](CoreDescribeZoneConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

