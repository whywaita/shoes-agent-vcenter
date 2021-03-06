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

// VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec struct for VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec
type VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec struct {
	// Public certificates of key server to trust. If unset, the trusted server certificates will not be updated.
	Certificates *[]string `json:"certificates,omitempty"`
}

// NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec {
	this := VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec{}
	return &this
}

// NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec {
	this := VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec{}
	return &this
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec) GetCertificates() []string {
	if o == nil || o.Certificates == nil {
		var ret []string
		return ret
	}
	return *o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec) GetCertificatesOk() (*[]string, bool) {
	if o == nil || o.Certificates == nil {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec) HasCertificates() bool {
	if o != nil && o.Certificates != nil {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []string and assigns it to the Certificates field.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec) SetCertificates(v []string) {
	o.Certificates = &v
}

func (o VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificates != nil {
		toSerialize["certificates"] = o.Certificates
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec struct {
	value *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec) Get() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec) Set(val *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec(val *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec) *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec {
	return &NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersTrustedPeerCertificatesUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


