# VcenterFolderFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folders** | Pointer to **[]string** | Identifiers of folders that can match the filter. If unset or empty, folders with any identifier match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Folder. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Folder. | [optional] 
**Names** | Pointer to **[]string** | Names that folders must have to match the filter (see Folder.Summary.name). If unset or empty, folders with any name match the filter. | [optional] 
**Type** | Pointer to [**VcenterFolderType**](VcenterFolderType.md) |  | [optional] 
**ParentFolders** | Pointer to **[]string** | Folders that must contain the folder for the folder to match the filter. If unset or empty, folder in any folder match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Folder. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Folder. | [optional] 
**Datacenters** | Pointer to **[]string** | Datacenters that must contain the folder for the folder to match the filter. If unset or empty, folder in any datacenter match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datacenter. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datacenter. | [optional] 

## Methods

### NewVcenterFolderFilterSpec

`func NewVcenterFolderFilterSpec() *VcenterFolderFilterSpec`

NewVcenterFolderFilterSpec instantiates a new VcenterFolderFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterFolderFilterSpecWithDefaults

`func NewVcenterFolderFilterSpecWithDefaults() *VcenterFolderFilterSpec`

NewVcenterFolderFilterSpecWithDefaults instantiates a new VcenterFolderFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolders

`func (o *VcenterFolderFilterSpec) GetFolders() []string`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *VcenterFolderFilterSpec) GetFoldersOk() (*[]string, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *VcenterFolderFilterSpec) SetFolders(v []string)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *VcenterFolderFilterSpec) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetNames

`func (o *VcenterFolderFilterSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *VcenterFolderFilterSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *VcenterFolderFilterSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *VcenterFolderFilterSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetType

`func (o *VcenterFolderFilterSpec) GetType() VcenterFolderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterFolderFilterSpec) GetTypeOk() (*VcenterFolderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterFolderFilterSpec) SetType(v VcenterFolderType)`

SetType sets Type field to given value.

### HasType

`func (o *VcenterFolderFilterSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParentFolders

`func (o *VcenterFolderFilterSpec) GetParentFolders() []string`

GetParentFolders returns the ParentFolders field if non-nil, zero value otherwise.

### GetParentFoldersOk

`func (o *VcenterFolderFilterSpec) GetParentFoldersOk() (*[]string, bool)`

GetParentFoldersOk returns a tuple with the ParentFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolders

`func (o *VcenterFolderFilterSpec) SetParentFolders(v []string)`

SetParentFolders sets ParentFolders field to given value.

### HasParentFolders

`func (o *VcenterFolderFilterSpec) HasParentFolders() bool`

HasParentFolders returns a boolean if a field has been set.

### GetDatacenters

`func (o *VcenterFolderFilterSpec) GetDatacenters() []string`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *VcenterFolderFilterSpec) GetDatacentersOk() (*[]string, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *VcenterFolderFilterSpec) SetDatacenters(v []string)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *VcenterFolderFilterSpec) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


