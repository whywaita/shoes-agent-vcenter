# VcenterVmGuestNetworkingInterfacesIpConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddresses** | [**[]VcenterVmGuestNetworkingInterfacesIpAddressInfo**](VcenterVmGuestNetworkingInterfacesIpAddressInfo.md) | IP addresses configured on the interface. | 
**Dhcp** | Pointer to [**VcenterVmGuestDhcpConfigInfo**](VcenterVmGuestDhcpConfigInfo.md) |  | [optional] 

## Methods

### NewVcenterVmGuestNetworkingInterfacesIpConfigInfo

`func NewVcenterVmGuestNetworkingInterfacesIpConfigInfo(ipAddresses []VcenterVmGuestNetworkingInterfacesIpAddressInfo, ) *VcenterVmGuestNetworkingInterfacesIpConfigInfo`

NewVcenterVmGuestNetworkingInterfacesIpConfigInfo instantiates a new VcenterVmGuestNetworkingInterfacesIpConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestNetworkingInterfacesIpConfigInfoWithDefaults

`func NewVcenterVmGuestNetworkingInterfacesIpConfigInfoWithDefaults() *VcenterVmGuestNetworkingInterfacesIpConfigInfo`

NewVcenterVmGuestNetworkingInterfacesIpConfigInfoWithDefaults instantiates a new VcenterVmGuestNetworkingInterfacesIpConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddresses

`func (o *VcenterVmGuestNetworkingInterfacesIpConfigInfo) GetIpAddresses() []VcenterVmGuestNetworkingInterfacesIpAddressInfo`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *VcenterVmGuestNetworkingInterfacesIpConfigInfo) GetIpAddressesOk() (*[]VcenterVmGuestNetworkingInterfacesIpAddressInfo, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *VcenterVmGuestNetworkingInterfacesIpConfigInfo) SetIpAddresses(v []VcenterVmGuestNetworkingInterfacesIpAddressInfo)`

SetIpAddresses sets IpAddresses field to given value.


### GetDhcp

`func (o *VcenterVmGuestNetworkingInterfacesIpConfigInfo) GetDhcp() VcenterVmGuestDhcpConfigInfo`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *VcenterVmGuestNetworkingInterfacesIpConfigInfo) GetDhcpOk() (*VcenterVmGuestDhcpConfigInfo, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *VcenterVmGuestNetworkingInterfacesIpConfigInfo) SetDhcp(v VcenterVmGuestDhcpConfigInfo)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *VcenterVmGuestNetworkingInterfacesIpConfigInfo) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


