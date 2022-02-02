# VcenterGuestAdapterMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | Pointer to **string** | The MAC address of a network adapter being customized. If unset, the customization process maps the the settings from the list of AdapterMappings.IPSettings in the CustomizationSpec.interfaces to the virtual machine&#39;s network adapters, in PCI slot order. The first virtual network adapter on the PCI bus is assigned interfaces[0].IPSettings, the second adapter is assigned interfaces[1].IPSettings, and so on. | [optional] 
**Adapter** | [**VcenterGuestIPSettings**](VcenterGuestIPSettings.md) |  | 

## Methods

### NewVcenterGuestAdapterMapping

`func NewVcenterGuestAdapterMapping(adapter VcenterGuestIPSettings, ) *VcenterGuestAdapterMapping`

NewVcenterGuestAdapterMapping instantiates a new VcenterGuestAdapterMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestAdapterMappingWithDefaults

`func NewVcenterGuestAdapterMappingWithDefaults() *VcenterGuestAdapterMapping`

NewVcenterGuestAdapterMappingWithDefaults instantiates a new VcenterGuestAdapterMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *VcenterGuestAdapterMapping) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VcenterGuestAdapterMapping) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VcenterGuestAdapterMapping) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VcenterGuestAdapterMapping) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetAdapter

`func (o *VcenterGuestAdapterMapping) GetAdapter() VcenterGuestIPSettings`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *VcenterGuestAdapterMapping) GetAdapterOk() (*VcenterGuestIPSettings, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *VcenterGuestAdapterMapping) SetAdapter(v VcenterGuestIPSettings)`

SetAdapter sets Adapter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


