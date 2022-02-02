# VcenterVmHardwareSerialInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Device label. | 
**YieldOnPoll** | **bool** | CPU yield behavior. If set to true, the virtual machine will periodically relinquish the processor if its sole task is polling the virtual serial port. The amount of time it takes to regain the processor will depend on the degree of other virtual machine activity on the host. | 
**Backing** | [**VcenterVmHardwareSerialBackingInfo**](VcenterVmHardwareSerialBackingInfo.md) |  | 
**State** | [**VcenterVmHardwareConnectionState**](VcenterVmHardwareConnectionState.md) |  | 
**StartConnected** | **bool** | Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. | 
**AllowGuestControl** | **bool** | Flag indicating whether the guest can connect and disconnect the device. | 

## Methods

### NewVcenterVmHardwareSerialInfo

`func NewVcenterVmHardwareSerialInfo(label string, yieldOnPoll bool, backing VcenterVmHardwareSerialBackingInfo, state VcenterVmHardwareConnectionState, startConnected bool, allowGuestControl bool, ) *VcenterVmHardwareSerialInfo`

NewVcenterVmHardwareSerialInfo instantiates a new VcenterVmHardwareSerialInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareSerialInfoWithDefaults

`func NewVcenterVmHardwareSerialInfoWithDefaults() *VcenterVmHardwareSerialInfo`

NewVcenterVmHardwareSerialInfoWithDefaults instantiates a new VcenterVmHardwareSerialInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *VcenterVmHardwareSerialInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VcenterVmHardwareSerialInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VcenterVmHardwareSerialInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetYieldOnPoll

`func (o *VcenterVmHardwareSerialInfo) GetYieldOnPoll() bool`

GetYieldOnPoll returns the YieldOnPoll field if non-nil, zero value otherwise.

### GetYieldOnPollOk

`func (o *VcenterVmHardwareSerialInfo) GetYieldOnPollOk() (*bool, bool)`

GetYieldOnPollOk returns a tuple with the YieldOnPoll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYieldOnPoll

`func (o *VcenterVmHardwareSerialInfo) SetYieldOnPoll(v bool)`

SetYieldOnPoll sets YieldOnPoll field to given value.


### GetBacking

`func (o *VcenterVmHardwareSerialInfo) GetBacking() VcenterVmHardwareSerialBackingInfo`

GetBacking returns the Backing field if non-nil, zero value otherwise.

### GetBackingOk

`func (o *VcenterVmHardwareSerialInfo) GetBackingOk() (*VcenterVmHardwareSerialBackingInfo, bool)`

GetBackingOk returns a tuple with the Backing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacking

`func (o *VcenterVmHardwareSerialInfo) SetBacking(v VcenterVmHardwareSerialBackingInfo)`

SetBacking sets Backing field to given value.


### GetState

`func (o *VcenterVmHardwareSerialInfo) GetState() VcenterVmHardwareConnectionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterVmHardwareSerialInfo) GetStateOk() (*VcenterVmHardwareConnectionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterVmHardwareSerialInfo) SetState(v VcenterVmHardwareConnectionState)`

SetState sets State field to given value.


### GetStartConnected

`func (o *VcenterVmHardwareSerialInfo) GetStartConnected() bool`

GetStartConnected returns the StartConnected field if non-nil, zero value otherwise.

### GetStartConnectedOk

`func (o *VcenterVmHardwareSerialInfo) GetStartConnectedOk() (*bool, bool)`

GetStartConnectedOk returns a tuple with the StartConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConnected

`func (o *VcenterVmHardwareSerialInfo) SetStartConnected(v bool)`

SetStartConnected sets StartConnected field to given value.


### GetAllowGuestControl

`func (o *VcenterVmHardwareSerialInfo) GetAllowGuestControl() bool`

GetAllowGuestControl returns the AllowGuestControl field if non-nil, zero value otherwise.

### GetAllowGuestControlOk

`func (o *VcenterVmHardwareSerialInfo) GetAllowGuestControlOk() (*bool, bool)`

GetAllowGuestControlOk returns a tuple with the AllowGuestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuestControl

`func (o *VcenterVmHardwareSerialInfo) SetAllowGuestControl(v bool)`

SetAllowGuestControl sets AllowGuestControl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


