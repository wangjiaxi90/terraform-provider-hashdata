# NotificationAddContactToGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contacts** | Pointer to [**[]NotificationDescribeContactResponse**](NotificationDescribeContactResponse.md) |  | [optional] 

## Methods

### NewNotificationAddContactToGroupRequest

`func NewNotificationAddContactToGroupRequest() *NotificationAddContactToGroupRequest`

NewNotificationAddContactToGroupRequest instantiates a new NotificationAddContactToGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationAddContactToGroupRequestWithDefaults

`func NewNotificationAddContactToGroupRequestWithDefaults() *NotificationAddContactToGroupRequest`

NewNotificationAddContactToGroupRequestWithDefaults instantiates a new NotificationAddContactToGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContacts

`func (o *NotificationAddContactToGroupRequest) GetContacts() []NotificationDescribeContactResponse`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *NotificationAddContactToGroupRequest) GetContactsOk() (*[]NotificationDescribeContactResponse, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *NotificationAddContactToGroupRequest) SetContacts(v []NotificationDescribeContactResponse)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *NotificationAddContactToGroupRequest) HasContacts() bool`

HasContacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


