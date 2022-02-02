# VcenterVmGuestFilesystemTransfersFileCreationAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int64** | The size in bytes of the file to be transferred into the guest. | 
**Overwrite** | Pointer to **bool** | Whether an existing file should be overwritten. If unset any existing file will not be overwritten. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time the file was last modified. If unset the value will be the time when the file is transferred into the guest. | [optional] 
**LastAccessed** | Pointer to **time.Time** | The date and time the file was last accessed. If unset the value will be the time when the file is transferred into the guest. | [optional] 
**Windows** | Pointer to [**VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec**](VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec.md) |  | [optional] 
**Posix** | Pointer to [**VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec**](VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec.md) |  | [optional] 

## Methods

### NewVcenterVmGuestFilesystemTransfersFileCreationAttributes

`func NewVcenterVmGuestFilesystemTransfersFileCreationAttributes(size int64, ) *VcenterVmGuestFilesystemTransfersFileCreationAttributes`

NewVcenterVmGuestFilesystemTransfersFileCreationAttributes instantiates a new VcenterVmGuestFilesystemTransfersFileCreationAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestFilesystemTransfersFileCreationAttributesWithDefaults

`func NewVcenterVmGuestFilesystemTransfersFileCreationAttributesWithDefaults() *VcenterVmGuestFilesystemTransfersFileCreationAttributes`

NewVcenterVmGuestFilesystemTransfersFileCreationAttributesWithDefaults instantiates a new VcenterVmGuestFilesystemTransfersFileCreationAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) SetSize(v int64)`

SetSize sets Size field to given value.


### GetOverwrite

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetLastModified

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastAccessed

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) GetLastAccessed() time.Time`

GetLastAccessed returns the LastAccessed field if non-nil, zero value otherwise.

### GetLastAccessedOk

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) GetLastAccessedOk() (*time.Time, bool)`

GetLastAccessedOk returns a tuple with the LastAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessed

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) SetLastAccessed(v time.Time)`

SetLastAccessed sets LastAccessed field to given value.

### HasLastAccessed

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) HasLastAccessed() bool`

HasLastAccessed returns a boolean if a field has been set.

### GetWindows

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) GetWindows() VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) GetWindowsOk() (*VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) SetWindows(v VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) HasWindows() bool`

HasWindows returns a boolean if a field has been set.

### GetPosix

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) GetPosix() VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec`

GetPosix returns the Posix field if non-nil, zero value otherwise.

### GetPosixOk

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) GetPosixOk() (*VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec, bool)`

GetPosixOk returns a tuple with the Posix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosix

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) SetPosix(v VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec)`

SetPosix sets Posix field to given value.

### HasPosix

`func (o *VcenterVmGuestFilesystemTransfersFileCreationAttributes) HasPosix() bool`

HasPosix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


