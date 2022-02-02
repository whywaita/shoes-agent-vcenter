# VcenterNamespacesInstancesSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description for the namespace. If unset, the description of the namespace will be cleared. | [optional] 
**ResourceSpec** | Pointer to **map[string]interface{}** | Resource quota for the namespace. This will replace the existing resource constraints on the namespace in entirety. Refer to vcenter.namespace_management.NamespaceResourceOptions.Info#createResourceQuotaType and use vcenter.namespace_management.NamespaceResourceOptions#get for retrieving the type for the value for this field. For an example of this, see ResourceQuotaOptionsV1. If unset, the resource constraints on the namespace will be cleared. | [optional] 
**AccessList** | Pointer to [**[]VcenterNamespacesInstancesAccess**](VcenterNamespacesInstancesAccess.md) | Access control associated with the namespace. If unset, the existing access controls on the namespace will be removed and users will not be able to access this namespace to create new pods. Existing pods from users will continue to run. | [optional] 
**StorageSpecs** | Pointer to [**[]VcenterNamespacesInstancesStorageSpec**](VcenterNamespacesInstancesStorageSpec.md) | Storage associated with the namespace. If unset, the existing storage policies will be disassociated with the namespace and existing limits will be cleared. Pods which are already using persistent storage from the earlier version of storage policies will be able to access them till the datastores are attached to the worker nodes. | [optional] 
**VmServiceSpec** | Pointer to [**VcenterNamespacesInstancesVMServiceSpec**](VcenterNamespacesInstancesVMServiceSpec.md) |  | [optional] 

## Methods

### NewVcenterNamespacesInstancesSetSpec

`func NewVcenterNamespacesInstancesSetSpec() *VcenterNamespacesInstancesSetSpec`

NewVcenterNamespacesInstancesSetSpec instantiates a new VcenterNamespacesInstancesSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesInstancesSetSpecWithDefaults

`func NewVcenterNamespacesInstancesSetSpecWithDefaults() *VcenterNamespacesInstancesSetSpec`

NewVcenterNamespacesInstancesSetSpecWithDefaults instantiates a new VcenterNamespacesInstancesSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VcenterNamespacesInstancesSetSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespacesInstancesSetSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespacesInstancesSetSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterNamespacesInstancesSetSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceSpec

`func (o *VcenterNamespacesInstancesSetSpec) GetResourceSpec() map[string]interface{}`

GetResourceSpec returns the ResourceSpec field if non-nil, zero value otherwise.

### GetResourceSpecOk

`func (o *VcenterNamespacesInstancesSetSpec) GetResourceSpecOk() (*map[string]interface{}, bool)`

GetResourceSpecOk returns a tuple with the ResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSpec

`func (o *VcenterNamespacesInstancesSetSpec) SetResourceSpec(v map[string]interface{})`

SetResourceSpec sets ResourceSpec field to given value.

### HasResourceSpec

`func (o *VcenterNamespacesInstancesSetSpec) HasResourceSpec() bool`

HasResourceSpec returns a boolean if a field has been set.

### GetAccessList

`func (o *VcenterNamespacesInstancesSetSpec) GetAccessList() []VcenterNamespacesInstancesAccess`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *VcenterNamespacesInstancesSetSpec) GetAccessListOk() (*[]VcenterNamespacesInstancesAccess, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *VcenterNamespacesInstancesSetSpec) SetAccessList(v []VcenterNamespacesInstancesAccess)`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *VcenterNamespacesInstancesSetSpec) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### GetStorageSpecs

`func (o *VcenterNamespacesInstancesSetSpec) GetStorageSpecs() []VcenterNamespacesInstancesStorageSpec`

GetStorageSpecs returns the StorageSpecs field if non-nil, zero value otherwise.

### GetStorageSpecsOk

`func (o *VcenterNamespacesInstancesSetSpec) GetStorageSpecsOk() (*[]VcenterNamespacesInstancesStorageSpec, bool)`

GetStorageSpecsOk returns a tuple with the StorageSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpecs

`func (o *VcenterNamespacesInstancesSetSpec) SetStorageSpecs(v []VcenterNamespacesInstancesStorageSpec)`

SetStorageSpecs sets StorageSpecs field to given value.

### HasStorageSpecs

`func (o *VcenterNamespacesInstancesSetSpec) HasStorageSpecs() bool`

HasStorageSpecs returns a boolean if a field has been set.

### GetVmServiceSpec

`func (o *VcenterNamespacesInstancesSetSpec) GetVmServiceSpec() VcenterNamespacesInstancesVMServiceSpec`

GetVmServiceSpec returns the VmServiceSpec field if non-nil, zero value otherwise.

### GetVmServiceSpecOk

`func (o *VcenterNamespacesInstancesSetSpec) GetVmServiceSpecOk() (*VcenterNamespacesInstancesVMServiceSpec, bool)`

GetVmServiceSpecOk returns a tuple with the VmServiceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmServiceSpec

`func (o *VcenterNamespacesInstancesSetSpec) SetVmServiceSpec(v VcenterNamespacesInstancesVMServiceSpec)`

SetVmServiceSpec sets VmServiceSpec field to given value.

### HasVmServiceSpec

`func (o *VcenterNamespacesInstancesSetSpec) HasVmServiceSpec() bool`

HasVmServiceSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


