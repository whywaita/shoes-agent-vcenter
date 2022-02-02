# Oauth2ErrorsInvalidScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**Oauth2ErrorsErrorType**](Oauth2ErrorsErrorType.md) |  | 
**ErrorDescription** | Pointer to **string** | Human-readable ASCII text providing additional information, used to assist the client developer in understanding the error that occurred. Values for the \&quot;error_description\&quot; parameter MUST NOT include characters outside the set %x20-21 / %x23-5B / %x5D-7E. if no additional information is available. | [optional] 
**ErrorUri** | Pointer to **string** | A URI identifying a human-readable web page with information about the error, used to provide the client developer with additional information about the error. if no such web-page is available. | [optional] 

## Methods

### NewOauth2ErrorsInvalidScope

`func NewOauth2ErrorsInvalidScope(error_ Oauth2ErrorsErrorType, ) *Oauth2ErrorsInvalidScope`

NewOauth2ErrorsInvalidScope instantiates a new Oauth2ErrorsInvalidScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2ErrorsInvalidScopeWithDefaults

`func NewOauth2ErrorsInvalidScopeWithDefaults() *Oauth2ErrorsInvalidScope`

NewOauth2ErrorsInvalidScopeWithDefaults instantiates a new Oauth2ErrorsInvalidScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *Oauth2ErrorsInvalidScope) GetError() Oauth2ErrorsErrorType`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Oauth2ErrorsInvalidScope) GetErrorOk() (*Oauth2ErrorsErrorType, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Oauth2ErrorsInvalidScope) SetError(v Oauth2ErrorsErrorType)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *Oauth2ErrorsInvalidScope) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *Oauth2ErrorsInvalidScope) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *Oauth2ErrorsInvalidScope) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *Oauth2ErrorsInvalidScope) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorUri

`func (o *Oauth2ErrorsInvalidScope) GetErrorUri() string`

GetErrorUri returns the ErrorUri field if non-nil, zero value otherwise.

### GetErrorUriOk

`func (o *Oauth2ErrorsInvalidScope) GetErrorUriOk() (*string, bool)`

GetErrorUriOk returns a tuple with the ErrorUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUri

`func (o *Oauth2ErrorsInvalidScope) SetErrorUri(v string)`

SetErrorUri sets ErrorUri field to given value.

### HasErrorUri

`func (o *Oauth2ErrorsInvalidScope) HasErrorUri() bool`

HasErrorUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


