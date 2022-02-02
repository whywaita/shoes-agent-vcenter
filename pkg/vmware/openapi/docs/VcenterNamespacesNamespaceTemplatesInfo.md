# VcenterNamespacesNamespaceTemplatesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** | Identifier for the cluster associated with namespace template. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | 
**Template** | **string** | Name of the namespace template. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespaces.NamespaceTemplate. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespaces.NamespaceTemplate. | 
**ResourceSpec** | Pointer to **map[string]interface{}** | Resource quotas that this template defines. Quotas on the namespace resources. Refer to vcenter.namespace_management.NamespaceResourceOptions#get for the type of the value for this field. If unset, no resource constraints are defined in the namespace template. | [optional] 
**StorageSpecs** | [**[]VcenterNamespacesInstancesStorageSpec**](VcenterNamespacesInstancesStorageSpec.md) | Storage that this template defines and will be associated with a namespace after namespace realization. | 
**Networks** | Pointer to **[]string** | vSphere Networks that this template captures and are associated with the namespace after namespace realization. This field is unset if the cluster hosting this namespace uses NSXT_CONTAINER_PLUGIN as its network provider. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ClusterComputeResource. | [optional] 
**Permissions** | Pointer to [**[]VcenterNamespacesNamespaceTemplatesSubject**](VcenterNamespacesNamespaceTemplatesSubject.md) | Permissions associated with namespace template. If unset, only users with the Administrator role can use this template; for example, this template is applied to the namespace created by self-service-users with the Administrator role. | [optional] 

## Methods

### NewVcenterNamespacesNamespaceTemplatesInfo

`func NewVcenterNamespacesNamespaceTemplatesInfo(cluster string, template string, storageSpecs []VcenterNamespacesInstancesStorageSpec, ) *VcenterNamespacesNamespaceTemplatesInfo`

NewVcenterNamespacesNamespaceTemplatesInfo instantiates a new VcenterNamespacesNamespaceTemplatesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesNamespaceTemplatesInfoWithDefaults

`func NewVcenterNamespacesNamespaceTemplatesInfoWithDefaults() *VcenterNamespacesNamespaceTemplatesInfo`

NewVcenterNamespacesNamespaceTemplatesInfoWithDefaults instantiates a new VcenterNamespacesNamespaceTemplatesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VcenterNamespacesNamespaceTemplatesInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterNamespacesNamespaceTemplatesInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterNamespacesNamespaceTemplatesInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetTemplate

`func (o *VcenterNamespacesNamespaceTemplatesInfo) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *VcenterNamespacesNamespaceTemplatesInfo) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *VcenterNamespacesNamespaceTemplatesInfo) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetResourceSpec

`func (o *VcenterNamespacesNamespaceTemplatesInfo) GetResourceSpec() map[string]interface{}`

GetResourceSpec returns the ResourceSpec field if non-nil, zero value otherwise.

### GetResourceSpecOk

`func (o *VcenterNamespacesNamespaceTemplatesInfo) GetResourceSpecOk() (*map[string]interface{}, bool)`

GetResourceSpecOk returns a tuple with the ResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSpec

`func (o *VcenterNamespacesNamespaceTemplatesInfo) SetResourceSpec(v map[string]interface{})`

SetResourceSpec sets ResourceSpec field to given value.

### HasResourceSpec

`func (o *VcenterNamespacesNamespaceTemplatesInfo) HasResourceSpec() bool`

HasResourceSpec returns a boolean if a field has been set.

### GetStorageSpecs

`func (o *VcenterNamespacesNamespaceTemplatesInfo) GetStorageSpecs() []VcenterNamespacesInstancesStorageSpec`

GetStorageSpecs returns the StorageSpecs field if non-nil, zero value otherwise.

### GetStorageSpecsOk

`func (o *VcenterNamespacesNamespaceTemplatesInfo) GetStorageSpecsOk() (*[]VcenterNamespacesInstancesStorageSpec, bool)`

GetStorageSpecsOk returns a tuple with the StorageSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpecs

`func (o *VcenterNamespacesNamespaceTemplatesInfo) SetStorageSpecs(v []VcenterNamespacesInstancesStorageSpec)`

SetStorageSpecs sets StorageSpecs field to given value.


### GetNetworks

`func (o *VcenterNamespacesNamespaceTemplatesInfo) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *VcenterNamespacesNamespaceTemplatesInfo) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *VcenterNamespacesNamespaceTemplatesInfo) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *VcenterNamespacesNamespaceTemplatesInfo) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetPermissions

`func (o *VcenterNamespacesNamespaceTemplatesInfo) GetPermissions() []VcenterNamespacesNamespaceTemplatesSubject`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *VcenterNamespacesNamespaceTemplatesInfo) GetPermissionsOk() (*[]VcenterNamespacesNamespaceTemplatesSubject, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *VcenterNamespacesNamespaceTemplatesInfo) SetPermissions(v []VcenterNamespacesNamespaceTemplatesSubject)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *VcenterNamespacesNamespaceTemplatesInfo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


