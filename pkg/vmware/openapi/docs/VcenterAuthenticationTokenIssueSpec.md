# VcenterAuthenticationTokenIssueSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** | The value of urn:ietf:params:oauth:grant-type:token-exchange indicates that a token exchange is being performed. When clients pass a value of this structure as a parameter, the field must be one of urn:ietf:params:oauth:grant-type:token-exchange. When operations return a value of this structure as a result, the field will be one of urn:ietf:params:oauth:grant-type:token-exchange. | 
**Resource** | Pointer to **string** | Indicates the location of the target service or resource where the client intends to use the requested security token. If unset, it is inferred from other arguments. | [optional] 
**Audience** | Pointer to **string** | The logical name of the target service where the client intends to use the requested security token. This serves a purpose similar to the Token.IssueSpec.resource field, but with the client providing a logical name rather than a location. If unset, it is inferred from other arguments. | [optional] 
**Scope** | Pointer to **string** | A list of space-delimited, case-sensitive strings, that allow the client to specify the desired scope of the requested security token in the context of the service or resource where the token will be used. If unset, it is inferred from other arguments. | [optional] 
**RequestedTokenType** | Pointer to **string** | An identifier for the type of the requested security token. If the requested type is unspecified, the issued token type is at the discretion of the server and may be dictated by knowledge of the requirements of the service or resource indicated by the Token.IssueSpec.resource or Token.IssueSpec.audience field. If unset, it is inferred from other arguments. | [optional] 
**SubjectToken** | Pointer to **string** | A security token that represents the identity of the party on behalf of whom exchange is being made. Typically, the subject of this token will be the subject of the security token issued. Token is base64-encoded.  The field is required when the value of the Token.IssueSpec.grant-type field is urn:ietf:params:oauth:grant-type:token-exchange.  This field is currently required. In the future, the structure may support grant-types other than urn:ietf:params:oauth:grant-type:token-exchange for which the value may be unset. | [optional] 
**SubjectTokenType** | Pointer to **string** | An identifier, that indicates the type of the security token in the Token.IssueSpec.subject-token field.  The field is required when the value of the Token.IssueSpec.grant-type field is urn:ietf:params:oauth:grant-type:token-exchange.  This field is currently required. In the future, the structure may support grant-types other than urn:ietf:params:oauth:grant-type:token-exchange for which the value may be unset. | [optional] 
**ActorToken** | Pointer to **string** | A security token that represents the identity of the acting party. Typically, this will be the party that is authorized to use the requested security token and act on behalf of the subject. unset if not needed for the specific case of exchange. | [optional] 
**ActorTokenType** | Pointer to **string** | An identifier, that indicates the type of the security token in the Token.IssueSpec.actor-token field. unset if Token.IssueSpec.actor-token field is not present. | [optional] 

## Methods

### NewVcenterAuthenticationTokenIssueSpec

`func NewVcenterAuthenticationTokenIssueSpec(grantType string, ) *VcenterAuthenticationTokenIssueSpec`

NewVcenterAuthenticationTokenIssueSpec instantiates a new VcenterAuthenticationTokenIssueSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterAuthenticationTokenIssueSpecWithDefaults

`func NewVcenterAuthenticationTokenIssueSpecWithDefaults() *VcenterAuthenticationTokenIssueSpec`

NewVcenterAuthenticationTokenIssueSpecWithDefaults instantiates a new VcenterAuthenticationTokenIssueSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *VcenterAuthenticationTokenIssueSpec) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *VcenterAuthenticationTokenIssueSpec) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *VcenterAuthenticationTokenIssueSpec) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetResource

`func (o *VcenterAuthenticationTokenIssueSpec) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *VcenterAuthenticationTokenIssueSpec) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *VcenterAuthenticationTokenIssueSpec) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *VcenterAuthenticationTokenIssueSpec) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetAudience

`func (o *VcenterAuthenticationTokenIssueSpec) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *VcenterAuthenticationTokenIssueSpec) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *VcenterAuthenticationTokenIssueSpec) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *VcenterAuthenticationTokenIssueSpec) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetScope

