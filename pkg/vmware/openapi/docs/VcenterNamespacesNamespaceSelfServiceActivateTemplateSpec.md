# VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | **string** | When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespaces.NamespaceTemplate. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespaces.NamespaceTemplate. | 
**Permissions** | [**[]VcenterNamespacesNamespaceTemplatesSubject**](VcenterNamespacesNamespaceTemplatesSubject.md) | Namespace Self Service permission to subjects. | 
**ResourceSpec** | **map[string]interface{}** | Resource quotas that the template defines. Resource quota on the namespace. Refer to vcenter.namespace_management.NamespaceResourceOptions.Info#createResourceQuotaType and use vcenter.namespace_management.NamespaceResourceOptions#get for retrieving the type for the value for this field. For an example of this, see ResourceQuotaOptionsV1. | 
**StorageSpecs** | [**[]VcenterNamespacesInstancesStorageSpec**](VcenterNamespacesInstancesStorageSpec.md) | Storage associated with the self service namespace. | 
**Networks** | Pointer to **[]string** | vSphere Namespaces network objects to be associated with the namespace. The values of this list need to reference names of pre-existing Networks.Info structures. The field must be left unset if the cluster hosting the namespace uses NSXT_CONTAINER_PLUGIN as the network provider, since the network(s) for this namespace will be managed by NSX-T Container Plugin. If field is unset when the cluster hosting the namespace uses VSPHERE_NETWORK as its network provider, the namespace will automatically be associated with the cluster&#39;s Supervisor Primary Workload Network. The field currently accepts at most only 1 vSphere Namespaces network object reference. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ClusterComputeResource. | [optional] 

## Methods

### NewVcenterNamespacesNamespaceSelfServiceActivateTemplateSpec

`func NewVcenterNamespacesNamespaceSelfServiceActivateTemplateSpec(template string, permissions []VcenterNamespacesNamespaceTemplatesSubject, resourceSpec map[string]interface{}, storageSpecs []VcenterNamespacesInstancesStorageSpec, ) *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec`

NewVcenterNamespacesNamespaceSelfServiceActivateTemplateSpec instantiates a new VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesNamespaceSelfServiceActivateTemplateSpecWithDefaults

`func NewVcenterNamespacesNamespaceSelfServiceActivateTemplateSpecWithDefaults() *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec`

NewVcenterNamespacesNamespaceSelfServiceActivateTemplateSpecWithDefaults instantiates a new VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetPermissions

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) GetPermissions() []VcenterNamespacesNamespaceTemplatesSubject`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) GetPermissionsOk() (*[]VcenterNamespacesNamespaceTemplatesSubject, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) SetPermissions(v []VcenterNamespacesNamespaceTemplatesSubject)`

SetPermissions sets Permissions field to given value.


### GetResourceSpec

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) GetResourceSpec() map[string]interface{}`

GetResourceSpec returns the ResourceSpec field if non-nil, zero value otherwise.

### GetResourceSpecOk

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) GetResourceSpecOk() (*map[string]interface{}, bool)`

GetResourceSpecOk returns a tuple with the ResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSpec

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) SetResourceSpec(v map[string]interface{})`

SetResourceSpec sets ResourceSpec field to given value.


### GetStorageSpecs

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) GetStorageSpecs() []VcenterNamespacesInstancesStorageSpec`

GetStorageSpecs returns the StorageSpecs field if non-nil, zero value otherwise.

### GetStorageSpecsOk

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) GetStorageSpecsOk() (*[]VcenterNamespacesInstancesStorageSpec, bool)`

GetStorageSpecsOk returns a tuple with the StorageSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpecs

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) SetStorageSpecs(v []VcenterNamespacesInstancesStorageSpec)`

SetStorageSpecs sets StorageSpecs field to given value.


### GetNetworks

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *VcenterNamespacesNamespaceSelfServiceActivateTemplateSpec) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


