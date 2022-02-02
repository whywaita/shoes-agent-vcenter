# Oauth2ErrorsInvalidRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**Oauth2ErrorsErrorType**](Oauth2ErrorsErrorType.md) |  | 
**ErrorDescription** | Pointer to **string** | Human-readable ASCII text providing additional information, used to assist the client developer in understanding the error that occurred. Values for the \&quot;error_description\&quot; parameter MUST NOT include characters outside the set %x20-21 / %x23-5B / %x5D-7E. if no additional information is available. | [optional] 
**ErrorUri** | Pointer to **string** | A URI identifying a human-readable web page with information about the error, used to provide the client developer with additional information about the error. if no such web-page is available. | [optional] 

## Methods

### NewOauth2ErrorsInvalidRequest

`func NewOauth2ErrorsInvalidRequest(error_ Oauth2ErrorsErrorType, ) *Oauth2ErrorsInvalidRequest`

NewOauth2ErrorsInvalidRequest instantiates a new Oauth2ErrorsInvalidRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2ErrorsInvalidRequestWithDefaults

`func NewOauth2ErrorsInvalidRequestWithDefaults() *Oauth2ErrorsInvalidRequest`

NewOauth2ErrorsInvalidRequestWithDefaults instantiates a new Oauth2ErrorsInvalidRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *Oauth2ErrorsInvalidRequest) GetError() Oauth2ErrorsErrorType`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Oauth2ErrorsInvalidRequest) GetErrorOk() (*Oauth2ErrorsErrorType, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Oauth2ErrorsInvalidRequest) SetError(v Oauth2ErrorsErrorType)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *Oauth2ErrorsInvalidRequest) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *Oauth2ErrorsInvalidRequest) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *Oauth2ErrorsInvalidRequest) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *Oauth2ErrorsInvalidRequest) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorUri

`func (o *Oauth2ErrorsInvalidRequest) GetErrorUri() string`

GetErrorUri returns the ErrorUri field if non-nil, zero value otherwise.

### GetErrorUriOk

`func (o *Oauth2ErrorsInvalidRequest) GetErrorUriOk() (*string, bool)`

GetErrorUriOk returns a tuple with the ErrorUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUri

`func (o *Oauth2ErrorsInvalidRequest) SetErrorUri(v string)`

SetErrorUri sets ErrorUri field to given value.

### HasErrorUri

`func (o *Oauth2ErrorsInvalidRequest) HasErrorUri() bool`

HasErrorUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


