# VcenterVchaClusterWitnessCheckSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcSpec** | Pointer to [**VcenterVchaCredentialsSpec**](VcenterVchaCredentialsSpec.md) |  | [optional] 
**Placement** | [**VcenterVchaPlacementSpec**](VcenterVchaPlacementSpec.md) |  | 

## Methods

### NewVcenterVchaClusterWitnessCheckSpec

`func NewVcenterVchaClusterWitnessCheckSpec(placement VcenterVchaPlacementSpec, ) *VcenterVchaClusterWitnessCheckSpec`

NewVcenterVchaClusterWitnessCheckSpec instantiates a new VcenterVchaClusterWitnessCheckSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterWitnessCheckSpecWithDefaults

`func NewVcenterVchaClusterWitnessCheckSpecWithDefaults() *VcenterVchaClusterWitnessCheckSpec`

NewVcenterVchaClusterWitnessCheckSpecWithDefaults instantiates a new VcenterVchaClusterWitnessCheckSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcSpec

`func (o *VcenterVchaClusterWitnessCheckSpec) GetVcSpec() VcenterVchaCredentialsSpec`

GetVcSpec returns the VcSpec field if non-nil, zero value otherwise.

### GetVcSpecOk

`func (o *VcenterVchaClusterWitnessCheckSpec) GetVcSpecOk() (*VcenterVchaCredentialsSpec, bool)`

GetVcSpecOk returns a tuple with the VcSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcSpec

`func (o *VcenterVchaClusterWitnessCheckSpec) SetVcSpec(v VcenterVchaCredentialsSpec)`

SetVcSpec sets VcSpec field to given value.

### HasVcSpec

`func (o *VcenterVchaClusterWitnessCheckSpec) HasVcSpec() bool`

HasVcSpec returns a boolean if a field has been set.

### GetPlacement

`func (o *VcenterVchaClusterWitnessCheckSpec) GetPlacement() VcenterVchaPlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVchaClusterWitnessCheckSpec) GetPlacementOk() (*VcenterVchaPlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVchaClusterWitnessCheckSpec) SetPlacement(v VcenterVchaPlacementSpec)`

SetPlacement sets Placement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


