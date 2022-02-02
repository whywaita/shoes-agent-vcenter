# VcenterVchaClusterWitnessRedeploySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcSpec** | Pointer to [**VcenterVchaCredentialsSpec**](VcenterVchaCredentialsSpec.md) |  | [optional] 
**Placement** | [**VcenterVchaPlacementSpec**](VcenterVchaPlacementSpec.md) |  | 
**HaIp** | Pointer to [**VcenterVchaIpSpec**](VcenterVchaIpSpec.md) |  | [optional] 

## Methods

### NewVcenterVchaClusterWitnessRedeploySpec

`func NewVcenterVchaClusterWitnessRedeploySpec(placement VcenterVchaPlacementSpec, ) *VcenterVchaClusterWitnessRedeploySpec`

NewVcenterVchaClusterWitnessRedeploySpec instantiates a new VcenterVchaClusterWitnessRedeploySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterWitnessRedeploySpecWithDefaults

`func NewVcenterVchaClusterWitnessRedeploySpecWithDefaults() *VcenterVchaClusterWitnessRedeploySpec`

NewVcenterVchaClusterWitnessRedeploySpecWithDefaults instantiates a new VcenterVchaClusterWitnessRedeploySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcSpec

`func (o *VcenterVchaClusterWitnessRedeploySpec) GetVcSpec() VcenterVchaCredentialsSpec`

GetVcSpec returns the VcSpec field if non-nil, zero value otherwise.

### GetVcSpecOk

`func (o *VcenterVchaClusterWitnessRedeploySpec) GetVcSpecOk() (*VcenterVchaCredentialsSpec, bool)`

GetVcSpecOk returns a tuple with the VcSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcSpec

`func (o *VcenterVchaClusterWitnessRedeploySpec) SetVcSpec(v VcenterVchaCredentialsSpec)`

SetVcSpec sets VcSpec field to given value.

### HasVcSpec

`func (o *VcenterVchaClusterWitnessRedeploySpec) HasVcSpec() bool`

HasVcSpec returns a boolean if a field has been set.

### GetPlacement

`func (o *VcenterVchaClusterWitnessRedeploySpec) GetPlacement() VcenterVchaPlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVchaClusterWitnessRedeploySpec) GetPlacementOk() (*VcenterVchaPlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVchaClusterWitnessRedeploySpec) SetPlacement(v VcenterVchaPlacementSpec)`

SetPlacement sets Placement field to given value.


### GetHaIp

`func (o *VcenterVchaClusterWitnessRedeploySpec) GetHaIp() VcenterVchaIpSpec`

GetHaIp returns the HaIp field if non-nil, zero value otherwise.

### GetHaIpOk

`func (o *VcenterVchaClusterWitnessRedeploySpec) GetHaIpOk() (*VcenterVchaIpSpec, bool)`

GetHaIpOk returns a tuple with the HaIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaIp

`func (o *VcenterVchaClusterWitnessRedeploySpec) SetHaIp(v VcenterVchaIpSpec)`

SetHaIp sets HaIp field to given value.

### HasHaIp

`func (o *VcenterVchaClusterWitnessRedeploySpec) HasHaIp() bool`

HasHaIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


