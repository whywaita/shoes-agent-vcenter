# VcenterNamespaceManagementSupervisorServicesCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomSpec** | Pointer to [**VcenterNamespaceManagementSupervisorServicesCustomCreateSpec**](VcenterNamespaceManagementSupervisorServicesCustomCreateSpec.md) |  | [optional] 
**VsphereSpec** | Pointer to [**VcenterNamespaceManagementSupervisorServicesVsphereCreateSpec**](VcenterNamespaceManagementSupervisorServicesVsphereCreateSpec.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesCreateSpec

`func NewVcenterNamespaceManagementSupervisorServicesCreateSpec() *VcenterNamespaceManagementSupervisorServicesCreateSpec`

NewVcenterNamespaceManagementSupervisorServicesCreateSpec instantiates a new VcenterNamespaceManagementSupervisorServicesCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesCreateSpecWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesCreateSpecWithDefaults() *VcenterNamespaceManagementSupervisorServicesCreateSpec`

NewVcenterNamespaceManagementSupervisorServicesCreateSpecWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomSpec

`func (o *VcenterNamespaceManagementSupervisorServicesCreateSpec) GetCustomSpec() VcenterNamespaceManagementSupervisorServicesCustomCreateSpec`

GetCustomSpec returns the CustomSpec field if non-nil, zero value otherwise.

### GetCustomSpecOk

`func (o *VcenterNamespaceManagementSupervisorServicesCreateSpec) GetCustomSpecOk() (*VcenterNamespaceManagementSupervisorServicesCustomCreateSpec, bool)`

GetCustomSpecOk returns a tuple with the CustomSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSpec

`func (o *VcenterNamespaceManagementSupervisorServicesCreateSpec) SetCustomSpec(v VcenterNamespaceManagementSupervisorServicesCustomCreateSpec)`

SetCustomSpec sets CustomSpec field to given value.

### HasCustomSpec

`func (o *VcenterNamespaceManagementSupervisorServicesCreateSpec) HasCustomSpec() bool`

HasCustomSpec returns a boolean if a field has been set.

### GetVsphereSpec

`func (o *VcenterNamespaceManagementSupervisorServicesCreateSpec) GetVsphereSpec() VcenterNamespaceManagementSupervisorServicesVsphereCreateSpec`

GetVsphereSpec returns the VsphereSpec field if non-nil, zero value otherwise.

### GetVsphereSpecOk

`func (o *VcenterNamespaceManagementSupervisorServicesCreateSpec) GetVsphereSpecOk() (*VcenterNamespaceManagementSupervisorServicesVsphereCreateSpec, bool)`

GetVsphereSpecOk returns a tuple with the VsphereSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereSpec

`func (o *VcenterNamespaceManagementSupervisorServicesCreateSpec) SetVsphereSpec(v VcenterNamespaceManagementSupervisorServicesVsphereCreateSpec)`

SetVsphereSpec sets VsphereSpec field to given value.

### HasVsphereSpec

`func (o *VcenterNamespaceManagementSupervisorServicesCreateSpec) HasVsphereSpec() bool`

HasVsphereSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


