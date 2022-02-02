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

// VcenterHvcLinksCreateSpec struct for VcenterHvcLinksCreateSpec
type VcenterHvcLinksCreateSpec struct {
	// The PSC hostname for the domain to be linked. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	PscHostname string `json:"psc_hostname"`
	// The HTTPS port of the PSC to be linked. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Port *string `json:"port,omitempty"`
	// The domain to which the PSC belongs. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	DomainName string `json:"domain_name"`
	// The administrator username of the PSC. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Username string `json:"username"`
	// The administrator password of the PSC. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Password string `json:"password"`
	// The ssl thumbprint of the server. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	SslThumbprint *string `json:"ssl_thumbprint,omitempty"`
	// List of groups to be added to enable administrator access to. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	AdminGroups *[]string `json:"admin_groups,omitempty"`
}

// NewVcenterHvcLinksCreateSpec instantiates a new VcenterHvcLinksCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterHvcLinksCreateSpec(pscHostname string, domainName string, username string, password string) *VcenterHvcLinksCreateSpec {
	this := VcenterHvcLinksCreateSpec{}
	this.PscHostname = pscHostname
	this.DomainName = domainName
	this.Username = username
	this.Password = password
	return &this
}

// NewVcenterHvcLinksCreateSpecWithDefaults instantiates a new VcenterHvcLinksCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterHvcLinksCreateSpecWithDefaults() *VcenterHvcLinksCreateSpec {
	this := VcenterHvcLinksCreateSpec{}
	return &this
}

// GetPscHostname returns the PscHostname field value
func (o *VcenterHvcLinksCreateSpec) GetPscHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PscHostname
}

// GetPscHostnameOk returns a tuple with the PscHostname field value
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksCreateSpec) GetPscHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PscHostname, true
}

// SetPscHostname sets field value
func (o *VcenterHvcLinksCreateSpec) SetPscHostname(v string) {
	o.PscHostname = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *VcenterHvcLinksCreateSpec) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksCreateSpec) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *VcenterHvcLinksCreateSpec) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *VcenterHvcLinksCreateSpec) SetPort(v string) {
	o.Port = &v
}

// GetDomainName returns the DomainName field value
func (o *VcenterHvcLinksCreateSpec) GetDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksCreateSpec) GetDomainNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DomainName, true
}

// SetDomainName sets field value
func (o *VcenterHvcLinksCreateSpec) SetDomainName(v string) {
	o.DomainName = v
}

// GetUsername returns the Username field value
func (o *VcenterHvcLinksCreateSpec) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksCreateSpec) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *VcenterHvcLinksCreateSpec) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *VcenterHvcLinksCreateSpec) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksCreateSpec) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *VcenterHvcLinksCreateSpec) SetPassword(v string) {
	o.Password = v
}

// GetSslThumbprint returns the SslThumbprint field value if set, zero value otherwise.
func (o *VcenterHvcLinksCreateSpec) GetSslThumbprint() string {
	if o == nil || o.SslThumbprint == nil {
		var ret string
		return ret
	}
	return *o.SslThumbprint
}

// GetSslThumbprintOk returns a tuple with the SslThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksCreateSpec) GetSslThumbprintOk() (*string, bool) {
	if o == nil || o.SslThumbprint == nil {
		return nil, false
	}
	return o.SslThumbprint, true
}

// HasSslThumbprint returns a boolean if a field has been set.
func (o *VcenterHvcLinksCreateSpec) HasSslThumbprint() bool {
	if o != nil && o.SslThumbprint != nil {
		return true
	}

	return false
}

// SetSslThumbprint gets a reference to the given string and assigns it to the SslThumbprint field.
func (o *VcenterHvcLinksCreateSpec) SetSslThumbprint(v string) {
	o.SslThumbprint = &v
}

// GetAdminGroups returns the AdminGroups field value if set, zero value otherwise.
func (o *VcenterHvcLinksCreateSpec) GetAdminGroups() []string {
	if o == nil || o.AdminGroups == nil {
		var ret []string
		return ret
	}
	return *o.AdminGroups
}

// GetAdminGroupsOk returns a tuple with the AdminGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksCreateSpec) GetAdminGroupsOk() (*[]string, bool) {
	if o == nil || o.AdminGroups == nil {
		return nil, false
	}
	return o.AdminGroups, true
}

// HasAdminGroups returns a boolean if a field has been set.
func (o *VcenterHvcLinksCreateSpec) HasAdminGroups() bool {
	if o != nil && o.AdminGroups != nil {
		return true
	}

	return false
}

// SetAdminGroups gets a reference to the given []string and assigns it to the AdminGroups field.
func (o *VcenterHvcLinksCreateSpec) SetAdminGroups(v []string) {
	o.AdminGroups = &v
}

func (o VcenterHvcLinksCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["psc_hostname"] = o.PscHostname
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["domain_name"] = o.DomainName
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.SslThumbprint != nil {
		toSerialize["ssl_thumbprint"] = o.SslThumbprint
	}
	if o.AdminGroups != nil {
		toSerialize["admin_groups"] = o.AdminGroups
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterHvcLinksCreateSpec struct {
	value *VcenterHvcLinksCreateSpec
	isSet bool
}

func (v NullableVcenterHvcLinksCreateSpec) Get() *VcenterHvcLinksCreateSpec {
	return v.value
}

func (v *NullableVcenterHvcLinksCreateSpec) Set(val *VcenterHvcLinksCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterHvcLinksCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterHvcLinksCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterHvcLinksCreateSpec(val *VcenterHvcLinksCreateSpec) *NullableVcenterHvcLinksCreateSpec {
	return &NullableVcenterHvcLinksCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterHvcLinksCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterHvcLinksCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

