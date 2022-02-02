# VapiStdLocalizationParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S** | Pointer to **string** | {@term String} value associated with the parameter. | [optional] 
**Dt** | Pointer to **time.Time** | Date and time value associated with the parameter. Use the {@name #format} {@term field} to specify date and time display style. | [optional] 
**I** | Pointer to **int64** | {@term long} value associated with the parameter. | [optional] 
**D** | Pointer to **float64** | The {@term double} value associated with the parameter. The number of displayed fractional digits is changed via {@name #precision} {@term field}. | [optional] 
**L** | Pointer to [**VapiStdNestedLocalizableMessage**](VapiStdNestedLocalizableMessage.md) |  | [optional] 
**Format** | Pointer to [**VapiStdLocalizationParamDateTimeFormat**](VapiStdLocalizationParamDateTimeFormat.md) |  | [optional] 
**Precision** | Pointer to **int64** | Number of fractional digits to include in formatted {@term double} value. | [optional] 

## Methods

### NewVapiStdLocalizationParam

`func NewVapiStdLocalizationParam() *VapiStdLocalizationParam`

NewVapiStdLocalizationParam instantiates a new VapiStdLocalizationParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVapiStdLocalizationParamWithDefaults

`func NewVapiStdLocalizationParamWithDefaults() *VapiStdLocalizationParam`

NewVapiStdLocalizationParamWithDefaults instantiates a new VapiStdLocalizationParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS

`func (o *VapiStdLocalizationParam) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *VapiStdLocalizationParam) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *VapiStdLocalizationParam) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *VapiStdLocalizationParam) HasS() bool`

HasS returns a boolean if a field has been set.

### GetDt

`func (o *VapiStdLocalizationParam) GetDt() time.Time`

GetDt returns the Dt field if non-nil, zero value otherwise.

### GetDtOk

`func (o *VapiStdLocalizationParam) GetDtOk() (*time.Time, bool)`

GetDtOk returns a tuple with the Dt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDt

`func (o *VapiStdLocalizationParam) SetDt(v time.Time)`

SetDt sets Dt field to given value.

### HasDt

`func (o *VapiStdLocalizationParam) HasDt() bool`

HasDt returns a boolean if a field has been set.

### GetI

`func (o *VapiStdLocalizationParam) GetI() int64`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *VapiStdLocalizationParam) GetIOk() (*int64, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *VapiStdLocalizationParam) SetI(v int64)`

SetI sets I field to given value.

### HasI

`func (o *VapiStdLocalizationParam) HasI() bool`

HasI returns a boolean if a field has been set.

### GetD

`func (o *VapiStdLocalizationParam) GetD() float64`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *VapiStdLocalizationParam) GetDOk() (*float64, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *VapiStdLocalizationParam) SetD(v float64)`

SetD sets D field to given value.

### HasD

`func (o *VapiStdLocalizationParam) HasD() bool`

HasD returns a boolean if a field has been set.

### GetL

`func (o *VapiStdLocalizationParam) GetL() VapiStdNestedLocalizableMessage`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *VapiStdLocalizationParam) GetLOk() (*VapiStdNestedLocalizableMessage, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *VapiStdLocalizationParam) SetL(v VapiStdNestedLocalizableMessage)`

SetL sets L field to given value.

### HasL

`func (o *VapiStdLocalizationParam) HasL() bool`

HasL returns a boolean if a field has been set.

### GetFormat

`func (o *VapiStdLocalizationParam) GetFormat() VapiStdLocalizationParamDateTimeFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *VapiStdLocalizationParam) GetFormatOk() (*VapiStdLocalizationParamDateTimeFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *VapiStdLocalizationParam) SetFormat(v VapiStdLocalizationParamDateTimeFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *VapiStdLocalizationParam) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPrecision

`func (o *VapiStdLocalizationParam) GetPrecision() int64`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *VapiStdLocalizationParam) GetPrecisionOk() (*int64, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *VapiStdLocalizationParam) SetPrecision(v int64)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *VapiStdLocalizationParam) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


