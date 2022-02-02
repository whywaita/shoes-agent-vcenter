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

// VcenterVchaClusterPassiveRedeploySpec struct for VcenterVchaClusterPassiveRedeploySpec
type VcenterVchaClusterPassiveRedeploySpec struct {
	VcSpec *VcenterVchaCredentialsSpec `json:"vc_spec,omitempty"`
	Placement VcenterVchaPlacementSpec `json:"placement"`
	HaIp *VcenterVchaIpSpec `json:"ha_ip,omitempty"`
	FailoverIp *VcenterVchaIpSpec `json:"failover_ip,omitempty"`
}

// NewVcenterVchaClusterPassiveRedeploySpec instantiates a new VcenterVchaClusterPassiveRedeploySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVchaClusterPassiveRedeploySpec(placement VcenterVchaPlacementSpec) *VcenterVchaClusterPassiveRedeploySpec {
	this := VcenterVchaClusterPassiveRedeploySpec{}
	this.Placement = placement
	return &this
}

// NewVcenterVchaClusterPassiveRedeploySpecWithDefaults instantiates a new VcenterVchaClusterPassiveRedeploySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVchaClusterPassiveRedeploySpecWithDefaults() *VcenterVchaClusterPassiveRedeploySpec {
	this := VcenterVchaClusterPassiveRedeploySpec{}
	return &this
}

// GetVcSpec returns the VcSpec field value if set, zero value otherwise.
func (o *VcenterVchaClusterPassiveRedeploySpec) GetVcSpec() VcenterVchaCredentialsSpec {
	if o == nil || o.VcSpec == nil {
		var ret VcenterVchaCredentialsSpec
		return ret
	}
	return *o.VcSpec
}

// GetVcSpecOk returns a tuple with the VcSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterPassiveRedeploySpec) GetVcSpecOk() (*VcenterVchaCredentialsSpec, bool) {
	if o == nil || o.VcSpec == nil {
		return nil, false
	}
	return o.VcSpec, true
}

// HasVcSpec returns a boolean if a field has been set.
func (o *VcenterVchaClusterPassiveRedeploySpec) HasVcSpec() bool {
	if o != nil && o.VcSpec != nil {
		return true
	}

	return false
}

// SetVcSpec gets a reference to the given VcenterVchaCredentialsSpec and assigns it to the VcSpec field.
func (o *VcenterVchaClusterPassiveRedeploySpec) SetVcSpec(v VcenterVchaCredentialsSpec) {
	o.VcSpec = &v
}

// GetPlacement returns the Placement field value
func (o *VcenterVchaClusterPassiveRedeploySpec) GetPlacement() VcenterVchaPlacementSpec {
	if o == nil {
		var ret VcenterVchaPlacementSpec
		return ret
	}

	return o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterPassiveRedeploySpec) GetPlacementOk() (*VcenterVchaPlacementSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Placement, true
}

// SetPlacement sets field value
func (o *VcenterVchaClusterPassiveRedeploySpec) SetPlacement(v VcenterVchaPlacementSpec) {
	o.Placement = v
}

// GetHaIp returns the HaIp field value if set, zero value otherwise.
func (o *VcenterVchaClusterPassiveRedeploySpec) GetHaIp() VcenterVchaIpSpec {
	if o == nil || o.HaIp == nil {
		var ret VcenterVchaIpSpec
		return ret
	}
	return *o.HaIp
}

// GetHaIpOk returns a tuple with the HaIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterPassiveRedeploySpec) GetHaIpOk() (*VcenterVchaIpSpec, bool) {
	if o == nil || o.HaIp == nil {
		return nil, false
	}
	return o.HaIp, true
}

// HasHaIp returns a boolean if a field has been set.
func (o *VcenterVchaClusterPassiveRedeploySpec) HasHaIp() bool {
	if o != nil && o.HaIp != nil {
		return true
	}

	return false
}

// SetHaIp gets a reference to the given VcenterVchaIpSpec and assigns it to the HaIp field.
func (o *VcenterVchaClusterPassiveRedeploySpec) SetHaIp(v VcenterVchaIpSpec) {
	o.HaIp = &v
}

// GetFailoverIp returns the FailoverIp field value if set, zero value otherwise.
func (o *VcenterVchaClusterPassiveRedeploySpec) GetFailoverIp() VcenterVchaIpSpec {
	if o == nil || o.FailoverIp == nil {
		var ret VcenterVchaIpSpec
		return ret
	}
	return *o.FailoverIp
}

// GetFailoverIpOk returns a tuple with the FailoverIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterPassiveRedeploySpec) GetFailoverIpOk() (*VcenterVchaIpSpec, bool) {
	if o == nil || o.FailoverIp == nil {
		return nil, false
	}
	return o.FailoverIp, true
}

// HasFailoverIp returns a boolean if a field has been set.
func (o *VcenterVchaClusterPassiveRedeploySpec) HasFailoverIp() bool {
	if o != nil && o.FailoverIp != nil {
		return true
	}

	return false
}

// SetFailoverIp gets a reference to the given VcenterVchaIpSpec and assigns it to the FailoverIp field.
func (o *VcenterVchaClusterPassiveRedeploySpec) SetFailoverIp(v VcenterVchaIpSpec) {
	o.FailoverIp = &v
}

func (o VcenterVchaClusterPassiveRedeploySpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VcSpec != nil {
		toSerialize["vc_spec"] = o.VcSpec
	}
	if true {
		toSerialize["placement"] = o.Placement
	}
	if o.HaIp != nil {
		toSerialize["ha_ip"] = o.HaIp
	}
	if o.FailoverIp != nil {
		toSerialize["failover_ip"] = o.FailoverIp
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVchaClusterPassiveRedeploySpec struct {
	value *VcenterVchaClusterPassiveRedeploySpec
	isSet bool
}

func (v NullableVcenterVchaClusterPassiveRedeploySpec) Get() *VcenterVchaClusterPassiveRedeploySpec {
	return v.value
}

func (v *NullableVcenterVchaClusterPassiveRedeploySpec) Set(val *VcenterVchaClusterPassiveRedeploySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVchaClusterPassiveRedeploySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVchaClusterPassiveRedeploySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVchaClusterPassiveRedeploySpec(val *VcenterVchaClusterPassiveRedeploySpec) *NullableVcenterVchaClusterPassiveRedeploySpec {
	return &NullableVcenterVchaClusterPassiveRedeploySpec{value: val, isSet: true}
}

func (v NullableVcenterVchaClusterPassiveRedeploySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVchaClusterPassiveRedeploySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


