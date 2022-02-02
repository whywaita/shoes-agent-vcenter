# VcenterIdentityProvidersOidcCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveryEndpoint** | **string** | Endpoint to retrieve the provider metadata | 
**ClientId** | **string** | Client identifier to connect to the provider | 
**ClientSecret** | **string** | The secret shared between the client and the provider | 
**ClaimMap** | [**[]VcenterIdentityProvidersOauth2CreateSpecClaimMap**](VcenterIdentityProvidersOauth2CreateSpecClaimMap.md) | The map used to transform an OAuth2 claim to a corresponding claim that vCenter Server understands. Currently only the key \&quot;perms\&quot; is supported. The key \&quot;perms\&quot; is used for mapping the \&quot;perms\&quot; claim of incoming JWT. The value is another map with an external group as the key and a vCenter Server group as value. | 

## Methods

### NewVcenterIdentityProvidersOidcCreateSpec

`func NewVcenterIdentityProvidersOidcCreateSpec(discoveryEndpoint string, clientId string, clientSecret string, claimMap []VcenterIdentityProvidersOauth2CreateSpecClaimMap, ) *VcenterIdentityProvidersOidcCreateSpec`

NewVcenterIdentityProvidersOidcCreateSpec instantiates a new VcenterIdentityProvidersOidcCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterIdentityProvidersOidcCreateSpecWithDefaults

`func NewVcenterIdentityProvidersOidcCreateSpecWithDefaults() *VcenterIdentityProvidersOidcCreateSpec`

NewVcenterIdentityProvidersOidcCreateSpecWithDefaults instantiates a new VcenterIdentityProvidersOidcCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveryEndpoint

`func (o *VcenterIdentityProvidersOidcCreateSpec) GetDiscoveryEndpoint() string`

GetDiscoveryEndpoint returns the DiscoveryEndpoint field if non-nil, zero value otherwise.

### GetDiscoveryEndpointOk

`func (o *VcenterIdentityProvidersOidcCreateSpec) GetDiscoveryEndpointOk() (*string, bool)`

GetDiscoveryEndpointOk returns a tuple with the DiscoveryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEndpoint

`func (o *VcenterIdentityProvidersOidcCreateSpec) SetDiscoveryEndpoint(v string)`

SetDiscoveryEndpoint sets DiscoveryEndpoint field to given value.


### GetClientId

`func (o *VcenterIdentityProvidersOidcCreateSpec) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *VcenterIdentityProvidersOidcCreateSpec) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *VcenterIdentityProvidersOidcCreateSpec) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *VcenterIdentityProvidersOidcCreateSpec) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *VcenterIdentityProvidersOidcCreateSpec) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *VcenterIdentityProvidersOidcCreateSpec) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetClaimMap

`func (o *VcenterIdentityProvidersOidcCreateSpec) GetClaimMap() []VcenterIdentityProvidersOauth2CreateSpecClaimMap`

GetClaimMap returns the ClaimMap field if non-nil, zero value otherwise.

### GetClaimMapOk

`func (o *VcenterIdentityProvidersOidcCreateSpec) GetClaimMapOk() (*[]VcenterIdentityProvidersOauth2CreateSpecClaimMap, bool)`

GetClaimMapOk returns a tuple with the ClaimMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimMap

`func (o *VcenterIdentityProvidersOidcCreateSpec) SetClaimMap(v []VcenterIdentityProvidersOauth2CreateSpecClaimMap)`

SetClaimMap sets ClaimMap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


