# VcenterDeploymentInstallInstallSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcsaEmbedded** | [**VcenterDeploymentInstallVcsaEmbeddedSpec**](VcenterDeploymentInstallVcsaEmbeddedSpec.md) |  | 
**AutoAnswer** | Pointer to **bool** | Use the default option for any questions that may come up during appliance configuration. If unset, will default to false. | [optional] 

## Methods

### NewVcenterDeploymentInstallInstallSpec

`func NewVcenterDeploymentInstallInstallSpec(vcsaEmbedded VcenterDeploymentInstallVcsaEmbeddedSpec, ) *VcenterDeploymentInstallInstallSpec`

NewVcenterDeploymentInstallInstallSpec instantiates a new VcenterDeploymentInstallInstallSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentInstallInstallSpecWithDefaults

`func NewVcenterDeploymentInstallInstallSpecWithDefaults() *VcenterDeploymentInstallInstallSpec`

NewVcenterDeploymentInstallInstallSpecWithDefaults instantiates a new VcenterDeploymentInstallInstallSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcsaEmbedded

`func (o *VcenterDeploymentInstallInstallSpec) GetVcsaEmbedded() VcenterDeploymentInstallVcsaEmbeddedSpec`

GetVcsaEmbedded returns the VcsaEmbedded field if non-nil, zero value otherwise.

### GetVcsaEmbeddedOk

`func (o *VcenterDeploymentInstallInstallSpec) GetVcsaEmbeddedOk() (*VcenterDeploymentInstallVcsaEmbeddedSpec, bool)`

GetVcsaEmbeddedOk returns a tuple with the VcsaEmbedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsaEmbedded

`func (o *VcenterDeploymentInstallInstallSpec) SetVcsaEmbedded(v VcenterDeploymentInstallVcsaEmbeddedSpec)`

SetVcsaEmbedded sets VcsaEmbedded field to given value.


### GetAutoAnswer

`func (o *VcenterDeploymentInstallInstallSpec) GetAutoAnswer() bool`

GetAutoAnswer returns the AutoAnswer field if non-nil, zero value otherwise.

### GetAutoAnswerOk

`func (o *VcenterDeploymentInstallInstallSpec) GetAutoAnswerOk() (*bool, bool)`

GetAutoAnswerOk returns a tuple with the AutoAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAnswer

`func (o *VcenterDeploymentInstallInstallSpec) SetAutoAnswer(v bool)`

SetAutoAnswer sets AutoAnswer field to given value.

### HasAutoAnswer

`func (o *VcenterDeploymentInstallInstallSpec) HasAutoAnswer() bool`

HasAutoAnswer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


