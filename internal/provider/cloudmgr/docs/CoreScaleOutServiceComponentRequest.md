# CoreScaleOutServiceComponentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iaas** | Pointer to [**CoreScaleOutIaasResource**](CoreScaleOutIaasResource.md) |  | [optional] 
**Reserved** | Pointer to [**CoreScaleOutReservedResource**](CoreScaleOutReservedResource.md) |  | [optional] 

## Methods

### NewCoreScaleOutServiceComponentRequest

`func NewCoreScaleOutServiceComponentRequest() *CoreScaleOutServiceComponentRequest`

NewCoreScaleOutServiceComponentRequest instantiates a new CoreScaleOutServiceComponentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreScaleOutServiceComponentRequestWithDefaults

`func NewCoreScaleOutServiceComponentRequestWithDefaults() *CoreScaleOutServiceComponentRequest`

NewCoreScaleOutServiceComponentRequestWithDefaults instantiates a new CoreScaleOutServiceComponentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIaas

`func (o *CoreScaleOutServiceComponentRequest) GetIaas() CoreScaleOutIaasResource`

GetIaas returns the Iaas field if non-nil, zero value otherwise.

### GetIaasOk

`func (o *CoreScaleOutServiceComponentRequest) GetIaasOk() (*CoreScaleOutIaasResource, bool)`

GetIaasOk returns a tuple with the Iaas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIaas

`func (o *CoreScaleOutServiceComponentRequest) SetIaas(v CoreScaleOutIaasResource)`

SetIaas sets Iaas field to given value.

### HasIaas

`func (o *CoreScaleOutServiceComponentRequest) HasIaas() bool`

HasIaas returns a boolean if a field has been set.

### GetReserved

`func (o *CoreScaleOutServiceComponentRequest) GetReserved() CoreScaleOutReservedResource`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *CoreScaleOutServiceComponentRequest) GetReservedOk() (*CoreScaleOutReservedResource, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *CoreScaleOutServiceComponentRequest) SetReserved(v CoreScaleOutReservedResource)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *CoreScaleOutServiceComponentRequest) HasReserved() bool`

HasReserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


