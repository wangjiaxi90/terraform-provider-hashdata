# CloudmgrcoreIaasResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**ElasticNic** | Pointer to [**CoreCreateElasticNicRequest**](CoreCreateElasticNicRequest.md) |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**InternetAccess** | Pointer to [**CoreCreateInternetAccessRequest**](CoreCreateInternetAccessRequest.md) |  | [optional] 
**VolumeSize** | Pointer to **int32** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudmgrcoreIaasResource

`func NewCloudmgrcoreIaasResource() *CloudmgrcoreIaasResource`

NewCloudmgrcoreIaasResource instantiates a new CloudmgrcoreIaasResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudmgrcoreIaasResourceWithDefaults

`func NewCloudmgrcoreIaasResourceWithDefaults() *CloudmgrcoreIaasResource`

NewCloudmgrcoreIaasResourceWithDefaults instantiates a new CloudmgrcoreIaasResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CloudmgrcoreIaasResource) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloudmgrcoreIaasResource) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloudmgrcoreIaasResource) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CloudmgrcoreIaasResource) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetElasticNic

`func (o *CloudmgrcoreIaasResource) GetElasticNic() CoreCreateElasticNicRequest`

GetElasticNic returns the ElasticNic field if non-nil, zero value otherwise.

### GetElasticNicOk

`func (o *CloudmgrcoreIaasResource) GetElasticNicOk() (*CoreCreateElasticNicRequest, bool)`

GetElasticNicOk returns a tuple with the ElasticNic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticNic

`func (o *CloudmgrcoreIaasResource) SetElasticNic(v CoreCreateElasticNicRequest)`

SetElasticNic sets ElasticNic field to given value.

### HasElasticNic

`func (o *CloudmgrcoreIaasResource) HasElasticNic() bool`

HasElasticNic returns a boolean if a field has been set.

### GetImage

`func (o *CloudmgrcoreIaasResource) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CloudmgrcoreIaasResource) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CloudmgrcoreIaasResource) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CloudmgrcoreIaasResource) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetInstanceType

`func (o *CloudmgrcoreIaasResource) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *CloudmgrcoreIaasResource) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *CloudmgrcoreIaasResource) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *CloudmgrcoreIaasResource) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInternetAccess

`func (o *CloudmgrcoreIaasResource) GetInternetAccess() CoreCreateInternetAccessRequest`

GetInternetAccess returns the InternetAccess field if non-nil, zero value otherwise.

### GetInternetAccessOk

`func (o *CloudmgrcoreIaasResource) GetInternetAccessOk() (*CoreCreateInternetAccessRequest, bool)`

GetInternetAccessOk returns a tuple with the InternetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetAccess

`func (o *CloudmgrcoreIaasResource) SetInternetAccess(v CoreCreateInternetAccessRequest)`

SetInternetAccess sets InternetAccess field to given value.

### HasInternetAccess

`func (o *CloudmgrcoreIaasResource) HasInternetAccess() bool`

HasInternetAccess returns a boolean if a field has been set.

### GetVolumeSize

`func (o *CloudmgrcoreIaasResource) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *CloudmgrcoreIaasResource) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *CloudmgrcoreIaasResource) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *CloudmgrcoreIaasResource) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.

### GetVolumeType

`func (o *CloudmgrcoreIaasResource) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CloudmgrcoreIaasResource) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CloudmgrcoreIaasResource) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *CloudmgrcoreIaasResource) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetZone

`func (o *CloudmgrcoreIaasResource) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CloudmgrcoreIaasResource) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CloudmgrcoreIaasResource) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CloudmgrcoreIaasResource) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


