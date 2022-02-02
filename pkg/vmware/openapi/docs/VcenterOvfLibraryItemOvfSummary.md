# VcenterOvfLibraryItemOvfSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Default name for the virtual machine or virtual appliance. | [optional] 
**Annotation** | Pointer to **string** | Default annotation for the virtual machine or virtual appliance. | [optional] 
**EULAs** | **[]string** | End User License Agreements specified in the OVF descriptor. All end user license agreements must be accepted in order for the {@name LibraryItem#deploy} {@term operation} to succeed. See {@link ResourcePoolDeploymentSpec#acceptAllEula}. | 
**Networks** | Pointer to **[]string** | Section identifiers for sections of type ovf:NetworkSection in the OVF descriptor. These identifiers can be used as keys in {@link ResourcePoolDeploymentSpec#networkMappings}. | [optional] 
**StorageGroups** | Pointer to **[]string** | Section identifiers for sections of type vmw:StorageGroupSection in the OVF descriptor. These identifiers can be used as keys in {@link ResourcePoolDeploymentSpec#storageMappings}. | [optional] 
**AdditionalParams** | Pointer to **[]map[string]interface{}** | Additional OVF parameters which can be specified for the deployment target. These OVF parameters can be inspected, optionally modified, and used as values in {@link ResourcePoolDeploymentSpec#additionalParameters} for the {@name LibraryItem#deploy} {@term operation}. | [optional] 

## Methods

### NewVcenterOvfLibraryItemOvfSummary

`func NewVcenterOvfLibraryItemOvfSummary(eULAs []string, ) *VcenterOvfLibraryItemOvfSummary`

NewVcenterOvfLibraryItemOvfSummary instantiates a new VcenterOvfLibraryItemOvfSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfLibraryItemOvfSummaryWithDefaults

`func NewVcenterOvfLibraryItemOvfSummaryWithDefaults() *VcenterOvfLibraryItemOvfSummary`

NewVcenterOvfLibraryItemOvfSummaryWithDefaults instantiates a new VcenterOvfLibraryItemOvfSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterOvfLibraryItemOvfSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterOvfLibraryItemOvfSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterOvfLibraryItemOvfSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterOvfLibraryItemOvfSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotation

`func (o *VcenterOvfLibraryItemOvfSummary) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *VcenterOvfLibraryItemOvfSummary) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *VcenterOvfLibraryItemOvfSummary) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *VcenterOvfLibraryItemOvfSummary) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetEULAs

`func (o *VcenterOvfLibraryItemOvfSummary) GetEULAs() []string`

GetEULAs returns the EULAs field if non-nil, zero value otherwise.

### GetEULAsOk

`func (o *VcenterOvfLibraryItemOvfSummary) GetEULAsOk() (*[]string, bool)`

GetEULAsOk returns a tuple with the EULAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEULAs

`func (o *VcenterOvfLibraryItemOvfSummary) SetEULAs(v []string)`

SetEULAs sets EULAs field to given value.


### GetNetworks

`func (o *VcenterOvfLibraryItemOvfSummary) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *VcenterOvfLibraryItemOvfSummary) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *VcenterOvfLibraryItemOvfSummary) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *VcenterOvfLibraryItemOvfSummary) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetStorageGroups

`func (o *VcenterOvfLibraryItemOvfSummary) GetStorageGroups() []string`

GetStorageGroups returns the StorageGroups field if non-nil, zero value otherwise.

### GetStorageGroupsOk

`func (o *VcenterOvfLibraryItemOvfSummary) GetStorageGroupsOk() (*[]string, bool)`

GetStorageGroupsOk returns a tuple with the StorageGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGroups

`func (o *VcenterOvfLibraryItemOvfSummary) SetStorageGroups(v []string)`

SetStorageGroups sets StorageGroups field to given value.

### HasStorageGroups

`func (o *VcenterOvfLibraryItemOvfSummary) HasStorageGroups() bool`

HasStorageGroups returns a boolean if a field has been set.

### GetAdditionalParams

`func (o *VcenterOvfLibraryItemOvfSummary) GetAdditionalParams() []map[string]interface{}`

GetAdditionalParams returns the AdditionalParams field if non-nil, zero value otherwise.

### GetAdditionalParamsOk

`func (o *VcenterOvfLibraryItemOvfSummary) GetAdditionalParamsOk() (*[]map[string]interface{}, bool)`

GetAdditionalParamsOk returns a tuple with the AdditionalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParams

`func (o *VcenterOvfLibraryItemOvfSummary) SetAdditionalParams(v []map[string]interface{})`

SetAdditionalParams sets AdditionalParams field to given value.

### HasAdditionalParams

`func (o *VcenterOvfLibraryItemOvfSummary) HasAdditionalParams() bool`

HasAdditionalParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


