# VcenterVmStoragePolicyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmHome** | Pointer to **string** | Storage Policy associated with virtual machine home. Ifunset, the virtual machine&#39;s home directory doesn&#39;t have any storage policy. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.StoragePolicy. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.StoragePolicy. | [optional] 
**Disks** | [**[]VcenterOvfLibraryItemResourcePoolDeploymentSpecNetworkMappings**](VcenterOvfLibraryItemResourcePoolDeploymentSpecNetworkMappings.md) | Storage policies associated with virtual disks. The values in this map are storage policy identifiers. They will be identifiers for the resource type:vcenter.StoragePolicy If the map is empty, the virtual machine does not have any disks or its disks are not associated with a storage policy. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Disk. | 

## Methods

### NewVcenterVmStoragePolicyInfo

`func NewVcenterVmStoragePolicyInfo(disks []VcenterOvfLibraryItemResourcePoolDeploymentSpecNetworkMappings, ) *VcenterVmStoragePolicyInfo`

NewVcenterVmStoragePolicyInfo instantiates a new VcenterVmStoragePolicyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmStoragePolicyInfoWithDefaults

`func NewVcenterVmStoragePolicyInfoWithDefaults() *VcenterVmStoragePolicyInfo`

NewVcenterVmStoragePolicyInfoWithDefaults instantiates a new VcenterVmStoragePolicyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmHome

`func (o *VcenterVmStoragePolicyInfo) GetVmHome() string`

GetVmHome returns the VmHome field if non-nil, zero value otherwise.

### GetVmHomeOk

`func (o *VcenterVmStoragePolicyInfo) GetVmHomeOk() (*string, bool)`

GetVmHomeOk returns a tuple with the VmHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHome

`func (o *VcenterVmStoragePolicyInfo) SetVmHome(v string)`

SetVmHome sets VmHome field to given value.

### HasVmHome

`func (o *VcenterVmStoragePolicyInfo) HasVmHome() bool`

HasVmHome returns a boolean if a field has been set.

### GetDisks

`func (o *VcenterVmStoragePolicyInfo) GetDisks() []VcenterOvfLibraryItemResourcePoolDeploymentSpecNetworkMappings`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VcenterVmStoragePolicyInfo) GetDisksOk() (*[]VcenterOvfLibraryItemResourcePoolDeploymentSpecNetworkMappings, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VcenterVmStoragePolicyInfo) SetDisks(v []VcenterOvfLibraryItemResourcePoolDeploymentSpecNetworkMappings)`

SetDisks sets Disks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


