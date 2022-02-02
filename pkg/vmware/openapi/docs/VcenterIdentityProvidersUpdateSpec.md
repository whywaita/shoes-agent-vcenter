# VcenterIdentityProvidersUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigTag** | [**VcenterIdentityProvidersConfigType**](VcenterIdentityProvidersConfigType.md) |  | 
**Oauth2** | Pointer to [**VcenterIdentityProvidersOauth2UpdateSpec**](VcenterIdentityProvidersOauth2UpdateSpec.md) |  | [optional] 
**Oidc** | Pointer to [**VcenterIdentityProvidersOidcUpdateSpec**](VcenterIdentityProvidersOidcUpdateSpec.md) |  | [optional] 
**OrgIds** | Pointer to **[]string** | The set orgIds as part of SDDC creation which provides the basis for tenancy If unset, leaves value unchanged. | [optional] 
**MakeDefault** | Pointer to **bool** | Specifies whether to make this the default provider. If Providers.UpdateSpec.make-default is set to true, this provider will be flagged as the default provider and any other providers that had previously been flagged as the default will be made non-default. If Providers.UpdateSpec.make-default is set to false, this provider&#39;s default flag will not be modified. If unset, leaves value unchanged. | [optional] 
**Name** | Pointer to **string** | The user friendly name for the provider. This name can be used for human-readable identification purposes, but it does not have to be unique, as the system will use internal UUIDs to differentiate providers. If unset, leaves value unchanged. | [optional] 
**DomainNames** | Pointer to **[]string** | Set of fully qualified domain names to trust when federating with this identity provider. Tokens from this identity provider will only be validated if the user belongs to one of these domains, and any domain-qualified groups in the tokens will be filtered to include only those groups that belong to one of these domains. If unset, leaves value unchanged. If domainNames is an empty set, domain validation behavior at login with this identity provider will be as follows: the user&#39;s domain will be parsed from the User Principal Name (UPN) value that is found in the tokens returned by the identity provider. This domain will then be implicitly trusted and used to filter any groups that are also provided in the tokens. | [optional] 
**AuthQueryParams** | Pointer to [**[]VcenterIdentityProvidersCreateSpecAuthQueryParams**](VcenterIdentityProvidersCreateSpecAuthQueryParams.md) | key/value pairs that are to be appended to the authEndpoint request. How to append to authEndpoint request: If the map is not empty, a \&quot;?\&quot; is added to the endpoint URL, and combination of each k and each string in the v is added with an \&quot;&amp;\&quot; delimiter. Details: If the value contains only one string, then the key is added with \&quot;k&#x3D;v\&quot;. If the value is an empty list, then the key is added without a \&quot;&#x3D;v\&quot;. If the value contains multiple strings, then the key is repeated in the query-string for each string in the value. If the map is empty, deletes all params. If unset, leaves value unchanged. | [optional] 
**IdmProtocol** | Pointer to [**VcenterIdentityProvidersIdmProtocol**](VcenterIdentityProvidersIdmProtocol.md) |  | [optional] 
**IdmEndpoints** | Pointer to **[]string** | Identity management endpoints. When specified, at least one endpoint must be provided. This field is optional and it is only relevant when the value of Providers.UpdateSpec.idm-protocol is one of REST, SCIM, or SCIM2_0. | [optional] 
**ActiveDirectoryOverLdap** | Pointer to [**VcenterIdentityProvidersActiveDirectoryOverLdap**](VcenterIdentityProvidersActiveDirectoryOverLdap.md) |  | [optional] 
**UpnClaim** | Pointer to **string** | Specifies which claim provides the user principal name (UPN) for the subject of the token. If unset, leaves value unchanged. | [optional] 
**ResetUpnClaim** | Pointer to **bool** | Flag indicating whether the user principal name (UPN) claim should be set back to its default value. If this field is set to true, the user principal name (UPN) claim will be set to &#39;acct&#39;, which is used for backwards compatibility with CSP. If this field is set to false, the existing user principal name (UPN) claim will be changed to the value specified in Providers.UpdateSpec.upn-claim, if any. If unset, the existing user principal name (UPN) claim will be changed to the value specified in Providers.UpdateSpec.upn-claim, if any. | [optional] 
**GroupsClaim** | Pointer to **string** | Specifies which claim provides the group membership for the token subject. If unset, leaves value unchanged. | [optional] 
**ResetGroupsClaim** | Pointer to **bool** | Flag indicating whether any existing groups claim value should be removed. If this field is set to true, the existing groups claim value is removed which defaults to backwards compatibility with CSP. In this case, the groups for the subject will be comprised of the groups in &#39;group_names&#39; and &#39;group_ids&#39; claims. If this field is set to false, the existing groups claim will be changed to the value specified in Providers.UpdateSpec.groups-claim, if any. If unset, the existing groups claim will be changed to the value specified in Providers.UpdateSpec.groups-claim, if any. | [optional] 

