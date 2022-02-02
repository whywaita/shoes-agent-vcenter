# VcenterVmHardwareFloppyUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backing** | Pointer to [**VcenterVmHardwareFloppyBackingSpec**](VcenterVmHardwareFloppyBackingSpec.md) |  | [optional] 
**StartConnected** | Pointer to **bool** | Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. If unset, the value is unchanged. | [optional] 
**AllowGuestControl** | Pointer to **bool** | Flag indicating whether the guest can connect and disconnect the device. If unset, the value is unchanged. | [optional] 

## Methods

### NewVcenterVmHardwareFloppyUpdateSpec

`func NewVcenterVmHardwareFloppyUpdateSpec() *VcenterVmHardwareFloppyUpdateSpec`

NewVcenterVmHardwareFloppyUpdateSpec instantiates a new VcenterVmHardwareFloppyUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareFloppyUpdateSpecWithDefaults

`func NewVcenterVmHardwareFloppyUpdateSpecWithDefaults() *VcenterVmHardwareFloppyUpdateSpec`

NewVcenterVmHardwareFloppyUpdateSpecWithDefaults instantiates a new VcenterVmHardwareFloppyUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBacking

`func (o *VcenterVmHardwareFloppyUpdateSpec) GetBacking() VcenterVmHardwareFloppyBackingSpec`

GetBacking returns the Backing field if non-nil, zero value otherwise.

### GetBackingOk

`func (o *VcenterVmHardwareFloppyUpdateSpec) GetBackingOk() (*VcenterVmHardwareFloppyBackingSpec, bool)`

GetBackingOk returns a tuple with the Backing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacking

`func (o *VcenterVmHardwareFloppyUpdateSpec) SetBacking(v VcenterVmHardwareFloppyBackingSpec)`

SetBacking sets Backing field to given value.

### HasBacking

`func (o *VcenterVmHardwareFloppyUpdateSpec) HasBacking() bool`

HasBacking returns a boolean if a field has been set.

### GetStartConnected

`func (o *VcenterVmHardwareFloppyUpdateSpec) GetStartConnected() bool`

GetStartConnected returns the StartConnected field if non-nil, zero value otherwise.

### GetStartConnectedOk

`func (o *VcenterVmHardwareFloppyUpdateSpec) GetStartConnectedOk() (*bool, bool)`

GetStartConnectedOk returns a tuple with the StartConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConnected

`func (o *VcenterVmHardwareFloppyUpdateSpec) SetStartConnected(v bool)`

SetStartConnected sets StartConnected field to given value.

### HasStartConnected

`func (o *VcenterVmHardwareFloppyUpdateSpec) HasStartConnected() bool`

HasStartConnected returns a boolean if a field has been set.

### GetAllowGuestControl

`func (o *VcenterVmHardwareFloppyUpdateSpec) GetAllowGuestControl() bool`

GetAllowGuestControl returns the AllowGuestControl field if non-nil, zero value otherwise.

### GetAllowGuestControlOk

`func (o *VcenterVmHardwareFloppyUpdateSpec) GetAllowGuestControlOk() (*bool, bool)`

GetAllowGuestControlOk returns a tuple with the AllowGuestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuestControl

`func (o *VcenterVmHardwareFloppyUpdateSpec) SetAllowGuestControl(v bool)`

SetAllowGuestControl sets AllowGuestControl field to given value.

### HasAllowGuestControl

`func (o *VcenterVmHardwareFloppyUpdateSpec) HasAllowGuestControl() bool`

HasAllowGuestControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


