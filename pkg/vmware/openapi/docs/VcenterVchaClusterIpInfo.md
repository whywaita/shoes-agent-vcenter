# VcenterVchaClusterIpInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpFamily** | [**VcenterVchaClusterIpFamily**](VcenterVchaClusterIpFamily.md) |  | 
**Ipv4** | Pointer to [**VcenterVchaClusterIpv4Info**](VcenterVchaClusterIpv4Info.md) |  | [optional] 
**Ipv6** | Pointer to [**VcenterVchaClusterIpv6Info**](VcenterVchaClusterIpv6Info.md) |  | [optional] 
**GatewayIp** | Pointer to **string** | Gateway IP address. If unset, no gateway is specified. | [optional] 

## Methods

### NewVcenterVchaClusterIpInfo

`func NewVcenterVchaClusterIpInfo(ipFamily VcenterVchaClusterIpFamily, ) *VcenterVchaClusterIpInfo`

NewVcenterVchaClusterIpInfo instantiates a new VcenterVchaClusterIpInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterIpInfoWithDefaults

`func NewVcenterVchaClusterIpInfoWithDefaults() *VcenterVchaClusterIpInfo`

NewVcenterVchaClusterIpInfoWithDefaults instantiates a new VcenterVchaClusterIpInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpFamily

`func (o *VcenterVchaClusterIpInfo) GetIpFamily() VcenterVchaClusterIpFamily`

GetIpFamily returns the IpFamily field if non-nil, zero value otherwise.

### GetIpFamilyOk

`func (o *VcenterVchaClusterIpInfo) GetIpFamilyOk() (*VcenterVchaClusterIpFamily, bool)`

GetIpFamilyOk returns a tuple with the IpFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFamily

`func (o *VcenterVchaClusterIpInfo) SetIpFamily(v VcenterVchaClusterIpFamily)`

SetIpFamily sets IpFamily field to given value.


### GetIpv4

`func (o *VcenterVchaClusterIpInfo) GetIpv4() VcenterVchaClusterIpv4Info`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *VcenterVchaClusterIpInfo) GetIpv4Ok() (*VcenterVchaClusterIpv4Info, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *VcenterVchaClusterIpInfo) SetIpv4(v VcenterVchaClusterIpv4Info)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *VcenterVchaClusterIpInfo) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *VcenterVchaClusterIpInfo) GetIpv6() VcenterVchaClusterIpv6Info`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *VcenterVchaClusterIpInfo) GetIpv6Ok() (*VcenterVchaClusterIpv6Info, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *VcenterVchaClusterIpInfo) SetIpv6(v VcenterVchaClusterIpv6Info)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *VcenterVchaClusterIpInfo) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetGatewayIp

`func (o *VcenterVchaClusterIpInfo) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *VcenterVchaClusterIpInfo) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *VcenterVchaClusterIpInfo) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *VcenterVchaClusterIpInfo) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


