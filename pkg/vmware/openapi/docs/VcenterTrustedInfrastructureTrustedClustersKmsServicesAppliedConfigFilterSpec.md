# VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to [**[]VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigHealth**](VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigHealth.md) | The health of the applied Key Provider Service configuration. If unset, no filtration will be performed by health. | [optional] 
**Address** | Pointer to [**[]VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) | The network address of the Key Provider Service configured for use in the Trusted Cluster. If unset, no filtration will be performed by network address. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec

`func NewVcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec() *VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec`

NewVcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec instantiates a new VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpecWithDefaults() *VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec`

NewVcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec) GetHealth() []VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec) GetHealthOk() (*[]VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec) SetHealth(v []VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec) GetAddress() []VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec) GetAddressOk() (*[]VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec) SetAddress(v []VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesAppliedConfigFilterSpec) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


