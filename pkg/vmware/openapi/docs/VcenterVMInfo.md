# VcenterVMInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuestOS** | [**VcenterVmGuestOS**](VcenterVmGuestOS.md) |  | 
**Name** | **string** | Virtual machine name. | 
**Identity** | Pointer to [**VcenterVmIdentityInfo**](VcenterVmIdentityInfo.md) |  | [optional] 
**PowerState** | [**VcenterVmPowerState**](VcenterVmPowerState.md) |  | 
**InstantCloneFrozen** | Pointer to **bool** | Indicates whether the virtual machine is frozen for instant clone, or not. This field is optional because it was added in a newer version than its parent node. | [optional] 
**Hardware** | [**VcenterVmHardwareInfo**](VcenterVmHardwareInfo.md) |  | 
**Boot** | [**VcenterVmHardwareBootInfo**](VcenterVmHardwareBootInfo.md) |  | 
**BootDevices** | [**[]VcenterVmHardwareBootDeviceEntry**](VcenterVmHardwareBootDeviceEntry.md) | Boot device configuration. If the list has no entries, a server-specific default boot sequence is used. | 
**Cpu** | [**VcenterVmHardwareCpuInfo**](VcenterVmHardwareCpuInfo.md) |  | 
**Memory** | [**VcenterVmHardwareMemoryInfo**](VcenterVmHardwareMemoryInfo.md) |  | 
**Disks** | [**[]VcenterVMInfoDisks**](VcenterVMInfoDisks.md) | List of disks. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Disk. | 
**Nics** | [**[]VcenterVMInfoNics**](VcenterVMInfoNics.md) | List of Ethernet adapters. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Ethernet. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Ethernet. | 
**Cdroms** | [**[]VcenterVMInfoCdroms**](VcenterVMInfoCdroms.md) | List of CD-ROMs. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Cdrom. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Cdrom. | 
**Floppies** | [**[]VcenterVMInfoFloppies**](VcenterVMInfoFloppies.md) | List of floppy drives. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Floppy. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Floppy. | 
**ParallelPorts** | [**[]VcenterVMInfoParallelPorts**](VcenterVMInfoParallelPorts.md) | List of parallel ports. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.ParallelPort. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.ParallelPort. | 
**SerialPorts** | [**[]VcenterVMInfoSerialPorts**](VcenterVMInfoSerialPorts.md) | List of serial ports. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.SerialPort. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.SerialPort. | 
**SataAdapters** | [**[]VcenterVMInfoSataAdapters**](VcenterVMInfoSataAdapters.md) | List of SATA adapters. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.SataAdapter. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.SataAdapter. | 
**ScsiAdapters** | [**[]VcenterVMInfoScsiAdapters**](VcenterVMInfoScsiAdapters.md) | List of SCSI adapters. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.ScsiAdapter. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.ScsiAdapter. | 
**NvmeAdapters** | Pointer to [**[]VcenterVMInfoNvmeAdapters**](VcenterVMInfoNvmeAdapters.md) | List of NVMe adapters. This field is optional because it was added in a newer version than its parent node. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.NvmeAdapter. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.NvmeAdapter. | [optional] 

## Methods

### NewVcenterVMInfo

`func NewVcenterVMInfo(guestOS VcenterVmGuestOS, name string, powerState VcenterVmPowerState, hardware VcenterVmHardwareInfo, boot VcenterVmHardwareBootInfo, bootDevices []VcenterVmHardwareBootDeviceEntry, cpu VcenterVmHardwareCpuInfo, memory VcenterVmHardwareMemoryInfo, disks []VcenterVMInfoDisks, nics []VcenterVMInfoNics, cdroms []VcenterVMInfoCdroms, floppies []VcenterVMInfoFloppies, parallelPorts []VcenterVMInfoParallelPorts, serialPorts []VcenterVMInfoSerialPorts, sataAdapters []VcenterVMInfoSataAdapters, scsiAdapters []VcenterVMInfoScsiAdapters, ) *VcenterVMInfo`

NewVcenterVMInfo instantiates a new VcenterVMInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMInfoWithDefaults

`func NewVcenterVMInfoWithDefaults() *VcenterVMInfo`

NewVcenterVMInfoWithDefaults instantiates a new VcenterVMInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuestOS

`func (o *VcenterVMInfo) GetGuestOS() VcenterVmGuestOS`

GetGuestOS returns the GuestOS field if non-nil, zero value otherwise.

### GetGuestOSOk

