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

// VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec struct for VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec
type VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec struct {
	// An administrator user name for accessing the HAProxy Data Plane API server. If unset, the existing username will not be modified.
	Username *string `json:"username,omitempty"`
	// The password for the administrator user. If unset, the existing password will not be modified.
	Password *string `json:"password,omitempty"`
	// CertificateAuthorityChain contains PEM-encoded CA chain which is used to verify x509 certificates received from the server. If unset, the existing PEM-encoded CA chain will not be modified.
	CertificateAuthorityChain *string `json:"certificate_authority_chain,omitempty"`
}

// NewVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec instantiates a new VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec() *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec {
	this := VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec{}
	return &this
}

// NewVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpecWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpecWithDefaults() *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec {
	this := VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) SetPassword(v string) {
	o.Password = &v
}

// GetCertificateAuthorityChain returns the CertificateAuthorityChain field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) GetCertificateAuthorityChain() string {
	if o == nil || o.CertificateAuthorityChain == nil {
		var ret string
		return ret
	}
	return *o.CertificateAuthorityChain
}

// GetCertificateAuthorityChainOk returns a tuple with the CertificateAuthorityChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) GetCertificateAuthorityChainOk() (*string, bool) {
	if o == nil || o.CertificateAuthorityChain == nil {
		return nil, false
	}
	return o.CertificateAuthorityChain, true
}

// HasCertificateAuthorityChain returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) HasCertificateAuthorityChain() bool {
	if o != nil && o.CertificateAuthorityChain != nil {
		return true
	}

	return false
}

// SetCertificateAuthorityChain gets a reference to the given string and assigns it to the CertificateAuthorityChain field.
func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) SetCertificateAuthorityChain(v string) {
	o.CertificateAuthorityChain = &v
}

func (o VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.CertificateAuthorityChain != nil {
		toSerialize["certificate_authority_chain"] = o.CertificateAuthorityChain
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec struct {
	value *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec
	isSet bool
}

func (v NullableVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) Get() *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec {
	return v.value
}

func (v *NullableVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) Set(val *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec(val *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) *NullableVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec {
	return &NullableVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

