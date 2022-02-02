# VcenterGuestWindowsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reboot** | Pointer to [**VcenterGuestWindowsConfigurationRebootOption**](VcenterGuestWindowsConfigurationRebootOption.md) |  | [optional] 
**Sysprep** | Pointer to [**VcenterGuestWindowsSysprep**](VcenterGuestWindowsSysprep.md) |  | [optional] 
**SysprepXml** | Pointer to **string** | All settings specified in a XML format. This is the content of a typical answer.xml file that is used by System administrators during the Windows image customization. Check https://docs.microsoft.com/en-us/windows-hardware/manufacture/desktop/update-windows-settings-and-scripts-create-your-own-answer-file-sxs Exactly one of WindowsConfiguration.sysprep or WindowsConfiguration.sysprep-xml must be specified. If unset, sysprep settings will not be applied to the windows guest operating system. | [optional] 

## Methods

### NewVcenterGuestWindowsConfiguration

`func NewVcenterGuestWindowsConfiguration() *VcenterGuestWindowsConfiguration`

NewVcenterGuestWindowsConfiguration instantiates a new VcenterGuestWindowsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestWindowsConfigurationWithDefaults

`func NewVcenterGuestWindowsConfigurationWithDefaults() *VcenterGuestWindowsConfiguration`

NewVcenterGuestWindowsConfigurationWithDefaults instantiates a new VcenterGuestWindowsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReboot

`func (o *VcenterGuestWindowsConfiguration) GetReboot() VcenterGuestWindowsConfigurationRebootOption`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *VcenterGuestWindowsConfiguration) GetRebootOk() (*VcenterGuestWindowsConfigurationRebootOption, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *VcenterGuestWindowsConfiguration) SetReboot(v VcenterGuestWindowsConfigurationRebootOption)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *VcenterGuestWindowsConfiguration) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetSysprep

`func (o *VcenterGuestWindowsConfiguration) GetSysprep() VcenterGuestWindowsSysprep`

GetSysprep returns the Sysprep field if non-nil, zero value otherwise.

### GetSysprepOk

`func (o *VcenterGuestWindowsConfiguration) GetSysprepOk() (*VcenterGuestWindowsSysprep, bool)`

GetSysprepOk returns a tuple with the Sysprep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysprep

`func (o *VcenterGuestWindowsConfiguration) SetSysprep(v VcenterGuestWindowsSysprep)`

SetSysprep sets Sysprep field to given value.

### HasSysprep

`func (o *VcenterGuestWindowsConfiguration) HasSysprep() bool`

HasSysprep returns a boolean if a field has been set.

### GetSysprepXml

`func (o *VcenterGuestWindowsConfiguration) GetSysprepXml() string`

GetSysprepXml returns the SysprepXml field if non-nil, zero value otherwise.

### GetSysprepXmlOk

`func (o *VcenterGuestWindowsConfiguration) GetSysprepXmlOk() (*string, bool)`

GetSysprepXmlOk returns a tuple with the SysprepXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysprepXml

`func (o *VcenterGuestWindowsConfiguration) SetSysprepXml(v string)`

SetSysprepXml sets SysprepXml field to given value.

### HasSysprepXml

`func (o *VcenterGuestWindowsConfiguration) HasSysprepXml() bool`

HasSysprepXml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


