# VcenterTopologyNodesFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Types** | Pointer to [**[]VcenterTopologyNodesApplianceType**](VcenterTopologyNodesApplianceType.md) | Types of the appliance that a vCenter and Platform Services Controller node must be to match the filter (see Nodes.ApplianceType. If unset or empty, node of any ApplianceType match the filter. | [optional] 

## Methods

### NewVcenterTopologyNodesFilterSpec

`func NewVcenterTopologyNodesFilterSpec() *VcenterTopologyNodesFilterSpec`

NewVcenterTopologyNodesFilterSpec instantiates a new VcenterTopologyNodesFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTopologyNodesFilterSpecWithDefaults

`func NewVcenterTopologyNodesFilterSpecWithDefaults() *VcenterTopologyNodesFilterSpec`

NewVcenterTopologyNodesFilterSpecWithDefaults instantiates a new VcenterTopologyNodesFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypes

`func (o *VcenterTopologyNodesFilterSpec) GetTypes() []VcenterTopologyNodesApplianceType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *VcenterTopologyNodesFilterSpec) GetTypesOk() (*[]VcenterTopologyNodesApplianceType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *VcenterTopologyNodesFilterSpec) SetTypes(v []VcenterTopologyNodesApplianceType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *VcenterTopologyNodesFilterSpec) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


