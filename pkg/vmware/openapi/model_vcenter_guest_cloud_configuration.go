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

// VcenterGuestCloudConfiguration struct for VcenterGuestCloudConfiguration
type VcenterGuestCloudConfiguration struct {
	Type VcenterGuestCloudConfigurationType `json:"type"`
	Cloudinit *VcenterGuestCloudinitConfiguration `json:"cloudinit,omitempty"`
}

// NewVcenterGuestCloudConfiguration instantiates a new VcenterGuestCloudConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterGuestCloudConfiguration(type_ VcenterGuestCloudConfigurationType) *VcenterGuestCloudConfiguration {
	this := VcenterGuestCloudConfiguration{}
	this.Type = type_
	return &this
}

// NewVcenterGuestCloudConfigurationWithDefaults instantiates a new VcenterGuestCloudConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterGuestCloudConfigurationWithDefaults() *VcenterGuestCloudConfiguration {
	this := VcenterGuestCloudConfiguration{}
	return &this
}

// GetType returns the Type field value
func (o *VcenterGuestCloudConfiguration) GetType() VcenterGuestCloudConfigurationType {
	if o == nil {
		var ret VcenterGuestCloudConfigurationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestCloudConfiguration) GetTypeOk() (*VcenterGuestCloudConfigurationType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterGuestCloudConfiguration) SetType(v VcenterGuestCloudConfigurationType) {
	o.Type = v
}

// GetCloudinit returns the Cloudinit field value if set, zero value otherwise.
func (o *VcenterGuestCloudConfiguration) GetCloudinit() VcenterGuestCloudinitConfiguration {
	if o == nil || o.Cloudinit == nil {
		var ret VcenterGuestCloudinitConfiguration
		return ret
	}
	return *o.Cloudinit
}

// GetCloudinitOk returns a tuple with the Cloudinit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterGuestCloudConfiguration) GetCloudinitOk() (*VcenterGuestCloudinitConfiguration, bool) {
	if o == nil || o.Cloudinit == nil {
		return nil, false
	}
	return o.Cloudinit, true
}

// HasCloudinit returns a boolean if a field has been set.
func (o *VcenterGuestCloudConfiguration) HasCloudinit() bool {
	if o != nil && o.Cloudinit != nil {
		return true
	}

	return false
}

// SetCloudinit gets a reference to the given VcenterGuestCloudinitConfiguration and assigns it to the Cloudinit field.
func (o *VcenterGuestCloudConfiguration) SetCloudinit(v VcenterGuestCloudinitConfiguration) {
	o.Cloudinit = &v
}

func (o VcenterGuestCloudConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Cloudinit != nil {
		toSerialize["cloudinit"] = o.Cloudinit
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterGuestCloudConfiguration struct {
	value *VcenterGuestCloudConfiguration
	isSet bool
}

func (v NullableVcenterGuestCloudConfiguration) Get() *VcenterGuestCloudConfiguration {
	return v.value
}

func (v *NullableVcenterGuestCloudConfiguration) Set(val *VcenterGuestCloudConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterGuestCloudConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterGuestCloudConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterGuestCloudConfiguration(val *VcenterGuestCloudConfiguration) *NullableVcenterGuestCloudConfiguration {
	return &NullableVcenterGuestCloudConfiguration{value: val, isSet: true}
}

func (v NullableVcenterGuestCloudConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterGuestCloudConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


