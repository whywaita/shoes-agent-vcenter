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

// VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec struct for VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec
type VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec struct {
	// Public certificate used by every host in the cluster.
	Certificate string `json:"certificate"`
	// Private part of the certificate. If unset, the update request is for a public/private client certificate pair, not for a signed CSR.
	PrivateKey *string `json:"private_key,omitempty"`
}

// NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec(certificate string) *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec {
	this := VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec{}
	this.Certificate = certificate
	return &this
}

// NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec {
	this := VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec{}
	return &this
}

// GetCertificate returns the Certificate field value
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) GetCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) SetCertificate(v string) {
	o.Certificate = v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

func (o VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["certificate"] = o.Certificate
	}
	if o.PrivateKey != nil {
		toSerialize["private_key"] = o.PrivateKey
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec struct {
	value *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) Get() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) Set(val *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec(val *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec {
	return &NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersClientCertificateUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


