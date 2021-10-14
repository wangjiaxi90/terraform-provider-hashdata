# CoreDescribeNodeStateMachineStateDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentState** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to **string** |  | [optional] 
**MachineTime** | Pointer to **string** |  | [optional] 
**SourceState** | Pointer to **string** |  | [optional] 
**StableRemaining** | Pointer to **int32** |  | [optional] 
**TimeoutRemaining** | Pointer to **int32** |  | [optional] 
**Transition** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreDescribeNodeStateMachineStateDetailResponse

`func NewCoreDescribeNodeStateMachineStateDetailResponse() *CoreDescribeNodeStateMachineStateDetailResponse`

NewCoreDescribeNodeStateMachineStateDetailResponse instantiates a new CoreDescribeNodeStateMachineStateDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeNodeStateMachineStateDetailResponseWithDefaults

`func NewCoreDescribeNodeStateMachineStateDetailResponseWithDefaults() *CoreDescribeNodeStateMachineStateDetailResponse`

NewCoreDescribeNodeStateMachineStateDetailResponseWithDefaults instantiates a new CoreDescribeNodeStateMachineStateDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentState

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetEvent

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetInstance

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetMachineTime

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetMachineTime() string`

GetMachineTime returns the MachineTime field if non-nil, zero value otherwise.

### GetMachineTimeOk

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetMachineTimeOk() (*string, bool)`

GetMachineTimeOk returns a tuple with the MachineTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTime

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) SetMachineTime(v string)`

SetMachineTime sets MachineTime field to given value.

### HasMachineTime

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) HasMachineTime() bool`

HasMachineTime returns a boolean if a field has been set.

### GetSourceState

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetSourceState() string`

GetSourceState returns the SourceState field if non-nil, zero value otherwise.

### GetSourceStateOk

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetSourceStateOk() (*string, bool)`

GetSourceStateOk returns a tuple with the SourceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceState

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) SetSourceState(v string)`

SetSourceState sets SourceState field to given value.

### HasSourceState

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) HasSourceState() bool`

HasSourceState returns a boolean if a field has been set.

### GetStableRemaining

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetStableRemaining() int32`

GetStableRemaining returns the StableRemaining field if non-nil, zero value otherwise.

### GetStableRemainingOk

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetStableRemainingOk() (*int32, bool)`

GetStableRemainingOk returns a tuple with the StableRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableRemaining

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) SetStableRemaining(v int32)`

SetStableRemaining sets StableRemaining field to given value.

### HasStableRemaining

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) HasStableRemaining() bool`

HasStableRemaining returns a boolean if a field has been set.

### GetTimeoutRemaining

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetTimeoutRemaining() int32`

GetTimeoutRemaining returns the TimeoutRemaining field if non-nil, zero value otherwise.

### GetTimeoutRemainingOk

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetTimeoutRemainingOk() (*int32, bool)`

GetTimeoutRemainingOk returns a tuple with the TimeoutRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutRemaining

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) SetTimeoutRemaining(v int32)`

SetTimeoutRemaining sets TimeoutRemaining field to given value.

### HasTimeoutRemaining

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) HasTimeoutRemaining() bool`

HasTimeoutRemaining returns a boolean if a field has been set.

### GetTransition

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetTransition() string`

GetTransition returns the Transition field if non-nil, zero value otherwise.

### GetTransitionOk

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) GetTransitionOk() (*string, bool)`

GetTransitionOk returns a tuple with the Transition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransition

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) SetTransition(v string)`

SetTransition sets Transition field to given value.

### HasTransition

`func (o *CoreDescribeNodeStateMachineStateDetailResponse) HasTransition() bool`

HasTransition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


