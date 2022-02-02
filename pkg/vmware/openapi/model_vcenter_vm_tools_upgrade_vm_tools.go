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

// VcenterVmToolsUpgradeVmTools struct for VcenterVmToolsUpgradeVmTools
type VcenterVmToolsUpgradeVmTools struct {
	// Command line options passed to the installer to modify the installation procedure for Tools. Set if any additional options are desired.
	CommandLineOptions *string `json:"command_line_options,omitempty"`
}

// NewVcenterVmToolsUpgradeVmTools instantiates a new VcenterVmToolsUpgradeVmTools object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmToolsUpgradeVmTools() *VcenterVmToolsUpgradeVmTools {
	this := VcenterVmToolsUpgradeVmTools{}
	return &this
}

// NewVcenterVmToolsUpgradeVmToolsWithDefaults instantiates a new VcenterVmToolsUpgradeVmTools object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmToolsUpgradeVmToolsWithDefaults() *VcenterVmToolsUpgradeVmTools {
	this := VcenterVmToolsUpgradeVmTools{}
	return &this
}

// GetCommandLineOptions returns the CommandLineOptions field value if set, zero value otherwise.
func (o *VcenterVmToolsUpgradeVmTools) GetCommandLineOptions() string {
	if o == nil || o.CommandLineOptions == nil {
		var ret string
		return ret
	}
	return *o.CommandLineOptions
}

// GetCommandLineOptionsOk returns a tuple with the CommandLineOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmToolsUpgradeVmTools) GetCommandLineOptionsOk() (*string, bool) {
	if o == nil || o.CommandLineOptions == nil {
		return nil, false
	}
	return o.CommandLineOptions, true
}

// HasCommandLineOptions returns a boolean if a field has been set.
func (o *VcenterVmToolsUpgradeVmTools) HasCommandLineOptions() bool {
	if o != nil && o.CommandLineOptions != nil {
		return true
	}

	return false
}

// SetCommandLineOptions gets a reference to the given string and assigns it to the CommandLineOptions field.
func (o *VcenterVmToolsUpgradeVmTools) SetCommandLineOptions(v string) {
	o.CommandLineOptions = &v
}

func (o VcenterVmToolsUpgradeVmTools) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommandLineOptions != nil {
		toSerialize["command_line_options"] = o.CommandLineOptions
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmToolsUpgradeVmTools struct {
	value *VcenterVmToolsUpgradeVmTools
	isSet bool
}

func (v NullableVcenterVmToolsUpgradeVmTools) Get() *VcenterVmToolsUpgradeVmTools {
	return v.value
}

func (v *NullableVcenterVmToolsUpgradeVmTools) Set(val *VcenterVmToolsUpgradeVmTools) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmToolsUpgradeVmTools) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmToolsUpgradeVmTools) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmToolsUpgradeVmTools(val *VcenterVmToolsUpgradeVmTools) *NullableVcenterVmToolsUpgradeVmTools {
	return &NullableVcenterVmToolsUpgradeVmTools{value: val, isSet: true}
}

func (v NullableVcenterVmToolsUpgradeVmTools) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmToolsUpgradeVmTools) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

