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

// VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains struct for VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains
type VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains struct {
	Spec *VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec `json:"spec,omitempty"`
}

// NewVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains instantiates a new VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains() *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains {
	this := VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains{}
	return &this
}

// NewVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsWithDefaults instantiates a new VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChainsWithDefaults() *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains {
	this := VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains) GetSpec() VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains) GetSpecOk() (*VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec and assigns it to the Spec field.
func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains) SetSpec(v VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec) {
	o.Spec = &v
}

func (o VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains struct {
	value *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains
	isSet bool
}

func (v NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains) Get() *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains {
	return v.value
}

func (v *NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains) Set(val *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains(val *VcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains) *NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains {
	return &NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains{value: val, isSet: true}
}

func (v NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterCertificateManagementVcenterTrustedRootChainsCreateCertificateManagementVcenterTrustedRootChains) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

