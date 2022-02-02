# VcenterVmHardwareNvmeAddressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bus** | **int64** | Bus number of the adapter to which the device is attached. | 
**Unit** | **int64** | Unit number of the device. | 

## Methods

### NewVcenterVmHardwareNvmeAddressInfo

`func NewVcenterVmHardwareNvmeAddressInfo(bus int64, unit int64, ) *VcenterVmHardwareNvmeAddressInfo`

NewVcenterVmHardwareNvmeAddressInfo instantiates a new VcenterVmHardwareNvmeAddressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareNvmeAddressInfoWithDefaults

`func NewVcenterVmHardwareNvmeAddressInfoWithDefaults() *VcenterVmHardwareNvmeAddressInfo`

NewVcenterVmHardwareNvmeAddressInfoWithDefaults instantiates a new VcenterVmHardwareNvmeAddressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBus

`func (o *VcenterVmHardwareNvmeAddressInfo) GetBus() int64`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *VcenterVmHardwareNvmeAddressInfo) GetBusOk() (*int64, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *VcenterVmHardwareNvmeAddressInfo) SetBus(v int64)`

SetBus sets Bus field to given value.


### GetUnit

`func (o *VcenterVmHardwareNvmeAddressInfo) GetUnit() int64`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *VcenterVmHardwareNvmeAddressInfo) GetUnitOk() (*int64, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *VcenterVmHardwareNvmeAddressInfo) SetUnit(v int64)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


