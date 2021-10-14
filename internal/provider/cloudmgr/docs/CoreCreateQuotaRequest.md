# CoreCreateQuotaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenant** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCoreCreateQuotaRequest

`func NewCoreCreateQuotaRequest() *CoreCreateQuotaRequest`

NewCoreCreateQuotaRequest instantiates a new CoreCreateQuotaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCreateQuotaRequestWithDefaults

`func NewCoreCreateQuotaRequestWithDefaults() *CoreCreateQuotaRequest`

NewCoreCreateQuotaRequestWithDefaults instantiates a new CoreCreateQuotaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenant

`func (o *CoreCreateQuotaRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CoreCreateQuotaRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CoreCreateQuotaRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CoreCreateQuotaRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetZone

`func (o *CoreCreateQuotaRequest) GetZone() []string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CoreCreateQuotaRequest) GetZoneOk() (*[]string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CoreCreateQuotaRequest) SetZone(v []string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CoreCreateQuotaRequest) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


