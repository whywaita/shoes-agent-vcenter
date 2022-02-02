# VcenterVMFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vms** | Pointer to **[]string** | Identifiers of virtual machines that can match the filter. If unset or empty, virtual machines with any identifier match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: VirtualMachine. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: VirtualMachine. | [optional] 
**Names** | Pointer to **[]string** | Names that virtual machines must have to match the filter (see VM.Info.name). If unset or empty, virtual machines with any name match the filter. | [optional] 
**Folders** | Pointer to **[]string** | Folders that must contain the virtual machine for the virtual machine to match the filter. If unset or empty, virtual machines in any folder match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Folder. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Folder. | [optional] 
**Datacenters** | Pointer to **[]string** | Datacenters that must contain the virtual machine for the virtual machine to match the filter. If unset or empty, virtual machines in any datacenter match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datacenter. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datacenter. | [optional] 
**Hosts** | Pointer to **[]string** | Hosts that must contain the virtual machine for the virtual machine to match the filter. If unset or empty, virtual machines on any host match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: HostSystem. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: HostSystem. | [optional] 
**Clusters** | Pointer to **[]string** | Clusters that must contain the virtual machine for the virtual machine to match the filter. If unset or empty, virtual machines in any cluster match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ClusterComputeResource. | [optional] 
**ResourcePools** | Pointer to **[]string** | Resource pools that must contain the virtual machine for the virtual machine to match the filter. If unset or empty, virtual machines in any resource pool match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ResourcePool. | [optional] 
**PowerStates** | Pointer to [**[]VcenterVmPowerState**](VcenterVmPowerState.md) | Power states that a virtual machine must be in to match the filter (see Power.Info.state. If unset or empty, virtual machines in any power state match the filter. | [optional] 

## Methods

### NewVcenterVMFilterSpec

`func NewVcenterVMFilterSpec() *VcenterVMFilterSpec`

NewVcenterVMFilterSpec instantiates a new VcenterVMFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMFilterSpecWithDefaults

`func NewVcenterVMFilterSpecWithDefaults() *VcenterVMFilterSpec`

NewVcenterVMFilterSpecWithDefaults instantiates a new VcenterVMFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVms

`func (o *VcenterVMFilterSpec) GetVms() []string`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *VcenterVMFilterSpec) GetVmsOk() (*[]string, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *VcenterVMFilterSpec) SetVms(v []string)`

SetVms sets Vms field to given value.

### HasVms

`func (o *VcenterVMFilterSpec) HasVms() bool`

HasVms returns a boolean if a field has been set.

### GetNames

`func (o *VcenterVMFilterSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *VcenterVMFilterSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *VcenterVMFilterSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *VcenterVMFilterSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetFolders

`func (o *VcenterVMFilterSpec) GetFolders() []string`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *VcenterVMFilterSpec) GetFoldersOk() (*[]string, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *VcenterVMFilterSpec) SetFolders(v []string)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *VcenterVMFilterSpec) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetDatacenters

`func (o *VcenterVMFilterSpec) GetDatacenters() []string`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *VcenterVMFilterSpec) GetDatacentersOk() (*[]string, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *VcenterVMFilterSpec) SetDatacenters(v []string)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *VcenterVMFilterSpec) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetHosts

`func (o *VcenterVMFilterSpec) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VcenterVMFilterSpec) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VcenterVMFilterSpec) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *VcenterVMFilterSpec) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetClusters

`func (o *VcenterVMFilterSpec) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *VcenterVMFilterSpec) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *VcenterVMFilterSpec) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *VcenterVMFilterSpec) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetResourcePools

`func (o *VcenterVMFilterSpec) GetResourcePools() []string`

GetResourcePools returns the ResourcePools field if non-nil, zero value otherwise.

### GetResourcePoolsOk

`func (o *VcenterVMFilterSpec) GetResourcePoolsOk() (*[]string, bool)`

GetResourcePoolsOk returns a tuple with the ResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePools

`func (o *VcenterVMFilterSpec) SetResourcePools(v []string)`

SetResourcePools sets ResourcePools field to given value.

### HasResourcePools

`func (o *VcenterVMFilterSpec) HasResourcePools() bool`

HasResourcePools returns a boolean if a field has been set.

### GetPowerStates

`func (o *VcenterVMFilterSpec) GetPowerStates() []VcenterVmPowerState`

GetPowerStates returns the PowerStates field if non-nil, zero value otherwise.

### GetPowerStatesOk

`func (o *VcenterVMFilterSpec) GetPowerStatesOk() (*[]VcenterVmPowerState, bool)`

GetPowerStatesOk returns a tuple with the PowerStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStates

`func (o *VcenterVMFilterSpec) SetPowerStates(v []VcenterVmPowerState)`

SetPowerStates sets PowerStates field to given value.

### HasPowerStates

`func (o *VcenterVMFilterSpec) HasPowerStates() bool`

HasPowerStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


