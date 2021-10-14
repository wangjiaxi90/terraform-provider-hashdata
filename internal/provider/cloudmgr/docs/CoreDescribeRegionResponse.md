# CoreDescribeRegionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**DefaultOssZone** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OssConnectivity** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreDescribeRegionResponse

`func NewCoreDescribeRegionResponse() *CoreDescribeRegionResponse`

NewCoreDescribeRegionResponse instantiates a new CoreDescribeRegionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeRegionResponseWithDefaults

`func NewCoreDescribeRegionResponseWithDefaults() *CoreDescribeRegionResponse`

NewCoreDescribeRegionResponseWithDefaults instantiates a new CoreDescribeRegionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CoreDescribeRegionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CoreDescribeRegionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CoreDescribeRegionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CoreDescribeRegionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultOssZone

`func (o *CoreDescribeRegionResponse) GetDefaultOssZone() string`

GetDefaultOssZone returns the DefaultOssZone field if non-nil, zero value otherwise.

### GetDefaultOssZoneOk

`func (o *CoreDescribeRegionResponse) GetDefaultOssZoneOk() (*string, bool)`

GetDefaultOssZoneOk returns a tuple with the DefaultOssZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOssZone

`func (o *CoreDescribeRegionResponse) SetDefaultOssZone(v string)`

SetDefaultOssZone sets DefaultOssZone field to given value.

### HasDefaultOssZone

`func (o *CoreDescribeRegionResponse) HasDefaultOssZone() bool`

HasDefaultOssZone returns a boolean if a field has been set.

### GetDisplayName

`func (o *CoreDescribeRegionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CoreDescribeRegionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CoreDescribeRegionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CoreDescribeRegionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEndpoint

`func (o *CoreDescribeRegionResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CoreDescribeRegionResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CoreDescribeRegionResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *CoreDescribeRegionResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetId

`func (o *CoreDescribeRegionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreDescribeRegionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreDescribeRegionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CoreDescribeRegionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CoreDescribeRegionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreDescribeRegionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreDescribeRegionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoreDescribeRegionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOssConnectivity

`func (o *CoreDescribeRegionResponse) GetOssConnectivity() []string`

GetOssConnectivity returns the OssConnectivity field if non-nil, zero value otherwise.

### GetOssConnectivityOk

`func (o *CoreDescribeRegionResponse) GetOssConnectivityOk() (*[]string, bool)`

GetOssConnectivityOk returns a tuple with the OssConnectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOssConnectivity

`func (o *CoreDescribeRegionResponse) SetOssConnectivity(v []string)`

SetOssConnectivity sets OssConnectivity field to given value.

### HasOssConnectivity

`func (o *CoreDescribeRegionResponse) HasOssConnectivity() bool`

HasOssConnectivity returns a boolean if a field has been set.

### GetProperties

`func (o *CoreDescribeRegionResponse) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CoreDescribeRegionResponse) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CoreDescribeRegionResponse) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CoreDescribeRegionResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRegion

`func (o *CoreDescribeRegionResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CoreDescribeRegionResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CoreDescribeRegionResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CoreDescribeRegionResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVendor

`func (o *CoreDescribeRegionResponse) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CoreDescribeRegionResponse) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CoreDescribeRegionResponse) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CoreDescribeRegionResponse) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


