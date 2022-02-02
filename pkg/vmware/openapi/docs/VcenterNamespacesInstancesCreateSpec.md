# VcenterNamespacesInstancesCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** | Identifier of the namespace. This has DNS_LABEL restrictions as specified in . This must be an alphanumeric (a-z and 0-9) string and with maximum length of 63 characters and with the &#39;-&#39; character allowed anywhere except the first or last character. This name is unique across all Namespaces in this vCenter server. In this version, this maps to the name of a Kubernetes namespace. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespaces.Instance. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespaces.Instance. | 
**Cluster** | **string** | Identifier of the cluster on which the namespace is being created. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | 
**Description** | Pointer to **string** | Description for the namespace. If unset, no description is added to the namespace. | [optional] 
**ResourceSpec** | Pointer to **map[string]interface{}** | Resource quota on the namespace. Refer to vcenter.namespace_management.NamespaceResourceOptions.Info#createResourceQuotaType and use vcenter.namespace_management.NamespaceResourceOptions#get for retrieving the type for the value for this field. For an example of this, see ResourceQuotaOptionsV1. If unset, no resource limits will be set on the namespace. | [optional] 
**AccessList** | Pointer to [**[]VcenterNamespacesInstancesAccess**](VcenterNamespacesInstancesAccess.md) | Access controls associated with the namespace. If unset, only users with Administrator role can access the namespace. | [optional] 
**StorageSpecs** | Pointer to [**[]VcenterNamespacesInstancesStorageSpec**](VcenterNamespacesInstancesStorageSpec.md) | Storage associated with the namespace. If unset, storage policies will not be associated with the namespace which will prevent users from being able to provision pods with persistent storage on the namespace. Users will be able to provision pods which use local storage. | [optional] 
**Networks** | Pointer to **[]string** | vSphere Namespaces network objects to be associated with the namespace. The values of this list need to reference names of pre-existing Networks.Info structures. The field must be left unset if the cluster hosting the namespace uses NSXT_CONTAINER_PLUGIN as the network provider, since the network(s) for this namespace will be managed by NSX-T Container Plugin. If field is unset when the cluster hosting the namespace uses VSPHERE_NETWORK as its network provider, the namespace will automatically be associated with the cluster&#39;s Supervisor Primary Workload Network. The field currently accepts at most only 1 vSphere Namespaces network object reference. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.namespaces.Instance. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.namespaces.Instance. | [optional] 
**VmServiceSpec** | Pointer to [**VcenterNamespacesInstancesVMServiceSpec**](VcenterNamespacesInstancesVMServiceSpec.md) |  | [optional] 
**Creator** | Pointer to [**VcenterNamespacesInstancesPrincipal**](VcenterNamespacesInstancesPrincipal.md) |  | [optional] 

## Methods

### NewVcenterNamespacesInstancesCreateSpec

`func NewVcenterNamespacesInstancesCreateSpec(namespace string, cluster string, ) *VcenterNamespacesInstancesCreateSpec`

NewVcenterNamespacesInstancesCreateSpec instantiates a new VcenterNamespacesInstancesCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesInstancesCreateSpecWithDefaults

`func NewVcenterNamespacesInstancesCreateSpecWithDefaults() *VcenterNamespacesInstancesCreateSpec`

NewVcenterNamespacesInstancesCreateSpecWithDefaults instantiates a new VcenterNamespacesInstancesCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *VcenterNamespacesInstancesCreateSpec) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *VcenterNamespacesInstancesCreateSpec) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *VcenterNamespacesInstancesCreateSpec) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetCluster

