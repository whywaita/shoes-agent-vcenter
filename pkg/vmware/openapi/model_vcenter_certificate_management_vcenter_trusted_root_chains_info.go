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

// VcenterCertificateManagementVcenterTrustedRootChainsInfo struct for VcenterCertificateManagementVcenterTrustedRootChainsInfo
type VcenterCertificateManagementVcenterTrustedRootChainsInfo struct {
	CertChain VcenterCertificateManagementX509CertChain `json:"cert_chain"`
}

// NewVcenterCertificateManagementVcenterTrustedRootChainsInfo instantiates a new VcenterCertificateManagementVcenterTrustedRootChainsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterCertificateManagementVcenterTrustedRootChainsInfo(certChain VcenterCertificateManagementX509CertChain) *VcenterCertificateManagementVcenterTrustedRootChainsInfo {
	this := VcenterCertificateManagementVcenterTrustedRootChainsInfo{}
	this.CertChain = certChain
	return &this
}

// NewVcenterCertificateManagementVcenterTrustedRootChainsInfoWithDefaults instantiates a new VcenterCertificateManagementVcenterTrustedRootChainsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterCertificateManagementVcenterTrustedRootChainsInfoWithDefaults() *VcenterCertificateManagementVcenterTrustedRootChainsInfo {
	this := VcenterCertificateManagementVcenterTrustedRootChainsInfo{}
	return &this
}

// GetCertChain returns the CertChain field value
func (o *VcenterCertificateManagementVcenterTrustedRootChainsInfo) GetCertChain() VcenterCertificateManagementX509CertChain {
	if o == nil {
		var ret VcenterCertificateManagementX509CertChain
		return ret
	}

	return o.CertChain
}

// GetCertChainOk returns a tuple with the CertChain field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTrustedRootChainsInfo) GetCertChainOk() (*VcenterCertificateManagementX509CertChain, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertChain, true
}

// SetCertChain sets field value
func (o *VcenterCertificateManagementVcenterTrustedRootChainsInfo) SetCertChain(v VcenterCertificateManagementX509CertChain) {
	o.CertChain = v
}

func (o VcenterCertificateManagementVcenterTrustedRootChainsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cert_chain"] = o.CertChain
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterCertificateManagementVcenterTrustedRootChainsInfo struct {
	value *VcenterCertificateManagementVcenterTrustedRootChainsInfo
	isSet bool
}

func (v NullableVcenterCertificateManagementVcenterTrustedRootChainsInfo) Get() *VcenterCertificateManagementVcenterTrustedRootChainsInfo {
	return v.value
}

func (v *NullableVcenterCertificateManagementVcenterTrustedRootChainsInfo) Set(val *VcenterCertificateManagementVcenterTrustedRootChainsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterCertificateManagementVcenterTrustedRootChainsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterCertificateManagementVcenterTrustedRootChainsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterCertificateManagementVcenterTrustedRootChainsInfo(val *VcenterCertificateManagementVcenterTrustedRootChainsInfo) *NullableVcenterCertificateManagementVcenterTrustedRootChainsInfo {
	return &NullableVcenterCertificateManagementVcenterTrustedRootChainsInfo{value: val, isSet: true}
}

func (v NullableVcenterCertificateManagementVcenterTrustedRootChainsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterCertificateManagementVcenterTrustedRootChainsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


