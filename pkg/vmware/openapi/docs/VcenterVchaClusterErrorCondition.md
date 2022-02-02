# VcenterVchaClusterErrorCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | 
**Recommendation** | Pointer to [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | [optional] 

## Methods

### NewVcenterVchaClusterErrorCondition

`func NewVcenterVchaClusterErrorCondition(error_ VapiStdLocalizableMessage, ) *VcenterVchaClusterErrorCondition`

NewVcenterVchaClusterErrorCondition instantiates a new VcenterVchaClusterErrorCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterErrorConditionWithDefaults

`func NewVcenterVchaClusterErrorConditionWithDefaults() *VcenterVchaClusterErrorCondition`

NewVcenterVchaClusterErrorConditionWithDefaults instantiates a new VcenterVchaClusterErrorCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *VcenterVchaClusterErrorCondition) GetError() VapiStdLocalizableMessage`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VcenterVchaClusterErrorCondition) GetErrorOk() (*VapiStdLocalizableMessage, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VcenterVchaClusterErrorCondition) SetError(v VapiStdLocalizableMessage)`

SetError sets Error field to given value.


### GetRecommendation

`func (o *VcenterVchaClusterErrorCondition) GetRecommendation() VapiStdLocalizableMessage`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *VcenterVchaClusterErrorCondition) GetRecommendationOk() (*VapiStdLocalizableMessage, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *VcenterVchaClusterErrorCondition) SetRecommendation(v VapiStdLocalizableMessage)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *VcenterVchaClusterErrorCondition) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


