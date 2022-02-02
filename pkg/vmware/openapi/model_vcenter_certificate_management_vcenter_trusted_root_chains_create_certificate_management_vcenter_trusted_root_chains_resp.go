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

// VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp struct for VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp
type VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp struct {
	Value string `json:"value"`
}

// NewVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp instantiates a new VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp(value string) *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp {
	this := VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp{}
	this.Value = value
	return &this
}

// NewVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsRespWithDefaults instantiates a new VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsRespWithDefaults() *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp {
	this := VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp) SetValue(v string) {
	o.Value = v
}

func (o VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp struct {
	value *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp
	isSet bool
}

func (v NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp) Get() *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp {
	return v.value
}

func (v *NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp) Set(val *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp(val *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp) *NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp {
	return &NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp{value: val, isSet: true}
}

func (v NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


