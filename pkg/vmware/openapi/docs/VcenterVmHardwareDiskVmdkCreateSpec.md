# VcenterVmHardwareDiskVmdkCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Base name of the VMDK file. The name should not include the &#39;.vmdk&#39; file extension. If unset, a name (derived from the name of the virtual machine) will be chosen by the server. | [optional] 
**Capacity** | Pointer to **int64** | Capacity of the virtual disk backing in bytes. If unset, defaults to a guest-specific capacity. | [optional] 
**StoragePolicy** | Pointer to [**VcenterVmHardwareDiskStoragePolicySpec**](VcenterVmHardwareDiskStoragePolicySpec.md) |  | [optional] 

## Methods

### NewVcenterVmHardwareDiskVmdkCreateSpec

`func NewVcenterVmHardwareDiskVmdkCreateSpec() *VcenterVmHardwareDiskVmdkCreateSpec`

NewVcenterVmHardwareDiskVmdkCreateSpec instantiates a new VcenterVmHardwareDiskVmdkCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareDiskVmdkCreateSpecWithDefaults

`func NewVcenterVmHardwareDiskVmdkCreateSpecWithDefaults() *VcenterVmHardwareDiskVmdkCreateSpec`

NewVcenterVmHardwareDiskVmdkCreateSpecWithDefaults instantiates a new VcenterVmHardwareDiskVmdkCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterVmHardwareDiskVmdkCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVmHardwareDiskVmdkCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVmHardwareDiskVmdkCreateSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterVmHardwareDiskVmdkCreateSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCapacity

`func (o *VcenterVmHardwareDiskVmdkCreateSpec) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VcenterVmHardwareDiskVmdkCreateSpec) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VcenterVmHardwareDiskVmdkCreateSpec) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VcenterVmHardwareDiskVmdkCreateSpec) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetStoragePolicy

`func (o *VcenterVmHardwareDiskVmdkCreateSpec) GetStoragePolicy() VcenterVmHardwareDiskStoragePolicySpec`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *VcenterVmHardwareDiskVmdkCreateSpec) GetStoragePolicyOk() (*VcenterVmHardwareDiskStoragePolicySpec, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *VcenterVmHardwareDiskVmdkCreateSpec) SetStoragePolicy(v VcenterVmHardwareDiskStoragePolicySpec)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *VcenterVmHardwareDiskVmdkCreateSpec) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


