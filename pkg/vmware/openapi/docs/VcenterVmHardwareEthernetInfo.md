# VcenterVmHardwareEthernetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Device label. | 
**Type** | [**VcenterVmHardwareEthernetEmulationType**](VcenterVmHardwareEthernetEmulationType.md) |  | 
**UptCompatibilityEnabled** | Pointer to **bool** | Flag indicating whether Universal Pass-Through (UPT) compatibility is enabled on this virtual Ethernet adapter. This field is optional and it is only relevant when the value of Ethernet.Info.type is VMXNET3. | [optional] 
**MacType** | [**VcenterVmHardwareEthernetMacAddressType**](VcenterVmHardwareEthernetMacAddressType.md) |  | 
**MacAddress** | Pointer to **string** | MAC address. May be unset if Ethernet.Info.mac-type is MANUAL and has not been specified, or if Ethernet.Info.mac-type is GENERATED and the virtual machine has never been powered on since the Ethernet adapter was created. | [optional] 
**PciSlotNumber** | Pointer to **int64** | Address of the virtual Ethernet adapter on the PCI bus. If the PCI address is invalid, the server will change it when the VM is started or as the device is hot added. May be unset if the virtual machine has never been powered on since the adapter was created. | [optional] 
**WakeOnLanEnabled** | **bool** | Flag indicating whether wake-on-LAN is enabled on this virtual Ethernet adapter. | 
**Backing** | [**VcenterVmHardwareEthernetBackingInfo**](VcenterVmHardwareEthernetBackingInfo.md) |  | 
**State** | [**VcenterVmHardwareConnectionState**](VcenterVmHardwareConnectionState.md) |  | 
**StartConnected** | **bool** | Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. | 
**AllowGuestControl** | **bool** | Flag indicating whether the guest can connect and disconnect the device. | 

## Methods

### NewVcenterVmHardwareEthernetInfo

`func NewVcenterVmHardwareEthernetInfo(label string, type_ VcenterVmHardwareEthernetEmulationType, macType VcenterVmHardwareEthernetMacAddressType, wakeOnLanEnabled bool, backing VcenterVmHardwareEthernetBackingInfo, state VcenterVmHardwareConnectionState, startConnected bool, allowGuestControl bool, ) *VcenterVmHardwareEthernetInfo`

NewVcenterVmHardwareEthernetInfo instantiates a new VcenterVmHardwareEthernetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareEthernetInfoWithDefaults

`func NewVcenterVmHardwareEthernetInfoWithDefaults() *VcenterVmHardwareEthernetInfo`

NewVcenterVmHardwareEthernetInfoWithDefaults instantiates a new VcenterVmHardwareEthernetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *VcenterVmHardwareEthernetInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VcenterVmHardwareEthernetInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VcenterVmHardwareEthernetInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *VcenterVmHardwareEthernetInfo) GetType() VcenterVmHardwareEthernetEmulationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareEthernetInfo) GetTypeOk() (*VcenterVmHardwareEthernetEmulationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareEthernetInfo) SetType(v VcenterVmHardwareEthernetEmulationType)`

SetType sets Type field to given value.


### GetUptCompatibilityEnabled

`func (o *VcenterVmHardwareEthernetInfo) GetUptCompatibilityEnabled() bool`

GetUptCompatibilityEnabled returns the UptCompatibilityEnabled field if non-nil, zero value otherwise.

### GetUptCompatibilityEnabledOk

`func (o *VcenterVmHardwareEthernetInfo) GetUptCompatibilityEnabledOk() (*bool, bool)`

GetUptCompatibilityEnabledOk returns a tuple with the UptCompatibilityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptCompatibilityEnabled

`func (o *VcenterVmHardwareEthernetInfo) SetUptCompatibilityEnabled(v bool)`

SetUptCompatibilityEnabled sets UptCompatibilityEnabled field to given value.

### HasUptCompatibilityEnabled

`func (o *VcenterVmHardwareEthernetInfo) HasUptCompatibilityEnabled() bool`

HasUptCompatibilityEnabled returns a boolean if a field has been set.

### GetMacType

`func (o *VcenterVmHardwareEthernetInfo) GetMacType() VcenterVmHardwareEthernetMacAddressType`

GetMacType returns the MacType field if non-nil, zero value otherwise.

### GetMacTypeOk

`func (o *VcenterVmHardwareEthernetInfo) GetMacTypeOk() (*VcenterVmHardwareEthernetMacAddressType, bool)`

GetMacTypeOk returns a tuple with the MacType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacType

`func (o *VcenterVmHardwareEthernetInfo) SetMacType(v VcenterVmHardwareEthernetMacAddressType)`

SetMacType sets MacType field to given value.


### GetMacAddress

`func (o *VcenterVmHardwareEthernetInfo) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VcenterVmHardwareEthernetInfo) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VcenterVmHardwareEthernetInfo) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VcenterVmHardwareEthernetInfo) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPciSlotNumber

