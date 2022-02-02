# VcenterOvfLibraryItemCreateOvfLibraryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientToken** | Pointer to **string** | Client-generated token used to retry a request if the client fails to get a response from the server. If the original request succeeded, the result of that request will be returned, otherwise the operation will be retried. | [optional] 
**Source** | Pointer to [**VcenterOvfLibraryItemDeployableIdentity**](VcenterOvfLibraryItemDeployableIdentity.md) |  | [optional] 
**Target** | Pointer to [**VcenterOvfLibraryItemCreateTarget**](VcenterOvfLibraryItemCreateTarget.md) |  | [optional] 
**CreateSpec** | Pointer to [**VcenterOvfLibraryItemCreateSpec**](VcenterOvfLibraryItemCreateSpec.md) |  | [optional] 

## Methods

### NewVcenterOvfLibraryItemCreateOvfLibraryItem

`func NewVcenterOvfLibraryItemCreateOvfLibraryItem() *VcenterOvfLibraryItemCreateOvfLibraryItem`

NewVcenterOvfLibraryItemCreateOvfLibraryItem instantiates a new VcenterOvfLibraryItemCreateOvfLibraryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfLibraryItemCreateOvfLibraryItemWithDefaults

`func NewVcenterOvfLibraryItemCreateOvfLibraryItemWithDefaults() *VcenterOvfLibraryItemCreateOvfLibraryItem`

NewVcenterOvfLibraryItemCreateOvfLibraryItemWithDefaults instantiates a new VcenterOvfLibraryItemCreateOvfLibraryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientToken

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) GetClientToken() string`

GetClientToken returns the ClientToken field if non-nil, zero value otherwise.

### GetClientTokenOk

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) GetClientTokenOk() (*string, bool)`

GetClientTokenOk returns a tuple with the ClientToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToken

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) SetClientToken(v string)`

SetClientToken sets ClientToken field to given value.

### HasClientToken

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) HasClientToken() bool`

HasClientToken returns a boolean if a field has been set.

### GetSource

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) GetSource() VcenterOvfLibraryItemDeployableIdentity`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) GetSourceOk() (*VcenterOvfLibraryItemDeployableIdentity, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) SetSource(v VcenterOvfLibraryItemDeployableIdentity)`

SetSource sets Source field to given value.

### HasSource

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) GetTarget() VcenterOvfLibraryItemCreateTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) GetTargetOk() (*VcenterOvfLibraryItemCreateTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) SetTarget(v VcenterOvfLibraryItemCreateTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetCreateSpec

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) GetCreateSpec() VcenterOvfLibraryItemCreateSpec`

GetCreateSpec returns the CreateSpec field if non-nil, zero value otherwise.

### GetCreateSpecOk

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) GetCreateSpecOk() (*VcenterOvfLibraryItemCreateSpec, bool)`

GetCreateSpecOk returns a tuple with the CreateSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSpec

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) SetCreateSpec(v VcenterOvfLibraryItemCreateSpec)`

SetCreateSpec sets CreateSpec field to given value.

### HasCreateSpec

`func (o *VcenterOvfLibraryItemCreateOvfLibraryItem) HasCreateSpec() bool`

HasCreateSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


