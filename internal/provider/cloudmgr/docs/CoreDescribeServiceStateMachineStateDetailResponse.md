# CoreDescribeServiceStateMachineStateDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentState** | Pointer to **string** |  | [optional] 
**ElementStatus** | Pointer to **map[string]interface{}** |  | [optional] 
**Elements** | Pointer to [**[]CoreDescribeComponentStateMachineStateDetailResponse**](CoreDescribeComponentStateMachineStateDetailResponse.md) |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**MachineTime** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**SourceState** | Pointer to **string** |  | [optional] 
**Transition** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreDescribeServiceStateMachineStateDetailResponse

`func NewCoreDescribeServiceStateMachineStateDetailResponse() *CoreDescribeServiceStateMachineStateDetailResponse`

NewCoreDescribeServiceStateMachineStateDetailResponse instantiates a new CoreDescribeServiceStateMachineStateDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeServiceStateMachineStateDetailResponseWithDefaults

`func NewCoreDescribeServiceStateMachineStateDetailResponseWithDefaults() *CoreDescribeServiceStateMachineStateDetailResponse`

NewCoreDescribeServiceStateMachineStateDetailResponseWithDefaults instantiates a new CoreDescribeServiceStateMachineStateDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentState

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetElementStatus

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetElementStatus() map[string]interface{}`

GetElementStatus returns the ElementStatus field if non-nil, zero value otherwise.

### GetElementStatusOk

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetElementStatusOk() (*map[string]interface{}, bool)`

GetElementStatusOk returns a tuple with the ElementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementStatus

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) SetElementStatus(v map[string]interface{})`

SetElementStatus sets ElementStatus field to given value.

### HasElementStatus

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) HasElementStatus() bool`

HasElementStatus returns a boolean if a field has been set.

### GetElements

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetElements() []CoreDescribeComponentStateMachineStateDetailResponse`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetElementsOk() (*[]CoreDescribeComponentStateMachineStateDetailResponse, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) SetElements(v []CoreDescribeComponentStateMachineStateDetailResponse)`

SetElements sets Elements field to given value.

### HasElements

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetEvent

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetMachineTime

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetMachineTime() string`

GetMachineTime returns the MachineTime field if non-nil, zero value otherwise.

### GetMachineTimeOk

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetMachineTimeOk() (*string, bool)`

GetMachineTimeOk returns a tuple with the MachineTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTime

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) SetMachineTime(v string)`

SetMachineTime sets MachineTime field to given value.

### HasMachineTime

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) HasMachineTime() bool`

HasMachineTime returns a boolean if a field has been set.

### GetService

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSourceState

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetSourceState() string`

GetSourceState returns the SourceState field if non-nil, zero value otherwise.

### GetSourceStateOk

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetSourceStateOk() (*string, bool)`

GetSourceStateOk returns a tuple with the SourceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceState

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) SetSourceState(v string)`

SetSourceState sets SourceState field to given value.

### HasSourceState

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) HasSourceState() bool`

HasSourceState returns a boolean if a field has been set.

### GetTransition

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetTransition() string`

GetTransition returns the Transition field if non-nil, zero value otherwise.

### GetTransitionOk

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) GetTransitionOk() (*string, bool)`

GetTransitionOk returns a tuple with the Transition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransition

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) SetTransition(v string)`

SetTransition sets Transition field to given value.

### HasTransition

`func (o *CoreDescribeServiceStateMachineStateDetailResponse) HasTransition() bool`

HasTransition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


