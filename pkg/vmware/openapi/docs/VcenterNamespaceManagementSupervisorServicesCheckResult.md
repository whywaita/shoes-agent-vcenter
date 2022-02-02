# VcenterNamespaceManagementSupervisorServicesCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**VcenterNamespaceManagementSupervisorServicesValidationStatus**](VcenterNamespaceManagementSupervisorServicesValidationStatus.md) |  | 
**ContentType** | Pointer to [**VcenterNamespaceManagementSupervisorServicesVersionsContentType**](VcenterNamespaceManagementSupervisorServicesVersionsContentType.md) |  | [optional] 
**VsphereAppsCheckResult** | Pointer to [**VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult**](VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult.md) |  | [optional] 
**WarningMessages** | Pointer to [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | A list of messages indicating why the content was considered valid but contains information that should be reviewed closely. This field is optional and it is only relevant when the value of SupervisorServices.CheckResult.status is VALID_WITH_WARNINGS. | [optional] 
**ErrorMessages** | Pointer to [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | A list of messages indicating why the content was considered invalid. This field is optional and it is only relevant when the value of SupervisorServices.CheckResult.status is INVALID. | [optional] 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesCheckResult

`func NewVcenterNamespaceManagementSupervisorServicesCheckResult(status VcenterNamespaceManagementSupervisorServicesValidationStatus, ) *VcenterNamespaceManagementSupervisorServicesCheckResult`

NewVcenterNamespaceManagementSupervisorServicesCheckResult instantiates a new VcenterNamespaceManagementSupervisorServicesCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesCheckResultWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesCheckResultWithDefaults() *VcenterNamespaceManagementSupervisorServicesCheckResult`

NewVcenterNamespaceManagementSupervisorServicesCheckResultWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) GetStatus() VcenterNamespaceManagementSupervisorServicesValidationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) GetStatusOk() (*VcenterNamespaceManagementSupervisorServicesValidationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) SetStatus(v VcenterNamespaceManagementSupervisorServicesValidationStatus)`

SetStatus sets Status field to given value.


### GetContentType

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) GetContentType() VcenterNamespaceManagementSupervisorServicesVersionsContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) GetContentTypeOk() (*VcenterNamespaceManagementSupervisorServicesVersionsContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) SetContentType(v VcenterNamespaceManagementSupervisorServicesVersionsContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetVsphereAppsCheckResult

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) GetVsphereAppsCheckResult() VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult`

GetVsphereAppsCheckResult returns the VsphereAppsCheckResult field if non-nil, zero value otherwise.

### GetVsphereAppsCheckResultOk

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) GetVsphereAppsCheckResultOk() (*VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult, bool)`

GetVsphereAppsCheckResultOk returns a tuple with the VsphereAppsCheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereAppsCheckResult

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) SetVsphereAppsCheckResult(v VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult)`

SetVsphereAppsCheckResult sets VsphereAppsCheckResult field to given value.

### HasVsphereAppsCheckResult

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) HasVsphereAppsCheckResult() bool`

HasVsphereAppsCheckResult returns a boolean if a field has been set.

### GetWarningMessages

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) GetWarningMessages() []VapiStdLocalizableMessage`

GetWarningMessages returns the WarningMessages field if non-nil, zero value otherwise.

### GetWarningMessagesOk

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) GetWarningMessagesOk() (*[]VapiStdLocalizableMessage, bool)`

GetWarningMessagesOk returns a tuple with the WarningMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMessages

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) SetWarningMessages(v []VapiStdLocalizableMessage)`

SetWarningMessages sets WarningMessages field to given value.

### HasWarningMessages

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) HasWarningMessages() bool`

HasWarningMessages returns a boolean if a field has been set.

### GetErrorMessages

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) GetErrorMessages() []VapiStdLocalizableMessage`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) GetErrorMessagesOk() (*[]VapiStdLocalizableMessage, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) SetErrorMessages(v []VapiStdLocalizableMessage)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *VcenterNamespaceManagementSupervisorServicesCheckResult) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


