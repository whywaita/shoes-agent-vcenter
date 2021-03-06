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

// VcenterNetworkSummary struct for VcenterNetworkSummary
type VcenterNetworkSummary struct {
	// Identifier of the network. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network.
	Network string `json:"network"`
	// Name of the network.
	Name string `json:"name"`
	Type VcenterNetworkType `json:"type"`
}

// NewVcenterNetworkSummary instantiates a new VcenterNetworkSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNetworkSummary(network string, name string, type_ VcenterNetworkType) *VcenterNetworkSummary {
	this := VcenterNetworkSummary{}
	this.Network = network
	this.Name = name
	this.Type = type_
	return &this
}

// NewVcenterNetworkSummaryWithDefaults instantiates a new VcenterNetworkSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNetworkSummaryWithDefaults() *VcenterNetworkSummary {
	this := VcenterNetworkSummary{}
	return &this
}

// GetNetwork returns the Network field value
func (o *VcenterNetworkSummary) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *VcenterNetworkSummary) GetNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *VcenterNetworkSummary) SetNetwork(v string) {
	o.Network = v
}

// GetName returns the Name field value
func (o *VcenterNetworkSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VcenterNetworkSummary) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VcenterNetworkSummary) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *VcenterNetworkSummary) GetType() VcenterNetworkType {
	if o == nil {
		var ret VcenterNetworkType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterNetworkSummary) GetTypeOk() (*VcenterNetworkType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterNetworkSummary) SetType(v VcenterNetworkType) {
	o.Type = v
}

func (o VcenterNetworkSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNetworkSummary struct {
	value *VcenterNetworkSummary
	isSet bool
}

func (v NullableVcenterNetworkSummary) Get() *VcenterNetworkSummary {
	return v.value
}

func (v *NullableVcenterNetworkSummary) Set(val *VcenterNetworkSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNetworkSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNetworkSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNetworkSummary(val *VcenterNetworkSummary) *NullableVcenterNetworkSummary {
	return &NullableVcenterNetworkSummary{value: val, isSet: true}
}

func (v NullableVcenterNetworkSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNetworkSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


