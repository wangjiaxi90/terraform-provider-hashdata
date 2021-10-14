# \CoreTeamServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServiceToTeam**](CoreTeamServiceApi.md#AddServiceToTeam) | **Post** /default/api/v2/teams/{team}/services | Add a service to a team
[**DeleteServiceFromTeam**](CoreTeamServiceApi.md#DeleteServiceFromTeam) | **Delete** /default/api/v2/teams/{team}/services/{service} | Remove service from a team.
[**ListServiceOfTeam**](CoreTeamServiceApi.md#ListServiceOfTeam) | **Get** /default/api/v2/teams/{id}/services | List all service of a team.



## AddServiceToTeam

> CommonActionResponse AddServiceToTeam(ctx, team).Body(body).Execute()

Add a service to a team

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
    team := "team_example" // string | 
    body := *openapiclient.NewCoreAddServiceToTeamRequest() // CoreAddServiceToTeamRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreTeamServiceApi.AddServiceToTeam(context.Background(), team).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreTeamServiceApi.AddServiceToTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddServiceToTeam`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreTeamServiceApi.AddServiceToTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**team** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddServiceToTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreAddServiceToTeamRequest**](CoreAddServiceToTeamRequest.md) |  | 

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


## DeleteServiceFromTeam

> CommonActionResponse DeleteServiceFromTeam(ctx, team, service).Execute()

Remove service from a team.

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
    team := "team_example" // string | 
    service := "service_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreTeamServiceApi.DeleteServiceFromTeam(context.Background(), team, service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreTeamServiceApi.DeleteServiceFromTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteServiceFromTeam`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreTeamServiceApi.DeleteServiceFromTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**team** | **string** |  | 
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceFromTeamRequest struct via the builder pattern


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


## ListServiceOfTeam

> CoreListServiceOfTeamResponse ListServiceOfTeam(ctx, id).Offset(offset).Limit(limit).Execute()

List all service of a team.

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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreTeamServiceApi.ListServiceOfTeam(context.Background(), id).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreTeamServiceApi.ListServiceOfTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceOfTeam`: CoreListServiceOfTeamResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreTeamServiceApi.ListServiceOfTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceOfTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListServiceOfTeamResponse**](CoreListServiceOfTeamResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

