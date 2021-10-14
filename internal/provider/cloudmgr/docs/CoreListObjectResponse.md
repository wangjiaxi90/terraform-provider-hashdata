# CoreListObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Marker** | Pointer to **string** |  | [optional] 
**NextMarker** | Pointer to **string** |  | [optional] 
**Objects** | Pointer to [**[]CoreDescribeObjectResponse**](CoreDescribeObjectResponse.md) |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreListObjectResponse

`func NewCoreListObjectResponse() *CoreListObjectResponse`

NewCoreListObjectResponse instantiates a new CoreListObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreListObjectResponseWithDefaults

`func NewCoreListObjectResponseWithDefaults() *CoreListObjectResponse`

NewCoreListObjectResponseWithDefaults instantiates a new CoreListObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarker

`func (o *CoreListObjectResponse) GetMarker() string`

GetMarker returns the Marker field if non-nil, zero value otherwise.

### GetMarkerOk

`func (o *CoreListObjectResponse) GetMarkerOk() (*string, bool)`

GetMarkerOk returns a tuple with the Marker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarker

`func (o *CoreListObjectResponse) SetMarker(v string)`

SetMarker sets Marker field to given value.

### HasMarker

`func (o *CoreListObjectResponse) HasMarker() bool`

HasMarker returns a boolean if a field has been set.

### GetNextMarker

`func (o *CoreListObjectResponse) GetNextMarker() string`

GetNextMarker returns the NextMarker field if non-nil, zero value otherwise.

### GetNextMarkerOk

`func (o *CoreListObjectResponse) GetNextMarkerOk() (*string, bool)`

GetNextMarkerOk returns a tuple with the NextMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextMarker

`func (o *CoreListObjectResponse) SetNextMarker(v string)`

SetNextMarker sets NextMarker field to given value.

### HasNextMarker

`func (o *CoreListObjectResponse) HasNextMarker() bool`

HasNextMarker returns a boolean if a field has been set.

### GetObjects

`func (o *CoreListObjectResponse) GetObjects() []CoreDescribeObjectResponse`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *CoreListObjectResponse) GetObjectsOk() (*[]CoreDescribeObjectResponse, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *CoreListObjectResponse) SetObjects(v []CoreDescribeObjectResponse)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *CoreListObjectResponse) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetPrefix

`func (o *CoreListObjectResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CoreListObjectResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CoreListObjectResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *CoreListObjectResponse) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


