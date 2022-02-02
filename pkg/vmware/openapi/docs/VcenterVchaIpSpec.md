# VcenterVchaIpSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpFamily** | [**VcenterVchaIpFamily**](VcenterVchaIpFamily.md) |  | 
**Ipv4** | Pointer to [**VcenterVchaIpv4Spec**](VcenterVchaIpv4Spec.md) |  | [optional] 
**Ipv6** | Pointer to [**VcenterVchaIpv6Spec**](VcenterVchaIpv6Spec.md) |  | [optional] 
**DefaultGateway** | Pointer to **string** | The IP address of the Gateway for this interface. If unset, gateway will not be used for the network interface. | [optional] 
**DnsServers** | Pointer to **[]string** | The list of IP addresses of the DNS servers for this interface. This list is a comma separated list. If unset, DNS servers will not be used for the network interface. | [optional] 

## Methods

### NewVcenterVchaIpSpec

`func NewVcenterVchaIpSpec(ipFamily VcenterVchaIpFamily, ) *VcenterVchaIpSpec`

NewVcenterVchaIpSpec instantiates a new VcenterVchaIpSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaIpSpecWithDefaults

`func NewVcenterVchaIpSpecWithDefaults() *VcenterVchaIpSpec`

NewVcenterVchaIpSpecWithDefaults instantiates a new VcenterVchaIpSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpFamily

`func (o *VcenterVchaIpSpec) GetIpFamily() VcenterVchaIpFamily`

GetIpFamily returns the IpFamily field if non-nil, zero value otherwise.

### GetIpFamilyOk

`func (o *VcenterVchaIpSpec) GetIpFamilyOk() (*VcenterVchaIpFamily, bool)`

GetIpFamilyOk returns a tuple with the IpFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFamily

`func (o *VcenterVchaIpSpec) SetIpFamily(v VcenterVchaIpFamily)`

SetIpFamily sets IpFamily field to given value.


### GetIpv4

`func (o *VcenterVchaIpSpec) GetIpv4() VcenterVchaIpv4Spec`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *VcenterVchaIpSpec) GetIpv4Ok() (*VcenterVchaIpv4Spec, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *VcenterVchaIpSpec) SetIpv4(v VcenterVchaIpv4Spec)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *VcenterVchaIpSpec) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *VcenterVchaIpSpec) GetIpv6() VcenterVchaIpv6Spec`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *VcenterVchaIpSpec) GetIpv6Ok() (*VcenterVchaIpv6Spec, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *VcenterVchaIpSpec) SetIpv6(v VcenterVchaIpv6Spec)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *VcenterVchaIpSpec) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetDefaultGateway

`func (o *VcenterVchaIpSpec) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *VcenterVchaIpSpec) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *VcenterVchaIpSpec) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *VcenterVchaIpSpec) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.

### GetDnsServers

`func (o *VcenterVchaIpSpec) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *VcenterVchaIpSpec) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *VcenterVchaIpSpec) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *VcenterVchaIpSpec) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


