# VcenterTopologyNodesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | **string** | Identifier for the vCenter or Platform Services Controller node. Identifier can be either IP address or DNS resolvable name of the node. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.VCenter.name. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.VCenter.name. | 
**Type** | [**VcenterTopologyNodesApplianceType**](VcenterTopologyNodesApplianceType.md) |  | 
**ReplicationPartners** | Pointer to **[]string** | List of replication partners&#39; node identifiers. Identifiers can be either IP address or DNS resolvable name of the partner node. This field is optional and it is only relevant when the value of Nodes.Summary.type is one of VCSA_EMBEDDED or PSC_EXTERNAL. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.VCenter.name. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.VCenter.name. | [optional] 
**ClientAffinity** | Pointer to **string** | Identifier of the affinitized Platform Services Controller node. Identifier can be either IP address or DNS resolvable name of the affinitized node. This field is optional and it is only relevant when the value of Nodes.Summary.type is VCSA_EXTERNAL. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.VCenter.name. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.VCenter.name. | [optional] 

## Methods

### NewVcenterTopologyNodesSummary

`func NewVcenterTopologyNodesSummary(node string, type_ VcenterTopologyNodesApplianceType, ) *VcenterTopologyNodesSummary`

NewVcenterTopologyNodesSummary instantiates a new VcenterTopologyNodesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTopologyNodesSummaryWithDefaults

`func NewVcenterTopologyNodesSummaryWithDefaults() *VcenterTopologyNodesSummary`

NewVcenterTopologyNodesSummaryWithDefaults instantiates a new VcenterTopologyNodesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *VcenterTopologyNodesSummary) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *VcenterTopologyNodesSummary) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *VcenterTopologyNodesSummary) SetNode(v string)`

SetNode sets Node field to given value.


### GetType

`func (o *VcenterTopologyNodesSummary) GetType() VcenterTopologyNodesApplianceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterTopologyNodesSummary) GetTypeOk() (*VcenterTopologyNodesApplianceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterTopologyNodesSummary) SetType(v VcenterTopologyNodesApplianceType)`

SetType sets Type field to given value.


### GetReplicationPartners

`func (o *VcenterTopologyNodesSummary) GetReplicationPartners() []string`

GetReplicationPartners returns the ReplicationPartners field if non-nil, zero value otherwise.

### GetReplicationPartnersOk

`func (o *VcenterTopologyNodesSummary) GetReplicationPartnersOk() (*[]string, bool)`

GetReplicationPartnersOk returns a tuple with the ReplicationPartners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPartners

`func (o *VcenterTopologyNodesSummary) SetReplicationPartners(v []string)`

SetReplicationPartners sets ReplicationPartners field to given value.

### HasReplicationPartners

`func (o *VcenterTopologyNodesSummary) HasReplicationPartners() bool`

HasReplicationPartners returns a boolean if a field has been set.

### GetClientAffinity

`func (o *VcenterTopologyNodesSummary) GetClientAffinity() string`

GetClientAffinity returns the ClientAffinity field if non-nil, zero value otherwise.

### GetClientAffinityOk

`func (o *VcenterTopologyNodesSummary) GetClientAffinityOk() (*string, bool)`

GetClientAffinityOk returns a tuple with the ClientAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAffinity

`func (o *VcenterTopologyNodesSummary) SetClientAffinity(v string)`

SetClientAffinity sets ClientAffinity field to given value.

### HasClientAffinity

`func (o *VcenterTopologyNodesSummary) HasClientAffinity() bool`

HasClientAffinity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


