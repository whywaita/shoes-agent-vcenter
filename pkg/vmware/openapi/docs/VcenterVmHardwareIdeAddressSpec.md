# VcenterVmHardwareIdeAddressSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | Pointer to **bool** | Flag specifying whether the device should be attached to the primary or secondary IDE adapter of the virtual machine. If unset, the server will choose a adapter with an available connection. If no IDE connections are available, the request will be rejected. | [optional] 
**Master** | Pointer to **bool** | Flag specifying whether the device should be the master or slave device on the IDE adapter. If unset, the server will choose an available connection type. If no IDE connections are available, the request will be rejected. | [optional] 

## Methods

### NewVcenterVmHardwareIdeAddressSpec

`func NewVcenterVmHardwareIdeAddressSpec() *VcenterVmHardwareIdeAddressSpec`

NewVcenterVmHardwareIdeAddressSpec instantiates a new VcenterVmHardwareIdeAddressSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareIdeAddressSpecWithDefaults

`func NewVcenterVmHardwareIdeAddressSpecWithDefaults() *VcenterVmHardwareIdeAddressSpec`

NewVcenterVmHardwareIdeAddressSpecWithDefaults instantiates a new VcenterVmHardwareIdeAddressSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *VcenterVmHardwareIdeAddressSpec) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *VcenterVmHardwareIdeAddressSpec) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *VcenterVmHardwareIdeAddressSpec) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *VcenterVmHardwareIdeAddressSpec) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetMaster

`func (o *VcenterVmHardwareIdeAddressSpec) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *VcenterVmHardwareIdeAddressSpec) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *VcenterVmHardwareIdeAddressSpec) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *VcenterVmHardwareIdeAddressSpec) HasMaster() bool`

HasMaster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


