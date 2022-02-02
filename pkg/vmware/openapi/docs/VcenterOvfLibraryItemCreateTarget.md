# VcenterOvfLibraryItemCreateTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LibraryId** | Pointer to **string** | Identifier of the library in which a new library item should be created. This {@term field} is not used if the {@name #libraryItemId} {@term field} is specified. | [optional] 
**LibraryItemId** | Pointer to **string** | Identifier of the library item that should be should be updated. | [optional] 

## Methods

### NewVcenterOvfLibraryItemCreateTarget

`func NewVcenterOvfLibraryItemCreateTarget() *VcenterOvfLibraryItemCreateTarget`

NewVcenterOvfLibraryItemCreateTarget instantiates a new VcenterOvfLibraryItemCreateTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfLibraryItemCreateTargetWithDefaults

`func NewVcenterOvfLibraryItemCreateTargetWithDefaults() *VcenterOvfLibraryItemCreateTarget`

NewVcenterOvfLibraryItemCreateTargetWithDefaults instantiates a new VcenterOvfLibraryItemCreateTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLibraryId

`func (o *VcenterOvfLibraryItemCreateTarget) GetLibraryId() string`

GetLibraryId returns the LibraryId field if non-nil, zero value otherwise.

### GetLibraryIdOk

`func (o *VcenterOvfLibraryItemCreateTarget) GetLibraryIdOk() (*string, bool)`

GetLibraryIdOk returns a tuple with the LibraryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryId

`func (o *VcenterOvfLibraryItemCreateTarget) SetLibraryId(v string)`

SetLibraryId sets LibraryId field to given value.

### HasLibraryId

`func (o *VcenterOvfLibraryItemCreateTarget) HasLibraryId() bool`

HasLibraryId returns a boolean if a field has been set.

### GetLibraryItemId

`func (o *VcenterOvfLibraryItemCreateTarget) GetLibraryItemId() string`

GetLibraryItemId returns the LibraryItemId field if non-nil, zero value otherwise.

### GetLibraryItemIdOk

`func (o *VcenterOvfLibraryItemCreateTarget) GetLibraryItemIdOk() (*string, bool)`

GetLibraryItemIdOk returns a tuple with the LibraryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryItemId

`func (o *VcenterOvfLibraryItemCreateTarget) SetLibraryItemId(v string)`

SetLibraryItemId sets LibraryItemId field to given value.

### HasLibraryItemId

`func (o *VcenterOvfLibraryItemCreateTarget) HasLibraryItemId() bool`

HasLibraryItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


