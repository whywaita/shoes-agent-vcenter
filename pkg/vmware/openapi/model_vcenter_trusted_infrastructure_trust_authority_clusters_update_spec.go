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

// VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec struct for VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec
type VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec struct {
	State *VcenterTrustedInfrastructureTrustAuthorityClustersState `json:"state,omitempty"`
}

// NewVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec() *VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec {
	this := VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec{}
	return &this
}

// NewVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec {
	this := VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec) GetState() VcenterTrustedInfrastructureTrustAuthorityClustersState {
	if o == nil || o.State == nil {
		var ret VcenterTrustedInfrastructureTrustAuthorityClustersState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec) GetStateOk() (*VcenterTrustedInfrastructureTrustAuthorityClustersState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given VcenterTrustedInfrastructureTrustAuthorityClustersState and assigns it to the State field.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec) SetState(v VcenterTrustedInfrastructureTrustAuthorityClustersState) {
	o.State = &v
}

func (o VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec struct {
	value *VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec) Get() *VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec) Set(val *VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec(val *VcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec) *NullableVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec {
	return &NullableVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

