/*
 * Cloud Manager API
 *
 * Cloud Manager Restful API Documentation.
 *
 * API version: v2.0
 * Contact: wang@hashdata.cn
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudmgr

import (
	"encoding/json"
)

// CoreListVpcSubnetsResponse struct for CoreListVpcSubnetsResponse
type CoreListVpcSubnetsResponse struct {
	Subnet *[]CoreDescribeSubnetResponse `json:"subnet,omitempty"`
}

// NewCoreListVpcSubnetsResponse instantiates a new CoreListVpcSubnetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreListVpcSubnetsResponse() *CoreListVpcSubnetsResponse {
	this := CoreListVpcSubnetsResponse{}
	return &this
}

// NewCoreListVpcSubnetsResponseWithDefaults instantiates a new CoreListVpcSubnetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreListVpcSubnetsResponseWithDefaults() *CoreListVpcSubnetsResponse {
	this := CoreListVpcSubnetsResponse{}
	return &this
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *CoreListVpcSubnetsResponse) GetSubnet() []CoreDescribeSubnetResponse {
	if o == nil || o.Subnet == nil {
		var ret []CoreDescribeSubnetResponse
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListVpcSubnetsResponse) GetSubnetOk() (*[]CoreDescribeSubnetResponse, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *CoreListVpcSubnetsResponse) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given []CoreDescribeSubnetResponse and assigns it to the Subnet field.
func (o *CoreListVpcSubnetsResponse) SetSubnet(v []CoreDescribeSubnetResponse) {
	o.Subnet = &v
}

func (o CoreListVpcSubnetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	return json.Marshal(toSerialize)
}

type NullableCoreListVpcSubnetsResponse struct {
	value *CoreListVpcSubnetsResponse
	isSet bool
}

func (v NullableCoreListVpcSubnetsResponse) Get() *CoreListVpcSubnetsResponse {
	return v.value
}

func (v *NullableCoreListVpcSubnetsResponse) Set(val *CoreListVpcSubnetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreListVpcSubnetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreListVpcSubnetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreListVpcSubnetsResponse(val *CoreListVpcSubnetsResponse) *NullableCoreListVpcSubnetsResponse {
	return &NullableCoreListVpcSubnetsResponse{value: val, isSet: true}
}

func (v NullableCoreListVpcSubnetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreListVpcSubnetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


