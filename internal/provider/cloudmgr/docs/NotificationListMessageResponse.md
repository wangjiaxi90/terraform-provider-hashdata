# NotificationListMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]NotificationDescribeMessageResponse**](NotificationDescribeMessageResponse.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewNotificationListMessageResponse

`func NewNotificationListMessageResponse() *NotificationListMessageResponse`

NewNotificationListMessageResponse instantiates a new NotificationListMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationListMessageResponseWithDefaults

`func NewNotificationListMessageResponseWithDefaults() *NotificationListMessageResponse`

NewNotificationListMessageResponseWithDefaults instantiates a new NotificationListMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *NotificationListMessageResponse) GetContent() []NotificationDescribeMessageResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NotificationListMessageResponse) GetContentOk() (*[]NotificationDescribeMessageResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NotificationListMessageResponse) SetContent(v []NotificationDescribeMessageResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *NotificationListMessageResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCount

`func (o *NotificationListMessageResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NotificationListMessageResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NotificationListMessageResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *NotificationListMessageResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *NotificationListMessageResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NotificationListMessageResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NotificationListMessageResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *NotificationListMessageResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


