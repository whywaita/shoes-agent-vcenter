# VcenterVmGuestDnsConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddresses** | **[]string** | The IP addresses of the DNS servers in order of use. IPv4 addresses are specified using dotted decimal notation. For example, \&quot;192.0.2.1\&quot;. IPv6 addresses are 128-bit addresses represented as eight fields of up to four hexadecimal digits. A colon separates each field (:). For example, 2001:DB8:101::230:6eff:fe04:d9ff. The address can also consist of the symbol &#39;::&#39; to represent multiple 16-bit groups of contiguous 0&#39;s only once in an address as described in RFC 2373. | 
**SearchDomains** | **[]string** | The domain in which to search for hosts, placed in order of preference. These are the domain name portion of the DNS names. | 

## Methods

### NewVcenterVmGuestDnsConfigInfo

`func NewVcenterVmGuestDnsConfigInfo(ipAddresses []string, searchDomains []string, ) *VcenterVmGuestDnsConfigInfo`

NewVcenterVmGuestDnsConfigInfo instantiates a new VcenterVmGuestDnsConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestDnsConfigInfoWithDefaults

`func NewVcenterVmGuestDnsConfigInfoWithDefaults() *VcenterVmGuestDnsConfigInfo`

NewVcenterVmGuestDnsConfigInfoWithDefaults instantiates a new VcenterVmGuestDnsConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddresses

`func (o *VcenterVmGuestDnsConfigInfo) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *VcenterVmGuestDnsConfigInfo) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *VcenterVmGuestDnsConfigInfo) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.


### GetSearchDomains

`func (o *VcenterVmGuestDnsConfigInfo) GetSearchDomains() []string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *VcenterVmGuestDnsConfigInfo) GetSearchDomainsOk() (*[]string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *VcenterVmGuestDnsConfigInfo) SetSearchDomains(v []string)`

SetSearchDomains sets SearchDomains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


