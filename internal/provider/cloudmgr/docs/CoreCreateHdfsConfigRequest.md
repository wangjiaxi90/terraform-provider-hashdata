# CoreCreateHdfsConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hdfs** | Pointer to **string** |  | [optional] 
**Hosts** | Pointer to [**[]CoreHost**](CoreHost.md) |  | [optional] 
**KeytabBase64Content** | Pointer to **string** |  | [optional] 
**Krb5** | Pointer to [**CoreKrb5**](CoreKrb5.md) |  | [optional] 
**ReferenceTimestamp** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreCreateHdfsConfigRequest

`func NewCoreCreateHdfsConfigRequest() *CoreCreateHdfsConfigRequest`

NewCoreCreateHdfsConfigRequest instantiates a new CoreCreateHdfsConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCreateHdfsConfigRequestWithDefaults

`func NewCoreCreateHdfsConfigRequestWithDefaults() *CoreCreateHdfsConfigRequest`

NewCoreCreateHdfsConfigRequestWithDefaults instantiates a new CoreCreateHdfsConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHdfs

`func (o *CoreCreateHdfsConfigRequest) GetHdfs() string`

GetHdfs returns the Hdfs field if non-nil, zero value otherwise.

### GetHdfsOk

`func (o *CoreCreateHdfsConfigRequest) GetHdfsOk() (*string, bool)`

GetHdfsOk returns a tuple with the Hdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfs

`func (o *CoreCreateHdfsConfigRequest) SetHdfs(v string)`

SetHdfs sets Hdfs field to given value.

### HasHdfs

`func (o *CoreCreateHdfsConfigRequest) HasHdfs() bool`

HasHdfs returns a boolean if a field has been set.

### GetHosts

`func (o *CoreCreateHdfsConfigRequest) GetHosts() []CoreHost`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *CoreCreateHdfsConfigRequest) GetHostsOk() (*[]CoreHost, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *CoreCreateHdfsConfigRequest) SetHosts(v []CoreHost)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *CoreCreateHdfsConfigRequest) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetKeytabBase64Content

`func (o *CoreCreateHdfsConfigRequest) GetKeytabBase64Content() string`

GetKeytabBase64Content returns the KeytabBase64Content field if non-nil, zero value otherwise.

### GetKeytabBase64ContentOk

`func (o *CoreCreateHdfsConfigRequest) GetKeytabBase64ContentOk() (*string, bool)`

GetKeytabBase64ContentOk returns a tuple with the KeytabBase64Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytabBase64Content

`func (o *CoreCreateHdfsConfigRequest) SetKeytabBase64Content(v string)`

SetKeytabBase64Content sets KeytabBase64Content field to given value.

### HasKeytabBase64Content

`func (o *CoreCreateHdfsConfigRequest) HasKeytabBase64Content() bool`

HasKeytabBase64Content returns a boolean if a field has been set.

### GetKrb5

`func (o *CoreCreateHdfsConfigRequest) GetKrb5() CoreKrb5`

GetKrb5 returns the Krb5 field if non-nil, zero value otherwise.

### GetKrb5Ok

`func (o *CoreCreateHdfsConfigRequest) GetKrb5Ok() (*CoreKrb5, bool)`

GetKrb5Ok returns a tuple with the Krb5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrb5

`func (o *CoreCreateHdfsConfigRequest) SetKrb5(v CoreKrb5)`

SetKrb5 sets Krb5 field to given value.

### HasKrb5

`func (o *CoreCreateHdfsConfigRequest) HasKrb5() bool`

HasKrb5 returns a boolean if a field has been set.

### GetReferenceTimestamp

`func (o *CoreCreateHdfsConfigRequest) GetReferenceTimestamp() string`

GetReferenceTimestamp returns the ReferenceTimestamp field if non-nil, zero value otherwise.

### GetReferenceTimestampOk

`func (o *CoreCreateHdfsConfigRequest) GetReferenceTimestampOk() (*string, bool)`

GetReferenceTimestampOk returns a tuple with the ReferenceTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTimestamp

`func (o *CoreCreateHdfsConfigRequest) SetReferenceTimestamp(v string)`

SetReferenceTimestamp sets ReferenceTimestamp field to given value.

### HasReferenceTimestamp

`func (o *CoreCreateHdfsConfigRequest) HasReferenceTimestamp() bool`

HasReferenceTimestamp returns a boolean if a field has been set.

### GetStatus

`func (o *CoreCreateHdfsConfigRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CoreCreateHdfsConfigRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CoreCreateHdfsConfigRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CoreCreateHdfsConfigRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


