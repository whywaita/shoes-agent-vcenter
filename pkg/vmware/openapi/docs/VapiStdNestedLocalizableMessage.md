# VapiStdNestedLocalizableMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the localizable string or message template. &lt;p&gt; This identifier is typically used to retrieve a locale-specific string or message template from a message catalog. | 
**Params** | Pointer to [**[]VapiStdLocalizableMessageParams**](VapiStdLocalizableMessageParams.md) | Named Arguments to be substituted into the message template. | [optional] 

## Methods

### NewVapiStdNestedLocalizableMessage

`func NewVapiStdNestedLocalizableMessage(id string, ) *VapiStdNestedLocalizableMessage`

NewVapiStdNestedLocalizableMessage instantiates a new VapiStdNestedLocalizableMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVapiStdNestedLocalizableMessageWithDefaults

`func NewVapiStdNestedLocalizableMessageWithDefaults() *VapiStdNestedLocalizableMessage`

NewVapiStdNestedLocalizableMessageWithDefaults instantiates a new VapiStdNestedLocalizableMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VapiStdNestedLocalizableMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VapiStdNestedLocalizableMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VapiStdNestedLocalizableMessage) SetId(v string)`

SetId sets Id field to given value.


### GetParams

`func (o *VapiStdNestedLocalizableMessage) GetParams() []VapiStdLocalizableMessageParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *VapiStdNestedLocalizableMessage) GetParamsOk() (*[]VapiStdLocalizableMessageParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *VapiStdNestedLocalizableMessage) SetParams(v []VapiStdLocalizableMessageParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *VapiStdNestedLocalizableMessage) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


