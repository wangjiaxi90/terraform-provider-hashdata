# CoreCreateInternetAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EipId** | Pointer to **[]string** |  | [optional] 
**MaxBandwidth** | Pointer to **int32** |  | [optional] 
**PayByBandwidth** | Pointer to **bool** |  | [optional] 

## Methods

### NewCoreCreateInternetAccessRequest

`func NewCoreCreateInternetAccessRequest() *CoreCreateInternetAccessRequest`

NewCoreCreateInternetAccessRequest instantiates a new CoreCreateInternetAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCreateInternetAccessRequestWithDefaults

`func NewCoreCreateInternetAccessRequestWithDefaults() *CoreCreateInternetAccessRequest`

NewCoreCreateInternetAccessRequestWithDefaults instantiates a new CoreCreateInternetAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEipId

`func (o *CoreCreateInternetAccessRequest) GetEipId() []string`

GetEipId returns the EipId field if non-nil, zero value otherwise.

### GetEipIdOk

`func (o *CoreCreateInternetAccessRequest) GetEipIdOk() (*[]string, bool)`

GetEipIdOk returns a tuple with the EipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEipId

`func (o *CoreCreateInternetAccessRequest) SetEipId(v []string)`

SetEipId sets EipId field to given value.

### HasEipId

`func (o *CoreCreateInternetAccessRequest) HasEipId() bool`

HasEipId returns a boolean if a field has been set.

### GetMaxBandwidth

`func (o *CoreCreateInternetAccessRequest) GetMaxBandwidth() int32`

GetMaxBandwidth returns the MaxBandwidth field if non-nil, zero value otherwise.

### GetMaxBandwidthOk

`func (o *CoreCreateInternetAccessRequest) GetMaxBandwidthOk() (*int32, bool)`

GetMaxBandwidthOk returns a tuple with the MaxBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBandwidth

`func (o *CoreCreateInternetAccessRequest) SetMaxBandwidth(v int32)`

SetMaxBandwidth sets MaxBandwidth field to given value.

### HasMaxBandwidth

`func (o *CoreCreateInternetAccessRequest) HasMaxBandwidth() bool`

HasMaxBandwidth returns a boolean if a field has been set.

### GetPayByBandwidth

`func (o *CoreCreateInternetAccessRequest) GetPayByBandwidth() bool`

GetPayByBandwidth returns the PayByBandwidth field if non-nil, zero value otherwise.

### GetPayByBandwidthOk

`func (o *CoreCreateInternetAccessRequest) GetPayByBandwidthOk() (*bool, bool)`

GetPayByBandwidthOk returns a tuple with the PayByBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayByBandwidth

`func (o *CoreCreateInternetAccessRequest) SetPayByBandwidth(v bool)`

SetPayByBandwidth sets PayByBandwidth field to given value.

### HasPayByBandwidth

`func (o *CoreCreateInternetAccessRequest) HasPayByBandwidth() bool`

HasPayByBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


