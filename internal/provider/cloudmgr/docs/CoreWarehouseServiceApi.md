# \CoreWarehouseServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCatalog**](CoreWarehouseServiceApi.md#CreateCatalog) | **Post** /default/api/v2/services/catalog | Create a catalog service.
[**CreateWarehouse**](CoreWarehouseServiceApi.md#CreateWarehouse) | **Post** /default/api/v2/services/warehouse | Create a database service.



## CreateCatalog

> CommonDescribeJobResponse CreateCatalog(ctx).Body(body).Execute()

Create a catalog service.

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
    body := *openapiclient.NewCoreCreateCatalogRequest() // CoreCreateCatalogRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreWarehouseServiceApi.CreateCatalog(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreWarehouseServiceApi.CreateCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCatalog`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreWarehouseServiceApi.CreateCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateCatalogRequest**](CoreCreateCatalogRequest.md) |  | 

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


## CreateWarehouse

> CommonDescribeJobResponse CreateWarehouse(ctx).Body(body).Execute()

Create a database service.

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
    body := *openapiclient.NewCoreCreateWarehouseRequest() // CoreCreateWarehouseRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreWarehouseServiceApi.CreateWarehouse(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreWarehouseServiceApi.CreateWarehouse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWarehouse`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreWarehouseServiceApi.CreateWarehouse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateWarehouseRequest**](CoreCreateWarehouseRequest.md) |  | 

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

