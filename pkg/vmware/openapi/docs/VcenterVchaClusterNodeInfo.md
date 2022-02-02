# VcenterVchaClusterNodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailoverIp** | Pointer to [**VcenterVchaClusterIpInfo**](VcenterVchaClusterIpInfo.md) |  | [optional] 
**HaIp** | [**VcenterVchaClusterIpInfo**](VcenterVchaClusterIpInfo.md) |  | 
**Runtime** | Pointer to [**VcenterVchaClusterNodeRuntimeInfo**](VcenterVchaClusterNodeRuntimeInfo.md) |  | [optional] 

## Methods

### NewVcenterVchaClusterNodeInfo

`func NewVcenterVchaClusterNodeInfo(haIp VcenterVchaClusterIpInfo, ) *VcenterVchaClusterNodeInfo`

NewVcenterVchaClusterNodeInfo instantiates a new VcenterVchaClusterNodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterNodeInfoWithDefaults

`func NewVcenterVchaClusterNodeInfoWithDefaults() *VcenterVchaClusterNodeInfo`

NewVcenterVchaClusterNodeInfoWithDefaults instantiates a new VcenterVchaClusterNodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailoverIp

`func (o *VcenterVchaClusterNodeInfo) GetFailoverIp() VcenterVchaClusterIpInfo`

GetFailoverIp returns the FailoverIp field if non-nil, zero value otherwise.

### GetFailoverIpOk

`func (o *VcenterVchaClusterNodeInfo) GetFailoverIpOk() (*VcenterVchaClusterIpInfo, bool)`

GetFailoverIpOk returns a tuple with the FailoverIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverIp

`func (o *VcenterVchaClusterNodeInfo) SetFailoverIp(v VcenterVchaClusterIpInfo)`

SetFailoverIp sets FailoverIp field to given value.

### HasFailoverIp

`func (o *VcenterVchaClusterNodeInfo) HasFailoverIp() bool`

HasFailoverIp returns a boolean if a field has been set.

### GetHaIp

`func (o *VcenterVchaClusterNodeInfo) GetHaIp() VcenterVchaClusterIpInfo`

GetHaIp returns the HaIp field if non-nil, zero value otherwise.

### GetHaIpOk

`func (o *VcenterVchaClusterNodeInfo) GetHaIpOk() (*VcenterVchaClusterIpInfo, bool)`

GetHaIpOk returns a tuple with the HaIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaIp

`func (o *VcenterVchaClusterNodeInfo) SetHaIp(v VcenterVchaClusterIpInfo)`

SetHaIp sets HaIp field to given value.


### GetRuntime

`func (o *VcenterVchaClusterNodeInfo) GetRuntime() VcenterVchaClusterNodeRuntimeInfo`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *VcenterVchaClusterNodeInfo) GetRuntimeOk() (*VcenterVchaClusterNodeRuntimeInfo, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *VcenterVchaClusterNodeInfo) SetRuntime(v VcenterVchaClusterNodeRuntimeInfo)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *VcenterVchaClusterNodeInfo) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


