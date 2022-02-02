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

// VcenterVchaClusterPassiveSpec struct for VcenterVchaClusterPassiveSpec
type VcenterVchaClusterPassiveSpec struct {
	Placement *VcenterVchaPlacementSpec `json:"placement,omitempty"`
	HaIp VcenterVchaIpSpec `json:"ha_ip"`
	FailoverIp *VcenterVchaIpSpec `json:"failover_ip,omitempty"`
}

// NewVcenterVchaClusterPassiveSpec instantiates a new VcenterVchaClusterPassiveSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVchaClusterPassiveSpec(haIp VcenterVchaIpSpec) *VcenterVchaClusterPassiveSpec {
	this := VcenterVchaClusterPassiveSpec{}
	this.HaIp = haIp
	return &this
}

// NewVcenterVchaClusterPassiveSpecWithDefaults instantiates a new VcenterVchaClusterPassiveSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVchaClusterPassiveSpecWithDefaults() *VcenterVchaClusterPassiveSpec {
	this := VcenterVchaClusterPassiveSpec{}
	return &this
}

// GetPlacement returns the Placement field value if set, zero value otherwise.
func (o *VcenterVchaClusterPassiveSpec) GetPlacement() VcenterVchaPlacementSpec {
	if o == nil || o.Placement == nil {
		var ret VcenterVchaPlacementSpec
		return ret
	}
	return *o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterPassiveSpec) GetPlacementOk() (*VcenterVchaPlacementSpec, bool) {
	if o == nil || o.Placement == nil {
		return nil, false
	}
	return o.Placement, true
}

// HasPlacement returns a boolean if a field has been set.
func (o *VcenterVchaClusterPassiveSpec) HasPlacement() bool {
	if o != nil && o.Placement != nil {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given VcenterVchaPlacementSpec and assigns it to the Placement field.
func (o *VcenterVchaClusterPassiveSpec) SetPlacement(v VcenterVchaPlacementSpec) {
	o.Placement = &v
}

// GetHaIp returns the HaIp field value
func (o *VcenterVchaClusterPassiveSpec) GetHaIp() VcenterVchaIpSpec {
	if o == nil {
		var ret VcenterVchaIpSpec
		return ret
	}

	return o.HaIp
}

// GetHaIpOk returns a tuple with the HaIp field value
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterPassiveSpec) GetHaIpOk() (*VcenterVchaIpSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HaIp, true
}

// SetHaIp sets field value
func (o *VcenterVchaClusterPassiveSpec) SetHaIp(v VcenterVchaIpSpec) {
	o.HaIp = v
}

// GetFailoverIp returns the FailoverIp field value if set, zero value otherwise.
func (o *VcenterVchaClusterPassiveSpec) GetFailoverIp() VcenterVchaIpSpec {
	if o == nil || o.FailoverIp == nil {
		var ret VcenterVchaIpSpec
		return ret
	}
	return *o.FailoverIp
}

// GetFailoverIpOk returns a tuple with the FailoverIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterPassiveSpec) GetFailoverIpOk() (*VcenterVchaIpSpec, bool) {
	if o == nil || o.FailoverIp == nil {
		return nil, false
	}
	return o.FailoverIp, true
}

// HasFailoverIp returns a boolean if a field has been set.
func (o *VcenterVchaClusterPassiveSpec) HasFailoverIp() bool {
	if o != nil && o.FailoverIp != nil {
		return true
	}

	return false
}

// SetFailoverIp gets a reference to the given VcenterVchaIpSpec and assigns it to the FailoverIp field.
func (o *VcenterVchaClusterPassiveSpec) SetFailoverIp(v VcenterVchaIpSpec) {
	o.FailoverIp = &v
}

func (o VcenterVchaClusterPassiveSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Placement != nil {
		toSerialize["placement"] = o.Placement
	}
	if true {
		toSerialize["ha_ip"] = o.HaIp
	}
	if o.FailoverIp != nil {
		toSerialize["failover_ip"] = o.FailoverIp
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVchaClusterPassiveSpec struct {
	value *VcenterVchaClusterPassiveSpec
	isSet bool
}

func (v NullableVcenterVchaClusterPassiveSpec) Get() *VcenterVchaClusterPassiveSpec {
	return v.value
}

func (v *NullableVcenterVchaClusterPassiveSpec) Set(val *VcenterVchaClusterPassiveSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVchaClusterPassiveSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVchaClusterPassiveSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVchaClusterPassiveSpec(val *VcenterVchaClusterPassiveSpec) *NullableVcenterVchaClusterPassiveSpec {
	return &NullableVcenterVchaClusterPassiveSpec{value: val, isSet: true}
}

func (v NullableVcenterVchaClusterPassiveSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVchaClusterPassiveSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


