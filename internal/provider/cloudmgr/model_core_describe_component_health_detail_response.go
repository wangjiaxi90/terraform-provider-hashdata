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

// CoreDescribeComponentHealthDetailResponse struct for CoreDescribeComponentHealthDetailResponse
type CoreDescribeComponentHealthDetailResponse struct {
	ComponentHealthStatus *string `json:"component_health_status,omitempty"`
	ComponentHealthStatusDescription *string `json:"component_health_status_description,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Instance *string `json:"instance,omitempty"`
	InstanceHealthStatus *string `json:"instance_health_status,omitempty"`
	InstanceHealthStatusDescription *string `json:"instance_health_status_description,omitempty"`
	Ipaddresses *[]string `json:"ipaddresses,omitempty"`
}

// NewCoreDescribeComponentHealthDetailResponse instantiates a new CoreDescribeComponentHealthDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreDescribeComponentHealthDetailResponse() *CoreDescribeComponentHealthDetailResponse {
	this := CoreDescribeComponentHealthDetailResponse{}
	return &this
}

// NewCoreDescribeComponentHealthDetailResponseWithDefaults instantiates a new CoreDescribeComponentHealthDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreDescribeComponentHealthDetailResponseWithDefaults() *CoreDescribeComponentHealthDetailResponse {
	this := CoreDescribeComponentHealthDetailResponse{}
	return &this
}

// GetComponentHealthStatus returns the ComponentHealthStatus field value if set, zero value otherwise.
func (o *CoreDescribeComponentHealthDetailResponse) GetComponentHealthStatus() string {
	if o == nil || o.ComponentHealthStatus == nil {
		var ret string
		return ret
	}
	return *o.ComponentHealthStatus
}

// GetComponentHealthStatusOk returns a tuple with the ComponentHealthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeComponentHealthDetailResponse) GetComponentHealthStatusOk() (*string, bool) {
	if o == nil || o.ComponentHealthStatus == nil {
		return nil, false
	}
	return o.ComponentHealthStatus, true
}

// HasComponentHealthStatus returns a boolean if a field has been set.
func (o *CoreDescribeComponentHealthDetailResponse) HasComponentHealthStatus() bool {
	if o != nil && o.ComponentHealthStatus != nil {
		return true
	}

	return false
}

// SetComponentHealthStatus gets a reference to the given string and assigns it to the ComponentHealthStatus field.
func (o *CoreDescribeComponentHealthDetailResponse) SetComponentHealthStatus(v string) {
	o.ComponentHealthStatus = &v
}

// GetComponentHealthStatusDescription returns the ComponentHealthStatusDescription field value if set, zero value otherwise.
func (o *CoreDescribeComponentHealthDetailResponse) GetComponentHealthStatusDescription() string {
	if o == nil || o.ComponentHealthStatusDescription == nil {
		var ret string
		return ret
	}
	return *o.ComponentHealthStatusDescription
}

// GetComponentHealthStatusDescriptionOk returns a tuple with the ComponentHealthStatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeComponentHealthDetailResponse) GetComponentHealthStatusDescriptionOk() (*string, bool) {
	if o == nil || o.ComponentHealthStatusDescription == nil {
		return nil, false
	}
	return o.ComponentHealthStatusDescription, true
}

// HasComponentHealthStatusDescription returns a boolean if a field has been set.
func (o *CoreDescribeComponentHealthDetailResponse) HasComponentHealthStatusDescription() bool {
	if o != nil && o.ComponentHealthStatusDescription != nil {
		return true
	}

	return false
}

// SetComponentHealthStatusDescription gets a reference to the given string and assigns it to the ComponentHealthStatusDescription field.
func (o *CoreDescribeComponentHealthDetailResponse) SetComponentHealthStatusDescription(v string) {
	o.ComponentHealthStatusDescription = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *CoreDescribeComponentHealthDetailResponse) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeComponentHealthDetailResponse) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *CoreDescribeComponentHealthDetailResponse) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *CoreDescribeComponentHealthDetailResponse) SetHostname(v string) {
	o.Hostname = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *CoreDescribeComponentHealthDetailResponse) GetInstance() string {
	if o == nil || o.Instance == nil {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeComponentHealthDetailResponse) GetInstanceOk() (*string, bool) {
	if o == nil || o.Instance == nil {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *CoreDescribeComponentHealthDetailResponse) HasInstance() bool {
	if o != nil && o.Instance != nil {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *CoreDescribeComponentHealthDetailResponse) SetInstance(v string) {
	o.Instance = &v
}

// GetInstanceHealthStatus returns the InstanceHealthStatus field value if set, zero value otherwise.
func (o *CoreDescribeComponentHealthDetailResponse) GetInstanceHealthStatus() string {
	if o == nil || o.InstanceHealthStatus == nil {
		var ret string
		return ret
	}
	return *o.InstanceHealthStatus
}

// GetInstanceHealthStatusOk returns a tuple with the InstanceHealthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeComponentHealthDetailResponse) GetInstanceHealthStatusOk() (*string, bool) {
	if o == nil || o.InstanceHealthStatus == nil {
		return nil, false
	}
	return o.InstanceHealthStatus, true
}

// HasInstanceHealthStatus returns a boolean if a field has been set.
func (o *CoreDescribeComponentHealthDetailResponse) HasInstanceHealthStatus() bool {
	if o != nil && o.InstanceHealthStatus != nil {
		return true
	}

	return false
}

// SetInstanceHealthStatus gets a reference to the given string and assigns it to the InstanceHealthStatus field.
func (o *CoreDescribeComponentHealthDetailResponse) SetInstanceHealthStatus(v string) {
	o.InstanceHealthStatus = &v
}

// GetInstanceHealthStatusDescription returns the InstanceHealthStatusDescription field value if set, zero value otherwise.
func (o *CoreDescribeComponentHealthDetailResponse) GetInstanceHealthStatusDescription() string {
	if o == nil || o.InstanceHealthStatusDescription == nil {
		var ret string
		return ret
	}
	return *o.InstanceHealthStatusDescription
}

// GetInstanceHealthStatusDescriptionOk returns a tuple with the InstanceHealthStatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeComponentHealthDetailResponse) GetInstanceHealthStatusDescriptionOk() (*string, bool) {
	if o == nil || o.InstanceHealthStatusDescription == nil {
		return nil, false
	}
	return o.InstanceHealthStatusDescription, true
}

// HasInstanceHealthStatusDescription returns a boolean if a field has been set.
func (o *CoreDescribeComponentHealthDetailResponse) HasInstanceHealthStatusDescription() bool {
	if o != nil && o.InstanceHealthStatusDescription != nil {
		return true
	}

	return false
}

// SetInstanceHealthStatusDescription gets a reference to the given string and assigns it to the InstanceHealthStatusDescription field.
func (o *CoreDescribeComponentHealthDetailResponse) SetInstanceHealthStatusDescription(v string) {
	o.InstanceHealthStatusDescription = &v
}

// GetIpaddresses returns the Ipaddresses field value if set, zero value otherwise.
func (o *CoreDescribeComponentHealthDetailResponse) GetIpaddresses() []string {
	if o == nil || o.Ipaddresses == nil {
		var ret []string
		return ret
	}
	return *o.Ipaddresses
}

// GetIpaddressesOk returns a tuple with the Ipaddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeComponentHealthDetailResponse) GetIpaddressesOk() (*[]string, bool) {
	if o == nil || o.Ipaddresses == nil {
		return nil, false
	}
	return o.Ipaddresses, true
}

// HasIpaddresses returns a boolean if a field has been set.
func (o *CoreDescribeComponentHealthDetailResponse) HasIpaddresses() bool {
	if o != nil && o.Ipaddresses != nil {
		return true
	}

	return false
}

// SetIpaddresses gets a reference to the given []string and assigns it to the Ipaddresses field.
func (o *CoreDescribeComponentHealthDetailResponse) SetIpaddresses(v []string) {
	o.Ipaddresses = &v
}

func (o CoreDescribeComponentHealthDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ComponentHealthStatus != nil {
		toSerialize["component_health_status"] = o.ComponentHealthStatus
	}
	if o.ComponentHealthStatusDescription != nil {
		toSerialize["component_health_status_description"] = o.ComponentHealthStatusDescription
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Instance != nil {
		toSerialize["instance"] = o.Instance
	}
	if o.InstanceHealthStatus != nil {
		toSerialize["instance_health_status"] = o.InstanceHealthStatus
	}
	if o.InstanceHealthStatusDescription != nil {
		toSerialize["instance_health_status_description"] = o.InstanceHealthStatusDescription
	}
	if o.Ipaddresses != nil {
		toSerialize["ipaddresses"] = o.Ipaddresses
	}
	return json.Marshal(toSerialize)
}

type NullableCoreDescribeComponentHealthDetailResponse struct {
	value *CoreDescribeComponentHealthDetailResponse
	isSet bool
}

func (v NullableCoreDescribeComponentHealthDetailResponse) Get() *CoreDescribeComponentHealthDetailResponse {
	return v.value
}

func (v *NullableCoreDescribeComponentHealthDetailResponse) Set(val *CoreDescribeComponentHealthDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreDescribeComponentHealthDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreDescribeComponentHealthDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreDescribeComponentHealthDetailResponse(val *CoreDescribeComponentHealthDetailResponse) *NullableCoreDescribeComponentHealthDetailResponse {
	return &NullableCoreDescribeComponentHealthDetailResponse{value: val, isSet: true}
}

func (v NullableCoreDescribeComponentHealthDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreDescribeComponentHealthDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


