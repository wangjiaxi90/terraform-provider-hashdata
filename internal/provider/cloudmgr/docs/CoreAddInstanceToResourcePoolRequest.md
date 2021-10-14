# CoreAddInstanceToResourcePoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]string** |  | [optional] 
**Instances** | Pointer to **[]string** |  | [optional] 
**RootPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreAddInstanceToResourcePoolRequest

`func NewCoreAddInstanceToResourcePoolRequest() *CoreAddInstanceToResourcePoolRequest`

NewCoreAddInstanceToResourcePoolRequest instantiates a new CoreAddInstanceToResourcePoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreAddInstanceToResourcePoolRequestWithDefaults

`func NewCoreAddInstanceToResourcePoolRequestWithDefaults() *CoreAddInstanceToResourcePoolRequest`

NewCoreAddInstanceToResourcePoolRequestWithDefaults instantiates a new CoreAddInstanceToResourcePoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *CoreAddInstanceToResourcePoolRequest) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *CoreAddInstanceToResourcePoolRequest) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *CoreAddInstanceToResourcePoolRequest) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *CoreAddInstanceToResourcePoolRequest) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetInstances

`func (o *CoreAddInstanceToResourcePoolRequest) GetInstances() []string`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *CoreAddInstanceToResourcePoolRequest) GetInstancesOk() (*[]string, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *CoreAddInstanceToResourcePoolRequest) SetInstances(v []string)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *CoreAddInstanceToResourcePoolRequest) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetRootPassword

`func (o *CoreAddInstanceToResourcePoolRequest) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *CoreAddInstanceToResourcePoolRequest) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *CoreAddInstanceToResourcePoolRequest) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *CoreAddInstanceToResourcePoolRequest) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


