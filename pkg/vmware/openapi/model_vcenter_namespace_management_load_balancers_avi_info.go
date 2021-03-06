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

// VcenterNamespaceManagementLoadBalancersAviInfo struct for VcenterNamespaceManagementLoadBalancersAviInfo
type VcenterNamespaceManagementLoadBalancersAviInfo struct {
	Server VcenterNamespaceManagementLoadBalancersServer `json:"server"`
	// An administrator user name for accessing the Avi Controller.
	Username string `json:"username"`
	// PEM-encoded CA certificate chain which is used to verify x509 certificates received from the server.
	CertificateAuthorityChain string `json:"certificate_authority_chain"`
}

// NewVcenterNamespaceManagementLoadBalancersAviInfo instantiates a new VcenterNamespaceManagementLoadBalancersAviInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementLoadBalancersAviInfo(server VcenterNamespaceManagementLoadBalancersServer, username string, certificateAuthorityChain string) *VcenterNamespaceManagementLoadBalancersAviInfo {
	this := VcenterNamespaceManagementLoadBalancersAviInfo{}
	this.Server = server
	this.Username = username
	this.CertificateAuthorityChain = certificateAuthorityChain
	return &this
}

// NewVcenterNamespaceManagementLoadBalancersAviInfoWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersAviInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementLoadBalancersAviInfoWithDefaults() *VcenterNamespaceManagementLoadBalancersAviInfo {
	this := VcenterNamespaceManagementLoadBalancersAviInfo{}
	return &this
}

// GetServer returns the Server field value
func (o *VcenterNamespaceManagementLoadBalancersAviInfo) GetServer() VcenterNamespaceManagementLoadBalancersServer {
	if o == nil {
		var ret VcenterNamespaceManagementLoadBalancersServer
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersAviInfo) GetServerOk() (*VcenterNamespaceManagementLoadBalancersServer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *VcenterNamespaceManagementLoadBalancersAviInfo) SetServer(v VcenterNamespaceManagementLoadBalancersServer) {
	o.Server = v
}

// GetUsername returns the Username field value
func (o *VcenterNamespaceManagementLoadBalancersAviInfo) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersAviInfo) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *VcenterNamespaceManagementLoadBalancersAviInfo) SetUsername(v string) {
	o.Username = v
}

// GetCertificateAuthorityChain returns the CertificateAuthorityChain field value
func (o *VcenterNamespaceManagementLoadBalancersAviInfo) GetCertificateAuthorityChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateAuthorityChain
}

// GetCertificateAuthorityChainOk returns a tuple with the CertificateAuthorityChain field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersAviInfo) GetCertificateAuthorityChainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CertificateAuthorityChain, true
}

// SetCertificateAuthorityChain sets field value
func (o *VcenterNamespaceManagementLoadBalancersAviInfo) SetCertificateAuthorityChain(v string) {
	o.CertificateAuthorityChain = v
}

func (o VcenterNamespaceManagementLoadBalancersAviInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["server"] = o.Server
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["certificate_authority_chain"] = o.CertificateAuthorityChain
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementLoadBalancersAviInfo struct {
	value *VcenterNamespaceManagementLoadBalancersAviInfo
	isSet bool
}

func (v NullableVcenterNamespaceManagementLoadBalancersAviInfo) Get() *VcenterNamespaceManagementLoadBalancersAviInfo {
	return v.value
}

func (v *NullableVcenterNamespaceManagementLoadBalancersAviInfo) Set(val *VcenterNamespaceManagementLoadBalancersAviInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementLoadBalancersAviInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementLoadBalancersAviInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementLoadBalancersAviInfo(val *VcenterNamespaceManagementLoadBalancersAviInfo) *NullableVcenterNamespaceManagementLoadBalancersAviInfo {
	return &NullableVcenterNamespaceManagementLoadBalancersAviInfo{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementLoadBalancersAviInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementLoadBalancersAviInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


