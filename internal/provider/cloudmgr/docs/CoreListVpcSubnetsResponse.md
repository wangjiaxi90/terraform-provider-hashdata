# CoreListVpcSubnetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | Pointer to [**[]CoreDescribeSubnetResponse**](CoreDescribeSubnetResponse.md) |  | [optional] 

## Methods

### NewCoreListVpcSubnetsResponse

`func NewCoreListVpcSubnetsResponse() *CoreListVpcSubnetsResponse`

NewCoreListVpcSubnetsResponse instantiates a new CoreListVpcSubnetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreListVpcSubnetsResponseWithDefaults

`func NewCoreListVpcSubnetsResponseWithDefaults() *CoreListVpcSubnetsResponse`

NewCoreListVpcSubnetsResponseWithDefaults instantiates a new CoreListVpcSubnetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *CoreListVpcSubnetsResponse) GetSubnet() []CoreDescribeSubnetResponse`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CoreListVpcSubnetsResponse) GetSubnetOk() (*[]CoreDescribeSubnetResponse, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CoreListVpcSubnetsResponse) SetSubnet(v []CoreDescribeSubnetResponse)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *CoreListVpcSubnetsResponse) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


