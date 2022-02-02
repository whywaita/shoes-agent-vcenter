# VcenterNamespacesUserInstancesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** | Identifier of the namespace. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespaces.Instance. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespaces.Instance. | 
**MasterHost** | **string** | IP address or FQDN of the API endpoint for the given namespace. | 

## Methods

### NewVcenterNamespacesUserInstancesSummary

`func NewVcenterNamespacesUserInstancesSummary(namespace string, masterHost string, ) *VcenterNamespacesUserInstancesSummary`

NewVcenterNamespacesUserInstancesSummary instantiates a new VcenterNamespacesUserInstancesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesUserInstancesSummaryWithDefaults

`func NewVcenterNamespacesUserInstancesSummaryWithDefaults() *VcenterNamespacesUserInstancesSummary`

NewVcenterNamespacesUserInstancesSummaryWithDefaults instantiates a new VcenterNamespacesUserInstancesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *VcenterNamespacesUserInstancesSummary) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *VcenterNamespacesUserInstancesSummary) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *VcenterNamespacesUserInstancesSummary) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetMasterHost

`func (o *VcenterNamespacesUserInstancesSummary) GetMasterHost() string`

GetMasterHost returns the MasterHost field if non-nil, zero value otherwise.

### GetMasterHostOk

`func (o *VcenterNamespacesUserInstancesSummary) GetMasterHostOk() (*string, bool)`

GetMasterHostOk returns a tuple with the MasterHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterHost

`func (o *VcenterNamespacesUserInstancesSummary) SetMasterHost(v string)`

SetMasterHost sets MasterHost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


