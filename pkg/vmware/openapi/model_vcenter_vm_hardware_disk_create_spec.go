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

// VcenterVmHardwareDiskCreateSpec struct for VcenterVmHardwareDiskCreateSpec
type VcenterVmHardwareDiskCreateSpec struct {
	Type *VcenterVmHardwareDiskHostBusAdapterType `json:"type,omitempty"`
	Ide *VcenterVmHardwareIdeAddressSpec `json:"ide,omitempty"`
	Scsi *VcenterVmHardwareScsiAddressSpec `json:"scsi,omitempty"`
	Sata *VcenterVmHardwareSataAddressSpec `json:"sata,omitempty"`
	Nvme *VcenterVmHardwareNvmeAddressSpec `json:"nvme,omitempty"`
	Backing *VcenterVmHardwareDiskBackingSpec `json:"backing,omitempty"`
	NewVmdk *VcenterVmHardwareDiskVmdkCreateSpec `json:"new_vmdk,omitempty"`
}

// NewVcenterVmHardwareDiskCreateSpec instantiates a new VcenterVmHardwareDiskCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareDiskCreateSpec() *VcenterVmHardwareDiskCreateSpec {
	this := VcenterVmHardwareDiskCreateSpec{}
	return &this
}

// NewVcenterVmHardwareDiskCreateSpecWithDefaults instantiates a new VcenterVmHardwareDiskCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareDiskCreateSpecWithDefaults() *VcenterVmHardwareDiskCreateSpec {
	this := VcenterVmHardwareDiskCreateSpec{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskCreateSpec) GetType() VcenterVmHardwareDiskHostBusAdapterType {
	if o == nil || o.Type == nil {
		var ret VcenterVmHardwareDiskHostBusAdapterType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskCreateSpec) GetTypeOk() (*VcenterVmHardwareDiskHostBusAdapterType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskCreateSpec) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given VcenterVmHardwareDiskHostBusAdapterType and assigns it to the Type field.
func (o *VcenterVmHardwareDiskCreateSpec) SetType(v VcenterVmHardwareDiskHostBusAdapterType) {
	o.Type = &v
}

// GetIde returns the Ide field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskCreateSpec) GetIde() VcenterVmHardwareIdeAddressSpec {
	if o == nil || o.Ide == nil {
		var ret VcenterVmHardwareIdeAddressSpec
		return ret
	}
	return *o.Ide
}

// GetIdeOk returns a tuple with the Ide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskCreateSpec) GetIdeOk() (*VcenterVmHardwareIdeAddressSpec, bool) {
	if o == nil || o.Ide == nil {
		return nil, false
	}
	return o.Ide, true
}

// HasIde returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskCreateSpec) HasIde() bool {
	if o != nil && o.Ide != nil {
		return true
	}

	return false
}

// SetIde gets a reference to the given VcenterVmHardwareIdeAddressSpec and assigns it to the Ide field.
func (o *VcenterVmHardwareDiskCreateSpec) SetIde(v VcenterVmHardwareIdeAddressSpec) {
	o.Ide = &v
}

// GetScsi returns the Scsi field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskCreateSpec) GetScsi() VcenterVmHardwareScsiAddressSpec {
	if o == nil || o.Scsi == nil {
		var ret VcenterVmHardwareScsiAddressSpec
		return ret
	}
	return *o.Scsi
}

// GetScsiOk returns a tuple with the Scsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskCreateSpec) GetScsiOk() (*VcenterVmHardwareScsiAddressSpec, bool) {
	if o == nil || o.Scsi == nil {
		return nil, false
	}
	return o.Scsi, true
}

// HasScsi returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskCreateSpec) HasScsi() bool {
	if o != nil && o.Scsi != nil {
		return true
	}

	return false
}

// SetScsi gets a reference to the given VcenterVmHardwareScsiAddressSpec and assigns it to the Scsi field.
func (o *VcenterVmHardwareDiskCreateSpec) SetScsi(v VcenterVmHardwareScsiAddressSpec) {
	o.Scsi = &v
}

// GetSata returns the Sata field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskCreateSpec) GetSata() VcenterVmHardwareSataAddressSpec {
	if o == nil || o.Sata == nil {
		var ret VcenterVmHardwareSataAddressSpec
		return ret
	}
	return *o.Sata
}

// GetSataOk returns a tuple with the Sata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskCreateSpec) GetSataOk() (*VcenterVmHardwareSataAddressSpec, bool) {
	if o == nil || o.Sata == nil {
		return nil, false
	}
	return o.Sata, true
}

// HasSata returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskCreateSpec) HasSata() bool {
	if o != nil && o.Sata != nil {
		return true
	}

	return false
}

