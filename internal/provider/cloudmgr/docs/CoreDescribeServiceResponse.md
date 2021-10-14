# CoreDescribeServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeChannel** | Pointer to **string** |  | [optional] 
**ChargeType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DestroyedAt** | Pointer to **string** |  | [optional] 
**DisableAutoRecovery** | Pointer to **bool** |  | [optional] 
**EnableAlert** | Pointer to **bool** |  | [optional] 
**EnableInternalCharge** | Pointer to **bool** |  | [optional] 
**Expired** | Pointer to **bool** |  | [optional] 
**ExpiredAt** | Pointer to **string** |  | [optional] 
**HealthStatus** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**ProtectedMode** | Pointer to **bool** |  | [optional] 
**Recovering** | Pointer to **bool** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ResourcePool** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreDescribeServiceResponse

`func NewCoreDescribeServiceResponse() *CoreDescribeServiceResponse`

NewCoreDescribeServiceResponse instantiates a new CoreDescribeServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeServiceResponseWithDefaults

`func NewCoreDescribeServiceResponseWithDefaults() *CoreDescribeServiceResponse`

NewCoreDescribeServiceResponseWithDefaults instantiates a new CoreDescribeServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeChannel

`func (o *CoreDescribeServiceResponse) GetChargeChannel() string`

GetChargeChannel returns the ChargeChannel field if non-nil, zero value otherwise.

### GetChargeChannelOk

`func (o *CoreDescribeServiceResponse) GetChargeChannelOk() (*string, bool)`

GetChargeChannelOk returns a tuple with the ChargeChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeChannel

`func (o *CoreDescribeServiceResponse) SetChargeChannel(v string)`

SetChargeChannel sets ChargeChannel field to given value.

### HasChargeChannel

`func (o *CoreDescribeServiceResponse) HasChargeChannel() bool`

HasChargeChannel returns a boolean if a field has been set.

### GetChargeType

`func (o *CoreDescribeServiceResponse) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *CoreDescribeServiceResponse) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *CoreDescribeServiceResponse) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *CoreDescribeServiceResponse) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CoreDescribeServiceResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CoreDescribeServiceResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CoreDescribeServiceResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CoreDescribeServiceResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CoreDescribeServiceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoreDescribeServiceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoreDescribeServiceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoreDescribeServiceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestroyedAt

`func (o *CoreDescribeServiceResponse) GetDestroyedAt() string`

GetDestroyedAt returns the DestroyedAt field if non-nil, zero value otherwise.

### GetDestroyedAtOk

`func (o *CoreDescribeServiceResponse) GetDestroyedAtOk() (*string, bool)`

GetDestroyedAtOk returns a tuple with the DestroyedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyedAt

`func (o *CoreDescribeServiceResponse) SetDestroyedAt(v string)`

SetDestroyedAt sets DestroyedAt field to given value.

### HasDestroyedAt

`func (o *CoreDescribeServiceResponse) HasDestroyedAt() bool`

HasDestroyedAt returns a boolean if a field has been set.

### GetDisableAutoRecovery

`func (o *CoreDescribeServiceResponse) GetDisableAutoRecovery() bool`

GetDisableAutoRecovery returns the DisableAutoRecovery field if non-nil, zero value otherwise.

### GetDisableAutoRecoveryOk

`func (o *CoreDescribeServiceResponse) GetDisableAutoRecoveryOk() (*bool, bool)`

GetDisableAutoRecoveryOk returns a tuple with the DisableAutoRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoRecovery

`func (o *CoreDescribeServiceResponse) SetDisableAutoRecovery(v bool)`

SetDisableAutoRecovery sets DisableAutoRecovery field to given value.

### HasDisableAutoRecovery

`func (o *CoreDescribeServiceResponse) HasDisableAutoRecovery() bool`

HasDisableAutoRecovery returns a boolean if a field has been set.

### GetEnableAlert

`func (o *CoreDescribeServiceResponse) GetEnableAlert() bool`

GetEnableAlert returns the EnableAlert field if non-nil, zero value otherwise.

### GetEnableAlertOk

`func (o *CoreDescribeServiceResponse) GetEnableAlertOk() (*bool, bool)`

GetEnableAlertOk returns a tuple with the EnableAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAlert

`func (o *CoreDescribeServiceResponse) SetEnableAlert(v bool)`

SetEnableAlert sets EnableAlert field to given value.

### HasEnableAlert

`func (o *CoreDescribeServiceResponse) HasEnableAlert() bool`

HasEnableAlert returns a boolean if a field has been set.

### GetEnableInternalCharge

`func (o *CoreDescribeServiceResponse) GetEnableInternalCharge() bool`

GetEnableInternalCharge returns the EnableInternalCharge field if non-nil, zero value otherwise.

### GetEnableInternalChargeOk

`func (o *CoreDescribeServiceResponse) GetEnableInternalChargeOk() (*bool, bool)`

GetEnableInternalChargeOk returns a tuple with the EnableInternalCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternalCharge

`func (o *CoreDescribeServiceResponse) SetEnableInternalCharge(v bool)`

SetEnableInternalCharge sets EnableInternalCharge field to given value.

### HasEnableInternalCharge

`func (o *CoreDescribeServiceResponse) HasEnableInternalCharge() bool`

HasEnableInternalCharge returns a boolean if a field has been set.

### GetExpired

`func (o *CoreDescribeServiceResponse) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *CoreDescribeServiceResponse) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *CoreDescribeServiceResponse) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *CoreDescribeServiceResponse) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetExpiredAt

