# VcenterVmTemplateLibraryItemsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuestOS** | [**VcenterVmGuestOS**](VcenterVmGuestOS.md) |  | 
**Cpu** | [**VcenterVmTemplateLibraryItemsCpuInfo**](VcenterVmTemplateLibraryItemsCpuInfo.md) |  | 
**Memory** | [**VcenterVmTemplateLibraryItemsMemoryInfo**](VcenterVmTemplateLibraryItemsMemoryInfo.md) |  | 
**VmHomeStorage** | [**VcenterVmTemplateLibraryItemsVmHomeStorageInfo**](VcenterVmTemplateLibraryItemsVmHomeStorageInfo.md) |  | 
**Disks** | [**[]VcenterVmTemplateLibraryItemsInfoDisks**](VcenterVmTemplateLibraryItemsInfoDisks.md) | Storage information about the virtual machine template&#39;s virtual disks. | 
**Nics** | [**[]VcenterVmTemplateLibraryItemsInfoNics**](VcenterVmTemplateLibraryItemsInfoNics.md) | Information about the virtual machine template&#39;s virtual ethernet adapters. | 
**VmTemplate** | **string** | Identifier of the latest virtual machine template contained in the library item. This {@term field} is the managed object identifier used to identify the virtual machine template in the vSphere Management (SOAP) API. | 

## Methods

### NewVcenterVmTemplateLibraryItemsInfo

`func NewVcenterVmTemplateLibraryItemsInfo(guestOS VcenterVmGuestOS, cpu VcenterVmTemplateLibraryItemsCpuInfo, memory VcenterVmTemplateLibraryItemsMemoryInfo, vmHomeStorage VcenterVmTemplateLibraryItemsVmHomeStorageInfo, disks []VcenterVmTemplateLibraryItemsInfoDisks, nics []VcenterVmTemplateLibraryItemsInfoNics, vmTemplate string, ) *VcenterVmTemplateLibraryItemsInfo`

NewVcenterVmTemplateLibraryItemsInfo instantiates a new VcenterVmTemplateLibraryItemsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsInfoWithDefaults

`func NewVcenterVmTemplateLibraryItemsInfoWithDefaults() *VcenterVmTemplateLibraryItemsInfo`

NewVcenterVmTemplateLibraryItemsInfoWithDefaults instantiates a new VcenterVmTemplateLibraryItemsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuestOS

`func (o *VcenterVmTemplateLibraryItemsInfo) GetGuestOS() VcenterVmGuestOS`

GetGuestOS returns the GuestOS field if non-nil, zero value otherwise.

### GetGuestOSOk

`func (o *VcenterVmTemplateLibraryItemsInfo) GetGuestOSOk() (*VcenterVmGuestOS, bool)`

GetGuestOSOk returns a tuple with the GuestOS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestOS

`func (o *VcenterVmTemplateLibraryItemsInfo) SetGuestOS(v VcenterVmGuestOS)`

SetGuestOS sets GuestOS field to given value.


### GetCpu

`func (o *VcenterVmTemplateLibraryItemsInfo) GetCpu() VcenterVmTemplateLibraryItemsCpuInfo`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *VcenterVmTemplateLibraryItemsInfo) GetCpuOk() (*VcenterVmTemplateLibraryItemsCpuInfo, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *VcenterVmTemplateLibraryItemsInfo) SetCpu(v VcenterVmTemplateLibraryItemsCpuInfo)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *VcenterVmTemplateLibraryItemsInfo) GetMemory() VcenterVmTemplateLibraryItemsMemoryInfo`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VcenterVmTemplateLibraryItemsInfo) GetMemoryOk() (*VcenterVmTemplateLibraryItemsMemoryInfo, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VcenterVmTemplateLibraryItemsInfo) SetMemory(v VcenterVmTemplateLibraryItemsMemoryInfo)`

SetMemory sets Memory field to given value.


### GetVmHomeStorage

`func (o *VcenterVmTemplateLibraryItemsInfo) GetVmHomeStorage() VcenterVmTemplateLibraryItemsVmHomeStorageInfo`

GetVmHomeStorage returns the VmHomeStorage field if non-nil, zero value otherwise.

### GetVmHomeStorageOk

`func (o *VcenterVmTemplateLibraryItemsInfo) GetVmHomeStorageOk() (*VcenterVmTemplateLibraryItemsVmHomeStorageInfo, bool)`

GetVmHomeStorageOk returns a tuple with the VmHomeStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHomeStorage

`func (o *VcenterVmTemplateLibraryItemsInfo) SetVmHomeStorage(v VcenterVmTemplateLibraryItemsVmHomeStorageInfo)`

SetVmHomeStorage sets VmHomeStorage field to given value.


### GetDisks

`func (o *VcenterVmTemplateLibraryItemsInfo) GetDisks() []VcenterVmTemplateLibraryItemsInfoDisks`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VcenterVmTemplateLibraryItemsInfo) GetDisksOk() (*[]VcenterVmTemplateLibraryItemsInfoDisks, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VcenterVmTemplateLibraryItemsInfo) SetDisks(v []VcenterVmTemplateLibraryItemsInfoDisks)`

SetDisks sets Disks field to given value.


### GetNics

`func (o *VcenterVmTemplateLibraryItemsInfo) GetNics() []VcenterVmTemplateLibraryItemsInfoNics`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *VcenterVmTemplateLibraryItemsInfo) GetNicsOk() (*[]VcenterVmTemplateLibraryItemsInfoNics, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *VcenterVmTemplateLibraryItemsInfo) SetNics(v []VcenterVmTemplateLibraryItemsInfoNics)`

SetNics sets Nics field to given value.


### GetVmTemplate

`func (o *VcenterVmTemplateLibraryItemsInfo) GetVmTemplate() string`

GetVmTemplate returns the VmTemplate field if non-nil, zero value otherwise.

### GetVmTemplateOk

`func (o *VcenterVmTemplateLibraryItemsInfo) GetVmTemplateOk() (*string, bool)`

GetVmTemplateOk returns a tuple with the VmTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplate

`func (o *VcenterVmTemplateLibraryItemsInfo) SetVmTemplate(v string)`

SetVmTemplate sets VmTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