`func (o *VcenterNamespacesInstancesCreateSpec) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterNamespacesInstancesCreateSpec) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterNamespacesInstancesCreateSpec) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetDescription

`func (o *VcenterNamespacesInstancesCreateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespacesInstancesCreateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespacesInstancesCreateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterNamespacesInstancesCreateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceSpec

`func (o *VcenterNamespacesInstancesCreateSpec) GetResourceSpec() map[string]interface{}`

GetResourceSpec returns the ResourceSpec field if non-nil, zero value otherwise.

### GetResourceSpecOk

`func (o *VcenterNamespacesInstancesCreateSpec) GetResourceSpecOk() (*map[string]interface{}, bool)`

GetResourceSpecOk returns a tuple with the ResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSpec

`func (o *VcenterNamespacesInstancesCreateSpec) SetResourceSpec(v map[string]interface{})`

SetResourceSpec sets ResourceSpec field to given value.

### HasResourceSpec

`func (o *VcenterNamespacesInstancesCreateSpec) HasResourceSpec() bool`

HasResourceSpec returns a boolean if a field has been set.

### GetAccessList

`func (o *VcenterNamespacesInstancesCreateSpec) GetAccessList() []VcenterNamespacesInstancesAccess`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *VcenterNamespacesInstancesCreateSpec) GetAccessListOk() (*[]VcenterNamespacesInstancesAccess, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *VcenterNamespacesInstancesCreateSpec) SetAccessList(v []VcenterNamespacesInstancesAccess)`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *VcenterNamespacesInstancesCreateSpec) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### GetStorageSpecs

`func (o *VcenterNamespacesInstancesCreateSpec) GetStorageSpecs() []VcenterNamespacesInstancesStorageSpec`

GetStorageSpecs returns the StorageSpecs field if non-nil, zero value otherwise.

### GetStorageSpecsOk

`func (o *VcenterNamespacesInstancesCreateSpec) GetStorageSpecsOk() (*[]VcenterNamespacesInstancesStorageSpec, bool)`

GetStorageSpecsOk returns a tuple with the StorageSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpecs

`func (o *VcenterNamespacesInstancesCreateSpec) SetStorageSpecs(v []VcenterNamespacesInstancesStorageSpec)`

SetStorageSpecs sets StorageSpecs field to given value.

### HasStorageSpecs

`func (o *VcenterNamespacesInstancesCreateSpec) HasStorageSpecs() bool`

HasStorageSpecs returns a boolean if a field has been set.

### GetNetworks

`func (o *VcenterNamespacesInstancesCreateSpec) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *VcenterNamespacesInstancesCreateSpec) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *VcenterNamespacesInstancesCreateSpec) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *VcenterNamespacesInstancesCreateSpec) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetVmServiceSpec

`func (o *VcenterNamespacesInstancesCreateSpec) GetVmServiceSpec() VcenterNamespacesInstancesVMServiceSpec`

GetVmServiceSpec returns the VmServiceSpec field if non-nil, zero value otherwise.

### GetVmServiceSpecOk

`func (o *VcenterNamespacesInstancesCreateSpec) GetVmServiceSpecOk() (*VcenterNamespacesInstancesVMServiceSpec, bool)`

GetVmServiceSpecOk returns a tuple with the VmServiceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmServiceSpec

`func (o *VcenterNamespacesInstancesCreateSpec) SetVmServiceSpec(v VcenterNamespacesInstancesVMServiceSpec)`

SetVmServiceSpec sets VmServiceSpec field to given value.

### HasVmServiceSpec

`func (o *VcenterNamespacesInstancesCreateSpec) HasVmServiceSpec() bool`

HasVmServiceSpec returns a boolean if a field has been set.

### GetCreator

`func (o *VcenterNamespacesInstancesCreateSpec) GetCreator() VcenterNamespacesInstancesPrincipal`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *VcenterNamespacesInstancesCreateSpec) GetCreatorOk() (*VcenterNamespacesInstancesPrincipal, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *VcenterNamespacesInstancesCreateSpec) SetCreator(v VcenterNamespacesInstancesPrincipal)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *VcenterNamespacesInstancesCreateSpec) HasCreator() bool`

HasCreator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


