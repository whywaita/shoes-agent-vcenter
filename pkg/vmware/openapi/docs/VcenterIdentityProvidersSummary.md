# VcenterIdentityProvidersSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | The identifier of the provider When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.identity.Providers. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.identity.Providers. | 
**Name** | Pointer to **string** | The user friendly name for the provider This field is optional because it was added in a newer version than its parent node. | [optional] 
**ConfigTag** | [**VcenterIdentityProvidersConfigType**](VcenterIdentityProvidersConfigType.md) |  | 
**Oauth2** | Pointer to [**VcenterIdentityProvidersOauth2Summary**](VcenterIdentityProvidersOauth2Summary.md) |  | [optional] 
**Oidc** | Pointer to [**VcenterIdentityProvidersOidcSummary**](VcenterIdentityProvidersOidcSummary.md) |  | [optional] 
**IsDefault** | **bool** | Specifies whether the provider is the default provider. | 
**DomainNames** | Pointer to **[]string** | Set of fully qualified domain names to trust when federating with this identity provider. Tokens from this identity provider will only be validated if the user belongs to one of these domains, and any domain-qualified groups in the tokens will be filtered to include only those groups that belong to one of these domains. If domainNames is an empty set, domain validation behavior at login with this identity provider will be as follows: the user&#39;s domain will be parsed from the User Principal Name (UPN) value that is found in the tokens returned by the identity provider. This domain will then be implicitly trusted and used to filter any groups that are also provided in the tokens. This field is optional because it was added in a newer version than its parent node. | [optional] 
**AuthQueryParams** | Pointer to [**[]VcenterIdentityProvidersCreateSpecAuthQueryParams**](VcenterIdentityProvidersCreateSpecAuthQueryParams.md) |  key/value pairs that are to be appended to the authEndpoint request.   How to append to authEndpoint request:  If the map is not empty, a \&quot;?\&quot; is added to the endpoint URL, and combination of each k and each string in the v is added with an \&quot;&amp;\&quot; delimiter. Details:    - If the value contains only one string, then the key is added with \&quot;k&#x3D;v\&quot;.    - If the value is an empty list, then the key is added without a \&quot;&#x3D;v\&quot;.    - If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.  This field is optional because it was added in a newer version than its parent node. | [optional] 

## Methods

### NewVcenterIdentityProvidersSummary

`func NewVcenterIdentityProvidersSummary(provider string, configTag VcenterIdentityProvidersConfigType, isDefault bool, ) *VcenterIdentityProvidersSummary`

NewVcenterIdentityProvidersSummary instantiates a new VcenterIdentityProvidersSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterIdentityProvidersSummaryWithDefaults

`func NewVcenterIdentityProvidersSummaryWithDefaults() *VcenterIdentityProvidersSummary`

NewVcenterIdentityProvidersSummaryWithDefaults instantiates a new VcenterIdentityProvidersSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *VcenterIdentityProvidersSummary) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VcenterIdentityProvidersSummary) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VcenterIdentityProvidersSummary) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetName

`func (o *VcenterIdentityProvidersSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterIdentityProvidersSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterIdentityProvidersSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterIdentityProvidersSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfigTag

`func (o *VcenterIdentityProvidersSummary) GetConfigTag() VcenterIdentityProvidersConfigType`

GetConfigTag returns the ConfigTag field if non-nil, zero value otherwise.

### GetConfigTagOk

`func (o *VcenterIdentityProvidersSummary) GetConfigTagOk() (*VcenterIdentityProvidersConfigType, bool)`

GetConfigTagOk returns a tuple with the ConfigTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTag

`func (o *VcenterIdentityProvidersSummary) SetConfigTag(v VcenterIdentityProvidersConfigType)`

SetConfigTag sets ConfigTag field to given value.


### GetOauth2

`func (o *VcenterIdentityProvidersSummary) GetOauth2() VcenterIdentityProvidersOauth2Summary`

GetOauth2 returns the Oauth2 field if non-nil, zero value otherwise.

### GetOauth2Ok

`func (o *VcenterIdentityProvidersSummary) GetOauth2Ok() (*VcenterIdentityProvidersOauth2Summary, bool)`

GetOauth2Ok returns a tuple with the Oauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2

`func (o *VcenterIdentityProvidersSummary) SetOauth2(v VcenterIdentityProvidersOauth2Summary)`

SetOauth2 sets Oauth2 field to given value.

### HasOauth2

`func (o *VcenterIdentityProvidersSummary) HasOauth2() bool`

HasOauth2 returns a boolean if a field has been set.

### GetOidc

`func (o *VcenterIdentityProvidersSummary) GetOidc() VcenterIdentityProvidersOidcSummary`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *VcenterIdentityProvidersSummary) GetOidcOk() (*VcenterIdentityProvidersOidcSummary, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *VcenterIdentityProvidersSummary) SetOidc(v VcenterIdentityProvidersOidcSummary)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *VcenterIdentityProvidersSummary) HasOidc() bool`

HasOidc returns a boolean if a field has been set.

### GetIsDefault

`func (o *VcenterIdentityProvidersSummary) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *VcenterIdentityProvidersSummary) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *VcenterIdentityProvidersSummary) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetDomainNames

`func (o *VcenterIdentityProvidersSummary) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *VcenterIdentityProvidersSummary) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *VcenterIdentityProvidersSummary) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *VcenterIdentityProvidersSummary) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetAuthQueryParams

`func (o *VcenterIdentityProvidersSummary) GetAuthQueryParams() []VcenterIdentityProvidersCreateSpecAuthQueryParams`

GetAuthQueryParams returns the AuthQueryParams field if non-nil, zero value otherwise.

### GetAuthQueryParamsOk

`func (o *VcenterIdentityProvidersSummary) GetAuthQueryParamsOk() (*[]VcenterIdentityProvidersCreateSpecAuthQueryParams, bool)`

GetAuthQueryParamsOk returns a tuple with the AuthQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthQueryParams

`func (o *VcenterIdentityProvidersSummary) SetAuthQueryParams(v []VcenterIdentityProvidersCreateSpecAuthQueryParams)`

SetAuthQueryParams sets AuthQueryParams field to given value.

### HasAuthQueryParams

`func (o *VcenterIdentityProvidersSummary) HasAuthQueryParams() bool`

HasAuthQueryParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


