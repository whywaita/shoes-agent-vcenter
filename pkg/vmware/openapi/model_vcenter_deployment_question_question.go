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

// VcenterDeploymentQuestionQuestion struct for VcenterDeploymentQuestionQuestion
type VcenterDeploymentQuestionQuestion struct {
	// Id of the question raised.
	Id string `json:"id"`
	Question VapiStdLocalizableMessage `json:"question"`
	Type VcenterDeploymentQuestionQuestionType `json:"type"`
	// Default answer value.
	DefaultAnswer string `json:"default_answer"`
	// Possible answers values.
	PossibleAnswers []string `json:"possible_answers"`
}

// NewVcenterDeploymentQuestionQuestion instantiates a new VcenterDeploymentQuestionQuestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentQuestionQuestion(id string, question VapiStdLocalizableMessage, type_ VcenterDeploymentQuestionQuestionType, defaultAnswer string, possibleAnswers []string) *VcenterDeploymentQuestionQuestion {
	this := VcenterDeploymentQuestionQuestion{}
	this.Id = id
	this.Question = question
	this.Type = type_
	this.DefaultAnswer = defaultAnswer
	this.PossibleAnswers = possibleAnswers
	return &this
}

// NewVcenterDeploymentQuestionQuestionWithDefaults instantiates a new VcenterDeploymentQuestionQuestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentQuestionQuestionWithDefaults() *VcenterDeploymentQuestionQuestion {
	this := VcenterDeploymentQuestionQuestion{}
	return &this
}

// GetId returns the Id field value
func (o *VcenterDeploymentQuestionQuestion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentQuestionQuestion) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VcenterDeploymentQuestionQuestion) SetId(v string) {
	o.Id = v
}

// GetQuestion returns the Question field value
func (o *VcenterDeploymentQuestionQuestion) GetQuestion() VapiStdLocalizableMessage {
	if o == nil {
		var ret VapiStdLocalizableMessage
		return ret
	}

	return o.Question
}

// GetQuestionOk returns a tuple with the Question field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentQuestionQuestion) GetQuestionOk() (*VapiStdLocalizableMessage, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Question, true
}

// SetQuestion sets field value
func (o *VcenterDeploymentQuestionQuestion) SetQuestion(v VapiStdLocalizableMessage) {
	o.Question = v
}

// GetType returns the Type field value
func (o *VcenterDeploymentQuestionQuestion) GetType() VcenterDeploymentQuestionQuestionType {
	if o == nil {
		var ret VcenterDeploymentQuestionQuestionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentQuestionQuestion) GetTypeOk() (*VcenterDeploymentQuestionQuestionType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterDeploymentQuestionQuestion) SetType(v VcenterDeploymentQuestionQuestionType) {
	o.Type = v
}

// GetDefaultAnswer returns the DefaultAnswer field value
func (o *VcenterDeploymentQuestionQuestion) GetDefaultAnswer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultAnswer
}

// GetDefaultAnswerOk returns a tuple with the DefaultAnswer field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentQuestionQuestion) GetDefaultAnswerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefaultAnswer, true
}

// SetDefaultAnswer sets field value
func (o *VcenterDeploymentQuestionQuestion) SetDefaultAnswer(v string) {
	o.DefaultAnswer = v
}

// GetPossibleAnswers returns the PossibleAnswers field value
func (o *VcenterDeploymentQuestionQuestion) GetPossibleAnswers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PossibleAnswers
}

// GetPossibleAnswersOk returns a tuple with the PossibleAnswers field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentQuestionQuestion) GetPossibleAnswersOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PossibleAnswers, true
}

// SetPossibleAnswers sets field value
func (o *VcenterDeploymentQuestionQuestion) SetPossibleAnswers(v []string) {
	o.PossibleAnswers = v
}

func (o VcenterDeploymentQuestionQuestion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["question"] = o.Question
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["default_answer"] = o.DefaultAnswer
	}
	if true {
		toSerialize["possible_answers"] = o.PossibleAnswers
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentQuestionQuestion struct {
	value *VcenterDeploymentQuestionQuestion
	isSet bool
}

func (v NullableVcenterDeploymentQuestionQuestion) Get() *VcenterDeploymentQuestionQuestion {
	return v.value
}

func (v *NullableVcenterDeploymentQuestionQuestion) Set(val *VcenterDeploymentQuestionQuestion) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentQuestionQuestion) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentQuestionQuestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentQuestionQuestion(val *VcenterDeploymentQuestionQuestion) *NullableVcenterDeploymentQuestionQuestion {
	return &NullableVcenterDeploymentQuestionQuestion{value: val, isSet: true}
}

func (v NullableVcenterDeploymentQuestionQuestion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentQuestionQuestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

