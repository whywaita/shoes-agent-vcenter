/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// VcenterVmTemplateLibraryItemsCheckOutsInfo struct for VcenterVmTemplateLibraryItemsCheckOutsInfo
type VcenterVmTemplateLibraryItemsCheckOutsInfo struct {
	// Date and time when the virtual machine was checked out.
	Time time.Time `json:"time"`
	// Name of the user who checked out the virtual machine.
	User string `json:"user"`
}

// NewVcenterVmTemplateLibraryItemsCheckOutsInfo instantiates a new VcenterVmTemplateLibraryItemsCheckOutsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmTemplateLibraryItemsCheckOutsInfo(time time.Time, user string) *VcenterVmTemplateLibraryItemsCheckOutsInfo {
	this := VcenterVmTemplateLibraryItemsCheckOutsInfo{}
	this.Time = time
	this.User = user
	return &this
}

// NewVcenterVmTemplateLibraryItemsCheckOutsInfoWithDefaults instantiates a new VcenterVmTemplateLibraryItemsCheckOutsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmTemplateLibraryItemsCheckOutsInfoWithDefaults() *VcenterVmTemplateLibraryItemsCheckOutsInfo {
	this := VcenterVmTemplateLibraryItemsCheckOutsInfo{}
	return &this
}

// GetTime returns the Time field value
func (o *VcenterVmTemplateLibraryItemsCheckOutsInfo) GetTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCheckOutsInfo) GetTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *VcenterVmTemplateLibraryItemsCheckOutsInfo) SetTime(v time.Time) {
	o.Time = v
}

// GetUser returns the User field value
func (o *VcenterVmTemplateLibraryItemsCheckOutsInfo) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCheckOutsInfo) GetUserOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *VcenterVmTemplateLibraryItemsCheckOutsInfo) SetUser(v string) {
	o.User = v
}

func (o VcenterVmTemplateLibraryItemsCheckOutsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmTemplateLibraryItemsCheckOutsInfo struct {
	value *VcenterVmTemplateLibraryItemsCheckOutsInfo
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsCheckOutsInfo) Get() *VcenterVmTemplateLibraryItemsCheckOutsInfo {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsCheckOutsInfo) Set(val *VcenterVmTemplateLibraryItemsCheckOutsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsCheckOutsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsCheckOutsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsCheckOutsInfo(val *VcenterVmTemplateLibraryItemsCheckOutsInfo) *NullableVcenterVmTemplateLibraryItemsCheckOutsInfo {
	return &NullableVcenterVmTemplateLibraryItemsCheckOutsInfo{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsCheckOutsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsCheckOutsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


