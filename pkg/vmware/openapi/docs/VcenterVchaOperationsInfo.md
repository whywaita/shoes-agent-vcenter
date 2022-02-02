# VcenterVchaOperationsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | **[]string** | Identifiers of the operations that are current disabled. These operation strings are one of \&quot;vcenter.vcha.cluster.deploy\&quot;, \&quot;vcenter.vcha.cluster.failover\&quot;, \&quot;vcenter.vcha.cluster.passive.redeploy\&quot;, \&quot;vcenter.vcha.cluster.witness.redeploy\&quot;, \&quot;vcenter.vcha.cluster.mode.set\&quot;, \&quot;vcenter.vcha.cluster.undeploy\&quot; and \&quot;vcenter.vcha.cluster.get\&quot;. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vapi.operation. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vapi.operation. | 
**Active** | **[]string** | Identifiers of the operations that are currently running. These operation strings are one of \&quot;vcenter.vcha.cluster.deploy\&quot;, \&quot;vcenter.vcha.cluster.failover\&quot;, \&quot;vcenter.vcha.cluster.passive.redeploy\&quot;, \&quot;vcenter.vcha.cluster.witness.redeploy\&quot;, \&quot;vcenter.vcha.cluster.mode.set\&quot;, and \&quot;vcenter.vcha.cluster.undeploy\&quot;. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vapi.operation. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vapi.operation. | 

## Methods

### NewVcenterVchaOperationsInfo

`func NewVcenterVchaOperationsInfo(disabled []string, active []string, ) *VcenterVchaOperationsInfo`

NewVcenterVchaOperationsInfo instantiates a new VcenterVchaOperationsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaOperationsInfoWithDefaults

`func NewVcenterVchaOperationsInfoWithDefaults() *VcenterVchaOperationsInfo`

NewVcenterVchaOperationsInfoWithDefaults instantiates a new VcenterVchaOperationsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *VcenterVchaOperationsInfo) GetDisabled() []string`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *VcenterVchaOperationsInfo) GetDisabledOk() (*[]string, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *VcenterVchaOperationsInfo) SetDisabled(v []string)`

SetDisabled sets Disabled field to given value.


### GetActive

`func (o *VcenterVchaOperationsInfo) GetActive() []string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *VcenterVchaOperationsInfo) GetActiveOk() (*[]string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *VcenterVchaOperationsInfo) SetActive(v []string)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


