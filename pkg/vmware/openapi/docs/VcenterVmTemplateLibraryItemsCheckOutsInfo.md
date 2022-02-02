# VcenterVmTemplateLibraryItemsCheckOutsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** | Date and time when the virtual machine was checked out. | 
**User** | **string** | Name of the user who checked out the virtual machine. | 

## Methods

### NewVcenterVmTemplateLibraryItemsCheckOutsInfo

`func NewVcenterVmTemplateLibraryItemsCheckOutsInfo(time time.Time, user string, ) *VcenterVmTemplateLibraryItemsCheckOutsInfo`

NewVcenterVmTemplateLibraryItemsCheckOutsInfo instantiates a new VcenterVmTemplateLibraryItemsCheckOutsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsCheckOutsInfoWithDefaults

`func NewVcenterVmTemplateLibraryItemsCheckOutsInfoWithDefaults() *VcenterVmTemplateLibraryItemsCheckOutsInfo`

NewVcenterVmTemplateLibraryItemsCheckOutsInfoWithDefaults instantiates a new VcenterVmTemplateLibraryItemsCheckOutsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *VcenterVmTemplateLibraryItemsCheckOutsInfo) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *VcenterVmTemplateLibraryItemsCheckOutsInfo) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *VcenterVmTemplateLibraryItemsCheckOutsInfo) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetUser

`func (o *VcenterVmTemplateLibraryItemsCheckOutsInfo) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *VcenterVmTemplateLibraryItemsCheckOutsInfo) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *VcenterVmTemplateLibraryItemsCheckOutsInfo) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


