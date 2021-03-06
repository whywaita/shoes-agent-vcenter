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

// VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec struct for VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec
type VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec struct {
	// Identifier of the Supervisor Service version which contains the service definition. This Supervisor Service version must be in the State#ACTIVATED state. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version.
	Version string `json:"version"`
	// A generic key-value map for additional configuration parameters required during service upgrade. As an example, a third party operator might reference a private registry using parameters such as \"registryName\" for the registry name, \"registryUsername\" and \"registryPassword\" for the registry credentials. If unset, no additional configuration parameters will be applied when upgrading a Supervisor Service in the vSphere Supervisor cluster.
	ServiceConfig *map[string]string `json:"service_config,omitempty"`
}

// NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec instantiates a new VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec(version string) *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec {
	this := VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec{}
	this.Version = version
	return &this
}

// NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpecWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpecWithDefaults() *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec {
	this := VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec{}
	return &this
}

// GetVersion returns the Version field value
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) SetVersion(v string) {
	o.Version = v
}

// GetServiceConfig returns the ServiceConfig field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) GetServiceConfig() map[string]string {
	if o == nil || o.ServiceConfig == nil {
		var ret map[string]string
		return ret
	}
	return *o.ServiceConfig
}

// GetServiceConfigOk returns a tuple with the ServiceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) GetServiceConfigOk() (*map[string]string, bool) {
	if o == nil || o.ServiceConfig == nil {
		return nil, false
	}
	return o.ServiceConfig, true
}

// HasServiceConfig returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) HasServiceConfig() bool {
	if o != nil && o.ServiceConfig != nil {
		return true
	}

	return false
}

// SetServiceConfig gets a reference to the given map[string]string and assigns it to the ServiceConfig field.
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) SetServiceConfig(v map[string]string) {
	o.ServiceConfig = &v
}

func (o VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.ServiceConfig != nil {
		toSerialize["service_config"] = o.ServiceConfig
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec struct {
	value *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec
	isSet bool
}

func (v NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) Get() *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec {
	return v.value
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) Set(val *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec(val *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) *NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec {
	return &NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSetSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


