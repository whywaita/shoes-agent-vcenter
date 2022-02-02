# VcenterVchaClusterPassiveSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Placement** | Pointer to [**VcenterVchaPlacementSpec**](VcenterVchaPlacementSpec.md) |  | [optional] 
**HaIp** | [**VcenterVchaIpSpec**](VcenterVchaIpSpec.md) |  | 
**FailoverIp** | Pointer to [**VcenterVchaIpSpec**](VcenterVchaIpSpec.md) |  | [optional] 

## Methods

### NewVcenterVchaClusterPassiveSpec

`func NewVcenterVchaClusterPassiveSpec(haIp VcenterVchaIpSpec, ) *VcenterVchaClusterPassiveSpec`

NewVcenterVchaClusterPassiveSpec instantiates a new VcenterVchaClusterPassiveSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterPassiveSpecWithDefaults

`func NewVcenterVchaClusterPassiveSpecWithDefaults() *VcenterVchaClusterPassiveSpec`

NewVcenterVchaClusterPassiveSpecWithDefaults instantiates a new VcenterVchaClusterPassiveSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlacement

`func (o *VcenterVchaClusterPassiveSpec) GetPlacement() VcenterVchaPlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVchaClusterPassiveSpec) GetPlacementOk() (*VcenterVchaPlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVchaClusterPassiveSpec) SetPlacement(v VcenterVchaPlacementSpec)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VcenterVchaClusterPassiveSpec) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetHaIp

`func (o *VcenterVchaClusterPassiveSpec) GetHaIp() VcenterVchaIpSpec`

GetHaIp returns the HaIp field if non-nil, zero value otherwise.

### GetHaIpOk

`func (o *VcenterVchaClusterPassiveSpec) GetHaIpOk() (*VcenterVchaIpSpec, bool)`

GetHaIpOk returns a tuple with the HaIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaIp

`func (o *VcenterVchaClusterPassiveSpec) SetHaIp(v VcenterVchaIpSpec)`

SetHaIp sets HaIp field to given value.


### GetFailoverIp

`func (o *VcenterVchaClusterPassiveSpec) GetFailoverIp() VcenterVchaIpSpec`

GetFailoverIp returns the FailoverIp field if non-nil, zero value otherwise.

### GetFailoverIpOk

`func (o *VcenterVchaClusterPassiveSpec) GetFailoverIpOk() (*VcenterVchaIpSpec, bool)`

GetFailoverIpOk returns a tuple with the FailoverIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverIp

`func (o *VcenterVchaClusterPassiveSpec) SetFailoverIp(v VcenterVchaIpSpec)`

SetFailoverIp sets FailoverIp field to given value.

### HasFailoverIp

`func (o *VcenterVchaClusterPassiveSpec) HasFailoverIp() bool`

HasFailoverIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


