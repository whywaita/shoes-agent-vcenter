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

// VcenterVMRegisterSpec struct for VcenterVMRegisterSpec
type VcenterVMRegisterSpec struct {
	// Identifier of the datastore on which the virtual machine's configuration state is stored. If unset, VM.RegisterSpec.path must also be unset and VM.RegisterSpec.datastore-path must be set. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datastore. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datastore.
	Datastore *string `json:"datastore,omitempty"`
	// Path to the virtual machine's configuration file on the datastore corresponding to {@link #datastore). If unset, VM.RegisterSpec.datastore must also be unset and VM.RegisterSpec.datastore-path must be set.
	Path *string `json:"path,omitempty"`
	// Datastore path for the virtual machine's configuration file in the format \"[datastore name] path\". For example \"[storage1] Test-VM/Test-VM.vmx\". If unset, both VM.RegisterSpec.datastore and VM.RegisterSpec.path must be set.
	DatastorePath *string `json:"datastore_path,omitempty"`
	// Virtual machine name. If unset, the display name from the virtual machine's configuration file will be used.
	Name *string `json:"name,omitempty"`
	Placement *VcenterVMRegisterPlacementSpec `json:"placement,omitempty"`
}

// NewVcenterVMRegisterSpec instantiates a new VcenterVMRegisterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMRegisterSpec() *VcenterVMRegisterSpec {
	this := VcenterVMRegisterSpec{}
	return &this
}

// NewVcenterVMRegisterSpecWithDefaults instantiates a new VcenterVMRegisterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMRegisterSpecWithDefaults() *VcenterVMRegisterSpec {
	this := VcenterVMRegisterSpec{}
	return &this
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *VcenterVMRegisterSpec) GetDatastore() string {
	if o == nil || o.Datastore == nil {
		var ret string
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMRegisterSpec) GetDatastoreOk() (*string, bool) {
	if o == nil || o.Datastore == nil {
		return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *VcenterVMRegisterSpec) HasDatastore() bool {
	if o != nil && o.Datastore != nil {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given string and assigns it to the Datastore field.
func (o *VcenterVMRegisterSpec) SetDatastore(v string) {
	o.Datastore = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *VcenterVMRegisterSpec) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMRegisterSpec) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *VcenterVMRegisterSpec) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *VcenterVMRegisterSpec) SetPath(v string) {
	o.Path = &v
}

// GetDatastorePath returns the DatastorePath field value if set, zero value otherwise.
func (o *VcenterVMRegisterSpec) GetDatastorePath() string {
	if o == nil || o.DatastorePath == nil {
		var ret string
		return ret
	}
	return *o.DatastorePath
}

// GetDatastorePathOk returns a tuple with the DatastorePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMRegisterSpec) GetDatastorePathOk() (*string, bool) {
	if o == nil || o.DatastorePath == nil {
		return nil, false
	}
	return o.DatastorePath, true
}

// HasDatastorePath returns a boolean if a field has been set.
func (o *VcenterVMRegisterSpec) HasDatastorePath() bool {
	if o != nil && o.DatastorePath != nil {
		return true
	}

	return false
}

// SetDatastorePath gets a reference to the given string and assigns it to the DatastorePath field.
func (o *VcenterVMRegisterSpec) SetDatastorePath(v string) {
	o.DatastorePath = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VcenterVMRegisterSpec) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMRegisterSpec) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VcenterVMRegisterSpec) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VcenterVMRegisterSpec) SetName(v string) {
	o.Name = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise.
func (o *VcenterVMRegisterSpec) GetPlacement() VcenterVMRegisterPlacementSpec {
	if o == nil || o.Placement == nil {
		var ret VcenterVMRegisterPlacementSpec
		return ret
	}
	return *o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMRegisterSpec) GetPlacementOk() (*VcenterVMRegisterPlacementSpec, bool) {
	if o == nil || o.Placement == nil {
		return nil, false
	}
	return o.Placement, true
}

// HasPlacement returns a boolean if a field has been set.
func (o *VcenterVMRegisterSpec) HasPlacement() bool {
	if o != nil && o.Placement != nil {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given VcenterVMRegisterPlacementSpec and assigns it to the Placement field.
func (o *VcenterVMRegisterSpec) SetPlacement(v VcenterVMRegisterPlacementSpec) {
	o.Placement = &v
}

func (o VcenterVMRegisterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Datastore != nil {
		toSerialize["datastore"] = o.Datastore
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.DatastorePath != nil {
		toSerialize["datastore_path"] = o.DatastorePath
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Placement != nil {
		toSerialize["placement"] = o.Placement
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMRegisterSpec struct {
	value *VcenterVMRegisterSpec
	isSet bool
}

func (v NullableVcenterVMRegisterSpec) Get() *VcenterVMRegisterSpec {
	return v.value
}

func (v *NullableVcenterVMRegisterSpec) Set(val *VcenterVMRegisterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMRegisterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMRegisterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMRegisterSpec(val *VcenterVMRegisterSpec) *NullableVcenterVMRegisterSpec {
	return &NullableVcenterVMRegisterSpec{value: val, isSet: true}
}

func (v NullableVcenterVMRegisterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMRegisterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


