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

// VcenterNamespacesNamespaceTemplatesUpdateSpec struct for VcenterNamespacesNamespaceTemplatesUpdateSpec
type VcenterNamespacesNamespaceTemplatesUpdateSpec struct {
	// Resource quota on the namespace. Refer to vcenter.namespace_management.NamespaceResourceOptions.Info#createResourceQuotaType and use vcenter.namespace_management.NamespaceResourceOptions#get for retrieving the type for the value for this field. For an example of this, see ResourceQuotaOptionsV1. If unset, no resource limits will be set on the namespace.
	ResourceSpec *map[string]interface{} `json:"resource_spec,omitempty"`
	// Storage that this template defines and will be associated with a namespace after namespace realization.
	StorageSpecs *[]VcenterNamespacesInstancesStorageSpec `json:"storage_specs,omitempty"`
	// vSphere Namespaces network objects to be associated with the namespace. The values of this list need to reference names of pre-existing Networks.Info structures. The field must be left unset if the cluster hosting the namespace uses NSXT_CONTAINER_PLUGIN as the network provider, since the network(s) for this namespace will be managed by NSX-T Container Plugin. If field is unset when the cluster hosting the namespace uses VSPHERE_NETWORK as its network provider, the namespace will automatically be associated with the cluster's Supervisor Primary Workload Network. The field currently accepts at most only 1 vSphere Namespaces network object reference. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ClusterComputeResource.
	Networks *[]string `json:"networks,omitempty"`
	// Permissions associated with namespace template. If unset, only users with the Administrator role can use this template; for example, this template is applied to the namespace created by self-service-users with the Administrator role.
	Permissions *[]VcenterNamespacesNamespaceTemplatesSubject `json:"permissions,omitempty"`
}

// NewVcenterNamespacesNamespaceTemplatesUpdateSpec instantiates a new VcenterNamespacesNamespaceTemplatesUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespacesNamespaceTemplatesUpdateSpec() *VcenterNamespacesNamespaceTemplatesUpdateSpec {
	this := VcenterNamespacesNamespaceTemplatesUpdateSpec{}
	return &this
}

// NewVcenterNamespacesNamespaceTemplatesUpdateSpecWithDefaults instantiates a new VcenterNamespacesNamespaceTemplatesUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespacesNamespaceTemplatesUpdateSpecWithDefaults() *VcenterNamespacesNamespaceTemplatesUpdateSpec {
	this := VcenterNamespacesNamespaceTemplatesUpdateSpec{}
	return &this
}

// GetResourceSpec returns the ResourceSpec field value if set, zero value otherwise.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetResourceSpec() map[string]interface{} {
	if o == nil || o.ResourceSpec == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ResourceSpec
}

// GetResourceSpecOk returns a tuple with the ResourceSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetResourceSpecOk() (*map[string]interface{}, bool) {
	if o == nil || o.ResourceSpec == nil {
		return nil, false
	}
	return o.ResourceSpec, true
}

// HasResourceSpec returns a boolean if a field has been set.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) HasResourceSpec() bool {
	if o != nil && o.ResourceSpec != nil {
		return true
	}

	return false
}

// SetResourceSpec gets a reference to the given map[string]interface{} and assigns it to the ResourceSpec field.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) SetResourceSpec(v map[string]interface{}) {
	o.ResourceSpec = &v
}

// GetStorageSpecs returns the StorageSpecs field value if set, zero value otherwise.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetStorageSpecs() []VcenterNamespacesInstancesStorageSpec {
	if o == nil || o.StorageSpecs == nil {
		var ret []VcenterNamespacesInstancesStorageSpec
		return ret
	}
	return *o.StorageSpecs
}

// GetStorageSpecsOk returns a tuple with the StorageSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetStorageSpecsOk() (*[]VcenterNamespacesInstancesStorageSpec, bool) {
	if o == nil || o.StorageSpecs == nil {
		return nil, false
	}
	return o.StorageSpecs, true
}

// HasStorageSpecs returns a boolean if a field has been set.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) HasStorageSpecs() bool {
	if o != nil && o.StorageSpecs != nil {
		return true
	}

	return false
}

// SetStorageSpecs gets a reference to the given []VcenterNamespacesInstancesStorageSpec and assigns it to the StorageSpecs field.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) SetStorageSpecs(v []VcenterNamespacesInstancesStorageSpec) {
	o.StorageSpecs = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetNetworks() []string {
	if o == nil || o.Networks == nil {
		var ret []string
		return ret
	}
	return *o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetNetworksOk() (*[]string, bool) {
	if o == nil || o.Networks == nil {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) HasNetworks() bool {
	if o != nil && o.Networks != nil {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []string and assigns it to the Networks field.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) SetNetworks(v []string) {
	o.Networks = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetPermissions() []VcenterNamespacesNamespaceTemplatesSubject {
	if o == nil || o.Permissions == nil {
		var ret []VcenterNamespacesNamespaceTemplatesSubject
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetPermissionsOk() (*[]VcenterNamespacesNamespaceTemplatesSubject, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []VcenterNamespacesNamespaceTemplatesSubject and assigns it to the Permissions field.
func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) SetPermissions(v []VcenterNamespacesNamespaceTemplatesSubject) {
	o.Permissions = &v
}

func (o VcenterNamespacesNamespaceTemplatesUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceSpec != nil {
		toSerialize["resource_spec"] = o.ResourceSpec
	}
	if o.StorageSpecs != nil {
		toSerialize["storage_specs"] = o.StorageSpecs
	}
	if o.Networks != nil {
		toSerialize["networks"] = o.Networks
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespacesNamespaceTemplatesUpdateSpec struct {
	value *VcenterNamespacesNamespaceTemplatesUpdateSpec
	isSet bool
}

func (v NullableVcenterNamespacesNamespaceTemplatesUpdateSpec) Get() *VcenterNamespacesNamespaceTemplatesUpdateSpec {
	return v.value
}

func (v *NullableVcenterNamespacesNamespaceTemplatesUpdateSpec) Set(val *VcenterNamespacesNamespaceTemplatesUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespacesNamespaceTemplatesUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespacesNamespaceTemplatesUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespacesNamespaceTemplatesUpdateSpec(val *VcenterNamespacesNamespaceTemplatesUpdateSpec) *NullableVcenterNamespacesNamespaceTemplatesUpdateSpec {
	return &NullableVcenterNamespacesNamespaceTemplatesUpdateSpec{value: val, isSet: true}
}

func (v NullableVcenterNamespacesNamespaceTemplatesUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespacesNamespaceTemplatesUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


