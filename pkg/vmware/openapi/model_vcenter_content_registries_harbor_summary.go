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

// VcenterContentRegistriesHarborSummary struct for VcenterContentRegistriesHarborSummary
type VcenterContentRegistriesHarborSummary struct {
	// Identifier of the cluster.
	Cluster *string `json:"cluster,omitempty"`
	// Identifier of the registry.
	Registry string `json:"registry"`
	// Version of the registry.
	Version string `json:"version"`
	// URL to access the UI of the registry.
	UiAccessUrl string `json:"ui_access_url"`
}

// NewVcenterContentRegistriesHarborSummary instantiates a new VcenterContentRegistriesHarborSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterContentRegistriesHarborSummary(registry string, version string, uiAccessUrl string) *VcenterContentRegistriesHarborSummary {
	this := VcenterContentRegistriesHarborSummary{}
	this.Registry = registry
	this.Version = version
	this.UiAccessUrl = uiAccessUrl
	return &this
}

// NewVcenterContentRegistriesHarborSummaryWithDefaults instantiates a new VcenterContentRegistriesHarborSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterContentRegistriesHarborSummaryWithDefaults() *VcenterContentRegistriesHarborSummary {
	this := VcenterContentRegistriesHarborSummary{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VcenterContentRegistriesHarborSummary) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborSummary) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VcenterContentRegistriesHarborSummary) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *VcenterContentRegistriesHarborSummary) SetCluster(v string) {
	o.Cluster = &v
}

// GetRegistry returns the Registry field value
func (o *VcenterContentRegistriesHarborSummary) GetRegistry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborSummary) GetRegistryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Registry, true
}

// SetRegistry sets field value
func (o *VcenterContentRegistriesHarborSummary) SetRegistry(v string) {
	o.Registry = v
}

// GetVersion returns the Version field value
func (o *VcenterContentRegistriesHarborSummary) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborSummary) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VcenterContentRegistriesHarborSummary) SetVersion(v string) {
	o.Version = v
}

// GetUiAccessUrl returns the UiAccessUrl field value
func (o *VcenterContentRegistriesHarborSummary) GetUiAccessUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UiAccessUrl
}

// GetUiAccessUrlOk returns a tuple with the UiAccessUrl field value
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborSummary) GetUiAccessUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UiAccessUrl, true
}

// SetUiAccessUrl sets field value
func (o *VcenterContentRegistriesHarborSummary) SetUiAccessUrl(v string) {
	o.UiAccessUrl = v
}

func (o VcenterContentRegistriesHarborSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if true {
		toSerialize["registry"] = o.Registry
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["ui_access_url"] = o.UiAccessUrl
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterContentRegistriesHarborSummary struct {
	value *VcenterContentRegistriesHarborSummary
	isSet bool
}

func (v NullableVcenterContentRegistriesHarborSummary) Get() *VcenterContentRegistriesHarborSummary {
	return v.value
}

func (v *NullableVcenterContentRegistriesHarborSummary) Set(val *VcenterContentRegistriesHarborSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterContentRegistriesHarborSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterContentRegistriesHarborSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterContentRegistriesHarborSummary(val *VcenterContentRegistriesHarborSummary) *NullableVcenterContentRegistriesHarborSummary {
	return &NullableVcenterContentRegistriesHarborSummary{value: val, isSet: true}
}

func (v NullableVcenterContentRegistriesHarborSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterContentRegistriesHarborSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


