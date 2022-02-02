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

// VcenterVmHardwareEthernetInfo struct for VcenterVmHardwareEthernetInfo
type VcenterVmHardwareEthernetInfo struct {
	// Device label.
	Label string `json:"label"`
	Type VcenterVmHardwareEthernetEmulationType `json:"type"`
	// Flag indicating whether Universal Pass-Through (UPT) compatibility is enabled on this virtual Ethernet adapter. This field is optional and it is only relevant when the value of Ethernet.Info.type is VMXNET3.
	UptCompatibilityEnabled *bool `json:"upt_compatibility_enabled,omitempty"`
	MacType VcenterVmHardwareEthernetMacAddressType `json:"mac_type"`
	// MAC address. May be unset if Ethernet.Info.mac-type is MANUAL and has not been specified, or if Ethernet.Info.mac-type is GENERATED and the virtual machine has never been powered on since the Ethernet adapter was created.
	MacAddress *string `json:"mac_address,omitempty"`
	// Address of the virtual Ethernet adapter on the PCI bus. If the PCI address is invalid, the server will change it when the VM is started or as the device is hot added. May be unset if the virtual machine has never been powered on since the adapter was created.
	PciSlotNumber *int64 `json:"pci_slot_number,omitempty"`
	// Flag indicating whether wake-on-LAN is enabled on this virtual Ethernet adapter.
	WakeOnLanEnabled bool `json:"wake_on_lan_enabled"`
	Backing VcenterVmHardwareEthernetBackingInfo `json:"backing"`
	State VcenterVmHardwareConnectionState `json:"state"`
	// Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected bool `json:"start_connected"`
	// Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl bool `json:"allow_guest_control"`
}

// NewVcenterVmHardwareEthernetInfo instantiates a new VcenterVmHardwareEthernetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareEthernetInfo(label string, type_ VcenterVmHardwareEthernetEmulationType, macType VcenterVmHardwareEthernetMacAddressType, wakeOnLanEnabled bool, backing VcenterVmHardwareEthernetBackingInfo, state VcenterVmHardwareConnectionState, startConnected bool, allowGuestControl bool) *VcenterVmHardwareEthernetInfo {
	this := VcenterVmHardwareEthernetInfo{}
	this.Label = label
	this.Type = type_
	this.MacType = macType
	this.WakeOnLanEnabled = wakeOnLanEnabled
	this.Backing = backing
	this.State = state
	this.StartConnected = startConnected
	this.AllowGuestControl = allowGuestControl
	return &this
}

// NewVcenterVmHardwareEthernetInfoWithDefaults instantiates a new VcenterVmHardwareEthernetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareEthernetInfoWithDefaults() *VcenterVmHardwareEthernetInfo {
	this := VcenterVmHardwareEthernetInfo{}
	return &this
}

// GetLabel returns the Label field value
func (o *VcenterVmHardwareEthernetInfo) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetInfo) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *VcenterVmHardwareEthernetInfo) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value
func (o *VcenterVmHardwareEthernetInfo) GetType() VcenterVmHardwareEthernetEmulationType {
	if o == nil {
		var ret VcenterVmHardwareEthernetEmulationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetInfo) GetTypeOk() (*VcenterVmHardwareEthernetEmulationType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterVmHardwareEthernetInfo) SetType(v VcenterVmHardwareEthernetEmulationType) {
	o.Type = v
}

// GetUptCompatibilityEnabled returns the UptCompatibilityEnabled field value if set, zero value otherwise.
func (o *VcenterVmHardwareEthernetInfo) GetUptCompatibilityEnabled() bool {
	if o == nil || o.UptCompatibilityEnabled == nil {
		var ret bool
		return ret
	}
	return *o.UptCompatibilityEnabled
}

// GetUptCompatibilityEnabledOk returns a tuple with the UptCompatibilityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetInfo) GetUptCompatibilityEnabledOk() (*bool, bool) {
	if o == nil || o.UptCompatibilityEnabled == nil {
		return nil, false
	}
	return o.UptCompatibilityEnabled, true
}

// HasUptCompatibilityEnabled returns a boolean if a field has been set.
func (o *VcenterVmHardwareEthernetInfo) HasUptCompatibilityEnabled() bool {
	if o != nil && o.UptCompatibilityEnabled != nil {
		return true
	}

	return false
}

// SetUptCompatibilityEnabled gets a reference to the given bool and assigns it to the UptCompatibilityEnabled field.
func (o *VcenterVmHardwareEthernetInfo) SetUptCompatibilityEnabled(v bool) {
	o.UptCompatibilityEnabled = &v
}

// GetMacType returns the MacType field value
func (o *VcenterVmHardwareEthernetInfo) GetMacType() VcenterVmHardwareEthernetMacAddressType {
	if o == nil {
		var ret VcenterVmHardwareEthernetMacAddressType
		return ret
	}

	return o.MacType
}

// GetMacTypeOk returns a tuple with the MacType field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetInfo) GetMacTypeOk() (*VcenterVmHardwareEthernetMacAddressType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MacType, true
}

