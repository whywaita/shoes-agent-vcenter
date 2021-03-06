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

// VcenterVmHardwareBootUpdateSpec struct for VcenterVmHardwareBootUpdateSpec
type VcenterVmHardwareBootUpdateSpec struct {
	Type *VcenterVmHardwareBootType `json:"type,omitempty"`
	// Flag indicating whether to use EFI legacy boot mode. If unset, the value is unchanged.
	EfiLegacyBoot *bool `json:"efi_legacy_boot,omitempty"`
	NetworkProtocol *VcenterVmHardwareBootNetworkProtocol `json:"network_protocol,omitempty"`
	// Delay in milliseconds before beginning the firmware boot process when the virtual machine is powered on. This delay may be used to provide a time window for users to connect to the virtual machine console and enter BIOS setup mode. If unset, the value is unchanged.
	Delay *int64 `json:"delay,omitempty"`
	// Flag indicating whether the virtual machine should automatically retry the boot process after a failure. If unset, the value is unchanged.
	Retry *bool `json:"retry,omitempty"`
	// Delay in milliseconds before retrying the boot process after a failure; applicable only when Boot.Info.retry is true. If unset, the value is unchanged.
	RetryDelay *int64 `json:"retry_delay,omitempty"`
	// Flag indicating whether the firmware boot process should automatically enter setup mode the next time the virtual machine boots. Note that this flag will automatically be reset to false once the virtual machine enters setup mode. If unset, the value is unchanged.
	EnterSetupMode *bool `json:"enter_setup_mode,omitempty"`
}

// NewVcenterVmHardwareBootUpdateSpec instantiates a new VcenterVmHardwareBootUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareBootUpdateSpec() *VcenterVmHardwareBootUpdateSpec {
	this := VcenterVmHardwareBootUpdateSpec{}
	return &this
}

// NewVcenterVmHardwareBootUpdateSpecWithDefaults instantiates a new VcenterVmHardwareBootUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareBootUpdateSpecWithDefaults() *VcenterVmHardwareBootUpdateSpec {
	this := VcenterVmHardwareBootUpdateSpec{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VcenterVmHardwareBootUpdateSpec) GetType() VcenterVmHardwareBootType {
	if o == nil || o.Type == nil {
		var ret VcenterVmHardwareBootType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootUpdateSpec) GetTypeOk() (*VcenterVmHardwareBootType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VcenterVmHardwareBootUpdateSpec) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given VcenterVmHardwareBootType and assigns it to the Type field.
func (o *VcenterVmHardwareBootUpdateSpec) SetType(v VcenterVmHardwareBootType) {
	o.Type = &v
}

// GetEfiLegacyBoot returns the EfiLegacyBoot field value if set, zero value otherwise.
func (o *VcenterVmHardwareBootUpdateSpec) GetEfiLegacyBoot() bool {
	if o == nil || o.EfiLegacyBoot == nil {
		var ret bool
		return ret
	}
	return *o.EfiLegacyBoot
}

// GetEfiLegacyBootOk returns a tuple with the EfiLegacyBoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootUpdateSpec) GetEfiLegacyBootOk() (*bool, bool) {
	if o == nil || o.EfiLegacyBoot == nil {
		return nil, false
	}
	return o.EfiLegacyBoot, true
}

// HasEfiLegacyBoot returns a boolean if a field has been set.
func (o *VcenterVmHardwareBootUpdateSpec) HasEfiLegacyBoot() bool {
	if o != nil && o.EfiLegacyBoot != nil {
		return true
	}

	return false
}

// SetEfiLegacyBoot gets a reference to the given bool and assigns it to the EfiLegacyBoot field.
func (o *VcenterVmHardwareBootUpdateSpec) SetEfiLegacyBoot(v bool) {
	o.EfiLegacyBoot = &v
}

// GetNetworkProtocol returns the NetworkProtocol field value if set, zero value otherwise.
func (o *VcenterVmHardwareBootUpdateSpec) GetNetworkProtocol() VcenterVmHardwareBootNetworkProtocol {
	if o == nil || o.NetworkProtocol == nil {
		var ret VcenterVmHardwareBootNetworkProtocol
		return ret
	}
	return *o.NetworkProtocol
}

// GetNetworkProtocolOk returns a tuple with the NetworkProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootUpdateSpec) GetNetworkProtocolOk() (*VcenterVmHardwareBootNetworkProtocol, bool) {
	if o == nil || o.NetworkProtocol == nil {
		return nil, false
	}
	return o.NetworkProtocol, true
}

// HasNetworkProtocol returns a boolean if a field has been set.
func (o *VcenterVmHardwareBootUpdateSpec) HasNetworkProtocol() bool {
	if o != nil && o.NetworkProtocol != nil {
		return true
	}

	return false
}