## Methods

### NewVcenterIdentityProvidersUpdateSpec

`func NewVcenterIdentityProvidersUpdateSpec(configTag VcenterIdentityProvidersConfigType, ) *VcenterIdentityProvidersUpdateSpec`

NewVcenterIdentityProvidersUpdateSpec instantiates a new VcenterIdentityProvidersUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterIdentityProvidersUpdateSpecWithDefaults

`func NewVcenterIdentityProvidersUpdateSpecWithDefaults() *VcenterIdentityProvidersUpdateSpec`

NewVcenterIdentityProvidersUpdateSpecWithDefaults instantiates a new VcenterIdentityProvidersUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigTag

`func (o *VcenterIdentityProvidersUpdateSpec) GetConfigTag() VcenterIdentityProvidersConfigType`

GetConfigTag returns the ConfigTag field if non-nil, zero value otherwise.

### GetConfigTagOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetConfigTagOk() (*VcenterIdentityProvidersConfigType, bool)`

GetConfigTagOk returns a tuple with the ConfigTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTag

`func (o *VcenterIdentityProvidersUpdateSpec) SetConfigTag(v VcenterIdentityProvidersConfigType)`

SetConfigTag sets ConfigTag field to given value.


### GetOauth2

`func (o *VcenterIdentityProvidersUpdateSpec) GetOauth2() VcenterIdentityProvidersOauth2UpdateSpec`

GetOauth2 returns the Oauth2 field if non-nil, zero value otherwise.

### GetOauth2Ok

`func (o *VcenterIdentityProvidersUpdateSpec) GetOauth2Ok() (*VcenterIdentityProvidersOauth2UpdateSpec, bool)`

GetOauth2Ok returns a tuple with the Oauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2

`func (o *VcenterIdentityProvidersUpdateSpec) SetOauth2(v VcenterIdentityProvidersOauth2UpdateSpec)`

SetOauth2 sets Oauth2 field to given value.

### HasOauth2

`func (o *VcenterIdentityProvidersUpdateSpec) HasOauth2() bool`

HasOauth2 returns a boolean if a field has been set.

### GetOidc

`func (o *VcenterIdentityProvidersUpdateSpec) GetOidc() VcenterIdentityProvidersOidcUpdateSpec`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetOidcOk() (*VcenterIdentityProvidersOidcUpdateSpec, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *VcenterIdentityProvidersUpdateSpec) SetOidc(v VcenterIdentityProvidersOidcUpdateSpec)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *VcenterIdentityProvidersUpdateSpec) HasOidc() bool`

HasOidc returns a boolean if a field has been set.

### GetOrgIds

`func (o *VcenterIdentityProvidersUpdateSpec) GetOrgIds() []string`

GetOrgIds returns the OrgIds field if non-nil, zero value otherwise.

### GetOrgIdsOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetOrgIdsOk() (*[]string, bool)`

GetOrgIdsOk returns a tuple with the OrgIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgIds

`func (o *VcenterIdentityProvidersUpdateSpec) SetOrgIds(v []string)`

