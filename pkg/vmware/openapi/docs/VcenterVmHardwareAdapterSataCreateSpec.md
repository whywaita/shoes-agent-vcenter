# VcenterVmHardwareAdapterSataCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**VcenterVmHardwareAdapterSataType**](VcenterVmHardwareAdapterSataType.md) |  | [optional] 
**Bus** | Pointer to **int64** | SATA bus number. If unset, the server will choose an available bus number; if none is available, the request will fail. | [optional] 
**PciSlotNumber** | Pointer to **int64** | Address of the SATA adapter on the PCI bus. If unset, the server will choose an available address when the virtual machine is powered on. | [optional] 

## Methods

### NewVcenterVmHardwareAdapterSataCreateSpec

`func NewVcenterVmHardwareAdapterSataCreateSpec() *VcenterVmHardwareAdapterSataCreateSpec`

NewVcenterVmHardwareAdapterSataCreateSpec instantiates a new VcenterVmHardwareAdapterSataCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareAdapterSataCreateSpecWithDefaults

`func NewVcenterVmHardwareAdapterSataCreateSpecWithDefaults() *VcenterVmHardwareAdapterSataCreateSpec`

NewVcenterVmHardwareAdapterSataCreateSpecWithDefaults instantiates a new VcenterVmHardwareAdapterSataCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareAdapterSataCreateSpec) GetType() VcenterVmHardwareAdapterSataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareAdapterSataCreateSpec) GetTypeOk() (*VcenterVmHardwareAdapterSataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareAdapterSataCreateSpec) SetType(v VcenterVmHardwareAdapterSataType)`

SetType sets Type field to given value.

### HasType

`func (o *VcenterVmHardwareAdapterSataCreateSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBus

`func (o *VcenterVmHardwareAdapterSataCreateSpec) GetBus() int64`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *VcenterVmHardwareAdapterSataCreateSpec) GetBusOk() (*int64, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *VcenterVmHardwareAdapterSataCreateSpec) SetBus(v int64)`

SetBus sets Bus field to given value.

### HasBus

`func (o *VcenterVmHardwareAdapterSataCreateSpec) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetPciSlotNumber

`func (o *VcenterVmHardwareAdapterSataCreateSpec) GetPciSlotNumber() int64`

GetPciSlotNumber returns the PciSlotNumber field if non-nil, zero value otherwise.

### GetPciSlotNumberOk

`func (o *VcenterVmHardwareAdapterSataCreateSpec) GetPciSlotNumberOk() (*int64, bool)`

GetPciSlotNumberOk returns a tuple with the PciSlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlotNumber

`func (o *VcenterVmHardwareAdapterSataCreateSpec) SetPciSlotNumber(v int64)`

SetPciSlotNumber sets PciSlotNumber field to given value.

### HasPciSlotNumber

`func (o *VcenterVmHardwareAdapterSataCreateSpec) HasPciSlotNumber() bool`

HasPciSlotNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


