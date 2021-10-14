# CoreDescribeInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**Components** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DestroyedAt** | Pointer to **string** |  | [optional] 
**ElasticNic** | Pointer to [**CoreDescribeElasticNicResponse**](CoreDescribeElasticNicResponse.md) |  | [optional] 
**HealthStatus** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Internet** | Pointer to [**CoreDescribeInternetResponse**](CoreDescribeInternetResponse.md) |  | [optional] 
**Ipaddresses** | Pointer to **[]string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ResourcePool** | Pointer to **string** |  | [optional] 
**ScaleType** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreDescribeInstanceResponse

`func NewCoreDescribeInstanceResponse() *CoreDescribeInstanceResponse`

NewCoreDescribeInstanceResponse instantiates a new CoreDescribeInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeInstanceResponseWithDefaults

`func NewCoreDescribeInstanceResponseWithDefaults() *CoreDescribeInstanceResponse`

NewCoreDescribeInstanceResponseWithDefaults instantiates a new CoreDescribeInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *CoreDescribeInstanceResponse) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *CoreDescribeInstanceResponse) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *CoreDescribeInstanceResponse) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *CoreDescribeInstanceResponse) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetComponents

`func (o *CoreDescribeInstanceResponse) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *CoreDescribeInstanceResponse) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *CoreDescribeInstanceResponse) SetComponents(v []string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *CoreDescribeInstanceResponse) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CoreDescribeInstanceResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CoreDescribeInstanceResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CoreDescribeInstanceResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CoreDescribeInstanceResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDestroyedAt

`func (o *CoreDescribeInstanceResponse) GetDestroyedAt() string`

GetDestroyedAt returns the DestroyedAt field if non-nil, zero value otherwise.

### GetDestroyedAtOk

`func (o *CoreDescribeInstanceResponse) GetDestroyedAtOk() (*string, bool)`

GetDestroyedAtOk returns a tuple with the DestroyedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyedAt

`func (o *CoreDescribeInstanceResponse) SetDestroyedAt(v string)`

SetDestroyedAt sets DestroyedAt field to given value.

### HasDestroyedAt

`func (o *CoreDescribeInstanceResponse) HasDestroyedAt() bool`

HasDestroyedAt returns a boolean if a field has been set.

### GetElasticNic

`func (o *CoreDescribeInstanceResponse) GetElasticNic() CoreDescribeElasticNicResponse`

GetElasticNic returns the ElasticNic field if non-nil, zero value otherwise.

### GetElasticNicOk

`func (o *CoreDescribeInstanceResponse) GetElasticNicOk() (*CoreDescribeElasticNicResponse, bool)`

GetElasticNicOk returns a tuple with the ElasticNic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticNic

`func (o *CoreDescribeInstanceResponse) SetElasticNic(v CoreDescribeElasticNicResponse)`

SetElasticNic sets ElasticNic field to given value.

### HasElasticNic

`func (o *CoreDescribeInstanceResponse) HasElasticNic() bool`

HasElasticNic returns a boolean if a field has been set.

### GetHealthStatus

`func (o *CoreDescribeInstanceResponse) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *CoreDescribeInstanceResponse) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *CoreDescribeInstanceResponse) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *CoreDescribeInstanceResponse) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetHostname

`func (o *CoreDescribeInstanceResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CoreDescribeInstanceResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CoreDescribeInstanceResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CoreDescribeInstanceResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *CoreDescribeInstanceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreDescribeInstanceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreDescribeInstanceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CoreDescribeInstanceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *CoreDescribeInstanceResponse) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CoreDescribeInstanceResponse) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CoreDescribeInstanceResponse) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CoreDescribeInstanceResponse) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetInternet

`func (o *CoreDescribeInstanceResponse) GetInternet() CoreDescribeInternetResponse`

GetInternet returns the Internet field if non-nil, zero value otherwise.

### GetInternetOk

`func (o *CoreDescribeInstanceResponse) GetInternetOk() (*CoreDescribeInternetResponse, bool)`

GetInternetOk returns a tuple with the Internet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternet

`func (o *CoreDescribeInstanceResponse) SetInternet(v CoreDescribeInternetResponse)`

SetInternet sets Internet field to given value.

### HasInternet

`func (o *CoreDescribeInstanceResponse) HasInternet() bool`

HasInternet returns a boolean if a field has been set.

### GetIpaddresses

`func (o *CoreDescribeInstanceResponse) GetIpaddresses() []string`

GetIpaddresses returns the Ipaddresses field if non-nil, zero value otherwise.

### GetIpaddressesOk

`func (o *CoreDescribeInstanceResponse) GetIpaddressesOk() (*[]string, bool)`

GetIpaddressesOk returns a tuple with the Ipaddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddresses

`func (o *CoreDescribeInstanceResponse) SetIpaddresses(v []string)`

SetIpaddresses sets Ipaddresses field to given value.

### HasIpaddresses

`func (o *CoreDescribeInstanceResponse) HasIpaddresses() bool`

HasIpaddresses returns a boolean if a field has been set.

### GetMessage

`func (o *CoreDescribeInstanceResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CoreDescribeInstanceResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CoreDescribeInstanceResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CoreDescribeInstanceResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *CoreDescribeInstanceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreDescribeInstanceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreDescribeInstanceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoreDescribeInstanceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourcePool

`func (o *CoreDescribeInstanceResponse) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *CoreDescribeInstanceResponse) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *CoreDescribeInstanceResponse) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *CoreDescribeInstanceResponse) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetScaleType

`func (o *CoreDescribeInstanceResponse) GetScaleType() string`

GetScaleType returns the ScaleType field if non-nil, zero value otherwise.

### GetScaleTypeOk

`func (o *CoreDescribeInstanceResponse) GetScaleTypeOk() (*string, bool)`

GetScaleTypeOk returns a tuple with the ScaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleType

`func (o *CoreDescribeInstanceResponse) SetScaleType(v string)`

SetScaleType sets ScaleType field to given value.

### HasScaleType

`func (o *CoreDescribeInstanceResponse) HasScaleType() bool`

HasScaleType returns a boolean if a field has been set.

### GetService

`func (o *CoreDescribeInstanceResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CoreDescribeInstanceResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CoreDescribeInstanceResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CoreDescribeInstanceResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *CoreDescribeInstanceResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CoreDescribeInstanceResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CoreDescribeInstanceResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CoreDescribeInstanceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *CoreDescribeInstanceResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CoreDescribeInstanceResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CoreDescribeInstanceResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CoreDescribeInstanceResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetType

`func (o *CoreDescribeInstanceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CoreDescribeInstanceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CoreDescribeInstanceResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CoreDescribeInstanceResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *CoreDescribeInstanceResponse) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CoreDescribeInstanceResponse) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CoreDescribeInstanceResponse) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CoreDescribeInstanceResponse) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetZone

`func (o *CoreDescribeInstanceResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CoreDescribeInstanceResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CoreDescribeInstanceResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CoreDescribeInstanceResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


