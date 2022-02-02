# VcenterNamespaceManagementClustersWorkloadNetworksInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupervisorPrimaryWorkloadNetwork** | [**VcenterNamespaceManagementNetworksInfo**](VcenterNamespaceManagementNetworksInfo.md) |  | 
**NetworkList** | Pointer to [**[]VcenterNamespaceManagementNetworksInfo**](VcenterNamespaceManagementNetworksInfo.md) | List of vSphere Namespaces networks associated with this cluster. | [optional] 

## Methods

### NewVcenterNamespaceManagementClustersWorkloadNetworksInfo

`func NewVcenterNamespaceManagementClustersWorkloadNetworksInfo(supervisorPrimaryWorkloadNetwork VcenterNamespaceManagementNetworksInfo, ) *VcenterNamespaceManagementClustersWorkloadNetworksInfo`

NewVcenterNamespaceManagementClustersWorkloadNetworksInfo instantiates a new VcenterNamespaceManagementClustersWorkloadNetworksInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersWorkloadNetworksInfoWithDefaults

`func NewVcenterNamespaceManagementClustersWorkloadNetworksInfoWithDefaults() *VcenterNamespaceManagementClustersWorkloadNetworksInfo`

NewVcenterNamespaceManagementClustersWorkloadNetworksInfoWithDefaults instantiates a new VcenterNamespaceManagementClustersWorkloadNetworksInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupervisorPrimaryWorkloadNetwork

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) GetSupervisorPrimaryWorkloadNetwork() VcenterNamespaceManagementNetworksInfo`

GetSupervisorPrimaryWorkloadNetwork returns the SupervisorPrimaryWorkloadNetwork field if non-nil, zero value otherwise.

### GetSupervisorPrimaryWorkloadNetworkOk

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) GetSupervisorPrimaryWorkloadNetworkOk() (*VcenterNamespaceManagementNetworksInfo, bool)`

GetSupervisorPrimaryWorkloadNetworkOk returns a tuple with the SupervisorPrimaryWorkloadNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorPrimaryWorkloadNetwork

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) SetSupervisorPrimaryWorkloadNetwork(v VcenterNamespaceManagementNetworksInfo)`

SetSupervisorPrimaryWorkloadNetwork sets SupervisorPrimaryWorkloadNetwork field to given value.


### GetNetworkList

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) GetNetworkList() []VcenterNamespaceManagementNetworksInfo`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) GetNetworkListOk() (*[]VcenterNamespaceManagementNetworksInfo, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) SetNetworkList(v []VcenterNamespaceManagementNetworksInfo)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


