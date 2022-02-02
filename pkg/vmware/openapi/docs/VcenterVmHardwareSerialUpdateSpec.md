# VcenterVmHardwareSerialUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**YieldOnPoll** | Pointer to **bool** | CPU yield behavior. If set to true, the virtual machine will periodically relinquish the processor if its sole task is polling the virtual serial port. The amount of time it takes to regain the processor will depend on the degree of other virtual machine activity on the host.   This field may be modified at any time, and changes applied to a connected virtual serial port take effect immediately.  If unset, the value is unchanged. | [optional] 
**Backing** | Pointer to [**VcenterVmHardwareSerialBackingSpec**](VcenterVmHardwareSerialBackingSpec.md) |  | [optional] 
**StartConnected** | Pointer to **bool** | Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. If unset, the value is unchanged. | [optional] 
**AllowGuestControl** | Pointer to **bool** | Flag indicating whether the guest can connect and disconnect the device. If unset, the value is unchanged. | [optional] 

## Methods

### NewVcenterVmHardwareSerialUpdateSpec

`func NewVcenterVmHardwareSerialUpdateSpec() *VcenterVmHardwareSerialUpdateSpec`

NewVcenterVmHardwareSerialUpdateSpec instantiates a new VcenterVmHardwareSerialUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareSerialUpdateSpecWithDefaults

`func NewVcenterVmHardwareSerialUpdateSpecWithDefaults() *VcenterVmHardwareSerialUpdateSpec`

NewVcenterVmHardwareSerialUpdateSpecWithDefaults instantiates a new VcenterVmHardwareSerialUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYieldOnPoll

`func (o *VcenterVmHardwareSerialUpdateSpec) GetYieldOnPoll() bool`

GetYieldOnPoll returns the YieldOnPoll field if non-nil, zero value otherwise.

### GetYieldOnPollOk

`func (o *VcenterVmHardwareSerialUpdateSpec) GetYieldOnPollOk() (*bool, bool)`

GetYieldOnPollOk returns a tuple with the YieldOnPoll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYieldOnPoll

`func (o *VcenterVmHardwareSerialUpdateSpec) SetYieldOnPoll(v bool)`

SetYieldOnPoll sets YieldOnPoll field to given value.

### HasYieldOnPoll

`func (o *VcenterVmHardwareSerialUpdateSpec) HasYieldOnPoll() bool`

HasYieldOnPoll returns a boolean if a field has been set.

### GetBacking

`func (o *VcenterVmHardwareSerialUpdateSpec) GetBacking() VcenterVmHardwareSerialBackingSpec`

GetBacking returns the Backing field if non-nil, zero value otherwise.

### GetBackingOk

`func (o *VcenterVmHardwareSerialUpdateSpec) GetBackingOk() (*VcenterVmHardwareSerialBackingSpec, bool)`

GetBackingOk returns a tuple with the Backing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacking

`func (o *VcenterVmHardwareSerialUpdateSpec) SetBacking(v VcenterVmHardwareSerialBackingSpec)`

SetBacking sets Backing field to given value.

### HasBacking

`func (o *VcenterVmHardwareSerialUpdateSpec) HasBacking() bool`

HasBacking returns a boolean if a field has been set.

### GetStartConnected

`func (o *VcenterVmHardwareSerialUpdateSpec) GetStartConnected() bool`

GetStartConnected returns the StartConnected field if non-nil, zero value otherwise.

### GetStartConnectedOk

`func (o *VcenterVmHardwareSerialUpdateSpec) GetStartConnectedOk() (*bool, bool)`

GetStartConnectedOk returns a tuple with the StartConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConnected

`func (o *VcenterVmHardwareSerialUpdateSpec) SetStartConnected(v bool)`

SetStartConnected sets StartConnected field to given value.

### HasStartConnected

`func (o *VcenterVmHardwareSerialUpdateSpec) HasStartConnected() bool`

HasStartConnected returns a boolean if a field has been set.

### GetAllowGuestControl

`func (o *VcenterVmHardwareSerialUpdateSpec) GetAllowGuestControl() bool`

GetAllowGuestControl returns the AllowGuestControl field if non-nil, zero value otherwise.

### GetAllowGuestControlOk

`func (o *VcenterVmHardwareSerialUpdateSpec) GetAllowGuestControlOk() (*bool, bool)`

GetAllowGuestControlOk returns a tuple with the AllowGuestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuestControl

`func (o *VcenterVmHardwareSerialUpdateSpec) SetAllowGuestControl(v bool)`

SetAllowGuestControl sets AllowGuestControl field to given value.

### HasAllowGuestControl

`func (o *VcenterVmHardwareSerialUpdateSpec) HasAllowGuestControl() bool`

HasAllowGuestControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


