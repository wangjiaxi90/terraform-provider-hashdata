# CoreDescribeComponentStateMachineStateDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] 
**CurrentState** | Pointer to **string** |  | [optional] 
**ElementStatus** | Pointer to **map[string]interface{}** |  | [optional] 
**Elements** | Pointer to [**[]CoreDescribeNodeStateMachineStateDetailResponse**](CoreDescribeNodeStateMachineStateDetailResponse.md) |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**MachineTime** | Pointer to **string** |  | [optional] 
**SourceState** | Pointer to **string** |  | [optional] 
**Transition** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreDescribeComponentStateMachineStateDetailResponse

`func NewCoreDescribeComponentStateMachineStateDetailResponse() *CoreDescribeComponentStateMachineStateDetailResponse`

NewCoreDescribeComponentStateMachineStateDetailResponse instantiates a new CoreDescribeComponentStateMachineStateDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeComponentStateMachineStateDetailResponseWithDefaults

`func NewCoreDescribeComponentStateMachineStateDetailResponseWithDefaults() *CoreDescribeComponentStateMachineStateDetailResponse`

NewCoreDescribeComponentStateMachineStateDetailResponseWithDefaults instantiates a new CoreDescribeComponentStateMachineStateDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCurrentState

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetElementStatus

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetElementStatus() map[string]interface{}`

GetElementStatus returns the ElementStatus field if non-nil, zero value otherwise.

### GetElementStatusOk

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetElementStatusOk() (*map[string]interface{}, bool)`

GetElementStatusOk returns a tuple with the ElementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementStatus

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) SetElementStatus(v map[string]interface{})`

SetElementStatus sets ElementStatus field to given value.

### HasElementStatus

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) HasElementStatus() bool`

HasElementStatus returns a boolean if a field has been set.

### GetElements

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetElements() []CoreDescribeNodeStateMachineStateDetailResponse`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetElementsOk() (*[]CoreDescribeNodeStateMachineStateDetailResponse, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) SetElements(v []CoreDescribeNodeStateMachineStateDetailResponse)`

SetElements sets Elements field to given value.

### HasElements

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetEvent

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetMachineTime

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetMachineTime() string`

GetMachineTime returns the MachineTime field if non-nil, zero value otherwise.

### GetMachineTimeOk

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetMachineTimeOk() (*string, bool)`

GetMachineTimeOk returns a tuple with the MachineTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTime

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) SetMachineTime(v string)`

SetMachineTime sets MachineTime field to given value.

### HasMachineTime

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) HasMachineTime() bool`

HasMachineTime returns a boolean if a field has been set.

### GetSourceState

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetSourceState() string`

GetSourceState returns the SourceState field if non-nil, zero value otherwise.

### GetSourceStateOk

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetSourceStateOk() (*string, bool)`

GetSourceStateOk returns a tuple with the SourceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceState

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) SetSourceState(v string)`

SetSourceState sets SourceState field to given value.

### HasSourceState

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) HasSourceState() bool`

HasSourceState returns a boolean if a field has been set.

### GetTransition

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetTransition() string`

GetTransition returns the Transition field if non-nil, zero value otherwise.

### GetTransitionOk

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) GetTransitionOk() (*string, bool)`

GetTransitionOk returns a tuple with the Transition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransition

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) SetTransition(v string)`

SetTransition sets Transition field to given value.

### HasTransition

`func (o *CoreDescribeComponentStateMachineStateDetailResponse) HasTransition() bool`

HasTransition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


