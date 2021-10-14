# \CoreAlertServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertRule**](CoreAlertServiceApi.md#CreateAlertRule) | **Post** /default/api/v2/services/{service}/alerts | Create a service alert rule.
[**DeleteAlertRule**](CoreAlertServiceApi.md#DeleteAlertRule) | **Delete** /default/api/v2/service/{service}/alerts/{alert} | Delete an alert rule of a service.
[**DeleteAlertRuleByService**](CoreAlertServiceApi.md#DeleteAlertRuleByService) | **Delete** /default/api/v2/service/{service}/alerts | Delete all alert rules of a service.
[**DescribeAlertRule**](CoreAlertServiceApi.md#DescribeAlertRule) | **Get** /default/api/v2/services/{service}/alerts/{alert} | Describe the detail of a given alert rule.
[**ListAlertRule**](CoreAlertServiceApi.md#ListAlertRule) | **Get** /default/api/v2/services/{service}/alerts | List the alerts of a given service.



## CreateAlertRule

> VmalertDescribeAlertRuleResponse CreateAlertRule(ctx, service).Body(body).Execute()

Create a service alert rule.

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
    service := "service_example" // string | 
    body := *openapiclient.NewVmalertCreateAlertRuleRequest() // VmalertCreateAlertRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreAlertServiceApi.CreateAlertRule(context.Background(), service).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreAlertServiceApi.CreateAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertRule`: VmalertDescribeAlertRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreAlertServiceApi.CreateAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VmalertCreateAlertRuleRequest**](VmalertCreateAlertRuleRequest.md) |  | 

### Return type

[**VmalertDescribeAlertRuleResponse**](VmalertDescribeAlertRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertRule

> CommonActionResponse DeleteAlertRule(ctx, service, alert).Execute()

Delete an alert rule of a service.

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
    service := "service_example" // string | 
    alert := "alert_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreAlertServiceApi.DeleteAlertRule(context.Background(), service, alert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreAlertServiceApi.DeleteAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlertRule`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreAlertServiceApi.DeleteAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**alert** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertRuleRequest struct via the builder pattern


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


## DeleteAlertRuleByService

> CommonActionResponse DeleteAlertRuleByService(ctx, service).Execute()

Delete all alert rules of a service.

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
    service := "service_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreAlertServiceApi.DeleteAlertRuleByService(context.Background(), service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreAlertServiceApi.DeleteAlertRuleByService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlertRuleByService`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreAlertServiceApi.DeleteAlertRuleByService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertRuleByServiceRequest struct via the builder pattern


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


## DescribeAlertRule

> VmalertDescribeAlertRuleResponse DescribeAlertRule(ctx, service, alert).Execute()

Describe the detail of a given alert rule.

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
    service := "service_example" // string | 
    alert := "alert_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreAlertServiceApi.DescribeAlertRule(context.Background(), service, alert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreAlertServiceApi.DescribeAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeAlertRule`: VmalertDescribeAlertRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreAlertServiceApi.DescribeAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**alert** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VmalertDescribeAlertRuleResponse**](VmalertDescribeAlertRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertRule

> VmalertListAlertRuleResponse ListAlertRule(ctx, service).Offset(offset).Limit(limit).Execute()

List the alerts of a given service.

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
    service := "service_example" // string | 
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreAlertServiceApi.ListAlertRule(context.Background(), service).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreAlertServiceApi.ListAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlertRule`: VmalertListAlertRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreAlertServiceApi.ListAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**VmalertListAlertRuleResponse**](VmalertListAlertRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

