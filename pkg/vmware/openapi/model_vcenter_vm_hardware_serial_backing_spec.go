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

// VcenterVmHardwareSerialBackingSpec struct for VcenterVmHardwareSerialBackingSpec
type VcenterVmHardwareSerialBackingSpec struct {
	Type VcenterVmHardwareSerialBackingType `json:"type"`
	// Path of the file backing the virtual serial port. This field is optional and it is only relevant when the value of Serial.BackingSpec.type is FILE.
	File *string `json:"file,omitempty"`
	// Name of the device backing the virtual serial port.    If unset, the virtual serial port will be configured to automatically detect a suitable host device.
	HostDevice *string `json:"host_device,omitempty"`
	// Name of the pipe backing the virtual serial port. This field is optional and it is only relevant when the value of Serial.BackingSpec.type is one of PIPE_SERVER or PIPE_CLIENT.
	Pipe *string `json:"pipe,omitempty"`
	// Flag that enables optimized data transfer over the pipe. When the value is true, the host buffers data to prevent data overrun. This allows the virtual machine to read all of the data transferred over the pipe with no data loss. If unset, defaults to false.
	NoRxLoss *bool `json:"no_rx_loss,omitempty"`
	// URI specifying the location of the network service backing the virtual serial port.     - If Serial.BackingSpec.type is NETWORK_SERVER, this field is the location used by clients to connect to this server. The hostname part of the URI should either be empty or should specify the address of the host on which the virtual machine is running.    - If Serial.BackingSpec.type is NETWORK_CLIENT, this field is the location used by the virtual machine to connect to the remote server.   This field is optional and it is only relevant when the value of Serial.BackingSpec.type is one of NETWORK_SERVER or NETWORK_CLIENT.
	NetworkLocation *string `json:"network_location,omitempty"`
	// Proxy service that provides network access to the network backing. If set, the virtual machine initiates a connection with the proxy service and forwards the traffic to the proxy. If unset, no proxy service should be used.
	Proxy *string `json:"proxy,omitempty"`
}

// NewVcenterVmHardwareSerialBackingSpec instantiates a new VcenterVmHardwareSerialBackingSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareSerialBackingSpec(type_ VcenterVmHardwareSerialBackingType) *VcenterVmHardwareSerialBackingSpec {
	this := VcenterVmHardwareSerialBackingSpec{}
	this.Type = type_
	return &this
}

// NewVcenterVmHardwareSerialBackingSpecWithDefaults instantiates a new VcenterVmHardwareSerialBackingSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareSerialBackingSpecWithDefaults() *VcenterVmHardwareSerialBackingSpec {
	this := VcenterVmHardwareSerialBackingSpec{}
	return &this
}

// GetType returns the Type field value
func (o *VcenterVmHardwareSerialBackingSpec) GetType() VcenterVmHardwareSerialBackingType {
	if o == nil {
		var ret VcenterVmHardwareSerialBackingType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareSerialBackingSpec) GetTypeOk() (*VcenterVmHardwareSerialBackingType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterVmHardwareSerialBackingSpec) SetType(v VcenterVmHardwareSerialBackingType) {
	o.Type = v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *VcenterVmHardwareSerialBackingSpec) GetFile() string {
	if o == nil || o.File == nil {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareSerialBackingSpec) GetFileOk() (*string, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *VcenterVmHardwareSerialBackingSpec) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *VcenterVmHardwareSerialBackingSpec) SetFile(v string) {
	o.File = &v
}

// GetHostDevice returns the HostDevice field value if set, zero value otherwise.
func (o *VcenterVmHardwareSerialBackingSpec) GetHostDevice() string {
	if o == nil || o.HostDevice == nil {
		var ret string
		return ret
	}
	return *o.HostDevice
}

// GetHostDeviceOk returns a tuple with the HostDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareSerialBackingSpec) GetHostDeviceOk() (*string, bool) {
	if o == nil || o.HostDevice == nil {
		return nil, false
	}
	return o.HostDevice, true
}

// HasHostDevice returns a boolean if a field has been set.
func (o *VcenterVmHardwareSerialBackingSpec) HasHostDevice() bool {
	if o != nil && o.HostDevice != nil {
		return true
	}

	return false
}

// SetHostDevice gets a reference to the given string and assigns it to the HostDevice field.
func (o *VcenterVmHardwareSerialBackingSpec) SetHostDevice(v string) {
	o.HostDevice = &v
}

// GetPipe returns the Pipe field value if set, zero value otherwise.
func (o *VcenterVmHardwareSerialBackingSpec) GetPipe() string {
	if o == nil || o.Pipe == nil {
		var ret string
		return ret
	}
	return *o.Pipe
}

