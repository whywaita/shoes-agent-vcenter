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

// VcenterVchaClusterWitnessInfo struct for VcenterVchaClusterWitnessInfo
type VcenterVchaClusterWitnessInfo struct {
	HaIp VcenterVchaClusterIpInfo `json:"ha_ip"`
	Runtime *VcenterVchaClusterNodeRuntimeInfo `json:"runtime,omitempty"`
}

// NewVcenterVchaClusterWitnessInfo instantiates a new VcenterVchaClusterWitnessInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVchaClusterWitnessInfo(haIp VcenterVchaClusterIpInfo) *VcenterVchaClusterWitnessInfo {
	this := VcenterVchaClusterWitnessInfo{}
	this.HaIp = haIp
	return &this
}

// NewVcenterVchaClusterWitnessInfoWithDefaults instantiates a new VcenterVchaClusterWitnessInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVchaClusterWitnessInfoWithDefaults() *VcenterVchaClusterWitnessInfo {
	this := VcenterVchaClusterWitnessInfo{}
	return &this
}

// GetHaIp returns the HaIp field value
func (o *VcenterVchaClusterWitnessInfo) GetHaIp() VcenterVchaClusterIpInfo {
	if o == nil {
		var ret VcenterVchaClusterIpInfo
		return ret
	}

	return o.HaIp
}

// GetHaIpOk returns a tuple with the HaIp field value
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterWitnessInfo) GetHaIpOk() (*VcenterVchaClusterIpInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HaIp, true
}

// SetHaIp sets field value
func (o *VcenterVchaClusterWitnessInfo) SetHaIp(v VcenterVchaClusterIpInfo) {
	o.HaIp = v
}

// GetRuntime returns the Runtime field value if set, zero value otherwise.
func (o *VcenterVchaClusterWitnessInfo) GetRuntime() VcenterVchaClusterNodeRuntimeInfo {
	if o == nil || o.Runtime == nil {
		var ret VcenterVchaClusterNodeRuntimeInfo
		return ret
	}
	return *o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterWitnessInfo) GetRuntimeOk() (*VcenterVchaClusterNodeRuntimeInfo, bool) {
	if o == nil || o.Runtime == nil {
		return nil, false
	}
	return o.Runtime, true
}

// HasRuntime returns a boolean if a field has been set.
func (o *VcenterVchaClusterWitnessInfo) HasRuntime() bool {
	if o != nil && o.Runtime != nil {
		return true
	}

	return false
}

// SetRuntime gets a reference to the given VcenterVchaClusterNodeRuntimeInfo and assigns it to the Runtime field.
func (o *VcenterVchaClusterWitnessInfo) SetRuntime(v VcenterVchaClusterNodeRuntimeInfo) {
	o.Runtime = &v
}

func (o VcenterVchaClusterWitnessInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ha_ip"] = o.HaIp
	}
	if o.Runtime != nil {
		toSerialize["runtime"] = o.Runtime
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVchaClusterWitnessInfo struct {
	value *VcenterVchaClusterWitnessInfo
	isSet bool
}

func (v NullableVcenterVchaClusterWitnessInfo) Get() *VcenterVchaClusterWitnessInfo {
	return v.value
}

func (v *NullableVcenterVchaClusterWitnessInfo) Set(val *VcenterVchaClusterWitnessInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVchaClusterWitnessInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVchaClusterWitnessInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVchaClusterWitnessInfo(val *VcenterVchaClusterWitnessInfo) *NullableVcenterVchaClusterWitnessInfo {
	return &NullableVcenterVchaClusterWitnessInfo{value: val, isSet: true}
}

func (v NullableVcenterVchaClusterWitnessInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVchaClusterWitnessInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


