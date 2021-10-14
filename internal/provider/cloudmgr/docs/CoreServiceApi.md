# \CoreServiceApi

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProjectTag**](CoreServiceApi.md#AddProjectTag) | **Put** /default/api/v2/services/projectTag | Describe the detail of a given service.
[**CheckServiceActionPermission**](CoreServiceApi.md#CheckServiceActionPermission) | **Put** /default/api/v2/services/{id}/action | Check if login user has permission to perform action on a given service.
[**CountService**](CoreServiceApi.md#CountService) | **Get** /default/api/v2/services/count | count services.
[**DeleteService**](CoreServiceApi.md#DeleteService) | **Delete** /default/api/v2/services/{id} | Destroy a given service.
[**DescribeService**](CoreServiceApi.md#DescribeService) | **Get** /default/api/v2/services/{id} | Describe the detail of a given service.
[**DescribeTunnelInfo**](CoreServiceApi.md#DescribeTunnelInfo) | **Get** /default/api/v2/services/tunnel | describe tunnel info of service.
[**EnableInternalCharge**](CoreServiceApi.md#EnableInternalCharge) | **Put** /default/api/v2/services/enableInternalCharge | enable internal charge for service.
[**ExportServiceStateMachineDetailInfo**](CoreServiceApi.md#ExportServiceStateMachineDetailInfo) | **Get** /default/api/v2/services/{id}/statemachine/info | List service and component state machine detail info.
[**InplaceUpgradeService**](CoreServiceApi.md#InplaceUpgradeService) | **Post** /default/api/v2/services/{service}/upgrade/inplace | Inplace upgrade service version.
[**ListBucket**](CoreServiceApi.md#ListBucket) | **Get** /default/api/v2/services/{service}/buckets | List all bucket the given service has.
[**ListFaults**](CoreServiceApi.md#ListFaults) | **Get** /default/api/v2/services/{id}/faults | List all faluts which service got.
[**ListMetadata**](CoreServiceApi.md#ListMetadata) | **Get** /default/api/v2/services/{service}/metadata | List service metadata.
[**ListMonitorService**](CoreServiceApi.md#ListMonitorService) | **Get** /default/api/v2/services/monitor | List monitor service.
[**ListRecoveries**](CoreServiceApi.md#ListRecoveries) | **Get** /default/api/v2/services/{id}/recoveries | List all recoveries which service got.
[**ListService**](CoreServiceApi.md#ListService) | **Get** /default/api/v2/services | List all service.
[**ListServiceComponentHeatlhStatus**](CoreServiceApi.md#ListServiceComponentHeatlhStatus) | **Get** /default/api/v2/services/{id}/component/health | List service Component health status.
[**ListServiceDependent**](CoreServiceApi.md#ListServiceDependent) | **Get** /default/api/v2/services/{id}/dependency | List all dependencies which catalog got.
[**ListServiceFeature**](CoreServiceApi.md#ListServiceFeature) | **Get** /default/api/v2/services/{id}/features | List all features which a service support.
[**ListServiceInstance**](CoreServiceApi.md#ListServiceInstance) | **Get** /default/api/v2/services/{service}/instances | List all service instances.
[**ListServiceRecoveryDetail**](CoreServiceApi.md#ListServiceRecoveryDetail) | **Get** /default/api/v2/service/recoveries/{id} | List recovery detail by recovery id.
[**ListServiceReference**](CoreServiceApi.md#ListServiceReference) | **Get** /default/api/v2/services/{id}/reference | List all references which computing service got.
[**ListServiceStatuses**](CoreServiceApi.md#ListServiceStatuses) | **Get** /default/api/v2/services/statuses | List all status of a service can be in.
[**ListTeamOfService**](CoreServiceApi.md#ListTeamOfService) | **Get** /default/api/v2/services/{service}/teams | List all teams the given service belongs to.
[**ListUnhandledFaults**](CoreServiceApi.md#ListUnhandledFaults) | **Get** /default/api/v2/services/{id}/unhandledfaults | List all Unhandled faults which service got.
[**ModifyChargeInfo**](CoreServiceApi.md#ModifyChargeInfo) | **Put** /default/api/v2/services/modifyChargeInfo/{id} | Modify Charge infomation.
[**ModifyService**](CoreServiceApi.md#ModifyService) | **Put** /default/api/v2/services/{id} | Modify attribute of a service.
[**ProtectService**](CoreServiceApi.md#ProtectService) | **Get** /default/api/v2/services/{id}/protect | Protect the given service from delete or stop.
[**ResetStateMachine**](CoreServiceApi.md#ResetStateMachine) | **Get** /default/api/v2/services/{id}/statemachine/reset | reset service state machine by a given service.
[**ResizeVolumes**](CoreServiceApi.md#ResizeVolumes) | **Post** /default/api/v2/services/{id}/volumes/resize | Adjust volume size of a given service.
[**RestartService**](CoreServiceApi.md#RestartService) | **Post** /default/api/v2/services/{id}/restart | Restart a service.
[**ScaleInService**](CoreServiceApi.md#ScaleInService) | **Post** /default/api/v2/services/{id}/scalein | Scale in a given service.
[**ScaleOutService**](CoreServiceApi.md#ScaleOutService) | **Post** /default/api/v2/services/{id}/scaleout | Scale out a given service.
[**StartService**](CoreServiceApi.md#StartService) | **Post** /default/api/v2/services/{id}/start | Start a service.
[**StopService**](CoreServiceApi.md#StopService) | **Post** /default/api/v2/services/{id}/stop | Stop a service.
[**ToggleAlertAvailability**](CoreServiceApi.md#ToggleAlertAvailability) | **Get** /default/api/v2/services/{id}/alert | Toggle alert availability.
[**ToggleAutoRecoveryAvailability**](CoreServiceApi.md#ToggleAutoRecoveryAvailability) | **Get** /default/api/v2/services/{id}/recovery | Toggle auto recovery availability.
[**UpgradeGuestAgent**](CoreServiceApi.md#UpgradeGuestAgent) | **Get** /default/api/v2/services/{id}/guestagent/upgrade | upgrade all guest agent by a given service.
[**UpgradeService**](CoreServiceApi.md#UpgradeService) | **Post** /default/api/v2/services/{service}/upgrade | Upgrade service version.



## AddProjectTag

> CoreDescribeServiceResponse AddProjectTag(ctx).Body(body).Execute()

Describe the detail of a given service.

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
    body := *openapiclient.NewCoreAddProjectTagRequest() // CoreAddProjectTagRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.AddProjectTag(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.AddProjectTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProjectTag`: CoreDescribeServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.AddProjectTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddProjectTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreAddProjectTagRequest**](CoreAddProjectTagRequest.md) |  | 

### Return type

[**CoreDescribeServiceResponse**](CoreDescribeServiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckServiceActionPermission

> CoreDescribeServiceResponse CheckServiceActionPermission(ctx, id).Execute()

Check if login user has permission to perform action on a given service.

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
    resp, r, err := api_client.CoreServiceApi.CheckServiceActionPermission(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.CheckServiceActionPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckServiceActionPermission`: CoreDescribeServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.CheckServiceActionPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckServiceActionPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeServiceResponse**](CoreDescribeServiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountService

> CoreCountServiceResponse CountService(ctx).Execute()

count services.

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
    resp, r, err := api_client.CoreServiceApi.CountService(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.CountService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountService`: CoreCountServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.CountService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCountServiceRequest struct via the builder pattern


### Return type

[**CoreCountServiceResponse**](CoreCountServiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteService

> CommonDescribeJobResponse DeleteService(ctx, id).Force(force).External(external).Execute()

Destroy a given service.

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
    force := true // bool |  (optional)
    external := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.DeleteService(context.Background(), id).Force(force).External(external).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.DeleteService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteService`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.DeleteService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** |  | 
 **external** | **bool** |  | 

### Return type

[**CommonDescribeJobResponse**](CommonDescribeJobResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeService

> CoreDescribeServiceResponse DescribeService(ctx, id).Execute()

Describe the detail of a given service.

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
    resp, r, err := api_client.CoreServiceApi.DescribeService(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.DescribeService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeService`: CoreDescribeServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.DescribeService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreDescribeServiceResponse**](CoreDescribeServiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTunnelInfo

> CoreDescribeTunnelInfoResponse DescribeTunnelInfo(ctx).Service(service).Execute()

describe tunnel info of service.

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
    service := "service_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.DescribeTunnelInfo(context.Background()).Service(service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.DescribeTunnelInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeTunnelInfo`: CoreDescribeTunnelInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.DescribeTunnelInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTunnelInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** |  | 

### Return type

[**CoreDescribeTunnelInfoResponse**](CoreDescribeTunnelInfoResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableInternalCharge

> CoreDescribeServiceResponse EnableInternalCharge(ctx).Body(body).Execute()

enable internal charge for service.

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
    body := *openapiclient.NewCoreEnableInternalChargeRequest() // CoreEnableInternalChargeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.EnableInternalCharge(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.EnableInternalCharge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableInternalCharge`: CoreDescribeServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.EnableInternalCharge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableInternalChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CoreEnableInternalChargeRequest**](CoreEnableInternalChargeRequest.md) |  | 

### Return type

[**CoreDescribeServiceResponse**](CoreDescribeServiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportServiceStateMachineDetailInfo

> CoreExportServiceStateMachineInfoReponse ExportServiceStateMachineDetailInfo(ctx, id).Execute()

List service and component state machine detail info.

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
    resp, r, err := api_client.CoreServiceApi.ExportServiceStateMachineDetailInfo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ExportServiceStateMachineDetailInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportServiceStateMachineDetailInfo`: CoreExportServiceStateMachineInfoReponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ExportServiceStateMachineDetailInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportServiceStateMachineDetailInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreExportServiceStateMachineInfoReponse**](CoreExportServiceStateMachineInfoReponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InplaceUpgradeService

> CommonDescribeJobResponse InplaceUpgradeService(ctx, service).Body(body).Execute()

Inplace upgrade service version.

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
    body := *openapiclient.NewCoreInplaceUpgradeServiceRequest() // CoreInplaceUpgradeServiceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.InplaceUpgradeService(context.Background(), service).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.InplaceUpgradeService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InplaceUpgradeService`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.InplaceUpgradeService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInplaceUpgradeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreInplaceUpgradeServiceRequest**](CoreInplaceUpgradeServiceRequest.md) |  | 

### Return type

[**CommonDescribeJobResponse**](CommonDescribeJobResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBucket

> CoreListServiceBucketResponse ListBucket(ctx, service).Offset(offset).Limit(limit).Execute()

List all bucket the given service has.

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
    resp, r, err := api_client.CoreServiceApi.ListBucket(context.Background(), service).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBucket`: CoreListServiceBucketResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListServiceBucketResponse**](CoreListServiceBucketResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFaults

> CoreListFaultResponse ListFaults(ctx, id).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all faluts which service got.

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
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ListFaults(context.Background(), id).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListFaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFaults`: CoreListFaultResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListFaults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListFaultResponse**](CoreListFaultResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetadata

> CoreListServiceMetadataResponse ListMetadata(ctx, service).Offset(offset).Limit(limit).Execute()

List service metadata.

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
    resp, r, err := api_client.CoreServiceApi.ListMetadata(context.Background(), service).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMetadata`: CoreListServiceMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListServiceMetadataResponse**](CoreListServiceMetadataResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMonitorService

> CoreListMonitorServiceResponse ListMonitorService(ctx).Type_(type_).Status(status).Offset(offset).Limit(limit).Execute()

List monitor service.

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
    type_ := []string{"Inner_example"} // []string |  (optional)
    status := []string{"Inner_example"} // []string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ListMonitorService(context.Background()).Type_(type_).Status(status).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListMonitorService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMonitorService`: CoreListMonitorServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListMonitorService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMonitorServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **[]string** |  | 
 **status** | **[]string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListMonitorServiceResponse**](CoreListMonitorServiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecoveries

> CoreListRecoveryResponse ListRecoveries(ctx, id).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all recoveries which service got.

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
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ListRecoveries(context.Background(), id).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListRecoveries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecoveries`: CoreListRecoveryResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListRecoveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRecoveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListRecoveryResponse**](CoreListRecoveryResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> CoreListServiceResponse ListService(ctx).Id(id).Tenant(tenant).Owner(owner).Name(name).Vendor(vendor).Zone(zone).Type_(type_).Status(status).Region(region).HealthStatus(healthStatus).Recovering(recovering).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()

List all service.

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
    owner := []string{"Inner_example"} // []string |  (optional)
    name := []string{"Inner_example"} // []string |  (optional)
    vendor := []string{"Inner_example"} // []string |  (optional)
    zone := []string{"Inner_example"} // []string |  (optional)
    type_ := []string{"Inner_example"} // []string |  (optional)
    status := []string{"Inner_example"} // []string |  (optional)
    region := []string{"Inner_example"} // []string |  (optional)
    healthStatus := []string{"Inner_example"} // []string |  (optional)
    recovering := []bool{false} // []bool |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ListService(context.Background()).Id(id).Tenant(tenant).Owner(owner).Name(name).Vendor(vendor).Zone(zone).Type_(type_).Status(status).Region(region).HealthStatus(healthStatus).Recovering(recovering).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListService`: CoreListServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **tenant** | **[]string** |  | 
 **owner** | **[]string** |  | 
 **name** | **[]string** |  | 
 **vendor** | **[]string** |  | 
 **zone** | **[]string** |  | 
 **type_** | **[]string** |  | 
 **status** | **[]string** |  | 
 **region** | **[]string** |  | 
 **healthStatus** | **[]string** |  | 
 **recovering** | **[]bool** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListServiceResponse**](CoreListServiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceComponentHeatlhStatus

> CoreListServiceComponentHealthResponse ListServiceComponentHeatlhStatus(ctx, id).ComponentType(componentType).Offset(offset).Limit(limit).Execute()

List service Component health status.

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
    componentType := "componentType_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ListServiceComponentHeatlhStatus(context.Background(), id).ComponentType(componentType).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListServiceComponentHeatlhStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceComponentHeatlhStatus`: CoreListServiceComponentHealthResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListServiceComponentHeatlhStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceComponentHeatlhStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentType** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListServiceComponentHealthResponse**](CoreListServiceComponentHealthResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceDependent

> CoreListServiceDependencyResponse ListServiceDependent(ctx, id).Offset(offset).Limit(limit).Execute()

List all dependencies which catalog got.

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
    resp, r, err := api_client.CoreServiceApi.ListServiceDependent(context.Background(), id).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListServiceDependent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceDependent`: CoreListServiceDependencyResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListServiceDependent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceDependentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListServiceDependencyResponse**](CoreListServiceDependencyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceFeature

> CoreListServiceFeatureResponse ListServiceFeature(ctx, id).Execute()

List all features which a service support.

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
    resp, r, err := api_client.CoreServiceApi.ListServiceFeature(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListServiceFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceFeature`: CoreListServiceFeatureResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListServiceFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreListServiceFeatureResponse**](CoreListServiceFeatureResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceInstance

> CoreListInstanceResponse ListServiceInstance(ctx, service).Component(component).Offset(offset).Limit(limit).Execute()

List all service instances.

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
    component := []string{"Inner_example"} // []string |  (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ListServiceInstance(context.Background(), service).Component(component).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListServiceInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceInstance`: CoreListInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListServiceInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **component** | **[]string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListInstanceResponse**](CoreListInstanceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceRecoveryDetail

> CoreListServiceRecoveryDetailResponse ListServiceRecoveryDetail(ctx, id).Execute()

List recovery detail by recovery id.

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
    resp, r, err := api_client.CoreServiceApi.ListServiceRecoveryDetail(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListServiceRecoveryDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceRecoveryDetail`: CoreListServiceRecoveryDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListServiceRecoveryDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceRecoveryDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CoreListServiceRecoveryDetailResponse**](CoreListServiceRecoveryDetailResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceReference

> CoreListServiceDependencyResponse ListServiceReference(ctx, id).Offset(offset).Limit(limit).Execute()

List all references which computing service got.

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
    resp, r, err := api_client.CoreServiceApi.ListServiceReference(context.Background(), id).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListServiceReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceReference`: CoreListServiceDependencyResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListServiceReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListServiceDependencyResponse**](CoreListServiceDependencyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceStatuses

> CoreListServiceStatusesResponse ListServiceStatuses(ctx).Execute()

List all status of a service can be in.

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
    resp, r, err := api_client.CoreServiceApi.ListServiceStatuses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListServiceStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceStatuses`: CoreListServiceStatusesResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListServiceStatuses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceStatusesRequest struct via the builder pattern


### Return type

[**CoreListServiceStatusesResponse**](CoreListServiceStatusesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamOfService

> CoreListTeamOfServiceResponse ListTeamOfService(ctx, service).Offset(offset).Limit(limit).Execute()

List all teams the given service belongs to.

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
    resp, r, err := api_client.CoreServiceApi.ListTeamOfService(context.Background(), service).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListTeamOfService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamOfService`: CoreListTeamOfServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListTeamOfService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamOfServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListTeamOfServiceResponse**](CoreListTeamOfServiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUnhandledFaults

> CoreListUnhandledFaultResponse ListUnhandledFaults(ctx, id).Offset(offset).Limit(limit).Execute()

List all Unhandled faults which service got.

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
    resp, r, err := api_client.CoreServiceApi.ListUnhandledFaults(context.Background(), id).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ListUnhandledFaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUnhandledFaults`: CoreListUnhandledFaultResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ListUnhandledFaults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUnhandledFaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**CoreListUnhandledFaultResponse**](CoreListUnhandledFaultResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyChargeInfo

> CoreDescribeServiceResponse ModifyChargeInfo(ctx, id).Body(body).Execute()

Modify Charge infomation.

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
    body := *openapiclient.NewCoreModifyChargeInfoRequest() // CoreModifyChargeInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ModifyChargeInfo(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ModifyChargeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyChargeInfo`: CoreDescribeServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ModifyChargeInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyChargeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreModifyChargeInfoRequest**](CoreModifyChargeInfoRequest.md) |  | 

### Return type

[**CoreDescribeServiceResponse**](CoreDescribeServiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyService

> CoreDescribeServiceResponse ModifyService(ctx, id).Body(body).Execute()

Modify attribute of a service.

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
    body := *openapiclient.NewCoreModifyServiceRequest() // CoreModifyServiceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ModifyService(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ModifyService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyService`: CoreDescribeServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ModifyService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreModifyServiceRequest**](CoreModifyServiceRequest.md) |  | 

### Return type

[**CoreDescribeServiceResponse**](CoreDescribeServiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProtectService

> CoreDescribeServiceResponse ProtectService(ctx, id).Protect(protect).Execute()

Protect the given service from delete or stop.

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
    protect := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ProtectService(context.Background(), id).Protect(protect).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ProtectService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProtectService`: CoreDescribeServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ProtectService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProtectServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **protect** | **bool** |  | 

### Return type

[**CoreDescribeServiceResponse**](CoreDescribeServiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetStateMachine

> string ResetStateMachine(ctx, id).Execute()

reset service state machine by a given service.

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
    resp, r, err := api_client.CoreServiceApi.ResetStateMachine(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ResetStateMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetStateMachine`: string
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ResetStateMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetStateMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeVolumes

> CommonDescribeJobResponse ResizeVolumes(ctx, id).Body(body).Execute()

Adjust volume size of a given service.

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
    body := *openapiclient.NewCoreResizeServiceVolumesRequest() // CoreResizeServiceVolumesRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ResizeVolumes(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ResizeVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResizeVolumes`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ResizeVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreResizeServiceVolumesRequest**](CoreResizeServiceVolumesRequest.md) |  | 

### Return type

[**CommonDescribeJobResponse**](CommonDescribeJobResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartService

> CommonDescribeJobResponse RestartService(ctx, id).Body(body).Execute()

Restart a service.

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
    body := map[string]interface{}(Object) // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.RestartService(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.RestartService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartService`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.RestartService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**CommonDescribeJobResponse**](CommonDescribeJobResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScaleInService

> CommonDescribeJobResponse ScaleInService(ctx, id).Body(body).Execute()

Scale in a given service.

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
    body := *openapiclient.NewCoreScaleInServiceRequest() // CoreScaleInServiceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ScaleInService(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ScaleInService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScaleInService`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ScaleInService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScaleInServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreScaleInServiceRequest**](CoreScaleInServiceRequest.md) |  | 

### Return type

[**CommonDescribeJobResponse**](CommonDescribeJobResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScaleOutService

> CommonDescribeJobResponse ScaleOutService(ctx, id).Body(body).Execute()

Scale out a given service.

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
    body := *openapiclient.NewCoreScaleOutServiceRequest() // CoreScaleOutServiceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ScaleOutService(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ScaleOutService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScaleOutService`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ScaleOutService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScaleOutServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreScaleOutServiceRequest**](CoreScaleOutServiceRequest.md) |  | 

### Return type

[**CommonDescribeJobResponse**](CommonDescribeJobResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartService

> CommonDescribeJobResponse StartService(ctx, id).Body(body).Execute()

Start a service.

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
    body := *openapiclient.NewCoreStartServiceRequest() // CoreStartServiceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.StartService(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.StartService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartService`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.StartService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreStartServiceRequest**](CoreStartServiceRequest.md) |  | 

### Return type

[**CommonDescribeJobResponse**](CommonDescribeJobResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopService

> CommonDescribeJobResponse StopService(ctx, id).Body(body).Execute()

Stop a service.

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
    body := *openapiclient.NewCoreStopServiceRequest() // CoreStopServiceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.StopService(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.StopService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopService`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.StopService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreStopServiceRequest**](CoreStopServiceRequest.md) |  | 

### Return type

[**CommonDescribeJobResponse**](CommonDescribeJobResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleAlertAvailability

> CoreDescribeServiceResponse ToggleAlertAvailability(ctx, id).EnableAlert(enableAlert).Execute()

Toggle alert availability.

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
    enableAlert := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ToggleAlertAvailability(context.Background(), id).EnableAlert(enableAlert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ToggleAlertAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToggleAlertAvailability`: CoreDescribeServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ToggleAlertAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleAlertAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableAlert** | **bool** |  | 

### Return type

[**CoreDescribeServiceResponse**](CoreDescribeServiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleAutoRecoveryAvailability

> CoreDescribeServiceResponse ToggleAutoRecoveryAvailability(ctx, id).DisableAutoRecovery(disableAutoRecovery).Execute()

Toggle auto recovery availability.

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
    disableAutoRecovery := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.ToggleAutoRecoveryAvailability(context.Background(), id).DisableAutoRecovery(disableAutoRecovery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.ToggleAutoRecoveryAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToggleAutoRecoveryAvailability`: CoreDescribeServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.ToggleAutoRecoveryAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleAutoRecoveryAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disableAutoRecovery** | **bool** |  | 

### Return type

[**CoreDescribeServiceResponse**](CoreDescribeServiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeGuestAgent

> CommonDescribeJobResponse UpgradeGuestAgent(ctx, id).Execute()

upgrade all guest agent by a given service.

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
    resp, r, err := api_client.CoreServiceApi.UpgradeGuestAgent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.UpgradeGuestAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeGuestAgent`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.UpgradeGuestAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeGuestAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommonDescribeJobResponse**](CommonDescribeJobResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeService

> CommonDescribeJobResponse UpgradeService(ctx, service).Body(body).Execute()

Upgrade service version.

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
    body := *openapiclient.NewCoreUpgradeServiceRequest() // CoreUpgradeServiceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CoreServiceApi.UpgradeService(context.Background(), service).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreServiceApi.UpgradeService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeService`: CommonDescribeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `CoreServiceApi.UpgradeService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CoreUpgradeServiceRequest**](CoreUpgradeServiceRequest.md) |  | 

### Return type

[**CommonDescribeJobResponse**](CommonDescribeJobResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

