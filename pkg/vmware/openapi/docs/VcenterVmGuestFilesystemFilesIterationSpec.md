# VcenterVmGuestFilesystemFilesIterationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int64** | Specifies the maximum number of results to return. If unset information about at most 50 files will be returned. | [optional] 
**Index** | Pointer to **int64** | Which result to start the list with. If this value exceeds the number of results, an empty list will be returned. If unset, the start of the list of files will be returned. | [optional] 

## Methods

### NewVcenterVmGuestFilesystemFilesIterationSpec

`func NewVcenterVmGuestFilesystemFilesIterationSpec() *VcenterVmGuestFilesystemFilesIterationSpec`

NewVcenterVmGuestFilesystemFilesIterationSpec instantiates a new VcenterVmGuestFilesystemFilesIterationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestFilesystemFilesIterationSpecWithDefaults

`func NewVcenterVmGuestFilesystemFilesIterationSpecWithDefaults() *VcenterVmGuestFilesystemFilesIterationSpec`

NewVcenterVmGuestFilesystemFilesIterationSpecWithDefaults instantiates a new VcenterVmGuestFilesystemFilesIterationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *VcenterVmGuestFilesystemFilesIterationSpec) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VcenterVmGuestFilesystemFilesIterationSpec) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VcenterVmGuestFilesystemFilesIterationSpec) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *VcenterVmGuestFilesystemFilesIterationSpec) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetIndex

`func (o *VcenterVmGuestFilesystemFilesIterationSpec) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *VcenterVmGuestFilesystemFilesIterationSpec) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *VcenterVmGuestFilesystemFilesIterationSpec) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *VcenterVmGuestFilesystemFilesIterationSpec) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


