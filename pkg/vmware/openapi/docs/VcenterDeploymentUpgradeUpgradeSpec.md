# VcenterDeploymentUpgradeUpgradeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAppliance** | [**VcenterDeploymentUpgradeSourceApplianceSpec**](VcenterDeploymentUpgradeSourceApplianceSpec.md) |  | 
**SourceLocation** | [**VcenterDeploymentLocationSpec**](VcenterDeploymentLocationSpec.md) |  | 
**History** | Pointer to [**VcenterDeploymentHistoryMigrationSpec**](VcenterDeploymentHistoryMigrationSpec.md) |  | [optional] 
**VcsaEmbedded** | Pointer to [**VcenterDeploymentUpgradeVcsaEmbeddedSpec**](VcenterDeploymentUpgradeVcsaEmbeddedSpec.md) |  | [optional] 
**Psc** | Pointer to [**VcenterDeploymentUpgradePscSpec**](VcenterDeploymentUpgradePscSpec.md) |  | [optional] 
**AutoAnswer** | Pointer to **bool** | Use the default option for any questions that may come up during appliance configuration. If unset, will default to false. | [optional] 

## Methods

### NewVcenterDeploymentUpgradeUpgradeSpec

`func NewVcenterDeploymentUpgradeUpgradeSpec(sourceAppliance VcenterDeploymentUpgradeSourceApplianceSpec, sourceLocation VcenterDeploymentLocationSpec, ) *VcenterDeploymentUpgradeUpgradeSpec`

NewVcenterDeploymentUpgradeUpgradeSpec instantiates a new VcenterDeploymentUpgradeUpgradeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentUpgradeUpgradeSpecWithDefaults

`func NewVcenterDeploymentUpgradeUpgradeSpecWithDefaults() *VcenterDeploymentUpgradeUpgradeSpec`

NewVcenterDeploymentUpgradeUpgradeSpecWithDefaults instantiates a new VcenterDeploymentUpgradeUpgradeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAppliance

`func (o *VcenterDeploymentUpgradeUpgradeSpec) GetSourceAppliance() VcenterDeploymentUpgradeSourceApplianceSpec`

GetSourceAppliance returns the SourceAppliance field if non-nil, zero value otherwise.

### GetSourceApplianceOk

`func (o *VcenterDeploymentUpgradeUpgradeSpec) GetSourceApplianceOk() (*VcenterDeploymentUpgradeSourceApplianceSpec, bool)`

GetSourceApplianceOk returns a tuple with the SourceAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAppliance

`func (o *VcenterDeploymentUpgradeUpgradeSpec) SetSourceAppliance(v VcenterDeploymentUpgradeSourceApplianceSpec)`

SetSourceAppliance sets SourceAppliance field to given value.


### GetSourceLocation

`func (o *VcenterDeploymentUpgradeUpgradeSpec) GetSourceLocation() VcenterDeploymentLocationSpec`

GetSourceLocation returns the SourceLocation field if non-nil, zero value otherwise.

### GetSourceLocationOk

`func (o *VcenterDeploymentUpgradeUpgradeSpec) GetSourceLocationOk() (*VcenterDeploymentLocationSpec, bool)`

GetSourceLocationOk returns a tuple with the SourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocation

`func (o *VcenterDeploymentUpgradeUpgradeSpec) SetSourceLocation(v VcenterDeploymentLocationSpec)`

SetSourceLocation sets SourceLocation field to given value.


### GetHistory

`func (o *VcenterDeploymentUpgradeUpgradeSpec) GetHistory() VcenterDeploymentHistoryMigrationSpec`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *VcenterDeploymentUpgradeUpgradeSpec) GetHistoryOk() (*VcenterDeploymentHistoryMigrationSpec, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *VcenterDeploymentUpgradeUpgradeSpec) SetHistory(v VcenterDeploymentHistoryMigrationSpec)`

SetHistory sets History field to given value.

### HasHistory

`func (o *VcenterDeploymentUpgradeUpgradeSpec) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetVcsaEmbedded

`func (o *VcenterDeploymentUpgradeUpgradeSpec) GetVcsaEmbedded() VcenterDeploymentUpgradeVcsaEmbeddedSpec`

GetVcsaEmbedded returns the VcsaEmbedded field if non-nil, zero value otherwise.

### GetVcsaEmbeddedOk

`func (o *VcenterDeploymentUpgradeUpgradeSpec) GetVcsaEmbeddedOk() (*VcenterDeploymentUpgradeVcsaEmbeddedSpec, bool)`

GetVcsaEmbeddedOk returns a tuple with the VcsaEmbedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsaEmbedded

`func (o *VcenterDeploymentUpgradeUpgradeSpec) SetVcsaEmbedded(v VcenterDeploymentUpgradeVcsaEmbeddedSpec)`

SetVcsaEmbedded sets VcsaEmbedded field to given value.

### HasVcsaEmbedded

`func (o *VcenterDeploymentUpgradeUpgradeSpec) HasVcsaEmbedded() bool`

HasVcsaEmbedded returns a boolean if a field has been set.

### GetPsc

`func (o *VcenterDeploymentUpgradeUpgradeSpec) GetPsc() VcenterDeploymentUpgradePscSpec`

GetPsc returns the Psc field if non-nil, zero value otherwise.

### GetPscOk

`func (o *VcenterDeploymentUpgradeUpgradeSpec) GetPscOk() (*VcenterDeploymentUpgradePscSpec, bool)`

GetPscOk returns a tuple with the Psc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsc

`func (o *VcenterDeploymentUpgradeUpgradeSpec) SetPsc(v VcenterDeploymentUpgradePscSpec)`

SetPsc sets Psc field to given value.

### HasPsc

`func (o *VcenterDeploymentUpgradeUpgradeSpec) HasPsc() bool`

HasPsc returns a boolean if a field has been set.

### GetAutoAnswer

`func (o *VcenterDeploymentUpgradeUpgradeSpec) GetAutoAnswer() bool`

GetAutoAnswer returns the AutoAnswer field if non-nil, zero value otherwise.

### GetAutoAnswerOk

`func (o *VcenterDeploymentUpgradeUpgradeSpec) GetAutoAnswerOk() (*bool, bool)`

GetAutoAnswerOk returns a tuple with the AutoAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAnswer

`func (o *VcenterDeploymentUpgradeUpgradeSpec) SetAutoAnswer(v bool)`

SetAutoAnswer sets AutoAnswer field to given value.

### HasAutoAnswer

`func (o *VcenterDeploymentUpgradeUpgradeSpec) HasAutoAnswer() bool`

HasAutoAnswer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


