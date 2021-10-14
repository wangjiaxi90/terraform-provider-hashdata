# CoreCreateUnmanagedServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Extra** | Pointer to [**CoreCreateServiceIaasExtraRequest**](CoreCreateServiceIaasExtraRequest.md) |  | [optional] 
**Iaas** | Pointer to [**CloudmgrcoreIaasResource**](CloudmgrcoreIaasResource.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreCreateUnmanagedServiceRequest

`func NewCoreCreateUnmanagedServiceRequest() *CoreCreateUnmanagedServiceRequest`

NewCoreCreateUnmanagedServiceRequest instantiates a new CoreCreateUnmanagedServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCreateUnmanagedServiceRequestWithDefaults

`func NewCoreCreateUnmanagedServiceRequestWithDefaults() *CoreCreateUnmanagedServiceRequest`

NewCoreCreateUnmanagedServiceRequestWithDefaults instantiates a new CoreCreateUnmanagedServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CoreCreateUnmanagedServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoreCreateUnmanagedServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoreCreateUnmanagedServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoreCreateUnmanagedServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtra

`func (o *CoreCreateUnmanagedServiceRequest) GetExtra() CoreCreateServiceIaasExtraRequest`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *CoreCreateUnmanagedServiceRequest) GetExtraOk() (*CoreCreateServiceIaasExtraRequest, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *CoreCreateUnmanagedServiceRequest) SetExtra(v CoreCreateServiceIaasExtraRequest)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *CoreCreateUnmanagedServiceRequest) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetIaas

`func (o *CoreCreateUnmanagedServiceRequest) GetIaas() CloudmgrcoreIaasResource`

GetIaas returns the Iaas field if non-nil, zero value otherwise.

### GetIaasOk

`func (o *CoreCreateUnmanagedServiceRequest) GetIaasOk() (*CloudmgrcoreIaasResource, bool)`

GetIaasOk returns a tuple with the Iaas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIaas

`func (o *CoreCreateUnmanagedServiceRequest) SetIaas(v CloudmgrcoreIaasResource)`

SetIaas sets Iaas field to given value.

### HasIaas

`func (o *CoreCreateUnmanagedServiceRequest) HasIaas() bool`

HasIaas returns a boolean if a field has been set.

### GetName

`func (o *CoreCreateUnmanagedServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreCreateUnmanagedServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreCreateUnmanagedServiceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoreCreateUnmanagedServiceRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


