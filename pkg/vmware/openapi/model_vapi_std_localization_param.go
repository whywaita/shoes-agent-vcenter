/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// VapiStdLocalizationParam struct for VapiStdLocalizationParam
type VapiStdLocalizationParam struct {
	// {@term String} value associated with the parameter.
	S *string `json:"s,omitempty"`
	// Date and time value associated with the parameter. Use the {@name #format} {@term field} to specify date and time display style.
	Dt *time.Time `json:"dt,omitempty"`
	// {@term long} value associated with the parameter.
	I *int64 `json:"i,omitempty"`
	// The {@term double} value associated with the parameter. The number of displayed fractional digits is changed via {@name #precision} {@term field}.
	D *float64 `json:"d,omitempty"`
	L *VapiStdNestedLocalizableMessage `json:"l,omitempty"`
	Format *VapiStdLocalizationParamDateTimeFormat `json:"format,omitempty"`
	// Number of fractional digits to include in formatted {@term double} value.
	Precision *int64 `json:"precision,omitempty"`
}

// NewVapiStdLocalizationParam instantiates a new VapiStdLocalizationParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVapiStdLocalizationParam() *VapiStdLocalizationParam {
	this := VapiStdLocalizationParam{}
	return &this
}

// NewVapiStdLocalizationParamWithDefaults instantiates a new VapiStdLocalizationParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVapiStdLocalizationParamWithDefaults() *VapiStdLocalizationParam {
	this := VapiStdLocalizationParam{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *VapiStdLocalizationParam) GetS() string {
	if o == nil || o.S == nil {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdLocalizationParam) GetSOk() (*string, bool) {
	if o == nil || o.S == nil {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *VapiStdLocalizationParam) HasS() bool {
	if o != nil && o.S != nil {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *VapiStdLocalizationParam) SetS(v string) {
	o.S = &v
}

// GetDt returns the Dt field value if set, zero value otherwise.
func (o *VapiStdLocalizationParam) GetDt() time.Time {
	if o == nil || o.Dt == nil {
		var ret time.Time
		return ret
	}
	return *o.Dt
}

// GetDtOk returns a tuple with the Dt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdLocalizationParam) GetDtOk() (*time.Time, bool) {
	if o == nil || o.Dt == nil {
		return nil, false
	}
	return o.Dt, true
}

// HasDt returns a boolean if a field has been set.
func (o *VapiStdLocalizationParam) HasDt() bool {
	if o != nil && o.Dt != nil {
		return true
	}

	return false
}

// SetDt gets a reference to the given time.Time and assigns it to the Dt field.
func (o *VapiStdLocalizationParam) SetDt(v time.Time) {
	o.Dt = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *VapiStdLocalizationParam) GetI() int64 {
	if o == nil || o.I == nil {
		var ret int64
		return ret
	}
	return *o.I
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdLocalizationParam) GetIOk() (*int64, bool) {
	if o == nil || o.I == nil {
		return nil, false
	}
	return o.I, true
}

// HasI returns a boolean if a field has been set.
func (o *VapiStdLocalizationParam) HasI() bool {
	if o != nil && o.I != nil {
		return true
	}

	return false
}

// SetI gets a reference to the given int64 and assigns it to the I field.
func (o *VapiStdLocalizationParam) SetI(v int64) {
	o.I = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *VapiStdLocalizationParam) GetD() float64 {
	if o == nil || o.D == nil {
		var ret float64
		return ret
	}
	return *o.D
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdLocalizationParam) GetDOk() (*float64, bool) {
	if o == nil || o.D == nil {
		return nil, false
	}
	return o.D, true
}

// HasD returns a boolean if a field has been set.
func (o *VapiStdLocalizationParam) HasD() bool {
	if o != nil && o.D != nil {
		return true
	}

	return false
}

// SetD gets a reference to the given float64 and assigns it to the D field.
func (o *VapiStdLocalizationParam) SetD(v float64) {
	o.D = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *VapiStdLocalizationParam) GetL() VapiStdNestedLocalizableMessage {
	if o == nil || o.L == nil {
		var ret VapiStdNestedLocalizableMessage
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdLocalizationParam) GetLOk() (*VapiStdNestedLocalizableMessage, bool) {
	if o == nil || o.L == nil {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *VapiStdLocalizationParam) HasL() bool {
	if o != nil && o.L != nil {
		return true
	}

	return false
}

// SetL gets a reference to the given VapiStdNestedLocalizableMessage and assigns it to the L field.
func (o *VapiStdLocalizationParam) SetL(v VapiStdNestedLocalizableMessage) {
	o.L = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *VapiStdLocalizationParam) GetFormat() VapiStdLocalizationParamDateTimeFormat {
	if o == nil || o.Format == nil {
		var ret VapiStdLocalizationParamDateTimeFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdLocalizationParam) GetFormatOk() (*VapiStdLocalizationParamDateTimeFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *VapiStdLocalizationParam) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given VapiStdLocalizationParamDateTimeFormat and assigns it to the Format field.
func (o *VapiStdLocalizationParam) SetFormat(v VapiStdLocalizationParamDateTimeFormat) {
	o.Format = &v
}

// GetPrecision returns the Precision field value if set, zero value otherwise.
func (o *VapiStdLocalizationParam) GetPrecision() int64 {
	if o == nil || o.Precision == nil {
		var ret int64
		return ret
	}
	return *o.Precision
}

// GetPrecisionOk returns a tuple with the Precision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdLocalizationParam) GetPrecisionOk() (*int64, bool) {
	if o == nil || o.Precision == nil {
		return nil, false
	}
	return o.Precision, true
}

// HasPrecision returns a boolean if a field has been set.
func (o *VapiStdLocalizationParam) HasPrecision() bool {
	if o != nil && o.Precision != nil {
		return true
	}

	return false
}

// SetPrecision gets a reference to the given int64 and assigns it to the Precision field.
func (o *VapiStdLocalizationParam) SetPrecision(v int64) {
	o.Precision = &v
}

func (o VapiStdLocalizationParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.S != nil {
		toSerialize["s"] = o.S
	}
	if o.Dt != nil {
		toSerialize["dt"] = o.Dt
	}
	if o.I != nil {
		toSerialize["i"] = o.I
	}
	if o.D != nil {
		toSerialize["d"] = o.D
	}
	if o.L != nil {
		toSerialize["l"] = o.L
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Precision != nil {
		toSerialize["precision"] = o.Precision
	}
	return json.Marshal(toSerialize)
}

type NullableVapiStdLocalizationParam struct {
	value *VapiStdLocalizationParam
	isSet bool
}

func (v NullableVapiStdLocalizationParam) Get() *VapiStdLocalizationParam {
	return v.value
}

func (v *NullableVapiStdLocalizationParam) Set(val *VapiStdLocalizationParam) {
	v.value = val
	v.isSet = true
}

func (v NullableVapiStdLocalizationParam) IsSet() bool {
	return v.isSet
}

func (v *NullableVapiStdLocalizationParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVapiStdLocalizationParam(val *VapiStdLocalizationParam) *NullableVapiStdLocalizationParam {
	return &NullableVapiStdLocalizationParam{value: val, isSet: true}
}

func (v NullableVapiStdLocalizationParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVapiStdLocalizationParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


