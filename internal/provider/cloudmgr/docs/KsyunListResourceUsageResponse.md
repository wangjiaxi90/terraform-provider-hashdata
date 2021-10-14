# KsyunListResourceUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]CommonIaasResource**](CommonIaasResource.md) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewKsyunListResourceUsageResponse

`func NewKsyunListResourceUsageResponse() *KsyunListResourceUsageResponse`

NewKsyunListResourceUsageResponse instantiates a new KsyunListResourceUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKsyunListResourceUsageResponseWithDefaults

`func NewKsyunListResourceUsageResponseWithDefaults() *KsyunListResourceUsageResponse`

NewKsyunListResourceUsageResponseWithDefaults instantiates a new KsyunListResourceUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *KsyunListResourceUsageResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KsyunListResourceUsageResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KsyunListResourceUsageResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KsyunListResourceUsageResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResult

`func (o *KsyunListResourceUsageResponse) GetResult() []CommonIaasResource`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *KsyunListResourceUsageResponse) GetResultOk() (*[]CommonIaasResource, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *KsyunListResourceUsageResponse) SetResult(v []CommonIaasResource)`

SetResult sets Result field to given value.

### HasResult

`func (o *KsyunListResourceUsageResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *KsyunListResourceUsageResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KsyunListResourceUsageResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KsyunListResourceUsageResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KsyunListResourceUsageResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


