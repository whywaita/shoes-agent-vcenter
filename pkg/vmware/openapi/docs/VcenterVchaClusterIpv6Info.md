# VcenterVchaClusterIpv6Info

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | IP address of the configured network interface. | 
**Prefix** | **int64** | The CIDR prefix of the interface. | 

## Methods

### NewVcenterVchaClusterIpv6Info

`func NewVcenterVchaClusterIpv6Info(address string, prefix int64, ) *VcenterVchaClusterIpv6Info`

NewVcenterVchaClusterIpv6Info instantiates a new VcenterVchaClusterIpv6Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterIpv6InfoWithDefaults

`func NewVcenterVchaClusterIpv6InfoWithDefaults() *VcenterVchaClusterIpv6Info`

NewVcenterVchaClusterIpv6InfoWithDefaults instantiates a new VcenterVchaClusterIpv6Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VcenterVchaClusterIpv6Info) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterVchaClusterIpv6Info) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterVchaClusterIpv6Info) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPrefix

`func (o *VcenterVchaClusterIpv6Info) GetPrefix() int64`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VcenterVchaClusterIpv6Info) GetPrefixOk() (*int64, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VcenterVchaClusterIpv6Info) SetPrefix(v int64)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


