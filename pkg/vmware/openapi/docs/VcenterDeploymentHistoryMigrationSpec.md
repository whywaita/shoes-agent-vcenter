# VcenterDeploymentHistoryMigrationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSet** | [**VcenterDeploymentHistoryMigrationOption**](VcenterDeploymentHistoryMigrationOption.md) |  | 
**DeferImport** | Pointer to **bool** | Defines how vCenter history will be migrated. If set to true, vCenter history will be migrated separately after successful upgrade(supported scenarios are upgrade from 6.0 or 6.5 to 6.7) or migration, otherwise it will be migrated along with core data during the upgrade or migration process. vCSA upgrade with deferred import is no longer supported for target version 7.0 and later. If unset, vCenter historical data won&#39;t be deferred and will be migrated along with core data. | [optional] 

## Methods

### NewVcenterDeploymentHistoryMigrationSpec

`func NewVcenterDeploymentHistoryMigrationSpec(dataSet VcenterDeploymentHistoryMigrationOption, ) *VcenterDeploymentHistoryMigrationSpec`

NewVcenterDeploymentHistoryMigrationSpec instantiates a new VcenterDeploymentHistoryMigrationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentHistoryMigrationSpecWithDefaults

`func NewVcenterDeploymentHistoryMigrationSpecWithDefaults() *VcenterDeploymentHistoryMigrationSpec`

NewVcenterDeploymentHistoryMigrationSpecWithDefaults instantiates a new VcenterDeploymentHistoryMigrationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSet

`func (o *VcenterDeploymentHistoryMigrationSpec) GetDataSet() VcenterDeploymentHistoryMigrationOption`

GetDataSet returns the DataSet field if non-nil, zero value otherwise.

### GetDataSetOk

`func (o *VcenterDeploymentHistoryMigrationSpec) GetDataSetOk() (*VcenterDeploymentHistoryMigrationOption, bool)`

GetDataSetOk returns a tuple with the DataSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSet

`func (o *VcenterDeploymentHistoryMigrationSpec) SetDataSet(v VcenterDeploymentHistoryMigrationOption)`

SetDataSet sets DataSet field to given value.


### GetDeferImport

`func (o *VcenterDeploymentHistoryMigrationSpec) GetDeferImport() bool`

GetDeferImport returns the DeferImport field if non-nil, zero value otherwise.

### GetDeferImportOk

`func (o *VcenterDeploymentHistoryMigrationSpec) GetDeferImportOk() (*bool, bool)`

GetDeferImportOk returns a tuple with the DeferImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferImport

`func (o *VcenterDeploymentHistoryMigrationSpec) SetDeferImport(v bool)`

SetDeferImport sets DeferImport field to given value.

### HasDeferImport

`func (o *VcenterDeploymentHistoryMigrationSpec) HasDeferImport() bool`

HasDeferImport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


