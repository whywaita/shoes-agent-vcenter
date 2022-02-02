# VcenterVmStoragePolicyVmHomePolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmStoragePolicyVmHomePolicySpecPolicyType**](VcenterVmStoragePolicyVmHomePolicySpecPolicyType.md) |  | 
**Policy** | Pointer to **string** | Storage Policy identification. This field is optional and it is only relevant when the value of Policy.VmHomePolicySpec.type is USE_SPECIFIED_POLICY. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.StoragePolicy. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.StoragePolicy. | [optional] 

## Methods

### NewVcenterVmStoragePolicyVmHomePolicySpec

`func NewVcenterVmStoragePolicyVmHomePolicySpec(type_ VcenterVmStoragePolicyVmHomePolicySpecPolicyType, ) *VcenterVmStoragePolicyVmHomePolicySpec`

NewVcenterVmStoragePolicyVmHomePolicySpec instantiates a new VcenterVmStoragePolicyVmHomePolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmStoragePolicyVmHomePolicySpecWithDefaults

`func NewVcenterVmStoragePolicyVmHomePolicySpecWithDefaults() *VcenterVmStoragePolicyVmHomePolicySpec`

NewVcenterVmStoragePolicyVmHomePolicySpecWithDefaults instantiates a new VcenterVmStoragePolicyVmHomePolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmStoragePolicyVmHomePolicySpec) GetType() VcenterVmStoragePolicyVmHomePolicySpecPolicyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmStoragePolicyVmHomePolicySpec) GetTypeOk() (*VcenterVmStoragePolicyVmHomePolicySpecPolicyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmStoragePolicyVmHomePolicySpec) SetType(v VcenterVmStoragePolicyVmHomePolicySpecPolicyType)`

SetType sets Type field to given value.


### GetPolicy

`func (o *VcenterVmStoragePolicyVmHomePolicySpec) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *VcenterVmStoragePolicyVmHomePolicySpec) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *VcenterVmStoragePolicyVmHomePolicySpec) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *VcenterVmStoragePolicyVmHomePolicySpec) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


