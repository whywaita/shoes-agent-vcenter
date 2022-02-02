# VcenterVchaIpv4Spec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | IPV4 address to be used to configure the interface. | 
**SubnetMask** | Pointer to **string** | The subnet mask for the interface. If unset and the Ipv4Spec.prefix field is unset, then an error will be reported.  If unset and the Ipv4Spec.prefix field is set, then the Ipv4Spec.prefix field will be used to create a subnet mask whose first prefix bits are 1 and the remaining bits 0.  If both the Ipv4Spec.subnet-mask field and the Ipv4Spec.prefix field are set and they do not represent the same value, then an error will be reported. | [optional] 
**Prefix** | Pointer to **int64** | The CIDR prefix for the interface. If unset and the Ipv4Spec.subnet-mask field is unset, this an error will be reported.  If unset and the Ipv4Spec.subnet-mask field is set, then the Ipv4Spec.subnet-mask field will be used.  If both the Ipv4Spec.subnet-mask field and the Ipv4Spec.prefix field are set and they do not represent the same value, then an error will be reported. | [optional] 

## Methods

### NewVcenterVchaIpv4Spec

`func NewVcenterVchaIpv4Spec(address string, ) *VcenterVchaIpv4Spec`

NewVcenterVchaIpv4Spec instantiates a new VcenterVchaIpv4Spec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaIpv4SpecWithDefaults

`func NewVcenterVchaIpv4SpecWithDefaults() *VcenterVchaIpv4Spec`

NewVcenterVchaIpv4SpecWithDefaults instantiates a new VcenterVchaIpv4Spec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VcenterVchaIpv4Spec) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterVchaIpv4Spec) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterVchaIpv4Spec) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetSubnetMask

`func (o *VcenterVchaIpv4Spec) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *VcenterVchaIpv4Spec) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *VcenterVchaIpv4Spec) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *VcenterVchaIpv4Spec) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetPrefix

`func (o *VcenterVchaIpv4Spec) GetPrefix() int64`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VcenterVchaIpv4Spec) GetPrefixOk() (*int64, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VcenterVchaIpv4Spec) SetPrefix(v int64)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *VcenterVchaIpv4Spec) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


