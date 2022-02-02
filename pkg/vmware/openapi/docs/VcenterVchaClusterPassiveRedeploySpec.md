# VcenterVchaClusterPassiveRedeploySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcSpec** | Pointer to [**VcenterVchaCredentialsSpec**](VcenterVchaCredentialsSpec.md) |  | [optional] 
**Placement** | [**VcenterVchaPlacementSpec**](VcenterVchaPlacementSpec.md) |  | 
**HaIp** | Pointer to [**VcenterVchaIpSpec**](VcenterVchaIpSpec.md) |  | [optional] 
**FailoverIp** | Pointer to [**VcenterVchaIpSpec**](VcenterVchaIpSpec.md) |  | [optional] 

## Methods

### NewVcenterVchaClusterPassiveRedeploySpec

`func NewVcenterVchaClusterPassiveRedeploySpec(placement VcenterVchaPlacementSpec, ) *VcenterVchaClusterPassiveRedeploySpec`

NewVcenterVchaClusterPassiveRedeploySpec instantiates a new VcenterVchaClusterPassiveRedeploySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterPassiveRedeploySpecWithDefaults

`func NewVcenterVchaClusterPassiveRedeploySpecWithDefaults() *VcenterVchaClusterPassiveRedeploySpec`

NewVcenterVchaClusterPassiveRedeploySpecWithDefaults instantiates a new VcenterVchaClusterPassiveRedeploySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcSpec

`func (o *VcenterVchaClusterPassiveRedeploySpec) GetVcSpec() VcenterVchaCredentialsSpec`

GetVcSpec returns the VcSpec field if non-nil, zero value otherwise.

### GetVcSpecOk

`func (o *VcenterVchaClusterPassiveRedeploySpec) GetVcSpecOk() (*VcenterVchaCredentialsSpec, bool)`

GetVcSpecOk returns a tuple with the VcSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcSpec

`func (o *VcenterVchaClusterPassiveRedeploySpec) SetVcSpec(v VcenterVchaCredentialsSpec)`

SetVcSpec sets VcSpec field to given value.

### HasVcSpec

`func (o *VcenterVchaClusterPassiveRedeploySpec) HasVcSpec() bool`

HasVcSpec returns a boolean if a field has been set.

### GetPlacement

`func (o *VcenterVchaClusterPassiveRedeploySpec) GetPlacement() VcenterVchaPlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVchaClusterPassiveRedeploySpec) GetPlacementOk() (*VcenterVchaPlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVchaClusterPassiveRedeploySpec) SetPlacement(v VcenterVchaPlacementSpec)`

SetPlacement sets Placement field to given value.


### GetHaIp

`func (o *VcenterVchaClusterPassiveRedeploySpec) GetHaIp() VcenterVchaIpSpec`

GetHaIp returns the HaIp field if non-nil, zero value otherwise.

### GetHaIpOk

`func (o *VcenterVchaClusterPassiveRedeploySpec) GetHaIpOk() (*VcenterVchaIpSpec, bool)`

GetHaIpOk returns a tuple with the HaIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaIp

`func (o *VcenterVchaClusterPassiveRedeploySpec) SetHaIp(v VcenterVchaIpSpec)`

SetHaIp sets HaIp field to given value.

### HasHaIp

`func (o *VcenterVchaClusterPassiveRedeploySpec) HasHaIp() bool`

HasHaIp returns a boolean if a field has been set.

### GetFailoverIp

`func (o *VcenterVchaClusterPassiveRedeploySpec) GetFailoverIp() VcenterVchaIpSpec`

GetFailoverIp returns the FailoverIp field if non-nil, zero value otherwise.

### GetFailoverIpOk

`func (o *VcenterVchaClusterPassiveRedeploySpec) GetFailoverIpOk() (*VcenterVchaIpSpec, bool)`

GetFailoverIpOk returns a tuple with the FailoverIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverIp

`func (o *VcenterVchaClusterPassiveRedeploySpec) SetFailoverIp(v VcenterVchaIpSpec)`

SetFailoverIp sets FailoverIp field to given value.

### HasFailoverIp

`func (o *VcenterVchaClusterPassiveRedeploySpec) HasFailoverIp() bool`

HasFailoverIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


