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

// CoreScaleOutServiceComponentRequest struct for CoreScaleOutServiceComponentRequest
type CoreScaleOutServiceComponentRequest struct {
	Iaas *CoreScaleOutIaasResource `json:"iaas,omitempty"`
	Reserved *CoreScaleOutReservedResource `json:"reserved,omitempty"`
}

// NewCoreScaleOutServiceComponentRequest instantiates a new CoreScaleOutServiceComponentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreScaleOutServiceComponentRequest() *CoreScaleOutServiceComponentRequest {
	this := CoreScaleOutServiceComponentRequest{}
	return &this
}

// NewCoreScaleOutServiceComponentRequestWithDefaults instantiates a new CoreScaleOutServiceComponentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreScaleOutServiceComponentRequestWithDefaults() *CoreScaleOutServiceComponentRequest {
	this := CoreScaleOutServiceComponentRequest{}
	return &this
}

// GetIaas returns the Iaas field value if set, zero value otherwise.
func (o *CoreScaleOutServiceComponentRequest) GetIaas() CoreScaleOutIaasResource {
	if o == nil || o.Iaas == nil {
		var ret CoreScaleOutIaasResource
		return ret
	}
	return *o.Iaas
}

// GetIaasOk returns a tuple with the Iaas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreScaleOutServiceComponentRequest) GetIaasOk() (*CoreScaleOutIaasResource, bool) {
	if o == nil || o.Iaas == nil {
		return nil, false
	}
	return o.Iaas, true
}

// HasIaas returns a boolean if a field has been set.
func (o *CoreScaleOutServiceComponentRequest) HasIaas() bool {
	if o != nil && o.Iaas != nil {
		return true
	}

	return false
}

// SetIaas gets a reference to the given CoreScaleOutIaasResource and assigns it to the Iaas field.
func (o *CoreScaleOutServiceComponentRequest) SetIaas(v CoreScaleOutIaasResource) {
	o.Iaas = &v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *CoreScaleOutServiceComponentRequest) GetReserved() CoreScaleOutReservedResource {
	if o == nil || o.Reserved == nil {
		var ret CoreScaleOutReservedResource
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreScaleOutServiceComponentRequest) GetReservedOk() (*CoreScaleOutReservedResource, bool) {
	if o == nil || o.Reserved == nil {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *CoreScaleOutServiceComponentRequest) HasReserved() bool {
	if o != nil && o.Reserved != nil {
		return true
	}

	return false
}

// SetReserved gets a reference to the given CoreScaleOutReservedResource and assigns it to the Reserved field.
func (o *CoreScaleOutServiceComponentRequest) SetReserved(v CoreScaleOutReservedResource) {
	o.Reserved = &v
}

func (o CoreScaleOutServiceComponentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Iaas != nil {
		toSerialize["iaas"] = o.Iaas
	}
	if o.Reserved != nil {
		toSerialize["reserved"] = o.Reserved
	}
	return json.Marshal(toSerialize)
}

type NullableCoreScaleOutServiceComponentRequest struct {
	value *CoreScaleOutServiceComponentRequest
	isSet bool
}

func (v NullableCoreScaleOutServiceComponentRequest) Get() *CoreScaleOutServiceComponentRequest {
	return v.value
}

func (v *NullableCoreScaleOutServiceComponentRequest) Set(val *CoreScaleOutServiceComponentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreScaleOutServiceComponentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreScaleOutServiceComponentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreScaleOutServiceComponentRequest(val *CoreScaleOutServiceComponentRequest) *NullableCoreScaleOutServiceComponentRequest {
	return &NullableCoreScaleOutServiceComponentRequest{value: val, isSet: true}
}

func (v NullableCoreScaleOutServiceComponentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreScaleOutServiceComponentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


