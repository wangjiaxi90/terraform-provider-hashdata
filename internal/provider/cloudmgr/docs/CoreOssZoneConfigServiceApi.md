# \CoreOssZoneConfigServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOssZoneConfig**](CoreOssZoneConfigServiceApi.md#CreateOssZoneConfig) | **Post** /default/api/v2/oss/zones/configs | Create an object storage service configuration.
[**DeleteOssZoneConfig**](CoreOssZoneConfigServiceApi.md#DeleteOssZoneConfig) | **Delete** /default/api/v2/oss/zones/configs/{id} | Delete an object storage service configuration.
[**DescribeOssZoneConfig**](CoreOssZoneConfigServiceApi.md#DescribeOssZoneConfig) | **Get** /default/api/v2/oss/zones/configs/{id} | Describe the detail of an object storage service configuration.
[**ListOssZoneConfig**](CoreOssZoneConfigServiceApi.md#ListOssZoneConfig) | **Get** /default/api/v2/oss/zones/configs | List all object storage service configuration.
[**UpdateOssZoneConfig**](CoreOssZoneConfigServiceApi.md#UpdateOssZoneConfig) | **Put** /default/api/v2/oss/zones/configs/{id} | Update n object storage service configuration.



## CreateOssZoneConfig

> CoreDescribeOssZoneConfigResponse CreateOssZoneConfig(ctx).Body(body).Execute()

Create an object storage service configuration.

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
    body := *openapiclient.NewCoreCreateOssZoneConfigRequest() // CoreCreateOssZoneConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreOssZoneConfigServiceApi.CreateOssZoneConfig(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreOssZoneConfigServiceApi.CreateOssZoneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOssZoneConfig`: CoreDescribeOssZoneConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreOssZoneConfigServiceApi.CreateOssZoneConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOssZoneConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateOssZoneConfigRequest**](CoreCreateOssZoneConfigRequest.md) |  | 

### Return type

[**CoreDescribeOssZoneConfigResponse**](CoreDescribeOssZoneConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOssZoneConfig

> CommonActionResponse DeleteOssZoneConfig(ctx, id).Execute()

Delete an object storage service configuration.

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
    resp, r, err := api_client.CoreOssZoneConfigServiceApi.DeleteOssZoneConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreOssZoneConfigServiceApi.DeleteOssZoneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOssZoneConfig`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreOssZoneConfigServiceApi.DeleteOssZoneConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOssZoneConfigRequest struct via the builder pattern


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


## DescribeOssZoneConfig

> CoreDescribeOssZoneConfigResponse DescribeOssZoneConfig(ctx, id).Execute()

Describe the detail of an object storage service configuration.

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
    resp, r, err := api_client.CoreOssZoneConfigServiceApi.DescribeOssZoneConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreOssZoneConfigServiceApi.DescribeOssZoneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeOssZoneConfig`: CoreDescribeOssZoneConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreOssZoneConfigServiceApi.DescribeOssZoneConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeOssZoneConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeOssZoneConfigResponse**](CoreDescribeOssZoneConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOssZoneConfig

> CoreListOssZoneConfigResponse ListOssZoneConfig(ctx).Id(id).Tenant(tenant).Zone(zone).Offset(offset).Limit(limit).Execute()

List all object storage service configuration.

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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreOssZoneConfigServiceApi.ListOssZoneConfig(context.Background()).Id(id).Tenant(tenant).Zone(zone).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreOssZoneConfigServiceApi.ListOssZoneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOssZoneConfig`: CoreListOssZoneConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreOssZoneConfigServiceApi.ListOssZoneConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOssZoneConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **tenant** | **[]string** |  | 
 **zone** | **[]string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListOssZoneConfigResponse**](CoreListOssZoneConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOssZoneConfig

> CoreDescribeOssZoneConfigResponse UpdateOssZoneConfig(ctx, id).Body(body).Execute()

Update n object storage service configuration.

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
    body := *openapiclient.NewCoreUpdateOssZoneConfigRequest() // CoreUpdateOssZoneConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreOssZoneConfigServiceApi.UpdateOssZoneConfig(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreOssZoneConfigServiceApi.UpdateOssZoneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOssZoneConfig`: CoreDescribeOssZoneConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreOssZoneConfigServiceApi.UpdateOssZoneConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOssZoneConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreUpdateOssZoneConfigRequest**](CoreUpdateOssZoneConfigRequest.md) |  | 

### Return type

[**CoreDescribeOssZoneConfigResponse**](CoreDescribeOssZoneConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

