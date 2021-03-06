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

// VcenterNamespacesNamespaceSelfServiceSummary struct for VcenterNamespacesNamespaceSelfServiceSummary
type VcenterNamespacesNamespaceSelfServiceSummary struct {
	// Identifier for the cluster to which namespace service is associated. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource.
	Cluster string `json:"cluster"`
	Capability VcenterNamespacesNamespaceSelfServiceCapability `json:"capability"`
	Status VcenterNamespacesNamespaceSelfServiceStatus `json:"status"`
}

// NewVcenterNamespacesNamespaceSelfServiceSummary instantiates a new VcenterNamespacesNamespaceSelfServiceSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespacesNamespaceSelfServiceSummary(cluster string, capability VcenterNamespacesNamespaceSelfServiceCapability, status VcenterNamespacesNamespaceSelfServiceStatus) *VcenterNamespacesNamespaceSelfServiceSummary {
	this := VcenterNamespacesNamespaceSelfServiceSummary{}
	this.Cluster = cluster
	this.Capability = capability
	this.Status = status
	return &this
}

// NewVcenterNamespacesNamespaceSelfServiceSummaryWithDefaults instantiates a new VcenterNamespacesNamespaceSelfServiceSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespacesNamespaceSelfServiceSummaryWithDefaults() *VcenterNamespacesNamespaceSelfServiceSummary {
	this := VcenterNamespacesNamespaceSelfServiceSummary{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *VcenterNamespacesNamespaceSelfServiceSummary) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesNamespaceSelfServiceSummary) GetClusterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *VcenterNamespacesNamespaceSelfServiceSummary) SetCluster(v string) {
	o.Cluster = v
}

// GetCapability returns the Capability field value
func (o *VcenterNamespacesNamespaceSelfServiceSummary) GetCapability() VcenterNamespacesNamespaceSelfServiceCapability {
	if o == nil {
		var ret VcenterNamespacesNamespaceSelfServiceCapability
		return ret
	}

	return o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesNamespaceSelfServiceSummary) GetCapabilityOk() (*VcenterNamespacesNamespaceSelfServiceCapability, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Capability, true
}

// SetCapability sets field value
func (o *VcenterNamespacesNamespaceSelfServiceSummary) SetCapability(v VcenterNamespacesNamespaceSelfServiceCapability) {
	o.Capability = v
}

// GetStatus returns the Status field value
func (o *VcenterNamespacesNamespaceSelfServiceSummary) GetStatus() VcenterNamespacesNamespaceSelfServiceStatus {
	if o == nil {
		var ret VcenterNamespacesNamespaceSelfServiceStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesNamespaceSelfServiceSummary) GetStatusOk() (*VcenterNamespacesNamespaceSelfServiceStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VcenterNamespacesNamespaceSelfServiceSummary) SetStatus(v VcenterNamespacesNamespaceSelfServiceStatus) {
	o.Status = v
}

func (o VcenterNamespacesNamespaceSelfServiceSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster"] = o.Cluster
	}
	if true {
		toSerialize["capability"] = o.Capability
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespacesNamespaceSelfServiceSummary struct {
	value *VcenterNamespacesNamespaceSelfServiceSummary
	isSet bool
}

func (v NullableVcenterNamespacesNamespaceSelfServiceSummary) Get() *VcenterNamespacesNamespaceSelfServiceSummary {
	return v.value
}

func (v *NullableVcenterNamespacesNamespaceSelfServiceSummary) Set(val *VcenterNamespacesNamespaceSelfServiceSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespacesNamespaceSelfServiceSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespacesNamespaceSelfServiceSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespacesNamespaceSelfServiceSummary(val *VcenterNamespacesNamespaceSelfServiceSummary) *NullableVcenterNamespacesNamespaceSelfServiceSummary {
	return &NullableVcenterNamespacesNamespaceSelfServiceSummary{value: val, isSet: true}
}

func (v NullableVcenterNamespacesNamespaceSelfServiceSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespacesNamespaceSelfServiceSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


