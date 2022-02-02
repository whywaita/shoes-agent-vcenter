# VcenterVmTemplateLibraryItemsEthernetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackingType** | [**VcenterVmTemplateLibraryItemsEthernetInfoNetworkBackingType**](VcenterVmTemplateLibraryItemsEthernetInfoNetworkBackingType.md) |  | 
**MacType** | [**VcenterVmTemplateLibraryItemsEthernetInfoMacAddressType**](VcenterVmTemplateLibraryItemsEthernetInfoMacAddressType.md) |  | 
**Network** | Pointer to **string** | Identifier of the network backing the virtual Ethernet adapter. | [optional] 

## Methods

### NewVcenterVmTemplateLibraryItemsEthernetInfo

`func NewVcenterVmTemplateLibraryItemsEthernetInfo(backingType VcenterVmTemplateLibraryItemsEthernetInfoNetworkBackingType, macType VcenterVmTemplateLibraryItemsEthernetInfoMacAddressType, ) *VcenterVmTemplateLibraryItemsEthernetInfo`

NewVcenterVmTemplateLibraryItemsEthernetInfo instantiates a new VcenterVmTemplateLibraryItemsEthernetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsEthernetInfoWithDefaults

`func NewVcenterVmTemplateLibraryItemsEthernetInfoWithDefaults() *VcenterVmTemplateLibraryItemsEthernetInfo`

NewVcenterVmTemplateLibraryItemsEthernetInfoWithDefaults instantiates a new VcenterVmTemplateLibraryItemsEthernetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackingType

`func (o *VcenterVmTemplateLibraryItemsEthernetInfo) GetBackingType() VcenterVmTemplateLibraryItemsEthernetInfoNetworkBackingType`

GetBackingType returns the BackingType field if non-nil, zero value otherwise.

### GetBackingTypeOk

`func (o *VcenterVmTemplateLibraryItemsEthernetInfo) GetBackingTypeOk() (*VcenterVmTemplateLibraryItemsEthernetInfoNetworkBackingType, bool)`

GetBackingTypeOk returns a tuple with the BackingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingType

`func (o *VcenterVmTemplateLibraryItemsEthernetInfo) SetBackingType(v VcenterVmTemplateLibraryItemsEthernetInfoNetworkBackingType)`

SetBackingType sets BackingType field to given value.


### GetMacType

`func (o *VcenterVmTemplateLibraryItemsEthernetInfo) GetMacType() VcenterVmTemplateLibraryItemsEthernetInfoMacAddressType`

GetMacType returns the MacType field if non-nil, zero value otherwise.

### GetMacTypeOk

`func (o *VcenterVmTemplateLibraryItemsEthernetInfo) GetMacTypeOk() (*VcenterVmTemplateLibraryItemsEthernetInfoMacAddressType, bool)`

GetMacTypeOk returns a tuple with the MacType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacType

`func (o *VcenterVmTemplateLibraryItemsEthernetInfo) SetMacType(v VcenterVmTemplateLibraryItemsEthernetInfoMacAddressType)`

SetMacType sets MacType field to given value.


### GetNetwork

`func (o *VcenterVmTemplateLibraryItemsEthernetInfo) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VcenterVmTemplateLibraryItemsEthernetInfo) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VcenterVmTemplateLibraryItemsEthernetInfo) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VcenterVmTemplateLibraryItemsEthernetInfo) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


