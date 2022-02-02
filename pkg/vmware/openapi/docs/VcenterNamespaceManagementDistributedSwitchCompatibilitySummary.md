# VcenterNamespaceManagementDistributedSwitchCompatibilitySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistributedSwitch** | **string** | Identifier of the switch. If networkProvider is either unset or is set to NSXT_CONTAINER_PLUGIN, the value of this field will refer to the UUID of a vim.DistributedVirtualSwitch. Otherwise, the value of the field will refer to the ID of a vim.DistributedVirtualSwitch. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vSphereDistributedSwitch. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vSphereDistributedSwitch. | 
**Compatible** | **bool** | Compatibility of this switch with vSphere Namespaces. | 
**IncompatibilityReasons** | Pointer to [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | List of reasons for incompatibility. If unset, this Distributed Switch is compatible. | [optional] 
**NetworkProvider** | Pointer to [**VcenterNamespaceManagementClustersNetworkProvider**](VcenterNamespaceManagementClustersNetworkProvider.md) |  | [optional] 
**CompatibleNetworks** | Pointer to **[]string** | List of compatible (PortGroup) Networks under the distributed switch. This field is optional because it was added in a newer version than its parent node. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Network. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Network. | [optional] 

## Methods

### NewVcenterNamespaceManagementDistributedSwitchCompatibilitySummary

`func NewVcenterNamespaceManagementDistributedSwitchCompatibilitySummary(distributedSwitch string, compatible bool, ) *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary`

NewVcenterNamespaceManagementDistributedSwitchCompatibilitySummary instantiates a new VcenterNamespaceManagementDistributedSwitchCompatibilitySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementDistributedSwitchCompatibilitySummaryWithDefaults

`func NewVcenterNamespaceManagementDistributedSwitchCompatibilitySummaryWithDefaults() *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary`

NewVcenterNamespaceManagementDistributedSwitchCompatibilitySummaryWithDefaults instantiates a new VcenterNamespaceManagementDistributedSwitchCompatibilitySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistributedSwitch

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetDistributedSwitch() string`

GetDistributedSwitch returns the DistributedSwitch field if non-nil, zero value otherwise.

### GetDistributedSwitchOk

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetDistributedSwitchOk() (*string, bool)`

GetDistributedSwitchOk returns a tuple with the DistributedSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedSwitch

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) SetDistributedSwitch(v string)`

SetDistributedSwitch sets DistributedSwitch field to given value.


### GetCompatible

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetCompatible() bool`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetCompatibleOk() (*bool, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) SetCompatible(v bool)`

SetCompatible sets Compatible field to given value.


### GetIncompatibilityReasons

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetIncompatibilityReasons() []VapiStdLocalizableMessage`

GetIncompatibilityReasons returns the IncompatibilityReasons field if non-nil, zero value otherwise.

### GetIncompatibilityReasonsOk

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetIncompatibilityReasonsOk() (*[]VapiStdLocalizableMessage, bool)`

GetIncompatibilityReasonsOk returns a tuple with the IncompatibilityReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibilityReasons

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) SetIncompatibilityReasons(v []VapiStdLocalizableMessage)`

SetIncompatibilityReasons sets IncompatibilityReasons field to given value.

### HasIncompatibilityReasons

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) HasIncompatibilityReasons() bool`

HasIncompatibilityReasons returns a boolean if a field has been set.

### GetNetworkProvider

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetNetworkProvider() VcenterNamespaceManagementClustersNetworkProvider`

GetNetworkProvider returns the NetworkProvider field if non-nil, zero value otherwise.

### GetNetworkProviderOk

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetNetworkProviderOk() (*VcenterNamespaceManagementClustersNetworkProvider, bool)`

GetNetworkProviderOk returns a tuple with the NetworkProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProvider

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) SetNetworkProvider(v VcenterNamespaceManagementClustersNetworkProvider)`

SetNetworkProvider sets NetworkProvider field to given value.

### HasNetworkProvider

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) HasNetworkProvider() bool`

HasNetworkProvider returns a boolean if a field has been set.

### GetCompatibleNetworks

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetCompatibleNetworks() []string`

GetCompatibleNetworks returns the CompatibleNetworks field if non-nil, zero value otherwise.

### GetCompatibleNetworksOk

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) GetCompatibleNetworksOk() (*[]string, bool)`

GetCompatibleNetworksOk returns a tuple with the CompatibleNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleNetworks

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) SetCompatibleNetworks(v []string)`

SetCompatibleNetworks sets CompatibleNetworks field to given value.

### HasCompatibleNetworks

`func (o *VcenterNamespaceManagementDistributedSwitchCompatibilitySummary) HasCompatibleNetworks() bool`

HasCompatibleNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


