# VcenterVmHardwareScsiAddressSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bus** | **int64** | Bus number of the adapter to which the device should be attached. | 
**Unit** | Pointer to **int64** | Unit number of the device. If unset, the server will choose an available unit number on the specified adapter. If there are no available connections on the adapter, the request will be rejected. | [optional] 

## Methods

### NewVcenterVmHardwareScsiAddressSpec

`func NewVcenterVmHardwareScsiAddressSpec(bus int64, ) *VcenterVmHardwareScsiAddressSpec`

NewVcenterVmHardwareScsiAddressSpec instantiates a new VcenterVmHardwareScsiAddressSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareScsiAddressSpecWithDefaults

`func NewVcenterVmHardwareScsiAddressSpecWithDefaults() *VcenterVmHardwareScsiAddressSpec`

NewVcenterVmHardwareScsiAddressSpecWithDefaults instantiates a new VcenterVmHardwareScsiAddressSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBus

`func (o *VcenterVmHardwareScsiAddressSpec) GetBus() int64`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *VcenterVmHardwareScsiAddressSpec) GetBusOk() (*int64, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *VcenterVmHardwareScsiAddressSpec) SetBus(v int64)`

SetBus sets Bus field to given value.


### GetUnit

`func (o *VcenterVmHardwareScsiAddressSpec) GetUnit() int64`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *VcenterVmHardwareScsiAddressSpec) GetUnitOk() (*int64, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *VcenterVmHardwareScsiAddressSpec) SetUnit(v int64)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *VcenterVmHardwareScsiAddressSpec) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


