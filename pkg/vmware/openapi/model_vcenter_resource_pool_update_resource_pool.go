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

// VcenterResourcePoolUpdateResourcePool struct for VcenterResourcePoolUpdateResourcePool
type VcenterResourcePoolUpdateResourcePool struct {
	Spec *VcenterResourcePoolUpdateSpec `json:"spec,omitempty"`
}

// NewVcenterResourcePoolUpdateResourcePool instantiates a new VcenterResourcePoolUpdateResourcePool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterResourcePoolUpdateResourcePool() *VcenterResourcePoolUpdateResourcePool {
	this := VcenterResourcePoolUpdateResourcePool{}
	return &this
}

// NewVcenterResourcePoolUpdateResourcePoolWithDefaults instantiates a new VcenterResourcePoolUpdateResourcePool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterResourcePoolUpdateResourcePoolWithDefaults() *VcenterResourcePoolUpdateResourcePool {
	this := VcenterResourcePoolUpdateResourcePool{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterResourcePoolUpdateResourcePool) GetSpec() VcenterResourcePoolUpdateSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterResourcePoolUpdateSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolUpdateResourcePool) GetSpecOk() (*VcenterResourcePoolUpdateSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterResourcePoolUpdateResourcePool) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterResourcePoolUpdateSpec and assigns it to the Spec field.
func (o *VcenterResourcePoolUpdateResourcePool) SetSpec(v VcenterResourcePoolUpdateSpec) {
	o.Spec = &v
}

func (o VcenterResourcePoolUpdateResourcePool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterResourcePoolUpdateResourcePool struct {
	value *VcenterResourcePoolUpdateResourcePool
	isSet bool
}

func (v NullableVcenterResourcePoolUpdateResourcePool) Get() *VcenterResourcePoolUpdateResourcePool {
	return v.value
}

func (v *NullableVcenterResourcePoolUpdateResourcePool) Set(val *VcenterResourcePoolUpdateResourcePool) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterResourcePoolUpdateResourcePool) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterResourcePoolUpdateResourcePool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterResourcePoolUpdateResourcePool(val *VcenterResourcePoolUpdateResourcePool) *NullableVcenterResourcePoolUpdateResourcePool {
	return &NullableVcenterResourcePoolUpdateResourcePool{value: val, isSet: true}
}

func (v NullableVcenterResourcePoolUpdateResourcePool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterResourcePoolUpdateResourcePool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


