# VcenterNamespaceManagementClusterCompatibilityFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compatible** | Pointer to **bool** | Compatibility criteria for matching the filter. If true, only clusters which are compatible for Namespaces match the filter. If false, all clusters match the filter. If unset, both compatible and incompatible clusters match the filter. | [optional] 
**NetworkProvider** | Pointer to [**VcenterNamespaceManagementClustersNetworkProvider**](VcenterNamespaceManagementClustersNetworkProvider.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementClusterCompatibilityFilterSpec

`func NewVcenterNamespaceManagementClusterCompatibilityFilterSpec() *VcenterNamespaceManagementClusterCompatibilityFilterSpec`

NewVcenterNamespaceManagementClusterCompatibilityFilterSpec instantiates a new VcenterNamespaceManagementClusterCompatibilityFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClusterCompatibilityFilterSpecWithDefaults

`func NewVcenterNamespaceManagementClusterCompatibilityFilterSpecWithDefaults() *VcenterNamespaceManagementClusterCompatibilityFilterSpec`

NewVcenterNamespaceManagementClusterCompatibilityFilterSpecWithDefaults instantiates a new VcenterNamespaceManagementClusterCompatibilityFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatible

`func (o *VcenterNamespaceManagementClusterCompatibilityFilterSpec) GetCompatible() bool`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *VcenterNamespaceManagementClusterCompatibilityFilterSpec) GetCompatibleOk() (*bool, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *VcenterNamespaceManagementClusterCompatibilityFilterSpec) SetCompatible(v bool)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *VcenterNamespaceManagementClusterCompatibilityFilterSpec) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.

### GetNetworkProvider

`func (o *VcenterNamespaceManagementClusterCompatibilityFilterSpec) GetNetworkProvider() VcenterNamespaceManagementClustersNetworkProvider`

GetNetworkProvider returns the NetworkProvider field if non-nil, zero value otherwise.

### GetNetworkProviderOk

`func (o *VcenterNamespaceManagementClusterCompatibilityFilterSpec) GetNetworkProviderOk() (*VcenterNamespaceManagementClustersNetworkProvider, bool)`

GetNetworkProviderOk returns a tuple with the NetworkProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProvider

`func (o *VcenterNamespaceManagementClusterCompatibilityFilterSpec) SetNetworkProvider(v VcenterNamespaceManagementClustersNetworkProvider)`

SetNetworkProvider sets NetworkProvider field to given value.

### HasNetworkProvider

`func (o *VcenterNamespaceManagementClusterCompatibilityFilterSpec) HasNetworkProvider() bool`

HasNetworkProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


