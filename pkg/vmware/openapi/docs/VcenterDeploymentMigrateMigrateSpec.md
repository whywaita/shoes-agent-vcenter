# VcenterDeploymentMigrateMigrateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceVcWindows** | [**VcenterDeploymentMigrateSourceVcWindows**](VcenterDeploymentMigrateSourceVcWindows.md) |  | 
**ExistingMigrationAssistant** | [**VcenterDeploymentMigrateMigrationAssistantSpec**](VcenterDeploymentMigrateMigrationAssistantSpec.md) |  | 
**History** | Pointer to [**VcenterDeploymentHistoryMigrationSpec**](VcenterDeploymentHistoryMigrationSpec.md) |  | [optional] 
**VcsaEmbedded** | Pointer to [**VcenterDeploymentMigrateVcsaEmbeddedSpec**](VcenterDeploymentMigrateVcsaEmbeddedSpec.md) |  | [optional] 
**Psc** | Pointer to [**VcenterDeploymentMigratePscSpec**](VcenterDeploymentMigratePscSpec.md) |  | [optional] 
**ActiveDirectory** | Pointer to [**VcenterDeploymentMigrateActiveDirectorySpec**](VcenterDeploymentMigrateActiveDirectorySpec.md) |  | [optional] 
**AutoAnswer** | Pointer to **bool** | Use the default option for any questions that may come up during appliance configuration. If unset, will default to false. | [optional] 

## Methods

### NewVcenterDeploymentMigrateMigrateSpec

`func NewVcenterDeploymentMigrateMigrateSpec(sourceVcWindows VcenterDeploymentMigrateSourceVcWindows, existingMigrationAssistant VcenterDeploymentMigrateMigrationAssistantSpec, ) *VcenterDeploymentMigrateMigrateSpec`

NewVcenterDeploymentMigrateMigrateSpec instantiates a new VcenterDeploymentMigrateMigrateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentMigrateMigrateSpecWithDefaults

`func NewVcenterDeploymentMigrateMigrateSpecWithDefaults() *VcenterDeploymentMigrateMigrateSpec`

NewVcenterDeploymentMigrateMigrateSpecWithDefaults instantiates a new VcenterDeploymentMigrateMigrateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceVcWindows

`func (o *VcenterDeploymentMigrateMigrateSpec) GetSourceVcWindows() VcenterDeploymentMigrateSourceVcWindows`

GetSourceVcWindows returns the SourceVcWindows field if non-nil, zero value otherwise.

### GetSourceVcWindowsOk

`func (o *VcenterDeploymentMigrateMigrateSpec) GetSourceVcWindowsOk() (*VcenterDeploymentMigrateSourceVcWindows, bool)`

GetSourceVcWindowsOk returns a tuple with the SourceVcWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVcWindows

`func (o *VcenterDeploymentMigrateMigrateSpec) SetSourceVcWindows(v VcenterDeploymentMigrateSourceVcWindows)`

SetSourceVcWindows sets SourceVcWindows field to given value.


### GetExistingMigrationAssistant

`func (o *VcenterDeploymentMigrateMigrateSpec) GetExistingMigrationAssistant() VcenterDeploymentMigrateMigrationAssistantSpec`

GetExistingMigrationAssistant returns the ExistingMigrationAssistant field if non-nil, zero value otherwise.

### GetExistingMigrationAssistantOk

`func (o *VcenterDeploymentMigrateMigrateSpec) GetExistingMigrationAssistantOk() (*VcenterDeploymentMigrateMigrationAssistantSpec, bool)`

GetExistingMigrationAssistantOk returns a tuple with the ExistingMigrationAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingMigrationAssistant

`func (o *VcenterDeploymentMigrateMigrateSpec) SetExistingMigrationAssistant(v VcenterDeploymentMigrateMigrationAssistantSpec)`

SetExistingMigrationAssistant sets ExistingMigrationAssistant field to given value.


### GetHistory

