# AccountModifyUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ExpiredAt** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**IsSupervisor** | Pointer to **bool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**PasswordExpiredAt** | Pointer to **string** |  | [optional] 
**Surname** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountModifyUserRequest

`func NewAccountModifyUserRequest() *AccountModifyUserRequest`

NewAccountModifyUserRequest instantiates a new AccountModifyUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountModifyUserRequestWithDefaults

`func NewAccountModifyUserRequestWithDefaults() *AccountModifyUserRequest`

NewAccountModifyUserRequestWithDefaults instantiates a new AccountModifyUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AccountModifyUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountModifyUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountModifyUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountModifyUserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *AccountModifyUserRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AccountModifyUserRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AccountModifyUserRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AccountModifyUserRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpiredAt

`func (o *AccountModifyUserRequest) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *AccountModifyUserRequest) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *AccountModifyUserRequest) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *AccountModifyUserRequest) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetGivenName

`func (o *AccountModifyUserRequest) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *AccountModifyUserRequest) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *AccountModifyUserRequest) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *AccountModifyUserRequest) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetIsAdmin

`func (o *AccountModifyUserRequest) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *AccountModifyUserRequest) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *AccountModifyUserRequest) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *AccountModifyUserRequest) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsSupervisor

`func (o *AccountModifyUserRequest) GetIsSupervisor() bool`

GetIsSupervisor returns the IsSupervisor field if non-nil, zero value otherwise.

### GetIsSupervisorOk

`func (o *AccountModifyUserRequest) GetIsSupervisorOk() (*bool, bool)`

GetIsSupervisorOk returns a tuple with the IsSupervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupervisor

`func (o *AccountModifyUserRequest) SetIsSupervisor(v bool)`

SetIsSupervisor sets IsSupervisor field to given value.

### HasIsSupervisor

`func (o *AccountModifyUserRequest) HasIsSupervisor() bool`

HasIsSupervisor returns a boolean if a field has been set.

### GetLocked

`func (o *AccountModifyUserRequest) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *AccountModifyUserRequest) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *AccountModifyUserRequest) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *AccountModifyUserRequest) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPasswordExpiredAt

`func (o *AccountModifyUserRequest) GetPasswordExpiredAt() string`

GetPasswordExpiredAt returns the PasswordExpiredAt field if non-nil, zero value otherwise.

### GetPasswordExpiredAtOk

`func (o *AccountModifyUserRequest) GetPasswordExpiredAtOk() (*string, bool)`

GetPasswordExpiredAtOk returns a tuple with the PasswordExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiredAt

`func (o *AccountModifyUserRequest) SetPasswordExpiredAt(v string)`

SetPasswordExpiredAt sets PasswordExpiredAt field to given value.

### HasPasswordExpiredAt

`func (o *AccountModifyUserRequest) HasPasswordExpiredAt() bool`

HasPasswordExpiredAt returns a boolean if a field has been set.

### GetSurname

`func (o *AccountModifyUserRequest) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *AccountModifyUserRequest) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *AccountModifyUserRequest) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *AccountModifyUserRequest) HasSurname() bool`

HasSurname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


