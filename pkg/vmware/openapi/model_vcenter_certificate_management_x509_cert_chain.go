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

// VcenterCertificateManagementX509CertChain struct for VcenterCertificateManagementX509CertChain
type VcenterCertificateManagementX509CertChain struct {
	// Certificate chain in base64 format
	CertChain []string `json:"cert_chain"`
}

// NewVcenterCertificateManagementX509CertChain instantiates a new VcenterCertificateManagementX509CertChain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterCertificateManagementX509CertChain(certChain []string) *VcenterCertificateManagementX509CertChain {
	this := VcenterCertificateManagementX509CertChain{}
	this.CertChain = certChain
	return &this
}

// NewVcenterCertificateManagementX509CertChainWithDefaults instantiates a new VcenterCertificateManagementX509CertChain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterCertificateManagementX509CertChainWithDefaults() *VcenterCertificateManagementX509CertChain {
	this := VcenterCertificateManagementX509CertChain{}
	return &this
}

// GetCertChain returns the CertChain field value
func (o *VcenterCertificateManagementX509CertChain) GetCertChain() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CertChain
}

// GetCertChainOk returns a tuple with the CertChain field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementX509CertChain) GetCertChainOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertChain, true
}

// SetCertChain sets field value
func (o *VcenterCertificateManagementX509CertChain) SetCertChain(v []string) {
	o.CertChain = v
}

func (o VcenterCertificateManagementX509CertChain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cert_chain"] = o.CertChain
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterCertificateManagementX509CertChain struct {
	value *VcenterCertificateManagementX509CertChain
	isSet bool
}

func (v NullableVcenterCertificateManagementX509CertChain) Get() *VcenterCertificateManagementX509CertChain {
	return v.value
}

func (v *NullableVcenterCertificateManagementX509CertChain) Set(val *VcenterCertificateManagementX509CertChain) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterCertificateManagementX509CertChain) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterCertificateManagementX509CertChain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterCertificateManagementX509CertChain(val *VcenterCertificateManagementX509CertChain) *NullableVcenterCertificateManagementX509CertChain {
	return &NullableVcenterCertificateManagementX509CertChain{value: val, isSet: true}
}

func (v NullableVcenterCertificateManagementX509CertChain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterCertificateManagementX509CertChain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


