# VapiStdLocalizableMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the localizable string or message template. &lt;p&gt; This identifier is typically used to retrieve a locale-specific string or message template from a message catalog. | 
**DefaultMessage** | **string** | The value of this localizable string or message template in the {@code en_US} (English) locale.  If {@link #id} refers to a message template, the default message will contain the substituted arguments. This value can be used by clients that do not need to display strings and messages in the native language of the user.  It could also be used as a fallback if a client is unable to access the appropriate message catalog. | 
**Args** | **[]string** | Positional arguments to be substituted into the message template. This list will be empty if the message uses named arguments or has no arguments. | 
**Params** | Pointer to [**[]VapiStdLocalizableMessageParams**](VapiStdLocalizableMessageParams.md) | Named arguments to be substituted into the message template. | [optional] 
**Localized** | Pointer to **string** | Localized string value as per request requirements. | [optional] 

## Methods

### NewVapiStdLocalizableMessage

`func NewVapiStdLocalizableMessage(id string, defaultMessage string, args []string, ) *VapiStdLocalizableMessage`

NewVapiStdLocalizableMessage instantiates a new VapiStdLocalizableMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVapiStdLocalizableMessageWithDefaults

`func NewVapiStdLocalizableMessageWithDefaults() *VapiStdLocalizableMessage`

NewVapiStdLocalizableMessageWithDefaults instantiates a new VapiStdLocalizableMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VapiStdLocalizableMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VapiStdLocalizableMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VapiStdLocalizableMessage) SetId(v string)`

SetId sets Id field to given value.


### GetDefaultMessage

`func (o *VapiStdLocalizableMessage) GetDefaultMessage() string`

GetDefaultMessage returns the DefaultMessage field if non-nil, zero value otherwise.

### GetDefaultMessageOk

`func (o *VapiStdLocalizableMessage) GetDefaultMessageOk() (*string, bool)`

GetDefaultMessageOk returns a tuple with the DefaultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMessage

`func (o *VapiStdLocalizableMessage) SetDefaultMessage(v string)`

SetDefaultMessage sets DefaultMessage field to given value.


### GetArgs

`func (o *VapiStdLocalizableMessage) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *VapiStdLocalizableMessage) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *VapiStdLocalizableMessage) SetArgs(v []string)`

SetArgs sets Args field to given value.


### GetParams

`func (o *VapiStdLocalizableMessage) GetParams() []VapiStdLocalizableMessageParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *VapiStdLocalizableMessage) GetParamsOk() (*[]VapiStdLocalizableMessageParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *VapiStdLocalizableMessage) SetParams(v []VapiStdLocalizableMessageParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *VapiStdLocalizableMessage) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetLocalized

`func (o *VapiStdLocalizableMessage) GetLocalized() string`

GetLocalized returns the Localized field if non-nil, zero value otherwise.

### GetLocalizedOk

`func (o *VapiStdLocalizableMessage) GetLocalizedOk() (*string, bool)`

GetLocalizedOk returns a tuple with the Localized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalized

`func (o *VapiStdLocalizableMessage) SetLocalized(v string)`

SetLocalized sets Localized field to given value.

### HasLocalized

`func (o *VapiStdLocalizableMessage) HasLocalized() bool`

HasLocalized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


