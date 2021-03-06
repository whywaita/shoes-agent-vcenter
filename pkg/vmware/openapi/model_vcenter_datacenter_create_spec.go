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

// VcenterDatacenterCreateSpec struct for VcenterDatacenterCreateSpec
type VcenterDatacenterCreateSpec struct {
	// The name of the datacenter to be created.
	Name string `json:"name"`
	// Datacenter folder in which the new datacenter should be created. This field is currently required. In the future, if this field is unset, the system will attempt to choose a suitable folder for the datacenter; if a folder cannot be chosen, the datacenter creation operation will fail. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder.
	Folder *string `json:"folder,omitempty"`
}

// NewVcenterDatacenterCreateSpec instantiates a new VcenterDatacenterCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDatacenterCreateSpec(name string) *VcenterDatacenterCreateSpec {
	this := VcenterDatacenterCreateSpec{}
	this.Name = name
	return &this
}

// NewVcenterDatacenterCreateSpecWithDefaults instantiates a new VcenterDatacenterCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDatacenterCreateSpecWithDefaults() *VcenterDatacenterCreateSpec {
	this := VcenterDatacenterCreateSpec{}
	return &this
}

// GetName returns the Name field value
func (o *VcenterDatacenterCreateSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VcenterDatacenterCreateSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VcenterDatacenterCreateSpec) SetName(v string) {
	o.Name = v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *VcenterDatacenterCreateSpec) GetFolder() string {
	if o == nil || o.Folder == nil {
		var ret string
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDatacenterCreateSpec) GetFolderOk() (*string, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *VcenterDatacenterCreateSpec) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given string and assigns it to the Folder field.
func (o *VcenterDatacenterCreateSpec) SetFolder(v string) {
	o.Folder = &v
}

func (o VcenterDatacenterCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Folder != nil {
		toSerialize["folder"] = o.Folder
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDatacenterCreateSpec struct {
	value *VcenterDatacenterCreateSpec
	isSet bool
}

func (v NullableVcenterDatacenterCreateSpec) Get() *VcenterDatacenterCreateSpec {
	return v.value
}

func (v *NullableVcenterDatacenterCreateSpec) Set(val *VcenterDatacenterCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDatacenterCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDatacenterCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDatacenterCreateSpec(val *VcenterDatacenterCreateSpec) *NullableVcenterDatacenterCreateSpec {
	return &NullableVcenterDatacenterCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterDatacenterCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDatacenterCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


