# \CoreBucketServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DescribeBucket**](CoreBucketServiceApi.md#DescribeBucket) | **Get** /default/api/v2/buckets/{id} | Describe the detail of a given bucket.
[**ListObject**](CoreBucketServiceApi.md#ListObject) | **Get** /default/api/v2/buckets/{bucket}/contents | List the contents of a given bucket.



## DescribeBucket

> CoreDescribeBucketResponse DescribeBucket(ctx, id).Execute()

Describe the detail of a given bucket.

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
    resp, r, err := api_client.CoreBucketServiceApi.DescribeBucket(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreBucketServiceApi.DescribeBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeBucket`: CoreDescribeBucketResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreBucketServiceApi.DescribeBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeBucketResponse**](CoreDescribeBucketResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObject

> CoreListObjectResponse ListObject(ctx, bucket).Prefix(prefix).Marker(marker).Limit(limit).Execute()

List the contents of a given bucket.

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
    bucket := "bucket_example" // string | 
    prefix := "prefix_example" // string |  (optional)
    marker := "marker_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreBucketServiceApi.ListObject(context.Background(), bucket).Prefix(prefix).Marker(marker).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreBucketServiceApi.ListObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListObject`: CoreListObjectResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreBucketServiceApi.ListObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **string** |  | 
 **marker** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListObjectResponse**](CoreListObjectResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

