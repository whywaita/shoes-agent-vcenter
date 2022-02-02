# VcenterVmHardwareCdromBackingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmHardwareCdromBackingType**](VcenterVmHardwareCdromBackingType.md) |  | 
**IsoFile** | Pointer to **string** | Path of the image file backing the virtual CD-ROM device. This field is optional and it is only relevant when the value of Cdrom.BackingInfo.type is ISO_FILE. | [optional] 
**HostDevice** | Pointer to **string** | Name of the host device backing the virtual CD-ROM device.    This field will be unset if Cdrom.BackingInfo.auto-detect is true and the virtual CD-ROM device is not connected or no suitable device is available on the host. | [optional] 
**AutoDetect** | Pointer to **bool** | Flag indicating whether the virtual CD-ROM device is configured to automatically detect a suitable host device. This field is optional and it is only relevant when the value of Cdrom.BackingInfo.type is HOST_DEVICE. | [optional] 
**DeviceAccessType** | Pointer to [**VcenterVmHardwareCdromDeviceAccessType**](VcenterVmHardwareCdromDeviceAccessType.md) |  | [optional] 

## Methods

### NewVcenterVmHardwareCdromBackingInfo

`func NewVcenterVmHardwareCdromBackingInfo(type_ VcenterVmHardwareCdromBackingType, ) *VcenterVmHardwareCdromBackingInfo`

NewVcenterVmHardwareCdromBackingInfo instantiates a new VcenterVmHardwareCdromBackingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareCdromBackingInfoWithDefaults

`func NewVcenterVmHardwareCdromBackingInfoWithDefaults() *VcenterVmHardwareCdromBackingInfo`

NewVcenterVmHardwareCdromBackingInfoWithDefaults instantiates a new VcenterVmHardwareCdromBackingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareCdromBackingInfo) GetType() VcenterVmHardwareCdromBackingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareCdromBackingInfo) GetTypeOk() (*VcenterVmHardwareCdromBackingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareCdromBackingInfo) SetType(v VcenterVmHardwareCdromBackingType)`

SetType sets Type field to given value.


### GetIsoFile

`func (o *VcenterVmHardwareCdromBackingInfo) GetIsoFile() string`

GetIsoFile returns the IsoFile field if non-nil, zero value otherwise.

### GetIsoFileOk

`func (o *VcenterVmHardwareCdromBackingInfo) GetIsoFileOk() (*string, bool)`

GetIsoFileOk returns a tuple with the IsoFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoFile

`func (o *VcenterVmHardwareCdromBackingInfo) SetIsoFile(v string)`

SetIsoFile sets IsoFile field to given value.

### HasIsoFile

`func (o *VcenterVmHardwareCdromBackingInfo) HasIsoFile() bool`

HasIsoFile returns a boolean if a field has been set.

### GetHostDevice

`func (o *VcenterVmHardwareCdromBackingInfo) GetHostDevice() string`

GetHostDevice returns the HostDevice field if non-nil, zero value otherwise.

### GetHostDeviceOk

`func (o *VcenterVmHardwareCdromBackingInfo) GetHostDeviceOk() (*string, bool)`

GetHostDeviceOk returns a tuple with the HostDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDevice

`func (o *VcenterVmHardwareCdromBackingInfo) SetHostDevice(v string)`

SetHostDevice sets HostDevice field to given value.

### HasHostDevice

`func (o *VcenterVmHardwareCdromBackingInfo) HasHostDevice() bool`

HasHostDevice returns a boolean if a field has been set.

### GetAutoDetect

`func (o *VcenterVmHardwareCdromBackingInfo) GetAutoDetect() bool`

GetAutoDetect returns the AutoDetect field if non-nil, zero value otherwise.

### GetAutoDetectOk

`func (o *VcenterVmHardwareCdromBackingInfo) GetAutoDetectOk() (*bool, bool)`

GetAutoDetectOk returns a tuple with the AutoDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDetect

`func (o *VcenterVmHardwareCdromBackingInfo) SetAutoDetect(v bool)`

SetAutoDetect sets AutoDetect field to given value.

### HasAutoDetect

`func (o *VcenterVmHardwareCdromBackingInfo) HasAutoDetect() bool`

HasAutoDetect returns a boolean if a field has been set.

### GetDeviceAccessType

`func (o *VcenterVmHardwareCdromBackingInfo) GetDeviceAccessType() VcenterVmHardwareCdromDeviceAccessType`

GetDeviceAccessType returns the DeviceAccessType field if non-nil, zero value otherwise.

### GetDeviceAccessTypeOk

`func (o *VcenterVmHardwareCdromBackingInfo) GetDeviceAccessTypeOk() (*VcenterVmHardwareCdromDeviceAccessType, bool)`

GetDeviceAccessTypeOk returns a tuple with the DeviceAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAccessType

`func (o *VcenterVmHardwareCdromBackingInfo) SetDeviceAccessType(v VcenterVmHardwareCdromDeviceAccessType)`

SetDeviceAccessType sets DeviceAccessType field to given value.

### HasDeviceAccessType

`func (o *VcenterVmHardwareCdromBackingInfo) HasDeviceAccessType() bool`

HasDeviceAccessType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


