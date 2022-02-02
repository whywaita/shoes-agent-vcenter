# VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupervisorService** | **string** | The identifier of the Supervisor Service. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.SupervisorService. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.SupervisorService. | 
**DesiredVersion** | **string** | The desired version of this Supervisor Service. | 
**ConfigStatus** | [**VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesConfigStatus**](VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesConfigStatus.md) |  | 
**CurrentVersion** | Pointer to **string** | The current version for the Supervisor Service. If unset, there is no version installed for the Supervisor Service. | [optional] 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary

`func NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary(supervisorService string, desiredVersion string, configStatus VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesConfigStatus, ) *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary`

NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary instantiates a new VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummaryWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummaryWithDefaults() *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary`

NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummaryWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupervisorService

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary) GetSupervisorService() string`

GetSupervisorService returns the SupervisorService field if non-nil, zero value otherwise.

### GetSupervisorServiceOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary) GetSupervisorServiceOk() (*string, bool)`

GetSupervisorServiceOk returns a tuple with the SupervisorService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorService

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary) SetSupervisorService(v string)`

SetSupervisorService sets SupervisorService field to given value.


### GetDesiredVersion

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary) GetDesiredVersion() string`

GetDesiredVersion returns the DesiredVersion field if non-nil, zero value otherwise.

### GetDesiredVersionOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary) GetDesiredVersionOk() (*string, bool)`

GetDesiredVersionOk returns a tuple with the DesiredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredVersion

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary) SetDesiredVersion(v string)`

SetDesiredVersion sets DesiredVersion field to given value.


### GetConfigStatus

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary) GetConfigStatus() VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesConfigStatus`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary) GetConfigStatusOk() (*VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesConfigStatus, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary) SetConfigStatus(v VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesConfigStatus)`

SetConfigStatus sets ConfigStatus field to given value.


### GetCurrentVersion

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesSummary) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


