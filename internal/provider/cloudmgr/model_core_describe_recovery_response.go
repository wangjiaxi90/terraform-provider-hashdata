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

// CoreDescribeRecoveryResponse struct for CoreDescribeRecoveryResponse
type CoreDescribeRecoveryResponse struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Jobs *[]string `json:"jobs,omitempty"`
	Message *string `json:"message,omitempty"`
	Policies *[]string `json:"policies,omitempty"`
	Properties *map[string]interface{} `json:"properties,omitempty"`
	RecoveryPolicyName *string `json:"recoveryPolicyName,omitempty"`
	Service *string `json:"service,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewCoreDescribeRecoveryResponse instantiates a new CoreDescribeRecoveryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreDescribeRecoveryResponse() *CoreDescribeRecoveryResponse {
	this := CoreDescribeRecoveryResponse{}
	return &this
}

// NewCoreDescribeRecoveryResponseWithDefaults instantiates a new CoreDescribeRecoveryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreDescribeRecoveryResponseWithDefaults() *CoreDescribeRecoveryResponse {
	this := CoreDescribeRecoveryResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CoreDescribeRecoveryResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeRecoveryResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CoreDescribeRecoveryResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CoreDescribeRecoveryResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreDescribeRecoveryResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeRecoveryResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreDescribeRecoveryResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CoreDescribeRecoveryResponse) SetId(v string) {
	o.Id = &v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *CoreDescribeRecoveryResponse) GetJobs() []string {
	if o == nil || o.Jobs == nil {
		var ret []string
		return ret
	}
	return *o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeRecoveryResponse) GetJobsOk() (*[]string, bool) {
	if o == nil || o.Jobs == nil {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *CoreDescribeRecoveryResponse) HasJobs() bool {
	if o != nil && o.Jobs != nil {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []string and assigns it to the Jobs field.
func (o *CoreDescribeRecoveryResponse) SetJobs(v []string) {
	o.Jobs = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CoreDescribeRecoveryResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeRecoveryResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CoreDescribeRecoveryResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CoreDescribeRecoveryResponse) SetMessage(v string) {
	o.Message = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *CoreDescribeRecoveryResponse) GetPolicies() []string {
	if o == nil || o.Policies == nil {
		var ret []string
		return ret
	}
	return *o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeRecoveryResponse) GetPoliciesOk() (*[]string, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *CoreDescribeRecoveryResponse) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []string and assigns it to the Policies field.
func (o *CoreDescribeRecoveryResponse) SetPolicies(v []string) {
	o.Policies = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CoreDescribeRecoveryResponse) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeRecoveryResponse) GetPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CoreDescribeRecoveryResponse) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *CoreDescribeRecoveryResponse) SetProperties(v map[string]interface{}) {
	o.Properties = &v
}

// GetRecoveryPolicyName returns the RecoveryPolicyName field value if set, zero value otherwise.
func (o *CoreDescribeRecoveryResponse) GetRecoveryPolicyName() string {
	if o == nil || o.RecoveryPolicyName == nil {
		var ret string
		return ret
	}
	return *o.RecoveryPolicyName
}

// GetRecoveryPolicyNameOk returns a tuple with the RecoveryPolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeRecoveryResponse) GetRecoveryPolicyNameOk() (*string, bool) {
	if o == nil || o.RecoveryPolicyName == nil {
		return nil, false
	}
	return o.RecoveryPolicyName, true
}

// HasRecoveryPolicyName returns a boolean if a field has been set.
func (o *CoreDescribeRecoveryResponse) HasRecoveryPolicyName() bool {
	if o != nil && o.RecoveryPolicyName != nil {
		return true
	}

	return false
}

// SetRecoveryPolicyName gets a reference to the given string and assigns it to the RecoveryPolicyName field.
func (o *CoreDescribeRecoveryResponse) SetRecoveryPolicyName(v string) {
	o.RecoveryPolicyName = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *CoreDescribeRecoveryResponse) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeRecoveryResponse) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *CoreDescribeRecoveryResponse) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *CoreDescribeRecoveryResponse) SetService(v string) {
	o.Service = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CoreDescribeRecoveryResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeRecoveryResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CoreDescribeRecoveryResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CoreDescribeRecoveryResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CoreDescribeRecoveryResponse) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeRecoveryResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CoreDescribeRecoveryResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *CoreDescribeRecoveryResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o CoreDescribeRecoveryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Jobs != nil {
		toSerialize["jobs"] = o.Jobs
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.RecoveryPolicyName != nil {
		toSerialize["recoveryPolicyName"] = o.RecoveryPolicyName
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableCoreDescribeRecoveryResponse struct {
	value *CoreDescribeRecoveryResponse
	isSet bool
}

func (v NullableCoreDescribeRecoveryResponse) Get() *CoreDescribeRecoveryResponse {
	return v.value
}

func (v *NullableCoreDescribeRecoveryResponse) Set(val *CoreDescribeRecoveryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreDescribeRecoveryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreDescribeRecoveryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreDescribeRecoveryResponse(val *CoreDescribeRecoveryResponse) *NullableCoreDescribeRecoveryResponse {
	return &NullableCoreDescribeRecoveryResponse{value: val, isSet: true}
}

func (v NullableCoreDescribeRecoveryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreDescribeRecoveryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


