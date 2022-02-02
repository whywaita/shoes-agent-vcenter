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

// VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec struct for VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec
type VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec struct {
	// The unqiue identifier of a connection profile. If unset, no filtration will be performed by ID. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: esx.authentication.clientprofile. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: esx.authentication.clientprofile.
	Id *[]string `json:"id,omitempty"`
	// The principal used by the vCenter to retrieve tokens. If unset, no filtration will be performed by principals.
	Principals *[]VcenterTrustedInfrastructureStsPrincipal `json:"principals,omitempty"`
	// The service which created and signed the security token. If unset, no filtration will be performed by issuer. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: esx.authentication.trust.security-token-issuer. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: esx.authentication.trust.security-token-issuer.
	Issuer *[]string `json:"issuer,omitempty"`
}

// NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec() *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec {
	this := VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec{}
	return &this
}

// NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec {
	this := VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) GetId() []string {
	if o == nil || o.Id == nil {
		var ret []string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) GetIdOk() (*[]string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given []string and assigns it to the Id field.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) SetId(v []string) {
	o.Id = &v
}

// GetPrincipals returns the Principals field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) GetPrincipals() []VcenterTrustedInfrastructureStsPrincipal {
	if o == nil || o.Principals == nil {
		var ret []VcenterTrustedInfrastructureStsPrincipal
		return ret
	}
	return *o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) GetPrincipalsOk() (*[]VcenterTrustedInfrastructureStsPrincipal, bool) {
	if o == nil || o.Principals == nil {
		return nil, false
	}
	return o.Principals, true
}

// HasPrincipals returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) HasPrincipals() bool {
	if o != nil && o.Principals != nil {
		return true
	}

	return false
}

// SetPrincipals gets a reference to the given []VcenterTrustedInfrastructureStsPrincipal and assigns it to the Principals field.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) SetPrincipals(v []VcenterTrustedInfrastructureStsPrincipal) {
	o.Principals = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) GetIssuer() []string {
	if o == nil || o.Issuer == nil {
		var ret []string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) GetIssuerOk() (*[]string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given []string and assigns it to the Issuer field.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) SetIssuer(v []string) {
	o.Issuer = &v
}

func (o VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Principals != nil {
		toSerialize["principals"] = o.Principals
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec struct {
	value *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) Get() *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) Set(val *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec(val *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) *NullableVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec {
	return &NullableVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

