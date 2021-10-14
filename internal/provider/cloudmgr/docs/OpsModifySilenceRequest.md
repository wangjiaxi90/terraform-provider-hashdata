# OpsModifySilenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**EndAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Matcher** | Pointer to [**[]OpsMatcher**](OpsMatcher.md) |  | [optional] 
**StartAt** | Pointer to **string** |  | [optional] 

## Methods

### NewOpsModifySilenceRequest

`func NewOpsModifySilenceRequest() *OpsModifySilenceRequest`

NewOpsModifySilenceRequest instantiates a new OpsModifySilenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsModifySilenceRequestWithDefaults

`func NewOpsModifySilenceRequestWithDefaults() *OpsModifySilenceRequest`

NewOpsModifySilenceRequestWithDefaults instantiates a new OpsModifySilenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *OpsModifySilenceRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *OpsModifySilenceRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *OpsModifySilenceRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *OpsModifySilenceRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEndAt

`func (o *OpsModifySilenceRequest) GetEndAt() string`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *OpsModifySilenceRequest) GetEndAtOk() (*string, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *OpsModifySilenceRequest) SetEndAt(v string)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *OpsModifySilenceRequest) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### GetId

`func (o *OpsModifySilenceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpsModifySilenceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpsModifySilenceRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OpsModifySilenceRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMatcher

`func (o *OpsModifySilenceRequest) GetMatcher() []OpsMatcher`

GetMatcher returns the Matcher field if non-nil, zero value otherwise.

### GetMatcherOk

`func (o *OpsModifySilenceRequest) GetMatcherOk() (*[]OpsMatcher, bool)`

GetMatcherOk returns a tuple with the Matcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcher

`func (o *OpsModifySilenceRequest) SetMatcher(v []OpsMatcher)`

SetMatcher sets Matcher field to given value.

### HasMatcher

`func (o *OpsModifySilenceRequest) HasMatcher() bool`

HasMatcher returns a boolean if a field has been set.

### GetStartAt

`func (o *OpsModifySilenceRequest) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *OpsModifySilenceRequest) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *OpsModifySilenceRequest) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *OpsModifySilenceRequest) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


