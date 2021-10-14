# CoreDescribeServiceTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to [**[]CoreDescribeServiceComponentResponse**](CoreDescribeServiceComponentResponse.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCoreDescribeServiceTypeResponse

`func NewCoreDescribeServiceTypeResponse() *CoreDescribeServiceTypeResponse`

NewCoreDescribeServiceTypeResponse instantiates a new CoreDescribeServiceTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeServiceTypeResponseWithDefaults

`func NewCoreDescribeServiceTypeResponseWithDefaults() *CoreDescribeServiceTypeResponse`

NewCoreDescribeServiceTypeResponseWithDefaults instantiates a new CoreDescribeServiceTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *CoreDescribeServiceTypeResponse) GetComponent() []CoreDescribeServiceComponentResponse`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *CoreDescribeServiceTypeResponse) GetComponentOk() (*[]CoreDescribeServiceComponentResponse, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *CoreDescribeServiceTypeResponse) SetComponent(v []CoreDescribeServiceComponentResponse)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *CoreDescribeServiceTypeResponse) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetEnabled

`func (o *CoreDescribeServiceTypeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CoreDescribeServiceTypeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CoreDescribeServiceTypeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CoreDescribeServiceTypeResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *CoreDescribeServiceTypeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreDescribeServiceTypeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreDescribeServiceTypeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoreDescribeServiceTypeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *CoreDescribeServiceTypeResponse) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CoreDescribeServiceTypeResponse) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CoreDescribeServiceTypeResponse) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CoreDescribeServiceTypeResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


