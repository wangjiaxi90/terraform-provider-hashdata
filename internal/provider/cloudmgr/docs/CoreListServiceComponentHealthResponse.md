# CoreListServiceComponentHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentHealthStatus** | Pointer to **string** |  | [optional] 
**ComponentHealthStatusDescription** | Pointer to **string** |  | [optional] 
**ComponentType** | Pointer to **string** |  | [optional] 
**Content** | Pointer to [**[]CoreDescribeComponentHealthDetailResponse**](CoreDescribeComponentHealthDetailResponse.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**HealthStatus** | Pointer to **string** |  | [optional] 
**HealthStatusDescription** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewCoreListServiceComponentHealthResponse

`func NewCoreListServiceComponentHealthResponse() *CoreListServiceComponentHealthResponse`

NewCoreListServiceComponentHealthResponse instantiates a new CoreListServiceComponentHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreListServiceComponentHealthResponseWithDefaults

`func NewCoreListServiceComponentHealthResponseWithDefaults() *CoreListServiceComponentHealthResponse`

NewCoreListServiceComponentHealthResponseWithDefaults instantiates a new CoreListServiceComponentHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentHealthStatus

`func (o *CoreListServiceComponentHealthResponse) GetComponentHealthStatus() string`

GetComponentHealthStatus returns the ComponentHealthStatus field if non-nil, zero value otherwise.

### GetComponentHealthStatusOk

`func (o *CoreListServiceComponentHealthResponse) GetComponentHealthStatusOk() (*string, bool)`

GetComponentHealthStatusOk returns a tuple with the ComponentHealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentHealthStatus

`func (o *CoreListServiceComponentHealthResponse) SetComponentHealthStatus(v string)`

SetComponentHealthStatus sets ComponentHealthStatus field to given value.

### HasComponentHealthStatus

`func (o *CoreListServiceComponentHealthResponse) HasComponentHealthStatus() bool`

HasComponentHealthStatus returns a boolean if a field has been set.

### GetComponentHealthStatusDescription

`func (o *CoreListServiceComponentHealthResponse) GetComponentHealthStatusDescription() string`

GetComponentHealthStatusDescription returns the ComponentHealthStatusDescription field if non-nil, zero value otherwise.

### GetComponentHealthStatusDescriptionOk

`func (o *CoreListServiceComponentHealthResponse) GetComponentHealthStatusDescriptionOk() (*string, bool)`

GetComponentHealthStatusDescriptionOk returns a tuple with the ComponentHealthStatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentHealthStatusDescription

`func (o *CoreListServiceComponentHealthResponse) SetComponentHealthStatusDescription(v string)`

SetComponentHealthStatusDescription sets ComponentHealthStatusDescription field to given value.

### HasComponentHealthStatusDescription

`func (o *CoreListServiceComponentHealthResponse) HasComponentHealthStatusDescription() bool`

HasComponentHealthStatusDescription returns a boolean if a field has been set.

### GetComponentType

`func (o *CoreListServiceComponentHealthResponse) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *CoreListServiceComponentHealthResponse) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *CoreListServiceComponentHealthResponse) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *CoreListServiceComponentHealthResponse) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetContent

`func (o *CoreListServiceComponentHealthResponse) GetContent() []CoreDescribeComponentHealthDetailResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CoreListServiceComponentHealthResponse) GetContentOk() (*[]CoreDescribeComponentHealthDetailResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CoreListServiceComponentHealthResponse) SetContent(v []CoreDescribeComponentHealthDetailResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *CoreListServiceComponentHealthResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCount

`func (o *CoreListServiceComponentHealthResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CoreListServiceComponentHealthResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CoreListServiceComponentHealthResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CoreListServiceComponentHealthResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHealthStatus

`func (o *CoreListServiceComponentHealthResponse) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *CoreListServiceComponentHealthResponse) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *CoreListServiceComponentHealthResponse) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *CoreListServiceComponentHealthResponse) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetHealthStatusDescription

`func (o *CoreListServiceComponentHealthResponse) GetHealthStatusDescription() string`

GetHealthStatusDescription returns the HealthStatusDescription field if non-nil, zero value otherwise.

### GetHealthStatusDescriptionOk

`func (o *CoreListServiceComponentHealthResponse) GetHealthStatusDescriptionOk() (*string, bool)`

GetHealthStatusDescriptionOk returns a tuple with the HealthStatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatusDescription

`func (o *CoreListServiceComponentHealthResponse) SetHealthStatusDescription(v string)`

SetHealthStatusDescription sets HealthStatusDescription field to given value.

### HasHealthStatusDescription

`func (o *CoreListServiceComponentHealthResponse) HasHealthStatusDescription() bool`

HasHealthStatusDescription returns a boolean if a field has been set.

### GetId

`func (o *CoreListServiceComponentHealthResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreListServiceComponentHealthResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreListServiceComponentHealthResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CoreListServiceComponentHealthResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTotal

`func (o *CoreListServiceComponentHealthResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CoreListServiceComponentHealthResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CoreListServiceComponentHealthResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CoreListServiceComponentHealthResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


