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

// VcenterVmHardwareCdromCreateSpec struct for VcenterVmHardwareCdromCreateSpec
type VcenterVmHardwareCdromCreateSpec struct {
	Type *VcenterVmHardwareCdromHostBusAdapterType `json:"type,omitempty"`
	Ide *VcenterVmHardwareIdeAddressSpec `json:"ide,omitempty"`
	Sata *VcenterVmHardwareSataAddressSpec `json:"sata,omitempty"`
	Backing *VcenterVmHardwareCdromBackingSpec `json:"backing,omitempty"`
	// Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. Defaults to false if unset.
	StartConnected *bool `json:"start_connected,omitempty"`
	// Flag indicating whether the guest can connect and disconnect the device. Defaults to false if unset.
	AllowGuestControl *bool `json:"allow_guest_control,omitempty"`
}

// NewVcenterVmHardwareCdromCreateSpec instantiates a new VcenterVmHardwareCdromCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareCdromCreateSpec() *VcenterVmHardwareCdromCreateSpec {
	this := VcenterVmHardwareCdromCreateSpec{}
	return &this
}

// NewVcenterVmHardwareCdromCreateSpecWithDefaults instantiates a new VcenterVmHardwareCdromCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareCdromCreateSpecWithDefaults() *VcenterVmHardwareCdromCreateSpec {
	this := VcenterVmHardwareCdromCreateSpec{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VcenterVmHardwareCdromCreateSpec) GetType() VcenterVmHardwareCdromHostBusAdapterType {
	if o == nil || o.Type == nil {
		var ret VcenterVmHardwareCdromHostBusAdapterType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCdromCreateSpec) GetTypeOk() (*VcenterVmHardwareCdromHostBusAdapterType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VcenterVmHardwareCdromCreateSpec) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given VcenterVmHardwareCdromHostBusAdapterType and assigns it to the Type field.
func (o *VcenterVmHardwareCdromCreateSpec) SetType(v VcenterVmHardwareCdromHostBusAdapterType) {
	o.Type = &v
}

// GetIde returns the Ide field value if set, zero value otherwise.
func (o *VcenterVmHardwareCdromCreateSpec) GetIde() VcenterVmHardwareIdeAddressSpec {
	if o == nil || o.Ide == nil {
		var ret VcenterVmHardwareIdeAddressSpec
		return ret
	}
	return *o.Ide
}

// GetIdeOk returns a tuple with the Ide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCdromCreateSpec) GetIdeOk() (*VcenterVmHardwareIdeAddressSpec, bool) {
	if o == nil || o.Ide == nil {
		return nil, false
	}
	return o.Ide, true
}

// HasIde returns a boolean if a field has been set.
func (o *VcenterVmHardwareCdromCreateSpec) HasIde() bool {
	if o != nil && o.Ide != nil {
		return true
	}

	return false
}

// SetIde gets a reference to the given VcenterVmHardwareIdeAddressSpec and assigns it to the Ide field.
func (o *VcenterVmHardwareCdromCreateSpec) SetIde(v VcenterVmHardwareIdeAddressSpec) {
	o.Ide = &v
}

// GetSata returns the Sata field value if set, zero value otherwise.
func (o *VcenterVmHardwareCdromCreateSpec) GetSata() VcenterVmHardwareSataAddressSpec {
	if o == nil || o.Sata == nil {
		var ret VcenterVmHardwareSataAddressSpec
		return ret
	}
	return *o.Sata
}

// GetSataOk returns a tuple with the Sata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCdromCreateSpec) GetSataOk() (*VcenterVmHardwareSataAddressSpec, bool) {
	if o == nil || o.Sata == nil {
		return nil, false
	}
	return o.Sata, true
}

// HasSata returns a boolean if a field has been set.
func (o *VcenterVmHardwareCdromCreateSpec) HasSata() bool {
	if o != nil && o.Sata != nil {
		return true
	}

	return false
}

// SetSata gets a reference to the given VcenterVmHardwareSataAddressSpec and assigns it to the Sata field.
func (o *VcenterVmHardwareCdromCreateSpec) SetSata(v VcenterVmHardwareSataAddressSpec) {
	o.Sata = &v
}

