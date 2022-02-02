# VcenterVchaIpv6Spec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | IPv6 address to be used to configure the interface. | 
**Prefix** | **int64** | The CIDR prefix for the interface. | 

## Methods

### NewVcenterVchaIpv6Spec

`func NewVcenterVchaIpv6Spec(address string, prefix int64, ) *VcenterVchaIpv6Spec`

NewVcenterVchaIpv6Spec instantiates a new VcenterVchaIpv6Spec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaIpv6SpecWithDefaults

`func NewVcenterVchaIpv6SpecWithDefaults() *VcenterVchaIpv6Spec`

NewVcenterVchaIpv6SpecWithDefaults instantiates a new VcenterVchaIpv6Spec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VcenterVchaIpv6Spec) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterVchaIpv6Spec) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterVchaIpv6Spec) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPrefix

`func (o *VcenterVchaIpv6Spec) GetPrefix() int64`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VcenterVchaIpv6Spec) GetPrefixOk() (*int64, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VcenterVchaIpv6Spec) SetPrefix(v int64)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


