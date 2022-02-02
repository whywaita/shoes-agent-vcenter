# VcenterDeploymentDataMigrationEstimate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedExportTime** | **int64** | The time estimated to export data from the source vCenter Server. | 
**EstimatedImportTime** | **int64** | The time estimated to import data to the upgraded vCenter Server. | 
**RequiredFreeDiskSpaceOnSource** | **float64** | The extra free space required on source vCenter Server. | 

## Methods

### NewVcenterDeploymentDataMigrationEstimate

`func NewVcenterDeploymentDataMigrationEstimate(estimatedExportTime int64, estimatedImportTime int64, requiredFreeDiskSpaceOnSource float64, ) *VcenterDeploymentDataMigrationEstimate`

NewVcenterDeploymentDataMigrationEstimate instantiates a new VcenterDeploymentDataMigrationEstimate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentDataMigrationEstimateWithDefaults

`func NewVcenterDeploymentDataMigrationEstimateWithDefaults() *VcenterDeploymentDataMigrationEstimate`

NewVcenterDeploymentDataMigrationEstimateWithDefaults instantiates a new VcenterDeploymentDataMigrationEstimate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedExportTime

`func (o *VcenterDeploymentDataMigrationEstimate) GetEstimatedExportTime() int64`

GetEstimatedExportTime returns the EstimatedExportTime field if non-nil, zero value otherwise.

### GetEstimatedExportTimeOk

`func (o *VcenterDeploymentDataMigrationEstimate) GetEstimatedExportTimeOk() (*int64, bool)`

GetEstimatedExportTimeOk returns a tuple with the EstimatedExportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedExportTime

`func (o *VcenterDeploymentDataMigrationEstimate) SetEstimatedExportTime(v int64)`

SetEstimatedExportTime sets EstimatedExportTime field to given value.


### GetEstimatedImportTime

`func (o *VcenterDeploymentDataMigrationEstimate) GetEstimatedImportTime() int64`

GetEstimatedImportTime returns the EstimatedImportTime field if non-nil, zero value otherwise.

### GetEstimatedImportTimeOk

`func (o *VcenterDeploymentDataMigrationEstimate) GetEstimatedImportTimeOk() (*int64, bool)`

GetEstimatedImportTimeOk returns a tuple with the EstimatedImportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedImportTime

`func (o *VcenterDeploymentDataMigrationEstimate) SetEstimatedImportTime(v int64)`

SetEstimatedImportTime sets EstimatedImportTime field to given value.


### GetRequiredFreeDiskSpaceOnSource

`func (o *VcenterDeploymentDataMigrationEstimate) GetRequiredFreeDiskSpaceOnSource() float64`

GetRequiredFreeDiskSpaceOnSource returns the RequiredFreeDiskSpaceOnSource field if non-nil, zero value otherwise.

### GetRequiredFreeDiskSpaceOnSourceOk

`func (o *VcenterDeploymentDataMigrationEstimate) GetRequiredFreeDiskSpaceOnSourceOk() (*float64, bool)`

GetRequiredFreeDiskSpaceOnSourceOk returns a tuple with the RequiredFreeDiskSpaceOnSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFreeDiskSpaceOnSource

`func (o *VcenterDeploymentDataMigrationEstimate) SetRequiredFreeDiskSpaceOnSource(v float64)`

SetRequiredFreeDiskSpaceOnSource sets RequiredFreeDiskSpaceOnSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


