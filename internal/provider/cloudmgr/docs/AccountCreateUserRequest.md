# AccountCreateUserRequest

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
**Password** | Pointer to **string** |  | [optional] 
**PasswordExpiredAt** | Pointer to **string** |  | [optional] 
**Surname** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountCreateUserRequest

`func NewAccountCreateUserRequest() *AccountCreateUserRequest`

NewAccountCreateUserRequest instantiates a new AccountCreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreateUserRequestWithDefaults

`func NewAccountCreateUserRequestWithDefaults() *AccountCreateUserRequest`

NewAccountCreateUserRequestWithDefaults instantiates a new AccountCreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AccountCreateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountCreateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountCreateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountCreateUserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *AccountCreateUserRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AccountCreateUserRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AccountCreateUserRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AccountCreateUserRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpiredAt

`func (o *AccountCreateUserRequest) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *AccountCreateUserRequest) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *AccountCreateUserRequest) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *AccountCreateUserRequest) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetGivenName

`func (o *AccountCreateUserRequest) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *AccountCreateUserRequest) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *AccountCreateUserRequest) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *AccountCreateUserRequest) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetIsAdmin

`func (o *AccountCreateUserRequest) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *AccountCreateUserRequest) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *AccountCreateUserRequest) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *AccountCreateUserRequest) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsSupervisor

`func (o *AccountCreateUserRequest) GetIsSupervisor() bool`

GetIsSupervisor returns the IsSupervisor field if non-nil, zero value otherwise.

### GetIsSupervisorOk

`func (o *AccountCreateUserRequest) GetIsSupervisorOk() (*bool, bool)`

GetIsSupervisorOk returns a tuple with the IsSupervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupervisor

`func (o *AccountCreateUserRequest) SetIsSupervisor(v bool)`

SetIsSupervisor sets IsSupervisor field to given value.

### HasIsSupervisor

`func (o *AccountCreateUserRequest) HasIsSupervisor() bool`

HasIsSupervisor returns a boolean if a field has been set.

### GetLocked

`func (o *AccountCreateUserRequest) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *AccountCreateUserRequest) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *AccountCreateUserRequest) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *AccountCreateUserRequest) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPassword

`func (o *AccountCreateUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AccountCreateUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AccountCreateUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AccountCreateUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordExpiredAt

`func (o *AccountCreateUserRequest) GetPasswordExpiredAt() string`

GetPasswordExpiredAt returns the PasswordExpiredAt field if non-nil, zero value otherwise.

### GetPasswordExpiredAtOk

`func (o *AccountCreateUserRequest) GetPasswordExpiredAtOk() (*string, bool)`

GetPasswordExpiredAtOk returns a tuple with the PasswordExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiredAt

`func (o *AccountCreateUserRequest) SetPasswordExpiredAt(v string)`

SetPasswordExpiredAt sets PasswordExpiredAt field to given value.

### HasPasswordExpiredAt

`func (o *AccountCreateUserRequest) HasPasswordExpiredAt() bool`

HasPasswordExpiredAt returns a boolean if a field has been set.

### GetSurname

`func (o *AccountCreateUserRequest) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *AccountCreateUserRequest) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *AccountCreateUserRequest) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *AccountCreateUserRequest) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetTenant

`func (o *AccountCreateUserRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AccountCreateUserRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AccountCreateUserRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AccountCreateUserRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUsername

`func (o *AccountCreateUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AccountCreateUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AccountCreateUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AccountCreateUserRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


