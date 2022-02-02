# VcenterIdentityProvidersCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigTag** | [**VcenterIdentityProvidersConfigType**](VcenterIdentityProvidersConfigType.md) |  | 
**Oauth2** | Pointer to [**VcenterIdentityProvidersOauth2CreateSpec**](VcenterIdentityProvidersOauth2CreateSpec.md) |  | [optional] 
**Oidc** | Pointer to [**VcenterIdentityProvidersOidcCreateSpec**](VcenterIdentityProvidersOidcCreateSpec.md) |  | [optional] 
**OrgIds** | Pointer to **[]string** | The set of orgIds as part of SDDC creation which provides the basis for tenancy If unset, the set will be empty. | [optional] 
**IsDefault** | Pointer to **bool** | Specifies whether the provider is the default provider. Setting Providers.CreateSpec.is-default of current provider to True makes all other providers non-default. If no other providers created in this vCenter Server before, this parameter will be disregarded, and the provider will always be set to the default. If unset the provider will be the default provider if it is the first provider that is created, and will not be the default provider otherwise. | [optional] 
**Name** | Pointer to **string** | The user friendly name for the provider. This name can be used for human-readable identification purposes, but it does not have to be unique, as the system will use internal UUIDs to differentiate providers. If unset, the name will be the empty string | [optional] 
**DomainNames** | Pointer to **[]string** | Set of fully qualified domain names to trust when federating with this identity provider. Tokens from this identity provider will only be validated if the user belongs to one of these domains, and any domain-qualified groups in the tokens will be filtered to include only those groups that belong to one of these domains. If unset, domainNames will be the empty set and the domain validation behavior at login with this identity provider will be as follows: the user&#39;s domain will be parsed from the User Principal Name (UPN) value that is found in the tokens returned by the identity provider. This domain will then be implicitly trusted and used to filter any groups that are also provided in the tokens. | [optional] 
**AuthQueryParams** | Pointer to [**[]VcenterIdentityProvidersCreateSpecAuthQueryParams**](VcenterIdentityProvidersCreateSpecAuthQueryParams.md) |  key/value pairs that are to be appended to the authEndpoint request.   How to append to authEndpoint request:  If the map is not empty, a \&quot;?\&quot; is added to the endpoint URL, and combination of each k and each string in the v is added with an \&quot;&amp;\&quot; delimiter. Details:    - If the value contains only one string, then the key is added with \&quot;k&#x3D;v\&quot;.    - If the value is an empty list, then the key is added without a \&quot;&#x3D;v\&quot;.    - If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.  If unset, the map will be empty. | [optional] 
**IdmProtocol** | Pointer to [**VcenterIdentityProvidersIdmProtocol**](VcenterIdentityProvidersIdmProtocol.md) |  | [optional] 
**IdmEndpoints** | Pointer to **[]string** | Identity management endpoints. When specified, at least one endpoint must be provided. This field is optional and it is only relevant when the value of Providers.CreateSpec.idm-protocol is one of REST, SCIM, or SCIM2_0. | [optional] 
**ActiveDirectoryOverLdap** | Pointer to [**VcenterIdentityProvidersActiveDirectoryOverLdap**](VcenterIdentityProvidersActiveDirectoryOverLdap.md) |  | [optional] 
**UpnClaim** | Pointer to **string** | Specifies which claim provides the user principal name (UPN) for the user. If unset, the claim named &#39;acct&#39; will be used to provide backwards compatibility with CSP. | [optional] 
**GroupsClaim** | Pointer to **string** | Specifies which claim provides the group membership for the token subject. These groups will be used for mapping to local groups per the claim map. If unset, the default behavior will be CSP backwards compatiblility. The groups for the subject will be comprised of the groups in &#39;group_names&#39; and &#39;group_ids&#39; claims. | [optional] 

## Methods

### NewVcenterIdentityProvidersCreateSpec

`func NewVcenterIdentityProvidersCreateSpec(configTag VcenterIdentityProvidersConfigType, ) *VcenterIdentityProvidersCreateSpec`

NewVcenterIdentityProvidersCreateSpec instantiates a new VcenterIdentityProvidersCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterIdentityProvidersCreateSpecWithDefaults

`func NewVcenterIdentityProvidersCreateSpecWithDefaults() *VcenterIdentityProvidersCreateSpec`

NewVcenterIdentityProvidersCreateSpecWithDefaults instantiates a new VcenterIdentityProvidersCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigTag

