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

// VcenterVmHardwareBootInfo struct for VcenterVmHardwareBootInfo
type VcenterVmHardwareBootInfo struct {
	Type VcenterVmHardwareBootType `json:"type"`
	// Flag indicating whether to use EFI legacy boot mode. This field is optional and it is only relevant when the value of Boot.Info.type is EFI.
	EfiLegacyBoot *bool `json:"efi_legacy_boot,omitempty"`
	NetworkProtocol *VcenterVmHardwareBootNetworkProtocol `json:"network_protocol,omitempty"`
	// Delay in milliseconds before beginning the firmware boot process when the virtual machine is powered on. This delay may be used to provide a time window for users to connect to the virtual machine console and enter BIOS setup mode.
	Delay int64 `json:"delay"`
	// Flag indicating whether the virtual machine will automatically retry the boot process after a failure.
	Retry bool `json:"retry"`
	// Delay in milliseconds before retrying the boot process after a failure; applicable only when Boot.Info.retry is true.
	RetryDelay int64 `json:"retry_delay"`
	// Flag indicating whether the firmware boot process will automatically enter setup mode the next time the virtual machine boots. Note that this flag will automatically be reset to false once the virtual machine enters setup mode.
	EnterSetupMode bool `json:"enter_setup_mode"`
}

// NewVcenterVmHardwareBootInfo instantiates a new VcenterVmHardwareBootInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareBootInfo(type_ VcenterVmHardwareBootType, delay int64, retry bool, retryDelay int64, enterSetupMode bool) *VcenterVmHardwareBootInfo {
	this := VcenterVmHardwareBootInfo{}
	this.Type = type_
	this.Delay = delay
	this.Retry = retry
	this.RetryDelay = retryDelay
	this.EnterSetupMode = enterSetupMode
	return &this
}

// NewVcenterVmHardwareBootInfoWithDefaults instantiates a new VcenterVmHardwareBootInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareBootInfoWithDefaults() *VcenterVmHardwareBootInfo {
	this := VcenterVmHardwareBootInfo{}
	return &this
}

// GetType returns the Type field value
func (o *VcenterVmHardwareBootInfo) GetType() VcenterVmHardwareBootType {
	if o == nil {
		var ret VcenterVmHardwareBootType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootInfo) GetTypeOk() (*VcenterVmHardwareBootType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterVmHardwareBootInfo) SetType(v VcenterVmHardwareBootType) {
	o.Type = v
}

// GetEfiLegacyBoot returns the EfiLegacyBoot field value if set, zero value otherwise.
func (o *VcenterVmHardwareBootInfo) GetEfiLegacyBoot() bool {
	if o == nil || o.EfiLegacyBoot == nil {
		var ret bool
		return ret
	}
	return *o.EfiLegacyBoot
}

// GetEfiLegacyBootOk returns a tuple with the EfiLegacyBoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootInfo) GetEfiLegacyBootOk() (*bool, bool) {
	if o == nil || o.EfiLegacyBoot == nil {
		return nil, false
	}
	return o.EfiLegacyBoot, true
}

// HasEfiLegacyBoot returns a boolean if a field has been set.
func (o *VcenterVmHardwareBootInfo) HasEfiLegacyBoot() bool {
	if o != nil && o.EfiLegacyBoot != nil {
		return true
	}

	return false
}

// SetEfiLegacyBoot gets a reference to the given bool and assigns it to the EfiLegacyBoot field.
func (o *VcenterVmHardwareBootInfo) SetEfiLegacyBoot(v bool) {
	o.EfiLegacyBoot = &v
}

// GetNetworkProtocol returns the NetworkProtocol field value if set, zero value otherwise.
func (o *VcenterVmHardwareBootInfo) GetNetworkProtocol() VcenterVmHardwareBootNetworkProtocol {
	if o == nil || o.NetworkProtocol == nil {
		var ret VcenterVmHardwareBootNetworkProtocol
		return ret
	}
	return *o.NetworkProtocol
}

// GetNetworkProtocolOk returns a tuple with the NetworkProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootInfo) GetNetworkProtocolOk() (*VcenterVmHardwareBootNetworkProtocol, bool) {
	if o == nil || o.NetworkProtocol == nil {
		return nil, false
	}
	return o.NetworkProtocol, true
}

// HasNetworkProtocol returns a boolean if a field has been set.
func (o *VcenterVmHardwareBootInfo) HasNetworkProtocol() bool {
	if o != nil && o.NetworkProtocol != nil {
		return true
	}

	return false
}

// SetNetworkProtocol gets a reference to the given VcenterVmHardwareBootNetworkProtocol and assigns it to the NetworkProtocol field.
func (o *VcenterVmHardwareBootInfo) SetNetworkProtocol(v VcenterVmHardwareBootNetworkProtocol) {
	o.NetworkProtocol = &v
}

// GetDelay returns the Delay field value
func (o *VcenterVmHardwareBootInfo) GetDelay() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Delay
}

// GetDelayOk returns a tuple with the Delay field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootInfo) GetDelayOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Delay, true
}

// SetDelay sets field value
func (o *VcenterVmHardwareBootInfo) SetDelay(v int64) {
	o.Delay = v
}

// GetRetry returns the Retry field value
func (o *VcenterVmHardwareBootInfo) GetRetry() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Retry
}

// GetRetryOk returns a tuple with the Retry field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootInfo) GetRetryOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Retry, true
}

// SetRetry sets field value
func (o *VcenterVmHardwareBootInfo) SetRetry(v bool) {
	o.Retry = v
}

// GetRetryDelay returns the RetryDelay field value
func (o *VcenterVmHardwareBootInfo) GetRetryDelay() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RetryDelay
}

// GetRetryDelayOk returns a tuple with the RetryDelay field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootInfo) GetRetryDelayOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RetryDelay, true
}

// SetRetryDelay sets field value
func (o *VcenterVmHardwareBootInfo) SetRetryDelay(v int64) {
	o.RetryDelay = v
}

// GetEnterSetupMode returns the EnterSetupMode field value
func (o *VcenterVmHardwareBootInfo) GetEnterSetupMode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnterSetupMode
}

// GetEnterSetupModeOk returns a tuple with the EnterSetupMode field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootInfo) GetEnterSetupModeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EnterSetupMode, true
}

// SetEnterSetupMode sets field value
func (o *VcenterVmHardwareBootInfo) SetEnterSetupMode(v bool) {
	o.EnterSetupMode = v
}

func (o VcenterVmHardwareBootInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.EfiLegacyBoot != nil {
		toSerialize["efi_legacy_boot"] = o.EfiLegacyBoot
	}
	if o.NetworkProtocol != nil {
		toSerialize["network_protocol"] = o.NetworkProtocol
	}
	if true {
		toSerialize["delay"] = o.Delay
	}
	if true {
		toSerialize["retry"] = o.Retry
	}
	if true {
		toSerialize["retry_delay"] = o.RetryDelay
	}
	if true {
		toSerialize["enter_setup_mode"] = o.EnterSetupMode
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareBootInfo struct {
	value *VcenterVmHardwareBootInfo
	isSet bool
}

func (v NullableVcenterVmHardwareBootInfo) Get() *VcenterVmHardwareBootInfo {
	return v.value
}

func (v *NullableVcenterVmHardwareBootInfo) Set(val *VcenterVmHardwareBootInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareBootInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareBootInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareBootInfo(val *VcenterVmHardwareBootInfo) *NullableVcenterVmHardwareBootInfo {
	return &NullableVcenterVmHardwareBootInfo{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareBootInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareBootInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


