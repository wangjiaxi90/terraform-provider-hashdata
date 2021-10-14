# \ContactGroupServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContactToGroup**](ContactGroupServiceApi.md#AddContactToGroup) | **Post** /notification/api/v2/groups/{group}/contacts | Add contacts to group.
[**CreateContactGroup**](ContactGroupServiceApi.md#CreateContactGroup) | **Post** /notification/api/v2/groups | Create a contact group.
[**DeleteContactGroup**](ContactGroupServiceApi.md#DeleteContactGroup) | **Delete** /notification/api/v2/groups/{id} | Delete a contact group.
[**DescribeContactGroup**](ContactGroupServiceApi.md#DescribeContactGroup) | **Get** /notification/api/v2/groups/{id} | Describe the detail of a given contact group.
[**ListContactGroup**](ContactGroupServiceApi.md#ListContactGroup) | **Get** /notification/api/v2/groups | List contact group.
[**RemoveContactFromGroup**](ContactGroupServiceApi.md#RemoveContactFromGroup) | **Delete** /notification/api/v2/groups/{group}/contacts/{contact} | Remove a contact from group.
[**TestContactGroup**](ContactGroupServiceApi.md#TestContactGroup) | **Put** /notification/api/v2/groups/{id} | Test a given contact group.



## AddContactToGroup

> NotificationDescribeContactGroupResponse AddContactToGroup(ctx, group).Body(body).Execute()

Add contacts to group.

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
    group := "group_example" // string | 
    body := *openapiclient.NewNotificationAddContactToGroupRequest() // NotificationAddContactToGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactGroupServiceApi.AddContactToGroup(context.Background(), group).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupServiceApi.AddContactToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddContactToGroup`: NotificationDescribeContactGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ContactGroupServiceApi.AddContactToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**group** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddContactToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NotificationAddContactToGroupRequest**](NotificationAddContactToGroupRequest.md) |  | 

### Return type

[**NotificationDescribeContactGroupResponse**](NotificationDescribeContactGroupResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContactGroup

> NotificationDescribeContactGroupResponse CreateContactGroup(ctx).Body(body).Execute()

Create a contact group.

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
    body := *openapiclient.NewNotificationCreateContactGroupRequest() // NotificationCreateContactGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactGroupServiceApi.CreateContactGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupServiceApi.CreateContactGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContactGroup`: NotificationDescribeContactGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ContactGroupServiceApi.CreateContactGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NotificationCreateContactGroupRequest**](NotificationCreateContactGroupRequest.md) |  | 

### Return type

[**NotificationDescribeContactGroupResponse**](NotificationDescribeContactGroupResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactGroup

> CommonActionResponse DeleteContactGroup(ctx, id).Execute()

Delete a contact group.

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
    resp, r, err := api_client.ContactGroupServiceApi.DeleteContactGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupServiceApi.DeleteContactGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContactGroup`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `ContactGroupServiceApi.DeleteContactGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactGroupRequest struct via the builder pattern


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


## DescribeContactGroup

> NotificationDescribeContactGroupResponse DescribeContactGroup(ctx, id).Execute()

Describe the detail of a given contact group.

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
    resp, r, err := api_client.ContactGroupServiceApi.DescribeContactGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupServiceApi.DescribeContactGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeContactGroup`: NotificationDescribeContactGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ContactGroupServiceApi.DescribeContactGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeContactGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationDescribeContactGroupResponse**](NotificationDescribeContactGroupResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContactGroup

> NotificationListContactGroupResponse ListContactGroup(ctx).Id(id).Name(name).User(user).Tenant(tenant).Offset(offset).Limit(limit).Execute()

List contact group.

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
    name := []string{"Inner_example"} // []string |  (optional)
    user := []string{"Inner_example"} // []string |  (optional)
    tenant := []string{"Inner_example"} // []string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactGroupServiceApi.ListContactGroup(context.Background()).Id(id).Name(name).User(user).Tenant(tenant).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupServiceApi.ListContactGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContactGroup`: NotificationListContactGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ContactGroupServiceApi.ListContactGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContactGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **name** | **[]string** |  | 
 **user** | **[]string** |  | 
 **tenant** | **[]string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**NotificationListContactGroupResponse**](NotificationListContactGroupResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveContactFromGroup

> NotificationDescribeContactGroupResponse RemoveContactFromGroup(ctx, group, contact).Execute()

Remove a contact from group.

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
    group := "group_example" // string | 
    contact := "contact_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactGroupServiceApi.RemoveContactFromGroup(context.Background(), group, contact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupServiceApi.RemoveContactFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveContactFromGroup`: NotificationDescribeContactGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ContactGroupServiceApi.RemoveContactFromGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**group** | **string** |  | 
**contact** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveContactFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NotificationDescribeContactGroupResponse**](NotificationDescribeContactGroupResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestContactGroup

> CommonActionResponse TestContactGroup(ctx, id).Execute()

Test a given contact group.

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
    resp, r, err := api_client.ContactGroupServiceApi.TestContactGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupServiceApi.TestContactGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestContactGroup`: CommonActionResponse
    fmt.Fprintf(os.Stdout, "Response from `ContactGroupServiceApi.TestContactGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestContactGroupRequest struct via the builder pattern


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

