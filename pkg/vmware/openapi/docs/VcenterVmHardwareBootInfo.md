# VcenterVmHardwareBootInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmHardwareBootType**](VcenterVmHardwareBootType.md) |  | 
**EfiLegacyBoot** | Pointer to **bool** | Flag indicating whether to use EFI legacy boot mode. This field is optional and it is only relevant when the value of Boot.Info.type is EFI. | [optional] 
**NetworkProtocol** | Pointer to [**VcenterVmHardwareBootNetworkProtocol**](VcenterVmHardwareBootNetworkProtocol.md) |  | [optional] 
**Delay** | **int64** | Delay in milliseconds before beginning the firmware boot process when the virtual machine is powered on. This delay may be used to provide a time window for users to connect to the virtual machine console and enter BIOS setup mode. | 
**Retry** | **bool** | Flag indicating whether the virtual machine will automatically retry the boot process after a failure. | 
**RetryDelay** | **int64** | Delay in milliseconds before retrying the boot process after a failure; applicable only when Boot.Info.retry is true. | 
**EnterSetupMode** | **bool** | Flag indicating whether the firmware boot process will automatically enter setup mode the next time the virtual machine boots. Note that this flag will automatically be reset to false once the virtual machine enters setup mode. | 

## Methods

### NewVcenterVmHardwareBootInfo

`func NewVcenterVmHardwareBootInfo(type_ VcenterVmHardwareBootType, delay int64, retry bool, retryDelay int64, enterSetupMode bool, ) *VcenterVmHardwareBootInfo`

NewVcenterVmHardwareBootInfo instantiates a new VcenterVmHardwareBootInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareBootInfoWithDefaults

`func NewVcenterVmHardwareBootInfoWithDefaults() *VcenterVmHardwareBootInfo`

NewVcenterVmHardwareBootInfoWithDefaults instantiates a new VcenterVmHardwareBootInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareBootInfo) GetType() VcenterVmHardwareBootType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareBootInfo) GetTypeOk() (*VcenterVmHardwareBootType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareBootInfo) SetType(v VcenterVmHardwareBootType)`

SetType sets Type field to given value.


### GetEfiLegacyBoot

`func (o *VcenterVmHardwareBootInfo) GetEfiLegacyBoot() bool`

GetEfiLegacyBoot returns the EfiLegacyBoot field if non-nil, zero value otherwise.

### GetEfiLegacyBootOk

`func (o *VcenterVmHardwareBootInfo) GetEfiLegacyBootOk() (*bool, bool)`

GetEfiLegacyBootOk returns a tuple with the EfiLegacyBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfiLegacyBoot

`func (o *VcenterVmHardwareBootInfo) SetEfiLegacyBoot(v bool)`

SetEfiLegacyBoot sets EfiLegacyBoot field to given value.

### HasEfiLegacyBoot

`func (o *VcenterVmHardwareBootInfo) HasEfiLegacyBoot() bool`

HasEfiLegacyBoot returns a boolean if a field has been set.

### GetNetworkProtocol

`func (o *VcenterVmHardwareBootInfo) GetNetworkProtocol() VcenterVmHardwareBootNetworkProtocol`

GetNetworkProtocol returns the NetworkProtocol field if non-nil, zero value otherwise.

### GetNetworkProtocolOk

`func (o *VcenterVmHardwareBootInfo) GetNetworkProtocolOk() (*VcenterVmHardwareBootNetworkProtocol, bool)`

GetNetworkProtocolOk returns a tuple with the NetworkProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtocol

`func (o *VcenterVmHardwareBootInfo) SetNetworkProtocol(v VcenterVmHardwareBootNetworkProtocol)`

SetNetworkProtocol sets NetworkProtocol field to given value.

### HasNetworkProtocol

`func (o *VcenterVmHardwareBootInfo) HasNetworkProtocol() bool`

HasNetworkProtocol returns a boolean if a field has been set.

### GetDelay

`func (o *VcenterVmHardwareBootInfo) GetDelay() int64`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *VcenterVmHardwareBootInfo) GetDelayOk() (*int64, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *VcenterVmHardwareBootInfo) SetDelay(v int64)`

SetDelay sets Delay field to given value.


### GetRetry

`func (o *VcenterVmHardwareBootInfo) GetRetry() bool`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *VcenterVmHardwareBootInfo) GetRetryOk() (*bool, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *VcenterVmHardwareBootInfo) SetRetry(v bool)`

SetRetry sets Retry field to given value.


### GetRetryDelay

`func (o *VcenterVmHardwareBootInfo) GetRetryDelay() int64`

GetRetryDelay returns the RetryDelay field if non-nil, zero value otherwise.

### GetRetryDelayOk

`func (o *VcenterVmHardwareBootInfo) GetRetryDelayOk() (*int64, bool)`

GetRetryDelayOk returns a tuple with the RetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelay

`func (o *VcenterVmHardwareBootInfo) SetRetryDelay(v int64)`

SetRetryDelay sets RetryDelay field to given value.


### GetEnterSetupMode

`func (o *VcenterVmHardwareBootInfo) GetEnterSetupMode() bool`

GetEnterSetupMode returns the EnterSetupMode field if non-nil, zero value otherwise.

### GetEnterSetupModeOk

`func (o *VcenterVmHardwareBootInfo) GetEnterSetupModeOk() (*bool, bool)`

GetEnterSetupModeOk returns a tuple with the EnterSetupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterSetupMode

`func (o *VcenterVmHardwareBootInfo) SetEnterSetupMode(v bool)`

SetEnterSetupMode sets EnterSetupMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


