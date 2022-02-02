# VcenterVmLibraryItemCheckOutInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LibraryItem** | **string** | Identifier of the library item that the virtual machine is checked out from. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: content.library.Item. When operations return a value of this structure as a result, the field will be an identifier for the resource type: content.library.Item. | 

## Methods

### NewVcenterVmLibraryItemCheckOutInfo

`func NewVcenterVmLibraryItemCheckOutInfo(libraryItem string, ) *VcenterVmLibraryItemCheckOutInfo`

NewVcenterVmLibraryItemCheckOutInfo instantiates a new VcenterVmLibraryItemCheckOutInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmLibraryItemCheckOutInfoWithDefaults

`func NewVcenterVmLibraryItemCheckOutInfoWithDefaults() *VcenterVmLibraryItemCheckOutInfo`

NewVcenterVmLibraryItemCheckOutInfoWithDefaults instantiates a new VcenterVmLibraryItemCheckOutInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLibraryItem

`func (o *VcenterVmLibraryItemCheckOutInfo) GetLibraryItem() string`

GetLibraryItem returns the LibraryItem field if non-nil, zero value otherwise.

### GetLibraryItemOk

`func (o *VcenterVmLibraryItemCheckOutInfo) GetLibraryItemOk() (*string, bool)`

GetLibraryItemOk returns a tuple with the LibraryItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryItem

`func (o *VcenterVmLibraryItemCheckOutInfo) SetLibraryItem(v string)`

SetLibraryItem sets LibraryItem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


