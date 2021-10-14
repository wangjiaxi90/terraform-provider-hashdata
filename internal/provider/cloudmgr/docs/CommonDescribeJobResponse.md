# CommonDescribeJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **int32** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Extras** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**RecoveryId** | Pointer to **string** |  | [optional] 
**ResourceIds** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewCommonDescribeJobResponse

`func NewCommonDescribeJobResponse() *CommonDescribeJobResponse`

NewCommonDescribeJobResponse instantiates a new CommonDescribeJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonDescribeJobResponseWithDefaults

`func NewCommonDescribeJobResponseWithDefaults() *CommonDescribeJobResponse`

NewCommonDescribeJobResponseWithDefaults instantiates a new CommonDescribeJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CommonDescribeJobResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CommonDescribeJobResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CommonDescribeJobResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CommonDescribeJobResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommonDescribeJobResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommonDescribeJobResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommonDescribeJobResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommonDescribeJobResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetErrorCode

`func (o *CommonDescribeJobResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CommonDescribeJobResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CommonDescribeJobResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CommonDescribeJobResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CommonDescribeJobResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CommonDescribeJobResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CommonDescribeJobResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CommonDescribeJobResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetExtras

`func (o *CommonDescribeJobResponse) GetExtras() string`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *CommonDescribeJobResponse) GetExtrasOk() (*string, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *CommonDescribeJobResponse) SetExtras(v string)`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *CommonDescribeJobResponse) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetId

`func (o *CommonDescribeJobResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonDescribeJobResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonDescribeJobResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommonDescribeJobResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecoveryId

`func (o *CommonDescribeJobResponse) GetRecoveryId() string`

GetRecoveryId returns the RecoveryId field if non-nil, zero value otherwise.

### GetRecoveryIdOk

`func (o *CommonDescribeJobResponse) GetRecoveryIdOk() (*string, bool)`

GetRecoveryIdOk returns a tuple with the RecoveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryId

`func (o *CommonDescribeJobResponse) SetRecoveryId(v string)`

SetRecoveryId sets RecoveryId field to given value.

### HasRecoveryId

`func (o *CommonDescribeJobResponse) HasRecoveryId() bool`

HasRecoveryId returns a boolean if a field has been set.

### GetResourceIds

`func (o *CommonDescribeJobResponse) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *CommonDescribeJobResponse) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *CommonDescribeJobResponse) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.

### HasResourceIds

`func (o *CommonDescribeJobResponse) HasResourceIds() bool`

HasResourceIds returns a boolean if a field has been set.

### GetStatus

`func (o *CommonDescribeJobResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonDescribeJobResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonDescribeJobResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommonDescribeJobResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommonDescribeJobResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommonDescribeJobResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommonDescribeJobResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommonDescribeJobResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *CommonDescribeJobResponse) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CommonDescribeJobResponse) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CommonDescribeJobResponse) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CommonDescribeJobResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


