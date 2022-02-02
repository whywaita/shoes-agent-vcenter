# VcenterNamespaceManagementLoadBalancersUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRanges** | Pointer to [**[]VcenterNamespaceManagementIPRange**](VcenterNamespaceManagementIPRange.md) | List of address ranges that will be used to derive frontend IP addresses for L4 virtual servers. At least one range must be provided. An update operation only allows for addition of new IP ranges to the existing list of IP ranges. If unset, the existing list of address ranges will not be modified. | [optional] 
**Provider** | Pointer to [**VcenterNamespaceManagementLoadBalancersProvider**](VcenterNamespaceManagementLoadBalancersProvider.md) |  | [optional] 
**HaProxyConfigUpdateSpec** | Pointer to [**VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec**](VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec.md) |  | [optional] 
**AviConfigUpdateSpec** | Pointer to [**VcenterNamespaceManagementLoadBalancersAviConfigUpdateSpec**](VcenterNamespaceManagementLoadBalancersAviConfigUpdateSpec.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementLoadBalancersUpdateSpec

`func NewVcenterNamespaceManagementLoadBalancersUpdateSpec() *VcenterNamespaceManagementLoadBalancersUpdateSpec`

NewVcenterNamespaceManagementLoadBalancersUpdateSpec instantiates a new VcenterNamespaceManagementLoadBalancersUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementLoadBalancersUpdateSpecWithDefaults

`func NewVcenterNamespaceManagementLoadBalancersUpdateSpecWithDefaults() *VcenterNamespaceManagementLoadBalancersUpdateSpec`

NewVcenterNamespaceManagementLoadBalancersUpdateSpecWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRanges

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) GetAddressRanges() []VcenterNamespaceManagementIPRange`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) GetAddressRangesOk() (*[]VcenterNamespaceManagementIPRange, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) SetAddressRanges(v []VcenterNamespaceManagementIPRange)`

SetAddressRanges sets AddressRanges field to given value.

### HasAddressRanges

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) HasAddressRanges() bool`

HasAddressRanges returns a boolean if a field has been set.

### GetProvider

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) GetProvider() VcenterNamespaceManagementLoadBalancersProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) GetProviderOk() (*VcenterNamespaceManagementLoadBalancersProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) SetProvider(v VcenterNamespaceManagementLoadBalancersProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetHaProxyConfigUpdateSpec

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) GetHaProxyConfigUpdateSpec() VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec`

GetHaProxyConfigUpdateSpec returns the HaProxyConfigUpdateSpec field if non-nil, zero value otherwise.

### GetHaProxyConfigUpdateSpecOk

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) GetHaProxyConfigUpdateSpecOk() (*VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec, bool)`

GetHaProxyConfigUpdateSpecOk returns a tuple with the HaProxyConfigUpdateSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaProxyConfigUpdateSpec

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) SetHaProxyConfigUpdateSpec(v VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec)`

SetHaProxyConfigUpdateSpec sets HaProxyConfigUpdateSpec field to given value.

### HasHaProxyConfigUpdateSpec

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) HasHaProxyConfigUpdateSpec() bool`

HasHaProxyConfigUpdateSpec returns a boolean if a field has been set.

### GetAviConfigUpdateSpec

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) GetAviConfigUpdateSpec() VcenterNamespaceManagementLoadBalancersAviConfigUpdateSpec`

GetAviConfigUpdateSpec returns the AviConfigUpdateSpec field if non-nil, zero value otherwise.

### GetAviConfigUpdateSpecOk

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) GetAviConfigUpdateSpecOk() (*VcenterNamespaceManagementLoadBalancersAviConfigUpdateSpec, bool)`

GetAviConfigUpdateSpecOk returns a tuple with the AviConfigUpdateSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAviConfigUpdateSpec

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) SetAviConfigUpdateSpec(v VcenterNamespaceManagementLoadBalancersAviConfigUpdateSpec)`

SetAviConfigUpdateSpec sets AviConfigUpdateSpec field to given value.

### HasAviConfigUpdateSpec

`func (o *VcenterNamespaceManagementLoadBalancersUpdateSpec) HasAviConfigUpdateSpec() bool`

HasAviConfigUpdateSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


