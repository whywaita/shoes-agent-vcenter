# VcenterVMInfoDisks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**VcenterVmHardwareDiskInfo**](VcenterVmHardwareDiskInfo.md) |  | [optional] 

## Methods

### NewVcenterVMInfoDisks

`func NewVcenterVMInfoDisks() *VcenterVMInfoDisks`

NewVcenterVMInfoDisks instantiates a new VcenterVMInfoDisks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMInfoDisksWithDefaults

`func NewVcenterVMInfoDisksWithDefaults() *VcenterVMInfoDisks`

NewVcenterVMInfoDisksWithDefaults instantiates a new VcenterVMInfoDisks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VcenterVMInfoDisks) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VcenterVMInfoDisks) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VcenterVMInfoDisks) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *VcenterVMInfoDisks) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *VcenterVMInfoDisks) GetValue() VcenterVmHardwareDiskInfo`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VcenterVMInfoDisks) GetValueOk() (*VcenterVmHardwareDiskInfo, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VcenterVMInfoDisks) SetValue(v VcenterVmHardwareDiskInfo)`

SetValue sets Value field to given value.

### HasValue

`func (o *VcenterVMInfoDisks) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


