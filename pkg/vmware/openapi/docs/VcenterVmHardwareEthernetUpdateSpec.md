# VcenterVmHardwareEthernetUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UptCompatibilityEnabled** | Pointer to **bool** | Flag indicating whether Universal Pass-Through (UPT) compatibility should be enabled on this virtual Ethernet adapter.   This field may be modified at any time, and changes will be applied the next time the virtual machine is powered on.  If unset, the value is unchanged. Must be unset if the emulation type of the virtual Ethernet adapter is not VMXNET3. | [optional] 
**MacType** | Pointer to [**VcenterVmHardwareEthernetMacAddressType**](VcenterVmHardwareEthernetMacAddressType.md) |  | [optional] 
**MacAddress** | Pointer to **string** | MAC address.   This field may be modified at any time, and changes will be applied the next time the virtual machine is powered on.  If unset, the value is unchanged. Must be specified if Ethernet.UpdateSpec.mac-type is MANUAL. Must be unset if the MAC address type is not MANUAL. | [optional] 
**WakeOnLanEnabled** | Pointer to **bool** | Flag indicating whether wake-on-LAN shoud be enabled on this virtual Ethernet adapter.   This field may be modified at any time, and changes will be applied the next time the virtual machine is powered on.  If unset, the value is unchanged. | [optional] 
**Backing** | Pointer to [**VcenterVmHardwareEthernetBackingSpec**](VcenterVmHardwareEthernetBackingSpec.md) |  | [optional] 
**StartConnected** | Pointer to **bool** | Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. If unset, the value is unchanged. | [optional] 
**AllowGuestControl** | Pointer to **bool** | Flag indicating whether the guest can connect and disconnect the device. If unset, the value is unchanged. | [optional] 

## Methods

### NewVcenterVmHardwareEthernetUpdateSpec

`func NewVcenterVmHardwareEthernetUpdateSpec() *VcenterVmHardwareEthernetUpdateSpec`

NewVcenterVmHardwareEthernetUpdateSpec instantiates a new VcenterVmHardwareEthernetUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareEthernetUpdateSpecWithDefaults

`func NewVcenterVmHardwareEthernetUpdateSpecWithDefaults() *VcenterVmHardwareEthernetUpdateSpec`

NewVcenterVmHardwareEthernetUpdateSpecWithDefaults instantiates a new VcenterVmHardwareEthernetUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUptCompatibilityEnabled

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetUptCompatibilityEnabled() bool`

GetUptCompatibilityEnabled returns the UptCompatibilityEnabled field if non-nil, zero value otherwise.

### GetUptCompatibilityEnabledOk

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetUptCompatibilityEnabledOk() (*bool, bool)`

GetUptCompatibilityEnabledOk returns a tuple with the UptCompatibilityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptCompatibilityEnabled

`func (o *VcenterVmHardwareEthernetUpdateSpec) SetUptCompatibilityEnabled(v bool)`

SetUptCompatibilityEnabled sets UptCompatibilityEnabled field to given value.

### HasUptCompatibilityEnabled

`func (o *VcenterVmHardwareEthernetUpdateSpec) HasUptCompatibilityEnabled() bool`

HasUptCompatibilityEnabled returns a boolean if a field has been set.

### GetMacType

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetMacType() VcenterVmHardwareEthernetMacAddressType`

GetMacType returns the MacType field if non-nil, zero value otherwise.

### GetMacTypeOk

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetMacTypeOk() (*VcenterVmHardwareEthernetMacAddressType, bool)`

GetMacTypeOk returns a tuple with the MacType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacType

`func (o *VcenterVmHardwareEthernetUpdateSpec) SetMacType(v VcenterVmHardwareEthernetMacAddressType)`

SetMacType sets MacType field to given value.

### HasMacType

`func (o *VcenterVmHardwareEthernetUpdateSpec) HasMacType() bool`

HasMacType returns a boolean if a field has been set.

### GetMacAddress

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VcenterVmHardwareEthernetUpdateSpec) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VcenterVmHardwareEthernetUpdateSpec) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetWakeOnLanEnabled

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetWakeOnLanEnabled() bool`

GetWakeOnLanEnabled returns the WakeOnLanEnabled field if non-nil, zero value otherwise.

### GetWakeOnLanEnabledOk

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetWakeOnLanEnabledOk() (*bool, bool)`

GetWakeOnLanEnabledOk returns a tuple with the WakeOnLanEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanEnabled

`func (o *VcenterVmHardwareEthernetUpdateSpec) SetWakeOnLanEnabled(v bool)`

SetWakeOnLanEnabled sets WakeOnLanEnabled field to given value.

### HasWakeOnLanEnabled

`func (o *VcenterVmHardwareEthernetUpdateSpec) HasWakeOnLanEnabled() bool`

HasWakeOnLanEnabled returns a boolean if a field has been set.

### GetBacking

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetBacking() VcenterVmHardwareEthernetBackingSpec`

GetBacking returns the Backing field if non-nil, zero value otherwise.

### GetBackingOk

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetBackingOk() (*VcenterVmHardwareEthernetBackingSpec, bool)`

GetBackingOk returns a tuple with the Backing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacking

`func (o *VcenterVmHardwareEthernetUpdateSpec) SetBacking(v VcenterVmHardwareEthernetBackingSpec)`

SetBacking sets Backing field to given value.

### HasBacking

`func (o *VcenterVmHardwareEthernetUpdateSpec) HasBacking() bool`

HasBacking returns a boolean if a field has been set.

### GetStartConnected

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetStartConnected() bool`

GetStartConnected returns the StartConnected field if non-nil, zero value otherwise.

### GetStartConnectedOk

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetStartConnectedOk() (*bool, bool)`

GetStartConnectedOk returns a tuple with the StartConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConnected

`func (o *VcenterVmHardwareEthernetUpdateSpec) SetStartConnected(v bool)`

SetStartConnected sets StartConnected field to given value.

### HasStartConnected

`func (o *VcenterVmHardwareEthernetUpdateSpec) HasStartConnected() bool`

HasStartConnected returns a boolean if a field has been set.

### GetAllowGuestControl

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetAllowGuestControl() bool`

GetAllowGuestControl returns the AllowGuestControl field if non-nil, zero value otherwise.

### GetAllowGuestControlOk

`func (o *VcenterVmHardwareEthernetUpdateSpec) GetAllowGuestControlOk() (*bool, bool)`

GetAllowGuestControlOk returns a tuple with the AllowGuestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuestControl

`func (o *VcenterVmHardwareEthernetUpdateSpec) SetAllowGuestControl(v bool)`

SetAllowGuestControl sets AllowGuestControl field to given value.

### HasAllowGuestControl

`func (o *VcenterVmHardwareEthernetUpdateSpec) HasAllowGuestControl() bool`

HasAllowGuestControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