// GetPipeOk returns a tuple with the Pipe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareSerialBackingSpec) GetPipeOk() (*string, bool) {
	if o == nil || o.Pipe == nil {
		return nil, false
	}
	return o.Pipe, true
}

// HasPipe returns a boolean if a field has been set.
func (o *VcenterVmHardwareSerialBackingSpec) HasPipe() bool {
	if o != nil && o.Pipe != nil {
		return true
	}

	return false
}

// SetPipe gets a reference to the given string and assigns it to the Pipe field.
func (o *VcenterVmHardwareSerialBackingSpec) SetPipe(v string) {
	o.Pipe = &v
}

// GetNoRxLoss returns the NoRxLoss field value if set, zero value otherwise.
func (o *VcenterVmHardwareSerialBackingSpec) GetNoRxLoss() bool {
	if o == nil || o.NoRxLoss == nil {
		var ret bool
		return ret
	}
	return *o.NoRxLoss
}

// GetNoRxLossOk returns a tuple with the NoRxLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareSerialBackingSpec) GetNoRxLossOk() (*bool, bool) {
	if o == nil || o.NoRxLoss == nil {
		return nil, false
	}
	return o.NoRxLoss, true
}

// HasNoRxLoss returns a boolean if a field has been set.
func (o *VcenterVmHardwareSerialBackingSpec) HasNoRxLoss() bool {
	if o != nil && o.NoRxLoss != nil {
		return true
	}

	return false
}

// SetNoRxLoss gets a reference to the given bool and assigns it to the NoRxLoss field.
func (o *VcenterVmHardwareSerialBackingSpec) SetNoRxLoss(v bool) {
	o.NoRxLoss = &v
}

// GetNetworkLocation returns the NetworkLocation field value if set, zero value otherwise.
func (o *VcenterVmHardwareSerialBackingSpec) GetNetworkLocation() string {
	if o == nil || o.NetworkLocation == nil {
		var ret string
		return ret
	}
	return *o.NetworkLocation
}

// GetNetworkLocationOk returns a tuple with the NetworkLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareSerialBackingSpec) GetNetworkLocationOk() (*string, bool) {
	if o == nil || o.NetworkLocation == nil {
		return nil, false
	}
	return o.NetworkLocation, true
}

// HasNetworkLocation returns a boolean if a field has been set.
func (o *VcenterVmHardwareSerialBackingSpec) HasNetworkLocation() bool {
	if o != nil && o.NetworkLocation != nil {
		return true
	}

	return false
}

// SetNetworkLocation gets a reference to the given string and assigns it to the NetworkLocation field.
func (o *VcenterVmHardwareSerialBackingSpec) SetNetworkLocation(v string) {
	o.NetworkLocation = &v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *VcenterVmHardwareSerialBackingSpec) GetProxy() string {
	if o == nil || o.Proxy == nil {
		var ret string
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareSerialBackingSpec) GetProxyOk() (*string, bool) {
	if o == nil || o.Proxy == nil {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *VcenterVmHardwareSerialBackingSpec) HasProxy() bool {
	if o != nil && o.Proxy != nil {
		return true
	}

	return false
}

// SetProxy gets a reference to the given string and assigns it to the Proxy field.
func (o *VcenterVmHardwareSerialBackingSpec) SetProxy(v string) {
	o.Proxy = &v
}

func (o VcenterVmHardwareSerialBackingSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.HostDevice != nil {
		toSerialize["host_device"] = o.HostDevice
	}
	if o.Pipe != nil {
		toSerialize["pipe"] = o.Pipe
	}
	if o.NoRxLoss != nil {
		toSerialize["no_rx_loss"] = o.NoRxLoss
	}
	if o.NetworkLocation != nil {
		toSerialize["network_location"] = o.NetworkLocation
	}
	if o.Proxy != nil {
		toSerialize["proxy"] = o.Proxy
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareSerialBackingSpec struct {
	value *VcenterVmHardwareSerialBackingSpec
	isSet bool
}

func (v NullableVcenterVmHardwareSerialBackingSpec) Get() *VcenterVmHardwareSerialBackingSpec {
	return v.value
}

func (v *NullableVcenterVmHardwareSerialBackingSpec) Set(val *VcenterVmHardwareSerialBackingSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareSerialBackingSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareSerialBackingSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareSerialBackingSpec(val *VcenterVmHardwareSerialBackingSpec) *NullableVcenterVmHardwareSerialBackingSpec {
	return &NullableVcenterVmHardwareSerialBackingSpec{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareSerialBackingSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareSerialBackingSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


