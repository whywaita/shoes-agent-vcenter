# VcenterVmTemplateLibraryItemsDiskUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | **int64** | Updated capacity of the virtual disk backing in bytes. This value has to be larger than the original capacity of the disk. | 

## Methods

### NewVcenterVmTemplateLibraryItemsDiskUpdateSpec

`func NewVcenterVmTemplateLibraryItemsDiskUpdateSpec(capacity int64, ) *VcenterVmTemplateLibraryItemsDiskUpdateSpec`

NewVcenterVmTemplateLibraryItemsDiskUpdateSpec instantiates a new VcenterVmTemplateLibraryItemsDiskUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsDiskUpdateSpecWithDefaults

`func NewVcenterVmTemplateLibraryItemsDiskUpdateSpecWithDefaults() *VcenterVmTemplateLibraryItemsDiskUpdateSpec`

NewVcenterVmTemplateLibraryItemsDiskUpdateSpecWithDefaults instantiates a new VcenterVmTemplateLibraryItemsDiskUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *VcenterVmTemplateLibraryItemsDiskUpdateSpec) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VcenterVmTemplateLibraryItemsDiskUpdateSpec) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VcenterVmTemplateLibraryItemsDiskUpdateSpec) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


