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

// VcenterGuestWindowsSysprep struct for VcenterGuestWindowsSysprep
type VcenterGuestWindowsSysprep struct {
	// A list of commands to run at first user logon, after customizing the guest. These commands are directly mapped to the GuiRunOnce key in the sysprep.xml answer file. These commands are transferred into the sysprep.xml file that VirtualCenter stores on the target virtual disk. For more information about performing unattended installation, check https://technet.microsoft.com/en-us/library/cc771830(v=ws.10).aspx The commands listed here ar executed when a user logs on the first time after customization completes. The logon may be driven by GuiUnattended.auto-logon setting. These commands are directly mapped to the GuiRunOnce key in the If unset, no commands are run.
	GuiRunOnceCommands *[]string `json:"gui_run_once_commands,omitempty"`
	UserData VcenterGuestUserData `json:"user_data"`
	Domain *VcenterGuestDomain `json:"domain,omitempty"`
	GuiUnattended VcenterGuestGuiUnattended `json:"gui_unattended"`
}

// NewVcenterGuestWindowsSysprep instantiates a new VcenterGuestWindowsSysprep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterGuestWindowsSysprep(userData VcenterGuestUserData, guiUnattended VcenterGuestGuiUnattended) *VcenterGuestWindowsSysprep {
	this := VcenterGuestWindowsSysprep{}
	this.UserData = userData
	this.GuiUnattended = guiUnattended
	return &this
}

// NewVcenterGuestWindowsSysprepWithDefaults instantiates a new VcenterGuestWindowsSysprep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterGuestWindowsSysprepWithDefaults() *VcenterGuestWindowsSysprep {
	this := VcenterGuestWindowsSysprep{}
	return &this
}

// GetGuiRunOnceCommands returns the GuiRunOnceCommands field value if set, zero value otherwise.
func (o *VcenterGuestWindowsSysprep) GetGuiRunOnceCommands() []string {
	if o == nil || o.GuiRunOnceCommands == nil {
		var ret []string
		return ret
	}
	return *o.GuiRunOnceCommands
}

// GetGuiRunOnceCommandsOk returns a tuple with the GuiRunOnceCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterGuestWindowsSysprep) GetGuiRunOnceCommandsOk() (*[]string, bool) {
	if o == nil || o.GuiRunOnceCommands == nil {
		return nil, false
	}
	return o.GuiRunOnceCommands, true
}

// HasGuiRunOnceCommands returns a boolean if a field has been set.
func (o *VcenterGuestWindowsSysprep) HasGuiRunOnceCommands() bool {
	if o != nil && o.GuiRunOnceCommands != nil {
		return true
	}

	return false
}

// SetGuiRunOnceCommands gets a reference to the given []string and assigns it to the GuiRunOnceCommands field.
func (o *VcenterGuestWindowsSysprep) SetGuiRunOnceCommands(v []string) {
	o.GuiRunOnceCommands = &v
}

// GetUserData returns the UserData field value
func (o *VcenterGuestWindowsSysprep) GetUserData() VcenterGuestUserData {
	if o == nil {
		var ret VcenterGuestUserData
		return ret
	}

	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestWindowsSysprep) GetUserDataOk() (*VcenterGuestUserData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserData, true
}

// SetUserData sets field value
func (o *VcenterGuestWindowsSysprep) SetUserData(v VcenterGuestUserData) {
	o.UserData = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *VcenterGuestWindowsSysprep) GetDomain() VcenterGuestDomain {
	if o == nil || o.Domain == nil {
		var ret VcenterGuestDomain
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterGuestWindowsSysprep) GetDomainOk() (*VcenterGuestDomain, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *VcenterGuestWindowsSysprep) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given VcenterGuestDomain and assigns it to the Domain field.
func (o *VcenterGuestWindowsSysprep) SetDomain(v VcenterGuestDomain) {
	o.Domain = &v
}

// GetGuiUnattended returns the GuiUnattended field value
func (o *VcenterGuestWindowsSysprep) GetGuiUnattended() VcenterGuestGuiUnattended {
	if o == nil {
		var ret VcenterGuestGuiUnattended
		return ret
	}

	return o.GuiUnattended
}

// GetGuiUnattendedOk returns a tuple with the GuiUnattended field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestWindowsSysprep) GetGuiUnattendedOk() (*VcenterGuestGuiUnattended, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GuiUnattended, true
}

// SetGuiUnattended sets field value
func (o *VcenterGuestWindowsSysprep) SetGuiUnattended(v VcenterGuestGuiUnattended) {
	o.GuiUnattended = v
}

func (o VcenterGuestWindowsSysprep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GuiRunOnceCommands != nil {
		toSerialize["gui_run_once_commands"] = o.GuiRunOnceCommands
	}
	if true {
		toSerialize["user_data"] = o.UserData
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["gui_unattended"] = o.GuiUnattended
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterGuestWindowsSysprep struct {
	value *VcenterGuestWindowsSysprep
	isSet bool
}

func (v NullableVcenterGuestWindowsSysprep) Get() *VcenterGuestWindowsSysprep {
	return v.value
}

func (v *NullableVcenterGuestWindowsSysprep) Set(val *VcenterGuestWindowsSysprep) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterGuestWindowsSysprep) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterGuestWindowsSysprep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterGuestWindowsSysprep(val *VcenterGuestWindowsSysprep) *NullableVcenterGuestWindowsSysprep {
	return &NullableVcenterGuestWindowsSysprep{value: val, isSet: true}
}

func (v NullableVcenterGuestWindowsSysprep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterGuestWindowsSysprep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

