# CoreDescribeHdfsConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hdfs** | Pointer to **string** |  | [optional] 
**Hosts** | Pointer to [**[]CoreHost**](CoreHost.md) |  | [optional] 
**KeytabName** | Pointer to **string** |  | [optional] 
**Krb5** | Pointer to [**CoreKrb5**](CoreKrb5.md) |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreDescribeHdfsConfigResponse

`func NewCoreDescribeHdfsConfigResponse() *CoreDescribeHdfsConfigResponse`

NewCoreDescribeHdfsConfigResponse instantiates a new CoreDescribeHdfsConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDescribeHdfsConfigResponseWithDefaults

`func NewCoreDescribeHdfsConfigResponseWithDefaults() *CoreDescribeHdfsConfigResponse`

NewCoreDescribeHdfsConfigResponseWithDefaults instantiates a new CoreDescribeHdfsConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHdfs

`func (o *CoreDescribeHdfsConfigResponse) GetHdfs() string`

GetHdfs returns the Hdfs field if non-nil, zero value otherwise.

### GetHdfsOk

`func (o *CoreDescribeHdfsConfigResponse) GetHdfsOk() (*string, bool)`

GetHdfsOk returns a tuple with the Hdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfs

`func (o *CoreDescribeHdfsConfigResponse) SetHdfs(v string)`

SetHdfs sets Hdfs field to given value.

### HasHdfs

`func (o *CoreDescribeHdfsConfigResponse) HasHdfs() bool`

HasHdfs returns a boolean if a field has been set.

### GetHosts

`func (o *CoreDescribeHdfsConfigResponse) GetHosts() []CoreHost`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *CoreDescribeHdfsConfigResponse) GetHostsOk() (*[]CoreHost, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *CoreDescribeHdfsConfigResponse) SetHosts(v []CoreHost)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *CoreDescribeHdfsConfigResponse) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetKeytabName

`func (o *CoreDescribeHdfsConfigResponse) GetKeytabName() string`

GetKeytabName returns the KeytabName field if non-nil, zero value otherwise.

### GetKeytabNameOk

`func (o *CoreDescribeHdfsConfigResponse) GetKeytabNameOk() (*string, bool)`

GetKeytabNameOk returns a tuple with the KeytabName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytabName

`func (o *CoreDescribeHdfsConfigResponse) SetKeytabName(v string)`

SetKeytabName sets KeytabName field to given value.

### HasKeytabName

`func (o *CoreDescribeHdfsConfigResponse) HasKeytabName() bool`

HasKeytabName returns a boolean if a field has been set.

### GetKrb5

`func (o *CoreDescribeHdfsConfigResponse) GetKrb5() CoreKrb5`

GetKrb5 returns the Krb5 field if non-nil, zero value otherwise.

### GetKrb5Ok

`func (o *CoreDescribeHdfsConfigResponse) GetKrb5Ok() (*CoreKrb5, bool)`

GetKrb5Ok returns a tuple with the Krb5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrb5

`func (o *CoreDescribeHdfsConfigResponse) SetKrb5(v CoreKrb5)`

SetKrb5 sets Krb5 field to given value.

### HasKrb5

`func (o *CoreDescribeHdfsConfigResponse) HasKrb5() bool`

HasKrb5 returns a boolean if a field has been set.

### GetService

`func (o *CoreDescribeHdfsConfigResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CoreDescribeHdfsConfigResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CoreDescribeHdfsConfigResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CoreDescribeHdfsConfigResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *CoreDescribeHdfsConfigResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CoreDescribeHdfsConfigResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CoreDescribeHdfsConfigResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CoreDescribeHdfsConfigResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *CoreDescribeHdfsConfigResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CoreDescribeHdfsConfigResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CoreDescribeHdfsConfigResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CoreDescribeHdfsConfigResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


