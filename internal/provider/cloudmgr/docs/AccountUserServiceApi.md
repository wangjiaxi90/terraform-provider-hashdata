# \AccountUserServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePassword**](AccountUserServiceApi.md#ChangePassword) | **Put** /account/api/v2/users/{userid}/password | Modify user&#39;s password.
[**ComplementExternalUserAndTenant**](AccountUserServiceApi.md#ComplementExternalUserAndTenant) | **Put** /account/api/v2/users/external/{user.id}/tenants/external/{tenant.id} | Complete external user&#39;s information and external tenant&#39;s information.
[**CreateExternalUser**](AccountUserServiceApi.md#CreateExternalUser) | **Post** /account/api/v2/users/external | Create a new user from outside.
[**CreateUser**](AccountUserServiceApi.md#CreateUser) | **Post** /account/api/v2/users | Create a new user.
[**DeleteUser**](AccountUserServiceApi.md#DeleteUser) | **Delete** /account/api/v2/users/{userid} | Delete an user.
[**DescribeExternalUser**](AccountUserServiceApi.md#DescribeExternalUser) | **Get** /account/api/v2/users/external | Describe the detail of an user.
[**DescribeSelf**](AccountUserServiceApi.md#DescribeSelf) | **Get** /account/api/v2/users/self | Describe current login user.
[**DescribeUser**](AccountUserServiceApi.md#DescribeUser) | **Get** /account/api/v2/users/{userid} | Describe the detail of an user.
[**ListTeamOfUser**](AccountUserServiceApi.md#ListTeamOfUser) | **Get** /account/api/v2/users/{userid}/teams | List all teams which the user belongs to.
[**ListUser**](AccountUserServiceApi.md#ListUser) | **Get** /account/api/v2/users | List all users.
[**ModifyUser**](AccountUserServiceApi.md#ModifyUser) | **Put** /account/api/v2/users/{userid} | Modify an user.



## ChangePassword

> CommonActionResponse ChangePassword(ctx, userid).Body(body).Execute()

Modify user's password.

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
    userid := "userid_example" // string | 
    body := *openapiclient.NewAccountChangePasswordRequest() // AccountChangePasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUserServiceApi.ChangePassword(context.Background(), userid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUserServiceApi.ChangePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangePassword`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountUserServiceApi.ChangePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AccountChangePasswordRequest**](AccountChangePasswordRequest.md) |  | 

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


## ComplementExternalUserAndTenant

> CommonActionResponse ComplementExternalUserAndTenant(ctx, userId, tenantId).Body(body).Execute()

Complete external user's information and external tenant's information.

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
    userId := "userId_example" // string | 
    tenantId := "tenantId_example" // string | 
    body := *openapiclient.NewAccountComplementExternalUserAndTenantRequest() // AccountComplementExternalUserAndTenantRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUserServiceApi.ComplementExternalUserAndTenant(context.Background(), userId, tenantId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUserServiceApi.ComplementExternalUserAndTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComplementExternalUserAndTenant`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountUserServiceApi.ComplementExternalUserAndTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComplementExternalUserAndTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AccountComplementExternalUserAndTenantRequest**](AccountComplementExternalUserAndTenantRequest.md) |  | 

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


## CreateExternalUser

> AccountDescribeUserResponse CreateExternalUser(ctx).Body(body).Execute()

Create a new user from outside.

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
    body := *openapiclient.NewAccountCreateExternalUserRequest() // AccountCreateExternalUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUserServiceApi.CreateExternalUser(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUserServiceApi.CreateExternalUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExternalUser`: AccountDescribeUserResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountUserServiceApi.CreateExternalUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccountCreateExternalUserRequest**](AccountCreateExternalUserRequest.md) |  | 

### Return type

[**AccountDescribeUserResponse**](AccountDescribeUserResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> AccountDescribeUserResponse CreateUser(ctx).Body(body).Execute()

Create a new user.

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
    body := *openapiclient.NewAccountCreateUserRequest() // AccountCreateUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUserServiceApi.CreateUser(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUserServiceApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: AccountDescribeUserResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountUserServiceApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccountCreateUserRequest**](AccountCreateUserRequest.md) |  | 

### Return type

[**AccountDescribeUserResponse**](AccountDescribeUserResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> CommonActionResponse DeleteUser(ctx, userid).Execute()

Delete an user.

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
    userid := "userid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUserServiceApi.DeleteUser(context.Background(), userid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUserServiceApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountUserServiceApi.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## DescribeExternalUser

> AccountDescribeUserResponse DescribeExternalUser(ctx).ExternalId(externalId).Vendor(vendor).Execute()

Describe the detail of an user.

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
    externalId := "externalId_example" // string |  (optional)
    vendor := "vendor_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUserServiceApi.DescribeExternalUser(context.Background()).ExternalId(externalId).Vendor(vendor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUserServiceApi.DescribeExternalUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeExternalUser`: AccountDescribeUserResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountUserServiceApi.DescribeExternalUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeExternalUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalId** | **string** |  | 
 **vendor** | **string** |  | 

### Return type

[**AccountDescribeUserResponse**](AccountDescribeUserResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeSelf

> AccountDescribeUserResponse DescribeSelf(ctx).Execute()

Describe current login user.

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
    resp, r, err := api_client.AccountUserServiceApi.DescribeSelf(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUserServiceApi.DescribeSelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeSelf`: AccountDescribeUserResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountUserServiceApi.DescribeSelf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeSelfRequest struct via the builder pattern


### Return type

[**AccountDescribeUserResponse**](AccountDescribeUserResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeUser

> AccountDescribeUserResponse DescribeUser(ctx, userid).Execute()

Describe the detail of an user.

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
    userid := "userid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUserServiceApi.DescribeUser(context.Background(), userid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUserServiceApi.DescribeUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeUser`: AccountDescribeUserResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountUserServiceApi.DescribeUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountDescribeUserResponse**](AccountDescribeUserResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamOfUser

> AccountListTeamOfUserResponse ListTeamOfUser(ctx, userid).Offset(offset).Limit(limit).Execute()

List all teams which the user belongs to.

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
    userid := "userid_example" // string | 
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUserServiceApi.ListTeamOfUser(context.Background(), userid).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUserServiceApi.ListTeamOfUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamOfUser`: AccountListTeamOfUserResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountUserServiceApi.ListTeamOfUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamOfUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**AccountListTeamOfUserResponse**](AccountListTeamOfUserResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUser

> AccountListUserResponse ListUser(ctx).Id(id).Tenant(tenant).Username(username).Surname(surname).GivenName(givenName).Email(email).Expired(expired).Locked(locked).External(external).PasswordExpired(passwordExpired).Enabled(enabled).ExternalId(externalId).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all users.

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
    username := []string{"Inner_example"} // []string |  (optional)
    surname := []string{"Inner_example"} // []string |  (optional)
    givenName := []string{"Inner_example"} // []string |  (optional)
    email := []string{"Inner_example"} // []string |  (optional)
    expired := true // bool |  (optional)
    locked := true // bool |  (optional)
    external := true // bool |  (optional)
    passwordExpired := true // bool |  (optional)
    enabled := true // bool |  (optional)
    externalId := "externalId_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUserServiceApi.ListUser(context.Background()).Id(id).Tenant(tenant).Username(username).Surname(surname).GivenName(givenName).Email(email).Expired(expired).Locked(locked).External(external).PasswordExpired(passwordExpired).Enabled(enabled).ExternalId(externalId).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUserServiceApi.ListUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUser`: AccountListUserResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountUserServiceApi.ListUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **tenant** | **[]string** |  | 
 **username** | **[]string** |  | 
 **surname** | **[]string** |  | 
 **givenName** | **[]string** |  | 
 **email** | **[]string** |  | 
 **expired** | **bool** |  | 
 **locked** | **bool** |  | 
 **external** | **bool** |  | 
 **passwordExpired** | **bool** |  | 
 **enabled** | **bool** |  | 
 **externalId** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**AccountListUserResponse**](AccountListUserResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyUser

> AccountDescribeUserResponse ModifyUser(ctx, userid).Body(body).Execute()

Modify an user.

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
    userid := "userid_example" // string | 
    body := *openapiclient.NewAccountModifyUserRequest() // AccountModifyUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUserServiceApi.ModifyUser(context.Background(), userid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUserServiceApi.ModifyUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyUser`: AccountDescribeUserResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountUserServiceApi.ModifyUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AccountModifyUserRequest**](AccountModifyUserRequest.md) |  | 

### Return type

[**AccountDescribeUserResponse**](AccountDescribeUserResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

