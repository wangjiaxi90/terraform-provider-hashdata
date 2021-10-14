# \AccountTeamServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToTeam**](AccountTeamServiceApi.md#AddUserToTeam) | **Post** /account/api/v2/teams/{teamid}/users | Add an user to the team with given role.
[**CreateTeam**](AccountTeamServiceApi.md#CreateTeam) | **Post** /account/api/v2/teams | Create a new team.
[**DeleteTeam**](AccountTeamServiceApi.md#DeleteTeam) | **Delete** /account/api/v2/teams/{teamid} | Delete an existent team.
[**DeleteUserFromTeam**](AccountTeamServiceApi.md#DeleteUserFromTeam) | **Delete** /account/api/v2/teams/{teamid}/users/{userid} | Remove an user from a team.
[**DescribeTeam**](AccountTeamServiceApi.md#DescribeTeam) | **Get** /account/api/v2/teams/{teamid} | Describe the detail of a team.
[**ListTeam**](AccountTeamServiceApi.md#ListTeam) | **Get** /account/api/v2/teams | List all teams.
[**ListUserOfTeam**](AccountTeamServiceApi.md#ListUserOfTeam) | **Get** /account/api/v2/teams/{teamid}/users | List all users of a team.
[**ModifyTeam**](AccountTeamServiceApi.md#ModifyTeam) | **Put** /account/api/v2/teams/{teamid} | Modify an existent team.



## AddUserToTeam

> CommonActionResponse AddUserToTeam(ctx, teamid).Body(body).Execute()

Add an user to the team with given role.

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
    teamid := "teamid_example" // string | 
    body := *openapiclient.NewAccountAddUserToTeamRequest() // AccountAddUserToTeamRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTeamServiceApi.AddUserToTeam(context.Background(), teamid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTeamServiceApi.AddUserToTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserToTeam`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTeamServiceApi.AddUserToTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AccountAddUserToTeamRequest**](AccountAddUserToTeamRequest.md) |  | 

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


## CreateTeam

> AccountDescribeTeamResponse CreateTeam(ctx).Body(body).Execute()

Create a new team.

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
    body := *openapiclient.NewAccountCreateTeamRequest() // AccountCreateTeamRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTeamServiceApi.CreateTeam(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTeamServiceApi.CreateTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTeam`: AccountDescribeTeamResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTeamServiceApi.CreateTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccountCreateTeamRequest**](AccountCreateTeamRequest.md) |  | 

### Return type

[**AccountDescribeTeamResponse**](AccountDescribeTeamResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeam

> CommonActionResponse DeleteTeam(ctx, teamid).Execute()

Delete an existent team.

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
    teamid := "teamid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTeamServiceApi.DeleteTeam(context.Background(), teamid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTeamServiceApi.DeleteTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTeam`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTeamServiceApi.DeleteTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamRequest struct via the builder pattern


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


## DeleteUserFromTeam

> CommonActionResponse DeleteUserFromTeam(ctx, teamid, userid).Execute()

Remove an user from a team.

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
    teamid := "teamid_example" // string | 
    userid := "userid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTeamServiceApi.DeleteUserFromTeam(context.Background(), teamid, userid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTeamServiceApi.DeleteUserFromTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserFromTeam`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTeamServiceApi.DeleteUserFromTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamid** | **string** |  | 
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserFromTeamRequest struct via the builder pattern


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


## DescribeTeam

> AccountDescribeTeamResponse DescribeTeam(ctx, teamid).Execute()

Describe the detail of a team.

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
    teamid := "teamid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTeamServiceApi.DescribeTeam(context.Background(), teamid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTeamServiceApi.DescribeTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeTeam`: AccountDescribeTeamResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTeamServiceApi.DescribeTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountDescribeTeamResponse**](AccountDescribeTeamResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeam

> AccountListTeamResponse ListTeam(ctx).Id(id).Tenant(tenant).User(user).Name(name).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all teams.

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
    id := []string{"Inner_example"} // []string |  (optional)
    tenant := []string{"Inner_example"} // []string |  (optional)
    user := []string{"Inner_example"} // []string |  (optional)
    name := []string{"Inner_example"} // []string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTeamServiceApi.ListTeam(context.Background()).Id(id).Tenant(tenant).User(user).Name(name).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTeamServiceApi.ListTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeam`: AccountListTeamResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTeamServiceApi.ListTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **tenant** | **[]string** |  | 
 **user** | **[]string** |  | 
 **name** | **[]string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**AccountListTeamResponse**](AccountListTeamResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserOfTeam

> AccountListUserOfTeamResponse ListUserOfTeam(ctx, teamid).Offset(offset).Limit(limit).Execute()

List all users of a team.

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
    teamid := "teamid_example" // string | 
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTeamServiceApi.ListUserOfTeam(context.Background(), teamid).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTeamServiceApi.ListUserOfTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserOfTeam`: AccountListUserOfTeamResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTeamServiceApi.ListUserOfTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserOfTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**AccountListUserOfTeamResponse**](AccountListUserOfTeamResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTeam

> AccountDescribeTeamResponse ModifyTeam(ctx, teamid).Body(body).Execute()

Modify an existent team.

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
    teamid := "teamid_example" // string | 
    body := *openapiclient.NewAccountModifyTeamRequest() // AccountModifyTeamRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTeamServiceApi.ModifyTeam(context.Background(), teamid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTeamServiceApi.ModifyTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyTeam`: AccountDescribeTeamResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTeamServiceApi.ModifyTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AccountModifyTeamRequest**](AccountModifyTeamRequest.md) |  | 

### Return type

[**AccountDescribeTeamResponse**](AccountDescribeTeamResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

