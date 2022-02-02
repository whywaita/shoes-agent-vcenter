# VcenterVchaClusterWitnessInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HaIp** | [**VcenterVchaClusterIpInfo**](VcenterVchaClusterIpInfo.md) |  | 
**Runtime** | Pointer to [**VcenterVchaClusterNodeRuntimeInfo**](VcenterVchaClusterNodeRuntimeInfo.md) |  | [optional] 

## Methods

### NewVcenterVchaClusterWitnessInfo

`func NewVcenterVchaClusterWitnessInfo(haIp VcenterVchaClusterIpInfo, ) *VcenterVchaClusterWitnessInfo`

NewVcenterVchaClusterWitnessInfo instantiates a new VcenterVchaClusterWitnessInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterWitnessInfoWithDefaults

`func NewVcenterVchaClusterWitnessInfoWithDefaults() *VcenterVchaClusterWitnessInfo`

NewVcenterVchaClusterWitnessInfoWithDefaults instantiates a new VcenterVchaClusterWitnessInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHaIp

`func (o *VcenterVchaClusterWitnessInfo) GetHaIp() VcenterVchaClusterIpInfo`

GetHaIp returns the HaIp field if non-nil, zero value otherwise.

### GetHaIpOk

`func (o *VcenterVchaClusterWitnessInfo) GetHaIpOk() (*VcenterVchaClusterIpInfo, bool)`

GetHaIpOk returns a tuple with the HaIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaIp

`func (o *VcenterVchaClusterWitnessInfo) SetHaIp(v VcenterVchaClusterIpInfo)`

SetHaIp sets HaIp field to given value.


### GetRuntime

`func (o *VcenterVchaClusterWitnessInfo) GetRuntime() VcenterVchaClusterNodeRuntimeInfo`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *VcenterVchaClusterWitnessInfo) GetRuntimeOk() (*VcenterVchaClusterNodeRuntimeInfo, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *VcenterVchaClusterWitnessInfo) SetRuntime(v VcenterVchaClusterNodeRuntimeInfo)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *VcenterVchaClusterWitnessInfo) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


