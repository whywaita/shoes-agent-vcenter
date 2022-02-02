# VcenterVmStoragePolicyComplianceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverallCompliance** | [**VcenterVmStoragePolicyComplianceStatus**](VcenterVmStoragePolicyComplianceStatus.md) |  | 
**VmHome** | Pointer to [**VcenterVmStoragePolicyComplianceVmComplianceInfo**](VcenterVmStoragePolicyComplianceVmComplianceInfo.md) |  | [optional] 
**Disks** | [**[]VcenterVmStoragePolicyComplianceInfoDisks**](VcenterVmStoragePolicyComplianceInfoDisks.md) | The compliance information Compliance.VmComplianceInfo for the virtual machine&#39;s virtual disks that are currently associated with a storage policy. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Disk. | 

## Methods

### NewVcenterVmStoragePolicyComplianceInfo

`func NewVcenterVmStoragePolicyComplianceInfo(overallCompliance VcenterVmStoragePolicyComplianceStatus, disks []VcenterVmStoragePolicyComplianceInfoDisks, ) *VcenterVmStoragePolicyComplianceInfo`

NewVcenterVmStoragePolicyComplianceInfo instantiates a new VcenterVmStoragePolicyComplianceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmStoragePolicyComplianceInfoWithDefaults

`func NewVcenterVmStoragePolicyComplianceInfoWithDefaults() *VcenterVmStoragePolicyComplianceInfo`

NewVcenterVmStoragePolicyComplianceInfoWithDefaults instantiates a new VcenterVmStoragePolicyComplianceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverallCompliance

`func (o *VcenterVmStoragePolicyComplianceInfo) GetOverallCompliance() VcenterVmStoragePolicyComplianceStatus`

GetOverallCompliance returns the OverallCompliance field if non-nil, zero value otherwise.

### GetOverallComplianceOk

`func (o *VcenterVmStoragePolicyComplianceInfo) GetOverallComplianceOk() (*VcenterVmStoragePolicyComplianceStatus, bool)`

GetOverallComplianceOk returns a tuple with the OverallCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallCompliance

`func (o *VcenterVmStoragePolicyComplianceInfo) SetOverallCompliance(v VcenterVmStoragePolicyComplianceStatus)`

SetOverallCompliance sets OverallCompliance field to given value.


### GetVmHome

`func (o *VcenterVmStoragePolicyComplianceInfo) GetVmHome() VcenterVmStoragePolicyComplianceVmComplianceInfo`

GetVmHome returns the VmHome field if non-nil, zero value otherwise.

### GetVmHomeOk

`func (o *VcenterVmStoragePolicyComplianceInfo) GetVmHomeOk() (*VcenterVmStoragePolicyComplianceVmComplianceInfo, bool)`

GetVmHomeOk returns a tuple with the VmHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHome

`func (o *VcenterVmStoragePolicyComplianceInfo) SetVmHome(v VcenterVmStoragePolicyComplianceVmComplianceInfo)`

SetVmHome sets VmHome field to given value.

### HasVmHome

`func (o *VcenterVmStoragePolicyComplianceInfo) HasVmHome() bool`

HasVmHome returns a boolean if a field has been set.

### GetDisks

`func (o *VcenterVmStoragePolicyComplianceInfo) GetDisks() []VcenterVmStoragePolicyComplianceInfoDisks`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VcenterVmStoragePolicyComplianceInfo) GetDisksOk() (*[]VcenterVmStoragePolicyComplianceInfoDisks, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VcenterVmStoragePolicyComplianceInfo) SetDisks(v []VcenterVmStoragePolicyComplianceInfoDisks)`

SetDisks sets Disks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


