# CoreCreateZoneConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenant** | Pointer to **string** |  | [optional] 
**Vxnet** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreCreateZoneConfigRequest

`func NewCoreCreateZoneConfigRequest() *CoreCreateZoneConfigRequest`

NewCoreCreateZoneConfigRequest instantiates a new CoreCreateZoneConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCreateZoneConfigRequestWithDefaults

`func NewCoreCreateZoneConfigRequestWithDefaults() *CoreCreateZoneConfigRequest`

NewCoreCreateZoneConfigRequestWithDefaults instantiates a new CoreCreateZoneConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenant

`func (o *CoreCreateZoneConfigRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CoreCreateZoneConfigRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CoreCreateZoneConfigRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CoreCreateZoneConfigRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetVxnet

`func (o *CoreCreateZoneConfigRequest) GetVxnet() string`

GetVxnet returns the Vxnet field if non-nil, zero value otherwise.

### GetVxnetOk

`func (o *CoreCreateZoneConfigRequest) GetVxnetOk() (*string, bool)`

GetVxnetOk returns a tuple with the Vxnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxnet

`func (o *CoreCreateZoneConfigRequest) SetVxnet(v string)`

SetVxnet sets Vxnet field to given value.

### HasVxnet

`func (o *CoreCreateZoneConfigRequest) HasVxnet() bool`

HasVxnet returns a boolean if a field has been set.

### GetZone

`func (o *CoreCreateZoneConfigRequest) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CoreCreateZoneConfigRequest) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CoreCreateZoneConfigRequest) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CoreCreateZoneConfigRequest) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


