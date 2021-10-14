# \CoreQuotaServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuota**](CoreQuotaServiceApi.md#CreateQuota) | **Post** /default/api/v2/quotas | Create default quota configure for given tenant.
[**ListQuota**](CoreQuotaServiceApi.md#ListQuota) | **Get** /default/api/v2/quotas | List all configured quota.
[**UpdateQuota**](CoreQuotaServiceApi.md#UpdateQuota) | **Put** /default/api/v2/quotas | Adjust quota for given tenant.



## CreateQuota

> CommonActionResponse CreateQuota(ctx).Body(body).Execute()

Create default quota configure for given tenant.

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
    body := *openapiclient.NewCoreCreateQuotaRequest() // CoreCreateQuotaRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreQuotaServiceApi.CreateQuota(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreQuotaServiceApi.CreateQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQuota`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreQuotaServiceApi.CreateQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreCreateQuotaRequest**](CoreCreateQuotaRequest.md) |  | 

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


## ListQuota

> CoreListQuotaResponse ListQuota(ctx).Tenant(tenant).Zone(zone).Offset(offset).Limit(limit).Region(region).Execute()

List all configured quota.

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
    tenant := "tenant_example" // string |  (optional)
    zone := []string{"Inner_example"} // []string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    region := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreQuotaServiceApi.ListQuota(context.Background()).Tenant(tenant).Zone(zone).Offset(offset).Limit(limit).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreQuotaServiceApi.ListQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQuota`: CoreListQuotaResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreQuotaServiceApi.ListQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant** | **string** |  | 
 **zone** | **[]string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **region** | **[]string** |  | 

### Return type

[**CoreListQuotaResponse**](CoreListQuotaResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuota

> CoreDescribeQuotaResponse UpdateQuota(ctx).Body(body).Execute()

Adjust quota for given tenant.

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
    body := *openapiclient.NewCoreUpdateQuotaRequest() // CoreUpdateQuotaRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreQuotaServiceApi.UpdateQuota(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreQuotaServiceApi.UpdateQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQuota`: CoreDescribeQuotaResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreQuotaServiceApi.UpdateQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreUpdateQuotaRequest**](CoreUpdateQuotaRequest.md) |  | 

### Return type

[**CoreDescribeQuotaResponse**](CoreDescribeQuotaResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

