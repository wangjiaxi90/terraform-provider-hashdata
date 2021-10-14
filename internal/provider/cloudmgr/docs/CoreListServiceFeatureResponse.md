# CoreListServiceFeatureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableBindNetworkCard** | Pointer to **bool** |  | [optional] 
**EnableBindPublicIp** | Pointer to **bool** |  | [optional] 
**EnableInstanceResize** | Pointer to **bool** |  | [optional] 
**EnableServiceAlert** | Pointer to **bool** |  | [optional] 
**EnableServiceExpand** | Pointer to **bool** |  | [optional] 
**EnableServiceMonitor** | Pointer to **bool** |  | [optional] 
**EnableServiceShrink** | Pointer to **bool** |  | [optional] 
**EnableServiceUpgrade** | Pointer to **bool** |  | [optional] 
**EnableVolumeScaleout** | Pointer to **bool** |  | [optional] 
**StopServiceBeforeExpandService** | Pointer to **bool** |  | [optional] 
**StopServiceBeforeResizeInstance** | Pointer to **bool** |  | [optional] 
**StopServiceBeforeResizeVolume** | Pointer to **bool** |  | [optional] 
**StopServiceBeforeShrinkService** | Pointer to **bool** |  | [optional] 

## Methods

### NewCoreListServiceFeatureResponse

`func NewCoreListServiceFeatureResponse() *CoreListServiceFeatureResponse`

NewCoreListServiceFeatureResponse instantiates a new CoreListServiceFeatureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreListServiceFeatureResponseWithDefaults

`func NewCoreListServiceFeatureResponseWithDefaults() *CoreListServiceFeatureResponse`

NewCoreListServiceFeatureResponseWithDefaults instantiates a new CoreListServiceFeatureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableBindNetworkCard

`func (o *CoreListServiceFeatureResponse) GetEnableBindNetworkCard() bool`

GetEnableBindNetworkCard returns the EnableBindNetworkCard field if non-nil, zero value otherwise.

### GetEnableBindNetworkCardOk

`func (o *CoreListServiceFeatureResponse) GetEnableBindNetworkCardOk() (*bool, bool)`

GetEnableBindNetworkCardOk returns a tuple with the EnableBindNetworkCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBindNetworkCard

`func (o *CoreListServiceFeatureResponse) SetEnableBindNetworkCard(v bool)`

SetEnableBindNetworkCard sets EnableBindNetworkCard field to given value.

### HasEnableBindNetworkCard

`func (o *CoreListServiceFeatureResponse) HasEnableBindNetworkCard() bool`

HasEnableBindNetworkCard returns a boolean if a field has been set.

### GetEnableBindPublicIp

`func (o *CoreListServiceFeatureResponse) GetEnableBindPublicIp() bool`

GetEnableBindPublicIp returns the EnableBindPublicIp field if non-nil, zero value otherwise.

### GetEnableBindPublicIpOk

`func (o *CoreListServiceFeatureResponse) GetEnableBindPublicIpOk() (*bool, bool)`

GetEnableBindPublicIpOk returns a tuple with the EnableBindPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBindPublicIp

`func (o *CoreListServiceFeatureResponse) SetEnableBindPublicIp(v bool)`

SetEnableBindPublicIp sets EnableBindPublicIp field to given value.

### HasEnableBindPublicIp

`func (o *CoreListServiceFeatureResponse) HasEnableBindPublicIp() bool`

HasEnableBindPublicIp returns a boolean if a field has been set.

### GetEnableInstanceResize

`func (o *CoreListServiceFeatureResponse) GetEnableInstanceResize() bool`

GetEnableInstanceResize returns the EnableInstanceResize field if non-nil, zero value otherwise.

### GetEnableInstanceResizeOk

`func (o *CoreListServiceFeatureResponse) GetEnableInstanceResizeOk() (*bool, bool)`

GetEnableInstanceResizeOk returns a tuple with the EnableInstanceResize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInstanceResize

`func (o *CoreListServiceFeatureResponse) SetEnableInstanceResize(v bool)`

SetEnableInstanceResize sets EnableInstanceResize field to given value.

### HasEnableInstanceResize

`func (o *CoreListServiceFeatureResponse) HasEnableInstanceResize() bool`

HasEnableInstanceResize returns a boolean if a field has been set.

### GetEnableServiceAlert

`func (o *CoreListServiceFeatureResponse) GetEnableServiceAlert() bool`

GetEnableServiceAlert returns the EnableServiceAlert field if non-nil, zero value otherwise.

