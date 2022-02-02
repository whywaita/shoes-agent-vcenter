# VcenterVmHardwareAdapterNvmeCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bus** | Pointer to **int64** | NVMe bus number. If unset, the server will choose an available bus number; if none is available, the request will fail. | [optional] 
**PciSlotNumber** | Pointer to **int64** | Address of the NVMe adapter on the PCI bus. If unset, the server will choose an available address when the virtual machine is powered on. | [optional] 

## Methods

### NewVcenterVmHardwareAdapterNvmeCreateSpec

`func NewVcenterVmHardwareAdapterNvmeCreateSpec() *VcenterVmHardwareAdapterNvmeCreateSpec`

NewVcenterVmHardwareAdapterNvmeCreateSpec instantiates a new VcenterVmHardwareAdapterNvmeCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareAdapterNvmeCreateSpecWithDefaults

`func NewVcenterVmHardwareAdapterNvmeCreateSpecWithDefaults() *VcenterVmHardwareAdapterNvmeCreateSpec`

NewVcenterVmHardwareAdapterNvmeCreateSpecWithDefaults instantiates a new VcenterVmHardwareAdapterNvmeCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBus

`func (o *VcenterVmHardwareAdapterNvmeCreateSpec) GetBus() int64`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *VcenterVmHardwareAdapterNvmeCreateSpec) GetBusOk() (*int64, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *VcenterVmHardwareAdapterNvmeCreateSpec) SetBus(v int64)`

SetBus sets Bus field to given value.

### HasBus

`func (o *VcenterVmHardwareAdapterNvmeCreateSpec) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetPciSlotNumber

`func (o *VcenterVmHardwareAdapterNvmeCreateSpec) GetPciSlotNumber() int64`

GetPciSlotNumber returns the PciSlotNumber field if non-nil, zero value otherwise.

### GetPciSlotNumberOk

`func (o *VcenterVmHardwareAdapterNvmeCreateSpec) GetPciSlotNumberOk() (*int64, bool)`

GetPciSlotNumberOk returns a tuple with the PciSlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlotNumber

`func (o *VcenterVmHardwareAdapterNvmeCreateSpec) SetPciSlotNumber(v int64)`

SetPciSlotNumber sets PciSlotNumber field to given value.

### HasPciSlotNumber

`func (o *VcenterVmHardwareAdapterNvmeCreateSpec) HasPciSlotNumber() bool`

HasPciSlotNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


