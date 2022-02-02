# VcenterVmGuestNetworkingInterfacesIpAddressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** | IPv4 address is specified using dotted decimal notation. For example, \&quot;192.0.2.1\&quot;. IPv6 addresses are 128-bit addresses specified using eight fields of up to four hexadecimal digits. A colon separates each field (:). For example, 2001:DB8:101::230:6eff:fe04:d9ff. The address can also consist of the symbol &#39;::&#39; to represent multiple 16-bit groups of contiguous 0&#39;s only once in an address as described in RFC 2373. | 
**PrefixLength** | **int64** | Denotes the length of a generic Internet network address prefix. Prefix length: the valid range of values is 0-32 for IPv4, and 0-128 for IPv6. A value of n corresponds to an IP address mask that has n contiguous 1-bits from the most significant bit (MSB), with all other bits set to 0. A value of zero is valid only if the calling context defines it. | 
**Origin** | Pointer to [**VcenterVmGuestNetworkingInterfacesIpAddressOrigin**](VcenterVmGuestNetworkingInterfacesIpAddressOrigin.md) |  | [optional] 
**State** | [**VcenterVmGuestNetworkingInterfacesIpAddressStatus**](VcenterVmGuestNetworkingInterfacesIpAddressStatus.md) |  | 

## Methods

### NewVcenterVmGuestNetworkingInterfacesIpAddressInfo

`func NewVcenterVmGuestNetworkingInterfacesIpAddressInfo(ipAddress string, prefixLength int64, state VcenterVmGuestNetworkingInterfacesIpAddressStatus, ) *VcenterVmGuestNetworkingInterfacesIpAddressInfo`

NewVcenterVmGuestNetworkingInterfacesIpAddressInfo instantiates a new VcenterVmGuestNetworkingInterfacesIpAddressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestNetworkingInterfacesIpAddressInfoWithDefaults

`func NewVcenterVmGuestNetworkingInterfacesIpAddressInfoWithDefaults() *VcenterVmGuestNetworkingInterfacesIpAddressInfo`

NewVcenterVmGuestNetworkingInterfacesIpAddressInfoWithDefaults instantiates a new VcenterVmGuestNetworkingInterfacesIpAddressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetPrefixLength

`func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetPrefixLength() int64`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetPrefixLengthOk() (*int64, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) SetPrefixLength(v int64)`

SetPrefixLength sets PrefixLength field to given value.


### GetOrigin

`func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetOrigin() VcenterVmGuestNetworkingInterfacesIpAddressOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetOriginOk() (*VcenterVmGuestNetworkingInterfacesIpAddressOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) SetOrigin(v VcenterVmGuestNetworkingInterfacesIpAddressOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetState

`func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetState() VcenterVmGuestNetworkingInterfacesIpAddressStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetStateOk() (*VcenterVmGuestNetworkingInterfacesIpAddressStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) SetState(v VcenterVmGuestNetworkingInterfacesIpAddressStatus)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


