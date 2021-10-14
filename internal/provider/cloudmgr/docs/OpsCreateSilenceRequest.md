# OpsCreateSilenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**EndAt** | Pointer to **string** |  | [optional] 
**Matcher** | Pointer to [**[]OpsMatcher**](OpsMatcher.md) |  | [optional] 
**StartAt** | Pointer to **string** |  | [optional] 

## Methods

### NewOpsCreateSilenceRequest

`func NewOpsCreateSilenceRequest() *OpsCreateSilenceRequest`

NewOpsCreateSilenceRequest instantiates a new OpsCreateSilenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsCreateSilenceRequestWithDefaults

`func NewOpsCreateSilenceRequestWithDefaults() *OpsCreateSilenceRequest`

NewOpsCreateSilenceRequestWithDefaults instantiates a new OpsCreateSilenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *OpsCreateSilenceRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *OpsCreateSilenceRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *OpsCreateSilenceRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *OpsCreateSilenceRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEndAt

`func (o *OpsCreateSilenceRequest) GetEndAt() string`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *OpsCreateSilenceRequest) GetEndAtOk() (*string, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *OpsCreateSilenceRequest) SetEndAt(v string)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *OpsCreateSilenceRequest) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### GetMatcher

`func (o *OpsCreateSilenceRequest) GetMatcher() []OpsMatcher`

GetMatcher returns the Matcher field if non-nil, zero value otherwise.

### GetMatcherOk

`func (o *OpsCreateSilenceRequest) GetMatcherOk() (*[]OpsMatcher, bool)`

GetMatcherOk returns a tuple with the Matcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcher

`func (o *OpsCreateSilenceRequest) SetMatcher(v []OpsMatcher)`

SetMatcher sets Matcher field to given value.

### HasMatcher

`func (o *OpsCreateSilenceRequest) HasMatcher() bool`

HasMatcher returns a boolean if a field has been set.

### GetStartAt

`func (o *OpsCreateSilenceRequest) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *OpsCreateSilenceRequest) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *OpsCreateSilenceRequest) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *OpsCreateSilenceRequest) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


