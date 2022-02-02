# VcenterVMDiskRelocateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to **string** | Destination datastore to relocate disk. This field is currently required. In the future, if this field is unset, disk will use the datastore specified in VM.RelocatePlacementSpec.datastore field of VM.RelocateSpec.placement. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datastore. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datastore. | [optional] 

## Methods

### NewVcenterVMDiskRelocateSpec

`func NewVcenterVMDiskRelocateSpec() *VcenterVMDiskRelocateSpec`

NewVcenterVMDiskRelocateSpec instantiates a new VcenterVMDiskRelocateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMDiskRelocateSpecWithDefaults

`func NewVcenterVMDiskRelocateSpecWithDefaults() *VcenterVMDiskRelocateSpec`

NewVcenterVMDiskRelocateSpecWithDefaults instantiates a new VcenterVMDiskRelocateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *VcenterVMDiskRelocateSpec) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVMDiskRelocateSpec) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVMDiskRelocateSpec) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VcenterVMDiskRelocateSpec) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


