# VcenterVMInstantCloneSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Virtual machine to InstantClone from. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: VirtualMachine. When operations return a value of this structure as a result, the field will be an identifier for the resource type: VirtualMachine. | 
**Name** | **string** | Name of the new virtual machine. | 
**Placement** | Pointer to [**VcenterVMInstantClonePlacementSpec**](VcenterVMInstantClonePlacementSpec.md) |  | [optional] 
**NicsToUpdate** | Pointer to [**[]VcenterVMInstantCloneSpecNicsToUpdate**](VcenterVMInstantCloneSpecNicsToUpdate.md) | Map of NICs to update. If unset, no NICs will be updated. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Ethernet. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Ethernet. | [optional] 
**DisconnectAllNics** | Pointer to **bool** | Indicates whether all NICs on the destination virtual machine should be disconnected from the newtwork If unset, connection status of all NICs on the destination virtual machine will be the same as on the source virtual machine. | [optional] 
**ParallelPortsToUpdate** | Pointer to [**[]VcenterVMInstantCloneSpecParallelPortsToUpdate**](VcenterVMInstantCloneSpecParallelPortsToUpdate.md) | Map of parallel ports to Update. If unset, no parallel ports will be updated. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.ParallelPort. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.ParallelPort. | [optional] 
**SerialPortsToUpdate** | Pointer to [**[]VcenterVMInstantCloneSpecSerialPortsToUpdate**](VcenterVMInstantCloneSpecSerialPortsToUpdate.md) | Map of serial ports to Update. If unset, no serial ports will be updated. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.SerialPort. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.SerialPort. | [optional] 
**BiosUuid** | Pointer to **string** | 128-bit SMBIOS UUID of a virtual machine represented as a hexadecimal string in \&quot;12345678-abcd-1234-cdef-123456789abc\&quot; format. If unset, will be generated. | [optional] 

## Methods

### NewVcenterVMInstantCloneSpec

`func NewVcenterVMInstantCloneSpec(source string, name string, ) *VcenterVMInstantCloneSpec`

NewVcenterVMInstantCloneSpec instantiates a new VcenterVMInstantCloneSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMInstantCloneSpecWithDefaults

`func NewVcenterVMInstantCloneSpecWithDefaults() *VcenterVMInstantCloneSpec`

NewVcenterVMInstantCloneSpecWithDefaults instantiates a new VcenterVMInstantCloneSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *VcenterVMInstantCloneSpec) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VcenterVMInstantCloneSpec) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VcenterVMInstantCloneSpec) SetSource(v string)`

SetSource sets Source field to given value.


### GetName

`func (o *VcenterVMInstantCloneSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVMInstantCloneSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVMInstantCloneSpec) SetName(v string)`

SetName sets Name field to given value.


### GetPlacement

`func (o *VcenterVMInstantCloneSpec) GetPlacement() VcenterVMInstantClonePlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVMInstantCloneSpec) GetPlacementOk() (*VcenterVMInstantClonePlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVMInstantCloneSpec) SetPlacement(v VcenterVMInstantClonePlacementSpec)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VcenterVMInstantCloneSpec) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetNicsToUpdate

`func (o *VcenterVMInstantCloneSpec) GetNicsToUpdate() []VcenterVMInstantCloneSpecNicsToUpdate`

GetNicsToUpdate returns the NicsToUpdate field if non-nil, zero value otherwise.

### GetNicsToUpdateOk

`func (o *VcenterVMInstantCloneSpec) GetNicsToUpdateOk() (*[]VcenterVMInstantCloneSpecNicsToUpdate, bool)`

GetNicsToUpdateOk returns a tuple with the NicsToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicsToUpdate

`func (o *VcenterVMInstantCloneSpec) SetNicsToUpdate(v []VcenterVMInstantCloneSpecNicsToUpdate)`

SetNicsToUpdate sets NicsToUpdate field to given value.

