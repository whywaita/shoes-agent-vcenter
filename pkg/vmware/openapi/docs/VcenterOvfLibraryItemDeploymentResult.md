# VcenterOvfLibraryItemDeploymentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Succeeded** | **bool** | Whether the {@name LibraryItem#deploy} {@term operation} completed successfully. | 
**ResourceId** | Pointer to [**VcenterOvfLibraryItemDeployableIdentity**](VcenterOvfLibraryItemDeployableIdentity.md) |  | [optional] 
**Error** | Pointer to [**VcenterOvfLibraryItemResultInfo**](VcenterOvfLibraryItemResultInfo.md) |  | [optional] 

## Methods

### NewVcenterOvfLibraryItemDeploymentResult

`func NewVcenterOvfLibraryItemDeploymentResult(succeeded bool, ) *VcenterOvfLibraryItemDeploymentResult`

NewVcenterOvfLibraryItemDeploymentResult instantiates a new VcenterOvfLibraryItemDeploymentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfLibraryItemDeploymentResultWithDefaults

`func NewVcenterOvfLibraryItemDeploymentResultWithDefaults() *VcenterOvfLibraryItemDeploymentResult`

NewVcenterOvfLibraryItemDeploymentResultWithDefaults instantiates a new VcenterOvfLibraryItemDeploymentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSucceeded

`func (o *VcenterOvfLibraryItemDeploymentResult) GetSucceeded() bool`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *VcenterOvfLibraryItemDeploymentResult) GetSucceededOk() (*bool, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *VcenterOvfLibraryItemDeploymentResult) SetSucceeded(v bool)`

SetSucceeded sets Succeeded field to given value.


### GetResourceId

`func (o *VcenterOvfLibraryItemDeploymentResult) GetResourceId() VcenterOvfLibraryItemDeployableIdentity`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *VcenterOvfLibraryItemDeploymentResult) GetResourceIdOk() (*VcenterOvfLibraryItemDeployableIdentity, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *VcenterOvfLibraryItemDeploymentResult) SetResourceId(v VcenterOvfLibraryItemDeployableIdentity)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *VcenterOvfLibraryItemDeploymentResult) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetError

`func (o *VcenterOvfLibraryItemDeploymentResult) GetError() VcenterOvfLibraryItemResultInfo`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VcenterOvfLibraryItemDeploymentResult) GetErrorOk() (*VcenterOvfLibraryItemResultInfo, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VcenterOvfLibraryItemDeploymentResult) SetError(v VcenterOvfLibraryItemResultInfo)`

SetError sets Error field to given value.

### HasError

`func (o *VcenterOvfLibraryItemDeploymentResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


