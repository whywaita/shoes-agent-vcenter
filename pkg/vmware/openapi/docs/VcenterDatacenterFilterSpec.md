# VcenterDatacenterFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datacenters** | Pointer to **[]string** | Identifiers of datacenters that can match the filter. If unset or empty, datacenters with any identifier match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datacenter. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datacenter. | [optional] 
**Names** | Pointer to **[]string** | Names that datacenters must have to match the filter (see Datacenter.Info.name). If unset or empty, datacenters with any name match the filter. | [optional] 
**Folders** | Pointer to **[]string** | Folders that must contain the datacenters for the datacenter to match the filter. If unset or empty, datacenters in any folder match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Folder. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Folder. | [optional] 

## Methods

### NewVcenterDatacenterFilterSpec

`func NewVcenterDatacenterFilterSpec() *VcenterDatacenterFilterSpec`

NewVcenterDatacenterFilterSpec instantiates a new VcenterDatacenterFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDatacenterFilterSpecWithDefaults

`func NewVcenterDatacenterFilterSpecWithDefaults() *VcenterDatacenterFilterSpec`

NewVcenterDatacenterFilterSpecWithDefaults instantiates a new VcenterDatacenterFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenters

`func (o *VcenterDatacenterFilterSpec) GetDatacenters() []string`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *VcenterDatacenterFilterSpec) GetDatacentersOk() (*[]string, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *VcenterDatacenterFilterSpec) SetDatacenters(v []string)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *VcenterDatacenterFilterSpec) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetNames

`func (o *VcenterDatacenterFilterSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *VcenterDatacenterFilterSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *VcenterDatacenterFilterSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *VcenterDatacenterFilterSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetFolders

`func (o *VcenterDatacenterFilterSpec) GetFolders() []string`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *VcenterDatacenterFilterSpec) GetFoldersOk() (*[]string, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *VcenterDatacenterFilterSpec) SetFolders(v []string)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *VcenterDatacenterFilterSpec) HasFolders() bool`

HasFolders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


