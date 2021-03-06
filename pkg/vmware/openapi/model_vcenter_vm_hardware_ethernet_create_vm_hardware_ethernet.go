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

// VcenterVmHardwareEthernetCreateVmHardwareEthernet struct for VcenterVmHardwareEthernetCreateVmHardwareEthernet
type VcenterVmHardwareEthernetCreateVmHardwareEthernet struct {
	Spec *VcenterVmHardwareEthernetCreateSpec `json:"spec,omitempty"`
}

// NewVcenterVmHardwareEthernetCreateVmHardwareEthernet instantiates a new VcenterVmHardwareEthernetCreateVmHardwareEthernet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareEthernetCreateVmHardwareEthernet() *VcenterVmHardwareEthernetCreateVmHardwareEthernet {
	this := VcenterVmHardwareEthernetCreateVmHardwareEthernet{}
	return &this
}

// NewVcenterVmHardwareEthernetCreateVmHardwareEthernetWithDefaults instantiates a new VcenterVmHardwareEthernetCreateVmHardwareEthernet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareEthernetCreateVmHardwareEthernetWithDefaults() *VcenterVmHardwareEthernetCreateVmHardwareEthernet {
	this := VcenterVmHardwareEthernetCreateVmHardwareEthernet{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterVmHardwareEthernetCreateVmHardwareEthernet) GetSpec() VcenterVmHardwareEthernetCreateSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterVmHardwareEthernetCreateSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetCreateVmHardwareEthernet) GetSpecOk() (*VcenterVmHardwareEthernetCreateSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterVmHardwareEthernetCreateVmHardwareEthernet) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterVmHardwareEthernetCreateSpec and assigns it to the Spec field.
func (o *VcenterVmHardwareEthernetCreateVmHardwareEthernet) SetSpec(v VcenterVmHardwareEthernetCreateSpec) {
	o.Spec = &v
}

func (o VcenterVmHardwareEthernetCreateVmHardwareEthernet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareEthernetCreateVmHardwareEthernet struct {
	value *VcenterVmHardwareEthernetCreateVmHardwareEthernet
	isSet bool
}

func (v NullableVcenterVmHardwareEthernetCreateVmHardwareEthernet) Get() *VcenterVmHardwareEthernetCreateVmHardwareEthernet {
	return v.value
}

func (v *NullableVcenterVmHardwareEthernetCreateVmHardwareEthernet) Set(val *VcenterVmHardwareEthernetCreateVmHardwareEthernet) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareEthernetCreateVmHardwareEthernet) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareEthernetCreateVmHardwareEthernet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareEthernetCreateVmHardwareEthernet(val *VcenterVmHardwareEthernetCreateVmHardwareEthernet) *NullableVcenterVmHardwareEthernetCreateVmHardwareEthernet {
	return &NullableVcenterVmHardwareEthernetCreateVmHardwareEthernet{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareEthernetCreateVmHardwareEthernet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareEthernetCreateVmHardwareEthernet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


