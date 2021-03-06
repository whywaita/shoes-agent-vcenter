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

// VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec struct for VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec
type VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec struct {
	// The identifier of the Supervisor Service version. This must be an alphanumeric (a-z and 0-9) string and with maximum length of 63 characters and with the '-' and '.' characters allowed anywhere except the first or last character. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version.
	Version string `json:"version"`
	// A human readable name of the Supervisor Service version.
	DisplayName string `json:"display_name"`
	// A human readable description of the Supervisor Service version. If unset, the description for the service version will be empty.
	Description *string `json:"description,omitempty"`
	// Inline content that contains all service definition of the version, which shall be base64 encoded. The service definition here doesn't follow the vSphere application service format.
	Content string `json:"content"`
	// Whether or not the Supervisor Service version is from a trusted provider, this field must be set to false if the service version is not from a trusted provider. If it is set to be true, but the Versions.CustomCreateSpec.content is not signed or the signature is invalid, an InvalidArgument will be thrown. If unset, the default value is true. In this case, the Versions.CustomCreateSpec.content must be signed and will be verified.
	TrustedProvider *bool `json:"trusted_provider,omitempty"`
}

// NewVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec instantiates a new VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec(version string, displayName string, content string) *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec {
	this := VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec{}
	this.Version = version
	this.DisplayName = displayName
	this.Content = content
	return &this
}

// NewVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpecWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpecWithDefaults() *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec {
	this := VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec{}
	return &this
}

// GetVersion returns the Version field value
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) SetVersion(v string) {
	o.Version = v
}

// GetDisplayName returns the DisplayName field value
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) SetDescription(v string) {
	o.Description = &v
}

// GetContent returns the Content field value
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) SetContent(v string) {
	o.Content = v
}

// GetTrustedProvider returns the TrustedProvider field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetTrustedProvider() bool {
	if o == nil || o.TrustedProvider == nil {
		var ret bool
		return ret
	}
	return *o.TrustedProvider
}

// GetTrustedProviderOk returns a tuple with the TrustedProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetTrustedProviderOk() (*bool, bool) {
	if o == nil || o.TrustedProvider == nil {
		return nil, false
	}
	return o.TrustedProvider, true
}

// HasTrustedProvider returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) HasTrustedProvider() bool {
	if o != nil && o.TrustedProvider != nil {
		return true
	}

	return false
}

// SetTrustedProvider gets a reference to the given bool and assigns it to the TrustedProvider field.
func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) SetTrustedProvider(v bool) {
	o.TrustedProvider = &v
}

func (o VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["content"] = o.Content
	}
	if o.TrustedProvider != nil {
		toSerialize["trusted_provider"] = o.TrustedProvider
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec struct {
	value *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec
	isSet bool
}

func (v NullableVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) Get() *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec {
	return v.value
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) Set(val *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec(val *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) *NullableVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec {
	return &NullableVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


