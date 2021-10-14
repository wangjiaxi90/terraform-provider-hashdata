# CoreInitializeTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenant** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreInitializeTenantRequest

`func NewCoreInitializeTenantRequest() *CoreInitializeTenantRequest`

NewCoreInitializeTenantRequest instantiates a new CoreInitializeTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreInitializeTenantRequestWithDefaults

`func NewCoreInitializeTenantRequestWithDefaults() *CoreInitializeTenantRequest`

NewCoreInitializeTenantRequestWithDefaults instantiates a new CoreInitializeTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenant

`func (o *CoreInitializeTenantRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CoreInitializeTenantRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CoreInitializeTenantRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CoreInitializeTenantRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetVendor

`func (o *CoreInitializeTenantRequest) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CoreInitializeTenantRequest) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CoreInitializeTenantRequest) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CoreInitializeTenantRequest) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


