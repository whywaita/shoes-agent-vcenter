# VcenterNamespaceManagementLoadBalancersSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRanges** | [**[]VcenterNamespaceManagementIPRange**](VcenterNamespaceManagementIPRange.md) | List of address ranges that will be used to derive frontend IP addresses for L4 virtual servers. At least one range must be provided. A set operation only allows for addition of new IP ranges to the existing list of IP ranges. | 
**Provider** | [**VcenterNamespaceManagementLoadBalancersProvider**](VcenterNamespaceManagementLoadBalancersProvider.md) |  | 
**HaProxyConfigSetSpec** | Pointer to [**VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec**](VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec.md) |  | [optional] 
**AviConfigSetSpec** | Pointer to [**VcenterNamespaceManagementLoadBalancersAviConfigSetSpec**](VcenterNamespaceManagementLoadBalancersAviConfigSetSpec.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementLoadBalancersSetSpec

`func NewVcenterNamespaceManagementLoadBalancersSetSpec(addressRanges []VcenterNamespaceManagementIPRange, provider VcenterNamespaceManagementLoadBalancersProvider, ) *VcenterNamespaceManagementLoadBalancersSetSpec`

NewVcenterNamespaceManagementLoadBalancersSetSpec instantiates a new VcenterNamespaceManagementLoadBalancersSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementLoadBalancersSetSpecWithDefaults

`func NewVcenterNamespaceManagementLoadBalancersSetSpecWithDefaults() *VcenterNamespaceManagementLoadBalancersSetSpec`

NewVcenterNamespaceManagementLoadBalancersSetSpecWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRanges

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) GetAddressRanges() []VcenterNamespaceManagementIPRange`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) GetAddressRangesOk() (*[]VcenterNamespaceManagementIPRange, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) SetAddressRanges(v []VcenterNamespaceManagementIPRange)`

SetAddressRanges sets AddressRanges field to given value.


### GetProvider

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) GetProvider() VcenterNamespaceManagementLoadBalancersProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) GetProviderOk() (*VcenterNamespaceManagementLoadBalancersProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) SetProvider(v VcenterNamespaceManagementLoadBalancersProvider)`

SetProvider sets Provider field to given value.


### GetHaProxyConfigSetSpec

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) GetHaProxyConfigSetSpec() VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec`

GetHaProxyConfigSetSpec returns the HaProxyConfigSetSpec field if non-nil, zero value otherwise.

### GetHaProxyConfigSetSpecOk

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) GetHaProxyConfigSetSpecOk() (*VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec, bool)`

GetHaProxyConfigSetSpecOk returns a tuple with the HaProxyConfigSetSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaProxyConfigSetSpec

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) SetHaProxyConfigSetSpec(v VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec)`

SetHaProxyConfigSetSpec sets HaProxyConfigSetSpec field to given value.

### HasHaProxyConfigSetSpec

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) HasHaProxyConfigSetSpec() bool`

HasHaProxyConfigSetSpec returns a boolean if a field has been set.

### GetAviConfigSetSpec

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) GetAviConfigSetSpec() VcenterNamespaceManagementLoadBalancersAviConfigSetSpec`

GetAviConfigSetSpec returns the AviConfigSetSpec field if non-nil, zero value otherwise.

### GetAviConfigSetSpecOk

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) GetAviConfigSetSpecOk() (*VcenterNamespaceManagementLoadBalancersAviConfigSetSpec, bool)`

GetAviConfigSetSpecOk returns a tuple with the AviConfigSetSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAviConfigSetSpec

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) SetAviConfigSetSpec(v VcenterNamespaceManagementLoadBalancersAviConfigSetSpec)`

SetAviConfigSetSpec sets AviConfigSetSpec field to given value.

### HasAviConfigSetSpec

`func (o *VcenterNamespaceManagementLoadBalancersSetSpec) HasAviConfigSetSpec() bool`

HasAviConfigSetSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


