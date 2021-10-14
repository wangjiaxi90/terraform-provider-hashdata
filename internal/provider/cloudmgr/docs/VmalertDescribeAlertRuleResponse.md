# VmalertDescribeAlertRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactGroups** | Pointer to **[]string** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Exceed** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 

## Methods

### NewVmalertDescribeAlertRuleResponse

`func NewVmalertDescribeAlertRuleResponse() *VmalertDescribeAlertRuleResponse`

NewVmalertDescribeAlertRuleResponse instantiates a new VmalertDescribeAlertRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmalertDescribeAlertRuleResponseWithDefaults

`func NewVmalertDescribeAlertRuleResponseWithDefaults() *VmalertDescribeAlertRuleResponse`

NewVmalertDescribeAlertRuleResponseWithDefaults instantiates a new VmalertDescribeAlertRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactGroups

`func (o *VmalertDescribeAlertRuleResponse) GetContactGroups() []string`

GetContactGroups returns the ContactGroups field if non-nil, zero value otherwise.

### GetContactGroupsOk

`func (o *VmalertDescribeAlertRuleResponse) GetContactGroupsOk() (*[]string, bool)`

GetContactGroupsOk returns a tuple with the ContactGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactGroups

`func (o *VmalertDescribeAlertRuleResponse) SetContactGroups(v []string)`

SetContactGroups sets ContactGroups field to given value.

### HasContactGroups

`func (o *VmalertDescribeAlertRuleResponse) HasContactGroups() bool`

HasContactGroups returns a boolean if a field has been set.

### GetDuration

`func (o *VmalertDescribeAlertRuleResponse) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VmalertDescribeAlertRuleResponse) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VmalertDescribeAlertRuleResponse) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VmalertDescribeAlertRuleResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetExceed

`func (o *VmalertDescribeAlertRuleResponse) GetExceed() bool`

GetExceed returns the Exceed field if non-nil, zero value otherwise.

### GetExceedOk

`func (o *VmalertDescribeAlertRuleResponse) GetExceedOk() (*bool, bool)`

GetExceedOk returns a tuple with the Exceed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceed

`func (o *VmalertDescribeAlertRuleResponse) SetExceed(v bool)`

SetExceed sets Exceed field to given value.

### HasExceed

`func (o *VmalertDescribeAlertRuleResponse) HasExceed() bool`

HasExceed returns a boolean if a field has been set.

### GetId

`func (o *VmalertDescribeAlertRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmalertDescribeAlertRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmalertDescribeAlertRuleResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VmalertDescribeAlertRuleResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetrics

`func (o *VmalertDescribeAlertRuleResponse) GetMetrics() string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *VmalertDescribeAlertRuleResponse) GetMetricsOk() (*string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *VmalertDescribeAlertRuleResponse) SetMetrics(v string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *VmalertDescribeAlertRuleResponse) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetService

`func (o *VmalertDescribeAlertRuleResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *VmalertDescribeAlertRuleResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *VmalertDescribeAlertRuleResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *VmalertDescribeAlertRuleResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VmalertDescribeAlertRuleResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VmalertDescribeAlertRuleResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VmalertDescribeAlertRuleResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VmalertDescribeAlertRuleResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *VmalertDescribeAlertRuleResponse) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VmalertDescribeAlertRuleResponse) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VmalertDescribeAlertRuleResponse) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *VmalertDescribeAlertRuleResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


