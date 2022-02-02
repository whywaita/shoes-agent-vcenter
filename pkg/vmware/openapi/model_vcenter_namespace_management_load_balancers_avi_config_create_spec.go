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

// VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec struct for VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec
type VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec struct {
	Server VcenterNamespaceManagementLoadBalancersServer `json:"server"`
	// An administrator user name for accessing the Avi Controller.
	Username string `json:"username"`
	// The password for the administrator user.
	Password string `json:"password"`
	// CertificateAuthorityChain contains PEM-encoded CA chain which is used to verify x509 certificates received from the server.
	CertificateAuthorityChain string `json:"certificate_authority_chain"`
}

// NewVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec instantiates a new VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec(server VcenterNamespaceManagementLoadBalancersServer, username string, password string, certificateAuthorityChain string) *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec {
	this := VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec{}
	this.Server = server
	this.Username = username
	this.Password = password
	this.CertificateAuthorityChain = certificateAuthorityChain
	return &this
}

// NewVcenterNamespaceManagementLoadBalancersAviConfigCreateSpecWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementLoadBalancersAviConfigCreateSpecWithDefaults() *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec {
	this := VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec{}
	return &this
}

// GetServer returns the Server field value
func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetServer() VcenterNamespaceManagementLoadBalancersServer {
	if o == nil {
		var ret VcenterNamespaceManagementLoadBalancersServer
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetServerOk() (*VcenterNamespaceManagementLoadBalancersServer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) SetServer(v VcenterNamespaceManagementLoadBalancersServer) {
	o.Server = v
}

// GetUsername returns the Username field value
func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) SetPassword(v string) {
	o.Password = v
}

// GetCertificateAuthorityChain returns the CertificateAuthorityChain field value
func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetCertificateAuthorityChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateAuthorityChain
}

// GetCertificateAuthorityChainOk returns a tuple with the CertificateAuthorityChain field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetCertificateAuthorityChainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertificateAuthorityChain, true
}

// SetCertificateAuthorityChain sets field value
func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) SetCertificateAuthorityChain(v string) {
	o.CertificateAuthorityChain = v
}

func (o VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["server"] = o.Server
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["certificate_authority_chain"] = o.CertificateAuthorityChain
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec struct {
	value *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec
	isSet bool
}

func (v NullableVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) Get() *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec {
	return v.value
}

func (v *NullableVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) Set(val *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec(val *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) *NullableVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec {
	return &NullableVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


