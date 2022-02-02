# VcenterVmHardwareParallelUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backing** | Pointer to [**VcenterVmHardwareParallelBackingSpec**](VcenterVmHardwareParallelBackingSpec.md) |  | [optional] 
**StartConnected** | Pointer to **bool** | Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. If unset, the value is unchanged. | [optional] 
**AllowGuestControl** | Pointer to **bool** | Flag indicating whether the guest can connect and disconnect the device. If unset, the value is unchanged. | [optional] 

## Methods

### NewVcenterVmHardwareParallelUpdateSpec

`func NewVcenterVmHardwareParallelUpdateSpec() *VcenterVmHardwareParallelUpdateSpec`

NewVcenterVmHardwareParallelUpdateSpec instantiates a new VcenterVmHardwareParallelUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareParallelUpdateSpecWithDefaults

`func NewVcenterVmHardwareParallelUpdateSpecWithDefaults() *VcenterVmHardwareParallelUpdateSpec`

NewVcenterVmHardwareParallelUpdateSpecWithDefaults instantiates a new VcenterVmHardwareParallelUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBacking

`func (o *VcenterVmHardwareParallelUpdateSpec) GetBacking() VcenterVmHardwareParallelBackingSpec`

GetBacking returns the Backing field if non-nil, zero value otherwise.

### GetBackingOk

`func (o *VcenterVmHardwareParallelUpdateSpec) GetBackingOk() (*VcenterVmHardwareParallelBackingSpec, bool)`

GetBackingOk returns a tuple with the Backing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacking

`func (o *VcenterVmHardwareParallelUpdateSpec) SetBacking(v VcenterVmHardwareParallelBackingSpec)`

SetBacking sets Backing field to given value.

### HasBacking

`func (o *VcenterVmHardwareParallelUpdateSpec) HasBacking() bool`

HasBacking returns a boolean if a field has been set.

### GetStartConnected

`func (o *VcenterVmHardwareParallelUpdateSpec) GetStartConnected() bool`

GetStartConnected returns the StartConnected field if non-nil, zero value otherwise.

### GetStartConnectedOk

`func (o *VcenterVmHardwareParallelUpdateSpec) GetStartConnectedOk() (*bool, bool)`

GetStartConnectedOk returns a tuple with the StartConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConnected

`func (o *VcenterVmHardwareParallelUpdateSpec) SetStartConnected(v bool)`

SetStartConnected sets StartConnected field to given value.

### HasStartConnected

`func (o *VcenterVmHardwareParallelUpdateSpec) HasStartConnected() bool`

HasStartConnected returns a boolean if a field has been set.

### GetAllowGuestControl

`func (o *VcenterVmHardwareParallelUpdateSpec) GetAllowGuestControl() bool`

GetAllowGuestControl returns the AllowGuestControl field if non-nil, zero value otherwise.

### GetAllowGuestControlOk

`func (o *VcenterVmHardwareParallelUpdateSpec) GetAllowGuestControlOk() (*bool, bool)`

GetAllowGuestControlOk returns a tuple with the AllowGuestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuestControl

`func (o *VcenterVmHardwareParallelUpdateSpec) SetAllowGuestControl(v bool)`

SetAllowGuestControl sets AllowGuestControl field to given value.

### HasAllowGuestControl

`func (o *VcenterVmHardwareParallelUpdateSpec) HasAllowGuestControl() bool`

HasAllowGuestControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


