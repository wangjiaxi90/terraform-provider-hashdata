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

// OpsDescribeTransactionResponse struct for OpsDescribeTransactionResponse
type OpsDescribeTransactionResponse struct {
	Amount *string `json:"amount,omitempty"`
	Id *string `json:"id,omitempty"`
	OrderNum *string `json:"order_num,omitempty"`
	QrCode *string `json:"qr_code,omitempty"`
	Status *string `json:"status,omitempty"`
	Tenant *string `json:"tenant,omitempty"`
	TenantName *string `json:"tenant_name,omitempty"`
	TransactionChannel *string `json:"transaction_channel,omitempty"`
	TransactionDate *string `json:"transaction_date,omitempty"`
	TransactionSource *string `json:"transaction_source,omitempty"`
	TransactionType *string `json:"transaction_type,omitempty"`
}

// NewOpsDescribeTransactionResponse instantiates a new OpsDescribeTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsDescribeTransactionResponse() *OpsDescribeTransactionResponse {
	this := OpsDescribeTransactionResponse{}
	return &this
}

// NewOpsDescribeTransactionResponseWithDefaults instantiates a new OpsDescribeTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsDescribeTransactionResponseWithDefaults() *OpsDescribeTransactionResponse {
	this := OpsDescribeTransactionResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *OpsDescribeTransactionResponse) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeTransactionResponse) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *OpsDescribeTransactionResponse) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *OpsDescribeTransactionResponse) SetAmount(v string) {
	o.Amount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OpsDescribeTransactionResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeTransactionResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OpsDescribeTransactionResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OpsDescribeTransactionResponse) SetId(v string) {
	o.Id = &v
}

// GetOrderNum returns the OrderNum field value if set, zero value otherwise.
func (o *OpsDescribeTransactionResponse) GetOrderNum() string {
	if o == nil || o.OrderNum == nil {
		var ret string
		return ret
	}
	return *o.OrderNum
}

// GetOrderNumOk returns a tuple with the OrderNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeTransactionResponse) GetOrderNumOk() (*string, bool) {
	if o == nil || o.OrderNum == nil {
		return nil, false
	}
	return o.OrderNum, true
}

// HasOrderNum returns a boolean if a field has been set.
func (o *OpsDescribeTransactionResponse) HasOrderNum() bool {
	if o != nil && o.OrderNum != nil {
		return true
	}

	return false
}

// SetOrderNum gets a reference to the given string and assigns it to the OrderNum field.
func (o *OpsDescribeTransactionResponse) SetOrderNum(v string) {
	o.OrderNum = &v
}

// GetQrCode returns the QrCode field value if set, zero value otherwise.
func (o *OpsDescribeTransactionResponse) GetQrCode() string {
	if o == nil || o.QrCode == nil {
		var ret string
		return ret
	}
	return *o.QrCode
}

// GetQrCodeOk returns a tuple with the QrCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeTransactionResponse) GetQrCodeOk() (*string, bool) {
	if o == nil || o.QrCode == nil {
		return nil, false
	}
	return o.QrCode, true
}

// HasQrCode returns a boolean if a field has been set.
func (o *OpsDescribeTransactionResponse) HasQrCode() bool {
	if o != nil && o.QrCode != nil {
		return true
	}

	return false
}

