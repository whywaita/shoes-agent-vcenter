# VcenterOvfLibraryItemDeployOvfLibraryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientToken** | Pointer to **string** | Client-generated token used to retry a request if the client fails to get a response from the server. If the original request succeeded, the result of that request will be returned, otherwise the operation will be retried. | [optional] 
**Target** | Pointer to [**VcenterOvfLibraryItemDeploymentTarget**](VcenterOvfLibraryItemDeploymentTarget.md) |  | [optional] 
**DeploymentSpec** | Pointer to [**VcenterOvfLibraryItemResourcePoolDeploymentSpec**](VcenterOvfLibraryItemResourcePoolDeploymentSpec.md) |  | [optional] 

## Methods

### NewVcenterOvfLibraryItemDeployOvfLibraryItem

`func NewVcenterOvfLibraryItemDeployOvfLibraryItem() *VcenterOvfLibraryItemDeployOvfLibraryItem`

NewVcenterOvfLibraryItemDeployOvfLibraryItem instantiates a new VcenterOvfLibraryItemDeployOvfLibraryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfLibraryItemDeployOvfLibraryItemWithDefaults

`func NewVcenterOvfLibraryItemDeployOvfLibraryItemWithDefaults() *VcenterOvfLibraryItemDeployOvfLibraryItem`

NewVcenterOvfLibraryItemDeployOvfLibraryItemWithDefaults instantiates a new VcenterOvfLibraryItemDeployOvfLibraryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientToken

`func (o *VcenterOvfLibraryItemDeployOvfLibraryItem) GetClientToken() string`

GetClientToken returns the ClientToken field if non-nil, zero value otherwise.

### GetClientTokenOk

`func (o *VcenterOvfLibraryItemDeployOvfLibraryItem) GetClientTokenOk() (*string, bool)`

GetClientTokenOk returns a tuple with the ClientToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToken

`func (o *VcenterOvfLibraryItemDeployOvfLibraryItem) SetClientToken(v string)`

SetClientToken sets ClientToken field to given value.

### HasClientToken

`func (o *VcenterOvfLibraryItemDeployOvfLibraryItem) HasClientToken() bool`

HasClientToken returns a boolean if a field has been set.

### GetTarget

`func (o *VcenterOvfLibraryItemDeployOvfLibraryItem) GetTarget() VcenterOvfLibraryItemDeploymentTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *VcenterOvfLibraryItemDeployOvfLibraryItem) GetTargetOk() (*VcenterOvfLibraryItemDeploymentTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *VcenterOvfLibraryItemDeployOvfLibraryItem) SetTarget(v VcenterOvfLibraryItemDeploymentTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *VcenterOvfLibraryItemDeployOvfLibraryItem) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetDeploymentSpec

`func (o *VcenterOvfLibraryItemDeployOvfLibraryItem) GetDeploymentSpec() VcenterOvfLibraryItemResourcePoolDeploymentSpec`

GetDeploymentSpec returns the DeploymentSpec field if non-nil, zero value otherwise.

### GetDeploymentSpecOk

`func (o *VcenterOvfLibraryItemDeployOvfLibraryItem) GetDeploymentSpecOk() (*VcenterOvfLibraryItemResourcePoolDeploymentSpec, bool)`

GetDeploymentSpecOk returns a tuple with the DeploymentSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSpec

`func (o *VcenterOvfLibraryItemDeployOvfLibraryItem) SetDeploymentSpec(v VcenterOvfLibraryItemResourcePoolDeploymentSpec)`

SetDeploymentSpec sets DeploymentSpec field to given value.

### HasDeploymentSpec

`func (o *VcenterOvfLibraryItemDeployOvfLibraryItem) HasDeploymentSpec() bool`

HasDeploymentSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


