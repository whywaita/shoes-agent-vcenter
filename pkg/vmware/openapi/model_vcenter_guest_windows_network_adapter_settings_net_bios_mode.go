/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode The WindowsNetworkAdapterSettings.NetBIOSMode enumerated type specifies different NetBIOS settings for Windows guest operating systems.
type VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode string

// List of vcenter.guest.windows_network_adapter_settings.net_BIOS_mode
const (
	VCENTERGUESTWINDOWSNETWORKADAPTERSETTINGSNETBIOSMODE_USE_DHCP VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode = "USE_DHCP"
	VCENTERGUESTWINDOWSNETWORKADAPTERSETTINGSNETBIOSMODE_ENABLE VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode = "ENABLE"
	VCENTERGUESTWINDOWSNETWORKADAPTERSETTINGSNETBIOSMODE_DISABLE VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode = "DISABLE"
)

// All allowed values of VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode enum
var AllowedVcenterGuestWindowsNetworkAdapterSettingsNetBIOSModeEnumValues = []VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode{
	"USE_DHCP",
	"ENABLE",
	"DISABLE",
}

func (v *VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode(value)
	for _, existing := range AllowedVcenterGuestWindowsNetworkAdapterSettingsNetBIOSModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode", value)
}

// NewVcenterGuestWindowsNetworkAdapterSettingsNetBIOSModeFromValue returns a pointer to a valid VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterGuestWindowsNetworkAdapterSettingsNetBIOSModeFromValue(v string) (*VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode, error) {
	ev := VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode: valid values are %v", v, AllowedVcenterGuestWindowsNetworkAdapterSettingsNetBIOSModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode) IsValid() bool {
	for _, existing := range AllowedVcenterGuestWindowsNetworkAdapterSettingsNetBIOSModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.guest.windows_network_adapter_settings.net_BIOS_mode value
func (v VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode) Ptr() *VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode {
	return &v
}

type NullableVcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode struct {
	value *VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode
	isSet bool
}

func (v NullableVcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode) Get() *VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode {
	return v.value
}

func (v *NullableVcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode) Set(val *VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode(val *VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode) *NullableVcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode {
	return &NullableVcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode{value: val, isSet: true}
}

func (v NullableVcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

