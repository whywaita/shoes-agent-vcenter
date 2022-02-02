# VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **int64** | The owner ID. If unset the value will not be changed. | [optional] 
**GroupId** | Pointer to **int64** | The group ID. If unset the value will not be changed. | [optional] 
**Permissions** | Pointer to **string** | The file permissions in chmod(2) format. This field is interpreted as octal. If unset the value will not be changed. | [optional] 

## Methods

### NewVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec

`func NewVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec() *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec`

NewVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec instantiates a new VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpecWithDefaults

`func NewVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpecWithDefaults() *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec`

NewVcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpecWithDefaults instantiates a new VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetGroupId

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetPermissions

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *VcenterVmGuestFilesystemFilesPosixFileAttributesUpdateSpec) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


