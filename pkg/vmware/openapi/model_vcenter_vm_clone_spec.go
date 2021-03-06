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

// VcenterVMCloneSpec struct for VcenterVMCloneSpec
type VcenterVMCloneSpec struct {
	// Virtual machine to clone from. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: VirtualMachine. When operations return a value of this structure as a result, the field will be an identifier for the resource type: VirtualMachine.
	Source string `json:"source"`
	// Virtual machine name.
	Name string `json:"name"`
	Placement *VcenterVMClonePlacementSpec `json:"placement,omitempty"`
	// Set of Disks to Remove. If unset, all disks will be copied. If the same identifier is in VM.CloneSpec.disks-to-update InvalidArgument fault will be returned. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.vm.hardware.Disk.
	DisksToRemove *[]string `json:"disks_to_remove,omitempty"`
	// Map of Disks to Update. If unset, all disks will copied to the datastore specified in the VM.ClonePlacementSpec.datastore field of VM.CloneSpec.placement. If the same identifier is in VM.CloneSpec.disks-to-remove InvalidArgument fault will be thrown. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Disk.
	DisksToUpdate *[]VcenterVMCloneSpecDisksToUpdate `json:"disks_to_update,omitempty"`
	// Attempt to perform a VM.CloneSpec.power-on after clone. If unset, the virtual machine will not be powered on.
	PowerOn *bool `json:"power_on,omitempty"`
	GuestCustomizationSpec *VcenterVMGuestCustomizationSpec `json:"guest_customization_spec,omitempty"`
}

// NewVcenterVMCloneSpec instantiates a new VcenterVMCloneSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMCloneSpec(source string, name string) *VcenterVMCloneSpec {
	this := VcenterVMCloneSpec{}
	this.Source = source
	this.Name = name
	return &this
}

// NewVcenterVMCloneSpecWithDefaults instantiates a new VcenterVMCloneSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMCloneSpecWithDefaults() *VcenterVMCloneSpec {
	this := VcenterVMCloneSpec{}
	return &this
}

// GetSource returns the Source field value
func (o *VcenterVMCloneSpec) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *VcenterVMCloneSpec) GetSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *VcenterVMCloneSpec) SetSource(v string) {
	o.Source = v
}

// GetName returns the Name field value
func (o *VcenterVMCloneSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VcenterVMCloneSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VcenterVMCloneSpec) SetName(v string) {
	o.Name = v
}

// GetPlacement returns the Placement field value if set, zero value otherwise.
func (o *VcenterVMCloneSpec) GetPlacement() VcenterVMClonePlacementSpec {
	if o == nil || o.Placement == nil {
		var ret VcenterVMClonePlacementSpec
		return ret
	}
	return *o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMCloneSpec) GetPlacementOk() (*VcenterVMClonePlacementSpec, bool) {
	if o == nil || o.Placement == nil {
		return nil, false
	}
	return o.Placement, true
}

// HasPlacement returns a boolean if a field has been set.
func (o *VcenterVMCloneSpec) HasPlacement() bool {
	if o != nil && o.Placement != nil {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given VcenterVMClonePlacementSpec and assigns it to the Placement field.
func (o *VcenterVMCloneSpec) SetPlacement(v VcenterVMClonePlacementSpec) {
	o.Placement = &v
}

// GetDisksToRemove returns the DisksToRemove field value if set, zero value otherwise.
func (o *VcenterVMCloneSpec) GetDisksToRemove() []string {
	if o == nil || o.DisksToRemove == nil {
		var ret []string
		return ret
	}
	return *o.DisksToRemove
}

// GetDisksToRemoveOk returns a tuple with the DisksToRemove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMCloneSpec) GetDisksToRemoveOk() (*[]string, bool) {
	if o == nil || o.DisksToRemove == nil {
		return nil, false
	}
	return o.DisksToRemove, true
}

// HasDisksToRemove returns a boolean if a field has been set.
func (o *VcenterVMCloneSpec) HasDisksToRemove() bool {
	if o != nil && o.DisksToRemove != nil {
		return true
	}

	return false
}

// SetDisksToRemove gets a reference to the given []string and assigns it to the DisksToRemove field.
func (o *VcenterVMCloneSpec) SetDisksToRemove(v []string) {
	o.DisksToRemove = &v
}

// GetDisksToUpdate returns the DisksToUpdate field value if set, zero value otherwise.
func (o *VcenterVMCloneSpec) GetDisksToUpdate() []VcenterVMCloneSpecDisksToUpdate {
	if o == nil || o.DisksToUpdate == nil {
		var ret []VcenterVMCloneSpecDisksToUpdate
		return ret
	}
	return *o.DisksToUpdate
}

// GetDisksToUpdateOk returns a tuple with the DisksToUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMCloneSpec) GetDisksToUpdateOk() (*[]VcenterVMCloneSpecDisksToUpdate, bool) {
	if o == nil || o.DisksToUpdate == nil {
		return nil, false
	}
	return o.DisksToUpdate, true
}

