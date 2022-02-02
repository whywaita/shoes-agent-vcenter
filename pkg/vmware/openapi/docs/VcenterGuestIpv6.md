# VcenterGuestIpv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterGuestIpv6Type**](VcenterGuestIpv6Type.md) |  | 
**Ipv6** | Pointer to [**[]VcenterGuestIpv6Address**](VcenterGuestIpv6Address.md) | IPv6 address This field is optional and it is only relevant when the value of Ipv6.type is STATIC. | [optional] 
**Gateways** | Pointer to **[]string** | gateways for the IPv6 address. If unset, no gateways are set. | [optional] 

## Methods

### NewVcenterGuestIpv6

`func NewVcenterGuestIpv6(type_ VcenterGuestIpv6Type, ) *VcenterGuestIpv6`

NewVcenterGuestIpv6 instantiates a new VcenterGuestIpv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestIpv6WithDefaults

`func NewVcenterGuestIpv6WithDefaults() *VcenterGuestIpv6`

NewVcenterGuestIpv6WithDefaults instantiates a new VcenterGuestIpv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterGuestIpv6) GetType() VcenterGuestIpv6Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterGuestIpv6) GetTypeOk() (*VcenterGuestIpv6Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterGuestIpv6) SetType(v VcenterGuestIpv6Type)`

SetType sets Type field to given value.


### GetIpv6

`func (o *VcenterGuestIpv6) GetIpv6() []VcenterGuestIpv6Address`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *VcenterGuestIpv6) GetIpv6Ok() (*[]VcenterGuestIpv6Address, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *VcenterGuestIpv6) SetIpv6(v []VcenterGuestIpv6Address)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *VcenterGuestIpv6) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetGateways

`func (o *VcenterGuestIpv6) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *VcenterGuestIpv6) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *VcenterGuestIpv6) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *VcenterGuestIpv6) HasGateways() bool`

HasGateways returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


