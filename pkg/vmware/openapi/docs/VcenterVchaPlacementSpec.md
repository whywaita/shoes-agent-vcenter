# VcenterVchaPlacementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the VCHA node to be used for the virtual machine name. | 
**Folder** | **string** | The identifier of the folder to deploy the VCHA node to. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder:VCenter. | 
**Host** | Pointer to **string** | The identifier of the host to deploy the VCHA node to. If unset, see vim.vm.RelocateSpec.host. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: HostSystem:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: HostSystem:VCenter. | [optional] 
**ResourcePool** | Pointer to **string** | The identifier of the resource pool to deploy the VCHA node to. If unset, then the active node&#39;s resource pool will be used. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ResourcePool:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ResourcePool:VCenter. | [optional] 
**HaNetworkType** | Pointer to [**VcenterVchaNetworkType**](VcenterVchaNetworkType.md) |  | [optional] 
**HaNetwork** | Pointer to **string** | The identifier of the Network object used for the HA network.  If the PlacementSpec.ha-network field is set, then the {#link #haNetworkType} field must be set.  If the PlacementSpec.ha-network field is unset, then the PlacementSpec.ha-network-type field is ignored. If unset and the PlacementSpec.ha-network-type field is unset, then the same network present on the Active node virtual machine is used to deploy the virtual machine with an assumption that the network is present on the destination.  When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network:VCenter. | [optional] 
**ManagementNetworkType** | Pointer to [**VcenterVchaNetworkType**](VcenterVchaNetworkType.md) |  | [optional] 
**ManagementNetwork** | Pointer to **string** | The identifier of the Network object used for the Management network. If the PlacementSpec.management-network field is set, then the PlacementSpec.management-network-type field must be set.  If the PlacementSpec.management-network field is unset, then the PlacementSpec.management-network-type field is ignored. If unset and the PlacementSpec.management-network-type field is unset, then the same network present on the Active node virtual machine is used to deploy the virtual machine with an assumption that the network is present on the destination.  When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network:VCenter. | [optional] 
**Storage** | Pointer to [**VcenterVchaDiskSpec**](VcenterVchaDiskSpec.md) |  | [optional] 

## Methods

### NewVcenterVchaPlacementSpec

`func NewVcenterVchaPlacementSpec(name string, folder string, ) *VcenterVchaPlacementSpec`

NewVcenterVchaPlacementSpec instantiates a new VcenterVchaPlacementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaPlacementSpecWithDefaults

`func NewVcenterVchaPlacementSpecWithDefaults() *VcenterVchaPlacementSpec`

NewVcenterVchaPlacementSpecWithDefaults instantiates a new VcenterVchaPlacementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterVchaPlacementSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVchaPlacementSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVchaPlacementSpec) SetName(v string)`

SetName sets Name field to given value.


### GetFolder

`func (o *VcenterVchaPlacementSpec) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VcenterVchaPlacementSpec) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VcenterVchaPlacementSpec) SetFolder(v string)`

SetFolder sets Folder field to given value.


### GetHost

`func (o *VcenterVchaPlacementSpec) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VcenterVchaPlacementSpec) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VcenterVchaPlacementSpec) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *VcenterVchaPlacementSpec) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetResourcePool

`func (o *VcenterVchaPlacementSpec) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *VcenterVchaPlacementSpec) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *VcenterVchaPlacementSpec) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *VcenterVchaPlacementSpec) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetHaNetworkType

`func (o *VcenterVchaPlacementSpec) GetHaNetworkType() VcenterVchaNetworkType`

GetHaNetworkType returns the HaNetworkType field if non-nil, zero value otherwise.

### GetHaNetworkTypeOk

`func (o *VcenterVchaPlacementSpec) GetHaNetworkTypeOk() (*VcenterVchaNetworkType, bool)`

GetHaNetworkTypeOk returns a tuple with the HaNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaNetworkType

`func (o *VcenterVchaPlacementSpec) SetHaNetworkType(v VcenterVchaNetworkType)`

SetHaNetworkType sets HaNetworkType field to given value.

### HasHaNetworkType

`func (o *VcenterVchaPlacementSpec) HasHaNetworkType() bool`

HasHaNetworkType returns a boolean if a field has been set.

### GetHaNetwork

`func (o *VcenterVchaPlacementSpec) GetHaNetwork() string`

GetHaNetwork returns the HaNetwork field if non-nil, zero value otherwise.

### GetHaNetworkOk

`func (o *VcenterVchaPlacementSpec) GetHaNetworkOk() (*string, bool)`

GetHaNetworkOk returns a tuple with the HaNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaNetwork

`func (o *VcenterVchaPlacementSpec) SetHaNetwork(v string)`

SetHaNetwork sets HaNetwork field to given value.

### HasHaNetwork

`func (o *VcenterVchaPlacementSpec) HasHaNetwork() bool`

HasHaNetwork returns a boolean if a field has been set.

### GetManagementNetworkType

`func (o *VcenterVchaPlacementSpec) GetManagementNetworkType() VcenterVchaNetworkType`

GetManagementNetworkType returns the ManagementNetworkType field if non-nil, zero value otherwise.

### GetManagementNetworkTypeOk

`func (o *VcenterVchaPlacementSpec) GetManagementNetworkTypeOk() (*VcenterVchaNetworkType, bool)`

GetManagementNetworkTypeOk returns a tuple with the ManagementNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNetworkType

`func (o *VcenterVchaPlacementSpec) SetManagementNetworkType(v VcenterVchaNetworkType)`

SetManagementNetworkType sets ManagementNetworkType field to given value.

### HasManagementNetworkType

`func (o *VcenterVchaPlacementSpec) HasManagementNetworkType() bool`

HasManagementNetworkType returns a boolean if a field has been set.

### GetManagementNetwork

`func (o *VcenterVchaPlacementSpec) GetManagementNetwork() string`

GetManagementNetwork returns the ManagementNetwork field if non-nil, zero value otherwise.

### GetManagementNetworkOk

`func (o *VcenterVchaPlacementSpec) GetManagementNetworkOk() (*string, bool)`

GetManagementNetworkOk returns a tuple with the ManagementNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNetwork

`func (o *VcenterVchaPlacementSpec) SetManagementNetwork(v string)`

SetManagementNetwork sets ManagementNetwork field to given value.

### HasManagementNetwork

`func (o *VcenterVchaPlacementSpec) HasManagementNetwork() bool`

HasManagementNetwork returns a boolean if a field has been set.

### GetStorage

`func (o *VcenterVchaPlacementSpec) GetStorage() VcenterVchaDiskSpec`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *VcenterVchaPlacementSpec) GetStorageOk() (*VcenterVchaDiskSpec, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *VcenterVchaPlacementSpec) SetStorage(v VcenterVchaDiskSpec)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *VcenterVchaPlacementSpec) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


