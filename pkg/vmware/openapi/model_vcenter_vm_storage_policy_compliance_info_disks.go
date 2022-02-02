/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VcenterVmStoragePolicyComplianceInfoDisks struct for VcenterVmStoragePolicyComplianceInfoDisks
type VcenterVmStoragePolicyComplianceInfoDisks struct {
	Key *string `json:"key,omitempty"`
	Value *VcenterVmStoragePolicyComplianceVmComplianceInfo `json:"value,omitempty"`
}

// NewVcenterVmStoragePolicyComplianceInfoDisks instantiates a new VcenterVmStoragePolicyComplianceInfoDisks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmStoragePolicyComplianceInfoDisks() *VcenterVmStoragePolicyComplianceInfoDisks {
	this := VcenterVmStoragePolicyComplianceInfoDisks{}
	return &this
}

// NewVcenterVmStoragePolicyComplianceInfoDisksWithDefaults instantiates a new VcenterVmStoragePolicyComplianceInfoDisks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmStoragePolicyComplianceInfoDisksWithDefaults() *VcenterVmStoragePolicyComplianceInfoDisks {
	this := VcenterVmStoragePolicyComplianceInfoDisks{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VcenterVmStoragePolicyComplianceInfoDisks) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmStoragePolicyComplianceInfoDisks) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VcenterVmStoragePolicyComplianceInfoDisks) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *VcenterVmStoragePolicyComplianceInfoDisks) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VcenterVmStoragePolicyComplianceInfoDisks) GetValue() VcenterVmStoragePolicyComplianceVmComplianceInfo {
	if o == nil || o.Value == nil {
		var ret VcenterVmStoragePolicyComplianceVmComplianceInfo
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmStoragePolicyComplianceInfoDisks) GetValueOk() (*VcenterVmStoragePolicyComplianceVmComplianceInfo, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VcenterVmStoragePolicyComplianceInfoDisks) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given VcenterVmStoragePolicyComplianceVmComplianceInfo and assigns it to the Value field.
func (o *VcenterVmStoragePolicyComplianceInfoDisks) SetValue(v VcenterVmStoragePolicyComplianceVmComplianceInfo) {
	o.Value = &v
}

func (o VcenterVmStoragePolicyComplianceInfoDisks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmStoragePolicyComplianceInfoDisks struct {
	value *VcenterVmStoragePolicyComplianceInfoDisks
	isSet bool
}

func (v NullableVcenterVmStoragePolicyComplianceInfoDisks) Get() *VcenterVmStoragePolicyComplianceInfoDisks {
	return v.value
}

func (v *NullableVcenterVmStoragePolicyComplianceInfoDisks) Set(val *VcenterVmStoragePolicyComplianceInfoDisks) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmStoragePolicyComplianceInfoDisks) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmStoragePolicyComplianceInfoDisks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmStoragePolicyComplianceInfoDisks(val *VcenterVmStoragePolicyComplianceInfoDisks) *NullableVcenterVmStoragePolicyComplianceInfoDisks {
	return &NullableVcenterVmStoragePolicyComplianceInfoDisks{value: val, isSet: true}
}

func (v NullableVcenterVmStoragePolicyComplianceInfoDisks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmStoragePolicyComplianceInfoDisks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

