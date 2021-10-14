# AccountDescribeUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DestroyedAt** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ExpiredAt** | Pointer to **string** |  | [optional] 
**External** | Pointer to **bool** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**PasswordExpiredAt** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Supervisor** | Pointer to **bool** |  | [optional] 
**Surname** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountDescribeUserResponse

`func NewAccountDescribeUserResponse() *AccountDescribeUserResponse`

NewAccountDescribeUserResponse instantiates a new AccountDescribeUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDescribeUserResponseWithDefaults

`func NewAccountDescribeUserResponseWithDefaults() *AccountDescribeUserResponse`

NewAccountDescribeUserResponseWithDefaults instantiates a new AccountDescribeUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *AccountDescribeUserResponse) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *AccountDescribeUserResponse) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *AccountDescribeUserResponse) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *AccountDescribeUserResponse) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountDescribeUserResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountDescribeUserResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountDescribeUserResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountDescribeUserResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDestroyedAt

`func (o *AccountDescribeUserResponse) GetDestroyedAt() string`

GetDestroyedAt returns the DestroyedAt field if non-nil, zero value otherwise.

### GetDestroyedAtOk

`func (o *AccountDescribeUserResponse) GetDestroyedAtOk() (*string, bool)`

GetDestroyedAtOk returns a tuple with the DestroyedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyedAt

`func (o *AccountDescribeUserResponse) SetDestroyedAt(v string)`

SetDestroyedAt sets DestroyedAt field to given value.

### HasDestroyedAt

`func (o *AccountDescribeUserResponse) HasDestroyedAt() bool`

HasDestroyedAt returns a boolean if a field has been set.

### GetEmail

`func (o *AccountDescribeUserResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountDescribeUserResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountDescribeUserResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountDescribeUserResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *AccountDescribeUserResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AccountDescribeUserResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AccountDescribeUserResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AccountDescribeUserResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpiredAt

`func (o *AccountDescribeUserResponse) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *AccountDescribeUserResponse) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *AccountDescribeUserResponse) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *AccountDescribeUserResponse) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetExternal

`func (o *AccountDescribeUserResponse) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *AccountDescribeUserResponse) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *AccountDescribeUserResponse) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *AccountDescribeUserResponse) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetExternalId

`func (o *AccountDescribeUserResponse) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AccountDescribeUserResponse) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AccountDescribeUserResponse) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AccountDescribeUserResponse) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetGivenName

`func (o *AccountDescribeUserResponse) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *AccountDescribeUserResponse) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *AccountDescribeUserResponse) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *AccountDescribeUserResponse) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetId

`func (o *AccountDescribeUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountDescribeUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountDescribeUserResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountDescribeUserResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocked

`func (o *AccountDescribeUserResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *AccountDescribeUserResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *AccountDescribeUserResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *AccountDescribeUserResponse) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPasswordExpiredAt

`func (o *AccountDescribeUserResponse) GetPasswordExpiredAt() string`

GetPasswordExpiredAt returns the PasswordExpiredAt field if non-nil, zero value otherwise.

### GetPasswordExpiredAtOk

`func (o *AccountDescribeUserResponse) GetPasswordExpiredAtOk() (*string, bool)`

GetPasswordExpiredAtOk returns a tuple with the PasswordExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiredAt

`func (o *AccountDescribeUserResponse) SetPasswordExpiredAt(v string)`

SetPasswordExpiredAt sets PasswordExpiredAt field to given value.

### HasPasswordExpiredAt

`func (o *AccountDescribeUserResponse) HasPasswordExpiredAt() bool`

HasPasswordExpiredAt returns a boolean if a field has been set.

### GetPlatform

`func (o *AccountDescribeUserResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AccountDescribeUserResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AccountDescribeUserResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AccountDescribeUserResponse) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSupervisor

`func (o *AccountDescribeUserResponse) GetSupervisor() bool`

GetSupervisor returns the Supervisor field if non-nil, zero value otherwise.

### GetSupervisorOk

`func (o *AccountDescribeUserResponse) GetSupervisorOk() (*bool, bool)`

GetSupervisorOk returns a tuple with the Supervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisor

`func (o *AccountDescribeUserResponse) SetSupervisor(v bool)`

SetSupervisor sets Supervisor field to given value.

### HasSupervisor

`func (o *AccountDescribeUserResponse) HasSupervisor() bool`

HasSupervisor returns a boolean if a field has been set.

### GetSurname

`func (o *AccountDescribeUserResponse) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *AccountDescribeUserResponse) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *AccountDescribeUserResponse) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *AccountDescribeUserResponse) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetTenant

`func (o *AccountDescribeUserResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AccountDescribeUserResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AccountDescribeUserResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AccountDescribeUserResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUsername

`func (o *AccountDescribeUserResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AccountDescribeUserResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AccountDescribeUserResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AccountDescribeUserResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


