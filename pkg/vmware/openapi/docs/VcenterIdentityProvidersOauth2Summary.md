# VcenterIdentityProvidersOauth2Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthEndpoint** | **string** | Authentication/authorization endpoint of the provider | 
**TokenEndpoint** | **string** | Token endpoint of the provider | 
**ClientId** | **string** | Client identifier to connect to the provider | 
**AuthenticationHeader** | **string** | The authentication data used as part of request header to acquire or refresh an OAuth2 token. The data format depends on the authentication method used. Example of basic authentication format: Authorization: Basic [base64Encode(clientId + \&quot;:\&quot; + secret)] | 
**AuthQueryParams** | [**[]VcenterIdentityProvidersCreateSpecAuthQueryParams**](VcenterIdentityProvidersCreateSpecAuthQueryParams.md) |  key/value pairs that are to be appended to the authEndpoint request.   How to append to authEndpoint request:  If the map is not empty, a \&quot;?\&quot; is added to the endpoint URL, and combination of each k and each string in the v is added with an \&quot;&amp;\&quot; delimiter. Details:    - If the value contains only one string, then the key is added with \&quot;k&#x3D;v\&quot;.    - If the value is an empty list, then the key is added without a \&quot;&#x3D;v\&quot;.    - If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.  | 

## Methods

### NewVcenterIdentityProvidersOauth2Summary

`func NewVcenterIdentityProvidersOauth2Summary(authEndpoint string, tokenEndpoint string, clientId string, authenticationHeader string, authQueryParams []VcenterIdentityProvidersCreateSpecAuthQueryParams, ) *VcenterIdentityProvidersOauth2Summary`

NewVcenterIdentityProvidersOauth2Summary instantiates a new VcenterIdentityProvidersOauth2Summary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterIdentityProvidersOauth2SummaryWithDefaults

`func NewVcenterIdentityProvidersOauth2SummaryWithDefaults() *VcenterIdentityProvidersOauth2Summary`

NewVcenterIdentityProvidersOauth2SummaryWithDefaults instantiates a new VcenterIdentityProvidersOauth2Summary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthEndpoint

`func (o *VcenterIdentityProvidersOauth2Summary) GetAuthEndpoint() string`

GetAuthEndpoint returns the AuthEndpoint field if non-nil, zero value otherwise.

### GetAuthEndpointOk

`func (o *VcenterIdentityProvidersOauth2Summary) GetAuthEndpointOk() (*string, bool)`

GetAuthEndpointOk returns a tuple with the AuthEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthEndpoint

`func (o *VcenterIdentityProvidersOauth2Summary) SetAuthEndpoint(v string)`

SetAuthEndpoint sets AuthEndpoint field to given value.


### GetTokenEndpoint

`func (o *VcenterIdentityProvidersOauth2Summary) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *VcenterIdentityProvidersOauth2Summary) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *VcenterIdentityProvidersOauth2Summary) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.


### GetClientId

`func (o *VcenterIdentityProvidersOauth2Summary) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *VcenterIdentityProvidersOauth2Summary) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *VcenterIdentityProvidersOauth2Summary) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetAuthenticationHeader

`func (o *VcenterIdentityProvidersOauth2Summary) GetAuthenticationHeader() string`

GetAuthenticationHeader returns the AuthenticationHeader field if non-nil, zero value otherwise.

### GetAuthenticationHeaderOk

`func (o *VcenterIdentityProvidersOauth2Summary) GetAuthenticationHeaderOk() (*string, bool)`

GetAuthenticationHeaderOk returns a tuple with the AuthenticationHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationHeader

`func (o *VcenterIdentityProvidersOauth2Summary) SetAuthenticationHeader(v string)`

SetAuthenticationHeader sets AuthenticationHeader field to given value.


### GetAuthQueryParams

`func (o *VcenterIdentityProvidersOauth2Summary) GetAuthQueryParams() []VcenterIdentityProvidersCreateSpecAuthQueryParams`

GetAuthQueryParams returns the AuthQueryParams field if non-nil, zero value otherwise.

### GetAuthQueryParamsOk

`func (o *VcenterIdentityProvidersOauth2Summary) GetAuthQueryParamsOk() (*[]VcenterIdentityProvidersCreateSpecAuthQueryParams, bool)`

GetAuthQueryParamsOk returns a tuple with the AuthQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthQueryParams

`func (o *VcenterIdentityProvidersOauth2Summary) SetAuthQueryParams(v []VcenterIdentityProvidersCreateSpecAuthQueryParams)`

SetAuthQueryParams sets AuthQueryParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