`func (o *VcenterAuthenticationTokenIssueSpec) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VcenterAuthenticationTokenIssueSpec) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VcenterAuthenticationTokenIssueSpec) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *VcenterAuthenticationTokenIssueSpec) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetRequestedTokenType

`func (o *VcenterAuthenticationTokenIssueSpec) GetRequestedTokenType() string`

GetRequestedTokenType returns the RequestedTokenType field if non-nil, zero value otherwise.

### GetRequestedTokenTypeOk

`func (o *VcenterAuthenticationTokenIssueSpec) GetRequestedTokenTypeOk() (*string, bool)`

GetRequestedTokenTypeOk returns a tuple with the RequestedTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTokenType

`func (o *VcenterAuthenticationTokenIssueSpec) SetRequestedTokenType(v string)`

SetRequestedTokenType sets RequestedTokenType field to given value.

### HasRequestedTokenType

`func (o *VcenterAuthenticationTokenIssueSpec) HasRequestedTokenType() bool`

HasRequestedTokenType returns a boolean if a field has been set.

### GetSubjectToken

`func (o *VcenterAuthenticationTokenIssueSpec) GetSubjectToken() string`

GetSubjectToken returns the SubjectToken field if non-nil, zero value otherwise.

### GetSubjectTokenOk

`func (o *VcenterAuthenticationTokenIssueSpec) GetSubjectTokenOk() (*string, bool)`

GetSubjectTokenOk returns a tuple with the SubjectToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectToken

`func (o *VcenterAuthenticationTokenIssueSpec) SetSubjectToken(v string)`

SetSubjectToken sets SubjectToken field to given value.

### HasSubjectToken

`func (o *VcenterAuthenticationTokenIssueSpec) HasSubjectToken() bool`

HasSubjectToken returns a boolean if a field has been set.

### GetSubjectTokenType

`func (o *VcenterAuthenticationTokenIssueSpec) GetSubjectTokenType() string`

GetSubjectTokenType returns the SubjectTokenType field if non-nil, zero value otherwise.

### GetSubjectTokenTypeOk

`func (o *VcenterAuthenticationTokenIssueSpec) GetSubjectTokenTypeOk() (*string, bool)`

GetSubjectTokenTypeOk returns a tuple with the SubjectTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTokenType

`func (o *VcenterAuthenticationTokenIssueSpec) SetSubjectTokenType(v string)`

SetSubjectTokenType sets SubjectTokenType field to given value.

### HasSubjectTokenType

`func (o *VcenterAuthenticationTokenIssueSpec) HasSubjectTokenType() bool`

HasSubjectTokenType returns a boolean if a field has been set.

### GetActorToken

`func (o *VcenterAuthenticationTokenIssueSpec) GetActorToken() string`

GetActorToken returns the ActorToken field if non-nil, zero value otherwise.

### GetActorTokenOk

`func (o *VcenterAuthenticationTokenIssueSpec) GetActorTokenOk() (*string, bool)`

GetActorTokenOk returns a tuple with the ActorToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorToken

`func (o *VcenterAuthenticationTokenIssueSpec) SetActorToken(v string)`

SetActorToken sets ActorToken field to given value.

### HasActorToken

`func (o *VcenterAuthenticationTokenIssueSpec) HasActorToken() bool`

HasActorToken returns a boolean if a field has been set.

### GetActorTokenType

`func (o *VcenterAuthenticationTokenIssueSpec) GetActorTokenType() string`

GetActorTokenType returns the ActorTokenType field if non-nil, zero value otherwise.

### GetActorTokenTypeOk

`func (o *VcenterAuthenticationTokenIssueSpec) GetActorTokenTypeOk() (*string, bool)`

GetActorTokenTypeOk returns a tuple with the ActorTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorTokenType

`func (o *VcenterAuthenticationTokenIssueSpec) SetActorTokenType(v string)`

SetActorTokenType sets ActorTokenType field to given value.

### HasActorTokenType

`func (o *VcenterAuthenticationTokenIssueSpec) HasActorTokenType() bool`

HasActorTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


