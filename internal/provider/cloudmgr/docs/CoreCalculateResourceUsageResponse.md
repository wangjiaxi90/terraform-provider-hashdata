# CoreCalculateResourceUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**CommonResourceConfig**](CommonResourceConfig.md) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewCoreCalculateResourceUsageResponse

`func NewCoreCalculateResourceUsageResponse() *CoreCalculateResourceUsageResponse`

NewCoreCalculateResourceUsageResponse instantiates a new CoreCalculateResourceUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCalculateResourceUsageResponseWithDefaults

`func NewCoreCalculateResourceUsageResponseWithDefaults() *CoreCalculateResourceUsageResponse`

NewCoreCalculateResourceUsageResponseWithDefaults instantiates a new CoreCalculateResourceUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CoreCalculateResourceUsageResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CoreCalculateResourceUsageResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CoreCalculateResourceUsageResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CoreCalculateResourceUsageResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResult

`func (o *CoreCalculateResourceUsageResponse) GetResult() CommonResourceConfig`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CoreCalculateResourceUsageResponse) GetResultOk() (*CommonResourceConfig, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CoreCalculateResourceUsageResponse) SetResult(v CommonResourceConfig)`

SetResult sets Result field to given value.

### HasResult

`func (o *CoreCalculateResourceUsageResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *CoreCalculateResourceUsageResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CoreCalculateResourceUsageResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CoreCalculateResourceUsageResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CoreCalculateResourceUsageResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


