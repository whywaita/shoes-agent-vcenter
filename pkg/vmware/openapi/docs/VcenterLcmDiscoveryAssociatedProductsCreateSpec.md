# VcenterLcmDiscoveryAssociatedProductsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductName** | **string** | The name of the product. | 
**Version** | **string** | Current product version. | 
**Deployments** | Pointer to **[]string** | The list of hostname/IPs of the instances of the VMware products deployed in the environment. | [optional] 

## Methods

### NewVcenterLcmDiscoveryAssociatedProductsCreateSpec

`func NewVcenterLcmDiscoveryAssociatedProductsCreateSpec(productName string, version string, ) *VcenterLcmDiscoveryAssociatedProductsCreateSpec`

NewVcenterLcmDiscoveryAssociatedProductsCreateSpec instantiates a new VcenterLcmDiscoveryAssociatedProductsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterLcmDiscoveryAssociatedProductsCreateSpecWithDefaults

`func NewVcenterLcmDiscoveryAssociatedProductsCreateSpecWithDefaults() *VcenterLcmDiscoveryAssociatedProductsCreateSpec`

NewVcenterLcmDiscoveryAssociatedProductsCreateSpecWithDefaults instantiates a new VcenterLcmDiscoveryAssociatedProductsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductName

`func (o *VcenterLcmDiscoveryAssociatedProductsCreateSpec) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *VcenterLcmDiscoveryAssociatedProductsCreateSpec) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *VcenterLcmDiscoveryAssociatedProductsCreateSpec) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetVersion

`func (o *VcenterLcmDiscoveryAssociatedProductsCreateSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterLcmDiscoveryAssociatedProductsCreateSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterLcmDiscoveryAssociatedProductsCreateSpec) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDeployments

`func (o *VcenterLcmDiscoveryAssociatedProductsCreateSpec) GetDeployments() []string`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *VcenterLcmDiscoveryAssociatedProductsCreateSpec) GetDeploymentsOk() (*[]string, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *VcenterLcmDiscoveryAssociatedProductsCreateSpec) SetDeployments(v []string)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *VcenterLcmDiscoveryAssociatedProductsCreateSpec) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


