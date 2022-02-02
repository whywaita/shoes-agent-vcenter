# VcenterNamespaceManagementEdgeClusterCompatibilityFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compatible** | Pointer to **bool** | Compatibility criteria for matching the filter. If true, only Edge Clusters which are compatible with vSphere Namespaces match the filter. If false, only Edge Clusters which are incompatible with vSphere Namespaces match the filter. If unset, both compatible and incompatible Edge Clusters match the filter. | [optional] 

## Methods

### NewVcenterNamespaceManagementEdgeClusterCompatibilityFilterSpec

`func NewVcenterNamespaceManagementEdgeClusterCompatibilityFilterSpec() *VcenterNamespaceManagementEdgeClusterCompatibilityFilterSpec`

NewVcenterNamespaceManagementEdgeClusterCompatibilityFilterSpec instantiates a new VcenterNamespaceManagementEdgeClusterCompatibilityFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementEdgeClusterCompatibilityFilterSpecWithDefaults

`func NewVcenterNamespaceManagementEdgeClusterCompatibilityFilterSpecWithDefaults() *VcenterNamespaceManagementEdgeClusterCompatibilityFilterSpec`

NewVcenterNamespaceManagementEdgeClusterCompatibilityFilterSpecWithDefaults instantiates a new VcenterNamespaceManagementEdgeClusterCompatibilityFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatible

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilityFilterSpec) GetCompatible() bool`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilityFilterSpec) GetCompatibleOk() (*bool, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilityFilterSpec) SetCompatible(v bool)`

SetCompatible sets Compatible field to given value.

### HasCompatible

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilityFilterSpec) HasCompatible() bool`

HasCompatible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


