# VcenterOvfLibraryItemStorageGroupMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterOvfLibraryItemStorageGroupMappingType**](VcenterOvfLibraryItemStorageGroupMappingType.md) |  | 
**DatastoreId** | Pointer to **string** | Target datastore to be used for the storage group. | [optional] 
**StorageProfileId** | Pointer to **string** | Target storage profile to be used for the storage group. | [optional] 
**Provisioning** | Pointer to [**VcenterOvfDiskProvisioningType**](VcenterOvfDiskProvisioningType.md) |  | [optional] 

## Methods

### NewVcenterOvfLibraryItemStorageGroupMapping

`func NewVcenterOvfLibraryItemStorageGroupMapping(type_ VcenterOvfLibraryItemStorageGroupMappingType, ) *VcenterOvfLibraryItemStorageGroupMapping`

NewVcenterOvfLibraryItemStorageGroupMapping instantiates a new VcenterOvfLibraryItemStorageGroupMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfLibraryItemStorageGroupMappingWithDefaults

`func NewVcenterOvfLibraryItemStorageGroupMappingWithDefaults() *VcenterOvfLibraryItemStorageGroupMapping`

NewVcenterOvfLibraryItemStorageGroupMappingWithDefaults instantiates a new VcenterOvfLibraryItemStorageGroupMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterOvfLibraryItemStorageGroupMapping) GetType() VcenterOvfLibraryItemStorageGroupMappingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterOvfLibraryItemStorageGroupMapping) GetTypeOk() (*VcenterOvfLibraryItemStorageGroupMappingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterOvfLibraryItemStorageGroupMapping) SetType(v VcenterOvfLibraryItemStorageGroupMappingType)`

SetType sets Type field to given value.


### GetDatastoreId

`func (o *VcenterOvfLibraryItemStorageGroupMapping) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *VcenterOvfLibraryItemStorageGroupMapping) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *VcenterOvfLibraryItemStorageGroupMapping) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *VcenterOvfLibraryItemStorageGroupMapping) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### GetStorageProfileId

`func (o *VcenterOvfLibraryItemStorageGroupMapping) GetStorageProfileId() string`

GetStorageProfileId returns the StorageProfileId field if non-nil, zero value otherwise.

### GetStorageProfileIdOk

`func (o *VcenterOvfLibraryItemStorageGroupMapping) GetStorageProfileIdOk() (*string, bool)`

GetStorageProfileIdOk returns a tuple with the StorageProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfileId

`func (o *VcenterOvfLibraryItemStorageGroupMapping) SetStorageProfileId(v string)`

SetStorageProfileId sets StorageProfileId field to given value.

### HasStorageProfileId

`func (o *VcenterOvfLibraryItemStorageGroupMapping) HasStorageProfileId() bool`

HasStorageProfileId returns a boolean if a field has been set.

### GetProvisioning

`func (o *VcenterOvfLibraryItemStorageGroupMapping) GetProvisioning() VcenterOvfDiskProvisioningType`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *VcenterOvfLibraryItemStorageGroupMapping) GetProvisioningOk() (*VcenterOvfDiskProvisioningType, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *VcenterOvfLibraryItemStorageGroupMapping) SetProvisioning(v VcenterOvfDiskProvisioningType)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *VcenterOvfLibraryItemStorageGroupMapping) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


