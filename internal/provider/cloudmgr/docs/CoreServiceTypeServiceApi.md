# \CoreServiceTypeServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListServiceType**](CoreServiceTypeServiceApi.md#ListServiceType) | **Get** /default/api/v2/services/types | List all type of service.



## ListServiceType

> CoreListServiceTypeResponse ListServiceType(ctx).Execute()

List all type of service.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceTypeServiceApi.ListServiceType(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceTypeServiceApi.ListServiceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceType`: CoreListServiceTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceTypeServiceApi.ListServiceType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceTypeRequest struct via the builder pattern


### Return type

[**CoreListServiceTypeResponse**](CoreListServiceTypeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

