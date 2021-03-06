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

// VcenterNamespaceManagementSupportBundleLocation struct for VcenterNamespaceManagementSupportBundleLocation
type VcenterNamespaceManagementSupportBundleLocation struct {
	// Support Bundle Download URL.
	Url string `json:"url"`
	WcpSupportBundleToken VcenterNamespaceManagementSupportBundleToken `json:"wcp_support_bundle_token"`
}

// NewVcenterNamespaceManagementSupportBundleLocation instantiates a new VcenterNamespaceManagementSupportBundleLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementSupportBundleLocation(url string, wcpSupportBundleToken VcenterNamespaceManagementSupportBundleToken) *VcenterNamespaceManagementSupportBundleLocation {
	this := VcenterNamespaceManagementSupportBundleLocation{}
	this.Url = url
	this.WcpSupportBundleToken = wcpSupportBundleToken
	return &this
}

// NewVcenterNamespaceManagementSupportBundleLocationWithDefaults instantiates a new VcenterNamespaceManagementSupportBundleLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementSupportBundleLocationWithDefaults() *VcenterNamespaceManagementSupportBundleLocation {
	this := VcenterNamespaceManagementSupportBundleLocation{}
	return &this
}

// GetUrl returns the Url field value
func (o *VcenterNamespaceManagementSupportBundleLocation) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSupportBundleLocation) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VcenterNamespaceManagementSupportBundleLocation) SetUrl(v string) {
	o.Url = v
}

// GetWcpSupportBundleToken returns the WcpSupportBundleToken field value
func (o *VcenterNamespaceManagementSupportBundleLocation) GetWcpSupportBundleToken() VcenterNamespaceManagementSupportBundleToken {
	if o == nil {
		var ret VcenterNamespaceManagementSupportBundleToken
		return ret
	}

	return o.WcpSupportBundleToken
}

// GetWcpSupportBundleTokenOk returns a tuple with the WcpSupportBundleToken field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSupportBundleLocation) GetWcpSupportBundleTokenOk() (*VcenterNamespaceManagementSupportBundleToken, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WcpSupportBundleToken, true
}

// SetWcpSupportBundleToken sets field value
func (o *VcenterNamespaceManagementSupportBundleLocation) SetWcpSupportBundleToken(v VcenterNamespaceManagementSupportBundleToken) {
	o.WcpSupportBundleToken = v
}

func (o VcenterNamespaceManagementSupportBundleLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["wcp_support_bundle_token"] = o.WcpSupportBundleToken
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementSupportBundleLocation struct {
	value *VcenterNamespaceManagementSupportBundleLocation
	isSet bool
}

func (v NullableVcenterNamespaceManagementSupportBundleLocation) Get() *VcenterNamespaceManagementSupportBundleLocation {
	return v.value
}

func (v *NullableVcenterNamespaceManagementSupportBundleLocation) Set(val *VcenterNamespaceManagementSupportBundleLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementSupportBundleLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementSupportBundleLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementSupportBundleLocation(val *VcenterNamespaceManagementSupportBundleLocation) *NullableVcenterNamespaceManagementSupportBundleLocation {
	return &NullableVcenterNamespaceManagementSupportBundleLocation{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementSupportBundleLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementSupportBundleLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


