# VmalertCreateAlertRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactGroups** | Pointer to **[]string** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Exceed** | Pointer to **bool** |  | [optional] 
**Metrics** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 

## Methods

### NewVmalertCreateAlertRuleRequest

`func NewVmalertCreateAlertRuleRequest() *VmalertCreateAlertRuleRequest`

NewVmalertCreateAlertRuleRequest instantiates a new VmalertCreateAlertRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmalertCreateAlertRuleRequestWithDefaults

`func NewVmalertCreateAlertRuleRequestWithDefaults() *VmalertCreateAlertRuleRequest`

NewVmalertCreateAlertRuleRequestWithDefaults instantiates a new VmalertCreateAlertRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactGroups

`func (o *VmalertCreateAlertRuleRequest) GetContactGroups() []string`

GetContactGroups returns the ContactGroups field if non-nil, zero value otherwise.

### GetContactGroupsOk

`func (o *VmalertCreateAlertRuleRequest) GetContactGroupsOk() (*[]string, bool)`

GetContactGroupsOk returns a tuple with the ContactGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactGroups

`func (o *VmalertCreateAlertRuleRequest) SetContactGroups(v []string)`

SetContactGroups sets ContactGroups field to given value.

### HasContactGroups

`func (o *VmalertCreateAlertRuleRequest) HasContactGroups() bool`

HasContactGroups returns a boolean if a field has been set.

### GetDuration

`func (o *VmalertCreateAlertRuleRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VmalertCreateAlertRuleRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VmalertCreateAlertRuleRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VmalertCreateAlertRuleRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetExceed

`func (o *VmalertCreateAlertRuleRequest) GetExceed() bool`

GetExceed returns the Exceed field if non-nil, zero value otherwise.

### GetExceedOk

`func (o *VmalertCreateAlertRuleRequest) GetExceedOk() (*bool, bool)`

GetExceedOk returns a tuple with the Exceed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceed

`func (o *VmalertCreateAlertRuleRequest) SetExceed(v bool)`

SetExceed sets Exceed field to given value.

### HasExceed

`func (o *VmalertCreateAlertRuleRequest) HasExceed() bool`

HasExceed returns a boolean if a field has been set.

### GetMetrics

`func (o *VmalertCreateAlertRuleRequest) GetMetrics() string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *VmalertCreateAlertRuleRequest) GetMetricsOk() (*string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *VmalertCreateAlertRuleRequest) SetMetrics(v string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *VmalertCreateAlertRuleRequest) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetValue

`func (o *VmalertCreateAlertRuleRequest) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VmalertCreateAlertRuleRequest) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VmalertCreateAlertRuleRequest) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *VmalertCreateAlertRuleRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


