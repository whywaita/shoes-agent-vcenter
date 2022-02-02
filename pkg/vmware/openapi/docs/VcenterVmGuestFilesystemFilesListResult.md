# VcenterVmGuestFilesystemFilesListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | [**[]VcenterVmGuestFilesystemFilesSummary**](VcenterVmGuestFilesystemFilesSummary.md) | A list of Files.Summary structures containing information for all the matching files. | 
**Total** | **int64** | The total number of results from the Files.list. This is a hint to the user of the iterator regarding how many items are available to be retrieved. The total could change if the inventory of items are being changed. | 
**StartIndex** | Pointer to **int64** | Positional index into the logical item list of the first item returned in the list of results. The first item in the logical item list has an index of 0. This is a hint to the user of the iterator regarding the logical position in the iteration. For example, this can be used to display to the user which page of the iteration is being shown. The total could change if the inventory of items are being changed. If unset no items were returned. | [optional] 
**EndIndex** | Pointer to **int64** | Positional index into the logical item list of the last item returned in the list of results. The first item in the logical item list has an index of 0. This is a hint to the user of the iterator regarding the logical position in the iteration. For example, this can be used to display to the user which page of the iteration is being shown. The total could change if the inventory of items are being changed. If unset no items were returned. | [optional] 
**Status** | [**VcenterVmGuestFilesystemFilesLastIterationStatus**](VcenterVmGuestFilesystemFilesLastIterationStatus.md) |  | 

## Methods

### NewVcenterVmGuestFilesystemFilesListResult

`func NewVcenterVmGuestFilesystemFilesListResult(files []VcenterVmGuestFilesystemFilesSummary, total int64, status VcenterVmGuestFilesystemFilesLastIterationStatus, ) *VcenterVmGuestFilesystemFilesListResult`

NewVcenterVmGuestFilesystemFilesListResult instantiates a new VcenterVmGuestFilesystemFilesListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestFilesystemFilesListResultWithDefaults

`func NewVcenterVmGuestFilesystemFilesListResultWithDefaults() *VcenterVmGuestFilesystemFilesListResult`

NewVcenterVmGuestFilesystemFilesListResultWithDefaults instantiates a new VcenterVmGuestFilesystemFilesListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *VcenterVmGuestFilesystemFilesListResult) GetFiles() []VcenterVmGuestFilesystemFilesSummary`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *VcenterVmGuestFilesystemFilesListResult) GetFilesOk() (*[]VcenterVmGuestFilesystemFilesSummary, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *VcenterVmGuestFilesystemFilesListResult) SetFiles(v []VcenterVmGuestFilesystemFilesSummary)`

SetFiles sets Files field to given value.


### GetTotal

`func (o *VcenterVmGuestFilesystemFilesListResult) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *VcenterVmGuestFilesystemFilesListResult) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *VcenterVmGuestFilesystemFilesListResult) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetStartIndex

`func (o *VcenterVmGuestFilesystemFilesListResult) GetStartIndex() int64`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *VcenterVmGuestFilesystemFilesListResult) GetStartIndexOk() (*int64, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *VcenterVmGuestFilesystemFilesListResult) SetStartIndex(v int64)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *VcenterVmGuestFilesystemFilesListResult) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetEndIndex

`func (o *VcenterVmGuestFilesystemFilesListResult) GetEndIndex() int64`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *VcenterVmGuestFilesystemFilesListResult) GetEndIndexOk() (*int64, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *VcenterVmGuestFilesystemFilesListResult) SetEndIndex(v int64)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *VcenterVmGuestFilesystemFilesListResult) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetStatus

`func (o *VcenterVmGuestFilesystemFilesListResult) GetStatus() VcenterVmGuestFilesystemFilesLastIterationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VcenterVmGuestFilesystemFilesListResult) GetStatusOk() (*VcenterVmGuestFilesystemFilesLastIterationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VcenterVmGuestFilesystemFilesListResult) SetStatus(v VcenterVmGuestFilesystemFilesLastIterationStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


