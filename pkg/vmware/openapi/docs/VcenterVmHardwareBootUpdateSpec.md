# VcenterVmHardwareBootUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**VcenterVmHardwareBootType**](VcenterVmHardwareBootType.md) |  | [optional] 
**EfiLegacyBoot** | Pointer to **bool** | Flag indicating whether to use EFI legacy boot mode. If unset, the value is unchanged. | [optional] 
**NetworkProtocol** | Pointer to [**VcenterVmHardwareBootNetworkProtocol**](VcenterVmHardwareBootNetworkProtocol.md) |  | [optional] 
**Delay** | Pointer to **int64** | Delay in milliseconds before beginning the firmware boot process when the virtual machine is powered on. This delay may be used to provide a time window for users to connect to the virtual machine console and enter BIOS setup mode. If unset, the value is unchanged. | [optional] 
**Retry** | Pointer to **bool** | Flag indicating whether the virtual machine should automatically retry the boot process after a failure. If unset, the value is unchanged. | [optional] 
**RetryDelay** | Pointer to **int64** | Delay in milliseconds before retrying the boot process after a failure; applicable only when Boot.Info.retry is true. If unset, the value is unchanged. | [optional] 
**EnterSetupMode** | Pointer to **bool** | Flag indicating whether the firmware boot process should automatically enter setup mode the next time the virtual machine boots. Note that this flag will automatically be reset to false once the virtual machine enters setup mode. If unset, the value is unchanged. | [optional] 

## Methods

### NewVcenterVmHardwareBootUpdateSpec

`func NewVcenterVmHardwareBootUpdateSpec() *VcenterVmHardwareBootUpdateSpec`

NewVcenterVmHardwareBootUpdateSpec instantiates a new VcenterVmHardwareBootUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareBootUpdateSpecWithDefaults

`func NewVcenterVmHardwareBootUpdateSpecWithDefaults() *VcenterVmHardwareBootUpdateSpec`

NewVcenterVmHardwareBootUpdateSpecWithDefaults instantiates a new VcenterVmHardwareBootUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareBootUpdateSpec) GetType() VcenterVmHardwareBootType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareBootUpdateSpec) GetTypeOk() (*VcenterVmHardwareBootType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareBootUpdateSpec) SetType(v VcenterVmHardwareBootType)`

SetType sets Type field to given value.

### HasType

`func (o *VcenterVmHardwareBootUpdateSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEfiLegacyBoot

`func (o *VcenterVmHardwareBootUpdateSpec) GetEfiLegacyBoot() bool`

GetEfiLegacyBoot returns the EfiLegacyBoot field if non-nil, zero value otherwise.

### GetEfiLegacyBootOk

`func (o *VcenterVmHardwareBootUpdateSpec) GetEfiLegacyBootOk() (*bool, bool)`

GetEfiLegacyBootOk returns a tuple with the EfiLegacyBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfiLegacyBoot

`func (o *VcenterVmHardwareBootUpdateSpec) SetEfiLegacyBoot(v bool)`

SetEfiLegacyBoot sets EfiLegacyBoot field to given value.

### HasEfiLegacyBoot

`func (o *VcenterVmHardwareBootUpdateSpec) HasEfiLegacyBoot() bool`

HasEfiLegacyBoot returns a boolean if a field has been set.

### GetNetworkProtocol

`func (o *VcenterVmHardwareBootUpdateSpec) GetNetworkProtocol() VcenterVmHardwareBootNetworkProtocol`

GetNetworkProtocol returns the NetworkProtocol field if non-nil, zero value otherwise.

### GetNetworkProtocolOk

`func (o *VcenterVmHardwareBootUpdateSpec) GetNetworkProtocolOk() (*VcenterVmHardwareBootNetworkProtocol, bool)`

GetNetworkProtocolOk returns a tuple with the NetworkProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtocol

`func (o *VcenterVmHardwareBootUpdateSpec) SetNetworkProtocol(v VcenterVmHardwareBootNetworkProtocol)`

SetNetworkProtocol sets NetworkProtocol field to given value.

### HasNetworkProtocol

`func (o *VcenterVmHardwareBootUpdateSpec) HasNetworkProtocol() bool`

HasNetworkProtocol returns a boolean if a field has been set.

### GetDelay

`func (o *VcenterVmHardwareBootUpdateSpec) GetDelay() int64`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *VcenterVmHardwareBootUpdateSpec) GetDelayOk() (*int64, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *VcenterVmHardwareBootUpdateSpec) SetDelay(v int64)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *VcenterVmHardwareBootUpdateSpec) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetRetry

`func (o *VcenterVmHardwareBootUpdateSpec) GetRetry() bool`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *VcenterVmHardwareBootUpdateSpec) GetRetryOk() (*bool, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *VcenterVmHardwareBootUpdateSpec) SetRetry(v bool)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *VcenterVmHardwareBootUpdateSpec) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetRetryDelay

`func (o *VcenterVmHardwareBootUpdateSpec) GetRetryDelay() int64`

GetRetryDelay returns the RetryDelay field if non-nil, zero value otherwise.

### GetRetryDelayOk

`func (o *VcenterVmHardwareBootUpdateSpec) GetRetryDelayOk() (*int64, bool)`

GetRetryDelayOk returns a tuple with the RetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelay

`func (o *VcenterVmHardwareBootUpdateSpec) SetRetryDelay(v int64)`

SetRetryDelay sets RetryDelay field to given value.

### HasRetryDelay

`func (o *VcenterVmHardwareBootUpdateSpec) HasRetryDelay() bool`

HasRetryDelay returns a boolean if a field has been set.

### GetEnterSetupMode

`func (o *VcenterVmHardwareBootUpdateSpec) GetEnterSetupMode() bool`

GetEnterSetupMode returns the EnterSetupMode field if non-nil, zero value otherwise.

### GetEnterSetupModeOk

`func (o *VcenterVmHardwareBootUpdateSpec) GetEnterSetupModeOk() (*bool, bool)`

GetEnterSetupModeOk returns a tuple with the EnterSetupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterSetupMode

`func (o *VcenterVmHardwareBootUpdateSpec) SetEnterSetupMode(v bool)`

SetEnterSetupMode sets EnterSetupMode field to given value.

### HasEnterSetupMode

`func (o *VcenterVmHardwareBootUpdateSpec) HasEnterSetupMode() bool`

HasEnterSetupMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


