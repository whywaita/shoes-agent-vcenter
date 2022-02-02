# VcenterVmHardwareDiskStoragePolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** | Identifier of the storage policy which should be associated with the VMDK file. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.StoragePolicy. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.StoragePolicy. | 

## Methods

### NewVcenterVmHardwareDiskStoragePolicySpec

`func NewVcenterVmHardwareDiskStoragePolicySpec(policy string, ) *VcenterVmHardwareDiskStoragePolicySpec`

NewVcenterVmHardwareDiskStoragePolicySpec instantiates a new VcenterVmHardwareDiskStoragePolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareDiskStoragePolicySpecWithDefaults

`func NewVcenterVmHardwareDiskStoragePolicySpecWithDefaults() *VcenterVmHardwareDiskStoragePolicySpec`

NewVcenterVmHardwareDiskStoragePolicySpecWithDefaults instantiates a new VcenterVmHardwareDiskStoragePolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *VcenterVmHardwareDiskStoragePolicySpec) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *VcenterVmHardwareDiskStoragePolicySpec) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *VcenterVmHardwareDiskStoragePolicySpec) SetPolicy(v string)`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


