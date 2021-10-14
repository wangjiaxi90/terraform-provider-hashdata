# CloudmgrcommonIaasResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Az** | Pointer to **string** |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**ClusterType** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**MachineTypeCode** | Pointer to **string** |  | [optional] 
**MachineTypeName** | Pointer to **string** |  | [optional] 
**ProductTypeCode** | Pointer to **string** |  | [optional] 
**ProjectList** | Pointer to [**[]CommonProjectInfo**](CommonProjectInfo.md) |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
**RegionName** | Pointer to **string** |  | [optional] 
**ResInfos** | Pointer to [**[]CommonResourceInfo**](CommonResourceInfo.md) |  | [optional] 
**SubAccountId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudmgrcommonIaasResource

`func NewCloudmgrcommonIaasResource() *CloudmgrcommonIaasResource`

NewCloudmgrcommonIaasResource instantiates a new CloudmgrcommonIaasResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudmgrcommonIaasResourceWithDefaults

`func NewCloudmgrcommonIaasResourceWithDefaults() *CloudmgrcommonIaasResource`

NewCloudmgrcommonIaasResourceWithDefaults instantiates a new CloudmgrcommonIaasResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAz

`func (o *CloudmgrcommonIaasResource) GetAz() string`

GetAz returns the Az field if non-nil, zero value otherwise.

### GetAzOk

`func (o *CloudmgrcommonIaasResource) GetAzOk() (*string, bool)`

GetAzOk returns a tuple with the Az field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAz

`func (o *CloudmgrcommonIaasResource) SetAz(v string)`

SetAz sets Az field to given value.

### HasAz

`func (o *CloudmgrcommonIaasResource) HasAz() bool`

HasAz returns a boolean if a field has been set.

### GetClusterName

`func (o *CloudmgrcommonIaasResource) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CloudmgrcommonIaasResource) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CloudmgrcommonIaasResource) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *CloudmgrcommonIaasResource) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetClusterType

`func (o *CloudmgrcommonIaasResource) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *CloudmgrcommonIaasResource) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *CloudmgrcommonIaasResource) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *CloudmgrcommonIaasResource) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetIp

`func (o *CloudmgrcommonIaasResource) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CloudmgrcommonIaasResource) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CloudmgrcommonIaasResource) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CloudmgrcommonIaasResource) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMachineTypeCode

`func (o *CloudmgrcommonIaasResource) GetMachineTypeCode() string`

GetMachineTypeCode returns the MachineTypeCode field if non-nil, zero value otherwise.

### GetMachineTypeCodeOk

`func (o *CloudmgrcommonIaasResource) GetMachineTypeCodeOk() (*string, bool)`

GetMachineTypeCodeOk returns a tuple with the MachineTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypeCode

`func (o *CloudmgrcommonIaasResource) SetMachineTypeCode(v string)`

SetMachineTypeCode sets MachineTypeCode field to given value.

### HasMachineTypeCode

`func (o *CloudmgrcommonIaasResource) HasMachineTypeCode() bool`

HasMachineTypeCode returns a boolean if a field has been set.

### GetMachineTypeName

`func (o *CloudmgrcommonIaasResource) GetMachineTypeName() string`

GetMachineTypeName returns the MachineTypeName field if non-nil, zero value otherwise.

### GetMachineTypeNameOk

`func (o *CloudmgrcommonIaasResource) GetMachineTypeNameOk() (*string, bool)`

GetMachineTypeNameOk returns a tuple with the MachineTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypeName

`func (o *CloudmgrcommonIaasResource) SetMachineTypeName(v string)`

SetMachineTypeName sets MachineTypeName field to given value.

### HasMachineTypeName

`func (o *CloudmgrcommonIaasResource) HasMachineTypeName() bool`

HasMachineTypeName returns a boolean if a field has been set.

### GetProductTypeCode

`func (o *CloudmgrcommonIaasResource) GetProductTypeCode() string`

GetProductTypeCode returns the ProductTypeCode field if non-nil, zero value otherwise.

### GetProductTypeCodeOk

`func (o *CloudmgrcommonIaasResource) GetProductTypeCodeOk() (*string, bool)`

GetProductTypeCodeOk returns a tuple with the ProductTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeCode

`func (o *CloudmgrcommonIaasResource) SetProductTypeCode(v string)`

SetProductTypeCode sets ProductTypeCode field to given value.

### HasProductTypeCode

`func (o *CloudmgrcommonIaasResource) HasProductTypeCode() bool`

HasProductTypeCode returns a boolean if a field has been set.

### GetProjectList

`func (o *CloudmgrcommonIaasResource) GetProjectList() []CommonProjectInfo`

GetProjectList returns the ProjectList field if non-nil, zero value otherwise.

### GetProjectListOk

`func (o *CloudmgrcommonIaasResource) GetProjectListOk() (*[]CommonProjectInfo, bool)`

GetProjectListOk returns a tuple with the ProjectList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectList

`func (o *CloudmgrcommonIaasResource) SetProjectList(v []CommonProjectInfo)`

SetProjectList sets ProjectList field to given value.

### HasProjectList

`func (o *CloudmgrcommonIaasResource) HasProjectList() bool`

HasProjectList returns a boolean if a field has been set.

### GetRegionCode

`func (o *CloudmgrcommonIaasResource) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *CloudmgrcommonIaasResource) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *CloudmgrcommonIaasResource) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *CloudmgrcommonIaasResource) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetRegionName

`func (o *CloudmgrcommonIaasResource) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *CloudmgrcommonIaasResource) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *CloudmgrcommonIaasResource) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *CloudmgrcommonIaasResource) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetResInfos

`func (o *CloudmgrcommonIaasResource) GetResInfos() []CommonResourceInfo`

GetResInfos returns the ResInfos field if non-nil, zero value otherwise.

### GetResInfosOk

`func (o *CloudmgrcommonIaasResource) GetResInfosOk() (*[]CommonResourceInfo, bool)`

GetResInfosOk returns a tuple with the ResInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResInfos

`func (o *CloudmgrcommonIaasResource) SetResInfos(v []CommonResourceInfo)`

SetResInfos sets ResInfos field to given value.

### HasResInfos

`func (o *CloudmgrcommonIaasResource) HasResInfos() bool`

HasResInfos returns a boolean if a field has been set.

### GetSubAccountId

`func (o *CloudmgrcommonIaasResource) GetSubAccountId() string`

GetSubAccountId returns the SubAccountId field if non-nil, zero value otherwise.

### GetSubAccountIdOk

`func (o *CloudmgrcommonIaasResource) GetSubAccountIdOk() (*string, bool)`

GetSubAccountIdOk returns a tuple with the SubAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountId

`func (o *CloudmgrcommonIaasResource) SetSubAccountId(v string)`

SetSubAccountId sets SubAccountId field to given value.

### HasSubAccountId

`func (o *CloudmgrcommonIaasResource) HasSubAccountId() bool`

HasSubAccountId returns a boolean if a field has been set.

### GetUserId

`func (o *CloudmgrcommonIaasResource) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CloudmgrcommonIaasResource) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CloudmgrcommonIaasResource) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CloudmgrcommonIaasResource) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


