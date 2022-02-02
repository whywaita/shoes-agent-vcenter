# VcenterNamespaceManagementNetworksInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | Identifier of the network. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.Network. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.Network. | 
**NetworkProvider** | [**VcenterNamespaceManagementClustersNetworkProvider**](VcenterNamespaceManagementClustersNetworkProvider.md) |  | 
**VsphereNetwork** | Pointer to [**VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo**](VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementNetworksInfo

`func NewVcenterNamespaceManagementNetworksInfo(network string, networkProvider VcenterNamespaceManagementClustersNetworkProvider, ) *VcenterNamespaceManagementNetworksInfo`

NewVcenterNamespaceManagementNetworksInfo instantiates a new VcenterNamespaceManagementNetworksInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementNetworksInfoWithDefaults

`func NewVcenterNamespaceManagementNetworksInfoWithDefaults() *VcenterNamespaceManagementNetworksInfo`

NewVcenterNamespaceManagementNetworksInfoWithDefaults instantiates a new VcenterNamespaceManagementNetworksInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *VcenterNamespaceManagementNetworksInfo) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VcenterNamespaceManagementNetworksInfo) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VcenterNamespaceManagementNetworksInfo) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetNetworkProvider

`func (o *VcenterNamespaceManagementNetworksInfo) GetNetworkProvider() VcenterNamespaceManagementClustersNetworkProvider`

GetNetworkProvider returns the NetworkProvider field if non-nil, zero value otherwise.

### GetNetworkProviderOk

`func (o *VcenterNamespaceManagementNetworksInfo) GetNetworkProviderOk() (*VcenterNamespaceManagementClustersNetworkProvider, bool)`

GetNetworkProviderOk returns a tuple with the NetworkProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProvider

`func (o *VcenterNamespaceManagementNetworksInfo) SetNetworkProvider(v VcenterNamespaceManagementClustersNetworkProvider)`

SetNetworkProvider sets NetworkProvider field to given value.


### GetVsphereNetwork

`func (o *VcenterNamespaceManagementNetworksInfo) GetVsphereNetwork() VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo`

GetVsphereNetwork returns the VsphereNetwork field if non-nil, zero value otherwise.

### GetVsphereNetworkOk

`func (o *VcenterNamespaceManagementNetworksInfo) GetVsphereNetworkOk() (*VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo, bool)`

GetVsphereNetworkOk returns a tuple with the VsphereNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereNetwork

`func (o *VcenterNamespaceManagementNetworksInfo) SetVsphereNetwork(v VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo)`

SetVsphereNetwork sets VsphereNetwork field to given value.

### HasVsphereNetwork

`func (o *VcenterNamespaceManagementNetworksInfo) HasVsphereNetwork() bool`

HasVsphereNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


