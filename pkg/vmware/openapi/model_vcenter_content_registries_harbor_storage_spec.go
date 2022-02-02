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

// VcenterContentRegistriesHarborStorageSpec struct for VcenterContentRegistriesHarborStorageSpec
type VcenterContentRegistriesHarborStorageSpec struct {
	// Identifier of the storage policy.
	Policy string `json:"policy"`
	// The maximum amount of storage (in mebibytes) which can be utilized by the registry for this specification.
	Limit *int64 `json:"limit,omitempty"`
}

// NewVcenterContentRegistriesHarborStorageSpec instantiates a new VcenterContentRegistriesHarborStorageSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterContentRegistriesHarborStorageSpec(policy string) *VcenterContentRegistriesHarborStorageSpec {
	this := VcenterContentRegistriesHarborStorageSpec{}
	this.Policy = policy
	return &this
}

// NewVcenterContentRegistriesHarborStorageSpecWithDefaults instantiates a new VcenterContentRegistriesHarborStorageSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterContentRegistriesHarborStorageSpecWithDefaults() *VcenterContentRegistriesHarborStorageSpec {
	this := VcenterContentRegistriesHarborStorageSpec{}
	return &this
}

// GetPolicy returns the Policy field value
func (o *VcenterContentRegistriesHarborStorageSpec) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborStorageSpec) GetPolicyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *VcenterContentRegistriesHarborStorageSpec) SetPolicy(v string) {
	o.Policy = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *VcenterContentRegistriesHarborStorageSpec) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborStorageSpec) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *VcenterContentRegistriesHarborStorageSpec) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *VcenterContentRegistriesHarborStorageSpec) SetLimit(v int64) {
	o.Limit = &v
}

func (o VcenterContentRegistriesHarborStorageSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["policy"] = o.Policy
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterContentRegistriesHarborStorageSpec struct {
	value *VcenterContentRegistriesHarborStorageSpec
	isSet bool
}

func (v NullableVcenterContentRegistriesHarborStorageSpec) Get() *VcenterContentRegistriesHarborStorageSpec {
	return v.value
}

func (v *NullableVcenterContentRegistriesHarborStorageSpec) Set(val *VcenterContentRegistriesHarborStorageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterContentRegistriesHarborStorageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterContentRegistriesHarborStorageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterContentRegistriesHarborStorageSpec(val *VcenterContentRegistriesHarborStorageSpec) *NullableVcenterContentRegistriesHarborStorageSpec {
	return &NullableVcenterContentRegistriesHarborStorageSpec{value: val, isSet: true}
}

func (v NullableVcenterContentRegistriesHarborStorageSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterContentRegistriesHarborStorageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


