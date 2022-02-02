# VcenterNamespaceManagementStatsTimeSeriesPodIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** | The namespace that the pod is running in. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespaces.Instance. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespaces.Instance. | 
**PodName** | **string** | The name of the pod itself. | 

## Methods

### NewVcenterNamespaceManagementStatsTimeSeriesPodIdentifier

`func NewVcenterNamespaceManagementStatsTimeSeriesPodIdentifier(namespace string, podName string, ) *VcenterNamespaceManagementStatsTimeSeriesPodIdentifier`

NewVcenterNamespaceManagementStatsTimeSeriesPodIdentifier instantiates a new VcenterNamespaceManagementStatsTimeSeriesPodIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementStatsTimeSeriesPodIdentifierWithDefaults

`func NewVcenterNamespaceManagementStatsTimeSeriesPodIdentifierWithDefaults() *VcenterNamespaceManagementStatsTimeSeriesPodIdentifier`

NewVcenterNamespaceManagementStatsTimeSeriesPodIdentifierWithDefaults instantiates a new VcenterNamespaceManagementStatsTimeSeriesPodIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *VcenterNamespaceManagementStatsTimeSeriesPodIdentifier) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *VcenterNamespaceManagementStatsTimeSeriesPodIdentifier) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *VcenterNamespaceManagementStatsTimeSeriesPodIdentifier) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetPodName

`func (o *VcenterNamespaceManagementStatsTimeSeriesPodIdentifier) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *VcenterNamespaceManagementStatsTimeSeriesPodIdentifier) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *VcenterNamespaceManagementStatsTimeSeriesPodIdentifier) SetPodName(v string)`

SetPodName sets PodName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


