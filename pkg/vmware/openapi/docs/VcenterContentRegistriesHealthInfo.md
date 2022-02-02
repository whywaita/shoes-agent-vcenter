# VcenterContentRegistriesHealthInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**VcenterContentRegistriesHealthStatus**](VcenterContentRegistriesHealthStatus.md) |  | 
**Details** | Pointer to [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | [optional] 

## Methods

### NewVcenterContentRegistriesHealthInfo

`func NewVcenterContentRegistriesHealthInfo(status VcenterContentRegistriesHealthStatus, ) *VcenterContentRegistriesHealthInfo`

NewVcenterContentRegistriesHealthInfo instantiates a new VcenterContentRegistriesHealthInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterContentRegistriesHealthInfoWithDefaults

`func NewVcenterContentRegistriesHealthInfoWithDefaults() *VcenterContentRegistriesHealthInfo`

NewVcenterContentRegistriesHealthInfoWithDefaults instantiates a new VcenterContentRegistriesHealthInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VcenterContentRegistriesHealthInfo) GetStatus() VcenterContentRegistriesHealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VcenterContentRegistriesHealthInfo) GetStatusOk() (*VcenterContentRegistriesHealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VcenterContentRegistriesHealthInfo) SetStatus(v VcenterContentRegistriesHealthStatus)`

SetStatus sets Status field to given value.


### GetDetails

`func (o *VcenterContentRegistriesHealthInfo) GetDetails() VapiStdLocalizableMessage`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VcenterContentRegistriesHealthInfo) GetDetailsOk() (*VapiStdLocalizableMessage, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VcenterContentRegistriesHealthInfo) SetDetails(v VapiStdLocalizableMessage)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VcenterContentRegistriesHealthInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


