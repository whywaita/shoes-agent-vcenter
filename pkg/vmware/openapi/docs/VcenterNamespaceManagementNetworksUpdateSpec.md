# VcenterNamespaceManagementNetworksUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkProvider** | [**VcenterNamespaceManagementClustersNetworkProvider**](VcenterNamespaceManagementClustersNetworkProvider.md) |  | 
**VsphereNetwork** | Pointer to [**VcenterNamespaceManagementNetworksVsphereDVPGNetworkUpdateSpec**](VcenterNamespaceManagementNetworksVsphereDVPGNetworkUpdateSpec.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementNetworksUpdateSpec

`func NewVcenterNamespaceManagementNetworksUpdateSpec(networkProvider VcenterNamespaceManagementClustersNetworkProvider, ) *VcenterNamespaceManagementNetworksUpdateSpec`

NewVcenterNamespaceManagementNetworksUpdateSpec instantiates a new VcenterNamespaceManagementNetworksUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementNetworksUpdateSpecWithDefaults

`func NewVcenterNamespaceManagementNetworksUpdateSpecWithDefaults() *VcenterNamespaceManagementNetworksUpdateSpec`

NewVcenterNamespaceManagementNetworksUpdateSpecWithDefaults instantiates a new VcenterNamespaceManagementNetworksUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkProvider

`func (o *VcenterNamespaceManagementNetworksUpdateSpec) GetNetworkProvider() VcenterNamespaceManagementClustersNetworkProvider`

GetNetworkProvider returns the NetworkProvider field if non-nil, zero value otherwise.

### GetNetworkProviderOk

`func (o *VcenterNamespaceManagementNetworksUpdateSpec) GetNetworkProviderOk() (*VcenterNamespaceManagementClustersNetworkProvider, bool)`

GetNetworkProviderOk returns a tuple with the NetworkProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProvider

`func (o *VcenterNamespaceManagementNetworksUpdateSpec) SetNetworkProvider(v VcenterNamespaceManagementClustersNetworkProvider)`

SetNetworkProvider sets NetworkProvider field to given value.


### GetVsphereNetwork

`func (o *VcenterNamespaceManagementNetworksUpdateSpec) GetVsphereNetwork() VcenterNamespaceManagementNetworksVsphereDVPGNetworkUpdateSpec`

GetVsphereNetwork returns the VsphereNetwork field if non-nil, zero value otherwise.

### GetVsphereNetworkOk

`func (o *VcenterNamespaceManagementNetworksUpdateSpec) GetVsphereNetworkOk() (*VcenterNamespaceManagementNetworksVsphereDVPGNetworkUpdateSpec, bool)`

GetVsphereNetworkOk returns a tuple with the VsphereNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereNetwork

`func (o *VcenterNamespaceManagementNetworksUpdateSpec) SetVsphereNetwork(v VcenterNamespaceManagementNetworksVsphereDVPGNetworkUpdateSpec)`

SetVsphereNetwork sets VsphereNetwork field to given value.

### HasVsphereNetwork

`func (o *VcenterNamespaceManagementNetworksUpdateSpec) HasVsphereNetwork() bool`

HasVsphereNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


