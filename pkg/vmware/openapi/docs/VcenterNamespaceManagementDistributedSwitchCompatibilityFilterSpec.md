# VcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compatible** | Pointer to **bool** | Compatibility criteria for matching the filter. If true, only Distributed Switches which are compatible with vSphere Namespaces match the filter. If false, only Distributed Switches which are incompatible with vSphere Namespaces match the filter. If unset, both compatible and incompatible Distributed Switches match the filter. | [optional] 
**NetworkProvider** | Pointer to [**VcenterNamespaceManagementClustersNetworkProvider**](VcenterNamespaceManagementClustersNetworkProvider.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec

`func NewVcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec() *VcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec`

NewVcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec instantiates a new VcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpecWithDefaults

`func NewVcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpecWithDefaults() *VcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec`

NewVcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpecWithDefaults instantiates a new VcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatible

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec) GetCompatible() bool`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec) GetCompatibleOk() (*bool, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec) SetCompatible(v bool)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetNetworkProvider

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec) GetNetworkProvider() VcenterNamespaceManagementClustersNetworkProvider`

GetNetworkProvider returns the NetworkProvider field if non-nil, zero value otherwise.

### GetNetworkProviderOk

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec) GetNetworkProviderOk() (*VcenterNamespaceManagementClustersNetworkProvider, bool)`

GetNetworkProviderOk returns a tuple with the NetworkProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProvider

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec) SetNetworkProvider(v VcenterNamespaceManagementClustersNetworkProvider)`

SetNetworkProvider sets NetworkProvider field to given value.

### HasNetworkProvider

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilityFilterSpec) HasNetworkProvider() bool`

HasNetworkProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


