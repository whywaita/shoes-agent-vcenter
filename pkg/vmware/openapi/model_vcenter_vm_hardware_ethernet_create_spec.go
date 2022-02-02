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

// VcenterVmHardwareEthernetCreateSpec struct for VcenterVmHardwareEthernetCreateSpec
type VcenterVmHardwareEthernetCreateSpec struct {
	Type *VcenterVmHardwareEthernetEmulationType `json:"type,omitempty"`
	// Flag indicating whether Universal Pass-Through (UPT) compatibility is enabled on this virtual Ethernet adapter. If unset, defaults to false.
	UptCompatibilityEnabled *bool `json:"upt_compatibility_enabled,omitempty"`
	MacType *VcenterVmHardwareEthernetMacAddressType `json:"mac_type,omitempty"`
	// MAC address. Workaround for PR1459647
	MacAddress *string `json:"mac_address,omitempty"`
	// Address of the virtual Ethernet adapter on the PCI bus. If the PCI address is invalid, the server will change when it the VM is started or as the device is hot added. If unset, the server will choose an available address when the virtual machine is powered on.
	PciSlotNumber *int64 `json:"pci_slot_number,omitempty"`
	// Flag indicating whether wake-on-LAN is enabled on this virtual Ethernet adapter. Defaults to false if unset.
	WakeOnLanEnabled *bool `json:"wake_on_lan_enabled,omitempty"`
	Backing *VcenterVmHardwareEthernetBackingSpec `json:"backing,omitempty"`
	// Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. Defaults to false if unset.
	StartConnected *bool `json:"start_connected,omitempty"`
	// Flag indicating whether the guest can connect and disconnect the device. Defaults to false if unset.
	AllowGuestControl *bool `json:"allow_guest_control,omitempty"`
}

// NewVcenterVmHardwareEthernetCreateSpec instantiates a new VcenterVmHardwareEthernetCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareEthernetCreateSpec() *VcenterVmHardwareEthernetCreateSpec {
	this := VcenterVmHardwareEthernetCreateSpec{}
	return &this
}

// NewVcenterVmHardwareEthernetCreateSpecWithDefaults instantiates a new VcenterVmHardwareEthernetCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareEthernetCreateSpecWithDefaults() *VcenterVmHardwareEthernetCreateSpec {
	this := VcenterVmHardwareEthernetCreateSpec{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VcenterVmHardwareEthernetCreateSpec) GetType() VcenterVmHardwareEthernetEmulationType {
	if o == nil || o.Type == nil {
		var ret VcenterVmHardwareEthernetEmulationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) GetTypeOk() (*VcenterVmHardwareEthernetEmulationType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given VcenterVmHardwareEthernetEmulationType and assigns it to the Type field.
func (o *VcenterVmHardwareEthernetCreateSpec) SetType(v VcenterVmHardwareEthernetEmulationType) {
	o.Type = &v
}

// GetUptCompatibilityEnabled returns the UptCompatibilityEnabled field value if set, zero value otherwise.
func (o *VcenterVmHardwareEthernetCreateSpec) GetUptCompatibilityEnabled() bool {
	if o == nil || o.UptCompatibilityEnabled == nil {
		var ret bool
		return ret
	}
	return *o.UptCompatibilityEnabled
}

// GetUptCompatibilityEnabledOk returns a tuple with the UptCompatibilityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) GetUptCompatibilityEnabledOk() (*bool, bool) {
	if o == nil || o.UptCompatibilityEnabled == nil {
		return nil, false
	}
	return o.UptCompatibilityEnabled, true
}

// HasUptCompatibilityEnabled returns a boolean if a field has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) HasUptCompatibilityEnabled() bool {
	if o != nil && o.UptCompatibilityEnabled != nil {
		return true
	}

	return false
}

