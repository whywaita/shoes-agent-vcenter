# VcenterVmHardwareBootDeviceEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmHardwareBootDeviceType**](VcenterVmHardwareBootDeviceType.md) |  | 
**Nic** | Pointer to **string** | Virtual Ethernet device. Ethernet device to use as boot device for this entry. This field is optional and it is only relevant when the value of Device.Entry.type is ETHERNET. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.Ethernet. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.Ethernet. | [optional] 
**Disks** | Pointer to **[]string** | Virtual disk device. List of virtual disks in boot order. This field is optional and it is only relevant when the value of Device.Entry.type is DISK. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.vm.hardware.Disk. | [optional] 

## Methods

### NewVcenterVmHardwareBootDeviceEntry

`func NewVcenterVmHardwareBootDeviceEntry(type_ VcenterVmHardwareBootDeviceType, ) *VcenterVmHardwareBootDeviceEntry`

NewVcenterVmHardwareBootDeviceEntry instantiates a new VcenterVmHardwareBootDeviceEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareBootDeviceEntryWithDefaults

`func NewVcenterVmHardwareBootDeviceEntryWithDefaults() *VcenterVmHardwareBootDeviceEntry`

NewVcenterVmHardwareBootDeviceEntryWithDefaults instantiates a new VcenterVmHardwareBootDeviceEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareBootDeviceEntry) GetType() VcenterVmHardwareBootDeviceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareBootDeviceEntry) GetTypeOk() (*VcenterVmHardwareBootDeviceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareBootDeviceEntry) SetType(v VcenterVmHardwareBootDeviceType)`

SetType sets Type field to given value.


### GetNic

`func (o *VcenterVmHardwareBootDeviceEntry) GetNic() string`

GetNic returns the Nic field if non-nil, zero value otherwise.

### GetNicOk

`func (o *VcenterVmHardwareBootDeviceEntry) GetNicOk() (*string, bool)`

GetNicOk returns a tuple with the Nic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNic

`func (o *VcenterVmHardwareBootDeviceEntry) SetNic(v string)`

SetNic sets Nic field to given value.

### HasNic

`func (o *VcenterVmHardwareBootDeviceEntry) HasNic() bool`

HasNic returns a boolean if a field has been set.

### GetDisks

`func (o *VcenterVmHardwareBootDeviceEntry) GetDisks() []string`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VcenterVmHardwareBootDeviceEntry) GetDisksOk() (*[]string, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VcenterVmHardwareBootDeviceEntry) SetDisks(v []string)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *VcenterVmHardwareBootDeviceEntry) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