// SetSata gets a reference to the given VcenterVmHardwareSataAddressSpec and assigns it to the Sata field.
func (o *VcenterVmHardwareDiskCreateSpec) SetSata(v VcenterVmHardwareSataAddressSpec) {
	o.Sata = &v
}

// GetNvme returns the Nvme field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskCreateSpec) GetNvme() VcenterVmHardwareNvmeAddressSpec {
	if o == nil || o.Nvme == nil {
		var ret VcenterVmHardwareNvmeAddressSpec
		return ret
	}
	return *o.Nvme
}

// GetNvmeOk returns a tuple with the Nvme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskCreateSpec) GetNvmeOk() (*VcenterVmHardwareNvmeAddressSpec, bool) {
	if o == nil || o.Nvme == nil {
		return nil, false
	}
	return o.Nvme, true
}

// HasNvme returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskCreateSpec) HasNvme() bool {
	if o != nil && o.Nvme != nil {
		return true
	}

	return false
}

// SetNvme gets a reference to the given VcenterVmHardwareNvmeAddressSpec and assigns it to the Nvme field.
func (o *VcenterVmHardwareDiskCreateSpec) SetNvme(v VcenterVmHardwareNvmeAddressSpec) {
	o.Nvme = &v
}

// GetBacking returns the Backing field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskCreateSpec) GetBacking() VcenterVmHardwareDiskBackingSpec {
	if o == nil || o.Backing == nil {
		var ret VcenterVmHardwareDiskBackingSpec
		return ret
	}
	return *o.Backing
}

// GetBackingOk returns a tuple with the Backing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskCreateSpec) GetBackingOk() (*VcenterVmHardwareDiskBackingSpec, bool) {
	if o == nil || o.Backing == nil {
		return nil, false
	}
	return o.Backing, true
}

// HasBacking returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskCreateSpec) HasBacking() bool {
	if o != nil && o.Backing != nil {
		return true
	}

	return false
}

// SetBacking gets a reference to the given VcenterVmHardwareDiskBackingSpec and assigns it to the Backing field.
func (o *VcenterVmHardwareDiskCreateSpec) SetBacking(v VcenterVmHardwareDiskBackingSpec) {
	o.Backing = &v
}

// GetNewVmdk returns the NewVmdk field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskCreateSpec) GetNewVmdk() VcenterVmHardwareDiskVmdkCreateSpec {
	if o == nil || o.NewVmdk == nil {
		var ret VcenterVmHardwareDiskVmdkCreateSpec
		return ret
	}
	return *o.NewVmdk
}

// GetNewVmdkOk returns a tuple with the NewVmdk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskCreateSpec) GetNewVmdkOk() (*VcenterVmHardwareDiskVmdkCreateSpec, bool) {
	if o == nil || o.NewVmdk == nil {
		return nil, false
	}
	return o.NewVmdk, true
}

// HasNewVmdk returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskCreateSpec) HasNewVmdk() bool {
	if o != nil && o.NewVmdk != nil {
		return true
	}

	return false
}

// SetNewVmdk gets a reference to the given VcenterVmHardwareDiskVmdkCreateSpec and assigns it to the NewVmdk field.
func (o *VcenterVmHardwareDiskCreateSpec) SetNewVmdk(v VcenterVmHardwareDiskVmdkCreateSpec) {
	o.NewVmdk = &v
}

func (o VcenterVmHardwareDiskCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Ide != nil {
		toSerialize["ide"] = o.Ide
	}
	if o.Scsi != nil {
		toSerialize["scsi"] = o.Scsi
	}
	if o.Sata != nil {
		toSerialize["sata"] = o.Sata
	}
	if o.Nvme != nil {
		toSerialize["nvme"] = o.Nvme
	}
	if o.Backing != nil {
		toSerialize["backing"] = o.Backing
	}
	if o.NewVmdk != nil {
		toSerialize["new_vmdk"] = o.NewVmdk
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareDiskCreateSpec struct {
	value *VcenterVmHardwareDiskCreateSpec
	isSet bool
}

func (v NullableVcenterVmHardwareDiskCreateSpec) Get() *VcenterVmHardwareDiskCreateSpec {
	return v.value
}

func (v *NullableVcenterVmHardwareDiskCreateSpec) Set(val *VcenterVmHardwareDiskCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareDiskCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareDiskCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareDiskCreateSpec(val *VcenterVmHardwareDiskCreateSpec) *NullableVcenterVmHardwareDiskCreateSpec {
	return &NullableVcenterVmHardwareDiskCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareDiskCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareDiskCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

