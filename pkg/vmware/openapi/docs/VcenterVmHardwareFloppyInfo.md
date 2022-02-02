# VcenterVmHardwareFloppyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Device label. | 
**Backing** | [**VcenterVmHardwareFloppyBackingInfo**](VcenterVmHardwareFloppyBackingInfo.md) |  | 
**State** | [**VcenterVmHardwareConnectionState**](VcenterVmHardwareConnectionState.md) |  | 
**StartConnected** | **bool** | Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. | 
**AllowGuestControl** | **bool** | Flag indicating whether the guest can connect and disconnect the device. | 

## Methods

### NewVcenterVmHardwareFloppyInfo

`func NewVcenterVmHardwareFloppyInfo(label string, backing VcenterVmHardwareFloppyBackingInfo, state VcenterVmHardwareConnectionState, startConnected bool, allowGuestControl bool, ) *VcenterVmHardwareFloppyInfo`

NewVcenterVmHardwareFloppyInfo instantiates a new VcenterVmHardwareFloppyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareFloppyInfoWithDefaults

`func NewVcenterVmHardwareFloppyInfoWithDefaults() *VcenterVmHardwareFloppyInfo`

NewVcenterVmHardwareFloppyInfoWithDefaults instantiates a new VcenterVmHardwareFloppyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *VcenterVmHardwareFloppyInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VcenterVmHardwareFloppyInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VcenterVmHardwareFloppyInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetBacking

`func (o *VcenterVmHardwareFloppyInfo) GetBacking() VcenterVmHardwareFloppyBackingInfo`

GetBacking returns the Backing field if non-nil, zero value otherwise.

### GetBackingOk

`func (o *VcenterVmHardwareFloppyInfo) GetBackingOk() (*VcenterVmHardwareFloppyBackingInfo, bool)`

GetBackingOk returns a tuple with the Backing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacking

`func (o *VcenterVmHardwareFloppyInfo) SetBacking(v VcenterVmHardwareFloppyBackingInfo)`

SetBacking sets Backing field to given value.


### GetState

`func (o *VcenterVmHardwareFloppyInfo) GetState() VcenterVmHardwareConnectionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterVmHardwareFloppyInfo) GetStateOk() (*VcenterVmHardwareConnectionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterVmHardwareFloppyInfo) SetState(v VcenterVmHardwareConnectionState)`

SetState sets State field to given value.


### GetStartConnected

`func (o *VcenterVmHardwareFloppyInfo) GetStartConnected() bool`

GetStartConnected returns the StartConnected field if non-nil, zero value otherwise.

### GetStartConnectedOk

`func (o *VcenterVmHardwareFloppyInfo) GetStartConnectedOk() (*bool, bool)`

GetStartConnectedOk returns a tuple with the StartConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConnected

`func (o *VcenterVmHardwareFloppyInfo) SetStartConnected(v bool)`

SetStartConnected sets StartConnected field to given value.


### GetAllowGuestControl

`func (o *VcenterVmHardwareFloppyInfo) GetAllowGuestControl() bool`

GetAllowGuestControl returns the AllowGuestControl field if non-nil, zero value otherwise.

### GetAllowGuestControlOk

`func (o *VcenterVmHardwareFloppyInfo) GetAllowGuestControlOk() (*bool, bool)`

GetAllowGuestControlOk returns a tuple with the AllowGuestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuestControl

`func (o *VcenterVmHardwareFloppyInfo) SetAllowGuestControl(v bool)`

SetAllowGuestControl sets AllowGuestControl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