`func (o *VcenterDeploymentMigrateMigrateSpec) GetHistory() VcenterDeploymentHistoryMigrationSpec`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *VcenterDeploymentMigrateMigrateSpec) GetHistoryOk() (*VcenterDeploymentHistoryMigrationSpec, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *VcenterDeploymentMigrateMigrateSpec) SetHistory(v VcenterDeploymentHistoryMigrationSpec)`

SetHistory sets History field to given value.

### HasHistory

`func (o *VcenterDeploymentMigrateMigrateSpec) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetVcsaEmbedded

`func (o *VcenterDeploymentMigrateMigrateSpec) GetVcsaEmbedded() VcenterDeploymentMigrateVcsaEmbeddedSpec`

GetVcsaEmbedded returns the VcsaEmbedded field if non-nil, zero value otherwise.

### GetVcsaEmbeddedOk

`func (o *VcenterDeploymentMigrateMigrateSpec) GetVcsaEmbeddedOk() (*VcenterDeploymentMigrateVcsaEmbeddedSpec, bool)`

GetVcsaEmbeddedOk returns a tuple with the VcsaEmbedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsaEmbedded

`func (o *VcenterDeploymentMigrateMigrateSpec) SetVcsaEmbedded(v VcenterDeploymentMigrateVcsaEmbeddedSpec)`

SetVcsaEmbedded sets VcsaEmbedded field to given value.

### HasVcsaEmbedded

`func (o *VcenterDeploymentMigrateMigrateSpec) HasVcsaEmbedded() bool`

HasVcsaEmbedded returns a boolean if a field has been set.

### GetPsc

`func (o *VcenterDeploymentMigrateMigrateSpec) GetPsc() VcenterDeploymentMigratePscSpec`

GetPsc returns the Psc field if non-nil, zero value otherwise.

### GetPscOk

`func (o *VcenterDeploymentMigrateMigrateSpec) GetPscOk() (*VcenterDeploymentMigratePscSpec, bool)`

GetPscOk returns a tuple with the Psc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsc

`func (o *VcenterDeploymentMigrateMigrateSpec) SetPsc(v VcenterDeploymentMigratePscSpec)`

SetPsc sets Psc field to given value.

### HasPsc

`func (o *VcenterDeploymentMigrateMigrateSpec) HasPsc() bool`

HasPsc returns a boolean if a field has been set.

### GetActiveDirectory

`func (o *VcenterDeploymentMigrateMigrateSpec) GetActiveDirectory() VcenterDeploymentMigrateActiveDirectorySpec`

GetActiveDirectory returns the ActiveDirectory field if non-nil, zero value otherwise.

### GetActiveDirectoryOk

`func (o *VcenterDeploymentMigrateMigrateSpec) GetActiveDirectoryOk() (*VcenterDeploymentMigrateActiveDirectorySpec, bool)`

GetActiveDirectoryOk returns a tuple with the ActiveDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectory

`func (o *VcenterDeploymentMigrateMigrateSpec) SetActiveDirectory(v VcenterDeploymentMigrateActiveDirectorySpec)`

SetActiveDirectory sets ActiveDirectory field to given value.

### HasActiveDirectory

`func (o *VcenterDeploymentMigrateMigrateSpec) HasActiveDirectory() bool`

HasActiveDirectory returns a boolean if a field has been set.

### GetAutoAnswer

`func (o *VcenterDeploymentMigrateMigrateSpec) GetAutoAnswer() bool`

GetAutoAnswer returns the AutoAnswer field if non-nil, zero value otherwise.

### GetAutoAnswerOk

`func (o *VcenterDeploymentMigrateMigrateSpec) GetAutoAnswerOk() (*bool, bool)`

GetAutoAnswerOk returns a tuple with the AutoAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAnswer

`func (o *VcenterDeploymentMigrateMigrateSpec) SetAutoAnswer(v bool)`

SetAutoAnswer sets AutoAnswer field to given value.

### HasAutoAnswer

`func (o *VcenterDeploymentMigrateMigrateSpec) HasAutoAnswer() bool`

HasAutoAnswer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