`func (o *VcenterVMInfo) GetGuestOSOk() (*VcenterVmGuestOS, bool)`

GetGuestOSOk returns a tuple with the GuestOS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestOS

`func (o *VcenterVMInfo) SetGuestOS(v VcenterVmGuestOS)`

SetGuestOS sets GuestOS field to given value.


### GetName

`func (o *VcenterVMInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVMInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVMInfo) SetName(v string)`

SetName sets Name field to given value.


### GetIdentity

`func (o *VcenterVMInfo) GetIdentity() VcenterVmIdentityInfo`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VcenterVMInfo) GetIdentityOk() (*VcenterVmIdentityInfo, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VcenterVMInfo) SetIdentity(v VcenterVmIdentityInfo)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VcenterVMInfo) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetPowerState

`func (o *VcenterVMInfo) GetPowerState() VcenterVmPowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *VcenterVMInfo) GetPowerStateOk() (*VcenterVmPowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *VcenterVMInfo) SetPowerState(v VcenterVmPowerState)`

SetPowerState sets PowerState field to given value.


### GetInstantCloneFrozen

`func (o *VcenterVMInfo) GetInstantCloneFrozen() bool`

GetInstantCloneFrozen returns the InstantCloneFrozen field if non-nil, zero value otherwise.

### GetInstantCloneFrozenOk

`func (o *VcenterVMInfo) GetInstantCloneFrozenOk() (*bool, bool)`

GetInstantCloneFrozenOk returns a tuple with the InstantCloneFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneFrozen

`func (o *VcenterVMInfo) SetInstantCloneFrozen(v bool)`

SetInstantCloneFrozen sets InstantCloneFrozen field to given value.

### HasInstantCloneFrozen

`func (o *VcenterVMInfo) HasInstantCloneFrozen() bool`

HasInstantCloneFrozen returns a boolean if a field has been set.

### GetHardware

`func (o *VcenterVMInfo) GetHardware() VcenterVmHardwareInfo`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *VcenterVMInfo) GetHardwareOk() (*VcenterVmHardwareInfo, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *VcenterVMInfo) SetHardware(v VcenterVmHardwareInfo)`

SetHardware sets Hardware field to given value.


### GetBoot

`func (o *VcenterVMInfo) GetBoot() VcenterVmHardwareBootInfo`

GetBoot returns the Boot field if non-nil, zero value otherwise.

### GetBootOk

`func (o *VcenterVMInfo) GetBootOk() (*VcenterVmHardwareBootInfo, bool)`

GetBootOk returns a tuple with the Boot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoot

`func (o *VcenterVMInfo) SetBoot(v VcenterVmHardwareBootInfo)`

SetBoot sets Boot field to given value.


### GetBootDevices

`func (o *VcenterVMInfo) GetBootDevices() []VcenterVmHardwareBootDeviceEntry`

GetBootDevices returns the BootDevices field if non-nil, zero value otherwise.

### GetBootDevicesOk

`func (o *VcenterVMInfo) GetBootDevicesOk() (*[]VcenterVmHardwareBootDeviceEntry, bool)`

GetBootDevicesOk returns a tuple with the BootDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDevices

`func (o *VcenterVMInfo) SetBootDevices(v []VcenterVmHardwareBootDeviceEntry)`

SetBootDevices sets BootDevices field to given value.


### GetCpu

`func (o *VcenterVMInfo) GetCpu() VcenterVmHardwareCpuInfo`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *VcenterVMInfo) GetCpuOk() (*VcenterVmHardwareCpuInfo, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *VcenterVMInfo) SetCpu(v VcenterVmHardwareCpuInfo)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *VcenterVMInfo) GetMemory() VcenterVmHardwareMemoryInfo`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VcenterVMInfo) GetMemoryOk() (*VcenterVmHardwareMemoryInfo, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VcenterVMInfo) SetMemory(v VcenterVmHardwareMemoryInfo)`

SetMemory sets Memory field to given value.


### GetDisks

`func (o *VcenterVMInfo) GetDisks() []VcenterVMInfoDisks`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VcenterVMInfo) GetDisksOk() (*[]VcenterVMInfoDisks, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VcenterVMInfo) SetDisks(v []VcenterVMInfoDisks)`

SetDisks sets Disks field to given value.


### GetNics

