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

// VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec struct for VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec
type VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec struct {
	// Master key identifier created for the provider.   A unique Key identifier.     If unset, masterKeyId will remain unchanged.
	MasterKeyId *string `json:"master_key_id,omitempty"`
	KeyServer *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec `json:"key_server,omitempty"`
}

// NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec {
	this := VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec{}
	return &this
}

// NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec {
	this := VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec{}
	return &this
}

// GetMasterKeyId returns the MasterKeyId field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) GetMasterKeyId() string {
	if o == nil || o.MasterKeyId == nil {
		var ret string
		return ret
	}
	return *o.MasterKeyId
}

// GetMasterKeyIdOk returns a tuple with the MasterKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) GetMasterKeyIdOk() (*string, bool) {
	if o == nil || o.MasterKeyId == nil {
		return nil, false
	}
	return o.MasterKeyId, true
}

// HasMasterKeyId returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) HasMasterKeyId() bool {
	if o != nil && o.MasterKeyId != nil {
		return true
	}

	return false
}

// SetMasterKeyId gets a reference to the given string and assigns it to the MasterKeyId field.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) SetMasterKeyId(v string) {
	o.MasterKeyId = &v
}

// GetKeyServer returns the KeyServer field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) GetKeyServer() VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec {
	if o == nil || o.KeyServer == nil {
		var ret VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec
		return ret
	}
	return *o.KeyServer
}

// GetKeyServerOk returns a tuple with the KeyServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) GetKeyServerOk() (*VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec, bool) {
	if o == nil || o.KeyServer == nil {
		return nil, false
	}
	return o.KeyServer, true
}

// HasKeyServer returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) HasKeyServer() bool {
	if o != nil && o.KeyServer != nil {
		return true
	}

	return false
}

// SetKeyServer gets a reference to the given VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec and assigns it to the KeyServer field.
func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) SetKeyServer(v VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) {
	o.KeyServer = &v
}

func (o VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MasterKeyId != nil {
		toSerialize["master_key_id"] = o.MasterKeyId
	}
	if o.KeyServer != nil {
		toSerialize["key_server"] = o.KeyServer
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec struct {
	value *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) Get() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) Set(val *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec(val *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec {
	return &NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


