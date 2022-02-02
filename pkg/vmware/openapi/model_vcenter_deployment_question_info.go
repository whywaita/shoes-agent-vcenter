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

// VcenterDeploymentQuestionInfo struct for VcenterDeploymentQuestionInfo
type VcenterDeploymentQuestionInfo struct {
	// One or more questions raised during the deployment.
	Questions []VcenterDeploymentQuestionQuestion `json:"questions"`
}

// NewVcenterDeploymentQuestionInfo instantiates a new VcenterDeploymentQuestionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentQuestionInfo(questions []VcenterDeploymentQuestionQuestion) *VcenterDeploymentQuestionInfo {
	this := VcenterDeploymentQuestionInfo{}
	this.Questions = questions
	return &this
}

// NewVcenterDeploymentQuestionInfoWithDefaults instantiates a new VcenterDeploymentQuestionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentQuestionInfoWithDefaults() *VcenterDeploymentQuestionInfo {
	this := VcenterDeploymentQuestionInfo{}
	return &this
}

// GetQuestions returns the Questions field value
func (o *VcenterDeploymentQuestionInfo) GetQuestions() []VcenterDeploymentQuestionQuestion {
	if o == nil {
		var ret []VcenterDeploymentQuestionQuestion
		return ret
	}

	return o.Questions
}

// GetQuestionsOk returns a tuple with the Questions field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentQuestionInfo) GetQuestionsOk() (*[]VcenterDeploymentQuestionQuestion, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Questions, true
}

// SetQuestions sets field value
func (o *VcenterDeploymentQuestionInfo) SetQuestions(v []VcenterDeploymentQuestionQuestion) {
	o.Questions = v
}

func (o VcenterDeploymentQuestionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["questions"] = o.Questions
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentQuestionInfo struct {
	value *VcenterDeploymentQuestionInfo
	isSet bool
}

func (v NullableVcenterDeploymentQuestionInfo) Get() *VcenterDeploymentQuestionInfo {
	return v.value
}

func (v *NullableVcenterDeploymentQuestionInfo) Set(val *VcenterDeploymentQuestionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentQuestionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentQuestionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentQuestionInfo(val *VcenterDeploymentQuestionInfo) *NullableVcenterDeploymentQuestionInfo {
	return &NullableVcenterDeploymentQuestionInfo{value: val, isSet: true}
}

func (v NullableVcenterDeploymentQuestionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentQuestionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

