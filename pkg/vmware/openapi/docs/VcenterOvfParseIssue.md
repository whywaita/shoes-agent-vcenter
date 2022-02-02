# VcenterOvfParseIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [**VcenterOvfParseIssueCategory**](VcenterOvfParseIssueCategory.md) |  | 
**File** | **string** | The name of the file in which the parse issue was found. | 
**LineNumber** | **int64** | The line number of the line in the file (see {@link #file}) where the parse issue was found (or -1 if not applicable). | 
**ColumnNumber** | **int64** | The position in the line (see {@link #lineNumber}) (or -1 if not applicable). | 
**Message** | [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | 

## Methods

### NewVcenterOvfParseIssue

`func NewVcenterOvfParseIssue(category VcenterOvfParseIssueCategory, file string, lineNumber int64, columnNumber int64, message VapiStdLocalizableMessage, ) *VcenterOvfParseIssue`

NewVcenterOvfParseIssue instantiates a new VcenterOvfParseIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfParseIssueWithDefaults

`func NewVcenterOvfParseIssueWithDefaults() *VcenterOvfParseIssue`

NewVcenterOvfParseIssueWithDefaults instantiates a new VcenterOvfParseIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *VcenterOvfParseIssue) GetCategory() VcenterOvfParseIssueCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *VcenterOvfParseIssue) GetCategoryOk() (*VcenterOvfParseIssueCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *VcenterOvfParseIssue) SetCategory(v VcenterOvfParseIssueCategory)`

SetCategory sets Category field to given value.


### GetFile

`func (o *VcenterOvfParseIssue) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *VcenterOvfParseIssue) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *VcenterOvfParseIssue) SetFile(v string)`

SetFile sets File field to given value.


### GetLineNumber

`func (o *VcenterOvfParseIssue) GetLineNumber() int64`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *VcenterOvfParseIssue) GetLineNumberOk() (*int64, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *VcenterOvfParseIssue) SetLineNumber(v int64)`

SetLineNumber sets LineNumber field to given value.


### GetColumnNumber

`func (o *VcenterOvfParseIssue) GetColumnNumber() int64`

GetColumnNumber returns the ColumnNumber field if non-nil, zero value otherwise.

### GetColumnNumberOk

`func (o *VcenterOvfParseIssue) GetColumnNumberOk() (*int64, bool)`

GetColumnNumberOk returns a tuple with the ColumnNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnNumber

`func (o *VcenterOvfParseIssue) SetColumnNumber(v int64)`

SetColumnNumber sets ColumnNumber field to given value.


### GetMessage

`func (o *VcenterOvfParseIssue) GetMessage() VapiStdLocalizableMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VcenterOvfParseIssue) GetMessageOk() (*VapiStdLocalizableMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VcenterOvfParseIssue) SetMessage(v VapiStdLocalizableMessage)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


