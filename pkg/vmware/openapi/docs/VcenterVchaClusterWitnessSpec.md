# VcenterVchaClusterWitnessSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Placement** | Pointer to [**VcenterVchaPlacementSpec**](VcenterVchaPlacementSpec.md) |  | [optional] 
**HaIp** | [**VcenterVchaIpSpec**](VcenterVchaIpSpec.md) |  | 

## Methods

### NewVcenterVchaClusterWitnessSpec

`func NewVcenterVchaClusterWitnessSpec(haIp VcenterVchaIpSpec, ) *VcenterVchaClusterWitnessSpec`

NewVcenterVchaClusterWitnessSpec instantiates a new VcenterVchaClusterWitnessSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterWitnessSpecWithDefaults

`func NewVcenterVchaClusterWitnessSpecWithDefaults() *VcenterVchaClusterWitnessSpec`

NewVcenterVchaClusterWitnessSpecWithDefaults instantiates a new VcenterVchaClusterWitnessSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlacement

`func (o *VcenterVchaClusterWitnessSpec) GetPlacement() VcenterVchaPlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVchaClusterWitnessSpec) GetPlacementOk() (*VcenterVchaPlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVchaClusterWitnessSpec) SetPlacement(v VcenterVchaPlacementSpec)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VcenterVchaClusterWitnessSpec) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetHaIp

`func (o *VcenterVchaClusterWitnessSpec) GetHaIp() VcenterVchaIpSpec`

GetHaIp returns the HaIp field if non-nil, zero value otherwise.

### GetHaIpOk

`func (o *VcenterVchaClusterWitnessSpec) GetHaIpOk() (*VcenterVchaIpSpec, bool)`

GetHaIpOk returns a tuple with the HaIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaIp

`func (o *VcenterVchaClusterWitnessSpec) SetHaIp(v VcenterVchaIpSpec)`

SetHaIp sets HaIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


