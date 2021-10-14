# CoreUpdateInstanceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**CompatibleVolumeType** | Pointer to **[]string** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**InstanceClass** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**InstanceTypeGroup** | Pointer to **string** |  | [optional] 
**Memory** | Pointer to **int32** |  | [optional] 
**Zone** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCoreUpdateInstanceTypeRequest

`func NewCoreUpdateInstanceTypeRequest() *CoreUpdateInstanceTypeRequest`

NewCoreUpdateInstanceTypeRequest instantiates a new CoreUpdateInstanceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreUpdateInstanceTypeRequestWithDefaults

`func NewCoreUpdateInstanceTypeRequestWithDefaults() *CoreUpdateInstanceTypeRequest`

NewCoreUpdateInstanceTypeRequestWithDefaults instantiates a new CoreUpdateInstanceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *CoreUpdateInstanceTypeRequest) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *CoreUpdateInstanceTypeRequest) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *CoreUpdateInstanceTypeRequest) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *CoreUpdateInstanceTypeRequest) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCompatibleVolumeType

`func (o *CoreUpdateInstanceTypeRequest) GetCompatibleVolumeType() []string`

GetCompatibleVolumeType returns the CompatibleVolumeType field if non-nil, zero value otherwise.

### GetCompatibleVolumeTypeOk

`func (o *CoreUpdateInstanceTypeRequest) GetCompatibleVolumeTypeOk() (*[]string, bool)`

GetCompatibleVolumeTypeOk returns a tuple with the CompatibleVolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleVolumeType

`func (o *CoreUpdateInstanceTypeRequest) SetCompatibleVolumeType(v []string)`

SetCompatibleVolumeType sets CompatibleVolumeType field to given value.

### HasCompatibleVolumeType

`func (o *CoreUpdateInstanceTypeRequest) HasCompatibleVolumeType() bool`

HasCompatibleVolumeType returns a boolean if a field has been set.

### GetCpu

`func (o *CoreUpdateInstanceTypeRequest) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *CoreUpdateInstanceTypeRequest) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *CoreUpdateInstanceTypeRequest) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *CoreUpdateInstanceTypeRequest) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDescription

`func (o *CoreUpdateInstanceTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoreUpdateInstanceTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoreUpdateInstanceTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoreUpdateInstanceTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CoreUpdateInstanceTypeRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CoreUpdateInstanceTypeRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CoreUpdateInstanceTypeRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CoreUpdateInstanceTypeRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetInstanceClass

`func (o *CoreUpdateInstanceTypeRequest) GetInstanceClass() string`

GetInstanceClass returns the InstanceClass field if non-nil, zero value otherwise.

### GetInstanceClassOk

`func (o *CoreUpdateInstanceTypeRequest) GetInstanceClassOk() (*string, bool)`

GetInstanceClassOk returns a tuple with the InstanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceClass

`func (o *CoreUpdateInstanceTypeRequest) SetInstanceClass(v string)`

SetInstanceClass sets InstanceClass field to given value.

### HasInstanceClass

`func (o *CoreUpdateInstanceTypeRequest) HasInstanceClass() bool`

HasInstanceClass returns a boolean if a field has been set.

### GetInstanceType

`func (o *CoreUpdateInstanceTypeRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *CoreUpdateInstanceTypeRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *CoreUpdateInstanceTypeRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *CoreUpdateInstanceTypeRequest) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInstanceTypeGroup

`func (o *CoreUpdateInstanceTypeRequest) GetInstanceTypeGroup() string`

GetInstanceTypeGroup returns the InstanceTypeGroup field if non-nil, zero value otherwise.

### GetInstanceTypeGroupOk

`func (o *CoreUpdateInstanceTypeRequest) GetInstanceTypeGroupOk() (*string, bool)`

GetInstanceTypeGroupOk returns a tuple with the InstanceTypeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeGroup

`func (o *CoreUpdateInstanceTypeRequest) SetInstanceTypeGroup(v string)`

SetInstanceTypeGroup sets InstanceTypeGroup field to given value.

### HasInstanceTypeGroup

`func (o *CoreUpdateInstanceTypeRequest) HasInstanceTypeGroup() bool`

HasInstanceTypeGroup returns a boolean if a field has been set.

### GetMemory

`func (o *CoreUpdateInstanceTypeRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CoreUpdateInstanceTypeRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CoreUpdateInstanceTypeRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CoreUpdateInstanceTypeRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetZone

`func (o *CoreUpdateInstanceTypeRequest) GetZone() []string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CoreUpdateInstanceTypeRequest) GetZoneOk() (*[]string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CoreUpdateInstanceTypeRequest) SetZone(v []string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CoreUpdateInstanceTypeRequest) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