// HasDisksToUpdate returns a boolean if a field has been set.
func (o *VcenterVMCloneSpec) HasDisksToUpdate() bool {
	if o != nil && o.DisksToUpdate != nil {
		return true
	}

	return false
}

// SetDisksToUpdate gets a reference to the given []VcenterVMCloneSpecDisksToUpdate and assigns it to the DisksToUpdate field.
func (o *VcenterVMCloneSpec) SetDisksToUpdate(v []VcenterVMCloneSpecDisksToUpdate) {
	o.DisksToUpdate = &v
}

// GetPowerOn returns the PowerOn field value if set, zero value otherwise.
func (o *VcenterVMCloneSpec) GetPowerOn() bool {
	if o == nil || o.PowerOn == nil {
		var ret bool
		return ret
	}
	return *o.PowerOn
}

// GetPowerOnOk returns a tuple with the PowerOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMCloneSpec) GetPowerOnOk() (*bool, bool) {
	if o == nil || o.PowerOn == nil {
		return nil, false
	}
	return o.PowerOn, true
}

// HasPowerOn returns a boolean if a field has been set.
func (o *VcenterVMCloneSpec) HasPowerOn() bool {
	if o != nil && o.PowerOn != nil {
		return true
	}

	return false
}

// SetPowerOn gets a reference to the given bool and assigns it to the PowerOn field.
func (o *VcenterVMCloneSpec) SetPowerOn(v bool) {
	o.PowerOn = &v
}

// GetGuestCustomizationSpec returns the GuestCustomizationSpec field value if set, zero value otherwise.
func (o *VcenterVMCloneSpec) GetGuestCustomizationSpec() VcenterVMGuestCustomizationSpec {
	if o == nil || o.GuestCustomizationSpec == nil {
		var ret VcenterVMGuestCustomizationSpec
		return ret
	}
	return *o.GuestCustomizationSpec
}

// GetGuestCustomizationSpecOk returns a tuple with the GuestCustomizationSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMCloneSpec) GetGuestCustomizationSpecOk() (*VcenterVMGuestCustomizationSpec, bool) {
	if o == nil || o.GuestCustomizationSpec == nil {
		return nil, false
	}
	return o.GuestCustomizationSpec, true
}

// HasGuestCustomizationSpec returns a boolean if a field has been set.
func (o *VcenterVMCloneSpec) HasGuestCustomizationSpec() bool {
	if o != nil && o.GuestCustomizationSpec != nil {
		return true
	}

	return false
}

// SetGuestCustomizationSpec gets a reference to the given VcenterVMGuestCustomizationSpec and assigns it to the GuestCustomizationSpec field.
func (o *VcenterVMCloneSpec) SetGuestCustomizationSpec(v VcenterVMGuestCustomizationSpec) {
	o.GuestCustomizationSpec = &v
}

func (o VcenterVMCloneSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Placement != nil {
		toSerialize["placement"] = o.Placement
	}
	if o.DisksToRemove != nil {
		toSerialize["disks_to_remove"] = o.DisksToRemove
	}
	if o.DisksToUpdate != nil {
		toSerialize["disks_to_update"] = o.DisksToUpdate
	}
	if o.PowerOn != nil {
		toSerialize["power_on"] = o.PowerOn
	}
	if o.GuestCustomizationSpec != nil {
		toSerialize["guest_customization_spec"] = o.GuestCustomizationSpec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMCloneSpec struct {
	value *VcenterVMCloneSpec
	isSet bool
}

func (v NullableVcenterVMCloneSpec) Get() *VcenterVMCloneSpec {
	return v.value
}

func (v *NullableVcenterVMCloneSpec) Set(val *VcenterVMCloneSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMCloneSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMCloneSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMCloneSpec(val *VcenterVMCloneSpec) *NullableVcenterVMCloneSpec {
	return &NullableVcenterVMCloneSpec{value: val, isSet: true}
}

func (v NullableVcenterVMCloneSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMCloneSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


