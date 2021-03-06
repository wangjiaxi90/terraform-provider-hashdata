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

// CoreCreateOssZoneRequest struct for CoreCreateOssZoneRequest
type CoreCreateOssZoneRequest struct {
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	InternalRegion *string `json:"internal_region,omitempty"`
	Name *string `json:"name,omitempty"`
	PublicRegion *string `json:"public_region,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
}

// NewCoreCreateOssZoneRequest instantiates a new CoreCreateOssZoneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCreateOssZoneRequest() *CoreCreateOssZoneRequest {
	this := CoreCreateOssZoneRequest{}
	return &this
}

// NewCoreCreateOssZoneRequestWithDefaults instantiates a new CoreCreateOssZoneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCreateOssZoneRequestWithDefaults() *CoreCreateOssZoneRequest {
	this := CoreCreateOssZoneRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CoreCreateOssZoneRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateOssZoneRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CoreCreateOssZoneRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CoreCreateOssZoneRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CoreCreateOssZoneRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateOssZoneRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CoreCreateOssZoneRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CoreCreateOssZoneRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *CoreCreateOssZoneRequest) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateOssZoneRequest) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *CoreCreateOssZoneRequest) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *CoreCreateOssZoneRequest) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetInternalRegion returns the InternalRegion field value if set, zero value otherwise.
func (o *CoreCreateOssZoneRequest) GetInternalRegion() string {
	if o == nil || o.InternalRegion == nil {
		var ret string
		return ret
	}
	return *o.InternalRegion
}

// GetInternalRegionOk returns a tuple with the InternalRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateOssZoneRequest) GetInternalRegionOk() (*string, bool) {
	if o == nil || o.InternalRegion == nil {
		return nil, false
	}
	return o.InternalRegion, true
}

// HasInternalRegion returns a boolean if a field has been set.
func (o *CoreCreateOssZoneRequest) HasInternalRegion() bool {
	if o != nil && o.InternalRegion != nil {
		return true
	}

	return false
}

// SetInternalRegion gets a reference to the given string and assigns it to the InternalRegion field.
func (o *CoreCreateOssZoneRequest) SetInternalRegion(v string) {
	o.InternalRegion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreCreateOssZoneRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateOssZoneRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreCreateOssZoneRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreCreateOssZoneRequest) SetName(v string) {
	o.Name = &v
}

// GetPublicRegion returns the PublicRegion field value if set, zero value otherwise.
func (o *CoreCreateOssZoneRequest) GetPublicRegion() string {
	if o == nil || o.PublicRegion == nil {
		var ret string
		return ret
	}
	return *o.PublicRegion
}

// GetPublicRegionOk returns a tuple with the PublicRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateOssZoneRequest) GetPublicRegionOk() (*string, bool) {
	if o == nil || o.PublicRegion == nil {
		return nil, false
	}
	return o.PublicRegion, true
}

// HasPublicRegion returns a boolean if a field has been set.
func (o *CoreCreateOssZoneRequest) HasPublicRegion() bool {
	if o != nil && o.PublicRegion != nil {
		return true
	}

	return false
}

// SetPublicRegion gets a reference to the given string and assigns it to the PublicRegion field.
func (o *CoreCreateOssZoneRequest) SetPublicRegion(v string) {
	o.PublicRegion = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *CoreCreateOssZoneRequest) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateOssZoneRequest) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *CoreCreateOssZoneRequest) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *CoreCreateOssZoneRequest) SetVendor(v string) {
	o.Vendor = &v
}

func (o CoreCreateOssZoneRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.InternalRegion != nil {
		toSerialize["internal_region"] = o.InternalRegion
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PublicRegion != nil {
		toSerialize["public_region"] = o.PublicRegion
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableCoreCreateOssZoneRequest struct {
	value *CoreCreateOssZoneRequest
	isSet bool
}

func (v NullableCoreCreateOssZoneRequest) Get() *CoreCreateOssZoneRequest {
	return v.value
}

func (v *NullableCoreCreateOssZoneRequest) Set(val *CoreCreateOssZoneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCreateOssZoneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCreateOssZoneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCreateOssZoneRequest(val *CoreCreateOssZoneRequest) *NullableCoreCreateOssZoneRequest {
	return &NullableCoreCreateOssZoneRequest{value: val, isSet: true}
}

func (v NullableCoreCreateOssZoneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCreateOssZoneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


