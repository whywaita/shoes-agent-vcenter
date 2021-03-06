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

// VcenterVMDiskRelocateSpec struct for VcenterVMDiskRelocateSpec
type VcenterVMDiskRelocateSpec struct {
	// Destination datastore to relocate disk. This field is currently required. In the future, if this field is unset, disk will use the datastore specified in VM.RelocatePlacementSpec.datastore field of VM.RelocateSpec.placement. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datastore. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datastore.
	Datastore *string `json:"datastore,omitempty"`
}

// NewVcenterVMDiskRelocateSpec instantiates a new VcenterVMDiskRelocateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMDiskRelocateSpec() *VcenterVMDiskRelocateSpec {
	this := VcenterVMDiskRelocateSpec{}
	return &this
}

// NewVcenterVMDiskRelocateSpecWithDefaults instantiates a new VcenterVMDiskRelocateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMDiskRelocateSpecWithDefaults() *VcenterVMDiskRelocateSpec {
	this := VcenterVMDiskRelocateSpec{}
	return &this
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *VcenterVMDiskRelocateSpec) GetDatastore() string {
	if o == nil || o.Datastore == nil {
		var ret string
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMDiskRelocateSpec) GetDatastoreOk() (*string, bool) {
	if o == nil || o.Datastore == nil {
		return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *VcenterVMDiskRelocateSpec) HasDatastore() bool {
	if o != nil && o.Datastore != nil {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given string and assigns it to the Datastore field.
func (o *VcenterVMDiskRelocateSpec) SetDatastore(v string) {
	o.Datastore = &v
}

func (o VcenterVMDiskRelocateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Datastore != nil {
		toSerialize["datastore"] = o.Datastore
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMDiskRelocateSpec struct {
	value *VcenterVMDiskRelocateSpec
	isSet bool
}

func (v NullableVcenterVMDiskRelocateSpec) Get() *VcenterVMDiskRelocateSpec {
	return v.value
}

func (v *NullableVcenterVMDiskRelocateSpec) Set(val *VcenterVMDiskRelocateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMDiskRelocateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMDiskRelocateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMDiskRelocateSpec(val *VcenterVMDiskRelocateSpec) *NullableVcenterVMDiskRelocateSpec {
	return &NullableVcenterVMDiskRelocateSpec{value: val, isSet: true}
}

func (v NullableVcenterVMDiskRelocateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMDiskRelocateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


