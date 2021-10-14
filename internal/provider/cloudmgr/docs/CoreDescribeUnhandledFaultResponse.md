# CoreDescribeUnhandledFaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**ComponentType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**FailureMessage** | Pointer to **string** |  | [optional] 
**FailureType** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreDescribeUnhandledFaultResponse

`func NewCoreDescribeUnhandledFaultResponse() *CoreDescribeUnhandledFaultResponse`

NewCoreDescribeUnhandledFaultResponse instantiates a new CoreDescribeUnhandledFaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeUnhandledFaultResponseWithDefaults

`func NewCoreDescribeUnhandledFaultResponseWithDefaults() *CoreDescribeUnhandledFaultResponse`

NewCoreDescribeUnhandledFaultResponseWithDefaults instantiates a new CoreDescribeUnhandledFaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CoreDescribeUnhandledFaultResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CoreDescribeUnhandledFaultResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CoreDescribeUnhandledFaultResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CoreDescribeUnhandledFaultResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetComponentType

`func (o *CoreDescribeUnhandledFaultResponse) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *CoreDescribeUnhandledFaultResponse) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *CoreDescribeUnhandledFaultResponse) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *CoreDescribeUnhandledFaultResponse) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CoreDescribeUnhandledFaultResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CoreDescribeUnhandledFaultResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CoreDescribeUnhandledFaultResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CoreDescribeUnhandledFaultResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFailureMessage

`func (o *CoreDescribeUnhandledFaultResponse) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *CoreDescribeUnhandledFaultResponse) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *CoreDescribeUnhandledFaultResponse) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *CoreDescribeUnhandledFaultResponse) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetFailureType

`func (o *CoreDescribeUnhandledFaultResponse) GetFailureType() string`

GetFailureType returns the FailureType field if non-nil, zero value otherwise.

### GetFailureTypeOk

`func (o *CoreDescribeUnhandledFaultResponse) GetFailureTypeOk() (*string, bool)`

GetFailureTypeOk returns a tuple with the FailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureType

`func (o *CoreDescribeUnhandledFaultResponse) SetFailureType(v string)`

SetFailureType sets FailureType field to given value.

### HasFailureType

`func (o *CoreDescribeUnhandledFaultResponse) HasFailureType() bool`

HasFailureType returns a boolean if a field has been set.

### GetInstance

`func (o *CoreDescribeUnhandledFaultResponse) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *CoreDescribeUnhandledFaultResponse) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *CoreDescribeUnhandledFaultResponse) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *CoreDescribeUnhandledFaultResponse) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetProperties

`func (o *CoreDescribeUnhandledFaultResponse) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CoreDescribeUnhandledFaultResponse) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CoreDescribeUnhandledFaultResponse) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CoreDescribeUnhandledFaultResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetService

`func (o *CoreDescribeUnhandledFaultResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CoreDescribeUnhandledFaultResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CoreDescribeUnhandledFaultResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CoreDescribeUnhandledFaultResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceType

`func (o *CoreDescribeUnhandledFaultResponse) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CoreDescribeUnhandledFaultResponse) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CoreDescribeUnhandledFaultResponse) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *CoreDescribeUnhandledFaultResponse) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


