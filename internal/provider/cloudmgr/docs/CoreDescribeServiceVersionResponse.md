# CoreDescribeServiceVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DestroyedAt** | Pointer to **string** |  | [optional] 
**DowngradeTo** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpgradeFrom** | Pointer to **[]string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreDescribeServiceVersionResponse

`func NewCoreDescribeServiceVersionResponse() *CoreDescribeServiceVersionResponse`

NewCoreDescribeServiceVersionResponse instantiates a new CoreDescribeServiceVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeServiceVersionResponseWithDefaults

`func NewCoreDescribeServiceVersionResponseWithDefaults() *CoreDescribeServiceVersionResponse`

NewCoreDescribeServiceVersionResponseWithDefaults instantiates a new CoreDescribeServiceVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *CoreDescribeServiceVersionResponse) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *CoreDescribeServiceVersionResponse) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *CoreDescribeServiceVersionResponse) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *CoreDescribeServiceVersionResponse) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CoreDescribeServiceVersionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CoreDescribeServiceVersionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CoreDescribeServiceVersionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CoreDescribeServiceVersionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDestroyedAt

`func (o *CoreDescribeServiceVersionResponse) GetDestroyedAt() string`

GetDestroyedAt returns the DestroyedAt field if non-nil, zero value otherwise.

### GetDestroyedAtOk

`func (o *CoreDescribeServiceVersionResponse) GetDestroyedAtOk() (*string, bool)`

GetDestroyedAtOk returns a tuple with the DestroyedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyedAt

`func (o *CoreDescribeServiceVersionResponse) SetDestroyedAt(v string)`

SetDestroyedAt sets DestroyedAt field to given value.

### HasDestroyedAt

`func (o *CoreDescribeServiceVersionResponse) HasDestroyedAt() bool`

HasDestroyedAt returns a boolean if a field has been set.

### GetDowngradeTo

`func (o *CoreDescribeServiceVersionResponse) GetDowngradeTo() []string`

GetDowngradeTo returns the DowngradeTo field if non-nil, zero value otherwise.

### GetDowngradeToOk

`func (o *CoreDescribeServiceVersionResponse) GetDowngradeToOk() (*[]string, bool)`

GetDowngradeToOk returns a tuple with the DowngradeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowngradeTo

`func (o *CoreDescribeServiceVersionResponse) SetDowngradeTo(v []string)`

SetDowngradeTo sets DowngradeTo field to given value.

### HasDowngradeTo

`func (o *CoreDescribeServiceVersionResponse) HasDowngradeTo() bool`

HasDowngradeTo returns a boolean if a field has been set.

### GetEnabled

`func (o *CoreDescribeServiceVersionResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CoreDescribeServiceVersionResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CoreDescribeServiceVersionResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CoreDescribeServiceVersionResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *CoreDescribeServiceVersionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreDescribeServiceVersionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreDescribeServiceVersionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CoreDescribeServiceVersionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *CoreDescribeServiceVersionResponse) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CoreDescribeServiceVersionResponse) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CoreDescribeServiceVersionResponse) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CoreDescribeServiceVersionResponse) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetRegion

`func (o *CoreDescribeServiceVersionResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CoreDescribeServiceVersionResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CoreDescribeServiceVersionResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CoreDescribeServiceVersionResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetType

`func (o *CoreDescribeServiceVersionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CoreDescribeServiceVersionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CoreDescribeServiceVersionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CoreDescribeServiceVersionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpgradeFrom

`func (o *CoreDescribeServiceVersionResponse) GetUpgradeFrom() []string`

GetUpgradeFrom returns the UpgradeFrom field if non-nil, zero value otherwise.

### GetUpgradeFromOk

`func (o *CoreDescribeServiceVersionResponse) GetUpgradeFromOk() (*[]string, bool)`

GetUpgradeFromOk returns a tuple with the UpgradeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeFrom

`func (o *CoreDescribeServiceVersionResponse) SetUpgradeFrom(v []string)`

SetUpgradeFrom sets UpgradeFrom field to given value.

### HasUpgradeFrom

`func (o *CoreDescribeServiceVersionResponse) HasUpgradeFrom() bool`

HasUpgradeFrom returns a boolean if a field has been set.

### GetVendor

`func (o *CoreDescribeServiceVersionResponse) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CoreDescribeServiceVersionResponse) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CoreDescribeServiceVersionResponse) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CoreDescribeServiceVersionResponse) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *CoreDescribeServiceVersionResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CoreDescribeServiceVersionResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CoreDescribeServiceVersionResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CoreDescribeServiceVersionResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


