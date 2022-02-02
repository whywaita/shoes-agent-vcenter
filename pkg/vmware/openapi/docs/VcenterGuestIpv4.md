# VcenterGuestIpv4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterGuestIpv4Type**](VcenterGuestIpv4Type.md) |  | 
**IpAddress** | Pointer to **string** | The static IPv4 address This field is optional and it is only relevant when the value of Ipv4.type is STATIC. | [optional] 
**Prefix** | Pointer to **int64** | The IPv4 CIDR prefix, for example, 24. See http://www.oav.net/mirrors/cidr.html for netmask-to-prefix conversion. This field is optional and it is only relevant when the value of Ipv4.type is STATIC. | [optional] 
**Gateways** | Pointer to **[]string** | Gateways for the IPv4 address. If unset, no gateways are set. | [optional] 

## Methods

### NewVcenterGuestIpv4

`func NewVcenterGuestIpv4(type_ VcenterGuestIpv4Type, ) *VcenterGuestIpv4`

NewVcenterGuestIpv4 instantiates a new VcenterGuestIpv4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestIpv4WithDefaults

`func NewVcenterGuestIpv4WithDefaults() *VcenterGuestIpv4`

NewVcenterGuestIpv4WithDefaults instantiates a new VcenterGuestIpv4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterGuestIpv4) GetType() VcenterGuestIpv4Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterGuestIpv4) GetTypeOk() (*VcenterGuestIpv4Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterGuestIpv4) SetType(v VcenterGuestIpv4Type)`

SetType sets Type field to given value.


### GetIpAddress

`func (o *VcenterGuestIpv4) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VcenterGuestIpv4) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VcenterGuestIpv4) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VcenterGuestIpv4) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPrefix

`func (o *VcenterGuestIpv4) GetPrefix() int64`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VcenterGuestIpv4) GetPrefixOk() (*int64, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VcenterGuestIpv4) SetPrefix(v int64)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *VcenterGuestIpv4) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetGateways

`func (o *VcenterGuestIpv4) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *VcenterGuestIpv4) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *VcenterGuestIpv4) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *VcenterGuestIpv4) HasGateways() bool`

HasGateways returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


