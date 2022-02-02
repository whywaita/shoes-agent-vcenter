# VcenterOvfLibraryItemDeploymentTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePoolId** | **string** | Identifier of the resource pool to which the virtual machine or virtual appliance should be attached. | 
**HostId** | Pointer to **string** | Identifier of the target host on which the virtual machine or virtual appliance will run. The target host must be a member of the cluster that contains the resource pool identified by {@link #resourcePoolId}. | [optional] 
**FolderId** | Pointer to **string** | Identifier of the vCenter folder that should contain the virtual machine or virtual appliance. The folder must be virtual machine folder. | [optional] 

## Methods

### NewVcenterOvfLibraryItemDeploymentTarget

`func NewVcenterOvfLibraryItemDeploymentTarget(resourcePoolId string, ) *VcenterOvfLibraryItemDeploymentTarget`

NewVcenterOvfLibraryItemDeploymentTarget instantiates a new VcenterOvfLibraryItemDeploymentTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfLibraryItemDeploymentTargetWithDefaults

`func NewVcenterOvfLibraryItemDeploymentTargetWithDefaults() *VcenterOvfLibraryItemDeploymentTarget`

NewVcenterOvfLibraryItemDeploymentTargetWithDefaults instantiates a new VcenterOvfLibraryItemDeploymentTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePoolId

`func (o *VcenterOvfLibraryItemDeploymentTarget) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *VcenterOvfLibraryItemDeploymentTarget) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *VcenterOvfLibraryItemDeploymentTarget) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.


### GetHostId

`func (o *VcenterOvfLibraryItemDeploymentTarget) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *VcenterOvfLibraryItemDeploymentTarget) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *VcenterOvfLibraryItemDeploymentTarget) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *VcenterOvfLibraryItemDeploymentTarget) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetFolderId

`func (o *VcenterOvfLibraryItemDeploymentTarget) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *VcenterOvfLibraryItemDeploymentTarget) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *VcenterOvfLibraryItemDeploymentTarget) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *VcenterOvfLibraryItemDeploymentTarget) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


