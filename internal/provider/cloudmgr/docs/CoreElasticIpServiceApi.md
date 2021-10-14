# \CoreElasticIpServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachElasticIp**](CoreElasticIpServiceApi.md#AttachElasticIp) | **Post** /default/api/v2/network/eips/{id}/attach | Attach an eip to specified instance.
[**CreateElasticIp**](CoreElasticIpServiceApi.md#CreateElasticIp) | **Post** /default/api/v2/network/eips | Create a new eip.
[**DeleteElasticIp**](CoreElasticIpServiceApi.md#DeleteElasticIp) | **Delete** /default/api/v2/network/eips/{id} | release an eip.
[**DescribeElasticIp**](CoreElasticIpServiceApi.md#DescribeElasticIp) | **Get** /default/api/v2/network/eips/{id} | Describe an eip detail.
[**DetachElasticIp**](CoreElasticIpServiceApi.md#DetachElasticIp) | **Get** /default/api/v2/network/eips/{id}/detach | Detach an eip from instance.
[**ListElasticIp**](CoreElasticIpServiceApi.md#ListElasticIp) | **Get** /default/api/v2/network/eips | List all eips.
[**ModifyElasticIpMaxBandWidth**](CoreElasticIpServiceApi.md#ModifyElasticIpMaxBandWidth) | **Put** /default/api/v2/network/eips/{id} | Modify maxBandWidth for specified eip.



## AttachElasticIp

> CommonActionResponse AttachElasticIp(ctx, id).Body(body).Execute()

Attach an eip to specified instance.

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
    body := *openapiclient.NewCoreAttachElasticIpRequest() // CoreAttachElasticIpRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreElasticIpServiceApi.AttachElasticIp(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreElasticIpServiceApi.AttachElasticIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachElasticIp`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreElasticIpServiceApi.AttachElasticIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachElasticIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreAttachElasticIpRequest**](CoreAttachElasticIpRequest.md) |  | 

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


## CreateElasticIp

> CommonActionResponse CreateElasticIp(ctx).Body(body).Execute()

Create a new eip.

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
    body := *openapiclient.NewCoreCreateElasticIpRequest() // CoreCreateElasticIpRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreElasticIpServiceApi.CreateElasticIp(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreElasticIpServiceApi.CreateElasticIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateElasticIp`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreElasticIpServiceApi.CreateElasticIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateElasticIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateElasticIpRequest**](CoreCreateElasticIpRequest.md) |  | 

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


## DeleteElasticIp

> CommonActionResponse DeleteElasticIp(ctx, id).Execute()

release an eip.

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
    resp, r, err := api_client.CoreElasticIpServiceApi.DeleteElasticIp(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreElasticIpServiceApi.DeleteElasticIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteElasticIp`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreElasticIpServiceApi.DeleteElasticIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteElasticIpRequest struct via the builder pattern


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


## DescribeElasticIp

> CoreDescribeElasticIpResponse DescribeElasticIp(ctx, id).Execute()

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
    resp, r, err := api_client.CoreElasticIpServiceApi.DescribeElasticIp(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreElasticIpServiceApi.DescribeElasticIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeElasticIp`: CoreDescribeElasticIpResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreElasticIpServiceApi.DescribeElasticIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeElasticIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeElasticIpResponse**](CoreDescribeElasticIpResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachElasticIp

> CommonActionResponse DetachElasticIp(ctx, id).Execute()

Detach an eip from instance.

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
    resp, r, err := api_client.CoreElasticIpServiceApi.DetachElasticIp(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreElasticIpServiceApi.DetachElasticIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachElasticIp`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreElasticIpServiceApi.DetachElasticIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachElasticIpRequest struct via the builder pattern


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


## ListElasticIp

> CoreListElasticIpResponse ListElasticIp(ctx).Id(id).Name(name).Region(region).Vendor(vendor).Tenant(tenant).Status(status).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all eips.

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
    region := []string{"Inner_example"} // []string |  (optional)
    vendor := []string{"Inner_example"} // []string |  (optional)
    tenant := []string{"Inner_example"} // []string |  (optional)
    status := []string{"Inner_example"} // []string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreElasticIpServiceApi.ListElasticIp(context.Background()).Id(id).Name(name).Region(region).Vendor(vendor).Tenant(tenant).Status(status).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreElasticIpServiceApi.ListElasticIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListElasticIp`: CoreListElasticIpResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreElasticIpServiceApi.ListElasticIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListElasticIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **name** | **[]string** |  | 
 **region** | **[]string** |  | 
 **vendor** | **[]string** |  | 
 **tenant** | **[]string** |  | 
 **status** | **[]string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListElasticIpResponse**](CoreListElasticIpResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyElasticIpMaxBandWidth

> CommonActionResponse ModifyElasticIpMaxBandWidth(ctx, id).Body(body).Execute()

Modify maxBandWidth for specified eip.

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
    body := *openapiclient.NewCoreModifyElasticIpMaxBandWidthRequest() // CoreModifyElasticIpMaxBandWidthRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreElasticIpServiceApi.ModifyElasticIpMaxBandWidth(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreElasticIpServiceApi.ModifyElasticIpMaxBandWidth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyElasticIpMaxBandWidth`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreElasticIpServiceApi.ModifyElasticIpMaxBandWidth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyElasticIpMaxBandWidthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreModifyElasticIpMaxBandWidthRequest**](CoreModifyElasticIpMaxBandWidthRequest.md) |  | 

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

