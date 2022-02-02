# VcenterOvfOvfError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [**VcenterOvfOvfMessageCategory**](VcenterOvfOvfMessageCategory.md) |  | 
**Issues** | Pointer to [**[]VcenterOvfParseIssue**](VcenterOvfParseIssue.md) | {@term List} of parse issues (see {@link ParseIssue}). | [optional] 
**Name** | Pointer to **string** | The name of input parameter. | [optional] 
**Value** | Pointer to **string** | The value of input parameter. | [optional] 
**Message** | Pointer to [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | [optional] 
**Error** | Pointer to **map[string]interface{}** | Represents a server {@link Error}. | [optional] 

## Methods

### NewVcenterOvfOvfError

`func NewVcenterOvfOvfError(category VcenterOvfOvfMessageCategory, ) *VcenterOvfOvfError`

NewVcenterOvfOvfError instantiates a new VcenterOvfOvfError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfOvfErrorWithDefaults

`func NewVcenterOvfOvfErrorWithDefaults() *VcenterOvfOvfError`

NewVcenterOvfOvfErrorWithDefaults instantiates a new VcenterOvfOvfError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *VcenterOvfOvfError) GetCategory() VcenterOvfOvfMessageCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *VcenterOvfOvfError) GetCategoryOk() (*VcenterOvfOvfMessageCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *VcenterOvfOvfError) SetCategory(v VcenterOvfOvfMessageCategory)`

SetCategory sets Category field to given value.


### GetIssues

`func (o *VcenterOvfOvfError) GetIssues() []VcenterOvfParseIssue`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *VcenterOvfOvfError) GetIssuesOk() (*[]VcenterOvfParseIssue, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *VcenterOvfOvfError) SetIssues(v []VcenterOvfParseIssue)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *VcenterOvfOvfError) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetName

`func (o *VcenterOvfOvfError) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterOvfOvfError) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterOvfOvfError) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterOvfOvfError) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *VcenterOvfOvfError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VcenterOvfOvfError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VcenterOvfOvfError) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VcenterOvfOvfError) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMessage

`func (o *VcenterOvfOvfError) GetMessage() VapiStdLocalizableMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VcenterOvfOvfError) GetMessageOk() (*VapiStdLocalizableMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VcenterOvfOvfError) SetMessage(v VapiStdLocalizableMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VcenterOvfOvfError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetError

`func (o *VcenterOvfOvfError) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VcenterOvfOvfError) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VcenterOvfOvfError) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *VcenterOvfOvfError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


