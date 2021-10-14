# CoreListFaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]CoreDescribeFaultResponse**](CoreDescribeFaultResponse.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewCoreListFaultResponse

`func NewCoreListFaultResponse() *CoreListFaultResponse`

NewCoreListFaultResponse instantiates a new CoreListFaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreListFaultResponseWithDefaults

`func NewCoreListFaultResponseWithDefaults() *CoreListFaultResponse`

NewCoreListFaultResponseWithDefaults instantiates a new CoreListFaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CoreListFaultResponse) GetContent() []CoreDescribeFaultResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CoreListFaultResponse) GetContentOk() (*[]CoreDescribeFaultResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CoreListFaultResponse) SetContent(v []CoreDescribeFaultResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *CoreListFaultResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCount

`func (o *CoreListFaultResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CoreListFaultResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CoreListFaultResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CoreListFaultResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *CoreListFaultResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CoreListFaultResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CoreListFaultResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CoreListFaultResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