`func (o *VcenterVmHardwareEthernetInfo) GetPciSlotNumber() int64`

GetPciSlotNumber returns the PciSlotNumber field if non-nil, zero value otherwise.

### GetPciSlotNumberOk

`func (o *VcenterVmHardwareEthernetInfo) GetPciSlotNumberOk() (*int64, bool)`

GetPciSlotNumberOk returns a tuple with the PciSlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlotNumber

`func (o *VcenterVmHardwareEthernetInfo) SetPciSlotNumber(v int64)`

SetPciSlotNumber sets PciSlotNumber field to given value.

### HasPciSlotNumber

`func (o *VcenterVmHardwareEthernetInfo) HasPciSlotNumber() bool`

HasPciSlotNumber returns a boolean if a field has been set.

### GetWakeOnLanEnabled

`func (o *VcenterVmHardwareEthernetInfo) GetWakeOnLanEnabled() bool`

GetWakeOnLanEnabled returns the WakeOnLanEnabled field if non-nil, zero value otherwise.

### GetWakeOnLanEnabledOk

`func (o *VcenterVmHardwareEthernetInfo) GetWakeOnLanEnabledOk() (*bool, bool)`

GetWakeOnLanEnabledOk returns a tuple with the WakeOnLanEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanEnabled

`func (o *VcenterVmHardwareEthernetInfo) SetWakeOnLanEnabled(v bool)`

SetWakeOnLanEnabled sets WakeOnLanEnabled field to given value.


### GetBacking

`func (o *VcenterVmHardwareEthernetInfo) GetBacking() VcenterVmHardwareEthernetBackingInfo`

GetBacking returns the Backing field if non-nil, zero value otherwise.

### GetBackingOk

`func (o *VcenterVmHardwareEthernetInfo) GetBackingOk() (*VcenterVmHardwareEthernetBackingInfo, bool)`

GetBackingOk returns a tuple with the Backing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacking

`func (o *VcenterVmHardwareEthernetInfo) SetBacking(v VcenterVmHardwareEthernetBackingInfo)`

SetBacking sets Backing field to given value.


### GetState

`func (o *VcenterVmHardwareEthernetInfo) GetState() VcenterVmHardwareConnectionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterVmHardwareEthernetInfo) GetStateOk() (*VcenterVmHardwareConnectionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterVmHardwareEthernetInfo) SetState(v VcenterVmHardwareConnectionState)`

SetState sets State field to given value.


### GetStartConnected

`func (o *VcenterVmHardwareEthernetInfo) GetStartConnected() bool`

GetStartConnected returns the StartConnected field if non-nil, zero value otherwise.

### GetStartConnectedOk

`func (o *VcenterVmHardwareEthernetInfo) GetStartConnectedOk() (*bool, bool)`

GetStartConnectedOk returns a tuple with the StartConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConnected

`func (o *VcenterVmHardwareEthernetInfo) SetStartConnected(v bool)`

SetStartConnected sets StartConnected field to given value.


### GetAllowGuestControl

`func (o *VcenterVmHardwareEthernetInfo) GetAllowGuestControl() bool`

GetAllowGuestControl returns the AllowGuestControl field if non-nil, zero value otherwise.

### GetAllowGuestControlOk

`func (o *VcenterVmHardwareEthernetInfo) GetAllowGuestControlOk() (*bool, bool)`

GetAllowGuestControlOk returns a tuple with the AllowGuestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuestControl

`func (o *VcenterVmHardwareEthernetInfo) SetAllowGuestControl(v bool)`

SetAllowGuestControl sets AllowGuestControl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


