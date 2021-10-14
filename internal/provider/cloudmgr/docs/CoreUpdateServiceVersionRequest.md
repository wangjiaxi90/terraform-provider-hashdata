# CoreUpdateServiceVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**DowngradeTo** | Pointer to [**CoreUpdateVersionDetailRequest**](CoreUpdateVersionDetailRequest.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**UpgradeFrom** | Pointer to [**CoreUpdateVersionDetailRequest**](CoreUpdateVersionDetailRequest.md) |  | [optional] 

## Methods

### NewCoreUpdateServiceVersionRequest

`func NewCoreUpdateServiceVersionRequest() *CoreUpdateServiceVersionRequest`

NewCoreUpdateServiceVersionRequest instantiates a new CoreUpdateServiceVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreUpdateServiceVersionRequestWithDefaults

`func NewCoreUpdateServiceVersionRequestWithDefaults() *CoreUpdateServiceVersionRequest`

NewCoreUpdateServiceVersionRequestWithDefaults instantiates a new CoreUpdateServiceVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *CoreUpdateServiceVersionRequest) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *CoreUpdateServiceVersionRequest) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *CoreUpdateServiceVersionRequest) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *CoreUpdateServiceVersionRequest) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetDowngradeTo

`func (o *CoreUpdateServiceVersionRequest) GetDowngradeTo() CoreUpdateVersionDetailRequest`

GetDowngradeTo returns the DowngradeTo field if non-nil, zero value otherwise.

### GetDowngradeToOk

`func (o *CoreUpdateServiceVersionRequest) GetDowngradeToOk() (*CoreUpdateVersionDetailRequest, bool)`

GetDowngradeToOk returns a tuple with the DowngradeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowngradeTo

`func (o *CoreUpdateServiceVersionRequest) SetDowngradeTo(v CoreUpdateVersionDetailRequest)`

SetDowngradeTo sets DowngradeTo field to given value.

### HasDowngradeTo

`func (o *CoreUpdateServiceVersionRequest) HasDowngradeTo() bool`

HasDowngradeTo returns a boolean if a field has been set.

### GetEnabled

`func (o *CoreUpdateServiceVersionRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CoreUpdateServiceVersionRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CoreUpdateServiceVersionRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CoreUpdateServiceVersionRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetImage

`func (o *CoreUpdateServiceVersionRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CoreUpdateServiceVersionRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CoreUpdateServiceVersionRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CoreUpdateServiceVersionRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetUpgradeFrom

`func (o *CoreUpdateServiceVersionRequest) GetUpgradeFrom() CoreUpdateVersionDetailRequest`

GetUpgradeFrom returns the UpgradeFrom field if non-nil, zero value otherwise.

### GetUpgradeFromOk

`func (o *CoreUpdateServiceVersionRequest) GetUpgradeFromOk() (*CoreUpdateVersionDetailRequest, bool)`

GetUpgradeFromOk returns a tuple with the UpgradeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeFrom

`func (o *CoreUpdateServiceVersionRequest) SetUpgradeFrom(v CoreUpdateVersionDetailRequest)`

SetUpgradeFrom sets UpgradeFrom field to given value.

### HasUpgradeFrom

`func (o *CoreUpdateServiceVersionRequest) HasUpgradeFrom() bool`

HasUpgradeFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


