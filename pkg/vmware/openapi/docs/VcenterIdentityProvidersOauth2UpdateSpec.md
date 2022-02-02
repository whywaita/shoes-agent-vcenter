# VcenterIdentityProvidersOauth2UpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthEndpoint** | Pointer to **string** | Authentication/authorization endpoint of the provider If unset, leaves value unchanged. | [optional] 
**TokenEndpoint** | Pointer to **string** | Token endpoint of the provider. If unset, leaves value unchanged. | [optional] 
**PublicKeyUri** | Pointer to **string** | Endpoint to retrieve the provider public key for validation If unset, leaves value unchanged. | [optional] 
**ClientId** | Pointer to **string** | Client identifier to connect to the provider If unset, leaves value unchanged. | [optional] 
**ClientSecret** | Pointer to **string** | Shared secret between identity provider and client If unset, leaves value unchanged. | [optional] 
**ClaimMap** | Pointer to [**[]VcenterIdentityProvidersOauth2CreateSpecClaimMap**](VcenterIdentityProvidersOauth2CreateSpecClaimMap.md) | The map used to transform an OAuth2 claim to a corresponding claim that vCenter Server understands. Currently only the key \&quot;perms\&quot; is supported. The key \&quot;perms\&quot; is used for mapping the \&quot;perms\&quot; claim of incoming JWT. The value is another map with an external group as the key and a vCenter Server group as value. If unset, leaves value unchanged. | [optional] 
**Issuer** | Pointer to **string** | The identity provider namespace. It is used to validate the issuer in the acquired OAuth2 token If unset, leaves value unchanged. | [optional] 
**AuthenticationMethod** | Pointer to [**VcenterIdentityProvidersOauth2AuthenticationMethod**](VcenterIdentityProvidersOauth2AuthenticationMethod.md) |  | [optional] 
**AuthQueryParams** | Pointer to [**[]VcenterIdentityProvidersCreateSpecAuthQueryParams**](VcenterIdentityProvidersCreateSpecAuthQueryParams.md) | key/value pairs that are to be appended to the authEndpoint request. How to append to authEndpoint request: If the map is not empty, a \&quot;?\&quot; is added to the endpoint URL, and combination of each k and each string in the v is added with an \&quot;&amp;\&quot; delimiter. Details: If the value contains only one string, then the key is added with \&quot;k&#x3D;v\&quot;. If the value is an empty list, then the key is added without a \&quot;&#x3D;v\&quot;. If the value contains multiple strings, then the key is repeated in the query-string for each string in the value. If the map is empty, deletes all params. If unset, leaves value unchanged. | [optional] 

## Methods

### NewVcenterIdentityProvidersOauth2UpdateSpec

`func NewVcenterIdentityProvidersOauth2UpdateSpec() *VcenterIdentityProvidersOauth2UpdateSpec`

NewVcenterIdentityProvidersOauth2UpdateSpec instantiates a new VcenterIdentityProvidersOauth2UpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterIdentityProvidersOauth2UpdateSpecWithDefaults

`func NewVcenterIdentityProvidersOauth2UpdateSpecWithDefaults() *VcenterIdentityProvidersOauth2UpdateSpec`

NewVcenterIdentityProvidersOauth2UpdateSpecWithDefaults instantiates a new VcenterIdentityProvidersOauth2UpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthEndpoint

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetAuthEndpoint() string`

GetAuthEndpoint returns the AuthEndpoint field if non-nil, zero value otherwise.

### GetAuthEndpointOk

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetAuthEndpointOk() (*string, bool)`

GetAuthEndpointOk returns a tuple with the AuthEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthEndpoint

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) SetAuthEndpoint(v string)`

SetAuthEndpoint sets AuthEndpoint field to given value.

### HasAuthEndpoint

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) HasAuthEndpoint() bool`

HasAuthEndpoint returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetPublicKeyUri

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetPublicKeyUri() string`

GetPublicKeyUri returns the PublicKeyUri field if non-nil, zero value otherwise.

### GetPublicKeyUriOk

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetPublicKeyUriOk() (*string, bool)`

GetPublicKeyUriOk returns a tuple with the PublicKeyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyUri

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) SetPublicKeyUri(v string)`

SetPublicKeyUri sets PublicKeyUri field to given value.

### HasPublicKeyUri

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) HasPublicKeyUri() bool`

HasPublicKeyUri returns a boolean if a field has been set.

### GetClientId

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClaimMap

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetClaimMap() []VcenterIdentityProvidersOauth2CreateSpecClaimMap`

GetClaimMap returns the ClaimMap field if non-nil, zero value otherwise.

### GetClaimMapOk

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetClaimMapOk() (*[]VcenterIdentityProvidersOauth2CreateSpecClaimMap, bool)`

GetClaimMapOk returns a tuple with the ClaimMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimMap

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) SetClaimMap(v []VcenterIdentityProvidersOauth2CreateSpecClaimMap)`

SetClaimMap sets ClaimMap field to given value.

### HasClaimMap

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) HasClaimMap() bool`

HasClaimMap returns a boolean if a field has been set.

### GetIssuer

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetAuthenticationMethod() VcenterIdentityProvidersOauth2AuthenticationMethod`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetAuthenticationMethodOk() (*VcenterIdentityProvidersOauth2AuthenticationMethod, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) SetAuthenticationMethod(v VcenterIdentityProvidersOauth2AuthenticationMethod)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetAuthQueryParams

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetAuthQueryParams() []VcenterIdentityProvidersCreateSpecAuthQueryParams`

GetAuthQueryParams returns the AuthQueryParams field if non-nil, zero value otherwise.

### GetAuthQueryParamsOk

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) GetAuthQueryParamsOk() (*[]VcenterIdentityProvidersCreateSpecAuthQueryParams, bool)`

GetAuthQueryParamsOk returns a tuple with the AuthQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthQueryParams

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) SetAuthQueryParams(v []VcenterIdentityProvidersCreateSpecAuthQueryParams)`

SetAuthQueryParams sets AuthQueryParams field to given value.

### HasAuthQueryParams

`func (o *VcenterIdentityProvidersOauth2UpdateSpec) HasAuthQueryParams() bool`

HasAuthQueryParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


