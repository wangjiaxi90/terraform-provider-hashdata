# OpsDescribeAlertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]interface{}** |  | [optional] 
**EndsAt** | Pointer to **string** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**GeneratorUrl** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **map[string]interface{}** |  | [optional] 
**Receivers** | Pointer to **[]string** |  | [optional] 
**StartsAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewOpsDescribeAlertResponse

`func NewOpsDescribeAlertResponse() *OpsDescribeAlertResponse`

NewOpsDescribeAlertResponse instantiates a new OpsDescribeAlertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsDescribeAlertResponseWithDefaults

`func NewOpsDescribeAlertResponseWithDefaults() *OpsDescribeAlertResponse`

NewOpsDescribeAlertResponseWithDefaults instantiates a new OpsDescribeAlertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *OpsDescribeAlertResponse) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *OpsDescribeAlertResponse) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *OpsDescribeAlertResponse) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *OpsDescribeAlertResponse) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetEndsAt

`func (o *OpsDescribeAlertResponse) GetEndsAt() string`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *OpsDescribeAlertResponse) GetEndsAtOk() (*string, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *OpsDescribeAlertResponse) SetEndsAt(v string)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *OpsDescribeAlertResponse) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetFingerprint

`func (o *OpsDescribeAlertResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *OpsDescribeAlertResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *OpsDescribeAlertResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *OpsDescribeAlertResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetGeneratorUrl

`func (o *OpsDescribeAlertResponse) GetGeneratorUrl() string`

GetGeneratorUrl returns the GeneratorUrl field if non-nil, zero value otherwise.

### GetGeneratorUrlOk

`func (o *OpsDescribeAlertResponse) GetGeneratorUrlOk() (*string, bool)`

GetGeneratorUrlOk returns a tuple with the GeneratorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorUrl

`func (o *OpsDescribeAlertResponse) SetGeneratorUrl(v string)`

SetGeneratorUrl sets GeneratorUrl field to given value.

### HasGeneratorUrl

`func (o *OpsDescribeAlertResponse) HasGeneratorUrl() bool`

HasGeneratorUrl returns a boolean if a field has been set.

### GetLabels

`func (o *OpsDescribeAlertResponse) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OpsDescribeAlertResponse) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OpsDescribeAlertResponse) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *OpsDescribeAlertResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetReceivers

`func (o *OpsDescribeAlertResponse) GetReceivers() []string`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *OpsDescribeAlertResponse) GetReceiversOk() (*[]string, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *OpsDescribeAlertResponse) SetReceivers(v []string)`

SetReceivers sets Receivers field to given value.

### HasReceivers

`func (o *OpsDescribeAlertResponse) HasReceivers() bool`

HasReceivers returns a boolean if a field has been set.

### GetStartsAt

`func (o *OpsDescribeAlertResponse) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *OpsDescribeAlertResponse) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *OpsDescribeAlertResponse) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *OpsDescribeAlertResponse) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetStatus

`func (o *OpsDescribeAlertResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpsDescribeAlertResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpsDescribeAlertResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpsDescribeAlertResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OpsDescribeAlertResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpsDescribeAlertResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpsDescribeAlertResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpsDescribeAlertResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


