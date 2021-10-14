# \CoreVpcServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListVpc**](CoreVpcServiceApi.md#ListVpc) | **Get** /default/api/v2/vpcs | List all available vpc.
[**ListVpcSubnets**](CoreVpcServiceApi.md#ListVpcSubnets) | **Get** /default/api/v2/vpcs/{vpc} | List all available subnet in a given vpc.



## ListVpc

> CoreListVpcResponse ListVpc(ctx).Zone(zone).AccessKey(accessKey).AccessKeySecret(accessKeySecret).Execute()

List all available vpc.

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
    zone := "zone_example" // string |  (optional)
    accessKey := "accessKey_example" // string |  (optional)
    accessKeySecret := "accessKeySecret_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreVpcServiceApi.ListVpc(context.Background()).Zone(zone).AccessKey(accessKey).AccessKeySecret(accessKeySecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreVpcServiceApi.ListVpc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVpc`: CoreListVpcResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreVpcServiceApi.ListVpc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string** |  | 
 **accessKey** | **string** |  | 
 **accessKeySecret** | **string** |  | 

### Return type

[**CoreListVpcResponse**](CoreListVpcResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVpcSubnets

> CoreListVpcSubnetsResponse ListVpcSubnets(ctx, vpc).Zone(zone).AccessKey(accessKey).AccessKeySecret(accessKeySecret).Execute()

List all available subnet in a given vpc.

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
    vpc := "vpc_example" // string | 
    zone := "zone_example" // string |  (optional)
    accessKey := "accessKey_example" // string |  (optional)
    accessKeySecret := "accessKeySecret_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreVpcServiceApi.ListVpcSubnets(context.Background(), vpc).Zone(zone).AccessKey(accessKey).AccessKeySecret(accessKeySecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreVpcServiceApi.ListVpcSubnets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVpcSubnets`: CoreListVpcSubnetsResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreVpcServiceApi.ListVpcSubnets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpc** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVpcSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zone** | **string** |  | 
 **accessKey** | **string** |  | 
 **accessKeySecret** | **string** |  | 

### Return type

[**CoreListVpcSubnetsResponse**](CoreListVpcSubnetsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

