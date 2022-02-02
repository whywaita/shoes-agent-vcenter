# VcenterVmHardwareIdeAddressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | **bool** | Flag specifying whether the device is attached to the primary or secondary IDE adapter of the virtual machine. | 
**Master** | **bool** | Flag specifying whether the device is the master or slave device on the IDE adapter. | 

## Methods

### NewVcenterVmHardwareIdeAddressInfo

`func NewVcenterVmHardwareIdeAddressInfo(primary bool, master bool, ) *VcenterVmHardwareIdeAddressInfo`

NewVcenterVmHardwareIdeAddressInfo instantiates a new VcenterVmHardwareIdeAddressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareIdeAddressInfoWithDefaults

`func NewVcenterVmHardwareIdeAddressInfoWithDefaults() *VcenterVmHardwareIdeAddressInfo`

NewVcenterVmHardwareIdeAddressInfoWithDefaults instantiates a new VcenterVmHardwareIdeAddressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *VcenterVmHardwareIdeAddressInfo) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *VcenterVmHardwareIdeAddressInfo) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *VcenterVmHardwareIdeAddressInfo) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


### GetMaster

`func (o *VcenterVmHardwareIdeAddressInfo) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *VcenterVmHardwareIdeAddressInfo) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *VcenterVmHardwareIdeAddressInfo) SetMaster(v bool)`

SetMaster sets Master field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


