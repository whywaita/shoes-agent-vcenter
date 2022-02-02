# VcenterGuestGlobalDNSSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsSuffixList** | Pointer to **[]string** | List of name resolution suffixes for the virtual network adapter. This list applies to both Windows and Linux guest customization. For Linux, this setting is global, whereas in Windows, this setting is listed on a per-adapter basis. If unset, no DNS suffixes are set. | [optional] 
**DnsServers** | Pointer to **[]string** | List of DNS servers, for a virtual network adapter with a static IP address. If this list is empty, then the guest operating system is expected to use a DHCP server to get its DNS server settings. These settings configure the virtual machine to use the specified DNS servers. These DNS server settings are listed in the order of preference. If unset, no DNS servers are set. | [optional] 

## Methods

### NewVcenterGuestGlobalDNSSettings

`func NewVcenterGuestGlobalDNSSettings() *VcenterGuestGlobalDNSSettings`

NewVcenterGuestGlobalDNSSettings instantiates a new VcenterGuestGlobalDNSSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestGlobalDNSSettingsWithDefaults

`func NewVcenterGuestGlobalDNSSettingsWithDefaults() *VcenterGuestGlobalDNSSettings`

NewVcenterGuestGlobalDNSSettingsWithDefaults instantiates a new VcenterGuestGlobalDNSSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsSuffixList

`func (o *VcenterGuestGlobalDNSSettings) GetDnsSuffixList() []string`

GetDnsSuffixList returns the DnsSuffixList field if non-nil, zero value otherwise.

### GetDnsSuffixListOk

`func (o *VcenterGuestGlobalDNSSettings) GetDnsSuffixListOk() (*[]string, bool)`

GetDnsSuffixListOk returns a tuple with the DnsSuffixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffixList

`func (o *VcenterGuestGlobalDNSSettings) SetDnsSuffixList(v []string)`

SetDnsSuffixList sets DnsSuffixList field to given value.

### HasDnsSuffixList

`func (o *VcenterGuestGlobalDNSSettings) HasDnsSuffixList() bool`

HasDnsSuffixList returns a boolean if a field has been set.

### GetDnsServers

`func (o *VcenterGuestGlobalDNSSettings) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *VcenterGuestGlobalDNSSettings) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *VcenterGuestGlobalDNSSettings) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *VcenterGuestGlobalDNSSettings) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


