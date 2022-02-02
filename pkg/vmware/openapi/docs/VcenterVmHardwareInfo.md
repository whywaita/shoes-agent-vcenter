# VcenterVmHardwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | [**VcenterVmHardwareVersion**](VcenterVmHardwareVersion.md) |  | 
**UpgradePolicy** | [**VcenterVmHardwareUpgradePolicy**](VcenterVmHardwareUpgradePolicy.md) |  | 
**UpgradeVersion** | Pointer to [**VcenterVmHardwareVersion**](VcenterVmHardwareVersion.md) |  | [optional] 
**UpgradeStatus** | [**VcenterVmHardwareUpgradeStatus**](VcenterVmHardwareUpgradeStatus.md) |  | 
**UpgradeError** | Pointer to **string** | Reason for the scheduled upgrade failure. This field is optional and it is only relevant when the value of Hardware.Info.upgrade-status is FAILED. | [optional] 

## Methods

### NewVcenterVmHardwareInfo

`func NewVcenterVmHardwareInfo(version VcenterVmHardwareVersion, upgradePolicy VcenterVmHardwareUpgradePolicy, upgradeStatus VcenterVmHardwareUpgradeStatus, ) *VcenterVmHardwareInfo`

NewVcenterVmHardwareInfo instantiates a new VcenterVmHardwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareInfoWithDefaults

`func NewVcenterVmHardwareInfoWithDefaults() *VcenterVmHardwareInfo`

NewVcenterVmHardwareInfoWithDefaults instantiates a new VcenterVmHardwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *VcenterVmHardwareInfo) GetVersion() VcenterVmHardwareVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterVmHardwareInfo) GetVersionOk() (*VcenterVmHardwareVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterVmHardwareInfo) SetVersion(v VcenterVmHardwareVersion)`

SetVersion sets Version field to given value.


### GetUpgradePolicy

`func (o *VcenterVmHardwareInfo) GetUpgradePolicy() VcenterVmHardwareUpgradePolicy`

GetUpgradePolicy returns the UpgradePolicy field if non-nil, zero value otherwise.

### GetUpgradePolicyOk

`func (o *VcenterVmHardwareInfo) GetUpgradePolicyOk() (*VcenterVmHardwareUpgradePolicy, bool)`

GetUpgradePolicyOk returns a tuple with the UpgradePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePolicy

`func (o *VcenterVmHardwareInfo) SetUpgradePolicy(v VcenterVmHardwareUpgradePolicy)`

SetUpgradePolicy sets UpgradePolicy field to given value.


### GetUpgradeVersion

`func (o *VcenterVmHardwareInfo) GetUpgradeVersion() VcenterVmHardwareVersion`

GetUpgradeVersion returns the UpgradeVersion field if non-nil, zero value otherwise.

### GetUpgradeVersionOk

`func (o *VcenterVmHardwareInfo) GetUpgradeVersionOk() (*VcenterVmHardwareVersion, bool)`

GetUpgradeVersionOk returns a tuple with the UpgradeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeVersion

`func (o *VcenterVmHardwareInfo) SetUpgradeVersion(v VcenterVmHardwareVersion)`

SetUpgradeVersion sets UpgradeVersion field to given value.

### HasUpgradeVersion

`func (o *VcenterVmHardwareInfo) HasUpgradeVersion() bool`

HasUpgradeVersion returns a boolean if a field has been set.

### GetUpgradeStatus

`func (o *VcenterVmHardwareInfo) GetUpgradeStatus() VcenterVmHardwareUpgradeStatus`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *VcenterVmHardwareInfo) GetUpgradeStatusOk() (*VcenterVmHardwareUpgradeStatus, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *VcenterVmHardwareInfo) SetUpgradeStatus(v VcenterVmHardwareUpgradeStatus)`

SetUpgradeStatus sets UpgradeStatus field to given value.


### GetUpgradeError

`func (o *VcenterVmHardwareInfo) GetUpgradeError() string`

GetUpgradeError returns the UpgradeError field if non-nil, zero value otherwise.

### GetUpgradeErrorOk

`func (o *VcenterVmHardwareInfo) GetUpgradeErrorOk() (*string, bool)`

GetUpgradeErrorOk returns a tuple with the UpgradeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeError

`func (o *VcenterVmHardwareInfo) SetUpgradeError(v string)`

SetUpgradeError sets UpgradeError field to given value.

### HasUpgradeError

`func (o *VcenterVmHardwareInfo) HasUpgradeError() bool`

HasUpgradeError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


