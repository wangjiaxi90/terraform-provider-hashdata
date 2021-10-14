# CoreListResourceUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]CloudmgrcommonIaasResource**](CloudmgrcommonIaasResource.md) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewCoreListResourceUsageResponse

`func NewCoreListResourceUsageResponse() *CoreListResourceUsageResponse`

NewCoreListResourceUsageResponse instantiates a new CoreListResourceUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreListResourceUsageResponseWithDefaults

`func NewCoreListResourceUsageResponseWithDefaults() *CoreListResourceUsageResponse`

NewCoreListResourceUsageResponseWithDefaults instantiates a new CoreListResourceUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CoreListResourceUsageResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CoreListResourceUsageResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CoreListResourceUsageResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CoreListResourceUsageResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResult

`func (o *CoreListResourceUsageResponse) GetResult() []CloudmgrcommonIaasResource`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CoreListResourceUsageResponse) GetResultOk() (*[]CloudmgrcommonIaasResource, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CoreListResourceUsageResponse) SetResult(v []CloudmgrcommonIaasResource)`

SetResult sets Result field to given value.

### HasResult

`func (o *CoreListResourceUsageResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *CoreListResourceUsageResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CoreListResourceUsageResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CoreListResourceUsageResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CoreListResourceUsageResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


