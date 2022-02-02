# VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **int64** | The owner ID. If this property is not specified when passing a Transfers.PosixFileAttributesCreateSpec object to Transfers.create, the default value will be the owner Id of the user who invoked the file transfer operation. Defaults to uid of user invoking the operation. | [optional] 
**GroupId** | Pointer to **int64** | The group ID. If this property is not specified when passing a Transfers.PosixFileAttributesCreateSpec object to Transfers.create, the default value will be the group Id of the user who invoked the file transfer operation. Defaults to gid of user invoking the operation. | [optional] 
**Permissions** | Pointer to **string** | The file permissions in chmod(2) format. If this property is not specified when passing a Transfers.PosixFileAttributesCreateSpec object to Transfers.create, the file will be created with 0644 permissions. This field is interpreted as octal. Defaults to 0644. | [optional] 

## Methods

### NewVcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec

`func NewVcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec() *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec`

NewVcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec instantiates a new VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpecWithDefaults

`func NewVcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpecWithDefaults() *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec`

NewVcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpecWithDefaults instantiates a new VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetGroupId

`func (o *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetPermissions

`func (o *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *VcenterVmGuestFilesystemTransfersPosixFileAttributesCreateSpec) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


