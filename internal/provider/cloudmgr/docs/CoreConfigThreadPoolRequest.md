# CoreConfigThreadPoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **int32** |  | [optional] 
**CorePoolSize** | Pointer to **int32** |  | [optional] 
**MaxPoolSize** | Pointer to **int32** |  | [optional] 
**PoolType** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreConfigThreadPoolRequest

`func NewCoreConfigThreadPoolRequest() *CoreConfigThreadPoolRequest`

NewCoreConfigThreadPoolRequest instantiates a new CoreConfigThreadPoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreConfigThreadPoolRequestWithDefaults

`func NewCoreConfigThreadPoolRequestWithDefaults() *CoreConfigThreadPoolRequest`

NewCoreConfigThreadPoolRequestWithDefaults instantiates a new CoreConfigThreadPoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *CoreConfigThreadPoolRequest) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *CoreConfigThreadPoolRequest) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *CoreConfigThreadPoolRequest) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *CoreConfigThreadPoolRequest) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCorePoolSize

`func (o *CoreConfigThreadPoolRequest) GetCorePoolSize() int32`

GetCorePoolSize returns the CorePoolSize field if non-nil, zero value otherwise.

### GetCorePoolSizeOk

`func (o *CoreConfigThreadPoolRequest) GetCorePoolSizeOk() (*int32, bool)`

GetCorePoolSizeOk returns a tuple with the CorePoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorePoolSize

`func (o *CoreConfigThreadPoolRequest) SetCorePoolSize(v int32)`

SetCorePoolSize sets CorePoolSize field to given value.

### HasCorePoolSize

`func (o *CoreConfigThreadPoolRequest) HasCorePoolSize() bool`

HasCorePoolSize returns a boolean if a field has been set.

### GetMaxPoolSize

`func (o *CoreConfigThreadPoolRequest) GetMaxPoolSize() int32`

GetMaxPoolSize returns the MaxPoolSize field if non-nil, zero value otherwise.

### GetMaxPoolSizeOk

`func (o *CoreConfigThreadPoolRequest) GetMaxPoolSizeOk() (*int32, bool)`

GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolSize

`func (o *CoreConfigThreadPoolRequest) SetMaxPoolSize(v int32)`

SetMaxPoolSize sets MaxPoolSize field to given value.

### HasMaxPoolSize

`func (o *CoreConfigThreadPoolRequest) HasMaxPoolSize() bool`

HasMaxPoolSize returns a boolean if a field has been set.

### GetPoolType

`func (o *CoreConfigThreadPoolRequest) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *CoreConfigThreadPoolRequest) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *CoreConfigThreadPoolRequest) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *CoreConfigThreadPoolRequest) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


