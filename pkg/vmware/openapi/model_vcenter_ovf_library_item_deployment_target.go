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

// VcenterOvfLibraryItemDeploymentTarget struct for VcenterOvfLibraryItemDeploymentTarget
type VcenterOvfLibraryItemDeploymentTarget struct {
	// Identifier of the resource pool to which the virtual machine or virtual appliance should be attached.
	ResourcePoolId string `json:"resource_pool_id"`
	// Identifier of the target host on which the virtual machine or virtual appliance will run. The target host must be a member of the cluster that contains the resource pool identified by {@link #resourcePoolId}.
	HostId *string `json:"host_id,omitempty"`
	// Identifier of the vCenter folder that should contain the virtual machine or virtual appliance. The folder must be virtual machine folder.
	FolderId *string `json:"folder_id,omitempty"`
}

// NewVcenterOvfLibraryItemDeploymentTarget instantiates a new VcenterOvfLibraryItemDeploymentTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterOvfLibraryItemDeploymentTarget(resourcePoolId string) *VcenterOvfLibraryItemDeploymentTarget {
	this := VcenterOvfLibraryItemDeploymentTarget{}
	this.ResourcePoolId = resourcePoolId
	return &this
}

// NewVcenterOvfLibraryItemDeploymentTargetWithDefaults instantiates a new VcenterOvfLibraryItemDeploymentTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterOvfLibraryItemDeploymentTargetWithDefaults() *VcenterOvfLibraryItemDeploymentTarget {
	this := VcenterOvfLibraryItemDeploymentTarget{}
	return &this
}

// GetResourcePoolId returns the ResourcePoolId field value
func (o *VcenterOvfLibraryItemDeploymentTarget) GetResourcePoolId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourcePoolId
}

// GetResourcePoolIdOk returns a tuple with the ResourcePoolId field value
// and a boolean to check if the value has been set.
func (o *VcenterOvfLibraryItemDeploymentTarget) GetResourcePoolIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourcePoolId, true
}

// SetResourcePoolId sets field value
func (o *VcenterOvfLibraryItemDeploymentTarget) SetResourcePoolId(v string) {
	o.ResourcePoolId = v
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *VcenterOvfLibraryItemDeploymentTarget) GetHostId() string {
	if o == nil || o.HostId == nil {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterOvfLibraryItemDeploymentTarget) GetHostIdOk() (*string, bool) {
	if o == nil || o.HostId == nil {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *VcenterOvfLibraryItemDeploymentTarget) HasHostId() bool {
	if o != nil && o.HostId != nil {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *VcenterOvfLibraryItemDeploymentTarget) SetHostId(v string) {
	o.HostId = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *VcenterOvfLibraryItemDeploymentTarget) GetFolderId() string {
	if o == nil || o.FolderId == nil {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterOvfLibraryItemDeploymentTarget) GetFolderIdOk() (*string, bool) {
	if o == nil || o.FolderId == nil {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *VcenterOvfLibraryItemDeploymentTarget) HasFolderId() bool {
	if o != nil && o.FolderId != nil {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *VcenterOvfLibraryItemDeploymentTarget) SetFolderId(v string) {
	o.FolderId = &v
}

func (o VcenterOvfLibraryItemDeploymentTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resource_pool_id"] = o.ResourcePoolId
	}
	if o.HostId != nil {
		toSerialize["host_id"] = o.HostId
	}
	if o.FolderId != nil {
		toSerialize["folder_id"] = o.FolderId
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterOvfLibraryItemDeploymentTarget struct {
	value *VcenterOvfLibraryItemDeploymentTarget
	isSet bool
}

func (v NullableVcenterOvfLibraryItemDeploymentTarget) Get() *VcenterOvfLibraryItemDeploymentTarget {
	return v.value
}

func (v *NullableVcenterOvfLibraryItemDeploymentTarget) Set(val *VcenterOvfLibraryItemDeploymentTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterOvfLibraryItemDeploymentTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterOvfLibraryItemDeploymentTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterOvfLibraryItemDeploymentTarget(val *VcenterOvfLibraryItemDeploymentTarget) *NullableVcenterOvfLibraryItemDeploymentTarget {
	return &NullableVcenterOvfLibraryItemDeploymentTarget{value: val, isSet: true}
}

func (v NullableVcenterOvfLibraryItemDeploymentTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterOvfLibraryItemDeploymentTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


