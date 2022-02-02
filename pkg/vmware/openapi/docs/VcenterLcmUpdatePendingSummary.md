# VcenterLcmUpdatePendingSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PendingUpdate** | **string** | Identifier of the given vSphere update When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.lcm.update.pending. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.lcm.update.pending. | 
**Version** | **string** | Version of the vSphere update or patch | 
**ReleaseDate** | **time.Time** | Release date of the vSphere update or patch | 
**Severity** | [**VcenterLcmUpdatePendingSeverityType**](VcenterLcmUpdatePendingSeverityType.md) |  | 
**Build** | **string** | Build number of the vCenter Release | 
**UpdateType** | [**VcenterLcmUpdatePendingUpdateType**](VcenterLcmUpdatePendingUpdateType.md) |  | 
**Category** | [**VcenterLcmUpdatePendingCategory**](VcenterLcmUpdatePendingCategory.md) |  | 
**RebootRequired** | **bool** | Flag to suggest a reboot after the release is applied | 
**ExecuteURL** | **string** | VAMI or ISO URL for update or upgrade execute phase redirection | 
**ReleaseNotes** | **[]string** | List of URI pointing to patch or update release notes | 

## Methods

### NewVcenterLcmUpdatePendingSummary

`func NewVcenterLcmUpdatePendingSummary(pendingUpdate string, version string, releaseDate time.Time, severity VcenterLcmUpdatePendingSeverityType, build string, updateType VcenterLcmUpdatePendingUpdateType, category VcenterLcmUpdatePendingCategory, rebootRequired bool, executeURL string, releaseNotes []string, ) *VcenterLcmUpdatePendingSummary`

NewVcenterLcmUpdatePendingSummary instantiates a new VcenterLcmUpdatePendingSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterLcmUpdatePendingSummaryWithDefaults

`func NewVcenterLcmUpdatePendingSummaryWithDefaults() *VcenterLcmUpdatePendingSummary`

NewVcenterLcmUpdatePendingSummaryWithDefaults instantiates a new VcenterLcmUpdatePendingSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPendingUpdate

`func (o *VcenterLcmUpdatePendingSummary) GetPendingUpdate() string`

GetPendingUpdate returns the PendingUpdate field if non-nil, zero value otherwise.

### GetPendingUpdateOk

`func (o *VcenterLcmUpdatePendingSummary) GetPendingUpdateOk() (*string, bool)`

GetPendingUpdateOk returns a tuple with the PendingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUpdate

`func (o *VcenterLcmUpdatePendingSummary) SetPendingUpdate(v string)`

SetPendingUpdate sets PendingUpdate field to given value.


### GetVersion

`func (o *VcenterLcmUpdatePendingSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterLcmUpdatePendingSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterLcmUpdatePendingSummary) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetReleaseDate

`func (o *VcenterLcmUpdatePendingSummary) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *VcenterLcmUpdatePendingSummary) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *VcenterLcmUpdatePendingSummary) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.


### GetSeverity

`func (o *VcenterLcmUpdatePendingSummary) GetSeverity() VcenterLcmUpdatePendingSeverityType`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *VcenterLcmUpdatePendingSummary) GetSeverityOk() (*VcenterLcmUpdatePendingSeverityType, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *VcenterLcmUpdatePendingSummary) SetSeverity(v VcenterLcmUpdatePendingSeverityType)`

SetSeverity sets Severity field to given value.


### GetBuild

`func (o *VcenterLcmUpdatePendingSummary) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *VcenterLcmUpdatePendingSummary) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *VcenterLcmUpdatePendingSummary) SetBuild(v string)`

SetBuild sets Build field to given value.


### GetUpdateType

`func (o *VcenterLcmUpdatePendingSummary) GetUpdateType() VcenterLcmUpdatePendingUpdateType`

GetUpdateType returns the UpdateType field if non-nil, zero value otherwise.

### GetUpdateTypeOk

`func (o *VcenterLcmUpdatePendingSummary) GetUpdateTypeOk() (*VcenterLcmUpdatePendingUpdateType, bool)`

GetUpdateTypeOk returns a tuple with the UpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateType

`func (o *VcenterLcmUpdatePendingSummary) SetUpdateType(v VcenterLcmUpdatePendingUpdateType)`

SetUpdateType sets UpdateType field to given value.


### GetCategory

`func (o *VcenterLcmUpdatePendingSummary) GetCategory() VcenterLcmUpdatePendingCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *VcenterLcmUpdatePendingSummary) GetCategoryOk() (*VcenterLcmUpdatePendingCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *VcenterLcmUpdatePendingSummary) SetCategory(v VcenterLcmUpdatePendingCategory)`

SetCategory sets Category field to given value.


### GetRebootRequired

`func (o *VcenterLcmUpdatePendingSummary) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *VcenterLcmUpdatePendingSummary) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *VcenterLcmUpdatePendingSummary) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.


### GetExecuteURL

`func (o *VcenterLcmUpdatePendingSummary) GetExecuteURL() string`

GetExecuteURL returns the ExecuteURL field if non-nil, zero value otherwise.

### GetExecuteURLOk

`func (o *VcenterLcmUpdatePendingSummary) GetExecuteURLOk() (*string, bool)`

GetExecuteURLOk returns a tuple with the ExecuteURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteURL

`func (o *VcenterLcmUpdatePendingSummary) SetExecuteURL(v string)`

SetExecuteURL sets ExecuteURL field to given value.


### GetReleaseNotes

`func (o *VcenterLcmUpdatePendingSummary) GetReleaseNotes() []string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *VcenterLcmUpdatePendingSummary) GetReleaseNotesOk() (*[]string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *VcenterLcmUpdatePendingSummary) SetReleaseNotes(v []string)`

SetReleaseNotes sets ReleaseNotes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