// GetBacking returns the Backing field value if set, zero value otherwise.
func (o *VcenterVmHardwareCdromCreateSpec) GetBacking() VcenterVmHardwareCdromBackingSpec {
	if o == nil || o.Backing == nil {
		var ret VcenterVmHardwareCdromBackingSpec
		return ret
	}
	return *o.Backing
}

// GetBackingOk returns a tuple with the Backing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCdromCreateSpec) GetBackingOk() (*VcenterVmHardwareCdromBackingSpec, bool) {
	if o == nil || o.Backing == nil {
		return nil, false
	}
	return o.Backing, true
}

// HasBacking returns a boolean if a field has been set.
func (o *VcenterVmHardwareCdromCreateSpec) HasBacking() bool {
	if o != nil && o.Backing != nil {
		return true
	}

	return false
}

// SetBacking gets a reference to the given VcenterVmHardwareCdromBackingSpec and assigns it to the Backing field.
func (o *VcenterVmHardwareCdromCreateSpec) SetBacking(v VcenterVmHardwareCdromBackingSpec) {
	o.Backing = &v
}

// GetStartConnected returns the StartConnected field value if set, zero value otherwise.
func (o *VcenterVmHardwareCdromCreateSpec) GetStartConnected() bool {
	if o == nil || o.StartConnected == nil {
		var ret bool
		return ret
	}
	return *o.StartConnected
}

// GetStartConnectedOk returns a tuple with the StartConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCdromCreateSpec) GetStartConnectedOk() (*bool, bool) {
	if o == nil || o.StartConnected == nil {
		return nil, false
	}
	return o.StartConnected, true
}

// HasStartConnected returns a boolean if a field has been set.
func (o *VcenterVmHardwareCdromCreateSpec) HasStartConnected() bool {
	if o != nil && o.StartConnected != nil {
		return true
	}

	return false
}

// SetStartConnected gets a reference to the given bool and assigns it to the StartConnected field.
func (o *VcenterVmHardwareCdromCreateSpec) SetStartConnected(v bool) {
	o.StartConnected = &v
}

// GetAllowGuestControl returns the AllowGuestControl field value if set, zero value otherwise.
func (o *VcenterVmHardwareCdromCreateSpec) GetAllowGuestControl() bool {
	if o == nil || o.AllowGuestControl == nil {
		var ret bool
		return ret
	}
	return *o.AllowGuestControl
}

// GetAllowGuestControlOk returns a tuple with the AllowGuestControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCdromCreateSpec) GetAllowGuestControlOk() (*bool, bool) {
	if o == nil || o.AllowGuestControl == nil {
		return nil, false
	}
	return o.AllowGuestControl, true
}

// HasAllowGuestControl returns a boolean if a field has been set.
func (o *VcenterVmHardwareCdromCreateSpec) HasAllowGuestControl() bool {
	if o != nil && o.AllowGuestControl != nil {
		return true
	}

	return false
}

// SetAllowGuestControl gets a reference to the given bool and assigns it to the AllowGuestControl field.
func (o *VcenterVmHardwareCdromCreateSpec) SetAllowGuestControl(v bool) {
	o.AllowGuestControl = &v
}

func (o VcenterVmHardwareCdromCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Ide != nil {
		toSerialize["ide"] = o.Ide
	}
	if o.Sata != nil {
		toSerialize["sata"] = o.Sata
	}
	if o.Backing != nil {
		toSerialize["backing"] = o.Backing
	}
	if o.StartConnected != nil {
		toSerialize["start_connected"] = o.StartConnected
	}
	if o.AllowGuestControl != nil {
		toSerialize["allow_guest_control"] = o.AllowGuestControl
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareCdromCreateSpec struct {
	value *VcenterVmHardwareCdromCreateSpec
	isSet bool
}

func (v NullableVcenterVmHardwareCdromCreateSpec) Get() *VcenterVmHardwareCdromCreateSpec {
	return v.value
}

func (v *NullableVcenterVmHardwareCdromCreateSpec) Set(val *VcenterVmHardwareCdromCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareCdromCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareCdromCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareCdromCreateSpec(val *VcenterVmHardwareCdromCreateSpec) *NullableVcenterVmHardwareCdromCreateSpec {
	return &NullableVcenterVmHardwareCdromCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareCdromCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareCdromCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

