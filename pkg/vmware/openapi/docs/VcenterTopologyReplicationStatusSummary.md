# VcenterTopologyReplicationStatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | **string** | Identifier for the vCenter or Platform Services Controller node. Identifier can be either IP address or DNS resolvable name of the node. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.VCenter.name. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.VCenter.name. | 
**ReplicationPartner** | **string** | Identifier for the vCenter or Platform Services Controller replication partner. Identifier can be either IP address or DNS resolvable name of the replication partner. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.VCenter.name. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.VCenter.name. | 
**PartnerAvailable** | **bool** | Indicates if the VMware Directory Service on partner is reachable or not. | 
**StatusAvailable** | **bool** | Indicates if the replication status for the node with respect to replication partner can be retrieved or not. | 
**Replicating** | Pointer to **bool** | Indicates if node is processing replication changes from the replication partner. This field will be unset if the partner host or replication status is not available, i.e, if ReplicationStatus.Summary.partner-available or ReplicationStatus.Summary.status-available is false. | [optional] 
**ChangeLag** | Pointer to **int64** | Number of replication changes node is behind the replication partner. This field will be unset if the partner host or replication status is not available, i.e, if ReplicationStatus.Summary.partner-available or ReplicationStatus.Summary.status-available is false. | [optional] 

## Methods

### NewVcenterTopologyReplicationStatusSummary

`func NewVcenterTopologyReplicationStatusSummary(node string, replicationPartner string, partnerAvailable bool, statusAvailable bool, ) *VcenterTopologyReplicationStatusSummary`

NewVcenterTopologyReplicationStatusSummary instantiates a new VcenterTopologyReplicationStatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTopologyReplicationStatusSummaryWithDefaults

`func NewVcenterTopologyReplicationStatusSummaryWithDefaults() *VcenterTopologyReplicationStatusSummary`

NewVcenterTopologyReplicationStatusSummaryWithDefaults instantiates a new VcenterTopologyReplicationStatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *VcenterTopologyReplicationStatusSummary) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *VcenterTopologyReplicationStatusSummary) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *VcenterTopologyReplicationStatusSummary) SetNode(v string)`

SetNode sets Node field to given value.


### GetReplicationPartner

`func (o *VcenterTopologyReplicationStatusSummary) GetReplicationPartner() string`

GetReplicationPartner returns the ReplicationPartner field if non-nil, zero value otherwise.

### GetReplicationPartnerOk

`func (o *VcenterTopologyReplicationStatusSummary) GetReplicationPartnerOk() (*string, bool)`

GetReplicationPartnerOk returns a tuple with the ReplicationPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPartner

`func (o *VcenterTopologyReplicationStatusSummary) SetReplicationPartner(v string)`

SetReplicationPartner sets ReplicationPartner field to given value.


### GetPartnerAvailable

`func (o *VcenterTopologyReplicationStatusSummary) GetPartnerAvailable() bool`

GetPartnerAvailable returns the PartnerAvailable field if non-nil, zero value otherwise.

### GetPartnerAvailableOk

`func (o *VcenterTopologyReplicationStatusSummary) GetPartnerAvailableOk() (*bool, bool)`

GetPartnerAvailableOk returns a tuple with the PartnerAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAvailable

`func (o *VcenterTopologyReplicationStatusSummary) SetPartnerAvailable(v bool)`

SetPartnerAvailable sets PartnerAvailable field to given value.


### GetStatusAvailable

`func (o *VcenterTopologyReplicationStatusSummary) GetStatusAvailable() bool`

GetStatusAvailable returns the StatusAvailable field if non-nil, zero value otherwise.

### GetStatusAvailableOk

`func (o *VcenterTopologyReplicationStatusSummary) GetStatusAvailableOk() (*bool, bool)`

GetStatusAvailableOk returns a tuple with the StatusAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAvailable

`func (o *VcenterTopologyReplicationStatusSummary) SetStatusAvailable(v bool)`

SetStatusAvailable sets StatusAvailable field to given value.


### GetReplicating

`func (o *VcenterTopologyReplicationStatusSummary) GetReplicating() bool`

GetReplicating returns the Replicating field if non-nil, zero value otherwise.

### GetReplicatingOk

`func (o *VcenterTopologyReplicationStatusSummary) GetReplicatingOk() (*bool, bool)`

GetReplicatingOk returns a tuple with the Replicating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicating

`func (o *VcenterTopologyReplicationStatusSummary) SetReplicating(v bool)`

SetReplicating sets Replicating field to given value.

### HasReplicating

`func (o *VcenterTopologyReplicationStatusSummary) HasReplicating() bool`

HasReplicating returns a boolean if a field has been set.

### GetChangeLag

`func (o *VcenterTopologyReplicationStatusSummary) GetChangeLag() int64`

GetChangeLag returns the ChangeLag field if non-nil, zero value otherwise.

### GetChangeLagOk

`func (o *VcenterTopologyReplicationStatusSummary) GetChangeLagOk() (*int64, bool)`

GetChangeLagOk returns a tuple with the ChangeLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLag

`func (o *VcenterTopologyReplicationStatusSummary) SetChangeLag(v int64)`

SetChangeLag sets ChangeLag field to given value.

### HasChangeLag

`func (o *VcenterTopologyReplicationStatusSummary) HasChangeLag() bool`

HasChangeLag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


