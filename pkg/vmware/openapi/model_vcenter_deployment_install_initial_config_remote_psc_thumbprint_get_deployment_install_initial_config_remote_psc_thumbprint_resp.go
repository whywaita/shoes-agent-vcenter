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

// VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp struct for VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp
type VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp struct {
	Value string `json:"value"`
}

// NewVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp instantiates a new VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp(value string) *VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp {
	this := VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp{}
	this.Value = value
	return &this
}

// NewVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintRespWithDefaults instantiates a new VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintRespWithDefaults() *VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp {
	this := VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp) SetValue(v string) {
	o.Value = v
}

func (o VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp struct {
	value *VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp
	isSet bool
}

func (v NullableVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp) Get() *VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp {
	return v.value
}

func (v *NullableVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp) Set(val *VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp(val *VcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp) *NullableVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp {
	return &NullableVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp{value: val, isSet: true}
}

func (v NullableVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentInstallInitialConfigRemotePscThumbprintGetDeploymentInstallInitialConfigRemotePscThumbprintResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


