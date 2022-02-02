# VcenterVchaClusterUndeploySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcSpec** | Pointer to [**VcenterVchaCredentialsSpec**](VcenterVchaCredentialsSpec.md) |  | [optional] 
**ForceDelete** | Pointer to **bool** | Flag controlling in what circumstances the virtual machines will be deleted. For this flag to take effect, the VCHA cluster should have been successfully configured using automatic deployment.     -  If true, the Cluster.UndeploySpec.vms field will be ignored, the VCHA cluster specific information is removed, and the passive and witness virtual machines will be deleted.    -  If false, the Cluster.UndeploySpec.vms field contains the information identifying the passive and witness virtual machines.        &#x3D;  If the Cluster.UndeploySpec.vms field is set, then it will be validated prior to deleting the passive and witness virtual machines and VCHA cluster specific information is removed.      &#x3D;  If the Cluster.UndeploySpec.vms field is unset, then the passive and witness virtual machines will not be deleted. The customer should delete them in order to cleanup completely. VCHA cluster specific information is removed.    If unset, the Cluster.UndeploySpec.vms field contains the information identifying the passive and witness virtual machines.     -  If the Cluster.UndeploySpec.vms field is set, then it will be validated prior to deleting the passive and witness virtual machines. VCHA cluster specific information is removed.    -  If the Cluster.UndeploySpec.vms field is unset, then the passive and witness virtual machines will not be deleted. The customer should delete them in order to cleanup completely. VCHA cluster specific information is removed.  | [optional] 
**Vms** | Pointer to [**VcenterVchaClusterVmInfo**](VcenterVchaClusterVmInfo.md) |  | [optional] 

## Methods

### NewVcenterVchaClusterUndeploySpec

`func NewVcenterVchaClusterUndeploySpec() *VcenterVchaClusterUndeploySpec`

NewVcenterVchaClusterUndeploySpec instantiates a new VcenterVchaClusterUndeploySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterUndeploySpecWithDefaults

`func NewVcenterVchaClusterUndeploySpecWithDefaults() *VcenterVchaClusterUndeploySpec`

NewVcenterVchaClusterUndeploySpecWithDefaults instantiates a new VcenterVchaClusterUndeploySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcSpec

`func (o *VcenterVchaClusterUndeploySpec) GetVcSpec() VcenterVchaCredentialsSpec`

GetVcSpec returns the VcSpec field if non-nil, zero value otherwise.

### GetVcSpecOk

`func (o *VcenterVchaClusterUndeploySpec) GetVcSpecOk() (*VcenterVchaCredentialsSpec, bool)`

GetVcSpecOk returns a tuple with the VcSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcSpec

`func (o *VcenterVchaClusterUndeploySpec) SetVcSpec(v VcenterVchaCredentialsSpec)`

SetVcSpec sets VcSpec field to given value.

### HasVcSpec

`func (o *VcenterVchaClusterUndeploySpec) HasVcSpec() bool`

HasVcSpec returns a boolean if a field has been set.

### GetForceDelete

`func (o *VcenterVchaClusterUndeploySpec) GetForceDelete() bool`

GetForceDelete returns the ForceDelete field if non-nil, zero value otherwise.

### GetForceDeleteOk

`func (o *VcenterVchaClusterUndeploySpec) GetForceDeleteOk() (*bool, bool)`

GetForceDeleteOk returns a tuple with the ForceDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDelete

`func (o *VcenterVchaClusterUndeploySpec) SetForceDelete(v bool)`

SetForceDelete sets ForceDelete field to given value.

### HasForceDelete

`func (o *VcenterVchaClusterUndeploySpec) HasForceDelete() bool`

HasForceDelete returns a boolean if a field has been set.

### GetVms

`func (o *VcenterVchaClusterUndeploySpec) GetVms() VcenterVchaClusterVmInfo`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *VcenterVchaClusterUndeploySpec) GetVmsOk() (*VcenterVchaClusterVmInfo, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *VcenterVchaClusterUndeploySpec) SetVms(v VcenterVchaClusterVmInfo)`

SetVms sets Vms field to given value.

### HasVms

`func (o *VcenterVchaClusterUndeploySpec) HasVms() bool`

HasVms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