SetOrgIds sets OrgIds field to given value.

### HasOrgIds

`func (o *VcenterIdentityProvidersUpdateSpec) HasOrgIds() bool`

HasOrgIds returns a boolean if a field has been set.

### GetMakeDefault

`func (o *VcenterIdentityProvidersUpdateSpec) GetMakeDefault() bool`

GetMakeDefault returns the MakeDefault field if non-nil, zero value otherwise.

### GetMakeDefaultOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetMakeDefaultOk() (*bool, bool)`

GetMakeDefaultOk returns a tuple with the MakeDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeDefault

`func (o *VcenterIdentityProvidersUpdateSpec) SetMakeDefault(v bool)`

SetMakeDefault sets MakeDefault field to given value.

### HasMakeDefault

`func (o *VcenterIdentityProvidersUpdateSpec) HasMakeDefault() bool`

HasMakeDefault returns a boolean if a field has been set.

### GetName

`func (o *VcenterIdentityProvidersUpdateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterIdentityProvidersUpdateSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterIdentityProvidersUpdateSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomainNames

`func (o *VcenterIdentityProvidersUpdateSpec) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *VcenterIdentityProvidersUpdateSpec) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *VcenterIdentityProvidersUpdateSpec) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetAuthQueryParams

`func (o *VcenterIdentityProvidersUpdateSpec) GetAuthQueryParams() []VcenterIdentityProvidersCreateSpecAuthQueryParams`

GetAuthQueryParams returns the AuthQueryParams field if non-nil, zero value otherwise.

### GetAuthQueryParamsOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetAuthQueryParamsOk() (*[]VcenterIdentityProvidersCreateSpecAuthQueryParams, bool)`

GetAuthQueryParamsOk returns a tuple with the AuthQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthQueryParams

`func (o *VcenterIdentityProvidersUpdateSpec) SetAuthQueryParams(v []VcenterIdentityProvidersCreateSpecAuthQueryParams)`

SetAuthQueryParams sets AuthQueryParams field to given value.

### HasAuthQueryParams

`func (o *VcenterIdentityProvidersUpdateSpec) HasAuthQueryParams() bool`

HasAuthQueryParams returns a boolean if a field has been set.

### GetIdmProtocol

`func (o *VcenterIdentityProvidersUpdateSpec) GetIdmProtocol() VcenterIdentityProvidersIdmProtocol`

GetIdmProtocol returns the IdmProtocol field if non-nil, zero value otherwise.

### GetIdmProtocolOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetIdmProtocolOk() (*VcenterIdentityProvidersIdmProtocol, bool)`

GetIdmProtocolOk returns a tuple with the IdmProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdmProtocol

`func (o *VcenterIdentityProvidersUpdateSpec) SetIdmProtocol(v VcenterIdentityProvidersIdmProtocol)`

SetIdmProtocol sets IdmProtocol field to given value.

### HasIdmProtocol

`func (o *VcenterIdentityProvidersUpdateSpec) HasIdmProtocol() bool`

HasIdmProtocol returns a boolean if a field has been set.

### GetIdmEndpoints

`func (o *VcenterIdentityProvidersUpdateSpec) GetIdmEndpoints() []string`

GetIdmEndpoints returns the IdmEndpoints field if non-nil, zero value otherwise.

### GetIdmEndpointsOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetIdmEndpointsOk() (*[]string, bool)`

GetIdmEndpointsOk returns a tuple with the IdmEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdmEndpoints

`func (o *VcenterIdentityProvidersUpdateSpec) SetIdmEndpoints(v []string)`

SetIdmEndpoints sets IdmEndpoints field to given value.

### HasIdmEndpoints

`func (o *VcenterIdentityProvidersUpdateSpec) HasIdmEndpoints() bool`

HasIdmEndpoints returns a boolean if a field has been set.

### GetActiveDirectoryOverLdap

`func (o *VcenterIdentityProvidersUpdateSpec) GetActiveDirectoryOverLdap() VcenterIdentityProvidersActiveDirectoryOverLdap`

