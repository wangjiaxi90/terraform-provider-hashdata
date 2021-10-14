# \CoreResourcePoolServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInstanceToResourcePool**](CoreResourcePoolServiceApi.md#AddInstanceToResourcePool) | **Post** /default/api/v2/resources/{id}/instances | Add instances to a resource pool
[**CreateResourcePool**](CoreResourcePoolServiceApi.md#CreateResourcePool) | **Post** /default/api/v2/resources | Create a new data center zone.
[**DeleteInstanceFromResourcePool**](CoreResourcePoolServiceApi.md#DeleteInstanceFromResourcePool) | **Delete** /default/api/v2/resources/{id}/instances/{instance} | Remove a instance from a resource pool.
[**DeleteResourcePool**](CoreResourcePoolServiceApi.md#DeleteResourcePool) | **Delete** /default/api/v2/resources/{id} | Delete a data center zone.
[**DescribeResourcePool**](CoreResourcePoolServiceApi.md#DescribeResourcePool) | **Get** /default/api/v2/resources/{id} | Describe the detail of a given resource pool.
[**ListPoolInstance**](CoreResourcePoolServiceApi.md#ListPoolInstance) | **Get** /default/api/v2/resources/{id}/instances | List all instances in resource pool.
[**ListResourcePool**](CoreResourcePoolServiceApi.md#ListResourcePool) | **Get** /default/api/v2/resources | List all resource pool.



## AddInstanceToResourcePool

> CommonDescribeJobResponse AddInstanceToResourcePool(ctx, id).Body(body).Execute()

Add instances to a resource pool

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
    body := *openapiclient.NewCoreAddInstanceToResourcePoolRequest() // CoreAddInstanceToResourcePoolRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreResourcePoolServiceApi.AddInstanceToResourcePool(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreResourcePoolServiceApi.AddInstanceToResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInstanceToResourcePool`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreResourcePoolServiceApi.AddInstanceToResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddInstanceToResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreAddInstanceToResourcePoolRequest**](CoreAddInstanceToResourcePoolRequest.md) |  | 

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


## CreateResourcePool

> CoreDescribeResourcePoolResponse CreateResourcePool(ctx).Body(body).Execute()

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
    body := *openapiclient.NewCoreCreateResourcePoolRequest() // CoreCreateResourcePoolRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreResourcePoolServiceApi.CreateResourcePool(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreResourcePoolServiceApi.CreateResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResourcePool`: CoreDescribeResourcePoolResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreResourcePoolServiceApi.CreateResourcePool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateResourcePoolRequest**](CoreCreateResourcePoolRequest.md) |  | 

### Return type

[**CoreDescribeResourcePoolResponse**](CoreDescribeResourcePoolResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstanceFromResourcePool

> CommonActionResponse DeleteInstanceFromResourcePool(ctx, id, instance).Execute()

Remove a instance from a resource pool.

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
    instance := "instance_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreResourcePoolServiceApi.DeleteInstanceFromResourcePool(context.Background(), id, instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreResourcePoolServiceApi.DeleteInstanceFromResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInstanceFromResourcePool`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreResourcePoolServiceApi.DeleteInstanceFromResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**instance** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceFromResourcePoolRequest struct via the builder pattern


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


## DeleteResourcePool

> CommonActionResponse DeleteResourcePool(ctx, id).Execute()

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreResourcePoolServiceApi.DeleteResourcePool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreResourcePoolServiceApi.DeleteResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteResourcePool`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreResourcePoolServiceApi.DeleteResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourcePoolRequest struct via the builder pattern


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


## DescribeResourcePool

> CoreDescribeResourcePoolResponse DescribeResourcePool(ctx, id).Execute()

Describe the detail of a given resource pool.

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
    resp, r, err := api_client.CoreResourcePoolServiceApi.DescribeResourcePool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreResourcePoolServiceApi.DescribeResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeResourcePool`: CoreDescribeResourcePoolResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreResourcePoolServiceApi.DescribeResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeResourcePoolResponse**](CoreDescribeResourcePoolResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPoolInstance

> CoreListInstanceResponse ListPoolInstance(ctx, id).Instance(instance).Status(status).Component(component).Offset(offset).Limit(limit).Execute()

List all instances in resource pool.

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
    instance := []string{"Inner_example"} // []string |  (optional)
    status := []string{"Inner_example"} // []string |  (optional)
    component := []string{"Inner_example"} // []string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreResourcePoolServiceApi.ListPoolInstance(context.Background(), id).Instance(instance).Status(status).Component(component).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreResourcePoolServiceApi.ListPoolInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPoolInstance`: CoreListInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreResourcePoolServiceApi.ListPoolInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPoolInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | **[]string** |  | 
 **status** | **[]string** |  | 
 **component** | **[]string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListInstanceResponse**](CoreListInstanceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourcePool

> CoreListResourcePoolResponse ListResourcePool(ctx).Id(id).Tenant(tenant).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all resource pool.

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
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreResourcePoolServiceApi.ListResourcePool(context.Background()).Id(id).Tenant(tenant).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreResourcePoolServiceApi.ListResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListResourcePool`: CoreListResourcePoolResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreResourcePoolServiceApi.ListResourcePool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **tenant** | **[]string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListResourcePoolResponse**](CoreListResourcePoolResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

