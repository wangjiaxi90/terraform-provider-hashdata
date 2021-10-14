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

// OpsDescribeCommentResponse struct for OpsDescribeCommentResponse
type OpsDescribeCommentResponse struct {
	Content *string `json:"content,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Id *string `json:"id,omitempty"`
	SenderId *string `json:"sender_id,omitempty"`
	SenderName *string `json:"sender_name,omitempty"`
}

// NewOpsDescribeCommentResponse instantiates a new OpsDescribeCommentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsDescribeCommentResponse() *OpsDescribeCommentResponse {
	this := OpsDescribeCommentResponse{}
	return &this
}

// NewOpsDescribeCommentResponseWithDefaults instantiates a new OpsDescribeCommentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsDescribeCommentResponseWithDefaults() *OpsDescribeCommentResponse {
	this := OpsDescribeCommentResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *OpsDescribeCommentResponse) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeCommentResponse) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *OpsDescribeCommentResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *OpsDescribeCommentResponse) SetContent(v string) {
	o.Content = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OpsDescribeCommentResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeCommentResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OpsDescribeCommentResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *OpsDescribeCommentResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OpsDescribeCommentResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeCommentResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OpsDescribeCommentResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OpsDescribeCommentResponse) SetId(v string) {
	o.Id = &v
}

// GetSenderId returns the SenderId field value if set, zero value otherwise.
func (o *OpsDescribeCommentResponse) GetSenderId() string {
	if o == nil || o.SenderId == nil {
		var ret string
		return ret
	}
	return *o.SenderId
}

// GetSenderIdOk returns a tuple with the SenderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeCommentResponse) GetSenderIdOk() (*string, bool) {
	if o == nil || o.SenderId == nil {
		return nil, false
	}
	return o.SenderId, true
}

// HasSenderId returns a boolean if a field has been set.
func (o *OpsDescribeCommentResponse) HasSenderId() bool {
	if o != nil && o.SenderId != nil {
		return true
	}

	return false
}

// SetSenderId gets a reference to the given string and assigns it to the SenderId field.
func (o *OpsDescribeCommentResponse) SetSenderId(v string) {
	o.SenderId = &v
}

// GetSenderName returns the SenderName field value if set, zero value otherwise.
func (o *OpsDescribeCommentResponse) GetSenderName() string {
	if o == nil || o.SenderName == nil {
		var ret string
		return ret
	}
	return *o.SenderName
}

// GetSenderNameOk returns a tuple with the SenderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeCommentResponse) GetSenderNameOk() (*string, bool) {
	if o == nil || o.SenderName == nil {
		return nil, false
	}
	return o.SenderName, true
}

// HasSenderName returns a boolean if a field has been set.
func (o *OpsDescribeCommentResponse) HasSenderName() bool {
	if o != nil && o.SenderName != nil {
		return true
	}

	return false
}

// SetSenderName gets a reference to the given string and assigns it to the SenderName field.
func (o *OpsDescribeCommentResponse) SetSenderName(v string) {
	o.SenderName = &v
}

func (o OpsDescribeCommentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SenderId != nil {
		toSerialize["sender_id"] = o.SenderId
	}
	if o.SenderName != nil {
		toSerialize["sender_name"] = o.SenderName
	}
	return json.Marshal(toSerialize)
}

type NullableOpsDescribeCommentResponse struct {
	value *OpsDescribeCommentResponse
	isSet bool
}

func (v NullableOpsDescribeCommentResponse) Get() *OpsDescribeCommentResponse {
	return v.value
}

func (v *NullableOpsDescribeCommentResponse) Set(val *OpsDescribeCommentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsDescribeCommentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsDescribeCommentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsDescribeCommentResponse(val *OpsDescribeCommentResponse) *NullableOpsDescribeCommentResponse {
	return &NullableOpsDescribeCommentResponse{value: val, isSet: true}
}

func (v NullableOpsDescribeCommentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsDescribeCommentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


