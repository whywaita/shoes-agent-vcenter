# VcenterVmGuestCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InteractiveSession** | **bool** | If set, the operation will interact with the logged-in desktop session in the guest. This requires that the logged-on user matches the user specified by the Credentials. This is currently only supported for USERNAME_PASSWORD. | 
**Type** | [**VcenterVmGuestCredentialsType**](VcenterVmGuestCredentialsType.md) |  | 
**UserName** | Pointer to **string** | For SAML_BEARER_TOKEN, this is the guest user to be associated with the credentials. For USERNAME_PASSWORD this is the guest username. If no user is specified for SAML_BEARER_TOKEN, a guest dependent mapping will decide what guest user account is applied. | [optional] 
**Password** | Pointer to **string** | password This field is optional and it is only relevant when the value of Credentials.type is USERNAME_PASSWORD. | [optional] 
**SamlToken** | Pointer to **string** | SAML Bearer Token This field is optional and it is only relevant when the value of Credentials.type is SAML_BEARER_TOKEN. | [optional] 

## Methods

### NewVcenterVmGuestCredentials

`func NewVcenterVmGuestCredentials(interactiveSession bool, type_ VcenterVmGuestCredentialsType, ) *VcenterVmGuestCredentials`

NewVcenterVmGuestCredentials instantiates a new VcenterVmGuestCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestCredentialsWithDefaults

`func NewVcenterVmGuestCredentialsWithDefaults() *VcenterVmGuestCredentials`

NewVcenterVmGuestCredentialsWithDefaults instantiates a new VcenterVmGuestCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInteractiveSession

`func (o *VcenterVmGuestCredentials) GetInteractiveSession() bool`

GetInteractiveSession returns the InteractiveSession field if non-nil, zero value otherwise.

### GetInteractiveSessionOk

`func (o *VcenterVmGuestCredentials) GetInteractiveSessionOk() (*bool, bool)`

GetInteractiveSessionOk returns a tuple with the InteractiveSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractiveSession

`func (o *VcenterVmGuestCredentials) SetInteractiveSession(v bool)`

SetInteractiveSession sets InteractiveSession field to given value.


### GetType

`func (o *VcenterVmGuestCredentials) GetType() VcenterVmGuestCredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmGuestCredentials) GetTypeOk() (*VcenterVmGuestCredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmGuestCredentials) SetType(v VcenterVmGuestCredentialsType)`

SetType sets Type field to given value.


### GetUserName

`func (o *VcenterVmGuestCredentials) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *VcenterVmGuestCredentials) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *VcenterVmGuestCredentials) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *VcenterVmGuestCredentials) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *VcenterVmGuestCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterVmGuestCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterVmGuestCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VcenterVmGuestCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSamlToken

`func (o *VcenterVmGuestCredentials) GetSamlToken() string`

GetSamlToken returns the SamlToken field if non-nil, zero value otherwise.

### GetSamlTokenOk

`func (o *VcenterVmGuestCredentials) GetSamlTokenOk() (*string, bool)`

GetSamlTokenOk returns a tuple with the SamlToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlToken

`func (o *VcenterVmGuestCredentials) SetSamlToken(v string)`

SetSamlToken sets SamlToken field to given value.

### HasSamlToken

`func (o *VcenterVmGuestCredentials) HasSamlToken() bool`

HasSamlToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


