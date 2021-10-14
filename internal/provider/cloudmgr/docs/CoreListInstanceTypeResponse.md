# CoreListInstanceTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]CoreDescribeInstanceTypeResponse**](CoreDescribeInstanceTypeResponse.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewCoreListInstanceTypeResponse

`func NewCoreListInstanceTypeResponse() *CoreListInstanceTypeResponse`

NewCoreListInstanceTypeResponse instantiates a new CoreListInstanceTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreListInstanceTypeResponseWithDefaults

`func NewCoreListInstanceTypeResponseWithDefaults() *CoreListInstanceTypeResponse`

NewCoreListInstanceTypeResponseWithDefaults instantiates a new CoreListInstanceTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CoreListInstanceTypeResponse) GetContent() []CoreDescribeInstanceTypeResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CoreListInstanceTypeResponse) GetContentOk() (*[]CoreDescribeInstanceTypeResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CoreListInstanceTypeResponse) SetContent(v []CoreDescribeInstanceTypeResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *CoreListInstanceTypeResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCount

`func (o *CoreListInstanceTypeResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CoreListInstanceTypeResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CoreListInstanceTypeResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CoreListInstanceTypeResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *CoreListInstanceTypeResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CoreListInstanceTypeResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CoreListInstanceTypeResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CoreListInstanceTypeResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


