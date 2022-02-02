# VcenterVmGuestNetworkingRoutesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | IP Address of the destination IP network. IPv4 address is specified using dotted decimal notation. For example, \&quot;192.0.2.1\&quot;. IPv6 addresses are 128-bit specified using as eight fields of up to four hexadecimal digits. A colon separates each field (:). For example, 2001:DB8:101::230:6eff:fe04:d9ff. The address can also consist of symbol &#39;::&#39; to represent multiple 16-bit groups of contiguous 0&#39;s only once in an address as described in RFC 2373. | 
**PrefixLength** | **int64** | The prefix length. For IPv4 the value range is 0-32. For IPv6 prefixLength is a decimal value range 0-128. The property represents the number of contiguous, higher-order bits of the address that make up the network portion of the IP address. | 
**GatewayAddress** | Pointer to **string** | Where to send the packets for this route. Unicast IP Address of the next hop router. IPv4 address is specified using dotted decimal notation. For example, \&quot;192.0.2.1\&quot;. IPv6 addresses are 128-bit specified using as eight fields of up to four hexadecimal digits. A colon separates each field (:). For example, 2001:DB8:101::230:6eff:fe04:d9ff. The address can also consist of symbol &#39;::&#39; to represent multiple 16-bit groups of contiguous 0&#39;s only once in an address as described in RFC 2373. If unset no gateway is set for the route. | [optional] 
**InterfaceIndex** | Pointer to **int64** | The network interface associated with this route. This is an index into the result of Interfaces.list The index refers to the relative position of an element in a list. For example, an index of 0 refers to the first element in the list while an index of 1 refers to the second element. If unset the route is not associated with a network interface. | [optional] 

## Methods

### NewVcenterVmGuestNetworkingRoutesInfo

`func NewVcenterVmGuestNetworkingRoutesInfo(network string, prefixLength int64, ) *VcenterVmGuestNetworkingRoutesInfo`

NewVcenterVmGuestNetworkingRoutesInfo instantiates a new VcenterVmGuestNetworkingRoutesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestNetworkingRoutesInfoWithDefaults

`func NewVcenterVmGuestNetworkingRoutesInfoWithDefaults() *VcenterVmGuestNetworkingRoutesInfo`

NewVcenterVmGuestNetworkingRoutesInfoWithDefaults instantiates a new VcenterVmGuestNetworkingRoutesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *VcenterVmGuestNetworkingRoutesInfo) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VcenterVmGuestNetworkingRoutesInfo) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VcenterVmGuestNetworkingRoutesInfo) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetPrefixLength

`func (o *VcenterVmGuestNetworkingRoutesInfo) GetPrefixLength() int64`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *VcenterVmGuestNetworkingRoutesInfo) GetPrefixLengthOk() (*int64, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *VcenterVmGuestNetworkingRoutesInfo) SetPrefixLength(v int64)`

SetPrefixLength sets PrefixLength field to given value.


### GetGatewayAddress

`func (o *VcenterVmGuestNetworkingRoutesInfo) GetGatewayAddress() string`

GetGatewayAddress returns the GatewayAddress field if non-nil, zero value otherwise.

### GetGatewayAddressOk

`func (o *VcenterVmGuestNetworkingRoutesInfo) GetGatewayAddressOk() (*string, bool)`

GetGatewayAddressOk returns a tuple with the GatewayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddress

`func (o *VcenterVmGuestNetworkingRoutesInfo) SetGatewayAddress(v string)`

SetGatewayAddress sets GatewayAddress field to given value.

### HasGatewayAddress

`func (o *VcenterVmGuestNetworkingRoutesInfo) HasGatewayAddress() bool`

HasGatewayAddress returns a boolean if a field has been set.

### GetInterfaceIndex

`func (o *VcenterVmGuestNetworkingRoutesInfo) GetInterfaceIndex() int64`

GetInterfaceIndex returns the InterfaceIndex field if non-nil, zero value otherwise.

### GetInterfaceIndexOk

`func (o *VcenterVmGuestNetworkingRoutesInfo) GetInterfaceIndexOk() (*int64, bool)`

GetInterfaceIndexOk returns a tuple with the InterfaceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIndex

`func (o *VcenterVmGuestNetworkingRoutesInfo) SetInterfaceIndex(v int64)`

SetInterfaceIndex sets InterfaceIndex field to given value.

### HasInterfaceIndex

`func (o *VcenterVmGuestNetworkingRoutesInfo) HasInterfaceIndex() bool`

HasInterfaceIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


