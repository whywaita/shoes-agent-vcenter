# VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredVersion** | **string** | The desired version of this Supervisor Service. | 
**ServiceNamespace** | Pointer to **string** | Identifier of the namespace to allocate the Supervisor Service&#39;s operators. If unset, there is an error when creating the service namespace or the namespace has not been created yet. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespaces.Instance. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespaces.Instance. | [optional] 
**ConfigStatus** | [**VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesConfigStatus**](VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesConfigStatus.md) |  | 
**Messages** | [**[]VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesMessage**](VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesMessage.md) | Current set of messages associated with the Supervisor Service on the vSphere Supervisor cluster. | 
**CurrentVersion** | Pointer to **string** | The current version for the Supervisor Service. If unset, there is no version installed for the Supervisor Service. | [optional] 
**DisplayName** | **string** | A human readable name of the Supervisor Service. | 
**Description** | Pointer to **string** | A human readable description of the Supervisor Service. If unset, the description for the service version is empty. | [optional] 
**Prefix** | Pointer to **string** | The prefix that will be added to the names of the Supervisor Service&#39;s kubernetes resources. If unset, the prefix is not assigned yet. | [optional] 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo

`func NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo(desiredVersion string, configStatus VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesConfigStatus, messages []VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesMessage, displayName string, ) *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo`

NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo instantiates a new VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfoWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfoWithDefaults() *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo`

NewVcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfoWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredVersion

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetDesiredVersion() string`

GetDesiredVersion returns the DesiredVersion field if non-nil, zero value otherwise.

### GetDesiredVersionOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetDesiredVersionOk() (*string, bool)`

GetDesiredVersionOk returns a tuple with the DesiredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredVersion

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) SetDesiredVersion(v string)`

SetDesiredVersion sets DesiredVersion field to given value.


### GetServiceNamespace

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetServiceNamespace() string`

GetServiceNamespace returns the ServiceNamespace field if non-nil, zero value otherwise.

### GetServiceNamespaceOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetServiceNamespaceOk() (*string, bool)`

GetServiceNamespaceOk returns a tuple with the ServiceNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNamespace

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) SetServiceNamespace(v string)`

SetServiceNamespace sets ServiceNamespace field to given value.

### HasServiceNamespace

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) HasServiceNamespace() bool`

HasServiceNamespace returns a boolean if a field has been set.

### GetConfigStatus

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetConfigStatus() VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesConfigStatus`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetConfigStatusOk() (*VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesConfigStatus, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) SetConfigStatus(v VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesConfigStatus)`

SetConfigStatus sets ConfigStatus field to given value.


### GetMessages

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetMessages() []VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetMessagesOk() (*[]VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) SetMessages(v []VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesMessage)`

SetMessages sets Messages field to given value.


### GetCurrentVersion

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrefix

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *VcenterNamespaceManagementSupervisorServicesClusterSupervisorServicesInfo) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