// SetUptCompatibilityEnabled gets a reference to the given bool and assigns it to the UptCompatibilityEnabled field.
func (o *VcenterVmHardwareEthernetCreateSpec) SetUptCompatibilityEnabled(v bool) {
	o.UptCompatibilityEnabled = &v
}

// GetMacType returns the MacType field value if set, zero value otherwise.
func (o *VcenterVmHardwareEthernetCreateSpec) GetMacType() VcenterVmHardwareEthernetMacAddressType {
	if o == nil || o.MacType == nil {
		var ret VcenterVmHardwareEthernetMacAddressType
		return ret
	}
	return *o.MacType
}

// GetMacTypeOk returns a tuple with the MacType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) GetMacTypeOk() (*VcenterVmHardwareEthernetMacAddressType, bool) {
	if o == nil || o.MacType == nil {
		return nil, false
	}
	return o.MacType, true
}

// HasMacType returns a boolean if a field has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) HasMacType() bool {
	if o != nil && o.MacType != nil {
		return true
	}

	return false
}

// SetMacType gets a reference to the given VcenterVmHardwareEthernetMacAddressType and assigns it to the MacType field.
func (o *VcenterVmHardwareEthernetCreateSpec) SetMacType(v VcenterVmHardwareEthernetMacAddressType) {
	o.MacType = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *VcenterVmHardwareEthernetCreateSpec) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *VcenterVmHardwareEthernetCreateSpec) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetPciSlotNumber returns the PciSlotNumber field value if set, zero value otherwise.
func (o *VcenterVmHardwareEthernetCreateSpec) GetPciSlotNumber() int64 {
	if o == nil || o.PciSlotNumber == nil {
		var ret int64
		return ret
	}
	return *o.PciSlotNumber
}

// GetPciSlotNumberOk returns a tuple with the PciSlotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) GetPciSlotNumberOk() (*int64, bool) {
	if o == nil || o.PciSlotNumber == nil {
		return nil, false
	}
	return o.PciSlotNumber, true
}

// HasPciSlotNumber returns a boolean if a field has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) HasPciSlotNumber() bool {
	if o != nil && o.PciSlotNumber != nil {
		return true
	}

	return false
}

// SetPciSlotNumber gets a reference to the given int64 and assigns it to the PciSlotNumber field.
func (o *VcenterVmHardwareEthernetCreateSpec) SetPciSlotNumber(v int64) {
	o.PciSlotNumber = &v
}

// GetWakeOnLanEnabled returns the WakeOnLanEnabled field value if set, zero value otherwise.
func (o *VcenterVmHardwareEthernetCreateSpec) GetWakeOnLanEnabled() bool {
	if o == nil || o.WakeOnLanEnabled == nil {
		var ret bool
		return ret
	}
	return *o.WakeOnLanEnabled
}

// GetWakeOnLanEnabledOk returns a tuple with the WakeOnLanEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) GetWakeOnLanEnabledOk() (*bool, bool) {
	if o == nil || o.WakeOnLanEnabled == nil {
		return nil, false
	}
	return o.WakeOnLanEnabled, true
}

// HasWakeOnLanEnabled returns a boolean if a field has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) HasWakeOnLanEnabled() bool {
	if o != nil && o.WakeOnLanEnabled != nil {
		return true
	}

	return false
}

// SetWakeOnLanEnabled gets a reference to the given bool and assigns it to the WakeOnLanEnabled field.
func (o *VcenterVmHardwareEthernetCreateSpec) SetWakeOnLanEnabled(v bool) {
	o.WakeOnLanEnabled = &v
}

// GetBacking returns the Backing field value if set, zero value otherwise.
func (o *VcenterVmHardwareEthernetCreateSpec) GetBacking() VcenterVmHardwareEthernetBackingSpec {
	if o == nil || o.Backing == nil {
		var ret VcenterVmHardwareEthernetBackingSpec
		return ret
	}
	return *o.Backing
}

