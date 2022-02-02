# VcenterNamespaceManagementSupervisorServicesCustomCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupervisorService** | **string** | The identifier of the Supervisor Service. This has DNS_LABEL restrictions as specified in . This must be an alphanumeric (a-z and 0-9) string and with maximum length of 63 characters and with the &#39;-&#39; character allowed anywhere except the first or last character. This identifier must be unique across all Namespaces in this vCenter server. Additionally, the ID &#39;namespaces&#39; is reserved and must not be used. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.SupervisorService. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.SupervisorService. | 
**DisplayName** | **string** | A human readable name of the Supervisor Service. | 
**Description** | Pointer to **string** | A human readable description of the Supervisor Service. If unset, the Supervisor Service description will be empty. | [optional] 
**VersionSpec** | [**VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec**](VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec.md) |  | 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesCustomCreateSpec

`func NewVcenterNamespaceManagementSupervisorServicesCustomCreateSpec(supervisorService string, displayName string, versionSpec VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec, ) *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec`

NewVcenterNamespaceManagementSupervisorServicesCustomCreateSpec instantiates a new VcenterNamespaceManagementSupervisorServicesCustomCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesCustomCreateSpecWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesCustomCreateSpecWithDefaults() *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec`

NewVcenterNamespaceManagementSupervisorServicesCustomCreateSpecWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesCustomCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupervisorService

`func (o *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec) GetSupervisorService() string`

GetSupervisorService returns the SupervisorService field if non-nil, zero value otherwise.

### GetSupervisorServiceOk

`func (o *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec) GetSupervisorServiceOk() (*string, bool)`

GetSupervisorServiceOk returns a tuple with the SupervisorService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorService

`func (o *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec) SetSupervisorService(v string)`

SetSupervisorService sets SupervisorService field to given value.


### GetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersionSpec

`func (o *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec) GetVersionSpec() VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec`

GetVersionSpec returns the VersionSpec field if non-nil, zero value otherwise.

### GetVersionSpecOk

`func (o *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec) GetVersionSpecOk() (*VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec, bool)`

GetVersionSpecOk returns a tuple with the VersionSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSpec

`func (o *VcenterNamespaceManagementSupervisorServicesCustomCreateSpec) SetVersionSpec(v VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec)`

SetVersionSpec sets VersionSpec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


