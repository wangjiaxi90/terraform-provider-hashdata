# CoreCreateCatalogRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to [**CoreCreateServiceComponentRequest**](CoreCreateServiceComponentRequest.md) |  | [optional] 
**ChargeChannel** | Pointer to **string** |  | [optional] 
**ChargeType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Etcd** | Pointer to [**CoreCreateServiceComponentRequest**](CoreCreateServiceComponentRequest.md) |  | [optional] 
**ExpiredAt** | Pointer to **string** |  | [optional] 
**Extras** | Pointer to [**CoreCreateServiceIaasExtraRequest**](CoreCreateServiceIaasExtraRequest.md) |  | [optional] 
**Fdb** | Pointer to [**CoreCreateServiceComponentRequest**](CoreCreateServiceComponentRequest.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Monitor** | Pointer to [**CoreCreateServiceComponentRequest**](CoreCreateServiceComponentRequest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Oss** | Pointer to [**CoreCreateServiceOssZoneRequest**](CoreCreateServiceOssZoneRequest.md) |  | [optional] 

## Methods

### NewCoreCreateCatalogRequest

`func NewCoreCreateCatalogRequest() *CoreCreateCatalogRequest`

NewCoreCreateCatalogRequest instantiates a new CoreCreateCatalogRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCreateCatalogRequestWithDefaults

`func NewCoreCreateCatalogRequestWithDefaults() *CoreCreateCatalogRequest`

NewCoreCreateCatalogRequestWithDefaults instantiates a new CoreCreateCatalogRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *CoreCreateCatalogRequest) GetCatalog() CoreCreateServiceComponentRequest`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *CoreCreateCatalogRequest) GetCatalogOk() (*CoreCreateServiceComponentRequest, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *CoreCreateCatalogRequest) SetCatalog(v CoreCreateServiceComponentRequest)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *CoreCreateCatalogRequest) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetChargeChannel

`func (o *CoreCreateCatalogRequest) GetChargeChannel() string`

GetChargeChannel returns the ChargeChannel field if non-nil, zero value otherwise.

### GetChargeChannelOk

`func (o *CoreCreateCatalogRequest) GetChargeChannelOk() (*string, bool)`

GetChargeChannelOk returns a tuple with the ChargeChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeChannel

`func (o *CoreCreateCatalogRequest) SetChargeChannel(v string)`

SetChargeChannel sets ChargeChannel field to given value.

### HasChargeChannel

`func (o *CoreCreateCatalogRequest) HasChargeChannel() bool`

HasChargeChannel returns a boolean if a field has been set.

### GetChargeType

`func (o *CoreCreateCatalogRequest) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *CoreCreateCatalogRequest) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *CoreCreateCatalogRequest) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *CoreCreateCatalogRequest) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetDescription

`func (o *CoreCreateCatalogRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoreCreateCatalogRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoreCreateCatalogRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoreCreateCatalogRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEtcd

`func (o *CoreCreateCatalogRequest) GetEtcd() CoreCreateServiceComponentRequest`

GetEtcd returns the Etcd field if non-nil, zero value otherwise.

### GetEtcdOk

`func (o *CoreCreateCatalogRequest) GetEtcdOk() (*CoreCreateServiceComponentRequest, bool)`

GetEtcdOk returns a tuple with the Etcd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtcd

`func (o *CoreCreateCatalogRequest) SetEtcd(v CoreCreateServiceComponentRequest)`

SetEtcd sets Etcd field to given value.

### HasEtcd

`func (o *CoreCreateCatalogRequest) HasEtcd() bool`

HasEtcd returns a boolean if a field has been set.

### GetExpiredAt

`func (o *CoreCreateCatalogRequest) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *CoreCreateCatalogRequest) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *CoreCreateCatalogRequest) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *CoreCreateCatalogRequest) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetExtras

`func (o *CoreCreateCatalogRequest) GetExtras() CoreCreateServiceIaasExtraRequest`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *CoreCreateCatalogRequest) GetExtrasOk() (*CoreCreateServiceIaasExtraRequest, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *CoreCreateCatalogRequest) SetExtras(v CoreCreateServiceIaasExtraRequest)`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *CoreCreateCatalogRequest) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetFdb

`func (o *CoreCreateCatalogRequest) GetFdb() CoreCreateServiceComponentRequest`

GetFdb returns the Fdb field if non-nil, zero value otherwise.

### GetFdbOk

`func (o *CoreCreateCatalogRequest) GetFdbOk() (*CoreCreateServiceComponentRequest, bool)`

GetFdbOk returns a tuple with the Fdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdb

`func (o *CoreCreateCatalogRequest) SetFdb(v CoreCreateServiceComponentRequest)`

SetFdb sets Fdb field to given value.

### HasFdb

`func (o *CoreCreateCatalogRequest) HasFdb() bool`

HasFdb returns a boolean if a field has been set.

### GetMetadata

`func (o *CoreCreateCatalogRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CoreCreateCatalogRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CoreCreateCatalogRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CoreCreateCatalogRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMonitor

`func (o *CoreCreateCatalogRequest) GetMonitor() CoreCreateServiceComponentRequest`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *CoreCreateCatalogRequest) GetMonitorOk() (*CoreCreateServiceComponentRequest, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *CoreCreateCatalogRequest) SetMonitor(v CoreCreateServiceComponentRequest)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *CoreCreateCatalogRequest) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetName

`func (o *CoreCreateCatalogRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreCreateCatalogRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreCreateCatalogRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoreCreateCatalogRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOss

`func (o *CoreCreateCatalogRequest) GetOss() CoreCreateServiceOssZoneRequest`

GetOss returns the Oss field if non-nil, zero value otherwise.

### GetOssOk

`func (o *CoreCreateCatalogRequest) GetOssOk() (*CoreCreateServiceOssZoneRequest, bool)`

GetOssOk returns a tuple with the Oss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOss

`func (o *CoreCreateCatalogRequest) SetOss(v CoreCreateServiceOssZoneRequest)`

SetOss sets Oss field to given value.

### HasOss

`func (o *CoreCreateCatalogRequest) HasOss() bool`

HasOss returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


