# \CoreVendorServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListVendor**](CoreVendorServiceApi.md#ListVendor) | **Get** /default/api/v2/vendors | List all cloud vendors.



## ListVendor

> CoreDescribeVendorListResponse ListVendor(ctx).Locale(locale).Execute()

List all cloud vendors.

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
    locale := "locale_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreVendorServiceApi.ListVendor(context.Background()).Locale(locale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreVendorServiceApi.ListVendor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVendor`: CoreDescribeVendorListResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreVendorServiceApi.ListVendor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVendorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** |  | 

### Return type

[**CoreDescribeVendorListResponse**](CoreDescribeVendorListResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

