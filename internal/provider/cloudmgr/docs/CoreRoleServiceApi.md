# \CoreRoleServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DescribeRole**](CoreRoleServiceApi.md#DescribeRole) | **Get** /default/api/v2/roles/{name} | Describe the detail of a given role.
[**ListRole**](CoreRoleServiceApi.md#ListRole) | **Get** /default/api/v2/roles | List all roles.



## DescribeRole

> CoreDescribeRoleResponse DescribeRole(ctx, name).Execute()

Describe the detail of a given role.

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
    resp, r, err := api_client.CoreRoleServiceApi.DescribeRole(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreRoleServiceApi.DescribeRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeRole`: CoreDescribeRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreRoleServiceApi.DescribeRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeRoleResponse**](CoreDescribeRoleResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRole

> CoreListRoleResponse ListRole(ctx).Name(name).Offset(offset).Limit(limit).Execute()

List all roles.

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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreRoleServiceApi.ListRole(context.Background()).Name(name).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreRoleServiceApi.ListRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRole`: CoreListRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreRoleServiceApi.ListRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **[]string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListRoleResponse**](CoreListRoleResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

