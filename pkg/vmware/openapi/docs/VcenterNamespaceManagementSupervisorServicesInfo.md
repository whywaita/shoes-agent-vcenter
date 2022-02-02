# VcenterNamespaceManagementSupervisorServicesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The human readable name of the Supervisor Service. | 
**State** | [**VcenterNamespaceManagementSupervisorServicesState**](VcenterNamespaceManagementSupervisorServicesState.md) |  | 
**Description** | Pointer to **string** | A human readable description of the Supervisor Service. If unset, no description is available for the Supervisor Service. | [optional] 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesInfo

`func NewVcenterNamespaceManagementSupervisorServicesInfo(displayName string, state VcenterNamespaceManagementSupervisorServicesState, ) *VcenterNamespaceManagementSupervisorServicesInfo`

NewVcenterNamespaceManagementSupervisorServicesInfo instantiates a new VcenterNamespaceManagementSupervisorServicesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesInfoWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesInfoWithDefaults() *VcenterNamespaceManagementSupervisorServicesInfo`

NewVcenterNamespaceManagementSupervisorServicesInfoWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VcenterNamespaceManagementSupervisorServicesInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetState

`func (o *VcenterNamespaceManagementSupervisorServicesInfo) GetState() VcenterNamespaceManagementSupervisorServicesState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterNamespaceManagementSupervisorServicesInfo) GetStateOk() (*VcenterNamespaceManagementSupervisorServicesState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterNamespaceManagementSupervisorServicesInfo) SetState(v VcenterNamespaceManagementSupervisorServicesState)`

SetState sets State field to given value.


### GetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespaceManagementSupervisorServicesInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterNamespaceManagementSupervisorServicesInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


