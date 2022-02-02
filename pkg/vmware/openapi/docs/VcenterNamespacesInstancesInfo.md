# VcenterNamespacesInstancesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** | Identifier for the cluster hosting the namespace. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | 
**ConfigStatus** | [**VcenterNamespacesInstancesConfigStatus**](VcenterNamespacesInstancesConfigStatus.md) |  | 
**Messages** | [**[]VcenterNamespacesInstancesMessage**](VcenterNamespacesInstancesMessage.md) | Current set of messages associated with the object. | 
**Stats** | [**VcenterNamespacesInstancesStats**](VcenterNamespacesInstancesStats.md) |  | 
**Description** | **string** | Description of the namespace. | 
**ResourceSpec** | Pointer to **map[string]interface{}** | Quotas on the namespace resources. Refer to vcenter.namespace_management.NamespaceResourceOptions#get for the type of the value for this field. If unset, no resource constraints are associated with the namespace. | [optional] 
**AccessList** | [**[]VcenterNamespacesInstancesAccess**](VcenterNamespacesInstancesAccess.md) | Access controls associated with the namespace. | 
**StorageSpecs** | [**[]VcenterNamespacesInstancesStorageSpec**](VcenterNamespacesInstancesStorageSpec.md) | Storage associated with the namespace. | 
**Networks** | Pointer to **[]string** | vSphere Networks associated with the namespace. This field is unset if the cluster hosting this namespace uses NSXT_CONTAINER_PLUGIN as its network provider. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.namespaces.Instance. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.namespaces.Instance. | [optional] 
**VmServiceSpec** | Pointer to [**VcenterNamespacesInstancesVMServiceSpec**](VcenterNamespacesInstancesVMServiceSpec.md) |  | [optional] 
**Creator** | Pointer to [**VcenterNamespacesInstancesPrincipal**](VcenterNamespacesInstancesPrincipal.md) |  | [optional] 
**SelfServiceNamespace** | Pointer to **bool** | Flag to indicate the self service namespace. If unset, the namespace is not marked as self service namespace. | [optional] 

## Methods

### NewVcenterNamespacesInstancesInfo

`func NewVcenterNamespacesInstancesInfo(cluster string, configStatus VcenterNamespacesInstancesConfigStatus, messages []VcenterNamespacesInstancesMessage, stats VcenterNamespacesInstancesStats, description string, accessList []VcenterNamespacesInstancesAccess, storageSpecs []VcenterNamespacesInstancesStorageSpec, ) *VcenterNamespacesInstancesInfo`

NewVcenterNamespacesInstancesInfo instantiates a new VcenterNamespacesInstancesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesInstancesInfoWithDefaults

`func NewVcenterNamespacesInstancesInfoWithDefaults() *VcenterNamespacesInstancesInfo`

NewVcenterNamespacesInstancesInfoWithDefaults instantiates a new VcenterNamespacesInstancesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VcenterNamespacesInstancesInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterNamespacesInstancesInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterNamespacesInstancesInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetConfigStatus

