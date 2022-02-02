# VcenterGuestIpv6Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** | Static IPv6 Address. | 
**Prefix** | **int64** | The CIDR prefix for the interface. | 

## Methods

### NewVcenterGuestIpv6Address

`func NewVcenterGuestIpv6Address(ipAddress string, prefix int64, ) *VcenterGuestIpv6Address`

NewVcenterGuestIpv6Address instantiates a new VcenterGuestIpv6Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestIpv6AddressWithDefaults

`func NewVcenterGuestIpv6AddressWithDefaults() *VcenterGuestIpv6Address`

NewVcenterGuestIpv6AddressWithDefaults instantiates a new VcenterGuestIpv6Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *VcenterGuestIpv6Address) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VcenterGuestIpv6Address) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VcenterGuestIpv6Address) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetPrefix

`func (o *VcenterGuestIpv6Address) GetPrefix() int64`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VcenterGuestIpv6Address) GetPrefixOk() (*int64, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VcenterGuestIpv6Address) SetPrefix(v int64)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


