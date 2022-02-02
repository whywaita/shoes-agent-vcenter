# VcenterVmGuestFilesystemFilesFileAttributesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastModified** | **time.Time** | The date and time the file was last modified. | 
**LastAccessed** | **time.Time** | The date and time the file was last accessed. | 
**SymlinkTarget** | Pointer to **string** | The target for the file if it&#39;s a symbolic link. This is currently only set for Posix guest operating systems, but may be supported in the future on Windows guest operating systems that support symbolic links. Set if the file is a symbolic link. | [optional] 
**FilesystemFamily** | [**VcenterVmGuestFilesystemFilesFilesystemFamily**](VcenterVmGuestFilesystemFilesFilesystemFamily.md) |  | 
**WinAttributes** | Pointer to [**VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo**](VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo.md) |  | [optional] 
**PosixAttributes** | Pointer to [**VcenterVmGuestFilesystemFilesPosixFileAttributesInfo**](VcenterVmGuestFilesystemFilesPosixFileAttributesInfo.md) |  | [optional] 

## Methods

### NewVcenterVmGuestFilesystemFilesFileAttributesInfo

`func NewVcenterVmGuestFilesystemFilesFileAttributesInfo(lastModified time.Time, lastAccessed time.Time, filesystemFamily VcenterVmGuestFilesystemFilesFilesystemFamily, ) *VcenterVmGuestFilesystemFilesFileAttributesInfo`

NewVcenterVmGuestFilesystemFilesFileAttributesInfo instantiates a new VcenterVmGuestFilesystemFilesFileAttributesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestFilesystemFilesFileAttributesInfoWithDefaults

`func NewVcenterVmGuestFilesystemFilesFileAttributesInfoWithDefaults() *VcenterVmGuestFilesystemFilesFileAttributesInfo`

NewVcenterVmGuestFilesystemFilesFileAttributesInfoWithDefaults instantiates a new VcenterVmGuestFilesystemFilesFileAttributesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModified

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetLastAccessed

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) GetLastAccessed() time.Time`

GetLastAccessed returns the LastAccessed field if non-nil, zero value otherwise.

### GetLastAccessedOk

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) GetLastAccessedOk() (*time.Time, bool)`

GetLastAccessedOk returns a tuple with the LastAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessed

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) SetLastAccessed(v time.Time)`

SetLastAccessed sets LastAccessed field to given value.


### GetSymlinkTarget

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) GetSymlinkTarget() string`

GetSymlinkTarget returns the SymlinkTarget field if non-nil, zero value otherwise.

### GetSymlinkTargetOk

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) GetSymlinkTargetOk() (*string, bool)`

GetSymlinkTargetOk returns a tuple with the SymlinkTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymlinkTarget

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) SetSymlinkTarget(v string)`

SetSymlinkTarget sets SymlinkTarget field to given value.

### HasSymlinkTarget

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) HasSymlinkTarget() bool`

HasSymlinkTarget returns a boolean if a field has been set.

### GetFilesystemFamily

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) GetFilesystemFamily() VcenterVmGuestFilesystemFilesFilesystemFamily`

GetFilesystemFamily returns the FilesystemFamily field if non-nil, zero value otherwise.

### GetFilesystemFamilyOk

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) GetFilesystemFamilyOk() (*VcenterVmGuestFilesystemFilesFilesystemFamily, bool)`

GetFilesystemFamilyOk returns a tuple with the FilesystemFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemFamily

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) SetFilesystemFamily(v VcenterVmGuestFilesystemFilesFilesystemFamily)`

SetFilesystemFamily sets FilesystemFamily field to given value.


### GetWinAttributes

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) GetWinAttributes() VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo`

GetWinAttributes returns the WinAttributes field if non-nil, zero value otherwise.

### GetWinAttributesOk

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) GetWinAttributesOk() (*VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo, bool)`

GetWinAttributesOk returns a tuple with the WinAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinAttributes

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) SetWinAttributes(v VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo)`

SetWinAttributes sets WinAttributes field to given value.

### HasWinAttributes

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) HasWinAttributes() bool`

HasWinAttributes returns a boolean if a field has been set.

### GetPosixAttributes

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) GetPosixAttributes() VcenterVmGuestFilesystemFilesPosixFileAttributesInfo`

GetPosixAttributes returns the PosixAttributes field if non-nil, zero value otherwise.

### GetPosixAttributesOk

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) GetPosixAttributesOk() (*VcenterVmGuestFilesystemFilesPosixFileAttributesInfo, bool)`

GetPosixAttributesOk returns a tuple with the PosixAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosixAttributes

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) SetPosixAttributes(v VcenterVmGuestFilesystemFilesPosixFileAttributesInfo)`

SetPosixAttributes sets PosixAttributes field to given value.

### HasPosixAttributes

`func (o *VcenterVmGuestFilesystemFilesFileAttributesInfo) HasPosixAttributes() bool`

HasPosixAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


