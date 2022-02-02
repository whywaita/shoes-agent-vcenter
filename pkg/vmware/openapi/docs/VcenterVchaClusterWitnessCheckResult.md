# VcenterVchaClusterWitnessCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warnings** | [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | A list of problems which may require attention, but which are not fatal. | 
**Errors** | [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | A list of problems which are fatal to the operation and the operation will fail. | 

## Methods

### NewVcenterVchaClusterWitnessCheckResult

`func NewVcenterVchaClusterWitnessCheckResult(warnings []VapiStdLocalizableMessage, errors []VapiStdLocalizableMessage, ) *VcenterVchaClusterWitnessCheckResult`

NewVcenterVchaClusterWitnessCheckResult instantiates a new VcenterVchaClusterWitnessCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterWitnessCheckResultWithDefaults

`func NewVcenterVchaClusterWitnessCheckResultWithDefaults() *VcenterVchaClusterWitnessCheckResult`

NewVcenterVchaClusterWitnessCheckResultWithDefaults instantiates a new VcenterVchaClusterWitnessCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarnings

`func (o *VcenterVchaClusterWitnessCheckResult) GetWarnings() []VapiStdLocalizableMessage`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *VcenterVchaClusterWitnessCheckResult) GetWarningsOk() (*[]VapiStdLocalizableMessage, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *VcenterVchaClusterWitnessCheckResult) SetWarnings(v []VapiStdLocalizableMessage)`

SetWarnings sets Warnings field to given value.


### GetErrors

`func (o *VcenterVchaClusterWitnessCheckResult) GetErrors() []VapiStdLocalizableMessage`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *VcenterVchaClusterWitnessCheckResult) GetErrorsOk() (*[]VapiStdLocalizableMessage, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *VcenterVchaClusterWitnessCheckResult) SetErrors(v []VapiStdLocalizableMessage)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


