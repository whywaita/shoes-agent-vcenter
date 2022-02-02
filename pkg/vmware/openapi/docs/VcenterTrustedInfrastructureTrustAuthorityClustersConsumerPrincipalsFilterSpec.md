# VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]string** | The unqiue identifier of a connection profile. If unset, no filtration will be performed by ID. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: esx.authentication.clientprofile. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: esx.authentication.clientprofile. | [optional] 
**Principals** | Pointer to [**[]VcenterTrustedInfrastructureStsPrincipal**](VcenterTrustedInfrastructureStsPrincipal.md) | The principal used by the vCenter to retrieve tokens. If unset, no filtration will be performed by principals. | [optional] 
**Issuer** | Pointer to **[]string** | The service which created and signed the security token. If unset, no filtration will be performed by issuer. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: esx.authentication.trust.security-token-issuer. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: esx.authentication.trust.security-token-issuer. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec() *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) GetId() []string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) GetIdOk() (*[]string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) SetId(v []string)`

SetId sets Id field to given value.

### HasId

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrincipals

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) GetPrincipals() []VcenterTrustedInfrastructureStsPrincipal`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) GetPrincipalsOk() (*[]VcenterTrustedInfrastructureStsPrincipal, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) SetPrincipals(v []VcenterTrustedInfrastructureStsPrincipal)`

SetPrincipals sets Principals field to given value.

### HasPrincipals

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) HasPrincipals() bool`

HasPrincipals returns a boolean if a field has been set.

### GetIssuer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) GetIssuer() []string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) GetIssuerOk() (*[]string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) SetIssuer(v []string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsFilterSpec) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


