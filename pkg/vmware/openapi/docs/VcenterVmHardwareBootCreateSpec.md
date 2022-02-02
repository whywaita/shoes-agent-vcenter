# VcenterVmHardwareBootCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**VcenterVmHardwareBootType**](VcenterVmHardwareBootType.md) |  | [optional] 
**EfiLegacyBoot** | Pointer to **bool** | Flag indicating whether to use EFI legacy boot mode. If unset, defaults to value that is recommended for the guest OS and is supported for the virtual hardware version. | [optional] 
**NetworkProtocol** | Pointer to [**VcenterVmHardwareBootNetworkProtocol**](VcenterVmHardwareBootNetworkProtocol.md) |  | [optional] 
**Delay** | Pointer to **int64** | Delay in milliseconds before beginning the firmware boot process when the virtual machine is powered on. This delay may be used to provide a time window for users to connect to the virtual machine console and enter BIOS setup mode. If unset, default value is 0. | [optional] 
**Retry** | Pointer to **bool** | Flag indicating whether the virtual machine should automatically retry the boot process after a failure. If unset, default value is false. | [optional] 
**RetryDelay** | Pointer to **int64** | Delay in milliseconds before retrying the boot process after a failure; applicable only when Boot.Info.retry is true. If unset, default value is 10000. | [optional] 
**EnterSetupMode** | Pointer to **bool** | Flag indicating whether the firmware boot process should automatically enter setup mode the next time the virtual machine boots. Note that this flag will automatically be reset to false once the virtual machine enters setup mode. If unset, the value is unchanged. | [optional] 

## Methods

### NewVcenterVmHardwareBootCreateSpec

`func NewVcenterVmHardwareBootCreateSpec() *VcenterVmHardwareBootCreateSpec`

NewVcenterVmHardwareBootCreateSpec instantiates a new VcenterVmHardwareBootCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareBootCreateSpecWithDefaults

`func NewVcenterVmHardwareBootCreateSpecWithDefaults() *VcenterVmHardwareBootCreateSpec`

NewVcenterVmHardwareBootCreateSpecWithDefaults instantiates a new VcenterVmHardwareBootCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareBootCreateSpec) GetType() VcenterVmHardwareBootType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareBootCreateSpec) GetTypeOk() (*VcenterVmHardwareBootType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareBootCreateSpec) SetType(v VcenterVmHardwareBootType)`

SetType sets Type field to given value.

### HasType

`func (o *VcenterVmHardwareBootCreateSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEfiLegacyBoot

`func (o *VcenterVmHardwareBootCreateSpec) GetEfiLegacyBoot() bool`

GetEfiLegacyBoot returns the EfiLegacyBoot field if non-nil, zero value otherwise.

### GetEfiLegacyBootOk

`func (o *VcenterVmHardwareBootCreateSpec) GetEfiLegacyBootOk() (*bool, bool)`

GetEfiLegacyBootOk returns a tuple with the EfiLegacyBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfiLegacyBoot

`func (o *VcenterVmHardwareBootCreateSpec) SetEfiLegacyBoot(v bool)`

SetEfiLegacyBoot sets EfiLegacyBoot field to given value.

### HasEfiLegacyBoot

`func (o *VcenterVmHardwareBootCreateSpec) HasEfiLegacyBoot() bool`

HasEfiLegacyBoot returns a boolean if a field has been set.

### GetNetworkProtocol

`func (o *VcenterVmHardwareBootCreateSpec) GetNetworkProtocol() VcenterVmHardwareBootNetworkProtocol`

GetNetworkProtocol returns the NetworkProtocol field if non-nil, zero value otherwise.

### GetNetworkProtocolOk

`func (o *VcenterVmHardwareBootCreateSpec) GetNetworkProtocolOk() (*VcenterVmHardwareBootNetworkProtocol, bool)`

GetNetworkProtocolOk returns a tuple with the NetworkProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtocol

`func (o *VcenterVmHardwareBootCreateSpec) SetNetworkProtocol(v VcenterVmHardwareBootNetworkProtocol)`

SetNetworkProtocol sets NetworkProtocol field to given value.

### HasNetworkProtocol

`func (o *VcenterVmHardwareBootCreateSpec) HasNetworkProtocol() bool`

HasNetworkProtocol returns a boolean if a field has been set.

### GetDelay

`func (o *VcenterVmHardwareBootCreateSpec) GetDelay() int64`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *VcenterVmHardwareBootCreateSpec) GetDelayOk() (*int64, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *VcenterVmHardwareBootCreateSpec) SetDelay(v int64)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *VcenterVmHardwareBootCreateSpec) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetRetry

`func (o *VcenterVmHardwareBootCreateSpec) GetRetry() bool`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *VcenterVmHardwareBootCreateSpec) GetRetryOk() (*bool, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *VcenterVmHardwareBootCreateSpec) SetRetry(v bool)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *VcenterVmHardwareBootCreateSpec) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetRetryDelay

`func (o *VcenterVmHardwareBootCreateSpec) GetRetryDelay() int64`

GetRetryDelay returns the RetryDelay field if non-nil, zero value otherwise.

### GetRetryDelayOk

`func (o *VcenterVmHardwareBootCreateSpec) GetRetryDelayOk() (*int64, bool)`

GetRetryDelayOk returns a tuple with the RetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelay

`func (o *VcenterVmHardwareBootCreateSpec) SetRetryDelay(v int64)`

SetRetryDelay sets RetryDelay field to given value.

### HasRetryDelay

`func (o *VcenterVmHardwareBootCreateSpec) HasRetryDelay() bool`

HasRetryDelay returns a boolean if a field has been set.

### GetEnterSetupMode

`func (o *VcenterVmHardwareBootCreateSpec) GetEnterSetupMode() bool`

GetEnterSetupMode returns the EnterSetupMode field if non-nil, zero value otherwise.

### GetEnterSetupModeOk

`func (o *VcenterVmHardwareBootCreateSpec) GetEnterSetupModeOk() (*bool, bool)`

GetEnterSetupModeOk returns a tuple with the EnterSetupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterSetupMode

`func (o *VcenterVmHardwareBootCreateSpec) SetEnterSetupMode(v bool)`

SetEnterSetupMode sets EnterSetupMode field to given value.

### HasEnterSetupMode

`func (o *VcenterVmHardwareBootCreateSpec) HasEnterSetupMode() bool`

HasEnterSetupMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


