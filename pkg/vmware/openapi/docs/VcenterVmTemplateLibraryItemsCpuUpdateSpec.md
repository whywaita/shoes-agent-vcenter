# VcenterVmTemplateLibraryItemsCpuUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumCpus** | Pointer to **int64** | Number of virtual processors in the deployed virtual machine. | [optional] 
**NumCoresPerSocket** | Pointer to **int64** | Number of cores among which to distribute CPUs in the deployed virtual machine. | [optional] 

## Methods

### NewVcenterVmTemplateLibraryItemsCpuUpdateSpec

`func NewVcenterVmTemplateLibraryItemsCpuUpdateSpec() *VcenterVmTemplateLibraryItemsCpuUpdateSpec`

NewVcenterVmTemplateLibraryItemsCpuUpdateSpec instantiates a new VcenterVmTemplateLibraryItemsCpuUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsCpuUpdateSpecWithDefaults

`func NewVcenterVmTemplateLibraryItemsCpuUpdateSpecWithDefaults() *VcenterVmTemplateLibraryItemsCpuUpdateSpec`

NewVcenterVmTemplateLibraryItemsCpuUpdateSpecWithDefaults instantiates a new VcenterVmTemplateLibraryItemsCpuUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumCpus

`func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) GetNumCpus() int64`

GetNumCpus returns the NumCpus field if non-nil, zero value otherwise.

### GetNumCpusOk

`func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) GetNumCpusOk() (*int64, bool)`

GetNumCpusOk returns a tuple with the NumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpus

`func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) SetNumCpus(v int64)`

SetNumCpus sets NumCpus field to given value.

### HasNumCpus

`func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) HasNumCpus() bool`

HasNumCpus returns a boolean if a field has been set.

### GetNumCoresPerSocket

`func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) GetNumCoresPerSocket() int64`

GetNumCoresPerSocket returns the NumCoresPerSocket field if non-nil, zero value otherwise.

### GetNumCoresPerSocketOk

`func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) GetNumCoresPerSocketOk() (*int64, bool)`

GetNumCoresPerSocketOk returns a tuple with the NumCoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCoresPerSocket

`func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) SetNumCoresPerSocket(v int64)`

SetNumCoresPerSocket sets NumCoresPerSocket field to given value.

### HasNumCoresPerSocket

`func (o *VcenterVmTemplateLibraryItemsCpuUpdateSpec) HasNumCoresPerSocket() bool`

HasNumCoresPerSocket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