`func (o *CoreDescribeServiceResponse) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *CoreDescribeServiceResponse) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *CoreDescribeServiceResponse) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *CoreDescribeServiceResponse) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetHealthStatus

`func (o *CoreDescribeServiceResponse) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *CoreDescribeServiceResponse) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *CoreDescribeServiceResponse) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *CoreDescribeServiceResponse) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetId

`func (o *CoreDescribeServiceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreDescribeServiceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreDescribeServiceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CoreDescribeServiceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceType

`func (o *CoreDescribeServiceResponse) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *CoreDescribeServiceResponse) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *CoreDescribeServiceResponse) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *CoreDescribeServiceResponse) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetName

`func (o *CoreDescribeServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreDescribeServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreDescribeServiceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoreDescribeServiceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *CoreDescribeServiceResponse) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CoreDescribeServiceResponse) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CoreDescribeServiceResponse) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CoreDescribeServiceResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetProtectedMode

`func (o *CoreDescribeServiceResponse) GetProtectedMode() bool`

GetProtectedMode returns the ProtectedMode field if non-nil, zero value otherwise.

### GetProtectedModeOk

`func (o *CoreDescribeServiceResponse) GetProtectedModeOk() (*bool, bool)`

GetProtectedModeOk returns a tuple with the ProtectedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedMode

`func (o *CoreDescribeServiceResponse) SetProtectedMode(v bool)`

SetProtectedMode sets ProtectedMode field to given value.

### HasProtectedMode

`func (o *CoreDescribeServiceResponse) HasProtectedMode() bool`

HasProtectedMode returns a boolean if a field has been set.

### GetRecovering

`func (o *CoreDescribeServiceResponse) GetRecovering() bool`

GetRecovering returns the Recovering field if non-nil, zero value otherwise.

### GetRecoveringOk

`func (o *CoreDescribeServiceResponse) GetRecoveringOk() (*bool, bool)`

GetRecoveringOk returns a tuple with the Recovering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovering

`func (o *CoreDescribeServiceResponse) SetRecovering(v bool)`

SetRecovering sets Recovering field to given value.

### HasRecovering

`func (o *CoreDescribeServiceResponse) HasRecovering() bool`

HasRecovering returns a boolean if a field has been set.

### GetRegion

`func (o *CoreDescribeServiceResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CoreDescribeServiceResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CoreDescribeServiceResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CoreDescribeServiceResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResourcePool

`func (o *CoreDescribeServiceResponse) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *CoreDescribeServiceResponse) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *CoreDescribeServiceResponse) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *CoreDescribeServiceResponse) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetStatus

`func (o *CoreDescribeServiceResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CoreDescribeServiceResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CoreDescribeServiceResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CoreDescribeServiceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *CoreDescribeServiceResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CoreDescribeServiceResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CoreDescribeServiceResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CoreDescribeServiceResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetType

`func (o *CoreDescribeServiceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CoreDescribeServiceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CoreDescribeServiceResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CoreDescribeServiceResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *CoreDescribeServiceResponse) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CoreDescribeServiceResponse) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CoreDescribeServiceResponse) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CoreDescribeServiceResponse) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *CoreDescribeServiceResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CoreDescribeServiceResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CoreDescribeServiceResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CoreDescribeServiceResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVolumeType

`func (o *CoreDescribeServiceResponse) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CoreDescribeServiceResponse) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CoreDescribeServiceResponse) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *CoreDescribeServiceResponse) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetZone

`func (o *CoreDescribeServiceResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CoreDescribeServiceResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CoreDescribeServiceResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CoreDescribeServiceResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


