# CoreListServiceRecoveryDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Jobs** | Pointer to [**[]CommonDescribeJobResponse**](CommonDescribeJobResponse.md) |  | [optional] 
**LastRecoveryPolicy** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Policies** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**RecoverSequence** | Pointer to **string** |  | [optional] 
**RecoveryPolicyName** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreListServiceRecoveryDetailResponse

`func NewCoreListServiceRecoveryDetailResponse() *CoreListServiceRecoveryDetailResponse`

NewCoreListServiceRecoveryDetailResponse instantiates a new CoreListServiceRecoveryDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreListServiceRecoveryDetailResponseWithDefaults

`func NewCoreListServiceRecoveryDetailResponseWithDefaults() *CoreListServiceRecoveryDetailResponse`

NewCoreListServiceRecoveryDetailResponseWithDefaults instantiates a new CoreListServiceRecoveryDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CoreListServiceRecoveryDetailResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CoreListServiceRecoveryDetailResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CoreListServiceRecoveryDetailResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CoreListServiceRecoveryDetailResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CoreListServiceRecoveryDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreListServiceRecoveryDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreListServiceRecoveryDetailResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CoreListServiceRecoveryDetailResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobs

`func (o *CoreListServiceRecoveryDetailResponse) GetJobs() []CommonDescribeJobResponse`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *CoreListServiceRecoveryDetailResponse) GetJobsOk() (*[]CommonDescribeJobResponse, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *CoreListServiceRecoveryDetailResponse) SetJobs(v []CommonDescribeJobResponse)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *CoreListServiceRecoveryDetailResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetLastRecoveryPolicy

`func (o *CoreListServiceRecoveryDetailResponse) GetLastRecoveryPolicy() string`

GetLastRecoveryPolicy returns the LastRecoveryPolicy field if non-nil, zero value otherwise.

### GetLastRecoveryPolicyOk

`func (o *CoreListServiceRecoveryDetailResponse) GetLastRecoveryPolicyOk() (*string, bool)`

GetLastRecoveryPolicyOk returns a tuple with the LastRecoveryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRecoveryPolicy

`func (o *CoreListServiceRecoveryDetailResponse) SetLastRecoveryPolicy(v string)`

SetLastRecoveryPolicy sets LastRecoveryPolicy field to given value.

### HasLastRecoveryPolicy

`func (o *CoreListServiceRecoveryDetailResponse) HasLastRecoveryPolicy() bool`

HasLastRecoveryPolicy returns a boolean if a field has been set.

### GetMessage

`func (o *CoreListServiceRecoveryDetailResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CoreListServiceRecoveryDetailResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CoreListServiceRecoveryDetailResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CoreListServiceRecoveryDetailResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPolicies

`func (o *CoreListServiceRecoveryDetailResponse) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CoreListServiceRecoveryDetailResponse) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CoreListServiceRecoveryDetailResponse) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CoreListServiceRecoveryDetailResponse) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetProperties

`func (o *CoreListServiceRecoveryDetailResponse) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CoreListServiceRecoveryDetailResponse) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CoreListServiceRecoveryDetailResponse) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CoreListServiceRecoveryDetailResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRecoverSequence

`func (o *CoreListServiceRecoveryDetailResponse) GetRecoverSequence() string`

GetRecoverSequence returns the RecoverSequence field if non-nil, zero value otherwise.

### GetRecoverSequenceOk

`func (o *CoreListServiceRecoveryDetailResponse) GetRecoverSequenceOk() (*string, bool)`

GetRecoverSequenceOk returns a tuple with the RecoverSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverSequence

`func (o *CoreListServiceRecoveryDetailResponse) SetRecoverSequence(v string)`

SetRecoverSequence sets RecoverSequence field to given value.

### HasRecoverSequence

`func (o *CoreListServiceRecoveryDetailResponse) HasRecoverSequence() bool`

HasRecoverSequence returns a boolean if a field has been set.

### GetRecoveryPolicyName

`func (o *CoreListServiceRecoveryDetailResponse) GetRecoveryPolicyName() string`

GetRecoveryPolicyName returns the RecoveryPolicyName field if non-nil, zero value otherwise.

### GetRecoveryPolicyNameOk

`func (o *CoreListServiceRecoveryDetailResponse) GetRecoveryPolicyNameOk() (*string, bool)`

GetRecoveryPolicyNameOk returns a tuple with the RecoveryPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPolicyName

`func (o *CoreListServiceRecoveryDetailResponse) SetRecoveryPolicyName(v string)`

SetRecoveryPolicyName sets RecoveryPolicyName field to given value.

### HasRecoveryPolicyName

`func (o *CoreListServiceRecoveryDetailResponse) HasRecoveryPolicyName() bool`

HasRecoveryPolicyName returns a boolean if a field has been set.

### GetService

`func (o *CoreListServiceRecoveryDetailResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CoreListServiceRecoveryDetailResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CoreListServiceRecoveryDetailResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CoreListServiceRecoveryDetailResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *CoreListServiceRecoveryDetailResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CoreListServiceRecoveryDetailResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CoreListServiceRecoveryDetailResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CoreListServiceRecoveryDetailResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CoreListServiceRecoveryDetailResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CoreListServiceRecoveryDetailResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CoreListServiceRecoveryDetailResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CoreListServiceRecoveryDetailResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


