# VcenterGuestGuiUnattended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoLogon** | **bool** | Flag to determine whether or not the machine automatically logs on as Administrator. See the GuiUnattended.password property. If GuiUnattended.auto-logon flag is set, then GuiUnattended.password must not be unset or the guest customization will fail. | 
**AutoLogonCount** | **int64** | If the GuiUnattended.auto-logon flag is set, then this property specifies the number of times the machine should automatically log on as Administrator. Generally it should be 1, but if the setup requires a number of reboots, you may want to increase it. This number may be determined by the list of commands. | 
**Password** | Pointer to **string** | The new administrator password for the machine. To specify that the password should be set to blank (that is, no password), leave it unset. If unset, blank password will be used. | [optional] 
**TimeZone** | **int64** | The time zone index for the virtual machine. Numbers correspond to time zones at https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values | 

## Methods

### NewVcenterGuestGuiUnattended

`func NewVcenterGuestGuiUnattended(autoLogon bool, autoLogonCount int64, timeZone int64, ) *VcenterGuestGuiUnattended`

NewVcenterGuestGuiUnattended instantiates a new VcenterGuestGuiUnattended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestGuiUnattendedWithDefaults

`func NewVcenterGuestGuiUnattendedWithDefaults() *VcenterGuestGuiUnattended`

NewVcenterGuestGuiUnattendedWithDefaults instantiates a new VcenterGuestGuiUnattended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoLogon

`func (o *VcenterGuestGuiUnattended) GetAutoLogon() bool`

GetAutoLogon returns the AutoLogon field if non-nil, zero value otherwise.

### GetAutoLogonOk

`func (o *VcenterGuestGuiUnattended) GetAutoLogonOk() (*bool, bool)`

GetAutoLogonOk returns a tuple with the AutoLogon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLogon

`func (o *VcenterGuestGuiUnattended) SetAutoLogon(v bool)`

SetAutoLogon sets AutoLogon field to given value.


### GetAutoLogonCount

`func (o *VcenterGuestGuiUnattended) GetAutoLogonCount() int64`

GetAutoLogonCount returns the AutoLogonCount field if non-nil, zero value otherwise.

### GetAutoLogonCountOk

`func (o *VcenterGuestGuiUnattended) GetAutoLogonCountOk() (*int64, bool)`

GetAutoLogonCountOk returns a tuple with the AutoLogonCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLogonCount

`func (o *VcenterGuestGuiUnattended) SetAutoLogonCount(v int64)`

SetAutoLogonCount sets AutoLogonCount field to given value.


### GetPassword

`func (o *VcenterGuestGuiUnattended) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterGuestGuiUnattended) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterGuestGuiUnattended) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VcenterGuestGuiUnattended) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTimeZone

`func (o *VcenterGuestGuiUnattended) GetTimeZone() int64`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *VcenterGuestGuiUnattended) GetTimeZoneOk() (*int64, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *VcenterGuestGuiUnattended) SetTimeZone(v int64)`

SetTimeZone sets TimeZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


