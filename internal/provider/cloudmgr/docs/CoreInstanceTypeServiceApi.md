# \CoreInstanceTypeServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstanceType**](CoreInstanceTypeServiceApi.md#CreateInstanceType) | **Post** /default/api/v2/instances/types | Create an instance type.
[**DeleteInstanceType**](CoreInstanceTypeServiceApi.md#DeleteInstanceType) | **Delete** /default/api/v2/instances/types/{name} | Delete an instance type.
[**DescribeInstanceType**](CoreInstanceTypeServiceApi.md#DescribeInstanceType) | **Get** /default/api/v2/instances/types/{name} | Describe the detail of a given instance type.
[**DescribeInstanceTypeByComponentType**](CoreInstanceTypeServiceApi.md#DescribeInstanceTypeByComponentType) | **Get** /default/api/v2/instances/types/component | Describe the detail of a given instance type.
[**ListInstanceType**](CoreInstanceTypeServiceApi.md#ListInstanceType) | **Get** /default/api/v2/instances/types | List all instance type.
[**UpdateInstanceType**](CoreInstanceTypeServiceApi.md#UpdateInstanceType) | **Put** /default/api/v2/instances/types/{name} | Update an instance type.



## CreateInstanceType

> CoreDescribeInstanceTypeResponse CreateInstanceType(ctx).Body(body).Execute()

Create an instance type.

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
    body := *openapiclient.NewCoreCreateInstanceTypeRequest() // CoreCreateInstanceTypeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreInstanceTypeServiceApi.CreateInstanceType(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceTypeServiceApi.CreateInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInstanceType`: CoreDescribeInstanceTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceTypeServiceApi.CreateInstanceType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateInstanceTypeRequest**](CoreCreateInstanceTypeRequest.md) |  | 

### Return type

[**CoreDescribeInstanceTypeResponse**](CoreDescribeInstanceTypeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstanceType

> CommonActionResponse DeleteInstanceType(ctx, name).Execute()

Delete an instance type.

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
    resp, r, err := api_client.CoreInstanceTypeServiceApi.DeleteInstanceType(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceTypeServiceApi.DeleteInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInstanceType`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceTypeServiceApi.DeleteInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceTypeRequest struct via the builder pattern


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


## DescribeInstanceType

> CoreDescribeInstanceTypeResponse DescribeInstanceType(ctx, name).Execute()

Describe the detail of a given instance type.

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
    resp, r, err := api_client.CoreInstanceTypeServiceApi.DescribeInstanceType(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceTypeServiceApi.DescribeInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeInstanceType`: CoreDescribeInstanceTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceTypeServiceApi.DescribeInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeInstanceTypeResponse**](CoreDescribeInstanceTypeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeInstanceTypeByComponentType

> CoreDescribeInstanceTypeResponse DescribeInstanceTypeByComponentType(ctx).Service(service).ComponentType(componentType).Execute()

Describe the detail of a given instance type.

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
    service := "service_example" // string |  (optional)
    componentType := "componentType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreInstanceTypeServiceApi.DescribeInstanceTypeByComponentType(context.Background()).Service(service).ComponentType(componentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceTypeServiceApi.DescribeInstanceTypeByComponentType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeInstanceTypeByComponentType`: CoreDescribeInstanceTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceTypeServiceApi.DescribeInstanceTypeByComponentType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeInstanceTypeByComponentTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** |  | 
 **componentType** | **string** |  | 

### Return type

[**CoreDescribeInstanceTypeResponse**](CoreDescribeInstanceTypeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstanceType

> CoreListInstanceTypeResponse ListInstanceType(ctx).Zone(zone).Vendor(vendor).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Arch(arch).Execute()

List all instance type.

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
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    arch := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreInstanceTypeServiceApi.ListInstanceType(context.Background()).Zone(zone).Vendor(vendor).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Arch(arch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceTypeServiceApi.ListInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstanceType`: CoreListInstanceTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceTypeServiceApi.ListInstanceType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **[]string** |  | 
 **vendor** | **[]string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **arch** | **[]string** |  | 

### Return type

[**CoreListInstanceTypeResponse**](CoreListInstanceTypeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstanceType

> CoreDescribeInstanceTypeResponse UpdateInstanceType(ctx, name).Body(body).Execute()

Update an instance type.

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
    body := *openapiclient.NewCoreUpdateInstanceTypeRequest() // CoreUpdateInstanceTypeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreInstanceTypeServiceApi.UpdateInstanceType(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceTypeServiceApi.UpdateInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInstanceType`: CoreDescribeInstanceTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceTypeServiceApi.UpdateInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreUpdateInstanceTypeRequest**](CoreUpdateInstanceTypeRequest.md) |  | 

### Return type

[**CoreDescribeInstanceTypeResponse**](CoreDescribeInstanceTypeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

