# VcenterVmHardwareEthernetBackingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmHardwareEthernetBackingType**](VcenterVmHardwareEthernetBackingType.md) |  | 
**Network** | Pointer to **string** | Identifier of the network that backs the virtual Ethernet adapter. This field is optional and it is only relevant when the value of Ethernet.BackingSpec.type is one of STANDARD_PORTGROUP, DISTRIBUTED_PORTGROUP, or OPAQUE_NETWORK. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network. | [optional] 
**DistributedPort** | Pointer to **string** | Key of the distributed virtual port that backs the virtual Ethernet adapter. Depending on the type of the Portgroup, the port may be specified using this field. If the portgroup type is early-binding (also known as static), a port is assigned when the Ethernet adapter is configured to use the port. The port may be either automatically or specifically assigned based on the value of this field. If the portgroup type is ephemeral, the port is created and assigned to a virtual machine when it is powered on and the Ethernet adapter is connected. This field cannot be specified as no free ports exist before use. May be used to specify a port when the network specified on the Ethernet.BackingSpec.network field is a static or early binding distributed portgroup. If unset, the port will be automatically assigned to the Ethernet adapter based on the policy embodied by the portgroup type. | [optional] 

## Methods

### NewVcenterVmHardwareEthernetBackingSpec

`func NewVcenterVmHardwareEthernetBackingSpec(type_ VcenterVmHardwareEthernetBackingType, ) *VcenterVmHardwareEthernetBackingSpec`

NewVcenterVmHardwareEthernetBackingSpec instantiates a new VcenterVmHardwareEthernetBackingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareEthernetBackingSpecWithDefaults

`func NewVcenterVmHardwareEthernetBackingSpecWithDefaults() *VcenterVmHardwareEthernetBackingSpec`

NewVcenterVmHardwareEthernetBackingSpecWithDefaults instantiates a new VcenterVmHardwareEthernetBackingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareEthernetBackingSpec) GetType() VcenterVmHardwareEthernetBackingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareEthernetBackingSpec) GetTypeOk() (*VcenterVmHardwareEthernetBackingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareEthernetBackingSpec) SetType(v VcenterVmHardwareEthernetBackingType)`

SetType sets Type field to given value.


### GetNetwork

`func (o *VcenterVmHardwareEthernetBackingSpec) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VcenterVmHardwareEthernetBackingSpec) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VcenterVmHardwareEthernetBackingSpec) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VcenterVmHardwareEthernetBackingSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetDistributedPort

`func (o *VcenterVmHardwareEthernetBackingSpec) GetDistributedPort() string`

GetDistributedPort returns the DistributedPort field if non-nil, zero value otherwise.

### GetDistributedPortOk

`func (o *VcenterVmHardwareEthernetBackingSpec) GetDistributedPortOk() (*string, bool)`

GetDistributedPortOk returns a tuple with the DistributedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedPort

`func (o *VcenterVmHardwareEthernetBackingSpec) SetDistributedPort(v string)`

SetDistributedPort sets DistributedPort field to given value.

### HasDistributedPort

`func (o *VcenterVmHardwareEthernetBackingSpec) HasDistributedPort() bool`

HasDistributedPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


