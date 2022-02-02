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

// VcenterNamespaceManagementLoadBalancersInfo struct for VcenterNamespaceManagementLoadBalancersInfo
type VcenterNamespaceManagementLoadBalancersInfo struct {
	// An DNS compliant identifier for a load balancer, which can be used to query or configure the load balancer properties. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.LoadBalancerConfig. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.LoadBalancerConfig.
	Id string `json:"id"`
	// IP address range from which virtual servers are assigned their IPs.
	AddressRanges []VcenterNamespaceManagementIPRange `json:"address_ranges"`
	Provider VcenterNamespaceManagementLoadBalancersProvider `json:"provider"`
	HaProxyInfo *VcenterNamespaceManagementLoadBalancersHAProxyInfo `json:"ha_proxy_info,omitempty"`
	AviInfo *VcenterNamespaceManagementLoadBalancersAviInfo `json:"avi_info,omitempty"`
}

// NewVcenterNamespaceManagementLoadBalancersInfo instantiates a new VcenterNamespaceManagementLoadBalancersInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementLoadBalancersInfo(id string, addressRanges []VcenterNamespaceManagementIPRange, provider VcenterNamespaceManagementLoadBalancersProvider) *VcenterNamespaceManagementLoadBalancersInfo {
	this := VcenterNamespaceManagementLoadBalancersInfo{}
	this.Id = id
	this.AddressRanges = addressRanges
	this.Provider = provider
	return &this
}

// NewVcenterNamespaceManagementLoadBalancersInfoWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementLoadBalancersInfoWithDefaults() *VcenterNamespaceManagementLoadBalancersInfo {
	this := VcenterNamespaceManagementLoadBalancersInfo{}
	return &this
}

// GetId returns the Id field value
func (o *VcenterNamespaceManagementLoadBalancersInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersInfo) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VcenterNamespaceManagementLoadBalancersInfo) SetId(v string) {
	o.Id = v
}

// GetAddressRanges returns the AddressRanges field value
func (o *VcenterNamespaceManagementLoadBalancersInfo) GetAddressRanges() []VcenterNamespaceManagementIPRange {
	if o == nil {
		var ret []VcenterNamespaceManagementIPRange
		return ret
	}

	return o.AddressRanges
}

// GetAddressRangesOk returns a tuple with the AddressRanges field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersInfo) GetAddressRangesOk() (*[]VcenterNamespaceManagementIPRange, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AddressRanges, true
}

// SetAddressRanges sets field value
func (o *VcenterNamespaceManagementLoadBalancersInfo) SetAddressRanges(v []VcenterNamespaceManagementIPRange) {
	o.AddressRanges = v
}

// GetProvider returns the Provider field value
func (o *VcenterNamespaceManagementLoadBalancersInfo) GetProvider() VcenterNamespaceManagementLoadBalancersProvider {
	if o == nil {
		var ret VcenterNamespaceManagementLoadBalancersProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersInfo) GetProviderOk() (*VcenterNamespaceManagementLoadBalancersProvider, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *VcenterNamespaceManagementLoadBalancersInfo) SetProvider(v VcenterNamespaceManagementLoadBalancersProvider) {
	o.Provider = v
}

// GetHaProxyInfo returns the HaProxyInfo field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementLoadBalancersInfo) GetHaProxyInfo() VcenterNamespaceManagementLoadBalancersHAProxyInfo {
	if o == nil || o.HaProxyInfo == nil {
		var ret VcenterNamespaceManagementLoadBalancersHAProxyInfo
		return ret
	}
	return *o.HaProxyInfo
}

// GetHaProxyInfoOk returns a tuple with the HaProxyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersInfo) GetHaProxyInfoOk() (*VcenterNamespaceManagementLoadBalancersHAProxyInfo, bool) {
	if o == nil || o.HaProxyInfo == nil {
		return nil, false
	}
	return o.HaProxyInfo, true
}

// HasHaProxyInfo returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementLoadBalancersInfo) HasHaProxyInfo() bool {
	if o != nil && o.HaProxyInfo != nil {
		return true
	}

	return false
}

// SetHaProxyInfo gets a reference to the given VcenterNamespaceManagementLoadBalancersHAProxyInfo and assigns it to the HaProxyInfo field.
func (o *VcenterNamespaceManagementLoadBalancersInfo) SetHaProxyInfo(v VcenterNamespaceManagementLoadBalancersHAProxyInfo) {
	o.HaProxyInfo = &v
}

// GetAviInfo returns the AviInfo field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementLoadBalancersInfo) GetAviInfo() VcenterNamespaceManagementLoadBalancersAviInfo {
	if o == nil || o.AviInfo == nil {
		var ret VcenterNamespaceManagementLoadBalancersAviInfo
		return ret
	}
	return *o.AviInfo
}

// GetAviInfoOk returns a tuple with the AviInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementLoadBalancersInfo) GetAviInfoOk() (*VcenterNamespaceManagementLoadBalancersAviInfo, bool) {
	if o == nil || o.AviInfo == nil {
		return nil, false
	}
	return o.AviInfo, true
}

// HasAviInfo returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementLoadBalancersInfo) HasAviInfo() bool {
	if o != nil && o.AviInfo != nil {
		return true
	}

	return false
}

// SetAviInfo gets a reference to the given VcenterNamespaceManagementLoadBalancersAviInfo and assigns it to the AviInfo field.
func (o *VcenterNamespaceManagementLoadBalancersInfo) SetAviInfo(v VcenterNamespaceManagementLoadBalancersAviInfo) {
	o.AviInfo = &v
}

func (o VcenterNamespaceManagementLoadBalancersInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["address_ranges"] = o.AddressRanges
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if o.HaProxyInfo != nil {
		toSerialize["ha_proxy_info"] = o.HaProxyInfo
	}
	if o.AviInfo != nil {
		toSerialize["avi_info"] = o.AviInfo
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementLoadBalancersInfo struct {
	value *VcenterNamespaceManagementLoadBalancersInfo
	isSet bool
}

func (v NullableVcenterNamespaceManagementLoadBalancersInfo) Get() *VcenterNamespaceManagementLoadBalancersInfo {
	return v.value
}

func (v *NullableVcenterNamespaceManagementLoadBalancersInfo) Set(val *VcenterNamespaceManagementLoadBalancersInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementLoadBalancersInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementLoadBalancersInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementLoadBalancersInfo(val *VcenterNamespaceManagementLoadBalancersInfo) *NullableVcenterNamespaceManagementLoadBalancersInfo {
	return &NullableVcenterNamespaceManagementLoadBalancersInfo{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementLoadBalancersInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementLoadBalancersInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


