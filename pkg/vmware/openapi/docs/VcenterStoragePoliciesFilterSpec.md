# VcenterStoragePoliciesFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to **[]string** | Identifiers of storage policies that can match the filter. If unset or empty, storage policies with any identifiers match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.StoragePolicy. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.StoragePolicy. | [optional] 

## Methods

### NewVcenterStoragePoliciesFilterSpec

`func NewVcenterStoragePoliciesFilterSpec() *VcenterStoragePoliciesFilterSpec`

NewVcenterStoragePoliciesFilterSpec instantiates a new VcenterStoragePoliciesFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterStoragePoliciesFilterSpecWithDefaults

`func NewVcenterStoragePoliciesFilterSpecWithDefaults() *VcenterStoragePoliciesFilterSpec`

NewVcenterStoragePoliciesFilterSpecWithDefaults instantiates a new VcenterStoragePoliciesFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *VcenterStoragePoliciesFilterSpec) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *VcenterStoragePoliciesFilterSpec) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *VcenterStoragePoliciesFilterSpec) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *VcenterStoragePoliciesFilterSpec) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


