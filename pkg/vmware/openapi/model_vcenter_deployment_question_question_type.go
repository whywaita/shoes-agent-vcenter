/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// VcenterDeploymentQuestionQuestionType The Question.QuestionType enumerated type defines the type of the question raised.
type VcenterDeploymentQuestionQuestionType string

// List of vcenter.deployment.question.question_type
const (
	VCENTERDEPLOYMENTQUESTIONQUESTIONTYPE_YES_NO VcenterDeploymentQuestionQuestionType = "YES_NO"
	VCENTERDEPLOYMENTQUESTIONQUESTIONTYPE_OK_CANCEL VcenterDeploymentQuestionQuestionType = "OK_CANCEL"
	VCENTERDEPLOYMENTQUESTIONQUESTIONTYPE_ABORT_RETRY_IGNORE VcenterDeploymentQuestionQuestionType = "ABORT_RETRY_IGNORE"
)

// All allowed values of VcenterDeploymentQuestionQuestionType enum
var AllowedVcenterDeploymentQuestionQuestionTypeEnumValues = []VcenterDeploymentQuestionQuestionType{
	"YES_NO",
	"OK_CANCEL",
	"ABORT_RETRY_IGNORE",
}

func (v *VcenterDeploymentQuestionQuestionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterDeploymentQuestionQuestionType(value)
	for _, existing := range AllowedVcenterDeploymentQuestionQuestionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterDeploymentQuestionQuestionType", value)
}

// NewVcenterDeploymentQuestionQuestionTypeFromValue returns a pointer to a valid VcenterDeploymentQuestionQuestionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterDeploymentQuestionQuestionTypeFromValue(v string) (*VcenterDeploymentQuestionQuestionType, error) {
	ev := VcenterDeploymentQuestionQuestionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterDeploymentQuestionQuestionType: valid values are %v", v, AllowedVcenterDeploymentQuestionQuestionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterDeploymentQuestionQuestionType) IsValid() bool {
	for _, existing := range AllowedVcenterDeploymentQuestionQuestionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.deployment.question.question_type value
func (v VcenterDeploymentQuestionQuestionType) Ptr() *VcenterDeploymentQuestionQuestionType {
	return &v
}

type NullableVcenterDeploymentQuestionQuestionType struct {
	value *VcenterDeploymentQuestionQuestionType
	isSet bool
}

func (v NullableVcenterDeploymentQuestionQuestionType) Get() *VcenterDeploymentQuestionQuestionType {
	return v.value
}

func (v *NullableVcenterDeploymentQuestionQuestionType) Set(val *VcenterDeploymentQuestionQuestionType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentQuestionQuestionType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentQuestionQuestionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentQuestionQuestionType(val *VcenterDeploymentQuestionQuestionType) *NullableVcenterDeploymentQuestionQuestionType {
	return &NullableVcenterDeploymentQuestionQuestionType{value: val, isSet: true}
}

func (v NullableVcenterDeploymentQuestionQuestionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentQuestionQuestionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
