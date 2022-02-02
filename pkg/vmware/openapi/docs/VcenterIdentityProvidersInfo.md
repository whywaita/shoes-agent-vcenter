# VcenterIdentityProvidersInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The user friendly name for the provider This field is optional because it was added in a newer version than its parent node. | [optional] 
**OrgIds** | **[]string** | The set of orgIds as part of SDDC creation which provides the basis for tenancy | 
**ConfigTag** | [**VcenterIdentityProvidersConfigType**](VcenterIdentityProvidersConfigType.md) |  | 
**Oauth2** | Pointer to [**VcenterIdentityProvidersOauth2Info**](VcenterIdentityProvidersOauth2Info.md) |  | [optional] 
**Oidc** | Pointer to [**VcenterIdentityProvidersOidcInfo**](VcenterIdentityProvidersOidcInfo.md) |  | [optional] 
**IsDefault** | **bool** | Specifies whether the provider is the default provider. | 
**DomainNames** | Pointer to **[]string** | Set of fully qualified domain names to trust when federating with this identity provider. Tokens from this identity provider will only be validated if the user belongs to one of these domains, and any domain-qualified groups in the tokens will be filtered to include only those groups that belong to one of these domains. If domainNames is an empty set, domain validation behavior at login with this identity provider will be as follows: the user&#39;s domain will be parsed from the User Principal Name (UPN) value that is found in the tokens returned by the identity provider. This domain will then be implicitly trusted and used to filter any groups that are also provided in the tokens. This field is optional because it was added in a newer version than its parent node. | [optional] 
**AuthQueryParams** | Pointer to [**[]VcenterIdentityProvidersCreateSpecAuthQueryParams**](VcenterIdentityProvidersCreateSpecAuthQueryParams.md) |  key/value pairs that are to be appended to the authEndpoint request.   How to append to authEndpoint request:  If the map is not empty, a \&quot;?\&quot; is added to the endpoint URL, and combination of each k and each string in the v is added with an \&quot;&amp;\&quot; delimiter. Details:    - If the value contains only one string, then the key is added with \&quot;k&#x3D;v\&quot;.    - If the value is an empty list, then the key is added without a \&quot;&#x3D;v\&quot;.    - If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.  This field is optional because it was added in a newer version than its parent node. | [optional] 
**IdmProtocol** | Pointer to [**VcenterIdentityProvidersIdmProtocol**](VcenterIdentityProvidersIdmProtocol.md) |  | [optional] 
**IdmEndpoints** | Pointer to **[]string** | Identity management endpoints. This field is optional and it is only relevant when the value of Providers.Info.idm-protocol is one of REST, SCIM, or SCIM2_0. | [optional] 
**ActiveDirectoryOverLdap** | Pointer to [**VcenterIdentityProvidersActiveDirectoryOverLdap**](VcenterIdentityProvidersActiveDirectoryOverLdap.md) |  | [optional] 
**UpnClaim** | Pointer to **string** | Specifies which claim provides the user principal name (UPN) for the user. This field is optional because it was added in a newer version than its parent node. | [optional] 
**GroupsClaim** | Pointer to **string** | Specifies which claim provides the group membership for the token subject. If empty, the default behavior for CSP is used. In this case, the groups for the subject will be comprised of the groups in &#39;group_names&#39; and &#39;group_ids&#39; claims. This field is optional because it was added in a newer version than its parent node. | [optional] 

## Methods

### NewVcenterIdentityProvidersInfo

`func NewVcenterIdentityProvidersInfo(orgIds []string, configTag VcenterIdentityProvidersConfigType, isDefault bool, ) *VcenterIdentityProvidersInfo`

NewVcenterIdentityProvidersInfo instantiates a new VcenterIdentityProvidersInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterIdentityProvidersInfoWithDefaults

`func NewVcenterIdentityProvidersInfoWithDefaults() *VcenterIdentityProvidersInfo`

NewVcenterIdentityProvidersInfoWithDefaults instantiates a new VcenterIdentityProvidersInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterIdentityProvidersInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterIdentityProvidersInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterIdentityProvidersInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterIdentityProvidersInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgIds

