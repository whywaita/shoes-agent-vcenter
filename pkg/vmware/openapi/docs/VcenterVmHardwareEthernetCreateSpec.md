# VcenterVmHardwareEthernetCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**VcenterVmHardwareEthernetEmulationType**](VcenterVmHardwareEthernetEmulationType.md) |  | [optional] 
**UptCompatibilityEnabled** | Pointer to **bool** | Flag indicating whether Universal Pass-Through (UPT) compatibility is enabled on this virtual Ethernet adapter. If unset, defaults to false. | [optional] 
**MacType** | Pointer to [**VcenterVmHardwareEthernetMacAddressType**](VcenterVmHardwareEthernetMacAddressType.md) |  | [optional] 
**MacAddress** | Pointer to **string** | MAC address. Workaround for PR1459647 | [optional] 
**PciSlotNumber** | Pointer to **int64** | Address of the virtual Ethernet adapter on the PCI bus. If the PCI address is invalid, the server will change when it the VM is started or as the device is hot added. If unset, the server will choose an available address when the virtual machine is powered on. | [optional] 
**WakeOnLanEnabled** | Pointer to **bool** | Flag indicating whether wake-on-LAN is enabled on this virtual Ethernet adapter. Defaults to false if unset. | [optional] 
**Backing** | Pointer to [**VcenterVmHardwareEthernetBackingSpec**](VcenterVmHardwareEthernetBackingSpec.md) |  | [optional] 
**StartConnected** | Pointer to **bool** | Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. Defaults to false if unset. | [optional] 
**AllowGuestControl** | Pointer to **bool** | Flag indicating whether the guest can connect and disconnect the device. Defaults to false if unset. | [optional] 

## Methods

### NewVcenterVmHardwareEthernetCreateSpec

`func NewVcenterVmHardwareEthernetCreateSpec() *VcenterVmHardwareEthernetCreateSpec`

NewVcenterVmHardwareEthernetCreateSpec instantiates a new VcenterVmHardwareEthernetCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareEthernetCreateSpecWithDefaults

`func NewVcenterVmHardwareEthernetCreateSpecWithDefaults() *VcenterVmHardwareEthernetCreateSpec`

