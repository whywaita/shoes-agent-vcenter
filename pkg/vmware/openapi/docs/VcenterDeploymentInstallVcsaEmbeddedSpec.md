# VcenterDeploymentInstallVcsaEmbeddedSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Standalone** | Pointer to [**VcenterDeploymentStandaloneSpec**](VcenterDeploymentStandaloneSpec.md) |  | [optional] 
**Replicated** | Pointer to [**VcenterDeploymentReplicatedSpec**](VcenterDeploymentReplicatedSpec.md) |  | [optional] 
**CeipEnabled** | **bool** | Whether CEIP should be enabled or disabled. | 

## Methods

### NewVcenterDeploymentInstallVcsaEmbeddedSpec

`func NewVcenterDeploymentInstallVcsaEmbeddedSpec(ceipEnabled bool, ) *VcenterDeploymentInstallVcsaEmbeddedSpec`

NewVcenterDeploymentInstallVcsaEmbeddedSpec instantiates a new VcenterDeploymentInstallVcsaEmbeddedSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentInstallVcsaEmbeddedSpecWithDefaults

`func NewVcenterDeploymentInstallVcsaEmbeddedSpecWithDefaults() *VcenterDeploymentInstallVcsaEmbeddedSpec`

NewVcenterDeploymentInstallVcsaEmbeddedSpecWithDefaults instantiates a new VcenterDeploymentInstallVcsaEmbeddedSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandalone

`func (o *VcenterDeploymentInstallVcsaEmbeddedSpec) GetStandalone() VcenterDeploymentStandaloneSpec`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *VcenterDeploymentInstallVcsaEmbeddedSpec) GetStandaloneOk() (*VcenterDeploymentStandaloneSpec, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *VcenterDeploymentInstallVcsaEmbeddedSpec) SetStandalone(v VcenterDeploymentStandaloneSpec)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *VcenterDeploymentInstallVcsaEmbeddedSpec) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetReplicated

`func (o *VcenterDeploymentInstallVcsaEmbeddedSpec) GetReplicated() VcenterDeploymentReplicatedSpec`

GetReplicated returns the Replicated field if non-nil, zero value otherwise.

### GetReplicatedOk

`func (o *VcenterDeploymentInstallVcsaEmbeddedSpec) GetReplicatedOk() (*VcenterDeploymentReplicatedSpec, bool)`

GetReplicatedOk returns a tuple with the Replicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicated

`func (o *VcenterDeploymentInstallVcsaEmbeddedSpec) SetReplicated(v VcenterDeploymentReplicatedSpec)`

SetReplicated sets Replicated field to given value.

### HasReplicated

`func (o *VcenterDeploymentInstallVcsaEmbeddedSpec) HasReplicated() bool`

HasReplicated returns a boolean if a field has been set.

### GetCeipEnabled

`func (o *VcenterDeploymentInstallVcsaEmbeddedSpec) GetCeipEnabled() bool`

GetCeipEnabled returns the CeipEnabled field if non-nil, zero value otherwise.

### GetCeipEnabledOk

`func (o *VcenterDeploymentInstallVcsaEmbeddedSpec) GetCeipEnabledOk() (*bool, bool)`

GetCeipEnabledOk returns a tuple with the CeipEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeipEnabled

`func (o *VcenterDeploymentInstallVcsaEmbeddedSpec) SetCeipEnabled(v bool)`

SetCeipEnabled sets CeipEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


