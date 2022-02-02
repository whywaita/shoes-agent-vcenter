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

// VcenterTaggingAssociationsLastIterationStatus The last status for the iterator. A field of this type is returned as part of the result and indicates to the caller of the API whether it can continue to make requests for more data.   The last status only reports on the state of the iteration at the time data was last returned. As a result, it not does guarantee if the next call will succeed in getting more data or not.    Failures to retrieve results will be returned as Error responses. These last statuses are only returned when the iterator is operating as expected. 
type VcenterTaggingAssociationsLastIterationStatus string

// List of VcenterTaggingAssociationsLastIterationStatus
const (
	VCENTERTAGGINGASSOCIATIONSLASTITERATIONSTATUS_READY VcenterTaggingAssociationsLastIterationStatus = "READY"
	VCENTERTAGGINGASSOCIATIONSLASTITERATIONSTATUS_END_OF_DATA VcenterTaggingAssociationsLastIterationStatus = "END_OF_DATA"
)

// All allowed values of VcenterTaggingAssociationsLastIterationStatus enum
var AllowedVcenterTaggingAssociationsLastIterationStatusEnumValues = []VcenterTaggingAssociationsLastIterationStatus{
	"READY",
	"END_OF_DATA",
}

func (v *VcenterTaggingAssociationsLastIterationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterTaggingAssociationsLastIterationStatus(value)
	for _, existing := range AllowedVcenterTaggingAssociationsLastIterationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterTaggingAssociationsLastIterationStatus", value)
}

// NewVcenterTaggingAssociationsLastIterationStatusFromValue returns a pointer to a valid VcenterTaggingAssociationsLastIterationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterTaggingAssociationsLastIterationStatusFromValue(v string) (*VcenterTaggingAssociationsLastIterationStatus, error) {
	ev := VcenterTaggingAssociationsLastIterationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterTaggingAssociationsLastIterationStatus: valid values are %v", v, AllowedVcenterTaggingAssociationsLastIterationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterTaggingAssociationsLastIterationStatus) IsValid() bool {
	for _, existing := range AllowedVcenterTaggingAssociationsLastIterationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VcenterTaggingAssociationsLastIterationStatus value
func (v VcenterTaggingAssociationsLastIterationStatus) Ptr() *VcenterTaggingAssociationsLastIterationStatus {
	return &v
}

type NullableVcenterTaggingAssociationsLastIterationStatus struct {
	value *VcenterTaggingAssociationsLastIterationStatus
	isSet bool
}

func (v NullableVcenterTaggingAssociationsLastIterationStatus) Get() *VcenterTaggingAssociationsLastIterationStatus {
	return v.value
}

func (v *NullableVcenterTaggingAssociationsLastIterationStatus) Set(val *VcenterTaggingAssociationsLastIterationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTaggingAssociationsLastIterationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTaggingAssociationsLastIterationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTaggingAssociationsLastIterationStatus(val *VcenterTaggingAssociationsLastIterationStatus) *NullableVcenterTaggingAssociationsLastIterationStatus {
	return &NullableVcenterTaggingAssociationsLastIterationStatus{value: val, isSet: true}
}

func (v NullableVcenterTaggingAssociationsLastIterationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTaggingAssociationsLastIterationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

