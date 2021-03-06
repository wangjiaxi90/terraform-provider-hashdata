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

// OpsListBucketInfoResponse struct for OpsListBucketInfoResponse
type OpsListBucketInfoResponse struct {
	BucketInfo *[]OpsDescribeBucketInfoResponse `json:"bucket_info,omitempty"`
}

// NewOpsListBucketInfoResponse instantiates a new OpsListBucketInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsListBucketInfoResponse() *OpsListBucketInfoResponse {
	this := OpsListBucketInfoResponse{}
	return &this
}

// NewOpsListBucketInfoResponseWithDefaults instantiates a new OpsListBucketInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsListBucketInfoResponseWithDefaults() *OpsListBucketInfoResponse {
	this := OpsListBucketInfoResponse{}
	return &this
}

// GetBucketInfo returns the BucketInfo field value if set, zero value otherwise.
func (o *OpsListBucketInfoResponse) GetBucketInfo() []OpsDescribeBucketInfoResponse {
	if o == nil || o.BucketInfo == nil {
		var ret []OpsDescribeBucketInfoResponse
		return ret
	}
	return *o.BucketInfo
}

// GetBucketInfoOk returns a tuple with the BucketInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsListBucketInfoResponse) GetBucketInfoOk() (*[]OpsDescribeBucketInfoResponse, bool) {
	if o == nil || o.BucketInfo == nil {
		return nil, false
	}
	return o.BucketInfo, true
}

// HasBucketInfo returns a boolean if a field has been set.
func (o *OpsListBucketInfoResponse) HasBucketInfo() bool {
	if o != nil && o.BucketInfo != nil {
		return true
	}

	return false
}

// SetBucketInfo gets a reference to the given []OpsDescribeBucketInfoResponse and assigns it to the BucketInfo field.
func (o *OpsListBucketInfoResponse) SetBucketInfo(v []OpsDescribeBucketInfoResponse) {
	o.BucketInfo = &v
}

func (o OpsListBucketInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BucketInfo != nil {
		toSerialize["bucket_info"] = o.BucketInfo
	}
	return json.Marshal(toSerialize)
}

type NullableOpsListBucketInfoResponse struct {
	value *OpsListBucketInfoResponse
	isSet bool
}

func (v NullableOpsListBucketInfoResponse) Get() *OpsListBucketInfoResponse {
	return v.value
}

func (v *NullableOpsListBucketInfoResponse) Set(val *OpsListBucketInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsListBucketInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsListBucketInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsListBucketInfoResponse(val *OpsListBucketInfoResponse) *NullableOpsListBucketInfoResponse {
	return &NullableOpsListBucketInfoResponse{value: val, isSet: true}
}

func (v NullableOpsListBucketInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsListBucketInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


