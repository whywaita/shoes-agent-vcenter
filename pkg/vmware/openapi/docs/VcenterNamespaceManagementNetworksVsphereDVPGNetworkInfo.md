# VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Portgroup** | **string** | Identifier of the vSphere Distributed Portgroup backing the vSphere network object. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network. | 
**AddressRanges** | [**[]VcenterNamespaceManagementIPRange**](VcenterNamespaceManagementIPRange.md) | Usable IP pools on this network. | 
**Gateway** | **string** | Gateway for the network. | 
**SubnetMask** | **string** | Subnet mask of the network. | 

## Methods

### NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo

`func NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo(portgroup string, addressRanges []VcenterNamespaceManagementIPRange, gateway string, subnetMask string, ) *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo`

NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo instantiates a new VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkInfoWithDefaults

`func NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkInfoWithDefaults() *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo`

NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkInfoWithDefaults instantiates a new VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortgroup

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo) GetPortgroup() string`

GetPortgroup returns the Portgroup field if non-nil, zero value otherwise.

### GetPortgroupOk

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo) GetPortgroupOk() (*string, bool)`

GetPortgroupOk returns a tuple with the Portgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortgroup

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo) SetPortgroup(v string)`

SetPortgroup sets Portgroup field to given value.


### GetAddressRanges

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo) GetAddressRanges() []VcenterNamespaceManagementIPRange`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo) GetAddressRangesOk() (*[]VcenterNamespaceManagementIPRange, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo) SetAddressRanges(v []VcenterNamespaceManagementIPRange)`

SetAddressRanges sets AddressRanges field to given value.


### GetGateway

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetSubnetMask

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkInfo) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


