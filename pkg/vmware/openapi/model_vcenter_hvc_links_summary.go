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

// VcenterHvcLinksSummary struct for VcenterHvcLinksSummary
type VcenterHvcLinksSummary struct {
	// Unique identifier for the link. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Link string `json:"link"`
	// The display name is set to the domain name which was set during create. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	DisplayName string `json:"display_name"`
}

// NewVcenterHvcLinksSummary instantiates a new VcenterHvcLinksSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterHvcLinksSummary(link string, displayName string) *VcenterHvcLinksSummary {
	this := VcenterHvcLinksSummary{}
	this.Link = link
	this.DisplayName = displayName
	return &this
}

// NewVcenterHvcLinksSummaryWithDefaults instantiates a new VcenterHvcLinksSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterHvcLinksSummaryWithDefaults() *VcenterHvcLinksSummary {
	this := VcenterHvcLinksSummary{}
	return &this
}

// GetLink returns the Link field value
func (o *VcenterHvcLinksSummary) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksSummary) GetLinkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *VcenterHvcLinksSummary) SetLink(v string) {
	o.Link = v
}

// GetDisplayName returns the DisplayName field value
func (o *VcenterHvcLinksSummary) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksSummary) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *VcenterHvcLinksSummary) SetDisplayName(v string) {
	o.DisplayName = v
}

func (o VcenterHvcLinksSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["link"] = o.Link
	}
	if true {
		toSerialize["display_name"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterHvcLinksSummary struct {
	value *VcenterHvcLinksSummary
	isSet bool
}

func (v NullableVcenterHvcLinksSummary) Get() *VcenterHvcLinksSummary {
	return v.value
}

func (v *NullableVcenterHvcLinksSummary) Set(val *VcenterHvcLinksSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterHvcLinksSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterHvcLinksSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterHvcLinksSummary(val *VcenterHvcLinksSummary) *NullableVcenterHvcLinksSummary {
	return &NullableVcenterHvcLinksSummary{value: val, isSet: true}
}

func (v NullableVcenterHvcLinksSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterHvcLinksSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