NewVcenterVmHardwareEthernetCreateSpecWithDefaults instantiates a new VcenterVmHardwareEthernetCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareEthernetCreateSpec) GetType() VcenterVmHardwareEthernetEmulationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareEthernetCreateSpec) GetTypeOk() (*VcenterVmHardwareEthernetEmulationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareEthernetCreateSpec) SetType(v VcenterVmHardwareEthernetEmulationType)`

SetType sets Type field to given value.

### HasType

`func (o *VcenterVmHardwareEthernetCreateSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptCompatibilityEnabled

`func (o *VcenterVmHardwareEthernetCreateSpec) GetUptCompatibilityEnabled() bool`

GetUptCompatibilityEnabled returns the UptCompatibilityEnabled field if non-nil, zero value otherwise.

### GetUptCompatibilityEnabledOk

`func (o *VcenterVmHardwareEthernetCreateSpec) GetUptCompatibilityEnabledOk() (*bool, bool)`

GetUptCompatibilityEnabledOk returns a tuple with the UptCompatibilityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptCompatibilityEnabled

`func (o *VcenterVmHardwareEthernetCreateSpec) SetUptCompatibilityEnabled(v bool)`

SetUptCompatibilityEnabled sets UptCompatibilityEnabled field to given value.

### HasUptCompatibilityEnabled

`func (o *VcenterVmHardwareEthernetCreateSpec) HasUptCompatibilityEnabled() bool`

HasUptCompatibilityEnabled returns a boolean if a field has been set.

### GetMacType

`func (o *VcenterVmHardwareEthernetCreateSpec) GetMacType() VcenterVmHardwareEthernetMacAddressType`

GetMacType returns the MacType field if non-nil, zero value otherwise.

### GetMacTypeOk

`func (o *VcenterVmHardwareEthernetCreateSpec) GetMacTypeOk() (*VcenterVmHardwareEthernetMacAddressType, bool)`

GetMacTypeOk returns a tuple with the MacType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacType

`func (o *VcenterVmHardwareEthernetCreateSpec) SetMacType(v VcenterVmHardwareEthernetMacAddressType)`

SetMacType sets MacType field to given value.

### HasMacType

`func (o *VcenterVmHardwareEthernetCreateSpec) HasMacType() bool`

HasMacType returns a boolean if a field has been set.

### GetMacAddress

`func (o *VcenterVmHardwareEthernetCreateSpec) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VcenterVmHardwareEthernetCreateSpec) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VcenterVmHardwareEthernetCreateSpec) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VcenterVmHardwareEthernetCreateSpec) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPciSlotNumber

`func (o *VcenterVmHardwareEthernetCreateSpec) GetPciSlotNumber() int64`

GetPciSlotNumber returns the PciSlotNumber field if non-nil, zero value otherwise.

### GetPciSlotNumberOk

`func (o *VcenterVmHardwareEthernetCreateSpec) GetPciSlotNumberOk() (*int64, bool)`

GetPciSlotNumberOk returns a tuple with the PciSlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlotNumber

`func (o *VcenterVmHardwareEthernetCreateSpec) SetPciSlotNumber(v int64)`

SetPciSlotNumber sets PciSlotNumber field to given value.

### HasPciSlotNumber

`func (o *VcenterVmHardwareEthernetCreateSpec) HasPciSlotNumber() bool`

HasPciSlotNumber returns a boolean if a field has been set.

### GetWakeOnLanEnabled

`func (o *VcenterVmHardwareEthernetCreateSpec) GetWakeOnLanEnabled() bool`

GetWakeOnLanEnabled returns the WakeOnLanEnabled field if non-nil, zero value otherwise.

### GetWakeOnLanEnabledOk

`func (o *VcenterVmHardwareEthernetCreateSpec) GetWakeOnLanEnabledOk() (*bool, bool)`

GetWakeOnLanEnabledOk returns a tuple with the WakeOnLanEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanEnabled

`func (o *VcenterVmHardwareEthernetCreateSpec) SetWakeOnLanEnabled(v bool)`

SetWakeOnLanEnabled sets WakeOnLanEnabled field to given value.

### HasWakeOnLanEnabled

`func (o *VcenterVmHardwareEthernetCreateSpec) HasWakeOnLanEnabled() bool`

HasWakeOnLanEnabled returns a boolean if a field has been set.

### GetBacking

`func (o *VcenterVmHardwareEthernetCreateSpec) GetBacking() VcenterVmHardwareEthernetBackingSpec`

GetBacking returns the Backing field if non-nil, zero value otherwise.

### GetBackingOk

`func (o *VcenterVmHardwareEthernetCreateSpec) GetBackingOk() (*VcenterVmHardwareEthernetBackingSpec, bool)`

GetBackingOk returns a tuple with the Backing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacking

`func (o *VcenterVmHardwareEthernetCreateSpec) SetBacking(v VcenterVmHardwareEthernetBackingSpec)`

SetBacking sets Backing field to given value.

### HasBacking

`func (o *VcenterVmHardwareEthernetCreateSpec) HasBacking() bool`

HasBacking returns a boolean if a field has been set.

### GetStartConnected

`func (o *VcenterVmHardwareEthernetCreateSpec) GetStartConnected() bool`

GetStartConnected returns the StartConnected field if non-nil, zero value otherwise.

### GetStartConnectedOk

`func (o *VcenterVmHardwareEthernetCreateSpec) GetStartConnectedOk() (*bool, bool)`

GetStartConnectedOk returns a tuple with the StartConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConnected

`func (o *VcenterVmHardwareEthernetCreateSpec) SetStartConnected(v bool)`

SetStartConnected sets StartConnected field to given value.

### HasStartConnected

`func (o *VcenterVmHardwareEthernetCreateSpec) HasStartConnected() bool`

HasStartConnected returns a boolean if a field has been set.

### GetAllowGuestControl

`func (o *VcenterVmHardwareEthernetCreateSpec) GetAllowGuestControl() bool`

GetAllowGuestControl returns the AllowGuestControl field if non-nil, zero value otherwise.

### GetAllowGuestControlOk

`func (o *VcenterVmHardwareEthernetCreateSpec) GetAllowGuestControlOk() (*bool, bool)`

GetAllowGuestControlOk returns a tuple with the AllowGuestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuestControl

`func (o *VcenterVmHardwareEthernetCreateSpec) SetAllowGuestControl(v bool)`

SetAllowGuestControl sets AllowGuestControl field to given value.

### HasAllowGuestControl

`func (o *VcenterVmHardwareEthernetCreateSpec) HasAllowGuestControl() bool`

HasAllowGuestControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


