# VcenterNamespaceManagementClustersImageStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoragePolicy** | **string** | Identifier of the storage policy. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: SpsStorageProfile. When operations return a value of this structure as a result, the field will be an identifier for the resource type: SpsStorageProfile. | 

## Methods

### NewVcenterNamespaceManagementClustersImageStorageSpec

`func NewVcenterNamespaceManagementClustersImageStorageSpec(storagePolicy string, ) *VcenterNamespaceManagementClustersImageStorageSpec`

NewVcenterNamespaceManagementClustersImageStorageSpec instantiates a new VcenterNamespaceManagementClustersImageStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersImageStorageSpecWithDefaults

`func NewVcenterNamespaceManagementClustersImageStorageSpecWithDefaults() *VcenterNamespaceManagementClustersImageStorageSpec`

NewVcenterNamespaceManagementClustersImageStorageSpecWithDefaults instantiates a new VcenterNamespaceManagementClustersImageStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoragePolicy

`func (o *VcenterNamespaceManagementClustersImageStorageSpec) GetStoragePolicy() string`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *VcenterNamespaceManagementClustersImageStorageSpec) GetStoragePolicyOk() (*string, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *VcenterNamespaceManagementClustersImageStorageSpec) SetStoragePolicy(v string)`

SetStoragePolicy sets StoragePolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


