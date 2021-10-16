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

// CoreListServiceBucketResponse struct for CoreListServiceBucketResponse
type CoreListServiceBucketResponse struct {
	Content *[]CoreDescribeBucketResponse `json:"content,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewCoreListServiceBucketResponse instantiates a new CoreListServiceBucketResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreListServiceBucketResponse() *CoreListServiceBucketResponse {
	this := CoreListServiceBucketResponse{}
	return &this
}

// NewCoreListServiceBucketResponseWithDefaults instantiates a new CoreListServiceBucketResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreListServiceBucketResponseWithDefaults() *CoreListServiceBucketResponse {
	this := CoreListServiceBucketResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *CoreListServiceBucketResponse) GetContent() []CoreDescribeBucketResponse {
	if o == nil || o.Content == nil {
		var ret []CoreDescribeBucketResponse
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListServiceBucketResponse) GetContentOk() (*[]CoreDescribeBucketResponse, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *CoreListServiceBucketResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []CoreDescribeBucketResponse and assigns it to the Content field.
func (o *CoreListServiceBucketResponse) SetContent(v []CoreDescribeBucketResponse) {
	o.Content = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CoreListServiceBucketResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListServiceBucketResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CoreListServiceBucketResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CoreListServiceBucketResponse) SetCount(v int32) {
	o.Count = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CoreListServiceBucketResponse) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListServiceBucketResponse) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *CoreListServiceBucketResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *CoreListServiceBucketResponse) SetTotal(v int32) {
	o.Total = &v
}

func (o CoreListServiceBucketResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableCoreListServiceBucketResponse struct {
	value *CoreListServiceBucketResponse
	isSet bool
}

func (v NullableCoreListServiceBucketResponse) Get() *CoreListServiceBucketResponse {
	return v.value
}

func (v *NullableCoreListServiceBucketResponse) Set(val *CoreListServiceBucketResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreListServiceBucketResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreListServiceBucketResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreListServiceBucketResponse(val *CoreListServiceBucketResponse) *NullableCoreListServiceBucketResponse {
	return &NullableCoreListServiceBucketResponse{value: val, isSet: true}
}

func (v NullableCoreListServiceBucketResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreListServiceBucketResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

