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

// VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp struct for VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp
type VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp struct {
	Value []VcenterCertificateManagementVcenterTrustedRootChainsSummary `json:"value"`
}

// NewVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp instantiates a new VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp(value []VcenterCertificateManagementVcenterTrustedRootChainsSummary) *VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp {
	this := VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp{}
	this.Value = value
	return &this
}

// NewVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsRespWithDefaults instantiates a new VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsRespWithDefaults() *VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp {
	this := VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp) GetValue() []VcenterCertificateManagementVcenterTrustedRootChainsSummary {
	if o == nil {
		var ret []VcenterCertificateManagementVcenterTrustedRootChainsSummary
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp) GetValueOk() (*[]VcenterCertificateManagementVcenterTrustedRootChainsSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp) SetValue(v []VcenterCertificateManagementVcenterTrustedRootChainsSummary) {
	o.Value = v
}

func (o VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp struct {
	value *VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp
	isSet bool
}

func (v NullableVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp) Get() *VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp {
	return v.value
}

func (v *NullableVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp) Set(val *VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp(val *VcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp) *NullableVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp {
	return &NullableVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp{value: val, isSet: true}
}

func (v NullableVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterCertificateManagementVcenterTrustedRootChainsListCertificateManagementVcenterTrustedRootChainsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

