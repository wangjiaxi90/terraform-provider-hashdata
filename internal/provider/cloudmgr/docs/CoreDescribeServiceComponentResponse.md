# CoreDescribeServiceComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxBandwidth** | Pointer to **int32** |  | [optional] 
**MaxCpuCount** | Pointer to **int32** |  | [optional] 
**MaxInitialInstance** | Pointer to **int32** |  | [optional] 
**MaxInstance** | Pointer to **int32** |  | [optional] 
**MaxMemoryMb** | Pointer to **int32** |  | [optional] 
**MaxVolumeSize** | Pointer to **int32** |  | [optional] 
**MinBandwidth** | Pointer to **int32** |  | [optional] 
**MinCpuCount** | Pointer to **int32** |  | [optional] 
**MinInitialInstance** | Pointer to **int32** |  | [optional] 
**MinInstance** | Pointer to **int32** |  | [optional] 
**MinMemoryMb** | Pointer to **int32** |  | [optional] 
**MinVolumeSize** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RequireDataVolume** | Pointer to **bool** |  | [optional] 
**RequirePortableVolume** | Pointer to **bool** |  | [optional] 

## Methods

### NewCoreDescribeServiceComponentResponse

`func NewCoreDescribeServiceComponentResponse() *CoreDescribeServiceComponentResponse`

NewCoreDescribeServiceComponentResponse instantiates a new CoreDescribeServiceComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeServiceComponentResponseWithDefaults

`func NewCoreDescribeServiceComponentResponseWithDefaults() *CoreDescribeServiceComponentResponse`

NewCoreDescribeServiceComponentResponseWithDefaults instantiates a new CoreDescribeServiceComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxBandwidth

`func (o *CoreDescribeServiceComponentResponse) GetMaxBandwidth() int32`

GetMaxBandwidth returns the MaxBandwidth field if non-nil, zero value otherwise.

### GetMaxBandwidthOk

`func (o *CoreDescribeServiceComponentResponse) GetMaxBandwidthOk() (*int32, bool)`

GetMaxBandwidthOk returns a tuple with the MaxBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBandwidth

`func (o *CoreDescribeServiceComponentResponse) SetMaxBandwidth(v int32)`

SetMaxBandwidth sets MaxBandwidth field to given value.

### HasMaxBandwidth

`func (o *CoreDescribeServiceComponentResponse) HasMaxBandwidth() bool`

HasMaxBandwidth returns a boolean if a field has been set.

### GetMaxCpuCount

`func (o *CoreDescribeServiceComponentResponse) GetMaxCpuCount() int32`

GetMaxCpuCount returns the MaxCpuCount field if non-nil, zero value otherwise.

### GetMaxCpuCountOk

`func (o *CoreDescribeServiceComponentResponse) GetMaxCpuCountOk() (*int32, bool)`

GetMaxCpuCountOk returns a tuple with the MaxCpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpuCount

`func (o *CoreDescribeServiceComponentResponse) SetMaxCpuCount(v int32)`

SetMaxCpuCount sets MaxCpuCount field to given value.

### HasMaxCpuCount

`func (o *CoreDescribeServiceComponentResponse) HasMaxCpuCount() bool`

HasMaxCpuCount returns a boolean if a field has been set.

### GetMaxInitialInstance

`func (o *CoreDescribeServiceComponentResponse) GetMaxInitialInstance() int32`

GetMaxInitialInstance returns the MaxInitialInstance field if non-nil, zero value otherwise.

### GetMaxInitialInstanceOk

`func (o *CoreDescribeServiceComponentResponse) GetMaxInitialInstanceOk() (*int32, bool)`

GetMaxInitialInstanceOk returns a tuple with the MaxInitialInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInitialInstance

`func (o *CoreDescribeServiceComponentResponse) SetMaxInitialInstance(v int32)`

SetMaxInitialInstance sets MaxInitialInstance field to given value.

### HasMaxInitialInstance

`func (o *CoreDescribeServiceComponentResponse) HasMaxInitialInstance() bool`

HasMaxInitialInstance returns a boolean if a field has been set.

### GetMaxInstance

`func (o *CoreDescribeServiceComponentResponse) GetMaxInstance() int32`

GetMaxInstance returns the MaxInstance field if non-nil, zero value otherwise.

