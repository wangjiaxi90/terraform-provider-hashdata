# CoreCreateRegionConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **string** |  | [optional] 
**AccessKeySecret** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**RecyclingSecurityGroup** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ReverseConnect** | Pointer to **bool** |  | [optional] 
**Tenant** | Pointer to **string** |  | [optional] 
**VendorTenant** | Pointer to **string** |  | [optional] 
**Vpc** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreCreateRegionConfigRequest

`func NewCoreCreateRegionConfigRequest() *CoreCreateRegionConfigRequest`

NewCoreCreateRegionConfigRequest instantiates a new CoreCreateRegionConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCreateRegionConfigRequestWithDefaults

`func NewCoreCreateRegionConfigRequestWithDefaults() *CoreCreateRegionConfigRequest`

NewCoreCreateRegionConfigRequestWithDefaults instantiates a new CoreCreateRegionConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *CoreCreateRegionConfigRequest) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *CoreCreateRegionConfigRequest) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *CoreCreateRegionConfigRequest) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *CoreCreateRegionConfigRequest) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetAccessKeySecret

`func (o *CoreCreateRegionConfigRequest) GetAccessKeySecret() string`

GetAccessKeySecret returns the AccessKeySecret field if non-nil, zero value otherwise.

### GetAccessKeySecretOk

`func (o *CoreCreateRegionConfigRequest) GetAccessKeySecretOk() (*string, bool)`

GetAccessKeySecretOk returns a tuple with the AccessKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeySecret

`func (o *CoreCreateRegionConfigRequest) SetAccessKeySecret(v string)`

SetAccessKeySecret sets AccessKeySecret field to given value.

### HasAccessKeySecret

`func (o *CoreCreateRegionConfigRequest) HasAccessKeySecret() bool`

HasAccessKeySecret returns a boolean if a field has been set.

### GetProject

`func (o *CoreCreateRegionConfigRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CoreCreateRegionConfigRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CoreCreateRegionConfigRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CoreCreateRegionConfigRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRecyclingSecurityGroup

`func (o *CoreCreateRegionConfigRequest) GetRecyclingSecurityGroup() string`

GetRecyclingSecurityGroup returns the RecyclingSecurityGroup field if non-nil, zero value otherwise.

### GetRecyclingSecurityGroupOk

`func (o *CoreCreateRegionConfigRequest) GetRecyclingSecurityGroupOk() (*string, bool)`

GetRecyclingSecurityGroupOk returns a tuple with the RecyclingSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclingSecurityGroup

`func (o *CoreCreateRegionConfigRequest) SetRecyclingSecurityGroup(v string)`

SetRecyclingSecurityGroup sets RecyclingSecurityGroup field to given value.

### HasRecyclingSecurityGroup

`func (o *CoreCreateRegionConfigRequest) HasRecyclingSecurityGroup() bool`

HasRecyclingSecurityGroup returns a boolean if a field has been set.

### GetRegion

`func (o *CoreCreateRegionConfigRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CoreCreateRegionConfigRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CoreCreateRegionConfigRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CoreCreateRegionConfigRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReverseConnect

`func (o *CoreCreateRegionConfigRequest) GetReverseConnect() bool`

GetReverseConnect returns the ReverseConnect field if non-nil, zero value otherwise.

### GetReverseConnectOk

`func (o *CoreCreateRegionConfigRequest) GetReverseConnectOk() (*bool, bool)`

GetReverseConnectOk returns a tuple with the ReverseConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseConnect

`func (o *CoreCreateRegionConfigRequest) SetReverseConnect(v bool)`

SetReverseConnect sets ReverseConnect field to given value.

### HasReverseConnect

`func (o *CoreCreateRegionConfigRequest) HasReverseConnect() bool`

HasReverseConnect returns a boolean if a field has been set.

### GetTenant

`func (o *CoreCreateRegionConfigRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CoreCreateRegionConfigRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CoreCreateRegionConfigRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CoreCreateRegionConfigRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetVendorTenant

`func (o *CoreCreateRegionConfigRequest) GetVendorTenant() string`

GetVendorTenant returns the VendorTenant field if non-nil, zero value otherwise.

### GetVendorTenantOk

`func (o *CoreCreateRegionConfigRequest) GetVendorTenantOk() (*string, bool)`

GetVendorTenantOk returns a tuple with the VendorTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorTenant

`func (o *CoreCreateRegionConfigRequest) SetVendorTenant(v string)`

SetVendorTenant sets VendorTenant field to given value.

### HasVendorTenant

`func (o *CoreCreateRegionConfigRequest) HasVendorTenant() bool`

HasVendorTenant returns a boolean if a field has been set.

### GetVpc

`func (o *CoreCreateRegionConfigRequest) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *CoreCreateRegionConfigRequest) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *CoreCreateRegionConfigRequest) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *CoreCreateRegionConfigRequest) HasVpc() bool`

HasVpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