// SetMacType sets field value
func (o *VcenterVmHardwareEthernetInfo) SetMacType(v VcenterVmHardwareEthernetMacAddressType) {
	o.MacType = v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *VcenterVmHardwareEthernetInfo) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetInfo) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VcenterVmHardwareEthernetInfo) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *VcenterVmHardwareEthernetInfo) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetPciSlotNumber returns the PciSlotNumber field value if set, zero value otherwise.
func (o *VcenterVmHardwareEthernetInfo) GetPciSlotNumber() int64 {
	if o == nil || o.PciSlotNumber == nil {
		var ret int64
		return ret
	}
	return *o.PciSlotNumber
}

// GetPciSlotNumberOk returns a tuple with the PciSlotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetInfo) GetPciSlotNumberOk() (*int64, bool) {
	if o == nil || o.PciSlotNumber == nil {
		return nil, false
	}
	return o.PciSlotNumber, true
}

// HasPciSlotNumber returns a boolean if a field has been set.
func (o *VcenterVmHardwareEthernetInfo) HasPciSlotNumber() bool {
	if o != nil && o.PciSlotNumber != nil {
		return true
	}

	return false
}

// SetPciSlotNumber gets a reference to the given int64 and assigns it to the PciSlotNumber field.
func (o *VcenterVmHardwareEthernetInfo) SetPciSlotNumber(v int64) {
	o.PciSlotNumber = &v
}

// GetWakeOnLanEnabled returns the WakeOnLanEnabled field value
func (o *VcenterVmHardwareEthernetInfo) GetWakeOnLanEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WakeOnLanEnabled
}

// GetWakeOnLanEnabledOk returns a tuple with the WakeOnLanEnabled field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetInfo) GetWakeOnLanEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WakeOnLanEnabled, true
}

// SetWakeOnLanEnabled sets field value
func (o *VcenterVmHardwareEthernetInfo) SetWakeOnLanEnabled(v bool) {
	o.WakeOnLanEnabled = v
}

// GetBacking returns the Backing field value
func (o *VcenterVmHardwareEthernetInfo) GetBacking() VcenterVmHardwareEthernetBackingInfo {
	if o == nil {
		var ret VcenterVmHardwareEthernetBackingInfo
		return ret
	}

	return o.Backing
}

// GetBackingOk returns a tuple with the Backing field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetInfo) GetBackingOk() (*VcenterVmHardwareEthernetBackingInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Backing, true
}

// SetBacking sets field value
func (o *VcenterVmHardwareEthernetInfo) SetBacking(v VcenterVmHardwareEthernetBackingInfo) {
	o.Backing = v
}

// GetState returns the State field value
func (o *VcenterVmHardwareEthernetInfo) GetState() VcenterVmHardwareConnectionState {
	if o == nil {
		var ret VcenterVmHardwareConnectionState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetInfo) GetStateOk() (*VcenterVmHardwareConnectionState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *VcenterVmHardwareEthernetInfo) SetState(v VcenterVmHardwareConnectionState) {
	o.State = v
}

// GetStartConnected returns the StartConnected field value
func (o *VcenterVmHardwareEthernetInfo) GetStartConnected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StartConnected
}

// GetStartConnectedOk returns a tuple with the StartConnected field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetInfo) GetStartConnectedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartConnected, true
}

// SetStartConnected sets field value
func (o *VcenterVmHardwareEthernetInfo) SetStartConnected(v bool) {
	o.StartConnected = v
}

// GetAllowGuestControl returns the AllowGuestControl field value
func (o *VcenterVmHardwareEthernetInfo) GetAllowGuestControl() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowGuestControl
}

// GetAllowGuestControlOk returns a tuple with the AllowGuestControl field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetInfo) GetAllowGuestControlOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AllowGuestControl, true
}

// SetAllowGuestControl sets field value
func (o *VcenterVmHardwareEthernetInfo) SetAllowGuestControl(v bool) {
	o.AllowGuestControl = v
}

func (o VcenterVmHardwareEthernetInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UptCompatibilityEnabled != nil {
		toSerialize["upt_compatibility_enabled"] = o.UptCompatibilityEnabled
	}
	if true {
		toSerialize["mac_type"] = o.MacType
	}
	if o.MacAddress != nil {
		toSerialize["mac_address"] = o.MacAddress
	}
	if o.PciSlotNumber != nil {
		toSerialize["pci_slot_number"] = o.PciSlotNumber
	}
	if true {
		toSerialize["wake_on_lan_enabled"] = o.WakeOnLanEnabled
	}
	if true {
		toSerialize["backing"] = o.Backing
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["start_connected"] = o.StartConnected
	}
	if true {
		toSerialize["allow_guest_control"] = o.AllowGuestControl
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareEthernetInfo struct {
	value *VcenterVmHardwareEthernetInfo
	isSet bool
}

func (v NullableVcenterVmHardwareEthernetInfo) Get() *VcenterVmHardwareEthernetInfo {
	return v.value
}

func (v *NullableVcenterVmHardwareEthernetInfo) Set(val *VcenterVmHardwareEthernetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareEthernetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareEthernetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareEthernetInfo(val *VcenterVmHardwareEthernetInfo) *NullableVcenterVmHardwareEthernetInfo {
	return &NullableVcenterVmHardwareEthernetInfo{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareEthernetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareEthernetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

