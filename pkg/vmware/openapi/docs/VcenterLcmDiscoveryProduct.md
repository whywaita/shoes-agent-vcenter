# VcenterLcmDiscoveryProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstalledProduct** | **string** | Identifies a product and a version uniquely.  The identifier consists of product internal name and version.  When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: PRODUCT. When operations return a value of this structure as a result, the field will be an identifier for the resource type: PRODUCT. | 
**Name** | **string** | A public official product name. | 
**Version** | **string** | Current product version. | 
**TargetVersion** | Pointer to **string** | Future version of the product after upgrade. Product.target-version may not be applicable. | [optional] 
**Deployments** | Pointer to **[]string** | The list of hostname/IPs of the instances of the VMware products deployed in the environment. This field would be empty for manually added products. | [optional] 
**Auto** | **bool** | Indicates if the product is auto-detected by the system or manually added. If it is set to true it means it is auto-detected. | 

## Methods

### NewVcenterLcmDiscoveryProduct

`func NewVcenterLcmDiscoveryProduct(installedProduct string, name string, version string, auto bool, ) *VcenterLcmDiscoveryProduct`

NewVcenterLcmDiscoveryProduct instantiates a new VcenterLcmDiscoveryProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterLcmDiscoveryProductWithDefaults

`func NewVcenterLcmDiscoveryProductWithDefaults() *VcenterLcmDiscoveryProduct`

NewVcenterLcmDiscoveryProductWithDefaults instantiates a new VcenterLcmDiscoveryProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstalledProduct

`func (o *VcenterLcmDiscoveryProduct) GetInstalledProduct() string`

GetInstalledProduct returns the InstalledProduct field if non-nil, zero value otherwise.

### GetInstalledProductOk

`func (o *VcenterLcmDiscoveryProduct) GetInstalledProductOk() (*string, bool)`

GetInstalledProductOk returns a tuple with the InstalledProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledProduct

`func (o *VcenterLcmDiscoveryProduct) SetInstalledProduct(v string)`

SetInstalledProduct sets InstalledProduct field to given value.


### GetName

`func (o *VcenterLcmDiscoveryProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterLcmDiscoveryProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterLcmDiscoveryProduct) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *VcenterLcmDiscoveryProduct) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterLcmDiscoveryProduct) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterLcmDiscoveryProduct) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetTargetVersion

`func (o *VcenterLcmDiscoveryProduct) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *VcenterLcmDiscoveryProduct) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *VcenterLcmDiscoveryProduct) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *VcenterLcmDiscoveryProduct) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetDeployments

`func (o *VcenterLcmDiscoveryProduct) GetDeployments() []string`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *VcenterLcmDiscoveryProduct) GetDeploymentsOk() (*[]string, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *VcenterLcmDiscoveryProduct) SetDeployments(v []string)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *VcenterLcmDiscoveryProduct) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetAuto

`func (o *VcenterLcmDiscoveryProduct) GetAuto() bool`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *VcenterLcmDiscoveryProduct) GetAutoOk() (*bool, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *VcenterLcmDiscoveryProduct) SetAuto(v bool)`

SetAuto sets Auto field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


