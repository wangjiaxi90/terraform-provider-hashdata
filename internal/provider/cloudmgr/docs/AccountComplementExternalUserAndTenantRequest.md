# AccountComplementExternalUserAndTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenant** | Pointer to [**AccountModifyExternalTenantRequest**](AccountModifyExternalTenantRequest.md) |  | [optional] 
**User** | Pointer to [**AccountModifyExternalUserRequest**](AccountModifyExternalUserRequest.md) |  | [optional] 

## Methods

### NewAccountComplementExternalUserAndTenantRequest

`func NewAccountComplementExternalUserAndTenantRequest() *AccountComplementExternalUserAndTenantRequest`

NewAccountComplementExternalUserAndTenantRequest instantiates a new AccountComplementExternalUserAndTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountComplementExternalUserAndTenantRequestWithDefaults

`func NewAccountComplementExternalUserAndTenantRequestWithDefaults() *AccountComplementExternalUserAndTenantRequest`

NewAccountComplementExternalUserAndTenantRequestWithDefaults instantiates a new AccountComplementExternalUserAndTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenant

`func (o *AccountComplementExternalUserAndTenantRequest) GetTenant() AccountModifyExternalTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AccountComplementExternalUserAndTenantRequest) GetTenantOk() (*AccountModifyExternalTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AccountComplementExternalUserAndTenantRequest) SetTenant(v AccountModifyExternalTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AccountComplementExternalUserAndTenantRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUser

`func (o *AccountComplementExternalUserAndTenantRequest) GetUser() AccountModifyExternalUserRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AccountComplementExternalUserAndTenantRequest) GetUserOk() (*AccountModifyExternalUserRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AccountComplementExternalUserAndTenantRequest) SetUser(v AccountModifyExternalUserRequest)`

SetUser sets User field to given value.

### HasUser

`func (o *AccountComplementExternalUserAndTenantRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


