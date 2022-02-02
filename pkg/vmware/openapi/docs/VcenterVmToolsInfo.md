# VcenterVmToolsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUpdateSupported** | **bool** | Set if the virtual machine supports auto-upgrading Tools via Tools.UpgradePolicy. | 
**InstallAttemptCount** | Pointer to **int64** | Number of attempts that have been made to install or upgrade the version of Tools installed on this virtual machine. This field will be unset if there have been no Tools install or upgrade attempt. | [optional] 
**Error** | Pointer to **string** | Error that happened, if any, during last attempt to upgrade or install Tools. This field will be unset if a the last Tools install or upgrade attempt succeeded. | [optional] 
**VersionNumber** | Pointer to **int64** | Version of VMware Tools installed on the guest operating system. This field wil be unset if VMWare Tools is not installed. This is an integer constructed as follows: (((MJR) &lt;&lt; 10) + ((MNR) &lt;&lt; 5) + (REV)) Where MJR is tha major verson, MNR is the minor version and REV is the revision. Tools version &#x3D; T Tools Version Major &#x3D; MJR &#x3D; (T / 1024) Tools Version Minor &#x3D; MNR &#x3D; ((T % 1024) / 32) Tools Version Revision &#x3D; BASE &#x3D; ((T % 1024) % 32) Tools actual version &#x3D; MJR.MNR.REV | [optional] 
**Version** | Pointer to **string** | Version of VMware Tools installed on the guest operating system. This is a human-readable value that should not be parsed. This field wil be unset if VMWare Tools is not installed. | [optional] 
**UpgradePolicy** | [**VcenterVmToolsUpgradePolicy**](VcenterVmToolsUpgradePolicy.md) |  | 
**VersionStatus** | Pointer to [**VcenterVmToolsVersionStatus**](VcenterVmToolsVersionStatus.md) |  | [optional] 
**InstallType** | Pointer to [**VcenterVmToolsToolsInstallType**](VcenterVmToolsToolsInstallType.md) |  | [optional] 
**RunState** | [**VcenterVmToolsRunState**](VcenterVmToolsRunState.md) |  | 

## Methods

### NewVcenterVmToolsInfo

`func NewVcenterVmToolsInfo(autoUpdateSupported bool, upgradePolicy VcenterVmToolsUpgradePolicy, runState VcenterVmToolsRunState, ) *VcenterVmToolsInfo`

NewVcenterVmToolsInfo instantiates a new VcenterVmToolsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmToolsInfoWithDefaults

`func NewVcenterVmToolsInfoWithDefaults() *VcenterVmToolsInfo`

NewVcenterVmToolsInfoWithDefaults instantiates a new VcenterVmToolsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUpdateSupported

`func (o *VcenterVmToolsInfo) GetAutoUpdateSupported() bool`

GetAutoUpdateSupported returns the AutoUpdateSupported field if non-nil, zero value otherwise.

### GetAutoUpdateSupportedOk

`func (o *VcenterVmToolsInfo) GetAutoUpdateSupportedOk() (*bool, bool)`

GetAutoUpdateSupportedOk returns a tuple with the AutoUpdateSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdateSupported

`func (o *VcenterVmToolsInfo) SetAutoUpdateSupported(v bool)`

SetAutoUpdateSupported sets AutoUpdateSupported field to given value.


### GetInstallAttemptCount

`func (o *VcenterVmToolsInfo) GetInstallAttemptCount() int64`

GetInstallAttemptCount returns the InstallAttemptCount field if non-nil, zero value otherwise.

### GetInstallAttemptCountOk

`func (o *VcenterVmToolsInfo) GetInstallAttemptCountOk() (*int64, bool)`

GetInstallAttemptCountOk returns a tuple with the InstallAttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAttemptCount

`func (o *VcenterVmToolsInfo) SetInstallAttemptCount(v int64)`

SetInstallAttemptCount sets InstallAttemptCount field to given value.

### HasInstallAttemptCount

`func (o *VcenterVmToolsInfo) HasInstallAttemptCount() bool`

HasInstallAttemptCount returns a boolean if a field has been set.

### GetError

`func (o *VcenterVmToolsInfo) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VcenterVmToolsInfo) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VcenterVmToolsInfo) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *VcenterVmToolsInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### GetVersionNumber

`func (o *VcenterVmToolsInfo) GetVersionNumber() int64`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *VcenterVmToolsInfo) GetVersionNumberOk() (*int64, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *VcenterVmToolsInfo) SetVersionNumber(v int64)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *VcenterVmToolsInfo) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### GetVersion

`func (o *VcenterVmToolsInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterVmToolsInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterVmToolsInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VcenterVmToolsInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUpgradePolicy

`func (o *VcenterVmToolsInfo) GetUpgradePolicy() VcenterVmToolsUpgradePolicy`

GetUpgradePolicy returns the UpgradePolicy field if non-nil, zero value otherwise.

### GetUpgradePolicyOk

`func (o *VcenterVmToolsInfo) GetUpgradePolicyOk() (*VcenterVmToolsUpgradePolicy, bool)`

GetUpgradePolicyOk returns a tuple with the UpgradePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePolicy

`func (o *VcenterVmToolsInfo) SetUpgradePolicy(v VcenterVmToolsUpgradePolicy)`

SetUpgradePolicy sets UpgradePolicy field to given value.


### GetVersionStatus

`func (o *VcenterVmToolsInfo) GetVersionStatus() VcenterVmToolsVersionStatus`

GetVersionStatus returns the VersionStatus field if non-nil, zero value otherwise.

### GetVersionStatusOk

`func (o *VcenterVmToolsInfo) GetVersionStatusOk() (*VcenterVmToolsVersionStatus, bool)`

GetVersionStatusOk returns a tuple with the VersionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionStatus

`func (o *VcenterVmToolsInfo) SetVersionStatus(v VcenterVmToolsVersionStatus)`

SetVersionStatus sets VersionStatus field to given value.

### HasVersionStatus

`func (o *VcenterVmToolsInfo) HasVersionStatus() bool`

HasVersionStatus returns a boolean if a field has been set.

### GetInstallType

`func (o *VcenterVmToolsInfo) GetInstallType() VcenterVmToolsToolsInstallType`

GetInstallType returns the InstallType field if non-nil, zero value otherwise.

### GetInstallTypeOk

`func (o *VcenterVmToolsInfo) GetInstallTypeOk() (*VcenterVmToolsToolsInstallType, bool)`

GetInstallTypeOk returns a tuple with the InstallType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallType

`func (o *VcenterVmToolsInfo) SetInstallType(v VcenterVmToolsToolsInstallType)`

SetInstallType sets InstallType field to given value.

### HasInstallType

`func (o *VcenterVmToolsInfo) HasInstallType() bool`

HasInstallType returns a boolean if a field has been set.

### GetRunState

`func (o *VcenterVmToolsInfo) GetRunState() VcenterVmToolsRunState`

GetRunState returns the RunState field if non-nil, zero value otherwise.

### GetRunStateOk

`func (o *VcenterVmToolsInfo) GetRunStateOk() (*VcenterVmToolsRunState, bool)`

GetRunStateOk returns a tuple with the RunState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunState

`func (o *VcenterVmToolsInfo) SetRunState(v VcenterVmToolsRunState)`

SetRunState sets RunState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


