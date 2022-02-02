# VcenterVmHardwareParallelBackingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmHardwareParallelBackingType**](VcenterVmHardwareParallelBackingType.md) |  | 
**File** | Pointer to **string** | Path of the file that should be used as the virtual parallel port backing. This field is optional and it is only relevant when the value of Parallel.BackingSpec.type is FILE. | [optional] 
**HostDevice** | Pointer to **string** | Name of the device that should be used as the virtual parallel port backing. If unset, the virtual parallel port will be configured to automatically detect a suitable host device. | [optional] 

## Methods

### NewVcenterVmHardwareParallelBackingSpec

`func NewVcenterVmHardwareParallelBackingSpec(type_ VcenterVmHardwareParallelBackingType, ) *VcenterVmHardwareParallelBackingSpec`

NewVcenterVmHardwareParallelBackingSpec instantiates a new VcenterVmHardwareParallelBackingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareParallelBackingSpecWithDefaults

`func NewVcenterVmHardwareParallelBackingSpecWithDefaults() *VcenterVmHardwareParallelBackingSpec`

NewVcenterVmHardwareParallelBackingSpecWithDefaults instantiates a new VcenterVmHardwareParallelBackingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareParallelBackingSpec) GetType() VcenterVmHardwareParallelBackingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareParallelBackingSpec) GetTypeOk() (*VcenterVmHardwareParallelBackingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareParallelBackingSpec) SetType(v VcenterVmHardwareParallelBackingType)`

SetType sets Type field to given value.


### GetFile

`func (o *VcenterVmHardwareParallelBackingSpec) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *VcenterVmHardwareParallelBackingSpec) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *VcenterVmHardwareParallelBackingSpec) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *VcenterVmHardwareParallelBackingSpec) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetHostDevice

`func (o *VcenterVmHardwareParallelBackingSpec) GetHostDevice() string`

GetHostDevice returns the HostDevice field if non-nil, zero value otherwise.

### GetHostDeviceOk

`func (o *VcenterVmHardwareParallelBackingSpec) GetHostDeviceOk() (*string, bool)`

GetHostDeviceOk returns a tuple with the HostDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDevice

`func (o *VcenterVmHardwareParallelBackingSpec) SetHostDevice(v string)`

SetHostDevice sets HostDevice field to given value.

### HasHostDevice

`func (o *VcenterVmHardwareParallelBackingSpec) HasHostDevice() bool`

HasHostDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


