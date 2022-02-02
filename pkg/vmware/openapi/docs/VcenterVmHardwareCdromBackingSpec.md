# VcenterVmHardwareCdromBackingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmHardwareCdromBackingType**](VcenterVmHardwareCdromBackingType.md) |  | 
**IsoFile** | Pointer to **string** | Path of the image file that should be used as the virtual CD-ROM device backing. This field is optional and it is only relevant when the value of Cdrom.BackingSpec.type is ISO_FILE. | [optional] 
**HostDevice** | Pointer to **string** | Name of the device that should be used as the virtual CD-ROM device backing. If unset, the virtual CD-ROM device will be configured to automatically detect a suitable host device. | [optional] 
**DeviceAccessType** | Pointer to [**VcenterVmHardwareCdromDeviceAccessType**](VcenterVmHardwareCdromDeviceAccessType.md) |  | [optional] 

## Methods

### NewVcenterVmHardwareCdromBackingSpec

`func NewVcenterVmHardwareCdromBackingSpec(type_ VcenterVmHardwareCdromBackingType, ) *VcenterVmHardwareCdromBackingSpec`

NewVcenterVmHardwareCdromBackingSpec instantiates a new VcenterVmHardwareCdromBackingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareCdromBackingSpecWithDefaults

`func NewVcenterVmHardwareCdromBackingSpecWithDefaults() *VcenterVmHardwareCdromBackingSpec`

NewVcenterVmHardwareCdromBackingSpecWithDefaults instantiates a new VcenterVmHardwareCdromBackingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareCdromBackingSpec) GetType() VcenterVmHardwareCdromBackingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareCdromBackingSpec) GetTypeOk() (*VcenterVmHardwareCdromBackingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareCdromBackingSpec) SetType(v VcenterVmHardwareCdromBackingType)`

SetType sets Type field to given value.


### GetIsoFile

`func (o *VcenterVmHardwareCdromBackingSpec) GetIsoFile() string`

GetIsoFile returns the IsoFile field if non-nil, zero value otherwise.

### GetIsoFileOk

`func (o *VcenterVmHardwareCdromBackingSpec) GetIsoFileOk() (*string, bool)`

GetIsoFileOk returns a tuple with the IsoFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoFile

`func (o *VcenterVmHardwareCdromBackingSpec) SetIsoFile(v string)`

SetIsoFile sets IsoFile field to given value.

### HasIsoFile

`func (o *VcenterVmHardwareCdromBackingSpec) HasIsoFile() bool`

HasIsoFile returns a boolean if a field has been set.

### GetHostDevice

`func (o *VcenterVmHardwareCdromBackingSpec) GetHostDevice() string`

GetHostDevice returns the HostDevice field if non-nil, zero value otherwise.

### GetHostDeviceOk

`func (o *VcenterVmHardwareCdromBackingSpec) GetHostDeviceOk() (*string, bool)`

GetHostDeviceOk returns a tuple with the HostDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDevice

`func (o *VcenterVmHardwareCdromBackingSpec) SetHostDevice(v string)`

SetHostDevice sets HostDevice field to given value.

### HasHostDevice

`func (o *VcenterVmHardwareCdromBackingSpec) HasHostDevice() bool`

HasHostDevice returns a boolean if a field has been set.

### GetDeviceAccessType

`func (o *VcenterVmHardwareCdromBackingSpec) GetDeviceAccessType() VcenterVmHardwareCdromDeviceAccessType`

GetDeviceAccessType returns the DeviceAccessType field if non-nil, zero value otherwise.

### GetDeviceAccessTypeOk

`func (o *VcenterVmHardwareCdromBackingSpec) GetDeviceAccessTypeOk() (*VcenterVmHardwareCdromDeviceAccessType, bool)`

GetDeviceAccessTypeOk returns a tuple with the DeviceAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAccessType

`func (o *VcenterVmHardwareCdromBackingSpec) SetDeviceAccessType(v VcenterVmHardwareCdromDeviceAccessType)`

SetDeviceAccessType sets DeviceAccessType field to given value.

### HasDeviceAccessType

`func (o *VcenterVmHardwareCdromBackingSpec) HasDeviceAccessType() bool`

HasDeviceAccessType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


