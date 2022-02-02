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

// VcenterDeploymentReplicatedSpec struct for VcenterDeploymentReplicatedSpec
type VcenterDeploymentReplicatedSpec struct {
	// The IP address or DNS resolvable name of the partner PSC appliance.
	PartnerHostname string `json:"partner_hostname"`
	// The HTTPS port of the external PSC appliance. If unset, port 443 will be used.
	HttpsPort *int64 `json:"https_port,omitempty"`
	// The SSO administrator account password.
	SsoAdminPassword string `json:"sso_admin_password"`
	// SHA1 thumbprint of the server SSL certificate will be used for verification. This field is only relevant if ReplicatedSpec.ssl-verify is unset or has the value true.
	SslThumbprint *string `json:"ssl_thumbprint,omitempty"`
	// SSL verification should be enabled or disabled. If unset, ssl_verify true will be used.
	SslVerify *bool `json:"ssl_verify,omitempty"`
}

// NewVcenterDeploymentReplicatedSpec instantiates a new VcenterDeploymentReplicatedSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentReplicatedSpec(partnerHostname string, ssoAdminPassword string) *VcenterDeploymentReplicatedSpec {
	this := VcenterDeploymentReplicatedSpec{}
	this.PartnerHostname = partnerHostname
	this.SsoAdminPassword = ssoAdminPassword
	return &this
}

// NewVcenterDeploymentReplicatedSpecWithDefaults instantiates a new VcenterDeploymentReplicatedSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentReplicatedSpecWithDefaults() *VcenterDeploymentReplicatedSpec {
	this := VcenterDeploymentReplicatedSpec{}
	return &this
}

// GetPartnerHostname returns the PartnerHostname field value
func (o *VcenterDeploymentReplicatedSpec) GetPartnerHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartnerHostname
}

// GetPartnerHostnameOk returns a tuple with the PartnerHostname field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentReplicatedSpec) GetPartnerHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PartnerHostname, true
}

// SetPartnerHostname sets field value
func (o *VcenterDeploymentReplicatedSpec) SetPartnerHostname(v string) {
	o.PartnerHostname = v
}

// GetHttpsPort returns the HttpsPort field value if set, zero value otherwise.
func (o *VcenterDeploymentReplicatedSpec) GetHttpsPort() int64 {
	if o == nil || o.HttpsPort == nil {
		var ret int64
		return ret
	}
	return *o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentReplicatedSpec) GetHttpsPortOk() (*int64, bool) {
	if o == nil || o.HttpsPort == nil {
		return nil, false
	}
	return o.HttpsPort, true
}

// HasHttpsPort returns a boolean if a field has been set.
func (o *VcenterDeploymentReplicatedSpec) HasHttpsPort() bool {
	if o != nil && o.HttpsPort != nil {
		return true
	}

	return false
}

// SetHttpsPort gets a reference to the given int64 and assigns it to the HttpsPort field.
func (o *VcenterDeploymentReplicatedSpec) SetHttpsPort(v int64) {
	o.HttpsPort = &v
}

// GetSsoAdminPassword returns the SsoAdminPassword field value
func (o *VcenterDeploymentReplicatedSpec) GetSsoAdminPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SsoAdminPassword
}

// GetSsoAdminPasswordOk returns a tuple with the SsoAdminPassword field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentReplicatedSpec) GetSsoAdminPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SsoAdminPassword, true
}

// SetSsoAdminPassword sets field value
func (o *VcenterDeploymentReplicatedSpec) SetSsoAdminPassword(v string) {
	o.SsoAdminPassword = v
}

// GetSslThumbprint returns the SslThumbprint field value if set, zero value otherwise.
func (o *VcenterDeploymentReplicatedSpec) GetSslThumbprint() string {
	if o == nil || o.SslThumbprint == nil {
		var ret string
		return ret
	}
	return *o.SslThumbprint
}

// GetSslThumbprintOk returns a tuple with the SslThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentReplicatedSpec) GetSslThumbprintOk() (*string, bool) {
	if o == nil || o.SslThumbprint == nil {
		return nil, false
	}
	return o.SslThumbprint, true
}

// HasSslThumbprint returns a boolean if a field has been set.
func (o *VcenterDeploymentReplicatedSpec) HasSslThumbprint() bool {
	if o != nil && o.SslThumbprint != nil {
		return true
	}

	return false
}

// SetSslThumbprint gets a reference to the given string and assigns it to the SslThumbprint field.
func (o *VcenterDeploymentReplicatedSpec) SetSslThumbprint(v string) {
	o.SslThumbprint = &v
}

// GetSslVerify returns the SslVerify field value if set, zero value otherwise.
func (o *VcenterDeploymentReplicatedSpec) GetSslVerify() bool {
	if o == nil || o.SslVerify == nil {
		var ret bool
		return ret
	}
	return *o.SslVerify
}

// GetSslVerifyOk returns a tuple with the SslVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentReplicatedSpec) GetSslVerifyOk() (*bool, bool) {
	if o == nil || o.SslVerify == nil {
		return nil, false
	}
	return o.SslVerify, true
}

// HasSslVerify returns a boolean if a field has been set.
func (o *VcenterDeploymentReplicatedSpec) HasSslVerify() bool {
	if o != nil && o.SslVerify != nil {
		return true
	}

	return false
}

// SetSslVerify gets a reference to the given bool and assigns it to the SslVerify field.
func (o *VcenterDeploymentReplicatedSpec) SetSslVerify(v bool) {
	o.SslVerify = &v
}

func (o VcenterDeploymentReplicatedSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["partner_hostname"] = o.PartnerHostname
	}
	if o.HttpsPort != nil {
		toSerialize["https_port"] = o.HttpsPort
	}
	if true {
		toSerialize["sso_admin_password"] = o.SsoAdminPassword
	}
	if o.SslThumbprint != nil {
		toSerialize["ssl_thumbprint"] = o.SslThumbprint
	}
	if o.SslVerify != nil {
		toSerialize["ssl_verify"] = o.SslVerify
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentReplicatedSpec struct {
	value *VcenterDeploymentReplicatedSpec
	isSet bool
}

func (v NullableVcenterDeploymentReplicatedSpec) Get() *VcenterDeploymentReplicatedSpec {
	return v.value
}

func (v *NullableVcenterDeploymentReplicatedSpec) Set(val *VcenterDeploymentReplicatedSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentReplicatedSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentReplicatedSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentReplicatedSpec(val *VcenterDeploymentReplicatedSpec) *NullableVcenterDeploymentReplicatedSpec {
	return &NullableVcenterDeploymentReplicatedSpec{value: val, isSet: true}
}

func (v NullableVcenterDeploymentReplicatedSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentReplicatedSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