`func (o *VcenterIdentityProvidersInfo) GetOrgIds() []string`

GetOrgIds returns the OrgIds field if non-nil, zero value otherwise.

### GetOrgIdsOk

`func (o *VcenterIdentityProvidersInfo) GetOrgIdsOk() (*[]string, bool)`

GetOrgIdsOk returns a tuple with the OrgIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgIds

`func (o *VcenterIdentityProvidersInfo) SetOrgIds(v []string)`

SetOrgIds sets OrgIds field to given value.


### GetConfigTag

`func (o *VcenterIdentityProvidersInfo) GetConfigTag() VcenterIdentityProvidersConfigType`

GetConfigTag returns the ConfigTag field if non-nil, zero value otherwise.

### GetConfigTagOk

`func (o *VcenterIdentityProvidersInfo) GetConfigTagOk() (*VcenterIdentityProvidersConfigType, bool)`

GetConfigTagOk returns a tuple with the ConfigTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTag

`func (o *VcenterIdentityProvidersInfo) SetConfigTag(v VcenterIdentityProvidersConfigType)`

SetConfigTag sets ConfigTag field to given value.


### GetOauth2

`func (o *VcenterIdentityProvidersInfo) GetOauth2() VcenterIdentityProvidersOauth2Info`

GetOauth2 returns the Oauth2 field if non-nil, zero value otherwise.

### GetOauth2Ok

`func (o *VcenterIdentityProvidersInfo) GetOauth2Ok() (*VcenterIdentityProvidersOauth2Info, bool)`

GetOauth2Ok returns a tuple with the Oauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2

`func (o *VcenterIdentityProvidersInfo) SetOauth2(v VcenterIdentityProvidersOauth2Info)`

SetOauth2 sets Oauth2 field to given value.

### HasOauth2

`func (o *VcenterIdentityProvidersInfo) HasOauth2() bool`

HasOauth2 returns a boolean if a field has been set.

### GetOidc

`func (o *VcenterIdentityProvidersInfo) GetOidc() VcenterIdentityProvidersOidcInfo`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *VcenterIdentityProvidersInfo) GetOidcOk() (*VcenterIdentityProvidersOidcInfo, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *VcenterIdentityProvidersInfo) SetOidc(v VcenterIdentityProvidersOidcInfo)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *VcenterIdentityProvidersInfo) HasOidc() bool`

HasOidc returns a boolean if a field has been set.

### GetIsDefault

`func (o *VcenterIdentityProvidersInfo) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *VcenterIdentityProvidersInfo) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *VcenterIdentityProvidersInfo) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetDomainNames

