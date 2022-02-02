# VcenterVmHardwareParallelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Device label. | 
**Backing** | [**VcenterVmHardwareParallelBackingInfo**](VcenterVmHardwareParallelBackingInfo.md) |  | 
**State** | [**VcenterVmHardwareConnectionState**](VcenterVmHardwareConnectionState.md) |  | 
**StartConnected** | **bool** | Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. | 
**AllowGuestControl** | **bool** | Flag indicating whether the guest can connect and disconnect the device. | 

## Methods

### NewVcenterVmHardwareParallelInfo

`func NewVcenterVmHardwareParallelInfo(label string, backing VcenterVmHardwareParallelBackingInfo, state VcenterVmHardwareConnectionState, startConnected bool, allowGuestControl bool, ) *VcenterVmHardwareParallelInfo`

NewVcenterVmHardwareParallelInfo instantiates a new VcenterVmHardwareParallelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareParallelInfoWithDefaults

`func NewVcenterVmHardwareParallelInfoWithDefaults() *VcenterVmHardwareParallelInfo`

NewVcenterVmHardwareParallelInfoWithDefaults instantiates a new VcenterVmHardwareParallelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *VcenterVmHardwareParallelInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VcenterVmHardwareParallelInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VcenterVmHardwareParallelInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetBacking

`func (o *VcenterVmHardwareParallelInfo) GetBacking() VcenterVmHardwareParallelBackingInfo`

GetBacking returns the Backing field if non-nil, zero value otherwise.

### GetBackingOk

`func (o *VcenterVmHardwareParallelInfo) GetBackingOk() (*VcenterVmHardwareParallelBackingInfo, bool)`

GetBackingOk returns a tuple with the Backing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacking

`func (o *VcenterVmHardwareParallelInfo) SetBacking(v VcenterVmHardwareParallelBackingInfo)`

SetBacking sets Backing field to given value.


### GetState

`func (o *VcenterVmHardwareParallelInfo) GetState() VcenterVmHardwareConnectionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterVmHardwareParallelInfo) GetStateOk() (*VcenterVmHardwareConnectionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterVmHardwareParallelInfo) SetState(v VcenterVmHardwareConnectionState)`

SetState sets State field to given value.


### GetStartConnected

`func (o *VcenterVmHardwareParallelInfo) GetStartConnected() bool`

GetStartConnected returns the StartConnected field if non-nil, zero value otherwise.

### GetStartConnectedOk

`func (o *VcenterVmHardwareParallelInfo) GetStartConnectedOk() (*bool, bool)`

GetStartConnectedOk returns a tuple with the StartConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConnected

`func (o *VcenterVmHardwareParallelInfo) SetStartConnected(v bool)`

SetStartConnected sets StartConnected field to given value.


### GetAllowGuestControl

`func (o *VcenterVmHardwareParallelInfo) GetAllowGuestControl() bool`

GetAllowGuestControl returns the AllowGuestControl field if non-nil, zero value otherwise.

### GetAllowGuestControlOk

`func (o *VcenterVmHardwareParallelInfo) GetAllowGuestControlOk() (*bool, bool)`

GetAllowGuestControlOk returns a tuple with the AllowGuestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuestControl

`func (o *VcenterVmHardwareParallelInfo) SetAllowGuestControl(v bool)`

SetAllowGuestControl sets AllowGuestControl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


