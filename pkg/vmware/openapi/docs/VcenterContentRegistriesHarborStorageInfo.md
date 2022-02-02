# VcenterContentRegistriesHarborStorageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** | Identifier of the storage policy. | 
**Capacity** | **int64** | Total capacity for the registry storage (in mebibytes). This is the storage limit set on the Harbor registry. If a storage limit was not set on the registry, the default registry capacity - 204800 mebibytes is used. | 
**Used** | **int64** | Overall storage used by the registry (in mebibytes). This is the sum of used storage associated with storage policies configured for the registry. | 

## Methods

### NewVcenterContentRegistriesHarborStorageInfo

`func NewVcenterContentRegistriesHarborStorageInfo(policy string, capacity int64, used int64, ) *VcenterContentRegistriesHarborStorageInfo`

NewVcenterContentRegistriesHarborStorageInfo instantiates a new VcenterContentRegistriesHarborStorageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterContentRegistriesHarborStorageInfoWithDefaults

`func NewVcenterContentRegistriesHarborStorageInfoWithDefaults() *VcenterContentRegistriesHarborStorageInfo`

NewVcenterContentRegistriesHarborStorageInfoWithDefaults instantiates a new VcenterContentRegistriesHarborStorageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *VcenterContentRegistriesHarborStorageInfo) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *VcenterContentRegistriesHarborStorageInfo) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *VcenterContentRegistriesHarborStorageInfo) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetCapacity

`func (o *VcenterContentRegistriesHarborStorageInfo) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VcenterContentRegistriesHarborStorageInfo) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VcenterContentRegistriesHarborStorageInfo) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.


### GetUsed

`func (o *VcenterContentRegistriesHarborStorageInfo) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *VcenterContentRegistriesHarborStorageInfo) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *VcenterContentRegistriesHarborStorageInfo) SetUsed(v int64)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


