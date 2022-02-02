# VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupervisorPrimaryWorkloadNetwork** | [**VcenterNamespaceManagementNetworksCreateSpec**](VcenterNamespaceManagementNetworksCreateSpec.md) |  | 
**NetworkList** | Pointer to [**[]VcenterNamespaceManagementNetworksCreateSpec**](VcenterNamespaceManagementNetworksCreateSpec.md) | CreateSpecs structurees for additional list of vSphere Namespaces networks to be associated with this cluster. | [optional] 

## Methods

### NewVcenterNamespaceManagementClustersWorkloadNetworksEnableSpec

`func NewVcenterNamespaceManagementClustersWorkloadNetworksEnableSpec(supervisorPrimaryWorkloadNetwork VcenterNamespaceManagementNetworksCreateSpec, ) *VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec`

NewVcenterNamespaceManagementClustersWorkloadNetworksEnableSpec instantiates a new VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersWorkloadNetworksEnableSpecWithDefaults

`func NewVcenterNamespaceManagementClustersWorkloadNetworksEnableSpecWithDefaults() *VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec`

NewVcenterNamespaceManagementClustersWorkloadNetworksEnableSpecWithDefaults instantiates a new VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupervisorPrimaryWorkloadNetwork

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec) GetSupervisorPrimaryWorkloadNetwork() VcenterNamespaceManagementNetworksCreateSpec`

GetSupervisorPrimaryWorkloadNetwork returns the SupervisorPrimaryWorkloadNetwork field if non-nil, zero value otherwise.

### GetSupervisorPrimaryWorkloadNetworkOk

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec) GetSupervisorPrimaryWorkloadNetworkOk() (*VcenterNamespaceManagementNetworksCreateSpec, bool)`

GetSupervisorPrimaryWorkloadNetworkOk returns a tuple with the SupervisorPrimaryWorkloadNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorPrimaryWorkloadNetwork

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec) SetSupervisorPrimaryWorkloadNetwork(v VcenterNamespaceManagementNetworksCreateSpec)`

SetSupervisorPrimaryWorkloadNetwork sets SupervisorPrimaryWorkloadNetwork field to given value.


### GetNetworkList

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec) GetNetworkList() []VcenterNamespaceManagementNetworksCreateSpec`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec) GetNetworkListOk() (*[]VcenterNamespaceManagementNetworksCreateSpec, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec) SetNetworkList(v []VcenterNamespaceManagementNetworksCreateSpec)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


