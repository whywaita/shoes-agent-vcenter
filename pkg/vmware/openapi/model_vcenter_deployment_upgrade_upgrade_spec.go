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

// VcenterDeploymentUpgradeUpgradeSpec struct for VcenterDeploymentUpgradeUpgradeSpec
type VcenterDeploymentUpgradeUpgradeSpec struct {
	SourceAppliance VcenterDeploymentUpgradeSourceApplianceSpec `json:"source_appliance"`
	SourceLocation VcenterDeploymentLocationSpec `json:"source_location"`
	History *VcenterDeploymentHistoryMigrationSpec `json:"history,omitempty"`
	VcsaEmbedded *VcenterDeploymentUpgradeVcsaEmbeddedSpec `json:"vcsa_embedded,omitempty"`
	Psc *VcenterDeploymentUpgradePscSpec `json:"psc,omitempty"`
	// Use the default option for any questions that may come up during appliance configuration. If unset, will default to false.
	AutoAnswer *bool `json:"auto_answer,omitempty"`
}

// NewVcenterDeploymentUpgradeUpgradeSpec instantiates a new VcenterDeploymentUpgradeUpgradeSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentUpgradeUpgradeSpec(sourceAppliance VcenterDeploymentUpgradeSourceApplianceSpec, sourceLocation VcenterDeploymentLocationSpec) *VcenterDeploymentUpgradeUpgradeSpec {
	this := VcenterDeploymentUpgradeUpgradeSpec{}
	this.SourceAppliance = sourceAppliance
	this.SourceLocation = sourceLocation
	return &this
}

// NewVcenterDeploymentUpgradeUpgradeSpecWithDefaults instantiates a new VcenterDeploymentUpgradeUpgradeSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentUpgradeUpgradeSpecWithDefaults() *VcenterDeploymentUpgradeUpgradeSpec {
	this := VcenterDeploymentUpgradeUpgradeSpec{}
	return &this
}

// GetSourceAppliance returns the SourceAppliance field value
func (o *VcenterDeploymentUpgradeUpgradeSpec) GetSourceAppliance() VcenterDeploymentUpgradeSourceApplianceSpec {
	if o == nil {
		var ret VcenterDeploymentUpgradeSourceApplianceSpec
		return ret
	}

	return o.SourceAppliance
}

// GetSourceApplianceOk returns a tuple with the SourceAppliance field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentUpgradeUpgradeSpec) GetSourceApplianceOk() (*VcenterDeploymentUpgradeSourceApplianceSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceAppliance, true
}

// SetSourceAppliance sets field value
func (o *VcenterDeploymentUpgradeUpgradeSpec) SetSourceAppliance(v VcenterDeploymentUpgradeSourceApplianceSpec) {
	o.SourceAppliance = v
}

// GetSourceLocation returns the SourceLocation field value
func (o *VcenterDeploymentUpgradeUpgradeSpec) GetSourceLocation() VcenterDeploymentLocationSpec {
	if o == nil {
		var ret VcenterDeploymentLocationSpec
		return ret
	}

	return o.SourceLocation
}

// GetSourceLocationOk returns a tuple with the SourceLocation field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentUpgradeUpgradeSpec) GetSourceLocationOk() (*VcenterDeploymentLocationSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceLocation, true
}

// SetSourceLocation sets field value
func (o *VcenterDeploymentUpgradeUpgradeSpec) SetSourceLocation(v VcenterDeploymentLocationSpec) {
	o.SourceLocation = v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *VcenterDeploymentUpgradeUpgradeSpec) GetHistory() VcenterDeploymentHistoryMigrationSpec {
	if o == nil || o.History == nil {
		var ret VcenterDeploymentHistoryMigrationSpec
		return ret
	}
	return *o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentUpgradeUpgradeSpec) GetHistoryOk() (*VcenterDeploymentHistoryMigrationSpec, bool) {
	if o == nil || o.History == nil {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *VcenterDeploymentUpgradeUpgradeSpec) HasHistory() bool {
	if o != nil && o.History != nil {
		return true
	}

	return false
}

// SetHistory gets a reference to the given VcenterDeploymentHistoryMigrationSpec and assigns it to the History field.
func (o *VcenterDeploymentUpgradeUpgradeSpec) SetHistory(v VcenterDeploymentHistoryMigrationSpec) {
	o.History = &v
}

// GetVcsaEmbedded returns the VcsaEmbedded field value if set, zero value otherwise.
func (o *VcenterDeploymentUpgradeUpgradeSpec) GetVcsaEmbedded() VcenterDeploymentUpgradeVcsaEmbeddedSpec {
	if o == nil || o.VcsaEmbedded == nil {
		var ret VcenterDeploymentUpgradeVcsaEmbeddedSpec
		return ret
	}
	return *o.VcsaEmbedded
}

// GetVcsaEmbeddedOk returns a tuple with the VcsaEmbedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentUpgradeUpgradeSpec) GetVcsaEmbeddedOk() (*VcenterDeploymentUpgradeVcsaEmbeddedSpec, bool) {
	if o == nil || o.VcsaEmbedded == nil {
		return nil, false
	}
	return o.VcsaEmbedded, true
}