// SetQrCode gets a reference to the given string and assigns it to the QrCode field.
func (o *OpsDescribeTransactionResponse) SetQrCode(v string) {
	o.QrCode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OpsDescribeTransactionResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeTransactionResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OpsDescribeTransactionResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OpsDescribeTransactionResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *OpsDescribeTransactionResponse) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeTransactionResponse) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *OpsDescribeTransactionResponse) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *OpsDescribeTransactionResponse) SetTenant(v string) {
	o.Tenant = &v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *OpsDescribeTransactionResponse) GetTenantName() string {
	if o == nil || o.TenantName == nil {
		var ret string
		return ret
	}
	return *o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeTransactionResponse) GetTenantNameOk() (*string, bool) {
	if o == nil || o.TenantName == nil {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *OpsDescribeTransactionResponse) HasTenantName() bool {
	if o != nil && o.TenantName != nil {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given string and assigns it to the TenantName field.
func (o *OpsDescribeTransactionResponse) SetTenantName(v string) {
	o.TenantName = &v
}

// GetTransactionChannel returns the TransactionChannel field value if set, zero value otherwise.
func (o *OpsDescribeTransactionResponse) GetTransactionChannel() string {
	if o == nil || o.TransactionChannel == nil {
		var ret string
		return ret
	}
	return *o.TransactionChannel
}

// GetTransactionChannelOk returns a tuple with the TransactionChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeTransactionResponse) GetTransactionChannelOk() (*string, bool) {
	if o == nil || o.TransactionChannel == nil {
		return nil, false
	}
	return o.TransactionChannel, true
}

// HasTransactionChannel returns a boolean if a field has been set.
func (o *OpsDescribeTransactionResponse) HasTransactionChannel() bool {
	if o != nil && o.TransactionChannel != nil {
		return true
	}

	return false
}

// SetTransactionChannel gets a reference to the given string and assigns it to the TransactionChannel field.
func (o *OpsDescribeTransactionResponse) SetTransactionChannel(v string) {
	o.TransactionChannel = &v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *OpsDescribeTransactionResponse) GetTransactionDate() string {
	if o == nil || o.TransactionDate == nil {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeTransactionResponse) GetTransactionDateOk() (*string, bool) {
	if o == nil || o.TransactionDate == nil {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *OpsDescribeTransactionResponse) HasTransactionDate() bool {
	if o != nil && o.TransactionDate != nil {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *OpsDescribeTransactionResponse) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

// GetTransactionSource returns the TransactionSource field value if set, zero value otherwise.
func (o *OpsDescribeTransactionResponse) GetTransactionSource() string {
	if o == nil || o.TransactionSource == nil {
		var ret string
		return ret
	}
	return *o.TransactionSource
}

// GetTransactionSourceOk returns a tuple with the TransactionSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeTransactionResponse) GetTransactionSourceOk() (*string, bool) {
	if o == nil || o.TransactionSource == nil {
		return nil, false
	}
	return o.TransactionSource, true
}

// HasTransactionSource returns a boolean if a field has been set.
func (o *OpsDescribeTransactionResponse) HasTransactionSource() bool {
	if o != nil && o.TransactionSource != nil {
		return true
	}

	return false
}

// SetTransactionSource gets a reference to the given string and assigns it to the TransactionSource field.
func (o *OpsDescribeTransactionResponse) SetTransactionSource(v string) {
	o.TransactionSource = &v
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *OpsDescribeTransactionResponse) GetTransactionType() string {
	if o == nil || o.TransactionType == nil {
		var ret string
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsDescribeTransactionResponse) GetTransactionTypeOk() (*string, bool) {
	if o == nil || o.TransactionType == nil {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *OpsDescribeTransactionResponse) HasTransactionType() bool {
	if o != nil && o.TransactionType != nil {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given string and assigns it to the TransactionType field.
func (o *OpsDescribeTransactionResponse) SetTransactionType(v string) {
	o.TransactionType = &v
}

func (o OpsDescribeTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrderNum != nil {
		toSerialize["order_num"] = o.OrderNum
	}
	if o.QrCode != nil {
		toSerialize["qr_code"] = o.QrCode
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.TenantName != nil {
		toSerialize["tenant_name"] = o.TenantName
	}
	if o.TransactionChannel != nil {
		toSerialize["transaction_channel"] = o.TransactionChannel
	}
	if o.TransactionDate != nil {
		toSerialize["transaction_date"] = o.TransactionDate
	}
	if o.TransactionSource != nil {
		toSerialize["transaction_source"] = o.TransactionSource
	}
	if o.TransactionType != nil {
		toSerialize["transaction_type"] = o.TransactionType
	}
	return json.Marshal(toSerialize)
}

type NullableOpsDescribeTransactionResponse struct {
	value *OpsDescribeTransactionResponse
	isSet bool
}

func (v NullableOpsDescribeTransactionResponse) Get() *OpsDescribeTransactionResponse {
	return v.value
}

func (v *NullableOpsDescribeTransactionResponse) Set(val *OpsDescribeTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsDescribeTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsDescribeTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsDescribeTransactionResponse(val *OpsDescribeTransactionResponse) *NullableOpsDescribeTransactionResponse {
	return &NullableOpsDescribeTransactionResponse{value: val, isSet: true}
}

func (v NullableOpsDescribeTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsDescribeTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


