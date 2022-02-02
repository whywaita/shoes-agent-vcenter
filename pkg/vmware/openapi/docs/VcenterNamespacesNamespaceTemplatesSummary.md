# VcenterNamespacesNamespaceTemplatesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** | Identifier for the cluster associated with namespace template. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | 
**Template** | **string** | Name of the namespace template. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespaces.NamespaceTemplate. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespaces.NamespaceTemplate. | 

## Methods

### NewVcenterNamespacesNamespaceTemplatesSummary

`func NewVcenterNamespacesNamespaceTemplatesSummary(cluster string, template string, ) *VcenterNamespacesNamespaceTemplatesSummary`

NewVcenterNamespacesNamespaceTemplatesSummary instantiates a new VcenterNamespacesNamespaceTemplatesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesNamespaceTemplatesSummaryWithDefaults

`func NewVcenterNamespacesNamespaceTemplatesSummaryWithDefaults() *VcenterNamespacesNamespaceTemplatesSummary`

NewVcenterNamespacesNamespaceTemplatesSummaryWithDefaults instantiates a new VcenterNamespacesNamespaceTemplatesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VcenterNamespacesNamespaceTemplatesSummary) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterNamespacesNamespaceTemplatesSummary) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterNamespacesNamespaceTemplatesSummary) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetTemplate

`func (o *VcenterNamespacesNamespaceTemplatesSummary) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *VcenterNamespacesNamespaceTemplatesSummary) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *VcenterNamespacesNamespaceTemplatesSummary) SetTemplate(v string)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


