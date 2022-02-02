# VcenterVchaPlacementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementVcenterName** | **string** | The hostname of the vCenter server that is managing the VCHA node. | 
**ManagementVcenterServerGuid** | Pointer to **string** | The unique identifier of the vCenter server that is managing the VCHA node. This field is optional because it was added in a newer version than its parent node. | [optional] 
**VmName** | **string** | The virtual machine name of the VCHA node. | 
**Datacenter** | **string** | The identifier of the datacenter of the VCHA node. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datacenter:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datacenter:VCenter. | 
**DatacenterName** | **string** | The name of the datacenter of the VCHA node. | 
**Host** | **string** | The identifier of the host of the VCHA node. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: HostSystem:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: HostSystem:VCenter. | 
**HostName** | **string** | The name of the host of the VCHA node. | 
**Cluster** | Pointer to **string** | The identifier of the cluster of which PlacementInfo.host is member. If unset, PlacementInfo.host is a standalone host. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource:VCenter. | [optional] 
**ClusterName** | Pointer to **string** | The name of the cluster of which PlacementInfo.host is member. If unset, PlacementInfo.host is a standalone host. | [optional] 
**HaNetwork** | Pointer to **string** | The identifier of the Network object used for the HA network. If unset, the information is currently unavailable or the haNetwork is not configured. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network:VCenter. | [optional] 
**HaNetworkName** | Pointer to **string** | The name of the Network object used for the HA network. If unset, the information is currently unavailable or the haNetwork is not configured. | [optional] 
**HaNetworkType** | Pointer to [**VcenterVchaNetworkType**](VcenterVchaNetworkType.md) |  | [optional] 
**ManagementNetwork** | **string** | The identifier of the Network object used for the Management network. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network:VCenter. | 
**ManagementNetworkName** | **string** | The name of the Network object used for the Management network. | 
**ManagementNetworkType** | [**VcenterVchaNetworkType**](VcenterVchaNetworkType.md) |  | 
**Storage** | [**VcenterVchaDiskInfo**](VcenterVchaDiskInfo.md) |  | 
**BiosUuid** | Pointer to **string** | BIOS UUID for the node. If unset, the information is currently unavailable. | [optional] 

## Methods

### NewVcenterVchaPlacementInfo

`func NewVcenterVchaPlacementInfo(managementVcenterName string, vmName string, datacenter string, datacenterName string, host string, hostName string, managementNetwork string, managementNetworkName string, managementNetworkType VcenterVchaNetworkType, storage VcenterVchaDiskInfo, ) *VcenterVchaPlacementInfo`

NewVcenterVchaPlacementInfo instantiates a new VcenterVchaPlacementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaPlacementInfoWithDefaults

`func NewVcenterVchaPlacementInfoWithDefaults() *VcenterVchaPlacementInfo`

NewVcenterVchaPlacementInfoWithDefaults instantiates a new VcenterVchaPlacementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementVcenterName

`func (o *VcenterVchaPlacementInfo) GetManagementVcenterName() string`

GetManagementVcenterName returns the ManagementVcenterName field if non-nil, zero value otherwise.

### GetManagementVcenterNameOk

`func (o *VcenterVchaPlacementInfo) GetManagementVcenterNameOk() (*string, bool)`

GetManagementVcenterNameOk returns a tuple with the ManagementVcenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementVcenterName

`func (o *VcenterVchaPlacementInfo) SetManagementVcenterName(v string)`

SetManagementVcenterName sets ManagementVcenterName field to given value.


### GetManagementVcenterServerGuid

`func (o *VcenterVchaPlacementInfo) GetManagementVcenterServerGuid() string`

GetManagementVcenterServerGuid returns the ManagementVcenterServerGuid field if non-nil, zero value otherwise.

### GetManagementVcenterServerGuidOk

`func (o *VcenterVchaPlacementInfo) GetManagementVcenterServerGuidOk() (*string, bool)`

GetManagementVcenterServerGuidOk returns a tuple with the ManagementVcenterServerGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementVcenterServerGuid

`func (o *VcenterVchaPlacementInfo) SetManagementVcenterServerGuid(v string)`

SetManagementVcenterServerGuid sets ManagementVcenterServerGuid field to given value.

### HasManagementVcenterServerGuid

`func (o *VcenterVchaPlacementInfo) HasManagementVcenterServerGuid() bool`

HasManagementVcenterServerGuid returns a boolean if a field has been set.

### GetVmName

`func (o *VcenterVchaPlacementInfo) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *VcenterVchaPlacementInfo) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *VcenterVchaPlacementInfo) SetVmName(v string)`

SetVmName sets VmName field to given value.


### GetDatacenter

`func (o *VcenterVchaPlacementInfo) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VcenterVchaPlacementInfo) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VcenterVchaPlacementInfo) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.


### GetDatacenterName

`func (o *VcenterVchaPlacementInfo) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *VcenterVchaPlacementInfo) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *VcenterVchaPlacementInfo) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.


### GetHost

`func (o *VcenterVchaPlacementInfo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VcenterVchaPlacementInfo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VcenterVchaPlacementInfo) SetHost(v string)`

SetHost sets Host field to given value.


### GetHostName

