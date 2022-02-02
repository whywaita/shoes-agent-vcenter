# VcenterVmHardwareFloppyBackingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmHardwareFloppyBackingType**](VcenterVmHardwareFloppyBackingType.md) |  | 
**ImageFile** | Pointer to **string** | Path of the image file backing the virtual floppy drive. This field is optional and it is only relevant when the value of Floppy.BackingInfo.type is IMAGE_FILE. | [optional] 
**HostDevice** | Pointer to **string** | Name of the host device backing the virtual floppy drive.    This field will be unset if Floppy.BackingInfo.auto-detect is true and the virtual floppy drive is not connected or no suitable device is available on the host. | [optional] 
**AutoDetect** | Pointer to **bool** | Flag indicating whether the virtual floppy drive is configured to automatically detect a suitable host device. This field is optional and it is only relevant when the value of Floppy.BackingInfo.type is HOST_DEVICE. | [optional] 

## Methods

### NewVcenterVmHardwareFloppyBackingInfo

`func NewVcenterVmHardwareFloppyBackingInfo(type_ VcenterVmHardwareFloppyBackingType, ) *VcenterVmHardwareFloppyBackingInfo`

NewVcenterVmHardwareFloppyBackingInfo instantiates a new VcenterVmHardwareFloppyBackingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareFloppyBackingInfoWithDefaults

`func NewVcenterVmHardwareFloppyBackingInfoWithDefaults() *VcenterVmHardwareFloppyBackingInfo`

NewVcenterVmHardwareFloppyBackingInfoWithDefaults instantiates a new VcenterVmHardwareFloppyBackingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareFloppyBackingInfo) GetType() VcenterVmHardwareFloppyBackingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareFloppyBackingInfo) GetTypeOk() (*VcenterVmHardwareFloppyBackingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareFloppyBackingInfo) SetType(v VcenterVmHardwareFloppyBackingType)`

SetType sets Type field to given value.


### GetImageFile

`func (o *VcenterVmHardwareFloppyBackingInfo) GetImageFile() string`

GetImageFile returns the ImageFile field if non-nil, zero value otherwise.

### GetImageFileOk

`func (o *VcenterVmHardwareFloppyBackingInfo) GetImageFileOk() (*string, bool)`

GetImageFileOk returns a tuple with the ImageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFile

`func (o *VcenterVmHardwareFloppyBackingInfo) SetImageFile(v string)`

SetImageFile sets ImageFile field to given value.

### HasImageFile

`func (o *VcenterVmHardwareFloppyBackingInfo) HasImageFile() bool`

HasImageFile returns a boolean if a field has been set.

### GetHostDevice

`func (o *VcenterVmHardwareFloppyBackingInfo) GetHostDevice() string`

GetHostDevice returns the HostDevice field if non-nil, zero value otherwise.

### GetHostDeviceOk

`func (o *VcenterVmHardwareFloppyBackingInfo) GetHostDeviceOk() (*string, bool)`

GetHostDeviceOk returns a tuple with the HostDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDevice

`func (o *VcenterVmHardwareFloppyBackingInfo) SetHostDevice(v string)`

SetHostDevice sets HostDevice field to given value.

### HasHostDevice

`func (o *VcenterVmHardwareFloppyBackingInfo) HasHostDevice() bool`

HasHostDevice returns a boolean if a field has been set.

### GetAutoDetect

`func (o *VcenterVmHardwareFloppyBackingInfo) GetAutoDetect() bool`

GetAutoDetect returns the AutoDetect field if non-nil, zero value otherwise.

### GetAutoDetectOk

`func (o *VcenterVmHardwareFloppyBackingInfo) GetAutoDetectOk() (*bool, bool)`

GetAutoDetectOk returns a tuple with the AutoDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDetect

`func (o *VcenterVmHardwareFloppyBackingInfo) SetAutoDetect(v bool)`

SetAutoDetect sets AutoDetect field to given value.

### HasAutoDetect

`func (o *VcenterVmHardwareFloppyBackingInfo) HasAutoDetect() bool`

HasAutoDetect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


