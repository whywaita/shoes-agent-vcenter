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

// VcenterDatacenterInfo struct for VcenterDatacenterInfo
type VcenterDatacenterInfo struct {
	// The name of the datacenter.
	Name string `json:"name"`
	// The root datastore folder associated with the datacenter. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder.
	DatastoreFolder string `json:"datastore_folder"`
	// The root host and cluster folder associated with the datacenter. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder.
	HostFolder string `json:"host_folder"`
	// The root network folder associated with the datacenter. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder.
	NetworkFolder string `json:"network_folder"`
	// The root virtual machine folder associated with the datacenter. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder.
	VmFolder string `json:"vm_folder"`
}

// NewVcenterDatacenterInfo instantiates a new VcenterDatacenterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDatacenterInfo(name string, datastoreFolder string, hostFolder string, networkFolder string, vmFolder string) *VcenterDatacenterInfo {
	this := VcenterDatacenterInfo{}
	this.Name = name
	this.DatastoreFolder = datastoreFolder
	this.HostFolder = hostFolder
	this.NetworkFolder = networkFolder
	this.VmFolder = vmFolder
	return &this
}

// NewVcenterDatacenterInfoWithDefaults instantiates a new VcenterDatacenterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDatacenterInfoWithDefaults() *VcenterDatacenterInfo {
	this := VcenterDatacenterInfo{}
	return &this
}

// GetName returns the Name field value
func (o *VcenterDatacenterInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VcenterDatacenterInfo) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VcenterDatacenterInfo) SetName(v string) {
	o.Name = v
}

// GetDatastoreFolder returns the DatastoreFolder field value
func (o *VcenterDatacenterInfo) GetDatastoreFolder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatastoreFolder
}

// GetDatastoreFolderOk returns a tuple with the DatastoreFolder field value
// and a boolean to check if the value has been set.
func (o *VcenterDatacenterInfo) GetDatastoreFolderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DatastoreFolder, true
}

// SetDatastoreFolder sets field value
func (o *VcenterDatacenterInfo) SetDatastoreFolder(v string) {
	o.DatastoreFolder = v
}

// GetHostFolder returns the HostFolder field value
func (o *VcenterDatacenterInfo) GetHostFolder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostFolder
}

// GetHostFolderOk returns a tuple with the HostFolder field value
// and a boolean to check if the value has been set.
func (o *VcenterDatacenterInfo) GetHostFolderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HostFolder, true
}

// SetHostFolder sets field value
func (o *VcenterDatacenterInfo) SetHostFolder(v string) {
	o.HostFolder = v
}

// GetNetworkFolder returns the NetworkFolder field value
func (o *VcenterDatacenterInfo) GetNetworkFolder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkFolder
}

// GetNetworkFolderOk returns a tuple with the NetworkFolder field value
// and a boolean to check if the value has been set.
func (o *VcenterDatacenterInfo) GetNetworkFolderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetworkFolder, true
}

// SetNetworkFolder sets field value
func (o *VcenterDatacenterInfo) SetNetworkFolder(v string) {
	o.NetworkFolder = v
}

// GetVmFolder returns the VmFolder field value
func (o *VcenterDatacenterInfo) GetVmFolder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmFolder
}

// GetVmFolderOk returns a tuple with the VmFolder field value
// and a boolean to check if the value has been set.
func (o *VcenterDatacenterInfo) GetVmFolderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VmFolder, true
}

// SetVmFolder sets field value
func (o *VcenterDatacenterInfo) SetVmFolder(v string) {
	o.VmFolder = v
}

func (o VcenterDatacenterInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["datastore_folder"] = o.DatastoreFolder
	}
	if true {
		toSerialize["host_folder"] = o.HostFolder
	}
	if true {
		toSerialize["network_folder"] = o.NetworkFolder
	}
	if true {
		toSerialize["vm_folder"] = o.VmFolder
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDatacenterInfo struct {
	value *VcenterDatacenterInfo
	isSet bool
}

func (v NullableVcenterDatacenterInfo) Get() *VcenterDatacenterInfo {
	return v.value
}

func (v *NullableVcenterDatacenterInfo) Set(val *VcenterDatacenterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDatacenterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDatacenterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDatacenterInfo(val *VcenterDatacenterInfo) *NullableVcenterDatacenterInfo {
	return &NullableVcenterDatacenterInfo{value: val, isSet: true}
}

func (v NullableVcenterDatacenterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDatacenterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

