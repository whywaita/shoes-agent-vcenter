# VcenterNamespaceManagementNetworksCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | Identifier of the network. This has DNS_LABEL restrictions as specified in . This must be an alphanumeric (a-z and 0-9) string and with maximum length of 63 characters and with the &#39;-&#39; character allowed anywhere except the first or last character. This name must be unique within a cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.Network. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.Network. | 
**NetworkProvider** | [**VcenterNamespaceManagementClustersNetworkProvider**](VcenterNamespaceManagementClustersNetworkProvider.md) |  | 
**VsphereNetwork** | Pointer to [**VcenterNamespaceManagementNetworksVsphereDVPGNetworkCreateSpec**](VcenterNamespaceManagementNetworksVsphereDVPGNetworkCreateSpec.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementNetworksCreateSpec

`func NewVcenterNamespaceManagementNetworksCreateSpec(network string, networkProvider VcenterNamespaceManagementClustersNetworkProvider, ) *VcenterNamespaceManagementNetworksCreateSpec`

NewVcenterNamespaceManagementNetworksCreateSpec instantiates a new VcenterNamespaceManagementNetworksCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementNetworksCreateSpecWithDefaults

`func NewVcenterNamespaceManagementNetworksCreateSpecWithDefaults() *VcenterNamespaceManagementNetworksCreateSpec`

NewVcenterNamespaceManagementNetworksCreateSpecWithDefaults instantiates a new VcenterNamespaceManagementNetworksCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *VcenterNamespaceManagementNetworksCreateSpec) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VcenterNamespaceManagementNetworksCreateSpec) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VcenterNamespaceManagementNetworksCreateSpec) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetNetworkProvider

`func (o *VcenterNamespaceManagementNetworksCreateSpec) GetNetworkProvider() VcenterNamespaceManagementClustersNetworkProvider`

GetNetworkProvider returns the NetworkProvider field if non-nil, zero value otherwise.

### GetNetworkProviderOk

`func (o *VcenterNamespaceManagementNetworksCreateSpec) GetNetworkProviderOk() (*VcenterNamespaceManagementClustersNetworkProvider, bool)`

GetNetworkProviderOk returns a tuple with the NetworkProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProvider

`func (o *VcenterNamespaceManagementNetworksCreateSpec) SetNetworkProvider(v VcenterNamespaceManagementClustersNetworkProvider)`

SetNetworkProvider sets NetworkProvider field to given value.


### GetVsphereNetwork

`func (o *VcenterNamespaceManagementNetworksCreateSpec) GetVsphereNetwork() VcenterNamespaceManagementNetworksVsphereDVPGNetworkCreateSpec`

GetVsphereNetwork returns the VsphereNetwork field if non-nil, zero value otherwise.

### GetVsphereNetworkOk

`func (o *VcenterNamespaceManagementNetworksCreateSpec) GetVsphereNetworkOk() (*VcenterNamespaceManagementNetworksVsphereDVPGNetworkCreateSpec, bool)`

GetVsphereNetworkOk returns a tuple with the VsphereNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereNetwork

`func (o *VcenterNamespaceManagementNetworksCreateSpec) SetVsphereNetwork(v VcenterNamespaceManagementNetworksVsphereDVPGNetworkCreateSpec)`

SetVsphereNetwork sets VsphereNetwork field to given value.

### HasVsphereNetwork

`func (o *VcenterNamespaceManagementNetworksCreateSpec) HasVsphereNetwork() bool`

HasVsphereNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


