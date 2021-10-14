# KsyunDescribeResourceUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]CommonIaasResource**](CommonIaasResource.md) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewKsyunDescribeResourceUsageResponse

`func NewKsyunDescribeResourceUsageResponse() *KsyunDescribeResourceUsageResponse`

NewKsyunDescribeResourceUsageResponse instantiates a new KsyunDescribeResourceUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKsyunDescribeResourceUsageResponseWithDefaults

`func NewKsyunDescribeResourceUsageResponseWithDefaults() *KsyunDescribeResourceUsageResponse`

NewKsyunDescribeResourceUsageResponseWithDefaults instantiates a new KsyunDescribeResourceUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *KsyunDescribeResourceUsageResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KsyunDescribeResourceUsageResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KsyunDescribeResourceUsageResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KsyunDescribeResourceUsageResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResult

`func (o *KsyunDescribeResourceUsageResponse) GetResult() []CommonIaasResource`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *KsyunDescribeResourceUsageResponse) GetResultOk() (*[]CommonIaasResource, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *KsyunDescribeResourceUsageResponse) SetResult(v []CommonIaasResource)`

SetResult sets Result field to given value.

### HasResult

`func (o *KsyunDescribeResourceUsageResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *KsyunDescribeResourceUsageResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KsyunDescribeResourceUsageResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KsyunDescribeResourceUsageResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KsyunDescribeResourceUsageResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


