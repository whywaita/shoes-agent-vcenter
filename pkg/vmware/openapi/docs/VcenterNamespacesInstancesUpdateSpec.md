# VcenterNamespacesInstancesUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description for the namespace. If unset, the description of the namespace will not be modified. | [optional] 
**ResourceSpec** | Pointer to **map[string]interface{}** | Resource quota updates on the namespace. Refer to vcenter.namespace_management.NamespaceResourceOptions.Info#updateResourceQuotaType and use vcenter.namespace_management.NamespaceResourceOptions#get for retrieving the type for the value for this field. For an example of this, see ResourceQuotaOptionsV1Update. If unset, the resource constraints on the namespace will not be modified. | [optional] 
**AccessList** | Pointer to [**[]VcenterNamespacesInstancesAccess**](VcenterNamespacesInstancesAccess.md) | Access control associated with the namespace. If unset, access controls on the namespace will not be modified. Existing pods from users will continue to run. | [optional] 
**StorageSpecs** | Pointer to [**[]VcenterNamespacesInstancesStorageSpec**](VcenterNamespacesInstancesStorageSpec.md) | Storage associated with the namespace. If unset, storage policies and their limit will not be modified. Pods which are already using persistent storage from the earlier version of storage policies will be able to access them till the datastores are attached to the worker nodes. | [optional] 
**VmServiceSpec** | Pointer to [**VcenterNamespacesInstancesVMServiceSpec**](VcenterNamespacesInstancesVMServiceSpec.md) |  | [optional] 

## Methods

### NewVcenterNamespacesInstancesUpdateSpec

`func NewVcenterNamespacesInstancesUpdateSpec() *VcenterNamespacesInstancesUpdateSpec`

NewVcenterNamespacesInstancesUpdateSpec instantiates a new VcenterNamespacesInstancesUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesInstancesUpdateSpecWithDefaults

`func NewVcenterNamespacesInstancesUpdateSpecWithDefaults() *VcenterNamespacesInstancesUpdateSpec`

NewVcenterNamespacesInstancesUpdateSpecWithDefaults instantiates a new VcenterNamespacesInstancesUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VcenterNamespacesInstancesUpdateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespacesInstancesUpdateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespacesInstancesUpdateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterNamespacesInstancesUpdateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceSpec

`func (o *VcenterNamespacesInstancesUpdateSpec) GetResourceSpec() map[string]interface{}`

GetResourceSpec returns the ResourceSpec field if non-nil, zero value otherwise.

### GetResourceSpecOk

`func (o *VcenterNamespacesInstancesUpdateSpec) GetResourceSpecOk() (*map[string]interface{}, bool)`

GetResourceSpecOk returns a tuple with the ResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSpec

`func (o *VcenterNamespacesInstancesUpdateSpec) SetResourceSpec(v map[string]interface{})`

SetResourceSpec sets ResourceSpec field to given value.

### HasResourceSpec

`func (o *VcenterNamespacesInstancesUpdateSpec) HasResourceSpec() bool`

HasResourceSpec returns a boolean if a field has been set.

### GetAccessList

`func (o *VcenterNamespacesInstancesUpdateSpec) GetAccessList() []VcenterNamespacesInstancesAccess`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *VcenterNamespacesInstancesUpdateSpec) GetAccessListOk() (*[]VcenterNamespacesInstancesAccess, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *VcenterNamespacesInstancesUpdateSpec) SetAccessList(v []VcenterNamespacesInstancesAccess)`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *VcenterNamespacesInstancesUpdateSpec) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### GetStorageSpecs

`func (o *VcenterNamespacesInstancesUpdateSpec) GetStorageSpecs() []VcenterNamespacesInstancesStorageSpec`

GetStorageSpecs returns the StorageSpecs field if non-nil, zero value otherwise.

### GetStorageSpecsOk

`func (o *VcenterNamespacesInstancesUpdateSpec) GetStorageSpecsOk() (*[]VcenterNamespacesInstancesStorageSpec, bool)`

GetStorageSpecsOk returns a tuple with the StorageSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpecs

`func (o *VcenterNamespacesInstancesUpdateSpec) SetStorageSpecs(v []VcenterNamespacesInstancesStorageSpec)`

SetStorageSpecs sets StorageSpecs field to given value.

### HasStorageSpecs

`func (o *VcenterNamespacesInstancesUpdateSpec) HasStorageSpecs() bool`

HasStorageSpecs returns a boolean if a field has been set.

### GetVmServiceSpec

`func (o *VcenterNamespacesInstancesUpdateSpec) GetVmServiceSpec() VcenterNamespacesInstancesVMServiceSpec`

GetVmServiceSpec returns the VmServiceSpec field if non-nil, zero value otherwise.

### GetVmServiceSpecOk

`func (o *VcenterNamespacesInstancesUpdateSpec) GetVmServiceSpecOk() (*VcenterNamespacesInstancesVMServiceSpec, bool)`

GetVmServiceSpecOk returns a tuple with the VmServiceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmServiceSpec

`func (o *VcenterNamespacesInstancesUpdateSpec) SetVmServiceSpec(v VcenterNamespacesInstancesVMServiceSpec)`

SetVmServiceSpec sets VmServiceSpec field to given value.

### HasVmServiceSpec

`func (o *VcenterNamespacesInstancesUpdateSpec) HasVmServiceSpec() bool`

HasVmServiceSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


