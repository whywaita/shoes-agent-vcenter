# VcenterVchaDiskSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to **string** | The identifier of the datastore to put all the virtual disks on. This field needs to be set. If unset, then see vim.vm.RelocateSpec.datastore. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datastore:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datastore:VCenter. | [optional] 

## Methods

### NewVcenterVchaDiskSpec

`func NewVcenterVchaDiskSpec() *VcenterVchaDiskSpec`

NewVcenterVchaDiskSpec instantiates a new VcenterVchaDiskSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaDiskSpecWithDefaults

`func NewVcenterVchaDiskSpecWithDefaults() *VcenterVchaDiskSpec`

NewVcenterVchaDiskSpecWithDefaults instantiates a new VcenterVchaDiskSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *VcenterVchaDiskSpec) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVchaDiskSpec) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVchaDiskSpec) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VcenterVchaDiskSpec) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


