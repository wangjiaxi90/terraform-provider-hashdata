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

// CoreCreateElasticNicRequest struct for CoreCreateElasticNicRequest
type CoreCreateElasticNicRequest struct {
	NicId *[]string `json:"nic_id,omitempty"`
}

// NewCoreCreateElasticNicRequest instantiates a new CoreCreateElasticNicRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCreateElasticNicRequest() *CoreCreateElasticNicRequest {
	this := CoreCreateElasticNicRequest{}
	return &this
}

// NewCoreCreateElasticNicRequestWithDefaults instantiates a new CoreCreateElasticNicRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCreateElasticNicRequestWithDefaults() *CoreCreateElasticNicRequest {
	this := CoreCreateElasticNicRequest{}
	return &this
}

// GetNicId returns the NicId field value if set, zero value otherwise.
func (o *CoreCreateElasticNicRequest) GetNicId() []string {
	if o == nil || o.NicId == nil {
		var ret []string
		return ret
	}
	return *o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateElasticNicRequest) GetNicIdOk() (*[]string, bool) {
	if o == nil || o.NicId == nil {
		return nil, false
	}
	return o.NicId, true
}

// HasNicId returns a boolean if a field has been set.
func (o *CoreCreateElasticNicRequest) HasNicId() bool {
	if o != nil && o.NicId != nil {
		return true
	}

	return false
}

// SetNicId gets a reference to the given []string and assigns it to the NicId field.
func (o *CoreCreateElasticNicRequest) SetNicId(v []string) {
	o.NicId = &v
}

func (o CoreCreateElasticNicRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NicId != nil {
		toSerialize["nic_id"] = o.NicId
	}
	return json.Marshal(toSerialize)
}

type NullableCoreCreateElasticNicRequest struct {
	value *CoreCreateElasticNicRequest
	isSet bool
}

func (v NullableCoreCreateElasticNicRequest) Get() *CoreCreateElasticNicRequest {
	return v.value
}

func (v *NullableCoreCreateElasticNicRequest) Set(val *CoreCreateElasticNicRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCreateElasticNicRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCreateElasticNicRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCreateElasticNicRequest(val *CoreCreateElasticNicRequest) *NullableCoreCreateElasticNicRequest {
	return &NullableCoreCreateElasticNicRequest{value: val, isSet: true}
}

func (v NullableCoreCreateElasticNicRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCreateElasticNicRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

