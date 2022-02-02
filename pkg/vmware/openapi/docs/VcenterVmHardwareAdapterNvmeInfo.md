# VcenterVmHardwareAdapterNvmeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Device label. | 
**Bus** | **int64** | NVMe bus number. | 
**PciSlotNumber** | Pointer to **int64** | Address of the NVMe adapter on the PCI bus. May be unset if the virtual machine has never been powered on since the adapter was created. | [optional] 

## Methods

### NewVcenterVmHardwareAdapterNvmeInfo

`func NewVcenterVmHardwareAdapterNvmeInfo(label string, bus int64, ) *VcenterVmHardwareAdapterNvmeInfo`

NewVcenterVmHardwareAdapterNvmeInfo instantiates a new VcenterVmHardwareAdapterNvmeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareAdapterNvmeInfoWithDefaults

`func NewVcenterVmHardwareAdapterNvmeInfoWithDefaults() *VcenterVmHardwareAdapterNvmeInfo`

NewVcenterVmHardwareAdapterNvmeInfoWithDefaults instantiates a new VcenterVmHardwareAdapterNvmeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *VcenterVmHardwareAdapterNvmeInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VcenterVmHardwareAdapterNvmeInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VcenterVmHardwareAdapterNvmeInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetBus

`func (o *VcenterVmHardwareAdapterNvmeInfo) GetBus() int64`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *VcenterVmHardwareAdapterNvmeInfo) GetBusOk() (*int64, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *VcenterVmHardwareAdapterNvmeInfo) SetBus(v int64)`

SetBus sets Bus field to given value.


### GetPciSlotNumber

`func (o *VcenterVmHardwareAdapterNvmeInfo) GetPciSlotNumber() int64`

GetPciSlotNumber returns the PciSlotNumber field if non-nil, zero value otherwise.

### GetPciSlotNumberOk

`func (o *VcenterVmHardwareAdapterNvmeInfo) GetPciSlotNumberOk() (*int64, bool)`

GetPciSlotNumberOk returns a tuple with the PciSlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlotNumber

`func (o *VcenterVmHardwareAdapterNvmeInfo) SetPciSlotNumber(v int64)`

SetPciSlotNumber sets PciSlotNumber field to given value.

### HasPciSlotNumber

`func (o *VcenterVmHardwareAdapterNvmeInfo) HasPciSlotNumber() bool`

HasPciSlotNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


