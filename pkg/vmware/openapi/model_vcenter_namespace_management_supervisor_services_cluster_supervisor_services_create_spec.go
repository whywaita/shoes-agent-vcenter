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

// VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec struct for VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec
type VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec struct {
	// Identifier of the Supervisor Service. This Supervisor Service must be in the State#ACTIVATED state. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.SupervisorService. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.SupervisorService.
	SupervisorService string `json:"supervisor_service"`
	// Identifier of the Supervisor Service version which contains the service definition. This Supervisor Service version must be in the State#ACTIVATED state. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version.
	Version string `json:"version"`
	// A generic key-value map for additional configuration parameters required during service creation. As an example, a third party operator might reference a private registry using parameters such as \"registryName\" for the registry name, \"registryUsername\" and \"registryPassword\" for the registry credentials. If unset, no additional configuration parameters will be applied when installing a Supervisor Service in the vSphere Supervisor cluster.
	ServiceConfig *map[string]string `json:"service_config,omitempty"`
}

// NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec instantiates a new VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec(supervisorService string, version string) *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec {
	this := VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec{}
	this.SupervisorService = supervisorService
	this.Version = version
	return &this
}

// NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpecWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpecWithDefaults() *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec {
	this := VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec{}
	return &this
}

// GetSupervisorService returns the SupervisorService field value
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) GetSupervisorService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupervisorService
}

// GetSupervisorServiceOk returns a tuple with the SupervisorService field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) GetSupervisorServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupervisorService, true
}

// SetSupervisorService sets field value
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) SetSupervisorService(v string) {
	o.SupervisorService = v
}

// GetVersion returns the Version field value
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) SetVersion(v string) {
	o.Version = v
}

// GetServiceConfig returns the ServiceConfig field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) GetServiceConfig() map[string]string {
	if o == nil || o.ServiceConfig == nil {
		var ret map[string]string
		return ret
	}
	return *o.ServiceConfig
}

// GetServiceConfigOk returns a tuple with the ServiceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) GetServiceConfigOk() (*map[string]string, bool) {
	if o == nil || o.ServiceConfig == nil {
		return nil, false
	}
	return o.ServiceConfig, true
}

// HasServiceConfig returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) HasServiceConfig() bool {
	if o != nil && o.ServiceConfig != nil {
		return true
	}

	return false
}

// SetServiceConfig gets a reference to the given map[string]string and assigns it to the ServiceConfig field.
func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) SetServiceConfig(v map[string]string) {
	o.ServiceConfig = &v
}

func (o VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supervisor_service"] = o.SupervisorService
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.ServiceConfig != nil {
		toSerialize["service_config"] = o.ServiceConfig
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec struct {
	value *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec
	isSet bool
}

func (v NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) Get() *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec {
	return v.value
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) Set(val *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec(val *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) *NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec {
	return &NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


