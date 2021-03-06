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

// VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp struct for VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp
type VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp struct {
	Value VcenterVmTemplateLibraryItemsVersionsInfo `json:"value"`
}

// NewVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp instantiates a new VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp(value VcenterVmTemplateLibraryItemsVersionsInfo) *VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp {
	this := VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp{}
	this.Value = value
	return &this
}

// NewVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsRespWithDefaults instantiates a new VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsRespWithDefaults() *VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp {
	this := VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp) GetValue() VcenterVmTemplateLibraryItemsVersionsInfo {
	if o == nil {
		var ret VcenterVmTemplateLibraryItemsVersionsInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp) GetValueOk() (*VcenterVmTemplateLibraryItemsVersionsInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp) SetValue(v VcenterVmTemplateLibraryItemsVersionsInfo) {
	o.Value = v
}

func (o VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp struct {
	value *VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp) Get() *VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp) Set(val *VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp(val *VcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp) *NullableVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp {
	return &NullableVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsVersionsGetVmTemplateLibraryItemsVersionsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


