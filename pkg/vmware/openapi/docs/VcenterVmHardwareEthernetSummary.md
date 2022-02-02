# VcenterVmHardwareEthernetSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nic** | **string** | Identifier of the virtual Ethernet adapter. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.Ethernet. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.Ethernet. | 

## Methods

### NewVcenterVmHardwareEthernetSummary

`func NewVcenterVmHardwareEthernetSummary(nic string, ) *VcenterVmHardwareEthernetSummary`

NewVcenterVmHardwareEthernetSummary instantiates a new VcenterVmHardwareEthernetSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareEthernetSummaryWithDefaults

`func NewVcenterVmHardwareEthernetSummaryWithDefaults() *VcenterVmHardwareEthernetSummary`

NewVcenterVmHardwareEthernetSummaryWithDefaults instantiates a new VcenterVmHardwareEthernetSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNic

`func (o *VcenterVmHardwareEthernetSummary) GetNic() string`

GetNic returns the Nic field if non-nil, zero value otherwise.

### GetNicOk

`func (o *VcenterVmHardwareEthernetSummary) GetNicOk() (*string, bool)`

GetNicOk returns a tuple with the Nic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNic

`func (o *VcenterVmHardwareEthernetSummary) SetNic(v string)`

SetNic sets Nic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


