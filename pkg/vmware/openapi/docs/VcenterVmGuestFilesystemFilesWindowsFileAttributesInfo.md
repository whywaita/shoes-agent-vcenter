# VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hidden** | **bool** | The file is hidden. | 
**ReadOnly** | **bool** | The file is read-only. | 
**Created** | **time.Time** | The date and time the file was created. | 

## Methods

### NewVcenterVmGuestFilesystemFilesWindowsFileAttributesInfo

`func NewVcenterVmGuestFilesystemFilesWindowsFileAttributesInfo(hidden bool, readOnly bool, created time.Time, ) *VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo`

NewVcenterVmGuestFilesystemFilesWindowsFileAttributesInfo instantiates a new VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestFilesystemFilesWindowsFileAttributesInfoWithDefaults

`func NewVcenterVmGuestFilesystemFilesWindowsFileAttributesInfoWithDefaults() *VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo`

NewVcenterVmGuestFilesystemFilesWindowsFileAttributesInfoWithDefaults instantiates a new VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHidden

`func (o *VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetReadOnly

`func (o *VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.


### GetCreated

`func (o *VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VcenterVmGuestFilesystemFilesWindowsFileAttributesInfo) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


