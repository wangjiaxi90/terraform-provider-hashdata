# CommonIaasResource

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

### NewCommonIaasResource

`func NewCommonIaasResource() *CommonIaasResource`

NewCommonIaasResource instantiates a new CommonIaasResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonIaasResourceWithDefaults

`func NewCommonIaasResourceWithDefaults() *CommonIaasResource`

NewCommonIaasResourceWithDefaults instantiates a new CommonIaasResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAz

`func (o *CommonIaasResource) GetAz() string`

GetAz returns the Az field if non-nil, zero value otherwise.

### GetAzOk

`func (o *CommonIaasResource) GetAzOk() (*string, bool)`

GetAzOk returns a tuple with the Az field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAz

`func (o *CommonIaasResource) SetAz(v string)`

SetAz sets Az field to given value.

### HasAz

`func (o *CommonIaasResource) HasAz() bool`

HasAz returns a boolean if a field has been set.

### GetClusterName

`func (o *CommonIaasResource) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CommonIaasResource) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CommonIaasResource) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *CommonIaasResource) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetClusterType

`func (o *CommonIaasResource) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *CommonIaasResource) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *CommonIaasResource) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *CommonIaasResource) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetIp

`func (o *CommonIaasResource) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CommonIaasResource) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CommonIaasResource) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CommonIaasResource) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMachineTypeCode

`func (o *CommonIaasResource) GetMachineTypeCode() string`

GetMachineTypeCode returns the MachineTypeCode field if non-nil, zero value otherwise.

### GetMachineTypeCodeOk

`func (o *CommonIaasResource) GetMachineTypeCodeOk() (*string, bool)`

GetMachineTypeCodeOk returns a tuple with the MachineTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypeCode

`func (o *CommonIaasResource) SetMachineTypeCode(v string)`

SetMachineTypeCode sets MachineTypeCode field to given value.

### HasMachineTypeCode

`func (o *CommonIaasResource) HasMachineTypeCode() bool`

HasMachineTypeCode returns a boolean if a field has been set.

### GetMachineTypeName

`func (o *CommonIaasResource) GetMachineTypeName() string`

GetMachineTypeName returns the MachineTypeName field if non-nil, zero value otherwise.

### GetMachineTypeNameOk

`func (o *CommonIaasResource) GetMachineTypeNameOk() (*string, bool)`

GetMachineTypeNameOk returns a tuple with the MachineTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypeName

`func (o *CommonIaasResource) SetMachineTypeName(v string)`

SetMachineTypeName sets MachineTypeName field to given value.

### HasMachineTypeName

`func (o *CommonIaasResource) HasMachineTypeName() bool`

HasMachineTypeName returns a boolean if a field has been set.

### GetProductTypeCode

`func (o *CommonIaasResource) GetProductTypeCode() string`

GetProductTypeCode returns the ProductTypeCode field if non-nil, zero value otherwise.

### GetProductTypeCodeOk

`func (o *CommonIaasResource) GetProductTypeCodeOk() (*string, bool)`

GetProductTypeCodeOk returns a tuple with the ProductTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeCode

`func (o *CommonIaasResource) SetProductTypeCode(v string)`

SetProductTypeCode sets ProductTypeCode field to given value.

### HasProductTypeCode

`func (o *CommonIaasResource) HasProductTypeCode() bool`

HasProductTypeCode returns a boolean if a field has been set.

### GetProjectList

`func (o *CommonIaasResource) GetProjectList() []CommonProjectInfo`

GetProjectList returns the ProjectList field if non-nil, zero value otherwise.

### GetProjectListOk

`func (o *CommonIaasResource) GetProjectListOk() (*[]CommonProjectInfo, bool)`

GetProjectListOk returns a tuple with the ProjectList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectList

`func (o *CommonIaasResource) SetProjectList(v []CommonProjectInfo)`

SetProjectList sets ProjectList field to given value.

### HasProjectList

`func (o *CommonIaasResource) HasProjectList() bool`

HasProjectList returns a boolean if a field has been set.

### GetRegionCode

`func (o *CommonIaasResource) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *CommonIaasResource) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *CommonIaasResource) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *CommonIaasResource) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetRegionName

`func (o *CommonIaasResource) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *CommonIaasResource) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *CommonIaasResource) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *CommonIaasResource) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetResInfos

`func (o *CommonIaasResource) GetResInfos() []CommonResourceInfo`

GetResInfos returns the ResInfos field if non-nil, zero value otherwise.

### GetResInfosOk

`func (o *CommonIaasResource) GetResInfosOk() (*[]CommonResourceInfo, bool)`

GetResInfosOk returns a tuple with the ResInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResInfos

`func (o *CommonIaasResource) SetResInfos(v []CommonResourceInfo)`

SetResInfos sets ResInfos field to given value.

### HasResInfos

`func (o *CommonIaasResource) HasResInfos() bool`

HasResInfos returns a boolean if a field has been set.

### GetSubAccountId

`func (o *CommonIaasResource) GetSubAccountId() string`

GetSubAccountId returns the SubAccountId field if non-nil, zero value otherwise.

### GetSubAccountIdOk

`func (o *CommonIaasResource) GetSubAccountIdOk() (*string, bool)`

GetSubAccountIdOk returns a tuple with the SubAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountId

`func (o *CommonIaasResource) SetSubAccountId(v string)`

SetSubAccountId sets SubAccountId field to given value.

### HasSubAccountId

`func (o *CommonIaasResource) HasSubAccountId() bool`

HasSubAccountId returns a boolean if a field has been set.

### GetUserId

`func (o *CommonIaasResource) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CommonIaasResource) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CommonIaasResource) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CommonIaasResource) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


