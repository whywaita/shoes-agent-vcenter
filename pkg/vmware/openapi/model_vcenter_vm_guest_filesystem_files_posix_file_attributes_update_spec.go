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

// VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec struct for VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec
type VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec struct {
	// The owner ID. If unset the value will not be changed.
	OwnerId *int64 `json:"owner_id,omitempty"`
	// The group ID. If unset the value will not be changed.
	GroupId *int64 `json:"group_id,omitempty"`
	// The file permissions in chmod(2) format. This field is interpreted as octal. If unset the value will not be changed.
	Permissions *string `json:"permissions,omitempty"`
}

// NewVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec instantiates a new VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec() *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec {
	this := VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec{}
	return &this
}

// NewVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpecWithDefaults instantiates a new VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpecWithDefaults() *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec {
	this := VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec{}
	return &this
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) GetOwnerId() int64 {
	if o == nil || o.OwnerId == nil {
		var ret int64
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) GetOwnerIdOk() (*int64, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given int64 and assigns it to the OwnerId field.
func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) SetOwnerId(v int64) {
	o.OwnerId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) GetGroupId() int64 {
	if o == nil || o.GroupId == nil {
		var ret int64
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) GetGroupIdOk() (*int64, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int64 and assigns it to the GroupId field.
func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) SetGroupId(v int64) {
	o.GroupId = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) GetPermissions() string {
	if o == nil || o.Permissions == nil {
		var ret string
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) GetPermissionsOk() (*string, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given string and assigns it to the Permissions field.
func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) SetPermissions(v string) {
	o.Permissions = &v
}

func (o VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OwnerId != nil {
		toSerialize["owner_id"] = o.OwnerId
	}
	if o.GroupId != nil {
		toSerialize["group_id"] = o.GroupId
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec struct {
	value *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec
	isSet bool
}

func (v NullableVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) Get() *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec {
	return v.value
}

func (v *NullableVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) Set(val *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec(val *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) *NullableVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec {
	return &NullableVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec{value: val, isSet: true}
}

func (v NullableVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


