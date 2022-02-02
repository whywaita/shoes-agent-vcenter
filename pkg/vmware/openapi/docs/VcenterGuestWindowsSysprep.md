# VcenterGuestWindowsSysprep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuiRunOnceCommands** | Pointer to **[]string** | A list of commands to run at first user logon, after customizing the guest. These commands are directly mapped to the GuiRunOnce key in the sysprep.xml answer file. These commands are transferred into the sysprep.xml file that VirtualCenter stores on the target virtual disk. For more information about performing unattended installation, check https://technet.microsoft.com/en-us/library/cc771830(v&#x3D;ws.10).aspx The commands listed here ar executed when a user logs on the first time after customization completes. The logon may be driven by GuiUnattended.auto-logon setting. These commands are directly mapped to the GuiRunOnce key in the If unset, no commands are run. | [optional] 
**UserData** | [**VcenterGuestUserData**](VcenterGuestUserData.md) |  | 
**Domain** | Pointer to [**VcenterGuestDomain**](VcenterGuestDomain.md) |  | [optional] 
**GuiUnattended** | [**VcenterGuestGuiUnattended**](VcenterGuestGuiUnattended.md) |  | 

## Methods

### NewVcenterGuestWindowsSysprep

`func NewVcenterGuestWindowsSysprep(userData VcenterGuestUserData, guiUnattended VcenterGuestGuiUnattended, ) *VcenterGuestWindowsSysprep`

NewVcenterGuestWindowsSysprep instantiates a new VcenterGuestWindowsSysprep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestWindowsSysprepWithDefaults

`func NewVcenterGuestWindowsSysprepWithDefaults() *VcenterGuestWindowsSysprep`

NewVcenterGuestWindowsSysprepWithDefaults instantiates a new VcenterGuestWindowsSysprep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuiRunOnceCommands

`func (o *VcenterGuestWindowsSysprep) GetGuiRunOnceCommands() []string`

GetGuiRunOnceCommands returns the GuiRunOnceCommands field if non-nil, zero value otherwise.

### GetGuiRunOnceCommandsOk

`func (o *VcenterGuestWindowsSysprep) GetGuiRunOnceCommandsOk() (*[]string, bool)`

GetGuiRunOnceCommandsOk returns a tuple with the GuiRunOnceCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiRunOnceCommands

`func (o *VcenterGuestWindowsSysprep) SetGuiRunOnceCommands(v []string)`

SetGuiRunOnceCommands sets GuiRunOnceCommands field to given value.

### HasGuiRunOnceCommands

`func (o *VcenterGuestWindowsSysprep) HasGuiRunOnceCommands() bool`

HasGuiRunOnceCommands returns a boolean if a field has been set.

### GetUserData

`func (o *VcenterGuestWindowsSysprep) GetUserData() VcenterGuestUserData`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *VcenterGuestWindowsSysprep) GetUserDataOk() (*VcenterGuestUserData, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *VcenterGuestWindowsSysprep) SetUserData(v VcenterGuestUserData)`

SetUserData sets UserData field to given value.


### GetDomain

`func (o *VcenterGuestWindowsSysprep) GetDomain() VcenterGuestDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *VcenterGuestWindowsSysprep) GetDomainOk() (*VcenterGuestDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *VcenterGuestWindowsSysprep) SetDomain(v VcenterGuestDomain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *VcenterGuestWindowsSysprep) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetGuiUnattended

`func (o *VcenterGuestWindowsSysprep) GetGuiUnattended() VcenterGuestGuiUnattended`

GetGuiUnattended returns the GuiUnattended field if non-nil, zero value otherwise.

### GetGuiUnattendedOk

`func (o *VcenterGuestWindowsSysprep) GetGuiUnattendedOk() (*VcenterGuestGuiUnattended, bool)`

GetGuiUnattendedOk returns a tuple with the GuiUnattended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiUnattended

`func (o *VcenterGuestWindowsSysprep) SetGuiUnattended(v VcenterGuestGuiUnattended)`

SetGuiUnattended sets GuiUnattended field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