// GetBackingOk returns a tuple with the Backing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) GetBackingOk() (*VcenterVmHardwareEthernetBackingSpec, bool) {
	if o == nil || o.Backing == nil {
		return nil, false
	}
	return o.Backing, true
}

// HasBacking returns a boolean if a field has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) HasBacking() bool {
	if o != nil && o.Backing != nil {
		return true
	}

	return false
}

// SetBacking gets a reference to the given VcenterVmHardwareEthernetBackingSpec and assigns it to the Backing field.
func (o *VcenterVmHardwareEthernetCreateSpec) SetBacking(v VcenterVmHardwareEthernetBackingSpec) {
	o.Backing = &v
}

// GetStartConnected returns the StartConnected field value if set, zero value otherwise.
func (o *VcenterVmHardwareEthernetCreateSpec) GetStartConnected() bool {
	if o == nil || o.StartConnected == nil {
		var ret bool
		return ret
	}
	return *o.StartConnected
}

// GetStartConnectedOk returns a tuple with the StartConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) GetStartConnectedOk() (*bool, bool) {
	if o == nil || o.StartConnected == nil {
		return nil, false
	}
	return o.StartConnected, true
}

// HasStartConnected returns a boolean if a field has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) HasStartConnected() bool {
	if o != nil && o.StartConnected != nil {
		return true
	}

	return false
}

// SetStartConnected gets a reference to the given bool and assigns it to the StartConnected field.
func (o *VcenterVmHardwareEthernetCreateSpec) SetStartConnected(v bool) {
	o.StartConnected = &v
}

// GetAllowGuestControl returns the AllowGuestControl field value if set, zero value otherwise.
func (o *VcenterVmHardwareEthernetCreateSpec) GetAllowGuestControl() bool {
	if o == nil || o.AllowGuestControl == nil {
		var ret bool
		return ret
	}
	return *o.AllowGuestControl
}

// GetAllowGuestControlOk returns a tuple with the AllowGuestControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) GetAllowGuestControlOk() (*bool, bool) {
	if o == nil || o.AllowGuestControl == nil {
		return nil, false
	}
	return o.AllowGuestControl, true
}

// HasAllowGuestControl returns a boolean if a field has been set.
func (o *VcenterVmHardwareEthernetCreateSpec) HasAllowGuestControl() bool {
	if o != nil && o.AllowGuestControl != nil {
		return true
	}

	return false
}

// SetAllowGuestControl gets a reference to the given bool and assigns it to the AllowGuestControl field.
func (o *VcenterVmHardwareEthernetCreateSpec) SetAllowGuestControl(v bool) {
	o.AllowGuestControl = &v
}

func (o VcenterVmHardwareEthernetCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UptCompatibilityEnabled != nil {
		toSerialize["upt_compatibility_enabled"] = o.UptCompatibilityEnabled
	}
	if o.MacType != nil {
		toSerialize["mac_type"] = o.MacType
	}
	if o.MacAddress != nil {
		toSerialize["mac_address"] = o.MacAddress
	}
	if o.PciSlotNumber != nil {
		toSerialize["pci_slot_number"] = o.PciSlotNumber
	}
	if o.WakeOnLanEnabled != nil {
		toSerialize["wake_on_lan_enabled"] = o.WakeOnLanEnabled
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

type NullableVcenterVmHardwareEthernetCreateSpec struct {
	value *VcenterVmHardwareEthernetCreateSpec
	isSet bool
}

func (v NullableVcenterVmHardwareEthernetCreateSpec) Get() *VcenterVmHardwareEthernetCreateSpec {
	return v.value
}

func (v *NullableVcenterVmHardwareEthernetCreateSpec) Set(val *VcenterVmHardwareEthernetCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareEthernetCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareEthernetCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareEthernetCreateSpec(val *VcenterVmHardwareEthernetCreateSpec) *NullableVcenterVmHardwareEthernetCreateSpec {
	return &NullableVcenterVmHardwareEthernetCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareEthernetCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareEthernetCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

