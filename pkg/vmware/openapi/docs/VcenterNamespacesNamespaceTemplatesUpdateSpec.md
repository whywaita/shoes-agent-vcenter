# VcenterNamespacesNamespaceTemplatesUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceSpec** | Pointer to **map[string]interface{}** | Resource quota on the namespace. Refer to vcenter.namespace_management.NamespaceResourceOptions.Info#createResourceQuotaType and use vcenter.namespace_management.NamespaceResourceOptions#get for retrieving the type for the value for this field. For an example of this, see ResourceQuotaOptionsV1. If unset, no resource limits will be set on the namespace. | [optional] 
**StorageSpecs** | Pointer to [**[]VcenterNamespacesInstancesStorageSpec**](VcenterNamespacesInstancesStorageSpec.md) | Storage that this template defines and will be associated with a namespace after namespace realization. | [optional] 
**Networks** | Pointer to **[]string** | vSphere Namespaces network objects to be associated with the namespace. The values of this list need to reference names of pre-existing Networks.Info structures. The field must be left unset if the cluster hosting the namespace uses NSXT_CONTAINER_PLUGIN as the network provider, since the network(s) for this namespace will be managed by NSX-T Container Plugin. If field is unset when the cluster hosting the namespace uses VSPHERE_NETWORK as its network provider, the namespace will automatically be associated with the cluster&#39;s Supervisor Primary Workload Network. The field currently accepts at most only 1 vSphere Namespaces network object reference. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ClusterComputeResource. | [optional] 
**Permissions** | Pointer to [**[]VcenterNamespacesNamespaceTemplatesSubject**](VcenterNamespacesNamespaceTemplatesSubject.md) | Permissions associated with namespace template. If unset, only users with the Administrator role can use this template; for example, this template is applied to the namespace created by self-service-users with the Administrator role. | [optional] 

## Methods

### NewVcenterNamespacesNamespaceTemplatesUpdateSpec

`func NewVcenterNamespacesNamespaceTemplatesUpdateSpec() *VcenterNamespacesNamespaceTemplatesUpdateSpec`

NewVcenterNamespacesNamespaceTemplatesUpdateSpec instantiates a new VcenterNamespacesNamespaceTemplatesUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesNamespaceTemplatesUpdateSpecWithDefaults

`func NewVcenterNamespacesNamespaceTemplatesUpdateSpecWithDefaults() *VcenterNamespacesNamespaceTemplatesUpdateSpec`

NewVcenterNamespacesNamespaceTemplatesUpdateSpecWithDefaults instantiates a new VcenterNamespacesNamespaceTemplatesUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceSpec

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetResourceSpec() map[string]interface{}`

GetResourceSpec returns the ResourceSpec field if non-nil, zero value otherwise.

### GetResourceSpecOk

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetResourceSpecOk() (*map[string]interface{}, bool)`

GetResourceSpecOk returns a tuple with the ResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSpec

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) SetResourceSpec(v map[string]interface{})`

SetResourceSpec sets ResourceSpec field to given value.

### HasResourceSpec

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) HasResourceSpec() bool`

HasResourceSpec returns a boolean if a field has been set.

### GetStorageSpecs

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetStorageSpecs() []VcenterNamespacesInstancesStorageSpec`

GetStorageSpecs returns the StorageSpecs field if non-nil, zero value otherwise.

### GetStorageSpecsOk

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetStorageSpecsOk() (*[]VcenterNamespacesInstancesStorageSpec, bool)`

GetStorageSpecsOk returns a tuple with the StorageSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpecs

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) SetStorageSpecs(v []VcenterNamespacesInstancesStorageSpec)`

SetStorageSpecs sets StorageSpecs field to given value.

### HasStorageSpecs

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) HasStorageSpecs() bool`

HasStorageSpecs returns a boolean if a field has been set.

### GetNetworks

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetPermissions

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetPermissions() []VcenterNamespacesNamespaceTemplatesSubject`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) GetPermissionsOk() (*[]VcenterNamespacesNamespaceTemplatesSubject, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) SetPermissions(v []VcenterNamespacesNamespaceTemplatesSubject)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *VcenterNamespacesNamespaceTemplatesUpdateSpec) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