### GetEnableServiceAlertOk

`func (o *CoreListServiceFeatureResponse) GetEnableServiceAlertOk() (*bool, bool)`

GetEnableServiceAlertOk returns a tuple with the EnableServiceAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableServiceAlert

`func (o *CoreListServiceFeatureResponse) SetEnableServiceAlert(v bool)`

SetEnableServiceAlert sets EnableServiceAlert field to given value.

### HasEnableServiceAlert

`func (o *CoreListServiceFeatureResponse) HasEnableServiceAlert() bool`

HasEnableServiceAlert returns a boolean if a field has been set.

### GetEnableServiceExpand

`func (o *CoreListServiceFeatureResponse) GetEnableServiceExpand() bool`

GetEnableServiceExpand returns the EnableServiceExpand field if non-nil, zero value otherwise.

### GetEnableServiceExpandOk

`func (o *CoreListServiceFeatureResponse) GetEnableServiceExpandOk() (*bool, bool)`

GetEnableServiceExpandOk returns a tuple with the EnableServiceExpand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableServiceExpand

`func (o *CoreListServiceFeatureResponse) SetEnableServiceExpand(v bool)`

SetEnableServiceExpand sets EnableServiceExpand field to given value.

### HasEnableServiceExpand

`func (o *CoreListServiceFeatureResponse) HasEnableServiceExpand() bool`

HasEnableServiceExpand returns a boolean if a field has been set.

### GetEnableServiceMonitor

`func (o *CoreListServiceFeatureResponse) GetEnableServiceMonitor() bool`

GetEnableServiceMonitor returns the EnableServiceMonitor field if non-nil, zero value otherwise.

### GetEnableServiceMonitorOk

`func (o *CoreListServiceFeatureResponse) GetEnableServiceMonitorOk() (*bool, bool)`

GetEnableServiceMonitorOk returns a tuple with the EnableServiceMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableServiceMonitor

`func (o *CoreListServiceFeatureResponse) SetEnableServiceMonitor(v bool)`

SetEnableServiceMonitor sets EnableServiceMonitor field to given value.

### HasEnableServiceMonitor

`func (o *CoreListServiceFeatureResponse) HasEnableServiceMonitor() bool`

HasEnableServiceMonitor returns a boolean if a field has been set.

### GetEnableServiceShrink

`func (o *CoreListServiceFeatureResponse) GetEnableServiceShrink() bool`

GetEnableServiceShrink returns the EnableServiceShrink field if non-nil, zero value otherwise.

### GetEnableServiceShrinkOk

`func (o *CoreListServiceFeatureResponse) GetEnableServiceShrinkOk() (*bool, bool)`

GetEnableServiceShrinkOk returns a tuple with the EnableServiceShrink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableServiceShrink

`func (o *CoreListServiceFeatureResponse) SetEnableServiceShrink(v bool)`

SetEnableServiceShrink sets EnableServiceShrink field to given value.

### HasEnableServiceShrink

`func (o *CoreListServiceFeatureResponse) HasEnableServiceShrink() bool`

HasEnableServiceShrink returns a boolean if a field has been set.

### GetEnableServiceUpgrade

`func (o *CoreListServiceFeatureResponse) GetEnableServiceUpgrade() bool`

GetEnableServiceUpgrade returns the EnableServiceUpgrade field if non-nil, zero value otherwise.

### GetEnableServiceUpgradeOk

`func (o *CoreListServiceFeatureResponse) GetEnableServiceUpgradeOk() (*bool, bool)`

GetEnableServiceUpgradeOk returns a tuple with the EnableServiceUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableServiceUpgrade

`func (o *CoreListServiceFeatureResponse) SetEnableServiceUpgrade(v bool)`

SetEnableServiceUpgrade sets EnableServiceUpgrade field to given value.

### HasEnableServiceUpgrade

`func (o *CoreListServiceFeatureResponse) HasEnableServiceUpgrade() bool`

HasEnableServiceUpgrade returns a boolean if a field has been set.

### GetEnableVolumeScaleout

`func (o *CoreListServiceFeatureResponse) GetEnableVolumeScaleout() bool`

GetEnableVolumeScaleout returns the EnableVolumeScaleout field if non-nil, zero value otherwise.

### GetEnableVolumeScaleoutOk

`func (o *CoreListServiceFeatureResponse) GetEnableVolumeScaleoutOk() (*bool, bool)`

GetEnableVolumeScaleoutOk returns a tuple with the EnableVolumeScaleout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVolumeScaleout

