# CoreCreateVolumeTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibleSystemVolumeType** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**IsInstanceAttached** | Pointer to **bool** |  | [optional] 
**MaxSize** | Pointer to **int32** |  | [optional] 
**MinSize** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumberOfDisk** | Pointer to **int32** |  | [optional] 
**Portable** | Pointer to **bool** |  | [optional] 
**Scalable** | Pointer to **bool** |  | [optional] 
**StepSize** | Pointer to **int32** |  | [optional] 
**SystemVolume** | Pointer to **bool** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCoreCreateVolumeTypeRequest

`func NewCoreCreateVolumeTypeRequest() *CoreCreateVolumeTypeRequest`

NewCoreCreateVolumeTypeRequest instantiates a new CoreCreateVolumeTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCreateVolumeTypeRequestWithDefaults

`func NewCoreCreateVolumeTypeRequestWithDefaults() *CoreCreateVolumeTypeRequest`

NewCoreCreateVolumeTypeRequestWithDefaults instantiates a new CoreCreateVolumeTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibleSystemVolumeType

`func (o *CoreCreateVolumeTypeRequest) GetCompatibleSystemVolumeType() []string`

GetCompatibleSystemVolumeType returns the CompatibleSystemVolumeType field if non-nil, zero value otherwise.

### GetCompatibleSystemVolumeTypeOk

`func (o *CoreCreateVolumeTypeRequest) GetCompatibleSystemVolumeTypeOk() (*[]string, bool)`

GetCompatibleSystemVolumeTypeOk returns a tuple with the CompatibleSystemVolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleSystemVolumeType

`func (o *CoreCreateVolumeTypeRequest) SetCompatibleSystemVolumeType(v []string)`

SetCompatibleSystemVolumeType sets CompatibleSystemVolumeType field to given value.

### HasCompatibleSystemVolumeType

`func (o *CoreCreateVolumeTypeRequest) HasCompatibleSystemVolumeType() bool`

HasCompatibleSystemVolumeType returns a boolean if a field has been set.

### GetDescription

`func (o *CoreCreateVolumeTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoreCreateVolumeTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoreCreateVolumeTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoreCreateVolumeTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CoreCreateVolumeTypeRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CoreCreateVolumeTypeRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CoreCreateVolumeTypeRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CoreCreateVolumeTypeRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsInstanceAttached

`func (o *CoreCreateVolumeTypeRequest) GetIsInstanceAttached() bool`

GetIsInstanceAttached returns the IsInstanceAttached field if non-nil, zero value otherwise.

### GetIsInstanceAttachedOk

`func (o *CoreCreateVolumeTypeRequest) GetIsInstanceAttachedOk() (*bool, bool)`

GetIsInstanceAttachedOk returns a tuple with the IsInstanceAttached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInstanceAttached

`func (o *CoreCreateVolumeTypeRequest) SetIsInstanceAttached(v bool)`

SetIsInstanceAttached sets IsInstanceAttached field to given value.

### HasIsInstanceAttached

`func (o *CoreCreateVolumeTypeRequest) HasIsInstanceAttached() bool`

HasIsInstanceAttached returns a boolean if a field has been set.

### GetMaxSize

`func (o *CoreCreateVolumeTypeRequest) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *CoreCreateVolumeTypeRequest) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *CoreCreateVolumeTypeRequest) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *CoreCreateVolumeTypeRequest) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetMinSize

`func (o *CoreCreateVolumeTypeRequest) GetMinSize() int32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *CoreCreateVolumeTypeRequest) GetMinSizeOk() (*int32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *CoreCreateVolumeTypeRequest) SetMinSize(v int32)`

SetMinSize sets MinSize field to given value.

### HasMinSize

`func (o *CoreCreateVolumeTypeRequest) HasMinSize() bool`

HasMinSize returns a boolean if a field has been set.

### GetName

