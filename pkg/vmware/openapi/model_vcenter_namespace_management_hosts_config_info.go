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

// VcenterNamespaceManagementHostsConfigInfo struct for VcenterNamespaceManagementHostsConfigInfo
type VcenterNamespaceManagementHostsConfigInfo struct {
	// True if vSphere Namespace feature is supported in this VC.
	NamespacesSupported bool `json:"namespaces_supported"`
	// True if vSphere Namespace feature is licensed on any hosts in this VC.
	NamespacesLicensed bool `json:"namespaces_licensed"`
}

// NewVcenterNamespaceManagementHostsConfigInfo instantiates a new VcenterNamespaceManagementHostsConfigInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementHostsConfigInfo(namespacesSupported bool, namespacesLicensed bool) *VcenterNamespaceManagementHostsConfigInfo {
	this := VcenterNamespaceManagementHostsConfigInfo{}
	this.NamespacesSupported = namespacesSupported
	this.NamespacesLicensed = namespacesLicensed
	return &this
}

// NewVcenterNamespaceManagementHostsConfigInfoWithDefaults instantiates a new VcenterNamespaceManagementHostsConfigInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementHostsConfigInfoWithDefaults() *VcenterNamespaceManagementHostsConfigInfo {
	this := VcenterNamespaceManagementHostsConfigInfo{}
	return &this
}

// GetNamespacesSupported returns the NamespacesSupported field value
func (o *VcenterNamespaceManagementHostsConfigInfo) GetNamespacesSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NamespacesSupported
}

// GetNamespacesSupportedOk returns a tuple with the NamespacesSupported field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementHostsConfigInfo) GetNamespacesSupportedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NamespacesSupported, true
}

// SetNamespacesSupported sets field value
func (o *VcenterNamespaceManagementHostsConfigInfo) SetNamespacesSupported(v bool) {
	o.NamespacesSupported = v
}

// GetNamespacesLicensed returns the NamespacesLicensed field value
func (o *VcenterNamespaceManagementHostsConfigInfo) GetNamespacesLicensed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NamespacesLicensed
}

// GetNamespacesLicensedOk returns a tuple with the NamespacesLicensed field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementHostsConfigInfo) GetNamespacesLicensedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NamespacesLicensed, true
}

// SetNamespacesLicensed sets field value
func (o *VcenterNamespaceManagementHostsConfigInfo) SetNamespacesLicensed(v bool) {
	o.NamespacesLicensed = v
}

func (o VcenterNamespaceManagementHostsConfigInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["namespaces_supported"] = o.NamespacesSupported
	}
	if true {
		toSerialize["namespaces_licensed"] = o.NamespacesLicensed
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementHostsConfigInfo struct {
	value *VcenterNamespaceManagementHostsConfigInfo
	isSet bool
}

func (v NullableVcenterNamespaceManagementHostsConfigInfo) Get() *VcenterNamespaceManagementHostsConfigInfo {
	return v.value
}

func (v *NullableVcenterNamespaceManagementHostsConfigInfo) Set(val *VcenterNamespaceManagementHostsConfigInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementHostsConfigInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementHostsConfigInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementHostsConfigInfo(val *VcenterNamespaceManagementHostsConfigInfo) *NullableVcenterNamespaceManagementHostsConfigInfo {
	return &NullableVcenterNamespaceManagementHostsConfigInfo{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementHostsConfigInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementHostsConfigInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


