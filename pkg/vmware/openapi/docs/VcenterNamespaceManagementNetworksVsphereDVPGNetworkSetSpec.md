# VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Portgroup** | **string** | Identifier of the vSphere Distributed Portgroup backing the vSphere network object. If the network object is associated with a Namespace or is Clusters.WorkloadNetworksEnableSpec.supervisor-primary-workload-network, modification to existing portgroup will result in the operation failing with a ResourceInUse error. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network. | 
**AddressRanges** | [**[]VcenterNamespaceManagementIPRange**](VcenterNamespaceManagementIPRange.md) | Usable IP pools on this network. If the network object is associated with a Namespace or is Clusters.WorkloadNetworksEnableSpec.supervisor-primary-workload-network, then replacement of or modification to any existing range in this list will result in operation failing with a ResourceInUse error. To add new address ranges to the list, existing address ranges have to be passed in without modifications. | 
**Gateway** | **string** | Gateway for the network. If the network object is associated with a Namespace or is Clusters.WorkloadNetworksEnableSpec.supervisor-primary-workload-network, then modification to existing gateway will result in the operation failing with a ResourceInUse error. | 
**SubnetMask** | **string** | Subnet mask of the network. If the network object is associated with a Namespace or is Clusters.WorkloadNetworksEnableSpec.supervisor-primary-workload-network, then modification to existing subnet mask will result in the operation failing with a ResourceInUse error. | 

## Methods

### NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec

`func NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec(portgroup string, addressRanges []VcenterNamespaceManagementIPRange, gateway string, subnetMask string, ) *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec`

NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec instantiates a new VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpecWithDefaults

`func NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpecWithDefaults() *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec`

NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpecWithDefaults instantiates a new VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortgroup

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetPortgroup() string`

GetPortgroup returns the Portgroup field if non-nil, zero value otherwise.

### GetPortgroupOk

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetPortgroupOk() (*string, bool)`

GetPortgroupOk returns a tuple with the Portgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortgroup

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) SetPortgroup(v string)`

SetPortgroup sets Portgroup field to given value.


### GetAddressRanges

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetAddressRanges() []VcenterNamespaceManagementIPRange`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetAddressRangesOk() (*[]VcenterNamespaceManagementIPRange, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) SetAddressRanges(v []VcenterNamespaceManagementIPRange)`

SetAddressRanges sets AddressRanges field to given value.


### GetGateway

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetSubnetMask

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


