# VcenterVmStoragePolicyComplianceVmComplianceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**VcenterVmStoragePolicyComplianceStatus**](VcenterVmStoragePolicyComplianceStatus.md) |  | 
**CheckTime** | **time.Time** | Date and time of the most recent compliance check. | 
**Policy** | Pointer to **string** | Identifier of the storage policy associated with the virtual machine. If unset SPBM is unable to retrieve or determine the associated policy, Compliance.VmComplianceInfo.failure-cause is set in such casses. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.StoragePolicy. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.StoragePolicy. | [optional] 
**FailureCause** | [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | The exception that caused the compliance check to fail. There can be more than one cause, since a policy can contain capabilities from multiple providers. If empty, it implies no failures while retrieving compliance. | 

## Methods

### NewVcenterVmStoragePolicyComplianceVmComplianceInfo

`func NewVcenterVmStoragePolicyComplianceVmComplianceInfo(status VcenterVmStoragePolicyComplianceStatus, checkTime time.Time, failureCause []VapiStdLocalizableMessage, ) *VcenterVmStoragePolicyComplianceVmComplianceInfo`

NewVcenterVmStoragePolicyComplianceVmComplianceInfo instantiates a new VcenterVmStoragePolicyComplianceVmComplianceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmStoragePolicyComplianceVmComplianceInfoWithDefaults

`func NewVcenterVmStoragePolicyComplianceVmComplianceInfoWithDefaults() *VcenterVmStoragePolicyComplianceVmComplianceInfo`

NewVcenterVmStoragePolicyComplianceVmComplianceInfoWithDefaults instantiates a new VcenterVmStoragePolicyComplianceVmComplianceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VcenterVmStoragePolicyComplianceVmComplianceInfo) GetStatus() VcenterVmStoragePolicyComplianceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VcenterVmStoragePolicyComplianceVmComplianceInfo) GetStatusOk() (*VcenterVmStoragePolicyComplianceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VcenterVmStoragePolicyComplianceVmComplianceInfo) SetStatus(v VcenterVmStoragePolicyComplianceStatus)`

SetStatus sets Status field to given value.


### GetCheckTime

`func (o *VcenterVmStoragePolicyComplianceVmComplianceInfo) GetCheckTime() time.Time`

GetCheckTime returns the CheckTime field if non-nil, zero value otherwise.

### GetCheckTimeOk

`func (o *VcenterVmStoragePolicyComplianceVmComplianceInfo) GetCheckTimeOk() (*time.Time, bool)`

GetCheckTimeOk returns a tuple with the CheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTime

`func (o *VcenterVmStoragePolicyComplianceVmComplianceInfo) SetCheckTime(v time.Time)`

SetCheckTime sets CheckTime field to given value.


### GetPolicy

`func (o *VcenterVmStoragePolicyComplianceVmComplianceInfo) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *VcenterVmStoragePolicyComplianceVmComplianceInfo) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *VcenterVmStoragePolicyComplianceVmComplianceInfo) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *VcenterVmStoragePolicyComplianceVmComplianceInfo) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetFailureCause

`func (o *VcenterVmStoragePolicyComplianceVmComplianceInfo) GetFailureCause() []VapiStdLocalizableMessage`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *VcenterVmStoragePolicyComplianceVmComplianceInfo) GetFailureCauseOk() (*[]VapiStdLocalizableMessage, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *VcenterVmStoragePolicyComplianceVmComplianceInfo) SetFailureCause(v []VapiStdLocalizableMessage)`

SetFailureCause sets FailureCause field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


