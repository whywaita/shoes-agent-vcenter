# VcenterNamespaceManagementEdgeClusterCompatibilitySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeCluster** | **string** | Identifier of the Edge Cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: NSXEdgeCluster. When operations return a value of this structure as a result, the field will be an identifier for the resource type: NSXEdgeCluster. | 
**DisplayName** | **string** | Display name of the Edge Cluster. | 
**Compatible** | **bool** | Compatibility of this Edge Cluster with Namespaces feature. | 
**IncompatibilityReasons** | Pointer to [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | List of reasons for incompatibility. If unset, this Edge Cluster is compatible. | [optional] 

## Methods

### NewVcenterNamespaceManagementEdgeClusterCompatibilitySummary

`func NewVcenterNamespaceManagementEdgeClusterCompatibilitySummary(edgeCluster string, displayName string, compatible bool, ) *VcenterNamespaceManagementEdgeClusterCompatibilitySummary`

NewVcenterNamespaceManagementEdgeClusterCompatibilitySummary instantiates a new VcenterNamespaceManagementEdgeClusterCompatibilitySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementEdgeClusterCompatibilitySummaryWithDefaults

`func NewVcenterNamespaceManagementEdgeClusterCompatibilitySummaryWithDefaults() *VcenterNamespaceManagementEdgeClusterCompatibilitySummary`

NewVcenterNamespaceManagementEdgeClusterCompatibilitySummaryWithDefaults instantiates a new VcenterNamespaceManagementEdgeClusterCompatibilitySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeCluster

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilitySummary) GetEdgeCluster() string`

GetEdgeCluster returns the EdgeCluster field if non-nil, zero value otherwise.

### GetEdgeClusterOk

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilitySummary) GetEdgeClusterOk() (*string, bool)`

GetEdgeClusterOk returns a tuple with the EdgeCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCluster

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilitySummary) SetEdgeCluster(v string)`

SetEdgeCluster sets EdgeCluster field to given value.


### GetDisplayName

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilitySummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilitySummary) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilitySummary) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCompatible

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilitySummary) GetCompatible() bool`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilitySummary) GetCompatibleOk() (*bool, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilitySummary) SetCompatible(v bool)`

SetCompatible sets Compatible field to given value.


### GetIncompatibilityReasons

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilitySummary) GetIncompatibilityReasons() []VapiStdLocalizableMessage`

GetIncompatibilityReasons returns the IncompatibilityReasons field if non-nil, zero value otherwise.

### GetIncompatibilityReasonsOk

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilitySummary) GetIncompatibilityReasonsOk() (*[]VapiStdLocalizableMessage, bool)`

GetIncompatibilityReasonsOk returns a tuple with the IncompatibilityReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibilityReasons

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilitySummary) SetIncompatibilityReasons(v []VapiStdLocalizableMessage)`

SetIncompatibilityReasons sets IncompatibilityReasons field to given value.

### HasIncompatibilityReasons

`func (o *VcenterNamespaceManagementEdgeClusterCompatibilitySummary) HasIncompatibilityReasons() bool`

HasIncompatibilityReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


