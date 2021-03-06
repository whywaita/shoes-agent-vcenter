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

// VcenterGuestConfigurationSpec struct for VcenterGuestConfigurationSpec
type VcenterGuestConfigurationSpec struct {
	WindowsConfig *VcenterGuestWindowsConfiguration `json:"windows_config,omitempty"`
	LinuxConfig *VcenterGuestLinuxConfiguration `json:"linux_config,omitempty"`
	CloudConfig *VcenterGuestCloudConfiguration `json:"cloud_config,omitempty"`
}

// NewVcenterGuestConfigurationSpec instantiates a new VcenterGuestConfigurationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterGuestConfigurationSpec() *VcenterGuestConfigurationSpec {
	this := VcenterGuestConfigurationSpec{}
	return &this
}

// NewVcenterGuestConfigurationSpecWithDefaults instantiates a new VcenterGuestConfigurationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterGuestConfigurationSpecWithDefaults() *VcenterGuestConfigurationSpec {
	this := VcenterGuestConfigurationSpec{}
	return &this
}

// GetWindowsConfig returns the WindowsConfig field value if set, zero value otherwise.
func (o *VcenterGuestConfigurationSpec) GetWindowsConfig() VcenterGuestWindowsConfiguration {
	if o == nil || o.WindowsConfig == nil {
		var ret VcenterGuestWindowsConfiguration
		return ret
	}
	return *o.WindowsConfig
}

// GetWindowsConfigOk returns a tuple with the WindowsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterGuestConfigurationSpec) GetWindowsConfigOk() (*VcenterGuestWindowsConfiguration, bool) {
	if o == nil || o.WindowsConfig == nil {
		return nil, false
	}
	return o.WindowsConfig, true
}

// HasWindowsConfig returns a boolean if a field has been set.
func (o *VcenterGuestConfigurationSpec) HasWindowsConfig() bool {
	if o != nil && o.WindowsConfig != nil {
		return true
	}

	return false
}

// SetWindowsConfig gets a reference to the given VcenterGuestWindowsConfiguration and assigns it to the WindowsConfig field.
func (o *VcenterGuestConfigurationSpec) SetWindowsConfig(v VcenterGuestWindowsConfiguration) {
	o.WindowsConfig = &v
}

// GetLinuxConfig returns the LinuxConfig field value if set, zero value otherwise.
func (o *VcenterGuestConfigurationSpec) GetLinuxConfig() VcenterGuestLinuxConfiguration {
	if o == nil || o.LinuxConfig == nil {
		var ret VcenterGuestLinuxConfiguration
		return ret
	}
	return *o.LinuxConfig
}

// GetLinuxConfigOk returns a tuple with the LinuxConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterGuestConfigurationSpec) GetLinuxConfigOk() (*VcenterGuestLinuxConfiguration, bool) {
	if o == nil || o.LinuxConfig == nil {
		return nil, false
	}
	return o.LinuxConfig, true
}

// HasLinuxConfig returns a boolean if a field has been set.
func (o *VcenterGuestConfigurationSpec) HasLinuxConfig() bool {
	if o != nil && o.LinuxConfig != nil {
		return true
	}

	return false
}

// SetLinuxConfig gets a reference to the given VcenterGuestLinuxConfiguration and assigns it to the LinuxConfig field.
func (o *VcenterGuestConfigurationSpec) SetLinuxConfig(v VcenterGuestLinuxConfiguration) {
	o.LinuxConfig = &v
}

// GetCloudConfig returns the CloudConfig field value if set, zero value otherwise.
func (o *VcenterGuestConfigurationSpec) GetCloudConfig() VcenterGuestCloudConfiguration {
	if o == nil || o.CloudConfig == nil {
		var ret VcenterGuestCloudConfiguration
		return ret
	}
	return *o.CloudConfig
}

// GetCloudConfigOk returns a tuple with the CloudConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterGuestConfigurationSpec) GetCloudConfigOk() (*VcenterGuestCloudConfiguration, bool) {
	if o == nil || o.CloudConfig == nil {
		return nil, false
	}
	return o.CloudConfig, true
}

// HasCloudConfig returns a boolean if a field has been set.
func (o *VcenterGuestConfigurationSpec) HasCloudConfig() bool {
	if o != nil && o.CloudConfig != nil {
		return true
	}

	return false
}

// SetCloudConfig gets a reference to the given VcenterGuestCloudConfiguration and assigns it to the CloudConfig field.
func (o *VcenterGuestConfigurationSpec) SetCloudConfig(v VcenterGuestCloudConfiguration) {
	o.CloudConfig = &v
}

func (o VcenterGuestConfigurationSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WindowsConfig != nil {
		toSerialize["windows_config"] = o.WindowsConfig
	}
	if o.LinuxConfig != nil {
		toSerialize["linux_config"] = o.LinuxConfig
	}
	if o.CloudConfig != nil {
		toSerialize["cloud_config"] = o.CloudConfig
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterGuestConfigurationSpec struct {
	value *VcenterGuestConfigurationSpec
	isSet bool
}

func (v NullableVcenterGuestConfigurationSpec) Get() *VcenterGuestConfigurationSpec {
	return v.value
}

func (v *NullableVcenterGuestConfigurationSpec) Set(val *VcenterGuestConfigurationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterGuestConfigurationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterGuestConfigurationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterGuestConfigurationSpec(val *VcenterGuestConfigurationSpec) *NullableVcenterGuestConfigurationSpec {
	return &NullableVcenterGuestConfigurationSpec{value: val, isSet: true}
}

func (v NullableVcenterGuestConfigurationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterGuestConfigurationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


