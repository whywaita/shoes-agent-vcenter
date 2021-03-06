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

// VcenterNamespaceManagementDistributedSwitchCompatibilitySummary struct for VcenterNamespaceManagementDistributedSwitchCompatibilitySummary
type VcenterNamespaceManagementDistributedSwitchCompatibilitySummary struct {
	// Identifier of the switch. If networkProvider is either unset or is set to NSXT_CONTAINER_PLUGIN, the value of this field will refer to the UUID of a vim.DistributedVirtualSwitch. Otherwise, the value of the field will refer to the ID of a vim.DistributedVirtualSwitch. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vSphereDistributedSwitch. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vSphereDistributedSwitch.
	DistributedSwitch string `json:"distributed_switch"`
	// Compatibility of this switch with vSphere Namespaces.
	Compatible bool `json:"compatible"`
	// List of reasons for incompatibility. If unset, this Distributed Switch is compatible.
	IncompatibilityReasons *[]VapiStdLocalizableMessage `json:"incompatibility_reasons,omitempty"`
	NetworkProvider *VcenterNamespaceManagementClustersNetworkProvider `json:"network_provider,omitempty"`
	// List of compatible (PortGroup) Networks under the distributed switch. This field is optional because it was added in a newer version than its parent node. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Network. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Network.
	CompatibleNetworks *[]string `json:"compatible_networks,omitempty"`
}

// NewVcenterNamespaceManagementDistributedSwitchCompatibilitySummary instantiates a new VcenterNamespaceManagementDistributedSwitchCompatibilitySummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementDistributedSwitchCompatibilitySummary(distributedSwitch string, compatible bool) *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary {
	this := VcenterNamespaceManagementDistributedSwitchCompatibilitySummary{}
	this.DistributedSwitch = distributedSwitch
	this.Compatible = compatible
	return &this
}

// NewVcenterNamespaceManagementDistributedSwitchCompatibilitySummaryWithDefaults instantiates a new VcenterNamespaceManagementDistributedSwitchCompatibilitySummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementDistributedSwitchCompatibilitySummaryWithDefaults() *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary {
	this := VcenterNamespaceManagementDistributedSwitchCompatibilitySummary{}
	return &this
}

// GetDistributedSwitch returns the DistributedSwitch field value
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetDistributedSwitch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DistributedSwitch
}

// GetDistributedSwitchOk returns a tuple with the DistributedSwitch field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetDistributedSwitchOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DistributedSwitch, true
}

// SetDistributedSwitch sets field value
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) SetDistributedSwitch(v string) {
	o.DistributedSwitch = v
}

// GetCompatible returns the Compatible field value
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetCompatible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Compatible
}

// GetCompatibleOk returns a tuple with the Compatible field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetCompatibleOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Compatible, true
}

// SetCompatible sets field value
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) SetCompatible(v bool) {
	o.Compatible = v
}

// GetIncompatibilityReasons returns the IncompatibilityReasons field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetIncompatibilityReasons() []VapiStdLocalizableMessage {
	if o == nil || o.IncompatibilityReasons == nil {
		var ret []VapiStdLocalizableMessage
		return ret
	}
	return *o.IncompatibilityReasons
}

// GetIncompatibilityReasonsOk returns a tuple with the IncompatibilityReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetIncompatibilityReasonsOk() (*[]VapiStdLocalizableMessage, bool) {
	if o == nil || o.IncompatibilityReasons == nil {
		return nil, false
	}
	return o.IncompatibilityReasons, true
}

// HasIncompatibilityReasons returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) HasIncompatibilityReasons() bool {
	if o != nil && o.IncompatibilityReasons != nil {
		return true
	}

	return false
}

// SetIncompatibilityReasons gets a reference to the given []VapiStdLocalizableMessage and assigns it to the IncompatibilityReasons field.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) SetIncompatibilityReasons(v []VapiStdLocalizableMessage) {
	o.IncompatibilityReasons = &v
}

// GetNetworkProvider returns the NetworkProvider field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetNetworkProvider() VcenterNamespaceManagementClustersNetworkProvider {
	if o == nil || o.NetworkProvider == nil {
		var ret VcenterNamespaceManagementClustersNetworkProvider
		return ret
	}
	return *o.NetworkProvider
}

// GetNetworkProviderOk returns a tuple with the NetworkProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetNetworkProviderOk() (*VcenterNamespaceManagementClustersNetworkProvider, bool) {
	if o == nil || o.NetworkProvider == nil {
		return nil, false
	}
	return o.NetworkProvider, true
}

// HasNetworkProvider returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) HasNetworkProvider() bool {
	if o != nil && o.NetworkProvider != nil {
		return true
	}

	return false
}

// SetNetworkProvider gets a reference to the given VcenterNamespaceManagementClustersNetworkProvider and assigns it to the NetworkProvider field.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) SetNetworkProvider(v VcenterNamespaceManagementClustersNetworkProvider) {
	o.NetworkProvider = &v
}

// GetCompatibleNetworks returns the CompatibleNetworks field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetCompatibleNetworks() []string {
	if o == nil || o.CompatibleNetworks == nil {
		var ret []string
		return ret
	}
	return *o.CompatibleNetworks
}

// GetCompatibleNetworksOk returns a tuple with the CompatibleNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetCompatibleNetworksOk() (*[]string, bool) {
	if o == nil || o.CompatibleNetworks == nil {
		return nil, false
	}
	return o.CompatibleNetworks, true
}

// HasCompatibleNetworks returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) HasCompatibleNetworks() bool {
	if o != nil && o.CompatibleNetworks != nil {
		return true
	}

	return false
}

// SetCompatibleNetworks gets a reference to the given []string and assigns it to the CompatibleNetworks field.
func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) SetCompatibleNetworks(v []string) {
	o.CompatibleNetworks = &v
}

func (o VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["distributed_switch"] = o.DistributedSwitch
	}
	if true {
		toSerialize["compatible"] = o.Compatible
	}
	if o.IncompatibilityReasons != nil {
		toSerialize["incompatibility_reasons"] = o.IncompatibilityReasons
	}
	if o.NetworkProvider != nil {
		toSerialize["network_provider"] = o.NetworkProvider
	}
	if o.CompatibleNetworks != nil {
		toSerialize["compatible_networks"] = o.CompatibleNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementDistributedSwitchCompatibilitySummary struct {
	value *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary
	isSet bool
}

func (v NullableVcenterNamespaceManagementDistributedSwitchCompatibilitySummary) Get() *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary {
	return v.value
}

func (v *NullableVcenterNamespaceManagementDistributedSwitchCompatibilitySummary) Set(val *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementDistributedSwitchCompatibilitySummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementDistributedSwitchCompatibilitySummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementDistributedSwitchCompatibilitySummary(val *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) *NullableVcenterNamespaceManagementDistributedSwitchCompatibilitySummary {
	return &NullableVcenterNamespaceManagementDistributedSwitchCompatibilitySummary{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementDistributedSwitchCompatibilitySummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementDistributedSwitchCompatibilitySummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


