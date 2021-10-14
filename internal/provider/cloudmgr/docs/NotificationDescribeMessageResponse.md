# NotificationDescribeMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Contacts** | Pointer to [**[]NotificationDescribeContactResponse**](NotificationDescribeContactResponse.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**SendAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationDescribeMessageResponse

`func NewNotificationDescribeMessageResponse() *NotificationDescribeMessageResponse`

NewNotificationDescribeMessageResponse instantiates a new NotificationDescribeMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDescribeMessageResponseWithDefaults

`func NewNotificationDescribeMessageResponseWithDefaults() *NotificationDescribeMessageResponse`

NewNotificationDescribeMessageResponseWithDefaults instantiates a new NotificationDescribeMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *NotificationDescribeMessageResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NotificationDescribeMessageResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NotificationDescribeMessageResponse) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NotificationDescribeMessageResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetContacts

`func (o *NotificationDescribeMessageResponse) GetContacts() []NotificationDescribeContactResponse`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *NotificationDescribeMessageResponse) GetContactsOk() (*[]NotificationDescribeContactResponse, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *NotificationDescribeMessageResponse) SetContacts(v []NotificationDescribeContactResponse)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *NotificationDescribeMessageResponse) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetContent

`func (o *NotificationDescribeMessageResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NotificationDescribeMessageResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NotificationDescribeMessageResponse) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *NotificationDescribeMessageResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationDescribeMessageResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationDescribeMessageResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationDescribeMessageResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationDescribeMessageResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetErrorMessage

`func (o *NotificationDescribeMessageResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NotificationDescribeMessageResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NotificationDescribeMessageResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NotificationDescribeMessageResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetId

`func (o *NotificationDescribeMessageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationDescribeMessageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationDescribeMessageResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationDescribeMessageResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSendAt

`func (o *NotificationDescribeMessageResponse) GetSendAt() string`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *NotificationDescribeMessageResponse) GetSendAtOk() (*string, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *NotificationDescribeMessageResponse) SetSendAt(v string)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *NotificationDescribeMessageResponse) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.

### GetStatus

`func (o *NotificationDescribeMessageResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationDescribeMessageResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationDescribeMessageResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NotificationDescribeMessageResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *NotificationDescribeMessageResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *NotificationDescribeMessageResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *NotificationDescribeMessageResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *NotificationDescribeMessageResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


