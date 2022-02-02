# VcenterNamespacesAccessInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | [**VcenterNamespacesAccessRole**](VcenterNamespacesAccessRole.md) |  | 
**Inherited** | Pointer to **bool** | Flag to indicate if the Access.Info.role is direct or inherited. The value is set to true if the Access.Info.role is inherited from group membership This field is optional because it was added in a newer version than its parent node. | [optional] 

## Methods

### NewVcenterNamespacesAccessInfo

`func NewVcenterNamespacesAccessInfo(role VcenterNamespacesAccessRole, ) *VcenterNamespacesAccessInfo`

NewVcenterNamespacesAccessInfo instantiates a new VcenterNamespacesAccessInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesAccessInfoWithDefaults

`func NewVcenterNamespacesAccessInfoWithDefaults() *VcenterNamespacesAccessInfo`

NewVcenterNamespacesAccessInfoWithDefaults instantiates a new VcenterNamespacesAccessInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *VcenterNamespacesAccessInfo) GetRole() VcenterNamespacesAccessRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VcenterNamespacesAccessInfo) GetRoleOk() (*VcenterNamespacesAccessRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VcenterNamespacesAccessInfo) SetRole(v VcenterNamespacesAccessRole)`

SetRole sets Role field to given value.


### GetInherited

`func (o *VcenterNamespacesAccessInfo) GetInherited() bool`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *VcenterNamespacesAccessInfo) GetInheritedOk() (*bool, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *VcenterNamespacesAccessInfo) SetInherited(v bool)`

SetInherited sets Inherited field to given value.

### HasInherited

`func (o *VcenterNamespacesAccessInfo) HasInherited() bool`

HasInherited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


