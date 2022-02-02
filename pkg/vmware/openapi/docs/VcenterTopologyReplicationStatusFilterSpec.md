# VcenterTopologyReplicationStatusFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to **[]string** | Identifier that a vCenter and Platform Services Controller node must have to match the filter. (see ReplicationStatus.Summary.node). If unset or empty, all vCenter and Platform Services Controller nodes of type VCSA_EMBEDDED/PSC_EXTERNAL match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.VCenter.name. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.VCenter.name. | [optional] 

## Methods

### NewVcenterTopologyReplicationStatusFilterSpec

`func NewVcenterTopologyReplicationStatusFilterSpec() *VcenterTopologyReplicationStatusFilterSpec`

NewVcenterTopologyReplicationStatusFilterSpec instantiates a new VcenterTopologyReplicationStatusFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTopologyReplicationStatusFilterSpecWithDefaults

`func NewVcenterTopologyReplicationStatusFilterSpecWithDefaults() *VcenterTopologyReplicationStatusFilterSpec`

NewVcenterTopologyReplicationStatusFilterSpecWithDefaults instantiates a new VcenterTopologyReplicationStatusFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *VcenterTopologyReplicationStatusFilterSpec) GetNodes() []string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *VcenterTopologyReplicationStatusFilterSpec) GetNodesOk() (*[]string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *VcenterTopologyReplicationStatusFilterSpec) SetNodes(v []string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *VcenterTopologyReplicationStatusFilterSpec) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


