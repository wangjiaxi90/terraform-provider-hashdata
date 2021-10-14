# NotificationCreateContactGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contacts** | Pointer to [**[]NotificationDescribeContactResponse**](NotificationDescribeContactResponse.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationCreateContactGroupRequest

`func NewNotificationCreateContactGroupRequest() *NotificationCreateContactGroupRequest`

NewNotificationCreateContactGroupRequest instantiates a new NotificationCreateContactGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationCreateContactGroupRequestWithDefaults

`func NewNotificationCreateContactGroupRequestWithDefaults() *NotificationCreateContactGroupRequest`

NewNotificationCreateContactGroupRequestWithDefaults instantiates a new NotificationCreateContactGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContacts

`func (o *NotificationCreateContactGroupRequest) GetContacts() []NotificationDescribeContactResponse`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *NotificationCreateContactGroupRequest) GetContactsOk() (*[]NotificationDescribeContactResponse, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *NotificationCreateContactGroupRequest) SetContacts(v []NotificationDescribeContactResponse)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *NotificationCreateContactGroupRequest) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetName

`func (o *NotificationCreateContactGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationCreateContactGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationCreateContactGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationCreateContactGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


