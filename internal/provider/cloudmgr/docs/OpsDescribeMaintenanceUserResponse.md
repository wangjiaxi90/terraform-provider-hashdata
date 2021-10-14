# OpsDescribeMaintenanceUserResponse

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
**GivenName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**PasswordExpiredAt** | Pointer to **string** |  | [optional] 
**Supervisor** | Pointer to **bool** |  | [optional] 
**Surname** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewOpsDescribeMaintenanceUserResponse

`func NewOpsDescribeMaintenanceUserResponse() *OpsDescribeMaintenanceUserResponse`

NewOpsDescribeMaintenanceUserResponse instantiates a new OpsDescribeMaintenanceUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsDescribeMaintenanceUserResponseWithDefaults

`func NewOpsDescribeMaintenanceUserResponseWithDefaults() *OpsDescribeMaintenanceUserResponse`

NewOpsDescribeMaintenanceUserResponseWithDefaults instantiates a new OpsDescribeMaintenanceUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *OpsDescribeMaintenanceUserResponse) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *OpsDescribeMaintenanceUserResponse) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *OpsDescribeMaintenanceUserResponse) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *OpsDescribeMaintenanceUserResponse) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OpsDescribeMaintenanceUserResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpsDescribeMaintenanceUserResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpsDescribeMaintenanceUserResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpsDescribeMaintenanceUserResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDestroyedAt

`func (o *OpsDescribeMaintenanceUserResponse) GetDestroyedAt() string`

GetDestroyedAt returns the DestroyedAt field if non-nil, zero value otherwise.

### GetDestroyedAtOk

`func (o *OpsDescribeMaintenanceUserResponse) GetDestroyedAtOk() (*string, bool)`

GetDestroyedAtOk returns a tuple with the DestroyedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyedAt

`func (o *OpsDescribeMaintenanceUserResponse) SetDestroyedAt(v string)`

SetDestroyedAt sets DestroyedAt field to given value.

### HasDestroyedAt

`func (o *OpsDescribeMaintenanceUserResponse) HasDestroyedAt() bool`

HasDestroyedAt returns a boolean if a field has been set.

### GetEmail

`func (o *OpsDescribeMaintenanceUserResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OpsDescribeMaintenanceUserResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OpsDescribeMaintenanceUserResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OpsDescribeMaintenanceUserResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *OpsDescribeMaintenanceUserResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OpsDescribeMaintenanceUserResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OpsDescribeMaintenanceUserResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OpsDescribeMaintenanceUserResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpiredAt

`func (o *OpsDescribeMaintenanceUserResponse) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *OpsDescribeMaintenanceUserResponse) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *OpsDescribeMaintenanceUserResponse) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *OpsDescribeMaintenanceUserResponse) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetExternal

`func (o *OpsDescribeMaintenanceUserResponse) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *OpsDescribeMaintenanceUserResponse) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *OpsDescribeMaintenanceUserResponse) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *OpsDescribeMaintenanceUserResponse) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetGivenName

`func (o *OpsDescribeMaintenanceUserResponse) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *OpsDescribeMaintenanceUserResponse) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *OpsDescribeMaintenanceUserResponse) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *OpsDescribeMaintenanceUserResponse) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetId

`func (o *OpsDescribeMaintenanceUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpsDescribeMaintenanceUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpsDescribeMaintenanceUserResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OpsDescribeMaintenanceUserResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocked

`func (o *OpsDescribeMaintenanceUserResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *OpsDescribeMaintenanceUserResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *OpsDescribeMaintenanceUserResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *OpsDescribeMaintenanceUserResponse) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPasswordExpiredAt

`func (o *OpsDescribeMaintenanceUserResponse) GetPasswordExpiredAt() string`

GetPasswordExpiredAt returns the PasswordExpiredAt field if non-nil, zero value otherwise.

### GetPasswordExpiredAtOk

`func (o *OpsDescribeMaintenanceUserResponse) GetPasswordExpiredAtOk() (*string, bool)`

GetPasswordExpiredAtOk returns a tuple with the PasswordExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiredAt

`func (o *OpsDescribeMaintenanceUserResponse) SetPasswordExpiredAt(v string)`

SetPasswordExpiredAt sets PasswordExpiredAt field to given value.

### HasPasswordExpiredAt

`func (o *OpsDescribeMaintenanceUserResponse) HasPasswordExpiredAt() bool`

HasPasswordExpiredAt returns a boolean if a field has been set.

### GetSupervisor

`func (o *OpsDescribeMaintenanceUserResponse) GetSupervisor() bool`

GetSupervisor returns the Supervisor field if non-nil, zero value otherwise.

### GetSupervisorOk

`func (o *OpsDescribeMaintenanceUserResponse) GetSupervisorOk() (*bool, bool)`

GetSupervisorOk returns a tuple with the Supervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisor

`func (o *OpsDescribeMaintenanceUserResponse) SetSupervisor(v bool)`

SetSupervisor sets Supervisor field to given value.

### HasSupervisor

`func (o *OpsDescribeMaintenanceUserResponse) HasSupervisor() bool`

HasSupervisor returns a boolean if a field has been set.

### GetSurname

`func (o *OpsDescribeMaintenanceUserResponse) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *OpsDescribeMaintenanceUserResponse) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *OpsDescribeMaintenanceUserResponse) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *OpsDescribeMaintenanceUserResponse) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetUsername

`func (o *OpsDescribeMaintenanceUserResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OpsDescribeMaintenanceUserResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OpsDescribeMaintenanceUserResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OpsDescribeMaintenanceUserResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


