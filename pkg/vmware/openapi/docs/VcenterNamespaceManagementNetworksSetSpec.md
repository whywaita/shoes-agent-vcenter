# VcenterNamespaceManagementNetworksSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkProvider** | [**VcenterNamespaceManagementClustersNetworkProvider**](VcenterNamespaceManagementClustersNetworkProvider.md) |  | 
**VsphereNetwork** | Pointer to [**VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec**](VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementNetworksSetSpec

`func NewVcenterNamespaceManagementNetworksSetSpec(networkProvider VcenterNamespaceManagementClustersNetworkProvider, ) *VcenterNamespaceManagementNetworksSetSpec`

NewVcenterNamespaceManagementNetworksSetSpec instantiates a new VcenterNamespaceManagementNetworksSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementNetworksSetSpecWithDefaults

`func NewVcenterNamespaceManagementNetworksSetSpecWithDefaults() *VcenterNamespaceManagementNetworksSetSpec`

NewVcenterNamespaceManagementNetworksSetSpecWithDefaults instantiates a new VcenterNamespaceManagementNetworksSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkProvider

`func (o *VcenterNamespaceManagementNetworksSetSpec) GetNetworkProvider() VcenterNamespaceManagementClustersNetworkProvider`

GetNetworkProvider returns the NetworkProvider field if non-nil, zero value otherwise.

### GetNetworkProviderOk

`func (o *VcenterNamespaceManagementNetworksSetSpec) GetNetworkProviderOk() (*VcenterNamespaceManagementClustersNetworkProvider, bool)`

GetNetworkProviderOk returns a tuple with the NetworkProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProvider

`func (o *VcenterNamespaceManagementNetworksSetSpec) SetNetworkProvider(v VcenterNamespaceManagementClustersNetworkProvider)`

SetNetworkProvider sets NetworkProvider field to given value.


### GetVsphereNetwork

`func (o *VcenterNamespaceManagementNetworksSetSpec) GetVsphereNetwork() VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec`

GetVsphereNetwork returns the VsphereNetwork field if non-nil, zero value otherwise.

### GetVsphereNetworkOk

`func (o *VcenterNamespaceManagementNetworksSetSpec) GetVsphereNetworkOk() (*VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec, bool)`

GetVsphereNetworkOk returns a tuple with the VsphereNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereNetwork

`func (o *VcenterNamespaceManagementNetworksSetSpec) SetVsphereNetwork(v VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec)`

SetVsphereNetwork sets VsphereNetwork field to given value.

### HasVsphereNetwork

`func (o *VcenterNamespaceManagementNetworksSetSpec) HasVsphereNetwork() bool`

HasVsphereNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


