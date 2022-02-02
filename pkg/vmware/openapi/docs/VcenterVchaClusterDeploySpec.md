# VcenterVchaClusterDeploySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcSpec** | Pointer to [**VcenterVchaCredentialsSpec**](VcenterVchaCredentialsSpec.md) |  | [optional] 
**Deployment** | [**VcenterVchaClusterType**](VcenterVchaClusterType.md) |  | 
**Active** | [**VcenterVchaClusterActiveSpec**](VcenterVchaClusterActiveSpec.md) |  | 
**Passive** | [**VcenterVchaClusterPassiveSpec**](VcenterVchaClusterPassiveSpec.md) |  | 
**Witness** | [**VcenterVchaClusterWitnessSpec**](VcenterVchaClusterWitnessSpec.md) |  | 

## Methods

### NewVcenterVchaClusterDeploySpec

`func NewVcenterVchaClusterDeploySpec(deployment VcenterVchaClusterType, active VcenterVchaClusterActiveSpec, passive VcenterVchaClusterPassiveSpec, witness VcenterVchaClusterWitnessSpec, ) *VcenterVchaClusterDeploySpec`

NewVcenterVchaClusterDeploySpec instantiates a new VcenterVchaClusterDeploySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterDeploySpecWithDefaults

`func NewVcenterVchaClusterDeploySpecWithDefaults() *VcenterVchaClusterDeploySpec`

NewVcenterVchaClusterDeploySpecWithDefaults instantiates a new VcenterVchaClusterDeploySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcSpec

`func (o *VcenterVchaClusterDeploySpec) GetVcSpec() VcenterVchaCredentialsSpec`

GetVcSpec returns the VcSpec field if non-nil, zero value otherwise.

### GetVcSpecOk

`func (o *VcenterVchaClusterDeploySpec) GetVcSpecOk() (*VcenterVchaCredentialsSpec, bool)`

GetVcSpecOk returns a tuple with the VcSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcSpec

`func (o *VcenterVchaClusterDeploySpec) SetVcSpec(v VcenterVchaCredentialsSpec)`

SetVcSpec sets VcSpec field to given value.

### HasVcSpec

`func (o *VcenterVchaClusterDeploySpec) HasVcSpec() bool`

HasVcSpec returns a boolean if a field has been set.

### GetDeployment

`func (o *VcenterVchaClusterDeploySpec) GetDeployment() VcenterVchaClusterType`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *VcenterVchaClusterDeploySpec) GetDeploymentOk() (*VcenterVchaClusterType, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *VcenterVchaClusterDeploySpec) SetDeployment(v VcenterVchaClusterType)`

SetDeployment sets Deployment field to given value.


### GetActive

`func (o *VcenterVchaClusterDeploySpec) GetActive() VcenterVchaClusterActiveSpec`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *VcenterVchaClusterDeploySpec) GetActiveOk() (*VcenterVchaClusterActiveSpec, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *VcenterVchaClusterDeploySpec) SetActive(v VcenterVchaClusterActiveSpec)`

SetActive sets Active field to given value.


### GetPassive

`func (o *VcenterVchaClusterDeploySpec) GetPassive() VcenterVchaClusterPassiveSpec`

GetPassive returns the Passive field if non-nil, zero value otherwise.

### GetPassiveOk

`func (o *VcenterVchaClusterDeploySpec) GetPassiveOk() (*VcenterVchaClusterPassiveSpec, bool)`

GetPassiveOk returns a tuple with the Passive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassive

`func (o *VcenterVchaClusterDeploySpec) SetPassive(v VcenterVchaClusterPassiveSpec)`

SetPassive sets Passive field to given value.


### GetWitness

`func (o *VcenterVchaClusterDeploySpec) GetWitness() VcenterVchaClusterWitnessSpec`

GetWitness returns the Witness field if non-nil, zero value otherwise.

### GetWitnessOk

`func (o *VcenterVchaClusterDeploySpec) GetWitnessOk() (*VcenterVchaClusterWitnessSpec, bool)`

GetWitnessOk returns a tuple with the Witness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWitness

`func (o *VcenterVchaClusterDeploySpec) SetWitness(v VcenterVchaClusterWitnessSpec)`

SetWitness sets Witness field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


