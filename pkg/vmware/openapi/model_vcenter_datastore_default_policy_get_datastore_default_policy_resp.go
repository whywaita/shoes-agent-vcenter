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

// VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp struct for VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp
type VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp struct {
	Value string `json:"value"`
}

// NewVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp instantiates a new VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp(value string) *VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp {
	this := VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp{}
	this.Value = value
	return &this
}

// NewVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyRespWithDefaults instantiates a new VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyRespWithDefaults() *VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp {
	this := VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp) SetValue(v string) {
	o.Value = v
}

func (o VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp struct {
	value *VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp
	isSet bool
}

func (v NullableVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp) Get() *VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp {
	return v.value
}

func (v *NullableVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp) Set(val *VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp(val *VcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp) *NullableVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp {
	return &NullableVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp{value: val, isSet: true}
}

func (v NullableVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDatastoreDefaultPolicyGetDatastoreDefaultPolicyResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

