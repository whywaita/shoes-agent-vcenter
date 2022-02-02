# VcenterNamespaceManagementSupervisorServicesVersionsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The identifier of the Supervisor Service version. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version. | 
**DisplayName** | **string** | A human readable name of the Supervisor Service version. | 
**State** | [**VcenterNamespaceManagementSupervisorServicesVersionsState**](VcenterNamespaceManagementSupervisorServicesVersionsState.md) |  | 
**Description** | Pointer to **string** | A human readable description of the Supervisor Service version. If unset, no description is available for the service version. | [optional] 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesVersionsSummary

`func NewVcenterNamespaceManagementSupervisorServicesVersionsSummary(version string, displayName string, state VcenterNamespaceManagementSupervisorServicesVersionsState, ) *VcenterNamespaceManagementSupervisorServicesVersionsSummary`

NewVcenterNamespaceManagementSupervisorServicesVersionsSummary instantiates a new VcenterNamespaceManagementSupervisorServicesVersionsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesVersionsSummaryWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesVersionsSummaryWithDefaults() *VcenterNamespaceManagementSupervisorServicesVersionsSummary`

NewVcenterNamespaceManagementSupervisorServicesVersionsSummaryWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesVersionsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsSummary) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsSummary) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsSummary) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetState

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsSummary) GetState() VcenterNamespaceManagementSupervisorServicesVersionsState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsSummary) GetStateOk() (*VcenterNamespaceManagementSupervisorServicesVersionsState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsSummary) SetState(v VcenterNamespaceManagementSupervisorServicesVersionsState)`

SetState sets State field to given value.


### GetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


