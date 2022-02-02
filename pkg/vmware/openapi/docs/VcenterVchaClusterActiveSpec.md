# VcenterVchaClusterActiveSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HaNetworkType** | Pointer to [**VcenterVchaNetworkType**](VcenterVchaNetworkType.md) |  | [optional] 
**HaNetwork** | Pointer to **string** | The identifier of the Network object used for the HA network.  If the Cluster.ActiveSpec.ha-network field is set, then the Cluster.ActiveSpec.ha-network-type field must be set.  If the Cluster.ActiveSpec.ha-network field is unset, then the Cluster.ActiveSpec.ha-network-type field is ignored. If unset and the Cluster.ActiveSpec.ha-network-type field is unset, then the second NIC is assumed to be already configured.  If unset and the Cluster.ActiveSpec.ha-network field is set, then an error is reported. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network:VCenter. | [optional] 
**HaIp** | [**VcenterVchaIpSpec**](VcenterVchaIpSpec.md) |  | 

## Methods

### NewVcenterVchaClusterActiveSpec

`func NewVcenterVchaClusterActiveSpec(haIp VcenterVchaIpSpec, ) *VcenterVchaClusterActiveSpec`

NewVcenterVchaClusterActiveSpec instantiates a new VcenterVchaClusterActiveSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterActiveSpecWithDefaults

`func NewVcenterVchaClusterActiveSpecWithDefaults() *VcenterVchaClusterActiveSpec`

NewVcenterVchaClusterActiveSpecWithDefaults instantiates a new VcenterVchaClusterActiveSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHaNetworkType

`func (o *VcenterVchaClusterActiveSpec) GetHaNetworkType() VcenterVchaNetworkType`

GetHaNetworkType returns the HaNetworkType field if non-nil, zero value otherwise.

### GetHaNetworkTypeOk

`func (o *VcenterVchaClusterActiveSpec) GetHaNetworkTypeOk() (*VcenterVchaNetworkType, bool)`

GetHaNetworkTypeOk returns a tuple with the HaNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaNetworkType

`func (o *VcenterVchaClusterActiveSpec) SetHaNetworkType(v VcenterVchaNetworkType)`

SetHaNetworkType sets HaNetworkType field to given value.

### HasHaNetworkType

`func (o *VcenterVchaClusterActiveSpec) HasHaNetworkType() bool`

HasHaNetworkType returns a boolean if a field has been set.

### GetHaNetwork

`func (o *VcenterVchaClusterActiveSpec) GetHaNetwork() string`

GetHaNetwork returns the HaNetwork field if non-nil, zero value otherwise.

### GetHaNetworkOk

`func (o *VcenterVchaClusterActiveSpec) GetHaNetworkOk() (*string, bool)`

GetHaNetworkOk returns a tuple with the HaNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaNetwork

`func (o *VcenterVchaClusterActiveSpec) SetHaNetwork(v string)`

SetHaNetwork sets HaNetwork field to given value.

### HasHaNetwork

`func (o *VcenterVchaClusterActiveSpec) HasHaNetwork() bool`

HasHaNetwork returns a boolean if a field has been set.

### GetHaIp

`func (o *VcenterVchaClusterActiveSpec) GetHaIp() VcenterVchaIpSpec`

GetHaIp returns the HaIp field if non-nil, zero value otherwise.

### GetHaIpOk

`func (o *VcenterVchaClusterActiveSpec) GetHaIpOk() (*VcenterVchaIpSpec, bool)`

GetHaIpOk returns a tuple with the HaIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaIp

`func (o *VcenterVchaClusterActiveSpec) SetHaIp(v VcenterVchaIpSpec)`

SetHaIp sets HaIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


