# VcenterVMRegisterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to **string** | Identifier of the datastore on which the virtual machine&#39;s configuration state is stored. If unset, VM.RegisterSpec.path must also be unset and VM.RegisterSpec.datastore-path must be set. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datastore. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datastore. | [optional] 
**Path** | Pointer to **string** | Path to the virtual machine&#39;s configuration file on the datastore corresponding to {@link #datastore). If unset, VM.RegisterSpec.datastore must also be unset and VM.RegisterSpec.datastore-path must be set. | [optional] 
**DatastorePath** | Pointer to **string** | Datastore path for the virtual machine&#39;s configuration file in the format \&quot;[datastore name] path\&quot;. For example \&quot;[storage1] Test-VM/Test-VM.vmx\&quot;. If unset, both VM.RegisterSpec.datastore and VM.RegisterSpec.path must be set. | [optional] 
**Name** | Pointer to **string** | Virtual machine name. If unset, the display name from the virtual machine&#39;s configuration file will be used. | [optional] 
**Placement** | Pointer to [**VcenterVMRegisterPlacementSpec**](VcenterVMRegisterPlacementSpec.md) |  | [optional] 

## Methods

### NewVcenterVMRegisterSpec

`func NewVcenterVMRegisterSpec() *VcenterVMRegisterSpec`

NewVcenterVMRegisterSpec instantiates a new VcenterVMRegisterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMRegisterSpecWithDefaults

`func NewVcenterVMRegisterSpecWithDefaults() *VcenterVMRegisterSpec`

NewVcenterVMRegisterSpecWithDefaults instantiates a new VcenterVMRegisterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *VcenterVMRegisterSpec) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVMRegisterSpec) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVMRegisterSpec) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VcenterVMRegisterSpec) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetPath

`func (o *VcenterVMRegisterSpec) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VcenterVMRegisterSpec) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VcenterVMRegisterSpec) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *VcenterVMRegisterSpec) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetDatastorePath

`func (o *VcenterVMRegisterSpec) GetDatastorePath() string`

GetDatastorePath returns the DatastorePath field if non-nil, zero value otherwise.

### GetDatastorePathOk

`func (o *VcenterVMRegisterSpec) GetDatastorePathOk() (*string, bool)`

GetDatastorePathOk returns a tuple with the DatastorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastorePath

`func (o *VcenterVMRegisterSpec) SetDatastorePath(v string)`

SetDatastorePath sets DatastorePath field to given value.

### HasDatastorePath

`func (o *VcenterVMRegisterSpec) HasDatastorePath() bool`

HasDatastorePath returns a boolean if a field has been set.

### GetName

`func (o *VcenterVMRegisterSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVMRegisterSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVMRegisterSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterVMRegisterSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *VcenterVMRegisterSpec) GetPlacement() VcenterVMRegisterPlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVMRegisterSpec) GetPlacementOk() (*VcenterVMRegisterPlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVMRegisterSpec) SetPlacement(v VcenterVMRegisterPlacementSpec)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VcenterVMRegisterSpec) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


