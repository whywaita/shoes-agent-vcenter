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

// VcenterNamespacesNamespaceTemplatesSummary struct for VcenterNamespacesNamespaceTemplatesSummary
type VcenterNamespacesNamespaceTemplatesSummary struct {
	// Identifier for the cluster associated with namespace template. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource.
	Cluster string `json:"cluster"`
	// Name of the namespace template. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespaces.NamespaceTemplate. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespaces.NamespaceTemplate.
	Template string `json:"template"`
}

// NewVcenterNamespacesNamespaceTemplatesSummary instantiates a new VcenterNamespacesNamespaceTemplatesSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespacesNamespaceTemplatesSummary(cluster string, template string) *VcenterNamespacesNamespaceTemplatesSummary {
	this := VcenterNamespacesNamespaceTemplatesSummary{}
	this.Cluster = cluster
	this.Template = template
	return &this
}

// NewVcenterNamespacesNamespaceTemplatesSummaryWithDefaults instantiates a new VcenterNamespacesNamespaceTemplatesSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespacesNamespaceTemplatesSummaryWithDefaults() *VcenterNamespacesNamespaceTemplatesSummary {
	this := VcenterNamespacesNamespaceTemplatesSummary{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *VcenterNamespacesNamespaceTemplatesSummary) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesNamespaceTemplatesSummary) GetClusterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *VcenterNamespacesNamespaceTemplatesSummary) SetCluster(v string) {
	o.Cluster = v
}

// GetTemplate returns the Template field value
func (o *VcenterNamespacesNamespaceTemplatesSummary) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesNamespaceTemplatesSummary) GetTemplateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *VcenterNamespacesNamespaceTemplatesSummary) SetTemplate(v string) {
	o.Template = v
}

func (o VcenterNamespacesNamespaceTemplatesSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster"] = o.Cluster
	}
	if true {
		toSerialize["template"] = o.Template
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespacesNamespaceTemplatesSummary struct {
	value *VcenterNamespacesNamespaceTemplatesSummary
	isSet bool
}

func (v NullableVcenterNamespacesNamespaceTemplatesSummary) Get() *VcenterNamespacesNamespaceTemplatesSummary {
	return v.value
}

func (v *NullableVcenterNamespacesNamespaceTemplatesSummary) Set(val *VcenterNamespacesNamespaceTemplatesSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespacesNamespaceTemplatesSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespacesNamespaceTemplatesSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespacesNamespaceTemplatesSummary(val *VcenterNamespacesNamespaceTemplatesSummary) *NullableVcenterNamespacesNamespaceTemplatesSummary {
	return &NullableVcenterNamespacesNamespaceTemplatesSummary{value: val, isSet: true}
}

func (v NullableVcenterNamespacesNamespaceTemplatesSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespacesNamespaceTemplatesSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

