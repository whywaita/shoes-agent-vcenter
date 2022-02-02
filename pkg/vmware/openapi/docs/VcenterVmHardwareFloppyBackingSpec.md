# VcenterVmHardwareFloppyBackingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmHardwareFloppyBackingType**](VcenterVmHardwareFloppyBackingType.md) |  | 
**ImageFile** | Pointer to **string** | Path of the image file that should be used as the virtual floppy drive backing. This field is optional and it is only relevant when the value of Floppy.BackingSpec.type is IMAGE_FILE. | [optional] 
**HostDevice** | Pointer to **string** | Name of the device that should be used as the virtual floppy drive backing. If unset, the virtual floppy drive will be configured to automatically detect a suitable host device. | [optional] 

## Methods

### NewVcenterVmHardwareFloppyBackingSpec

`func NewVcenterVmHardwareFloppyBackingSpec(type_ VcenterVmHardwareFloppyBackingType, ) *VcenterVmHardwareFloppyBackingSpec`

NewVcenterVmHardwareFloppyBackingSpec instantiates a new VcenterVmHardwareFloppyBackingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareFloppyBackingSpecWithDefaults

`func NewVcenterVmHardwareFloppyBackingSpecWithDefaults() *VcenterVmHardwareFloppyBackingSpec`

NewVcenterVmHardwareFloppyBackingSpecWithDefaults instantiates a new VcenterVmHardwareFloppyBackingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareFloppyBackingSpec) GetType() VcenterVmHardwareFloppyBackingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareFloppyBackingSpec) GetTypeOk() (*VcenterVmHardwareFloppyBackingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareFloppyBackingSpec) SetType(v VcenterVmHardwareFloppyBackingType)`

SetType sets Type field to given value.


### GetImageFile

`func (o *VcenterVmHardwareFloppyBackingSpec) GetImageFile() string`

GetImageFile returns the ImageFile field if non-nil, zero value otherwise.

### GetImageFileOk

`func (o *VcenterVmHardwareFloppyBackingSpec) GetImageFileOk() (*string, bool)`

GetImageFileOk returns a tuple with the ImageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFile

`func (o *VcenterVmHardwareFloppyBackingSpec) SetImageFile(v string)`

SetImageFile sets ImageFile field to given value.

### HasImageFile

`func (o *VcenterVmHardwareFloppyBackingSpec) HasImageFile() bool`

HasImageFile returns a boolean if a field has been set.

### GetHostDevice

`func (o *VcenterVmHardwareFloppyBackingSpec) GetHostDevice() string`

GetHostDevice returns the HostDevice field if non-nil, zero value otherwise.

### GetHostDeviceOk

`func (o *VcenterVmHardwareFloppyBackingSpec) GetHostDeviceOk() (*string, bool)`

GetHostDeviceOk returns a tuple with the HostDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDevice

`func (o *VcenterVmHardwareFloppyBackingSpec) SetHostDevice(v string)`

SetHostDevice sets HostDevice field to given value.

### HasHostDevice

`func (o *VcenterVmHardwareFloppyBackingSpec) HasHostDevice() bool`

HasHostDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


