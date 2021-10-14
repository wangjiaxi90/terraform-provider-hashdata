# CoreUpdateQuotaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quota** | Pointer to **int32** |  | [optional] 
**Tenant** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreUpdateQuotaRequest

`func NewCoreUpdateQuotaRequest() *CoreUpdateQuotaRequest`

NewCoreUpdateQuotaRequest instantiates a new CoreUpdateQuotaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreUpdateQuotaRequestWithDefaults

`func NewCoreUpdateQuotaRequestWithDefaults() *CoreUpdateQuotaRequest`

NewCoreUpdateQuotaRequestWithDefaults instantiates a new CoreUpdateQuotaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuota

`func (o *CoreUpdateQuotaRequest) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *CoreUpdateQuotaRequest) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *CoreUpdateQuotaRequest) SetQuota(v int32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *CoreUpdateQuotaRequest) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetTenant

`func (o *CoreUpdateQuotaRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CoreUpdateQuotaRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CoreUpdateQuotaRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CoreUpdateQuotaRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetType

`func (o *CoreUpdateQuotaRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CoreUpdateQuotaRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CoreUpdateQuotaRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CoreUpdateQuotaRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetZone

`func (o *CoreUpdateQuotaRequest) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CoreUpdateQuotaRequest) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CoreUpdateQuotaRequest) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CoreUpdateQuotaRequest) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


