# VcenterComputePoliciesTagUsageFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to **[]string** | Identifiers that compute policies must have to match the filter. | [optional] 
**Tags** | Pointer to **[]string** | Identifiers that tags must have to match the filter. | [optional] 
**TagTypes** | Pointer to **[]string** | Identifiers that tag types must have to match the filter. | [optional] 

## Methods

### NewVcenterComputePoliciesTagUsageFilterSpec

`func NewVcenterComputePoliciesTagUsageFilterSpec() *VcenterComputePoliciesTagUsageFilterSpec`

NewVcenterComputePoliciesTagUsageFilterSpec instantiates a new VcenterComputePoliciesTagUsageFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterComputePoliciesTagUsageFilterSpecWithDefaults

`func NewVcenterComputePoliciesTagUsageFilterSpecWithDefaults() *VcenterComputePoliciesTagUsageFilterSpec`

NewVcenterComputePoliciesTagUsageFilterSpecWithDefaults instantiates a new VcenterComputePoliciesTagUsageFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *VcenterComputePoliciesTagUsageFilterSpec) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *VcenterComputePoliciesTagUsageFilterSpec) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *VcenterComputePoliciesTagUsageFilterSpec) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *VcenterComputePoliciesTagUsageFilterSpec) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetTags

`func (o *VcenterComputePoliciesTagUsageFilterSpec) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VcenterComputePoliciesTagUsageFilterSpec) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VcenterComputePoliciesTagUsageFilterSpec) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VcenterComputePoliciesTagUsageFilterSpec) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTagTypes

`func (o *VcenterComputePoliciesTagUsageFilterSpec) GetTagTypes() []string`

GetTagTypes returns the TagTypes field if non-nil, zero value otherwise.

### GetTagTypesOk

`func (o *VcenterComputePoliciesTagUsageFilterSpec) GetTagTypesOk() (*[]string, bool)`

GetTagTypesOk returns a tuple with the TagTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTypes

`func (o *VcenterComputePoliciesTagUsageFilterSpec) SetTagTypes(v []string)`

SetTagTypes sets TagTypes field to given value.

### HasTagTypes

`func (o *VcenterComputePoliciesTagUsageFilterSpec) HasTagTypes() bool`

HasTagTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