`func (o *VcenterVMInfo) GetNics() []VcenterVMInfoNics`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *VcenterVMInfo) GetNicsOk() (*[]VcenterVMInfoNics, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *VcenterVMInfo) SetNics(v []VcenterVMInfoNics)`

SetNics sets Nics field to given value.


### GetCdroms

`func (o *VcenterVMInfo) GetCdroms() []VcenterVMInfoCdroms`

GetCdroms returns the Cdroms field if non-nil, zero value otherwise.

### GetCdromsOk

`func (o *VcenterVMInfo) GetCdromsOk() (*[]VcenterVMInfoCdroms, bool)`

GetCdromsOk returns a tuple with the Cdroms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdroms

`func (o *VcenterVMInfo) SetCdroms(v []VcenterVMInfoCdroms)`

SetCdroms sets Cdroms field to given value.


### GetFloppies

`func (o *VcenterVMInfo) GetFloppies() []VcenterVMInfoFloppies`

GetFloppies returns the Floppies field if non-nil, zero value otherwise.

### GetFloppiesOk

`func (o *VcenterVMInfo) GetFloppiesOk() (*[]VcenterVMInfoFloppies, bool)`

GetFloppiesOk returns a tuple with the Floppies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloppies

`func (o *VcenterVMInfo) SetFloppies(v []VcenterVMInfoFloppies)`

SetFloppies sets Floppies field to given value.


### GetParallelPorts

`func (o *VcenterVMInfo) GetParallelPorts() []VcenterVMInfoParallelPorts`

GetParallelPorts returns the ParallelPorts field if non-nil, zero value otherwise.

### GetParallelPortsOk

`func (o *VcenterVMInfo) GetParallelPortsOk() (*[]VcenterVMInfoParallelPorts, bool)`

GetParallelPortsOk returns a tuple with the ParallelPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelPorts

`func (o *VcenterVMInfo) SetParallelPorts(v []VcenterVMInfoParallelPorts)`

SetParallelPorts sets ParallelPorts field to given value.


### GetSerialPorts

`func (o *VcenterVMInfo) GetSerialPorts() []VcenterVMInfoSerialPorts`

GetSerialPorts returns the SerialPorts field if non-nil, zero value otherwise.

### GetSerialPortsOk

`func (o *VcenterVMInfo) GetSerialPortsOk() (*[]VcenterVMInfoSerialPorts, bool)`

GetSerialPortsOk returns a tuple with the SerialPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialPorts

`func (o *VcenterVMInfo) SetSerialPorts(v []VcenterVMInfoSerialPorts)`

SetSerialPorts sets SerialPorts field to given value.


### GetSataAdapters

`func (o *VcenterVMInfo) GetSataAdapters() []VcenterVMInfoSataAdapters`

GetSataAdapters returns the SataAdapters field if non-nil, zero value otherwise.

### GetSataAdaptersOk

`func (o *VcenterVMInfo) GetSataAdaptersOk() (*[]VcenterVMInfoSataAdapters, bool)`

GetSataAdaptersOk returns a tuple with the SataAdapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSataAdapters

`func (o *VcenterVMInfo) SetSataAdapters(v []VcenterVMInfoSataAdapters)`

SetSataAdapters sets SataAdapters field to given value.


### GetScsiAdapters

`func (o *VcenterVMInfo) GetScsiAdapters() []VcenterVMInfoScsiAdapters`

GetScsiAdapters returns the ScsiAdapters field if non-nil, zero value otherwise.

### GetScsiAdaptersOk

`func (o *VcenterVMInfo) GetScsiAdaptersOk() (*[]VcenterVMInfoScsiAdapters, bool)`

GetScsiAdaptersOk returns a tuple with the ScsiAdapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsiAdapters

`func (o *VcenterVMInfo) SetScsiAdapters(v []VcenterVMInfoScsiAdapters)`

SetScsiAdapters sets ScsiAdapters field to given value.


### GetNvmeAdapters

`func (o *VcenterVMInfo) GetNvmeAdapters() []VcenterVMInfoNvmeAdapters`

GetNvmeAdapters returns the NvmeAdapters field if non-nil, zero value otherwise.

### GetNvmeAdaptersOk

`func (o *VcenterVMInfo) GetNvmeAdaptersOk() (*[]VcenterVMInfoNvmeAdapters, bool)`

GetNvmeAdaptersOk returns a tuple with the NvmeAdapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvmeAdapters

`func (o *VcenterVMInfo) SetNvmeAdapters(v []VcenterVMInfoNvmeAdapters)`

SetNvmeAdapters sets NvmeAdapters field to given value.

### HasNvmeAdapters

`func (o *VcenterVMInfo) HasNvmeAdapters() bool`

HasNvmeAdapters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


