# VcenterNamespaceManagementSupervisorServicesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupervisorService** | **string** | The identifier of the Supervisor Service. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.SupervisorService. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.SupervisorService. | 
**DisplayName** | **string** | The human readable name of the Supervisor Service. | 
**State** | [**VcenterNamespaceManagementSupervisorServicesState**](VcenterNamespaceManagementSupervisorServicesState.md) |  | 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesSummary

`func NewVcenterNamespaceManagementSupervisorServicesSummary(supervisorService string, displayName string, state VcenterNamespaceManagementSupervisorServicesState, ) *VcenterNamespaceManagementSupervisorServicesSummary`

NewVcenterNamespaceManagementSupervisorServicesSummary instantiates a new VcenterNamespaceManagementSupervisorServicesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesSummaryWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesSummaryWithDefaults() *VcenterNamespaceManagementSupervisorServicesSummary`

NewVcenterNamespaceManagementSupervisorServicesSummaryWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupervisorService

`func (o *VcenterNamespaceManagementSupervisorServicesSummary) GetSupervisorService() string`

GetSupervisorService returns the SupervisorService field if non-nil, zero value otherwise.

### GetSupervisorServiceOk

`func (o *VcenterNamespaceManagementSupervisorServicesSummary) GetSupervisorServiceOk() (*string, bool)`

GetSupervisorServiceOk returns a tuple with the SupervisorService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorService

`func (o *VcenterNamespaceManagementSupervisorServicesSummary) SetSupervisorService(v string)`

SetSupervisorService sets SupervisorService field to given value.


### GetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VcenterNamespaceManagementSupervisorServicesSummary) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesSummary) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetState

`func (o *VcenterNamespaceManagementSupervisorServicesSummary) GetState() VcenterNamespaceManagementSupervisorServicesState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterNamespaceManagementSupervisorServicesSummary) GetStateOk() (*VcenterNamespaceManagementSupervisorServicesState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterNamespaceManagementSupervisorServicesSummary) SetState(v VcenterNamespaceManagementSupervisorServicesState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


