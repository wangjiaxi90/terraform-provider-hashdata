# CoreCreateServiceComponentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iaas** | Pointer to [**CloudmgrcoreIaasResource**](CloudmgrcoreIaasResource.md) |  | [optional] 
**Reserved** | Pointer to [**CoreReservedResource**](CoreReservedResource.md) |  | [optional] 

## Methods

### NewCoreCreateServiceComponentRequest

`func NewCoreCreateServiceComponentRequest() *CoreCreateServiceComponentRequest`

NewCoreCreateServiceComponentRequest instantiates a new CoreCreateServiceComponentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCreateServiceComponentRequestWithDefaults

`func NewCoreCreateServiceComponentRequestWithDefaults() *CoreCreateServiceComponentRequest`

NewCoreCreateServiceComponentRequestWithDefaults instantiates a new CoreCreateServiceComponentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIaas

`func (o *CoreCreateServiceComponentRequest) GetIaas() CloudmgrcoreIaasResource`

GetIaas returns the Iaas field if non-nil, zero value otherwise.

### GetIaasOk

`func (o *CoreCreateServiceComponentRequest) GetIaasOk() (*CloudmgrcoreIaasResource, bool)`

GetIaasOk returns a tuple with the Iaas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIaas

`func (o *CoreCreateServiceComponentRequest) SetIaas(v CloudmgrcoreIaasResource)`

SetIaas sets Iaas field to given value.

### HasIaas

`func (o *CoreCreateServiceComponentRequest) HasIaas() bool`

HasIaas returns a boolean if a field has been set.

### GetReserved

`func (o *CoreCreateServiceComponentRequest) GetReserved() CoreReservedResource`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *CoreCreateServiceComponentRequest) GetReservedOk() (*CoreReservedResource, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *CoreCreateServiceComponentRequest) SetReserved(v CoreReservedResource)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *CoreCreateServiceComponentRequest) HasReserved() bool`

HasReserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


