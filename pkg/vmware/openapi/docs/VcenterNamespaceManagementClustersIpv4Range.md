# VcenterNamespaceManagementClustersIpv4Range

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartingAddress** | **string** | The IPv4 address denoting the start of the range. | 
**AddressCount** | **int64** | The number of IP addresses in the range. Addresses are derived by incrementing Clusters.Ipv4Range.starting-address. | 
**SubnetMask** | **string** | Subnet mask to be set. | 
**Gateway** | **string** | The IPv4 address of the gateway associated with the range indicated by Clusters.Ipv4Range.starting-address and Clusters.Ipv4Range.address-count. | 

## Methods

### NewVcenterNamespaceManagementClustersIpv4Range

`func NewVcenterNamespaceManagementClustersIpv4Range(startingAddress string, addressCount int64, subnetMask string, gateway string, ) *VcenterNamespaceManagementClustersIpv4Range`

NewVcenterNamespaceManagementClustersIpv4Range instantiates a new VcenterNamespaceManagementClustersIpv4Range object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersIpv4RangeWithDefaults

`func NewVcenterNamespaceManagementClustersIpv4RangeWithDefaults() *VcenterNamespaceManagementClustersIpv4Range`

NewVcenterNamespaceManagementClustersIpv4RangeWithDefaults instantiates a new VcenterNamespaceManagementClustersIpv4Range object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartingAddress

`func (o *VcenterNamespaceManagementClustersIpv4Range) GetStartingAddress() string`

GetStartingAddress returns the StartingAddress field if non-nil, zero value otherwise.

### GetStartingAddressOk

`func (o *VcenterNamespaceManagementClustersIpv4Range) GetStartingAddressOk() (*string, bool)`

GetStartingAddressOk returns a tuple with the StartingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingAddress

`func (o *VcenterNamespaceManagementClustersIpv4Range) SetStartingAddress(v string)`

SetStartingAddress sets StartingAddress field to given value.


### GetAddressCount

`func (o *VcenterNamespaceManagementClustersIpv4Range) GetAddressCount() int64`

GetAddressCount returns the AddressCount field if non-nil, zero value otherwise.

### GetAddressCountOk

`func (o *VcenterNamespaceManagementClustersIpv4Range) GetAddressCountOk() (*int64, bool)`

GetAddressCountOk returns a tuple with the AddressCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCount

`func (o *VcenterNamespaceManagementClustersIpv4Range) SetAddressCount(v int64)`

SetAddressCount sets AddressCount field to given value.


### GetSubnetMask

`func (o *VcenterNamespaceManagementClustersIpv4Range) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *VcenterNamespaceManagementClustersIpv4Range) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *VcenterNamespaceManagementClustersIpv4Range) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.


### GetGateway

`func (o *VcenterNamespaceManagementClustersIpv4Range) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *VcenterNamespaceManagementClustersIpv4Range) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *VcenterNamespaceManagementClustersIpv4Range) SetGateway(v string)`

SetGateway sets Gateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


