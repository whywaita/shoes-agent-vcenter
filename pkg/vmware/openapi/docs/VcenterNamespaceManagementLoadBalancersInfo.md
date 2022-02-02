# VcenterNamespaceManagementLoadBalancersInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An DNS compliant identifier for a load balancer, which can be used to query or configure the load balancer properties. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.LoadBalancerConfig. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.LoadBalancerConfig. | 
**AddressRanges** | [**[]VcenterNamespaceManagementIPRange**](VcenterNamespaceManagementIPRange.md) | IP address range from which virtual servers are assigned their IPs. | 
**Provider** | [**VcenterNamespaceManagementLoadBalancersProvider**](VcenterNamespaceManagementLoadBalancersProvider.md) |  | 
**HaProxyInfo** | Pointer to [**VcenterNamespaceManagementLoadBalancersHAProxyInfo**](VcenterNamespaceManagementLoadBalancersHAProxyInfo.md) |  | [optional] 
**AviInfo** | Pointer to [**VcenterNamespaceManagementLoadBalancersAviInfo**](VcenterNamespaceManagementLoadBalancersAviInfo.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementLoadBalancersInfo

`func NewVcenterNamespaceManagementLoadBalancersInfo(id string, addressRanges []VcenterNamespaceManagementIPRange, provider VcenterNamespaceManagementLoadBalancersProvider, ) *VcenterNamespaceManagementLoadBalancersInfo`

NewVcenterNamespaceManagementLoadBalancersInfo instantiates a new VcenterNamespaceManagementLoadBalancersInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementLoadBalancersInfoWithDefaults

`func NewVcenterNamespaceManagementLoadBalancersInfoWithDefaults() *VcenterNamespaceManagementLoadBalancersInfo`

NewVcenterNamespaceManagementLoadBalancersInfoWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VcenterNamespaceManagementLoadBalancersInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VcenterNamespaceManagementLoadBalancersInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VcenterNamespaceManagementLoadBalancersInfo) SetId(v string)`

SetId sets Id field to given value.


### GetAddressRanges

`func (o *VcenterNamespaceManagementLoadBalancersInfo) GetAddressRanges() []VcenterNamespaceManagementIPRange`

GetAddressRanges returns the AddressRanges field if non-nil, zero value otherwise.

### GetAddressRangesOk

`func (o *VcenterNamespaceManagementLoadBalancersInfo) GetAddressRangesOk() (*[]VcenterNamespaceManagementIPRange, bool)`

GetAddressRangesOk returns a tuple with the AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRanges

`func (o *VcenterNamespaceManagementLoadBalancersInfo) SetAddressRanges(v []VcenterNamespaceManagementIPRange)`

SetAddressRanges sets AddressRanges field to given value.


### GetProvider

`func (o *VcenterNamespaceManagementLoadBalancersInfo) GetProvider() VcenterNamespaceManagementLoadBalancersProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VcenterNamespaceManagementLoadBalancersInfo) GetProviderOk() (*VcenterNamespaceManagementLoadBalancersProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VcenterNamespaceManagementLoadBalancersInfo) SetProvider(v VcenterNamespaceManagementLoadBalancersProvider)`

SetProvider sets Provider field to given value.


### GetHaProxyInfo

`func (o *VcenterNamespaceManagementLoadBalancersInfo) GetHaProxyInfo() VcenterNamespaceManagementLoadBalancersHAProxyInfo`

GetHaProxyInfo returns the HaProxyInfo field if non-nil, zero value otherwise.

### GetHaProxyInfoOk

`func (o *VcenterNamespaceManagementLoadBalancersInfo) GetHaProxyInfoOk() (*VcenterNamespaceManagementLoadBalancersHAProxyInfo, bool)`

GetHaProxyInfoOk returns a tuple with the HaProxyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaProxyInfo

`func (o *VcenterNamespaceManagementLoadBalancersInfo) SetHaProxyInfo(v VcenterNamespaceManagementLoadBalancersHAProxyInfo)`

SetHaProxyInfo sets HaProxyInfo field to given value.

### HasHaProxyInfo

`func (o *VcenterNamespaceManagementLoadBalancersInfo) HasHaProxyInfo() bool`

HasHaProxyInfo returns a boolean if a field has been set.

### GetAviInfo

`func (o *VcenterNamespaceManagementLoadBalancersInfo) GetAviInfo() VcenterNamespaceManagementLoadBalancersAviInfo`

GetAviInfo returns the AviInfo field if non-nil, zero value otherwise.

### GetAviInfoOk

`func (o *VcenterNamespaceManagementLoadBalancersInfo) GetAviInfoOk() (*VcenterNamespaceManagementLoadBalancersAviInfo, bool)`

GetAviInfoOk returns a tuple with the AviInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAviInfo

`func (o *VcenterNamespaceManagementLoadBalancersInfo) SetAviInfo(v VcenterNamespaceManagementLoadBalancersAviInfo)`

SetAviInfo sets AviInfo field to given value.

### HasAviInfo

`func (o *VcenterNamespaceManagementLoadBalancersInfo) HasAviInfo() bool`

HasAviInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