GetActiveDirectoryOverLdap returns the ActiveDirectoryOverLdap field if non-nil, zero value otherwise.

### GetActiveDirectoryOverLdapOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetActiveDirectoryOverLdapOk() (*VcenterIdentityProvidersActiveDirectoryOverLdap, bool)`

GetActiveDirectoryOverLdapOk returns a tuple with the ActiveDirectoryOverLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryOverLdap

`func (o *VcenterIdentityProvidersUpdateSpec) SetActiveDirectoryOverLdap(v VcenterIdentityProvidersActiveDirectoryOverLdap)`

SetActiveDirectoryOverLdap sets ActiveDirectoryOverLdap field to given value.

### HasActiveDirectoryOverLdap

`func (o *VcenterIdentityProvidersUpdateSpec) HasActiveDirectoryOverLdap() bool`

HasActiveDirectoryOverLdap returns a boolean if a field has been set.

### GetUpnClaim

`func (o *VcenterIdentityProvidersUpdateSpec) GetUpnClaim() string`

GetUpnClaim returns the UpnClaim field if non-nil, zero value otherwise.

### GetUpnClaimOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetUpnClaimOk() (*string, bool)`

GetUpnClaimOk returns a tuple with the UpnClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpnClaim

`func (o *VcenterIdentityProvidersUpdateSpec) SetUpnClaim(v string)`

SetUpnClaim sets UpnClaim field to given value.

### HasUpnClaim

`func (o *VcenterIdentityProvidersUpdateSpec) HasUpnClaim() bool`

HasUpnClaim returns a boolean if a field has been set.

### GetResetUpnClaim

`func (o *VcenterIdentityProvidersUpdateSpec) GetResetUpnClaim() bool`

GetResetUpnClaim returns the ResetUpnClaim field if non-nil, zero value otherwise.

### GetResetUpnClaimOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetResetUpnClaimOk() (*bool, bool)`

GetResetUpnClaimOk returns a tuple with the ResetUpnClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetUpnClaim

`func (o *VcenterIdentityProvidersUpdateSpec) SetResetUpnClaim(v bool)`

SetResetUpnClaim sets ResetUpnClaim field to given value.

### HasResetUpnClaim

`func (o *VcenterIdentityProvidersUpdateSpec) HasResetUpnClaim() bool`

HasResetUpnClaim returns a boolean if a field has been set.

### GetGroupsClaim

`func (o *VcenterIdentityProvidersUpdateSpec) GetGroupsClaim() string`

GetGroupsClaim returns the GroupsClaim field if non-nil, zero value otherwise.

### GetGroupsClaimOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetGroupsClaimOk() (*string, bool)`

GetGroupsClaimOk returns a tuple with the GroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsClaim

`func (o *VcenterIdentityProvidersUpdateSpec) SetGroupsClaim(v string)`

SetGroupsClaim sets GroupsClaim field to given value.

### HasGroupsClaim

`func (o *VcenterIdentityProvidersUpdateSpec) HasGroupsClaim() bool`

HasGroupsClaim returns a boolean if a field has been set.

### GetResetGroupsClaim

`func (o *VcenterIdentityProvidersUpdateSpec) GetResetGroupsClaim() bool`

GetResetGroupsClaim returns the ResetGroupsClaim field if non-nil, zero value otherwise.

### GetResetGroupsClaimOk

`func (o *VcenterIdentityProvidersUpdateSpec) GetResetGroupsClaimOk() (*bool, bool)`

GetResetGroupsClaimOk returns a tuple with the ResetGroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetGroupsClaim

`func (o *VcenterIdentityProvidersUpdateSpec) SetResetGroupsClaim(v bool)`

SetResetGroupsClaim sets ResetGroupsClaim field to given value.

### HasResetGroupsClaim

`func (o *VcenterIdentityProvidersUpdateSpec) HasResetGroupsClaim() bool`

HasResetGroupsClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


