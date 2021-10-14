# CoreDescribeFaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**FailureMessage** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**RecoveryId** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreDescribeFaultResponse

`func NewCoreDescribeFaultResponse() *CoreDescribeFaultResponse`

NewCoreDescribeFaultResponse instantiates a new CoreDescribeFaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeFaultResponseWithDefaults

`func NewCoreDescribeFaultResponseWithDefaults() *CoreDescribeFaultResponse`

NewCoreDescribeFaultResponseWithDefaults instantiates a new CoreDescribeFaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentType

`func (o *CoreDescribeFaultResponse) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *CoreDescribeFaultResponse) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *CoreDescribeFaultResponse) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *CoreDescribeFaultResponse) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CoreDescribeFaultResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CoreDescribeFaultResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CoreDescribeFaultResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CoreDescribeFaultResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFailureMessage

`func (o *CoreDescribeFaultResponse) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *CoreDescribeFaultResponse) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *CoreDescribeFaultResponse) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *CoreDescribeFaultResponse) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetId

`func (o *CoreDescribeFaultResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreDescribeFaultResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreDescribeFaultResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CoreDescribeFaultResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstance

`func (o *CoreDescribeFaultResponse) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *CoreDescribeFaultResponse) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *CoreDescribeFaultResponse) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *CoreDescribeFaultResponse) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetMessage

`func (o *CoreDescribeFaultResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CoreDescribeFaultResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CoreDescribeFaultResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CoreDescribeFaultResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetProperties

`func (o *CoreDescribeFaultResponse) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CoreDescribeFaultResponse) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CoreDescribeFaultResponse) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CoreDescribeFaultResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRecoveryId

`func (o *CoreDescribeFaultResponse) GetRecoveryId() string`

GetRecoveryId returns the RecoveryId field if non-nil, zero value otherwise.

### GetRecoveryIdOk

`func (o *CoreDescribeFaultResponse) GetRecoveryIdOk() (*string, bool)`

GetRecoveryIdOk returns a tuple with the RecoveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryId

`func (o *CoreDescribeFaultResponse) SetRecoveryId(v string)`

SetRecoveryId sets RecoveryId field to given value.

### HasRecoveryId

`func (o *CoreDescribeFaultResponse) HasRecoveryId() bool`

HasRecoveryId returns a boolean if a field has been set.

### GetService

`func (o *CoreDescribeFaultResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CoreDescribeFaultResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CoreDescribeFaultResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CoreDescribeFaultResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceType

`func (o *CoreDescribeFaultResponse) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CoreDescribeFaultResponse) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CoreDescribeFaultResponse) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *CoreDescribeFaultResponse) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStatus

`func (o *CoreDescribeFaultResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CoreDescribeFaultResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CoreDescribeFaultResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CoreDescribeFaultResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CoreDescribeFaultResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CoreDescribeFaultResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CoreDescribeFaultResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CoreDescribeFaultResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


