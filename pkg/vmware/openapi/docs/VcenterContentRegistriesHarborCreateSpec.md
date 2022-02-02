# VcenterContentRegistriesHarborCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to **string** | Identifier of the cluster hosting the registry. | [optional] 
**GarbageCollection** | Pointer to [**VcenterContentRegistriesHarborGarbageCollection**](VcenterContentRegistriesHarborGarbageCollection.md) |  | [optional] 
**Storage** | [**[]VcenterContentRegistriesHarborStorageSpec**](VcenterContentRegistriesHarborStorageSpec.md) | Storage associated with the Harbor registry. The list contains only one storage backing in this version. | 

## Methods

### NewVcenterContentRegistriesHarborCreateSpec

`func NewVcenterContentRegistriesHarborCreateSpec(storage []VcenterContentRegistriesHarborStorageSpec, ) *VcenterContentRegistriesHarborCreateSpec`

NewVcenterContentRegistriesHarborCreateSpec instantiates a new VcenterContentRegistriesHarborCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterContentRegistriesHarborCreateSpecWithDefaults

`func NewVcenterContentRegistriesHarborCreateSpecWithDefaults() *VcenterContentRegistriesHarborCreateSpec`

NewVcenterContentRegistriesHarborCreateSpecWithDefaults instantiates a new VcenterContentRegistriesHarborCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VcenterContentRegistriesHarborCreateSpec) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterContentRegistriesHarborCreateSpec) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterContentRegistriesHarborCreateSpec) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VcenterContentRegistriesHarborCreateSpec) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetGarbageCollection

`func (o *VcenterContentRegistriesHarborCreateSpec) GetGarbageCollection() VcenterContentRegistriesHarborGarbageCollection`

GetGarbageCollection returns the GarbageCollection field if non-nil, zero value otherwise.

### GetGarbageCollectionOk

`func (o *VcenterContentRegistriesHarborCreateSpec) GetGarbageCollectionOk() (*VcenterContentRegistriesHarborGarbageCollection, bool)`

GetGarbageCollectionOk returns a tuple with the GarbageCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarbageCollection

`func (o *VcenterContentRegistriesHarborCreateSpec) SetGarbageCollection(v VcenterContentRegistriesHarborGarbageCollection)`

SetGarbageCollection sets GarbageCollection field to given value.

### HasGarbageCollection

`func (o *VcenterContentRegistriesHarborCreateSpec) HasGarbageCollection() bool`

HasGarbageCollection returns a boolean if a field has been set.

### GetStorage

`func (o *VcenterContentRegistriesHarborCreateSpec) GetStorage() []VcenterContentRegistriesHarborStorageSpec`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *VcenterContentRegistriesHarborCreateSpec) GetStorageOk() (*[]VcenterContentRegistriesHarborStorageSpec, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *VcenterContentRegistriesHarborCreateSpec) SetStorage(v []VcenterContentRegistriesHarborStorageSpec)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


