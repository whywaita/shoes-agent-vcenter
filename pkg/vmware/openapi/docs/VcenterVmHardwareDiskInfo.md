# VcenterVmHardwareDiskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Device label. | 
**Type** | [**VcenterVmHardwareDiskHostBusAdapterType**](VcenterVmHardwareDiskHostBusAdapterType.md) |  | 
**Ide** | Pointer to [**VcenterVmHardwareIdeAddressInfo**](VcenterVmHardwareIdeAddressInfo.md) |  | [optional] 
**Scsi** | Pointer to [**VcenterVmHardwareScsiAddressInfo**](VcenterVmHardwareScsiAddressInfo.md) |  | [optional] 
**Sata** | Pointer to [**VcenterVmHardwareSataAddressInfo**](VcenterVmHardwareSataAddressInfo.md) |  | [optional] 
**Nvme** | Pointer to [**VcenterVmHardwareNvmeAddressInfo**](VcenterVmHardwareNvmeAddressInfo.md) |  | [optional] 
**Backing** | [**VcenterVmHardwareDiskBackingInfo**](VcenterVmHardwareDiskBackingInfo.md) |  | 
**Capacity** | Pointer to **int64** | Capacity of the virtual disk in bytes. If unset, virtual disk is inaccessible or disk capacity is 0. | [optional] 

## Methods

### NewVcenterVmHardwareDiskInfo

`func NewVcenterVmHardwareDiskInfo(label string, type_ VcenterVmHardwareDiskHostBusAdapterType, backing VcenterVmHardwareDiskBackingInfo, ) *VcenterVmHardwareDiskInfo`

NewVcenterVmHardwareDiskInfo instantiates a new VcenterVmHardwareDiskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareDiskInfoWithDefaults

`func NewVcenterVmHardwareDiskInfoWithDefaults() *VcenterVmHardwareDiskInfo`

NewVcenterVmHardwareDiskInfoWithDefaults instantiates a new VcenterVmHardwareDiskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *VcenterVmHardwareDiskInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VcenterVmHardwareDiskInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VcenterVmHardwareDiskInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *VcenterVmHardwareDiskInfo) GetType() VcenterVmHardwareDiskHostBusAdapterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareDiskInfo) GetTypeOk() (*VcenterVmHardwareDiskHostBusAdapterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareDiskInfo) SetType(v VcenterVmHardwareDiskHostBusAdapterType)`

SetType sets Type field to given value.


### GetIde

`func (o *VcenterVmHardwareDiskInfo) GetIde() VcenterVmHardwareIdeAddressInfo`

GetIde returns the Ide field if non-nil, zero value otherwise.

### GetIdeOk

`func (o *VcenterVmHardwareDiskInfo) GetIdeOk() (*VcenterVmHardwareIdeAddressInfo, bool)`

GetIdeOk returns a tuple with the Ide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde

`func (o *VcenterVmHardwareDiskInfo) SetIde(v VcenterVmHardwareIdeAddressInfo)`

SetIde sets Ide field to given value.

### HasIde

`func (o *VcenterVmHardwareDiskInfo) HasIde() bool`

HasIde returns a boolean if a field has been set.

### GetScsi

`func (o *VcenterVmHardwareDiskInfo) GetScsi() VcenterVmHardwareScsiAddressInfo`

GetScsi returns the Scsi field if non-nil, zero value otherwise.

### GetScsiOk

`func (o *VcenterVmHardwareDiskInfo) GetScsiOk() (*VcenterVmHardwareScsiAddressInfo, bool)`

GetScsiOk returns a tuple with the Scsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi

`func (o *VcenterVmHardwareDiskInfo) SetScsi(v VcenterVmHardwareScsiAddressInfo)`

SetScsi sets Scsi field to given value.

### HasScsi

`func (o *VcenterVmHardwareDiskInfo) HasScsi() bool`

HasScsi returns a boolean if a field has been set.

### GetSata

`func (o *VcenterVmHardwareDiskInfo) GetSata() VcenterVmHardwareSataAddressInfo`

GetSata returns the Sata field if non-nil, zero value otherwise.

### GetSataOk

`func (o *VcenterVmHardwareDiskInfo) GetSataOk() (*VcenterVmHardwareSataAddressInfo, bool)`

GetSataOk returns a tuple with the Sata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata

`func (o *VcenterVmHardwareDiskInfo) SetSata(v VcenterVmHardwareSataAddressInfo)`

SetSata sets Sata field to given value.

### HasSata

`func (o *VcenterVmHardwareDiskInfo) HasSata() bool`

HasSata returns a boolean if a field has been set.

### GetNvme

`func (o *VcenterVmHardwareDiskInfo) GetNvme() VcenterVmHardwareNvmeAddressInfo`

GetNvme returns the Nvme field if non-nil, zero value otherwise.

### GetNvmeOk

`func (o *VcenterVmHardwareDiskInfo) GetNvmeOk() (*VcenterVmHardwareNvmeAddressInfo, bool)`

GetNvmeOk returns a tuple with the Nvme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvme

`func (o *VcenterVmHardwareDiskInfo) SetNvme(v VcenterVmHardwareNvmeAddressInfo)`

SetNvme sets Nvme field to given value.

### HasNvme

`func (o *VcenterVmHardwareDiskInfo) HasNvme() bool`

HasNvme returns a boolean if a field has been set.

### GetBacking

`func (o *VcenterVmHardwareDiskInfo) GetBacking() VcenterVmHardwareDiskBackingInfo`

GetBacking returns the Backing field if non-nil, zero value otherwise.

### GetBackingOk

`func (o *VcenterVmHardwareDiskInfo) GetBackingOk() (*VcenterVmHardwareDiskBackingInfo, bool)`

GetBackingOk returns a tuple with the Backing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacking

`func (o *VcenterVmHardwareDiskInfo) SetBacking(v VcenterVmHardwareDiskBackingInfo)`

SetBacking sets Backing field to given value.


### GetCapacity

`func (o *VcenterVmHardwareDiskInfo) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VcenterVmHardwareDiskInfo) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VcenterVmHardwareDiskInfo) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VcenterVmHardwareDiskInfo) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


