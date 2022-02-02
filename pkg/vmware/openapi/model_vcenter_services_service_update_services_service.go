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

// VcenterServicesServiceUpdateServicesService struct for VcenterServicesServiceUpdateServicesService
type VcenterServicesServiceUpdateServicesService struct {
	Spec *VcenterServicesServiceUpdateSpec `json:"spec,omitempty"`
}

// NewVcenterServicesServiceUpdateServicesService instantiates a new VcenterServicesServiceUpdateServicesService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterServicesServiceUpdateServicesService() *VcenterServicesServiceUpdateServicesService {
	this := VcenterServicesServiceUpdateServicesService{}
	return &this
}

// NewVcenterServicesServiceUpdateServicesServiceWithDefaults instantiates a new VcenterServicesServiceUpdateServicesService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterServicesServiceUpdateServicesServiceWithDefaults() *VcenterServicesServiceUpdateServicesService {
	this := VcenterServicesServiceUpdateServicesService{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterServicesServiceUpdateServicesService) GetSpec() VcenterServicesServiceUpdateSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterServicesServiceUpdateSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterServicesServiceUpdateServicesService) GetSpecOk() (*VcenterServicesServiceUpdateSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterServicesServiceUpdateServicesService) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterServicesServiceUpdateSpec and assigns it to the Spec field.
func (o *VcenterServicesServiceUpdateServicesService) SetSpec(v VcenterServicesServiceUpdateSpec) {
	o.Spec = &v
}

func (o VcenterServicesServiceUpdateServicesService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterServicesServiceUpdateServicesService struct {
	value *VcenterServicesServiceUpdateServicesService
	isSet bool
}

func (v NullableVcenterServicesServiceUpdateServicesService) Get() *VcenterServicesServiceUpdateServicesService {
	return v.value
}

func (v *NullableVcenterServicesServiceUpdateServicesService) Set(val *VcenterServicesServiceUpdateServicesService) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterServicesServiceUpdateServicesService) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterServicesServiceUpdateServicesService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterServicesServiceUpdateServicesService(val *VcenterServicesServiceUpdateServicesService) *NullableVcenterServicesServiceUpdateServicesService {
	return &NullableVcenterServicesServiceUpdateServicesService{value: val, isSet: true}
}

func (v NullableVcenterServicesServiceUpdateServicesService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterServicesServiceUpdateServicesService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