### GetMaxInstanceOk

`func (o *CoreDescribeServiceComponentResponse) GetMaxInstanceOk() (*int32, bool)`

GetMaxInstanceOk returns a tuple with the MaxInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstance

`func (o *CoreDescribeServiceComponentResponse) SetMaxInstance(v int32)`

SetMaxInstance sets MaxInstance field to given value.

### HasMaxInstance

`func (o *CoreDescribeServiceComponentResponse) HasMaxInstance() bool`

HasMaxInstance returns a boolean if a field has been set.

### GetMaxMemoryMb

`func (o *CoreDescribeServiceComponentResponse) GetMaxMemoryMb() int32`

GetMaxMemoryMb returns the MaxMemoryMb field if non-nil, zero value otherwise.

### GetMaxMemoryMbOk

`func (o *CoreDescribeServiceComponentResponse) GetMaxMemoryMbOk() (*int32, bool)`

GetMaxMemoryMbOk returns a tuple with the MaxMemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryMb

`func (o *CoreDescribeServiceComponentResponse) SetMaxMemoryMb(v int32)`

SetMaxMemoryMb sets MaxMemoryMb field to given value.

### HasMaxMemoryMb

`func (o *CoreDescribeServiceComponentResponse) HasMaxMemoryMb() bool`

HasMaxMemoryMb returns a boolean if a field has been set.

### GetMaxVolumeSize

`func (o *CoreDescribeServiceComponentResponse) GetMaxVolumeSize() int32`

GetMaxVolumeSize returns the MaxVolumeSize field if non-nil, zero value otherwise.

### GetMaxVolumeSizeOk

`func (o *CoreDescribeServiceComponentResponse) GetMaxVolumeSizeOk() (*int32, bool)`

GetMaxVolumeSizeOk returns a tuple with the MaxVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVolumeSize

`func (o *CoreDescribeServiceComponentResponse) SetMaxVolumeSize(v int32)`

SetMaxVolumeSize sets MaxVolumeSize field to given value.

### HasMaxVolumeSize

`func (o *CoreDescribeServiceComponentResponse) HasMaxVolumeSize() bool`

HasMaxVolumeSize returns a boolean if a field has been set.

### GetMinBandwidth

`func (o *CoreDescribeServiceComponentResponse) GetMinBandwidth() int32`

GetMinBandwidth returns the MinBandwidth field if non-nil, zero value otherwise.

### GetMinBandwidthOk

`func (o *CoreDescribeServiceComponentResponse) GetMinBandwidthOk() (*int32, bool)`

GetMinBandwidthOk returns a tuple with the MinBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBandwidth

`func (o *CoreDescribeServiceComponentResponse) SetMinBandwidth(v int32)`

SetMinBandwidth sets MinBandwidth field to given value.

### HasMinBandwidth

`func (o *CoreDescribeServiceComponentResponse) HasMinBandwidth() bool`

HasMinBandwidth returns a boolean if a field has been set.

### GetMinCpuCount

`func (o *CoreDescribeServiceComponentResponse) GetMinCpuCount() int32`

GetMinCpuCount returns the MinCpuCount field if non-nil, zero value otherwise.

### GetMinCpuCountOk

`func (o *CoreDescribeServiceComponentResponse) GetMinCpuCountOk() (*int32, bool)`

GetMinCpuCountOk returns a tuple with the MinCpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpuCount

`func (o *CoreDescribeServiceComponentResponse) SetMinCpuCount(v int32)`

SetMinCpuCount sets MinCpuCount field to given value.

### HasMinCpuCount

`func (o *CoreDescribeServiceComponentResponse) HasMinCpuCount() bool`

HasMinCpuCount returns a boolean if a field has been set.

### GetMinInitialInstance

`func (o *CoreDescribeServiceComponentResponse) GetMinInitialInstance() int32`

GetMinInitialInstance returns the MinInitialInstance field if non-nil, zero value otherwise.

### GetMinInitialInstanceOk

`func (o *CoreDescribeServiceComponentResponse) GetMinInitialInstanceOk() (*int32, bool)`

GetMinInitialInstanceOk returns a tuple with the MinInitialInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInitialInstance

`func (o *CoreDescribeServiceComponentResponse) SetMinInitialInstance(v int32)`

