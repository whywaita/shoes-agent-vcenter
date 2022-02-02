# VcenterVMDiskCloneSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to **string** | Destination datastore to clone disk. This field is currently required. In the future, if this field is unset disk will be copied to the datastore specified in the VM.ClonePlacementSpec.datastore field of VM.CloneSpec.placement. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datastore. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datastore. | [optional] 

## Methods

### NewVcenterVMDiskCloneSpec

`func NewVcenterVMDiskCloneSpec() *VcenterVMDiskCloneSpec`

NewVcenterVMDiskCloneSpec instantiates a new VcenterVMDiskCloneSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMDiskCloneSpecWithDefaults

`func NewVcenterVMDiskCloneSpecWithDefaults() *VcenterVMDiskCloneSpec`

NewVcenterVMDiskCloneSpecWithDefaults instantiates a new VcenterVMDiskCloneSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *VcenterVMDiskCloneSpec) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVMDiskCloneSpec) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVMDiskCloneSpec) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VcenterVMDiskCloneSpec) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


