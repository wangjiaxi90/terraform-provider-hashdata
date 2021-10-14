# AccountChangePasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPassword** | Pointer to **string** |  | [optional] 
**OldPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountChangePasswordRequest

`func NewAccountChangePasswordRequest() *AccountChangePasswordRequest`

NewAccountChangePasswordRequest instantiates a new AccountChangePasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountChangePasswordRequestWithDefaults

`func NewAccountChangePasswordRequestWithDefaults() *AccountChangePasswordRequest`

NewAccountChangePasswordRequestWithDefaults instantiates a new AccountChangePasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPassword

`func (o *AccountChangePasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *AccountChangePasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *AccountChangePasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *AccountChangePasswordRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetOldPassword

`func (o *AccountChangePasswordRequest) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *AccountChangePasswordRequest) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *AccountChangePasswordRequest) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *AccountChangePasswordRequest) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