`func (o *VcenterIdentityProvidersCreateSpec) GetConfigTag() VcenterIdentityProvidersConfigType`

GetConfigTag returns the ConfigTag field if non-nil, zero value otherwise.

### GetConfigTagOk

`func (o *VcenterIdentityProvidersCreateSpec) GetConfigTagOk() (*VcenterIdentityProvidersConfigType, bool)`

GetConfigTagOk returns a tuple with the ConfigTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTag

`func (o *VcenterIdentityProvidersCreateSpec) SetConfigTag(v VcenterIdentityProvidersConfigType)`

SetConfigTag sets ConfigTag field to given value.


### GetOauth2

`func (o *VcenterIdentityProvidersCreateSpec) GetOauth2() VcenterIdentityProvidersOauth2CreateSpec`

GetOauth2 returns the Oauth2 field if non-nil, zero value otherwise.

### GetOauth2Ok

`func (o *VcenterIdentityProvidersCreateSpec) GetOauth2Ok() (*VcenterIdentityProvidersOauth2CreateSpec, bool)`

GetOauth2Ok returns a tuple with the Oauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2

`func (o *VcenterIdentityProvidersCreateSpec) SetOauth2(v VcenterIdentityProvidersOauth2CreateSpec)`

SetOauth2 sets Oauth2 field to given value.

### HasOauth2

`func (o *VcenterIdentityProvidersCreateSpec) HasOauth2() bool`

HasOauth2 returns a boolean if a field has been set.

### GetOidc

`func (o *VcenterIdentityProvidersCreateSpec) GetOidc() VcenterIdentityProvidersOidcCreateSpec`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *VcenterIdentityProvidersCreateSpec) GetOidcOk() (*VcenterIdentityProvidersOidcCreateSpec, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *VcenterIdentityProvidersCreateSpec) SetOidc(v VcenterIdentityProvidersOidcCreateSpec)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *VcenterIdentityProvidersCreateSpec) HasOidc() bool`

HasOidc returns a boolean if a field has been set.

### GetOrgIds

`func (o *VcenterIdentityProvidersCreateSpec) GetOrgIds() []string`

GetOrgIds returns the OrgIds field if non-nil, zero value otherwise.

### GetOrgIdsOk

`func (o *VcenterIdentityProvidersCreateSpec) GetOrgIdsOk() (*[]string, bool)`

GetOrgIdsOk returns a tuple with the OrgIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgIds

`func (o *VcenterIdentityProvidersCreateSpec) SetOrgIds(v []string)`

SetOrgIds sets OrgIds field to given value.

### HasOrgIds

`func (o *VcenterIdentityProvidersCreateSpec) HasOrgIds() bool`

HasOrgIds returns a boolean if a field has been set.

### GetIsDefault

`func (o *VcenterIdentityProvidersCreateSpec) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *VcenterIdentityProvidersCreateSpec) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *VcenterIdentityProvidersCreateSpec) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *VcenterIdentityProvidersCreateSpec) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *VcenterIdentityProvidersCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterIdentityProvidersCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterIdentityProvidersCreateSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterIdentityProvidersCreateSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomainNames

