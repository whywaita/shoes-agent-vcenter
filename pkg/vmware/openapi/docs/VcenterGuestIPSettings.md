# VcenterGuestIPSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4** | Pointer to [**VcenterGuestIpv4**](VcenterGuestIpv4.md) |  | [optional] 
**Ipv6** | Pointer to [**VcenterGuestIpv6**](VcenterGuestIpv6.md) |  | [optional] 
**Windows** | Pointer to [**VcenterGuestWindowsNetworkAdapterSettings**](VcenterGuestWindowsNetworkAdapterSettings.md) |  | [optional] 

## Methods

### NewVcenterGuestIPSettings

`func NewVcenterGuestIPSettings() *VcenterGuestIPSettings`

NewVcenterGuestIPSettings instantiates a new VcenterGuestIPSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestIPSettingsWithDefaults

`func NewVcenterGuestIPSettingsWithDefaults() *VcenterGuestIPSettings`

NewVcenterGuestIPSettingsWithDefaults instantiates a new VcenterGuestIPSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4

`func (o *VcenterGuestIPSettings) GetIpv4() VcenterGuestIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *VcenterGuestIPSettings) GetIpv4Ok() (*VcenterGuestIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *VcenterGuestIPSettings) SetIpv4(v VcenterGuestIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *VcenterGuestIPSettings) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *VcenterGuestIPSettings) GetIpv6() VcenterGuestIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *VcenterGuestIPSettings) GetIpv6Ok() (*VcenterGuestIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *VcenterGuestIPSettings) SetIpv6(v VcenterGuestIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *VcenterGuestIPSettings) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetWindows

`func (o *VcenterGuestIPSettings) GetWindows() VcenterGuestWindowsNetworkAdapterSettings`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *VcenterGuestIPSettings) GetWindowsOk() (*VcenterGuestWindowsNetworkAdapterSettings, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *VcenterGuestIPSettings) SetWindows(v VcenterGuestWindowsNetworkAdapterSettings)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *VcenterGuestIPSettings) HasWindows() bool`

HasWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


