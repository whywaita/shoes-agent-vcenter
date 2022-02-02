# VcenterNetworkFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Networks** | Pointer to **[]string** | Identifiers of networks that can match the filter. If unset or empty, networks with any identifier match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Network. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Network. | [optional] 
**Names** | Pointer to **[]string** | Names that networks must have to match the filter (see Network.Summary.name). If unset or empty, networks with any name match the filter. | [optional] 
**Types** | Pointer to [**[]VcenterNetworkType**](VcenterNetworkType.md) | Types that networks must have to match the filter (see Network.Summary.type). If unset, networks with any type match the filter. | [optional] 
**Folders** | Pointer to **[]string** | Folders that must contain the network for the network to match the filter. If unset or empty, networks in any folder match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Folder. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Folder. | [optional] 
**Datacenters** | Pointer to **[]string** | Datacenters that must contain the network for the network to match the filter. If unset or empty, networks in any datacenter match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datacenter. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datacenter. | [optional] 

## Methods

### NewVcenterNetworkFilterSpec

`func NewVcenterNetworkFilterSpec() *VcenterNetworkFilterSpec`

NewVcenterNetworkFilterSpec instantiates a new VcenterNetworkFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNetworkFilterSpecWithDefaults

`func NewVcenterNetworkFilterSpecWithDefaults() *VcenterNetworkFilterSpec`

NewVcenterNetworkFilterSpecWithDefaults instantiates a new VcenterNetworkFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworks

`func (o *VcenterNetworkFilterSpec) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *VcenterNetworkFilterSpec) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *VcenterNetworkFilterSpec) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *VcenterNetworkFilterSpec) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNames

`func (o *VcenterNetworkFilterSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *VcenterNetworkFilterSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *VcenterNetworkFilterSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *VcenterNetworkFilterSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetTypes

`func (o *VcenterNetworkFilterSpec) GetTypes() []VcenterNetworkType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *VcenterNetworkFilterSpec) GetTypesOk() (*[]VcenterNetworkType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *VcenterNetworkFilterSpec) SetTypes(v []VcenterNetworkType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *VcenterNetworkFilterSpec) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetFolders

`func (o *VcenterNetworkFilterSpec) GetFolders() []string`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *VcenterNetworkFilterSpec) GetFoldersOk() (*[]string, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *VcenterNetworkFilterSpec) SetFolders(v []string)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *VcenterNetworkFilterSpec) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetDatacenters

`func (o *VcenterNetworkFilterSpec) GetDatacenters() []string`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *VcenterNetworkFilterSpec) GetDatacentersOk() (*[]string, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *VcenterNetworkFilterSpec) SetDatacenters(v []string)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *VcenterNetworkFilterSpec) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


