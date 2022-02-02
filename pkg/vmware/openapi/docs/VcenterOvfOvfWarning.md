# VcenterOvfOvfWarning

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

### NewVcenterOvfOvfWarning

`func NewVcenterOvfOvfWarning(category VcenterOvfOvfMessageCategory, ) *VcenterOvfOvfWarning`

NewVcenterOvfOvfWarning instantiates a new VcenterOvfOvfWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfOvfWarningWithDefaults

`func NewVcenterOvfOvfWarningWithDefaults() *VcenterOvfOvfWarning`

NewVcenterOvfOvfWarningWithDefaults instantiates a new VcenterOvfOvfWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *VcenterOvfOvfWarning) GetCategory() VcenterOvfOvfMessageCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *VcenterOvfOvfWarning) GetCategoryOk() (*VcenterOvfOvfMessageCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *VcenterOvfOvfWarning) SetCategory(v VcenterOvfOvfMessageCategory)`

SetCategory sets Category field to given value.


### GetIssues

`func (o *VcenterOvfOvfWarning) GetIssues() []VcenterOvfParseIssue`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *VcenterOvfOvfWarning) GetIssuesOk() (*[]VcenterOvfParseIssue, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *VcenterOvfOvfWarning) SetIssues(v []VcenterOvfParseIssue)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *VcenterOvfOvfWarning) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetName

`func (o *VcenterOvfOvfWarning) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterOvfOvfWarning) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterOvfOvfWarning) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterOvfOvfWarning) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *VcenterOvfOvfWarning) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VcenterOvfOvfWarning) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VcenterOvfOvfWarning) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VcenterOvfOvfWarning) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMessage

`func (o *VcenterOvfOvfWarning) GetMessage() VapiStdLocalizableMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VcenterOvfOvfWarning) GetMessageOk() (*VapiStdLocalizableMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VcenterOvfOvfWarning) SetMessage(v VapiStdLocalizableMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VcenterOvfOvfWarning) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetError

`func (o *VcenterOvfOvfWarning) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VcenterOvfOvfWarning) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VcenterOvfOvfWarning) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *VcenterOvfOvfWarning) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


