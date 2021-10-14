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

// AccountCreateExternalTenantRequest struct for AccountCreateExternalTenantRequest
type AccountCreateExternalTenantRequest struct {
	ExternalId *string `json:"external_id,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
}

// NewAccountCreateExternalTenantRequest instantiates a new AccountCreateExternalTenantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCreateExternalTenantRequest() *AccountCreateExternalTenantRequest {
	this := AccountCreateExternalTenantRequest{}
	return &this
}

// NewAccountCreateExternalTenantRequestWithDefaults instantiates a new AccountCreateExternalTenantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCreateExternalTenantRequestWithDefaults() *AccountCreateExternalTenantRequest {
	this := AccountCreateExternalTenantRequest{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AccountCreateExternalTenantRequest) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateExternalTenantRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AccountCreateExternalTenantRequest) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AccountCreateExternalTenantRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *AccountCreateExternalTenantRequest) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateExternalTenantRequest) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *AccountCreateExternalTenantRequest) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *AccountCreateExternalTenantRequest) SetVendor(v string) {
	o.Vendor = &v
}

func (o AccountCreateExternalTenantRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableAccountCreateExternalTenantRequest struct {
	value *AccountCreateExternalTenantRequest
	isSet bool
}

func (v NullableAccountCreateExternalTenantRequest) Get() *AccountCreateExternalTenantRequest {
	return v.value
}

func (v *NullableAccountCreateExternalTenantRequest) Set(val *AccountCreateExternalTenantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCreateExternalTenantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCreateExternalTenantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCreateExternalTenantRequest(val *AccountCreateExternalTenantRequest) *NullableAccountCreateExternalTenantRequest {
	return &NullableAccountCreateExternalTenantRequest{value: val, isSet: true}
}

func (v NullableAccountCreateExternalTenantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCreateExternalTenantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


