# VcenterVchaClusterPassiveCheckSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcSpec** | Pointer to [**VcenterVchaCredentialsSpec**](VcenterVchaCredentialsSpec.md) |  | [optional] 
**Placement** | [**VcenterVchaPlacementSpec**](VcenterVchaPlacementSpec.md) |  | 

## Methods

### NewVcenterVchaClusterPassiveCheckSpec

`func NewVcenterVchaClusterPassiveCheckSpec(placement VcenterVchaPlacementSpec, ) *VcenterVchaClusterPassiveCheckSpec`

NewVcenterVchaClusterPassiveCheckSpec instantiates a new VcenterVchaClusterPassiveCheckSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterPassiveCheckSpecWithDefaults

`func NewVcenterVchaClusterPassiveCheckSpecWithDefaults() *VcenterVchaClusterPassiveCheckSpec`

NewVcenterVchaClusterPassiveCheckSpecWithDefaults instantiates a new VcenterVchaClusterPassiveCheckSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcSpec

`func (o *VcenterVchaClusterPassiveCheckSpec) GetVcSpec() VcenterVchaCredentialsSpec`

GetVcSpec returns the VcSpec field if non-nil, zero value otherwise.

### GetVcSpecOk

`func (o *VcenterVchaClusterPassiveCheckSpec) GetVcSpecOk() (*VcenterVchaCredentialsSpec, bool)`

GetVcSpecOk returns a tuple with the VcSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcSpec

`func (o *VcenterVchaClusterPassiveCheckSpec) SetVcSpec(v VcenterVchaCredentialsSpec)`

SetVcSpec sets VcSpec field to given value.

### HasVcSpec

`func (o *VcenterVchaClusterPassiveCheckSpec) HasVcSpec() bool`

HasVcSpec returns a boolean if a field has been set.

### GetPlacement

`func (o *VcenterVchaClusterPassiveCheckSpec) GetPlacement() VcenterVchaPlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVchaClusterPassiveCheckSpec) GetPlacementOk() (*VcenterVchaPlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVchaClusterPassiveCheckSpec) SetPlacement(v VcenterVchaPlacementSpec)`

SetPlacement sets Placement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


