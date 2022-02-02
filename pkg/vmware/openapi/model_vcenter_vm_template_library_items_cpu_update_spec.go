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

// VcenterVmTemplateLibraryItemsCpuUpdateSpec struct for VcenterVmTemplateLibraryItemsCpuUpdateSpec
type VcenterVmTemplateLibraryItemsCpuUpdateSpec struct {
	// Number of virtual processors in the deployed virtual machine.
	NumCpus *int64 `json:"num_cpus,omitempty"`
	// Number of cores among which to distribute CPUs in the deployed virtual machine.
	NumCoresPerSocket *int64 `json:"num_cores_per_socket,omitempty"`
}

// NewVcenterVmTemplateLibraryItemsCpuUpdateSpec instantiates a new VcenterVmTemplateLibraryItemsCpuUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmTemplateLibraryItemsCpuUpdateSpec() *VcenterVmTemplateLibraryItemsCpuUpdateSpec {
	this := VcenterVmTemplateLibraryItemsCpuUpdateSpec{}
	return &this
}

// NewVcenterVmTemplateLibraryItemsCpuUpdateSpecWithDefaults instantiates a new VcenterVmTemplateLibraryItemsCpuUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmTemplateLibraryItemsCpuUpdateSpecWithDefaults() *VcenterVmTemplateLibraryItemsCpuUpdateSpec {
	this := VcenterVmTemplateLibraryItemsCpuUpdateSpec{}
	return &this
}

// GetNumCpus returns the NumCpus field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) GetNumCpus() int64 {
	if o == nil || o.NumCpus == nil {
		var ret int64
		return ret
	}
	return *o.NumCpus
}

// GetNumCpusOk returns a tuple with the NumCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) GetNumCpusOk() (*int64, bool) {
	if o == nil || o.NumCpus == nil {
		return nil, false
	}
	return o.NumCpus, true
}

// HasNumCpus returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) HasNumCpus() bool {
	if o != nil && o.NumCpus != nil {
		return true
	}

	return false
}

// SetNumCpus gets a reference to the given int64 and assigns it to the NumCpus field.
func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) SetNumCpus(v int64) {
	o.NumCpus = &v
}

// GetNumCoresPerSocket returns the NumCoresPerSocket field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) GetNumCoresPerSocket() int64 {
	if o == nil || o.NumCoresPerSocket == nil {
		var ret int64
		return ret
	}
	return *o.NumCoresPerSocket
}

// GetNumCoresPerSocketOk returns a tuple with the NumCoresPerSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) GetNumCoresPerSocketOk() (*int64, bool) {
	if o == nil || o.NumCoresPerSocket == nil {
		return nil, false
	}
	return o.NumCoresPerSocket, true
}

// HasNumCoresPerSocket returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) HasNumCoresPerSocket() bool {
	if o != nil && o.NumCoresPerSocket != nil {
		return true
	}

	return false
}

// SetNumCoresPerSocket gets a reference to the given int64 and assigns it to the NumCoresPerSocket field.
func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) SetNumCoresPerSocket(v int64) {
	o.NumCoresPerSocket = &v
}

func (o VcenterVmTemplateLibraryItemsCpuUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NumCpus != nil {
		toSerialize["num_cpus"] = o.NumCpus
	}
	if o.NumCoresPerSocket != nil {
		toSerialize["num_cores_per_socket"] = o.NumCoresPerSocket
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmTemplateLibraryItemsCpuUpdateSpec struct {
	value *VcenterVmTemplateLibraryItemsCpuUpdateSpec
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsCpuUpdateSpec) Get() *VcenterVmTemplateLibraryItemsCpuUpdateSpec {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsCpuUpdateSpec) Set(val *VcenterVmTemplateLibraryItemsCpuUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsCpuUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsCpuUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsCpuUpdateSpec(val *VcenterVmTemplateLibraryItemsCpuUpdateSpec) *NullableVcenterVmTemplateLibraryItemsCpuUpdateSpec {
	return &NullableVcenterVmTemplateLibraryItemsCpuUpdateSpec{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsCpuUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsCpuUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


