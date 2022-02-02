# VcenterVchaClusterIpv4Info

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | IP address of the configured network interface. | 
**SubnetMask** | **string** | The subnet mask of the interface. | 
**Prefix** | Pointer to **int64** | The CIDR prefix of the interface. If unset , then the subnet mask is invalid. | [optional] 

## Methods

### NewVcenterVchaClusterIpv4Info

`func NewVcenterVchaClusterIpv4Info(address string, subnetMask string, ) *VcenterVchaClusterIpv4Info`

NewVcenterVchaClusterIpv4Info instantiates a new VcenterVchaClusterIpv4Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterIpv4InfoWithDefaults

`func NewVcenterVchaClusterIpv4InfoWithDefaults() *VcenterVchaClusterIpv4Info`

NewVcenterVchaClusterIpv4InfoWithDefaults instantiates a new VcenterVchaClusterIpv4Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VcenterVchaClusterIpv4Info) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterVchaClusterIpv4Info) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterVchaClusterIpv4Info) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetSubnetMask

`func (o *VcenterVchaClusterIpv4Info) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *VcenterVchaClusterIpv4Info) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *VcenterVchaClusterIpv4Info) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.


### GetPrefix

`func (o *VcenterVchaClusterIpv4Info) GetPrefix() int64`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VcenterVchaClusterIpv4Info) GetPrefixOk() (*int64, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VcenterVchaClusterIpv4Info) SetPrefix(v int64)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *VcenterVchaClusterIpv4Info) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