`func (o *VcenterIdentityProvidersCreateSpec) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *VcenterIdentityProvidersCreateSpec) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *VcenterIdentityProvidersCreateSpec) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *VcenterIdentityProvidersCreateSpec) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetAuthQueryParams

`func (o *VcenterIdentityProvidersCreateSpec) GetAuthQueryParams() []VcenterIdentityProvidersCreateSpecAuthQueryParams`

GetAuthQueryParams returns the AuthQueryParams field if non-nil, zero value otherwise.

### GetAuthQueryParamsOk

`func (o *VcenterIdentityProvidersCreateSpec) GetAuthQueryParamsOk() (*[]VcenterIdentityProvidersCreateSpecAuthQueryParams, bool)`

GetAuthQueryParamsOk returns a tuple with the AuthQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthQueryParams

`func (o *VcenterIdentityProvidersCreateSpec) SetAuthQueryParams(v []VcenterIdentityProvidersCreateSpecAuthQueryParams)`

SetAuthQueryParams sets AuthQueryParams field to given value.

### HasAuthQueryParams

`func (o *VcenterIdentityProvidersCreateSpec) HasAuthQueryParams() bool`

HasAuthQueryParams returns a boolean if a field has been set.

### GetIdmProtocol

`func (o *VcenterIdentityProvidersCreateSpec) GetIdmProtocol() VcenterIdentityProvidersIdmProtocol`

GetIdmProtocol returns the IdmProtocol field if non-nil, zero value otherwise.

### GetIdmProtocolOk

`func (o *VcenterIdentityProvidersCreateSpec) GetIdmProtocolOk() (*VcenterIdentityProvidersIdmProtocol, bool)`

GetIdmProtocolOk returns a tuple with the IdmProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdmProtocol

`func (o *VcenterIdentityProvidersCreateSpec) SetIdmProtocol(v VcenterIdentityProvidersIdmProtocol)`

SetIdmProtocol sets IdmProtocol field to given value.

### HasIdmProtocol

`func (o *VcenterIdentityProvidersCreateSpec) HasIdmProtocol() bool`

HasIdmProtocol returns a boolean if a field has been set.

### GetIdmEndpoints

`func (o *VcenterIdentityProvidersCreateSpec) GetIdmEndpoints() []string`

GetIdmEndpoints returns the IdmEndpoints field if non-nil, zero value otherwise.

### GetIdmEndpointsOk

`func (o *VcenterIdentityProvidersCreateSpec) GetIdmEndpointsOk() (*[]string, bool)`

GetIdmEndpointsOk returns a tuple with the IdmEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdmEndpoints

`func (o *VcenterIdentityProvidersCreateSpec) SetIdmEndpoints(v []string)`

SetIdmEndpoints sets IdmEndpoints field to given value.

### HasIdmEndpoints

`func (o *VcenterIdentityProvidersCreateSpec) HasIdmEndpoints() bool`

HasIdmEndpoints returns a boolean if a field has been set.

### GetActiveDirectoryOverLdap

`func (o *VcenterIdentityProvidersCreateSpec) GetActiveDirectoryOverLdap() VcenterIdentityProvidersActiveDirectoryOverLdap`

GetActiveDirectoryOverLdap returns the ActiveDirectoryOverLdap field if non-nil, zero value otherwise.

### GetActiveDirectoryOverLdapOk

`func (o *VcenterIdentityProvidersCreateSpec) GetActiveDirectoryOverLdapOk() (*VcenterIdentityProvidersActiveDirectoryOverLdap, bool)`

GetActiveDirectoryOverLdapOk returns a tuple with the ActiveDirectoryOverLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryOverLdap

`func (o *VcenterIdentityProvidersCreateSpec) SetActiveDirectoryOverLdap(v VcenterIdentityProvidersActiveDirectoryOverLdap)`

SetActiveDirectoryOverLdap sets ActiveDirectoryOverLdap field to given value.

### HasActiveDirectoryOverLdap

`func (o *VcenterIdentityProvidersCreateSpec) HasActiveDirectoryOverLdap() bool`

HasActiveDirectoryOverLdap returns a boolean if a field has been set.

### GetUpnClaim

`func (o *VcenterIdentityProvidersCreateSpec) GetUpnClaim() string`

GetUpnClaim returns the UpnClaim field if non-nil, zero value otherwise.

### GetUpnClaimOk

`func (o *VcenterIdentityProvidersCreateSpec) GetUpnClaimOk() (*string, bool)`

GetUpnClaimOk returns a tuple with the UpnClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpnClaim

`func (o *VcenterIdentityProvidersCreateSpec) SetUpnClaim(v string)`

SetUpnClaim sets UpnClaim field to given value.

### HasUpnClaim

`func (o *VcenterIdentityProvidersCreateSpec) HasUpnClaim() bool`

HasUpnClaim returns a boolean if a field has been set.

### GetGroupsClaim

`func (o *VcenterIdentityProvidersCreateSpec) GetGroupsClaim() string`

GetGroupsClaim returns the GroupsClaim field if non-nil, zero value otherwise.

### GetGroupsClaimOk

`func (o *VcenterIdentityProvidersCreateSpec) GetGroupsClaimOk() (*string, bool)`

GetGroupsClaimOk returns a tuple with the GroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsClaim

`func (o *VcenterIdentityProvidersCreateSpec) SetGroupsClaim(v string)`

SetGroupsClaim sets GroupsClaim field to given value.

### HasGroupsClaim

`func (o *VcenterIdentityProvidersCreateSpec) HasGroupsClaim() bool`

HasGroupsClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


