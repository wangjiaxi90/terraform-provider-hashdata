# AccountCreateClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenValiditySeconds** | Pointer to **int32** |  | [optional] 
**AuthorizedGrantTypes** | Pointer to **[]string** |  | [optional] 
**AutoApproveScopes** | Pointer to **[]string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**RefreshTokenValiditySeconds** | Pointer to **int32** |  | [optional] 
**RegisteredRedirectUris** | Pointer to **[]string** |  | [optional] 
**ResourceIds** | Pointer to **[]string** |  | [optional] 
**Scope** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAccountCreateClientRequest

`func NewAccountCreateClientRequest() *AccountCreateClientRequest`

NewAccountCreateClientRequest instantiates a new AccountCreateClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreateClientRequestWithDefaults

`func NewAccountCreateClientRequestWithDefaults() *AccountCreateClientRequest`

NewAccountCreateClientRequestWithDefaults instantiates a new AccountCreateClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenValiditySeconds

`func (o *AccountCreateClientRequest) GetAccessTokenValiditySeconds() int32`

GetAccessTokenValiditySeconds returns the AccessTokenValiditySeconds field if non-nil, zero value otherwise.

### GetAccessTokenValiditySecondsOk

`func (o *AccountCreateClientRequest) GetAccessTokenValiditySecondsOk() (*int32, bool)`

GetAccessTokenValiditySecondsOk returns a tuple with the AccessTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValiditySeconds

`func (o *AccountCreateClientRequest) SetAccessTokenValiditySeconds(v int32)`

SetAccessTokenValiditySeconds sets AccessTokenValiditySeconds field to given value.

### HasAccessTokenValiditySeconds

`func (o *AccountCreateClientRequest) HasAccessTokenValiditySeconds() bool`

HasAccessTokenValiditySeconds returns a boolean if a field has been set.

### GetAuthorizedGrantTypes

`func (o *AccountCreateClientRequest) GetAuthorizedGrantTypes() []string`

GetAuthorizedGrantTypes returns the AuthorizedGrantTypes field if non-nil, zero value otherwise.

### GetAuthorizedGrantTypesOk

`func (o *AccountCreateClientRequest) GetAuthorizedGrantTypesOk() (*[]string, bool)`

GetAuthorizedGrantTypesOk returns a tuple with the AuthorizedGrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedGrantTypes

`func (o *AccountCreateClientRequest) SetAuthorizedGrantTypes(v []string)`

SetAuthorizedGrantTypes sets AuthorizedGrantTypes field to given value.

### HasAuthorizedGrantTypes

`func (o *AccountCreateClientRequest) HasAuthorizedGrantTypes() bool`

HasAuthorizedGrantTypes returns a boolean if a field has been set.

### GetAutoApproveScopes

`func (o *AccountCreateClientRequest) GetAutoApproveScopes() []string`

GetAutoApproveScopes returns the AutoApproveScopes field if non-nil, zero value otherwise.

### GetAutoApproveScopesOk

`func (o *AccountCreateClientRequest) GetAutoApproveScopesOk() (*[]string, bool)`

GetAutoApproveScopesOk returns a tuple with the AutoApproveScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveScopes

`func (o *AccountCreateClientRequest) SetAutoApproveScopes(v []string)`

SetAutoApproveScopes sets AutoApproveScopes field to given value.

### HasAutoApproveScopes

`func (o *AccountCreateClientRequest) HasAutoApproveScopes() bool`

HasAutoApproveScopes returns a boolean if a field has been set.

### GetClientSecret

`func (o *AccountCreateClientRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AccountCreateClientRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AccountCreateClientRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AccountCreateClientRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetRefreshTokenValiditySeconds

`func (o *AccountCreateClientRequest) GetRefreshTokenValiditySeconds() int32`

GetRefreshTokenValiditySeconds returns the RefreshTokenValiditySeconds field if non-nil, zero value otherwise.

### GetRefreshTokenValiditySecondsOk

`func (o *AccountCreateClientRequest) GetRefreshTokenValiditySecondsOk() (*int32, bool)`

GetRefreshTokenValiditySecondsOk returns a tuple with the RefreshTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenValiditySeconds

`func (o *AccountCreateClientRequest) SetRefreshTokenValiditySeconds(v int32)`

SetRefreshTokenValiditySeconds sets RefreshTokenValiditySeconds field to given value.

### HasRefreshTokenValiditySeconds

`func (o *AccountCreateClientRequest) HasRefreshTokenValiditySeconds() bool`

HasRefreshTokenValiditySeconds returns a boolean if a field has been set.

### GetRegisteredRedirectUris

`func (o *AccountCreateClientRequest) GetRegisteredRedirectUris() []string`

GetRegisteredRedirectUris returns the RegisteredRedirectUris field if non-nil, zero value otherwise.

### GetRegisteredRedirectUrisOk

`func (o *AccountCreateClientRequest) GetRegisteredRedirectUrisOk() (*[]string, bool)`

GetRegisteredRedirectUrisOk returns a tuple with the RegisteredRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredRedirectUris

`func (o *AccountCreateClientRequest) SetRegisteredRedirectUris(v []string)`

SetRegisteredRedirectUris sets RegisteredRedirectUris field to given value.

### HasRegisteredRedirectUris

`func (o *AccountCreateClientRequest) HasRegisteredRedirectUris() bool`

HasRegisteredRedirectUris returns a boolean if a field has been set.

### GetResourceIds

`func (o *AccountCreateClientRequest) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *AccountCreateClientRequest) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *AccountCreateClientRequest) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.

### HasResourceIds

`func (o *AccountCreateClientRequest) HasResourceIds() bool`

HasResourceIds returns a boolean if a field has been set.

### GetScope

`func (o *AccountCreateClientRequest) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AccountCreateClientRequest) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AccountCreateClientRequest) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AccountCreateClientRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


