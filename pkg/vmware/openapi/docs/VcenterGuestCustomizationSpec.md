# VcenterGuestCustomizationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationSpec** | [**VcenterGuestConfigurationSpec**](VcenterGuestConfigurationSpec.md) |  | 
**GlobalDNSSettings** | [**VcenterGuestGlobalDNSSettings**](VcenterGuestGlobalDNSSettings.md) |  | 
**Interfaces** | [**[]VcenterGuestAdapterMapping**](VcenterGuestAdapterMapping.md) | IP settings that are specific to a particular virtual network adapter. The AdapterMapping structure maps a network adapter&#39;s MAC address to its IPSettings. May be empty if there are no network adapters, else should match number of network adapters configured for the VM. | 

## Methods

### NewVcenterGuestCustomizationSpec

`func NewVcenterGuestCustomizationSpec(configurationSpec VcenterGuestConfigurationSpec, globalDNSSettings VcenterGuestGlobalDNSSettings, interfaces []VcenterGuestAdapterMapping, ) *VcenterGuestCustomizationSpec`

NewVcenterGuestCustomizationSpec instantiates a new VcenterGuestCustomizationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestCustomizationSpecWithDefaults

`func NewVcenterGuestCustomizationSpecWithDefaults() *VcenterGuestCustomizationSpec`

NewVcenterGuestCustomizationSpecWithDefaults instantiates a new VcenterGuestCustomizationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationSpec

`func (o *VcenterGuestCustomizationSpec) GetConfigurationSpec() VcenterGuestConfigurationSpec`

GetConfigurationSpec returns the ConfigurationSpec field if non-nil, zero value otherwise.

### GetConfigurationSpecOk

`func (o *VcenterGuestCustomizationSpec) GetConfigurationSpecOk() (*VcenterGuestConfigurationSpec, bool)`

GetConfigurationSpecOk returns a tuple with the ConfigurationSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationSpec

`func (o *VcenterGuestCustomizationSpec) SetConfigurationSpec(v VcenterGuestConfigurationSpec)`

SetConfigurationSpec sets ConfigurationSpec field to given value.


### GetGlobalDNSSettings

`func (o *VcenterGuestCustomizationSpec) GetGlobalDNSSettings() VcenterGuestGlobalDNSSettings`

GetGlobalDNSSettings returns the GlobalDNSSettings field if non-nil, zero value otherwise.

### GetGlobalDNSSettingsOk

`func (o *VcenterGuestCustomizationSpec) GetGlobalDNSSettingsOk() (*VcenterGuestGlobalDNSSettings, bool)`

GetGlobalDNSSettingsOk returns a tuple with the GlobalDNSSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDNSSettings

`func (o *VcenterGuestCustomizationSpec) SetGlobalDNSSettings(v VcenterGuestGlobalDNSSettings)`

SetGlobalDNSSettings sets GlobalDNSSettings field to given value.


### GetInterfaces

`func (o *VcenterGuestCustomizationSpec) GetInterfaces() []VcenterGuestAdapterMapping`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *VcenterGuestCustomizationSpec) GetInterfacesOk() (*[]VcenterGuestAdapterMapping, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *VcenterGuestCustomizationSpec) SetInterfaces(v []VcenterGuestAdapterMapping)`

SetInterfaces sets Interfaces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