`func (o *CoreListServiceFeatureResponse) SetEnableVolumeScaleout(v bool)`

SetEnableVolumeScaleout sets EnableVolumeScaleout field to given value.

### HasEnableVolumeScaleout

`func (o *CoreListServiceFeatureResponse) HasEnableVolumeScaleout() bool`

HasEnableVolumeScaleout returns a boolean if a field has been set.

### GetStopServiceBeforeExpandService

`func (o *CoreListServiceFeatureResponse) GetStopServiceBeforeExpandService() bool`

GetStopServiceBeforeExpandService returns the StopServiceBeforeExpandService field if non-nil, zero value otherwise.

### GetStopServiceBeforeExpandServiceOk

`func (o *CoreListServiceFeatureResponse) GetStopServiceBeforeExpandServiceOk() (*bool, bool)`

GetStopServiceBeforeExpandServiceOk returns a tuple with the StopServiceBeforeExpandService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopServiceBeforeExpandService

`func (o *CoreListServiceFeatureResponse) SetStopServiceBeforeExpandService(v bool)`

SetStopServiceBeforeExpandService sets StopServiceBeforeExpandService field to given value.

### HasStopServiceBeforeExpandService

`func (o *CoreListServiceFeatureResponse) HasStopServiceBeforeExpandService() bool`

HasStopServiceBeforeExpandService returns a boolean if a field has been set.

### GetStopServiceBeforeResizeInstance

`func (o *CoreListServiceFeatureResponse) GetStopServiceBeforeResizeInstance() bool`

GetStopServiceBeforeResizeInstance returns the StopServiceBeforeResizeInstance field if non-nil, zero value otherwise.

### GetStopServiceBeforeResizeInstanceOk

`func (o *CoreListServiceFeatureResponse) GetStopServiceBeforeResizeInstanceOk() (*bool, bool)`

GetStopServiceBeforeResizeInstanceOk returns a tuple with the StopServiceBeforeResizeInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopServiceBeforeResizeInstance

`func (o *CoreListServiceFeatureResponse) SetStopServiceBeforeResizeInstance(v bool)`

SetStopServiceBeforeResizeInstance sets StopServiceBeforeResizeInstance field to given value.

### HasStopServiceBeforeResizeInstance

`func (o *CoreListServiceFeatureResponse) HasStopServiceBeforeResizeInstance() bool`

HasStopServiceBeforeResizeInstance returns a boolean if a field has been set.

### GetStopServiceBeforeResizeVolume

`func (o *CoreListServiceFeatureResponse) GetStopServiceBeforeResizeVolume() bool`

GetStopServiceBeforeResizeVolume returns the StopServiceBeforeResizeVolume field if non-nil, zero value otherwise.

### GetStopServiceBeforeResizeVolumeOk

`func (o *CoreListServiceFeatureResponse) GetStopServiceBeforeResizeVolumeOk() (*bool, bool)`

GetStopServiceBeforeResizeVolumeOk returns a tuple with the StopServiceBeforeResizeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopServiceBeforeResizeVolume

`func (o *CoreListServiceFeatureResponse) SetStopServiceBeforeResizeVolume(v bool)`

SetStopServiceBeforeResizeVolume sets StopServiceBeforeResizeVolume field to given value.

### HasStopServiceBeforeResizeVolume

`func (o *CoreListServiceFeatureResponse) HasStopServiceBeforeResizeVolume() bool`

HasStopServiceBeforeResizeVolume returns a boolean if a field has been set.

### GetStopServiceBeforeShrinkService

`func (o *CoreListServiceFeatureResponse) GetStopServiceBeforeShrinkService() bool`

GetStopServiceBeforeShrinkService returns the StopServiceBeforeShrinkService field if non-nil, zero value otherwise.

### GetStopServiceBeforeShrinkServiceOk

`func (o *CoreListServiceFeatureResponse) GetStopServiceBeforeShrinkServiceOk() (*bool, bool)`

GetStopServiceBeforeShrinkServiceOk returns a tuple with the StopServiceBeforeShrinkService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopServiceBeforeShrinkService

`func (o *CoreListServiceFeatureResponse) SetStopServiceBeforeShrinkService(v bool)`

SetStopServiceBeforeShrinkService sets StopServiceBeforeShrinkService field to given value.

### HasStopServiceBeforeShrinkService

`func (o *CoreListServiceFeatureResponse) HasStopServiceBeforeShrinkService() bool`

HasStopServiceBeforeShrinkService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


