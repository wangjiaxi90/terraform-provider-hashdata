# CoreDescribeVolumeTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibleSystemVolumeType** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
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

### NewCoreDescribeVolumeTypeResponse

`func NewCoreDescribeVolumeTypeResponse() *CoreDescribeVolumeTypeResponse`

NewCoreDescribeVolumeTypeResponse instantiates a new CoreDescribeVolumeTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeVolumeTypeResponseWithDefaults

`func NewCoreDescribeVolumeTypeResponseWithDefaults() *CoreDescribeVolumeTypeResponse`

NewCoreDescribeVolumeTypeResponseWithDefaults instantiates a new CoreDescribeVolumeTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibleSystemVolumeType

`func (o *CoreDescribeVolumeTypeResponse) GetCompatibleSystemVolumeType() []string`

GetCompatibleSystemVolumeType returns the CompatibleSystemVolumeType field if non-nil, zero value otherwise.

### GetCompatibleSystemVolumeTypeOk

`func (o *CoreDescribeVolumeTypeResponse) GetCompatibleSystemVolumeTypeOk() (*[]string, bool)`

GetCompatibleSystemVolumeTypeOk returns a tuple with the CompatibleSystemVolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleSystemVolumeType

`func (o *CoreDescribeVolumeTypeResponse) SetCompatibleSystemVolumeType(v []string)`

SetCompatibleSystemVolumeType sets CompatibleSystemVolumeType field to given value.

### HasCompatibleSystemVolumeType

`func (o *CoreDescribeVolumeTypeResponse) HasCompatibleSystemVolumeType() bool`

HasCompatibleSystemVolumeType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CoreDescribeVolumeTypeResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CoreDescribeVolumeTypeResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CoreDescribeVolumeTypeResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CoreDescribeVolumeTypeResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CoreDescribeVolumeTypeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoreDescribeVolumeTypeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoreDescribeVolumeTypeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoreDescribeVolumeTypeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CoreDescribeVolumeTypeResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CoreDescribeVolumeTypeResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CoreDescribeVolumeTypeResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CoreDescribeVolumeTypeResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *CoreDescribeVolumeTypeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreDescribeVolumeTypeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreDescribeVolumeTypeResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CoreDescribeVolumeTypeResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsInstanceAttached

`func (o *CoreDescribeVolumeTypeResponse) GetIsInstanceAttached() bool`

GetIsInstanceAttached returns the IsInstanceAttached field if non-nil, zero value otherwise.

### GetIsInstanceAttachedOk

`func (o *CoreDescribeVolumeTypeResponse) GetIsInstanceAttachedOk() (*bool, bool)`

GetIsInstanceAttachedOk returns a tuple with the IsInstanceAttached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInstanceAttached

`func (o *CoreDescribeVolumeTypeResponse) SetIsInstanceAttached(v bool)`

SetIsInstanceAttached sets IsInstanceAttached field to given value.

### HasIsInstanceAttached

`func (o *CoreDescribeVolumeTypeResponse) HasIsInstanceAttached() bool`

HasIsInstanceAttached returns a boolean if a field has been set.

### GetMaxSize

