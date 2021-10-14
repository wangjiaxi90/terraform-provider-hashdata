# AccountListTenantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]AccountDescribeTenantResponse**](AccountDescribeTenantResponse.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewAccountListTenantResponse

`func NewAccountListTenantResponse() *AccountListTenantResponse`

NewAccountListTenantResponse instantiates a new AccountListTenantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountListTenantResponseWithDefaults

`func NewAccountListTenantResponseWithDefaults() *AccountListTenantResponse`

NewAccountListTenantResponseWithDefaults instantiates a new AccountListTenantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *AccountListTenantResponse) GetContent() []AccountDescribeTenantResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AccountListTenantResponse) GetContentOk() (*[]AccountDescribeTenantResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AccountListTenantResponse) SetContent(v []AccountDescribeTenantResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *AccountListTenantResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCount

`func (o *AccountListTenantResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AccountListTenantResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AccountListTenantResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AccountListTenantResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *AccountListTenantResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccountListTenantResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccountListTenantResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AccountListTenantResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


