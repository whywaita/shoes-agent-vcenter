# VcenterVmHardwareAdapterScsiCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**VcenterVmHardwareAdapterScsiType**](VcenterVmHardwareAdapterScsiType.md) |  | [optional] 
**Bus** | Pointer to **int64** | SCSI bus number. If unset, the server will choose an available bus number; if none is available, the request will fail. | [optional] 
**PciSlotNumber** | Pointer to **int64** | Address of the SCSI adapter on the PCI bus. If the PCI address is invalid, the server will change it when the VM is started or as the device is hot added. If unset, the server will choose an available address when the virtual machine is powered on. | [optional] 
**Sharing** | Pointer to [**VcenterVmHardwareAdapterScsiSharing**](VcenterVmHardwareAdapterScsiSharing.md) |  | [optional] 

## Methods

### NewVcenterVmHardwareAdapterScsiCreateSpec

`func NewVcenterVmHardwareAdapterScsiCreateSpec() *VcenterVmHardwareAdapterScsiCreateSpec`

NewVcenterVmHardwareAdapterScsiCreateSpec instantiates a new VcenterVmHardwareAdapterScsiCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareAdapterScsiCreateSpecWithDefaults

`func NewVcenterVmHardwareAdapterScsiCreateSpecWithDefaults() *VcenterVmHardwareAdapterScsiCreateSpec`

NewVcenterVmHardwareAdapterScsiCreateSpecWithDefaults instantiates a new VcenterVmHardwareAdapterScsiCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetType() VcenterVmHardwareAdapterScsiType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetTypeOk() (*VcenterVmHardwareAdapterScsiType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) SetType(v VcenterVmHardwareAdapterScsiType)`

SetType sets Type field to given value.

### HasType

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBus

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetBus() int64`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetBusOk() (*int64, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) SetBus(v int64)`

SetBus sets Bus field to given value.

### HasBus

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetPciSlotNumber

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetPciSlotNumber() int64`

GetPciSlotNumber returns the PciSlotNumber field if non-nil, zero value otherwise.

### GetPciSlotNumberOk

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetPciSlotNumberOk() (*int64, bool)`

GetPciSlotNumberOk returns a tuple with the PciSlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlotNumber

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) SetPciSlotNumber(v int64)`

SetPciSlotNumber sets PciSlotNumber field to given value.

### HasPciSlotNumber

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) HasPciSlotNumber() bool`

HasPciSlotNumber returns a boolean if a field has been set.

### GetSharing

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetSharing() VcenterVmHardwareAdapterScsiSharing`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) GetSharingOk() (*VcenterVmHardwareAdapterScsiSharing, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) SetSharing(v VcenterVmHardwareAdapterScsiSharing)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *VcenterVmHardwareAdapterScsiCreateSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