`func (o *VcenterVchaPlacementInfo) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *VcenterVchaPlacementInfo) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *VcenterVchaPlacementInfo) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetCluster

`func (o *VcenterVchaPlacementInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterVchaPlacementInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterVchaPlacementInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VcenterVchaPlacementInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterName

`func (o *VcenterVchaPlacementInfo) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *VcenterVchaPlacementInfo) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *VcenterVchaPlacementInfo) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *VcenterVchaPlacementInfo) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetHaNetwork

`func (o *VcenterVchaPlacementInfo) GetHaNetwork() string`

GetHaNetwork returns the HaNetwork field if non-nil, zero value otherwise.

### GetHaNetworkOk

`func (o *VcenterVchaPlacementInfo) GetHaNetworkOk() (*string, bool)`

GetHaNetworkOk returns a tuple with the HaNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaNetwork

`func (o *VcenterVchaPlacementInfo) SetHaNetwork(v string)`

SetHaNetwork sets HaNetwork field to given value.

### HasHaNetwork

`func (o *VcenterVchaPlacementInfo) HasHaNetwork() bool`

HasHaNetwork returns a boolean if a field has been set.

### GetHaNetworkName

`func (o *VcenterVchaPlacementInfo) GetHaNetworkName() string`

GetHaNetworkName returns the HaNetworkName field if non-nil, zero value otherwise.

### GetHaNetworkNameOk

`func (o *VcenterVchaPlacementInfo) GetHaNetworkNameOk() (*string, bool)`

GetHaNetworkNameOk returns a tuple with the HaNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaNetworkName

`func (o *VcenterVchaPlacementInfo) SetHaNetworkName(v string)`

SetHaNetworkName sets HaNetworkName field to given value.

### HasHaNetworkName

`func (o *VcenterVchaPlacementInfo) HasHaNetworkName() bool`

HasHaNetworkName returns a boolean if a field has been set.

### GetHaNetworkType

`func (o *VcenterVchaPlacementInfo) GetHaNetworkType() VcenterVchaNetworkType`

GetHaNetworkType returns the HaNetworkType field if non-nil, zero value otherwise.

### GetHaNetworkTypeOk

`func (o *VcenterVchaPlacementInfo) GetHaNetworkTypeOk() (*VcenterVchaNetworkType, bool)`

GetHaNetworkTypeOk returns a tuple with the HaNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaNetworkType

`func (o *VcenterVchaPlacementInfo) SetHaNetworkType(v VcenterVchaNetworkType)`

SetHaNetworkType sets HaNetworkType field to given value.

### HasHaNetworkType

`func (o *VcenterVchaPlacementInfo) HasHaNetworkType() bool`

HasHaNetworkType returns a boolean if a field has been set.

### GetManagementNetwork

`func (o *VcenterVchaPlacementInfo) GetManagementNetwork() string`

GetManagementNetwork returns the ManagementNetwork field if non-nil, zero value otherwise.

### GetManagementNetworkOk

`func (o *VcenterVchaPlacementInfo) GetManagementNetworkOk() (*string, bool)`

GetManagementNetworkOk returns a tuple with the ManagementNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNetwork

`func (o *VcenterVchaPlacementInfo) SetManagementNetwork(v string)`

SetManagementNetwork sets ManagementNetwork field to given value.


### GetManagementNetworkName

`func (o *VcenterVchaPlacementInfo) GetManagementNetworkName() string`

GetManagementNetworkName returns the ManagementNetworkName field if non-nil, zero value otherwise.

### GetManagementNetworkNameOk

`func (o *VcenterVchaPlacementInfo) GetManagementNetworkNameOk() (*string, bool)`

GetManagementNetworkNameOk returns a tuple with the ManagementNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNetworkName

`func (o *VcenterVchaPlacementInfo) SetManagementNetworkName(v string)`

SetManagementNetworkName sets ManagementNetworkName field to given value.


### GetManagementNetworkType

`func (o *VcenterVchaPlacementInfo) GetManagementNetworkType() VcenterVchaNetworkType`

GetManagementNetworkType returns the ManagementNetworkType field if non-nil, zero value otherwise.

### GetManagementNetworkTypeOk

`func (o *VcenterVchaPlacementInfo) GetManagementNetworkTypeOk() (*VcenterVchaNetworkType, bool)`

GetManagementNetworkTypeOk returns a tuple with the ManagementNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNetworkType

`func (o *VcenterVchaPlacementInfo) SetManagementNetworkType(v VcenterVchaNetworkType)`

SetManagementNetworkType sets ManagementNetworkType field to given value.


### GetStorage

`func (o *VcenterVchaPlacementInfo) GetStorage() VcenterVchaDiskInfo`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *VcenterVchaPlacementInfo) GetStorageOk() (*VcenterVchaDiskInfo, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *VcenterVchaPlacementInfo) SetStorage(v VcenterVchaDiskInfo)`

SetStorage sets Storage field to given value.


### GetBiosUuid

`func (o *VcenterVchaPlacementInfo) GetBiosUuid() string`

GetBiosUuid returns the BiosUuid field if non-nil, zero value otherwise.

### GetBiosUuidOk

`func (o *VcenterVchaPlacementInfo) GetBiosUuidOk() (*string, bool)`

GetBiosUuidOk returns a tuple with the BiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUuid

`func (o *VcenterVchaPlacementInfo) SetBiosUuid(v string)`

SetBiosUuid sets BiosUuid field to given value.

### HasBiosUuid

`func (o *VcenterVchaPlacementInfo) HasBiosUuid() bool`

HasBiosUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


