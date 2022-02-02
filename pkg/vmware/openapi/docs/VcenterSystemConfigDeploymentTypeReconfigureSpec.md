# VcenterSystemConfigDeploymentTypeReconfigureSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterDeploymentApplianceType**](VcenterDeploymentApplianceType.md) |  | 
**RemotePsc** | Pointer to [**VcenterDeploymentRemotePscSpec**](VcenterDeploymentRemotePscSpec.md) |  | [optional] 

## Methods

### NewVcenterSystemConfigDeploymentTypeReconfigureSpec

`func NewVcenterSystemConfigDeploymentTypeReconfigureSpec(type_ VcenterDeploymentApplianceType, ) *VcenterSystemConfigDeploymentTypeReconfigureSpec`

NewVcenterSystemConfigDeploymentTypeReconfigureSpec instantiates a new VcenterSystemConfigDeploymentTypeReconfigureSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterSystemConfigDeploymentTypeReconfigureSpecWithDefaults

`func NewVcenterSystemConfigDeploymentTypeReconfigureSpecWithDefaults() *VcenterSystemConfigDeploymentTypeReconfigureSpec`

NewVcenterSystemConfigDeploymentTypeReconfigureSpecWithDefaults instantiates a new VcenterSystemConfigDeploymentTypeReconfigureSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterSystemConfigDeploymentTypeReconfigureSpec) GetType() VcenterDeploymentApplianceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterSystemConfigDeploymentTypeReconfigureSpec) GetTypeOk() (*VcenterDeploymentApplianceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterSystemConfigDeploymentTypeReconfigureSpec) SetType(v VcenterDeploymentApplianceType)`

SetType sets Type field to given value.


### GetRemotePsc

`func (o *VcenterSystemConfigDeploymentTypeReconfigureSpec) GetRemotePsc() VcenterDeploymentRemotePscSpec`

GetRemotePsc returns the RemotePsc field if non-nil, zero value otherwise.

### GetRemotePscOk

`func (o *VcenterSystemConfigDeploymentTypeReconfigureSpec) GetRemotePscOk() (*VcenterDeploymentRemotePscSpec, bool)`

GetRemotePscOk returns a tuple with the RemotePsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePsc

`func (o *VcenterSystemConfigDeploymentTypeReconfigureSpec) SetRemotePsc(v VcenterDeploymentRemotePscSpec)`

SetRemotePsc sets RemotePsc field to given value.

### HasRemotePsc

`func (o *VcenterSystemConfigDeploymentTypeReconfigureSpec) HasRemotePsc() bool`

HasRemotePsc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


