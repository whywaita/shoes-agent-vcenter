# VcenterVmGuestFilesystemFilesPosixFileAttributesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **int64** | The owner ID. | 
**Group** | **int64** | The group ID. | 
**Permissions** | **string** | The file permissions in chmod(2) format. This field is presented as octal. | 

## Methods

### NewVcenterVmGuestFilesystemFilesPosixFileAttributesInfo

`func NewVcenterVmGuestFilesystemFilesPosixFileAttributesInfo(owner int64, group int64, permissions string, ) *VcenterVmGuestFilesystemFilesPosixFileAttributesInfo`

NewVcenterVmGuestFilesystemFilesPosixFileAttributesInfo instantiates a new VcenterVmGuestFilesystemFilesPosixFileAttributesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestFilesystemFilesPosixFileAttributesInfoWithDefaults

`func NewVcenterVmGuestFilesystemFilesPosixFileAttributesInfoWithDefaults() *VcenterVmGuestFilesystemFilesPosixFileAttributesInfo`

NewVcenterVmGuestFilesystemFilesPosixFileAttributesInfoWithDefaults instantiates a new VcenterVmGuestFilesystemFilesPosixFileAttributesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesInfo) GetOwner() int64`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesInfo) GetOwnerOk() (*int64, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesInfo) SetOwner(v int64)`

SetOwner sets Owner field to given value.


### GetGroup

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesInfo) GetGroup() int64`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesInfo) GetGroupOk() (*int64, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesInfo) SetGroup(v int64)`

SetGroup sets Group field to given value.


### GetPermissions

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesInfo) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesInfo) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesInfo) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


