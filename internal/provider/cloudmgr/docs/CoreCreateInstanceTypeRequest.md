# CoreCreateInstanceTypeRequest

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
**Name** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCoreCreateInstanceTypeRequest

`func NewCoreCreateInstanceTypeRequest() *CoreCreateInstanceTypeRequest`

NewCoreCreateInstanceTypeRequest instantiates a new CoreCreateInstanceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCreateInstanceTypeRequestWithDefaults

`func NewCoreCreateInstanceTypeRequestWithDefaults() *CoreCreateInstanceTypeRequest`

NewCoreCreateInstanceTypeRequestWithDefaults instantiates a new CoreCreateInstanceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *CoreCreateInstanceTypeRequest) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *CoreCreateInstanceTypeRequest) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *CoreCreateInstanceTypeRequest) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *CoreCreateInstanceTypeRequest) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCompatibleVolumeType

`func (o *CoreCreateInstanceTypeRequest) GetCompatibleVolumeType() []string`

GetCompatibleVolumeType returns the CompatibleVolumeType field if non-nil, zero value otherwise.

### GetCompatibleVolumeTypeOk

`func (o *CoreCreateInstanceTypeRequest) GetCompatibleVolumeTypeOk() (*[]string, bool)`

GetCompatibleVolumeTypeOk returns a tuple with the CompatibleVolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleVolumeType

`func (o *CoreCreateInstanceTypeRequest) SetCompatibleVolumeType(v []string)`

SetCompatibleVolumeType sets CompatibleVolumeType field to given value.

### HasCompatibleVolumeType

`func (o *CoreCreateInstanceTypeRequest) HasCompatibleVolumeType() bool`

HasCompatibleVolumeType returns a boolean if a field has been set.

### GetCpu

`func (o *CoreCreateInstanceTypeRequest) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *CoreCreateInstanceTypeRequest) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *CoreCreateInstanceTypeRequest) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *CoreCreateInstanceTypeRequest) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDescription

`func (o *CoreCreateInstanceTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoreCreateInstanceTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoreCreateInstanceTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoreCreateInstanceTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CoreCreateInstanceTypeRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CoreCreateInstanceTypeRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CoreCreateInstanceTypeRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CoreCreateInstanceTypeRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetInstanceClass

`func (o *CoreCreateInstanceTypeRequest) GetInstanceClass() string`

GetInstanceClass returns the InstanceClass field if non-nil, zero value otherwise.

### GetInstanceClassOk

`func (o *CoreCreateInstanceTypeRequest) GetInstanceClassOk() (*string, bool)`

GetInstanceClassOk returns a tuple with the InstanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceClass

`func (o *CoreCreateInstanceTypeRequest) SetInstanceClass(v string)`

SetInstanceClass sets InstanceClass field to given value.

### HasInstanceClass

`func (o *CoreCreateInstanceTypeRequest) HasInstanceClass() bool`

HasInstanceClass returns a boolean if a field has been set.

### GetInstanceType

`func (o *CoreCreateInstanceTypeRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *CoreCreateInstanceTypeRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *CoreCreateInstanceTypeRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *CoreCreateInstanceTypeRequest) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInstanceTypeGroup

`func (o *CoreCreateInstanceTypeRequest) GetInstanceTypeGroup() string`

GetInstanceTypeGroup returns the InstanceTypeGroup field if non-nil, zero value otherwise.

### GetInstanceTypeGroupOk

`func (o *CoreCreateInstanceTypeRequest) GetInstanceTypeGroupOk() (*string, bool)`

GetInstanceTypeGroupOk returns a tuple with the InstanceTypeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeGroup

`func (o *CoreCreateInstanceTypeRequest) SetInstanceTypeGroup(v string)`

SetInstanceTypeGroup sets InstanceTypeGroup field to given value.

### HasInstanceTypeGroup

`func (o *CoreCreateInstanceTypeRequest) HasInstanceTypeGroup() bool`

HasInstanceTypeGroup returns a boolean if a field has been set.

### GetMemory

`func (o *CoreCreateInstanceTypeRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CoreCreateInstanceTypeRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CoreCreateInstanceTypeRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CoreCreateInstanceTypeRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetName

`func (o *CoreCreateInstanceTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreCreateInstanceTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreCreateInstanceTypeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoreCreateInstanceTypeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVendor

`func (o *CoreCreateInstanceTypeRequest) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CoreCreateInstanceTypeRequest) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CoreCreateInstanceTypeRequest) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CoreCreateInstanceTypeRequest) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetZone

`func (o *CoreCreateInstanceTypeRequest) GetZone() []string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CoreCreateInstanceTypeRequest) GetZoneOk() (*[]string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CoreCreateInstanceTypeRequest) SetZone(v []string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CoreCreateInstanceTypeRequest) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


