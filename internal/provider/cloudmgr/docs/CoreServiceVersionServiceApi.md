# \CoreServiceVersionServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceVersion**](CoreServiceVersionServiceApi.md#CreateServiceVersion) | **Post** /default/api/v2/services/versions | Create a new service version.
[**DeleteServiceVersion**](CoreServiceVersionServiceApi.md#DeleteServiceVersion) | **Delete** /default/api/v2/services/versions/{id} | Delete a service version.
[**DescribeServiceVersion**](CoreServiceVersionServiceApi.md#DescribeServiceVersion) | **Get** /default/api/v2/services/versions/{id} | Describe the detail of a given service version.
[**ListServiceVersion**](CoreServiceVersionServiceApi.md#ListServiceVersion) | **Get** /default/api/v2/services/versions | List all service versions.
[**UpdateServiceVersion**](CoreServiceVersionServiceApi.md#UpdateServiceVersion) | **Put** /default/api/v2/services/versions/{id} | Update a service version.



## CreateServiceVersion

> CoreDescribeServiceVersionResponse CreateServiceVersion(ctx).Body(body).Execute()

Create a new service version.

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
    body := *openapiclient.NewCoreCreateServiceVersionRequest() // CoreCreateServiceVersionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceVersionServiceApi.CreateServiceVersion(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceVersionServiceApi.CreateServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceVersion`: CoreDescribeServiceVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceVersionServiceApi.CreateServiceVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateServiceVersionRequest**](CoreCreateServiceVersionRequest.md) |  | 

### Return type

[**CoreDescribeServiceVersionResponse**](CoreDescribeServiceVersionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceVersion

> CommonActionResponse DeleteServiceVersion(ctx, id).Execute()

Delete a service version.

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
    resp, r, err := api_client.CoreServiceVersionServiceApi.DeleteServiceVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceVersionServiceApi.DeleteServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteServiceVersion`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceVersionServiceApi.DeleteServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceVersionRequest struct via the builder pattern


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


## DescribeServiceVersion

> CoreDescribeServiceVersionResponse DescribeServiceVersion(ctx, id).Execute()

Describe the detail of a given service version.

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
    resp, r, err := api_client.CoreServiceVersionServiceApi.DescribeServiceVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceVersionServiceApi.DescribeServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeServiceVersion`: CoreDescribeServiceVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceVersionServiceApi.DescribeServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeServiceVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeServiceVersionResponse**](CoreDescribeServiceVersionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceVersion

> CoreListServiceVersionResponse ListServiceVersion(ctx).Id(id).Type_(type_).Region(region).Vendor(vendor).LowerThan(lowerThan).UpgradeFrom(upgradeFrom).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Arch(arch).Execute()

List all service versions.

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
    type_ := []string{"Inner_example"} // []string |  (optional)
    region := []string{"Inner_example"} // []string |  (optional)
    vendor := []string{"Inner_example"} // []string |  (optional)
    lowerThan := "lowerThan_example" // string |  (optional)
    upgradeFrom := "upgradeFrom_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    arch := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceVersionServiceApi.ListServiceVersion(context.Background()).Id(id).Type_(type_).Region(region).Vendor(vendor).LowerThan(lowerThan).UpgradeFrom(upgradeFrom).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Arch(arch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceVersionServiceApi.ListServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceVersion`: CoreListServiceVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceVersionServiceApi.ListServiceVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServiceVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **type_** | **[]string** |  | 
 **region** | **[]string** |  | 
 **vendor** | **[]string** |  | 
 **lowerThan** | **string** |  | 
 **upgradeFrom** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **arch** | **[]string** |  | 

### Return type

[**CoreListServiceVersionResponse**](CoreListServiceVersionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceVersion

> CoreDescribeServiceVersionResponse UpdateServiceVersion(ctx, id).Body(body).Execute()

Update a service version.

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
    body := *openapiclient.NewCoreUpdateServiceVersionRequest() // CoreUpdateServiceVersionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceVersionServiceApi.UpdateServiceVersion(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceVersionServiceApi.UpdateServiceVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceVersion`: CoreDescribeServiceVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceVersionServiceApi.UpdateServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreUpdateServiceVersionRequest**](CoreUpdateServiceVersionRequest.md) |  | 

### Return type

[**CoreDescribeServiceVersionResponse**](CoreDescribeServiceVersionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