`func (o *VcenterNamespacesInstancesInfo) GetConfigStatus() VcenterNamespacesInstancesConfigStatus`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *VcenterNamespacesInstancesInfo) GetConfigStatusOk() (*VcenterNamespacesInstancesConfigStatus, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *VcenterNamespacesInstancesInfo) SetConfigStatus(v VcenterNamespacesInstancesConfigStatus)`

SetConfigStatus sets ConfigStatus field to given value.


### GetMessages

`func (o *VcenterNamespacesInstancesInfo) GetMessages() []VcenterNamespacesInstancesMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *VcenterNamespacesInstancesInfo) GetMessagesOk() (*[]VcenterNamespacesInstancesMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *VcenterNamespacesInstancesInfo) SetMessages(v []VcenterNamespacesInstancesMessage)`

SetMessages sets Messages field to given value.


### GetStats

`func (o *VcenterNamespacesInstancesInfo) GetStats() VcenterNamespacesInstancesStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *VcenterNamespacesInstancesInfo) GetStatsOk() (*VcenterNamespacesInstancesStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *VcenterNamespacesInstancesInfo) SetStats(v VcenterNamespacesInstancesStats)`

SetStats sets Stats field to given value.


### GetDescription

`func (o *VcenterNamespacesInstancesInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespacesInstancesInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespacesInstancesInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetResourceSpec

`func (o *VcenterNamespacesInstancesInfo) GetResourceSpec() map[string]interface{}`

GetResourceSpec returns the ResourceSpec field if non-nil, zero value otherwise.

### GetResourceSpecOk

`func (o *VcenterNamespacesInstancesInfo) GetResourceSpecOk() (*map[string]interface{}, bool)`

GetResourceSpecOk returns a tuple with the ResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSpec

`func (o *VcenterNamespacesInstancesInfo) SetResourceSpec(v map[string]interface{})`

SetResourceSpec sets ResourceSpec field to given value.

### HasResourceSpec

`func (o *VcenterNamespacesInstancesInfo) HasResourceSpec() bool`

HasResourceSpec returns a boolean if a field has been set.

### GetAccessList

`func (o *VcenterNamespacesInstancesInfo) GetAccessList() []VcenterNamespacesInstancesAccess`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *VcenterNamespacesInstancesInfo) GetAccessListOk() (*[]VcenterNamespacesInstancesAccess, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *VcenterNamespacesInstancesInfo) SetAccessList(v []VcenterNamespacesInstancesAccess)`

SetAccessList sets AccessList field to given value.


### GetStorageSpecs

`func (o *VcenterNamespacesInstancesInfo) GetStorageSpecs() []VcenterNamespacesInstancesStorageSpec`

GetStorageSpecs returns the StorageSpecs field if non-nil, zero value otherwise.

### GetStorageSpecsOk

`func (o *VcenterNamespacesInstancesInfo) GetStorageSpecsOk() (*[]VcenterNamespacesInstancesStorageSpec, bool)`

GetStorageSpecsOk returns a tuple with the StorageSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpecs

`func (o *VcenterNamespacesInstancesInfo) SetStorageSpecs(v []VcenterNamespacesInstancesStorageSpec)`

SetStorageSpecs sets StorageSpecs field to given value.


### GetNetworks

`func (o *VcenterNamespacesInstancesInfo) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *VcenterNamespacesInstancesInfo) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *VcenterNamespacesInstancesInfo) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *VcenterNamespacesInstancesInfo) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetVmServiceSpec

`func (o *VcenterNamespacesInstancesInfo) GetVmServiceSpec() VcenterNamespacesInstancesVMServiceSpec`

GetVmServiceSpec returns the VmServiceSpec field if non-nil, zero value otherwise.

### GetVmServiceSpecOk

`func (o *VcenterNamespacesInstancesInfo) GetVmServiceSpecOk() (*VcenterNamespacesInstancesVMServiceSpec, bool)`

GetVmServiceSpecOk returns a tuple with the VmServiceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmServiceSpec

`func (o *VcenterNamespacesInstancesInfo) SetVmServiceSpec(v VcenterNamespacesInstancesVMServiceSpec)`

SetVmServiceSpec sets VmServiceSpec field to given value.

### HasVmServiceSpec

`func (o *VcenterNamespacesInstancesInfo) HasVmServiceSpec() bool`

HasVmServiceSpec returns a boolean if a field has been set.

### GetCreator

`func (o *VcenterNamespacesInstancesInfo) GetCreator() VcenterNamespacesInstancesPrincipal`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *VcenterNamespacesInstancesInfo) GetCreatorOk() (*VcenterNamespacesInstancesPrincipal, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *VcenterNamespacesInstancesInfo) SetCreator(v VcenterNamespacesInstancesPrincipal)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *VcenterNamespacesInstancesInfo) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetSelfServiceNamespace

`func (o *VcenterNamespacesInstancesInfo) GetSelfServiceNamespace() bool`

GetSelfServiceNamespace returns the SelfServiceNamespace field if non-nil, zero value otherwise.

### GetSelfServiceNamespaceOk

`func (o *VcenterNamespacesInstancesInfo) GetSelfServiceNamespaceOk() (*bool, bool)`

GetSelfServiceNamespaceOk returns a tuple with the SelfServiceNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceNamespace

`func (o *VcenterNamespacesInstancesInfo) SetSelfServiceNamespace(v bool)`

SetSelfServiceNamespace sets SelfServiceNamespace field to given value.

### HasSelfServiceNamespace

`func (o *VcenterNamespacesInstancesInfo) HasSelfServiceNamespace() bool`

HasSelfServiceNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


