# VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to [**[]VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigHealth**](VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigHealth.md) | The health of the applied Attestation Service configuration. If unset, no filtration will be performed by health. | [optional] 
**Address** | Pointer to [**[]VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) | The network address of the Attestation Service configured for use in the Trusted Cluster. If unset, no filtration will be performed by network address. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec

`func NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec() *VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec`

NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec instantiates a new VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpecWithDefaults() *VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec`

NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec) GetHealth() []VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec) GetHealthOk() (*[]VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec) SetHealth(v []VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec) GetAddress() []VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec) GetAddressOk() (*[]VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec) SetAddress(v []VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesAppliedConfigFilterSpec) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


