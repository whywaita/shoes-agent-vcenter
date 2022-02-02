# VcenterDatastoreFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastores** | Pointer to **[]string** | Identifiers of datastores that can match the filter. If unset or empty, datastores with any identifier match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datastore. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datastore. | [optional] 
**Names** | Pointer to **[]string** | Names that datastores must have to match the filter (see Datastore.Info.name). If unset or empty, datastores with any name match the filter. | [optional] 
**Types** | Pointer to [**[]VcenterDatastoreType**](VcenterDatastoreType.md) | Types that datastores must have to match the filter (see Datastore.Summary.type). If unset or empty, datastores with any type match the filter. | [optional] 
**Folders** | Pointer to **[]string** | Folders that must contain the datastore for the datastore to match the filter. If unset or empty, datastores in any folder match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Folder. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Folder. | [optional] 
**Datacenters** | Pointer to **[]string** | Datacenters that must contain the datastore for the datastore to match the filter. If unset or empty, datastores in any datacenter match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datacenter. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datacenter. | [optional] 

## Methods

### NewVcenterDatastoreFilterSpec

`func NewVcenterDatastoreFilterSpec() *VcenterDatastoreFilterSpec`

NewVcenterDatastoreFilterSpec instantiates a new VcenterDatastoreFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDatastoreFilterSpecWithDefaults

`func NewVcenterDatastoreFilterSpecWithDefaults() *VcenterDatastoreFilterSpec`

NewVcenterDatastoreFilterSpecWithDefaults instantiates a new VcenterDatastoreFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastores

`func (o *VcenterDatastoreFilterSpec) GetDatastores() []string`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *VcenterDatastoreFilterSpec) GetDatastoresOk() (*[]string, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *VcenterDatastoreFilterSpec) SetDatastores(v []string)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *VcenterDatastoreFilterSpec) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetNames

`func (o *VcenterDatastoreFilterSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *VcenterDatastoreFilterSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *VcenterDatastoreFilterSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *VcenterDatastoreFilterSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetTypes

`func (o *VcenterDatastoreFilterSpec) GetTypes() []VcenterDatastoreType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *VcenterDatastoreFilterSpec) GetTypesOk() (*[]VcenterDatastoreType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *VcenterDatastoreFilterSpec) SetTypes(v []VcenterDatastoreType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *VcenterDatastoreFilterSpec) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetFolders

`func (o *VcenterDatastoreFilterSpec) GetFolders() []string`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *VcenterDatastoreFilterSpec) GetFoldersOk() (*[]string, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *VcenterDatastoreFilterSpec) SetFolders(v []string)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *VcenterDatastoreFilterSpec) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetDatacenters

`func (o *VcenterDatastoreFilterSpec) GetDatacenters() []string`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *VcenterDatastoreFilterSpec) GetDatacentersOk() (*[]string, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *VcenterDatastoreFilterSpec) SetDatacenters(v []string)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *VcenterDatastoreFilterSpec) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


