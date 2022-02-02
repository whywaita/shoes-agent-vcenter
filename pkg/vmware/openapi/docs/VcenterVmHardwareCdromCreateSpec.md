# VcenterVmHardwareCdromCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**VcenterVmHardwareCdromHostBusAdapterType**](VcenterVmHardwareCdromHostBusAdapterType.md) |  | [optional] 
**Ide** | Pointer to [**VcenterVmHardwareIdeAddressSpec**](VcenterVmHardwareIdeAddressSpec.md) |  | [optional] 
**Sata** | Pointer to [**VcenterVmHardwareSataAddressSpec**](VcenterVmHardwareSataAddressSpec.md) |  | [optional] 
**Backing** | Pointer to [**VcenterVmHardwareCdromBackingSpec**](VcenterVmHardwareCdromBackingSpec.md) |  | [optional] 
**StartConnected** | Pointer to **bool** | Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. Defaults to false if unset. | [optional] 
**AllowGuestControl** | Pointer to **bool** | Flag indicating whether the guest can connect and disconnect the device. Defaults to false if unset. | [optional] 

## Methods

### NewVcenterVmHardwareCdromCreateSpec

`func NewVcenterVmHardwareCdromCreateSpec() *VcenterVmHardwareCdromCreateSpec`

NewVcenterVmHardwareCdromCreateSpec instantiates a new VcenterVmHardwareCdromCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareCdromCreateSpecWithDefaults

`func NewVcenterVmHardwareCdromCreateSpecWithDefaults() *VcenterVmHardwareCdromCreateSpec`

NewVcenterVmHardwareCdromCreateSpecWithDefaults instantiates a new VcenterVmHardwareCdromCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareCdromCreateSpec) GetType() VcenterVmHardwareCdromHostBusAdapterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareCdromCreateSpec) GetTypeOk() (*VcenterVmHardwareCdromHostBusAdapterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareCdromCreateSpec) SetType(v VcenterVmHardwareCdromHostBusAdapterType)`

SetType sets Type field to given value.

### HasType

`func (o *VcenterVmHardwareCdromCreateSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIde

`func (o *VcenterVmHardwareCdromCreateSpec) GetIde() VcenterVmHardwareIdeAddressSpec`

GetIde returns the Ide field if non-nil, zero value otherwise.

### GetIdeOk

`func (o *VcenterVmHardwareCdromCreateSpec) GetIdeOk() (*VcenterVmHardwareIdeAddressSpec, bool)`

GetIdeOk returns a tuple with the Ide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde

`func (o *VcenterVmHardwareCdromCreateSpec) SetIde(v VcenterVmHardwareIdeAddressSpec)`

SetIde sets Ide field to given value.

### HasIde

`func (o *VcenterVmHardwareCdromCreateSpec) HasIde() bool`

HasIde returns a boolean if a field has been set.

### GetSata

`func (o *VcenterVmHardwareCdromCreateSpec) GetSata() VcenterVmHardwareSataAddressSpec`

GetSata returns the Sata field if non-nil, zero value otherwise.

### GetSataOk

`func (o *VcenterVmHardwareCdromCreateSpec) GetSataOk() (*VcenterVmHardwareSataAddressSpec, bool)`

GetSataOk returns a tuple with the Sata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata

`func (o *VcenterVmHardwareCdromCreateSpec) SetSata(v VcenterVmHardwareSataAddressSpec)`

SetSata sets Sata field to given value.

### HasSata

`func (o *VcenterVmHardwareCdromCreateSpec) HasSata() bool`

HasSata returns a boolean if a field has been set.

### GetBacking

`func (o *VcenterVmHardwareCdromCreateSpec) GetBacking() VcenterVmHardwareCdromBackingSpec`

GetBacking returns the Backing field if non-nil, zero value otherwise.

### GetBackingOk

`func (o *VcenterVmHardwareCdromCreateSpec) GetBackingOk() (*VcenterVmHardwareCdromBackingSpec, bool)`

GetBackingOk returns a tuple with the Backing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacking

`func (o *VcenterVmHardwareCdromCreateSpec) SetBacking(v VcenterVmHardwareCdromBackingSpec)`

SetBacking sets Backing field to given value.

### HasBacking

`func (o *VcenterVmHardwareCdromCreateSpec) HasBacking() bool`

HasBacking returns a boolean if a field has been set.

### GetStartConnected

`func (o *VcenterVmHardwareCdromCreateSpec) GetStartConnected() bool`

GetStartConnected returns the StartConnected field if non-nil, zero value otherwise.

### GetStartConnectedOk

`func (o *VcenterVmHardwareCdromCreateSpec) GetStartConnectedOk() (*bool, bool)`

GetStartConnectedOk returns a tuple with the StartConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConnected

`func (o *VcenterVmHardwareCdromCreateSpec) SetStartConnected(v bool)`

SetStartConnected sets StartConnected field to given value.

### HasStartConnected

`func (o *VcenterVmHardwareCdromCreateSpec) HasStartConnected() bool`

HasStartConnected returns a boolean if a field has been set.

### GetAllowGuestControl

`func (o *VcenterVmHardwareCdromCreateSpec) GetAllowGuestControl() bool`

GetAllowGuestControl returns the AllowGuestControl field if non-nil, zero value otherwise.

### GetAllowGuestControlOk

`func (o *VcenterVmHardwareCdromCreateSpec) GetAllowGuestControlOk() (*bool, bool)`

GetAllowGuestControlOk returns a tuple with the AllowGuestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuestControl

`func (o *VcenterVmHardwareCdromCreateSpec) SetAllowGuestControl(v bool)`

SetAllowGuestControl sets AllowGuestControl field to given value.

### HasAllowGuestControl

`func (o *VcenterVmHardwareCdromCreateSpec) HasAllowGuestControl() bool`

HasAllowGuestControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


