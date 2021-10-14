# AccountAddUserToTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Maintainer** | Pointer to **bool** |  | [optional] 
**Userid** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountAddUserToTeamRequest

`func NewAccountAddUserToTeamRequest() *AccountAddUserToTeamRequest`

NewAccountAddUserToTeamRequest instantiates a new AccountAddUserToTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAddUserToTeamRequestWithDefaults

`func NewAccountAddUserToTeamRequestWithDefaults() *AccountAddUserToTeamRequest`

NewAccountAddUserToTeamRequestWithDefaults instantiates a new AccountAddUserToTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintainer

`func (o *AccountAddUserToTeamRequest) GetMaintainer() bool`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *AccountAddUserToTeamRequest) GetMaintainerOk() (*bool, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *AccountAddUserToTeamRequest) SetMaintainer(v bool)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *AccountAddUserToTeamRequest) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetUserid

`func (o *AccountAddUserToTeamRequest) GetUserid() string`

GetUserid returns the Userid field if non-nil, zero value otherwise.

### GetUseridOk

`func (o *AccountAddUserToTeamRequest) GetUseridOk() (*string, bool)`

GetUseridOk returns a tuple with the Userid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserid

`func (o *AccountAddUserToTeamRequest) SetUserid(v string)`

SetUserid sets Userid field to given value.

### HasUserid

`func (o *AccountAddUserToTeamRequest) HasUserid() bool`

HasUserid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