SetMinInitialInstance sets MinInitialInstance field to given value.

### HasMinInitialInstance

`func (o *CoreDescribeServiceComponentResponse) HasMinInitialInstance() bool`

HasMinInitialInstance returns a boolean if a field has been set.

### GetMinInstance

`func (o *CoreDescribeServiceComponentResponse) GetMinInstance() int32`

GetMinInstance returns the MinInstance field if non-nil, zero value otherwise.

### GetMinInstanceOk

`func (o *CoreDescribeServiceComponentResponse) GetMinInstanceOk() (*int32, bool)`

GetMinInstanceOk returns a tuple with the MinInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInstance

`func (o *CoreDescribeServiceComponentResponse) SetMinInstance(v int32)`

SetMinInstance sets MinInstance field to given value.

### HasMinInstance

`func (o *CoreDescribeServiceComponentResponse) HasMinInstance() bool`

HasMinInstance returns a boolean if a field has been set.

### GetMinMemoryMb

`func (o *CoreDescribeServiceComponentResponse) GetMinMemoryMb() int32`

GetMinMemoryMb returns the MinMemoryMb field if non-nil, zero value otherwise.

### GetMinMemoryMbOk

`func (o *CoreDescribeServiceComponentResponse) GetMinMemoryMbOk() (*int32, bool)`

GetMinMemoryMbOk returns a tuple with the MinMemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemoryMb

`func (o *CoreDescribeServiceComponentResponse) SetMinMemoryMb(v int32)`

SetMinMemoryMb sets MinMemoryMb field to given value.

### HasMinMemoryMb

`func (o *CoreDescribeServiceComponentResponse) HasMinMemoryMb() bool`

HasMinMemoryMb returns a boolean if a field has been set.

### GetMinVolumeSize

`func (o *CoreDescribeServiceComponentResponse) GetMinVolumeSize() int32`

GetMinVolumeSize returns the MinVolumeSize field if non-nil, zero value otherwise.

### GetMinVolumeSizeOk

`func (o *CoreDescribeServiceComponentResponse) GetMinVolumeSizeOk() (*int32, bool)`

GetMinVolumeSizeOk returns a tuple with the MinVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVolumeSize

`func (o *CoreDescribeServiceComponentResponse) SetMinVolumeSize(v int32)`

SetMinVolumeSize sets MinVolumeSize field to given value.

### HasMinVolumeSize

`func (o *CoreDescribeServiceComponentResponse) HasMinVolumeSize() bool`

HasMinVolumeSize returns a boolean if a field has been set.

### GetName

`func (o *CoreDescribeServiceComponentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreDescribeServiceComponentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreDescribeServiceComponentResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoreDescribeServiceComponentResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequireDataVolume

`func (o *CoreDescribeServiceComponentResponse) GetRequireDataVolume() bool`

GetRequireDataVolume returns the RequireDataVolume field if non-nil, zero value otherwise.

### GetRequireDataVolumeOk

`func (o *CoreDescribeServiceComponentResponse) GetRequireDataVolumeOk() (*bool, bool)`

GetRequireDataVolumeOk returns a tuple with the RequireDataVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireDataVolume

`func (o *CoreDescribeServiceComponentResponse) SetRequireDataVolume(v bool)`

SetRequireDataVolume sets RequireDataVolume field to given value.

### HasRequireDataVolume

`func (o *CoreDescribeServiceComponentResponse) HasRequireDataVolume() bool`

HasRequireDataVolume returns a boolean if a field has been set.

### GetRequirePortableVolume

`func (o *CoreDescribeServiceComponentResponse) GetRequirePortableVolume() bool`

GetRequirePortableVolume returns the RequirePortableVolume field if non-nil, zero value otherwise.

### GetRequirePortableVolumeOk

`func (o *CoreDescribeServiceComponentResponse) GetRequirePortableVolumeOk() (*bool, bool)`

GetRequirePortableVolumeOk returns a tuple with the RequirePortableVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePortableVolume

`func (o *CoreDescribeServiceComponentResponse) SetRequirePortableVolume(v bool)`

SetRequirePortableVolume sets RequirePortableVolume field to given value.

### HasRequirePortableVolume

`func (o *CoreDescribeServiceComponentResponse) HasRequirePortableVolume() bool`

HasRequirePortableVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


