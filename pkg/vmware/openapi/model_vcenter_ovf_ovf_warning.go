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

// VcenterOvfOvfWarning struct for VcenterOvfOvfWarning
type VcenterOvfOvfWarning struct {
	Category VcenterOvfOvfMessageCategory `json:"category"`
	// {@term List} of parse issues (see {@link ParseIssue}).
	Issues *[]VcenterOvfParseIssue `json:"issues,omitempty"`
	// The name of input parameter.
	Name *string `json:"name,omitempty"`
	// The value of input parameter.
	Value *string `json:"value,omitempty"`
	Message *VapiStdLocalizableMessage `json:"message,omitempty"`
	// Represents a server {@link Error}.
	Error *map[string]interface{} `json:"error,omitempty"`
}

// NewVcenterOvfOvfWarning instantiates a new VcenterOvfOvfWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterOvfOvfWarning(category VcenterOvfOvfMessageCategory) *VcenterOvfOvfWarning {
	this := VcenterOvfOvfWarning{}
	this.Category = category
	return &this
}

// NewVcenterOvfOvfWarningWithDefaults instantiates a new VcenterOvfOvfWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterOvfOvfWarningWithDefaults() *VcenterOvfOvfWarning {
	this := VcenterOvfOvfWarning{}
	return &this
}

// GetCategory returns the Category field value
func (o *VcenterOvfOvfWarning) GetCategory() VcenterOvfOvfMessageCategory {
	if o == nil {
		var ret VcenterOvfOvfMessageCategory
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *VcenterOvfOvfWarning) GetCategoryOk() (*VcenterOvfOvfMessageCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *VcenterOvfOvfWarning) SetCategory(v VcenterOvfOvfMessageCategory) {
	o.Category = v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *VcenterOvfOvfWarning) GetIssues() []VcenterOvfParseIssue {
	if o == nil || o.Issues == nil {
		var ret []VcenterOvfParseIssue
		return ret
	}
	return *o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterOvfOvfWarning) GetIssuesOk() (*[]VcenterOvfParseIssue, bool) {
	if o == nil || o.Issues == nil {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *VcenterOvfOvfWarning) HasIssues() bool {
	if o != nil && o.Issues != nil {
		return true
	}

	return false
}

// SetIssues gets a reference to the given []VcenterOvfParseIssue and assigns it to the Issues field.
func (o *VcenterOvfOvfWarning) SetIssues(v []VcenterOvfParseIssue) {
	o.Issues = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VcenterOvfOvfWarning) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterOvfOvfWarning) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VcenterOvfOvfWarning) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VcenterOvfOvfWarning) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VcenterOvfOvfWarning) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterOvfOvfWarning) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VcenterOvfOvfWarning) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *VcenterOvfOvfWarning) SetValue(v string) {
	o.Value = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *VcenterOvfOvfWarning) GetMessage() VapiStdLocalizableMessage {
	if o == nil || o.Message == nil {
		var ret VapiStdLocalizableMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterOvfOvfWarning) GetMessageOk() (*VapiStdLocalizableMessage, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *VcenterOvfOvfWarning) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given VapiStdLocalizableMessage and assigns it to the Message field.
func (o *VcenterOvfOvfWarning) SetMessage(v VapiStdLocalizableMessage) {
	o.Message = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *VcenterOvfOvfWarning) GetError() map[string]interface{} {
	if o == nil || o.Error == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterOvfOvfWarning) GetErrorOk() (*map[string]interface{}, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *VcenterOvfOvfWarning) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given map[string]interface{} and assigns it to the Error field.
func (o *VcenterOvfOvfWarning) SetError(v map[string]interface{}) {
	o.Error = &v
}

func (o VcenterOvfOvfWarning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["category"] = o.Category
	}
	if o.Issues != nil {
		toSerialize["issues"] = o.Issues
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterOvfOvfWarning struct {
	value *VcenterOvfOvfWarning
	isSet bool
}

func (v NullableVcenterOvfOvfWarning) Get() *VcenterOvfOvfWarning {
	return v.value
}

func (v *NullableVcenterOvfOvfWarning) Set(val *VcenterOvfOvfWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterOvfOvfWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterOvfOvfWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterOvfOvfWarning(val *VcenterOvfOvfWarning) *NullableVcenterOvfOvfWarning {
	return &NullableVcenterOvfOvfWarning{value: val, isSet: true}
}

func (v NullableVcenterOvfOvfWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterOvfOvfWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