`func (o *CoreCreateVolumeTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreCreateVolumeTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreCreateVolumeTypeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoreCreateVolumeTypeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfDisk

`func (o *CoreCreateVolumeTypeRequest) GetNumberOfDisk() int32`

GetNumberOfDisk returns the NumberOfDisk field if non-nil, zero value otherwise.

### GetNumberOfDiskOk

`func (o *CoreCreateVolumeTypeRequest) GetNumberOfDiskOk() (*int32, bool)`

GetNumberOfDiskOk returns a tuple with the NumberOfDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDisk

`func (o *CoreCreateVolumeTypeRequest) SetNumberOfDisk(v int32)`

SetNumberOfDisk sets NumberOfDisk field to given value.

### HasNumberOfDisk

`func (o *CoreCreateVolumeTypeRequest) HasNumberOfDisk() bool`

HasNumberOfDisk returns a boolean if a field has been set.

### GetPortable

`func (o *CoreCreateVolumeTypeRequest) GetPortable() bool`

GetPortable returns the Portable field if non-nil, zero value otherwise.

### GetPortableOk

`func (o *CoreCreateVolumeTypeRequest) GetPortableOk() (*bool, bool)`

GetPortableOk returns a tuple with the Portable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortable

`func (o *CoreCreateVolumeTypeRequest) SetPortable(v bool)`

SetPortable sets Portable field to given value.

### HasPortable

`func (o *CoreCreateVolumeTypeRequest) HasPortable() bool`

HasPortable returns a boolean if a field has been set.

### GetScalable

`func (o *CoreCreateVolumeTypeRequest) GetScalable() bool`

GetScalable returns the Scalable field if non-nil, zero value otherwise.

### GetScalableOk

`func (o *CoreCreateVolumeTypeRequest) GetScalableOk() (*bool, bool)`

GetScalableOk returns a tuple with the Scalable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalable

`func (o *CoreCreateVolumeTypeRequest) SetScalable(v bool)`

SetScalable sets Scalable field to given value.

### HasScalable

`func (o *CoreCreateVolumeTypeRequest) HasScalable() bool`

HasScalable returns a boolean if a field has been set.

### GetStepSize

`func (o *CoreCreateVolumeTypeRequest) GetStepSize() int32`

GetStepSize returns the StepSize field if non-nil, zero value otherwise.

### GetStepSizeOk

`func (o *CoreCreateVolumeTypeRequest) GetStepSizeOk() (*int32, bool)`

GetStepSizeOk returns a tuple with the StepSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSize

`func (o *CoreCreateVolumeTypeRequest) SetStepSize(v int32)`

SetStepSize sets StepSize field to given value.

### HasStepSize

`func (o *CoreCreateVolumeTypeRequest) HasStepSize() bool`

HasStepSize returns a boolean if a field has been set.

### GetSystemVolume

`func (o *CoreCreateVolumeTypeRequest) GetSystemVolume() bool`

GetSystemVolume returns the SystemVolume field if non-nil, zero value otherwise.

### GetSystemVolumeOk

`func (o *CoreCreateVolumeTypeRequest) GetSystemVolumeOk() (*bool, bool)`

GetSystemVolumeOk returns a tuple with the SystemVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemVolume

`func (o *CoreCreateVolumeTypeRequest) SetSystemVolume(v bool)`

SetSystemVolume sets SystemVolume field to given value.

### HasSystemVolume

`func (o *CoreCreateVolumeTypeRequest) HasSystemVolume() bool`

HasSystemVolume returns a boolean if a field has been set.

### GetVendor

`func (o *CoreCreateVolumeTypeRequest) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CoreCreateVolumeTypeRequest) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CoreCreateVolumeTypeRequest) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CoreCreateVolumeTypeRequest) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVolumeType

`func (o *CoreCreateVolumeTypeRequest) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CoreCreateVolumeTypeRequest) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CoreCreateVolumeTypeRequest) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *CoreCreateVolumeTypeRequest) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetZone

`func (o *CoreCreateVolumeTypeRequest) GetZone() []string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CoreCreateVolumeTypeRequest) GetZoneOk() (*[]string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CoreCreateVolumeTypeRequest) SetZone(v []string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CoreCreateVolumeTypeRequest) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


