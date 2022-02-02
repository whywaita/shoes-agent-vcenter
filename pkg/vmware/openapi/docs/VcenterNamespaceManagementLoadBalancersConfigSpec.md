# VcenterNamespaceManagementLoadBalancersConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An identifier that identifies a load balancer and can be used to query or configure load balancer properties via these resources. The identifier has DNS_LABEL restrictions as specified in . This must be an alphanumeric (a-z and 0-9) string, with a maximum length of 63 characters and with the &#39;-&#39; character allowed anywhere except the first or last character. This name is unique across all Namespaces in this vCenter server. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.LoadBalancerConfig. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.LoadBalancerConfig. | 
**AddressRanges** | [**[]VcenterNamespaceManagementIPRange**](VcenterNamespaceManagementIPRange.md) | List of address ranges that will be used to derive frontend IP addresses for L4 virtual servers. At least one range must be provided. | 
**Provider** | [**VcenterNamespaceManagementLoadBalancersProvider**](VcenterNamespaceManagementLoadBalancersProvider.md) |  | 
**HaProxyConfigCreateSpec** | Pointer to [**VcenterNamespaceManagementLoadBalancersHAProxyConfigCreateSpec**](VcenterNamespaceManagementLoadBalancersHAProxyConfigCreateSpec.md) |  | [optional] 
**AviConfigCreateSpec** | Pointer to [**VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec**](VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementLoadBalancersConfigSpec

`func NewVcenterNamespaceManagementLoadBalancersConfigSpec(id string, addressRanges []VcenterNamespaceManagementIPRange, provider VcenterNamespaceManagementLoadBalancersProvider, ) *VcenterNamespaceManagementLoadBalancersConfigSpec`

NewVcenterNamespaceManagementLoadBalancersConfigSpec instantiates a new VcenterNamespaceManagementLoadBalancersConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementLoadBalancersConfigSpecWithDefaults

`func NewVcenterNamespaceManagementLoadBalancersConfigSpecWithDefaults() *VcenterNamespaceManagementLoadBalancersConfigSpec`

NewVcenterNamespaceManagementLoadBalancersConfigSpecWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) SetId(v string)`

SetId sets Id field to given value.


### GetAddressRanges

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) GetAddressRanges() []VcenterNamespaceManagementIPRange`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) GetAddressRangesOk() (*[]VcenterNamespaceManagementIPRange, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) SetAddressRanges(v []VcenterNamespaceManagementIPRange)`

SetAddressRanges sets AddressRanges field to given value.


### GetProvider

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) GetProvider() VcenterNamespaceManagementLoadBalancersProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) GetProviderOk() (*VcenterNamespaceManagementLoadBalancersProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) SetProvider(v VcenterNamespaceManagementLoadBalancersProvider)`

SetProvider sets Provider field to given value.


### GetHaProxyConfigCreateSpec

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) GetHaProxyConfigCreateSpec() VcenterNamespaceManagementLoadBalancersHAProxyConfigCreateSpec`

GetHaProxyConfigCreateSpec returns the HaProxyConfigCreateSpec field if non-nil, zero value otherwise.

### GetHaProxyConfigCreateSpecOk

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) GetHaProxyConfigCreateSpecOk() (*VcenterNamespaceManagementLoadBalancersHAProxyConfigCreateSpec, bool)`

GetHaProxyConfigCreateSpecOk returns a tuple with the HaProxyConfigCreateSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaProxyConfigCreateSpec

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) SetHaProxyConfigCreateSpec(v VcenterNamespaceManagementLoadBalancersHAProxyConfigCreateSpec)`

SetHaProxyConfigCreateSpec sets HaProxyConfigCreateSpec field to given value.

### HasHaProxyConfigCreateSpec

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) HasHaProxyConfigCreateSpec() bool`

HasHaProxyConfigCreateSpec returns a boolean if a field has been set.

### GetAviConfigCreateSpec

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) GetAviConfigCreateSpec() VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec`

GetAviConfigCreateSpec returns the AviConfigCreateSpec field if non-nil, zero value otherwise.

### GetAviConfigCreateSpecOk

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) GetAviConfigCreateSpecOk() (*VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec, bool)`

GetAviConfigCreateSpecOk returns a tuple with the AviConfigCreateSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAviConfigCreateSpec

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) SetAviConfigCreateSpec(v VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec)`

SetAviConfigCreateSpec sets AviConfigCreateSpec field to given value.

### HasAviConfigCreateSpec

`func (o *VcenterNamespaceManagementLoadBalancersConfigSpec) HasAviConfigCreateSpec() bool`

HasAviConfigCreateSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


