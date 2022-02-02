# VcenterContentRegistriesHarborStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** | Identifier of the storage policy. | 
**Limit** | Pointer to **int64** | The maximum amount of storage (in mebibytes) which can be utilized by the registry for this specification. | [optional] 

## Methods

### NewVcenterContentRegistriesHarborStorageSpec

`func NewVcenterContentRegistriesHarborStorageSpec(policy string, ) *VcenterContentRegistriesHarborStorageSpec`

NewVcenterContentRegistriesHarborStorageSpec instantiates a new VcenterContentRegistriesHarborStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterContentRegistriesHarborStorageSpecWithDefaults

`func NewVcenterContentRegistriesHarborStorageSpecWithDefaults() *VcenterContentRegistriesHarborStorageSpec`

NewVcenterContentRegistriesHarborStorageSpecWithDefaults instantiates a new VcenterContentRegistriesHarborStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *VcenterContentRegistriesHarborStorageSpec) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *VcenterContentRegistriesHarborStorageSpec) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *VcenterContentRegistriesHarborStorageSpec) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetLimit

`func (o *VcenterContentRegistriesHarborStorageSpec) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VcenterContentRegistriesHarborStorageSpec) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VcenterContentRegistriesHarborStorageSpec) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *VcenterContentRegistriesHarborStorageSpec) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


