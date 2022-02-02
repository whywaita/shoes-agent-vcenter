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

// VcenterInventoryDatastoreInfo struct for VcenterInventoryDatastoreInfo
type VcenterInventoryDatastoreInfo struct {
	// Type of the datastore.
	Type string `json:"type"`
}

// NewVcenterInventoryDatastoreInfo instantiates a new VcenterInventoryDatastoreInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterInventoryDatastoreInfo(type_ string) *VcenterInventoryDatastoreInfo {
	this := VcenterInventoryDatastoreInfo{}
	this.Type = type_
	return &this
}

// NewVcenterInventoryDatastoreInfoWithDefaults instantiates a new VcenterInventoryDatastoreInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterInventoryDatastoreInfoWithDefaults() *VcenterInventoryDatastoreInfo {
	this := VcenterInventoryDatastoreInfo{}
	return &this
}

// GetType returns the Type field value
func (o *VcenterInventoryDatastoreInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterInventoryDatastoreInfo) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterInventoryDatastoreInfo) SetType(v string) {
	o.Type = v
}

func (o VcenterInventoryDatastoreInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterInventoryDatastoreInfo struct {
	value *VcenterInventoryDatastoreInfo
	isSet bool
}

func (v NullableVcenterInventoryDatastoreInfo) Get() *VcenterInventoryDatastoreInfo {
	return v.value
}

func (v *NullableVcenterInventoryDatastoreInfo) Set(val *VcenterInventoryDatastoreInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterInventoryDatastoreInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterInventoryDatastoreInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterInventoryDatastoreInfo(val *VcenterInventoryDatastoreInfo) *NullableVcenterInventoryDatastoreInfo {
	return &NullableVcenterInventoryDatastoreInfo{value: val, isSet: true}
}

func (v NullableVcenterInventoryDatastoreInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterInventoryDatastoreInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


