# VcenterVMInfoSerialPorts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**VcenterVmHardwareSerialInfo**](VcenterVmHardwareSerialInfo.md) |  | [optional] 

## Methods

### NewVcenterVMInfoSerialPorts

`func NewVcenterVMInfoSerialPorts() *VcenterVMInfoSerialPorts`

NewVcenterVMInfoSerialPorts instantiates a new VcenterVMInfoSerialPorts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMInfoSerialPortsWithDefaults

`func NewVcenterVMInfoSerialPortsWithDefaults() *VcenterVMInfoSerialPorts`

NewVcenterVMInfoSerialPortsWithDefaults instantiates a new VcenterVMInfoSerialPorts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VcenterVMInfoSerialPorts) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VcenterVMInfoSerialPorts) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VcenterVMInfoSerialPorts) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *VcenterVMInfoSerialPorts) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *VcenterVMInfoSerialPorts) GetValue() VcenterVmHardwareSerialInfo`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VcenterVMInfoSerialPorts) GetValueOk() (*VcenterVmHardwareSerialInfo, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VcenterVMInfoSerialPorts) SetValue(v VcenterVmHardwareSerialInfo)`

SetValue sets Value field to given value.

### HasValue

`func (o *VcenterVMInfoSerialPorts) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


