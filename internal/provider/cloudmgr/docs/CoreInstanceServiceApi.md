# \CoreInstanceServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DescribeInstance**](CoreInstanceServiceApi.md#DescribeInstance) | **Get** /default/api/v2/instances/{id} | Describe the detail of a given instance.
[**ListInstance**](CoreInstanceServiceApi.md#ListInstance) | **Get** /default/api/v2/instances | List all instance.
[**ListVolume**](CoreInstanceServiceApi.md#ListVolume) | **Get** /default/api/v2/instances/{instance}/volumes | List all volumes of of a given instance.



## DescribeInstance

> CoreDescribeInstanceResponse DescribeInstance(ctx, id).Execute()

Describe the detail of a given instance.

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
    resp, r, err := api_client.CoreInstanceServiceApi.DescribeInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceServiceApi.DescribeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeInstance`: CoreDescribeInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceServiceApi.DescribeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeInstanceResponse**](CoreDescribeInstanceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstance

> CoreListInstanceResponse ListInstance(ctx).Id(id).Owner(owner).Name(name).Tenant(tenant).Vendor(vendor).Zone(zone).Status(status).Offset(offset).Limit(limit).Execute()

List all instance.

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
    owner := []string{"Inner_example"} // []string |  (optional)
    name := []string{"Inner_example"} // []string |  (optional)
    tenant := []string{"Inner_example"} // []string |  (optional)
    vendor := []string{"Inner_example"} // []string |  (optional)
    zone := []string{"Inner_example"} // []string |  (optional)
    status := []string{"Inner_example"} // []string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreInstanceServiceApi.ListInstance(context.Background()).Id(id).Owner(owner).Name(name).Tenant(tenant).Vendor(vendor).Zone(zone).Status(status).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceServiceApi.ListInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstance`: CoreListInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceServiceApi.ListInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **owner** | **[]string** |  | 
 **name** | **[]string** |  | 
 **tenant** | **[]string** |  | 
 **vendor** | **[]string** |  | 
 **zone** | **[]string** |  | 
 **status** | **[]string** |  | 
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


## ListVolume

> CoreListVolumeResponse ListVolume(ctx, instance).Id(id).Name(name).Offset(offset).Limit(limit).Execute()

List all volumes of of a given instance.

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
    instance := "instance_example" // string | 
    id := []string{"Inner_example"} // []string |  (optional)
    name := []string{"Inner_example"} // []string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreInstanceServiceApi.ListVolume(context.Background(), instance).Id(id).Name(name).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreInstanceServiceApi.ListVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVolume`: CoreListVolumeResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreInstanceServiceApi.ListVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instance** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **[]string** |  | 
 **name** | **[]string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListVolumeResponse**](CoreListVolumeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

