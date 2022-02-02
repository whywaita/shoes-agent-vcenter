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

// VcenterTrustedInfrastructureX509CertChain struct for VcenterTrustedInfrastructureX509CertChain
type VcenterTrustedInfrastructureX509CertChain struct {
	// Certificate chain in base64 format
	CertChain []string `json:"cert_chain"`
}

// NewVcenterTrustedInfrastructureX509CertChain instantiates a new VcenterTrustedInfrastructureX509CertChain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTrustedInfrastructureX509CertChain(certChain []string) *VcenterTrustedInfrastructureX509CertChain {
	this := VcenterTrustedInfrastructureX509CertChain{}
	this.CertChain = certChain
	return &this
}

// NewVcenterTrustedInfrastructureX509CertChainWithDefaults instantiates a new VcenterTrustedInfrastructureX509CertChain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTrustedInfrastructureX509CertChainWithDefaults() *VcenterTrustedInfrastructureX509CertChain {
	this := VcenterTrustedInfrastructureX509CertChain{}
	return &this
}

// GetCertChain returns the CertChain field value
func (o *VcenterTrustedInfrastructureX509CertChain) GetCertChain() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CertChain
}

// GetCertChainOk returns a tuple with the CertChain field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureX509CertChain) GetCertChainOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertChain, true
}

// SetCertChain sets field value
func (o *VcenterTrustedInfrastructureX509CertChain) SetCertChain(v []string) {
	o.CertChain = v
}

func (o VcenterTrustedInfrastructureX509CertChain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cert_chain"] = o.CertChain
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTrustedInfrastructureX509CertChain struct {
	value *VcenterTrustedInfrastructureX509CertChain
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureX509CertChain) Get() *VcenterTrustedInfrastructureX509CertChain {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureX509CertChain) Set(val *VcenterTrustedInfrastructureX509CertChain) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureX509CertChain) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureX509CertChain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureX509CertChain(val *VcenterTrustedInfrastructureX509CertChain) *NullableVcenterTrustedInfrastructureX509CertChain {
	return &NullableVcenterTrustedInfrastructureX509CertChain{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureX509CertChain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureX509CertChain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