`func (o *VcenterIdentityProvidersInfo) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *VcenterIdentityProvidersInfo) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *VcenterIdentityProvidersInfo) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *VcenterIdentityProvidersInfo) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetAuthQueryParams

`func (o *VcenterIdentityProvidersInfo) GetAuthQueryParams() []VcenterIdentityProvidersCreateSpecAuthQueryParams`

GetAuthQueryParams returns the AuthQueryParams field if non-nil, zero value otherwise.

### GetAuthQueryParamsOk

`func (o *VcenterIdentityProvidersInfo) GetAuthQueryParamsOk() (*[]VcenterIdentityProvidersCreateSpecAuthQueryParams, bool)`

GetAuthQueryParamsOk returns a tuple with the AuthQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthQueryParams

`func (o *VcenterIdentityProvidersInfo) SetAuthQueryParams(v []VcenterIdentityProvidersCreateSpecAuthQueryParams)`

SetAuthQueryParams sets AuthQueryParams field to given value.

### HasAuthQueryParams

`func (o *VcenterIdentityProvidersInfo) HasAuthQueryParams() bool`

HasAuthQueryParams returns a boolean if a field has been set.

### GetIdmProtocol

`func (o *VcenterIdentityProvidersInfo) GetIdmProtocol() VcenterIdentityProvidersIdmProtocol`

GetIdmProtocol returns the IdmProtocol field if non-nil, zero value otherwise.

### GetIdmProtocolOk

`func (o *VcenterIdentityProvidersInfo) GetIdmProtocolOk() (*VcenterIdentityProvidersIdmProtocol, bool)`

GetIdmProtocolOk returns a tuple with the IdmProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdmProtocol

`func (o *VcenterIdentityProvidersInfo) SetIdmProtocol(v VcenterIdentityProvidersIdmProtocol)`

SetIdmProtocol sets IdmProtocol field to given value.

### HasIdmProtocol

`func (o *VcenterIdentityProvidersInfo) HasIdmProtocol() bool`

HasIdmProtocol returns a boolean if a field has been set.

### GetIdmEndpoints

`func (o *VcenterIdentityProvidersInfo) GetIdmEndpoints() []string`

GetIdmEndpoints returns the IdmEndpoints field if non-nil, zero value otherwise.

### GetIdmEndpointsOk

`func (o *VcenterIdentityProvidersInfo) GetIdmEndpointsOk() (*[]string, bool)`

GetIdmEndpointsOk returns a tuple with the IdmEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdmEndpoints

`func (o *VcenterIdentityProvidersInfo) SetIdmEndpoints(v []string)`

SetIdmEndpoints sets IdmEndpoints field to given value.

### HasIdmEndpoints

`func (o *VcenterIdentityProvidersInfo) HasIdmEndpoints() bool`

HasIdmEndpoints returns a boolean if a field has been set.

### GetActiveDirectoryOverLdap

`func (o *VcenterIdentityProvidersInfo) GetActiveDirectoryOverLdap() VcenterIdentityProvidersActiveDirectoryOverLdap`

GetActiveDirectoryOverLdap returns the ActiveDirectoryOverLdap field if non-nil, zero value otherwise.

### GetActiveDirectoryOverLdapOk

`func (o *VcenterIdentityProvidersInfo) GetActiveDirectoryOverLdapOk() (*VcenterIdentityProvidersActiveDirectoryOverLdap, bool)`

GetActiveDirectoryOverLdapOk returns a tuple with the ActiveDirectoryOverLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryOverLdap

`func (o *VcenterIdentityProvidersInfo) SetActiveDirectoryOverLdap(v VcenterIdentityProvidersActiveDirectoryOverLdap)`

SetActiveDirectoryOverLdap sets ActiveDirectoryOverLdap field to given value.

### HasActiveDirectoryOverLdap

`func (o *VcenterIdentityProvidersInfo) HasActiveDirectoryOverLdap() bool`

HasActiveDirectoryOverLdap returns a boolean if a field has been set.

### GetUpnClaim

`func (o *VcenterIdentityProvidersInfo) GetUpnClaim() string`

GetUpnClaim returns the UpnClaim field if non-nil, zero value otherwise.

### GetUpnClaimOk

`func (o *VcenterIdentityProvidersInfo) GetUpnClaimOk() (*string, bool)`

GetUpnClaimOk returns a tuple with the UpnClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpnClaim

`func (o *VcenterIdentityProvidersInfo) SetUpnClaim(v string)`

SetUpnClaim sets UpnClaim field to given value.

### HasUpnClaim

`func (o *VcenterIdentityProvidersInfo) HasUpnClaim() bool`

HasUpnClaim returns a boolean if a field has been set.

### GetGroupsClaim

`func (o *VcenterIdentityProvidersInfo) GetGroupsClaim() string`

GetGroupsClaim returns the GroupsClaim field if non-nil, zero value otherwise.

### GetGroupsClaimOk

`func (o *VcenterIdentityProvidersInfo) GetGroupsClaimOk() (*string, bool)`

GetGroupsClaimOk returns a tuple with the GroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsClaim

`func (o *VcenterIdentityProvidersInfo) SetGroupsClaim(v string)`

SetGroupsClaim sets GroupsClaim field to given value.

### HasGroupsClaim

`func (o *VcenterIdentityProvidersInfo) HasGroupsClaim() bool`

HasGroupsClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


