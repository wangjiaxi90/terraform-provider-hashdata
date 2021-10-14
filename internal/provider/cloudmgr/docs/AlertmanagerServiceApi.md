# \AlertmanagerServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSilence**](AlertmanagerServiceApi.md#CreateSilence) | **Post** /ops/api/v2/silence | create silence of alertmanagement.
[**DeleteSilence**](AlertmanagerServiceApi.md#DeleteSilence) | **Delete** /ops/api/v2/silence | delete silence of alertmanagement.
[**DescribeSilence**](AlertmanagerServiceApi.md#DescribeSilence) | **Get** /ops/api/v2/silence/{id} | decribe silence of alertmanagement.
[**ListAlert**](AlertmanagerServiceApi.md#ListAlert) | **Get** /ops/api/v2/alerts | list alerts of alertmanagement.
[**ListSilence**](AlertmanagerServiceApi.md#ListSilence) | **Get** /ops/api/v2/silences | list silence of alertmanagement.
[**ModifySilence**](AlertmanagerServiceApi.md#ModifySilence) | **Put** /ops/api/v2/silence | create silence of alertmanagement.



## CreateSilence

> OpsCreateSilenceResponse CreateSilence(ctx).Body(body).Execute()

create silence of alertmanagement.

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
    body := *openapiclient.NewOpsCreateSilenceRequest() // OpsCreateSilenceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertmanagerServiceApi.CreateSilence(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertmanagerServiceApi.CreateSilence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSilence`: OpsCreateSilenceResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertmanagerServiceApi.CreateSilence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSilenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OpsCreateSilenceRequest**](OpsCreateSilenceRequest.md) |  | 

### Return type

[**OpsCreateSilenceResponse**](OpsCreateSilenceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSilence

> map[string]interface{} DeleteSilence(ctx).Id(id).Execute()

delete silence of alertmanagement.

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
    id := "id_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertmanagerServiceApi.DeleteSilence(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertmanagerServiceApi.DeleteSilence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSilence`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AlertmanagerServiceApi.DeleteSilence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSilenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

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


## DescribeSilence

> OpsDescribeSilenceResponse DescribeSilence(ctx, id).Execute()

decribe silence of alertmanagement.

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
    resp, r, err := api_client.AlertmanagerServiceApi.DescribeSilence(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertmanagerServiceApi.DescribeSilence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeSilence`: OpsDescribeSilenceResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertmanagerServiceApi.DescribeSilence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeSilenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpsDescribeSilenceResponse**](OpsDescribeSilenceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlert

> OpsListAlertResponse ListAlert(ctx).Active(active).Silenced(silenced).Inhibited(inhibited).Unprocessed(unprocessed).Receiver(receiver).Filter(filter).Execute()

list alerts of alertmanagement.

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
    active := true // bool |  (optional)
    silenced := true // bool |  (optional)
    inhibited := true // bool |  (optional)
    unprocessed := true // bool |  (optional)
    receiver := "receiver_example" // string |  (optional)
    filter := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertmanagerServiceApi.ListAlert(context.Background()).Active(active).Silenced(silenced).Inhibited(inhibited).Unprocessed(unprocessed).Receiver(receiver).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertmanagerServiceApi.ListAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlert`: OpsListAlertResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertmanagerServiceApi.ListAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **bool** |  | 
 **silenced** | **bool** |  | 
 **inhibited** | **bool** |  | 
 **unprocessed** | **bool** |  | 
 **receiver** | **string** |  | 
 **filter** | **[]string** |  | 

### Return type

[**OpsListAlertResponse**](OpsListAlertResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSilence

> OpsListSilenceResponse ListSilence(ctx).Filter(filter).Execute()

list silence of alertmanagement.

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
    filter := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertmanagerServiceApi.ListSilence(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertmanagerServiceApi.ListSilence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSilence`: OpsListSilenceResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertmanagerServiceApi.ListSilence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSilenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **[]string** |  | 

### Return type

[**OpsListSilenceResponse**](OpsListSilenceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySilence

> OpsModifySilenceResponse ModifySilence(ctx).Body(body).Execute()

create silence of alertmanagement.

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
    body := *openapiclient.NewOpsModifySilenceRequest() // OpsModifySilenceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertmanagerServiceApi.ModifySilence(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertmanagerServiceApi.ModifySilence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifySilence`: OpsModifySilenceResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertmanagerServiceApi.ModifySilence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifySilenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OpsModifySilenceRequest**](OpsModifySilenceRequest.md) |  | 

### Return type

[**OpsModifySilenceResponse**](OpsModifySilenceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

