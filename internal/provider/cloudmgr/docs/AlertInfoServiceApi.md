# \AlertInfoServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmAlertInfo**](AlertInfoServiceApi.md#ConfirmAlertInfo) | **Put** /ops/api/v2/alertinfo/confirm/{id} | confirm a alert info.
[**CreateAlertInfo**](AlertInfoServiceApi.md#CreateAlertInfo) | **Post** /ops/api/v2/alertinfo/create | create a alert info.
[**DescribeAlertInfo**](AlertInfoServiceApi.md#DescribeAlertInfo) | **Get** /ops/api/v2/alertinfo/{id} | get a alert info.
[**ListAlertInfo**](AlertInfoServiceApi.md#ListAlertInfo) | **Get** /ops/api/v2/alertinfos | list alert infos.
[**ProcessAlertInfo**](AlertInfoServiceApi.md#ProcessAlertInfo) | **Put** /ops/api/v2/alertinfo/process/{id} | processed a alert info.



## ConfirmAlertInfo

> CommonActionResponse ConfirmAlertInfo(ctx, id).Execute()

confirm a alert info.

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
    resp, r, err := api_client.AlertInfoServiceApi.ConfirmAlertInfo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoServiceApi.ConfirmAlertInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmAlertInfo`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertInfoServiceApi.ConfirmAlertInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmAlertInfoRequest struct via the builder pattern


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


## CreateAlertInfo

> map[string]interface{} CreateAlertInfo(ctx).Execute()

create a alert info.

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
    resp, r, err := api_client.AlertInfoServiceApi.CreateAlertInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoServiceApi.CreateAlertInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AlertInfoServiceApi.CreateAlertInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertInfoRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeAlertInfo

> OpsDescribeAlertInfoResponse DescribeAlertInfo(ctx, id).Execute()

get a alert info.

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
    resp, r, err := api_client.AlertInfoServiceApi.DescribeAlertInfo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoServiceApi.DescribeAlertInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeAlertInfo`: OpsDescribeAlertInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertInfoServiceApi.DescribeAlertInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeAlertInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsDescribeAlertInfoResponse**](OpsDescribeAlertInfoResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertInfo

> OpsListAlertInfoResponse ListAlertInfo(ctx).Name(name).Value(value).Processor(processor).AlertStatus(alertStatus).AlertType(alertType).StartAt(startAt).EndAt(endAt).Offset(offset).Limit(limit).Execute()

list alert infos.

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
    name := "name_example" // string |  (optional)
    value := "value_example" // string |  (optional)
    processor := "processor_example" // string |  (optional)
    alertStatus := []string{"Inner_example"} // []string |  (optional)
    alertType := "alertType_example" // string |  (optional)
    startAt := "startAt_example" // string |  (optional)
    endAt := "endAt_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertInfoServiceApi.ListAlertInfo(context.Background()).Name(name).Value(value).Processor(processor).AlertStatus(alertStatus).AlertType(alertType).StartAt(startAt).EndAt(endAt).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoServiceApi.ListAlertInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlertInfo`: OpsListAlertInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertInfoServiceApi.ListAlertInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **value** | **string** |  | 
 **processor** | **string** |  | 
 **alertStatus** | **[]string** |  | 
 **alertType** | **string** |  | 
 **startAt** | **string** |  | 
 **endAt** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**OpsListAlertInfoResponse**](OpsListAlertInfoResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessAlertInfo

> CommonActionResponse ProcessAlertInfo(ctx, id).Execute()

processed a alert info.

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
    resp, r, err := api_client.AlertInfoServiceApi.ProcessAlertInfo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoServiceApi.ProcessAlertInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessAlertInfo`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertInfoServiceApi.ProcessAlertInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessAlertInfoRequest struct via the builder pattern


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

