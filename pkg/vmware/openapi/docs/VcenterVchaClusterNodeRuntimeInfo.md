# VcenterVchaClusterNodeRuntimeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**VcenterVchaClusterNodeState**](VcenterVchaClusterNodeState.md) |  | [optional] 
**Role** | Pointer to [**VcenterVchaClusterNodeRole**](VcenterVchaClusterNodeRole.md) |  | [optional] 
**Placement** | Pointer to [**VcenterVchaPlacementInfo**](VcenterVchaPlacementInfo.md) |  | [optional] 

## Methods

### NewVcenterVchaClusterNodeRuntimeInfo

`func NewVcenterVchaClusterNodeRuntimeInfo() *VcenterVchaClusterNodeRuntimeInfo`

NewVcenterVchaClusterNodeRuntimeInfo instantiates a new VcenterVchaClusterNodeRuntimeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterNodeRuntimeInfoWithDefaults

`func NewVcenterVchaClusterNodeRuntimeInfoWithDefaults() *VcenterVchaClusterNodeRuntimeInfo`

NewVcenterVchaClusterNodeRuntimeInfoWithDefaults instantiates a new VcenterVchaClusterNodeRuntimeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *VcenterVchaClusterNodeRuntimeInfo) GetState() VcenterVchaClusterNodeState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterVchaClusterNodeRuntimeInfo) GetStateOk() (*VcenterVchaClusterNodeState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterVchaClusterNodeRuntimeInfo) SetState(v VcenterVchaClusterNodeState)`

SetState sets State field to given value.

### HasState

`func (o *VcenterVchaClusterNodeRuntimeInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRole

`func (o *VcenterVchaClusterNodeRuntimeInfo) GetRole() VcenterVchaClusterNodeRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VcenterVchaClusterNodeRuntimeInfo) GetRoleOk() (*VcenterVchaClusterNodeRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VcenterVchaClusterNodeRuntimeInfo) SetRole(v VcenterVchaClusterNodeRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *VcenterVchaClusterNodeRuntimeInfo) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetPlacement

`func (o *VcenterVchaClusterNodeRuntimeInfo) GetPlacement() VcenterVchaPlacementInfo`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVchaClusterNodeRuntimeInfo) GetPlacementOk() (*VcenterVchaPlacementInfo, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVchaClusterNodeRuntimeInfo) SetPlacement(v VcenterVchaPlacementInfo)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VcenterVchaClusterNodeRuntimeInfo) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


