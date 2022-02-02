# VcenterVmGuestFilesystemFilesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | The name of the file or directory with any leading directories removed. | 
**Type** | [**VcenterVmGuestFilesystemFilesType**](VcenterVmGuestFilesystemFilesType.md) |  | 
**Size** | **int64** | The file size in bytes. | 

## Methods

### NewVcenterVmGuestFilesystemFilesSummary

`func NewVcenterVmGuestFilesystemFilesSummary(filename string, type_ VcenterVmGuestFilesystemFilesType, size int64, ) *VcenterVmGuestFilesystemFilesSummary`

NewVcenterVmGuestFilesystemFilesSummary instantiates a new VcenterVmGuestFilesystemFilesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestFilesystemFilesSummaryWithDefaults

`func NewVcenterVmGuestFilesystemFilesSummaryWithDefaults() *VcenterVmGuestFilesystemFilesSummary`

NewVcenterVmGuestFilesystemFilesSummaryWithDefaults instantiates a new VcenterVmGuestFilesystemFilesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *VcenterVmGuestFilesystemFilesSummary) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *VcenterVmGuestFilesystemFilesSummary) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *VcenterVmGuestFilesystemFilesSummary) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetType

`func (o *VcenterVmGuestFilesystemFilesSummary) GetType() VcenterVmGuestFilesystemFilesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmGuestFilesystemFilesSummary) GetTypeOk() (*VcenterVmGuestFilesystemFilesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmGuestFilesystemFilesSummary) SetType(v VcenterVmGuestFilesystemFilesType)`

SetType sets Type field to given value.


### GetSize

`func (o *VcenterVmGuestFilesystemFilesSummary) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VcenterVmGuestFilesystemFilesSummary) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VcenterVmGuestFilesystemFilesSummary) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


