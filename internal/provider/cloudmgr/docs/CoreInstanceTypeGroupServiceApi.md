# \CoreInstanceTypeGroupServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstanceTypeGroup**](CoreInstanceTypeGroupServiceApi.md#CreateInstanceTypeGroup) | **Post** /default/api/v2/instances/types/groups | Create an instance type group.
[**DeleteInstanceTypeGroup**](CoreInstanceTypeGroupServiceApi.md#DeleteInstanceTypeGroup) | **Delete** /default/api/v2/instances/types/groups/{name} | Delete an instance type group.
[**DescribeInstanceTypeGroup**](CoreInstanceTypeGroupServiceApi.md#DescribeInstanceTypeGroup) | **Get** /default/api/v2/instances/types/groups/{name} | Describe the detail of a given instance type group.
[**ListInstanceTypeGroup**](CoreInstanceTypeGroupServiceApi.md#ListInstanceTypeGroup) | **Get** /default/api/v2/instances/types/groups | List all instance type group.
[**UpdateInstanceTypeGroup**](CoreInstanceTypeGroupServiceApi.md#UpdateInstanceTypeGroup) | **Put** /default/api/v2/instances/types/groups/{name} | Update an instance type group.



## CreateInstanceTypeGroup

> CoreDescribeInstanceTypeGroupResponse CreateInstanceTypeGroup(ctx).Body(body).Execute()

Create an instance type group.

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
    body := *openapiclient.NewCoreCreateInstanceTypeGroupRequest() // CoreCreateInstanceTypeGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreInstanceTypeGroupServiceApi.CreateInstanceTypeGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceTypeGroupServiceApi.CreateInstanceTypeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInstanceTypeGroup`: CoreDescribeInstanceTypeGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceTypeGroupServiceApi.CreateInstanceTypeGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceTypeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateInstanceTypeGroupRequest**](CoreCreateInstanceTypeGroupRequest.md) |  | 

### Return type

[**CoreDescribeInstanceTypeGroupResponse**](CoreDescribeInstanceTypeGroupResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstanceTypeGroup

> CommonActionResponse DeleteInstanceTypeGroup(ctx, name).Execute()

Delete an instance type group.

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
    resp, r, err := api_client.CoreInstanceTypeGroupServiceApi.DeleteInstanceTypeGroup(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceTypeGroupServiceApi.DeleteInstanceTypeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInstanceTypeGroup`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceTypeGroupServiceApi.DeleteInstanceTypeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceTypeGroupRequest struct via the builder pattern


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


## DescribeInstanceTypeGroup

> CoreDescribeInstanceTypeGroupResponse DescribeInstanceTypeGroup(ctx, name).Execute()

Describe the detail of a given instance type group.

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
    resp, r, err := api_client.CoreInstanceTypeGroupServiceApi.DescribeInstanceTypeGroup(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceTypeGroupServiceApi.DescribeInstanceTypeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeInstanceTypeGroup`: CoreDescribeInstanceTypeGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceTypeGroupServiceApi.DescribeInstanceTypeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeInstanceTypeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeInstanceTypeGroupResponse**](CoreDescribeInstanceTypeGroupResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstanceTypeGroup

> CoreListInstanceTypeGroupResponse ListInstanceTypeGroup(ctx).Vendor(vendor).Offset(offset).Limit(limit).Arches(arches).Execute()

List all instance type group.

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
    vendor := []string{"Inner_example"} // []string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    arches := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreInstanceTypeGroupServiceApi.ListInstanceTypeGroup(context.Background()).Vendor(vendor).Offset(offset).Limit(limit).Arches(arches).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceTypeGroupServiceApi.ListInstanceTypeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstanceTypeGroup`: CoreListInstanceTypeGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceTypeGroupServiceApi.ListInstanceTypeGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceTypeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vendor** | **[]string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **arches** | **[]string** |  | 

### Return type

[**CoreListInstanceTypeGroupResponse**](CoreListInstanceTypeGroupResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstanceTypeGroup

> CoreDescribeInstanceTypeGroupResponse UpdateInstanceTypeGroup(ctx, name).Body(body).Execute()

Update an instance type group.

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
    body := *openapiclient.NewCoreUpdateInstanceTypeGroupRequest() // CoreUpdateInstanceTypeGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreInstanceTypeGroupServiceApi.UpdateInstanceTypeGroup(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceTypeGroupServiceApi.UpdateInstanceTypeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInstanceTypeGroup`: CoreDescribeInstanceTypeGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceTypeGroupServiceApi.UpdateInstanceTypeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceTypeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreUpdateInstanceTypeGroupRequest**](CoreUpdateInstanceTypeGroupRequest.md) |  | 

### Return type

[**CoreDescribeInstanceTypeGroupResponse**](CoreDescribeInstanceTypeGroupResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

