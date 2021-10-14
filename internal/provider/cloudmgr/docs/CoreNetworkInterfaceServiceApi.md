# \CoreNetworkInterfaceServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachNetworkInterface**](CoreNetworkInterfaceServiceApi.md#AttachNetworkInterface) | **Post** /default/api/v2/network/nics/{id}/attach | Attach a nic to specified instance.
[**CreateNetworkInterface**](CoreNetworkInterfaceServiceApi.md#CreateNetworkInterface) | **Post** /default/api/v2/network/nics | Create a new nic.
[**DeleteNetworkInterface**](CoreNetworkInterfaceServiceApi.md#DeleteNetworkInterface) | **Delete** /default/api/v2/network/nics/{id} | release a nic.
[**DescribeNetworkInterface**](CoreNetworkInterfaceServiceApi.md#DescribeNetworkInterface) | **Get** /default/api/v2/network/nics/{id} | Describe an eip detail.
[**DetachNetworkInterface**](CoreNetworkInterfaceServiceApi.md#DetachNetworkInterface) | **Get** /default/api/v2/network/nics/{id}/detach | Detach a nic from instance.
[**ListNetworkInterface**](CoreNetworkInterfaceServiceApi.md#ListNetworkInterface) | **Get** /default/api/v2/network/nics | List all nics.



## AttachNetworkInterface

> CommonActionResponse AttachNetworkInterface(ctx, id).Body(body).Execute()

Attach a nic to specified instance.

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
    body := *openapiclient.NewCoreAttachNetworkInterfaceRequest() // CoreAttachNetworkInterfaceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreNetworkInterfaceServiceApi.AttachNetworkInterface(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreNetworkInterfaceServiceApi.AttachNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachNetworkInterface`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreNetworkInterfaceServiceApi.AttachNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreAttachNetworkInterfaceRequest**](CoreAttachNetworkInterfaceRequest.md) |  | 

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


## CreateNetworkInterface

> CommonActionResponse CreateNetworkInterface(ctx).Body(body).Execute()

Create a new nic.

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
    body := *openapiclient.NewCoreCreateNetworkInterfaceRequest() // CoreCreateNetworkInterfaceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreNetworkInterfaceServiceApi.CreateNetworkInterface(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreNetworkInterfaceServiceApi.CreateNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkInterface`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreNetworkInterfaceServiceApi.CreateNetworkInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateNetworkInterfaceRequest**](CoreCreateNetworkInterfaceRequest.md) |  | 

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


## DeleteNetworkInterface

> CommonActionResponse DeleteNetworkInterface(ctx, id).Execute()

release a nic.

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
    resp, r, err := api_client.CoreNetworkInterfaceServiceApi.DeleteNetworkInterface(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreNetworkInterfaceServiceApi.DeleteNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkInterface`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreNetworkInterfaceServiceApi.DeleteNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkInterfaceRequest struct via the builder pattern


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


## DescribeNetworkInterface

> CoreDescribeNetworkInterfaceResponse DescribeNetworkInterface(ctx, id).Execute()

Describe an eip detail.

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
    resp, r, err := api_client.CoreNetworkInterfaceServiceApi.DescribeNetworkInterface(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreNetworkInterfaceServiceApi.DescribeNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeNetworkInterface`: CoreDescribeNetworkInterfaceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreNetworkInterfaceServiceApi.DescribeNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeNetworkInterfaceResponse**](CoreDescribeNetworkInterfaceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachNetworkInterface

> CommonActionResponse DetachNetworkInterface(ctx, id).Execute()

Detach a nic from instance.

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
    resp, r, err := api_client.CoreNetworkInterfaceServiceApi.DetachNetworkInterface(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreNetworkInterfaceServiceApi.DetachNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachNetworkInterface`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreNetworkInterfaceServiceApi.DetachNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachNetworkInterfaceRequest struct via the builder pattern


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


## ListNetworkInterface

> CoreListNetworkInterfaceResponse ListNetworkInterface(ctx).Id(id).Name(name).Tenant(tenant).Vendor(vendor).Region(region).Zone(zone).Status(status).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all nics.

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
    name := []string{"Inner_example"} // []string |  (optional)
    tenant := []string{"Inner_example"} // []string |  (optional)
    vendor := []string{"Inner_example"} // []string |  (optional)
    region := []string{"Inner_example"} // []string |  (optional)
    zone := []string{"Inner_example"} // []string |  (optional)
    status := []string{"Inner_example"} // []string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreNetworkInterfaceServiceApi.ListNetworkInterface(context.Background()).Id(id).Name(name).Tenant(tenant).Vendor(vendor).Region(region).Zone(zone).Status(status).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreNetworkInterfaceServiceApi.ListNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkInterface`: CoreListNetworkInterfaceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreNetworkInterfaceServiceApi.ListNetworkInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **name** | **[]string** |  | 
 **tenant** | **[]string** |  | 
 **vendor** | **[]string** |  | 
 **region** | **[]string** |  | 
 **zone** | **[]string** |  | 
 **status** | **[]string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListNetworkInterfaceResponse**](CoreListNetworkInterfaceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

