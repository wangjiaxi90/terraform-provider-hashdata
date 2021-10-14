# CoreDescribeRecoveryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Jobs** | Pointer to **[]string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Policies** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**RecoveryPolicyName** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreDescribeRecoveryResponse

`func NewCoreDescribeRecoveryResponse() *CoreDescribeRecoveryResponse`

NewCoreDescribeRecoveryResponse instantiates a new CoreDescribeRecoveryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeRecoveryResponseWithDefaults

`func NewCoreDescribeRecoveryResponseWithDefaults() *CoreDescribeRecoveryResponse`

NewCoreDescribeRecoveryResponseWithDefaults instantiates a new CoreDescribeRecoveryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CoreDescribeRecoveryResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CoreDescribeRecoveryResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CoreDescribeRecoveryResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CoreDescribeRecoveryResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CoreDescribeRecoveryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreDescribeRecoveryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreDescribeRecoveryResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CoreDescribeRecoveryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobs

`func (o *CoreDescribeRecoveryResponse) GetJobs() []string`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *CoreDescribeRecoveryResponse) GetJobsOk() (*[]string, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *CoreDescribeRecoveryResponse) SetJobs(v []string)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *CoreDescribeRecoveryResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetMessage

`func (o *CoreDescribeRecoveryResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CoreDescribeRecoveryResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CoreDescribeRecoveryResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CoreDescribeRecoveryResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPolicies

`func (o *CoreDescribeRecoveryResponse) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CoreDescribeRecoveryResponse) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CoreDescribeRecoveryResponse) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CoreDescribeRecoveryResponse) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetProperties

`func (o *CoreDescribeRecoveryResponse) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CoreDescribeRecoveryResponse) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CoreDescribeRecoveryResponse) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CoreDescribeRecoveryResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRecoveryPolicyName

`func (o *CoreDescribeRecoveryResponse) GetRecoveryPolicyName() string`

GetRecoveryPolicyName returns the RecoveryPolicyName field if non-nil, zero value otherwise.

### GetRecoveryPolicyNameOk

`func (o *CoreDescribeRecoveryResponse) GetRecoveryPolicyNameOk() (*string, bool)`

GetRecoveryPolicyNameOk returns a tuple with the RecoveryPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPolicyName

`func (o *CoreDescribeRecoveryResponse) SetRecoveryPolicyName(v string)`

SetRecoveryPolicyName sets RecoveryPolicyName field to given value.

### HasRecoveryPolicyName

`func (o *CoreDescribeRecoveryResponse) HasRecoveryPolicyName() bool`

HasRecoveryPolicyName returns a boolean if a field has been set.

### GetService

`func (o *CoreDescribeRecoveryResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CoreDescribeRecoveryResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CoreDescribeRecoveryResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CoreDescribeRecoveryResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *CoreDescribeRecoveryResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CoreDescribeRecoveryResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CoreDescribeRecoveryResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CoreDescribeRecoveryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CoreDescribeRecoveryResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CoreDescribeRecoveryResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CoreDescribeRecoveryResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CoreDescribeRecoveryResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


