# VcenterGuestWindowsNetworkAdapterSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsServers** | Pointer to **[]string** | A list of server IP addresses to use for DNS lookup in a Windows guest operating system.   These servers should be specified in the order of preference. If this list is not empty, and if a DHCP IP address is used, then these settings override the DHCP settings.  If unset, no DNS servers are set. | [optional] 
**DnsDomain** | Pointer to **string** | A DNS domain suffix such as vmware.com. If unset, no DNS domain is set. | [optional] 
**WinsServers** | Pointer to **[]string** | List of WINS Servers to set for the Windows guest operating system. A Maximum of two IP addresses can be specified in this list. The first IP address will be set as the primary WINS server. The second IP address will be set as the secondary WINS server. If unset, no WINS Servers are set. | [optional] 
**NetBIOSMode** | Pointer to [**VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode**](VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode.md) |  | [optional] 

## Methods

### NewVcenterGuestWindowsNetworkAdapterSettings

`func NewVcenterGuestWindowsNetworkAdapterSettings() *VcenterGuestWindowsNetworkAdapterSettings`

NewVcenterGuestWindowsNetworkAdapterSettings instantiates a new VcenterGuestWindowsNetworkAdapterSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestWindowsNetworkAdapterSettingsWithDefaults

`func NewVcenterGuestWindowsNetworkAdapterSettingsWithDefaults() *VcenterGuestWindowsNetworkAdapterSettings`

NewVcenterGuestWindowsNetworkAdapterSettingsWithDefaults instantiates a new VcenterGuestWindowsNetworkAdapterSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsServers

`func (o *VcenterGuestWindowsNetworkAdapterSettings) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *VcenterGuestWindowsNetworkAdapterSettings) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *VcenterGuestWindowsNetworkAdapterSettings) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *VcenterGuestWindowsNetworkAdapterSettings) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsDomain

`func (o *VcenterGuestWindowsNetworkAdapterSettings) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *VcenterGuestWindowsNetworkAdapterSettings) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *VcenterGuestWindowsNetworkAdapterSettings) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *VcenterGuestWindowsNetworkAdapterSettings) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.

### GetWinsServers

`func (o *VcenterGuestWindowsNetworkAdapterSettings) GetWinsServers() []string`

GetWinsServers returns the WinsServers field if non-nil, zero value otherwise.

### GetWinsServersOk

`func (o *VcenterGuestWindowsNetworkAdapterSettings) GetWinsServersOk() (*[]string, bool)`

GetWinsServersOk returns a tuple with the WinsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsServers

`func (o *VcenterGuestWindowsNetworkAdapterSettings) SetWinsServers(v []string)`

SetWinsServers sets WinsServers field to given value.

### HasWinsServers

`func (o *VcenterGuestWindowsNetworkAdapterSettings) HasWinsServers() bool`

HasWinsServers returns a boolean if a field has been set.

### GetNetBIOSMode

`func (o *VcenterGuestWindowsNetworkAdapterSettings) GetNetBIOSMode() VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode`

GetNetBIOSMode returns the NetBIOSMode field if non-nil, zero value otherwise.

### GetNetBIOSModeOk

`func (o *VcenterGuestWindowsNetworkAdapterSettings) GetNetBIOSModeOk() (*VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode, bool)`

GetNetBIOSModeOk returns a tuple with the NetBIOSMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetBIOSMode

`func (o *VcenterGuestWindowsNetworkAdapterSettings) SetNetBIOSMode(v VcenterGuestWindowsNetworkAdapterSettingsNetBIOSMode)`

SetNetBIOSMode sets NetBIOSMode field to given value.

### HasNetBIOSMode

`func (o *VcenterGuestWindowsNetworkAdapterSettings) HasNetBIOSMode() bool`

HasNetBIOSMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


