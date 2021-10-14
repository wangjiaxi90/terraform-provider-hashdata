# CoreCreateResourcePoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultOssZone** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**OssConnectivity** | Pointer to **[]string** |  | [optional] 
**Tenant** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreCreateResourcePoolRequest

`func NewCoreCreateResourcePoolRequest() *CoreCreateResourcePoolRequest`

NewCoreCreateResourcePoolRequest instantiates a new CoreCreateResourcePoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCreateResourcePoolRequestWithDefaults

`func NewCoreCreateResourcePoolRequestWithDefaults() *CoreCreateResourcePoolRequest`

NewCoreCreateResourcePoolRequestWithDefaults instantiates a new CoreCreateResourcePoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultOssZone

`func (o *CoreCreateResourcePoolRequest) GetDefaultOssZone() string`

GetDefaultOssZone returns the DefaultOssZone field if non-nil, zero value otherwise.

### GetDefaultOssZoneOk

`func (o *CoreCreateResourcePoolRequest) GetDefaultOssZoneOk() (*string, bool)`

GetDefaultOssZoneOk returns a tuple with the DefaultOssZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOssZone

`func (o *CoreCreateResourcePoolRequest) SetDefaultOssZone(v string)`

SetDefaultOssZone sets DefaultOssZone field to given value.

### HasDefaultOssZone

`func (o *CoreCreateResourcePoolRequest) HasDefaultOssZone() bool`

HasDefaultOssZone returns a boolean if a field has been set.

### GetDescription

`func (o *CoreCreateResourcePoolRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoreCreateResourcePoolRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoreCreateResourcePoolRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoreCreateResourcePoolRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CoreCreateResourcePoolRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CoreCreateResourcePoolRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CoreCreateResourcePoolRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CoreCreateResourcePoolRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOssConnectivity

`func (o *CoreCreateResourcePoolRequest) GetOssConnectivity() []string`

GetOssConnectivity returns the OssConnectivity field if non-nil, zero value otherwise.

### GetOssConnectivityOk

`func (o *CoreCreateResourcePoolRequest) GetOssConnectivityOk() (*[]string, bool)`

GetOssConnectivityOk returns a tuple with the OssConnectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOssConnectivity

`func (o *CoreCreateResourcePoolRequest) SetOssConnectivity(v []string)`

SetOssConnectivity sets OssConnectivity field to given value.

### HasOssConnectivity

`func (o *CoreCreateResourcePoolRequest) HasOssConnectivity() bool`

HasOssConnectivity returns a boolean if a field has been set.

### GetTenant

`func (o *CoreCreateResourcePoolRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CoreCreateResourcePoolRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CoreCreateResourcePoolRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CoreCreateResourcePoolRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


