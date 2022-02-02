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

// VcenterVmHardwareAdapterScsiCreateSpec struct for VcenterVmHardwareAdapterScsiCreateSpec
type VcenterVmHardwareAdapterScsiCreateSpec struct {
	Type *VcenterVmHardwareAdapterScsiType `json:"type,omitempty"`
	// SCSI bus number. If unset, the server will choose an available bus number; if none is available, the request will fail.
	Bus *int64 `json:"bus,omitempty"`
	// Address of the SCSI adapter on the PCI bus. If the PCI address is invalid, the server will change it when the VM is started or as the device is hot added. If unset, the server will choose an available address when the virtual machine is powered on.
	PciSlotNumber *int64 `json:"pci_slot_number,omitempty"`
	Sharing *VcenterVmHardwareAdapterScsiSharing `json:"sharing,omitempty"`
}

// NewVcenterVmHardwareAdapterScsiCreateSpec instantiates a new VcenterVmHardwareAdapterScsiCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareAdapterScsiCreateSpec() *VcenterVmHardwareAdapterScsiCreateSpec {
	this := VcenterVmHardwareAdapterScsiCreateSpec{}
	return &this
}

// NewVcenterVmHardwareAdapterScsiCreateSpecWithDefaults instantiates a new VcenterVmHardwareAdapterScsiCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareAdapterScsiCreateSpecWithDefaults() *VcenterVmHardwareAdapterScsiCreateSpec {
	this := VcenterVmHardwareAdapterScsiCreateSpec{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetType() VcenterVmHardwareAdapterScsiType {
	if o == nil || o.Type == nil {
		var ret VcenterVmHardwareAdapterScsiType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetTypeOk() (*VcenterVmHardwareAdapterScsiType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given VcenterVmHardwareAdapterScsiType and assigns it to the Type field.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) SetType(v VcenterVmHardwareAdapterScsiType) {
	o.Type = &v
}

// GetBus returns the Bus field value if set, zero value otherwise.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetBus() int64 {
	if o == nil || o.Bus == nil {
		var ret int64
		return ret
	}
	return *o.Bus
}

// GetBusOk returns a tuple with the Bus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetBusOk() (*int64, bool) {
	if o == nil || o.Bus == nil {
		return nil, false
	}
	return o.Bus, true
}

// HasBus returns a boolean if a field has been set.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) HasBus() bool {
	if o != nil && o.Bus != nil {
		return true
	}

	return false
}

// SetBus gets a reference to the given int64 and assigns it to the Bus field.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) SetBus(v int64) {
	o.Bus = &v
}

// GetPciSlotNumber returns the PciSlotNumber field value if set, zero value otherwise.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetPciSlotNumber() int64 {
	if o == nil || o.PciSlotNumber == nil {
		var ret int64
		return ret
	}
	return *o.PciSlotNumber
}

// GetPciSlotNumberOk returns a tuple with the PciSlotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetPciSlotNumberOk() (*int64, bool) {
	if o == nil || o.PciSlotNumber == nil {
		return nil, false
	}
	return o.PciSlotNumber, true
}

// HasPciSlotNumber returns a boolean if a field has been set.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) HasPciSlotNumber() bool {
	if o != nil && o.PciSlotNumber != nil {
		return true
	}

	return false
}

// SetPciSlotNumber gets a reference to the given int64 and assigns it to the PciSlotNumber field.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) SetPciSlotNumber(v int64) {
	o.PciSlotNumber = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetSharing() VcenterVmHardwareAdapterScsiSharing {
	if o == nil || o.Sharing == nil {
		var ret VcenterVmHardwareAdapterScsiSharing
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetSharingOk() (*VcenterVmHardwareAdapterScsiSharing, bool) {
	if o == nil || o.Sharing == nil {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) HasSharing() bool {
	if o != nil && o.Sharing != nil {
		return true
	}

	return false
}

// SetSharing gets a reference to the given VcenterVmHardwareAdapterScsiSharing and assigns it to the Sharing field.
func (o *VcenterVmHardwareAdapterScsiCreateSpec) SetSharing(v VcenterVmHardwareAdapterScsiSharing) {
	o.Sharing = &v
}

func (o VcenterVmHardwareAdapterScsiCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Bus != nil {
		toSerialize["bus"] = o.Bus
	}
	if o.PciSlotNumber != nil {
		toSerialize["pci_slot_number"] = o.PciSlotNumber
	}
	if o.Sharing != nil {
		toSerialize["sharing"] = o.Sharing
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareAdapterScsiCreateSpec struct {
	value *VcenterVmHardwareAdapterScsiCreateSpec
	isSet bool
}

func (v NullableVcenterVmHardwareAdapterScsiCreateSpec) Get() *VcenterVmHardwareAdapterScsiCreateSpec {
	return v.value
}

func (v *NullableVcenterVmHardwareAdapterScsiCreateSpec) Set(val *VcenterVmHardwareAdapterScsiCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareAdapterScsiCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareAdapterScsiCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareAdapterScsiCreateSpec(val *VcenterVmHardwareAdapterScsiCreateSpec) *NullableVcenterVmHardwareAdapterScsiCreateSpec {
	return &NullableVcenterVmHardwareAdapterScsiCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareAdapterScsiCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareAdapterScsiCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

