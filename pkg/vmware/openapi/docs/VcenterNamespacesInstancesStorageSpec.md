# VcenterNamespacesInstancesStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** | ID of the storage policy. A Kubernetes storage class is created for this storage policy if it does not exist already. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: SpsStorageProfile. When operations return a value of this structure as a result, the field will be an identifier for the resource type: SpsStorageProfile. | 
**Limit** | Pointer to **int64** | The maximum amount of storage (in mebibytes) which can be utilized by the namespace for this specification. If unset, no limits are placed. | [optional] 

## Methods

### NewVcenterNamespacesInstancesStorageSpec

`func NewVcenterNamespacesInstancesStorageSpec(policy string, ) *VcenterNamespacesInstancesStorageSpec`

NewVcenterNamespacesInstancesStorageSpec instantiates a new VcenterNamespacesInstancesStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesInstancesStorageSpecWithDefaults

`func NewVcenterNamespacesInstancesStorageSpecWithDefaults() *VcenterNamespacesInstancesStorageSpec`

NewVcenterNamespacesInstancesStorageSpecWithDefaults instantiates a new VcenterNamespacesInstancesStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *VcenterNamespacesInstancesStorageSpec) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *VcenterNamespacesInstancesStorageSpec) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *VcenterNamespacesInstancesStorageSpec) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetLimit

`func (o *VcenterNamespacesInstancesStorageSpec) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VcenterNamespacesInstancesStorageSpec) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VcenterNamespacesInstancesStorageSpec) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *VcenterNamespacesInstancesStorageSpec) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


