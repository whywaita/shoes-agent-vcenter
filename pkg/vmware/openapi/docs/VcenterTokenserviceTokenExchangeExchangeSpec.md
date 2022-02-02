# VcenterTokenserviceTokenExchangeExchangeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** | The value of {@link TokenExchange#TOKEN_EXCHANGE_GRANT} indicates that a token exchange is being performed. | 
**Resource** | Pointer to **string** | Indicates the location of the target service or resource where the client intends to use the requested security token. | [optional] 
**Audience** | Pointer to **string** | The logical name of the target service where the client intends to use the requested security token. This serves a purpose similar to the {@link ExchangeSpec#resource} parameter, but with the client providing a logical name rather than a location. | [optional] 
**Scope** | Pointer to **string** | A list of space-delimited, case-sensitive strings, that allow the client to specify the desired scope of the requested security token in the context of the service or resource where the token will be used. | [optional] 
**RequestedTokenType** | Pointer to **string** | An identifier for the type of the requested security token. If the requested type is unspecified, the issued token type is at the discretion of the server and may be dictated by knowledge of the requirements of the service or resource indicated by the {@link ExchangeSpec#resource} or {@link ExchangeSpec#audience} parameter. | [optional] 
**SubjectToken** | **string** | A security token that represents the identity of the party on behalf of whom exchange is being made. Typically, the subject of this token will be the subject of the security token issued. Token is base64-encoded. | 
**SubjectTokenType** | **string** | An identifier, that indicates the type of the security token in the {@link ExchangeSpec#subject_token} parameter. | 
**ActorToken** | Pointer to **string** | A security token that represents the identity of the acting party. Typically, this will be the party that is authorized to use the requested security token and act on behalf of the subject. | [optional] 
**ActorTokenType** | Pointer to **string** | An identifier, that indicates the type of the security token in the {@link ExchangeSpec#actor_token} parameter. | [optional] 

## Methods

### NewVcenterTokenserviceTokenExchangeExchangeSpec

`func NewVcenterTokenserviceTokenExchangeExchangeSpec(grantType string, subjectToken string, subjectTokenType string, ) *VcenterTokenserviceTokenExchangeExchangeSpec`

NewVcenterTokenserviceTokenExchangeExchangeSpec instantiates a new VcenterTokenserviceTokenExchangeExchangeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTokenserviceTokenExchangeExchangeSpecWithDefaults

`func NewVcenterTokenserviceTokenExchangeExchangeSpecWithDefaults() *VcenterTokenserviceTokenExchangeExchangeSpec`

NewVcenterTokenserviceTokenExchangeExchangeSpecWithDefaults instantiates a new VcenterTokenserviceTokenExchangeExchangeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetResource

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetAudience

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetScope

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetRequestedTokenType

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetRequestedTokenType() string`

GetRequestedTokenType returns the RequestedTokenType field if non-nil, zero value otherwise.

### GetRequestedTokenTypeOk

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetRequestedTokenTypeOk() (*string, bool)`

GetRequestedTokenTypeOk returns a tuple with the RequestedTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTokenType

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) SetRequestedTokenType(v string)`

SetRequestedTokenType sets RequestedTokenType field to given value.

### HasRequestedTokenType

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) HasRequestedTokenType() bool`

HasRequestedTokenType returns a boolean if a field has been set.

### GetSubjectToken

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetSubjectToken() string`

GetSubjectToken returns the SubjectToken field if non-nil, zero value otherwise.

### GetSubjectTokenOk

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetSubjectTokenOk() (*string, bool)`

GetSubjectTokenOk returns a tuple with the SubjectToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectToken

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) SetSubjectToken(v string)`

SetSubjectToken sets SubjectToken field to given value.


### GetSubjectTokenType

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetSubjectTokenType() string`

GetSubjectTokenType returns the SubjectTokenType field if non-nil, zero value otherwise.

### GetSubjectTokenTypeOk

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetSubjectTokenTypeOk() (*string, bool)`

GetSubjectTokenTypeOk returns a tuple with the SubjectTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTokenType

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) SetSubjectTokenType(v string)`

SetSubjectTokenType sets SubjectTokenType field to given value.


### GetActorToken

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetActorToken() string`

GetActorToken returns the ActorToken field if non-nil, zero value otherwise.

### GetActorTokenOk

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetActorTokenOk() (*string, bool)`

GetActorTokenOk returns a tuple with the ActorToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorToken

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) SetActorToken(v string)`

SetActorToken sets ActorToken field to given value.

### HasActorToken

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) HasActorToken() bool`

HasActorToken returns a boolean if a field has been set.

### GetActorTokenType

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetActorTokenType() string`

GetActorTokenType returns the ActorTokenType field if non-nil, zero value otherwise.

### GetActorTokenTypeOk

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) GetActorTokenTypeOk() (*string, bool)`

GetActorTokenTypeOk returns a tuple with the ActorTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorTokenType

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) SetActorTokenType(v string)`

SetActorTokenType sets ActorTokenType field to given value.

### HasActorTokenType

`func (o *VcenterTokenserviceTokenExchangeExchangeSpec) HasActorTokenType() bool`

HasActorTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


