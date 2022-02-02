# VcenterStoragePoliciesCheckCompatibilityStoragePolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastores** | Pointer to **[]string** | Datastores used to check compatibility against a storage policy. The number of datastores is limited to 1024. The parameter must contain identifiers for the resource type: Datastore. | [optional] 

## Methods

### NewVcenterStoragePoliciesCheckCompatibilityStoragePolicies

`func NewVcenterStoragePoliciesCheckCompatibilityStoragePolicies() *VcenterStoragePoliciesCheckCompatibilityStoragePolicies`

NewVcenterStoragePoliciesCheckCompatibilityStoragePolicies instantiates a new VcenterStoragePoliciesCheckCompatibilityStoragePolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterStoragePoliciesCheckCompatibilityStoragePoliciesWithDefaults

`func NewVcenterStoragePoliciesCheckCompatibilityStoragePoliciesWithDefaults() *VcenterStoragePoliciesCheckCompatibilityStoragePolicies`

NewVcenterStoragePoliciesCheckCompatibilityStoragePoliciesWithDefaults instantiates a new VcenterStoragePoliciesCheckCompatibilityStoragePolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastores

`func (o *VcenterStoragePoliciesCheckCompatibilityStoragePolicies) GetDatastores() []string`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *VcenterStoragePoliciesCheckCompatibilityStoragePolicies) GetDatastoresOk() (*[]string, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *VcenterStoragePoliciesCheckCompatibilityStoragePolicies) SetDatastores(v []string)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *VcenterStoragePoliciesCheckCompatibilityStoragePolicies) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


