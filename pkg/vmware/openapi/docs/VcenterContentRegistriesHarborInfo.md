# VcenterContentRegistriesHarborInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to **string** | Identifier of the cluster. | [optional] 
**Namespace** | Pointer to **string** | Identifier of the Harbor namespace in case it is created in a Kubernetes environment. | [optional] 
**Version** | **string** | Version of the registry. | 
**CreationTime** | **time.Time** | The date and time when the harbor registry was created. | 
**UiAccessUrl** | **string** | URL to access the UI of the registry. | 
**CertChain** | **[]string** | Harbor certificate chain in base64 format. | 
**GarbageCollection** | [**VcenterContentRegistriesHarborGarbageCollection**](VcenterContentRegistriesHarborGarbageCollection.md) |  | 
**Storage** | [**[]VcenterContentRegistriesHarborStorageInfo**](VcenterContentRegistriesHarborStorageInfo.md) | Storage information associated with the registry. | 
**Health** | [**VcenterContentRegistriesHealthInfo**](VcenterContentRegistriesHealthInfo.md) |  | 

## Methods

### NewVcenterContentRegistriesHarborInfo

`func NewVcenterContentRegistriesHarborInfo(version string, creationTime time.Time, uiAccessUrl string, certChain []string, garbageCollection VcenterContentRegistriesHarborGarbageCollection, storage []VcenterContentRegistriesHarborStorageInfo, health VcenterContentRegistriesHealthInfo, ) *VcenterContentRegistriesHarborInfo`

NewVcenterContentRegistriesHarborInfo instantiates a new VcenterContentRegistriesHarborInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterContentRegistriesHarborInfoWithDefaults

`func NewVcenterContentRegistriesHarborInfoWithDefaults() *VcenterContentRegistriesHarborInfo`

NewVcenterContentRegistriesHarborInfoWithDefaults instantiates a new VcenterContentRegistriesHarborInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VcenterContentRegistriesHarborInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterContentRegistriesHarborInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterContentRegistriesHarborInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VcenterContentRegistriesHarborInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetNamespace

`func (o *VcenterContentRegistriesHarborInfo) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *VcenterContentRegistriesHarborInfo) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *VcenterContentRegistriesHarborInfo) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *VcenterContentRegistriesHarborInfo) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetVersion

`func (o *VcenterContentRegistriesHarborInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterContentRegistriesHarborInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterContentRegistriesHarborInfo) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetCreationTime

`func (o *VcenterContentRegistriesHarborInfo) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *VcenterContentRegistriesHarborInfo) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *VcenterContentRegistriesHarborInfo) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetUiAccessUrl

`func (o *VcenterContentRegistriesHarborInfo) GetUiAccessUrl() string`

GetUiAccessUrl returns the UiAccessUrl field if non-nil, zero value otherwise.

### GetUiAccessUrlOk

`func (o *VcenterContentRegistriesHarborInfo) GetUiAccessUrlOk() (*string, bool)`

GetUiAccessUrlOk returns a tuple with the UiAccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiAccessUrl

`func (o *VcenterContentRegistriesHarborInfo) SetUiAccessUrl(v string)`

SetUiAccessUrl sets UiAccessUrl field to given value.


### GetCertChain

`func (o *VcenterContentRegistriesHarborInfo) GetCertChain() []string`

GetCertChain returns the CertChain field if non-nil, zero value otherwise.

### GetCertChainOk

`func (o *VcenterContentRegistriesHarborInfo) GetCertChainOk() (*[]string, bool)`

GetCertChainOk returns a tuple with the CertChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertChain

`func (o *VcenterContentRegistriesHarborInfo) SetCertChain(v []string)`

SetCertChain sets CertChain field to given value.


### GetGarbageCollection

`func (o *VcenterContentRegistriesHarborInfo) GetGarbageCollection() VcenterContentRegistriesHarborGarbageCollection`

GetGarbageCollection returns the GarbageCollection field if non-nil, zero value otherwise.

### GetGarbageCollectionOk

`func (o *VcenterContentRegistriesHarborInfo) GetGarbageCollectionOk() (*VcenterContentRegistriesHarborGarbageCollection, bool)`

GetGarbageCollectionOk returns a tuple with the GarbageCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarbageCollection

`func (o *VcenterContentRegistriesHarborInfo) SetGarbageCollection(v VcenterContentRegistriesHarborGarbageCollection)`

SetGarbageCollection sets GarbageCollection field to given value.


### GetStorage

`func (o *VcenterContentRegistriesHarborInfo) GetStorage() []VcenterContentRegistriesHarborStorageInfo`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *VcenterContentRegistriesHarborInfo) GetStorageOk() (*[]VcenterContentRegistriesHarborStorageInfo, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *VcenterContentRegistriesHarborInfo) SetStorage(v []VcenterContentRegistriesHarborStorageInfo)`

SetStorage sets Storage field to given value.


### GetHealth

`func (o *VcenterContentRegistriesHarborInfo) GetHealth() VcenterContentRegistriesHealthInfo`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VcenterContentRegistriesHarborInfo) GetHealthOk() (*VcenterContentRegistriesHealthInfo, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VcenterContentRegistriesHarborInfo) SetHealth(v VcenterContentRegistriesHealthInfo)`

SetHealth sets Health field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


