# VcenterVmHardwareCdromInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmHardwareCdromHostBusAdapterType**](VcenterVmHardwareCdromHostBusAdapterType.md) |  | 
**Label** | **string** | Device label. | 
**Ide** | Pointer to [**VcenterVmHardwareIdeAddressInfo**](VcenterVmHardwareIdeAddressInfo.md) |  | [optional] 
**Sata** | Pointer to [**VcenterVmHardwareSataAddressInfo**](VcenterVmHardwareSataAddressInfo.md) |  | [optional] 
**Backing** | [**VcenterVmHardwareCdromBackingInfo**](VcenterVmHardwareCdromBackingInfo.md) |  | 
**State** | [**VcenterVmHardwareConnectionState**](VcenterVmHardwareConnectionState.md) |  | 
**StartConnected** | **bool** | Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. | 
**AllowGuestControl** | **bool** | Flag indicating whether the guest can connect and disconnect the device. | 

## Methods

### NewVcenterVmHardwareCdromInfo

`func NewVcenterVmHardwareCdromInfo(type_ VcenterVmHardwareCdromHostBusAdapterType, label string, backing VcenterVmHardwareCdromBackingInfo, state VcenterVmHardwareConnectionState, startConnected bool, allowGuestControl bool, ) *VcenterVmHardwareCdromInfo`

NewVcenterVmHardwareCdromInfo instantiates a new VcenterVmHardwareCdromInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareCdromInfoWithDefaults

`func NewVcenterVmHardwareCdromInfoWithDefaults() *VcenterVmHardwareCdromInfo`

NewVcenterVmHardwareCdromInfoWithDefaults instantiates a new VcenterVmHardwareCdromInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareCdromInfo) GetType() VcenterVmHardwareCdromHostBusAdapterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareCdromInfo) GetTypeOk() (*VcenterVmHardwareCdromHostBusAdapterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareCdromInfo) SetType(v VcenterVmHardwareCdromHostBusAdapterType)`

SetType sets Type field to given value.


### GetLabel

`func (o *VcenterVmHardwareCdromInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VcenterVmHardwareCdromInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VcenterVmHardwareCdromInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetIde

`func (o *VcenterVmHardwareCdromInfo) GetIde() VcenterVmHardwareIdeAddressInfo`

GetIde returns the Ide field if non-nil, zero value otherwise.

### GetIdeOk

`func (o *VcenterVmHardwareCdromInfo) GetIdeOk() (*VcenterVmHardwareIdeAddressInfo, bool)`

GetIdeOk returns a tuple with the Ide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde

`func (o *VcenterVmHardwareCdromInfo) SetIde(v VcenterVmHardwareIdeAddressInfo)`

SetIde sets Ide field to given value.

### HasIde

`func (o *VcenterVmHardwareCdromInfo) HasIde() bool`

HasIde returns a boolean if a field has been set.

### GetSata

`func (o *VcenterVmHardwareCdromInfo) GetSata() VcenterVmHardwareSataAddressInfo`

GetSata returns the Sata field if non-nil, zero value otherwise.

### GetSataOk

`func (o *VcenterVmHardwareCdromInfo) GetSataOk() (*VcenterVmHardwareSataAddressInfo, bool)`

GetSataOk returns a tuple with the Sata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata

`func (o *VcenterVmHardwareCdromInfo) SetSata(v VcenterVmHardwareSataAddressInfo)`

SetSata sets Sata field to given value.

### HasSata

`func (o *VcenterVmHardwareCdromInfo) HasSata() bool`

HasSata returns a boolean if a field has been set.

### GetBacking

`func (o *VcenterVmHardwareCdromInfo) GetBacking() VcenterVmHardwareCdromBackingInfo`

GetBacking returns the Backing field if non-nil, zero value otherwise.

### GetBackingOk

`func (o *VcenterVmHardwareCdromInfo) GetBackingOk() (*VcenterVmHardwareCdromBackingInfo, bool)`

GetBackingOk returns a tuple with the Backing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacking

`func (o *VcenterVmHardwareCdromInfo) SetBacking(v VcenterVmHardwareCdromBackingInfo)`

SetBacking sets Backing field to given value.


### GetState

`func (o *VcenterVmHardwareCdromInfo) GetState() VcenterVmHardwareConnectionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterVmHardwareCdromInfo) GetStateOk() (*VcenterVmHardwareConnectionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterVmHardwareCdromInfo) SetState(v VcenterVmHardwareConnectionState)`

SetState sets State field to given value.


### GetStartConnected

`func (o *VcenterVmHardwareCdromInfo) GetStartConnected() bool`

GetStartConnected returns the StartConnected field if non-nil, zero value otherwise.

### GetStartConnectedOk

`func (o *VcenterVmHardwareCdromInfo) GetStartConnectedOk() (*bool, bool)`

GetStartConnectedOk returns a tuple with the StartConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConnected

`func (o *VcenterVmHardwareCdromInfo) SetStartConnected(v bool)`

SetStartConnected sets StartConnected field to given value.


### GetAllowGuestControl

`func (o *VcenterVmHardwareCdromInfo) GetAllowGuestControl() bool`

GetAllowGuestControl returns the AllowGuestControl field if non-nil, zero value otherwise.

### GetAllowGuestControlOk

`func (o *VcenterVmHardwareCdromInfo) GetAllowGuestControlOk() (*bool, bool)`

GetAllowGuestControlOk returns a tuple with the AllowGuestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuestControl

`func (o *VcenterVmHardwareCdromInfo) SetAllowGuestControl(v bool)`

SetAllowGuestControl sets AllowGuestControl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


