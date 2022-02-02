# VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupervisorService** | **string** | The identifier of the Supervisor Service. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.SupervisorService. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.SupervisorService. | 
**Version** | **string** | The identifier of the Supervisor Service version. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version. | 
**DisplayName** | **string** | A human readable name of the Supervisor Service version. | 
**Description** | Pointer to **string** | A human readable description of the Supervisor Service version. If unset, no description is available for the Supervisor Service version. | [optional] 
**Eula** | Pointer to **string** | The End User License Agreement (EULA) of the Supervisor Service version. If unset, no EULA is available for the Supervisor Service version. | [optional] 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult

`func NewVcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult(supervisorService string, version string, displayName string, ) *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult`

NewVcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult instantiates a new VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResultWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResultWithDefaults() *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult`

NewVcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResultWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupervisorService

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) GetSupervisorService() string`

GetSupervisorService returns the SupervisorService field if non-nil, zero value otherwise.

### GetSupervisorServiceOk

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) GetSupervisorServiceOk() (*string, bool)`

GetSupervisorServiceOk returns a tuple with the SupervisorService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorService

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) SetSupervisorService(v string)`

SetSupervisorService sets SupervisorService field to given value.


### GetVersion

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEula

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) GetEula() string`

GetEula returns the Eula field if non-nil, zero value otherwise.

### GetEulaOk

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) GetEulaOk() (*string, bool)`

GetEulaOk returns a tuple with the Eula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEula

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) SetEula(v string)`

SetEula sets Eula field to given value.

### HasEula

`func (o *VcenterNamespaceManagementSupervisorServicesVsphereAppsCheckResult) HasEula() bool`

HasEula returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


