# VcenterNamespacesNamespaceSelfServiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** | Identifier for the cluster to which namespace service is associated. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | 
**Capability** | [**VcenterNamespacesNamespaceSelfServiceCapability**](VcenterNamespacesNamespaceSelfServiceCapability.md) |  | 
**Status** | [**VcenterNamespacesNamespaceSelfServiceStatus**](VcenterNamespacesNamespaceSelfServiceStatus.md) |  | 
**Messages** | [**[]VcenterNamespacesInstancesMessage**](VcenterNamespacesInstancesMessage.md) | Current set of messages associated with the object. | 

## Methods

### NewVcenterNamespacesNamespaceSelfServiceInfo

`func NewVcenterNamespacesNamespaceSelfServiceInfo(cluster string, capability VcenterNamespacesNamespaceSelfServiceCapability, status VcenterNamespacesNamespaceSelfServiceStatus, messages []VcenterNamespacesInstancesMessage, ) *VcenterNamespacesNamespaceSelfServiceInfo`

NewVcenterNamespacesNamespaceSelfServiceInfo instantiates a new VcenterNamespacesNamespaceSelfServiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesNamespaceSelfServiceInfoWithDefaults

`func NewVcenterNamespacesNamespaceSelfServiceInfoWithDefaults() *VcenterNamespacesNamespaceSelfServiceInfo`

NewVcenterNamespacesNamespaceSelfServiceInfoWithDefaults instantiates a new VcenterNamespacesNamespaceSelfServiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VcenterNamespacesNamespaceSelfServiceInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterNamespacesNamespaceSelfServiceInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterNamespacesNamespaceSelfServiceInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetCapability

`func (o *VcenterNamespacesNamespaceSelfServiceInfo) GetCapability() VcenterNamespacesNamespaceSelfServiceCapability`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *VcenterNamespacesNamespaceSelfServiceInfo) GetCapabilityOk() (*VcenterNamespacesNamespaceSelfServiceCapability, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *VcenterNamespacesNamespaceSelfServiceInfo) SetCapability(v VcenterNamespacesNamespaceSelfServiceCapability)`

SetCapability sets Capability field to given value.


### GetStatus

`func (o *VcenterNamespacesNamespaceSelfServiceInfo) GetStatus() VcenterNamespacesNamespaceSelfServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VcenterNamespacesNamespaceSelfServiceInfo) GetStatusOk() (*VcenterNamespacesNamespaceSelfServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VcenterNamespacesNamespaceSelfServiceInfo) SetStatus(v VcenterNamespacesNamespaceSelfServiceStatus)`

SetStatus sets Status field to given value.


### GetMessages

`func (o *VcenterNamespacesNamespaceSelfServiceInfo) GetMessages() []VcenterNamespacesInstancesMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *VcenterNamespacesNamespaceSelfServiceInfo) GetMessagesOk() (*[]VcenterNamespacesInstancesMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *VcenterNamespacesNamespaceSelfServiceInfo) SetMessages(v []VcenterNamespacesInstancesMessage)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


