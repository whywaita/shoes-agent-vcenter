# VcenterNamespaceManagementClustersNetworkSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FloatingIP** | Pointer to **string** | Optionally specify the Floating IP used by the HA master cluster in the DHCP case. This field is optional and it is only relevant when the value of Clusters.NetworkSpec.mode is DHCP. | [optional] 
**Network** | **string** | Identifier for the network. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network. | 
**Mode** | [**VcenterNamespaceManagementClustersNetworkSpecIpv4Mode**](VcenterNamespaceManagementClustersNetworkSpecIpv4Mode.md) |  | 
**AddressRange** | Pointer to [**VcenterNamespaceManagementClustersIpv4Range**](VcenterNamespaceManagementClustersIpv4Range.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementClustersNetworkSpec

`func NewVcenterNamespaceManagementClustersNetworkSpec(network string, mode VcenterNamespaceManagementClustersNetworkSpecIpv4Mode, ) *VcenterNamespaceManagementClustersNetworkSpec`

NewVcenterNamespaceManagementClustersNetworkSpec instantiates a new VcenterNamespaceManagementClustersNetworkSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersNetworkSpecWithDefaults

`func NewVcenterNamespaceManagementClustersNetworkSpecWithDefaults() *VcenterNamespaceManagementClustersNetworkSpec`

NewVcenterNamespaceManagementClustersNetworkSpecWithDefaults instantiates a new VcenterNamespaceManagementClustersNetworkSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloatingIP

`func (o *VcenterNamespaceManagementClustersNetworkSpec) GetFloatingIP() string`

GetFloatingIP returns the FloatingIP field if non-nil, zero value otherwise.

### GetFloatingIPOk

`func (o *VcenterNamespaceManagementClustersNetworkSpec) GetFloatingIPOk() (*string, bool)`

GetFloatingIPOk returns a tuple with the FloatingIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIP

`func (o *VcenterNamespaceManagementClustersNetworkSpec) SetFloatingIP(v string)`

SetFloatingIP sets FloatingIP field to given value.

### HasFloatingIP

`func (o *VcenterNamespaceManagementClustersNetworkSpec) HasFloatingIP() bool`

HasFloatingIP returns a boolean if a field has been set.

### GetNetwork

`func (o *VcenterNamespaceManagementClustersNetworkSpec) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VcenterNamespaceManagementClustersNetworkSpec) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VcenterNamespaceManagementClustersNetworkSpec) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetMode

`func (o *VcenterNamespaceManagementClustersNetworkSpec) GetMode() VcenterNamespaceManagementClustersNetworkSpecIpv4Mode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VcenterNamespaceManagementClustersNetworkSpec) GetModeOk() (*VcenterNamespaceManagementClustersNetworkSpecIpv4Mode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VcenterNamespaceManagementClustersNetworkSpec) SetMode(v VcenterNamespaceManagementClustersNetworkSpecIpv4Mode)`

SetMode sets Mode field to given value.


### GetAddressRange

`func (o *VcenterNamespaceManagementClustersNetworkSpec) GetAddressRange() VcenterNamespaceManagementClustersIpv4Range`

GetAddressRange returns the AddressRange field if non-nil, zero value otherwise.

### GetAddressRangeOk

`func (o *VcenterNamespaceManagementClustersNetworkSpec) GetAddressRangeOk() (*VcenterNamespaceManagementClustersIpv4Range, bool)`

GetAddressRangeOk returns a tuple with the AddressRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRange

`func (o *VcenterNamespaceManagementClustersNetworkSpec) SetAddressRange(v VcenterNamespaceManagementClustersIpv4Range)`

SetAddressRange sets AddressRange field to given value.

### HasAddressRange

`func (o *VcenterNamespaceManagementClustersNetworkSpec) HasAddressRange() bool`

HasAddressRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


