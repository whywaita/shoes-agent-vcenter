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

// VcenterVmHardwareDiskInfo struct for VcenterVmHardwareDiskInfo
type VcenterVmHardwareDiskInfo struct {
	// Device label.
	Label string `json:"label"`
	Type VcenterVmHardwareDiskHostBusAdapterType `json:"type"`
	Ide *VcenterVmHardwareIdeAddressInfo `json:"ide,omitempty"`
	Scsi *VcenterVmHardwareScsiAddressInfo `json:"scsi,omitempty"`
	Sata *VcenterVmHardwareSataAddressInfo `json:"sata,omitempty"`
	Nvme *VcenterVmHardwareNvmeAddressInfo `json:"nvme,omitempty"`
	Backing VcenterVmHardwareDiskBackingInfo `json:"backing"`
	// Capacity of the virtual disk in bytes. If unset, virtual disk is inaccessible or disk capacity is 0.
	Capacity *int64 `json:"capacity,omitempty"`
}

// NewVcenterVmHardwareDiskInfo instantiates a new VcenterVmHardwareDiskInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareDiskInfo(label string, type_ VcenterVmHardwareDiskHostBusAdapterType, backing VcenterVmHardwareDiskBackingInfo) *VcenterVmHardwareDiskInfo {
	this := VcenterVmHardwareDiskInfo{}
	this.Label = label
	this.Type = type_
	this.Backing = backing
	return &this
}

// NewVcenterVmHardwareDiskInfoWithDefaults instantiates a new VcenterVmHardwareDiskInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareDiskInfoWithDefaults() *VcenterVmHardwareDiskInfo {
	this := VcenterVmHardwareDiskInfo{}
	return &this
}

// GetLabel returns the Label field value
func (o *VcenterVmHardwareDiskInfo) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskInfo) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *VcenterVmHardwareDiskInfo) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value
func (o *VcenterVmHardwareDiskInfo) GetType() VcenterVmHardwareDiskHostBusAdapterType {
	if o == nil {
		var ret VcenterVmHardwareDiskHostBusAdapterType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskInfo) GetTypeOk() (*VcenterVmHardwareDiskHostBusAdapterType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterVmHardwareDiskInfo) SetType(v VcenterVmHardwareDiskHostBusAdapterType) {
	o.Type = v
}

// GetIde returns the Ide field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskInfo) GetIde() VcenterVmHardwareIdeAddressInfo {
	if o == nil || o.Ide == nil {
		var ret VcenterVmHardwareIdeAddressInfo
		return ret
	}
	return *o.Ide
}

// GetIdeOk returns a tuple with the Ide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskInfo) GetIdeOk() (*VcenterVmHardwareIdeAddressInfo, bool) {
	if o == nil || o.Ide == nil {
		return nil, false
	}
	return o.Ide, true
}

// HasIde returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskInfo) HasIde() bool {
	if o != nil && o.Ide != nil {
		return true
	}

	return false
}

// SetIde gets a reference to the given VcenterVmHardwareIdeAddressInfo and assigns it to the Ide field.
func (o *VcenterVmHardwareDiskInfo) SetIde(v VcenterVmHardwareIdeAddressInfo) {
	o.Ide = &v
}

// GetScsi returns the Scsi field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskInfo) GetScsi() VcenterVmHardwareScsiAddressInfo {
	if o == nil || o.Scsi == nil {
		var ret VcenterVmHardwareScsiAddressInfo
		return ret
	}
	return *o.Scsi
}

// GetScsiOk returns a tuple with the Scsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskInfo) GetScsiOk() (*VcenterVmHardwareScsiAddressInfo, bool) {
	if o == nil || o.Scsi == nil {
		return nil, false
	}
	return o.Scsi, true
}

// HasScsi returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskInfo) HasScsi() bool {
	if o != nil && o.Scsi != nil {
		return true
	}

	return false
}

// SetScsi gets a reference to the given VcenterVmHardwareScsiAddressInfo and assigns it to the Scsi field.
func (o *VcenterVmHardwareDiskInfo) SetScsi(v VcenterVmHardwareScsiAddressInfo) {
	o.Scsi = &v
}

// GetSata returns the Sata field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskInfo) GetSata() VcenterVmHardwareSataAddressInfo {
	if o == nil || o.Sata == nil {
		var ret VcenterVmHardwareSataAddressInfo
		return ret
	}
	return *o.Sata
}

// GetSataOk returns a tuple with the Sata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskInfo) GetSataOk() (*VcenterVmHardwareSataAddressInfo, bool) {
	if o == nil || o.Sata == nil {
		return nil, false
	}
	return o.Sata, true
}

// HasSata returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskInfo) HasSata() bool {
	if o != nil && o.Sata != nil {
		return true
	}

	return false
}

// SetSata gets a reference to the given VcenterVmHardwareSataAddressInfo and assigns it to the Sata field.
func (o *VcenterVmHardwareDiskInfo) SetSata(v VcenterVmHardwareSataAddressInfo) {
	o.Sata = &v
}

// GetNvme returns the Nvme field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskInfo) GetNvme() VcenterVmHardwareNvmeAddressInfo {
	if o == nil || o.Nvme == nil {
		var ret VcenterVmHardwareNvmeAddressInfo
		return ret
	}
	return *o.Nvme
}

// GetNvmeOk returns a tuple with the Nvme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskInfo) GetNvmeOk() (*VcenterVmHardwareNvmeAddressInfo, bool) {
	if o == nil || o.Nvme == nil {
		return nil, false
	}
	return o.Nvme, true
}

// HasNvme returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskInfo) HasNvme() bool {
	if o != nil && o.Nvme != nil {
		return true
	}

	return false
}

// SetNvme gets a reference to the given VcenterVmHardwareNvmeAddressInfo and assigns it to the Nvme field.
func (o *VcenterVmHardwareDiskInfo) SetNvme(v VcenterVmHardwareNvmeAddressInfo) {
	o.Nvme = &v
}

// GetBacking returns the Backing field value
func (o *VcenterVmHardwareDiskInfo) GetBacking() VcenterVmHardwareDiskBackingInfo {
	if o == nil {
		var ret VcenterVmHardwareDiskBackingInfo
		return ret
	}

	return o.Backing
}

// GetBackingOk returns a tuple with the Backing field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskInfo) GetBackingOk() (*VcenterVmHardwareDiskBackingInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Backing, true
}

// SetBacking sets field value
func (o *VcenterVmHardwareDiskInfo) SetBacking(v VcenterVmHardwareDiskBackingInfo) {
	o.Backing = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskInfo) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskInfo) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskInfo) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *VcenterVmHardwareDiskInfo) SetCapacity(v int64) {
	o.Capacity = &v
}

func (o VcenterVmHardwareDiskInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
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
	if true {
		toSerialize["backing"] = o.Backing
	}
	if o.Capacity != nil {
		toSerialize["capacity"] = o.Capacity
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareDiskInfo struct {
	value *VcenterVmHardwareDiskInfo
	isSet bool
}

func (v NullableVcenterVmHardwareDiskInfo) Get() *VcenterVmHardwareDiskInfo {
	return v.value
}

func (v *NullableVcenterVmHardwareDiskInfo) Set(val *VcenterVmHardwareDiskInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareDiskInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareDiskInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareDiskInfo(val *VcenterVmHardwareDiskInfo) *NullableVcenterVmHardwareDiskInfo {
	return &NullableVcenterVmHardwareDiskInfo{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareDiskInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareDiskInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


