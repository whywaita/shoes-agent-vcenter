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

// VcenterInventoryDatastoreFindInventoryDatastoreResp struct for VcenterInventoryDatastoreFindInventoryDatastoreResp
type VcenterInventoryDatastoreFindInventoryDatastoreResp struct {
	Value []VcenterInventoryDatastoreFindInventoryDatastoreRespValue `json:"value"`
}

// NewVcenterInventoryDatastoreFindInventoryDatastoreResp instantiates a new VcenterInventoryDatastoreFindInventoryDatastoreResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterInventoryDatastoreFindInventoryDatastoreResp(value []VcenterInventoryDatastoreFindInventoryDatastoreRespValue) *VcenterInventoryDatastoreFindInventoryDatastoreResp {
	this := VcenterInventoryDatastoreFindInventoryDatastoreResp{}
	this.Value = value
	return &this
}

// NewVcenterInventoryDatastoreFindInventoryDatastoreRespWithDefaults instantiates a new VcenterInventoryDatastoreFindInventoryDatastoreResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterInventoryDatastoreFindInventoryDatastoreRespWithDefaults() *VcenterInventoryDatastoreFindInventoryDatastoreResp {
	this := VcenterInventoryDatastoreFindInventoryDatastoreResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterInventoryDatastoreFindInventoryDatastoreResp) GetValue() []VcenterInventoryDatastoreFindInventoryDatastoreRespValue {
	if o == nil {
		var ret []VcenterInventoryDatastoreFindInventoryDatastoreRespValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterInventoryDatastoreFindInventoryDatastoreResp) GetValueOk() (*[]VcenterInventoryDatastoreFindInventoryDatastoreRespValue, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterInventoryDatastoreFindInventoryDatastoreResp) SetValue(v []VcenterInventoryDatastoreFindInventoryDatastoreRespValue) {
	o.Value = v
}

func (o VcenterInventoryDatastoreFindInventoryDatastoreResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterInventoryDatastoreFindInventoryDatastoreResp struct {
	value *VcenterInventoryDatastoreFindInventoryDatastoreResp
	isSet bool
}

func (v NullableVcenterInventoryDatastoreFindInventoryDatastoreResp) Get() *VcenterInventoryDatastoreFindInventoryDatastoreResp {
	return v.value
}

func (v *NullableVcenterInventoryDatastoreFindInventoryDatastoreResp) Set(val *VcenterInventoryDatastoreFindInventoryDatastoreResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterInventoryDatastoreFindInventoryDatastoreResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterInventoryDatastoreFindInventoryDatastoreResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterInventoryDatastoreFindInventoryDatastoreResp(val *VcenterInventoryDatastoreFindInventoryDatastoreResp) *NullableVcenterInventoryDatastoreFindInventoryDatastoreResp {
	return &NullableVcenterInventoryDatastoreFindInventoryDatastoreResp{value: val, isSet: true}
}

func (v NullableVcenterInventoryDatastoreFindInventoryDatastoreResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterInventoryDatastoreFindInventoryDatastoreResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