`func (o *CoreDescribeVolumeTypeResponse) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *CoreDescribeVolumeTypeResponse) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *CoreDescribeVolumeTypeResponse) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *CoreDescribeVolumeTypeResponse) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetMinSize

`func (o *CoreDescribeVolumeTypeResponse) GetMinSize() int32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *CoreDescribeVolumeTypeResponse) GetMinSizeOk() (*int32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *CoreDescribeVolumeTypeResponse) SetMinSize(v int32)`

SetMinSize sets MinSize field to given value.

### HasMinSize

`func (o *CoreDescribeVolumeTypeResponse) HasMinSize() bool`

HasMinSize returns a boolean if a field has been set.

### GetName

`func (o *CoreDescribeVolumeTypeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreDescribeVolumeTypeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreDescribeVolumeTypeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoreDescribeVolumeTypeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfDisk

`func (o *CoreDescribeVolumeTypeResponse) GetNumberOfDisk() int32`

GetNumberOfDisk returns the NumberOfDisk field if non-nil, zero value otherwise.

### GetNumberOfDiskOk

`func (o *CoreDescribeVolumeTypeResponse) GetNumberOfDiskOk() (*int32, bool)`

GetNumberOfDiskOk returns a tuple with the NumberOfDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDisk

`func (o *CoreDescribeVolumeTypeResponse) SetNumberOfDisk(v int32)`

SetNumberOfDisk sets NumberOfDisk field to given value.

### HasNumberOfDisk

`func (o *CoreDescribeVolumeTypeResponse) HasNumberOfDisk() bool`

HasNumberOfDisk returns a boolean if a field has been set.

### GetPortable

`func (o *CoreDescribeVolumeTypeResponse) GetPortable() bool`

GetPortable returns the Portable field if non-nil, zero value otherwise.

### GetPortableOk

`func (o *CoreDescribeVolumeTypeResponse) GetPortableOk() (*bool, bool)`

GetPortableOk returns a tuple with the Portable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortable

`func (o *CoreDescribeVolumeTypeResponse) SetPortable(v bool)`

SetPortable sets Portable field to given value.

### HasPortable

`func (o *CoreDescribeVolumeTypeResponse) HasPortable() bool`

HasPortable returns a boolean if a field has been set.

### GetScalable

`func (o *CoreDescribeVolumeTypeResponse) GetScalable() bool`

GetScalable returns the Scalable field if non-nil, zero value otherwise.

### GetScalableOk

`func (o *CoreDescribeVolumeTypeResponse) GetScalableOk() (*bool, bool)`

GetScalableOk returns a tuple with the Scalable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalable

`func (o *CoreDescribeVolumeTypeResponse) SetScalable(v bool)`

SetScalable sets Scalable field to given value.

### HasScalable

`func (o *CoreDescribeVolumeTypeResponse) HasScalable() bool`

HasScalable returns a boolean if a field has been set.

### GetStepSize

`func (o *CoreDescribeVolumeTypeResponse) GetStepSize() int32`

GetStepSize returns the StepSize field if non-nil, zero value otherwise.

### GetStepSizeOk

`func (o *CoreDescribeVolumeTypeResponse) GetStepSizeOk() (*int32, bool)`

GetStepSizeOk returns a tuple with the StepSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSize

`func (o *CoreDescribeVolumeTypeResponse) SetStepSize(v int32)`

SetStepSize sets StepSize field to given value.

### HasStepSize

`func (o *CoreDescribeVolumeTypeResponse) HasStepSize() bool`

HasStepSize returns a boolean if a field has been set.

### GetSystemVolume

`func (o *CoreDescribeVolumeTypeResponse) GetSystemVolume() bool`

GetSystemVolume returns the SystemVolume field if non-nil, zero value otherwise.

### GetSystemVolumeOk

`func (o *CoreDescribeVolumeTypeResponse) GetSystemVolumeOk() (*bool, bool)`

GetSystemVolumeOk returns a tuple with the SystemVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemVolume

`func (o *CoreDescribeVolumeTypeResponse) SetSystemVolume(v bool)`

SetSystemVolume sets SystemVolume field to given value.

### HasSystemVolume

`func (o *CoreDescribeVolumeTypeResponse) HasSystemVolume() bool`

HasSystemVolume returns a boolean if a field has been set.

### GetVendor

`func (o *CoreDescribeVolumeTypeResponse) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CoreDescribeVolumeTypeResponse) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CoreDescribeVolumeTypeResponse) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CoreDescribeVolumeTypeResponse) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVolumeType

`func (o *CoreDescribeVolumeTypeResponse) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CoreDescribeVolumeTypeResponse) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CoreDescribeVolumeTypeResponse) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *CoreDescribeVolumeTypeResponse) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetZone

`func (o *CoreDescribeVolumeTypeResponse) GetZone() []string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CoreDescribeVolumeTypeResponse) GetZoneOk() (*[]string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CoreDescribeVolumeTypeResponse) SetZone(v []string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CoreDescribeVolumeTypeResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


