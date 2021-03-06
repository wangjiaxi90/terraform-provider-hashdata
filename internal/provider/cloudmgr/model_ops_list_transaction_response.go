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

// OpsListTransactionResponse struct for OpsListTransactionResponse
type OpsListTransactionResponse struct {
	Content *[]OpsDescribeTransactionResponse `json:"content,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewOpsListTransactionResponse instantiates a new OpsListTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsListTransactionResponse() *OpsListTransactionResponse {
	this := OpsListTransactionResponse{}
	return &this
}

// NewOpsListTransactionResponseWithDefaults instantiates a new OpsListTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsListTransactionResponseWithDefaults() *OpsListTransactionResponse {
	this := OpsListTransactionResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *OpsListTransactionResponse) GetContent() []OpsDescribeTransactionResponse {
	if o == nil || o.Content == nil {
		var ret []OpsDescribeTransactionResponse
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsListTransactionResponse) GetContentOk() (*[]OpsDescribeTransactionResponse, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *OpsListTransactionResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []OpsDescribeTransactionResponse and assigns it to the Content field.
func (o *OpsListTransactionResponse) SetContent(v []OpsDescribeTransactionResponse) {
	o.Content = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *OpsListTransactionResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsListTransactionResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *OpsListTransactionResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *OpsListTransactionResponse) SetCount(v int32) {
	o.Count = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OpsListTransactionResponse) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsListTransactionResponse) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OpsListTransactionResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *OpsListTransactionResponse) SetTotal(v int32) {
	o.Total = &v
}

func (o OpsListTransactionResponse) MarshalJSON() ([]byte, error) {
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

type NullableOpsListTransactionResponse struct {
	value *OpsListTransactionResponse
	isSet bool
}

func (v NullableOpsListTransactionResponse) Get() *OpsListTransactionResponse {
	return v.value
}

func (v *NullableOpsListTransactionResponse) Set(val *OpsListTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsListTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsListTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsListTransactionResponse(val *OpsListTransactionResponse) *NullableOpsListTransactionResponse {
	return &NullableOpsListTransactionResponse{value: val, isSet: true}
}

func (v NullableOpsListTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsListTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


