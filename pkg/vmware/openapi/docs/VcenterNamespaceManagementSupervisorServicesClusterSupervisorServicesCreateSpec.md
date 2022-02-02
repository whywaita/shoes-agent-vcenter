# VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupervisorService** | **string** | Identifier of the Supervisor Service. This Supervisor Service must be in the State#ACTIVATED state. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.SupervisorService. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.SupervisorService. | 
**Version** | **string** | Identifier of the Supervisor Service version which contains the service definition. This Supervisor Service version must be in the State#ACTIVATED state. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version. | 
**ServiceConfig** | Pointer to **map[string]string** | A generic key-value map for additional configuration parameters required during service creation. As an example, a third party operator might reference a private registry using parameters such as \&quot;registryName\&quot; for the registry name, \&quot;registryUsername\&quot; and \&quot;registryPassword\&quot; for the registry credentials. If unset, no additional configuration parameters will be applied when installing a Supervisor Service in the vSphere Supervisor cluster. | [optional] 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec

`func NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec(supervisorService string, version string, ) *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec`

NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec instantiates a new VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpecWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpecWithDefaults() *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec`

NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpecWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupervisorService

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) GetSupervisorService() string`

GetSupervisorService returns the SupervisorService field if non-nil, zero value otherwise.

### GetSupervisorServiceOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) GetSupervisorServiceOk() (*string, bool)`

GetSupervisorServiceOk returns a tuple with the SupervisorService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorService

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) SetSupervisorService(v string)`

SetSupervisorService sets SupervisorService field to given value.


### GetVersion

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetServiceConfig

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) GetServiceConfig() map[string]string`

GetServiceConfig returns the ServiceConfig field if non-nil, zero value otherwise.

### GetServiceConfigOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) GetServiceConfigOk() (*map[string]string, bool)`

GetServiceConfigOk returns a tuple with the ServiceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfig

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) SetServiceConfig(v map[string]string)`

SetServiceConfig sets ServiceConfig field to given value.

### HasServiceConfig

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesCreateSpec) HasServiceConfig() bool`

HasServiceConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


