# VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastModified** | Pointer to **time.Time** | The date and time the file was last modified. If unset the value will not be changed. | [optional] 
**LastAccessed** | Pointer to **time.Time** | The date and time the file was last accessed. If unset the value will not be changed. | [optional] 
**Windows** | Pointer to [**VcenterVmGuestFilesystemFilesWindowsFileAttributesUpdateSpec**](VcenterVmGuestFilesystemFilesWindowsFileAttributesUpdateSpec.md) |  | [optional] 
**Posix** | Pointer to [**VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec**](VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec.md) |  | [optional] 

## Methods

### NewVcenterVmGuestFilesystemFilesFileAttributesUpdateSpec

`func NewVcenterVmGuestFilesystemFilesFileAttributesUpdateSpec() *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec`

NewVcenterVmGuestFilesystemFilesFileAttributesUpdateSpec instantiates a new VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestFilesystemFilesFileAttributesUpdateSpecWithDefaults

`func NewVcenterVmGuestFilesystemFilesFileAttributesUpdateSpecWithDefaults() *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec`

NewVcenterVmGuestFilesystemFilesFileAttributesUpdateSpecWithDefaults instantiates a new VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModified

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastAccessed

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) GetLastAccessed() time.Time`

GetLastAccessed returns the LastAccessed field if non-nil, zero value otherwise.

### GetLastAccessedOk

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) GetLastAccessedOk() (*time.Time, bool)`

GetLastAccessedOk returns a tuple with the LastAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessed

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) SetLastAccessed(v time.Time)`

SetLastAccessed sets LastAccessed field to given value.

### HasLastAccessed

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) HasLastAccessed() bool`

HasLastAccessed returns a boolean if a field has been set.

### GetWindows

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) GetWindows() VcenterVmGuestFilesystemFilesWindowsFileAttributesUpdateSpec`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) GetWindowsOk() (*VcenterVmGuestFilesystemFilesWindowsFileAttributesUpdateSpec, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) SetWindows(v VcenterVmGuestFilesystemFilesWindowsFileAttributesUpdateSpec)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) HasWindows() bool`

HasWindows returns a boolean if a field has been set.

### GetPosix

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) GetPosix() VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec`

GetPosix returns the Posix field if non-nil, zero value otherwise.

### GetPosixOk

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) GetPosixOk() (*VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec, bool)`

GetPosixOk returns a tuple with the Posix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosix

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) SetPosix(v VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec)`

SetPosix sets Posix field to given value.

### HasPosix

`func (o *VcenterVmGuestFilesystemFilesFileAttributesUpdateSpec) HasPosix() bool`

HasPosix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


