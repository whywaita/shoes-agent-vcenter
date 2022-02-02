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

// VcenterDatastoreSummary struct for VcenterDatastoreSummary
type VcenterDatastoreSummary struct {
	// Identifier of the datastore. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datastore. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datastore.
	Datastore string `json:"datastore"`
	// Name of the datastore.
	Name string `json:"name"`
	Type VcenterDatastoreType `json:"type"`
	// Available space of this datastore, in bytes.   The server periodically updates this value.  This field will be unset if the available space of this datastore is not known.
	FreeSpace *int64 `json:"free_space,omitempty"`
	// Capacity of this datastore, in bytes.   The server periodically updates this value.  This field will be unset if the capacity of this datastore is not known.
	Capacity *int64 `json:"capacity,omitempty"`
}

// NewVcenterDatastoreSummary instantiates a new VcenterDatastoreSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDatastoreSummary(datastore string, name string, type_ VcenterDatastoreType) *VcenterDatastoreSummary {
	this := VcenterDatastoreSummary{}
	this.Datastore = datastore
	this.Name = name
	this.Type = type_
	return &this
}

// NewVcenterDatastoreSummaryWithDefaults instantiates a new VcenterDatastoreSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDatastoreSummaryWithDefaults() *VcenterDatastoreSummary {
	this := VcenterDatastoreSummary{}
	return &this
}

// GetDatastore returns the Datastore field value
func (o *VcenterDatastoreSummary) GetDatastore() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value
// and a boolean to check if the value has been set.
func (o *VcenterDatastoreSummary) GetDatastoreOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Datastore, true
}

// SetDatastore sets field value
func (o *VcenterDatastoreSummary) SetDatastore(v string) {
	o.Datastore = v
}

// GetName returns the Name field value
func (o *VcenterDatastoreSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VcenterDatastoreSummary) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VcenterDatastoreSummary) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *VcenterDatastoreSummary) GetType() VcenterDatastoreType {
	if o == nil {
		var ret VcenterDatastoreType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterDatastoreSummary) GetTypeOk() (*VcenterDatastoreType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterDatastoreSummary) SetType(v VcenterDatastoreType) {
	o.Type = v
}

// GetFreeSpace returns the FreeSpace field value if set, zero value otherwise.
func (o *VcenterDatastoreSummary) GetFreeSpace() int64 {
	if o == nil || o.FreeSpace == nil {
		var ret int64
		return ret
	}
	return *o.FreeSpace
}

// GetFreeSpaceOk returns a tuple with the FreeSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDatastoreSummary) GetFreeSpaceOk() (*int64, bool) {
	if o == nil || o.FreeSpace == nil {
		return nil, false
	}
	return o.FreeSpace, true
}

// HasFreeSpace returns a boolean if a field has been set.
func (o *VcenterDatastoreSummary) HasFreeSpace() bool {
	if o != nil && o.FreeSpace != nil {
		return true
	}

	return false
}

// SetFreeSpace gets a reference to the given int64 and assigns it to the FreeSpace field.
func (o *VcenterDatastoreSummary) SetFreeSpace(v int64) {
	o.FreeSpace = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VcenterDatastoreSummary) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDatastoreSummary) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VcenterDatastoreSummary) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *VcenterDatastoreSummary) SetCapacity(v int64) {
	o.Capacity = &v
}

func (o VcenterDatastoreSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["datastore"] = o.Datastore
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.FreeSpace != nil {
		toSerialize["free_space"] = o.FreeSpace
	}
	if o.Capacity != nil {
		toSerialize["capacity"] = o.Capacity
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDatastoreSummary struct {
	value *VcenterDatastoreSummary
	isSet bool
}

func (v NullableVcenterDatastoreSummary) Get() *VcenterDatastoreSummary {
	return v.value
}

func (v *NullableVcenterDatastoreSummary) Set(val *VcenterDatastoreSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDatastoreSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDatastoreSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDatastoreSummary(val *VcenterDatastoreSummary) *NullableVcenterDatastoreSummary {
	return &NullableVcenterDatastoreSummary{value: val, isSet: true}
}

func (v NullableVcenterDatastoreSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDatastoreSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

