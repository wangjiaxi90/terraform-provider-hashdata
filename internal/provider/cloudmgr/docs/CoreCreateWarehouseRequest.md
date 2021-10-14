# CoreCreateWarehouseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to **string** |  | [optional] 
**ChargeChannel** | Pointer to **string** |  | [optional] 
**ChargeType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExpiredAt** | Pointer to **string** |  | [optional] 
**Extras** | Pointer to [**CoreCreateServiceIaasExtraRequest**](CoreCreateServiceIaasExtraRequest.md) |  | [optional] 
**Features** | Pointer to [**CoreCreateServiceFeatureRequest**](CoreCreateServiceFeatureRequest.md) |  | [optional] 
**Master** | Pointer to [**CoreCreateServiceComponentRequest**](CoreCreateServiceComponentRequest.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Oss** | Pointer to [**CoreCreateServiceOssZoneRequest**](CoreCreateServiceOssZoneRequest.md) |  | [optional] 
**Segment** | Pointer to [**CoreCreateServiceComponentRequest**](CoreCreateServiceComponentRequest.md) |  | [optional] 
**Standby** | Pointer to [**CoreCreateServiceComponentRequest**](CoreCreateServiceComponentRequest.md) |  | [optional] 

## Methods

### NewCoreCreateWarehouseRequest

`func NewCoreCreateWarehouseRequest() *CoreCreateWarehouseRequest`

NewCoreCreateWarehouseRequest instantiates a new CoreCreateWarehouseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCreateWarehouseRequestWithDefaults

`func NewCoreCreateWarehouseRequestWithDefaults() *CoreCreateWarehouseRequest`

NewCoreCreateWarehouseRequestWithDefaults instantiates a new CoreCreateWarehouseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *CoreCreateWarehouseRequest) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *CoreCreateWarehouseRequest) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *CoreCreateWarehouseRequest) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *CoreCreateWarehouseRequest) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetChargeChannel

`func (o *CoreCreateWarehouseRequest) GetChargeChannel() string`

GetChargeChannel returns the ChargeChannel field if non-nil, zero value otherwise.

### GetChargeChannelOk

`func (o *CoreCreateWarehouseRequest) GetChargeChannelOk() (*string, bool)`

GetChargeChannelOk returns a tuple with the ChargeChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeChannel

`func (o *CoreCreateWarehouseRequest) SetChargeChannel(v string)`

SetChargeChannel sets ChargeChannel field to given value.

### HasChargeChannel

`func (o *CoreCreateWarehouseRequest) HasChargeChannel() bool`

HasChargeChannel returns a boolean if a field has been set.

### GetChargeType

`func (o *CoreCreateWarehouseRequest) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *CoreCreateWarehouseRequest) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *CoreCreateWarehouseRequest) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *CoreCreateWarehouseRequest) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetDescription

`func (o *CoreCreateWarehouseRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoreCreateWarehouseRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoreCreateWarehouseRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoreCreateWarehouseRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiredAt

`func (o *CoreCreateWarehouseRequest) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *CoreCreateWarehouseRequest) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *CoreCreateWarehouseRequest) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *CoreCreateWarehouseRequest) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetExtras

`func (o *CoreCreateWarehouseRequest) GetExtras() CoreCreateServiceIaasExtraRequest`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *CoreCreateWarehouseRequest) GetExtrasOk() (*CoreCreateServiceIaasExtraRequest, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *CoreCreateWarehouseRequest) SetExtras(v CoreCreateServiceIaasExtraRequest)`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *CoreCreateWarehouseRequest) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetFeatures

`func (o *CoreCreateWarehouseRequest) GetFeatures() CoreCreateServiceFeatureRequest`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CoreCreateWarehouseRequest) GetFeaturesOk() (*CoreCreateServiceFeatureRequest, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CoreCreateWarehouseRequest) SetFeatures(v CoreCreateServiceFeatureRequest)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CoreCreateWarehouseRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetMaster

`func (o *CoreCreateWarehouseRequest) GetMaster() CoreCreateServiceComponentRequest`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *CoreCreateWarehouseRequest) GetMasterOk() (*CoreCreateServiceComponentRequest, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *CoreCreateWarehouseRequest) SetMaster(v CoreCreateServiceComponentRequest)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *CoreCreateWarehouseRequest) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetMetadata

`func (o *CoreCreateWarehouseRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CoreCreateWarehouseRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CoreCreateWarehouseRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CoreCreateWarehouseRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CoreCreateWarehouseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreCreateWarehouseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreCreateWarehouseRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoreCreateWarehouseRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOss

`func (o *CoreCreateWarehouseRequest) GetOss() CoreCreateServiceOssZoneRequest`

GetOss returns the Oss field if non-nil, zero value otherwise.

### GetOssOk

`func (o *CoreCreateWarehouseRequest) GetOssOk() (*CoreCreateServiceOssZoneRequest, bool)`

GetOssOk returns a tuple with the Oss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOss

`func (o *CoreCreateWarehouseRequest) SetOss(v CoreCreateServiceOssZoneRequest)`

SetOss sets Oss field to given value.

### HasOss

`func (o *CoreCreateWarehouseRequest) HasOss() bool`

HasOss returns a boolean if a field has been set.

### GetSegment

`func (o *CoreCreateWarehouseRequest) GetSegment() CoreCreateServiceComponentRequest`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *CoreCreateWarehouseRequest) GetSegmentOk() (*CoreCreateServiceComponentRequest, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *CoreCreateWarehouseRequest) SetSegment(v CoreCreateServiceComponentRequest)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *CoreCreateWarehouseRequest) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetStandby

`func (o *CoreCreateWarehouseRequest) GetStandby() CoreCreateServiceComponentRequest`

GetStandby returns the Standby field if non-nil, zero value otherwise.

### GetStandbyOk

`func (o *CoreCreateWarehouseRequest) GetStandbyOk() (*CoreCreateServiceComponentRequest, bool)`

GetStandbyOk returns a tuple with the Standby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandby

`func (o *CoreCreateWarehouseRequest) SetStandby(v CoreCreateServiceComponentRequest)`

SetStandby sets Standby field to given value.

### HasStandby

`func (o *CoreCreateWarehouseRequest) HasStandby() bool`

HasStandby returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


