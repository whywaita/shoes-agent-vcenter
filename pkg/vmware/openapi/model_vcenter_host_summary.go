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

// VcenterHostSummary struct for VcenterHostSummary
type VcenterHostSummary struct {
	// Identifier of the host. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: HostSystem. When operations return a value of this structure as a result, the field will be an identifier for the resource type: HostSystem.
	Host string `json:"host"`
	// Name of the host.
	Name string `json:"name"`
	ConnectionState VcenterHostConnectionState `json:"connection_state"`
	PowerState *VcenterHostPowerState `json:"power_state,omitempty"`
}

// NewVcenterHostSummary instantiates a new VcenterHostSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterHostSummary(host string, name string, connectionState VcenterHostConnectionState) *VcenterHostSummary {
	this := VcenterHostSummary{}
	this.Host = host
	this.Name = name
	this.ConnectionState = connectionState
	return &this
}

// NewVcenterHostSummaryWithDefaults instantiates a new VcenterHostSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterHostSummaryWithDefaults() *VcenterHostSummary {
	this := VcenterHostSummary{}
	return &this
}

// GetHost returns the Host field value
func (o *VcenterHostSummary) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *VcenterHostSummary) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *VcenterHostSummary) SetHost(v string) {
	o.Host = v
}

// GetName returns the Name field value
func (o *VcenterHostSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VcenterHostSummary) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VcenterHostSummary) SetName(v string) {
	o.Name = v
}

// GetConnectionState returns the ConnectionState field value
func (o *VcenterHostSummary) GetConnectionState() VcenterHostConnectionState {
	if o == nil {
		var ret VcenterHostConnectionState
		return ret
	}

	return o.ConnectionState
}

// GetConnectionStateOk returns a tuple with the ConnectionState field value
// and a boolean to check if the value has been set.
func (o *VcenterHostSummary) GetConnectionStateOk() (*VcenterHostConnectionState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConnectionState, true
}

// SetConnectionState sets field value
func (o *VcenterHostSummary) SetConnectionState(v VcenterHostConnectionState) {
	o.ConnectionState = v
}

// GetPowerState returns the PowerState field value if set, zero value otherwise.
func (o *VcenterHostSummary) GetPowerState() VcenterHostPowerState {
	if o == nil || o.PowerState == nil {
		var ret VcenterHostPowerState
		return ret
	}
	return *o.PowerState
}

// GetPowerStateOk returns a tuple with the PowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterHostSummary) GetPowerStateOk() (*VcenterHostPowerState, bool) {
	if o == nil || o.PowerState == nil {
		return nil, false
	}
	return o.PowerState, true
}

// HasPowerState returns a boolean if a field has been set.
func (o *VcenterHostSummary) HasPowerState() bool {
	if o != nil && o.PowerState != nil {
		return true
	}

	return false
}

// SetPowerState gets a reference to the given VcenterHostPowerState and assigns it to the PowerState field.
func (o *VcenterHostSummary) SetPowerState(v VcenterHostPowerState) {
	o.PowerState = &v
}

func (o VcenterHostSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["connection_state"] = o.ConnectionState
	}
	if o.PowerState != nil {
		toSerialize["power_state"] = o.PowerState
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterHostSummary struct {
	value *VcenterHostSummary
	isSet bool
}

func (v NullableVcenterHostSummary) Get() *VcenterHostSummary {
	return v.value
}

func (v *NullableVcenterHostSummary) Set(val *VcenterHostSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterHostSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterHostSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterHostSummary(val *VcenterHostSummary) *NullableVcenterHostSummary {
	return &NullableVcenterHostSummary{value: val, isSet: true}
}

func (v NullableVcenterHostSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterHostSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


