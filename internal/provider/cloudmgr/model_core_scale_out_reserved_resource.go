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

// CoreScaleOutReservedResource struct for CoreScaleOutReservedResource
type CoreScaleOutReservedResource struct {
	Instances *[]string `json:"instances,omitempty"`
}

// NewCoreScaleOutReservedResource instantiates a new CoreScaleOutReservedResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreScaleOutReservedResource() *CoreScaleOutReservedResource {
	this := CoreScaleOutReservedResource{}
	return &this
}

// NewCoreScaleOutReservedResourceWithDefaults instantiates a new CoreScaleOutReservedResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreScaleOutReservedResourceWithDefaults() *CoreScaleOutReservedResource {
	this := CoreScaleOutReservedResource{}
	return &this
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *CoreScaleOutReservedResource) GetInstances() []string {
	if o == nil || o.Instances == nil {
		var ret []string
		return ret
	}
	return *o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreScaleOutReservedResource) GetInstancesOk() (*[]string, bool) {
	if o == nil || o.Instances == nil {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *CoreScaleOutReservedResource) HasInstances() bool {
	if o != nil && o.Instances != nil {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []string and assigns it to the Instances field.
func (o *CoreScaleOutReservedResource) SetInstances(v []string) {
	o.Instances = &v
}

func (o CoreScaleOutReservedResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}
	return json.Marshal(toSerialize)
}

type NullableCoreScaleOutReservedResource struct {
	value *CoreScaleOutReservedResource
	isSet bool
}

func (v NullableCoreScaleOutReservedResource) Get() *CoreScaleOutReservedResource {
	return v.value
}

func (v *NullableCoreScaleOutReservedResource) Set(val *CoreScaleOutReservedResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreScaleOutReservedResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreScaleOutReservedResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreScaleOutReservedResource(val *CoreScaleOutReservedResource) *NullableCoreScaleOutReservedResource {
	return &NullableCoreScaleOutReservedResource{value: val, isSet: true}
}

func (v NullableCoreScaleOutReservedResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreScaleOutReservedResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