// HasVcsaEmbedded returns a boolean if a field has been set.
func (o *VcenterDeploymentUpgradeUpgradeSpec) HasVcsaEmbedded() bool {
	if o != nil && o.VcsaEmbedded != nil {
		return true
	}

	return false
}

// SetVcsaEmbedded gets a reference to the given VcenterDeploymentUpgradeVcsaEmbeddedSpec and assigns it to the VcsaEmbedded field.
func (o *VcenterDeploymentUpgradeUpgradeSpec) SetVcsaEmbedded(v VcenterDeploymentUpgradeVcsaEmbeddedSpec) {
	o.VcsaEmbedded = &v
}

// GetPsc returns the Psc field value if set, zero value otherwise.
func (o *VcenterDeploymentUpgradeUpgradeSpec) GetPsc() VcenterDeploymentUpgradePscSpec {
	if o == nil || o.Psc == nil {
		var ret VcenterDeploymentUpgradePscSpec
		return ret
	}
	return *o.Psc
}

// GetPscOk returns a tuple with the Psc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentUpgradeUpgradeSpec) GetPscOk() (*VcenterDeploymentUpgradePscSpec, bool) {
	if o == nil || o.Psc == nil {
		return nil, false
	}
	return o.Psc, true
}

// HasPsc returns a boolean if a field has been set.
func (o *VcenterDeploymentUpgradeUpgradeSpec) HasPsc() bool {
	if o != nil && o.Psc != nil {
		return true
	}

	return false
}

// SetPsc gets a reference to the given VcenterDeploymentUpgradePscSpec and assigns it to the Psc field.
func (o *VcenterDeploymentUpgradeUpgradeSpec) SetPsc(v VcenterDeploymentUpgradePscSpec) {
	o.Psc = &v
}

// GetAutoAnswer returns the AutoAnswer field value if set, zero value otherwise.
func (o *VcenterDeploymentUpgradeUpgradeSpec) GetAutoAnswer() bool {
	if o == nil || o.AutoAnswer == nil {
		var ret bool
		return ret
	}
	return *o.AutoAnswer
}

// GetAutoAnswerOk returns a tuple with the AutoAnswer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentUpgradeUpgradeSpec) GetAutoAnswerOk() (*bool, bool) {
	if o == nil || o.AutoAnswer == nil {
		return nil, false
	}
	return o.AutoAnswer, true
}

// HasAutoAnswer returns a boolean if a field has been set.
func (o *VcenterDeploymentUpgradeUpgradeSpec) HasAutoAnswer() bool {
	if o != nil && o.AutoAnswer != nil {
		return true
	}

	return false
}

// SetAutoAnswer gets a reference to the given bool and assigns it to the AutoAnswer field.
func (o *VcenterDeploymentUpgradeUpgradeSpec) SetAutoAnswer(v bool) {
	o.AutoAnswer = &v
}

func (o VcenterDeploymentUpgradeUpgradeSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source_appliance"] = o.SourceAppliance
	}
	if true {
		toSerialize["source_location"] = o.SourceLocation
	}
	if o.History != nil {
		toSerialize["history"] = o.History
	}
	if o.VcsaEmbedded != nil {
		toSerialize["vcsa_embedded"] = o.VcsaEmbedded
	}
	if o.Psc != nil {
		toSerialize["psc"] = o.Psc
	}
	if o.AutoAnswer != nil {
		toSerialize["auto_answer"] = o.AutoAnswer
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentUpgradeUpgradeSpec struct {
	value *VcenterDeploymentUpgradeUpgradeSpec
	isSet bool
}

func (v NullableVcenterDeploymentUpgradeUpgradeSpec) Get() *VcenterDeploymentUpgradeUpgradeSpec {
	return v.value
}

func (v *NullableVcenterDeploymentUpgradeUpgradeSpec) Set(val *VcenterDeploymentUpgradeUpgradeSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentUpgradeUpgradeSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentUpgradeUpgradeSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentUpgradeUpgradeSpec(val *VcenterDeploymentUpgradeUpgradeSpec) *NullableVcenterDeploymentUpgradeUpgradeSpec {
	return &NullableVcenterDeploymentUpgradeUpgradeSpec{value: val, isSet: true}
}

func (v NullableVcenterDeploymentUpgradeUpgradeSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentUpgradeUpgradeSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

