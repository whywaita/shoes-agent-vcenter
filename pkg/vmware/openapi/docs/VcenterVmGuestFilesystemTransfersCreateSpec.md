# VcenterVmGuestFilesystemTransfersCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The complete destination path in the guest to transfer the file to or from the client. It cannot be a path to a directory or a symbolic link. | 
**Attributes** | Pointer to [**VcenterVmGuestFilesystemTransfersFileCreationAttributes**](VcenterVmGuestFilesystemTransfersFileCreationAttributes.md) |  | [optional] 

## Methods

### NewVcenterVmGuestFilesystemTransfersCreateSpec

`func NewVcenterVmGuestFilesystemTransfersCreateSpec(path string, ) *VcenterVmGuestFilesystemTransfersCreateSpec`

NewVcenterVmGuestFilesystemTransfersCreateSpec instantiates a new VcenterVmGuestFilesystemTransfersCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestFilesystemTransfersCreateSpecWithDefaults

`func NewVcenterVmGuestFilesystemTransfersCreateSpecWithDefaults() *VcenterVmGuestFilesystemTransfersCreateSpec`

NewVcenterVmGuestFilesystemTransfersCreateSpecWithDefaults instantiates a new VcenterVmGuestFilesystemTransfersCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *VcenterVmGuestFilesystemTransfersCreateSpec) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VcenterVmGuestFilesystemTransfersCreateSpec) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VcenterVmGuestFilesystemTransfersCreateSpec) SetPath(v string)`

SetPath sets Path field to given value.


### GetAttributes

`func (o *VcenterVmGuestFilesystemTransfersCreateSpec) GetAttributes() VcenterVmGuestFilesystemTransfersFileCreationAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VcenterVmGuestFilesystemTransfersCreateSpec) GetAttributesOk() (*VcenterVmGuestFilesystemTransfersFileCreationAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VcenterVmGuestFilesystemTransfersCreateSpec) SetAttributes(v VcenterVmGuestFilesystemTransfersFileCreationAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *VcenterVmGuestFilesystemTransfersCreateSpec) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


