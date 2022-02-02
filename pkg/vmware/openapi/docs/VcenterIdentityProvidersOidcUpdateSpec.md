# VcenterIdentityProvidersOidcUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveryEndpoint** | Pointer to **string** | Endpoint to retrieve the provider metadata If unset, leaves value unchanged. | [optional] 
**ClientId** | Pointer to **string** | Client identifier to connect to the provider If unset, leaves value unchanged. | [optional] 
**ClientSecret** | Pointer to **string** | The secret shared between the client and the provider If unset, leaves value unchanged. | [optional] 
**ClaimMap** | Pointer to [**[]VcenterIdentityProvidersOauth2CreateSpecClaimMap**](VcenterIdentityProvidersOauth2CreateSpecClaimMap.md) | The map used to transform an OAuth2 claim to a corresponding claim that vCenter Server understands. Currently only the key \&quot;perms\&quot; is supported. The key \&quot;perms\&quot; is used for mapping the \&quot;perms\&quot; claim of incoming JWT. The value is another map with an external group as the key and a vCenter Server group as value. If unset, leaves value unchanged. | [optional] 

## Methods

### NewVcenterIdentityProvidersOidcUpdateSpec

`func NewVcenterIdentityProvidersOidcUpdateSpec() *VcenterIdentityProvidersOidcUpdateSpec`

NewVcenterIdentityProvidersOidcUpdateSpec instantiates a new VcenterIdentityProvidersOidcUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterIdentityProvidersOidcUpdateSpecWithDefaults

`func NewVcenterIdentityProvidersOidcUpdateSpecWithDefaults() *VcenterIdentityProvidersOidcUpdateSpec`

NewVcenterIdentityProvidersOidcUpdateSpecWithDefaults instantiates a new VcenterIdentityProvidersOidcUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveryEndpoint

`func (o *VcenterIdentityProvidersOidcUpdateSpec) GetDiscoveryEndpoint() string`

GetDiscoveryEndpoint returns the DiscoveryEndpoint field if non-nil, zero value otherwise.

### GetDiscoveryEndpointOk

`func (o *VcenterIdentityProvidersOidcUpdateSpec) GetDiscoveryEndpointOk() (*string, bool)`

GetDiscoveryEndpointOk returns a tuple with the DiscoveryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEndpoint

`func (o *VcenterIdentityProvidersOidcUpdateSpec) SetDiscoveryEndpoint(v string)`

SetDiscoveryEndpoint sets DiscoveryEndpoint field to given value.

### HasDiscoveryEndpoint

`func (o *VcenterIdentityProvidersOidcUpdateSpec) HasDiscoveryEndpoint() bool`

HasDiscoveryEndpoint returns a boolean if a field has been set.

### GetClientId

`func (o *VcenterIdentityProvidersOidcUpdateSpec) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *VcenterIdentityProvidersOidcUpdateSpec) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *VcenterIdentityProvidersOidcUpdateSpec) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *VcenterIdentityProvidersOidcUpdateSpec) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *VcenterIdentityProvidersOidcUpdateSpec) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *VcenterIdentityProvidersOidcUpdateSpec) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *VcenterIdentityProvidersOidcUpdateSpec) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *VcenterIdentityProvidersOidcUpdateSpec) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClaimMap

`func (o *VcenterIdentityProvidersOidcUpdateSpec) GetClaimMap() []VcenterIdentityProvidersOauth2CreateSpecClaimMap`

GetClaimMap returns the ClaimMap field if non-nil, zero value otherwise.

### GetClaimMapOk

`func (o *VcenterIdentityProvidersOidcUpdateSpec) GetClaimMapOk() (*[]VcenterIdentityProvidersOauth2CreateSpecClaimMap, bool)`

GetClaimMapOk returns a tuple with the ClaimMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimMap

`func (o *VcenterIdentityProvidersOidcUpdateSpec) SetClaimMap(v []VcenterIdentityProvidersOauth2CreateSpecClaimMap)`

SetClaimMap sets ClaimMap field to given value.

### HasClaimMap

`func (o *VcenterIdentityProvidersOidcUpdateSpec) HasClaimMap() bool`

HasClaimMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