// SetNetworkProtocol gets a reference to the given VcenterVmHardwareBootNetworkProtocol and assigns it to the NetworkProtocol field.
func (o *VcenterVmHardwareBootUpdateSpec) SetNetworkProtocol(v VcenterVmHardwareBootNetworkProtocol) {
	o.NetworkProtocol = &v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *VcenterVmHardwareBootUpdateSpec) GetDelay() int64 {
	if o == nil || o.Delay == nil {
		var ret int64
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootUpdateSpec) GetDelayOk() (*int64, bool) {
	if o == nil || o.Delay == nil {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *VcenterVmHardwareBootUpdateSpec) HasDelay() bool {
	if o != nil && o.Delay != nil {
		return true
	}

	return false
}

// SetDelay gets a reference to the given int64 and assigns it to the Delay field.
func (o *VcenterVmHardwareBootUpdateSpec) SetDelay(v int64) {
	o.Delay = &v
}

// GetRetry returns the Retry field value if set, zero value otherwise.
func (o *VcenterVmHardwareBootUpdateSpec) GetRetry() bool {
	if o == nil || o.Retry == nil {
		var ret bool
		return ret
	}
	return *o.Retry
}

// GetRetryOk returns a tuple with the Retry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootUpdateSpec) GetRetryOk() (*bool, bool) {
	if o == nil || o.Retry == nil {
		return nil, false
	}
	return o.Retry, true
}

// HasRetry returns a boolean if a field has been set.
func (o *VcenterVmHardwareBootUpdateSpec) HasRetry() bool {
	if o != nil && o.Retry != nil {
		return true
	}

	return false
}

// SetRetry gets a reference to the given bool and assigns it to the Retry field.
func (o *VcenterVmHardwareBootUpdateSpec) SetRetry(v bool) {
	o.Retry = &v
}

// GetRetryDelay returns the RetryDelay field value if set, zero value otherwise.
func (o *VcenterVmHardwareBootUpdateSpec) GetRetryDelay() int64 {
	if o == nil || o.RetryDelay == nil {
		var ret int64
		return ret
	}
	return *o.RetryDelay
}

// GetRetryDelayOk returns a tuple with the RetryDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootUpdateSpec) GetRetryDelayOk() (*int64, bool) {
	if o == nil || o.RetryDelay == nil {
		return nil, false
	}
	return o.RetryDelay, true
}

// HasRetryDelay returns a boolean if a field has been set.
func (o *VcenterVmHardwareBootUpdateSpec) HasRetryDelay() bool {
	if o != nil && o.RetryDelay != nil {
		return true
	}

	return false
}

// SetRetryDelay gets a reference to the given int64 and assigns it to the RetryDelay field.
func (o *VcenterVmHardwareBootUpdateSpec) SetRetryDelay(v int64) {
	o.RetryDelay = &v
}

// GetEnterSetupMode returns the EnterSetupMode field value if set, zero value otherwise.
func (o *VcenterVmHardwareBootUpdateSpec) GetEnterSetupMode() bool {
	if o == nil || o.EnterSetupMode == nil {
		var ret bool
		return ret
	}
	return *o.EnterSetupMode
}

// GetEnterSetupModeOk returns a tuple with the EnterSetupMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootUpdateSpec) GetEnterSetupModeOk() (*bool, bool) {
	if o == nil || o.EnterSetupMode == nil {
		return nil, false
	}
	return o.EnterSetupMode, true
}

// HasEnterSetupMode returns a boolean if a field has been set.
func (o *VcenterVmHardwareBootUpdateSpec) HasEnterSetupMode() bool {
	if o != nil && o.EnterSetupMode != nil {
		return true
	}

	return false
}

// SetEnterSetupMode gets a reference to the given bool and assigns it to the EnterSetupMode field.
func (o *VcenterVmHardwareBootUpdateSpec) SetEnterSetupMode(v bool) {
	o.EnterSetupMode = &v
}

func (o VcenterVmHardwareBootUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.EfiLegacyBoot != nil {
		toSerialize["efi_legacy_boot"] = o.EfiLegacyBoot
	}
	if o.NetworkProtocol != nil {
		toSerialize["network_protocol"] = o.NetworkProtocol
	}
	if o.Delay != nil {
		toSerialize["delay"] = o.Delay
	}
	if o.Retry != nil {
		toSerialize["retry"] = o.Retry
	}
	if o.RetryDelay != nil {
		toSerialize["retry_delay"] = o.RetryDelay
	}
	if o.EnterSetupMode != nil {
		toSerialize["enter_setup_mode"] = o.EnterSetupMode
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareBootUpdateSpec struct {
	value *VcenterVmHardwareBootUpdateSpec
	isSet bool
}

func (v NullableVcenterVmHardwareBootUpdateSpec) Get() *VcenterVmHardwareBootUpdateSpec {
	return v.value
}

func (v *NullableVcenterVmHardwareBootUpdateSpec) Set(val *VcenterVmHardwareBootUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareBootUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareBootUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareBootUpdateSpec(val *VcenterVmHardwareBootUpdateSpec) *NullableVcenterVmHardwareBootUpdateSpec {
	return &NullableVcenterVmHardwareBootUpdateSpec{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareBootUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareBootUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