### HasNicsToUpdate

`func (o *VcenterVMInstantCloneSpec) HasNicsToUpdate() bool`

HasNicsToUpdate returns a boolean if a field has been set.

### GetDisconnectAllNics

`func (o *VcenterVMInstantCloneSpec) GetDisconnectAllNics() bool`

GetDisconnectAllNics returns the DisconnectAllNics field if non-nil, zero value otherwise.

### GetDisconnectAllNicsOk

`func (o *VcenterVMInstantCloneSpec) GetDisconnectAllNicsOk() (*bool, bool)`

GetDisconnectAllNicsOk returns a tuple with the DisconnectAllNics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectAllNics

`func (o *VcenterVMInstantCloneSpec) SetDisconnectAllNics(v bool)`

SetDisconnectAllNics sets DisconnectAllNics field to given value.

### HasDisconnectAllNics

`func (o *VcenterVMInstantCloneSpec) HasDisconnectAllNics() bool`

HasDisconnectAllNics returns a boolean if a field has been set.

### GetParallelPortsToUpdate

`func (o *VcenterVMInstantCloneSpec) GetParallelPortsToUpdate() []VcenterVMInstantCloneSpecParallelPortsToUpdate`

GetParallelPortsToUpdate returns the ParallelPortsToUpdate field if non-nil, zero value otherwise.

### GetParallelPortsToUpdateOk

`func (o *VcenterVMInstantCloneSpec) GetParallelPortsToUpdateOk() (*[]VcenterVMInstantCloneSpecParallelPortsToUpdate, bool)`

GetParallelPortsToUpdateOk returns a tuple with the ParallelPortsToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelPortsToUpdate

`func (o *VcenterVMInstantCloneSpec) SetParallelPortsToUpdate(v []VcenterVMInstantCloneSpecParallelPortsToUpdate)`

SetParallelPortsToUpdate sets ParallelPortsToUpdate field to given value.

### HasParallelPortsToUpdate

`func (o *VcenterVMInstantCloneSpec) HasParallelPortsToUpdate() bool`

HasParallelPortsToUpdate returns a boolean if a field has been set.

### GetSerialPortsToUpdate

`func (o *VcenterVMInstantCloneSpec) GetSerialPortsToUpdate() []VcenterVMInstantCloneSpecSerialPortsToUpdate`

GetSerialPortsToUpdate returns the SerialPortsToUpdate field if non-nil, zero value otherwise.

### GetSerialPortsToUpdateOk

`func (o *VcenterVMInstantCloneSpec) GetSerialPortsToUpdateOk() (*[]VcenterVMInstantCloneSpecSerialPortsToUpdate, bool)`

GetSerialPortsToUpdateOk returns a tuple with the SerialPortsToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialPortsToUpdate

`func (o *VcenterVMInstantCloneSpec) SetSerialPortsToUpdate(v []VcenterVMInstantCloneSpecSerialPortsToUpdate)`

SetSerialPortsToUpdate sets SerialPortsToUpdate field to given value.

### HasSerialPortsToUpdate

`func (o *VcenterVMInstantCloneSpec) HasSerialPortsToUpdate() bool`

HasSerialPortsToUpdate returns a boolean if a field has been set.

### GetBiosUuid

`func (o *VcenterVMInstantCloneSpec) GetBiosUuid() string`

GetBiosUuid returns the BiosUuid field if non-nil, zero value otherwise.

### GetBiosUuidOk

`func (o *VcenterVMInstantCloneSpec) GetBiosUuidOk() (*string, bool)`

GetBiosUuidOk returns a tuple with the BiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUuid

`func (o *VcenterVMInstantCloneSpec) SetBiosUuid(v string)`

SetBiosUuid sets BiosUuid field to given value.

### HasBiosUuid

`func (o *VcenterVMInstantCloneSpec) HasBiosUuid() bool`

HasBiosUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


