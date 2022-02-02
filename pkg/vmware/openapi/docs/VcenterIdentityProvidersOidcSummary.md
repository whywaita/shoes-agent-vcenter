# VcenterIdentityProvidersOidcSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveryEndpoint** | Pointer to **string** | Endpoint to retrieve the provider metadata This field is optional because it was added in a newer version than its parent node. | [optional] 
**LogoutEndpoint** | Pointer to **string** | The endpoint to use for terminating the user&#39;s session at the identity provider. This value is automatically derived from the metadata information provided by the OIDC discovery endpoint. This field is optional because it was added in a newer version than its parent node. | [optional] 
**AuthEndpoint** | **string** | Authentication/authorization endpoint of the provider | 
**TokenEndpoint** | **string** | Token endpoint of the provider | 
**ClientId** | **string** | Client identifier to connect to the provider | 
**AuthenticationHeader** | **string** | The authentication data used as part of request header to acquire or refresh an OAuth2 token. The data format depends on the authentication method used. Example of basic authentication format: Authorization: Basic [base64Encode(clientId + \&quot;:\&quot; + secret)] | 
**AuthQueryParams** | [**[]VcenterIdentityProvidersCreateSpecAuthQueryParams**](VcenterIdentityProvidersCreateSpecAuthQueryParams.md) |  key/value pairs that are to be appended to the authEndpoint request.   How to append to authEndpoint request:  If the map is not empty, a \&quot;?\&quot; is added to the endpoint URL, and combination of each k and each string in the v is added with an \&quot;&amp;\&quot; delimiter. Details:    - If the value contains only one string, then the key is added with \&quot;k&#x3D;v\&quot;.    - If the value is an empty list, then the key is added without a \&quot;&#x3D;v\&quot;.    - If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.  | 

## Methods

### NewVcenterIdentityProvidersOidcSummary

`func NewVcenterIdentityProvidersOidcSummary(authEndpoint string, tokenEndpoint string, clientId string, authenticationHeader string, authQueryParams []VcenterIdentityProvidersCreateSpecAuthQueryParams, ) *VcenterIdentityProvidersOidcSummary`

NewVcenterIdentityProvidersOidcSummary instantiates a new VcenterIdentityProvidersOidcSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterIdentityProvidersOidcSummaryWithDefaults

`func NewVcenterIdentityProvidersOidcSummaryWithDefaults() *VcenterIdentityProvidersOidcSummary`

NewVcenterIdentityProvidersOidcSummaryWithDefaults instantiates a new VcenterIdentityProvidersOidcSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveryEndpoint

`func (o *VcenterIdentityProvidersOidcSummary) GetDiscoveryEndpoint() string`

GetDiscoveryEndpoint returns the DiscoveryEndpoint field if non-nil, zero value otherwise.

### GetDiscoveryEndpointOk

`func (o *VcenterIdentityProvidersOidcSummary) GetDiscoveryEndpointOk() (*string, bool)`

GetDiscoveryEndpointOk returns a tuple with the DiscoveryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEndpoint

`func (o *VcenterIdentityProvidersOidcSummary) SetDiscoveryEndpoint(v string)`

SetDiscoveryEndpoint sets DiscoveryEndpoint field to given value.

### HasDiscoveryEndpoint

`func (o *VcenterIdentityProvidersOidcSummary) HasDiscoveryEndpoint() bool`

HasDiscoveryEndpoint returns a boolean if a field has been set.

### GetLogoutEndpoint

`func (o *VcenterIdentityProvidersOidcSummary) GetLogoutEndpoint() string`

GetLogoutEndpoint returns the LogoutEndpoint field if non-nil, zero value otherwise.

### GetLogoutEndpointOk

`func (o *VcenterIdentityProvidersOidcSummary) GetLogoutEndpointOk() (*string, bool)`

GetLogoutEndpointOk returns a tuple with the LogoutEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutEndpoint

`func (o *VcenterIdentityProvidersOidcSummary) SetLogoutEndpoint(v string)`

SetLogoutEndpoint sets LogoutEndpoint field to given value.

### HasLogoutEndpoint

`func (o *VcenterIdentityProvidersOidcSummary) HasLogoutEndpoint() bool`

HasLogoutEndpoint returns a boolean if a field has been set.

### GetAuthEndpoint

`func (o *VcenterIdentityProvidersOidcSummary) GetAuthEndpoint() string`

GetAuthEndpoint returns the AuthEndpoint field if non-nil, zero value otherwise.

### GetAuthEndpointOk

`func (o *VcenterIdentityProvidersOidcSummary) GetAuthEndpointOk() (*string, bool)`

GetAuthEndpointOk returns a tuple with the AuthEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthEndpoint

`func (o *VcenterIdentityProvidersOidcSummary) SetAuthEndpoint(v string)`

SetAuthEndpoint sets AuthEndpoint field to given value.


### GetTokenEndpoint

`func (o *VcenterIdentityProvidersOidcSummary) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *VcenterIdentityProvidersOidcSummary) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *VcenterIdentityProvidersOidcSummary) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.


### GetClientId

`func (o *VcenterIdentityProvidersOidcSummary) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *VcenterIdentityProvidersOidcSummary) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *VcenterIdentityProvidersOidcSummary) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetAuthenticationHeader

`func (o *VcenterIdentityProvidersOidcSummary) GetAuthenticationHeader() string`

GetAuthenticationHeader returns the AuthenticationHeader field if non-nil, zero value otherwise.

### GetAuthenticationHeaderOk

`func (o *VcenterIdentityProvidersOidcSummary) GetAuthenticationHeaderOk() (*string, bool)`

GetAuthenticationHeaderOk returns a tuple with the AuthenticationHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationHeader

`func (o *VcenterIdentityProvidersOidcSummary) SetAuthenticationHeader(v string)`

SetAuthenticationHeader sets AuthenticationHeader field to given value.


### GetAuthQueryParams

`func (o *VcenterIdentityProvidersOidcSummary) GetAuthQueryParams() []VcenterIdentityProvidersCreateSpecAuthQueryParams`

GetAuthQueryParams returns the AuthQueryParams field if non-nil, zero value otherwise.

### GetAuthQueryParamsOk

`func (o *VcenterIdentityProvidersOidcSummary) GetAuthQueryParamsOk() (*[]VcenterIdentityProvidersCreateSpecAuthQueryParams, bool)`

GetAuthQueryParamsOk returns a tuple with the AuthQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthQueryParams

`func (o *VcenterIdentityProvidersOidcSummary) SetAuthQueryParams(v []VcenterIdentityProvidersCreateSpecAuthQueryParams)`

SetAuthQueryParams sets AuthQueryParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


