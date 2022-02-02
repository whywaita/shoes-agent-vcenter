# VcenterNamespaceManagementClusterCompatibilitySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** | Identifier of the cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | 
**Compatible** | **bool** | Compatibility of this cluster. | 
**IncompatibilityReasons** | [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | Reasons for incompatibility. | 

## Methods

### NewVcenterNamespaceManagementClusterCompatibilitySummary

`func NewVcenterNamespaceManagementClusterCompatibilitySummary(cluster string, compatible bool, incompatibilityReasons []VapiStdLocalizableMessage, ) *VcenterNamespaceManagementClusterCompatibilitySummary`

NewVcenterNamespaceManagementClusterCompatibilitySummary instantiates a new VcenterNamespaceManagementClusterCompatibilitySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClusterCompatibilitySummaryWithDefaults

`func NewVcenterNamespaceManagementClusterCompatibilitySummaryWithDefaults() *VcenterNamespaceManagementClusterCompatibilitySummary`

NewVcenterNamespaceManagementClusterCompatibilitySummaryWithDefaults instantiates a new VcenterNamespaceManagementClusterCompatibilitySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VcenterNamespaceManagementClusterCompatibilitySummary) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterNamespaceManagementClusterCompatibilitySummary) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterNamespaceManagementClusterCompatibilitySummary) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetCompatible

`func (o *VcenterNamespaceManagementClusterCompatibilitySummary) GetCompatible() bool`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *VcenterNamespaceManagementClusterCompatibilitySummary) GetCompatibleOk() (*bool, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *VcenterNamespaceManagementClusterCompatibilitySummary) SetCompatible(v bool)`

SetCompatible sets Compatible field to given value.


### GetIncompatibilityReasons

`func (o *VcenterNamespaceManagementClusterCompatibilitySummary) GetIncompatibilityReasons() []VapiStdLocalizableMessage`

GetIncompatibilityReasons returns the IncompatibilityReasons field if non-nil, zero value otherwise.

### GetIncompatibilityReasonsOk

`func (o *VcenterNamespaceManagementClusterCompatibilitySummary) GetIncompatibilityReasonsOk() (*[]VapiStdLocalizableMessage, bool)`

GetIncompatibilityReasonsOk returns a tuple with the IncompatibilityReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibilityReasons

`func (o *VcenterNamespaceManagementClusterCompatibilitySummary) SetIncompatibilityReasons(v []VapiStdLocalizableMessage)`

SetIncompatibilityReasons sets IncompatibilityReasons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


