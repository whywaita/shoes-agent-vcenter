# VcenterVmHardwareAdapterSataInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Device label. | 
**Type** | [**VcenterVmHardwareAdapterSataType**](VcenterVmHardwareAdapterSataType.md) |  | 
**Bus** | **int64** | SATA bus number. | 
**PciSlotNumber** | Pointer to **int64** | Address of the SATA adapter on the PCI bus. May be unset if the virtual machine has never been powered on since the adapter was created. | [optional] 

## Methods

### NewVcenterVmHardwareAdapterSataInfo

`func NewVcenterVmHardwareAdapterSataInfo(label string, type_ VcenterVmHardwareAdapterSataType, bus int64, ) *VcenterVmHardwareAdapterSataInfo`

NewVcenterVmHardwareAdapterSataInfo instantiates a new VcenterVmHardwareAdapterSataInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareAdapterSataInfoWithDefaults

`func NewVcenterVmHardwareAdapterSataInfoWithDefaults() *VcenterVmHardwareAdapterSataInfo`

NewVcenterVmHardwareAdapterSataInfoWithDefaults instantiates a new VcenterVmHardwareAdapterSataInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *VcenterVmHardwareAdapterSataInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VcenterVmHardwareAdapterSataInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VcenterVmHardwareAdapterSataInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *VcenterVmHardwareAdapterSataInfo) GetType() VcenterVmHardwareAdapterSataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareAdapterSataInfo) GetTypeOk() (*VcenterVmHardwareAdapterSataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareAdapterSataInfo) SetType(v VcenterVmHardwareAdapterSataType)`

SetType sets Type field to given value.


### GetBus

`func (o *VcenterVmHardwareAdapterSataInfo) GetBus() int64`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *VcenterVmHardwareAdapterSataInfo) GetBusOk() (*int64, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *VcenterVmHardwareAdapterSataInfo) SetBus(v int64)`

SetBus sets Bus field to given value.


### GetPciSlotNumber

`func (o *VcenterVmHardwareAdapterSataInfo) GetPciSlotNumber() int64`

GetPciSlotNumber returns the PciSlotNumber field if non-nil, zero value otherwise.

### GetPciSlotNumberOk

`func (o *VcenterVmHardwareAdapterSataInfo) GetPciSlotNumberOk() (*int64, bool)`

GetPciSlotNumberOk returns a tuple with the PciSlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlotNumber

`func (o *VcenterVmHardwareAdapterSataInfo) SetPciSlotNumber(v int64)`

SetPciSlotNumber sets PciSlotNumber field to given value.

### HasPciSlotNumber

`func (o *VcenterVmHardwareAdapterSataInfo) HasPciSlotNumber() bool`

HasPciSlotNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


